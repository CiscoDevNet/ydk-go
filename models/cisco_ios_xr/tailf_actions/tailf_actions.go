// This module contains YANG definitions
// for Cisco IOS-XR and Tail-F action support.
// 
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package tailf_actions

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_actions"))
    ydk.RegisterEntity("{http://tail-f.com/ns/netconf/actions/1.0 action}", reflect.TypeOf(Action{}))
    ydk.RegisterEntity("tailf-actions:action", reflect.TypeOf(Action{}))
}

// Action
// Support Tail-F actions rpc format.
type Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Action_Input

    
    Output Action_Output
}

func (action *Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "tailf-actions"
    action.EntityData.SegmentPath = "tailf-actions:action"
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Children.Append("input", types.YChild{"Input", &action.Input})
    action.EntityData.Children.Append("output", types.YChild{"Output", &action.Output})
    action.EntityData.Leafs = types.NewOrderedMap()

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

// Action_Input
type Action_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data section of the YANG action hierarchy. The type is string.
    Data interface{}
}

func (input *Action_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "action"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("data", types.YLeaf{"Data", input.Data})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Action_Output
type Action_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data and messages returned by the Tail-F ConfD agent. The type is string.
    Data interface{}
}

func (output *Action_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "action"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("data", types.YLeaf{"Data", output.Data})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

