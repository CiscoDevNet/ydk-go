// This module contains a collection of YANG definitions
// for Cisco IOS-XR ISIS action package configuration.
// 
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsisProcess_Input
}

func (clearIsisProcess *ClearIsisProcess) GetEntityData() *types.CommonEntityData {
    clearIsisProcess.EntityData.YFilter = clearIsisProcess.YFilter
    clearIsisProcess.EntityData.YangName = "clear-isis-process"
    clearIsisProcess.EntityData.BundleName = "cisco_ios_xr"
    clearIsisProcess.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsisProcess.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis-process"
    clearIsisProcess.EntityData.AbsolutePath = clearIsisProcess.EntityData.SegmentPath
    clearIsisProcess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisProcess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisProcess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisProcess.EntityData.Children = types.NewOrderedMap()
    clearIsisProcess.EntityData.Children.Append("input", types.YChild{"Input", &clearIsisProcess.Input})
    clearIsisProcess.EntityData.Leafs = types.NewOrderedMap()

    clearIsisProcess.EntityData.YListKeys = []string {}

    return &(clearIsisProcess.EntityData)
}

// ClearIsisProcess_Input
type ClearIsisProcess_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear all IS-IS data structures. The type is interface{}.
    Process interface{}

    // Clear data from single IS-IS instance.
    Instance ClearIsisProcess_Input_Instance
}

func (input *ClearIsisProcess_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis-process"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-process/" + input.EntityData.SegmentPath
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

// ClearIsisProcess_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisProcess_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisProcess_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-process/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearIsisRoute
// Clear IS-IS routes
type ClearIsisRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsisRoute_Input
}

func (clearIsisRoute *ClearIsisRoute) GetEntityData() *types.CommonEntityData {
    clearIsisRoute.EntityData.YFilter = clearIsisRoute.YFilter
    clearIsisRoute.EntityData.YangName = "clear-isis-route"
    clearIsisRoute.EntityData.BundleName = "cisco_ios_xr"
    clearIsisRoute.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsisRoute.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis-route"
    clearIsisRoute.EntityData.AbsolutePath = clearIsisRoute.EntityData.SegmentPath
    clearIsisRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisRoute.EntityData.Children = types.NewOrderedMap()
    clearIsisRoute.EntityData.Children.Append("input", types.YChild{"Input", &clearIsisRoute.Input})
    clearIsisRoute.EntityData.Leafs = types.NewOrderedMap()

    clearIsisRoute.EntityData.YListKeys = []string {}

    return &(clearIsisRoute.EntityData)
}

// ClearIsisRoute_Input
type ClearIsisRoute_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear IS-IS routes. The type is interface{}.
    Route interface{}

    // Clear data from single IS-IS instance.
    Instance ClearIsisRoute_Input_Instance
}

func (input *ClearIsisRoute_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis-route"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-route/" + input.EntityData.SegmentPath
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

// ClearIsisRoute_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisRoute_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisRoute_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-route/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearIsisStat
// Clear IS-IS protocol statistics
type ClearIsisStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsisStat_Input
}

func (clearIsisStat *ClearIsisStat) GetEntityData() *types.CommonEntityData {
    clearIsisStat.EntityData.YFilter = clearIsisStat.YFilter
    clearIsisStat.EntityData.YangName = "clear-isis-stat"
    clearIsisStat.EntityData.BundleName = "cisco_ios_xr"
    clearIsisStat.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsisStat.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis-stat"
    clearIsisStat.EntityData.AbsolutePath = clearIsisStat.EntityData.SegmentPath
    clearIsisStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisStat.EntityData.Children = types.NewOrderedMap()
    clearIsisStat.EntityData.Children.Append("input", types.YChild{"Input", &clearIsisStat.Input})
    clearIsisStat.EntityData.Leafs = types.NewOrderedMap()

    clearIsisStat.EntityData.YListKeys = []string {}

    return &(clearIsisStat.EntityData)
}

// ClearIsisStat_Input
type ClearIsisStat_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear data from single IS-IS instance.
    Instance ClearIsisStat_Input_Instance

    // Clear IS-IS protocol statistics.
    Statistics ClearIsisStat_Input_Statistics
}

func (input *ClearIsisStat_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis-stat"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-stat/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Children.Append("statistics", types.YChild{"Statistics", &input.Statistics})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearIsisStat_Input_Instance
// Clear data from single IS-IS instance
type ClearIsisStat_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisStat_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-stat/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearIsisStat_Input_Statistics
// Clear IS-IS protocol statistics
type ClearIsisStat_Input_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+. This
    // attribute is mandatory.
    InterfaceName interface{}
}

func (statistics *ClearIsisStat_Input_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "input"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-stat/input/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statistics.InterfaceName})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// ClearIsisDist
// Reset BGP-LS topology distribution
type ClearIsisDist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsisDist_Input
}

