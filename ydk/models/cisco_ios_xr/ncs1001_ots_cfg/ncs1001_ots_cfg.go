// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1001-ots package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: NCS1k HW module config
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// OtsOtdrDirection represents Ots otdr direction
type OtsOtdrDirection string

const (
    // Tx
    OtsOtdrDirection_tx OtsOtdrDirection = "tx"

    // Rx
    OtsOtdrDirection_rx OtsOtdrDirection = "rx"
)

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is slice of HardwareModule_Node.
    Node []*HardwareModule_Node
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1001-ots-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module"
    hardwareModule.EntityData.AbsolutePath = hardwareModule.EntityData.SegmentPath
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = types.NewOrderedMap()
    hardwareModule.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range hardwareModule.Node {
        hardwareModule.EntityData.Children.Append(types.GetSegmentPath(hardwareModule.Node[i]), types.YChild{"Node", hardwareModule.Node[i]})
    }
    hardwareModule.EntityData.Leafs = types.NewOrderedMap()

    hardwareModule.EntityData.YListKeys = []string {}

    return &(hardwareModule.EntityData)
}

// HardwareModule_Node
// Node
type HardwareModule_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Location interface{}

    // Slot Id. The type is slice of HardwareModule_Node_Slot.
    Slot []*HardwareModule_Node_Slot
}

func (node *HardwareModule_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "hardware-module"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Location, "location")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range node.Slot {
        node.EntityData.Children.Append(types.GetSegmentPath(node.Slot[i]), types.YChild{"Slot", node.Slot[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("location", types.YLeaf{"Location", node.Location})

    node.EntityData.YListKeys = []string {"Location"}

    return &(node.EntityData)
}

// HardwareModule_Node_Slot
// Slot Id
type HardwareModule_Node_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set Slot. The type is interface{} with range:
    // 1..3.
    SlotId interface{}

    // OTDR Configs.
    Otdrs HardwareModule_Node_Slot_Otdrs

    // Otdr Threshold.
    OtdrThresholds HardwareModule_Node_Slot_OtdrThresholds

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
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.SlotId, "slot-id")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("otdrs", types.YChild{"Otdrs", &slot.Otdrs})
    slot.EntityData.Children.Append("otdr-thresholds", types.YChild{"OtdrThresholds", &slot.OtdrThresholds})
    slot.EntityData.Children.Append("amplifier", types.YChild{"Amplifier", &slot.Amplifier})
    slot.EntityData.Children.Append("psm", types.YChild{"Psm", &slot.Psm})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", slot.SlotId})

    slot.EntityData.YListKeys = []string {"SlotId"}

    return &(slot.EntityData)
}

// HardwareModule_Node_Slot_Otdrs
// OTDR Configs
type HardwareModule_Node_Slot_Otdrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Id. The type is slice of HardwareModule_Node_Slot_Otdrs_Otdr.
    Otdr []*HardwareModule_Node_Slot_Otdrs_Otdr
}

func (otdrs *HardwareModule_Node_Slot_Otdrs) GetEntityData() *types.CommonEntityData {
    otdrs.EntityData.YFilter = otdrs.YFilter
    otdrs.EntityData.YangName = "otdrs"
    otdrs.EntityData.BundleName = "cisco_ios_xr"
    otdrs.EntityData.ParentYangName = "slot"
    otdrs.EntityData.SegmentPath = "otdrs"
    otdrs.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/" + otdrs.EntityData.SegmentPath
    otdrs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrs.EntityData.Children = types.NewOrderedMap()
    otdrs.EntityData.Children.Append("otdr", types.YChild{"Otdr", nil})
    for i := range otdrs.Otdr {
        otdrs.EntityData.Children.Append(types.GetSegmentPath(otdrs.Otdr[i]), types.YChild{"Otdr", otdrs.Otdr[i]})
    }
    otdrs.EntityData.Leafs = types.NewOrderedMap()

    otdrs.EntityData.YListKeys = []string {}

    return &(otdrs.EntityData)
}

