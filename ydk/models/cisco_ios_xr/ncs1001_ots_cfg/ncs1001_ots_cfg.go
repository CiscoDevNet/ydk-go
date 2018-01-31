// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1001-ots package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: NCS1k HW module config
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1001_ots_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1001_ots_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1001-ots-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// OtsPsmLockoutFrom represents Ots psm lockout from
type OtsPsmLockoutFrom string

const (
    // Working port
    OtsPsmLockoutFrom_working OtsPsmLockoutFrom = "working"

    // Protected port
    OtsPsmLockoutFrom_protected OtsPsmLockoutFrom = "protected"
)

// OtsAmplifierGridMode represents Ots amplifier grid mode
type OtsAmplifierGridMode string

const (
    // 100GHz mode
    OtsAmplifierGridMode_Y_100g_hz OtsAmplifierGridMode = "100g-hz"

    // 50GHz mode
    OtsAmplifierGridMode_Y_50g_hz OtsAmplifierGridMode = "50g-hz"

    // Gridless mode
    OtsAmplifierGridMode_gr_idle_ss OtsAmplifierGridMode = "gr-idle-ss"
)

// OtsPsmManualSwitch represents Ots psm manual switch
type OtsPsmManualSwitch string

const (
    // Working port
    OtsPsmManualSwitch_working OtsPsmManualSwitch = "working"

    // Protected port
    OtsPsmManualSwitch_protected OtsPsmManualSwitch = "protected"
)

// OtsAmplifierNode represents Ots amplifier node
type OtsAmplifierNode string

const (
    // Nodetype TERM
    OtsAmplifierNode_term OtsAmplifierNode = "term"

    // Nodetype InLine Amplifier
    OtsAmplifierNode_ila OtsAmplifierNode = "ila"

    // Nodetype ROADM
    OtsAmplifierNode_roadm OtsAmplifierNode = "roadm"
)

// HardwareModule
// NCS1k HW module config
type HardwareModule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is slice of HardwareModule_Node.
    Node []HardwareModule_Node
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range hardwareModule.Node {
            if hardwareModule.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Node{}
        hardwareModule.Node = append(hardwareModule.Node, child)
        return &hardwareModule.Node[len(hardwareModule.Node)-1]
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareModule.Node {
        children[hardwareModule.Node[i].GetSegmentPath()] = &hardwareModule.Node[i]
    }
    return children
}

func (hardwareModule *HardwareModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModule *HardwareModule) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModule *HardwareModule) GetYangName() string { return "hardware-module" }

func (hardwareModule *HardwareModule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModule *HardwareModule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModule *HardwareModule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModule *HardwareModule) SetParent(parent types.Entity) { hardwareModule.parent = parent }

func (hardwareModule *HardwareModule) GetParent() types.Entity { return hardwareModule.parent }

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-ncs1001-ots-cfg" }

// HardwareModule_Node
// Node
type HardwareModule_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Location interface{}

    // Slot Id. The type is slice of HardwareModule_Node_Slot.
    Slot []HardwareModule_Node_Slot
}

func (node *HardwareModule_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModule_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModule_Node) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "slot" { return "Slot" }
    return ""
}

func (node *HardwareModule_Node) GetSegmentPath() string {
    return "node" + "[location='" + fmt.Sprintf("%v", node.Location) + "']"
}

func (node *HardwareModule_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range node.Slot {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Node_Slot{}
        node.Slot = append(node.Slot, child)
        return &node.Slot[len(node.Slot)-1]
    }
    return nil
}

func (node *HardwareModule_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.Slot {
        children[node.Slot[i].GetSegmentPath()] = &node.Slot[i]
    }
    return children
}

func (node *HardwareModule_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = node.Location
    return leafs
}

func (node *HardwareModule_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModule_Node) GetYangName() string { return "node" }

func (node *HardwareModule_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModule_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModule_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModule_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModule_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModule_Node) GetParentYangName() string { return "hardware-module" }

// HardwareModule_Node_Slot
// Slot Id
type HardwareModule_Node_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Set Slot. The type is interface{} with range:
    // 1..3.
    SlotId interface{}

    // Amplifier Configs.
    Amplifier HardwareModule_Node_Slot_Amplifier

    // PSM Configs.
    Psm HardwareModule_Node_Slot_Psm
}

func (slot *HardwareModule_Node_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *HardwareModule_Node_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *HardwareModule_Node_Slot) GetGoName(yname string) string {
    if yname == "slot-id" { return "SlotId" }
    if yname == "amplifier" { return "Amplifier" }
    if yname == "psm" { return "Psm" }
    return ""
}

func (slot *HardwareModule_Node_Slot) GetSegmentPath() string {
    return "slot" + "[slot-id='" + fmt.Sprintf("%v", slot.SlotId) + "']"
}

func (slot *HardwareModule_Node_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "amplifier" {
        return &slot.Amplifier
    }
    if childYangName == "psm" {
        return &slot.Psm
    }
    return nil
}

