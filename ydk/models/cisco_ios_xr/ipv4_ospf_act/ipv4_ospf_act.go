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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfRoutes_Input
}

func (clearOspfRoutes *ClearOspfRoutes) GetEntityData() *types.CommonEntityData {
    clearOspfRoutes.EntityData.YFilter = clearOspfRoutes.YFilter
    clearOspfRoutes.EntityData.YangName = "clear-ospf-routes"
    clearOspfRoutes.EntityData.BundleName = "cisco_ios_xr"
    clearOspfRoutes.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfRoutes.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-routes"
    clearOspfRoutes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfRoutes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfRoutes.EntityData.Children = make(map[string]types.YChild)
    clearOspfRoutes.EntityData.Children["input"] = types.YChild{"Input", &clearOspfRoutes.Input}
    clearOspfRoutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfRoutes.EntityData)
}

// ClearOspfRoutes_Input
type ClearOspfRoutes_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear OSPF route table. The type is interface{}. This attribute is
    // mandatory.
    Route interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfRoutes_Input_Instance
}

func (input *ClearOspfRoutes_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-routes"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["route"] = types.YLeaf{"Route", input.Route}
    return &(input.EntityData)
}

// ClearOspfRoutes_Input_Instance
// Clear data from OSPF instance
type ClearOspfRoutes_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfRoutes_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfRedistribution
// Clear OSPF route redistribution
type ClearOspfRedistribution struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfRedistribution_Input
}

func (clearOspfRedistribution *ClearOspfRedistribution) GetEntityData() *types.CommonEntityData {
    clearOspfRedistribution.EntityData.YFilter = clearOspfRedistribution.YFilter
    clearOspfRedistribution.EntityData.YangName = "clear-ospf-redistribution"
    clearOspfRedistribution.EntityData.BundleName = "cisco_ios_xr"
    clearOspfRedistribution.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfRedistribution.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-redistribution"
    clearOspfRedistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfRedistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfRedistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfRedistribution.EntityData.Children = make(map[string]types.YChild)
    clearOspfRedistribution.EntityData.Children["input"] = types.YChild{"Input", &clearOspfRedistribution.Input}
    clearOspfRedistribution.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfRedistribution.EntityData)
}

// ClearOspfRedistribution_Input
type ClearOspfRedistribution_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear OSPF route redistribution. The type is interface{}. This attribute is
    // mandatory.
    Redistribution interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfRedistribution_Input_Instance
}

func (input *ClearOspfRedistribution_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-redistribution"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["redistribution"] = types.YLeaf{"Redistribution", input.Redistribution}
    return &(input.EntityData)
}

// ClearOspfRedistribution_Input_Instance
// Clear data from OSPF instance
type ClearOspfRedistribution_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfRedistribution_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfStatistics
// Clear OSPF counters and statistics
type ClearOspfStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfStatistics_Input
}

func (clearOspfStatistics *ClearOspfStatistics) GetEntityData() *types.CommonEntityData {
    clearOspfStatistics.EntityData.YFilter = clearOspfStatistics.YFilter
    clearOspfStatistics.EntityData.YangName = "clear-ospf-statistics"
    clearOspfStatistics.EntityData.BundleName = "cisco_ios_xr"
    clearOspfStatistics.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfStatistics.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics"
    clearOspfStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfStatistics.EntityData.Children = make(map[string]types.YChild)
    clearOspfStatistics.EntityData.Children["input"] = types.YChild{"Input", &clearOspfStatistics.Input}
    clearOspfStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfStatistics.EntityData)
}

