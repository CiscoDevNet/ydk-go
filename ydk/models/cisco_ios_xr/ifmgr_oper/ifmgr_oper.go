// This module contains a collection of YANG definitions
// for Cisco IOS-XR ifmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   interface-dampening: Interface dampening data
//   interface-properties: interface properties
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ifmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ifmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ifmgr-oper interface-dampening}", reflect.TypeOf(InterfaceDampening{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ifmgr-oper:interface-dampening", reflect.TypeOf(InterfaceDampening{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ifmgr-oper interface-properties}", reflect.TypeOf(InterfaceProperties{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ifmgr-oper:interface-properties", reflect.TypeOf(InterfaceProperties{}))
}

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// InterfaceDampening
// Interface dampening data
type InterfaceDampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface list for which dampening info is available.
    Interfaces InterfaceDampening_Interfaces

    // The location of the interface(s) being queried.
    Nodes InterfaceDampening_Nodes
}

func (interfaceDampening *InterfaceDampening) GetEntityData() *types.CommonEntityData {
    interfaceDampening.EntityData.YFilter = interfaceDampening.YFilter
    interfaceDampening.EntityData.YangName = "interface-dampening"
    interfaceDampening.EntityData.BundleName = "cisco_ios_xr"
    interfaceDampening.EntityData.ParentYangName = "Cisco-IOS-XR-ifmgr-oper"
    interfaceDampening.EntityData.SegmentPath = "Cisco-IOS-XR-ifmgr-oper:interface-dampening"
    interfaceDampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDampening.EntityData.Children = make(map[string]types.YChild)
    interfaceDampening.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceDampening.Interfaces}
    interfaceDampening.EntityData.Children["nodes"] = types.YChild{"Nodes", &interfaceDampening.Nodes}
    interfaceDampening.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceDampening.EntityData)
}

// InterfaceDampening_Interfaces
// The interface list for which dampening info is
// available
type InterfaceDampening_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface for which dampening info is being queried. The type is slice
    // of InterfaceDampening_Interfaces_Interface_.
    Interface_ []InterfaceDampening_Interfaces_Interface
}

func (interfaces *InterfaceDampening_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-dampening"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// InterfaceDampening_Interfaces_Interface
// The interface for which dampening info is being
// queried
type InterfaceDampening_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Dampening info for the interface.
    IfDampening InterfaceDampening_Interfaces_Interface_IfDampening
}

func (self *InterfaceDampening_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["if-dampening"] = types.YChild{"IfDampening", &self.IfDampening}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// InterfaceDampening_Interfaces_Interface_IfDampening
// Dampening info for the interface
type InterfaceDampening_Interfaces_Interface_IfDampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times the state has changed. The type is interface{} with
    // range: 0..4294967295.
    StateTransitionCount interface{}

    // The time elasped after the last state transition. The type is interface{}
    // with range: 0..4294967295.
    LastStateTransitionTime interface{}

    // Flag showing if dampening is enabled. The type is bool.
    IsDampeningEnabled interface{}

    // Configured decay half life in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    HalfLife interface{}

    // Configured reuse threshold. The type is interface{} with range:
    // 0..4294967295.
    ReuseThreshold interface{}

    // Value of suppress threshold. The type is interface{} with range:
    // 0..4294967295.
    SuppressThreshold interface{}

    // Maximum suppress time in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    MaximumSuppressTime interface{}

    // Configured restart penalty. The type is interface{} with range:
    // 0..4294967295.
    RestartPenalty interface{}

    // Interface dampening.
    InterfaceDampening InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening_

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation.
    Capsulation []InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetEntityData() *types.CommonEntityData {
    ifDampening.EntityData.YFilter = ifDampening.YFilter
    ifDampening.EntityData.YangName = "if-dampening"
    ifDampening.EntityData.BundleName = "cisco_ios_xr"
    ifDampening.EntityData.ParentYangName = "interface"
    ifDampening.EntityData.SegmentPath = "if-dampening"
    ifDampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifDampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifDampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifDampening.EntityData.Children = make(map[string]types.YChild)
    ifDampening.EntityData.Children["interface-dampening"] = types.YChild{"InterfaceDampening", &ifDampening.InterfaceDampening}
    ifDampening.EntityData.Children["capsulation"] = types.YChild{"Capsulation", nil}
    for i := range ifDampening.Capsulation {
        ifDampening.EntityData.Children[types.GetSegmentPath(&ifDampening.Capsulation[i])] = types.YChild{"Capsulation", &ifDampening.Capsulation[i]}
    }
    ifDampening.EntityData.Leafs = make(map[string]types.YLeaf)
    ifDampening.EntityData.Leafs["state-transition-count"] = types.YLeaf{"StateTransitionCount", ifDampening.StateTransitionCount}
    ifDampening.EntityData.Leafs["last-state-transition-time"] = types.YLeaf{"LastStateTransitionTime", ifDampening.LastStateTransitionTime}
    ifDampening.EntityData.Leafs["is-dampening-enabled"] = types.YLeaf{"IsDampeningEnabled", ifDampening.IsDampeningEnabled}
    ifDampening.EntityData.Leafs["half-life"] = types.YLeaf{"HalfLife", ifDampening.HalfLife}
    ifDampening.EntityData.Leafs["reuse-threshold"] = types.YLeaf{"ReuseThreshold", ifDampening.ReuseThreshold}
    ifDampening.EntityData.Leafs["suppress-threshold"] = types.YLeaf{"SuppressThreshold", ifDampening.SuppressThreshold}
    ifDampening.EntityData.Leafs["maximum-suppress-time"] = types.YLeaf{"MaximumSuppressTime", ifDampening.MaximumSuppressTime}
    ifDampening.EntityData.Leafs["restart-penalty"] = types.YLeaf{"RestartPenalty", ifDampening.RestartPenalty}
    return &(ifDampening.EntityData)
}

// InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening_
// Interface dampening
type InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (interfaceDampening_ *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening_) GetEntityData() *types.CommonEntityData {
    interfaceDampening_.EntityData.YFilter = interfaceDampening_.YFilter
    interfaceDampening_.EntityData.YangName = "interface-dampening"
    interfaceDampening_.EntityData.BundleName = "cisco_ios_xr"
    interfaceDampening_.EntityData.ParentYangName = "if-dampening"
    interfaceDampening_.EntityData.SegmentPath = "interface-dampening"
    interfaceDampening_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDampening_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDampening_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDampening_.EntityData.Children = make(map[string]types.YChild)
    interfaceDampening_.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDampening_.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", interfaceDampening_.Penalty}
    interfaceDampening_.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", interfaceDampening_.IsSuppressedEnabled}
    interfaceDampening_.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", interfaceDampening_.SecondsRemaining}
    interfaceDampening_.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", interfaceDampening_.Flaps}
    interfaceDampening_.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceDampening_.State}
    return &(interfaceDampening_.EntityData)
}

// InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetEntityData() *types.CommonEntityData {
    capsulation.EntityData.YFilter = capsulation.YFilter
    capsulation.EntityData.YangName = "capsulation"
    capsulation.EntityData.BundleName = "cisco_ios_xr"
    capsulation.EntityData.ParentYangName = "if-dampening"
    capsulation.EntityData.SegmentPath = "capsulation"
    capsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulation.EntityData.Children = make(map[string]types.YChild)
    capsulation.EntityData.Children["capsulation-dampening"] = types.YChild{"CapsulationDampening", &capsulation.CapsulationDampening}
    capsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulation.EntityData.Leafs["capsulation-number"] = types.YLeaf{"CapsulationNumber", capsulation.CapsulationNumber}
    return &(capsulation.EntityData)
}

// InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetEntityData() *types.CommonEntityData {
    capsulationDampening.EntityData.YFilter = capsulationDampening.YFilter
    capsulationDampening.EntityData.YangName = "capsulation-dampening"
    capsulationDampening.EntityData.BundleName = "cisco_ios_xr"
    capsulationDampening.EntityData.ParentYangName = "capsulation"
    capsulationDampening.EntityData.SegmentPath = "capsulation-dampening"
    capsulationDampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulationDampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulationDampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulationDampening.EntityData.Children = make(map[string]types.YChild)
    capsulationDampening.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulationDampening.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", capsulationDampening.Penalty}
    capsulationDampening.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", capsulationDampening.IsSuppressedEnabled}
    capsulationDampening.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", capsulationDampening.SecondsRemaining}
    capsulationDampening.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", capsulationDampening.Flaps}
    capsulationDampening.EntityData.Leafs["state"] = types.YLeaf{"State", capsulationDampening.State}
    return &(capsulationDampening.EntityData)
}

// InterfaceDampening_Nodes
// The location of the interface(s) being queried
type InterfaceDampening_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The location of the interface(s) being queried. The type is slice of
    // InterfaceDampening_Nodes_Node.
    Node []InterfaceDampening_Nodes_Node
}

func (nodes *InterfaceDampening_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "interface-dampening"
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

// InterfaceDampening_Nodes_Node
// The location of the interface(s) being queried
type InterfaceDampening_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The location of the interface(s). The type is
    // string with pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Show details for the interfaces.
    Show InterfaceDampening_Nodes_Node_Show
}

func (node *InterfaceDampening_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["show"] = types.YChild{"Show", &node.Show}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// InterfaceDampening_Nodes_Node_Show
// Show details for the interfaces
type InterfaceDampening_Nodes_Node_Show struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening information of the interface(s) being queried.
    Dampening InterfaceDampening_Nodes_Node_Show_Dampening
}

