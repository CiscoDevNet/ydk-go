// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-sc-diag package operational data.
// 
// This module contains definitions
// for the following management objects:
//   diag: Diag admin operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_sc_diag_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_sc_diag_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-sc-diag-oper diag}", reflect.TypeOf(Diag{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-sc-diag-oper:diag", reflect.TypeOf(Diag{}))
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diag operational data for available racks.
    Racks Diag_Racks
}

func (diag *Diag) GetEntityData() *types.CommonEntityData {
    diag.EntityData.YFilter = diag.YFilter
    diag.EntityData.YangName = "diag"
    diag.EntityData.BundleName = "cisco_ios_xr"
    diag.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-sc-diag-oper"
    diag.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag"
    diag.EntityData.AbsolutePath = diag.EntityData.SegmentPath
    diag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diag.EntityData.Children = types.NewOrderedMap()
    diag.EntityData.Children.Append("racks", types.YChild{"Racks", &diag.Racks})
    diag.EntityData.Leafs = types.NewOrderedMap()

    diag.EntityData.YListKeys = []string {}

    return &(diag.EntityData)
}

// Diag_Racks
// Diag operational data for available racks
type Diag_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diag operational data for a particular rack. The type is slice of
    // Diag_Racks_Rack.
    Rack []*Diag_Racks_Rack
}

func (racks *Diag_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "diag"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/" + racks.EntityData.SegmentPath
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

// Diag_Racks_Rack
// Diag operational data for a particular rack
type Diag_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Rack name. The type is interface{} with range:
    // 0..4294967295.
    RackName interface{}

    // Diag operational data for all available slots.
    Slots Diag_Racks_Rack_Slots

    // Summary information.
    Summary Diag_Racks_Rack_Summary
}

func (rack *Diag_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.RackName, "rack-name")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Children.Append("summary", types.YChild{"Summary", &rack.Summary})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("rack-name", types.YLeaf{"RackName", rack.RackName})

    rack.EntityData.YListKeys = []string {"RackName"}

    return &(rack.EntityData)
}

// Diag_Racks_Rack_Slots
// Diag operational data for all available slots
type Diag_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diag operational data for a particular slot. The type is slice of
    // Diag_Racks_Rack_Slots_Slot.
    Slot []*Diag_Racks_Rack_Slots_Slot
}

func (slots *Diag_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/" + slots.EntityData.SegmentPath
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

// Diag_Racks_Rack_Slots_Slot
// Diag operational data for a particular slot
type Diag_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    SlotName interface{}

    // Slot detailed information.
    Detail Diag_Racks_Rack_Slots_Slot_Detail

    // Diag operational data for all available instances.
    Instances Diag_Racks_Rack_Slots_Slot_Instances
}

func (slot *Diag_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.SlotName, "slot-name")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("detail", types.YChild{"Detail", &slot.Detail})
    slot.EntityData.Children.Append("instances", types.YChild{"Instances", &slot.Instances})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("slot-name", types.YLeaf{"SlotName", slot.SlotName})

    slot.EntityData.YListKeys = []string {"SlotName"}

    return &(slot.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail
// Slot detailed information
type Diag_Racks_Rack_Slots_Slot_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail information for slot. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail.
    NodeDetail []*Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail

    // Detail information for spa. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail.
    SpaDetail []*Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail
}

