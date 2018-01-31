// This module contains a collection of YANG definitions
// for Cisco IOS-XR syncc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   timing-controller: Timing controller operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package syncc_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package syncc_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-syncc-oper timing-controller}", reflect.TypeOf(TimingController{}))
    ydk.RegisterEntity("Cisco-IOS-XR-syncc-oper:timing-controller", reflect.TypeOf(TimingController{}))
}

// Smode1 represents First mode type
type Smode1 string

const (
    // Extended Superframe(ESF)
    Smode1_extended_super_frame Smode1 = "extended-super-frame"

    // D4 channel unit
    Smode1_d4 Smode1 = "d4"

    // Non CRC 4 mode
    Smode1_non_crc4 Smode1 = "non-crc4"

    // CRC 4
    Smode1_crc4 Smode1 = "crc4"

    // No mode is selected
    Smode1_submode1_none Smode1 = "submode1-none"
)

// Smode2 represents Second mode type
type Smode2 string

const (
    // AMI
    Smode2_ami_mode Smode2 = "ami-mode"

    // B8ZS submode
    Smode2_b8zs Smode2 = "b8zs"

    // HDB3 submode
    Smode2_hdb3 Smode2 = "hdb3"

    // No mode is selected
    Smode2_submode2_none Smode2 = "submode2-none"
)

// ClockModes represents Different clock modes
type ClockModes string

const (
    // T1 mode
    ClockModes_t1 ClockModes = "t1"

    // E1 mode
    ClockModes_e1 ClockModes = "e1"

    // 2048M mode
    ClockModes_two_m ClockModes = "two-m"

    // 64kCC input
    ClockModes_input64k ClockModes = "input64k"

    // 6312M output
    ClockModes_output6m ClockModes = "output6m"

    // Universal transport interface(UTI)
    ClockModes_uti ClockModes = "uti"

    // No mode is selected
    ClockModes_none ClockModes = "none"
)

// QlOption1 represents Quality level options
type QlOption1 string

const (
    // No value
    QlOption1_quality_level_none QlOption1 = "quality-level-none"

    // Option 1
    QlOption1_o1 QlOption1 = "o1"

    // Option 2 Gen 1
    QlOption1_o2_g1 QlOption1 = "o2-g1"

    // Option 2 Gen 2
    QlOption1_o2_g2 QlOption1 = "o2-g2"
)

// Direct represents Direction status
type Direct string

const (
    // Receive or transmit
    Direct_receive_transmit Direct = "receive-transmit"

    // Transmit
    Direct_transmit Direct = "transmit"

    // Receive
    Direct_receive Direct = "receive"
)

// SourceStateName represents Syncc source state name
type SourceStateName string

const (
    // Invalid
    SourceStateName_not_valid SourceStateName = "not-valid"

    // Unqualified state
    SourceStateName_unqualified SourceStateName = "unqualified"

    // Available state
    SourceStateName_available SourceStateName = "available"

    // Failed state
    SourceStateName_failed SourceStateName = "failed"

    // Unmonitored state
    SourceStateName_unmonitored SourceStateName = "unmonitored"

    // Error state
    SourceStateName_error SourceStateName = "error"
)

// Source represents Syncc source type
type Source string

const (
    // Invalid state
    Source_invalid Source = "invalid"

    // Ethernet interface 
    Source_ethernet_line_interface Source = "ethernet-line-interface"

    // SONET interface 
    Source_sonet_line_interface Source = "sonet-line-interface"

    // Clock interface state
    Source_clock_interface Source = "clock-interface"

    // Internal state
    Source_internal Source = "internal"
)

// NodeState represents Different modes of a node
type NodeState string

const (
    // Node in active mode
    NodeState_active NodeState = "active"

    // Node in standby mode
    NodeState_standby NodeState = "standby"
)

// InterfaceState represents Interface state
type InterfaceState string

const (
    // Up state
    InterfaceState_up InterfaceState = "up"

    // Down state
    InterfaceState_down InterfaceState = "down"

    // Admin down state
    InterfaceState_admin_down InterfaceState = "admin-down"
)

// SynccStates represents Different syncc states
type SynccStates string

const (
    // Initial  state
    SynccStates_initializing SynccStates = "initializing"

    // Running  state
    SynccStates_running SynccStates = "running"

    // Normal  state
    SynccStates_normal SynccStates = "normal"

    // Shutdown  state
    SynccStates_shutdown SynccStates = "shutdown"
)

// TimingController
// Timing controller operational data
type TimingController struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes applicable to timing controller.
    Nodes TimingController_Nodes
}

func (timingController *TimingController) GetFilter() yfilter.YFilter { return timingController.YFilter }

func (timingController *TimingController) SetFilter(yf yfilter.YFilter) { timingController.YFilter = yf }

func (timingController *TimingController) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (timingController *TimingController) GetSegmentPath() string {
    return "Cisco-IOS-XR-syncc-oper:timing-controller"
}

func (timingController *TimingController) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &timingController.Nodes
    }
    return nil
}

