// This module contains a collection of YANG definitions
// for Cisco IOS-XR syncc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   timing-controller: Timing controller operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    SourceStateName_error_ SourceStateName = "error"
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

    timingController.EntityData.Children = types.NewOrderedMap()
    timingController.EntityData.Children.Append("nodes", types.YChild{"Nodes", &timingController.Nodes})
    timingController.EntityData.Leafs = types.NewOrderedMap()

    timingController.EntityData.YListKeys = []string {}

    return &(timingController.EntityData)
}

// TimingController_Nodes
// List of nodes applicable to timing controller
type TimingController_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syncc operational data for a single node. The type is slice of
    // TimingController_Nodes_Node.
    Node []*TimingController_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// TimingController_Nodes_Node
// Syncc operational data for a single node
type TimingController_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *TimingController_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("state", types.YChild{"State", &node.State})
    node.EntityData.Children.Append("clock", types.YChild{"Clock", &node.Clock})
    node.EntityData.Children.Append("timing-source", types.YChild{"TimingSource", &node.TimingSource})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// TimingController_Nodes_Node_State
// Syncc state for a node
type TimingController_Nodes_Node_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc states. The type is slice of
    // TimingController_Nodes_Node_State_SynccInstance.
    SynccInstance []*TimingController_Nodes_Node_State_SynccInstance
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("syncc-instance", types.YChild{"SynccInstance", nil})
    for i := range state.SynccInstance {
        state.EntityData.Children.Append(types.GetSegmentPath(state.SynccInstance[i]), types.YChild{"SynccInstance", state.SynccInstance[i]})
    }
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

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

    synccInstance.EntityData.Children = types.NewOrderedMap()
    synccInstance.EntityData.Leafs = types.NewOrderedMap()
    synccInstance.EntityData.Leafs.Append("controller-state", types.YLeaf{"ControllerState", synccInstance.ControllerState})
    synccInstance.EntityData.Leafs.Append("syncc-node-state", types.YLeaf{"SynccNodeState", synccInstance.SynccNodeState})
    synccInstance.EntityData.Leafs.Append("verbose-level", types.YLeaf{"VerboseLevel", synccInstance.VerboseLevel})
    synccInstance.EntityData.Leafs.Append("initial-count", types.YLeaf{"InitialCount", synccInstance.InitialCount})
    synccInstance.EntityData.Leafs.Append("shutdown-count", types.YLeaf{"ShutdownCount", synccInstance.ShutdownCount})
    synccInstance.EntityData.Leafs.Append("set-input-count", types.YLeaf{"SetInputCount", synccInstance.SetInputCount})
    synccInstance.EntityData.Leafs.Append("set-capability-count", types.YLeaf{"SetCapabilityCount", synccInstance.SetCapabilityCount})
    synccInstance.EntityData.Leafs.Append("get-clock-count", types.YLeaf{"GetClockCount", synccInstance.GetClockCount})
    synccInstance.EntityData.Leafs.Append("set-clock-out-count", types.YLeaf{"SetClockOutCount", synccInstance.SetClockOutCount})
    synccInstance.EntityData.Leafs.Append("sync-enable-count", types.YLeaf{"SyncEnableCount", synccInstance.SyncEnableCount})
    synccInstance.EntityData.Leafs.Append("sync-disable-count", types.YLeaf{"SyncDisableCount", synccInstance.SyncDisableCount})
    synccInstance.EntityData.Leafs.Append("capability-count", types.YLeaf{"CapabilityCount", synccInstance.CapabilityCount})
    synccInstance.EntityData.Leafs.Append("set-quality-level-count", types.YLeaf{"SetQualityLevelCount", synccInstance.SetQualityLevelCount})
    synccInstance.EntityData.Leafs.Append("input-notification", types.YLeaf{"InputNotification", synccInstance.InputNotification})
    synccInstance.EntityData.Leafs.Append("capability-notification", types.YLeaf{"CapabilityNotification", synccInstance.CapabilityNotification})
    synccInstance.EntityData.Leafs.Append("status-notification", types.YLeaf{"StatusNotification", synccInstance.StatusNotification})
    synccInstance.EntityData.Leafs.Append("resync-notification", types.YLeaf{"ResyncNotification", synccInstance.ResyncNotification})

    synccInstance.EntityData.YListKeys = []string {}

    return &(synccInstance.EntityData)
}

