// This module contains a collection of YANG definitions
// for Cisco IOS-XR OSPF action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_ospf_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_ospf_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-routes}", reflect.TypeOf(ClearOspfRoutes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-routes", reflect.TypeOf(ClearOspfRoutes{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-redistribution}", reflect.TypeOf(ClearOspfRedistribution{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-redistribution", reflect.TypeOf(ClearOspfRedistribution{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-statistics}", reflect.TypeOf(ClearOspfStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics", reflect.TypeOf(ClearOspfStatistics{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-statistics-neighbor}", reflect.TypeOf(ClearOspfStatisticsNeighbor{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-neighbor", reflect.TypeOf(ClearOspfStatisticsNeighbor{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-statistics-interface}", reflect.TypeOf(ClearOspfStatisticsInterface{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-interface", reflect.TypeOf(ClearOspfStatisticsInterface{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-process}", reflect.TypeOf(ClearOspfProcess{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-process", reflect.TypeOf(ClearOspfProcess{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ospf-act clear-ospf-instance-vrf}", reflect.TypeOf(ClearOspfInstanceVrf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-instance-vrf", reflect.TypeOf(ClearOspfInstanceVrf{}))
}

// ClearOspfRoutes
// Clear OSPF route table
type ClearOspfRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfRoutes_Input
}

func (clearOspfRoutes *ClearOspfRoutes) GetFilter() yfilter.YFilter { return clearOspfRoutes.YFilter }

func (clearOspfRoutes *ClearOspfRoutes) SetFilter(yf yfilter.YFilter) { clearOspfRoutes.YFilter = yf }

func (clearOspfRoutes *ClearOspfRoutes) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfRoutes *ClearOspfRoutes) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-routes"
}

func (clearOspfRoutes *ClearOspfRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfRoutes.Input
    }
    return nil
}

func (clearOspfRoutes *ClearOspfRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfRoutes.Input
    return children
}

func (clearOspfRoutes *ClearOspfRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfRoutes *ClearOspfRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfRoutes *ClearOspfRoutes) GetYangName() string { return "clear-ospf-routes" }

func (clearOspfRoutes *ClearOspfRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfRoutes *ClearOspfRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfRoutes *ClearOspfRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfRoutes *ClearOspfRoutes) SetParent(parent types.Entity) { clearOspfRoutes.parent = parent }

func (clearOspfRoutes *ClearOspfRoutes) GetParent() types.Entity { return clearOspfRoutes.parent }

func (clearOspfRoutes *ClearOspfRoutes) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfRoutes_Input
type ClearOspfRoutes_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear OSPF route table. The type is interface{}. This attribute is
    // mandatory.
    Route interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfRoutes_Input_Instance
}

func (input *ClearOspfRoutes_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfRoutes_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfRoutes_Input) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfRoutes_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfRoutes_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfRoutes_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfRoutes_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = input.Route
    return leafs
}

func (input *ClearOspfRoutes_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfRoutes_Input) GetYangName() string { return "input" }

func (input *ClearOspfRoutes_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfRoutes_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfRoutes_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfRoutes_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfRoutes_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfRoutes_Input) GetParentYangName() string { return "clear-ospf-routes" }

// ClearOspfRoutes_Input_Instance
// Clear data from OSPF instance
type ClearOspfRoutes_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfRoutes_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfRoutes_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfRoutes_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfRoutes_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfRoutes_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfRoutes_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfRoutes_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfRoutes_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfRoutes_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfRoutes_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfRoutes_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfRoutes_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfRoutes_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfRoutes_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfRoutes_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfRedistribution
// Clear OSPF route redistribution
type ClearOspfRedistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfRedistribution_Input
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetFilter() yfilter.YFilter { return clearOspfRedistribution.YFilter }

func (clearOspfRedistribution *ClearOspfRedistribution) SetFilter(yf yfilter.YFilter) { clearOspfRedistribution.YFilter = yf }

