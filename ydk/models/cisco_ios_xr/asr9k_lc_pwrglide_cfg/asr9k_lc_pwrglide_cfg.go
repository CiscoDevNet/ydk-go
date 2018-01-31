// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-pwrglide package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-port-mode: HW module port-mode config
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lc_pwrglide_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lc_pwrglide_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lc-pwrglide-cfg hardware-module-port-mode}", reflect.TypeOf(HardwareModulePortMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode", reflect.TypeOf(HardwareModulePortMode{}))
}

// HardwareModulePortMode
// HW module port-mode config
type HardwareModulePortMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Active or Pre configuration. The type is slice of
    // HardwareModulePortMode_ConfigMode.
    ConfigMode []HardwareModulePortMode_ConfigMode
}

func (hardwareModulePortMode *HardwareModulePortMode) GetFilter() yfilter.YFilter { return hardwareModulePortMode.YFilter }

func (hardwareModulePortMode *HardwareModulePortMode) SetFilter(yf yfilter.YFilter) { hardwareModulePortMode.YFilter = yf }

func (hardwareModulePortMode *HardwareModulePortMode) GetGoName(yname string) string {
    if yname == "config-mode" { return "ConfigMode" }
    return ""
}

func (hardwareModulePortMode *HardwareModulePortMode) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode"
}

func (hardwareModulePortMode *HardwareModulePortMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-mode" {
        for _, c := range hardwareModulePortMode.ConfigMode {
            if hardwareModulePortMode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModulePortMode_ConfigMode{}
        hardwareModulePortMode.ConfigMode = append(hardwareModulePortMode.ConfigMode, child)
        return &hardwareModulePortMode.ConfigMode[len(hardwareModulePortMode.ConfigMode)-1]
    }
    return nil
}

func (hardwareModulePortMode *HardwareModulePortMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareModulePortMode.ConfigMode {
        children[hardwareModulePortMode.ConfigMode[i].GetSegmentPath()] = &hardwareModulePortMode.ConfigMode[i]
    }
    return children
}

func (hardwareModulePortMode *HardwareModulePortMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModulePortMode *HardwareModulePortMode) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModulePortMode *HardwareModulePortMode) GetYangName() string { return "hardware-module-port-mode" }

func (hardwareModulePortMode *HardwareModulePortMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModulePortMode *HardwareModulePortMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModulePortMode *HardwareModulePortMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModulePortMode *HardwareModulePortMode) SetParent(parent types.Entity) { hardwareModulePortMode.parent = parent }

func (hardwareModulePortMode *HardwareModulePortMode) GetParent() types.Entity { return hardwareModulePortMode.parent }

func (hardwareModulePortMode *HardwareModulePortMode) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg" }

// HardwareModulePortMode_ConfigMode
// Active or Pre configuration
type HardwareModulePortMode_ConfigMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. act- or pre-config. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // A node. The type is slice of HardwareModulePortMode_ConfigMode_Node.
    Node []HardwareModulePortMode_ConfigMode_Node
}

func (configMode *HardwareModulePortMode_ConfigMode) GetFilter() yfilter.YFilter { return configMode.YFilter }

func (configMode *HardwareModulePortMode_ConfigMode) SetFilter(yf yfilter.YFilter) { configMode.YFilter = yf }

func (configMode *HardwareModulePortMode_ConfigMode) GetGoName(yname string) string {
    if yname == "id1" { return "Id1" }
    if yname == "node" { return "Node" }
    return ""
}

func (configMode *HardwareModulePortMode_ConfigMode) GetSegmentPath() string {
    return "config-mode" + "[id1='" + fmt.Sprintf("%v", configMode.Id1) + "']"
}

func (configMode *HardwareModulePortMode_ConfigMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range configMode.Node {
            if configMode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModulePortMode_ConfigMode_Node{}
        configMode.Node = append(configMode.Node, child)
        return &configMode.Node[len(configMode.Node)-1]
    }
    return nil
}

func (configMode *HardwareModulePortMode_ConfigMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configMode.Node {
        children[configMode.Node[i].GetSegmentPath()] = &configMode.Node[i]
    }
    return children
}

func (configMode *HardwareModulePortMode_ConfigMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id1"] = configMode.Id1
    return leafs
}

func (configMode *HardwareModulePortMode_ConfigMode) GetBundleName() string { return "cisco_ios_xr" }

func (configMode *HardwareModulePortMode_ConfigMode) GetYangName() string { return "config-mode" }

func (configMode *HardwareModulePortMode_ConfigMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMode *HardwareModulePortMode_ConfigMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMode *HardwareModulePortMode_ConfigMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMode *HardwareModulePortMode_ConfigMode) SetParent(parent types.Entity) { configMode.parent = parent }

func (configMode *HardwareModulePortMode_ConfigMode) GetParent() types.Entity { return configMode.parent }

func (configMode *HardwareModulePortMode_ConfigMode) GetParentYangName() string { return "hardware-module-port-mode" }

// HardwareModulePortMode_ConfigMode_Node
// A node
type HardwareModulePortMode_ConfigMode_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Id2 interface{}

    // Linecard port-mode.
    PortMode HardwareModulePortMode_ConfigMode_Node_PortMode
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModulePortMode_ConfigMode_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModulePortMode_ConfigMode_Node) GetGoName(yname string) string {
    if yname == "id2" { return "Id2" }
    if yname == "port-mode" { return "PortMode" }
    return ""
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetSegmentPath() string {
    return "node" + "[id2='" + fmt.Sprintf("%v", node.Id2) + "']"
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-mode" {
        return &node.PortMode
    }
    return nil
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-mode"] = &node.PortMode
    return children
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id2"] = node.Id2
    return leafs
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModulePortMode_ConfigMode_Node) GetYangName() string { return "node" }

func (node *HardwareModulePortMode_ConfigMode_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModulePortMode_ConfigMode_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModulePortMode_ConfigMode_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModulePortMode_ConfigMode_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModulePortMode_ConfigMode_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModulePortMode_ConfigMode_Node) GetParentYangName() string { return "config-mode" }

// HardwareModulePortMode_ConfigMode_Node_PortMode
// Linecard port-mode
type HardwareModulePortMode_ConfigMode_Node_PortMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Linecard interface port-mode. The type is string.
    IfPortMode interface{}
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetFilter() yfilter.YFilter { return portMode.YFilter }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) SetFilter(yf yfilter.YFilter) { portMode.YFilter = yf }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetGoName(yname string) string {
    if yname == "if-port-mode" { return "IfPortMode" }
    return ""
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetSegmentPath() string {
    return "port-mode"
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["if-port-mode"] = portMode.IfPortMode
    return leafs
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetBundleName() string { return "cisco_ios_xr" }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetYangName() string { return "port-mode" }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) SetParent(parent types.Entity) { portMode.parent = parent }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetParent() types.Entity { return portMode.parent }

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetParentYangName() string { return "node" }