// HardwareModule_Node_Slot_Otdrs_Otdr
// Otdr Id
type HardwareModule_Node_Slot_Otdrs_Otdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port Id. The type is interface{} with range: 1..2.
    Port interface{}

    // This attribute is a key. Direction Id. The type is OtsOtdrDirection.
    Direction interface{}

    // Otdr Insertion Loss, supported values [0 ,500] in units of 0.1dB. The type
    // is interface{} with range: 0..500.
    TotalLoss interface{}

    // Mode Auto.
    ModeAuto HardwareModule_Node_Slot_Otdrs_Otdr_ModeAuto

    // Mode Expert.
    ModeExpert HardwareModule_Node_Slot_Otdrs_Otdr_ModeExpert
}

func (otdr *HardwareModule_Node_Slot_Otdrs_Otdr) GetEntityData() *types.CommonEntityData {
    otdr.EntityData.YFilter = otdr.YFilter
    otdr.EntityData.YangName = "otdr"
    otdr.EntityData.BundleName = "cisco_ios_xr"
    otdr.EntityData.ParentYangName = "otdrs"
    otdr.EntityData.SegmentPath = "otdr" + types.AddKeyToken(otdr.Port, "port") + types.AddKeyToken(otdr.Direction, "direction")
    otdr.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/otdrs/" + otdr.EntityData.SegmentPath
    otdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdr.EntityData.Children = types.NewOrderedMap()
    otdr.EntityData.Children.Append("mode-auto", types.YChild{"ModeAuto", &otdr.ModeAuto})
    otdr.EntityData.Children.Append("mode-expert", types.YChild{"ModeExpert", &otdr.ModeExpert})
    otdr.EntityData.Leafs = types.NewOrderedMap()
    otdr.EntityData.Leafs.Append("port", types.YLeaf{"Port", otdr.Port})
    otdr.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdr.Direction})
    otdr.EntityData.Leafs.Append("total-loss", types.YLeaf{"TotalLoss", otdr.TotalLoss})

    otdr.EntityData.YListKeys = []string {"Port", "Direction"}

    return &(otdr.EntityData)
}

// HardwareModule_Node_Slot_Otdrs_Otdr_ModeAuto
// Mode Auto
type HardwareModule_Node_Slot_Otdrs_Otdr_ModeAuto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Loss Sensitivity, supported values [4 ,50] in units of 0.1dB. The type
    // is interface{} with range: 4..50.
    LossSensitivity interface{}

    // Otdr Reflection Sensitivity, supported values [-400,-140] in units of
    // 0.1dB. The type is interface{} with range: -400..-140.
    ReflectionSensitivity interface{}
}

func (modeAuto *HardwareModule_Node_Slot_Otdrs_Otdr_ModeAuto) GetEntityData() *types.CommonEntityData {
    modeAuto.EntityData.YFilter = modeAuto.YFilter
    modeAuto.EntityData.YangName = "mode-auto"
    modeAuto.EntityData.BundleName = "cisco_ios_xr"
    modeAuto.EntityData.ParentYangName = "otdr"
    modeAuto.EntityData.SegmentPath = "mode-auto"
    modeAuto.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/otdrs/otdr/" + modeAuto.EntityData.SegmentPath
    modeAuto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modeAuto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modeAuto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modeAuto.EntityData.Children = types.NewOrderedMap()
    modeAuto.EntityData.Leafs = types.NewOrderedMap()
    modeAuto.EntityData.Leafs.Append("loss-sensitivity", types.YLeaf{"LossSensitivity", modeAuto.LossSensitivity})
    modeAuto.EntityData.Leafs.Append("reflection-sensitivity", types.YLeaf{"ReflectionSensitivity", modeAuto.ReflectionSensitivity})

    modeAuto.EntityData.YListKeys = []string {}

    return &(modeAuto.EntityData)
}

