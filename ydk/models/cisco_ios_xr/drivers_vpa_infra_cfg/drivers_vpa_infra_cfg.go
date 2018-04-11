// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-vpa-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: Configure subslot h/w module
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package drivers_vpa_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package drivers_vpa_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-drivers-vpa-infra-cfg hardware-module}", reflect.TypeOf(HardwareModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module", reflect.TypeOf(HardwareModule{}))
}

// HwModuleShutdownPowerMode represents Hw module shutdown power mode
type HwModuleShutdownPowerMode string

const (
    // Keep the showdown module powered (default)
    HwModuleShutdownPowerMode_powered HwModuleShutdownPowerMode = "powered"
)

// HardwareModule
// Configure subslot h/w module
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // subslot h/w module.
    Nodes HardwareModule_Nodes
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-drivers-vpa-infra-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module"
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = make(map[string]types.YChild)
    hardwareModule.EntityData.Children["nodes"] = types.YChild{"Nodes", &hardwareModule.Nodes}
    hardwareModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareModule.EntityData)
}

// HardwareModule_Nodes
//  subslot h/w module
type HardwareModule_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The identifier for a SPA node. The type is slice of
    // HardwareModule_Nodes_Node.
    Node []HardwareModule_Nodes_Node
}

func (nodes *HardwareModule_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// HardwareModule_Nodes_Node
// The identifier for a SPA node
type HardwareModule_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A SPA node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Shutdown a subslot h/w module. The type is HwModuleShutdownPowerMode.
    Shutdown interface{}
}

func (node *HardwareModule_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["shutdown"] = types.YLeaf{"Shutdown", node.Shutdown}
    return &(node.EntityData)
}