func (show *InterfaceDampening_Nodes_Node_Show) GetEntityData() *types.CommonEntityData {
    show.EntityData.YFilter = show.YFilter
    show.EntityData.YangName = "show"
    show.EntityData.BundleName = "cisco_ios_xr"
    show.EntityData.ParentYangName = "node"
    show.EntityData.SegmentPath = "show"
    show.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    show.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    show.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    show.EntityData.Children = make(map[string]types.YChild)
    show.EntityData.Children["dampening"] = types.YChild{"Dampening", &show.Dampening}
    show.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(show.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening
// Dampening information of the interface(s)
// being queried
type InterfaceDampening_Nodes_Node_Show_Dampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface handle for which dampening info queried.
    IfHandles InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles

    // Table of interfaces for which dampening info can be queried.
    Interfaces InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetEntityData() *types.CommonEntityData {
    dampening.EntityData.YFilter = dampening.YFilter
    dampening.EntityData.YangName = "dampening"
    dampening.EntityData.BundleName = "cisco_ios_xr"
    dampening.EntityData.ParentYangName = "show"
    dampening.EntityData.SegmentPath = "dampening"
    dampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dampening.EntityData.Children = make(map[string]types.YChild)
    dampening.EntityData.Children["if-handles"] = types.YChild{"IfHandles", &dampening.IfHandles}
    dampening.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &dampening.Interfaces}
    dampening.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dampening.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles
// Interface handle for which dampening info
// queried
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening info for the interface handle. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle.
    IfHandle []InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetEntityData() *types.CommonEntityData {
    ifHandles.EntityData.YFilter = ifHandles.YFilter
    ifHandles.EntityData.YangName = "if-handles"
    ifHandles.EntityData.BundleName = "cisco_ios_xr"
    ifHandles.EntityData.ParentYangName = "dampening"
    ifHandles.EntityData.SegmentPath = "if-handles"
    ifHandles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifHandles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifHandles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifHandles.EntityData.Children = make(map[string]types.YChild)
    ifHandles.EntityData.Children["if-handle"] = types.YChild{"IfHandle", nil}
    for i := range ifHandles.IfHandle {
        ifHandles.EntityData.Children[types.GetSegmentPath(&ifHandles.IfHandle[i])] = types.YChild{"IfHandle", &ifHandles.IfHandle[i]}
    }
    ifHandles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ifHandles.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle
// Dampening info for the interface handle
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface handle. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    InterfaceHandleName interface{}

    // The number of times the state has changed. The type is interface{} with
    // range: 0..4294967295.
    StateTransitionCount interface{}

    // The time elasped after the last state transition. The type is interface{}
    // with range: 0..4294967295.
    LastStateTransitionTime interface{}

    // Flag showing if dampening is enabled. The type is bool.
    IsDampeningEnabled interface{}

    // Configured decay half life in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    HalfLife interface{}

    // Configured reuse threshold. The type is interface{} with range:
    // 0..4294967295.
    ReuseThreshold interface{}

    // Value of suppress threshold. The type is interface{} with range:
    // 0..4294967295.
    SuppressThreshold interface{}

    // Maximum suppress time in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    MaximumSuppressTime interface{}

    // Configured restart penalty. The type is interface{} with range:
    // 0..4294967295.
    RestartPenalty interface{}

    // Interface dampening.
    InterfaceDampening InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening_

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation.
    Capsulation []InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetEntityData() *types.CommonEntityData {
    ifHandle.EntityData.YFilter = ifHandle.YFilter
    ifHandle.EntityData.YangName = "if-handle"
    ifHandle.EntityData.BundleName = "cisco_ios_xr"
    ifHandle.EntityData.ParentYangName = "if-handles"
    ifHandle.EntityData.SegmentPath = "if-handle" + "[interface-handle-name='" + fmt.Sprintf("%v", ifHandle.InterfaceHandleName) + "']"
    ifHandle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifHandle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifHandle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifHandle.EntityData.Children = make(map[string]types.YChild)
    ifHandle.EntityData.Children["interface-dampening"] = types.YChild{"InterfaceDampening", &ifHandle.InterfaceDampening}
    ifHandle.EntityData.Children["capsulation"] = types.YChild{"Capsulation", nil}
    for i := range ifHandle.Capsulation {
        ifHandle.EntityData.Children[types.GetSegmentPath(&ifHandle.Capsulation[i])] = types.YChild{"Capsulation", &ifHandle.Capsulation[i]}
    }
    ifHandle.EntityData.Leafs = make(map[string]types.YLeaf)
    ifHandle.EntityData.Leafs["interface-handle-name"] = types.YLeaf{"InterfaceHandleName", ifHandle.InterfaceHandleName}
    ifHandle.EntityData.Leafs["state-transition-count"] = types.YLeaf{"StateTransitionCount", ifHandle.StateTransitionCount}
    ifHandle.EntityData.Leafs["last-state-transition-time"] = types.YLeaf{"LastStateTransitionTime", ifHandle.LastStateTransitionTime}
    ifHandle.EntityData.Leafs["is-dampening-enabled"] = types.YLeaf{"IsDampeningEnabled", ifHandle.IsDampeningEnabled}
    ifHandle.EntityData.Leafs["half-life"] = types.YLeaf{"HalfLife", ifHandle.HalfLife}
    ifHandle.EntityData.Leafs["reuse-threshold"] = types.YLeaf{"ReuseThreshold", ifHandle.ReuseThreshold}
    ifHandle.EntityData.Leafs["suppress-threshold"] = types.YLeaf{"SuppressThreshold", ifHandle.SuppressThreshold}
    ifHandle.EntityData.Leafs["maximum-suppress-time"] = types.YLeaf{"MaximumSuppressTime", ifHandle.MaximumSuppressTime}
    ifHandle.EntityData.Leafs["restart-penalty"] = types.YLeaf{"RestartPenalty", ifHandle.RestartPenalty}
    return &(ifHandle.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening_
// Interface dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (interfaceDampening_ *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening_) GetEntityData() *types.CommonEntityData {
    interfaceDampening_.EntityData.YFilter = interfaceDampening_.YFilter
    interfaceDampening_.EntityData.YangName = "interface-dampening"
    interfaceDampening_.EntityData.BundleName = "cisco_ios_xr"
    interfaceDampening_.EntityData.ParentYangName = "if-handle"
    interfaceDampening_.EntityData.SegmentPath = "interface-dampening"
    interfaceDampening_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDampening_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDampening_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDampening_.EntityData.Children = make(map[string]types.YChild)
    interfaceDampening_.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDampening_.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", interfaceDampening_.Penalty}
    interfaceDampening_.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", interfaceDampening_.IsSuppressedEnabled}
    interfaceDampening_.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", interfaceDampening_.SecondsRemaining}
    interfaceDampening_.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", interfaceDampening_.Flaps}
    interfaceDampening_.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceDampening_.State}
    return &(interfaceDampening_.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetEntityData() *types.CommonEntityData {
    capsulation.EntityData.YFilter = capsulation.YFilter
    capsulation.EntityData.YangName = "capsulation"
    capsulation.EntityData.BundleName = "cisco_ios_xr"
    capsulation.EntityData.ParentYangName = "if-handle"
    capsulation.EntityData.SegmentPath = "capsulation"
    capsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulation.EntityData.Children = make(map[string]types.YChild)
    capsulation.EntityData.Children["capsulation-dampening"] = types.YChild{"CapsulationDampening", &capsulation.CapsulationDampening}
    capsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulation.EntityData.Leafs["capsulation-number"] = types.YLeaf{"CapsulationNumber", capsulation.CapsulationNumber}
    return &(capsulation.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetEntityData() *types.CommonEntityData {
    capsulationDampening.EntityData.YFilter = capsulationDampening.YFilter
    capsulationDampening.EntityData.YangName = "capsulation-dampening"
    capsulationDampening.EntityData.BundleName = "cisco_ios_xr"
    capsulationDampening.EntityData.ParentYangName = "capsulation"
    capsulationDampening.EntityData.SegmentPath = "capsulation-dampening"
    capsulationDampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulationDampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulationDampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulationDampening.EntityData.Children = make(map[string]types.YChild)
    capsulationDampening.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulationDampening.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", capsulationDampening.Penalty}
    capsulationDampening.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", capsulationDampening.IsSuppressedEnabled}
    capsulationDampening.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", capsulationDampening.SecondsRemaining}
    capsulationDampening.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", capsulationDampening.Flaps}
    capsulationDampening.EntityData.Leafs["state"] = types.YLeaf{"State", capsulationDampening.State}
    return &(capsulationDampening.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces
// Table of interfaces for which dampening info
// can be queried
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening info for the interface. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_.
    Interface_ []InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "dampening"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface
// Dampening info for the interface
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // The number of times the state has changed. The type is interface{} with
    // range: 0..4294967295.
    StateTransitionCount interface{}

    // The time elasped after the last state transition. The type is interface{}
    // with range: 0..4294967295.
    LastStateTransitionTime interface{}

    // Flag showing if dampening is enabled. The type is bool.
    IsDampeningEnabled interface{}

    // Configured decay half life in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    HalfLife interface{}

    // Configured reuse threshold. The type is interface{} with range:
    // 0..4294967295.
    ReuseThreshold interface{}

    // Value of suppress threshold. The type is interface{} with range:
    // 0..4294967295.
    SuppressThreshold interface{}

    // Maximum suppress time in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    MaximumSuppressTime interface{}

    // Configured restart penalty. The type is interface{} with range:
    // 0..4294967295.
    RestartPenalty interface{}

    // Interface dampening.
    InterfaceDampening InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening_

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation.
    Capsulation []InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["interface-dampening"] = types.YChild{"InterfaceDampening", &self.InterfaceDampening}
    self.EntityData.Children["capsulation"] = types.YChild{"Capsulation", nil}
    for i := range self.Capsulation {
        self.EntityData.Children[types.GetSegmentPath(&self.Capsulation[i])] = types.YChild{"Capsulation", &self.Capsulation[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["state-transition-count"] = types.YLeaf{"StateTransitionCount", self.StateTransitionCount}
    self.EntityData.Leafs["last-state-transition-time"] = types.YLeaf{"LastStateTransitionTime", self.LastStateTransitionTime}
    self.EntityData.Leafs["is-dampening-enabled"] = types.YLeaf{"IsDampeningEnabled", self.IsDampeningEnabled}
    self.EntityData.Leafs["half-life"] = types.YLeaf{"HalfLife", self.HalfLife}
    self.EntityData.Leafs["reuse-threshold"] = types.YLeaf{"ReuseThreshold", self.ReuseThreshold}
    self.EntityData.Leafs["suppress-threshold"] = types.YLeaf{"SuppressThreshold", self.SuppressThreshold}
    self.EntityData.Leafs["maximum-suppress-time"] = types.YLeaf{"MaximumSuppressTime", self.MaximumSuppressTime}
    self.EntityData.Leafs["restart-penalty"] = types.YLeaf{"RestartPenalty", self.RestartPenalty}
    return &(self.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening_
// Interface dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (interfaceDampening_ *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening_) GetEntityData() *types.CommonEntityData {
    interfaceDampening_.EntityData.YFilter = interfaceDampening_.YFilter
    interfaceDampening_.EntityData.YangName = "interface-dampening"
    interfaceDampening_.EntityData.BundleName = "cisco_ios_xr"
    interfaceDampening_.EntityData.ParentYangName = "interface"
    interfaceDampening_.EntityData.SegmentPath = "interface-dampening"
    interfaceDampening_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDampening_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDampening_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDampening_.EntityData.Children = make(map[string]types.YChild)
    interfaceDampening_.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceDampening_.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", interfaceDampening_.Penalty}
    interfaceDampening_.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", interfaceDampening_.IsSuppressedEnabled}
    interfaceDampening_.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", interfaceDampening_.SecondsRemaining}
    interfaceDampening_.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", interfaceDampening_.Flaps}
    interfaceDampening_.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceDampening_.State}
    return &(interfaceDampening_.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetEntityData() *types.CommonEntityData {
    capsulation.EntityData.YFilter = capsulation.YFilter
    capsulation.EntityData.YangName = "capsulation"
    capsulation.EntityData.BundleName = "cisco_ios_xr"
    capsulation.EntityData.ParentYangName = "interface"
    capsulation.EntityData.SegmentPath = "capsulation"
    capsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulation.EntityData.Children = make(map[string]types.YChild)
    capsulation.EntityData.Children["capsulation-dampening"] = types.YChild{"CapsulationDampening", &capsulation.CapsulationDampening}
    capsulation.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulation.EntityData.Leafs["capsulation-number"] = types.YLeaf{"CapsulationNumber", capsulation.CapsulationNumber}
    return &(capsulation.EntityData)
}

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampening penalty of the interface. The type is interface{} with range:
    // 0..4294967295.
    Penalty interface{}

    // Flag showing if state is suppressed. The type is bool.
    IsSuppressedEnabled interface{}

    // Remaining period of suppression in secs. The type is interface{} with
    // range: 0..4294967295. Units are second.
    SecondsRemaining interface{}

    // Number of underlying state flaps. The type is interface{} with range:
    // 0..4294967295.
    Flaps interface{}

    // Underlying state of the node. The type is ImStateEnum.
    State interface{}
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetEntityData() *types.CommonEntityData {
    capsulationDampening.EntityData.YFilter = capsulationDampening.YFilter
    capsulationDampening.EntityData.YangName = "capsulation-dampening"
    capsulationDampening.EntityData.BundleName = "cisco_ios_xr"
    capsulationDampening.EntityData.ParentYangName = "capsulation"
    capsulationDampening.EntityData.SegmentPath = "capsulation-dampening"
    capsulationDampening.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capsulationDampening.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capsulationDampening.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capsulationDampening.EntityData.Children = make(map[string]types.YChild)
    capsulationDampening.EntityData.Leafs = make(map[string]types.YLeaf)
    capsulationDampening.EntityData.Leafs["penalty"] = types.YLeaf{"Penalty", capsulationDampening.Penalty}
    capsulationDampening.EntityData.Leafs["is-suppressed-enabled"] = types.YLeaf{"IsSuppressedEnabled", capsulationDampening.IsSuppressedEnabled}
    capsulationDampening.EntityData.Leafs["seconds-remaining"] = types.YLeaf{"SecondsRemaining", capsulationDampening.SecondsRemaining}
    capsulationDampening.EntityData.Leafs["flaps"] = types.YLeaf{"Flaps", capsulationDampening.Flaps}
    capsulationDampening.EntityData.Leafs["state"] = types.YLeaf{"State", capsulationDampening.State}
    return &(capsulationDampening.EntityData)
}

// InterfaceProperties
// interface properties
type InterfaceProperties struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for interfaces.
    DataNodes InterfaceProperties_DataNodes
}

func (interfaceProperties *InterfaceProperties) GetEntityData() *types.CommonEntityData {
    interfaceProperties.EntityData.YFilter = interfaceProperties.YFilter
    interfaceProperties.EntityData.YangName = "interface-properties"
    interfaceProperties.EntityData.BundleName = "cisco_ios_xr"
    interfaceProperties.EntityData.ParentYangName = "Cisco-IOS-XR-ifmgr-oper"
    interfaceProperties.EntityData.SegmentPath = "Cisco-IOS-XR-ifmgr-oper:interface-properties"
    interfaceProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProperties.EntityData.Children = make(map[string]types.YChild)
    interfaceProperties.EntityData.Children["data-nodes"] = types.YChild{"DataNodes", &interfaceProperties.DataNodes}
    interfaceProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceProperties.EntityData)
}

// InterfaceProperties_DataNodes
// Operational data for interfaces
type InterfaceProperties_DataNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The location of a (D)RP in the same LR as the interface being queried. The
    // type is slice of InterfaceProperties_DataNodes_DataNode.
    DataNode []InterfaceProperties_DataNodes_DataNode
}

func (dataNodes *InterfaceProperties_DataNodes) GetEntityData() *types.CommonEntityData {
    dataNodes.EntityData.YFilter = dataNodes.YFilter
    dataNodes.EntityData.YangName = "data-nodes"
    dataNodes.EntityData.BundleName = "cisco_ios_xr"
    dataNodes.EntityData.ParentYangName = "interface-properties"
    dataNodes.EntityData.SegmentPath = "data-nodes"
    dataNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataNodes.EntityData.Children = make(map[string]types.YChild)
    dataNodes.EntityData.Children["data-node"] = types.YChild{"DataNode", nil}
    for i := range dataNodes.DataNode {
        dataNodes.EntityData.Children[types.GetSegmentPath(&dataNodes.DataNode[i])] = types.YChild{"DataNode", &dataNodes.DataNode[i]}
    }
    dataNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataNodes.EntityData)
}

// InterfaceProperties_DataNodes_DataNode
// The location of a (D)RP in the same LR as the
// interface being queried
type InterfaceProperties_DataNodes_DataNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The location of the (D)RP. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    DataNodeName interface{}

    // Location-specific view of interface operational data.
    Locationviews InterfaceProperties_DataNodes_DataNode_Locationviews

    // Partially qualified Location-specific view of interface operational data.
    PqNodeLocations InterfaceProperties_DataNodes_DataNode_PqNodeLocations

    // System-wide view of interface operational data.
    SystemView InterfaceProperties_DataNodes_DataNode_SystemView
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetEntityData() *types.CommonEntityData {
    dataNode.EntityData.YFilter = dataNode.YFilter
    dataNode.EntityData.YangName = "data-node"
    dataNode.EntityData.BundleName = "cisco_ios_xr"
    dataNode.EntityData.ParentYangName = "data-nodes"
    dataNode.EntityData.SegmentPath = "data-node" + "[data-node-name='" + fmt.Sprintf("%v", dataNode.DataNodeName) + "']"
    dataNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataNode.EntityData.Children = make(map[string]types.YChild)
    dataNode.EntityData.Children["locationviews"] = types.YChild{"Locationviews", &dataNode.Locationviews}
    dataNode.EntityData.Children["pq-node-locations"] = types.YChild{"PqNodeLocations", &dataNode.PqNodeLocations}
    dataNode.EntityData.Children["system-view"] = types.YChild{"SystemView", &dataNode.SystemView}
    dataNode.EntityData.Leafs = make(map[string]types.YLeaf)
    dataNode.EntityData.Leafs["data-node-name"] = types.YLeaf{"DataNodeName", dataNode.DataNodeName}
    return &(dataNode.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_Locationviews
// Location-specific view of interface
// operational data
type InterfaceProperties_DataNodes_DataNode_Locationviews struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers on a particular node.
    // The type is slice of
    // InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview.
    Locationview []InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetEntityData() *types.CommonEntityData {
    locationviews.EntityData.YFilter = locationviews.YFilter
    locationviews.EntityData.YangName = "locationviews"
    locationviews.EntityData.BundleName = "cisco_ios_xr"
    locationviews.EntityData.ParentYangName = "data-node"
    locationviews.EntityData.SegmentPath = "locationviews"
    locationviews.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationviews.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationviews.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationviews.EntityData.Children = make(map[string]types.YChild)
    locationviews.EntityData.Children["locationview"] = types.YChild{"Locationview", nil}
    for i := range locationviews.Locationview {
        locationviews.EntityData.Children[types.GetSegmentPath(&locationviews.Locationview[i])] = types.YChild{"Locationview", &locationviews.Locationview[i]}
    }
    locationviews.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(locationviews.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview
// Operational data for all interfaces and
// controllers on a particular node
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The location to filter on. The type is string with
    // pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    LocationviewName interface{}

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetEntityData() *types.CommonEntityData {
    locationview.EntityData.YFilter = locationview.YFilter
    locationview.EntityData.YangName = "locationview"
    locationview.EntityData.BundleName = "cisco_ios_xr"
    locationview.EntityData.ParentYangName = "locationviews"
    locationview.EntityData.SegmentPath = "locationview" + "[locationview-name='" + fmt.Sprintf("%v", locationview.LocationviewName) + "']"
    locationview.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locationview.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locationview.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locationview.EntityData.Children = make(map[string]types.YChild)
    locationview.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &locationview.Interfaces}
    locationview.EntityData.Leafs = make(map[string]types.YLeaf)
    locationview.EntityData.Leafs["locationview-name"] = types.YLeaf{"LocationviewName", locationview.LocationviewName}
    return &(locationview.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface_.
    Interface_ []InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "locationview"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Parent Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type_ interface{}

    // Operational state. The type is ImStateEnum.
    State interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualState interface{}

    // Line protocol state. The type is ImStateEnum.
    LineState interface{}

    // Line protocol state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualLineState interface{}

    // Interface encapsulation. The type is string.
    Encapsulation interface{}

    // Interface encapsulation description string. The type is string with length:
    // 0..32.
    EncapsulationTypeString interface{}

    // MTU in bytes. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    Mtu interface{}

    // Subif MTU overhead. The type is interface{} with range: 0..4294967295.
    SubInterfaceMtuOverhead interface{}

    // L2 transport. The type is bool.
    L2Transport interface{}

    // Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    Bandwidth interface{}
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", self.ParentInterface}
    self.EntityData.Leafs["type"] = types.YLeaf{"Type_", self.Type_}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["actual-state"] = types.YLeaf{"ActualState", self.ActualState}
    self.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", self.LineState}
    self.EntityData.Leafs["actual-line-state"] = types.YLeaf{"ActualLineState", self.ActualLineState}
    self.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", self.Encapsulation}
    self.EntityData.Leafs["encapsulation-type-string"] = types.YLeaf{"EncapsulationTypeString", self.EncapsulationTypeString}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["sub-interface-mtu-overhead"] = types.YLeaf{"SubInterfaceMtuOverhead", self.SubInterfaceMtuOverhead}
    self.EntityData.Leafs["l2-transport"] = types.YLeaf{"L2Transport", self.L2Transport}
    self.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", self.Bandwidth}
    return &(self.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations
// Partially qualified Location-specific view of
// interface operational data
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers on a particular
    // pq_node. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation.
    PqNodeLocation []InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetEntityData() *types.CommonEntityData {
    pqNodeLocations.EntityData.YFilter = pqNodeLocations.YFilter
    pqNodeLocations.EntityData.YangName = "pq-node-locations"
    pqNodeLocations.EntityData.BundleName = "cisco_ios_xr"
    pqNodeLocations.EntityData.ParentYangName = "data-node"
    pqNodeLocations.EntityData.SegmentPath = "pq-node-locations"
    pqNodeLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pqNodeLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pqNodeLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pqNodeLocations.EntityData.Children = make(map[string]types.YChild)
    pqNodeLocations.EntityData.Children["pq-node-location"] = types.YChild{"PqNodeLocation", nil}
    for i := range pqNodeLocations.PqNodeLocation {
        pqNodeLocations.EntityData.Children[types.GetSegmentPath(&pqNodeLocations.PqNodeLocation[i])] = types.YChild{"PqNodeLocation", &pqNodeLocations.PqNodeLocation[i]}
    }
    pqNodeLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pqNodeLocations.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation
// Operational data for all interfaces and
// controllers on a particular pq_node
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The partially qualified location to filter on. The
    // type is string with pattern:
    // b'((([a-zA-Z0-9_]*\\d+)|(\\*))/){2}(([a-zA-Z0-9_]*\\d+)|(\\*))'.
    PqNodeName interface{}

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetEntityData() *types.CommonEntityData {
    pqNodeLocation.EntityData.YFilter = pqNodeLocation.YFilter
    pqNodeLocation.EntityData.YangName = "pq-node-location"
    pqNodeLocation.EntityData.BundleName = "cisco_ios_xr"
    pqNodeLocation.EntityData.ParentYangName = "pq-node-locations"
    pqNodeLocation.EntityData.SegmentPath = "pq-node-location" + "[pq-node-name='" + fmt.Sprintf("%v", pqNodeLocation.PqNodeName) + "']"
    pqNodeLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pqNodeLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pqNodeLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pqNodeLocation.EntityData.Children = make(map[string]types.YChild)
    pqNodeLocation.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &pqNodeLocation.Interfaces}
    pqNodeLocation.EntityData.Leafs = make(map[string]types.YLeaf)
    pqNodeLocation.EntityData.Leafs["pq-node-name"] = types.YLeaf{"PqNodeName", pqNodeLocation.PqNodeName}
    return &(pqNodeLocation.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface_.
    Interface_ []InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "pq-node-location"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Parent Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type_ interface{}

    // Operational state. The type is ImStateEnum.
    State interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualState interface{}

    // Line protocol state. The type is ImStateEnum.
    LineState interface{}

    // Line protocol state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualLineState interface{}

    // Interface encapsulation. The type is string.
    Encapsulation interface{}

    // Interface encapsulation description string. The type is string with length:
    // 0..32.
    EncapsulationTypeString interface{}

    // MTU in bytes. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    Mtu interface{}

    // Subif MTU overhead. The type is interface{} with range: 0..4294967295.
    SubInterfaceMtuOverhead interface{}

    // L2 transport. The type is bool.
    L2Transport interface{}

    // Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    Bandwidth interface{}
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", self.ParentInterface}
    self.EntityData.Leafs["type"] = types.YLeaf{"Type_", self.Type_}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["actual-state"] = types.YLeaf{"ActualState", self.ActualState}
    self.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", self.LineState}
    self.EntityData.Leafs["actual-line-state"] = types.YLeaf{"ActualLineState", self.ActualLineState}
    self.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", self.Encapsulation}
    self.EntityData.Leafs["encapsulation-type-string"] = types.YLeaf{"EncapsulationTypeString", self.EncapsulationTypeString}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["sub-interface-mtu-overhead"] = types.YLeaf{"SubInterfaceMtuOverhead", self.SubInterfaceMtuOverhead}
    self.EntityData.Leafs["l2-transport"] = types.YLeaf{"L2Transport", self.L2Transport}
    self.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", self.Bandwidth}
    return &(self.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_SystemView
// System-wide view of interface operational data
type InterfaceProperties_DataNodes_DataNode_SystemView struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetEntityData() *types.CommonEntityData {
    systemView.EntityData.YFilter = systemView.YFilter
    systemView.EntityData.YangName = "system-view"
    systemView.EntityData.BundleName = "cisco_ios_xr"
    systemView.EntityData.ParentYangName = "data-node"
    systemView.EntityData.SegmentPath = "system-view"
    systemView.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemView.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemView.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemView.EntityData.Children = make(map[string]types.YChild)
    systemView.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &systemView.Interfaces}
    systemView.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(systemView.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface_.
    Interface_ []InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "system-view"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Parent Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type_ interface{}

    // Operational state. The type is ImStateEnum.
    State interface{}

    // Operational state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualState interface{}

    // Line protocol state. The type is ImStateEnum.
    LineState interface{}

    // Line protocol state with no translation of error disable or shutdown. The
    // type is ImStateEnum.
    ActualLineState interface{}

    // Interface encapsulation. The type is string.
    Encapsulation interface{}

    // Interface encapsulation description string. The type is string with length:
    // 0..32.
    EncapsulationTypeString interface{}

    // MTU in bytes. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    Mtu interface{}

    // Subif MTU overhead. The type is interface{} with range: 0..4294967295.
    SubInterfaceMtuOverhead interface{}

    // L2 transport. The type is bool.
    L2Transport interface{}

    // Interface bandwidth (Kb/s). The type is interface{} with range:
    // 0..4294967295.
    Bandwidth interface{}
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["parent-interface"] = types.YLeaf{"ParentInterface", self.ParentInterface}
    self.EntityData.Leafs["type"] = types.YLeaf{"Type_", self.Type_}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["actual-state"] = types.YLeaf{"ActualState", self.ActualState}
    self.EntityData.Leafs["line-state"] = types.YLeaf{"LineState", self.LineState}
    self.EntityData.Leafs["actual-line-state"] = types.YLeaf{"ActualLineState", self.ActualLineState}
    self.EntityData.Leafs["encapsulation"] = types.YLeaf{"Encapsulation", self.Encapsulation}
    self.EntityData.Leafs["encapsulation-type-string"] = types.YLeaf{"EncapsulationTypeString", self.EncapsulationTypeString}
    self.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", self.Mtu}
    self.EntityData.Leafs["sub-interface-mtu-overhead"] = types.YLeaf{"SubInterfaceMtuOverhead", self.SubInterfaceMtuOverhead}
    self.EntityData.Leafs["l2-transport"] = types.YLeaf{"L2Transport", self.L2Transport}
    self.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", self.Bandwidth}
    return &(self.EntityData)
}