// ClearOspfStatistics_Input
type ClearOspfStatistics_Input struct {
    EntityData types.CommonEntityData
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

func (input *ClearOspfStatistics_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-statistics"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["all"] = types.YLeaf{"All", input.All}
    input.EntityData.Leafs["message-queue"] = types.YLeaf{"MessageQueue", input.MessageQueue}
    input.EntityData.Leafs["spf"] = types.YLeaf{"Spf", input.Spf}
    input.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", input.Neighbor}
    input.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", input.InterfaceName}
    return &(input.EntityData)
}

// ClearOspfStatistics_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatistics_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatistics_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfStatisticsNeighbor
// Clear OSPF neighbor statistics per interface or neighbor id
type ClearOspfStatisticsNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfStatisticsNeighbor_Input
}

func (clearOspfStatisticsNeighbor *ClearOspfStatisticsNeighbor) GetEntityData() *types.CommonEntityData {
    clearOspfStatisticsNeighbor.EntityData.YFilter = clearOspfStatisticsNeighbor.YFilter
    clearOspfStatisticsNeighbor.EntityData.YangName = "clear-ospf-statistics-neighbor"
    clearOspfStatisticsNeighbor.EntityData.BundleName = "cisco_ios_xr"
    clearOspfStatisticsNeighbor.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfStatisticsNeighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-neighbor"
    clearOspfStatisticsNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfStatisticsNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfStatisticsNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfStatisticsNeighbor.EntityData.Children = make(map[string]types.YChild)
    clearOspfStatisticsNeighbor.EntityData.Children["input"] = types.YChild{"Input", &clearOspfStatisticsNeighbor.Input}
    clearOspfStatisticsNeighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfStatisticsNeighbor.EntityData)
}

// ClearOspfStatisticsNeighbor_Input
type ClearOspfStatisticsNeighbor_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear data from OSPF instance.
    Instance ClearOspfStatisticsNeighbor_Input_Instance

    
    Neighbor ClearOspfStatisticsNeighbor_Input_Neighbor
}

func (input *ClearOspfStatisticsNeighbor_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-statistics-neighbor"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Children["neighbor"] = types.YChild{"Neighbor", &input.Neighbor}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// ClearOspfStatisticsNeighbor_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatisticsNeighbor_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatisticsNeighbor_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfStatisticsNeighbor_Input_Neighbor
type ClearOspfStatisticsNeighbor_Input_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (neighbor *ClearOspfStatisticsNeighbor_Input_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "input"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    neighbor.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", neighbor.InterfaceName}
    return &(neighbor.EntityData)
}

// ClearOspfStatisticsInterface
// Clear OSPF interface statistics
type ClearOspfStatisticsInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfStatisticsInterface_Input
}

func (clearOspfStatisticsInterface *ClearOspfStatisticsInterface) GetEntityData() *types.CommonEntityData {
    clearOspfStatisticsInterface.EntityData.YFilter = clearOspfStatisticsInterface.YFilter
    clearOspfStatisticsInterface.EntityData.YangName = "clear-ospf-statistics-interface"
    clearOspfStatisticsInterface.EntityData.BundleName = "cisco_ios_xr"
    clearOspfStatisticsInterface.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfStatisticsInterface.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-statistics-interface"
    clearOspfStatisticsInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfStatisticsInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfStatisticsInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfStatisticsInterface.EntityData.Children = make(map[string]types.YChild)
    clearOspfStatisticsInterface.EntityData.Children["input"] = types.YChild{"Input", &clearOspfStatisticsInterface.Input}
    clearOspfStatisticsInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfStatisticsInterface.EntityData)
}

// ClearOspfStatisticsInterface_Input
type ClearOspfStatisticsInterface_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear data from OSPF instance.
    Instance ClearOspfStatisticsInterface_Input_Instance

    
    Interface_ ClearOspfStatisticsInterface_Input_Interface
}

func (input *ClearOspfStatisticsInterface_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-statistics-interface"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Children["interface"] = types.YChild{"Interface_", &input.Interface_}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// ClearOspfStatisticsInterface_Input_Instance
// Clear data from OSPF instance
type ClearOspfStatisticsInterface_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfStatisticsInterface_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfStatisticsInterface_Input_Interface
type ClearOspfStatisticsInterface_Input_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfStatisticsInterface_Input_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "input"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfProcess
// Clear (reset) OSPF process
type ClearOspfProcess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfProcess_Input
}

