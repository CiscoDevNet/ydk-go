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
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface list for which dampening info is available.
    Interfaces InterfaceDampening_Interfaces

    // The location of the interface(s) being queried.
    Nodes InterfaceDampening_Nodes
}

func (interfaceDampening *InterfaceDampening) GetFilter() yfilter.YFilter { return interfaceDampening.YFilter }

func (interfaceDampening *InterfaceDampening) SetFilter(yf yfilter.YFilter) { interfaceDampening.YFilter = yf }

func (interfaceDampening *InterfaceDampening) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (interfaceDampening *InterfaceDampening) GetSegmentPath() string {
    return "Cisco-IOS-XR-ifmgr-oper:interface-dampening"
}

func (interfaceDampening *InterfaceDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceDampening.Interfaces
    }
    if childYangName == "nodes" {
        return &interfaceDampening.Nodes
    }
    return nil
}

func (interfaceDampening *InterfaceDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceDampening.Interfaces
    children["nodes"] = &interfaceDampening.Nodes
    return children
}

func (interfaceDampening *InterfaceDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceDampening *InterfaceDampening) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDampening *InterfaceDampening) GetYangName() string { return "interface-dampening" }

func (interfaceDampening *InterfaceDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDampening *InterfaceDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDampening *InterfaceDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDampening *InterfaceDampening) SetParent(parent types.Entity) { interfaceDampening.parent = parent }

func (interfaceDampening *InterfaceDampening) GetParent() types.Entity { return interfaceDampening.parent }

func (interfaceDampening *InterfaceDampening) GetParentYangName() string { return "Cisco-IOS-XR-ifmgr-oper" }

// InterfaceDampening_Interfaces
// The interface list for which dampening info is
// available
type InterfaceDampening_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface for which dampening info is being queried. The type is slice
    // of InterfaceDampening_Interfaces_Interface.
    Interface []InterfaceDampening_Interfaces_Interface
}

func (interfaces *InterfaceDampening_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InterfaceDampening_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InterfaceDampening_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InterfaceDampening_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InterfaceDampening_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InterfaceDampening_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InterfaceDampening_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InterfaceDampening_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InterfaceDampening_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InterfaceDampening_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InterfaceDampening_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InterfaceDampening_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InterfaceDampening_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InterfaceDampening_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InterfaceDampening_Interfaces) GetParentYangName() string { return "interface-dampening" }

// InterfaceDampening_Interfaces_Interface
// The interface for which dampening info is being
// queried
type InterfaceDampening_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Dampening info for the interface.
    IfDampening InterfaceDampening_Interfaces_Interface_IfDampening
}

func (self *InterfaceDampening_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfaceDampening_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfaceDampening_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "if-dampening" { return "IfDampening" }
    return ""
}

func (self *InterfaceDampening_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InterfaceDampening_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-dampening" {
        return &self.IfDampening
    }
    return nil
}

func (self *InterfaceDampening_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["if-dampening"] = &self.IfDampening
    return children
}

func (self *InterfaceDampening_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *InterfaceDampening_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InterfaceDampening_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InterfaceDampening_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InterfaceDampening_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InterfaceDampening_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InterfaceDampening_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfaceDampening_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfaceDampening_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// InterfaceDampening_Interfaces_Interface_IfDampening
// Dampening info for the interface
type InterfaceDampening_Interfaces_Interface_IfDampening struct {
    parent types.Entity
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
    InterfaceDampening InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation.
    Capsulation []InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetFilter() yfilter.YFilter { return ifDampening.YFilter }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) SetFilter(yf yfilter.YFilter) { ifDampening.YFilter = yf }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetGoName(yname string) string {
    if yname == "state-transition-count" { return "StateTransitionCount" }
    if yname == "last-state-transition-time" { return "LastStateTransitionTime" }
    if yname == "is-dampening-enabled" { return "IsDampeningEnabled" }
    if yname == "half-life" { return "HalfLife" }
    if yname == "reuse-threshold" { return "ReuseThreshold" }
    if yname == "suppress-threshold" { return "SuppressThreshold" }
    if yname == "maximum-suppress-time" { return "MaximumSuppressTime" }
    if yname == "restart-penalty" { return "RestartPenalty" }
    if yname == "interface-dampening" { return "InterfaceDampening" }
    if yname == "capsulation" { return "Capsulation" }
    return ""
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetSegmentPath() string {
    return "if-dampening"
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-dampening" {
        return &ifDampening.InterfaceDampening
    }
    if childYangName == "capsulation" {
        for _, c := range ifDampening.Capsulation {
            if ifDampening.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation{}
        ifDampening.Capsulation = append(ifDampening.Capsulation, child)
        return &ifDampening.Capsulation[len(ifDampening.Capsulation)-1]
    }
    return nil
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-dampening"] = &ifDampening.InterfaceDampening
    for i := range ifDampening.Capsulation {
        children[ifDampening.Capsulation[i].GetSegmentPath()] = &ifDampening.Capsulation[i]
    }
    return children
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state-transition-count"] = ifDampening.StateTransitionCount
    leafs["last-state-transition-time"] = ifDampening.LastStateTransitionTime
    leafs["is-dampening-enabled"] = ifDampening.IsDampeningEnabled
    leafs["half-life"] = ifDampening.HalfLife
    leafs["reuse-threshold"] = ifDampening.ReuseThreshold
    leafs["suppress-threshold"] = ifDampening.SuppressThreshold
    leafs["maximum-suppress-time"] = ifDampening.MaximumSuppressTime
    leafs["restart-penalty"] = ifDampening.RestartPenalty
    return leafs
}

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetBundleName() string { return "cisco_ios_xr" }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetYangName() string { return "if-dampening" }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) SetParent(parent types.Entity) { ifDampening.parent = parent }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetParent() types.Entity { return ifDampening.parent }

func (ifDampening *InterfaceDampening_Interfaces_Interface_IfDampening) GetParentYangName() string { return "interface" }

// InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening
// Interface dampening
type InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening struct {
    parent types.Entity
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

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetFilter() yfilter.YFilter { return interfaceDampening.YFilter }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) SetFilter(yf yfilter.YFilter) { interfaceDampening.YFilter = yf }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetSegmentPath() string {
    return "interface-dampening"
}

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = interfaceDampening.Penalty
    leafs["is-suppressed-enabled"] = interfaceDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = interfaceDampening.SecondsRemaining
    leafs["flaps"] = interfaceDampening.Flaps
    leafs["state"] = interfaceDampening.State
    return leafs
}

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetYangName() string { return "interface-dampening" }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) SetParent(parent types.Entity) { interfaceDampening.parent = parent }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetParent() types.Entity { return interfaceDampening.parent }