func (timingController *TimingController) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &timingController.Nodes
    return children
}

func (timingController *TimingController) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timingController *TimingController) GetBundleName() string { return "cisco_ios_xr" }

func (timingController *TimingController) GetYangName() string { return "timing-controller" }

func (timingController *TimingController) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timingController *TimingController) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timingController *TimingController) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timingController *TimingController) SetParent(parent types.Entity) { timingController.parent = parent }

func (timingController *TimingController) GetParent() types.Entity { return timingController.parent }

func (timingController *TimingController) GetParentYangName() string { return "Cisco-IOS-XR-syncc-oper" }

// TimingController_Nodes
// List of nodes applicable to timing controller
type TimingController_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syncc operational data for a single node. The type is slice of
    // TimingController_Nodes_Node.
    Node []TimingController_Nodes_Node
}

func (nodes *TimingController_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *TimingController_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *TimingController_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *TimingController_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *TimingController_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *TimingController_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *TimingController_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *TimingController_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *TimingController_Nodes) GetYangName() string { return "nodes" }

func (nodes *TimingController_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *TimingController_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *TimingController_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *TimingController_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *TimingController_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *TimingController_Nodes) GetParentYangName() string { return "timing-controller" }

// TimingController_Nodes_Node
// Syncc operational data for a single node
type TimingController_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node Name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Syncc state for a node.
    State TimingController_Nodes_Node_State

    // Syncc clock information for a node.
    Clock TimingController_Nodes_Node_Clock

    // Syncc timing information for a node.
    TimingSource TimingController_Nodes_Node_TimingSource
}

func (node *TimingController_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *TimingController_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *TimingController_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "state" { return "State" }
    if yname == "clock" { return "Clock" }
    if yname == "timing-source" { return "TimingSource" }
    return ""
}

func (node *TimingController_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *TimingController_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        return &node.State
    }
    if childYangName == "clock" {
        return &node.Clock
    }
    if childYangName == "timing-source" {
        return &node.TimingSource
    }
    return nil
}

func (node *TimingController_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["state"] = &node.State
    children["clock"] = &node.Clock
    children["timing-source"] = &node.TimingSource
    return children
}

func (node *TimingController_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *TimingController_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *TimingController_Nodes_Node) GetYangName() string { return "node" }

func (node *TimingController_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *TimingController_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *TimingController_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *TimingController_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *TimingController_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *TimingController_Nodes_Node) GetParentYangName() string { return "nodes" }

// TimingController_Nodes_Node_State
// Syncc state for a node
type TimingController_Nodes_Node_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of syncc states. The type is slice of
    // TimingController_Nodes_Node_State_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_State_SynccInstance
}

func (state *TimingController_Nodes_Node_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *TimingController_Nodes_Node_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *TimingController_Nodes_Node_State) GetGoName(yname string) string {
    if yname == "syncc-instance" { return "SynccInstance" }
    return ""
}

func (state *TimingController_Nodes_Node_State) GetSegmentPath() string {
    return "state"
}

