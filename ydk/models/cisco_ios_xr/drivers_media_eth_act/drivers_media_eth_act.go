// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package drivers_media_eth_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package drivers_media_eth_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-drivers-media-eth-act clear-controller-counters}", reflect.TypeOf(ClearControllerCounters{}))
    ydk.RegisterEntity("Cisco-IOS-XR-drivers-media-eth-act:clear-controller-counters", reflect.TypeOf(ClearControllerCounters{}))
}

// ClearControllerCounters
// Clear Ethernet MAC ASIC statistics.
type ClearControllerCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearControllerCounters_Input
}

func (clearControllerCounters *ClearControllerCounters) GetEntityData() *types.CommonEntityData {
    clearControllerCounters.EntityData.YFilter = clearControllerCounters.YFilter
    clearControllerCounters.EntityData.YangName = "clear-controller-counters"
    clearControllerCounters.EntityData.BundleName = "cisco_ios_xr"
    clearControllerCounters.EntityData.ParentYangName = "Cisco-IOS-XR-drivers-media-eth-act"
    clearControllerCounters.EntityData.SegmentPath = "Cisco-IOS-XR-drivers-media-eth-act:clear-controller-counters"
    clearControllerCounters.EntityData.AbsolutePath = clearControllerCounters.EntityData.SegmentPath
    clearControllerCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearControllerCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearControllerCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearControllerCounters.EntityData.Children = types.NewOrderedMap()
    clearControllerCounters.EntityData.Children.Append("input", types.YChild{"Input", &clearControllerCounters.Input})
    clearControllerCounters.EntityData.Leafs = types.NewOrderedMap()

    clearControllerCounters.EntityData.YListKeys = []string {}

    return &(clearControllerCounters.EntityData)
}

// ClearControllerCounters_Input
type ClearControllerCounters_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controller name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'. This
    // attribute is mandatory.
    ControllerName interface{}
}

func (input *ClearControllerCounters_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "clear-controller-counters"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-drivers-media-eth-act:clear-controller-counters/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", input.ControllerName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