func (interfaceDampening *InterfaceDampening_Interfaces_Interface_IfDampening_InterfaceDampening) GetParentYangName() string { return "if-dampening" }

// InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetFilter() yfilter.YFilter { return capsulation.YFilter }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) SetFilter(yf yfilter.YFilter) { capsulation.YFilter = yf }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetGoName(yname string) string {
    if yname == "capsulation-number" { return "CapsulationNumber" }
    if yname == "capsulation-dampening" { return "CapsulationDampening" }
    return ""
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetSegmentPath() string {
    return "capsulation"
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capsulation-dampening" {
        return &capsulation.CapsulationDampening
    }
    return nil
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capsulation-dampening"] = &capsulation.CapsulationDampening
    return children
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["capsulation-number"] = capsulation.CapsulationNumber
    return leafs
}

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetBundleName() string { return "cisco_ios_xr" }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetYangName() string { return "capsulation" }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) SetParent(parent types.Entity) { capsulation.parent = parent }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetParent() types.Entity { return capsulation.parent }

func (capsulation *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation) GetParentYangName() string { return "if-dampening" }

// InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening struct {
    parent types.Entity
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

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetFilter() yfilter.YFilter { return capsulationDampening.YFilter }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) SetFilter(yf yfilter.YFilter) { capsulationDampening.YFilter = yf }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetSegmentPath() string {
    return "capsulation-dampening"
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = capsulationDampening.Penalty
    leafs["is-suppressed-enabled"] = capsulationDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = capsulationDampening.SecondsRemaining
    leafs["flaps"] = capsulationDampening.Flaps
    leafs["state"] = capsulationDampening.State
    return leafs
}

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetBundleName() string { return "cisco_ios_xr" }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetYangName() string { return "capsulation-dampening" }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) SetParent(parent types.Entity) { capsulationDampening.parent = parent }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetParent() types.Entity { return capsulationDampening.parent }

func (capsulationDampening *InterfaceDampening_Interfaces_Interface_IfDampening_Capsulation_CapsulationDampening) GetParentYangName() string { return "capsulation" }

// InterfaceDampening_Nodes
// The location of the interface(s) being queried
type InterfaceDampening_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The location of the interface(s) being queried. The type is slice of
    // InterfaceDampening_Nodes_Node.
    Node []InterfaceDampening_Nodes_Node
}

func (nodes *InterfaceDampening_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *InterfaceDampening_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *InterfaceDampening_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *InterfaceDampening_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *InterfaceDampening_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *InterfaceDampening_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *InterfaceDampening_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *InterfaceDampening_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *InterfaceDampening_Nodes) GetYangName() string { return "nodes" }

func (nodes *InterfaceDampening_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *InterfaceDampening_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *InterfaceDampening_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *InterfaceDampening_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *InterfaceDampening_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *InterfaceDampening_Nodes) GetParentYangName() string { return "interface-dampening" }

// InterfaceDampening_Nodes_Node
// The location of the interface(s) being queried
type InterfaceDampening_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The location of the interface(s). The type is
    // string with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Show details for the interfaces.
    Show InterfaceDampening_Nodes_Node_Show
}

func (node *InterfaceDampening_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *InterfaceDampening_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *InterfaceDampening_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "show" { return "Show" }
    return ""
}

func (node *InterfaceDampening_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *InterfaceDampening_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "show" {
        return &node.Show
    }
    return nil
}

func (node *InterfaceDampening_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["show"] = &node.Show
    return children
}

func (node *InterfaceDampening_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *InterfaceDampening_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *InterfaceDampening_Nodes_Node) GetYangName() string { return "node" }

func (node *InterfaceDampening_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *InterfaceDampening_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *InterfaceDampening_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *InterfaceDampening_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *InterfaceDampening_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *InterfaceDampening_Nodes_Node) GetParentYangName() string { return "nodes" }

