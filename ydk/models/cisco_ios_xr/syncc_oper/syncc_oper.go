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

// NodeState represents Different modes of a node
type NodeState string

const (
    // Node in active mode
    NodeState_active NodeState = "active"

    // Node in standby mode
    NodeState_standby NodeState = "standby"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes applicable to timing controller.
    Nodes TimingController_Nodes
}

func (timingController *TimingController) GetEntityData() *types.CommonEntityData {
    timingController.EntityData.YFilter = timingController.YFilter
    timingController.EntityData.YangName = "timing-controller"
    timingController.EntityData.BundleName = "cisco_ios_xr"
    timingController.EntityData.ParentYangName = "Cisco-IOS-XR-syncc-oper"
    timingController.EntityData.SegmentPath = "Cisco-IOS-XR-syncc-oper:timing-controller"
    timingController.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingController.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingController.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingController.EntityData.Children = make(map[string]types.YChild)
    timingController.EntityData.Children["nodes"] = types.YChild{"Nodes", &timingController.Nodes}
    timingController.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timingController.EntityData)
}

// TimingController_Nodes
// List of nodes applicable to timing controller
type TimingController_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syncc operational data for a single node. The type is slice of
    // TimingController_Nodes_Node.
    Node []TimingController_Nodes_Node
}

func (nodes *TimingController_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "timing-controller"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// TimingController_Nodes_Node
// Syncc operational data for a single node
type TimingController_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node Name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Syncc state for a node.
    State TimingController_Nodes_Node_State

    // Syncc clock information for a node.
    Clock TimingController_Nodes_Node_Clock

    // Syncc timing information for a node.
    TimingSource TimingController_Nodes_Node_TimingSource
}

func (node *TimingController_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["state"] = types.YChild{"State", &node.State}
    node.EntityData.Children["clock"] = types.YChild{"Clock", &node.Clock}
    node.EntityData.Children["timing-source"] = types.YChild{"TimingSource", &node.TimingSource}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// TimingController_Nodes_Node_State
// Syncc state for a node
type TimingController_Nodes_Node_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc states. The type is slice of
    // TimingController_Nodes_Node_State_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_State_SynccInstance
}

func (state *TimingController_Nodes_Node_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "node"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["syncc-instance"] = types.YChild{"SynccInstance", nil}
    for i := range state.SynccInstance {
        state.EntityData.Children[types.GetSegmentPath(&state.SynccInstance[i])] = types.YChild{"SynccInstance", &state.SynccInstance[i]}
    }
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(state.EntityData)
}