func (clearOspfRedistribution *ClearOspfRedistribution) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-redistribution"
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfRedistribution.Input
    }
    return nil
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfRedistribution.Input
    return children
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfRedistribution *ClearOspfRedistribution) GetYangName() string { return "clear-ospf-redistribution" }

func (clearOspfRedistribution *ClearOspfRedistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfRedistribution *ClearOspfRedistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfRedistribution *ClearOspfRedistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfRedistribution *ClearOspfRedistribution) SetParent(parent types.Entity) { clearOspfRedistribution.parent = parent }

func (clearOspfRedistribution *ClearOspfRedistribution) GetParent() types.Entity { return clearOspfRedistribution.parent }

func (clearOspfRedistribution *ClearOspfRedistribution) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfRedistribution_Input
type ClearOspfRedistribution_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear OSPF route redistribution. The type is interface{}. This attribute is
    // mandatory.
    Redistribution interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfRedistribution_Input_Instance
}

func (input *ClearOspfRedistribution_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfRedistribution_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfRedistribution_Input) GetGoName(yname string) string {
    if yname == "redistribution" { return "Redistribution" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfRedistribution_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfRedistribution_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfRedistribution_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfRedistribution_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["redistribution"] = input.Redistribution
    return leafs
}

func (input *ClearOspfRedistribution_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfRedistribution_Input) GetYangName() string { return "input" }

func (input *ClearOspfRedistribution_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfRedistribution_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfRedistribution_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfRedistribution_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfRedistribution_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfRedistribution_Input) GetParentYangName() string { return "clear-ospf-redistribution" }

// ClearOspfRedistribution_Input_Instance
// Clear data from OSPF instance
type ClearOspfRedistribution_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfRedistribution_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfRedistribution_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfRedistribution_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfRedistribution_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfRedistribution_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfRedistribution_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfRedistribution_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfRedistribution_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfRedistribution_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfRedistribution_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfRedistribution_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfRedistribution_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfRedistribution_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfRedistribution_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfRedistribution_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfStatistics
// Clear OSPF counters and statistics
type ClearOspfStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfStatistics_Input
}

func (clearOspfStatistics *ClearOspfStatistics) GetFilter() yfilter.YFilter { return clearOspfStatistics.YFilter }

func (clearOspfStatistics *ClearOspfStatistics) SetFilter(yf yfilter.YFilter) { clearOspfStatistics.YFilter = yf }

func (clearOspfStatistics *ClearOspfStatistics) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfStatistics *ClearOspfStatistics) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics"
}

func (clearOspfStatistics *ClearOspfStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfStatistics.Input
    }
    return nil
}

func (clearOspfStatistics *ClearOspfStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfStatistics.Input
    return children
}

func (clearOspfStatistics *ClearOspfStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfStatistics *ClearOspfStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfStatistics *ClearOspfStatistics) GetYangName() string { return "clear-ospf-statistics" }

func (clearOspfStatistics *ClearOspfStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfStatistics *ClearOspfStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfStatistics *ClearOspfStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfStatistics *ClearOspfStatistics) SetParent(parent types.Entity) { clearOspfStatistics.parent = parent }

func (clearOspfStatistics *ClearOspfStatistics) GetParent() types.Entity { return clearOspfStatistics.parent }

func (clearOspfStatistics *ClearOspfStatistics) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfStatistics_Input
type ClearOspfStatistics_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All OSPF counters and statistics. The type is interface{}.
    All interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Neighbor statistics per neighbor id. The type is interface{}.
    Neighbor interface{}

    // OSPF interface statistics. The type is interface{}.
    InterfaceName interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfStatistics_Input_Instance
}

func (input *ClearOspfStatistics_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfStatistics_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfStatistics_Input) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "message-queue" { return "MessageQueue" }
    if yname == "spf" { return "Spf" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfStatistics_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfStatistics_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfStatistics_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfStatistics_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = input.All
    leafs["message-queue"] = input.MessageQueue
    leafs["spf"] = input.Spf
    leafs["neighbor"] = input.Neighbor
    leafs["interface-name"] = input.InterfaceName
    return leafs
}