// InterfaceDampening_Nodes_Node_Show
// Show details for the interfaces
type InterfaceDampening_Nodes_Node_Show struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampening information of the interface(s) being queried.
    Dampening InterfaceDampening_Nodes_Node_Show_Dampening
}

func (show *InterfaceDampening_Nodes_Node_Show) GetFilter() yfilter.YFilter { return show.YFilter }

func (show *InterfaceDampening_Nodes_Node_Show) SetFilter(yf yfilter.YFilter) { show.YFilter = yf }

func (show *InterfaceDampening_Nodes_Node_Show) GetGoName(yname string) string {
    if yname == "dampening" { return "Dampening" }
    return ""
}

func (show *InterfaceDampening_Nodes_Node_Show) GetSegmentPath() string {
    return "show"
}

func (show *InterfaceDampening_Nodes_Node_Show) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dampening" {
        return &show.Dampening
    }
    return nil
}

func (show *InterfaceDampening_Nodes_Node_Show) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dampening"] = &show.Dampening
    return children
}

func (show *InterfaceDampening_Nodes_Node_Show) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (show *InterfaceDampening_Nodes_Node_Show) GetBundleName() string { return "cisco_ios_xr" }

func (show *InterfaceDampening_Nodes_Node_Show) GetYangName() string { return "show" }

func (show *InterfaceDampening_Nodes_Node_Show) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (show *InterfaceDampening_Nodes_Node_Show) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (show *InterfaceDampening_Nodes_Node_Show) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (show *InterfaceDampening_Nodes_Node_Show) SetParent(parent types.Entity) { show.parent = parent }

func (show *InterfaceDampening_Nodes_Node_Show) GetParent() types.Entity { return show.parent }

func (show *InterfaceDampening_Nodes_Node_Show) GetParentYangName() string { return "node" }

// InterfaceDampening_Nodes_Node_Show_Dampening
// Dampening information of the interface(s)
// being queried
type InterfaceDampening_Nodes_Node_Show_Dampening struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface handle for which dampening info queried.
    IfHandles InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles

    // Table of interfaces for which dampening info can be queried.
    Interfaces InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetFilter() yfilter.YFilter { return dampening.YFilter }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) SetFilter(yf yfilter.YFilter) { dampening.YFilter = yf }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetGoName(yname string) string {
    if yname == "if-handles" { return "IfHandles" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetSegmentPath() string {
    return "dampening"
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-handles" {
        return &dampening.IfHandles
    }
    if childYangName == "interfaces" {
        return &dampening.Interfaces
    }
    return nil
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["if-handles"] = &dampening.IfHandles
    children["interfaces"] = &dampening.Interfaces
    return children
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetBundleName() string { return "cisco_ios_xr" }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetYangName() string { return "dampening" }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) SetParent(parent types.Entity) { dampening.parent = parent }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetParent() types.Entity { return dampening.parent }

func (dampening *InterfaceDampening_Nodes_Node_Show_Dampening) GetParentYangName() string { return "show" }

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles
// Interface handle for which dampening info
// queried
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampening info for the interface handle. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle.
    IfHandle []InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetFilter() yfilter.YFilter { return ifHandles.YFilter }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) SetFilter(yf yfilter.YFilter) { ifHandles.YFilter = yf }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetGoName(yname string) string {
    if yname == "if-handle" { return "IfHandle" }
    return ""
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetSegmentPath() string {
    return "if-handles"
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-handle" {
        for _, c := range ifHandles.IfHandle {
            if ifHandles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle{}
        ifHandles.IfHandle = append(ifHandles.IfHandle, child)
        return &ifHandles.IfHandle[len(ifHandles.IfHandle)-1]
    }
    return nil
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ifHandles.IfHandle {
        children[ifHandles.IfHandle[i].GetSegmentPath()] = &ifHandles.IfHandle[i]
    }
    return children
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetBundleName() string { return "cisco_ios_xr" }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetYangName() string { return "if-handles" }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) SetParent(parent types.Entity) { ifHandles.parent = parent }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetParent() types.Entity { return ifHandles.parent }

func (ifHandles *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles) GetParentYangName() string { return "dampening" }

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle
// Dampening info for the interface handle
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface handle. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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
    InterfaceDampening InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation.
    Capsulation []InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetFilter() yfilter.YFilter { return ifHandle.YFilter }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) SetFilter(yf yfilter.YFilter) { ifHandle.YFilter = yf }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetGoName(yname string) string {
    if yname == "interface-handle-name" { return "InterfaceHandleName" }
    if yname == "state-transition-count" { return "StateTransitionCount" }
    if yname == "last-state-transition-time" { return "LastStateTransitionTime" }
    if yname == "is-dampening-enabled" { return "IsDampeningEnabled" }
    if yname == "half-life" { return "HalfLife" }
    if yname == "reuse-threshold" { return "ReuseThreshold" }
    if yname == "suppress-threshold" { return "SuppressThreshold" }
    if yname == "maximum-suppress-time" { return "MaximumSuppressTime" }
    if yname == "restart-penalty" { return "RestartPenalty" }
    if yname == "interface-dampening" { return "InterfaceDampening" }
    if yname == "capsulation" { return "Capsulation" }
    return ""
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetSegmentPath() string {
    return "if-handle" + "[interface-handle-name='" + fmt.Sprintf("%v", ifHandle.InterfaceHandleName) + "']"
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-dampening" {
        return &ifHandle.InterfaceDampening
    }
    if childYangName == "capsulation" {
        for _, c := range ifHandle.Capsulation {
            if ifHandle.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation{}
        ifHandle.Capsulation = append(ifHandle.Capsulation, child)
        return &ifHandle.Capsulation[len(ifHandle.Capsulation)-1]
    }
    return nil
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-dampening"] = &ifHandle.InterfaceDampening
    for i := range ifHandle.Capsulation {
        children[ifHandle.Capsulation[i].GetSegmentPath()] = &ifHandle.Capsulation[i]
    }
    return children
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-handle-name"] = ifHandle.InterfaceHandleName
    leafs["state-transition-count"] = ifHandle.StateTransitionCount
    leafs["last-state-transition-time"] = ifHandle.LastStateTransitionTime
    leafs["is-dampening-enabled"] = ifHandle.IsDampeningEnabled
    leafs["half-life"] = ifHandle.HalfLife
    leafs["reuse-threshold"] = ifHandle.ReuseThreshold
    leafs["suppress-threshold"] = ifHandle.SuppressThreshold
    leafs["maximum-suppress-time"] = ifHandle.MaximumSuppressTime
    leafs["restart-penalty"] = ifHandle.RestartPenalty
    return leafs
}

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetBundleName() string { return "cisco_ios_xr" }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetYangName() string { return "if-handle" }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) SetParent(parent types.Entity) { ifHandle.parent = parent }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetParent() types.Entity { return ifHandle.parent }

