// This module contains a collection of YANG definitions
// for Cisco IOS-XR ISIS action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package isis_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package isis_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis-process}", reflect.TypeOf(ClearIsisProcess{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis-process", reflect.TypeOf(ClearIsisProcess{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis-route}", reflect.TypeOf(ClearIsisRoute{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis-route", reflect.TypeOf(ClearIsisRoute{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis-stat}", reflect.TypeOf(ClearIsisStat{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis-stat", reflect.TypeOf(ClearIsisStat{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis-dist}", reflect.TypeOf(ClearIsisDist{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis-dist", reflect.TypeOf(ClearIsisDist{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis-local-lsp}", reflect.TypeOf(ClearIsisLocalLsp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis-local-lsp", reflect.TypeOf(ClearIsisLocalLsp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-isis-act clear-isis}", reflect.TypeOf(ClearIsis{}))
    ydk.RegisterEntity("Cisco-IOS-XR-isis-act:clear-isis", reflect.TypeOf(ClearIsis{}))
}

// ClearIsisProcess
// Clear all IS-IS data structures
type ClearIsisProcess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsisProcess_Input
}

func (clearIsisProcess *ClearIsisProcess) GetFilter() yfilter.YFilter { return clearIsisProcess.YFilter }

func (clearIsisProcess *ClearIsisProcess) SetFilter(yf yfilter.YFilter) { clearIsisProcess.YFilter = yf }

func (clearIsisProcess *ClearIsisProcess) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsisProcess *ClearIsisProcess) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis-process"
}

func (clearIsisProcess *ClearIsisProcess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsisProcess.Input
    }
    return nil
}

func (clearIsisProcess *ClearIsisProcess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsisProcess.Input
    return children
}

func (clearIsisProcess *ClearIsisProcess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsisProcess *ClearIsisProcess) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsisProcess *ClearIsisProcess) GetYangName() string { return "clear-isis-process" }

func (clearIsisProcess *ClearIsisProcess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsisProcess *ClearIsisProcess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsisProcess *ClearIsisProcess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsisProcess *ClearIsisProcess) SetParent(parent types.Entity) { clearIsisProcess.parent = parent }

func (clearIsisProcess *ClearIsisProcess) GetParent() types.Entity { return clearIsisProcess.parent }

func (clearIsisProcess *ClearIsisProcess) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsisProcess_Input
type ClearIsisProcess_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear all IS-IS data structures. The type is interface{}.
    Process interface{}

    // Clear data from single IS-IS instance.
    Instance ClearIsisProcess_Input_Instance
}

func (input *ClearIsisProcess_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsisProcess_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsisProcess_Input) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearIsisProcess_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsisProcess_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearIsisProcess_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearIsisProcess_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = input.Process
    return leafs
}

func (input *ClearIsisProcess_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsisProcess_Input) GetYangName() string { return "input" }

func (input *ClearIsisProcess_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsisProcess_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsisProcess_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsisProcess_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsisProcess_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsisProcess_Input) GetParentYangName() string { return "clear-isis-process" }

// ClearIsisProcess_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisProcess_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisProcess_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsisProcess_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsisProcess_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsisProcess_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsisProcess_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsisProcess_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsisProcess_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsisProcess_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsisProcess_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsisProcess_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsisProcess_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsisProcess_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsisProcess_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsisProcess_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsisProcess_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsisRoute
// Clear IS-IS routes
type ClearIsisRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsisRoute_Input
}

func (clearIsisRoute *ClearIsisRoute) GetFilter() yfilter.YFilter { return clearIsisRoute.YFilter }

func (clearIsisRoute *ClearIsisRoute) SetFilter(yf yfilter.YFilter) { clearIsisRoute.YFilter = yf }

func (clearIsisRoute *ClearIsisRoute) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsisRoute *ClearIsisRoute) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis-route"
}

func (clearIsisRoute *ClearIsisRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsisRoute.Input
    }
    return nil
}

func (clearIsisRoute *ClearIsisRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsisRoute.Input
    return children
}

func (clearIsisRoute *ClearIsisRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsisRoute *ClearIsisRoute) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsisRoute *ClearIsisRoute) GetYangName() string { return "clear-isis-route" }

func (clearIsisRoute *ClearIsisRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsisRoute *ClearIsisRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsisRoute *ClearIsisRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsisRoute *ClearIsisRoute) SetParent(parent types.Entity) { clearIsisRoute.parent = parent }