func (input *ClearOspfStatistics_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfStatistics_Input) GetYangName() string { return "input" }

func (input *ClearOspfStatistics_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfStatistics_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfStatistics_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfStatistics_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfStatistics_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfStatistics_Input) GetParentYangName() string { return "clear-ospf-statistics" }

// ClearOspfStatistics_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatistics_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatistics_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfStatistics_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfStatistics_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfStatistics_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfStatistics_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfStatistics_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfStatistics_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfStatistics_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfStatistics_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfStatistics_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfStatistics_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfStatistics_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfStatistics_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfStatistics_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfStatistics_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfStatisticsNeighbor
// Clear OSPF neighbor statistics per interface or neighbor id
type ClearOspfStatisticsNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfStatisticsNeighbor_Input
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetFilter() yfilter.YFilter { return clearOspfStatisticsNeighbor.YFilter }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) SetFilter(yf yfilter.YFilter) { clearOspfStatisticsNeighbor.YFilter = yf }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-neighbor"
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfStatisticsNeighbor.Input
    }
    return nil
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfStatisticsNeighbor.Input
    return children
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetYangName() string { return "clear-ospf-statistics-neighbor" }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) SetParent(parent types.Entity) { clearOspfStatisticsNeighbor.parent = parent }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetParent() types.Entity { return clearOspfStatisticsNeighbor.parent }

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfStatisticsNeighbor_Input
type ClearOspfStatisticsNeighbor_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear data from OSPF instance.
    Instance ClearOspfStatisticsNeighbor_Input_Instance

    
    Neighbor ClearOspfStatisticsNeighbor_Input_Neighbor
}

func (input *ClearOspfStatisticsNeighbor_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfStatisticsNeighbor_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfStatisticsNeighbor_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (input *ClearOspfStatisticsNeighbor_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfStatisticsNeighbor_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    if childYangName == "neighbor" {
        return &input.Neighbor
    }
    return nil
}

func (input *ClearOspfStatisticsNeighbor_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    children["neighbor"] = &input.Neighbor
    return children
}

func (input *ClearOspfStatisticsNeighbor_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearOspfStatisticsNeighbor_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfStatisticsNeighbor_Input) GetYangName() string { return "input" }

func (input *ClearOspfStatisticsNeighbor_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfStatisticsNeighbor_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfStatisticsNeighbor_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfStatisticsNeighbor_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfStatisticsNeighbor_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfStatisticsNeighbor_Input) GetParentYangName() string { return "clear-ospf-statistics-neighbor" }

// ClearOspfStatisticsNeighbor_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatisticsNeighbor_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfStatisticsNeighbor_Input_Neighbor
type ClearOspfStatisticsNeighbor_Input_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    leafs["interface-name"] = neighbor.InterfaceName
    return leafs
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetParentYangName() string { return "input" }

// ClearOspfStatisticsInterface
// Clear OSPF interface statistics
type ClearOspfStatisticsInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfStatisticsInterface_Input
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetFilter() yfilter.YFilter { return clearOspfStatisticsInterface.YFilter }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) SetFilter(yf yfilter.YFilter) { clearOspfStatisticsInterface.YFilter = yf }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-interface"
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfStatisticsInterface.Input
    }
    return nil
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfStatisticsInterface.Input
    return children
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetYangName() string { return "clear-ospf-statistics-interface" }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) SetParent(parent types.Entity) { clearOspfStatisticsInterface.parent = parent }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetParent() types.Entity { return clearOspfStatisticsInterface.parent }

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfStatisticsInterface_Input
type ClearOspfStatisticsInterface_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear data from OSPF instance.
    Instance ClearOspfStatisticsInterface_Input_Instance

    
    Interface ClearOspfStatisticsInterface_Input_Interface
}

