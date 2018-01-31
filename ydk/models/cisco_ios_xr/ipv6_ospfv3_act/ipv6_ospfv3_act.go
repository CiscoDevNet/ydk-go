// This module contains a collection of YANG definitions
// for Cisco IOS-XR IPv6 OSPFv3 action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_ospfv3_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_ospfv3_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-routes}", reflect.TypeOf(ClearOspfv3Routes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-routes", reflect.TypeOf(ClearOspfv3Routes{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-redistribution}", reflect.TypeOf(ClearOspfv3Redistribution{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-redistribution", reflect.TypeOf(ClearOspfv3Redistribution{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-process}", reflect.TypeOf(ClearOspfv3Process{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-process", reflect.TypeOf(ClearOspfv3Process{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-statistics-neighbor}", reflect.TypeOf(ClearOspfv3StatisticsNeighbor{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics-neighbor", reflect.TypeOf(ClearOspfv3StatisticsNeighbor{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-statistics}", reflect.TypeOf(ClearOspfv3Statistics{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics", reflect.TypeOf(ClearOspfv3Statistics{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-act clear-ospfv3-instance-vrf}", reflect.TypeOf(ClearOspfv3InstanceVrf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-instance-vrf", reflect.TypeOf(ClearOspfv3InstanceVrf{}))
}

// ClearOspfv3Routes
// Clear OSPFv3 route table
type ClearOspfv3Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Routes_Input
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetFilter() yfilter.YFilter { return clearOspfv3Routes.YFilter }

func (clearOspfv3Routes *ClearOspfv3Routes) SetFilter(yf yfilter.YFilter) { clearOspfv3Routes.YFilter = yf }

func (clearOspfv3Routes *ClearOspfv3Routes) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-routes"
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3Routes.Input
    }
    return nil
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3Routes.Input
    return children
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3Routes *ClearOspfv3Routes) GetYangName() string { return "clear-ospfv3-routes" }

func (clearOspfv3Routes *ClearOspfv3Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3Routes *ClearOspfv3Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3Routes *ClearOspfv3Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3Routes *ClearOspfv3Routes) SetParent(parent types.Entity) { clearOspfv3Routes.parent = parent }

func (clearOspfv3Routes *ClearOspfv3Routes) GetParent() types.Entity { return clearOspfv3Routes.parent }

func (clearOspfv3Routes *ClearOspfv3Routes) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3Routes_Input
type ClearOspfv3Routes_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear OSPFv3 route table. The type is interface{}. This attribute is
    // mandatory.
    Route interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Routes_Input_Instance
}

func (input *ClearOspfv3Routes_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3Routes_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3Routes_Input) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfv3Routes_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3Routes_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfv3Routes_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfv3Routes_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = input.Route
    return leafs
}

func (input *ClearOspfv3Routes_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3Routes_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3Routes_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3Routes_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3Routes_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3Routes_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3Routes_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3Routes_Input) GetParentYangName() string { return "clear-ospfv3-routes" }

// ClearOspfv3Routes_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Routes_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Routes_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3Routes_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3Routes_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfv3Routes_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3Routes_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfv3Routes_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfv3Routes_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3Routes_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3Routes_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3Routes_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3Routes_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3Routes_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3Routes_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3Routes_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3Routes_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3Redistribution
// Clear OSPFv3 route redistribution
type ClearOspfv3Redistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Redistribution_Input
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetFilter() yfilter.YFilter { return clearOspfv3Redistribution.YFilter }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) SetFilter(yf yfilter.YFilter) { clearOspfv3Redistribution.YFilter = yf }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-redistribution"
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3Redistribution.Input
    }
    return nil
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3Redistribution.Input
    return children
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetYangName() string { return "clear-ospfv3-redistribution" }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) SetParent(parent types.Entity) { clearOspfv3Redistribution.parent = parent }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetParent() types.Entity { return clearOspfv3Redistribution.parent }

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3Redistribution_Input
type ClearOspfv3Redistribution_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear OSPFv3 route redistribution. The type is interface{}. This attribute
    // is mandatory.
    Redistribution interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Redistribution_Input_Instance
}

