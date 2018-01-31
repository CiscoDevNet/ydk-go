// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-prm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-qos-mode: QoS mode in hardware module port(s)
//   hardware-module-tcp-mss-adjust: hardware module tcp mss adjust
//   hardware-module-load-balance: hardware module load balance
//   hardware-module-tcam: hardware module tcam
//   hardware-module-efd: hardware module efd
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_prm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_prm_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-qos-mode}", reflect.TypeOf(HardwareModuleQosMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode", reflect.TypeOf(HardwareModuleQosMode{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcp-mss-adjust}", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-load-balance}", reflect.TypeOf(HardwareModuleLoadBalance{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-load-balance", reflect.TypeOf(HardwareModuleLoadBalance{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcam}", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-efd}", reflect.TypeOf(HardwareModuleEfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd", reflect.TypeOf(HardwareModuleEfd{}))
}

// Asr9kEfdMode represents Asr9k efd mode
type Asr9kEfdMode string

const (
    // Only check outer encap
    Asr9kEfdMode_only_outer_encap Asr9kEfdMode = "only-outer-encap"

    // Check outer and inner encap
    Asr9kEfdMode_include_inner_encap Asr9kEfdMode = "include-inner-encap"
)

// Asr9kEfdOperation represents Asr9k efd operation
type Asr9kEfdOperation string

const (
    // Less than
    Asr9kEfdOperation_less_than Asr9kEfdOperation = "less-than"

    // Greater than or equal
    Asr9kEfdOperation_greater_than_or_equal Asr9kEfdOperation = "greater-than-or-equal"
)

// PrmTcamProfile represents Prm tcam profile
type PrmTcamProfile string

const (
    // Profile 0
    PrmTcamProfile_profile0 PrmTcamProfile = "profile0"

    // Profile 1
    PrmTcamProfile_profile1 PrmTcamProfile = "profile1"

    // Profile 2
    PrmTcamProfile_profile2 PrmTcamProfile = "profile2"
)

// HardwareModuleQosMode
// QoS mode in hardware module port(s)
type HardwareModuleQosMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS applicable nodes.
    Nodes HardwareModuleQosMode_Nodes
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetFilter() yfilter.YFilter { return hardwareModuleQosMode.YFilter }

func (hardwareModuleQosMode *HardwareModuleQosMode) SetFilter(yf yfilter.YFilter) { hardwareModuleQosMode.YFilter = yf }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode"
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModuleQosMode.Nodes
    }
    return nil
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModuleQosMode.Nodes
    return children
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetYangName() string { return "hardware-module-qos-mode" }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleQosMode *HardwareModuleQosMode) SetParent(parent types.Entity) { hardwareModuleQosMode.parent = parent }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetParent() types.Entity { return hardwareModuleQosMode.parent }

func (hardwareModuleQosMode *HardwareModuleQosMode) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-prm-cfg" }

// HardwareModuleQosMode_Nodes
// QoS applicable nodes
type HardwareModuleQosMode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleQosMode_Nodes_Node.
    Node []HardwareModuleQosMode_Nodes_Node
}

func (nodes *HardwareModuleQosMode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModuleQosMode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModuleQosMode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModuleQosMode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModuleQosMode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleQosMode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModuleQosMode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModuleQosMode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModuleQosMode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModuleQosMode_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModuleQosMode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModuleQosMode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModuleQosMode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModuleQosMode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModuleQosMode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModuleQosMode_Nodes) GetParentYangName() string { return "hardware-module-qos-mode" }

// HardwareModuleQosMode_Nodes_Node
// A node
type HardwareModuleQosMode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Disable child level/flat policy shaping. The type is interface{}.
    ChildShapingDisable interface{}

    // Enable low burst mode for TM entity. The type is interface{}.
    LowburstEnable interface{}
}

func (node *HardwareModuleQosMode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModuleQosMode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModuleQosMode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "child-shaping-disable" { return "ChildShapingDisable" }
    if yname == "lowburst-enable" { return "LowburstEnable" }
    return ""
}

func (node *HardwareModuleQosMode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModuleQosMode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *HardwareModuleQosMode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *HardwareModuleQosMode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["child-shaping-disable"] = node.ChildShapingDisable
    leafs["lowburst-enable"] = node.LowburstEnable
    return leafs
}

