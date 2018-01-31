// This module contains a collection of YANG definitions
// for Cisco IOS-XR prm-hwmod-sr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: HardwareModule
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Segment Routing.
    SegmentRouting HardwareModule_SegmentRouting
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "segment-routing" { return "SegmentRouting" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-prm-hwmod-sr-cfg:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "segment-routing" {
        return &hardwareModule.SegmentRouting
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["segment-routing"] = &hardwareModule.SegmentRouting
    return children
}

func (hardwareModule *HardwareModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModule *HardwareModule) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModule *HardwareModule) GetYangName() string { return "hardware-module" }

func (hardwareModule *HardwareModule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModule *HardwareModule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModule *HardwareModule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModule *HardwareModule) SetParent(parent types.Entity) { hardwareModule.parent = parent }

func (hardwareModule *HardwareModule) GetParent() types.Entity { return hardwareModule.parent }

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-prm-hwmod-sr-cfg" }

// HardwareModule_SegmentRouting
// Segment Routing
type HardwareModule_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reserve.
    Reserve HardwareModule_SegmentRouting_Reserve
}

func (segmentRouting *HardwareModule_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *HardwareModule_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *HardwareModule_SegmentRouting) GetGoName(yname string) string {
    if yname == "reserve" { return "Reserve" }
    return ""
}

func (segmentRouting *HardwareModule_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *HardwareModule_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reserve" {
        return &segmentRouting.Reserve
    }
    return nil
}

func (segmentRouting *HardwareModule_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reserve"] = &segmentRouting.Reserve
    return children
}

func (segmentRouting *HardwareModule_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (segmentRouting *HardwareModule_SegmentRouting) GetBundleName() string { return "cisco_ios_xr" }

func (segmentRouting *HardwareModule_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *HardwareModule_SegmentRouting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (segmentRouting *HardwareModule_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (segmentRouting *HardwareModule_SegmentRouting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (segmentRouting *HardwareModule_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *HardwareModule_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *HardwareModule_SegmentRouting) GetParentYangName() string { return "hardware-module" }

// HardwareModule_SegmentRouting_Reserve
// Reserve
type HardwareModule_SegmentRouting_Reserve struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service Label.
    ServiceLabel HardwareModule_SegmentRouting_Reserve_ServiceLabel
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetFilter() yfilter.YFilter { return reserve.YFilter }

func (reserve *HardwareModule_SegmentRouting_Reserve) SetFilter(yf yfilter.YFilter) { reserve.YFilter = yf }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetGoName(yname string) string {
    if yname == "service-label" { return "ServiceLabel" }
    return ""
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetSegmentPath() string {
    return "reserve"
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-label" {
        return &reserve.ServiceLabel
    }
    return nil
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-label"] = &reserve.ServiceLabel
    return children
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reserve *HardwareModule_SegmentRouting_Reserve) GetBundleName() string { return "cisco_ios_xr" }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetYangName() string { return "reserve" }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reserve *HardwareModule_SegmentRouting_Reserve) SetParent(parent types.Entity) { reserve.parent = parent }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetParent() types.Entity { return reserve.parent }

func (reserve *HardwareModule_SegmentRouting_Reserve) GetParentYangName() string { return "segment-routing" }

// HardwareModule_SegmentRouting_Reserve_ServiceLabel
// Service Label
type HardwareModule_SegmentRouting_Reserve_ServiceLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable. The type is interface{}.
    Enable interface{}
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetFilter() yfilter.YFilter { return serviceLabel.YFilter }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) SetFilter(yf yfilter.YFilter) { serviceLabel.YFilter = yf }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetSegmentPath() string {
    return "service-label"
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = serviceLabel.Enable
    return leafs
}

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetBundleName() string { return "cisco_ios_xr" }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetYangName() string { return "service-label" }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) SetParent(parent types.Entity) { serviceLabel.parent = parent }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetParent() types.Entity { return serviceLabel.parent }

func (serviceLabel *HardwareModule_SegmentRouting_Reserve_ServiceLabel) GetParentYangName() string { return "reserve" }