func (input *ClearOspfv3Redistribution_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3Redistribution_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3Redistribution_Input) GetGoName(yname string) string {
    if yname == "redistribution" { return "Redistribution" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfv3Redistribution_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3Redistribution_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfv3Redistribution_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfv3Redistribution_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["redistribution"] = input.Redistribution
    return leafs
}

func (input *ClearOspfv3Redistribution_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3Redistribution_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3Redistribution_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3Redistribution_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3Redistribution_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3Redistribution_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3Redistribution_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3Redistribution_Input) GetParentYangName() string { return "clear-ospfv3-redistribution" }

// ClearOspfv3Redistribution_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Redistribution_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3Redistribution_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3Redistribution_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3Redistribution_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3Process
// Clear (reset) OSPFv3 Process
type ClearOspfv3Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Process_Input
}

func (clearOspfv3Process *ClearOspfv3Process) GetFilter() yfilter.YFilter { return clearOspfv3Process.YFilter }

func (clearOspfv3Process *ClearOspfv3Process) SetFilter(yf yfilter.YFilter) { clearOspfv3Process.YFilter = yf }

func (clearOspfv3Process *ClearOspfv3Process) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3Process *ClearOspfv3Process) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-process"
}

func (clearOspfv3Process *ClearOspfv3Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3Process.Input
    }
    return nil
}

func (clearOspfv3Process *ClearOspfv3Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3Process.Input
    return children
}

func (clearOspfv3Process *ClearOspfv3Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3Process *ClearOspfv3Process) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3Process *ClearOspfv3Process) GetYangName() string { return "clear-ospfv3-process" }

func (clearOspfv3Process *ClearOspfv3Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3Process *ClearOspfv3Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3Process *ClearOspfv3Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3Process *ClearOspfv3Process) SetParent(parent types.Entity) { clearOspfv3Process.parent = parent }

func (clearOspfv3Process *ClearOspfv3Process) GetParent() types.Entity { return clearOspfv3Process.parent }

func (clearOspfv3Process *ClearOspfv3Process) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3Process_Input
type ClearOspfv3Process_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPFv3 process. The type is interface{}. This attribute is mandatory.
    Process interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Process_Input_Instance
}

func (input *ClearOspfv3Process_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3Process_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3Process_Input) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfv3Process_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3Process_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfv3Process_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfv3Process_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = input.Process
    return leafs
}

func (input *ClearOspfv3Process_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3Process_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3Process_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3Process_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3Process_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3Process_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3Process_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3Process_Input) GetParentYangName() string { return "clear-ospfv3-process" }

// ClearOspfv3Process_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Process_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Process_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3Process_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3Process_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfv3Process_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3Process_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfv3Process_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfv3Process_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3Process_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3Process_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3Process_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3Process_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3Process_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3Process_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3Process_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3Process_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3StatisticsNeighbor
// Clear OSPFv3 neighbor statistics per interface or neighbor id
type ClearOspfv3StatisticsNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3StatisticsNeighbor_Input
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetFilter() yfilter.YFilter { return clearOspfv3StatisticsNeighbor.YFilter }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) SetFilter(yf yfilter.YFilter) { clearOspfv3StatisticsNeighbor.YFilter = yf }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics-neighbor"
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3StatisticsNeighbor.Input
    }
    return nil
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3StatisticsNeighbor.Input
    return children
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetYangName() string { return "clear-ospfv3-statistics-neighbor" }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) SetParent(parent types.Entity) { clearOspfv3StatisticsNeighbor.parent = parent }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetParent() types.Entity { return clearOspfv3StatisticsNeighbor.parent }

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3StatisticsNeighbor_Input
type ClearOspfv3StatisticsNeighbor_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3StatisticsNeighbor_Input_Instance

    
    Neighbor ClearOspfv3StatisticsNeighbor_Input_Neighbor
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3StatisticsNeighbor_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    if childYangName == "neighbor" {
        return &input.Neighbor
    }
    return nil
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    children["neighbor"] = &input.Neighbor
    return children
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3StatisticsNeighbor_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3StatisticsNeighbor_Input) GetParentYangName() string { return "clear-ospfv3-statistics-neighbor" }

