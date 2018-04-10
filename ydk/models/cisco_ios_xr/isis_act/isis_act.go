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
    clearIsisProcess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisProcess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisProcess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisProcess.EntityData.Children = make(map[string]types.YChild)
    clearIsisProcess.EntityData.Children["input"] = types.YChild{"Input", &clearIsisProcess.Input}
    clearIsisProcess.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["process"] = types.YLeaf{"Process", input.Process}
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
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
    clearIsisRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisRoute.EntityData.Children = make(map[string]types.YChild)
    clearIsisRoute.EntityData.Children["input"] = types.YChild{"Input", &clearIsisRoute.Input}
    clearIsisRoute.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["route"] = types.YLeaf{"Route", input.Route}
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
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
    clearIsisStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisStat.EntityData.Children = make(map[string]types.YChild)
    clearIsisStat.EntityData.Children["input"] = types.YChild{"Input", &clearIsisStat.Input}
    clearIsisStat.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Children["statistics"] = types.YChild{"Statistics", &input.Statistics}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
    return &(instance.EntityData)
}

// ClearIsisStat_Input_Statistics
// Clear IS-IS protocol statistics
type ClearIsisStat_Input_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'. This
    // attribute is mandatory.
    InterfaceName interface{}
}

func (statistics *ClearIsisStat_Input_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "input"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", statistics.InterfaceName}
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
    clearIsisDist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisDist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisDist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisDist.EntityData.Children = make(map[string]types.YChild)
    clearIsisDist.EntityData.Children["input"] = types.YChild{"Input", &clearIsisDist.Input}
    clearIsisDist.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["distribution"] = types.YLeaf{"Distribution", input.Distribution}
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
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
    clearIsisLocalLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsisLocalLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsisLocalLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsisLocalLsp.EntityData.Children = make(map[string]types.YChild)
    clearIsisLocalLsp.EntityData.Children["input"] = types.YChild{"Input", &clearIsisLocalLsp.Input}
    clearIsisLocalLsp.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["local-lsp"] = types.YLeaf{"LocalLsp", input.LocalLsp}
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
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
    clearIsis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearIsis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearIsis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearIsis.EntityData.Children = make(map[string]types.YChild)
    clearIsis.EntityData.Children["input"] = types.YChild{"Input", &clearIsis.Input}
    clearIsis.EntityData.Leafs = make(map[string]types.YLeaf)
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
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["instance"] = types.YChild{"Instance", &input.Instance}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["rt-type"] = types.YLeaf{"RtType", input.RtType}
    input.EntityData.Leafs["route"] = types.YLeaf{"Route", input.Route}
    input.EntityData.Leafs["topology"] = types.YLeaf{"Topology", input.Topology}
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
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-identifier"] = types.YLeaf{"InstanceIdentifier", instance.InstanceIdentifier}
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