// HardwareModule_Node_Slot_Otdrs_Otdr_ModeExpert
// Mode Expert
type HardwareModule_Node_Slot_Otdrs_Otdr_ModeExpert struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fiber Resolution in meters, supported values [0,100]. The type is
    // interface{} with range: 0..100.
    FiberResolution interface{}

    // Capture Length in Km, supported values [0 ,150]. The type is interface{}
    // with range: 0..150.
    CaptureLength interface{}

    // Pulse Width in nanoseconds, supported values [8,100000]. The type is
    // interface{} with range: 8..100000. Units are nanosecond.
    PulseWidth interface{}

    // Measure time in seconds, supported values [0,360]. The type is interface{}
    // with range: 0..360. Units are second.
    MeasureTime interface{}

    // Otdr Loss Sensitivity, supported values [4 ,50] in units of 0.1dB. The type
    // is interface{} with range: 4..50.
    LossSensitivity interface{}

    // Capture Offset in Km, supported values [0 ,150]. The type is interface{}
    // with range: 0..150.
    CaptureOffset interface{}

    // Span Length in Km, supported values [0,150]. The type is interface{} with
    // range: 0..150.
    SpanLength interface{}

    // Otdr Reflection Sensitivity, supported values [-400,-140] in units of
    // 0.1dB. The type is interface{} with range: -400..-140.
    ReflectionSensitivity interface{}
}

func (modeExpert *HardwareModule_Node_Slot_Otdrs_Otdr_ModeExpert) GetEntityData() *types.CommonEntityData {
    modeExpert.EntityData.YFilter = modeExpert.YFilter
    modeExpert.EntityData.YangName = "mode-expert"
    modeExpert.EntityData.BundleName = "cisco_ios_xr"
    modeExpert.EntityData.ParentYangName = "otdr"
    modeExpert.EntityData.SegmentPath = "mode-expert"
    modeExpert.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/otdrs/otdr/" + modeExpert.EntityData.SegmentPath
    modeExpert.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modeExpert.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modeExpert.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modeExpert.EntityData.Children = types.NewOrderedMap()
    modeExpert.EntityData.Leafs = types.NewOrderedMap()
    modeExpert.EntityData.Leafs.Append("fiber-resolution", types.YLeaf{"FiberResolution", modeExpert.FiberResolution})
    modeExpert.EntityData.Leafs.Append("capture-length", types.YLeaf{"CaptureLength", modeExpert.CaptureLength})
    modeExpert.EntityData.Leafs.Append("pulse-width", types.YLeaf{"PulseWidth", modeExpert.PulseWidth})
    modeExpert.EntityData.Leafs.Append("measure-time", types.YLeaf{"MeasureTime", modeExpert.MeasureTime})
    modeExpert.EntityData.Leafs.Append("loss-sensitivity", types.YLeaf{"LossSensitivity", modeExpert.LossSensitivity})
    modeExpert.EntityData.Leafs.Append("capture-offset", types.YLeaf{"CaptureOffset", modeExpert.CaptureOffset})
    modeExpert.EntityData.Leafs.Append("span-length", types.YLeaf{"SpanLength", modeExpert.SpanLength})
    modeExpert.EntityData.Leafs.Append("reflection-sensitivity", types.YLeaf{"ReflectionSensitivity", modeExpert.ReflectionSensitivity})

    modeExpert.EntityData.YListKeys = []string {}

    return &(modeExpert.EntityData)
}

// HardwareModule_Node_Slot_OtdrThresholds
// Otdr Threshold
type HardwareModule_Node_Slot_OtdrThresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Threshold. The type is slice of
    // HardwareModule_Node_Slot_OtdrThresholds_OtdrThreshold.
    OtdrThreshold []*HardwareModule_Node_Slot_OtdrThresholds_OtdrThreshold
}

