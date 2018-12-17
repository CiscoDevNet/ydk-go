// This module contains a collection of YANG definitions
// for Cisco IOS-XR Controller AINS configuration.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_ains_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_ains_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-controller-ains-act controller-ains-soak}", reflect.TypeOf(ControllerAinsSoak{}))
    ydk.RegisterEntity("Cisco-IOS-XR-controller-ains-act:controller-ains-soak", reflect.TypeOf(ControllerAinsSoak{}))
}

// ControllerAinsSoak
// Execute ains soak configuration operations
type ControllerAinsSoak struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ControllerAinsSoak_Input
}

func (controllerAinsSoak *ControllerAinsSoak) GetEntityData() *types.CommonEntityData {
    controllerAinsSoak.EntityData.YFilter = controllerAinsSoak.YFilter
    controllerAinsSoak.EntityData.YangName = "controller-ains-soak"
    controllerAinsSoak.EntityData.BundleName = "cisco_ios_xr"
    controllerAinsSoak.EntityData.ParentYangName = "Cisco-IOS-XR-controller-ains-act"
    controllerAinsSoak.EntityData.SegmentPath = "Cisco-IOS-XR-controller-ains-act:controller-ains-soak"
    controllerAinsSoak.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllerAinsSoak.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllerAinsSoak.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllerAinsSoak.EntityData.Children = types.NewOrderedMap()
    controllerAinsSoak.EntityData.Children.Append("input", types.YChild{"Input", &controllerAinsSoak.Input})
    controllerAinsSoak.EntityData.Leafs = types.NewOrderedMap()

    controllerAinsSoak.EntityData.YListKeys = []string {}

    return &(controllerAinsSoak.EntityData)
}

// ControllerAinsSoak_Input
type ControllerAinsSoak_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller name in R/S/I/P format. The type is string.
    Controller interface{}

    // Hours in range of 0-48. The type is interface{} with range: 0..48.
    Hours interface{}

    // Minutes in range of 0-59. The type is interface{} with range: 0..59.
    Minutes interface{}
}

func (input *ControllerAinsSoak_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "controller-ains-soak"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("controller", types.YLeaf{"Controller", input.Controller})
    input.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", input.Hours})
    input.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", input.Minutes})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