// TimingController_Nodes_Node_State_SynccInstance
// List of syncc states
type TimingController_Nodes_Node_State_SynccInstance struct {
    EntityData types.CommonEntityData
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

func (synccInstance *TimingController_Nodes_Node_State_SynccInstance) GetEntityData() *types.CommonEntityData {
    synccInstance.EntityData.YFilter = synccInstance.YFilter
    synccInstance.EntityData.YangName = "syncc-instance"
    synccInstance.EntityData.BundleName = "cisco_ios_xr"
    synccInstance.EntityData.ParentYangName = "state"
    synccInstance.EntityData.SegmentPath = "syncc-instance"
    synccInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    synccInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    synccInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    synccInstance.EntityData.Children = make(map[string]types.YChild)
    synccInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    synccInstance.EntityData.Leafs["controller-state"] = types.YLeaf{"ControllerState", synccInstance.ControllerState}
    synccInstance.EntityData.Leafs["syncc-node-state"] = types.YLeaf{"SynccNodeState", synccInstance.SynccNodeState}
    synccInstance.EntityData.Leafs["verbose-level"] = types.YLeaf{"VerboseLevel", synccInstance.VerboseLevel}
    synccInstance.EntityData.Leafs["initial-count"] = types.YLeaf{"InitialCount", synccInstance.InitialCount}
    synccInstance.EntityData.Leafs["shutdown-count"] = types.YLeaf{"ShutdownCount", synccInstance.ShutdownCount}
    synccInstance.EntityData.Leafs["set-input-count"] = types.YLeaf{"SetInputCount", synccInstance.SetInputCount}
    synccInstance.EntityData.Leafs["set-capability-count"] = types.YLeaf{"SetCapabilityCount", synccInstance.SetCapabilityCount}
    synccInstance.EntityData.Leafs["get-clock-count"] = types.YLeaf{"GetClockCount", synccInstance.GetClockCount}
    synccInstance.EntityData.Leafs["set-clock-out-count"] = types.YLeaf{"SetClockOutCount", synccInstance.SetClockOutCount}
    synccInstance.EntityData.Leafs["sync-enable-count"] = types.YLeaf{"SyncEnableCount", synccInstance.SyncEnableCount}
    synccInstance.EntityData.Leafs["sync-disable-count"] = types.YLeaf{"SyncDisableCount", synccInstance.SyncDisableCount}
    synccInstance.EntityData.Leafs["capability-count"] = types.YLeaf{"CapabilityCount", synccInstance.CapabilityCount}
    synccInstance.EntityData.Leafs["set-quality-level-count"] = types.YLeaf{"SetQualityLevelCount", synccInstance.SetQualityLevelCount}
    synccInstance.EntityData.Leafs["input-notification"] = types.YLeaf{"InputNotification", synccInstance.InputNotification}
    synccInstance.EntityData.Leafs["capability-notification"] = types.YLeaf{"CapabilityNotification", synccInstance.CapabilityNotification}
    synccInstance.EntityData.Leafs["status-notification"] = types.YLeaf{"StatusNotification", synccInstance.StatusNotification}
    synccInstance.EntityData.Leafs["resync-notification"] = types.YLeaf{"ResyncNotification", synccInstance.ResyncNotification}
    return &(synccInstance.EntityData)
}

// TimingController_Nodes_Node_Clock
// Syncc clock information for a node
type TimingController_Nodes_Node_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc clock information . The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_Clock_SynccInstance
}

func (clock *TimingController_Nodes_Node_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "node"
    clock.EntityData.SegmentPath = "clock"
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = make(map[string]types.YChild)
    clock.EntityData.Children["syncc-instance"] = types.YChild{"SynccInstance", nil}
    for i := range clock.SynccInstance {
        clock.EntityData.Children[types.GetSegmentPath(&clock.SynccInstance[i])] = types.YChild{"SynccInstance", &clock.SynccInstance[i]}
    }
    clock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clock.EntityData)
}

// TimingController_Nodes_Node_Clock_SynccInstance
// List of syncc clock information 
type TimingController_Nodes_Node_Clock_SynccInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock table for an RP. The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance_Clock.
    Clock []TimingController_Nodes_Node_Clock_SynccInstance_Clock_
}

func (synccInstance *TimingController_Nodes_Node_Clock_SynccInstance) GetEntityData() *types.CommonEntityData {
    synccInstance.EntityData.YFilter = synccInstance.YFilter
    synccInstance.EntityData.YangName = "syncc-instance"
    synccInstance.EntityData.BundleName = "cisco_ios_xr"
    synccInstance.EntityData.ParentYangName = "clock"
    synccInstance.EntityData.SegmentPath = "syncc-instance"
    synccInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    synccInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    synccInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    synccInstance.EntityData.Children = make(map[string]types.YChild)
    synccInstance.EntityData.Children["clock"] = types.YChild{"Clock", nil}
    for i := range synccInstance.Clock {
        synccInstance.EntityData.Children[types.GetSegmentPath(&synccInstance.Clock[i])] = types.YChild{"Clock", &synccInstance.Clock[i]}
    }
    synccInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(synccInstance.EntityData)
}