func (ifHandle *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle) GetParentYangName() string { return "if-handles" }

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening
// Interface dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening struct {
    parent types.Entity
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

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetFilter() yfilter.YFilter { return interfaceDampening.YFilter }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) SetFilter(yf yfilter.YFilter) { interfaceDampening.YFilter = yf }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetSegmentPath() string {
    return "interface-dampening"
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = interfaceDampening.Penalty
    leafs["is-suppressed-enabled"] = interfaceDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = interfaceDampening.SecondsRemaining
    leafs["flaps"] = interfaceDampening.Flaps
    leafs["state"] = interfaceDampening.State
    return leafs
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetYangName() string { return "interface-dampening" }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) SetParent(parent types.Entity) { interfaceDampening.parent = parent }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetParent() types.Entity { return interfaceDampening.parent }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_InterfaceDampening) GetParentYangName() string { return "if-handle" }

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetFilter() yfilter.YFilter { return capsulation.YFilter }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) SetFilter(yf yfilter.YFilter) { capsulation.YFilter = yf }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetGoName(yname string) string {
    if yname == "capsulation-number" { return "CapsulationNumber" }
    if yname == "capsulation-dampening" { return "CapsulationDampening" }
    return ""
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetSegmentPath() string {
    return "capsulation"
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capsulation-dampening" {
        return &capsulation.CapsulationDampening
    }
    return nil
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capsulation-dampening"] = &capsulation.CapsulationDampening
    return children
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["capsulation-number"] = capsulation.CapsulationNumber
    return leafs
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetBundleName() string { return "cisco_ios_xr" }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetYangName() string { return "capsulation" }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) SetParent(parent types.Entity) { capsulation.parent = parent }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetParent() types.Entity { return capsulation.parent }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation) GetParentYangName() string { return "if-handle" }

// InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening struct {
    parent types.Entity
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

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetFilter() yfilter.YFilter { return capsulationDampening.YFilter }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) SetFilter(yf yfilter.YFilter) { capsulationDampening.YFilter = yf }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetSegmentPath() string {
    return "capsulation-dampening"
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = capsulationDampening.Penalty
    leafs["is-suppressed-enabled"] = capsulationDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = capsulationDampening.SecondsRemaining
    leafs["flaps"] = capsulationDampening.Flaps
    leafs["state"] = capsulationDampening.State
    return leafs
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetBundleName() string { return "cisco_ios_xr" }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetYangName() string { return "capsulation-dampening" }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) SetParent(parent types.Entity) { capsulationDampening.parent = parent }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetParent() types.Entity { return capsulationDampening.parent }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_IfHandles_IfHandle_Capsulation_CapsulationDampening) GetParentYangName() string { return "capsulation" }

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces
// Table of interfaces for which dampening info
// can be queried
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampening info for the interface. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface.
    Interface []InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces) GetParentYangName() string { return "dampening" }

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface
// Dampening info for the interface
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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
    InterfaceDampening InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening

    // Dampening information for capsulations. The type is slice of
    // InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation.
    Capsulation []InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "state-transition-count" { return "StateTransitionCount" }
    if yname == "last-state-transition-time" { return "LastStateTransitionTime" }
    if yname == "is-dampening-enabled" { return "IsDampeningEnabled" }
    if yname == "half-life" { return "HalfLife" }
    if yname == "reuse-threshold" { return "ReuseThreshold" }
    if yname == "suppress-threshold" { return "SuppressThreshold" }
    if yname == "maximum-suppress-time" { return "MaximumSuppressTime" }
    if yname == "restart-penalty" { return "RestartPenalty" }
    if yname == "interface-dampening" { return "InterfaceDampening" }
    if yname == "capsulation" { return "Capsulation" }
    return ""
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-dampening" {
        return &self.InterfaceDampening
    }
    if childYangName == "capsulation" {
        for _, c := range self.Capsulation {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation{}
        self.Capsulation = append(self.Capsulation, child)
        return &self.Capsulation[len(self.Capsulation)-1]
    }
    return nil
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-dampening"] = &self.InterfaceDampening
    for i := range self.Capsulation {
        children[self.Capsulation[i].GetSegmentPath()] = &self.Capsulation[i]
    }
    return children
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["state-transition-count"] = self.StateTransitionCount
    leafs["last-state-transition-time"] = self.LastStateTransitionTime
    leafs["is-dampening-enabled"] = self.IsDampeningEnabled
    leafs["half-life"] = self.HalfLife
    leafs["reuse-threshold"] = self.ReuseThreshold
    leafs["suppress-threshold"] = self.SuppressThreshold
    leafs["maximum-suppress-time"] = self.MaximumSuppressTime
    leafs["restart-penalty"] = self.RestartPenalty
    return leafs
}

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening
// Interface dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening struct {
    parent types.Entity
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

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetFilter() yfilter.YFilter { return interfaceDampening.YFilter }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) SetFilter(yf yfilter.YFilter) { interfaceDampening.YFilter = yf }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetSegmentPath() string {
    return "interface-dampening"
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = interfaceDampening.Penalty
    leafs["is-suppressed-enabled"] = interfaceDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = interfaceDampening.SecondsRemaining
    leafs["flaps"] = interfaceDampening.Flaps
    leafs["state"] = interfaceDampening.State
    return leafs
}

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetYangName() string { return "interface-dampening" }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) SetParent(parent types.Entity) { interfaceDampening.parent = parent }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetParent() types.Entity { return interfaceDampening.parent }

func (interfaceDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_InterfaceDampening) GetParentYangName() string { return "interface" }

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation
// Dampening information for capsulations
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Capsulation number. The type is string.
    CapsulationNumber interface{}

    // Capsulation dampening.
    CapsulationDampening InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetFilter() yfilter.YFilter { return capsulation.YFilter }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) SetFilter(yf yfilter.YFilter) { capsulation.YFilter = yf }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetGoName(yname string) string {
    if yname == "capsulation-number" { return "CapsulationNumber" }
    if yname == "capsulation-dampening" { return "CapsulationDampening" }
    return ""
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetSegmentPath() string {
    return "capsulation"
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capsulation-dampening" {
        return &capsulation.CapsulationDampening
    }
    return nil
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["capsulation-dampening"] = &capsulation.CapsulationDampening
    return children
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["capsulation-number"] = capsulation.CapsulationNumber
    return leafs
}

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetBundleName() string { return "cisco_ios_xr" }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetYangName() string { return "capsulation" }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) SetParent(parent types.Entity) { capsulation.parent = parent }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetParent() types.Entity { return capsulation.parent }

func (capsulation *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation) GetParentYangName() string { return "interface" }

// InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening
// Capsulation dampening
type InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening struct {
    parent types.Entity
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

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetFilter() yfilter.YFilter { return capsulationDampening.YFilter }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) SetFilter(yf yfilter.YFilter) { capsulationDampening.YFilter = yf }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetGoName(yname string) string {
    if yname == "penalty" { return "Penalty" }
    if yname == "is-suppressed-enabled" { return "IsSuppressedEnabled" }
    if yname == "seconds-remaining" { return "SecondsRemaining" }
    if yname == "flaps" { return "Flaps" }
    if yname == "state" { return "State" }
    return ""
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetSegmentPath() string {
    return "capsulation-dampening"
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["penalty"] = capsulationDampening.Penalty
    leafs["is-suppressed-enabled"] = capsulationDampening.IsSuppressedEnabled
    leafs["seconds-remaining"] = capsulationDampening.SecondsRemaining
    leafs["flaps"] = capsulationDampening.Flaps
    leafs["state"] = capsulationDampening.State
    return leafs
}

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetBundleName() string { return "cisco_ios_xr" }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetYangName() string { return "capsulation-dampening" }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) SetParent(parent types.Entity) { capsulationDampening.parent = parent }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetParent() types.Entity { return capsulationDampening.parent }

func (capsulationDampening *InterfaceDampening_Nodes_Node_Show_Dampening_Interfaces_Interface_Capsulation_CapsulationDampening) GetParentYangName() string { return "capsulation" }

// InterfaceProperties
// interface properties
type InterfaceProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for interfaces.
    DataNodes InterfaceProperties_DataNodes
}

func (interfaceProperties *InterfaceProperties) GetFilter() yfilter.YFilter { return interfaceProperties.YFilter }

func (interfaceProperties *InterfaceProperties) SetFilter(yf yfilter.YFilter) { interfaceProperties.YFilter = yf }

func (interfaceProperties *InterfaceProperties) GetGoName(yname string) string {
    if yname == "data-nodes" { return "DataNodes" }
    return ""
}