func (node *HardwareModuleQosMode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModuleQosMode_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModuleQosMode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModuleQosMode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModuleQosMode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModuleQosMode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModuleQosMode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModuleQosMode_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModuleTcpMssAdjust
// hardware module tcp mss adjust
type HardwareModuleTcpMssAdjust struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP MSS Adjust applicable nodes.
    Nodes HardwareModuleTcpMssAdjust_Nodes
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetFilter() yfilter.YFilter { return hardwareModuleTcpMssAdjust.YFilter }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) SetFilter(yf yfilter.YFilter) { hardwareModuleTcpMssAdjust.YFilter = yf }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust"
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModuleTcpMssAdjust.Nodes
    }
    return nil
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModuleTcpMssAdjust.Nodes
    return children
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetYangName() string { return "hardware-module-tcp-mss-adjust" }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) SetParent(parent types.Entity) { hardwareModuleTcpMssAdjust.parent = parent }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetParent() types.Entity { return hardwareModuleTcpMssAdjust.parent }

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-prm-cfg" }

// HardwareModuleTcpMssAdjust_Nodes
// TCP MSS Adjust applicable nodes
type HardwareModuleTcpMssAdjust_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleTcpMssAdjust_Nodes_Node.
    Node []HardwareModuleTcpMssAdjust_Nodes_Node
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleTcpMssAdjust_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetParentYangName() string { return "hardware-module-tcp-mss-adjust" }