func (clearIsisRoute *ClearIsisRoute) GetParent() types.Entity { return clearIsisRoute.parent }

func (clearIsisRoute *ClearIsisRoute) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsisRoute_Input
type ClearIsisRoute_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear IS-IS routes. The type is interface{}.
    Route interface{}

    // Clear data from single IS-IS instance.
    Instance ClearIsisRoute_Input_Instance
}

func (input *ClearIsisRoute_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsisRoute_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsisRoute_Input) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearIsisRoute_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsisRoute_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearIsisRoute_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearIsisRoute_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = input.Route
    return leafs
}

func (input *ClearIsisRoute_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsisRoute_Input) GetYangName() string { return "input" }

func (input *ClearIsisRoute_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsisRoute_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsisRoute_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsisRoute_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsisRoute_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsisRoute_Input) GetParentYangName() string { return "clear-isis-route" }

// ClearIsisRoute_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisRoute_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisRoute_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsisRoute_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsisRoute_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsisRoute_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsisRoute_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsisRoute_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsisRoute_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsisRoute_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsisRoute_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsisRoute_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsisRoute_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsisRoute_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsisRoute_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsisRoute_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsisRoute_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsisStat
// Clear IS-IS protocol statistics
type ClearIsisStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsisStat_Input
}

func (clearIsisStat *ClearIsisStat) GetFilter() yfilter.YFilter { return clearIsisStat.YFilter }

func (clearIsisStat *ClearIsisStat) SetFilter(yf yfilter.YFilter) { clearIsisStat.YFilter = yf }

func (clearIsisStat *ClearIsisStat) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsisStat *ClearIsisStat) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis-stat"
}

func (clearIsisStat *ClearIsisStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsisStat.Input
    }
    return nil
}

func (clearIsisStat *ClearIsisStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsisStat.Input
    return children
}

func (clearIsisStat *ClearIsisStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsisStat *ClearIsisStat) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsisStat *ClearIsisStat) GetYangName() string { return "clear-isis-stat" }

func (clearIsisStat *ClearIsisStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsisStat *ClearIsisStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsisStat *ClearIsisStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsisStat *ClearIsisStat) SetParent(parent types.Entity) { clearIsisStat.parent = parent }

func (clearIsisStat *ClearIsisStat) GetParent() types.Entity { return clearIsisStat.parent }

func (clearIsisStat *ClearIsisStat) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsisStat_Input
type ClearIsisStat_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear data from single IS-IS instance.
    Instance ClearIsisStat_Input_Instance

    // Clear IS-IS protocol statistics.
    Statistics ClearIsisStat_Input_Statistics
}

func (input *ClearIsisStat_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsisStat_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsisStat_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (input *ClearIsisStat_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsisStat_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    if childYangName == "statistics" {
        return &input.Statistics
    }
    return nil
}

func (input *ClearIsisStat_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    children["statistics"] = &input.Statistics
    return children
}

func (input *ClearIsisStat_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearIsisStat_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsisStat_Input) GetYangName() string { return "input" }

func (input *ClearIsisStat_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsisStat_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsisStat_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsisStat_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsisStat_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsisStat_Input) GetParentYangName() string { return "clear-isis-stat" }

// ClearIsisStat_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisStat_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisStat_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsisStat_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsisStat_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsisStat_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsisStat_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsisStat_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsisStat_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsisStat_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsisStat_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsisStat_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsisStat_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsisStat_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsisStat_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsisStat_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsisStat_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsisStat_Input_Statistics
// Clear IS-IS protocol statistics
type ClearIsisStat_Input_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+. This
    // attribute is mandatory.
    InterfaceName interface{}
}

func (statistics *ClearIsisStat_Input_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *ClearIsisStat_Input_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *ClearIsisStat_Input_Statistics) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (statistics *ClearIsisStat_Input_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *ClearIsisStat_Input_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *ClearIsisStat_Input_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *ClearIsisStat_Input_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = statistics.InterfaceName
    return leafs
}

func (statistics *ClearIsisStat_Input_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *ClearIsisStat_Input_Statistics) GetYangName() string { return "statistics" }

func (statistics *ClearIsisStat_Input_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *ClearIsisStat_Input_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *ClearIsisStat_Input_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *ClearIsisStat_Input_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *ClearIsisStat_Input_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *ClearIsisStat_Input_Statistics) GetParentYangName() string { return "input" }