// ClearOspfv3StatisticsNeighbor_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3StatisticsNeighbor_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3StatisticsNeighbor_Input_Neighbor
type ClearOspfv3StatisticsNeighbor_Input_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    leafs["interface-name"] = neighbor.InterfaceName
    return leafs
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetParentYangName() string { return "input" }

// ClearOspfv3Statistics
// Clear OSPFv3 counters and statistics
type ClearOspfv3Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Statistics_Input
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetFilter() yfilter.YFilter { return clearOspfv3Statistics.YFilter }

func (clearOspfv3Statistics *ClearOspfv3Statistics) SetFilter(yf yfilter.YFilter) { clearOspfv3Statistics.YFilter = yf }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics"
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3Statistics.Input
    }
    return nil
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3Statistics.Input
    return children
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetYangName() string { return "clear-ospfv3-statistics" }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3Statistics *ClearOspfv3Statistics) SetParent(parent types.Entity) { clearOspfv3Statistics.parent = parent }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetParent() types.Entity { return clearOspfv3Statistics.parent }

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3Statistics_Input
type ClearOspfv3Statistics_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All OSPFv3 counters and statistics. The type is interface{}.
    PrefixPriority interface{}

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Neighbor statistics per neighbor id. The type is interface{}.
    Neighbor interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Statistics_Input_Instance
}

func (input *ClearOspfv3Statistics_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3Statistics_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3Statistics_Input) GetGoName(yname string) string {
    if yname == "prefix-priority" { return "PrefixPriority" }
    if yname == "spf" { return "Spf" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfv3Statistics_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3Statistics_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfv3Statistics_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfv3Statistics_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-priority"] = input.PrefixPriority
    leafs["spf"] = input.Spf
    leafs["neighbor"] = input.Neighbor
    return leafs
}

func (input *ClearOspfv3Statistics_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3Statistics_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3Statistics_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3Statistics_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3Statistics_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3Statistics_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3Statistics_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3Statistics_Input) GetParentYangName() string { return "clear-ospfv3-statistics" }

// ClearOspfv3Statistics_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Statistics_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3Statistics_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3Statistics_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3Statistics_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3Statistics_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3Statistics_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3Statistics_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3Statistics_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3Statistics_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3Statistics_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3InstanceVrf
// Clear one or more non-default OSPFv3 VRFs in process
type ClearOspfv3InstanceVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfv3InstanceVrf_Input
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetFilter() yfilter.YFilter { return clearOspfv3InstanceVrf.YFilter }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) SetFilter(yf yfilter.YFilter) { clearOspfv3InstanceVrf.YFilter = yf }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-instance-vrf"
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfv3InstanceVrf.Input
    }
    return nil
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfv3InstanceVrf.Input
    return children
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetYangName() string { return "clear-ospfv3-instance-vrf" }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) SetParent(parent types.Entity) { clearOspfv3InstanceVrf.parent = parent }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetParent() types.Entity { return clearOspfv3InstanceVrf.parent }

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-ospfv3-act" }

// ClearOspfv3InstanceVrf_Input
type ClearOspfv3InstanceVrf_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 instance name.
    Instance ClearOspfv3InstanceVrf_Input_Instance
}

func (input *ClearOspfv3InstanceVrf_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfv3InstanceVrf_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfv3InstanceVrf_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfv3InstanceVrf_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfv3InstanceVrf_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfv3InstanceVrf_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfv3InstanceVrf_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearOspfv3InstanceVrf_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfv3InstanceVrf_Input) GetYangName() string { return "input" }

func (input *ClearOspfv3InstanceVrf_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfv3InstanceVrf_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfv3InstanceVrf_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfv3InstanceVrf_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfv3InstanceVrf_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfv3InstanceVrf_Input) GetParentYangName() string { return "clear-ospfv3-instance-vrf" }