func (interfaceProperties *InterfaceProperties) GetSegmentPath() string {
    return "Cisco-IOS-XR-ifmgr-oper:interface-properties"
}

func (interfaceProperties *InterfaceProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-nodes" {
        return &interfaceProperties.DataNodes
    }
    return nil
}

func (interfaceProperties *InterfaceProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data-nodes"] = &interfaceProperties.DataNodes
    return children
}

func (interfaceProperties *InterfaceProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceProperties *InterfaceProperties) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceProperties *InterfaceProperties) GetYangName() string { return "interface-properties" }

func (interfaceProperties *InterfaceProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceProperties *InterfaceProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceProperties *InterfaceProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceProperties *InterfaceProperties) SetParent(parent types.Entity) { interfaceProperties.parent = parent }

func (interfaceProperties *InterfaceProperties) GetParent() types.Entity { return interfaceProperties.parent }

func (interfaceProperties *InterfaceProperties) GetParentYangName() string { return "Cisco-IOS-XR-ifmgr-oper" }

// InterfaceProperties_DataNodes
// Operational data for interfaces
type InterfaceProperties_DataNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The location of a (D)RP in the same LR as the interface being queried. The
    // type is slice of InterfaceProperties_DataNodes_DataNode.
    DataNode []InterfaceProperties_DataNodes_DataNode
}

func (dataNodes *InterfaceProperties_DataNodes) GetFilter() yfilter.YFilter { return dataNodes.YFilter }

func (dataNodes *InterfaceProperties_DataNodes) SetFilter(yf yfilter.YFilter) { dataNodes.YFilter = yf }

func (dataNodes *InterfaceProperties_DataNodes) GetGoName(yname string) string {
    if yname == "data-node" { return "DataNode" }
    return ""
}

func (dataNodes *InterfaceProperties_DataNodes) GetSegmentPath() string {
    return "data-nodes"
}

func (dataNodes *InterfaceProperties_DataNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-node" {
        for _, c := range dataNodes.DataNode {
            if dataNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode{}
        dataNodes.DataNode = append(dataNodes.DataNode, child)
        return &dataNodes.DataNode[len(dataNodes.DataNode)-1]
    }
    return nil
}

func (dataNodes *InterfaceProperties_DataNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dataNodes.DataNode {
        children[dataNodes.DataNode[i].GetSegmentPath()] = &dataNodes.DataNode[i]
    }
    return children
}

func (dataNodes *InterfaceProperties_DataNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataNodes *InterfaceProperties_DataNodes) GetBundleName() string { return "cisco_ios_xr" }

func (dataNodes *InterfaceProperties_DataNodes) GetYangName() string { return "data-nodes" }

func (dataNodes *InterfaceProperties_DataNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataNodes *InterfaceProperties_DataNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataNodes *InterfaceProperties_DataNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataNodes *InterfaceProperties_DataNodes) SetParent(parent types.Entity) { dataNodes.parent = parent }

func (dataNodes *InterfaceProperties_DataNodes) GetParent() types.Entity { return dataNodes.parent }

func (dataNodes *InterfaceProperties_DataNodes) GetParentYangName() string { return "interface-properties" }

// InterfaceProperties_DataNodes_DataNode
// The location of a (D)RP in the same LR as the
// interface being queried
type InterfaceProperties_DataNodes_DataNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The location of the (D)RP. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    DataNodeName interface{}

    // Location-specific view of interface operational data.
    Locationviews InterfaceProperties_DataNodes_DataNode_Locationviews

    // Partially qualified Location-specific view of interface operational data.
    PqNodeLocations InterfaceProperties_DataNodes_DataNode_PqNodeLocations

    // System-wide view of interface operational data.
    SystemView InterfaceProperties_DataNodes_DataNode_SystemView
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetFilter() yfilter.YFilter { return dataNode.YFilter }

func (dataNode *InterfaceProperties_DataNodes_DataNode) SetFilter(yf yfilter.YFilter) { dataNode.YFilter = yf }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetGoName(yname string) string {
    if yname == "data-node-name" { return "DataNodeName" }
    if yname == "locationviews" { return "Locationviews" }
    if yname == "pq-node-locations" { return "PqNodeLocations" }
    if yname == "system-view" { return "SystemView" }
    return ""
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetSegmentPath() string {
    return "data-node" + "[data-node-name='" + fmt.Sprintf("%v", dataNode.DataNodeName) + "']"
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "locationviews" {
        return &dataNode.Locationviews
    }
    if childYangName == "pq-node-locations" {
        return &dataNode.PqNodeLocations
    }
    if childYangName == "system-view" {
        return &dataNode.SystemView
    }
    return nil
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["locationviews"] = &dataNode.Locationviews
    children["pq-node-locations"] = &dataNode.PqNodeLocations
    children["system-view"] = &dataNode.SystemView
    return children
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-node-name"] = dataNode.DataNodeName
    return leafs
}

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetBundleName() string { return "cisco_ios_xr" }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetYangName() string { return "data-node" }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataNode *InterfaceProperties_DataNodes_DataNode) SetParent(parent types.Entity) { dataNode.parent = parent }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetParent() types.Entity { return dataNode.parent }