// TimingController_Nodes_Node_Clock
// Syncc clock information for a node
type TimingController_Nodes_Node_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc clock information . The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance.
    SynccInstance []*TimingController_Nodes_Node_Clock_SynccInstance
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

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("syncc-instance", types.YChild{"SynccInstance", nil})
    for i := range clock.SynccInstance {
        clock.EntityData.Children.Append(types.GetSegmentPath(clock.SynccInstance[i]), types.YChild{"SynccInstance", clock.SynccInstance[i]})
    }
    clock.EntityData.Leafs = types.NewOrderedMap()

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// TimingController_Nodes_Node_Clock_SynccInstance
// List of syncc clock information 
type TimingController_Nodes_Node_Clock_SynccInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clock table for an RP. The type is slice of
    // TimingController_Nodes_Node_Clock_SynccInstance_Clock.
    Clock []*TimingController_Nodes_Node_Clock_SynccInstance_Clock
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

    synccInstance.EntityData.Children = types.NewOrderedMap()
    synccInstance.EntityData.Children.Append("clock", types.YChild{"Clock", nil})
    for i := range synccInstance.Clock {
        synccInstance.EntityData.Children.Append(types.GetSegmentPath(synccInstance.Clock[i]), types.YChild{"Clock", synccInstance.Clock[i]})
    }
    synccInstance.EntityData.Leafs = types.NewOrderedMap()

    synccInstance.EntityData.YListKeys = []string {}

    return &(synccInstance.EntityData)
}

