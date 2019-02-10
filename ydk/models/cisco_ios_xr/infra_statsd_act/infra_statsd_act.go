// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_statsd_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_statsd_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-statsd-act clear-counters-controller}", reflect.TypeOf(ClearCountersController{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-statsd-act:clear-counters-controller", reflect.TypeOf(ClearCountersController{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-statsd-act clear-counters-all}", reflect.TypeOf(ClearCountersAll{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-statsd-act:clear-counters-all", reflect.TypeOf(ClearCountersAll{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-statsd-act clear-counters-interface}", reflect.TypeOf(ClearCountersInterface{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-statsd-act:clear-counters-interface", reflect.TypeOf(ClearCountersInterface{}))
}

// ClearCountersController
// Controller name.
// 
type ClearCountersController struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearCountersController_Input
}

func (clearCountersController *ClearCountersController) GetEntityData() *types.CommonEntityData {
    clearCountersController.EntityData.YFilter = clearCountersController.YFilter
    clearCountersController.EntityData.YangName = "clear-counters-controller"
    clearCountersController.EntityData.BundleName = "cisco_ios_xr"
    clearCountersController.EntityData.ParentYangName = "Cisco-IOS-XR-infra-statsd-act"
    clearCountersController.EntityData.SegmentPath = "Cisco-IOS-XR-infra-statsd-act:clear-counters-controller"
    clearCountersController.EntityData.AbsolutePath = clearCountersController.EntityData.SegmentPath
    clearCountersController.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearCountersController.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearCountersController.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearCountersController.EntityData.Children = types.NewOrderedMap()
    clearCountersController.EntityData.Children.Append("input", types.YChild{"Input", &clearCountersController.Input})
    clearCountersController.EntityData.Leafs = types.NewOrderedMap()

    clearCountersController.EntityData.YListKeys = []string {}

    return &(clearCountersController.EntityData)
}

// ClearCountersController_Input
type ClearCountersController_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller name. The type is string with pattern: [a-zA-Z0-9._/-]+. This
    // attribute is mandatory.
    ControllerName interface{}
}

func (input *ClearCountersController_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-counters-controller"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-statsd-act:clear-counters-controller/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", input.ControllerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearCountersAll
// Clear counters on all interfaces.
// 
type ClearCountersAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (clearCountersAll *ClearCountersAll) GetEntityData() *types.CommonEntityData {
    clearCountersAll.EntityData.YFilter = clearCountersAll.YFilter
    clearCountersAll.EntityData.YangName = "clear-counters-all"
    clearCountersAll.EntityData.BundleName = "cisco_ios_xr"
    clearCountersAll.EntityData.ParentYangName = "Cisco-IOS-XR-infra-statsd-act"
    clearCountersAll.EntityData.SegmentPath = "Cisco-IOS-XR-infra-statsd-act:clear-counters-all"
    clearCountersAll.EntityData.AbsolutePath = clearCountersAll.EntityData.SegmentPath
    clearCountersAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearCountersAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearCountersAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearCountersAll.EntityData.Children = types.NewOrderedMap()
    clearCountersAll.EntityData.Leafs = types.NewOrderedMap()

    clearCountersAll.EntityData.YListKeys = []string {}

    return &(clearCountersAll.EntityData)
}

// ClearCountersInterface
// Clear counters for interface.
// 
type ClearCountersInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearCountersInterface_Input
}

func (clearCountersInterface *ClearCountersInterface) GetEntityData() *types.CommonEntityData {
    clearCountersInterface.EntityData.YFilter = clearCountersInterface.YFilter
    clearCountersInterface.EntityData.YangName = "clear-counters-interface"
    clearCountersInterface.EntityData.BundleName = "cisco_ios_xr"
    clearCountersInterface.EntityData.ParentYangName = "Cisco-IOS-XR-infra-statsd-act"
    clearCountersInterface.EntityData.SegmentPath = "Cisco-IOS-XR-infra-statsd-act:clear-counters-interface"
    clearCountersInterface.EntityData.AbsolutePath = clearCountersInterface.EntityData.SegmentPath
    clearCountersInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearCountersInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearCountersInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearCountersInterface.EntityData.Children = types.NewOrderedMap()
    clearCountersInterface.EntityData.Children.Append("input", types.YChild{"Input", &clearCountersInterface.Input})
    clearCountersInterface.EntityData.Leafs = types.NewOrderedMap()

    clearCountersInterface.EntityData.YListKeys = []string {}

    return &(clearCountersInterface.EntityData)
}

// ClearCountersInterface_Input
type ClearCountersInterface_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+. This
    // attribute is mandatory.
    InterfaceName interface{}
}

func (input *ClearCountersInterface_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-counters-interface"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-statsd-act:clear-counters-interface/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", input.InterfaceName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