func (clearIsisDist *ClearIsisDist) GetEntityData() *types.CommonEntityData {
    clearIsisDist.EntityData.YFilter = clearIsisDist.YFilter
    clearIsisDist.EntityData.YangName = "clear-isis-dist"
    clearIsisDist.EntityData.BundleName = "cisco_ios_xr"
    clearIsisDist.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsisDist.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis-dist"
    clearIsisDist.EntityData.AbsolutePath = clearIsisDist.EntityData.SegmentPath
    clearIsisDist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisDist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisDist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisDist.EntityData.Children = types.NewOrderedMap()
    clearIsisDist.EntityData.Children.Append("input", types.YChild{"Input", &clearIsisDist.Input})
    clearIsisDist.EntityData.Leafs = types.NewOrderedMap()

    clearIsisDist.EntityData.YListKeys = []string {}

    return &(clearIsisDist.EntityData)
}

// ClearIsisDist_Input
type ClearIsisDist_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reset BGP-LS topology distribution. The type is interface{}.
    Distribution interface{}

    // Reset BGP-LS topology from single IS-IS instance.
    Instance ClearIsisDist_Input_Instance
}

func (input *ClearIsisDist_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis-dist"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-dist/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("distribution", types.YLeaf{"Distribution", input.Distribution})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearIsisDist_Input_Instance
// Reset BGP-LS topology from single IS-IS instance
type ClearIsisDist_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisDist_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-dist/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearIsisLocalLsp
// Clean and regenerate local LSPs
type ClearIsisLocalLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsisLocalLsp_Input
}

func (clearIsisLocalLsp *ClearIsisLocalLsp) GetEntityData() *types.CommonEntityData {
    clearIsisLocalLsp.EntityData.YFilter = clearIsisLocalLsp.YFilter
    clearIsisLocalLsp.EntityData.YangName = "clear-isis-local-lsp"
    clearIsisLocalLsp.EntityData.BundleName = "cisco_ios_xr"
    clearIsisLocalLsp.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsisLocalLsp.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis-local-lsp"
    clearIsisLocalLsp.EntityData.AbsolutePath = clearIsisLocalLsp.EntityData.SegmentPath
    clearIsisLocalLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisLocalLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisLocalLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisLocalLsp.EntityData.Children = types.NewOrderedMap()
    clearIsisLocalLsp.EntityData.Children.Append("input", types.YChild{"Input", &clearIsisLocalLsp.Input})
    clearIsisLocalLsp.EntityData.Leafs = types.NewOrderedMap()

    clearIsisLocalLsp.EntityData.YListKeys = []string {}

    return &(clearIsisLocalLsp.EntityData)
}

// ClearIsisLocalLsp_Input
type ClearIsisLocalLsp_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clean and regenerate local LSPs. The type is interface{}.
    LocalLsp interface{}

    // Clean and regenerate local LSPs from single IS-IS instance.
    Instance ClearIsisLocalLsp_Input_Instance
}

func (input *ClearIsisLocalLsp_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis-local-lsp"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-local-lsp/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("local-lsp", types.YLeaf{"LocalLsp", input.LocalLsp})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearIsisLocalLsp_Input_Instance
// Clean and regenerate local LSPs from single IS-IS instance
type ClearIsisLocalLsp_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsisLocalLsp_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis-local-lsp/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// ClearIsis
// Clear IS-IS data structures
type ClearIsis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearIsis_Input
}

func (clearIsis *ClearIsis) GetEntityData() *types.CommonEntityData {
    clearIsis.EntityData.YFilter = clearIsis.YFilter
    clearIsis.EntityData.YangName = "clear-isis"
    clearIsis.EntityData.BundleName = "cisco_ios_xr"
    clearIsis.EntityData.ParentYangName = "Cisco-IOS-XR-isis-act"
    clearIsis.EntityData.SegmentPath = "Cisco-IOS-XR-isis-act:clear-isis"
    clearIsis.EntityData.AbsolutePath = clearIsis.EntityData.SegmentPath
    clearIsis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsis.EntityData.Children = types.NewOrderedMap()
    clearIsis.EntityData.Children.Append("input", types.YChild{"Input", &clearIsis.Input})
    clearIsis.EntityData.Leafs = types.NewOrderedMap()

    clearIsis.EntityData.YListKeys = []string {}

    return &(clearIsis.EntityData)
}

// ClearIsis_Input
type ClearIsis_Input struct {
    EntityData types.CommonEntityData
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

func (input *ClearIsis_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-isis"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("instance", types.YChild{"Instance", &input.Instance})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("rt-type", types.YLeaf{"RtType", input.RtType})
    input.EntityData.Leafs.Append("route", types.YLeaf{"Route", input.Route})
    input.EntityData.Leafs.Append("topology", types.YLeaf{"Topology", input.Topology})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearIsis_Input_Instance
// Clear data from single IS-IS instance
type ClearIsis_Input_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS process instance identifier. The type is string.
    InstanceIdentifier interface{}
}

func (instance *ClearIsis_Input_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "input"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-isis-act:clear-isis/input/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("instance-identifier", types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

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