// ClearIsisDist
// Reset BGP-LS topology distribution
type ClearIsisDist struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsisDist_Input
}

func (clearIsisDist *ClearIsisDist) GetFilter() yfilter.YFilter { return clearIsisDist.YFilter }

func (clearIsisDist *ClearIsisDist) SetFilter(yf yfilter.YFilter) { clearIsisDist.YFilter = yf }

func (clearIsisDist *ClearIsisDist) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsisDist *ClearIsisDist) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis-dist"
}

func (clearIsisDist *ClearIsisDist) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsisDist.Input
    }
    return nil
}

func (clearIsisDist *ClearIsisDist) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsisDist.Input
    return children
}

func (clearIsisDist *ClearIsisDist) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsisDist *ClearIsisDist) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsisDist *ClearIsisDist) GetYangName() string { return "clear-isis-dist" }

func (clearIsisDist *ClearIsisDist) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsisDist *ClearIsisDist) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsisDist *ClearIsisDist) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsisDist *ClearIsisDist) SetParent(parent types.Entity) { clearIsisDist.parent = parent }

func (clearIsisDist *ClearIsisDist) GetParent() types.Entity { return clearIsisDist.parent }

func (clearIsisDist *ClearIsisDist) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsisDist_Input
type ClearIsisDist_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset BGP-LS topology distribution. The type is interface{}.
    Distribution interface{}

    // Reset BGP-LS topology from single IS-IS instance.
    Instance ClearIsisDist_Input_Instance
}

func (input *ClearIsisDist_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsisDist_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsisDist_Input) GetGoName(yname string) string {
    if yname == "distribution" { return "Distribution" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearIsisDist_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsisDist_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearIsisDist_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearIsisDist_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["distribution"] = input.Distribution
    return leafs
}

func (input *ClearIsisDist_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsisDist_Input) GetYangName() string { return "input" }

func (input *ClearIsisDist_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsisDist_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsisDist_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsisDist_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsisDist_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsisDist_Input) GetParentYangName() string { return "clear-isis-dist" }

// ClearIsisDist_Input_Instance
// Reset BGP-LS topology from single IS-IS instance
type ClearIsisDist_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisDist_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsisDist_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsisDist_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsisDist_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsisDist_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsisDist_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsisDist_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsisDist_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsisDist_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsisDist_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsisDist_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsisDist_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsisDist_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsisDist_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsisDist_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsisLocalLsp
// Clean and regenerate local LSPs
type ClearIsisLocalLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsisLocalLsp_Input
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetFilter() yfilter.YFilter { return clearIsisLocalLsp.YFilter }

func (clearIsisLocalLsp *ClearIsisLocalLsp) SetFilter(yf yfilter.YFilter) { clearIsisLocalLsp.YFilter = yf }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis-local-lsp"
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsisLocalLsp.Input
    }
    return nil
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsisLocalLsp.Input
    return children
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetYangName() string { return "clear-isis-local-lsp" }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsisLocalLsp *ClearIsisLocalLsp) SetParent(parent types.Entity) { clearIsisLocalLsp.parent = parent }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetParent() types.Entity { return clearIsisLocalLsp.parent }

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsisLocalLsp_Input
type ClearIsisLocalLsp_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clean and regenerate local LSPs. The type is interface{}.
    LocalLsp interface{}

    // Clean and regenerate local LSPs from single IS-IS instance.
    Instance ClearIsisLocalLsp_Input_Instance
}

func (input *ClearIsisLocalLsp_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsisLocalLsp_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsisLocalLsp_Input) GetGoName(yname string) string {
    if yname == "local-lsp" { return "LocalLsp" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearIsisLocalLsp_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsisLocalLsp_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearIsisLocalLsp_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearIsisLocalLsp_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-lsp"] = input.LocalLsp
    return leafs
}

func (input *ClearIsisLocalLsp_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsisLocalLsp_Input) GetYangName() string { return "input" }

func (input *ClearIsisLocalLsp_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsisLocalLsp_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsisLocalLsp_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsisLocalLsp_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsisLocalLsp_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsisLocalLsp_Input) GetParentYangName() string { return "clear-isis-local-lsp" }

// ClearIsisLocalLsp_Input_Instance
// Clean and regenerate local LSPs from single IS-IS instance
type ClearIsisLocalLsp_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsisLocalLsp_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsisLocalLsp_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsisLocalLsp_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsisLocalLsp_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsisLocalLsp_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsisLocalLsp_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsisLocalLsp_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsisLocalLsp_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsisLocalLsp_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsis
// Clear IS-IS data structures
type ClearIsis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearIsis_Input
}

