// This module contains a collection of YANG definitions
// for Cisco IOS-XR Controller to clear alarm counters.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package clear_counters_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package clear_counters_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-clear-counters-act clear-counters}", reflect.TypeOf(ClearCounters{}))
    ydk.RegisterEntity("Cisco-IOS-XR-clear-counters-act:clear-counters", reflect.TypeOf(ClearCounters{}))
}

// ClearCounters
// Execute clear counters operations
type ClearCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearCounters_Input
}

func (clearCounters *ClearCounters) GetEntityData() *types.CommonEntityData {
    clearCounters.EntityData.YFilter = clearCounters.YFilter
    clearCounters.EntityData.YangName = "clear-counters"
    clearCounters.EntityData.BundleName = "cisco_ios_xr"
    clearCounters.EntityData.ParentYangName = "Cisco-IOS-XR-clear-counters-act"
    clearCounters.EntityData.SegmentPath = "Cisco-IOS-XR-clear-counters-act:clear-counters"
    clearCounters.EntityData.AbsolutePath = clearCounters.EntityData.SegmentPath
    clearCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearCounters.EntityData.Children = types.NewOrderedMap()
    clearCounters.EntityData.Children.Append("input", types.YChild{"Input", &clearCounters.Input})
    clearCounters.EntityData.Leafs = types.NewOrderedMap()

    clearCounters.EntityData.YListKeys = []string {}

    return &(clearCounters.EntityData)
}

// ClearCounters_Input
type ClearCounters_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller name in R/S/I/P format. The type is string.
    Controller interface{}
}

func (input *ClearCounters_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-counters"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-clear-counters-act:clear-counters/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("controller", types.YLeaf{"Controller", input.Controller})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