// HardwareModuleTcpMssAdjust_Nodes_Node
// A node
type HardwareModuleTcpMssAdjust_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // TCP MSS Adjust NPs.
    Nps HardwareModuleTcpMssAdjust_Nodes_Node_Nps
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "nps" { return "Nps" }
    return ""
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nps" {
        return &node.Nps
    }
    return nil
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nps"] = &node.Nps
    return children
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModuleTcpMssAdjust_Nodes_Node_Nps
// TCP MSS Adjust NPs
type HardwareModuleTcpMssAdjust_Nodes_Node_Nps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NP number. The type is slice of
    // HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np.
    Np []HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetFilter() yfilter.YFilter { return nps.YFilter }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) SetFilter(yf yfilter.YFilter) { nps.YFilter = yf }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetGoName(yname string) string {
    if yname == "np" { return "Np" }
    return ""
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetSegmentPath() string {
    return "nps"
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        for _, c := range nps.Np {
            if nps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np{}
        nps.Np = append(nps.Np, child)
        return &nps.Np[len(nps.Np)-1]
    }
    return nil
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nps.Np {
        children[nps.Np[i].GetSegmentPath()] = &nps.Np[i]
    }
    return children
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetBundleName() string { return "cisco_ios_xr" }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetYangName() string { return "nps" }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) SetParent(parent types.Entity) { nps.parent = parent }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetParent() types.Entity { return nps.parent }

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetParentYangName() string { return "node" }

// HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np
// NP number
type HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Number between 0-7. The type is interface{} with
    // range: 0..7.
    NpId interface{}

    // TCP MSS Adjust value. The type is interface{} with range: 1280..1535. Units
    // are byte.
    AdjustValue interface{}
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetGoName(yname string) string {
    if yname == "np-id" { return "NpId" }
    if yname == "adjust-value" { return "AdjustValue" }
    return ""
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetSegmentPath() string {
    return "np" + "[np-id='" + fmt.Sprintf("%v", np.NpId) + "']"
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["np-id"] = np.NpId
    leafs["adjust-value"] = np.AdjustValue
    return leafs
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetYangName() string { return "np" }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetParent() types.Entity { return np.parent }

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetParentYangName() string { return "nps" }

// HardwareModuleLoadBalance
// hardware module load balance
type HardwareModuleLoadBalance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bundle load balance options.
    Bundle HardwareModuleLoadBalance_Bundle
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetFilter() yfilter.YFilter { return hardwareModuleLoadBalance.YFilter }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) SetFilter(yf yfilter.YFilter) { hardwareModuleLoadBalance.YFilter = yf }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetGoName(yname string) string {
    if yname == "bundle" { return "Bundle" }
    return ""
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-load-balance"
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle" {
        return &hardwareModuleLoadBalance.Bundle
    }
    return nil
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle"] = &hardwareModuleLoadBalance.Bundle
    return children
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetYangName() string { return "hardware-module-load-balance" }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) SetParent(parent types.Entity) { hardwareModuleLoadBalance.parent = parent }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetParent() types.Entity { return hardwareModuleLoadBalance.parent }

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-prm-cfg" }

// HardwareModuleLoadBalance_Bundle
// Bundle load balance options
type HardwareModuleLoadBalance_Bundle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load balance for L2 services.
    L2Service HardwareModuleLoadBalance_Bundle_L2Service
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetFilter() yfilter.YFilter { return bundle.YFilter }

func (bundle *HardwareModuleLoadBalance_Bundle) SetFilter(yf yfilter.YFilter) { bundle.YFilter = yf }

func (bundle *HardwareModuleLoadBalance_Bundle) GetGoName(yname string) string {
    if yname == "l2-service" { return "L2Service" }
    return ""
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetSegmentPath() string {
    return "bundle"
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "l2-service" {
        return &bundle.L2Service
    }
    return nil
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["l2-service"] = &bundle.L2Service
    return children
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetBundleName() string { return "cisco_ios_xr" }

func (bundle *HardwareModuleLoadBalance_Bundle) GetYangName() string { return "bundle" }

func (bundle *HardwareModuleLoadBalance_Bundle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundle *HardwareModuleLoadBalance_Bundle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundle *HardwareModuleLoadBalance_Bundle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundle *HardwareModuleLoadBalance_Bundle) SetParent(parent types.Entity) { bundle.parent = parent }

func (bundle *HardwareModuleLoadBalance_Bundle) GetParent() types.Entity { return bundle.parent }

func (bundle *HardwareModuleLoadBalance_Bundle) GetParentYangName() string { return "hardware-module-load-balance" }

// HardwareModuleLoadBalance_Bundle_L2Service
// Load balance for L2 services
type HardwareModuleLoadBalance_Bundle_L2Service struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load balance L2 services over bundle with L3 parameters. The type is
    // interface{}.
    L3Parameters interface{}
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetFilter() yfilter.YFilter { return l2Service.YFilter }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) SetFilter(yf yfilter.YFilter) { l2Service.YFilter = yf }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetGoName(yname string) string {
    if yname == "l3-parameters" { return "L3Parameters" }
    return ""
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetSegmentPath() string {
    return "l2-service"
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["l3-parameters"] = l2Service.L3Parameters
    return leafs
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetBundleName() string { return "cisco_ios_xr" }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetYangName() string { return "l2-service" }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) SetParent(parent types.Entity) { l2Service.parent = parent }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetParent() types.Entity { return l2Service.parent }

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetParentYangName() string { return "bundle" }

// HardwareModuleTcam
// hardware module tcam
type HardwareModuleTcam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global TCAM partition profile for all TCAM applicable nodes. The type is
    // PrmTcamProfile. The default value is profile0.
    GlobalProfile interface{}

    // TCAM applicable nodes.
    Nodes HardwareModuleTcam_Nodes
}

func (hardwareModuleTcam *HardwareModuleTcam) GetFilter() yfilter.YFilter { return hardwareModuleTcam.YFilter }

func (hardwareModuleTcam *HardwareModuleTcam) SetFilter(yf yfilter.YFilter) { hardwareModuleTcam.YFilter = yf }

func (hardwareModuleTcam *HardwareModuleTcam) GetGoName(yname string) string {
    if yname == "global-profile" { return "GlobalProfile" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModuleTcam *HardwareModuleTcam) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam"
}

func (hardwareModuleTcam *HardwareModuleTcam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &hardwareModuleTcam.Nodes
    }
    return nil
}

func (hardwareModuleTcam *HardwareModuleTcam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &hardwareModuleTcam.Nodes
    return children
}

func (hardwareModuleTcam *HardwareModuleTcam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-profile"] = hardwareModuleTcam.GlobalProfile
    return leafs
}

func (hardwareModuleTcam *HardwareModuleTcam) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleTcam *HardwareModuleTcam) GetYangName() string { return "hardware-module-tcam" }

func (hardwareModuleTcam *HardwareModuleTcam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleTcam *HardwareModuleTcam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleTcam *HardwareModuleTcam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleTcam *HardwareModuleTcam) SetParent(parent types.Entity) { hardwareModuleTcam.parent = parent }

func (hardwareModuleTcam *HardwareModuleTcam) GetParent() types.Entity { return hardwareModuleTcam.parent }

func (hardwareModuleTcam *HardwareModuleTcam) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-prm-cfg" }