func (dataNode *InterfaceProperties_DataNodes_DataNode) GetParentYangName() string { return "data-nodes" }

// InterfaceProperties_DataNodes_DataNode_Locationviews
// Location-specific view of interface
// operational data
type InterfaceProperties_DataNodes_DataNode_Locationviews struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers on a particular node.
    // The type is slice of
    // InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview.
    Locationview []InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetFilter() yfilter.YFilter { return locationviews.YFilter }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) SetFilter(yf yfilter.YFilter) { locationviews.YFilter = yf }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetGoName(yname string) string {
    if yname == "locationview" { return "Locationview" }
    return ""
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetSegmentPath() string {
    return "locationviews"
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "locationview" {
        for _, c := range locationviews.Locationview {
            if locationviews.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview{}
        locationviews.Locationview = append(locationviews.Locationview, child)
        return &locationviews.Locationview[len(locationviews.Locationview)-1]
    }
    return nil
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locationviews.Locationview {
        children[locationviews.Locationview[i].GetSegmentPath()] = &locationviews.Locationview[i]
    }
    return children
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetBundleName() string { return "cisco_ios_xr" }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetYangName() string { return "locationviews" }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) SetParent(parent types.Entity) { locationviews.parent = parent }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetParent() types.Entity { return locationviews.parent }

func (locationviews *InterfaceProperties_DataNodes_DataNode_Locationviews) GetParentYangName() string { return "data-node" }

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview
// Operational data for all interfaces and
// controllers on a particular node
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The location to filter on. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    LocationviewName interface{}

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetFilter() yfilter.YFilter { return locationview.YFilter }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) SetFilter(yf yfilter.YFilter) { locationview.YFilter = yf }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetGoName(yname string) string {
    if yname == "locationview-name" { return "LocationviewName" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetSegmentPath() string {
    return "locationview" + "[locationview-name='" + fmt.Sprintf("%v", locationview.LocationviewName) + "']"
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &locationview.Interfaces
    }
    return nil
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &locationview.Interfaces
    return children
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["locationview-name"] = locationview.LocationviewName
    return leafs
}

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetBundleName() string { return "cisco_ios_xr" }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetYangName() string { return "locationview" }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) SetParent(parent types.Entity) { locationview.parent = parent }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetParent() types.Entity { return locationview.parent }

func (locationview *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview) GetParentYangName() string { return "locationviews" }

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface.
    Interface []InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces) GetParentYangName() string { return "locationview" }

// InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type interface{}

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

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "type" { return "Type" }
    if yname == "state" { return "State" }
    if yname == "actual-state" { return "ActualState" }
    if yname == "line-state" { return "LineState" }
    if yname == "actual-line-state" { return "ActualLineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "encapsulation-type-string" { return "EncapsulationTypeString" }
    if yname == "mtu" { return "Mtu" }
    if yname == "sub-interface-mtu-overhead" { return "SubInterfaceMtuOverhead" }
    if yname == "l2-transport" { return "L2Transport" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["parent-interface"] = self.ParentInterface
    leafs["type"] = self.Type
    leafs["state"] = self.State
    leafs["actual-state"] = self.ActualState
    leafs["line-state"] = self.LineState
    leafs["actual-line-state"] = self.ActualLineState
    leafs["encapsulation"] = self.Encapsulation
    leafs["encapsulation-type-string"] = self.EncapsulationTypeString
    leafs["mtu"] = self.Mtu
    leafs["sub-interface-mtu-overhead"] = self.SubInterfaceMtuOverhead
    leafs["l2-transport"] = self.L2Transport
    leafs["bandwidth"] = self.Bandwidth
    return leafs
}

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfaceProperties_DataNodes_DataNode_Locationviews_Locationview_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations
// Partially qualified Location-specific view of
// interface operational data
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers on a particular
    // pq_node. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation.
    PqNodeLocation []InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetFilter() yfilter.YFilter { return pqNodeLocations.YFilter }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) SetFilter(yf yfilter.YFilter) { pqNodeLocations.YFilter = yf }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetGoName(yname string) string {
    if yname == "pq-node-location" { return "PqNodeLocation" }
    return ""
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetSegmentPath() string {
    return "pq-node-locations"
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pq-node-location" {
        for _, c := range pqNodeLocations.PqNodeLocation {
            if pqNodeLocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation{}
        pqNodeLocations.PqNodeLocation = append(pqNodeLocations.PqNodeLocation, child)
        return &pqNodeLocations.PqNodeLocation[len(pqNodeLocations.PqNodeLocation)-1]
    }
    return nil
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pqNodeLocations.PqNodeLocation {
        children[pqNodeLocations.PqNodeLocation[i].GetSegmentPath()] = &pqNodeLocations.PqNodeLocation[i]
    }
    return children
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetBundleName() string { return "cisco_ios_xr" }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetYangName() string { return "pq-node-locations" }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) SetParent(parent types.Entity) { pqNodeLocations.parent = parent }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetParent() types.Entity { return pqNodeLocations.parent }

func (pqNodeLocations *InterfaceProperties_DataNodes_DataNode_PqNodeLocations) GetParentYangName() string { return "data-node" }

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation
// Operational data for all interfaces and
// controllers on a particular pq_node
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The partially qualified location to filter on. The
    // type is string with pattern:
    // ((([a-zA-Z0-9_]*\d+)|(\*))/){2}(([a-zA-Z0-9_]*\d+)|(\*)).
    PqNodeName interface{}

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetFilter() yfilter.YFilter { return pqNodeLocation.YFilter }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) SetFilter(yf yfilter.YFilter) { pqNodeLocation.YFilter = yf }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetGoName(yname string) string {
    if yname == "pq-node-name" { return "PqNodeName" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetSegmentPath() string {
    return "pq-node-location" + "[pq-node-name='" + fmt.Sprintf("%v", pqNodeLocation.PqNodeName) + "']"
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &pqNodeLocation.Interfaces
    }
    return nil
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &pqNodeLocation.Interfaces
    return children
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pq-node-name"] = pqNodeLocation.PqNodeName
    return leafs
}

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetBundleName() string { return "cisco_ios_xr" }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetYangName() string { return "pq-node-location" }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) SetParent(parent types.Entity) { pqNodeLocation.parent = parent }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetParent() types.Entity { return pqNodeLocation.parent }