func (otdrThresholds *HardwareModule_Node_Slot_OtdrThresholds) GetEntityData() *types.CommonEntityData {
    otdrThresholds.EntityData.YFilter = otdrThresholds.YFilter
    otdrThresholds.EntityData.YangName = "otdr-thresholds"
    otdrThresholds.EntityData.BundleName = "cisco_ios_xr"
    otdrThresholds.EntityData.ParentYangName = "slot"
    otdrThresholds.EntityData.SegmentPath = "otdr-thresholds"
    otdrThresholds.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/" + otdrThresholds.EntityData.SegmentPath
    otdrThresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrThresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrThresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrThresholds.EntityData.Children = types.NewOrderedMap()
    otdrThresholds.EntityData.Children.Append("otdr-threshold", types.YChild{"OtdrThreshold", nil})
    for i := range otdrThresholds.OtdrThreshold {
        otdrThresholds.EntityData.Children.Append(types.GetSegmentPath(otdrThresholds.OtdrThreshold[i]), types.YChild{"OtdrThreshold", otdrThresholds.OtdrThreshold[i]})
    }
    otdrThresholds.EntityData.Leafs = types.NewOrderedMap()

    otdrThresholds.EntityData.YListKeys = []string {}

    return &(otdrThresholds.EntityData)
}

// HardwareModule_Node_Slot_OtdrThresholds_OtdrThreshold
// Otdr Threshold
type HardwareModule_Node_Slot_OtdrThresholds_OtdrThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port Id. The type is interface{} with range:
    // 0..4294967295.
    Port interface{}

    // Absolute Loss Threshold, supported values [1 ,300] in units of 0.1dB. The
    // type is interface{} with range: 1..300.
    LossAbsThreshold interface{}

    // Absolute Reflection Threshold, supported values [-500,0] in units of 0.1dB.
    // The type is interface{} with range: -500..0.
    ReflectionAbsThreshold interface{}

    // Absolute Orl Threshold, supported values [140,400] in units of 0.1dB. The
    // type is interface{} with range: 140..400.
    OrlAbsThreshold interface{}
}

func (otdrThreshold *HardwareModule_Node_Slot_OtdrThresholds_OtdrThreshold) GetEntityData() *types.CommonEntityData {
    otdrThreshold.EntityData.YFilter = otdrThreshold.YFilter
    otdrThreshold.EntityData.YangName = "otdr-threshold"
    otdrThreshold.EntityData.BundleName = "cisco_ios_xr"
    otdrThreshold.EntityData.ParentYangName = "otdr-thresholds"
    otdrThreshold.EntityData.SegmentPath = "otdr-threshold" + types.AddKeyToken(otdrThreshold.Port, "port")
    otdrThreshold.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/otdr-thresholds/" + otdrThreshold.EntityData.SegmentPath
    otdrThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrThreshold.EntityData.Children = types.NewOrderedMap()
    otdrThreshold.EntityData.Leafs = types.NewOrderedMap()
    otdrThreshold.EntityData.Leafs.Append("port", types.YLeaf{"Port", otdrThreshold.Port})
    otdrThreshold.EntityData.Leafs.Append("loss-abs-threshold", types.YLeaf{"LossAbsThreshold", otdrThreshold.LossAbsThreshold})
    otdrThreshold.EntityData.Leafs.Append("reflection-abs-threshold", types.YLeaf{"ReflectionAbsThreshold", otdrThreshold.ReflectionAbsThreshold})
    otdrThreshold.EntityData.Leafs.Append("orl-abs-threshold", types.YLeaf{"OrlAbsThreshold", otdrThreshold.OrlAbsThreshold})

    otdrThreshold.EntityData.YListKeys = []string {"Port"}

    return &(otdrThreshold.EntityData)
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

    // SpanLoss configuration. The type is bool.
    SpanLoss interface{}

    // RemoteNode Configuration.
    RemoteNode HardwareModule_Node_Slot_Amplifier_RemoteNode
}