func (clearOspfProcess *ClearOspfProcess) GetEntityData() *types.CommonEntityData {
    clearOspfProcess.EntityData.YFilter = clearOspfProcess.YFilter
    clearOspfProcess.EntityData.YangName = "clear-ospf-process"
    clearOspfProcess.EntityData.BundleName = "cisco_ios_xr"
    clearOspfProcess.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfProcess.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-process"
    clearOspfProcess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfProcess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfProcess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfProcess.EntityData.Children = make(map[string]types.YChild)
    clearOspfProcess.EntityData.Children["input"] = types.YChild{"Input", &clearOspfProcess.Input}
    clearOspfProcess.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfProcess.EntityData)
}

// ClearOspfProcess_Input
type ClearOspfProcess_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reset OSPF process. The type is interface{}. This attribute is mandatory.
    Process interface{}

    // Clear data from OSPF instance.
    Instance ClearOspfProcess_Input_Instance
}

func (input *ClearOspfProcess_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-process"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["process"] = types.YLeaf{"Process", input.Process}
    return &(input.EntityData)
}

// ClearOspfProcess_Input_Instance
// Clear data from OSPF instance
type ClearOspfProcess_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfProcess_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfInstanceVrf
// Clear one or more non-default OSPF VRFs in process
type ClearOspfInstanceVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfInstanceVrf_Input
}

func (clearOspfInstanceVrf *ClearOspfInstanceVrf) GetEntityData() *types.CommonEntityData {
    clearOspfInstanceVrf.EntityData.YFilter = clearOspfInstanceVrf.YFilter
    clearOspfInstanceVrf.EntityData.YangName = "clear-ospf-instance-vrf"
    clearOspfInstanceVrf.EntityData.BundleName = "cisco_ios_xr"
    clearOspfInstanceVrf.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ospf-act"
    clearOspfInstanceVrf.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-act:clear-ospf-instance-vrf"
    clearOspfInstanceVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfInstanceVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfInstanceVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfInstanceVrf.EntityData.Children = make(map[string]types.YChild)
    clearOspfInstanceVrf.EntityData.Children["input"] = types.YChild{"Input", &clearOspfInstanceVrf.Input}
    clearOspfInstanceVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearOspfInstanceVrf.EntityData)
}

// ClearOspfInstanceVrf_Input
type ClearOspfInstanceVrf_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF instance name.
    Instance ClearOspfInstanceVrf_Input_Instance
}