func (slot *HardwareModule_Node_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["amplifier"] = &slot.Amplifier
    children["psm"] = &slot.Psm
    return children
}

func (slot *HardwareModule_Node_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot-id"] = slot.SlotId
    return leafs
}

func (slot *HardwareModule_Node_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *HardwareModule_Node_Slot) GetYangName() string { return "slot" }

func (slot *HardwareModule_Node_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *HardwareModule_Node_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *HardwareModule_Node_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *HardwareModule_Node_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *HardwareModule_Node_Slot) GetParent() types.Entity { return slot.parent }

func (slot *HardwareModule_Node_Slot) GetParentYangName() string { return "node" }

// HardwareModule_Node_Slot_Amplifier
// Amplifier Configs
type HardwareModule_Node_Slot_Amplifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Define the type of node in which the amplifier is set to work. The type is
    // OtsAmplifierNode.
    NodeType interface{}

    // Define the working mode for the optical module. The type is
    // OtsAmplifierGridMode.
    GridMode interface{}

    // Define the VLAN ID in range <2-4080>. The type is interface{} with range:
    // 2..4080.
    UdcVlan interface{}
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetFilter() yfilter.YFilter { return amplifier.YFilter }

func (amplifier *HardwareModule_Node_Slot_Amplifier) SetFilter(yf yfilter.YFilter) { amplifier.YFilter = yf }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetGoName(yname string) string {
    if yname == "node-type" { return "NodeType" }
    if yname == "grid-mode" { return "GridMode" }
    if yname == "udc-vlan" { return "UdcVlan" }
    return ""
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetSegmentPath() string {
    return "amplifier"
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-type"] = amplifier.NodeType
    leafs["grid-mode"] = amplifier.GridMode
    leafs["udc-vlan"] = amplifier.UdcVlan
    return leafs
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetBundleName() string { return "cisco_ios_xr" }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetYangName() string { return "amplifier" }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (amplifier *HardwareModule_Node_Slot_Amplifier) SetParent(parent types.Entity) { amplifier.parent = parent }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetParent() types.Entity { return amplifier.parent }

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetParentYangName() string { return "slot" }

// HardwareModule_Node_Slot_Psm
// PSM Configs
type HardwareModule_Node_Slot_Psm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Psm Uni directional configuration. The type is bool.
    MonoDir interface{}

    // Psm Automatic Threshold Setting. The type is bool.
    AutoThreshold interface{}

    // Psm path protection configuration. The type is bool.
    PathProtection interface{}

    // Psm section protection configuration. The type is bool.
    SectionProtection interface{}

    // Exclude selected port from protection. The type is OtsPsmLockoutFrom.
    LockoutFrom interface{}

    // Switch active path on selected port. The type is OtsPsmManualSwitch.
    ManualSwitchTo interface{}
}

func (psm *HardwareModule_Node_Slot_Psm) GetFilter() yfilter.YFilter { return psm.YFilter }

func (psm *HardwareModule_Node_Slot_Psm) SetFilter(yf yfilter.YFilter) { psm.YFilter = yf }

func (psm *HardwareModule_Node_Slot_Psm) GetGoName(yname string) string {
    if yname == "mono-dir" { return "MonoDir" }
    if yname == "auto-threshold" { return "AutoThreshold" }
    if yname == "path-protection" { return "PathProtection" }
    if yname == "section-protection" { return "SectionProtection" }
    if yname == "lockout-from" { return "LockoutFrom" }
    if yname == "manual-switch-to" { return "ManualSwitchTo" }
    return ""
}

func (psm *HardwareModule_Node_Slot_Psm) GetSegmentPath() string {
    return "psm"
}

func (psm *HardwareModule_Node_Slot_Psm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (psm *HardwareModule_Node_Slot_Psm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (psm *HardwareModule_Node_Slot_Psm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mono-dir"] = psm.MonoDir
    leafs["auto-threshold"] = psm.AutoThreshold
    leafs["path-protection"] = psm.PathProtection
    leafs["section-protection"] = psm.SectionProtection
    leafs["lockout-from"] = psm.LockoutFrom
    leafs["manual-switch-to"] = psm.ManualSwitchTo
    return leafs
}

func (psm *HardwareModule_Node_Slot_Psm) GetBundleName() string { return "cisco_ios_xr" }

func (psm *HardwareModule_Node_Slot_Psm) GetYangName() string { return "psm" }

func (psm *HardwareModule_Node_Slot_Psm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (psm *HardwareModule_Node_Slot_Psm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (psm *HardwareModule_Node_Slot_Psm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (psm *HardwareModule_Node_Slot_Psm) SetParent(parent types.Entity) { psm.parent = parent }

func (psm *HardwareModule_Node_Slot_Psm) GetParent() types.Entity { return psm.parent }

func (psm *HardwareModule_Node_Slot_Psm) GetParentYangName() string { return "slot" }