// HardwareModuleTcam_Nodes
// TCAM applicable nodes
type HardwareModuleTcam_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A TCAM applicable node. The type is slice of HardwareModuleTcam_Nodes_Node.
    Node []HardwareModuleTcam_Nodes_Node
}

func (nodes *HardwareModuleTcam_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModuleTcam_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModuleTcam_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModuleTcam_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModuleTcam_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleTcam_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModuleTcam_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModuleTcam_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModuleTcam_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModuleTcam_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModuleTcam_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModuleTcam_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModuleTcam_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModuleTcam_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModuleTcam_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModuleTcam_Nodes) GetParentYangName() string { return "hardware-module-tcam" }

// HardwareModuleTcam_Nodes_Node
// A TCAM applicable node
type HardwareModuleTcam_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // A TCAM partition profile. The type is PrmTcamProfile. The default value is
    // profile0.
    Profile interface{}
}

func (node *HardwareModuleTcam_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModuleTcam_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModuleTcam_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "profile" { return "Profile" }
    return ""
}

func (node *HardwareModuleTcam_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModuleTcam_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *HardwareModuleTcam_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *HardwareModuleTcam_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["profile"] = node.Profile
    return leafs
}

func (node *HardwareModuleTcam_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModuleTcam_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModuleTcam_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModuleTcam_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModuleTcam_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModuleTcam_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModuleTcam_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModuleTcam_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModuleEfd
// hardware module efd
type HardwareModuleEfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All nodes.
    NodeAll HardwareModuleEfd_NodeAll

    // EFD applicable nodes.
    Nodes HardwareModuleEfd_Nodes
}

func (hardwareModuleEfd *HardwareModuleEfd) GetFilter() yfilter.YFilter { return hardwareModuleEfd.YFilter }

func (hardwareModuleEfd *HardwareModuleEfd) SetFilter(yf yfilter.YFilter) { hardwareModuleEfd.YFilter = yf }

func (hardwareModuleEfd *HardwareModuleEfd) GetGoName(yname string) string {
    if yname == "node-all" { return "NodeAll" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (hardwareModuleEfd *HardwareModuleEfd) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd"
}

func (hardwareModuleEfd *HardwareModuleEfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-all" {
        return &hardwareModuleEfd.NodeAll
    }
    if childYangName == "nodes" {
        return &hardwareModuleEfd.Nodes
    }
    return nil
}

func (hardwareModuleEfd *HardwareModuleEfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-all"] = &hardwareModuleEfd.NodeAll
    children["nodes"] = &hardwareModuleEfd.Nodes
    return children
}

func (hardwareModuleEfd *HardwareModuleEfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareModuleEfd *HardwareModuleEfd) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleEfd *HardwareModuleEfd) GetYangName() string { return "hardware-module-efd" }

func (hardwareModuleEfd *HardwareModuleEfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleEfd *HardwareModuleEfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleEfd *HardwareModuleEfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleEfd *HardwareModuleEfd) SetParent(parent types.Entity) { hardwareModuleEfd.parent = parent }

func (hardwareModuleEfd *HardwareModuleEfd) GetParent() types.Entity { return hardwareModuleEfd.parent }

func (hardwareModuleEfd *HardwareModuleEfd) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-prm-cfg" }