func (input *ClearOspfStatisticsInterface_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfStatisticsInterface_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfStatisticsInterface_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (input *ClearOspfStatisticsInterface_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfStatisticsInterface_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    if childYangName == "interface" {
        return &input.Interface
    }
    return nil
}

func (input *ClearOspfStatisticsInterface_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    children["interface"] = &input.Interface
    return children
}

func (input *ClearOspfStatisticsInterface_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearOspfStatisticsInterface_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfStatisticsInterface_Input) GetYangName() string { return "input" }

func (input *ClearOspfStatisticsInterface_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfStatisticsInterface_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfStatisticsInterface_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfStatisticsInterface_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfStatisticsInterface_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfStatisticsInterface_Input) GetParentYangName() string { return "clear-ospf-statistics-interface" }

// ClearOspfStatisticsInterface_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatisticsInterface_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfStatisticsInterface_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfStatisticsInterface_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfStatisticsInterface_Input_Interface
type ClearOspfStatisticsInterface_Input_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfStatisticsInterface_Input_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfStatisticsInterface_Input_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfStatisticsInterface_Input_Interface) GetParentYangName() string { return "input" }

// ClearOspfProcess
// Clear (reset) OSPF process
type ClearOspfProcess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfProcess_Input
}

func (clearOspfProcess *ClearOspfProcess) GetFilter() yfilter.YFilter { return clearOspfProcess.YFilter }

func (clearOspfProcess *ClearOspfProcess) SetFilter(yf yfilter.YFilter) { clearOspfProcess.YFilter = yf }

func (clearOspfProcess *ClearOspfProcess) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfProcess *ClearOspfProcess) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-process"
}

func (clearOspfProcess *ClearOspfProcess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfProcess.Input
    }
    return nil
}

func (clearOspfProcess *ClearOspfProcess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfProcess.Input
    return children
}

func (clearOspfProcess *ClearOspfProcess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfProcess *ClearOspfProcess) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfProcess *ClearOspfProcess) GetYangName() string { return "clear-ospf-process" }

func (clearOspfProcess *ClearOspfProcess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfProcess *ClearOspfProcess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfProcess *ClearOspfProcess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfProcess *ClearOspfProcess) SetParent(parent types.Entity) { clearOspfProcess.parent = parent }

func (clearOspfProcess *ClearOspfProcess) GetParent() types.Entity { return clearOspfProcess.parent }

func (clearOspfProcess *ClearOspfProcess) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfProcess_Input
type ClearOspfProcess_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPF process. The type is interface{}. This attribute is mandatory.
    Process interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfProcess_Input_Instance
}

func (input *ClearOspfProcess_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfProcess_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfProcess_Input) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfProcess_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfProcess_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfProcess_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfProcess_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = input.Process
    return leafs
}

func (input *ClearOspfProcess_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfProcess_Input) GetYangName() string { return "input" }

func (input *ClearOspfProcess_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfProcess_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfProcess_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfProcess_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfProcess_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfProcess_Input) GetParentYangName() string { return "clear-ospf-process" }

// ClearOspfProcess_Input_Instance
// Clear data from OSPF instance
type ClearOspfProcess_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfProcess_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfProcess_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfProcess_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    return ""
}

func (instance *ClearOspfProcess_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfProcess_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *ClearOspfProcess_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *ClearOspfProcess_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfProcess_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfProcess_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfProcess_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfProcess_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfProcess_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfProcess_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfProcess_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfProcess_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfInstanceVrf
// Clear one or more non-default OSPF VRFs in process
type ClearOspfInstanceVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearOspfInstanceVrf_Input
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetFilter() yfilter.YFilter { return clearOspfInstanceVrf.YFilter }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) SetFilter(yf yfilter.YFilter) { clearOspfInstanceVrf.YFilter = yf }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-instance-vrf"
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearOspfInstanceVrf.Input
    }
    return nil
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearOspfInstanceVrf.Input
    return children
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetBundleName() string { return "cisco_ios_xr" }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetYangName() string { return "clear-ospf-instance-vrf" }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) SetParent(parent types.Entity) { clearOspfInstanceVrf.parent = parent }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetParent() types.Entity { return clearOspfInstanceVrf.parent }

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ospf-act" }