// TimingController_Nodes_Node_Clock_SynccInstance_Clock
// Clock table for an RP
type TimingController_Nodes_Node_Clock_SynccInstance_Clock struct {
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

func (clock *TimingController_Nodes_Node_Clock_SynccInstance_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "syncc-instance"
    clock.EntityData.SegmentPath = "clock"
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Leafs = types.NewOrderedMap()
    clock.EntityData.Leafs.Append("is-configured-port0", types.YLeaf{"IsConfiguredPort0", clock.IsConfiguredPort0})
    clock.EntityData.Leafs.Append("is-configured-port1", types.YLeaf{"IsConfiguredPort1", clock.IsConfiguredPort1})
    clock.EntityData.Leafs.Append("is-configured-port2", types.YLeaf{"IsConfiguredPort2", clock.IsConfiguredPort2})
    clock.EntityData.Leafs.Append("is-configured-port3", types.YLeaf{"IsConfiguredPort3", clock.IsConfiguredPort3})
    clock.EntityData.Leafs.Append("mode-port0", types.YLeaf{"ModePort0", clock.ModePort0})
    clock.EntityData.Leafs.Append("mode-port1", types.YLeaf{"ModePort1", clock.ModePort1})
    clock.EntityData.Leafs.Append("mode-port2", types.YLeaf{"ModePort2", clock.ModePort2})
    clock.EntityData.Leafs.Append("mode-port3", types.YLeaf{"ModePort3", clock.ModePort3})
    clock.EntityData.Leafs.Append("submode1-port0", types.YLeaf{"Submode1Port0", clock.Submode1Port0})
    clock.EntityData.Leafs.Append("submode1-port1", types.YLeaf{"Submode1Port1", clock.Submode1Port1})
    clock.EntityData.Leafs.Append("submode1-port2", types.YLeaf{"Submode1Port2", clock.Submode1Port2})
    clock.EntityData.Leafs.Append("submode1-port3", types.YLeaf{"Submode1Port3", clock.Submode1Port3})
    clock.EntityData.Leafs.Append("submode2-port0", types.YLeaf{"Submode2Port0", clock.Submode2Port0})
    clock.EntityData.Leafs.Append("submode2-port1", types.YLeaf{"Submode2Port1", clock.Submode2Port1})
    clock.EntityData.Leafs.Append("submode2-port2", types.YLeaf{"Submode2Port2", clock.Submode2Port2})
    clock.EntityData.Leafs.Append("submode2-port3", types.YLeaf{"Submode2Port3", clock.Submode2Port3})
    clock.EntityData.Leafs.Append("submode3-port0", types.YLeaf{"Submode3Port0", clock.Submode3Port0})
    clock.EntityData.Leafs.Append("submode3-port1", types.YLeaf{"Submode3Port1", clock.Submode3Port1})
    clock.EntityData.Leafs.Append("submode3-port2", types.YLeaf{"Submode3Port2", clock.Submode3Port2})
    clock.EntityData.Leafs.Append("submode3-port3", types.YLeaf{"Submode3Port3", clock.Submode3Port3})
    clock.EntityData.Leafs.Append("shutdown-port0", types.YLeaf{"ShutdownPort0", clock.ShutdownPort0})
    clock.EntityData.Leafs.Append("shutdown-port1", types.YLeaf{"ShutdownPort1", clock.ShutdownPort1})
    clock.EntityData.Leafs.Append("shutdown-port2", types.YLeaf{"ShutdownPort2", clock.ShutdownPort2})
    clock.EntityData.Leafs.Append("shutdown-port3", types.YLeaf{"ShutdownPort3", clock.ShutdownPort3})
    clock.EntityData.Leafs.Append("direction-port0", types.YLeaf{"DirectionPort0", clock.DirectionPort0})
    clock.EntityData.Leafs.Append("direction-port1", types.YLeaf{"DirectionPort1", clock.DirectionPort1})
    clock.EntityData.Leafs.Append("direction-port2", types.YLeaf{"DirectionPort2", clock.DirectionPort2})
    clock.EntityData.Leafs.Append("direction-port3", types.YLeaf{"DirectionPort3", clock.DirectionPort3})
    clock.EntityData.Leafs.Append("baudrate-port0", types.YLeaf{"BaudratePort0", clock.BaudratePort0})
    clock.EntityData.Leafs.Append("baudrate-port1", types.YLeaf{"BaudratePort1", clock.BaudratePort1})
    clock.EntityData.Leafs.Append("baudrate-port2", types.YLeaf{"BaudratePort2", clock.BaudratePort2})
    clock.EntityData.Leafs.Append("baudrate-port3", types.YLeaf{"BaudratePort3", clock.BaudratePort3})
    clock.EntityData.Leafs.Append("quality-option-port0", types.YLeaf{"QualityOptionPort0", clock.QualityOptionPort0})
    clock.EntityData.Leafs.Append("quality-option-port1", types.YLeaf{"QualityOptionPort1", clock.QualityOptionPort1})
    clock.EntityData.Leafs.Append("quality-option-port2", types.YLeaf{"QualityOptionPort2", clock.QualityOptionPort2})
    clock.EntityData.Leafs.Append("quality-option-port3", types.YLeaf{"QualityOptionPort3", clock.QualityOptionPort3})
    clock.EntityData.Leafs.Append("transmit-ssm-port0", types.YLeaf{"TransmitSsmPort0", clock.TransmitSsmPort0})
    clock.EntityData.Leafs.Append("transmit-ssm-port1", types.YLeaf{"TransmitSsmPort1", clock.TransmitSsmPort1})
    clock.EntityData.Leafs.Append("transmit-ssm-port2", types.YLeaf{"TransmitSsmPort2", clock.TransmitSsmPort2})
    clock.EntityData.Leafs.Append("transmit-ssm-port3", types.YLeaf{"TransmitSsmPort3", clock.TransmitSsmPort3})
    clock.EntityData.Leafs.Append("recieve-ssm-port0", types.YLeaf{"RecieveSsmPort0", clock.RecieveSsmPort0})
    clock.EntityData.Leafs.Append("recieve-ssm-port1", types.YLeaf{"RecieveSsmPort1", clock.RecieveSsmPort1})
    clock.EntityData.Leafs.Append("recieve-ssm-port2", types.YLeaf{"RecieveSsmPort2", clock.RecieveSsmPort2})
    clock.EntityData.Leafs.Append("recieve-ssm-port3", types.YLeaf{"RecieveSsmPort3", clock.RecieveSsmPort3})
    clock.EntityData.Leafs.Append("interface-state-port0", types.YLeaf{"InterfaceStatePort0", clock.InterfaceStatePort0})
    clock.EntityData.Leafs.Append("interface-state-port1", types.YLeaf{"InterfaceStatePort1", clock.InterfaceStatePort1})
    clock.EntityData.Leafs.Append("interface-state-port2", types.YLeaf{"InterfaceStatePort2", clock.InterfaceStatePort2})
    clock.EntityData.Leafs.Append("interface-state-port3", types.YLeaf{"InterfaceStatePort3", clock.InterfaceStatePort3})

    clock.EntityData.YListKeys = []string {}

    return &(clock.EntityData)
}

// TimingController_Nodes_Node_TimingSource
// Syncc timing information for a node
type TimingController_Nodes_Node_TimingSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of syncc timing table information. The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance.
    SynccInstance []*TimingController_Nodes_Node_TimingSource_SynccInstance
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

    timingSource.EntityData.Children = types.NewOrderedMap()
    timingSource.EntityData.Children.Append("syncc-instance", types.YChild{"SynccInstance", nil})
    for i := range timingSource.SynccInstance {
        timingSource.EntityData.Children.Append(types.GetSegmentPath(timingSource.SynccInstance[i]), types.YChild{"SynccInstance", timingSource.SynccInstance[i]})
    }
    timingSource.EntityData.Leafs = types.NewOrderedMap()

    timingSource.EntityData.YListKeys = []string {}

    return &(timingSource.EntityData)
}