// HardwareModuleEfd_NodeAll
// All nodes
type HardwareModuleEfd_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable EFD for this node. The type is interface{}.
    Enable interface{}

    // EFD mode parameter. The type is Asr9kEfdMode.
    Mode interface{}

    // VLAN Priority Mask.
    VlanPriorityMask HardwareModuleEfd_NodeAll_VlanPriorityMask

    // EFD IP parameters.
    IpPrecedence HardwareModuleEfd_NodeAll_IpPrecedence

    // EFD VLAN parameters.
    VlanCos HardwareModuleEfd_NodeAll_VlanCos

    // IP Priority Mask.
    IpPriorityMask HardwareModuleEfd_NodeAll_IpPriorityMask

    // MPLS Priority Mask.
    MplsPriorityMask HardwareModuleEfd_NodeAll_MplsPriorityMask

    // EFD MPLS parameters.
    MplsExp HardwareModuleEfd_NodeAll_MplsExp
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *HardwareModuleEfd_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *HardwareModuleEfd_NodeAll) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "mode" { return "Mode" }
    if yname == "vlan-priority-mask" { return "VlanPriorityMask" }
    if yname == "ip-precedence" { return "IpPrecedence" }
    if yname == "vlan-cos" { return "VlanCos" }
    if yname == "ip-priority-mask" { return "IpPriorityMask" }
    if yname == "mpls-priority-mask" { return "MplsPriorityMask" }
    if yname == "mpls-exp" { return "MplsExp" }
    return ""
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-priority-mask" {
        return &nodeAll.VlanPriorityMask
    }
    if childYangName == "ip-precedence" {
        return &nodeAll.IpPrecedence
    }
    if childYangName == "vlan-cos" {
        return &nodeAll.VlanCos
    }
    if childYangName == "ip-priority-mask" {
        return &nodeAll.IpPriorityMask
    }
    if childYangName == "mpls-priority-mask" {
        return &nodeAll.MplsPriorityMask
    }
    if childYangName == "mpls-exp" {
        return &nodeAll.MplsExp
    }
    return nil
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vlan-priority-mask"] = &nodeAll.VlanPriorityMask
    children["ip-precedence"] = &nodeAll.IpPrecedence
    children["vlan-cos"] = &nodeAll.VlanCos
    children["ip-priority-mask"] = &nodeAll.IpPriorityMask
    children["mpls-priority-mask"] = &nodeAll.MplsPriorityMask
    children["mpls-exp"] = &nodeAll.MplsExp
    return children
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = nodeAll.Enable
    leafs["mode"] = nodeAll.Mode
    return leafs
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *HardwareModuleEfd_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *HardwareModuleEfd_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *HardwareModuleEfd_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *HardwareModuleEfd_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *HardwareModuleEfd_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *HardwareModuleEfd_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *HardwareModuleEfd_NodeAll) GetParentYangName() string { return "hardware-module-efd" }