// ClearOspfInstanceVrf_Input
type ClearOspfInstanceVrf_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF instance name.
    Instance ClearOspfInstanceVrf_Input_Instance
}

func (input *ClearOspfInstanceVrf_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearOspfInstanceVrf_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearOspfInstanceVrf_Input) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (input *ClearOspfInstanceVrf_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearOspfInstanceVrf_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &input.Instance
    }
    return nil
}

func (input *ClearOspfInstanceVrf_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &input.Instance
    return children
}

func (input *ClearOspfInstanceVrf_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *ClearOspfInstanceVrf_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *ClearOspfInstanceVrf_Input) GetYangName() string { return "input" }

func (input *ClearOspfInstanceVrf_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *ClearOspfInstanceVrf_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *ClearOspfInstanceVrf_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *ClearOspfInstanceVrf_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearOspfInstanceVrf_Input) GetParent() types.Entity { return input.parent }

func (input *ClearOspfInstanceVrf_Input) GetParentYangName() string { return "clear-ospf-instance-vrf" }

// ClearOspfInstanceVrf_Input_Instance
// OSPF instance name
type ClearOspfInstanceVrf_Input_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string. This attribute is
    // mandatory.
    InstanceIdentifier interface{}

    // Clear one or more non-default OSPF VRFs in process.
    Vrf ClearOspfInstanceVrf_Input_Instance_Vrf

    // Clear all non-default OSPF VRFs.
    All ClearOspfInstanceVrf_Input_Instance_All

    // Clear all non-default and default OSPF VRFs.
    AllInclusive ClearOspfInstanceVrf_Input_Instance_AllInclusive
}