func (detail *Diag_Racks_Rack_Slots_Slot_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "slot"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("node-detail", types.YChild{"NodeDetail", nil})
    for i := range detail.NodeDetail {
        types.SetYListKey(detail.NodeDetail[i], i)
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.NodeDetail[i]), types.YChild{"NodeDetail", detail.NodeDetail[i]})
    }
    detail.EntityData.Children.Append("spa-detail", types.YChild{"SpaDetail", nil})
    for i := range detail.SpaDetail {
        types.SetYListKey(detail.SpaDetail[i], i)
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.SpaDetail[i]), types.YChild{"SpaDetail", detail.SpaDetail[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail
// Detail information for slot
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (nodeDetail *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail) GetEntityData() *types.CommonEntityData {
    nodeDetail.EntityData.YFilter = nodeDetail.YFilter
    nodeDetail.EntityData.YangName = "node-detail"
    nodeDetail.EntityData.BundleName = "cisco_ios_xr"
    nodeDetail.EntityData.ParentYangName = "detail"
    nodeDetail.EntityData.SegmentPath = "node-detail" + types.AddNoKeyToken(nodeDetail)
    nodeDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/" + nodeDetail.EntityData.SegmentPath
    nodeDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeDetail.EntityData.Children = types.NewOrderedMap()
    nodeDetail.EntityData.Children.Append("pld", types.YChild{"Pld", &nodeDetail.Pld})
    nodeDetail.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", &nodeDetail.HardwareRevision})
    nodeDetail.EntityData.Children.Append("cbc-active-partition", types.YChild{"CbcActivePartition", &nodeDetail.CbcActivePartition})
    nodeDetail.EntityData.Children.Append("cbc-inactive-partition", types.YChild{"CbcInactivePartition", &nodeDetail.CbcInactivePartition})
    nodeDetail.EntityData.Leafs = types.NewOrderedMap()
    nodeDetail.EntityData.Leafs.Append("description", types.YLeaf{"Description", nodeDetail.Description})
    nodeDetail.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", nodeDetail.SerialNumber})
    nodeDetail.EntityData.Leafs.Append("tan", types.YLeaf{"Tan", nodeDetail.Tan})
    nodeDetail.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", nodeDetail.Pid})
    nodeDetail.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", nodeDetail.Vid})
    nodeDetail.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", nodeDetail.ChipHardwareRevision})
    nodeDetail.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", nodeDetail.NewDeviationNumber})
    nodeDetail.EntityData.Leafs.Append("clei", types.YLeaf{"Clei", nodeDetail.Clei})
    nodeDetail.EntityData.Leafs.Append("board-state", types.YLeaf{"BoardState", nodeDetail.BoardState})
    nodeDetail.EntityData.Leafs.Append("pld-motherboard", types.YLeaf{"PldMotherboard", nodeDetail.PldMotherboard})
    nodeDetail.EntityData.Leafs.Append("pld-power", types.YLeaf{"PldPower", nodeDetail.PldPower})
    nodeDetail.EntityData.Leafs.Append("monlib", types.YLeaf{"Monlib", nodeDetail.Monlib})
    nodeDetail.EntityData.Leafs.Append("rommon", types.YLeaf{"Rommon", nodeDetail.Rommon})
    nodeDetail.EntityData.Leafs.Append("cpu0", types.YLeaf{"Cpu0", nodeDetail.Cpu0})

    nodeDetail.EntityData.YListKeys = []string {}

    return &(nodeDetail.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld
// Programmable logic device information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_Pld) GetEntityData() *types.CommonEntityData {
    pld.EntityData.YFilter = pld.YFilter
    pld.EntityData.YangName = "pld"
    pld.EntityData.BundleName = "cisco_ios_xr"
    pld.EntityData.ParentYangName = "node-detail"
    pld.EntityData.SegmentPath = "pld"
    pld.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/" + pld.EntityData.SegmentPath
    pld.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pld.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pld.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pld.EntityData.Children = types.NewOrderedMap()
    pld.EntityData.Leafs = types.NewOrderedMap()
    pld.EntityData.Leafs.Append("type", types.YLeaf{"Type", pld.Type})
    pld.EntityData.Leafs.Append("processor-higher-version", types.YLeaf{"ProcessorHigherVersion", pld.ProcessorHigherVersion})
    pld.EntityData.Leafs.Append("processor-lower-version", types.YLeaf{"ProcessorLowerVersion", pld.ProcessorLowerVersion})

    pld.EntityData.YListKeys = []string {}

    return &(pld.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision.
    HardwareRevision []*Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "node-detail"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision"
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", nil})
    for i := range hardwareRevision.HardwareRevision {
        types.SetYListKey(hardwareRevision.HardwareRevision[i], i)
        hardwareRevision.EntityData.Children.Append(types.GetSegmentPath(hardwareRevision.HardwareRevision[i]), types.YChild{"HardwareRevision", hardwareRevision.HardwareRevision[i]})
    }
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "hardware-revision"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision" + types.AddNoKeyToken(hardwareRevision)
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hw-rev", types.YChild{"HwRev", &hardwareRevision.HwRev})
    hardwareRevision.EntityData.Children.Append("fw-rev", types.YChild{"FwRev", &hardwareRevision.FwRev})
    hardwareRevision.EntityData.Children.Append("sw-rev", types.YChild{"SwRev", &hardwareRevision.SwRev})
    hardwareRevision.EntityData.Children.Append("dimm-rev", types.YChild{"DimmRev", &hardwareRevision.DimmRev})
    hardwareRevision.EntityData.Children.Append("ssd-rev", types.YChild{"SsdRev", &hardwareRevision.SsdRev})
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()
    hardwareRevision.EntityData.Leafs.Append("node-description", types.YLeaf{"NodeDescription", hardwareRevision.NodeDescription})
    hardwareRevision.EntityData.Leafs.Append("version", types.YLeaf{"Version", hardwareRevision.Version})

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_HwRev) GetEntityData() *types.CommonEntityData {
    hwRev.EntityData.YFilter = hwRev.YFilter
    hwRev.EntityData.YangName = "hw-rev"
    hwRev.EntityData.BundleName = "cisco_ios_xr"
    hwRev.EntityData.ParentYangName = "hardware-revision"
    hwRev.EntityData.SegmentPath = "hw-rev"
    hwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/hardware-revision/" + hwRev.EntityData.SegmentPath
    hwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwRev.EntityData.Children = types.NewOrderedMap()
    hwRev.EntityData.Leafs = types.NewOrderedMap()
    hwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", hwRev.MajorRevision})
    hwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", hwRev.MinorRevision})

    hwRev.EntityData.YListKeys = []string {}

    return &(hwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_FwRev) GetEntityData() *types.CommonEntityData {
    fwRev.EntityData.YFilter = fwRev.YFilter
    fwRev.EntityData.YangName = "fw-rev"
    fwRev.EntityData.BundleName = "cisco_ios_xr"
    fwRev.EntityData.ParentYangName = "hardware-revision"
    fwRev.EntityData.SegmentPath = "fw-rev"
    fwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/hardware-revision/" + fwRev.EntityData.SegmentPath
    fwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwRev.EntityData.Children = types.NewOrderedMap()
    fwRev.EntityData.Leafs = types.NewOrderedMap()
    fwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", fwRev.MajorRevision})
    fwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", fwRev.MinorRevision})

    fwRev.EntityData.YListKeys = []string {}

    return &(fwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SwRev) GetEntityData() *types.CommonEntityData {
    swRev.EntityData.YFilter = swRev.YFilter
    swRev.EntityData.YangName = "sw-rev"
    swRev.EntityData.BundleName = "cisco_ios_xr"
    swRev.EntityData.ParentYangName = "hardware-revision"
    swRev.EntityData.SegmentPath = "sw-rev"
    swRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/hardware-revision/" + swRev.EntityData.SegmentPath
    swRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swRev.EntityData.Children = types.NewOrderedMap()
    swRev.EntityData.Leafs = types.NewOrderedMap()
    swRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", swRev.MajorRevision})
    swRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", swRev.MinorRevision})

    swRev.EntityData.YListKeys = []string {}

    return &(swRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev struct {
    EntityData types.CommonEntityData
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

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_DimmRev) GetEntityData() *types.CommonEntityData {
    dimmRev.EntityData.YFilter = dimmRev.YFilter
    dimmRev.EntityData.YangName = "dimm-rev"
    dimmRev.EntityData.BundleName = "cisco_ios_xr"
    dimmRev.EntityData.ParentYangName = "hardware-revision"
    dimmRev.EntityData.SegmentPath = "dimm-rev"
    dimmRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/hardware-revision/" + dimmRev.EntityData.SegmentPath
    dimmRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dimmRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dimmRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dimmRev.EntityData.Children = types.NewOrderedMap()
    dimmRev.EntityData.Leafs = types.NewOrderedMap()
    dimmRev.EntityData.Leafs.Append("size", types.YLeaf{"Size", dimmRev.Size})
    dimmRev.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", dimmRev.Speed})
    dimmRev.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", dimmRev.Locator})
    dimmRev.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", dimmRev.Cas})

    dimmRev.EntityData.YListKeys = []string {}

    return &(dimmRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_HardwareRevision_HardwareRevision_SsdRev) GetEntityData() *types.CommonEntityData {
    ssdRev.EntityData.YFilter = ssdRev.YFilter
    ssdRev.EntityData.YangName = "ssd-rev"
    ssdRev.EntityData.BundleName = "cisco_ios_xr"
    ssdRev.EntityData.ParentYangName = "hardware-revision"
    ssdRev.EntityData.SegmentPath = "ssd-rev"
    ssdRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/hardware-revision/hardware-revision/" + ssdRev.EntityData.SegmentPath
    ssdRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssdRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssdRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssdRev.EntityData.Children = types.NewOrderedMap()
    ssdRev.EntityData.Leafs = types.NewOrderedMap()
    ssdRev.EntityData.Leafs.Append("number", types.YLeaf{"Number", ssdRev.Number})
    ssdRev.EntityData.Leafs.Append("fw-rev", types.YLeaf{"FwRev", ssdRev.FwRev})
    ssdRev.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", ssdRev.SerialNumber})

    ssdRev.EntityData.YListKeys = []string {}

    return &(ssdRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcActivePartition) GetEntityData() *types.CommonEntityData {
    cbcActivePartition.EntityData.YFilter = cbcActivePartition.YFilter
    cbcActivePartition.EntityData.YangName = "cbc-active-partition"
    cbcActivePartition.EntityData.BundleName = "cisco_ios_xr"
    cbcActivePartition.EntityData.ParentYangName = "node-detail"
    cbcActivePartition.EntityData.SegmentPath = "cbc-active-partition"
    cbcActivePartition.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/" + cbcActivePartition.EntityData.SegmentPath
    cbcActivePartition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbcActivePartition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbcActivePartition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbcActivePartition.EntityData.Children = types.NewOrderedMap()
    cbcActivePartition.EntityData.Leafs = types.NewOrderedMap()
    cbcActivePartition.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", cbcActivePartition.MajorRevision})
    cbcActivePartition.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", cbcActivePartition.MinorRevision})

    cbcActivePartition.EntityData.YListKeys = []string {}

    return &(cbcActivePartition.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Detail_NodeDetail_CbcInactivePartition) GetEntityData() *types.CommonEntityData {
    cbcInactivePartition.EntityData.YFilter = cbcInactivePartition.YFilter
    cbcInactivePartition.EntityData.YangName = "cbc-inactive-partition"
    cbcInactivePartition.EntityData.BundleName = "cisco_ios_xr"
    cbcInactivePartition.EntityData.ParentYangName = "node-detail"
    cbcInactivePartition.EntityData.SegmentPath = "cbc-inactive-partition"
    cbcInactivePartition.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/node-detail/" + cbcInactivePartition.EntityData.SegmentPath
    cbcInactivePartition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbcInactivePartition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbcInactivePartition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbcInactivePartition.EntityData.Children = types.NewOrderedMap()
    cbcInactivePartition.EntityData.Leafs = types.NewOrderedMap()
    cbcInactivePartition.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", cbcInactivePartition.MajorRevision})
    cbcInactivePartition.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", cbcInactivePartition.MinorRevision})

    cbcInactivePartition.EntityData.YListKeys = []string {}

    return &(cbcInactivePartition.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail
// Detail information for spa
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (spaDetail *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail) GetEntityData() *types.CommonEntityData {
    spaDetail.EntityData.YFilter = spaDetail.YFilter
    spaDetail.EntityData.YangName = "spa-detail"
    spaDetail.EntityData.BundleName = "cisco_ios_xr"
    spaDetail.EntityData.ParentYangName = "detail"
    spaDetail.EntityData.SegmentPath = "spa-detail" + types.AddNoKeyToken(spaDetail)
    spaDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/" + spaDetail.EntityData.SegmentPath
    spaDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaDetail.EntityData.Children = types.NewOrderedMap()
    spaDetail.EntityData.Children.Append("main", types.YChild{"Main", &spaDetail.Main})
    spaDetail.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", &spaDetail.HardwareRevision})
    spaDetail.EntityData.Leafs = types.NewOrderedMap()
    spaDetail.EntityData.Leafs.Append("node", types.YLeaf{"Node", spaDetail.Node})
    spaDetail.EntityData.Leafs.Append("name", types.YLeaf{"Name", spaDetail.Name})
    spaDetail.EntityData.Leafs.Append("pca-unit-number", types.YLeaf{"PcaUnitNumber", spaDetail.PcaUnitNumber})
    spaDetail.EntityData.Leafs.Append("pca-revision", types.YLeaf{"PcaRevision", spaDetail.PcaRevision})
    spaDetail.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", spaDetail.Pid})
    spaDetail.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", spaDetail.Vid})
    spaDetail.EntityData.Leafs.Append("clei", types.YLeaf{"Clei", spaDetail.Clei})
    spaDetail.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", spaDetail.NodeState})

    spaDetail.EntityData.YListKeys = []string {}

    return &(spaDetail.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main
// Main
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main struct {
    EntityData types.CommonEntityData
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

func (main *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_Main) GetEntityData() *types.CommonEntityData {
    main.EntityData.YFilter = main.YFilter
    main.EntityData.YangName = "main"
    main.EntityData.BundleName = "cisco_ios_xr"
    main.EntityData.ParentYangName = "spa-detail"
    main.EntityData.SegmentPath = "main"
    main.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/" + main.EntityData.SegmentPath
    main.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    main.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    main.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    main.EntityData.Children = types.NewOrderedMap()
    main.EntityData.Leafs = types.NewOrderedMap()
    main.EntityData.Leafs.Append("board-type", types.YLeaf{"BoardType", main.BoardType})
    main.EntityData.Leafs.Append("tan", types.YLeaf{"Tan", main.Tan})
    main.EntityData.Leafs.Append("tan-revision", types.YLeaf{"TanRevision", main.TanRevision})
    main.EntityData.Leafs.Append("deviation-number", types.YLeaf{"DeviationNumber", main.DeviationNumber})
    main.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", main.SerialNumber})

    main.EntityData.YListKeys = []string {}

    return &(main.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision.
    HardwareRevision []*Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "spa-detail"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision"
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", nil})
    for i := range hardwareRevision.HardwareRevision {
        types.SetYListKey(hardwareRevision.HardwareRevision[i], i)
        hardwareRevision.EntityData.Children.Append(types.GetSegmentPath(hardwareRevision.HardwareRevision[i]), types.YChild{"HardwareRevision", hardwareRevision.HardwareRevision[i]})
    }
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "hardware-revision"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision" + types.AddNoKeyToken(hardwareRevision)
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hw-rev", types.YChild{"HwRev", &hardwareRevision.HwRev})
    hardwareRevision.EntityData.Children.Append("fw-rev", types.YChild{"FwRev", &hardwareRevision.FwRev})
    hardwareRevision.EntityData.Children.Append("sw-rev", types.YChild{"SwRev", &hardwareRevision.SwRev})
    hardwareRevision.EntityData.Children.Append("dimm-rev", types.YChild{"DimmRev", &hardwareRevision.DimmRev})
    hardwareRevision.EntityData.Children.Append("ssd-rev", types.YChild{"SsdRev", &hardwareRevision.SsdRev})
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()
    hardwareRevision.EntityData.Leafs.Append("node-description", types.YLeaf{"NodeDescription", hardwareRevision.NodeDescription})
    hardwareRevision.EntityData.Leafs.Append("version", types.YLeaf{"Version", hardwareRevision.Version})

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_HwRev) GetEntityData() *types.CommonEntityData {
    hwRev.EntityData.YFilter = hwRev.YFilter
    hwRev.EntityData.YangName = "hw-rev"
    hwRev.EntityData.BundleName = "cisco_ios_xr"
    hwRev.EntityData.ParentYangName = "hardware-revision"
    hwRev.EntityData.SegmentPath = "hw-rev"
    hwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/hardware-revision/" + hwRev.EntityData.SegmentPath
    hwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwRev.EntityData.Children = types.NewOrderedMap()
    hwRev.EntityData.Leafs = types.NewOrderedMap()
    hwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", hwRev.MajorRevision})
    hwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", hwRev.MinorRevision})

    hwRev.EntityData.YListKeys = []string {}

    return &(hwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_FwRev) GetEntityData() *types.CommonEntityData {
    fwRev.EntityData.YFilter = fwRev.YFilter
    fwRev.EntityData.YangName = "fw-rev"
    fwRev.EntityData.BundleName = "cisco_ios_xr"
    fwRev.EntityData.ParentYangName = "hardware-revision"
    fwRev.EntityData.SegmentPath = "fw-rev"
    fwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/hardware-revision/" + fwRev.EntityData.SegmentPath
    fwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwRev.EntityData.Children = types.NewOrderedMap()
    fwRev.EntityData.Leafs = types.NewOrderedMap()
    fwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", fwRev.MajorRevision})
    fwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", fwRev.MinorRevision})

    fwRev.EntityData.YListKeys = []string {}

    return &(fwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SwRev) GetEntityData() *types.CommonEntityData {
    swRev.EntityData.YFilter = swRev.YFilter
    swRev.EntityData.YangName = "sw-rev"
    swRev.EntityData.BundleName = "cisco_ios_xr"
    swRev.EntityData.ParentYangName = "hardware-revision"
    swRev.EntityData.SegmentPath = "sw-rev"
    swRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/hardware-revision/" + swRev.EntityData.SegmentPath
    swRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swRev.EntityData.Children = types.NewOrderedMap()
    swRev.EntityData.Leafs = types.NewOrderedMap()
    swRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", swRev.MajorRevision})
    swRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", swRev.MinorRevision})

    swRev.EntityData.YListKeys = []string {}

    return &(swRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev struct {
    EntityData types.CommonEntityData
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

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_DimmRev) GetEntityData() *types.CommonEntityData {
    dimmRev.EntityData.YFilter = dimmRev.YFilter
    dimmRev.EntityData.YangName = "dimm-rev"
    dimmRev.EntityData.BundleName = "cisco_ios_xr"
    dimmRev.EntityData.ParentYangName = "hardware-revision"
    dimmRev.EntityData.SegmentPath = "dimm-rev"
    dimmRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/hardware-revision/" + dimmRev.EntityData.SegmentPath
    dimmRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dimmRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dimmRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dimmRev.EntityData.Children = types.NewOrderedMap()
    dimmRev.EntityData.Leafs = types.NewOrderedMap()
    dimmRev.EntityData.Leafs.Append("size", types.YLeaf{"Size", dimmRev.Size})
    dimmRev.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", dimmRev.Speed})
    dimmRev.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", dimmRev.Locator})
    dimmRev.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", dimmRev.Cas})

    dimmRev.EntityData.YListKeys = []string {}

    return &(dimmRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Detail_SpaDetail_HardwareRevision_HardwareRevision_SsdRev) GetEntityData() *types.CommonEntityData {
    ssdRev.EntityData.YFilter = ssdRev.YFilter
    ssdRev.EntityData.YangName = "ssd-rev"
    ssdRev.EntityData.BundleName = "cisco_ios_xr"
    ssdRev.EntityData.ParentYangName = "hardware-revision"
    ssdRev.EntityData.SegmentPath = "ssd-rev"
    ssdRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/detail/spa-detail/hardware-revision/hardware-revision/" + ssdRev.EntityData.SegmentPath
    ssdRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssdRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssdRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssdRev.EntityData.Children = types.NewOrderedMap()
    ssdRev.EntityData.Leafs = types.NewOrderedMap()
    ssdRev.EntityData.Leafs.Append("number", types.YLeaf{"Number", ssdRev.Number})
    ssdRev.EntityData.Leafs.Append("fw-rev", types.YLeaf{"FwRev", ssdRev.FwRev})
    ssdRev.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", ssdRev.SerialNumber})

    ssdRev.EntityData.YListKeys = []string {}

    return &(ssdRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances
// Diag operational data for all available
// instances
type Diag_Racks_Rack_Slots_Slot_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diag operational data for a particular instance. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []*Diag_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Diag_Racks_Rack_Slots_Slot_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "slot"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/" + instances.EntityData.SegmentPath
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

// Diag_Racks_Rack_Slots_Slot_Instances_Instance
// Diag operational data for a particular
// instance
type Diag_Racks_Rack_Slots_Slot_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Instance name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Diag detailed information.
    Detail Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
}

func (instance *Diag_Racks_Rack_Slots_Slot_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.InstanceName, "instance-name")
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("detail", types.YChild{"Detail", &instance.Detail})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", instance.InstanceName})

    instance.EntityData.YListKeys = []string {"InstanceName"}

    return &(instance.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail
// Diag detailed information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information.
    Node Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node

    // SPA information.
    Spa Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa
}

func (detail *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "instance"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("node", types.YChild{"Node", &detail.Node})
    detail.EntityData.Children.Append("spa", types.YChild{"Spa", &detail.Spa})
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node
// Node information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node struct {
    EntityData types.CommonEntityData
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

func (node *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "detail"
    node.EntityData.SegmentPath = "node"
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("pld", types.YChild{"Pld", &node.Pld})
    node.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", &node.HardwareRevision})
    node.EntityData.Children.Append("cbc-active-partition", types.YChild{"CbcActivePartition", &node.CbcActivePartition})
    node.EntityData.Children.Append("cbc-inactive-partition", types.YChild{"CbcInactivePartition", &node.CbcInactivePartition})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("description", types.YLeaf{"Description", node.Description})
    node.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", node.SerialNumber})
    node.EntityData.Leafs.Append("tan", types.YLeaf{"Tan", node.Tan})
    node.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", node.Pid})
    node.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", node.Vid})
    node.EntityData.Leafs.Append("chip-hardware-revision", types.YLeaf{"ChipHardwareRevision", node.ChipHardwareRevision})
    node.EntityData.Leafs.Append("new-deviation-number", types.YLeaf{"NewDeviationNumber", node.NewDeviationNumber})
    node.EntityData.Leafs.Append("clei", types.YLeaf{"Clei", node.Clei})
    node.EntityData.Leafs.Append("board-state", types.YLeaf{"BoardState", node.BoardState})
    node.EntityData.Leafs.Append("pld-motherboard", types.YLeaf{"PldMotherboard", node.PldMotherboard})
    node.EntityData.Leafs.Append("pld-power", types.YLeaf{"PldPower", node.PldPower})
    node.EntityData.Leafs.Append("monlib", types.YLeaf{"Monlib", node.Monlib})
    node.EntityData.Leafs.Append("rommon", types.YLeaf{"Rommon", node.Rommon})
    node.EntityData.Leafs.Append("cpu0", types.YLeaf{"Cpu0", node.Cpu0})

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld
// Programmable logic device information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Processor PLD version. The type is interface{} with range: 0..4294967295.
    Type interface{}

    // HigherVerion. The type is interface{} with range: 0..4294967295.
    ProcessorHigherVersion interface{}

    // LowerVersion. The type is interface{} with range: 0..4294967295.
    ProcessorLowerVersion interface{}
}

func (pld *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_Pld) GetEntityData() *types.CommonEntityData {
    pld.EntityData.YFilter = pld.YFilter
    pld.EntityData.YangName = "pld"
    pld.EntityData.BundleName = "cisco_ios_xr"
    pld.EntityData.ParentYangName = "node"
    pld.EntityData.SegmentPath = "pld"
    pld.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/" + pld.EntityData.SegmentPath
    pld.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pld.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pld.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pld.EntityData.Children = types.NewOrderedMap()
    pld.EntityData.Leafs = types.NewOrderedMap()
    pld.EntityData.Leafs.Append("type", types.YLeaf{"Type", pld.Type})
    pld.EntityData.Leafs.Append("processor-higher-version", types.YLeaf{"ProcessorHigherVersion", pld.ProcessorHigherVersion})
    pld.EntityData.Leafs.Append("processor-lower-version", types.YLeaf{"ProcessorLowerVersion", pld.ProcessorLowerVersion})

    pld.EntityData.YListKeys = []string {}

    return &(pld.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision.
    HardwareRevision []*Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "node"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision"
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", nil})
    for i := range hardwareRevision.HardwareRevision {
        types.SetYListKey(hardwareRevision.HardwareRevision[i], i)
        hardwareRevision.EntityData.Children.Append(types.GetSegmentPath(hardwareRevision.HardwareRevision[i]), types.YChild{"HardwareRevision", hardwareRevision.HardwareRevision[i]})
    }
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "hardware-revision"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision" + types.AddNoKeyToken(hardwareRevision)
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hw-rev", types.YChild{"HwRev", &hardwareRevision.HwRev})
    hardwareRevision.EntityData.Children.Append("fw-rev", types.YChild{"FwRev", &hardwareRevision.FwRev})
    hardwareRevision.EntityData.Children.Append("sw-rev", types.YChild{"SwRev", &hardwareRevision.SwRev})
    hardwareRevision.EntityData.Children.Append("dimm-rev", types.YChild{"DimmRev", &hardwareRevision.DimmRev})
    hardwareRevision.EntityData.Children.Append("ssd-rev", types.YChild{"SsdRev", &hardwareRevision.SsdRev})
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()
    hardwareRevision.EntityData.Leafs.Append("node-description", types.YLeaf{"NodeDescription", hardwareRevision.NodeDescription})
    hardwareRevision.EntityData.Leafs.Append("version", types.YLeaf{"Version", hardwareRevision.Version})

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_HwRev) GetEntityData() *types.CommonEntityData {
    hwRev.EntityData.YFilter = hwRev.YFilter
    hwRev.EntityData.YangName = "hw-rev"
    hwRev.EntityData.BundleName = "cisco_ios_xr"
    hwRev.EntityData.ParentYangName = "hardware-revision"
    hwRev.EntityData.SegmentPath = "hw-rev"
    hwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/hardware-revision/" + hwRev.EntityData.SegmentPath
    hwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwRev.EntityData.Children = types.NewOrderedMap()
    hwRev.EntityData.Leafs = types.NewOrderedMap()
    hwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", hwRev.MajorRevision})
    hwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", hwRev.MinorRevision})

    hwRev.EntityData.YListKeys = []string {}

    return &(hwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_FwRev) GetEntityData() *types.CommonEntityData {
    fwRev.EntityData.YFilter = fwRev.YFilter
    fwRev.EntityData.YangName = "fw-rev"
    fwRev.EntityData.BundleName = "cisco_ios_xr"
    fwRev.EntityData.ParentYangName = "hardware-revision"
    fwRev.EntityData.SegmentPath = "fw-rev"
    fwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/hardware-revision/" + fwRev.EntityData.SegmentPath
    fwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwRev.EntityData.Children = types.NewOrderedMap()
    fwRev.EntityData.Leafs = types.NewOrderedMap()
    fwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", fwRev.MajorRevision})
    fwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", fwRev.MinorRevision})

    fwRev.EntityData.YListKeys = []string {}

    return &(fwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SwRev) GetEntityData() *types.CommonEntityData {
    swRev.EntityData.YFilter = swRev.YFilter
    swRev.EntityData.YangName = "sw-rev"
    swRev.EntityData.BundleName = "cisco_ios_xr"
    swRev.EntityData.ParentYangName = "hardware-revision"
    swRev.EntityData.SegmentPath = "sw-rev"
    swRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/hardware-revision/" + swRev.EntityData.SegmentPath
    swRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swRev.EntityData.Children = types.NewOrderedMap()
    swRev.EntityData.Leafs = types.NewOrderedMap()
    swRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", swRev.MajorRevision})
    swRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", swRev.MinorRevision})

    swRev.EntityData.YListKeys = []string {}

    return &(swRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev struct {
    EntityData types.CommonEntityData
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

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_DimmRev) GetEntityData() *types.CommonEntityData {
    dimmRev.EntityData.YFilter = dimmRev.YFilter
    dimmRev.EntityData.YangName = "dimm-rev"
    dimmRev.EntityData.BundleName = "cisco_ios_xr"
    dimmRev.EntityData.ParentYangName = "hardware-revision"
    dimmRev.EntityData.SegmentPath = "dimm-rev"
    dimmRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/hardware-revision/" + dimmRev.EntityData.SegmentPath
    dimmRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dimmRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dimmRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dimmRev.EntityData.Children = types.NewOrderedMap()
    dimmRev.EntityData.Leafs = types.NewOrderedMap()
    dimmRev.EntityData.Leafs.Append("size", types.YLeaf{"Size", dimmRev.Size})
    dimmRev.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", dimmRev.Speed})
    dimmRev.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", dimmRev.Locator})
    dimmRev.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", dimmRev.Cas})

    dimmRev.EntityData.YListKeys = []string {}

    return &(dimmRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_HardwareRevision_HardwareRevision_SsdRev) GetEntityData() *types.CommonEntityData {
    ssdRev.EntityData.YFilter = ssdRev.YFilter
    ssdRev.EntityData.YangName = "ssd-rev"
    ssdRev.EntityData.BundleName = "cisco_ios_xr"
    ssdRev.EntityData.ParentYangName = "hardware-revision"
    ssdRev.EntityData.SegmentPath = "ssd-rev"
    ssdRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/hardware-revision/hardware-revision/" + ssdRev.EntityData.SegmentPath
    ssdRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssdRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssdRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssdRev.EntityData.Children = types.NewOrderedMap()
    ssdRev.EntityData.Leafs = types.NewOrderedMap()
    ssdRev.EntityData.Leafs.Append("number", types.YLeaf{"Number", ssdRev.Number})
    ssdRev.EntityData.Leafs.Append("fw-rev", types.YLeaf{"FwRev", ssdRev.FwRev})
    ssdRev.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", ssdRev.SerialNumber})

    ssdRev.EntityData.YListKeys = []string {}

    return &(ssdRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition
// CBC active partition
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcActivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcActivePartition) GetEntityData() *types.CommonEntityData {
    cbcActivePartition.EntityData.YFilter = cbcActivePartition.YFilter
    cbcActivePartition.EntityData.YangName = "cbc-active-partition"
    cbcActivePartition.EntityData.BundleName = "cisco_ios_xr"
    cbcActivePartition.EntityData.ParentYangName = "node"
    cbcActivePartition.EntityData.SegmentPath = "cbc-active-partition"
    cbcActivePartition.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/" + cbcActivePartition.EntityData.SegmentPath
    cbcActivePartition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbcActivePartition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbcActivePartition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbcActivePartition.EntityData.Children = types.NewOrderedMap()
    cbcActivePartition.EntityData.Leafs = types.NewOrderedMap()
    cbcActivePartition.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", cbcActivePartition.MajorRevision})
    cbcActivePartition.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", cbcActivePartition.MinorRevision})

    cbcActivePartition.EntityData.YListKeys = []string {}

    return &(cbcActivePartition.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition
// CBC inactive partition
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (cbcInactivePartition *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Node_CbcInactivePartition) GetEntityData() *types.CommonEntityData {
    cbcInactivePartition.EntityData.YFilter = cbcInactivePartition.YFilter
    cbcInactivePartition.EntityData.YangName = "cbc-inactive-partition"
    cbcInactivePartition.EntityData.BundleName = "cisco_ios_xr"
    cbcInactivePartition.EntityData.ParentYangName = "node"
    cbcInactivePartition.EntityData.SegmentPath = "cbc-inactive-partition"
    cbcInactivePartition.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/node/" + cbcInactivePartition.EntityData.SegmentPath
    cbcInactivePartition.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbcInactivePartition.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbcInactivePartition.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbcInactivePartition.EntityData.Children = types.NewOrderedMap()
    cbcInactivePartition.EntityData.Leafs = types.NewOrderedMap()
    cbcInactivePartition.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", cbcInactivePartition.MajorRevision})
    cbcInactivePartition.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", cbcInactivePartition.MinorRevision})

    cbcInactivePartition.EntityData.YListKeys = []string {}

    return &(cbcInactivePartition.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa
// SPA information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa struct {
    EntityData types.CommonEntityData
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

func (spa *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa) GetEntityData() *types.CommonEntityData {
    spa.EntityData.YFilter = spa.YFilter
    spa.EntityData.YangName = "spa"
    spa.EntityData.BundleName = "cisco_ios_xr"
    spa.EntityData.ParentYangName = "detail"
    spa.EntityData.SegmentPath = "spa"
    spa.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/" + spa.EntityData.SegmentPath
    spa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spa.EntityData.Children = types.NewOrderedMap()
    spa.EntityData.Children.Append("main", types.YChild{"Main", &spa.Main})
    spa.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", &spa.HardwareRevision})
    spa.EntityData.Leafs = types.NewOrderedMap()
    spa.EntityData.Leafs.Append("node", types.YLeaf{"Node", spa.Node})
    spa.EntityData.Leafs.Append("name", types.YLeaf{"Name", spa.Name})
    spa.EntityData.Leafs.Append("pca-unit-number", types.YLeaf{"PcaUnitNumber", spa.PcaUnitNumber})
    spa.EntityData.Leafs.Append("pca-revision", types.YLeaf{"PcaRevision", spa.PcaRevision})
    spa.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", spa.Pid})
    spa.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", spa.Vid})
    spa.EntityData.Leafs.Append("clei", types.YLeaf{"Clei", spa.Clei})
    spa.EntityData.Leafs.Append("node-state", types.YLeaf{"NodeState", spa.NodeState})

    spa.EntityData.YListKeys = []string {}

    return &(spa.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main
// Main
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main struct {
    EntityData types.CommonEntityData
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

func (main *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_Main) GetEntityData() *types.CommonEntityData {
    main.EntityData.YFilter = main.YFilter
    main.EntityData.YangName = "main"
    main.EntityData.BundleName = "cisco_ios_xr"
    main.EntityData.ParentYangName = "spa"
    main.EntityData.SegmentPath = "main"
    main.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/" + main.EntityData.SegmentPath
    main.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    main.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    main.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    main.EntityData.Children = types.NewOrderedMap()
    main.EntityData.Leafs = types.NewOrderedMap()
    main.EntityData.Leafs.Append("board-type", types.YLeaf{"BoardType", main.BoardType})
    main.EntityData.Leafs.Append("tan", types.YLeaf{"Tan", main.Tan})
    main.EntityData.Leafs.Append("tan-revision", types.YLeaf{"TanRevision", main.TanRevision})
    main.EntityData.Leafs.Append("deviation-number", types.YLeaf{"DeviationNumber", main.DeviationNumber})
    main.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", main.SerialNumber})

    main.EntityData.YListKeys = []string {}

    return &(main.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision
// Hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Board FPGA/CPLD/ASIC hardware revision. The type is slice of
    // Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision.
    HardwareRevision []*Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision
}

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "spa"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision"
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hardware-revision", types.YChild{"HardwareRevision", nil})
    for i := range hardwareRevision.HardwareRevision {
        types.SetYListKey(hardwareRevision.HardwareRevision[i], i)
        hardwareRevision.EntityData.Children.Append(types.GetSegmentPath(hardwareRevision.HardwareRevision[i]), types.YChild{"HardwareRevision", hardwareRevision.HardwareRevision[i]})
    }
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision
// Board FPGA/CPLD/ASIC hardware revision
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hardwareRevision *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision) GetEntityData() *types.CommonEntityData {
    hardwareRevision.EntityData.YFilter = hardwareRevision.YFilter
    hardwareRevision.EntityData.YangName = "hardware-revision"
    hardwareRevision.EntityData.BundleName = "cisco_ios_xr"
    hardwareRevision.EntityData.ParentYangName = "hardware-revision"
    hardwareRevision.EntityData.SegmentPath = "hardware-revision" + types.AddNoKeyToken(hardwareRevision)
    hardwareRevision.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/" + hardwareRevision.EntityData.SegmentPath
    hardwareRevision.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareRevision.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareRevision.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareRevision.EntityData.Children = types.NewOrderedMap()
    hardwareRevision.EntityData.Children.Append("hw-rev", types.YChild{"HwRev", &hardwareRevision.HwRev})
    hardwareRevision.EntityData.Children.Append("fw-rev", types.YChild{"FwRev", &hardwareRevision.FwRev})
    hardwareRevision.EntityData.Children.Append("sw-rev", types.YChild{"SwRev", &hardwareRevision.SwRev})
    hardwareRevision.EntityData.Children.Append("dimm-rev", types.YChild{"DimmRev", &hardwareRevision.DimmRev})
    hardwareRevision.EntityData.Children.Append("ssd-rev", types.YChild{"SsdRev", &hardwareRevision.SsdRev})
    hardwareRevision.EntityData.Leafs = types.NewOrderedMap()
    hardwareRevision.EntityData.Leafs.Append("node-description", types.YLeaf{"NodeDescription", hardwareRevision.NodeDescription})
    hardwareRevision.EntityData.Leafs.Append("version", types.YLeaf{"Version", hardwareRevision.Version})

    hardwareRevision.EntityData.YListKeys = []string {}

    return &(hardwareRevision.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev
// Hardware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (hwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_HwRev) GetEntityData() *types.CommonEntityData {
    hwRev.EntityData.YFilter = hwRev.YFilter
    hwRev.EntityData.YangName = "hw-rev"
    hwRev.EntityData.BundleName = "cisco_ios_xr"
    hwRev.EntityData.ParentYangName = "hardware-revision"
    hwRev.EntityData.SegmentPath = "hw-rev"
    hwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/hardware-revision/" + hwRev.EntityData.SegmentPath
    hwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwRev.EntityData.Children = types.NewOrderedMap()
    hwRev.EntityData.Leafs = types.NewOrderedMap()
    hwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", hwRev.MajorRevision})
    hwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", hwRev.MinorRevision})

    hwRev.EntityData.YListKeys = []string {}

    return &(hwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev
// Firmware version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (fwRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_FwRev) GetEntityData() *types.CommonEntityData {
    fwRev.EntityData.YFilter = fwRev.YFilter
    fwRev.EntityData.YangName = "fw-rev"
    fwRev.EntityData.BundleName = "cisco_ios_xr"
    fwRev.EntityData.ParentYangName = "hardware-revision"
    fwRev.EntityData.SegmentPath = "fw-rev"
    fwRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/hardware-revision/" + fwRev.EntityData.SegmentPath
    fwRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fwRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fwRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fwRev.EntityData.Children = types.NewOrderedMap()
    fwRev.EntityData.Leafs = types.NewOrderedMap()
    fwRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", fwRev.MajorRevision})
    fwRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", fwRev.MinorRevision})

    fwRev.EntityData.YListKeys = []string {}

    return &(fwRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev
// Software version
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Major revision. The type is interface{} with range: 0..4294967295.
    MajorRevision interface{}

    // Minor revision. The type is interface{} with range: 0..4294967295.
    MinorRevision interface{}
}

func (swRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SwRev) GetEntityData() *types.CommonEntityData {
    swRev.EntityData.YFilter = swRev.YFilter
    swRev.EntityData.YangName = "sw-rev"
    swRev.EntityData.BundleName = "cisco_ios_xr"
    swRev.EntityData.ParentYangName = "hardware-revision"
    swRev.EntityData.SegmentPath = "sw-rev"
    swRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/hardware-revision/" + swRev.EntityData.SegmentPath
    swRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swRev.EntityData.Children = types.NewOrderedMap()
    swRev.EntityData.Leafs = types.NewOrderedMap()
    swRev.EntityData.Leafs.Append("major-revision", types.YLeaf{"MajorRevision", swRev.MajorRevision})
    swRev.EntityData.Leafs.Append("minor-revision", types.YLeaf{"MinorRevision", swRev.MinorRevision})

    swRev.EntityData.YListKeys = []string {}

    return &(swRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev
// DIMM version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev struct {
    EntityData types.CommonEntityData
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

func (dimmRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_DimmRev) GetEntityData() *types.CommonEntityData {
    dimmRev.EntityData.YFilter = dimmRev.YFilter
    dimmRev.EntityData.YangName = "dimm-rev"
    dimmRev.EntityData.BundleName = "cisco_ios_xr"
    dimmRev.EntityData.ParentYangName = "hardware-revision"
    dimmRev.EntityData.SegmentPath = "dimm-rev"
    dimmRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/hardware-revision/" + dimmRev.EntityData.SegmentPath
    dimmRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dimmRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dimmRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dimmRev.EntityData.Children = types.NewOrderedMap()
    dimmRev.EntityData.Leafs = types.NewOrderedMap()
    dimmRev.EntityData.Leafs.Append("size", types.YLeaf{"Size", dimmRev.Size})
    dimmRev.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", dimmRev.Speed})
    dimmRev.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", dimmRev.Locator})
    dimmRev.EntityData.Leafs.Append("cas", types.YLeaf{"Cas", dimmRev.Cas})

    dimmRev.EntityData.YListKeys = []string {}

    return &(dimmRev.EntityData)
}

// Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev
// SSD version information
type Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSD number. The type is string.
    Number interface{}

    // Firmware revision. The type is string.
    FwRev interface{}

    // Serial number. The type is string.
    SerialNumber interface{}
}

func (ssdRev *Diag_Racks_Rack_Slots_Slot_Instances_Instance_Detail_Spa_HardwareRevision_HardwareRevision_SsdRev) GetEntityData() *types.CommonEntityData {
    ssdRev.EntityData.YFilter = ssdRev.YFilter
    ssdRev.EntityData.YangName = "ssd-rev"
    ssdRev.EntityData.BundleName = "cisco_ios_xr"
    ssdRev.EntityData.ParentYangName = "hardware-revision"
    ssdRev.EntityData.SegmentPath = "ssd-rev"
    ssdRev.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/slots/slot/instances/instance/detail/spa/hardware-revision/hardware-revision/" + ssdRev.EntityData.SegmentPath
    ssdRev.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssdRev.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssdRev.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssdRev.EntityData.Children = types.NewOrderedMap()
    ssdRev.EntityData.Leafs = types.NewOrderedMap()
    ssdRev.EntityData.Leafs.Append("number", types.YLeaf{"Number", ssdRev.Number})
    ssdRev.EntityData.Leafs.Append("fw-rev", types.YLeaf{"FwRev", ssdRev.FwRev})
    ssdRev.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", ssdRev.SerialNumber})

    ssdRev.EntityData.YListKeys = []string {}

    return &(ssdRev.EntityData)
}

// Diag_Racks_Rack_Summary
// Summary information
type Diag_Racks_Rack_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary data. The type is slice of Diag_Racks_Rack_Summary_Summary.
    Summary []*Diag_Racks_Rack_Summary_Summary
}

func (summary *Diag_Racks_Rack_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "rack"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range summary.Summary {
        types.SetYListKey(summary.Summary[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.Summary[i]), types.YChild{"Summary", summary.Summary[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Diag_Racks_Rack_Summary_Summary
// Summary data
type Diag_Racks_Rack_Summary_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (summary *Diag_Racks_Rack_Summary_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "summary"
    summary.EntityData.SegmentPath = "summary" + types.AddNoKeyToken(summary)
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-sc-diag-oper:diag/racks/rack/summary/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("type", types.YLeaf{"Type", summary.Type})
    summary.EntityData.Leafs.Append("node", types.YLeaf{"Node", summary.Node})
    summary.EntityData.Leafs.Append("slot-type", types.YLeaf{"SlotType", summary.SlotType})
    summary.EntityData.Leafs.Append("description", types.YLeaf{"Description", summary.Description})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

