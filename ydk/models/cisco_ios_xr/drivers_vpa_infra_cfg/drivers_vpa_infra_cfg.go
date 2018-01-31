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
    parent types.Entity
    YFilter yfilter.YFilter

    // subslot h/w module.
    Nodes HardwareModule_Nodes
}

func (hardwareModule *HardwareModule) GetFilter() yfilter.YFilter { return hardwareModule.YFilter }

func (hardwareModule *HardwareModule) SetFilter(yf yfilter.YFilter) { hardwareModule.YFilter = yf }

func (hardwareModule *HardwareModule) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModule *HardwareModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-drivers-vpa-infra-cfg:hardware-module"
}

func (hardwareModule *HardwareModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModule.Nodes
    }
    return nil
}

func (hardwareModule *HardwareModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModule.Nodes
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

func (hardwareModule *HardwareModule) GetParentYangName() string { return "Cisco-IOS-XR-drivers-vpa-infra-cfg" }

// HardwareModule_Nodes
//  subslot h/w module
type HardwareModule_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The identifier for a SPA node. The type is slice of
    // HardwareModule_Nodes_Node.
    Node []HardwareModule_Nodes_Node
}

func (nodes *HardwareModule_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModule_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModule_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModule_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModule_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModule_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModule_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModule_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModule_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModule_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModule_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModule_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModule_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModule_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModule_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModule_Nodes) GetParentYangName() string { return "hardware-module" }

// HardwareModule_Nodes_Node
// The identifier for a SPA node
type HardwareModule_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A SPA node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Shutdown a subslot h/w module. The type is HwModuleShutdownPowerMode.
    Shutdown interface{}
}

func (node *HardwareModule_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModule_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModule_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "shutdown" { return "Shutdown" }
    return ""
}

func (node *HardwareModule_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModule_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *HardwareModule_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *HardwareModule_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["shutdown"] = node.Shutdown
    return leafs
}

func (node *HardwareModule_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModule_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModule_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModule_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModule_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModule_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModule_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModule_Nodes_Node) GetParentYangName() string { return "nodes" }