func (instance *ClearOspfInstanceVrf_Input_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *ClearOspfInstanceVrf_Input_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetGoName(yname string) string {
    if yname == "instance-identifier" { return "InstanceIdentifier" }
    if yname == "vrf" { return "Vrf" }
    if yname == "all" { return "All" }
    if yname == "all-inclusive" { return "AllInclusive" }
    return ""
}

func (instance *ClearOspfInstanceVrf_Input_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *ClearOspfInstanceVrf_Input_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
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

func (instance *ClearOspfInstanceVrf_Input_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf"] = &instance.Vrf
    children["all"] = &instance.All
    children["all-inclusive"] = &instance.AllInclusive
    return children
}

func (instance *ClearOspfInstanceVrf_Input_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-identifier"] = instance.InstanceIdentifier
    return leafs
}

func (instance *ClearOspfInstanceVrf_Input_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetYangName() string { return "instance" }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *ClearOspfInstanceVrf_Input_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetParent() types.Entity { return instance.parent }

func (instance *ClearOspfInstanceVrf_Input_Instance) GetParentYangName() string { return "input" }

// ClearOspfInstanceVrf_Input_Instance_Vrf
// Clear one or more non-default OSPF VRFs in process
type ClearOspfInstanceVrf_Input_Instance_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF VRF name. The type is string.
    VrfName interface{}

    // Reset OSPF process. The type is interface{}.
    Process interface{}

    // Clear OSPF route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPF route table. The type is interface{}.
    Route interface{}

    // OSPF counters and statistics.
    Stats ClearOspfInstanceVrf_Input_Instance_Vrf_Stats
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetSegmentPath() string {
    return "vrf"
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &vrf.Stats
    }
    return nil
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &vrf.Stats
    return children
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["process"] = vrf.Process
    leafs["redistribution"] = vrf.Redistribution
    leafs["route"] = vrf.Route
    return leafs
}

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetYangName() string { return "vrf" }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetParentYangName() string { return "instance" }

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "message-queue" { return "MessageQueue" }
    if yname == "interface" { return "Interface" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &stats.Interface
    }
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &stats.Interface
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["message-queue"] = stats.MessageQueue
    return leafs
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetParentYangName() string { return "vrf" }

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

// ClearOspfInstanceVrf_Input_Instance_All
// Clear all non-default OSPF VRFs
type ClearOspfInstanceVrf_Input_Instance_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPF process. The type is interface{}.
    Process interface{}

    // Clear OSPF route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPF route table. The type is interface{}.
    Route interface{}

    // OSPF counters and statistics.
    Stats ClearOspfInstanceVrf_Input_Instance_All_Stats
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *ClearOspfInstanceVrf_Input_Instance_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetSegmentPath() string {
    return "all"
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &all.Stats
    }
    return nil
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &all.Stats
    return children
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = all.Process
    leafs["redistribution"] = all.Redistribution
    leafs["route"] = all.Route
    return leafs
}

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetYangName() string { return "all" }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *ClearOspfInstanceVrf_Input_Instance_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetParent() types.Entity { return all.parent }

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetParentYangName() string { return "instance" }

// ClearOspfInstanceVrf_Input_Instance_All_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_All_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "message-queue" { return "MessageQueue" }
    if yname == "interface" { return "Interface" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &stats.Interface
    }
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &stats.Interface
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["message-queue"] = stats.MessageQueue
    return leafs
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetParentYangName() string { return "all" }

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

// ClearOspfInstanceVrf_Input_Instance_AllInclusive
// Clear all non-default and default OSPF VRFs
type ClearOspfInstanceVrf_Input_Instance_AllInclusive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reset OSPF process. The type is interface{}.
    Process interface{}

    // Clear OSPF route redistrbution. The type is interface{}.
    Redistribution interface{}

    // Clear OSPF route table. The type is interface{}.
    Route interface{}

    // OSPF counters and statistics.
    Stats ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetFilter() yfilter.YFilter { return allInclusive.YFilter }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) SetFilter(yf yfilter.YFilter) { allInclusive.YFilter = yf }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetGoName(yname string) string {
    if yname == "process" { return "Process" }
    if yname == "redistribution" { return "Redistribution" }
    if yname == "route" { return "Route" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetSegmentPath() string {
    return "all-inclusive"
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &allInclusive.Stats
    }
    return nil
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &allInclusive.Stats
    return children
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process"] = allInclusive.Process
    leafs["redistribution"] = allInclusive.Redistribution
    leafs["route"] = allInclusive.Route
    return leafs
}

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetBundleName() string { return "cisco_ios_xr" }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetYangName() string { return "all-inclusive" }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) SetParent(parent types.Entity) { allInclusive.parent = parent }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetParent() types.Entity { return allInclusive.parent }

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetParentYangName() string { return "instance" }

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetGoName(yname string) string {
    if yname == "spf" { return "Spf" }
    if yname == "message-queue" { return "MessageQueue" }
    if yname == "interface" { return "Interface" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &stats.Interface
    }
    if childYangName == "neighbor" {
        return &stats.Neighbor
    }
    return nil
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &stats.Interface
    children["neighbor"] = &stats.Neighbor
    return children
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spf"] = stats.Spf
    leafs["message-queue"] = stats.MessageQueue
    return leafs
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetYangName() string { return "stats" }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetParent() types.Entity { return stats.parent }

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetParentYangName() string { return "all-inclusive" }

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetGoName(yname string) string {
    if yname == "neighbor-id" { return "NeighborId" }
    if yname == "interface" { return "Interface" }
    return ""
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetSegmentPath() string {
    return "neighbor"
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        return &neighbor.Interface
    }
    return nil
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface"] = &neighbor.Interface
    return children
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-id"] = neighbor.NeighborId
    return leafs
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetParentYangName() string { return "stats" }

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetSegmentPath() string {
    return "interface"
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetYangName() string { return "interface" }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetParent() types.Entity { return self.parent }

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetParentYangName() string { return "neighbor" }

