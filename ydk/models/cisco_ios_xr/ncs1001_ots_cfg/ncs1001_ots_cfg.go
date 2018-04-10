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

// OtsPsmLockoutFrom represents Ots psm lockout from
type OtsPsmLockoutFrom string

const (
    // Working port
    OtsPsmLockoutFrom_working OtsPsmLockoutFrom = "working"

    // Protected port
    OtsPsmLockoutFrom_protected OtsPsmLockoutFrom = "protected"
)

// HardwareModule
// NCS1k HW module config
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is slice of HardwareModule_Node.
    Node []HardwareModule_Node
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1001-ots-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = make(map[string]types.YChild)
    hardwareModule.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range hardwareModule.Node {
        hardwareModule.EntityData.Children[types.GetSegmentPath(&hardwareModule.Node[i])] = types.YChild{"Node", &hardwareModule.Node[i]}
    }
    hardwareModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareModule.EntityData)
}

// HardwareModule_Node
// Node
type HardwareModule_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Location interface{}

    // Slot Id. The type is slice of HardwareModule_Node_Slot.
    Slot []HardwareModule_Node_Slot
}

func (node *HardwareModule_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "hardware-module"
    node.EntityData.SegmentPath = "node" + "[location='" + fmt.Sprintf("%v", node.Location) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range node.Slot {
        node.EntityData.Children[types.GetSegmentPath(&node.Slot[i])] = types.YChild{"Slot", &node.Slot[i]}
    }
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["location"] = types.YLeaf{"Location", node.Location}
    return &(node.EntityData)
}

// HardwareModule_Node_Slot
// Slot Id
type HardwareModule_Node_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Set Slot. The type is interface{} with range:
    // 1..3.
    SlotId interface{}

    // Amplifier Configs.
    Amplifier HardwareModule_Node_Slot_Amplifier

    // PSM Configs.
    Psm HardwareModule_Node_Slot_Psm
}

func (slot *HardwareModule_Node_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "node"
    slot.EntityData.SegmentPath = "slot" + "[slot-id='" + fmt.Sprintf("%v", slot.SlotId) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["amplifier"] = types.YChild{"Amplifier", &slot.Amplifier}
    slot.EntityData.Children["psm"] = types.YChild{"Psm", &slot.Psm}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot-id"] = types.YLeaf{"SlotId", slot.SlotId}
    return &(slot.EntityData)
}

// HardwareModule_Node_Slot_Amplifier
// Amplifier Configs
type HardwareModule_Node_Slot_Amplifier struct {
    EntityData types.CommonEntityData
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

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetEntityData() *types.CommonEntityData {
    amplifier.EntityData.YFilter = amplifier.YFilter
    amplifier.EntityData.YangName = "amplifier"
    amplifier.EntityData.BundleName = "cisco_ios_xr"
    amplifier.EntityData.ParentYangName = "slot"
    amplifier.EntityData.SegmentPath = "amplifier"
    amplifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifier.EntityData.Children = make(map[string]types.YChild)
    amplifier.EntityData.Leafs = make(map[string]types.YLeaf)
    amplifier.EntityData.Leafs["node-type"] = types.YLeaf{"NodeType", amplifier.NodeType}
    amplifier.EntityData.Leafs["grid-mode"] = types.YLeaf{"GridMode", amplifier.GridMode}
    amplifier.EntityData.Leafs["udc-vlan"] = types.YLeaf{"UdcVlan", amplifier.UdcVlan}
    return &(amplifier.EntityData)
}

// HardwareModule_Node_Slot_Psm
// PSM Configs
type HardwareModule_Node_Slot_Psm struct {
    EntityData types.CommonEntityData
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
}

func (psm *HardwareModule_Node_Slot_Psm) GetEntityData() *types.CommonEntityData {
    psm.EntityData.YFilter = psm.YFilter
    psm.EntityData.YangName = "psm"
    psm.EntityData.BundleName = "cisco_ios_xr"
    psm.EntityData.ParentYangName = "slot"
    psm.EntityData.SegmentPath = "psm"
    psm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    psm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    psm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    psm.EntityData.Children = make(map[string]types.YChild)
    psm.EntityData.Leafs = make(map[string]types.YLeaf)
    psm.EntityData.Leafs["mono-dir"] = types.YLeaf{"MonoDir", psm.MonoDir}
    psm.EntityData.Leafs["auto-threshold"] = types.YLeaf{"AutoThreshold", psm.AutoThreshold}
    psm.EntityData.Leafs["path-protection"] = types.YLeaf{"PathProtection", psm.PathProtection}
    psm.EntityData.Leafs["section-protection"] = types.YLeaf{"SectionProtection", psm.SectionProtection}
    psm.EntityData.Leafs["lockout-from"] = types.YLeaf{"LockoutFrom", psm.LockoutFrom}
    return &(psm.EntityData)
}