// TimingController_Nodes_Node_TimingSource_SynccInstance
// List of syncc timing table information
type TimingController_Nodes_Node_TimingSource_SynccInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Scheduling PLL T0 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0.
    TimingStatusT0 []*TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT0

    // Scheduling PLL T4 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4.
    TimingStatusT4 []*TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatusT4

    // Scheduling PLL 1588 . The type is slice of
    // TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588.
    TimingStatus1588 []*TimingController_Nodes_Node_TimingSource_SynccInstance_TimingStatus1588
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

    synccInstance.EntityData.Children = types.NewOrderedMap()
    synccInstance.EntityData.Children.Append("timing-status-t0", types.YChild{"TimingStatusT0", nil})
    for i := range synccInstance.TimingStatusT0 {
        synccInstance.EntityData.Children.Append(types.GetSegmentPath(synccInstance.TimingStatusT0[i]), types.YChild{"TimingStatusT0", synccInstance.TimingStatusT0[i]})
    }
    synccInstance.EntityData.Children.Append("timing-status-t4", types.YChild{"TimingStatusT4", nil})
    for i := range synccInstance.TimingStatusT4 {
        synccInstance.EntityData.Children.Append(types.GetSegmentPath(synccInstance.TimingStatusT4[i]), types.YChild{"TimingStatusT4", synccInstance.TimingStatusT4[i]})
    }
    synccInstance.EntityData.Children.Append("timing-status1588", types.YChild{"TimingStatus1588", nil})
    for i := range synccInstance.TimingStatus1588 {
        synccInstance.EntityData.Children.Append(types.GetSegmentPath(synccInstance.TimingStatus1588[i]), types.YChild{"TimingStatus1588", synccInstance.TimingStatus1588[i]})
    }
    synccInstance.EntityData.Leafs = types.NewOrderedMap()

    synccInstance.EntityData.YListKeys = []string {}

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

    timingStatusT0.EntityData.Children = types.NewOrderedMap()
    timingStatusT0.EntityData.Leafs = types.NewOrderedMap()
    timingStatusT0.EntityData.Leafs.Append("input", types.YLeaf{"Input", timingStatusT0.Input})
    timingStatusT0.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", timingStatusT0.Slot})
    timingStatusT0.EntityData.Leafs.Append("port", types.YLeaf{"Port", timingStatusT0.Port})
    timingStatusT0.EntityData.Leafs.Append("clock-source", types.YLeaf{"ClockSource", timingStatusT0.ClockSource})
    timingStatusT0.EntityData.Leafs.Append("rank", types.YLeaf{"Rank", timingStatusT0.Rank})
    timingStatusT0.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", timingStatusT0.QualityLevelOption})
    timingStatusT0.EntityData.Leafs.Append("quality-level-value", types.YLeaf{"QualityLevelValue", timingStatusT0.QualityLevelValue})
    timingStatusT0.EntityData.Leafs.Append("user-priority", types.YLeaf{"UserPriority", timingStatusT0.UserPriority})
    timingStatusT0.EntityData.Leafs.Append("clock-state", types.YLeaf{"ClockState", timingStatusT0.ClockState})
    timingStatusT0.EntityData.Leafs.Append("is-select", types.YLeaf{"IsSelect", timingStatusT0.IsSelect})

    timingStatusT0.EntityData.YListKeys = []string {}

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

    timingStatusT4.EntityData.Children = types.NewOrderedMap()
    timingStatusT4.EntityData.Leafs = types.NewOrderedMap()
    timingStatusT4.EntityData.Leafs.Append("input", types.YLeaf{"Input", timingStatusT4.Input})
    timingStatusT4.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", timingStatusT4.Slot})
    timingStatusT4.EntityData.Leafs.Append("port", types.YLeaf{"Port", timingStatusT4.Port})
    timingStatusT4.EntityData.Leafs.Append("clock-source", types.YLeaf{"ClockSource", timingStatusT4.ClockSource})
    timingStatusT4.EntityData.Leafs.Append("rank", types.YLeaf{"Rank", timingStatusT4.Rank})
    timingStatusT4.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", timingStatusT4.QualityLevelOption})
    timingStatusT4.EntityData.Leafs.Append("quality-level-value", types.YLeaf{"QualityLevelValue", timingStatusT4.QualityLevelValue})
    timingStatusT4.EntityData.Leafs.Append("user-priority", types.YLeaf{"UserPriority", timingStatusT4.UserPriority})
    timingStatusT4.EntityData.Leafs.Append("clock-state", types.YLeaf{"ClockState", timingStatusT4.ClockState})
    timingStatusT4.EntityData.Leafs.Append("is-select", types.YLeaf{"IsSelect", timingStatusT4.IsSelect})

    timingStatusT4.EntityData.YListKeys = []string {}

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

    timingStatus1588.EntityData.Children = types.NewOrderedMap()
    timingStatus1588.EntityData.Leafs = types.NewOrderedMap()
    timingStatus1588.EntityData.Leafs.Append("input", types.YLeaf{"Input", timingStatus1588.Input})
    timingStatus1588.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", timingStatus1588.Slot})
    timingStatus1588.EntityData.Leafs.Append("port", types.YLeaf{"Port", timingStatus1588.Port})
    timingStatus1588.EntityData.Leafs.Append("clock-source", types.YLeaf{"ClockSource", timingStatus1588.ClockSource})
    timingStatus1588.EntityData.Leafs.Append("rank", types.YLeaf{"Rank", timingStatus1588.Rank})
    timingStatus1588.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", timingStatus1588.QualityLevelOption})
    timingStatus1588.EntityData.Leafs.Append("quality-level-value", types.YLeaf{"QualityLevelValue", timingStatus1588.QualityLevelValue})
    timingStatus1588.EntityData.Leafs.Append("user-priority", types.YLeaf{"UserPriority", timingStatus1588.UserPriority})
    timingStatus1588.EntityData.Leafs.Append("clock-state", types.YLeaf{"ClockState", timingStatus1588.ClockState})
    timingStatus1588.EntityData.Leafs.Append("is-select", types.YLeaf{"IsSelect", timingStatus1588.IsSelect})

    timingStatus1588.EntityData.YListKeys = []string {}

    return &(timingStatus1588.EntityData)
}

