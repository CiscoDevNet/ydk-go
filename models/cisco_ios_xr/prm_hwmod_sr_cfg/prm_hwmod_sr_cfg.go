// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-hwmod-sr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: HardwareModule
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package prm_hwmod_sr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package prm_hwmod_sr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-prm-hwmod-sr-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-prm-hwmod-sr-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// HardwareModule
// HardwareModule
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing.
    SegmentRouting HardwareModule_SegmentRouting
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-prm-hwmod-sr-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-prm-hwmod-sr-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = types.NewOrderedMap()
    hardwareModule.EntityData.Children.Append("segment-routing", types.YChild{"SegmentRouting", &hardwareModule.SegmentRouting})
    hardwareModule.EntityData.Leafs = types.NewOrderedMap()

    hardwareModule.EntityData.YListKeys = []string {}

    return &(hardwareModule.EntityData)
}

// HardwareModule_SegmentRouting
// Segment Routing
type HardwareModule_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reserve.
    Reserve HardwareModule_SegmentRouting_Reserve
}

func (segmentRouting *HardwareModule_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "hardware-module"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = types.NewOrderedMap()
    segmentRouting.EntityData.Children.Append("reserve", types.YChild{"Reserve", &segmentRouting.Reserve})
    segmentRouting.EntityData.Leafs = types.NewOrderedMap()

    segmentRouting.EntityData.YListKeys = []string {}

    return &(segmentRouting.EntityData)
}

// HardwareModule_SegmentRouting_Reserve
// Reserve
type HardwareModule_SegmentRouting_Reserve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service Label.
    ServiceLabel HardwareModule_SegmentRouting_Reserve_ServiceLabel
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetEntityData() *types.CommonEntityData {
    reserve.EntityData.YFilter = reserve.YFilter
    reserve.EntityData.YangName = "reserve"
    reserve.EntityData.BundleName = "cisco_ios_xr"
    reserve.EntityData.ParentYangName = "segment-routing"
    reserve.EntityData.SegmentPath = "reserve"
    reserve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reserve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reserve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reserve.EntityData.Children = types.NewOrderedMap()
    reserve.EntityData.Children.Append("service-label", types.YChild{"ServiceLabel", &reserve.ServiceLabel})
    reserve.EntityData.Leafs = types.NewOrderedMap()

    reserve.EntityData.YListKeys = []string {}

    return &(reserve.EntityData)
}

// HardwareModule_SegmentRouting_Reserve_ServiceLabel
// Service Label
type HardwareModule_SegmentRouting_Reserve_ServiceLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable. The type is interface{}.
    Enable interface{}
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetEntityData() *types.CommonEntityData {
    serviceLabel.EntityData.YFilter = serviceLabel.YFilter
    serviceLabel.EntityData.YangName = "service-label"
    serviceLabel.EntityData.BundleName = "cisco_ios_xr"
    serviceLabel.EntityData.ParentYangName = "reserve"
    serviceLabel.EntityData.SegmentPath = "service-label"
    serviceLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceLabel.EntityData.Children = types.NewOrderedMap()
    serviceLabel.EntityData.Leafs = types.NewOrderedMap()
    serviceLabel.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", serviceLabel.Enable})

    serviceLabel.EntityData.YListKeys = []string {}

    return &(serviceLabel.EntityData)
}