// HardwareModuleEfd_NodeAll_VlanPriorityMask
// VLAN Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_VlanPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetFilter() yfilter.YFilter { return vlanPriorityMask.YFilter }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) SetFilter(yf yfilter.YFilter) { vlanPriorityMask.YFilter = yf }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetSegmentPath() string {
    return "vlan-priority-mask"
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = vlanPriorityMask.Prec0
    leafs["prec1"] = vlanPriorityMask.Prec1
    leafs["prec2"] = vlanPriorityMask.Prec2
    leafs["prec3"] = vlanPriorityMask.Prec3
    leafs["prec4"] = vlanPriorityMask.Prec4
    leafs["prec5"] = vlanPriorityMask.Prec5
    leafs["prec6"] = vlanPriorityMask.Prec6
    leafs["prec7"] = vlanPriorityMask.Prec7
    return leafs
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetYangName() string { return "vlan-priority-mask" }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) SetParent(parent types.Entity) { vlanPriorityMask.parent = parent }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetParent() types.Entity { return vlanPriorityMask.parent }

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_NodeAll_IpPrecedence
// EFD IP parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_IpPrecedence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP TOS precedence threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Precedence interface{}

    // IP operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetFilter() yfilter.YFilter { return ipPrecedence.YFilter }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) SetFilter(yf yfilter.YFilter) { ipPrecedence.YFilter = yf }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetGoName(yname string) string {
    if yname == "precedence" { return "Precedence" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetSegmentPath() string {
    return "ip-precedence"
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["precedence"] = ipPrecedence.Precedence
    leafs["operation"] = ipPrecedence.Operation
    return leafs
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetBundleName() string { return "cisco_ios_xr" }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetYangName() string { return "ip-precedence" }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) SetParent(parent types.Entity) { ipPrecedence.parent = parent }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetParent() types.Entity { return ipPrecedence.parent }

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_NodeAll_VlanCos
// EFD VLAN parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_VlanCos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN COS threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Cos interface{}

    // VLAN operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetFilter() yfilter.YFilter { return vlanCos.YFilter }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) SetFilter(yf yfilter.YFilter) { vlanCos.YFilter = yf }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetGoName(yname string) string {
    if yname == "cos" { return "Cos" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetSegmentPath() string {
    return "vlan-cos"
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cos"] = vlanCos.Cos
    leafs["operation"] = vlanCos.Operation
    return leafs
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetBundleName() string { return "cisco_ios_xr" }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetYangName() string { return "vlan-cos" }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) SetParent(parent types.Entity) { vlanCos.parent = parent }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetParent() types.Entity { return vlanCos.parent }

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_NodeAll_IpPriorityMask
// IP Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_IpPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetFilter() yfilter.YFilter { return ipPriorityMask.YFilter }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) SetFilter(yf yfilter.YFilter) { ipPriorityMask.YFilter = yf }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetSegmentPath() string {
    return "ip-priority-mask"
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = ipPriorityMask.Prec0
    leafs["prec1"] = ipPriorityMask.Prec1
    leafs["prec2"] = ipPriorityMask.Prec2
    leafs["prec3"] = ipPriorityMask.Prec3
    leafs["prec4"] = ipPriorityMask.Prec4
    leafs["prec5"] = ipPriorityMask.Prec5
    leafs["prec6"] = ipPriorityMask.Prec6
    leafs["prec7"] = ipPriorityMask.Prec7
    return leafs
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetYangName() string { return "ip-priority-mask" }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) SetParent(parent types.Entity) { ipPriorityMask.parent = parent }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetParent() types.Entity { return ipPriorityMask.parent }

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_NodeAll_MplsPriorityMask
// MPLS Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_MplsPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetFilter() yfilter.YFilter { return mplsPriorityMask.YFilter }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) SetFilter(yf yfilter.YFilter) { mplsPriorityMask.YFilter = yf }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetSegmentPath() string {
    return "mpls-priority-mask"
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = mplsPriorityMask.Prec0
    leafs["prec1"] = mplsPriorityMask.Prec1
    leafs["prec2"] = mplsPriorityMask.Prec2
    leafs["prec3"] = mplsPriorityMask.Prec3
    leafs["prec4"] = mplsPriorityMask.Prec4
    leafs["prec5"] = mplsPriorityMask.Prec5
    leafs["prec6"] = mplsPriorityMask.Prec6
    leafs["prec7"] = mplsPriorityMask.Prec7
    return leafs
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetYangName() string { return "mpls-priority-mask" }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) SetParent(parent types.Entity) { mplsPriorityMask.parent = parent }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetParent() types.Entity { return mplsPriorityMask.parent }

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_NodeAll_MplsExp
// EFD MPLS parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_MplsExp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS EXP threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Exp interface{}

    // MPLS operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetFilter() yfilter.YFilter { return mplsExp.YFilter }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) SetFilter(yf yfilter.YFilter) { mplsExp.YFilter = yf }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetGoName(yname string) string {
    if yname == "exp" { return "Exp" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetSegmentPath() string {
    return "mpls-exp"
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exp"] = mplsExp.Exp
    leafs["operation"] = mplsExp.Operation
    return leafs
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetYangName() string { return "mpls-exp" }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) SetParent(parent types.Entity) { mplsExp.parent = parent }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetParent() types.Entity { return mplsExp.parent }

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetParentYangName() string { return "node-all" }

// HardwareModuleEfd_Nodes
// EFD applicable nodes
type HardwareModuleEfd_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleEfd_Nodes_Node.
    Node []HardwareModuleEfd_Nodes_Node
}

func (nodes *HardwareModuleEfd_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *HardwareModuleEfd_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *HardwareModuleEfd_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *HardwareModuleEfd_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *HardwareModuleEfd_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HardwareModuleEfd_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *HardwareModuleEfd_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *HardwareModuleEfd_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *HardwareModuleEfd_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *HardwareModuleEfd_Nodes) GetYangName() string { return "nodes" }

func (nodes *HardwareModuleEfd_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *HardwareModuleEfd_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *HardwareModuleEfd_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *HardwareModuleEfd_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *HardwareModuleEfd_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *HardwareModuleEfd_Nodes) GetParentYangName() string { return "hardware-module-efd" }

// HardwareModuleEfd_Nodes_Node
// A node
type HardwareModuleEfd_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node Name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Enable EFD for this node. The type is interface{}.
    Enable interface{}

    // EFD mode parameter. The type is Asr9kEfdMode.
    Mode interface{}

    // VLAN Priority Mask.
    VlanPriorityMask HardwareModuleEfd_Nodes_Node_VlanPriorityMask

    // EFD IP parameters.
    IpPrecedence HardwareModuleEfd_Nodes_Node_IpPrecedence

    // EFD VLAN parameters.
    VlanCos HardwareModuleEfd_Nodes_Node_VlanCos

    // IP Priority Mask.
    IpPriorityMask HardwareModuleEfd_Nodes_Node_IpPriorityMask

    // MPLS Priority Mask.
    MplsPriorityMask HardwareModuleEfd_Nodes_Node_MplsPriorityMask

    // EFD MPLS parameters.
    MplsExp HardwareModuleEfd_Nodes_Node_MplsExp
}