func (input *ClearOspfInstanceVrf_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospf-instance-vrf"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance
// OSPF instance name
type ClearOspfInstanceVrf_Input_Instance struct {
    EntityData types.CommonEntityData
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

func (instance *ClearOspfInstanceVrf_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["vrf"] = types.YChild{"Vrf", &instance.Vrf}
    instance.EntityData.Children["all"] = types.YChild{"All", &instance.All}
    instance.EntityData.Children["all-inclusive"] = types.YChild{"AllInclusive", &instance.AllInclusive}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_Vrf
// Clear one or more non-default OSPF VRFs in process
type ClearOspfInstanceVrf_Input_Instance_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *ClearOspfInstanceVrf_Input_Instance_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "instance"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["stats"] = types.YChild{"Stats", &vrf.Stats}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["process"] = types.YLeaf{"Process", vrf.Process}
    vrf.EntityData.Leafs["redistribution"] = types.YLeaf{"Redistribution", vrf.Redistribution}
    vrf.EntityData.Leafs["route"] = types.YLeaf{"Route", vrf.Route}
    return &(vrf.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface_ ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "vrf"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["interface"] = types.YChild{"Interface_", &stats.Interface_}
    stats.EntityData.Children["neighbor"] = types.YChild{"Neighbor", &stats.Neighbor}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["spf"] = types.YLeaf{"Spf", stats.Spf}
    stats.EntityData.Leafs["message-queue"] = types.YLeaf{"MessageQueue", stats.MessageQueue}
    return &(stats.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "stats"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    
    Interface_ ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["interface"] = types.YChild{"Interface_", &neighbor.Interface_}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    return &(neighbor.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_All
// Clear all non-default OSPF VRFs
type ClearOspfInstanceVrf_Input_Instance_All struct {
    EntityData types.CommonEntityData
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

func (all *ClearOspfInstanceVrf_Input_Instance_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "instance"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["stats"] = types.YChild{"Stats", &all.Stats}
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    all.EntityData.Leafs["process"] = types.YLeaf{"Process", all.Process}
    all.EntityData.Leafs["redistribution"] = types.YLeaf{"Redistribution", all.Redistribution}
    all.EntityData.Leafs["route"] = types.YLeaf{"Route", all.Route}
    return &(all.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_All_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_All_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface_ ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_All_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "all"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["interface"] = types.YChild{"Interface_", &stats.Interface_}
    stats.EntityData.Children["neighbor"] = types.YChild{"Neighbor", &stats.Neighbor}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["spf"] = types.YLeaf{"Spf", stats.Spf}
    stats.EntityData.Leafs["message-queue"] = types.YLeaf{"MessageQueue", stats.MessageQueue}
    return &(stats.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "stats"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    
    Interface_ ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["interface"] = types.YChild{"Interface_", &neighbor.Interface_}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    return &(neighbor.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_AllInclusive
// Clear all non-default and default OSPF VRFs
type ClearOspfInstanceVrf_Input_Instance_AllInclusive struct {
    EntityData types.CommonEntityData
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

func (allInclusive *ClearOspfInstanceVrf_Input_Instance_AllInclusive) GetEntityData() *types.CommonEntityData {
    allInclusive.EntityData.YFilter = allInclusive.YFilter
    allInclusive.EntityData.YangName = "all-inclusive"
    allInclusive.EntityData.BundleName = "cisco_ios_xr"
    allInclusive.EntityData.ParentYangName = "instance"
    allInclusive.EntityData.SegmentPath = "all-inclusive"
    allInclusive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInclusive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInclusive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInclusive.EntityData.Children = make(map[string]types.YChild)
    allInclusive.EntityData.Children["stats"] = types.YChild{"Stats", &allInclusive.Stats}
    allInclusive.EntityData.Leafs = make(map[string]types.YLeaf)
    allInclusive.EntityData.Leafs["process"] = types.YLeaf{"Process", allInclusive.Process}
    allInclusive.EntityData.Leafs["redistribution"] = types.YLeaf{"Redistribution", allInclusive.Redistribution}
    allInclusive.EntityData.Leafs["route"] = types.YLeaf{"Route", allInclusive.Route}
    return &(allInclusive.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats
// OSPF counters and statistics
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // Message-queue statistics. The type is interface{}.
    MessageQueue interface{}

    // OSPF interface statistics.
    Interface_ ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
}

func (stats *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "all-inclusive"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["interface"] = types.YChild{"Interface_", &stats.Interface_}
    stats.EntityData.Children["neighbor"] = types.YChild{"Neighbor", &stats.Neighbor}
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    stats.EntityData.Leafs["spf"] = types.YLeaf{"Spf", stats.Spf}
    stats.EntityData.Leafs["message-queue"] = types.YLeaf{"MessageQueue", stats.MessageQueue}
    return &(stats.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface
// OSPF interface statistics
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "stats"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborId interface{}

    
    Interface_ ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Children["interface"] = types.YChild{"Interface_", &neighbor.Interface_}
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-id"] = types.YLeaf{"NeighborId", neighbor.NeighborId}
    return &(neighbor.EntityData)
}

// ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
type ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF interface statistics. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *ClearOspfInstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