func (state *TimingController_Nodes_Node_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "syncc-instance" {
        for _, c := range state.SynccInstance {
            if state.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_State_SynccInstance{}
        state.SynccInstance = append(state.SynccInstance, child)
        return &state.SynccInstance[len(state.SynccInstance)-1]
    }
    return nil
}

func (state *TimingController_Nodes_Node_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range state.SynccInstance {
        children[state.SynccInstance[i].GetSegmentPath()] = &state.SynccInstance[i]
    }
    return children
}

func (state *TimingController_Nodes_Node_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *TimingController_Nodes_Node_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *TimingController_Nodes_Node_State) GetYangName() string { return "state" }

func (state *TimingController_Nodes_Node_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *TimingController_Nodes_Node_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *TimingController_Nodes_Node_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *TimingController_Nodes_Node_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *TimingController_Nodes_Node_State) GetParent() types.Entity { return state.parent }

func (state *TimingController_Nodes_Node_State) GetParentYangName() string { return "node" }

// TimingController_Nodes_Node_State_SynccInstance
// List of syncc states
type TimingController_Nodes_Node_State_SynccInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syncc controller state. The type is SynccStates.
    ControllerState interface{}

    // Status of syncc node mode. The type is NodeState.
    SynccNodeState interface{}

    // Verbose level number. The type is interface{} with range: 0..4294967295.
    VerboseLevel interface{}

    // Initial count number. The type is interface{} with range: 0..4294967295.
    InitialCount interface{}

    // Shutdown count number. The type is interface{} with range: 0..4294967295.
    ShutdownCount interface{}

    // Set the value of input count. The type is interface{} with range:
    // 0..4294967295.
    SetInputCount interface{}

    // Set the value of display count. The type is interface{} with range:
    // 0..4294967295.
    SetCapabilityCount interface{}

    // Clock count number. The type is interface{} with range: 0..4294967295.
    GetClockCount interface{}

    // Set clock count. The type is interface{} with range: 0..4294967295.
    SetClockOutCount interface{}

    // Sync enable count number. The type is interface{} with range:
    // 0..4294967295.
    SyncEnableCount interface{}

    // Sync disable count number. The type is interface{} with range:
    // 0..4294967295.
    SyncDisableCount interface{}

    // Interface capability count. The type is interface{} with range:
    // 0..4294967295.
    CapabilityCount interface{}

    // Value of quality level count. The type is interface{} with range:
    // 0..4294967295.
    SetQualityLevelCount interface{}

    // Selects proper input result notification. The type is interface{} with
    // range: 0..4294967295.
    InputNotification interface{}

    // Value of interface capability notification. The type is interface{} with
    // range: 0..4294967295.
    CapabilityNotification interface{}

    // Notification of source status. The type is interface{} with range:
    // 0..4294967295.
    StatusNotification interface{}

    // Value of resync notification. The type is interface{} with range:
    // 0..4294967295.
    ResyncNotification interface{}
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetFilter() yfilter.YFilter { return synccInstance.YFilter }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) SetFilter(yf yfilter.YFilter) { synccInstance.YFilter = yf }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetGoName(yname string) string {
    if yname == "controller-state" { return "ControllerState" }
    if yname == "syncc-node-state" { return "SynccNodeState" }
    if yname == "verbose-level" { return "VerboseLevel" }
    if yname == "initial-count" { return "InitialCount" }
    if yname == "shutdown-count" { return "ShutdownCount" }
    if yname == "set-input-count" { return "SetInputCount" }
    if yname == "set-capability-count" { return "SetCapabilityCount" }
    if yname == "get-clock-count" { return "GetClockCount" }
    if yname == "set-clock-out-count" { return "SetClockOutCount" }
    if yname == "sync-enable-count" { return "SyncEnableCount" }
    if yname == "sync-disable-count" { return "SyncDisableCount" }
    if yname == "capability-count" { return "CapabilityCount" }
    if yname == "set-quality-level-count" { return "SetQualityLevelCount" }
    if yname == "input-notification" { return "InputNotification" }
    if yname == "capability-notification" { return "CapabilityNotification" }
    if yname == "status-notification" { return "StatusNotification" }
    if yname == "resync-notification" { return "ResyncNotification" }
    return ""
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetSegmentPath() string {
    return "syncc-instance"
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["controller-state"] = synccInstance.ControllerState
    leafs["syncc-node-state"] = synccInstance.SynccNodeState
    leafs["verbose-level"] = synccInstance.VerboseLevel
    leafs["initial-count"] = synccInstance.InitialCount
    leafs["shutdown-count"] = synccInstance.ShutdownCount
    leafs["set-input-count"] = synccInstance.SetInputCount
    leafs["set-capability-count"] = synccInstance.SetCapabilityCount
    leafs["get-clock-count"] = synccInstance.GetClockCount
    leafs["set-clock-out-count"] = synccInstance.SetClockOutCount
    leafs["sync-enable-count"] = synccInstance.SyncEnableCount
    leafs["sync-disable-count"] = synccInstance.SyncDisableCount
    leafs["capability-count"] = synccInstance.CapabilityCount
    leafs["set-quality-level-count"] = synccInstance.SetQualityLevelCount
    leafs["input-notification"] = synccInstance.InputNotification
    leafs["capability-notification"] = synccInstance.CapabilityNotification
    leafs["status-notification"] = synccInstance.StatusNotification
    leafs["resync-notification"] = synccInstance.ResyncNotification
    return leafs
}

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetBundleName() string { return "cisco_ios_xr" }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetYangName() string { return "syncc-instance" }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) SetParent(parent types.Entity) { synccInstance.parent = parent }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetParent() types.Entity { return synccInstance.parent }

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetParentYangName() string { return "state" }

// TimingController_Nodes_Node_Clock
// Syncc clock information for a node
type TimingController_Nodes_Node_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of syncc clock information . The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_Clock_SynccInstance
}

func (clock *TimingController_Nodes_Node_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *TimingController_Nodes_Node_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *TimingController_Nodes_Node_Clock) GetGoName(yname string) string {
    if yname == "syncc-instance" { return "SynccInstance" }
    return ""
}

func (clock *TimingController_Nodes_Node_Clock) GetSegmentPath() string {
    return "clock"
}