func (pqNodeLocation *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation) GetParentYangName() string { return "pq-node-locations" }

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface.
    Interface []InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces) GetParentYangName() string { return "pq-node-location" }

// InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type interface{}

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

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "type" { return "Type" }
    if yname == "state" { return "State" }
    if yname == "actual-state" { return "ActualState" }
    if yname == "line-state" { return "LineState" }
    if yname == "actual-line-state" { return "ActualLineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "encapsulation-type-string" { return "EncapsulationTypeString" }
    if yname == "mtu" { return "Mtu" }
    if yname == "sub-interface-mtu-overhead" { return "SubInterfaceMtuOverhead" }
    if yname == "l2-transport" { return "L2Transport" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["parent-interface"] = self.ParentInterface
    leafs["type"] = self.Type
    leafs["state"] = self.State
    leafs["actual-state"] = self.ActualState
    leafs["line-state"] = self.LineState
    leafs["actual-line-state"] = self.ActualLineState
    leafs["encapsulation"] = self.Encapsulation
    leafs["encapsulation-type-string"] = self.EncapsulationTypeString
    leafs["mtu"] = self.Mtu
    leafs["sub-interface-mtu-overhead"] = self.SubInterfaceMtuOverhead
    leafs["l2-transport"] = self.L2Transport
    leafs["bandwidth"] = self.Bandwidth
    return leafs
}

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfaceProperties_DataNodes_DataNode_PqNodeLocations_PqNodeLocation_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// InterfaceProperties_DataNodes_DataNode_SystemView
// System-wide view of interface operational data
type InterfaceProperties_DataNodes_DataNode_SystemView struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for all interfaces and controllers.
    Interfaces InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetFilter() yfilter.YFilter { return systemView.YFilter }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) SetFilter(yf yfilter.YFilter) { systemView.YFilter = yf }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetSegmentPath() string {
    return "system-view"
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &systemView.Interfaces
    }
    return nil
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &systemView.Interfaces
    return children
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetBundleName() string { return "cisco_ios_xr" }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetYangName() string { return "system-view" }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) SetParent(parent types.Entity) { systemView.parent = parent }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetParent() types.Entity { return systemView.parent }

func (systemView *InterfaceProperties_DataNodes_DataNode_SystemView) GetParentYangName() string { return "data-node" }

// InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces
// Operational data for all interfaces and
// controllers
type InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The operational attributes for a particular interface. The type is slice of
    // InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface.
    Interface []InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces) GetParentYangName() string { return "system-view" }

// InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface
// The operational attributes for a particular
// interface
type InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Parent Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    ParentInterface interface{}

    // Interface type. The type is string.
    Type interface{}

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

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface" { return "Interface" }
    if yname == "parent-interface" { return "ParentInterface" }
    if yname == "type" { return "Type" }
    if yname == "state" { return "State" }
    if yname == "actual-state" { return "ActualState" }
    if yname == "line-state" { return "LineState" }
    if yname == "actual-line-state" { return "ActualLineState" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "encapsulation-type-string" { return "EncapsulationTypeString" }
    if yname == "mtu" { return "Mtu" }
    if yname == "sub-interface-mtu-overhead" { return "SubInterfaceMtuOverhead" }
    if yname == "l2-transport" { return "L2Transport" }
    if yname == "bandwidth" { return "Bandwidth" }
    return ""
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface"] = self.Interface
    leafs["parent-interface"] = self.ParentInterface
    leafs["type"] = self.Type
    leafs["state"] = self.State
    leafs["actual-state"] = self.ActualState
    leafs["line-state"] = self.LineState
    leafs["actual-line-state"] = self.ActualLineState
    leafs["encapsulation"] = self.Encapsulation
    leafs["encapsulation-type-string"] = self.EncapsulationTypeString
    leafs["mtu"] = self.Mtu
    leafs["sub-interface-mtu-overhead"] = self.SubInterfaceMtuOverhead
    leafs["l2-transport"] = self.L2Transport
    leafs["bandwidth"] = self.Bandwidth
    return leafs
}

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *InterfaceProperties_DataNodes_DataNode_SystemView_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

