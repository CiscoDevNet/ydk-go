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
    parent types.Entity
    YFilter yfilter.YFilter

    // Loadbalance option.
    Loadbalancing HardwareModule_Loadbalancing
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "loadbalancing" { return "Loadbalancing" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-prm-hwmod-loadbalance-cfg:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "loadbalancing" {
        return &hardwareModule.Loadbalancing
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["loadbalancing"] = &hardwareModule.Loadbalancing
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

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-prm-hwmod-loadbalance-cfg" }

// HardwareModule_Loadbalancing
// Loadbalance option
type HardwareModule_Loadbalancing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP LU.
    Bgp3107 HardwareModule_Loadbalancing_Bgp3107
}

func (loadbalancing *HardwareModule_Loadbalancing) GetFilter() yfilter.YFilter { return loadbalancing.YFilter }

func (loadbalancing *HardwareModule_Loadbalancing) SetFilter(yf yfilter.YFilter) { loadbalancing.YFilter = yf }

func (loadbalancing *HardwareModule_Loadbalancing) GetGoName(yname string) string {
    if yname == "bgp3107" { return "Bgp3107" }
    return ""
}

func (loadbalancing *HardwareModule_Loadbalancing) GetSegmentPath() string {
    return "loadbalancing"
}

func (loadbalancing *HardwareModule_Loadbalancing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp3107" {
        return &loadbalancing.Bgp3107
    }
    return nil
}

func (loadbalancing *HardwareModule_Loadbalancing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp3107"] = &loadbalancing.Bgp3107
    return children
}

func (loadbalancing *HardwareModule_Loadbalancing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadbalancing *HardwareModule_Loadbalancing) GetBundleName() string { return "cisco_ios_xr" }

func (loadbalancing *HardwareModule_Loadbalancing) GetYangName() string { return "loadbalancing" }

func (loadbalancing *HardwareModule_Loadbalancing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (loadbalancing *HardwareModule_Loadbalancing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (loadbalancing *HardwareModule_Loadbalancing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (loadbalancing *HardwareModule_Loadbalancing) SetParent(parent types.Entity) { loadbalancing.parent = parent }

func (loadbalancing *HardwareModule_Loadbalancing) GetParent() types.Entity { return loadbalancing.parent }

func (loadbalancing *HardwareModule_Loadbalancing) GetParentYangName() string { return "hardware-module" }

// HardwareModule_Loadbalancing_Bgp3107
// BGP LU
type HardwareModule_Loadbalancing_Bgp3107 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ECMP .
    Ecmp HardwareModule_Loadbalancing_Bgp3107_Ecmp
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetFilter() yfilter.YFilter { return bgp3107.YFilter }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) SetFilter(yf yfilter.YFilter) { bgp3107.YFilter = yf }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetGoName(yname string) string {
    if yname == "ecmp" { return "Ecmp" }
    return ""
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetSegmentPath() string {
    return "bgp3107"
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ecmp" {
        return &bgp3107.Ecmp
    }
    return nil
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ecmp"] = &bgp3107.Ecmp
    return children
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetBundleName() string { return "cisco_ios_xr" }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetYangName() string { return "bgp3107" }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) SetParent(parent types.Entity) { bgp3107.parent = parent }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetParent() types.Entity { return bgp3107.parent }

func (bgp3107 *HardwareModule_Loadbalancing_Bgp3107) GetParentYangName() string { return "loadbalancing" }

// HardwareModule_Loadbalancing_Bgp3107_Ecmp
// ECMP 
type HardwareModule_Loadbalancing_Bgp3107_Ecmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Option. The type is interface{}.
    Enable interface{}
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetFilter() yfilter.YFilter { return ecmp.YFilter }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) SetFilter(yf yfilter.YFilter) { ecmp.YFilter = yf }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetSegmentPath() string {
    return "ecmp"
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ecmp.Enable
    return leafs
}

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetBundleName() string { return "cisco_ios_xr" }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetYangName() string { return "ecmp" }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) SetParent(parent types.Entity) { ecmp.parent = parent }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetParent() types.Entity { return ecmp.parent }

func (ecmp *HardwareModule_Loadbalancing_Bgp3107_Ecmp) GetParentYangName() string { return "bgp3107" }