func (clock *TimingController_Nodes_Node_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "syncc-instance" {
        for _, c := range clock.SynccInstance {
            if clock.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_Clock_SynccInstance{}
        clock.SynccInstance = append(clock.SynccInstance, child)
        return &clock.SynccInstance[len(clock.SynccInstance)-1]
    }
    return nil
}

func (clock *TimingController_Nodes_Node_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clock.SynccInstance {
        children[clock.SynccInstance[i].GetSegmentPath()] = &clock.SynccInstance[i]
    }
    return children
}

func (clock *TimingController_Nodes_Node_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clock *TimingController_Nodes_Node_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *TimingController_Nodes_Node_Clock) GetYangName() string { return "clock" }

func (clock *TimingController_Nodes_Node_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *TimingController_Nodes_Node_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *TimingController_Nodes_Node_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *TimingController_Nodes_Node_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *TimingController_Nodes_Node_Clock) GetParent() types.Entity { return clock.parent }

func (clock *TimingController_Nodes_Node_Clock) GetParentYangName() string { return "node" }

// TimingController_Nodes_Node_Clock_SynccInstance
// List of syncc clock information 
type TimingController_Nodes_Node_Clock_SynccInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clock table for an RP. The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance_Clock.
    Clock []TimingController_Nodes_Node_Clock_SynccInstance_Clock
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetFilter() yfilter.YFilter { return synccInstance.YFilter }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) SetFilter(yf yfilter.YFilter) { synccInstance.YFilter = yf }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    return ""
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetSegmentPath() string {
    return "syncc-instance"
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        for _, c := range synccInstance.Clock {
            if synccInstance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_Clock_SynccInstance_Clock{}
        synccInstance.Clock = append(synccInstance.Clock, child)
        return &synccInstance.Clock[len(synccInstance.Clock)-1]
    }
    return nil
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range synccInstance.Clock {
        children[synccInstance.Clock[i].GetSegmentPath()] = &synccInstance.Clock[i]
    }
    return children
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetBundleName() string { return "cisco_ios_xr" }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetYangName() string { return "syncc-instance" }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) SetParent(parent types.Entity) { synccInstance.parent = parent }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetParent() types.Entity { return synccInstance.parent }

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetParentYangName() string { return "clock" }

