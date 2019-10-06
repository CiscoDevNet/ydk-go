// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-pwrglide package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-port-mode: HW module port-mode config
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active or Pre configuration. The type is slice of
    // HardwareModulePortMode_ConfigMode.
    ConfigMode []*HardwareModulePortMode_ConfigMode
}

func (hardwareModulePortMode *HardwareModulePortMode) GetEntityData() *types.CommonEntityData {
    hardwareModulePortMode.EntityData.YFilter = hardwareModulePortMode.YFilter
    hardwareModulePortMode.EntityData.YangName = "hardware-module-port-mode"
    hardwareModulePortMode.EntityData.BundleName = "cisco_ios_xr"
    hardwareModulePortMode.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg"
    hardwareModulePortMode.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode"
    hardwareModulePortMode.EntityData.AbsolutePath = hardwareModulePortMode.EntityData.SegmentPath
    hardwareModulePortMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModulePortMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModulePortMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModulePortMode.EntityData.Children = types.NewOrderedMap()
    hardwareModulePortMode.EntityData.Children.Append("config-mode", types.YChild{"ConfigMode", nil})
    for i := range hardwareModulePortMode.ConfigMode {
        hardwareModulePortMode.EntityData.Children.Append(types.GetSegmentPath(hardwareModulePortMode.ConfigMode[i]), types.YChild{"ConfigMode", hardwareModulePortMode.ConfigMode[i]})
    }
    hardwareModulePortMode.EntityData.Leafs = types.NewOrderedMap()

    hardwareModulePortMode.EntityData.YListKeys = []string {}

    return &(hardwareModulePortMode.EntityData)
}

// HardwareModulePortMode_ConfigMode
// Active or Pre configuration
type HardwareModulePortMode_ConfigMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. act- or pre-config. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // A node. The type is slice of HardwareModulePortMode_ConfigMode_Node.
    Node []*HardwareModulePortMode_ConfigMode_Node
}

func (configMode *HardwareModulePortMode_ConfigMode) GetEntityData() *types.CommonEntityData {
    configMode.EntityData.YFilter = configMode.YFilter
    configMode.EntityData.YangName = "config-mode"
    configMode.EntityData.BundleName = "cisco_ios_xr"
    configMode.EntityData.ParentYangName = "hardware-module-port-mode"
    configMode.EntityData.SegmentPath = "config-mode" + types.AddKeyToken(configMode.Id1, "id1")
    configMode.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode/" + configMode.EntityData.SegmentPath
    configMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMode.EntityData.Children = types.NewOrderedMap()
    configMode.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range configMode.Node {
        configMode.EntityData.Children.Append(types.GetSegmentPath(configMode.Node[i]), types.YChild{"Node", configMode.Node[i]})
    }
    configMode.EntityData.Leafs = types.NewOrderedMap()
    configMode.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", configMode.Id1})

    configMode.EntityData.YListKeys = []string {"Id1"}

    return &(configMode.EntityData)
}

// HardwareModulePortMode_ConfigMode_Node
// A node
type HardwareModulePortMode_ConfigMode_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fully qualified line card specification. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Id2 interface{}

    // Linecard port-mode.
    PortMode HardwareModulePortMode_ConfigMode_Node_PortMode
}

func (node *HardwareModulePortMode_ConfigMode_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "config-mode"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Id2, "id2")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode/config-mode/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("port-mode", types.YChild{"PortMode", &node.PortMode})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("id2", types.YLeaf{"Id2", node.Id2})

    node.EntityData.YListKeys = []string {"Id2"}

    return &(node.EntityData)
}

// HardwareModulePortMode_ConfigMode_Node_PortMode
// Linecard port-mode
type HardwareModulePortMode_ConfigMode_Node_PortMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Linecard interface port-mode. The type is string.
    IfPortMode interface{}
}

func (portMode *HardwareModulePortMode_ConfigMode_Node_PortMode) GetEntityData() *types.CommonEntityData {
    portMode.EntityData.YFilter = portMode.YFilter
    portMode.EntityData.YangName = "port-mode"
    portMode.EntityData.BundleName = "cisco_ios_xr"
    portMode.EntityData.ParentYangName = "node"
    portMode.EntityData.SegmentPath = "port-mode"
    portMode.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-pwrglide-cfg:hardware-module-port-mode/config-mode/node/" + portMode.EntityData.SegmentPath
    portMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portMode.EntityData.Children = types.NewOrderedMap()
    portMode.EntityData.Leafs = types.NewOrderedMap()
    portMode.EntityData.Leafs.Append("if-port-mode", types.YLeaf{"IfPortMode", portMode.IfPortMode})

    portMode.EntityData.YListKeys = []string {}

    return &(portMode.EntityData)
}