func (node *HardwareModuleEfd_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *HardwareModuleEfd_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *HardwareModuleEfd_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "enable" { return "Enable" }
    if yname == "mode" { return "Mode" }
    if yname == "vlan-priority-mask" { return "VlanPriorityMask" }
    if yname == "ip-precedence" { return "IpPrecedence" }
    if yname == "vlan-cos" { return "VlanCos" }
    if yname == "ip-priority-mask" { return "IpPriorityMask" }
    if yname == "mpls-priority-mask" { return "MplsPriorityMask" }
    if yname == "mpls-exp" { return "MplsExp" }
    return ""
}

func (node *HardwareModuleEfd_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *HardwareModuleEfd_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-priority-mask" {
        return &node.VlanPriorityMask
    }
    if childYangName == "ip-precedence" {
        return &node.IpPrecedence
    }
    if childYangName == "vlan-cos" {
        return &node.VlanCos
    }
    if childYangName == "ip-priority-mask" {
        return &node.IpPriorityMask
    }
    if childYangName == "mpls-priority-mask" {
        return &node.MplsPriorityMask
    }
    if childYangName == "mpls-exp" {
        return &node.MplsExp
    }
    return nil
}

func (node *HardwareModuleEfd_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vlan-priority-mask"] = &node.VlanPriorityMask
    children["ip-precedence"] = &node.IpPrecedence
    children["vlan-cos"] = &node.VlanCos
    children["ip-priority-mask"] = &node.IpPriorityMask
    children["mpls-priority-mask"] = &node.MplsPriorityMask
    children["mpls-exp"] = &node.MplsExp
    return children
}

func (node *HardwareModuleEfd_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["enable"] = node.Enable
    leafs["mode"] = node.Mode
    return leafs
}

func (node *HardwareModuleEfd_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *HardwareModuleEfd_Nodes_Node) GetYangName() string { return "node" }

func (node *HardwareModuleEfd_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *HardwareModuleEfd_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *HardwareModuleEfd_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *HardwareModuleEfd_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *HardwareModuleEfd_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *HardwareModuleEfd_Nodes_Node) GetParentYangName() string { return "nodes" }

// HardwareModuleEfd_Nodes_Node_VlanPriorityMask
// VLAN Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_VlanPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetFilter() yfilter.YFilter { return vlanPriorityMask.YFilter }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) SetFilter(yf yfilter.YFilter) { vlanPriorityMask.YFilter = yf }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetSegmentPath() string {
    return "vlan-priority-mask"
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = vlanPriorityMask.Prec0
    leafs["prec1"] = vlanPriorityMask.Prec1
    leafs["prec2"] = vlanPriorityMask.Prec2
    leafs["prec3"] = vlanPriorityMask.Prec3
    leafs["prec4"] = vlanPriorityMask.Prec4
    leafs["prec5"] = vlanPriorityMask.Prec5
    leafs["prec6"] = vlanPriorityMask.Prec6
    leafs["prec7"] = vlanPriorityMask.Prec7
    return leafs
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetYangName() string { return "vlan-priority-mask" }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) SetParent(parent types.Entity) { vlanPriorityMask.parent = parent }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetParent() types.Entity { return vlanPriorityMask.parent }

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetParentYangName() string { return "node" }