// ClearOspfv3InstanceVrf_Input_Instance
// OSPFv3 instance name
type ClearOspfv3InstanceVrf_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string. This attribute is
    // mandatory.
    InstanceIdentifier interface{}

    // Clear one or more non-default OSPFv3 VRFs in process.
    Vrf ClearOspfv3InstanceVrf_Input_Instance_Vrf

    // Clear all non-default OSPFv3 VRFs.
    All ClearOspfv3InstanceVrf_Input_Instance_All

    // Clear all non-default and default OSPFv3 VRFs.
    AllInclusive ClearOspfv3InstanceVrf_Input_Instance_AllInclusive
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    if yname == "vrf" { return "Vrf" }
    if yname == "all" { return "All" }
    if yname == "all-inclusive" { return "AllInclusive" }
    return ""
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        return &instance.Vrf
    }
    if childYangName == "all" {
        return &instance.All
    }
    if childYangName == "all-inclusive" {
        return &instance.AllInclusive
    }
    return nil
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf"] = &instance.Vrf
    children["all"] = &instance.All
    children["all-inclusive"] = &instance.AllInclusive
    return children
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfv3InstanceVrf_Input_Instance_Vrf
// Clear one or more non-default OSPFv3 VRFs in process
type ClearOspfv3InstanceVrf_Input_Instance_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 VRF name. The type is string. This attribute is mandatory.
    VrfName interface{}

    // Reset OSPFv3 process. The type is interface{}.
    Process interface{}

    // Clear OSPFv3 route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPFv3 route table. The type is interface{}.
    Route interface{}

    // OSPFv3 counters and statistics.
    Stats ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &vrf.Stats
    }
    return nil
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &vrf.Stats
    return children
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["process"] = vrf.Process
    leafs["redistribution"] = vrf.Redistribution
    leafs["route"] = vrf.Route
    return leafs
}

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetYangName() string { return "vrf" }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetParentYangName() string { return "instance" }

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "prefix-priority" { return "PrefixPriority" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["prefix-priority"] = stats.PrefixPriority
    return leafs
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetParentYangName() string { return "vrf" }

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

// ClearOspfv3InstanceVrf_Input_Instance_All
// Clear all non-default OSPFv3 VRFs
type ClearOspfv3InstanceVrf_Input_Instance_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPFv3 process. The type is interface{}.
    Process interface{}

    // Clear OSPFv3 route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPFv3 route table. The type is interface{}.
    Route interface{}

    // OSPFv3 counters and statistics.
    Stats ClearOspfv3InstanceVrf_Input_Instance_All_Stats
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetSegmentPath() string {
    return "all"
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &all.Stats
    }
    return nil
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &all.Stats
    return children
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = all.Process
    leafs["redistribution"] = all.Redistribution
    leafs["route"] = all.Route
    return leafs
}

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetYangName() string { return "all" }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetParent() types.Entity { return all.parent }

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetParentYangName() string { return "instance" }

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "prefix-priority" { return "PrefixPriority" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["prefix-priority"] = stats.PrefixPriority
    return leafs
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetParentYangName() string { return "all" }

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive
// Clear all non-default and default OSPFv3 VRFs
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPFv3 process. The type is interface{}.
    Process interface{}

    // Clear OSPFv3 route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPFv3 route table. The type is interface{}.
    Route interface{}

    // OSPFv3 counters and statistics.
    Stats ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetFilter() yfilter.YFilter { return allInclusive.YFilter }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) SetFilter(yf yfilter.YFilter) { allInclusive.YFilter = yf }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetSegmentPath() string {
    return "all-inclusive"
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &allInclusive.Stats
    }
    return nil
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &allInclusive.Stats
    return children
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = allInclusive.Process
    leafs["redistribution"] = allInclusive.Redistribution
    leafs["route"] = allInclusive.Route
    return leafs
}

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetBundleName() string { return "cisco_ios_xr" }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetYangName() string { return "all-inclusive" }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) SetParent(parent types.Entity) { allInclusive.parent = parent }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetParent() types.Entity { return allInclusive.parent }

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetParentYangName() string { return "instance" }

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "prefix-priority" { return "PrefixPriority" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["prefix-priority"] = stats.PrefixPriority
    return leafs
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetParentYangName() string { return "all-inclusive" }

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

