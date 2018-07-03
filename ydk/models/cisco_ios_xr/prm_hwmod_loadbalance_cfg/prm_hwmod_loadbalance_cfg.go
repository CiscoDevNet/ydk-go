// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-hwmod-loadbalance package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: HardwareModule
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package prm_hwmod_loadbalance_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package prm_hwmod_loadbalance_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-hwmod-loadbalance-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-hwmod-loadbalance-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// HardwareModule
// HardwareModule
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Loadbalance option.
    Loadbalancing HardwareModule_Loadbalancing
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-prm-hwmod-loadbalance-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-prm-hwmod-loadbalance-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = types.NewOrderedMap()
    hardwareModule.EntityData.Children.Append("loadbalancing", types.YChild{"Loadbalancing", &hardwareModule.Loadbalancing})
    hardwareModule.EntityData.Leafs = types.NewOrderedMap()

    hardwareModule.EntityData.YListKeys = []string {}

    return &(hardwareModule.EntityData)
}

// HardwareModule_Loadbalancing
// Loadbalance option
type HardwareModule_Loadbalancing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP LU.
    Bgp3107 HardwareModule_Loadbalancing_Bgp3107
}

func (loadbalancing *HardwareModule_Loadbalancing) GetEntityData() *types.CommonEntityData {
    loadbalancing.EntityData.YFilter = loadbalancing.YFilter
    loadbalancing.EntityData.YangName = "loadbalancing"
    loadbalancing.EntityData.BundleName = "cisco_ios_xr"
    loadbalancing.EntityData.ParentYangName = "hardware-module"
    loadbalancing.EntityData.SegmentPath = "loadbalancing"
    loadbalancing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadbalancing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadbalancing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadbalancing.EntityData.Children = types.NewOrderedMap()
    loadbalancing.EntityData.Children.Append("bgp3107", types.YChild{"Bgp3107", &loadbalancing.Bgp3107})
    loadbalancing.EntityData.Leafs = types.NewOrderedMap()

    loadbalancing.EntityData.YListKeys = []string {}

    return &(loadbalancing.EntityData)
}

// HardwareModule_Loadbalancing_Bgp3107
// BGP LU
type HardwareModule_Loadbalancing_Bgp3107 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ECMP .
    Ecmp HardwareModule_Loadbalancing_Bgp3107_Ecmp
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetEntityData() *types.CommonEntityData {
    bgp3107.EntityData.YFilter = bgp3107.YFilter
    bgp3107.EntityData.YangName = "bgp3107"
    bgp3107.EntityData.BundleName = "cisco_ios_xr"
    bgp3107.EntityData.ParentYangName = "loadbalancing"
    bgp3107.EntityData.SegmentPath = "bgp3107"
    bgp3107.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp3107.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp3107.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp3107.EntityData.Children = types.NewOrderedMap()
    bgp3107.EntityData.Children.Append("ecmp", types.YChild{"Ecmp", &bgp3107.Ecmp})
    bgp3107.EntityData.Leafs = types.NewOrderedMap()

    bgp3107.EntityData.YListKeys = []string {}

    return &(bgp3107.EntityData)
}

// HardwareModule_Loadbalancing_Bgp3107_Ecmp
// ECMP 
type HardwareModule_Loadbalancing_Bgp3107_Ecmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Option. The type is interface{}.
    Enable interface{}
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetEntityData() *types.CommonEntityData {
    ecmp.EntityData.YFilter = ecmp.YFilter
    ecmp.EntityData.YangName = "ecmp"
    ecmp.EntityData.BundleName = "cisco_ios_xr"
    ecmp.EntityData.ParentYangName = "bgp3107"
    ecmp.EntityData.SegmentPath = "ecmp"
    ecmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ecmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ecmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ecmp.EntityData.Children = types.NewOrderedMap()
    ecmp.EntityData.Leafs = types.NewOrderedMap()
    ecmp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ecmp.Enable})

    ecmp.EntityData.YListKeys = []string {}

    return &(ecmp.EntityData)
}

