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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Routes_Input
}

func (clearOspfv3Routes *ClearOspfv3Routes) GetEntityData() *types.CommonEntityData {
    clearOspfv3Routes.EntityData.YFilter = clearOspfv3Routes.YFilter
    clearOspfv3Routes.EntityData.YangName = "clear-ospfv3-routes"
    clearOspfv3Routes.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3Routes.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3Routes.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-routes"
    clearOspfv3Routes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3Routes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3Routes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3Routes.EntityData.Children = types.NewOrderedMap()
    clearOspfv3Routes.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3Routes.Input})
    clearOspfv3Routes.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3Routes.EntityData.YListKeys = []string {}

    return &(clearOspfv3Routes.EntityData)
}

// ClearOspfv3Routes_Input
type ClearOspfv3Routes_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear OSPFv3 route table. The type is interface{}. This attribute is
    // mandatory.
    Route interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Routes_Input_Instance
}

func (input *ClearOspfv3Routes_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-routes"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("route", types.YLeaf{"Route", input.Route})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3Routes_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Routes_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Routes_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3Redistribution
// Clear OSPFv3 route redistribution
type ClearOspfv3Redistribution struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Redistribution_Input
}

func (clearOspfv3Redistribution *ClearOspfv3Redistribution) GetEntityData() *types.CommonEntityData {
    clearOspfv3Redistribution.EntityData.YFilter = clearOspfv3Redistribution.YFilter
    clearOspfv3Redistribution.EntityData.YangName = "clear-ospfv3-redistribution"
    clearOspfv3Redistribution.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3Redistribution.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3Redistribution.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-redistribution"
    clearOspfv3Redistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3Redistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3Redistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3Redistribution.EntityData.Children = types.NewOrderedMap()
    clearOspfv3Redistribution.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3Redistribution.Input})
    clearOspfv3Redistribution.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3Redistribution.EntityData.YListKeys = []string {}

    return &(clearOspfv3Redistribution.EntityData)
}

// ClearOspfv3Redistribution_Input
type ClearOspfv3Redistribution_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear OSPFv3 route redistribution. The type is interface{}. This attribute
    // is mandatory.
    Redistribution interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Redistribution_Input_Instance
}

func (input *ClearOspfv3Redistribution_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-redistribution"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("redistribution", types.YLeaf{"Redistribution", input.Redistribution})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3Redistribution_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Redistribution_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Redistribution_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3Process
// Clear (reset) OSPFv3 Process
type ClearOspfv3Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Process_Input
}

func (clearOspfv3Process *ClearOspfv3Process) GetEntityData() *types.CommonEntityData {
    clearOspfv3Process.EntityData.YFilter = clearOspfv3Process.YFilter
    clearOspfv3Process.EntityData.YangName = "clear-ospfv3-process"
    clearOspfv3Process.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3Process.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3Process.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-process"
    clearOspfv3Process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3Process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3Process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3Process.EntityData.Children = types.NewOrderedMap()
    clearOspfv3Process.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3Process.Input})
    clearOspfv3Process.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3Process.EntityData.YListKeys = []string {}

    return &(clearOspfv3Process.EntityData)
}

// ClearOspfv3Process_Input
type ClearOspfv3Process_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reset OSPFv3 process. The type is interface{}. This attribute is mandatory.
    Process interface{}

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3Process_Input_Instance
}

func (input *ClearOspfv3Process_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-process"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("process", types.YLeaf{"Process", input.Process})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3Process_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Process_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Process_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3StatisticsNeighbor
// Clear OSPFv3 neighbor statistics per interface or neighbor id
type ClearOspfv3StatisticsNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3StatisticsNeighbor_Input
}

func (clearOspfv3StatisticsNeighbor *ClearOspfv3StatisticsNeighbor) GetEntityData() *types.CommonEntityData {
    clearOspfv3StatisticsNeighbor.EntityData.YFilter = clearOspfv3StatisticsNeighbor.YFilter
    clearOspfv3StatisticsNeighbor.EntityData.YangName = "clear-ospfv3-statistics-neighbor"
    clearOspfv3StatisticsNeighbor.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3StatisticsNeighbor.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3StatisticsNeighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics-neighbor"
    clearOspfv3StatisticsNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3StatisticsNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3StatisticsNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3StatisticsNeighbor.EntityData.Children = types.NewOrderedMap()
    clearOspfv3StatisticsNeighbor.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3StatisticsNeighbor.Input})
    clearOspfv3StatisticsNeighbor.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3StatisticsNeighbor.EntityData.YListKeys = []string {}

    return &(clearOspfv3StatisticsNeighbor.EntityData)
}

