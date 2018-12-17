// This module contains a collection of YANG definitions
// for Cisco IOS-XR MPA reload action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package hwmod_mpa_reload_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package hwmod_mpa_reload_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-hwmod-mpa-reload-act hw-module-subslot}", reflect.TypeOf(HwModuleSubslot{}))
    ydk.RegisterEntity("Cisco-IOS-XR-hwmod-mpa-reload-act:hw-module-subslot", reflect.TypeOf(HwModuleSubslot{}))
}

// HwModuleSubslot
// Execute subslot h/w module operations
type HwModuleSubslot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input HwModuleSubslot_Input
}

func (hwModuleSubslot *HwModuleSubslot) GetEntityData() *types.CommonEntityData {
    hwModuleSubslot.EntityData.YFilter = hwModuleSubslot.YFilter
    hwModuleSubslot.EntityData.YangName = "hw-module-subslot"
    hwModuleSubslot.EntityData.BundleName = "cisco_ios_xr"
    hwModuleSubslot.EntityData.ParentYangName = "Cisco-IOS-XR-hwmod-mpa-reload-act"
    hwModuleSubslot.EntityData.SegmentPath = "Cisco-IOS-XR-hwmod-mpa-reload-act:hw-module-subslot"
    hwModuleSubslot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleSubslot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleSubslot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleSubslot.EntityData.Children = types.NewOrderedMap()
    hwModuleSubslot.EntityData.Children.Append("input", types.YChild{"Input", &hwModuleSubslot.Input})
    hwModuleSubslot.EntityData.Leafs = types.NewOrderedMap()

    hwModuleSubslot.EntityData.YListKeys = []string {}

    return &(hwModuleSubslot.EntityData)
}

// HwModuleSubslot_Input
type HwModuleSubslot_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fully qualified location specification. The type is string.
    Subslot interface{}

    // Cycle subslot h/w reset. The type is interface{}.
    Reload interface{}
}

func (input *HwModuleSubslot_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "hw-module-subslot"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("subslot", types.YLeaf{"Subslot", input.Subslot})
    input.EntityData.Leafs.Append("reload", types.YLeaf{"Reload", input.Reload})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