// TimingController_Nodes_Node_Clock_SynccInstance_Clock_
// Clock table for an RP
type TimingController_Nodes_Node_Clock_SynccInstance_Clock_ struct {
    EntityData types.CommonEntityData
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

func (clock_ *TimingController_Nodes_Node_Clock_SynccInstance_Clock_) GetEntityData() *types.CommonEntityData {
    clock_.EntityData.YFilter = clock_.YFilter
    clock_.EntityData.YangName = "clock"
    clock_.EntityData.BundleName = "cisco_ios_xr"
    clock_.EntityData.ParentYangName = "syncc-instance"
    clock_.EntityData.SegmentPath = "clock"
    clock_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock_.EntityData.Children = make(map[string]types.YChild)
    clock_.EntityData.Leafs = make(map[string]types.YLeaf)
    clock_.EntityData.Leafs["is-configured-port0"] = types.YLeaf{"IsConfiguredPort0", clock_.IsConfiguredPort0}
    clock_.EntityData.Leafs["is-configured-port1"] = types.YLeaf{"IsConfiguredPort1", clock_.IsConfiguredPort1}
    clock_.EntityData.Leafs["is-configured-port2"] = types.YLeaf{"IsConfiguredPort2", clock_.IsConfiguredPort2}
    clock_.EntityData.Leafs["is-configured-port3"] = types.YLeaf{"IsConfiguredPort3", clock_.IsConfiguredPort3}
    clock_.EntityData.Leafs["mode-port0"] = types.YLeaf{"ModePort0", clock_.ModePort0}
    clock_.EntityData.Leafs["mode-port1"] = types.YLeaf{"ModePort1", clock_.ModePort1}
    clock_.EntityData.Leafs["mode-port2"] = types.YLeaf{"ModePort2", clock_.ModePort2}
    clock_.EntityData.Leafs["mode-port3"] = types.YLeaf{"ModePort3", clock_.ModePort3}
    clock_.EntityData.Leafs["submode1-port0"] = types.YLeaf{"Submode1Port0", clock_.Submode1Port0}
    clock_.EntityData.Leafs["submode1-port1"] = types.YLeaf{"Submode1Port1", clock_.Submode1Port1}
    clock_.EntityData.Leafs["submode1-port2"] = types.YLeaf{"Submode1Port2", clock_.Submode1Port2}
    clock_.EntityData.Leafs["submode1-port3"] = types.YLeaf{"Submode1Port3", clock_.Submode1Port3}
    clock_.EntityData.Leafs["submode2-port0"] = types.YLeaf{"Submode2Port0", clock_.Submode2Port0}
    clock_.EntityData.Leafs["submode2-port1"] = types.YLeaf{"Submode2Port1", clock_.Submode2Port1}
    clock_.EntityData.Leafs["submode2-port2"] = types.YLeaf{"Submode2Port2", clock_.Submode2Port2}
    clock_.EntityData.Leafs["submode2-port3"] = types.YLeaf{"Submode2Port3", clock_.Submode2Port3}
    clock_.EntityData.Leafs["submode3-port0"] = types.YLeaf{"Submode3Port0", clock_.Submode3Port0}
    clock_.EntityData.Leafs["submode3-port1"] = types.YLeaf{"Submode3Port1", clock_.Submode3Port1}
    clock_.EntityData.Leafs["submode3-port2"] = types.YLeaf{"Submode3Port2", clock_.Submode3Port2}
    clock_.EntityData.Leafs["submode3-port3"] = types.YLeaf{"Submode3Port3", clock_.Submode3Port3}
    clock_.EntityData.Leafs["shutdown-port0"] = types.YLeaf{"ShutdownPort0", clock_.ShutdownPort0}
    clock_.EntityData.Leafs["shutdown-port1"] = types.YLeaf{"ShutdownPort1", clock_.ShutdownPort1}
    clock_.EntityData.Leafs["shutdown-port2"] = types.YLeaf{"ShutdownPort2", clock_.ShutdownPort2}
    clock_.EntityData.Leafs["shutdown-port3"] = types.YLeaf{"ShutdownPort3", clock_.ShutdownPort3}
    clock_.EntityData.Leafs["direction-port0"] = types.YLeaf{"DirectionPort0", clock_.DirectionPort0}
    clock_.EntityData.Leafs["direction-port1"] = types.YLeaf{"DirectionPort1", clock_.DirectionPort1}
    clock_.EntityData.Leafs["direction-port2"] = types.YLeaf{"DirectionPort2", clock_.DirectionPort2}
    clock_.EntityData.Leafs["direction-port3"] = types.YLeaf{"DirectionPort3", clock_.DirectionPort3}
    clock_.EntityData.Leafs["baudrate-port0"] = types.YLeaf{"BaudratePort0", clock_.BaudratePort0}
    clock_.EntityData.Leafs["baudrate-port1"] = types.YLeaf{"BaudratePort1", clock_.BaudratePort1}
    clock_.EntityData.Leafs["baudrate-port2"] = types.YLeaf{"BaudratePort2", clock_.BaudratePort2}
    clock_.EntityData.Leafs["baudrate-port3"] = types.YLeaf{"BaudratePort3", clock_.BaudratePort3}
    clock_.EntityData.Leafs["quality-option-port0"] = types.YLeaf{"QualityOptionPort0", clock_.QualityOptionPort0}
    clock_.EntityData.Leafs["quality-option-port1"] = types.YLeaf{"QualityOptionPort1", clock_.QualityOptionPort1}
    clock_.EntityData.Leafs["quality-option-port2"] = types.YLeaf{"QualityOptionPort2", clock_.QualityOptionPort2}
    clock_.EntityData.Leafs["quality-option-port3"] = types.YLeaf{"QualityOptionPort3", clock_.QualityOptionPort3}
    clock_.EntityData.Leafs["transmit-ssm-port0"] = types.YLeaf{"TransmitSsmPort0", clock_.TransmitSsmPort0}
    clock_.EntityData.Leafs["transmit-ssm-port1"] = types.YLeaf{"TransmitSsmPort1", clock_.TransmitSsmPort1}
    clock_.EntityData.Leafs["transmit-ssm-port2"] = types.YLeaf{"TransmitSsmPort2", clock_.TransmitSsmPort2}
    clock_.EntityData.Leafs["transmit-ssm-port3"] = types.YLeaf{"TransmitSsmPort3", clock_.TransmitSsmPort3}
    clock_.EntityData.Leafs["recieve-ssm-port0"] = types.YLeaf{"RecieveSsmPort0", clock_.RecieveSsmPort0}
    clock_.EntityData.Leafs["recieve-ssm-port1"] = types.YLeaf{"RecieveSsmPort1", clock_.RecieveSsmPort1}
    clock_.EntityData.Leafs["recieve-ssm-port2"] = types.YLeaf{"RecieveSsmPort2", clock_.RecieveSsmPort2}
    clock_.EntityData.Leafs["recieve-ssm-port3"] = types.YLeaf{"RecieveSsmPort3", clock_.RecieveSsmPort3}
    clock_.EntityData.Leafs["interface-state-port0"] = types.YLeaf{"InterfaceStatePort0", clock_.InterfaceStatePort0}
    clock_.EntityData.Leafs["interface-state-port1"] = types.YLeaf{"InterfaceStatePort1", clock_.InterfaceStatePort1}
    clock_.EntityData.Leafs["interface-state-port2"] = types.YLeaf{"InterfaceStatePort2", clock_.InterfaceStatePort2}
    clock_.EntityData.Leafs["interface-state-port3"] = types.YLeaf{"InterfaceStatePort3", clock_.InterfaceStatePort3}
    return &(clock_.EntityData)
}

// TimingController_Nodes_Node_TimingSource
// Syncc timing information for a node
type TimingController_Nodes_Node_TimingSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc timing table information. The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance.
    SynccInstance []TimingController_Nodes_Node_TimingSource_SynccInstance
}