// ClearOspfv3StatisticsNeighbor_Input
type ClearOspfv3StatisticsNeighbor_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear data from OSPFv3 instance.
    Instance ClearOspfv3StatisticsNeighbor_Input_Instance

    
    Neighbor ClearOspfv3StatisticsNeighbor_Input_Neighbor
}

func (input *ClearOspfv3StatisticsNeighbor_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-statistics-neighbor"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &input.Neighbor})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3StatisticsNeighbor_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3StatisticsNeighbor_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3StatisticsNeighbor_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3StatisticsNeighbor_Input_Neighbor
type ClearOspfv3StatisticsNeighbor_Input_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (neighbor *ClearOspfv3StatisticsNeighbor_Input_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "input"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-id", types.YLeaf{"NeighborId", neighbor.NeighborId})
    neighbor.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", neighbor.InterfaceName})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// ClearOspfv3Statistics
// Clear OSPFv3 counters and statistics
type ClearOspfv3Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3Statistics_Input
}

func (clearOspfv3Statistics *ClearOspfv3Statistics) GetEntityData() *types.CommonEntityData {
    clearOspfv3Statistics.EntityData.YFilter = clearOspfv3Statistics.YFilter
    clearOspfv3Statistics.EntityData.YangName = "clear-ospfv3-statistics"
    clearOspfv3Statistics.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3Statistics.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3Statistics.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-statistics"
    clearOspfv3Statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3Statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3Statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3Statistics.EntityData.Children = types.NewOrderedMap()
    clearOspfv3Statistics.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3Statistics.Input})
    clearOspfv3Statistics.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3Statistics.EntityData.YListKeys = []string {}

    return &(clearOspfv3Statistics.EntityData)
}

// ClearOspfv3Statistics_Input
type ClearOspfv3Statistics_Input struct {
    EntityData types.CommonEntityData
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

func (input *ClearOspfv3Statistics_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-statistics"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("prefix-priority", types.YLeaf{"PrefixPriority", input.PrefixPriority})
    input.EntityData.Leafs.Append("spf", types.YLeaf{"Spf", input.Spf})
    input.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", input.Neighbor})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3Statistics_Input_Instance
// Clear data from OSPFv3 instance
type ClearOspfv3Statistics_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearOspfv3Statistics_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3InstanceVrf
// Clear one or more non-default OSPFv3 VRFs in process
type ClearOspfv3InstanceVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearOspfv3InstanceVrf_Input
}

func (clearOspfv3InstanceVrf *ClearOspfv3InstanceVrf) GetEntityData() *types.CommonEntityData {
    clearOspfv3InstanceVrf.EntityData.YFilter = clearOspfv3InstanceVrf.YFilter
    clearOspfv3InstanceVrf.EntityData.YangName = "clear-ospfv3-instance-vrf"
    clearOspfv3InstanceVrf.EntityData.BundleName = "cisco_ios_xr"
    clearOspfv3InstanceVrf.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-act"
    clearOspfv3InstanceVrf.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-act:clear-ospfv3-instance-vrf"
    clearOspfv3InstanceVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearOspfv3InstanceVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearOspfv3InstanceVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearOspfv3InstanceVrf.EntityData.Children = types.NewOrderedMap()
    clearOspfv3InstanceVrf.EntityData.Children.Append("input", types.YChild{"Input", &clearOspfv3InstanceVrf.Input})
    clearOspfv3InstanceVrf.EntityData.Leafs = types.NewOrderedMap()

    clearOspfv3InstanceVrf.EntityData.YListKeys = []string {}

    return &(clearOspfv3InstanceVrf.EntityData)
}

// ClearOspfv3InstanceVrf_Input
type ClearOspfv3InstanceVrf_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 instance name.
    Instance ClearOspfv3InstanceVrf_Input_Instance
}

