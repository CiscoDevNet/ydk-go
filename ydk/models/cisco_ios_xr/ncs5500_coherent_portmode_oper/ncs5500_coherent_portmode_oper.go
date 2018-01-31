// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-portmode package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controller-port-mode: Coherent PortMode  operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5500_coherent_portmode_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5500_coherent_portmode_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs5500-coherent-portmode-oper controller-port-mode}", reflect.TypeOf(ControllerPortMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs5500-coherent-portmode-oper:controller-port-mode", reflect.TypeOf(ControllerPortMode{}))
}

// ControllerPortMode
// Coherent PortMode  operational data
type ControllerPortMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of optics controller. The type is slice of
    // ControllerPortMode_OpticsName.
    OpticsName []ControllerPortMode_OpticsName
}

func (controllerPortMode *ControllerPortMode) GetFilter() yfilter.YFilter { return controllerPortMode.YFilter }

func (controllerPortMode *ControllerPortMode) SetFilter(yf yfilter.YFilter) { controllerPortMode.YFilter = yf }

func (controllerPortMode *ControllerPortMode) GetGoName(yname string) string {
    if yname == "optics-name" { return "OpticsName" }
    return ""
}

func (controllerPortMode *ControllerPortMode) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs5500-coherent-portmode-oper:controller-port-mode"
}

func (controllerPortMode *ControllerPortMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "optics-name" {
        for _, c := range controllerPortMode.OpticsName {
            if controllerPortMode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ControllerPortMode_OpticsName{}
        controllerPortMode.OpticsName = append(controllerPortMode.OpticsName, child)
        return &controllerPortMode.OpticsName[len(controllerPortMode.OpticsName)-1]
    }
    return nil
}

func (controllerPortMode *ControllerPortMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range controllerPortMode.OpticsName {
        children[controllerPortMode.OpticsName[i].GetSegmentPath()] = &controllerPortMode.OpticsName[i]
    }
    return children
}

func (controllerPortMode *ControllerPortMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (controllerPortMode *ControllerPortMode) GetBundleName() string { return "cisco_ios_xr" }

func (controllerPortMode *ControllerPortMode) GetYangName() string { return "controller-port-mode" }

func (controllerPortMode *ControllerPortMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controllerPortMode *ControllerPortMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controllerPortMode *ControllerPortMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controllerPortMode *ControllerPortMode) SetParent(parent types.Entity) { controllerPortMode.parent = parent }

func (controllerPortMode *ControllerPortMode) GetParent() types.Entity { return controllerPortMode.parent }

func (controllerPortMode *ControllerPortMode) GetParentYangName() string { return "Cisco-IOS-XR-ncs5500-coherent-portmode-oper" }

// ControllerPortMode_OpticsName
// Name of optics controller
type ControllerPortMode_OpticsName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // PortMode  operational data.
    PortModeInfo ControllerPortMode_OpticsName_PortModeInfo
}

func (opticsName *ControllerPortMode_OpticsName) GetFilter() yfilter.YFilter { return opticsName.YFilter }

func (opticsName *ControllerPortMode_OpticsName) SetFilter(yf yfilter.YFilter) { opticsName.YFilter = yf }

func (opticsName *ControllerPortMode_OpticsName) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "port-mode-info" { return "PortModeInfo" }
    return ""
}

func (opticsName *ControllerPortMode_OpticsName) GetSegmentPath() string {
    return "optics-name" + "[interface-name='" + fmt.Sprintf("%v", opticsName.InterfaceName) + "']"
}

func (opticsName *ControllerPortMode_OpticsName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-mode-info" {
        return &opticsName.PortModeInfo
    }
    return nil
}

func (opticsName *ControllerPortMode_OpticsName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-mode-info"] = &opticsName.PortModeInfo
    return children
}

func (opticsName *ControllerPortMode_OpticsName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = opticsName.InterfaceName
    return leafs
}

func (opticsName *ControllerPortMode_OpticsName) GetBundleName() string { return "cisco_ios_xr" }

func (opticsName *ControllerPortMode_OpticsName) GetYangName() string { return "optics-name" }

func (opticsName *ControllerPortMode_OpticsName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticsName *ControllerPortMode_OpticsName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticsName *ControllerPortMode_OpticsName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticsName *ControllerPortMode_OpticsName) SetParent(parent types.Entity) { opticsName.parent = parent }

func (opticsName *ControllerPortMode_OpticsName) GetParent() types.Entity { return opticsName.parent }

func (opticsName *ControllerPortMode_OpticsName) GetParentYangName() string { return "controller-port-mode" }

// ControllerPortMode_OpticsName_PortModeInfo
// PortMode  operational data
type ControllerPortMode_OpticsName_PortModeInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // intf name. The type is string with length: 0..128.
    IntfName interface{}

    // speed. The type is string with length: 0..128.
    Speed interface{}

    // fec. The type is string with length: 0..128.
    Fec interface{}

    // diff. The type is string with length: 0..128.
    Diff interface{}

    // modulation. The type is string with length: 0..128.
    Modulation interface{}
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetFilter() yfilter.YFilter { return portModeInfo.YFilter }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) SetFilter(yf yfilter.YFilter) { portModeInfo.YFilter = yf }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetGoName(yname string) string {
    if yname == "intf-name" { return "IntfName" }
    if yname == "speed" { return "Speed" }
    if yname == "fec" { return "Fec" }
    if yname == "diff" { return "Diff" }
    if yname == "modulation" { return "Modulation" }
    return ""
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetSegmentPath() string {
    return "port-mode-info"
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intf-name"] = portModeInfo.IntfName
    leafs["speed"] = portModeInfo.Speed
    leafs["fec"] = portModeInfo.Fec
    leafs["diff"] = portModeInfo.Diff
    leafs["modulation"] = portModeInfo.Modulation
    return leafs
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetBundleName() string { return "cisco_ios_xr" }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetYangName() string { return "port-mode-info" }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) SetParent(parent types.Entity) { portModeInfo.parent = parent }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetParent() types.Entity { return portModeInfo.parent }

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetParentYangName() string { return "optics-name" }