// HardwareModuleEfd_Nodes_Node_IpPrecedence
// EFD IP parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_IpPrecedence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP TOS precedence threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Precedence interface{}

    // IP operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetFilter() yfilter.YFilter { return ipPrecedence.YFilter }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) SetFilter(yf yfilter.YFilter) { ipPrecedence.YFilter = yf }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetGoName(yname string) string {
    if yname == "precedence" { return "Precedence" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetSegmentPath() string {
    return "ip-precedence"
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["precedence"] = ipPrecedence.Precedence
    leafs["operation"] = ipPrecedence.Operation
    return leafs
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetBundleName() string { return "cisco_ios_xr" }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetYangName() string { return "ip-precedence" }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) SetParent(parent types.Entity) { ipPrecedence.parent = parent }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetParent() types.Entity { return ipPrecedence.parent }

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetParentYangName() string { return "node" }

// HardwareModuleEfd_Nodes_Node_VlanCos
// EFD VLAN parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_VlanCos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN COS threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Cos interface{}

    // VLAN operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetFilter() yfilter.YFilter { return vlanCos.YFilter }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) SetFilter(yf yfilter.YFilter) { vlanCos.YFilter = yf }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetGoName(yname string) string {
    if yname == "cos" { return "Cos" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetSegmentPath() string {
    return "vlan-cos"
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cos"] = vlanCos.Cos
    leafs["operation"] = vlanCos.Operation
    return leafs
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetBundleName() string { return "cisco_ios_xr" }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetYangName() string { return "vlan-cos" }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) SetParent(parent types.Entity) { vlanCos.parent = parent }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetParent() types.Entity { return vlanCos.parent }

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetParentYangName() string { return "node" }

// HardwareModuleEfd_Nodes_Node_IpPriorityMask
// IP Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_IpPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetFilter() yfilter.YFilter { return ipPriorityMask.YFilter }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) SetFilter(yf yfilter.YFilter) { ipPriorityMask.YFilter = yf }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetSegmentPath() string {
    return "ip-priority-mask"
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = ipPriorityMask.Prec0
    leafs["prec1"] = ipPriorityMask.Prec1
    leafs["prec2"] = ipPriorityMask.Prec2
    leafs["prec3"] = ipPriorityMask.Prec3
    leafs["prec4"] = ipPriorityMask.Prec4
    leafs["prec5"] = ipPriorityMask.Prec5
    leafs["prec6"] = ipPriorityMask.Prec6
    leafs["prec7"] = ipPriorityMask.Prec7
    return leafs
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetYangName() string { return "ip-priority-mask" }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) SetParent(parent types.Entity) { ipPriorityMask.parent = parent }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetParent() types.Entity { return ipPriorityMask.parent }

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetParentYangName() string { return "node" }

// HardwareModuleEfd_Nodes_Node_MplsPriorityMask
// MPLS Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_MplsPriorityMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetFilter() yfilter.YFilter { return mplsPriorityMask.YFilter }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) SetFilter(yf yfilter.YFilter) { mplsPriorityMask.YFilter = yf }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetGoName(yname string) string {
    if yname == "prec0" { return "Prec0" }
    if yname == "prec1" { return "Prec1" }
    if yname == "prec2" { return "Prec2" }
    if yname == "prec3" { return "Prec3" }
    if yname == "prec4" { return "Prec4" }
    if yname == "prec5" { return "Prec5" }
    if yname == "prec6" { return "Prec6" }
    if yname == "prec7" { return "Prec7" }
    return ""
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetSegmentPath() string {
    return "mpls-priority-mask"
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prec0"] = mplsPriorityMask.Prec0
    leafs["prec1"] = mplsPriorityMask.Prec1
    leafs["prec2"] = mplsPriorityMask.Prec2
    leafs["prec3"] = mplsPriorityMask.Prec3
    leafs["prec4"] = mplsPriorityMask.Prec4
    leafs["prec5"] = mplsPriorityMask.Prec5
    leafs["prec6"] = mplsPriorityMask.Prec6
    leafs["prec7"] = mplsPriorityMask.Prec7
    return leafs
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetBundleName() string { return "cisco_ios_xr" }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetYangName() string { return "mpls-priority-mask" }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) SetParent(parent types.Entity) { mplsPriorityMask.parent = parent }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetParent() types.Entity { return mplsPriorityMask.parent }

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetParentYangName() string { return "node" }

// HardwareModuleEfd_Nodes_Node_MplsExp
// EFD MPLS parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_MplsExp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS EXP threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Exp interface{}

    // MPLS operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetFilter() yfilter.YFilter { return mplsExp.YFilter }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) SetFilter(yf yfilter.YFilter) { mplsExp.YFilter = yf }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetGoName(yname string) string {
    if yname == "exp" { return "Exp" }
    if yname == "operation" { return "Operation" }
    return ""
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetSegmentPath() string {
    return "mpls-exp"
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exp"] = mplsExp.Exp
    leafs["operation"] = mplsExp.Operation
    return leafs
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetYangName() string { return "mpls-exp" }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) SetParent(parent types.Entity) { mplsExp.parent = parent }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetParent() types.Entity { return mplsExp.parent }

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetParentYangName() string { return "node" }

