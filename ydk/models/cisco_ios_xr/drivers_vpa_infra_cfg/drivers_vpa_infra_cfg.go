// This module contains a collection of YANG definitions
// for Cisco IOS-XR drivers-vpa-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module: Configure subslot h/w module
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// HwModuleSpaPhysicalMode represents Hw module spa physical mode
type HwModuleSpaPhysicalMode string

const (
    // CEM mode type
    HwModuleSpaPhysicalMode_cem HwModuleSpaPhysicalMode = "cem"
)

// HwModuleShutdownPowerMode represents Hw module shutdown power mode
type HwModuleShutdownPowerMode string

const (
    // Keep the showdown module unpowered
    HwModuleShutdownPowerMode_unpowered HwModuleShutdownPowerMode = "unpowered"

    // Keep the showdown module powered (default)
    HwModuleShutdownPowerMode_powered HwModuleShutdownPowerMode = "powered"
)

// HwModuleSpaPhysicalInterface represents Hw module spa physical interface
type HwModuleSpaPhysicalInterface string

const (
    // T3 interface type(default)
    HwModuleSpaPhysicalInterface_t3 HwModuleSpaPhysicalInterface = "t3"

    // E3 interface type
    HwModuleSpaPhysicalInterface_e3 HwModuleSpaPhysicalInterface = "e3"

    // T1 interface type
    HwModuleSpaPhysicalInterface_t1 HwModuleSpaPhysicalInterface = "t1"

    // E1 interface type
    HwModuleSpaPhysicalInterface_e1 HwModuleSpaPhysicalInterface = "e1"

    // Sonet interface type
    HwModuleSpaPhysicalInterface_sonet HwModuleSpaPhysicalInterface = "sonet"

    // SDH interface type
    HwModuleSpaPhysicalInterface_sdh HwModuleSpaPhysicalInterface = "sdh"
)

// HardwareModule
// Configure subslot h/w module
type HardwareModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging configuration.
    Logging HardwareModule_Logging

    // subslot h/w module.
    Nodes HardwareModule_Nodes
}

func (hardwareModule *HardwareModule) GetEntityData() *types.CommonEntityData {
    hardwareModule.EntityData.YFilter = hardwareModule.YFilter
    hardwareModule.EntityData.YangName = "hardware-module"
    hardwareModule.EntityData.BundleName = "cisco_ios_xr"
    hardwareModule.EntityData.ParentYangName = "Cisco-IOS-XR-drivers-vpa-infra-cfg"
    hardwareModule.EntityData.SegmentPath = "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module"
    hardwareModule.EntityData.AbsolutePath = hardwareModule.EntityData.SegmentPath
    hardwareModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModule.EntityData.Children = types.NewOrderedMap()
    hardwareModule.EntityData.Children.Append("logging", types.YChild{"Logging", &hardwareModule.Logging})
    hardwareModule.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModule.Nodes})
    hardwareModule.EntityData.Leafs = types.NewOrderedMap()

    hardwareModule.EntityData.YListKeys = []string {}

    return &(hardwareModule.EntityData)
}

// HardwareModule_Logging
// Logging configuration
type HardwareModule_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (logging *HardwareModule_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "hardware-module"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// HardwareModule_Nodes
//  subslot h/w module
type HardwareModule_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The identifier for a SPA node. The type is slice of
    // HardwareModule_Nodes_Node.
    Node []*HardwareModule_Nodes_Node
}

func (nodes *HardwareModule_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// HardwareModule_Nodes_Node
// The identifier for a SPA node
type HardwareModule_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A SPA node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Configure the SPA physical interface type. The type is
    // HwModuleSpaPhysicalInterface.
    CardType interface{}

    // Configure the SPA mode. The type is HwModuleSpaPhysicalMode.
    Mode interface{}

    // Shutdown a subslot h/w module. The type is HwModuleShutdownPowerMode.
    Shutdown interface{}
}

func (node *HardwareModule_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", node.CardType})
    node.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", node.Mode})
    node.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", node.Shutdown})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