func (amplifier *HardwareModule_Node_Slot_Amplifier) GetEntityData() *types.CommonEntityData {
    amplifier.EntityData.YFilter = amplifier.YFilter
    amplifier.EntityData.YangName = "amplifier"
    amplifier.EntityData.BundleName = "cisco_ios_xr"
    amplifier.EntityData.ParentYangName = "slot"
    amplifier.EntityData.SegmentPath = "amplifier"
    amplifier.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/" + amplifier.EntityData.SegmentPath
    amplifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifier.EntityData.Children = types.NewOrderedMap()
    amplifier.EntityData.Children.Append("remote-node", types.YChild{"RemoteNode", &amplifier.RemoteNode})
    amplifier.EntityData.Leafs = types.NewOrderedMap()
    amplifier.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", amplifier.NodeType})
    amplifier.EntityData.Leafs.Append("grid-mode", types.YLeaf{"GridMode", amplifier.GridMode})
    amplifier.EntityData.Leafs.Append("udc-vlan", types.YLeaf{"UdcVlan", amplifier.UdcVlan})
    amplifier.EntityData.Leafs.Append("span-loss", types.YLeaf{"SpanLoss", amplifier.SpanLoss})

    amplifier.EntityData.YListKeys = []string {}

    return &(amplifier.EntityData)
}

// HardwareModule_Node_Slot_Amplifier_RemoteNode
// RemoteNode Configuration
type HardwareModule_Node_Slot_Amplifier_RemoteNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of local host. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalIpAddress interface{}

    // IP address of remote host. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteIpAddress interface{}

    // Set remote Slot. The type is interface{} with range: 1..3.
    RemoteSlotId interface{}
}

func (remoteNode *HardwareModule_Node_Slot_Amplifier_RemoteNode) GetEntityData() *types.CommonEntityData {
    remoteNode.EntityData.YFilter = remoteNode.YFilter
    remoteNode.EntityData.YangName = "remote-node"
    remoteNode.EntityData.BundleName = "cisco_ios_xr"
    remoteNode.EntityData.ParentYangName = "amplifier"
    remoteNode.EntityData.SegmentPath = "remote-node"
    remoteNode.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/amplifier/" + remoteNode.EntityData.SegmentPath
    remoteNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNode.EntityData.Children = types.NewOrderedMap()
    remoteNode.EntityData.Leafs = types.NewOrderedMap()
    remoteNode.EntityData.Leafs.Append("local-ip-address", types.YLeaf{"LocalIpAddress", remoteNode.LocalIpAddress})
    remoteNode.EntityData.Leafs.Append("remote-ip-address", types.YLeaf{"RemoteIpAddress", remoteNode.RemoteIpAddress})
    remoteNode.EntityData.Leafs.Append("remote-slot-id", types.YLeaf{"RemoteSlotId", remoteNode.RemoteSlotId})

    remoteNode.EntityData.YListKeys = []string {}

    return &(remoteNode.EntityData)
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
    psm.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-ots-cfg:hardware-module/node/slot/" + psm.EntityData.SegmentPath
    psm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    psm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    psm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    psm.EntityData.Children = types.NewOrderedMap()
    psm.EntityData.Leafs = types.NewOrderedMap()
    psm.EntityData.Leafs.Append("mono-dir", types.YLeaf{"MonoDir", psm.MonoDir})
    psm.EntityData.Leafs.Append("auto-threshold", types.YLeaf{"AutoThreshold", psm.AutoThreshold})
    psm.EntityData.Leafs.Append("path-protection", types.YLeaf{"PathProtection", psm.PathProtection})
    psm.EntityData.Leafs.Append("section-protection", types.YLeaf{"SectionProtection", psm.SectionProtection})
    psm.EntityData.Leafs.Append("lockout-from", types.YLeaf{"LockoutFrom", psm.LockoutFrom})

    psm.EntityData.YListKeys = []string {}

    return &(psm.EntityData)
}