func (timingSource *TimingController_Nodes_Node_TimingSource) GetEntityData() *types.CommonEntityData {
    timingSource.EntityData.YFilter = timingSource.YFilter
    timingSource.EntityData.YangName = "timing-source"
    timingSource.EntityData.BundleName = "cisco_ios_xr"
    timingSource.EntityData.ParentYangName = "node"
    timingSource.EntityData.SegmentPath = "timing-source"
    timingSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingSource.EntityData.Children = make(map[string]types.YChild)
    timingSource.EntityData.Children["syncc-instance"] = types.YChild{"SynccInstance", nil}
    for i := range timingSource.SynccInstance {
        timingSource.EntityData.Children[types.GetSegmentPath(&timingSource.SynccInstance[i])] = types.YChild{"SynccInstance", &timingSource.SynccInstance[i]}
    }
    timingSource.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timingSource.EntityData)
}

// TimingController_Nodes_Node_TimingSource_SynccInstance
// List of syncc timing table information
type TimingController_Nodes_Node_TimingSource_SynccInstance struct {
    EntityData types.CommonEntityData
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

func (synccInstance *TimingController_Nodes_Node_TimingSource_SynccInstance) GetEntityData() *types.CommonEntityData {
    synccInstance.EntityData.YFilter = synccInstance.YFilter
    synccInstance.EntityData.YangName = "syncc-instance"
    synccInstance.EntityData.BundleName = "cisco_ios_xr"
    synccInstance.EntityData.ParentYangName = "timing-source"
    synccInstance.EntityData.SegmentPath = "syncc-instance"
    synccInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    synccInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    synccInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    synccInstance.EntityData.Children = make(map[string]types.YChild)
    synccInstance.EntityData.Children["timing-status-t0"] = types.YChild{"TimingStatusT0", nil}
    for i := range synccInstance.TimingStatusT0 {
        synccInstance.EntityData.Children[types.GetSegmentPath(&synccInstance.TimingStatusT0[i])] = types.YChild{"TimingStatusT0", &synccInstance.TimingStatusT0[i]}
    }
    synccInstance.EntityData.Children["timing-status-t4"] = types.YChild{"TimingStatusT4", nil}
    for i := range synccInstance.TimingStatusT4 {
        synccInstance.EntityData.Children[types.GetSegmentPath(&synccInstance.TimingStatusT4[i])] = types.YChild{"TimingStatusT4", &synccInstance.TimingStatusT4[i]}
    }
    synccInstance.EntityData.Children["timing-status1588"] = types.YChild{"TimingStatus1588", nil}
    for i := range synccInstance.TimingStatus1588 {
        synccInstance.EntityData.Children[types.GetSegmentPath(&synccInstance.TimingStatus1588[i])] = types.YChild{"TimingStatus1588", &synccInstance.TimingStatus1588[i]}
    }
    synccInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(synccInstance.EntityData)
}

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0
// Scheduling PLL T0 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0 struct {
    EntityData types.CommonEntityData
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

func (timingStatusT0 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0) GetEntityData() *types.CommonEntityData {
    timingStatusT0.EntityData.YFilter = timingStatusT0.YFilter
    timingStatusT0.EntityData.YangName = "timing-status-t0"
    timingStatusT0.EntityData.BundleName = "cisco_ios_xr"
    timingStatusT0.EntityData.ParentYangName = "syncc-instance"
    timingStatusT0.EntityData.SegmentPath = "timing-status-t0"
    timingStatusT0.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingStatusT0.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingStatusT0.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingStatusT0.EntityData.Children = make(map[string]types.YChild)
    timingStatusT0.EntityData.Leafs = make(map[string]types.YLeaf)
    timingStatusT0.EntityData.Leafs["input"] = types.YLeaf{"Input", timingStatusT0.Input}
    timingStatusT0.EntityData.Leafs["slot"] = types.YLeaf{"Slot", timingStatusT0.Slot}
    timingStatusT0.EntityData.Leafs["port"] = types.YLeaf{"Port", timingStatusT0.Port}
    timingStatusT0.EntityData.Leafs["clock-source"] = types.YLeaf{"ClockSource", timingStatusT0.ClockSource}
    timingStatusT0.EntityData.Leafs["rank"] = types.YLeaf{"Rank", timingStatusT0.Rank}
    timingStatusT0.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", timingStatusT0.QualityLevelOption}
    timingStatusT0.EntityData.Leafs["quality-level-value"] = types.YLeaf{"QualityLevelValue", timingStatusT0.QualityLevelValue}
    timingStatusT0.EntityData.Leafs["user-priority"] = types.YLeaf{"UserPriority", timingStatusT0.UserPriority}
    timingStatusT0.EntityData.Leafs["clock-state"] = types.YLeaf{"ClockState", timingStatusT0.ClockState}
    timingStatusT0.EntityData.Leafs["is-select"] = types.YLeaf{"IsSelect", timingStatusT0.IsSelect}
    return &(timingStatusT0.EntityData)
}

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4
// Scheduling PLL T4 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4 struct {
    EntityData types.CommonEntityData
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

func (timingStatusT4 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4) GetEntityData() *types.CommonEntityData {
    timingStatusT4.EntityData.YFilter = timingStatusT4.YFilter
    timingStatusT4.EntityData.YangName = "timing-status-t4"
    timingStatusT4.EntityData.BundleName = "cisco_ios_xr"
    timingStatusT4.EntityData.ParentYangName = "syncc-instance"
    timingStatusT4.EntityData.SegmentPath = "timing-status-t4"
    timingStatusT4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingStatusT4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingStatusT4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingStatusT4.EntityData.Children = make(map[string]types.YChild)
    timingStatusT4.EntityData.Leafs = make(map[string]types.YLeaf)
    timingStatusT4.EntityData.Leafs["input"] = types.YLeaf{"Input", timingStatusT4.Input}
    timingStatusT4.EntityData.Leafs["slot"] = types.YLeaf{"Slot", timingStatusT4.Slot}
    timingStatusT4.EntityData.Leafs["port"] = types.YLeaf{"Port", timingStatusT4.Port}
    timingStatusT4.EntityData.Leafs["clock-source"] = types.YLeaf{"ClockSource", timingStatusT4.ClockSource}
    timingStatusT4.EntityData.Leafs["rank"] = types.YLeaf{"Rank", timingStatusT4.Rank}
    timingStatusT4.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", timingStatusT4.QualityLevelOption}
    timingStatusT4.EntityData.Leafs["quality-level-value"] = types.YLeaf{"QualityLevelValue", timingStatusT4.QualityLevelValue}
    timingStatusT4.EntityData.Leafs["user-priority"] = types.YLeaf{"UserPriority", timingStatusT4.UserPriority}
    timingStatusT4.EntityData.Leafs["clock-state"] = types.YLeaf{"ClockState", timingStatusT4.ClockState}
    timingStatusT4.EntityData.Leafs["is-select"] = types.YLeaf{"IsSelect", timingStatusT4.IsSelect}
    return &(timingStatusT4.EntityData)
}

// TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588
// Scheduling PLL 1588 
type TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588 struct {
    EntityData types.CommonEntityData
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

func (timingStatus1588 *TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588) GetEntityData() *types.CommonEntityData {
    timingStatus1588.EntityData.YFilter = timingStatus1588.YFilter
    timingStatus1588.EntityData.YangName = "timing-status1588"
    timingStatus1588.EntityData.BundleName = "cisco_ios_xr"
    timingStatus1588.EntityData.ParentYangName = "syncc-instance"
    timingStatus1588.EntityData.SegmentPath = "timing-status1588"
    timingStatus1588.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timingStatus1588.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timingStatus1588.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timingStatus1588.EntityData.Children = make(map[string]types.YChild)
    timingStatus1588.EntityData.Leafs = make(map[string]types.YLeaf)
    timingStatus1588.EntityData.Leafs["input"] = types.YLeaf{"Input", timingStatus1588.Input}
    timingStatus1588.EntityData.Leafs["slot"] = types.YLeaf{"Slot", timingStatus1588.Slot}
    timingStatus1588.EntityData.Leafs["port"] = types.YLeaf{"Port", timingStatus1588.Port}
    timingStatus1588.EntityData.Leafs["clock-source"] = types.YLeaf{"ClockSource", timingStatus1588.ClockSource}
    timingStatus1588.EntityData.Leafs["rank"] = types.YLeaf{"Rank", timingStatus1588.Rank}
    timingStatus1588.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", timingStatus1588.QualityLevelOption}
    timingStatus1588.EntityData.Leafs["quality-level-value"] = types.YLeaf{"QualityLevelValue", timingStatus1588.QualityLevelValue}
    timingStatus1588.EntityData.Leafs["user-priority"] = types.YLeaf{"UserPriority", timingStatus1588.UserPriority}
    timingStatus1588.EntityData.Leafs["clock-state"] = types.YLeaf{"ClockState", timingStatus1588.ClockState}
    timingStatus1588.EntityData.Leafs["is-select"] = types.YLeaf{"IsSelect", timingStatus1588.IsSelect}
    return &(timingStatus1588.EntityData)
}