func (clearIsis *ClearIsis) GetFilter() yfilter.YFilter { return clearIsis.YFilter }

func (clearIsis *ClearIsis) SetFilter(yf yfilter.YFilter) { clearIsis.YFilter = yf }

func (clearIsis *ClearIsis) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearIsis *ClearIsis) GetSegmentPath() string {
    return "Cisco-IOS-XR-isis-act:clear-isis"
}

func (clearIsis *ClearIsis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearIsis.Input
    }
    return nil
}

func (clearIsis *ClearIsis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearIsis.Input
    return children
}

func (clearIsis *ClearIsis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearIsis *ClearIsis) GetBundleName() string { return "cisco_ios_xr" }

func (clearIsis *ClearIsis) GetYangName() string { return "clear-isis" }

func (clearIsis *ClearIsis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearIsis *ClearIsis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearIsis *ClearIsis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearIsis *ClearIsis) SetParent(parent types.Entity) { clearIsis.parent = parent }

func (clearIsis *ClearIsis) GetParent() types.Entity { return clearIsis.parent }

func (clearIsis *ClearIsis) GetParentYangName() string { return "Cisco-IOS-XR-isis-act" }

// ClearIsis_Input
type ClearIsis_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear data for these route types. The type is RtType.
    RtType interface{}

    // Clear IS-IS routes. The type is interface{}.
    Route interface{}

    // Topology table information. The type is string.
    Topology interface{}

    // Clear data from single IS-IS instance.
    Instance ClearIsis_Input_Instance
}

func (input *ClearIsis_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearIsis_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearIsis_Input) GetGoName(yname string) string {
    if yname == "rt-type" { return "RtType" }
    if yname == "route" { return "Route" }
    if yname == "topology" { return "Topology" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearIsis_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearIsis_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearIsis_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearIsis_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rt-type"] = input.RtType
    leafs["route"] = input.Route
    leafs["topology"] = input.Topology
    return leafs
}

func (input *ClearIsis_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearIsis_Input) GetYangName() string { return "input" }

func (input *ClearIsis_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearIsis_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearIsis_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearIsis_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearIsis_Input) GetParent() types.Entity { return input.parent }

func (input *ClearIsis_Input) GetParentYangName() string { return "clear-isis" }

// ClearIsis_Input_Instance
// Clear data from single IS-IS instance
type ClearIsis_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsis_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearIsis_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearIsis_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearIsis_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearIsis_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearIsis_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearIsis_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearIsis_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearIsis_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearIsis_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearIsis_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearIsis_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearIsis_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearIsis_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearIsis_Input_Instance) GetParentYangName() string { return "input" }

// ClearIsis_Input_RtType represents Clear data for these route types
type ClearIsis_Input_RtType string

const (
    ClearIsis_Input_RtType_AFI_ALL_MULTICAST ClearIsis_Input_RtType = "AFI-ALL-MULTICAST"

    ClearIsis_Input_RtType_AFI_ALL_SAFI_ALL ClearIsis_Input_RtType = "AFI-ALL-SAFI-ALL"

    ClearIsis_Input_RtType_AFI_ALL_UNICAST ClearIsis_Input_RtType = "AFI-ALL-UNICAST"

    ClearIsis_Input_RtType_IPv4_MULTICAST ClearIsis_Input_RtType = "IPv4-MULTICAST"

    ClearIsis_Input_RtType_IPv4_SAFI_ALL ClearIsis_Input_RtType = "IPv4-SAFI-ALL"

    ClearIsis_Input_RtType_IPv4_UNICAST ClearIsis_Input_RtType = "IPv4-UNICAST"

    ClearIsis_Input_RtType_IPv6_MULTICAST ClearIsis_Input_RtType = "IPv6-MULTICAST"

    ClearIsis_Input_RtType_IPv6_SAFI_ALL ClearIsis_Input_RtType = "IPv6-SAFI-ALL"

    ClearIsis_Input_RtType_IPv6_UNICAST ClearIsis_Input_RtType = "IPv6-UNICAST"
)