// TimingController_Nodes_Node_Clock_SynccInstance_Clock
// Clock table for an RP
type TimingController_Nodes_Node_Clock_SynccInstance_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // True if clock is configured for port 0. The type is bool.
    IsConfiguredPort0 interface{}

    // True if clock is configured for port 1. The type is bool.
    IsConfiguredPort1 interface{}

    // True if clock is configured for port 2. The type is bool.
    IsConfiguredPort2 interface{}

    // True if clock is configured for port 3. The type is bool.
    IsConfiguredPort3 interface{}

    // Clock setting mode for port 0. The type is ClockModes.
    ModePort0 interface{}

    // Clock setting mode for port 1. The type is ClockModes.
    ModePort1 interface{}

    // Clock setting mode for port 2. The type is ClockModes.
    ModePort2 interface{}

    // Clock setting mode for port 3. The type is ClockModes.
    ModePort3 interface{}

    // First submode for port 0. The type is Smode1.
    Submode1Port0 interface{}

    // First submode for port 1. The type is Smode1.
    Submode1Port1 interface{}

    // First submode for port 2. The type is Smode1.
    Submode1Port2 interface{}

    // First submode for port 3. The type is Smode1.
    Submode1Port3 interface{}

    // Second submode for port 0. The type is Smode2.
    Submode2Port0 interface{}

    // Second submode for port 1. The type is Smode2.
    Submode2Port1 interface{}

    // Second submode for port 2. The type is Smode2.
    Submode2Port2 interface{}

    // Second submode for port 3. The type is Smode2.
    Submode2Port3 interface{}

    // Third submode for port 0. The type is interface{} with range:
    // 0..4294967295.
    Submode3Port0 interface{}

    // Third submode for port 1. The type is interface{} with range:
    // 0..4294967295.
    Submode3Port1 interface{}

    // Third submode for port 2. The type is interface{} with range:
    // 0..4294967295.
    Submode3Port2 interface{}

    // Third submode for port 3. The type is interface{} with range:
    // 0..4294967295.
    Submode3Port3 interface{}

    // Configure disable value for port 0. The type is interface{} with range:
    // 0..4294967295.
    ShutdownPort0 interface{}

    // Configure disable value for port 1. The type is interface{} with range:
    // 0..4294967295.
    ShutdownPort1 interface{}

    // Configure disable value for port 2. The type is interface{} with range:
    // 0..4294967295.
    ShutdownPort2 interface{}

    // Configure disable value for port 3. The type is interface{} with range:
    // 0..4294967295.
    ShutdownPort3 interface{}

    // Direction of interface for port 0. The type is Direct.
    DirectionPort0 interface{}

    // Direction of interface for port 1. The type is Direct.
    DirectionPort1 interface{}

    // Direction of interface for port 2. The type is Direct.
    DirectionPort2 interface{}

    // Direction of interface for port 3. The type is Direct.
    DirectionPort3 interface{}

    // Baudrate for port 0. The type is interface{} with range: 0..4294967295.
    BaudratePort0 interface{}

    // Baudrate for port 1. The type is interface{} with range: 0..4294967295.
    BaudratePort1 interface{}

    // Baudrate for port 2. The type is interface{} with range: 0..4294967295.
    BaudratePort2 interface{}

    // Baudrate for port 3. The type is interface{} with range: 0..4294967295.
    BaudratePort3 interface{}

    // Quality Level option of port 0. The type is QlOption1.
    QualityOptionPort0 interface{}

    // Quality Level option of the port 1. The type is QlOption1.
    QualityOptionPort1 interface{}

    // Quality Level option of the port 2. The type is QlOption1.
    QualityOptionPort2 interface{}

    // Quality Level option of the port 3. The type is QlOption1.
    QualityOptionPort3 interface{}

    // Transmit SSM for port 0. The type is interface{} with range: 0..4294967295.
    TransmitSsmPort0 interface{}

    // Transmit SSM for port 1. The type is interface{} with range: 0..4294967295.
    TransmitSsmPort1 interface{}

    // Transmit SSM for port 2. The type is interface{} with range: 0..4294967295.
    TransmitSsmPort2 interface{}

    // Transmit SSM for port 3. The type is interface{} with range: 0..4294967295.
    TransmitSsmPort3 interface{}

    // Receive SSM for port 0. The type is interface{} with range: 0..4294967295.
    RecieveSsmPort0 interface{}

    // Receive SSM for port 1. The type is interface{} with range: 0..4294967295.
    RecieveSsmPort1 interface{}

    // Receive SSM for port 2. The type is interface{} with range: 0..4294967295.
    RecieveSsmPort2 interface{}

    // Receive SSM for port 3. The type is interface{} with range: 0..4294967295.
    RecieveSsmPort3 interface{}

    // Interface state for port 0. The type is InterfaceState.
    InterfaceStatePort0 interface{}

    // Interface state for port 1. The type is InterfaceState.
    InterfaceStatePort1 interface{}

    // Interface state for port 2. The type is InterfaceState.
    InterfaceStatePort2 interface{}

    // Interface state for port 3. The type is InterfaceState.
    InterfaceStatePort3 interface{}
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetGoName(yname string) string {
    if yname == "is-configured-port0" { return "IsConfiguredPort0" }
    if yname == "is-configured-port1" { return "IsConfiguredPort1" }
    if yname == "is-configured-port2" { return "IsConfiguredPort2" }
    if yname == "is-configured-port3" { return "IsConfiguredPort3" }
    if yname == "mode-port0" { return "ModePort0" }
    if yname == "mode-port1" { return "ModePort1" }
    if yname == "mode-port2" { return "ModePort2" }
    if yname == "mode-port3" { return "ModePort3" }
    if yname == "submode1-port0" { return "Submode1Port0" }
    if yname == "submode1-port1" { return "Submode1Port1" }
    if yname == "submode1-port2" { return "Submode1Port2" }
    if yname == "submode1-port3" { return "Submode1Port3" }
    if yname == "submode2-port0" { return "Submode2Port0" }
    if yname == "submode2-port1" { return "Submode2Port1" }
    if yname == "submode2-port2" { return "Submode2Port2" }
    if yname == "submode2-port3" { return "Submode2Port3" }
    if yname == "submode3-port0" { return "Submode3Port0" }
    if yname == "submode3-port1" { return "Submode3Port1" }
    if yname == "submode3-port2" { return "Submode3Port2" }
    if yname == "submode3-port3" { return "Submode3Port3" }
    if yname == "shutdown-port0" { return "ShutdownPort0" }
    if yname == "shutdown-port1" { return "ShutdownPort1" }
    if yname == "shutdown-port2" { return "ShutdownPort2" }
    if yname == "shutdown-port3" { return "ShutdownPort3" }
    if yname == "direction-port0" { return "DirectionPort0" }
    if yname == "direction-port1" { return "DirectionPort1" }
    if yname == "direction-port2" { return "DirectionPort2" }
    if yname == "direction-port3" { return "DirectionPort3" }
    if yname == "baudrate-port0" { return "BaudratePort0" }
    if yname == "baudrate-port1" { return "BaudratePort1" }
    if yname == "baudrate-port2" { return "BaudratePort2" }
    if yname == "baudrate-port3" { return "BaudratePort3" }
    if yname == "quality-option-port0" { return "QualityOptionPort0" }
    if yname == "quality-option-port1" { return "QualityOptionPort1" }
    if yname == "quality-option-port2" { return "QualityOptionPort2" }
    if yname == "quality-option-port3" { return "QualityOptionPort3" }
    if yname == "transmit-ssm-port0" { return "TransmitSsmPort0" }
    if yname == "transmit-ssm-port1" { return "TransmitSsmPort1" }
    if yname == "transmit-ssm-port2" { return "TransmitSsmPort2" }
    if yname == "transmit-ssm-port3" { return "TransmitSsmPort3" }
    if yname == "recieve-ssm-port0" { return "RecieveSsmPort0" }
    if yname == "recieve-ssm-port1" { return "RecieveSsmPort1" }
    if yname == "recieve-ssm-port2" { return "RecieveSsmPort2" }
    if yname == "recieve-ssm-port3" { return "RecieveSsmPort3" }
    if yname == "interface-state-port0" { return "InterfaceStatePort0" }
    if yname == "interface-state-port1" { return "InterfaceStatePort1" }
    if yname == "interface-state-port2" { return "InterfaceStatePort2" }
    if yname == "interface-state-port3" { return "InterfaceStatePort3" }
    return ""
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetSegmentPath() string {
    return "clock"
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-configured-port0"] = clock.IsConfiguredPort0
    leafs["is-configured-port1"] = clock.IsConfiguredPort1
    leafs["is-configured-port2"] = clock.IsConfiguredPort2
    leafs["is-configured-port3"] = clock.IsConfiguredPort3
    leafs["mode-port0"] = clock.ModePort0
    leafs["mode-port1"] = clock.ModePort1
    leafs["mode-port2"] = clock.ModePort2
    leafs["mode-port3"] = clock.ModePort3
    leafs["submode1-port0"] = clock.Submode1Port0
    leafs["submode1-port1"] = clock.Submode1Port1
    leafs["submode1-port2"] = clock.Submode1Port2
    leafs["submode1-port3"] = clock.Submode1Port3
    leafs["submode2-port0"] = clock.Submode2Port0
    leafs["submode2-port1"] = clock.Submode2Port1
    leafs["submode2-port2"] = clock.Submode2Port2
    leafs["submode2-port3"] = clock.Submode2Port3
    leafs["submode3-port0"] = clock.Submode3Port0
    leafs["submode3-port1"] = clock.Submode3Port1
    leafs["submode3-port2"] = clock.Submode3Port2
    leafs["submode3-port3"] = clock.Submode3Port3
    leafs["shutdown-port0"] = clock.ShutdownPort0
    leafs["shutdown-port1"] = clock.ShutdownPort1
    leafs["shutdown-port2"] = clock.ShutdownPort2
    leafs["shutdown-port3"] = clock.ShutdownPort3
    leafs["direction-port0"] = clock.DirectionPort0
    leafs["direction-port1"] = clock.DirectionPort1
    leafs["direction-port2"] = clock.DirectionPort2
    leafs["direction-port3"] = clock.DirectionPort3
    leafs["baudrate-port0"] = clock.BaudratePort0
    leafs["baudrate-port1"] = clock.BaudratePort1
    leafs["baudrate-port2"] = clock.BaudratePort2
    leafs["baudrate-port3"] = clock.BaudratePort3
    leafs["quality-option-port0"] = clock.QualityOptionPort0
    leafs["quality-option-port1"] = clock.QualityOptionPort1
    leafs["quality-option-port2"] = clock.QualityOptionPort2
    leafs["quality-option-port3"] = clock.QualityOptionPort3
    leafs["transmit-ssm-port0"] = clock.TransmitSsmPort0
    leafs["transmit-ssm-port1"] = clock.TransmitSsmPort1
    leafs["transmit-ssm-port2"] = clock.TransmitSsmPort2
    leafs["transmit-ssm-port3"] = clock.TransmitSsmPort3
    leafs["recieve-ssm-port0"] = clock.RecieveSsmPort0
    leafs["recieve-ssm-port1"] = clock.RecieveSsmPort1
    leafs["recieve-ssm-port2"] = clock.RecieveSsmPort2
    leafs["recieve-ssm-port3"] = clock.RecieveSsmPort3
    leafs["interface-state-port0"] = clock.InterfaceStatePort0
    leafs["interface-state-port1"] = clock.InterfaceStatePort1
    leafs["interface-state-port2"] = clock.InterfaceStatePort2
    leafs["interface-state-port3"] = clock.InterfaceStatePort3
    return leafs
}

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetYangName() string { return "clock" }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetParent() types.Entity { return clock.parent }

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetParentYangName() string { return "syncc-instance" }

// TimingController_Nodes_Node_TimingSource
// Syncc timing information for a node
type TimingController_Nodes_Node_TimingSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of syncc timing table information. The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_TimingSource_SynccInstance
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetFilter() yfilter.YFilter { return timingSource.YFilter }

func (timingSource *TimingController_Nodes_Node_TimingSource) SetFilter(yf yfilter.YFilter) { timingSource.YFilter = yf }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetGoName(yname string) string {
    if yname == "syncc-instance" { return "SynccInstance" }
    return ""
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetSegmentPath() string {
    return "timing-source"
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "syncc-instance" {
        for _, c := range timingSource.SynccInstance {
            if timingSource.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_TimingSource_SynccInstance{}
        timingSource.SynccInstance = append(timingSource.SynccInstance, child)
        return &timingSource.SynccInstance[len(timingSource.SynccInstance)-1]
    }
    return nil
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range timingSource.SynccInstance {
        children[timingSource.SynccInstance[i].GetSegmentPath()] = &timingSource.SynccInstance[i]
    }
    return children
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetBundleName() string { return "cisco_ios_xr" }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetYangName() string { return "timing-source" }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timingSource *TimingController_Nodes_Node_TimingSource) SetParent(parent types.Entity) { timingSource.parent = parent }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetParent() types.Entity { return timingSource.parent }

func (timingSource *TimingController_Nodes_Node_TimingSource) GetParentYangName() string { return "node" }

// TimingController_Nodes_Node_TimingSource_SynccInstance
// List of syncc timing table information
type TimingController_Nodes_Node_TimingSource_SynccInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Scheduling PLL T0 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0.
    TimingStatusT0 []TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0

    // Scheduling PLL T4 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4.
    TimingStatusT4 []TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4

    // Scheduling PLL 1588 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588.
    TimingStatus1588 []TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetFilter() yfilter.YFilter { return synccInstance.YFilter }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) SetFilter(yf yfilter.YFilter) { synccInstance.YFilter = yf }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetGoName(yname string) string {
    if yname == "timing-status-t0" { return "TimingStatusT0" }
    if yname == "timing-status-t4" { return "TimingStatusT4" }
    if yname == "timing-status1588" { return "TimingStatus1588" }
    return ""
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetSegmentPath() string {
    return "syncc-instance"
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timing-status-t0" {
        for _, c := range synccInstance.TimingStatusT0 {
            if synccInstance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0{}
        synccInstance.TimingStatusT0 = append(synccInstance.TimingStatusT0, child)
        return &synccInstance.TimingStatusT0[len(synccInstance.TimingStatusT0)-1]
    }
    if childYangName == "timing-status-t4" {
        for _, c := range synccInstance.TimingStatusT4 {
            if synccInstance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4{}
        synccInstance.TimingStatusT4 = append(synccInstance.TimingStatusT4, child)
        return &synccInstance.TimingStatusT4[len(synccInstance.TimingStatusT4)-1]
    }
    if childYangName == "timing-status1588" {
        for _, c := range synccInstance.TimingStatus1588 {
            if synccInstance.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588{}
        synccInstance.TimingStatus1588 = append(synccInstance.TimingStatus1588, child)
        return &synccInstance.TimingStatus1588[len(synccInstance.TimingStatus1588)-1]
    }
    return nil
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range synccInstance.TimingStatusT0 {
        children[synccInstance.TimingStatusT0[i].GetSegmentPath()] = &synccInstance.TimingStatusT0[i]
    }
    for i := range synccInstance.TimingStatusT4 {
        children[synccInstance.TimingStatusT4[i].GetSegmentPath()] = &synccInstance.TimingStatusT4[i]
    }
    for i := range synccInstance.TimingStatus1588 {
        children[synccInstance.TimingStatus1588[i].GetSegmentPath()] = &synccInstance.TimingStatus1588[i]
    }
    return children
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetBundleName() string { return "cisco_ios_xr" }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetYangName() string { return "syncc-instance" }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) SetParent(parent types.Entity) { synccInstance.parent = parent }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetParent() types.Entity { return synccInstance.parent }

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetParentYangName() string { return "timing-source" }

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0
// Scheduling PLL T0 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Input number. The type is interface{} with range: 0..255.
    Input interface{}

    // Slot number. The type is interface{} with range: 0..255.
    Slot interface{}

    // Port number. The type is interface{} with range: 0..255.
    Port interface{}

    // Status of syncc source type. The type is Source.
    ClockSource interface{}

    // Rank of sync timing source table. The type is interface{} with range:
    // 0..255.
    Rank interface{}

    // Quality level option. The type is interface{} with range: 0..255.
    QualityLevelOption interface{}

    // Quality level value. The type is interface{} with range: 0..255.
    QualityLevelValue interface{}

    // User priority of sync timing source table. The type is interface{} with
    // range: 0..255.
    UserPriority interface{}

    // Status of clock state. The type is SourceStateName.
    ClockState interface{}

    // True if selected. The type is bool.
    IsSelect interface{}
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetFilter() yfilter.YFilter { return timingStatusT0.YFilter }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) SetFilter(yf yfilter.YFilter) { timingStatusT0.YFilter = yf }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "slot" { return "Slot" }
    if yname == "port" { return "Port" }
    if yname == "clock-source" { return "ClockSource" }
    if yname == "rank" { return "Rank" }
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "quality-level-value" { return "QualityLevelValue" }
    if yname == "user-priority" { return "UserPriority" }
    if yname == "clock-state" { return "ClockState" }
    if yname == "is-select" { return "IsSelect" }
    return ""
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetSegmentPath() string {
    return "timing-status-t0"
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = timingStatusT0.Input
    leafs["slot"] = timingStatusT0.Slot
    leafs["port"] = timingStatusT0.Port
    leafs["clock-source"] = timingStatusT0.ClockSource
    leafs["rank"] = timingStatusT0.Rank
    leafs["quality-level-option"] = timingStatusT0.QualityLevelOption
    leafs["quality-level-value"] = timingStatusT0.QualityLevelValue
    leafs["user-priority"] = timingStatusT0.UserPriority
    leafs["clock-state"] = timingStatusT0.ClockState
    leafs["is-select"] = timingStatusT0.IsSelect
    return leafs
}

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetBundleName() string { return "cisco_ios_xr" }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetYangName() string { return "timing-status-t0" }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) SetParent(parent types.Entity) { timingStatusT0.parent = parent }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetParent() types.Entity { return timingStatusT0.parent }

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetParentYangName() string { return "syncc-instance" }

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4
// Scheduling PLL T4 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Input number. The type is interface{} with range: 0..255.
    Input interface{}

    // Slot number. The type is interface{} with range: 0..255.
    Slot interface{}

    // Port number. The type is interface{} with range: 0..255.
    Port interface{}

    // Status of syncc source type. The type is Source.
    ClockSource interface{}

    // Rank of sync timing source table. The type is interface{} with range:
    // 0..255.
    Rank interface{}

    // Quality level option. The type is interface{} with range: 0..255.
    QualityLevelOption interface{}

    // Quality level value. The type is interface{} with range: 0..255.
    QualityLevelValue interface{}

    // User priority of sync timing source table. The type is interface{} with
    // range: 0..255.
    UserPriority interface{}

    // Status of clock state. The type is SourceStateName.
    ClockState interface{}

    // True if selected. The type is bool.
    IsSelect interface{}
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetFilter() yfilter.YFilter { return timingStatusT4.YFilter }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) SetFilter(yf yfilter.YFilter) { timingStatusT4.YFilter = yf }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "slot" { return "Slot" }
    if yname == "port" { return "Port" }
    if yname == "clock-source" { return "ClockSource" }
    if yname == "rank" { return "Rank" }
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "quality-level-value" { return "QualityLevelValue" }
    if yname == "user-priority" { return "UserPriority" }
    if yname == "clock-state" { return "ClockState" }
    if yname == "is-select" { return "IsSelect" }
    return ""
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetSegmentPath() string {
    return "timing-status-t4"
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = timingStatusT4.Input
    leafs["slot"] = timingStatusT4.Slot
    leafs["port"] = timingStatusT4.Port
    leafs["clock-source"] = timingStatusT4.ClockSource
    leafs["rank"] = timingStatusT4.Rank
    leafs["quality-level-option"] = timingStatusT4.QualityLevelOption
    leafs["quality-level-value"] = timingStatusT4.QualityLevelValue
    leafs["user-priority"] = timingStatusT4.UserPriority
    leafs["clock-state"] = timingStatusT4.ClockState
    leafs["is-select"] = timingStatusT4.IsSelect
    return leafs
}

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetBundleName() string { return "cisco_ios_xr" }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetYangName() string { return "timing-status-t4" }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) SetParent(parent types.Entity) { timingStatusT4.parent = parent }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetParent() types.Entity { return timingStatusT4.parent }

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetParentYangName() string { return "syncc-instance" }

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588
// Scheduling PLL 1588 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Input number. The type is interface{} with range: 0..255.
    Input interface{}

    // Slot number. The type is interface{} with range: 0..255.
    Slot interface{}

    // Port number. The type is interface{} with range: 0..255.
    Port interface{}

    // Status of syncc source type. The type is Source.
    ClockSource interface{}

    // Rank of sync timing source table. The type is interface{} with range:
    // 0..255.
    Rank interface{}

    // Quality level option. The type is interface{} with range: 0..255.
    QualityLevelOption interface{}

    // Quality level value. The type is interface{} with range: 0..255.
    QualityLevelValue interface{}

    // User priority of sync timing source table. The type is interface{} with
    // range: 0..255.
    UserPriority interface{}

    // Status of clock state. The type is SourceStateName.
    ClockState interface{}

    // True if selected. The type is bool.
    IsSelect interface{}
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetFilter() yfilter.YFilter { return timingStatus1588.YFilter }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) SetFilter(yf yfilter.YFilter) { timingStatus1588.YFilter = yf }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "slot" { return "Slot" }
    if yname == "port" { return "Port" }
    if yname == "clock-source" { return "ClockSource" }
    if yname == "rank" { return "Rank" }
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "quality-level-value" { return "QualityLevelValue" }
    if yname == "user-priority" { return "UserPriority" }
    if yname == "clock-state" { return "ClockState" }
    if yname == "is-select" { return "IsSelect" }
    return ""
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetSegmentPath() string {
    return "timing-status1588"
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = timingStatus1588.Input
    leafs["slot"] = timingStatus1588.Slot
    leafs["port"] = timingStatus1588.Port
    leafs["clock-source"] = timingStatus1588.ClockSource
    leafs["rank"] = timingStatus1588.Rank
    leafs["quality-level-option"] = timingStatus1588.QualityLevelOption
    leafs["quality-level-value"] = timingStatus1588.QualityLevelValue
    leafs["user-priority"] = timingStatus1588.UserPriority
    leafs["clock-state"] = timingStatus1588.ClockState
    leafs["is-select"] = timingStatus1588.IsSelect
    return leafs
}

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetBundleName() string { return "cisco_ios_xr" }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetYangName() string { return "timing-status1588" }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) SetParent(parent types.Entity) { timingStatus1588.parent = parent }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetParent() types.Entity { return timingStatus1588.parent }

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetParentYangName() string { return "syncc-instance" }

