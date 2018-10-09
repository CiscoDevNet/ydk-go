// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-hwmod-profile package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: HardwareModule
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package prm_hwmod_profile_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package prm_hwmod_profile_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-hwmod-profile-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-hwmod-profile-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// ProfileTypeData represents Profile type data
type ProfileTypeData string

const (
    // SP Profile
    ProfileTypeData_sp ProfileTypeData = "sp"

    // DC Profile
    ProfileTypeData_dc ProfileTypeData = "dc"
)

// HardwareModule
// HardwareModule
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Profile type. The type is ProfileTypeData.
    Profile interface{}
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-prm-hwmod-profile-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-prm-hwmod-profile-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = types.NewOrderedMap()
    hardwareModule.EntityData.Leafs = types.NewOrderedMap()
    hardwareModule.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", hardwareModule.Profile})

    hardwareModule.EntityData.YListKeys = []string {}

    return &(hardwareModule.EntityData)
}