func (input *ClearOspfv3InstanceVrf_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-ospfv3-instance-vrf"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance
// OSPFv3 instance name
type ClearOspfv3InstanceVrf_Input_Instance struct {
    EntityData types.CommonEntityData
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

func (instance *ClearOspfv3InstanceVrf_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("vrf", types.YChild{"Vrf", &instance.Vrf})
    instance.EntityData.Children.Append("all", types.YChild{"All", &instance.All})
    instance.EntityData.Children.Append("all-inclusive", types.YChild{"AllInclusive", &instance.AllInclusive})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_Vrf
// Clear one or more non-default OSPFv3 VRFs in process
type ClearOspfv3InstanceVrf_Input_Instance_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *ClearOspfv3InstanceVrf_Input_Instance_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "instance"
    vrf.EntityData.SegmentPath = "vrf"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("stats", types.YChild{"Stats", &vrf.Stats})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("process", types.YLeaf{"Process", vrf.Process})
    vrf.EntityData.Leafs.Append("redistribution", types.YLeaf{"Redistribution", vrf.Redistribution})
    vrf.EntityData.Leafs.Append("route", types.YLeaf{"Route", vrf.Route})

    vrf.EntityData.YListKeys = []string {}

    return &(vrf.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "vrf"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &stats.Neighbor})
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("spf", types.YLeaf{"Spf", stats.Spf})
    stats.EntityData.Leafs.Append("prefix-priority", types.YLeaf{"PrefixPriority", stats.PrefixPriority})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("interface", types.YChild{"Interface", &neighbor.Interface})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-id", types.YLeaf{"NeighborId", neighbor.NeighborId})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_Vrf_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_All
// Clear all non-default OSPFv3 VRFs
type ClearOspfv3InstanceVrf_Input_Instance_All struct {
    EntityData types.CommonEntityData
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

func (all *ClearOspfv3InstanceVrf_Input_Instance_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "instance"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("stats", types.YChild{"Stats", &all.Stats})
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("process", types.YLeaf{"Process", all.Process})
    all.EntityData.Leafs.Append("redistribution", types.YLeaf{"Redistribution", all.Redistribution})
    all.EntityData.Leafs.Append("route", types.YLeaf{"Route", all.Route})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_All_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "all"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &stats.Neighbor})
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("spf", types.YLeaf{"Spf", stats.Spf})
    stats.EntityData.Leafs.Append("prefix-priority", types.YLeaf{"PrefixPriority", stats.PrefixPriority})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("interface", types.YChild{"Interface", &neighbor.Interface})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-id", types.YLeaf{"NeighborId", neighbor.NeighborId})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_All_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive
// Clear all non-default and default OSPFv3 VRFs
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive struct {
    EntityData types.CommonEntityData
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

func (allInclusive *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive) GetEntityData() *types.CommonEntityData {
    allInclusive.EntityData.YFilter = allInclusive.YFilter
    allInclusive.EntityData.YangName = "all-inclusive"
    allInclusive.EntityData.BundleName = "cisco_ios_xr"
    allInclusive.EntityData.ParentYangName = "instance"
    allInclusive.EntityData.SegmentPath = "all-inclusive"
    allInclusive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allInclusive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allInclusive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allInclusive.EntityData.Children = types.NewOrderedMap()
    allInclusive.EntityData.Children.Append("stats", types.YChild{"Stats", &allInclusive.Stats})
    allInclusive.EntityData.Leafs = types.NewOrderedMap()
    allInclusive.EntityData.Leafs.Append("process", types.YLeaf{"Process", allInclusive.Process})
    allInclusive.EntityData.Leafs.Append("redistribution", types.YLeaf{"Redistribution", allInclusive.Redistribution})
    allInclusive.EntityData.Leafs.Append("route", types.YLeaf{"Route", allInclusive.Route})

    allInclusive.EntityData.YListKeys = []string {}

    return &(allInclusive.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats
// OSPFv3 counters and statistics
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPF statistics. The type is interface{}.
    Spf interface{}

    // SPF Prefix Priority statistics. The type is interface{}.
    PrefixPriority interface{}

    // Neighbor statistics per interface or neighbor id.
    Neighbor ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
}

func (stats *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "all-inclusive"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", &stats.Neighbor})
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("spf", types.YLeaf{"Spf", stats.Spf})
    stats.EntityData.Leafs.Append("prefix-priority", types.YLeaf{"PrefixPriority", stats.PrefixPriority})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor
// Neighbor statistics per interface or neighbor id
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NeighborId interface{}

    
    Interface ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
}

func (neighbor *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "stats"
    neighbor.EntityData.SegmentPath = "neighbor"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("interface", types.YChild{"Interface", &neighbor.Interface})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-id", types.YLeaf{"NeighborId", neighbor.NeighborId})

    neighbor.EntityData.YListKeys = []string {}

    return &(neighbor.EntityData)
}

// ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface
type ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface statistics. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *ClearOspfv3InstanceVrf_Input_Instance_AllInclusive_Stats_Neighbor_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "neighbor"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

