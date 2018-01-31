// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-lldp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lldp: Link Layer Discovery Protocol operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_lldp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_lldp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-lldp-oper lldp}", reflect.TypeOf(Lldp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-lldp-oper:lldp", reflect.TypeOf(Lldp{}))
}

// LldpL3AddrProtocol represents Lldp l3 addr protocol
type LldpL3AddrProtocol string

const (
    // IPv4
    LldpL3AddrProtocol_ipv4 LldpL3AddrProtocol = "ipv4"

    // IPv6
    LldpL3AddrProtocol_ipv6 LldpL3AddrProtocol = "ipv6"
)

// Lldp
// Link Layer Discovery Protocol operational data
type Lldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global LLDP data.
    GlobalLldp Lldp_GlobalLldp

    // Per node LLDP operational data.
    Nodes Lldp_Nodes
}

func (lldp *Lldp) GetFilter() yfilter.YFilter { return lldp.YFilter }

func (lldp *Lldp) SetFilter(yf yfilter.YFilter) { lldp.YFilter = yf }

func (lldp *Lldp) GetGoName(yname string) string {
    if yname == "global-lldp" { return "GlobalLldp" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (lldp *Lldp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-lldp-oper:lldp"
}

func (lldp *Lldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-lldp" {
        return &lldp.GlobalLldp
    }
    if childYangName == "nodes" {
        return &lldp.Nodes
    }
    return nil
}

func (lldp *Lldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-lldp"] = &lldp.GlobalLldp
    children["nodes"] = &lldp.Nodes
    return children
}

func (lldp *Lldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lldp *Lldp) GetBundleName() string { return "cisco_ios_xr" }

func (lldp *Lldp) GetYangName() string { return "lldp" }

func (lldp *Lldp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldp *Lldp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldp *Lldp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldp *Lldp) SetParent(parent types.Entity) { lldp.parent = parent }

func (lldp *Lldp) GetParent() types.Entity { return lldp.parent }

func (lldp *Lldp) GetParentYangName() string { return "Cisco-IOS-XR-ethernet-lldp-oper" }

// Lldp_GlobalLldp
// Global LLDP data
type Lldp_GlobalLldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The LLDP Global Information of this box.
    LldpInfo Lldp_GlobalLldp_LldpInfo
}

func (globalLldp *Lldp_GlobalLldp) GetFilter() yfilter.YFilter { return globalLldp.YFilter }

func (globalLldp *Lldp_GlobalLldp) SetFilter(yf yfilter.YFilter) { globalLldp.YFilter = yf }

func (globalLldp *Lldp_GlobalLldp) GetGoName(yname string) string {
    if yname == "lldp-info" { return "LldpInfo" }
    return ""
}

func (globalLldp *Lldp_GlobalLldp) GetSegmentPath() string {
    return "global-lldp"
}

func (globalLldp *Lldp_GlobalLldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-info" {
        return &globalLldp.LldpInfo
    }
    return nil
}

func (globalLldp *Lldp_GlobalLldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lldp-info"] = &globalLldp.LldpInfo
    return children
}

func (globalLldp *Lldp_GlobalLldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalLldp *Lldp_GlobalLldp) GetBundleName() string { return "cisco_ios_xr" }

func (globalLldp *Lldp_GlobalLldp) GetYangName() string { return "global-lldp" }

func (globalLldp *Lldp_GlobalLldp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalLldp *Lldp_GlobalLldp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalLldp *Lldp_GlobalLldp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalLldp *Lldp_GlobalLldp) SetParent(parent types.Entity) { globalLldp.parent = parent }

func (globalLldp *Lldp_GlobalLldp) GetParent() types.Entity { return globalLldp.parent }

func (globalLldp *Lldp_GlobalLldp) GetParentYangName() string { return "lldp" }

// Lldp_GlobalLldp_LldpInfo
// The LLDP Global Information of this box
type Lldp_GlobalLldp_LldpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Chassis identifier. The type is string.
    ChassisId interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // System Name. The type is string.
    SystemName interface{}

    // Rate at which LLDP packets re sent (in sec). The type is interface{} with
    // range: 0..4294967295.
    Timer interface{}

    // Length  of time  (in sec)that receiver must keep thispacket. The type is
    // interface{} with range: 0..4294967295.
    HoldTime interface{}

    // Delay (in sec) for LLDPinitialization on anyinterface. The type is
    // interface{} with range: 0..4294967295.
    ReInit interface{}
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetFilter() yfilter.YFilter { return lldpInfo.YFilter }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) SetFilter(yf yfilter.YFilter) { lldpInfo.YFilter = yf }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetGoName(yname string) string {
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "chassis-id-sub-type" { return "ChassisIdSubType" }
    if yname == "system-name" { return "SystemName" }
    if yname == "timer" { return "Timer" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "re-init" { return "ReInit" }
    return ""
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetSegmentPath() string {
    return "lldp-info"
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id"] = lldpInfo.ChassisId
    leafs["chassis-id-sub-type"] = lldpInfo.ChassisIdSubType
    leafs["system-name"] = lldpInfo.SystemName
    leafs["timer"] = lldpInfo.Timer
    leafs["hold-time"] = lldpInfo.HoldTime
    leafs["re-init"] = lldpInfo.ReInit
    return leafs
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetYangName() string { return "lldp-info" }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) SetParent(parent types.Entity) { lldpInfo.parent = parent }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetParent() types.Entity { return lldpInfo.parent }

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetParentYangName() string { return "global-lldp" }

// Lldp_Nodes
// Per node LLDP operational data
type Lldp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The LLDP operational data for a particular node. The type is slice of
    // Lldp_Nodes_Node.
    Node []Lldp_Nodes_Node
}

func (nodes *Lldp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Lldp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Lldp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Lldp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Lldp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Lldp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Lldp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Lldp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Lldp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Lldp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Lldp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Lldp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Lldp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Lldp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Lldp_Nodes) GetParentYangName() string { return "lldp" }

// Lldp_Nodes_Node
// The LLDP operational data for a particular node
type Lldp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // The LLDP neighbor tables on this node.
    Neighbors Lldp_Nodes_Node_Neighbors

    // The table of interfaces on which LLDP is running on this node.
    Interfaces Lldp_Nodes_Node_Interfaces

    // The LLDP traffic statistics for this node.
    Statistics Lldp_Nodes_Node_Statistics
}

func (node *Lldp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Lldp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Lldp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Lldp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Lldp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &node.Neighbors
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Lldp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &node.Neighbors
    children["interfaces"] = &node.Interfaces
    children["statistics"] = &node.Statistics
    return children
}

func (node *Lldp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Lldp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Lldp_Nodes_Node) GetYangName() string { return "node" }

func (node *Lldp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Lldp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Lldp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Lldp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Lldp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Lldp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Lldp_Nodes_Node_Neighbors
// The LLDP neighbor tables on this node
type Lldp_Nodes_Node_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The detailed LLDP neighbor table on this device.
    Devices Lldp_Nodes_Node_Neighbors_Devices

    // The detailed LLDP neighbor table.
    Details Lldp_Nodes_Node_Neighbors_Details

    // The LLDP neighbor summary table.
    Summaries Lldp_Nodes_Node_Neighbors_Summaries
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Lldp_Nodes_Node_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetGoName(yname string) string {
    if yname == "devices" { return "Devices" }
    if yname == "details" { return "Details" }
    if yname == "summaries" { return "Summaries" }
    return ""
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "devices" {
        return &neighbors.Devices
    }
    if childYangName == "details" {
        return &neighbors.Details
    }
    if childYangName == "summaries" {
        return &neighbors.Summaries
    }
    return nil
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["devices"] = &neighbors.Devices
    children["details"] = &neighbors.Details
    children["summaries"] = &neighbors.Summaries
    return children
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *Lldp_Nodes_Node_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Lldp_Nodes_Node_Neighbors) GetParentYangName() string { return "node" }

// Lldp_Nodes_Node_Neighbors_Devices
// The detailed LLDP neighbor table on this
// device
type Lldp_Nodes_Node_Neighbors_Devices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device.
    Device []Lldp_Nodes_Node_Neighbors_Devices_Device
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetFilter() yfilter.YFilter { return devices.YFilter }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) SetFilter(yf yfilter.YFilter) { devices.YFilter = yf }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetGoName(yname string) string {
    if yname == "device" { return "Device" }
    return ""
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetSegmentPath() string {
    return "devices"
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device" {
        for _, c := range devices.Device {
            if devices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Devices_Device{}
        devices.Device = append(devices.Device, child)
        return &devices.Device[len(devices.Device)-1]
    }
    return nil
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range devices.Device {
        children[devices.Device[i].GetSegmentPath()] = &devices.Device[i]
    }
    return children
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetBundleName() string { return "cisco_ios_xr" }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetYangName() string { return "devices" }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) SetParent(parent types.Entity) { devices.parent = parent }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetParent() types.Entity { return devices.parent }

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetParentYangName() string { return "neighbors" }

// Lldp_Nodes_Node_Neighbors_Devices_Device
// Detailed information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Devices_Device struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor.
    LldpNeighbor []Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetFilter() yfilter.YFilter { return device.YFilter }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) SetFilter(yf yfilter.YFilter) { device.YFilter = yf }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetGoName(yname string) string {
    if yname == "device-id" { return "DeviceId" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "lldp-neighbor" { return "LldpNeighbor" }
    return ""
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetSegmentPath() string {
    return "device"
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-neighbor" {
        for _, c := range device.LldpNeighbor {
            if device.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor{}
        device.LldpNeighbor = append(device.LldpNeighbor, child)
        return &device.LldpNeighbor[len(device.LldpNeighbor)-1]
    }
    return nil
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range device.LldpNeighbor {
        children[device.LldpNeighbor[i].GetSegmentPath()] = &device.LldpNeighbor[i]
    }
    return children
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-id"] = device.DeviceId
    leafs["interface-name"] = device.InterfaceName
    return leafs
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetBundleName() string { return "cisco_ios_xr" }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetYangName() string { return "device" }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) SetParent(parent types.Entity) { device.parent = parent }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetParent() types.Entity { return device.parent }

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetParentYangName() string { return "devices" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetFilter() yfilter.YFilter { return lldpNeighbor.YFilter }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) SetFilter(yf yfilter.YFilter) { lldpNeighbor.YFilter = yf }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "receiving-parent-interface-name" { return "ReceivingParentInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "port-id-detail" { return "PortIdDetail" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    if yname == "mib" { return "Mib" }
    return ""
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetSegmentPath() string {
    return "lldp-neighbor"
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &lldpNeighbor.Detail
    }
    if childYangName == "mib" {
        return &lldpNeighbor.Mib
    }
    return nil
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &lldpNeighbor.Detail
    children["mib"] = &lldpNeighbor.Mib
    return children
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = lldpNeighbor.ReceivingInterfaceName
    leafs["receiving-parent-interface-name"] = lldpNeighbor.ReceivingParentInterfaceName
    leafs["device-id"] = lldpNeighbor.DeviceId
    leafs["chassis-id"] = lldpNeighbor.ChassisId
    leafs["port-id-detail"] = lldpNeighbor.PortIdDetail
    leafs["header-version"] = lldpNeighbor.HeaderVersion
    leafs["hold-time"] = lldpNeighbor.HoldTime
    leafs["enabled-capabilities"] = lldpNeighbor.EnabledCapabilities
    leafs["platform"] = lldpNeighbor.Platform
    return leafs
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetYangName() string { return "lldp-neighbor" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) SetParent(parent types.Entity) { lldpNeighbor.parent = parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetParent() types.Entity { return lldpNeighbor.parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetParentYangName() string { return "device" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "peer-mac-address" { return "PeerMacAddress" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "system-name" { return "SystemName" }
    if yname == "system-description" { return "SystemDescription" }
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "auto-negotiation" { return "AutoNegotiation" }
    if yname == "physical-media-capabilities" { return "PhysicalMediaCapabilities" }
    if yname == "media-attachment-unit-type" { return "MediaAttachmentUnitType" }
    if yname == "port-vlan-id" { return "PortVlanId" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    return ""
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    return nil
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    return children
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-mac-address"] = detail.PeerMacAddress
    leafs["port-description"] = detail.PortDescription
    leafs["system-name"] = detail.SystemName
    leafs["system-description"] = detail.SystemDescription
    leafs["time-remaining"] = detail.TimeRemaining
    leafs["system-capabilities"] = detail.SystemCapabilities
    leafs["enabled-capabilities"] = detail.EnabledCapabilities
    leafs["auto-negotiation"] = detail.AutoNegotiation
    leafs["physical-media-capabilities"] = detail.PhysicalMediaCapabilities
    leafs["media-attachment-unit-type"] = detail.MediaAttachmentUnitType
    leafs["port-vlan-id"] = detail.PortVlanId
    return leafs
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "lldp-addr-entry" { return "LldpAddrEntry" }
    return ""
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-addr-entry" {
        for _, c := range networkAddresses.LldpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry{}
        networkAddresses.LldpAddrEntry = append(networkAddresses.LldpAddrEntry, child)
        return &networkAddresses.LldpAddrEntry[len(networkAddresses.LldpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.LldpAddrEntry {
        children[networkAddresses.LldpAddrEntry[i].GetSegmentPath()] = &networkAddresses.LldpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetFilter() yfilter.YFilter { return lldpAddrEntry.YFilter }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetFilter(yf yfilter.YFilter) { lldpAddrEntry.YFilter = yf }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetGoName(yname string) string {
    if yname == "ma-subtype" { return "MaSubtype" }
    if yname == "if-num" { return "IfNum" }
    if yname == "address" { return "Address" }
    return ""
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetSegmentPath() string {
    return "lldp-addr-entry"
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &lldpAddrEntry.Address
    }
    return nil
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &lldpAddrEntry.Address
    return children
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ma-subtype"] = lldpAddrEntry.MaSubtype
    leafs["if-num"] = lldpAddrEntry.IfNum
    return leafs
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetYangName() string { return "lldp-addr-entry" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetParent(parent types.Entity) { lldpAddrEntry.parent = parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParent() types.Entity { return lldpAddrEntry.parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParentYangName() string { return "lldp-addr-entry" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetFilter() yfilter.YFilter { return mib.YFilter }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) SetFilter(yf yfilter.YFilter) { mib.YFilter = yf }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetGoName(yname string) string {
    if yname == "rem-time-mark" { return "RemTimeMark" }
    if yname == "rem-local-port-num" { return "RemLocalPortNum" }
    if yname == "rem-index" { return "RemIndex" }
    if yname == "chassis-id-sub-type" { return "ChassisIdSubType" }
    if yname == "chassis-id-len" { return "ChassisIdLen" }
    if yname == "port-id-sub-type" { return "PortIdSubType" }
    if yname == "port-id-len" { return "PortIdLen" }
    if yname == "combined-capabilities" { return "CombinedCapabilities" }
    if yname == "unknown-tlv-list" { return "UnknownTlvList" }
    if yname == "org-def-tlv-list" { return "OrgDefTlvList" }
    return ""
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetSegmentPath() string {
    return "mib"
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv-list" {
        return &mib.UnknownTlvList
    }
    if childYangName == "org-def-tlv-list" {
        return &mib.OrgDefTlvList
    }
    return nil
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unknown-tlv-list"] = &mib.UnknownTlvList
    children["org-def-tlv-list"] = &mib.OrgDefTlvList
    return children
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rem-time-mark"] = mib.RemTimeMark
    leafs["rem-local-port-num"] = mib.RemLocalPortNum
    leafs["rem-index"] = mib.RemIndex
    leafs["chassis-id-sub-type"] = mib.ChassisIdSubType
    leafs["chassis-id-len"] = mib.ChassisIdLen
    leafs["port-id-sub-type"] = mib.PortIdSubType
    leafs["port-id-len"] = mib.PortIdLen
    leafs["combined-capabilities"] = mib.CombinedCapabilities
    return leafs
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetBundleName() string { return "cisco_ios_xr" }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetYangName() string { return "mib" }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) SetParent(parent types.Entity) { mib.parent = parent }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetParent() types.Entity { return mib.parent }

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetFilter() yfilter.YFilter { return unknownTlvList.YFilter }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) SetFilter(yf yfilter.YFilter) { unknownTlvList.YFilter = yf }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetGoName(yname string) string {
    if yname == "lldp-unknown-tlv-entry" { return "LldpUnknownTlvEntry" }
    return ""
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetSegmentPath() string {
    return "unknown-tlv-list"
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-unknown-tlv-entry" {
        for _, c := range unknownTlvList.LldpUnknownTlvEntry {
            if unknownTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry{}
        unknownTlvList.LldpUnknownTlvEntry = append(unknownTlvList.LldpUnknownTlvEntry, child)
        return &unknownTlvList.LldpUnknownTlvEntry[len(unknownTlvList.LldpUnknownTlvEntry)-1]
    }
    return nil
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        children[unknownTlvList.LldpUnknownTlvEntry[i].GetSegmentPath()] = &unknownTlvList.LldpUnknownTlvEntry[i]
    }
    return children
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetYangName() string { return "unknown-tlv-list" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) SetParent(parent types.Entity) { unknownTlvList.parent = parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetParent() types.Entity { return unknownTlvList.parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetFilter() yfilter.YFilter { return lldpUnknownTlvEntry.YFilter }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetFilter(yf yfilter.YFilter) { lldpUnknownTlvEntry.YFilter = yf }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetSegmentPath() string {
    return "lldp-unknown-tlv-entry"
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = lldpUnknownTlvEntry.TlvType
    leafs["tlv-value"] = lldpUnknownTlvEntry.TlvValue
    return leafs
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetYangName() string { return "lldp-unknown-tlv-entry" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetParent(parent types.Entity) { lldpUnknownTlvEntry.parent = parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParent() types.Entity { return lldpUnknownTlvEntry.parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParentYangName() string { return "unknown-tlv-list" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetFilter() yfilter.YFilter { return orgDefTlvList.YFilter }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) SetFilter(yf yfilter.YFilter) { orgDefTlvList.YFilter = yf }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetGoName(yname string) string {
    if yname == "lldp-org-def-tlv-entry" { return "LldpOrgDefTlvEntry" }
    return ""
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetSegmentPath() string {
    return "org-def-tlv-list"
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-org-def-tlv-entry" {
        for _, c := range orgDefTlvList.LldpOrgDefTlvEntry {
            if orgDefTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry{}
        orgDefTlvList.LldpOrgDefTlvEntry = append(orgDefTlvList.LldpOrgDefTlvEntry, child)
        return &orgDefTlvList.LldpOrgDefTlvEntry[len(orgDefTlvList.LldpOrgDefTlvEntry)-1]
    }
    return nil
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        children[orgDefTlvList.LldpOrgDefTlvEntry[i].GetSegmentPath()] = &orgDefTlvList.LldpOrgDefTlvEntry[i]
    }
    return children
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetYangName() string { return "org-def-tlv-list" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) SetParent(parent types.Entity) { orgDefTlvList.parent = parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetParent() types.Entity { return orgDefTlvList.parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetFilter() yfilter.YFilter { return lldpOrgDefTlvEntry.YFilter }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetFilter(yf yfilter.YFilter) { lldpOrgDefTlvEntry.YFilter = yf }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "tlv-subtype" { return "TlvSubtype" }
    if yname == "tlv-info-indes" { return "TlvInfoIndes" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetSegmentPath() string {
    return "lldp-org-def-tlv-entry"
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = lldpOrgDefTlvEntry.Oui
    leafs["tlv-subtype"] = lldpOrgDefTlvEntry.TlvSubtype
    leafs["tlv-info-indes"] = lldpOrgDefTlvEntry.TlvInfoIndes
    leafs["tlv-value"] = lldpOrgDefTlvEntry.TlvValue
    return leafs
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetYangName() string { return "lldp-org-def-tlv-entry" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetParent(parent types.Entity) { lldpOrgDefTlvEntry.parent = parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParent() types.Entity { return lldpOrgDefTlvEntry.parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParentYangName() string { return "org-def-tlv-list" }

// Lldp_Nodes_Node_Neighbors_Details
// The detailed LLDP neighbor table
type Lldp_Nodes_Node_Neighbors_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail.
    Detail []Lldp_Nodes_Node_Neighbors_Details_Detail
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Lldp_Nodes_Node_Neighbors_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetGoName(yname string) string {
    if yname == "detail" { return "Detail" }
    return ""
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetSegmentPath() string {
    return "details"
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        for _, c := range details.Detail {
            if details.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Details_Detail{}
        details.Detail = append(details.Detail, child)
        return &details.Detail[len(details.Detail)-1]
    }
    return nil
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range details.Detail {
        children[details.Detail[i].GetSegmentPath()] = &details.Detail[i]
    }
    return children
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetYangName() string { return "details" }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Lldp_Nodes_Node_Neighbors_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetParent() types.Entity { return details.parent }

func (details *Lldp_Nodes_Node_Neighbors_Details) GetParentYangName() string { return "neighbors" }

// Lldp_Nodes_Node_Neighbors_Details_Detail
// Detailed information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Details_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor.
    LldpNeighbor []Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "lldp-neighbor" { return "LldpNeighbor" }
    return ""
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-neighbor" {
        for _, c := range detail.LldpNeighbor {
            if detail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor{}
        detail.LldpNeighbor = append(detail.LldpNeighbor, child)
        return &detail.LldpNeighbor[len(detail.LldpNeighbor)-1]
    }
    return nil
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detail.LldpNeighbor {
        children[detail.LldpNeighbor[i].GetSegmentPath()] = &detail.LldpNeighbor[i]
    }
    return children
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = detail.InterfaceName
    leafs["device-id"] = detail.DeviceId
    return leafs
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetYangName() string { return "detail" }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetParentYangName() string { return "details" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetFilter() yfilter.YFilter { return lldpNeighbor.YFilter }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) SetFilter(yf yfilter.YFilter) { lldpNeighbor.YFilter = yf }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "receiving-parent-interface-name" { return "ReceivingParentInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "port-id-detail" { return "PortIdDetail" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    if yname == "mib" { return "Mib" }
    return ""
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetSegmentPath() string {
    return "lldp-neighbor"
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &lldpNeighbor.Detail
    }
    if childYangName == "mib" {
        return &lldpNeighbor.Mib
    }
    return nil
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &lldpNeighbor.Detail
    children["mib"] = &lldpNeighbor.Mib
    return children
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = lldpNeighbor.ReceivingInterfaceName
    leafs["receiving-parent-interface-name"] = lldpNeighbor.ReceivingParentInterfaceName
    leafs["device-id"] = lldpNeighbor.DeviceId
    leafs["chassis-id"] = lldpNeighbor.ChassisId
    leafs["port-id-detail"] = lldpNeighbor.PortIdDetail
    leafs["header-version"] = lldpNeighbor.HeaderVersion
    leafs["hold-time"] = lldpNeighbor.HoldTime
    leafs["enabled-capabilities"] = lldpNeighbor.EnabledCapabilities
    leafs["platform"] = lldpNeighbor.Platform
    return leafs
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetYangName() string { return "lldp-neighbor" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) SetParent(parent types.Entity) { lldpNeighbor.parent = parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetParent() types.Entity { return lldpNeighbor.parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetParentYangName() string { return "detail" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "peer-mac-address" { return "PeerMacAddress" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "system-name" { return "SystemName" }
    if yname == "system-description" { return "SystemDescription" }
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "auto-negotiation" { return "AutoNegotiation" }
    if yname == "physical-media-capabilities" { return "PhysicalMediaCapabilities" }
    if yname == "media-attachment-unit-type" { return "MediaAttachmentUnitType" }
    if yname == "port-vlan-id" { return "PortVlanId" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    return ""
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    return nil
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    return children
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-mac-address"] = detail.PeerMacAddress
    leafs["port-description"] = detail.PortDescription
    leafs["system-name"] = detail.SystemName
    leafs["system-description"] = detail.SystemDescription
    leafs["time-remaining"] = detail.TimeRemaining
    leafs["system-capabilities"] = detail.SystemCapabilities
    leafs["enabled-capabilities"] = detail.EnabledCapabilities
    leafs["auto-negotiation"] = detail.AutoNegotiation
    leafs["physical-media-capabilities"] = detail.PhysicalMediaCapabilities
    leafs["media-attachment-unit-type"] = detail.MediaAttachmentUnitType
    leafs["port-vlan-id"] = detail.PortVlanId
    return leafs
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "lldp-addr-entry" { return "LldpAddrEntry" }
    return ""
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-addr-entry" {
        for _, c := range networkAddresses.LldpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry{}
        networkAddresses.LldpAddrEntry = append(networkAddresses.LldpAddrEntry, child)
        return &networkAddresses.LldpAddrEntry[len(networkAddresses.LldpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.LldpAddrEntry {
        children[networkAddresses.LldpAddrEntry[i].GetSegmentPath()] = &networkAddresses.LldpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetFilter() yfilter.YFilter { return lldpAddrEntry.YFilter }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetFilter(yf yfilter.YFilter) { lldpAddrEntry.YFilter = yf }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetGoName(yname string) string {
    if yname == "ma-subtype" { return "MaSubtype" }
    if yname == "if-num" { return "IfNum" }
    if yname == "address" { return "Address" }
    return ""
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetSegmentPath() string {
    return "lldp-addr-entry"
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &lldpAddrEntry.Address
    }
    return nil
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &lldpAddrEntry.Address
    return children
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ma-subtype"] = lldpAddrEntry.MaSubtype
    leafs["if-num"] = lldpAddrEntry.IfNum
    return leafs
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetYangName() string { return "lldp-addr-entry" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetParent(parent types.Entity) { lldpAddrEntry.parent = parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParent() types.Entity { return lldpAddrEntry.parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParentYangName() string { return "lldp-addr-entry" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetFilter() yfilter.YFilter { return mib.YFilter }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) SetFilter(yf yfilter.YFilter) { mib.YFilter = yf }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetGoName(yname string) string {
    if yname == "rem-time-mark" { return "RemTimeMark" }
    if yname == "rem-local-port-num" { return "RemLocalPortNum" }
    if yname == "rem-index" { return "RemIndex" }
    if yname == "chassis-id-sub-type" { return "ChassisIdSubType" }
    if yname == "chassis-id-len" { return "ChassisIdLen" }
    if yname == "port-id-sub-type" { return "PortIdSubType" }
    if yname == "port-id-len" { return "PortIdLen" }
    if yname == "combined-capabilities" { return "CombinedCapabilities" }
    if yname == "unknown-tlv-list" { return "UnknownTlvList" }
    if yname == "org-def-tlv-list" { return "OrgDefTlvList" }
    return ""
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetSegmentPath() string {
    return "mib"
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv-list" {
        return &mib.UnknownTlvList
    }
    if childYangName == "org-def-tlv-list" {
        return &mib.OrgDefTlvList
    }
    return nil
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unknown-tlv-list"] = &mib.UnknownTlvList
    children["org-def-tlv-list"] = &mib.OrgDefTlvList
    return children
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rem-time-mark"] = mib.RemTimeMark
    leafs["rem-local-port-num"] = mib.RemLocalPortNum
    leafs["rem-index"] = mib.RemIndex
    leafs["chassis-id-sub-type"] = mib.ChassisIdSubType
    leafs["chassis-id-len"] = mib.ChassisIdLen
    leafs["port-id-sub-type"] = mib.PortIdSubType
    leafs["port-id-len"] = mib.PortIdLen
    leafs["combined-capabilities"] = mib.CombinedCapabilities
    return leafs
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetBundleName() string { return "cisco_ios_xr" }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetYangName() string { return "mib" }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) SetParent(parent types.Entity) { mib.parent = parent }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetParent() types.Entity { return mib.parent }

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetFilter() yfilter.YFilter { return unknownTlvList.YFilter }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) SetFilter(yf yfilter.YFilter) { unknownTlvList.YFilter = yf }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetGoName(yname string) string {
    if yname == "lldp-unknown-tlv-entry" { return "LldpUnknownTlvEntry" }
    return ""
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetSegmentPath() string {
    return "unknown-tlv-list"
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-unknown-tlv-entry" {
        for _, c := range unknownTlvList.LldpUnknownTlvEntry {
            if unknownTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry{}
        unknownTlvList.LldpUnknownTlvEntry = append(unknownTlvList.LldpUnknownTlvEntry, child)
        return &unknownTlvList.LldpUnknownTlvEntry[len(unknownTlvList.LldpUnknownTlvEntry)-1]
    }
    return nil
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        children[unknownTlvList.LldpUnknownTlvEntry[i].GetSegmentPath()] = &unknownTlvList.LldpUnknownTlvEntry[i]
    }
    return children
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetYangName() string { return "unknown-tlv-list" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) SetParent(parent types.Entity) { unknownTlvList.parent = parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetParent() types.Entity { return unknownTlvList.parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetFilter() yfilter.YFilter { return lldpUnknownTlvEntry.YFilter }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetFilter(yf yfilter.YFilter) { lldpUnknownTlvEntry.YFilter = yf }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetSegmentPath() string {
    return "lldp-unknown-tlv-entry"
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = lldpUnknownTlvEntry.TlvType
    leafs["tlv-value"] = lldpUnknownTlvEntry.TlvValue
    return leafs
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetYangName() string { return "lldp-unknown-tlv-entry" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetParent(parent types.Entity) { lldpUnknownTlvEntry.parent = parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParent() types.Entity { return lldpUnknownTlvEntry.parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParentYangName() string { return "unknown-tlv-list" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetFilter() yfilter.YFilter { return orgDefTlvList.YFilter }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) SetFilter(yf yfilter.YFilter) { orgDefTlvList.YFilter = yf }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetGoName(yname string) string {
    if yname == "lldp-org-def-tlv-entry" { return "LldpOrgDefTlvEntry" }
    return ""
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetSegmentPath() string {
    return "org-def-tlv-list"
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-org-def-tlv-entry" {
        for _, c := range orgDefTlvList.LldpOrgDefTlvEntry {
            if orgDefTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry{}
        orgDefTlvList.LldpOrgDefTlvEntry = append(orgDefTlvList.LldpOrgDefTlvEntry, child)
        return &orgDefTlvList.LldpOrgDefTlvEntry[len(orgDefTlvList.LldpOrgDefTlvEntry)-1]
    }
    return nil
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        children[orgDefTlvList.LldpOrgDefTlvEntry[i].GetSegmentPath()] = &orgDefTlvList.LldpOrgDefTlvEntry[i]
    }
    return children
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetYangName() string { return "org-def-tlv-list" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) SetParent(parent types.Entity) { orgDefTlvList.parent = parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetParent() types.Entity { return orgDefTlvList.parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetFilter() yfilter.YFilter { return lldpOrgDefTlvEntry.YFilter }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetFilter(yf yfilter.YFilter) { lldpOrgDefTlvEntry.YFilter = yf }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "tlv-subtype" { return "TlvSubtype" }
    if yname == "tlv-info-indes" { return "TlvInfoIndes" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetSegmentPath() string {
    return "lldp-org-def-tlv-entry"
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = lldpOrgDefTlvEntry.Oui
    leafs["tlv-subtype"] = lldpOrgDefTlvEntry.TlvSubtype
    leafs["tlv-info-indes"] = lldpOrgDefTlvEntry.TlvInfoIndes
    leafs["tlv-value"] = lldpOrgDefTlvEntry.TlvValue
    return leafs
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetYangName() string { return "lldp-org-def-tlv-entry" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetParent(parent types.Entity) { lldpOrgDefTlvEntry.parent = parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParent() types.Entity { return lldpOrgDefTlvEntry.parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParentYangName() string { return "org-def-tlv-list" }

// Lldp_Nodes_Node_Neighbors_Summaries
// The LLDP neighbor summary table
type Lldp_Nodes_Node_Neighbors_Summaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary.
    Summary []Lldp_Nodes_Node_Neighbors_Summaries_Summary
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetFilter() yfilter.YFilter { return summaries.YFilter }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) SetFilter(yf yfilter.YFilter) { summaries.YFilter = yf }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    return ""
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetSegmentPath() string {
    return "summaries"
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        for _, c := range summaries.Summary {
            if summaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Summaries_Summary{}
        summaries.Summary = append(summaries.Summary, child)
        return &summaries.Summary[len(summaries.Summary)-1]
    }
    return nil
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summaries.Summary {
        children[summaries.Summary[i].GetSegmentPath()] = &summaries.Summary[i]
    }
    return children
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetBundleName() string { return "cisco_ios_xr" }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetYangName() string { return "summaries" }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) SetParent(parent types.Entity) { summaries.parent = parent }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetParent() types.Entity { return summaries.parent }

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetParentYangName() string { return "neighbors" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary
// Brief information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor.
    LldpNeighbor []Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "lldp-neighbor" { return "LldpNeighbor" }
    return ""
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-neighbor" {
        for _, c := range summary.LldpNeighbor {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor{}
        summary.LldpNeighbor = append(summary.LldpNeighbor, child)
        return &summary.LldpNeighbor[len(summary.LldpNeighbor)-1]
    }
    return nil
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.LldpNeighbor {
        children[summary.LldpNeighbor[i].GetSegmentPath()] = &summary.LldpNeighbor[i]
    }
    return children
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = summary.InterfaceName
    leafs["device-id"] = summary.DeviceId
    return leafs
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetYangName() string { return "summary" }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetParentYangName() string { return "summaries" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetFilter() yfilter.YFilter { return lldpNeighbor.YFilter }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) SetFilter(yf yfilter.YFilter) { lldpNeighbor.YFilter = yf }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetGoName(yname string) string {
    if yname == "receiving-interface-name" { return "ReceivingInterfaceName" }
    if yname == "receiving-parent-interface-name" { return "ReceivingParentInterfaceName" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "port-id-detail" { return "PortIdDetail" }
    if yname == "header-version" { return "HeaderVersion" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "platform" { return "Platform" }
    if yname == "detail" { return "Detail" }
    if yname == "mib" { return "Mib" }
    return ""
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetSegmentPath() string {
    return "lldp-neighbor"
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail" {
        return &lldpNeighbor.Detail
    }
    if childYangName == "mib" {
        return &lldpNeighbor.Mib
    }
    return nil
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail"] = &lldpNeighbor.Detail
    children["mib"] = &lldpNeighbor.Mib
    return children
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["receiving-interface-name"] = lldpNeighbor.ReceivingInterfaceName
    leafs["receiving-parent-interface-name"] = lldpNeighbor.ReceivingParentInterfaceName
    leafs["device-id"] = lldpNeighbor.DeviceId
    leafs["chassis-id"] = lldpNeighbor.ChassisId
    leafs["port-id-detail"] = lldpNeighbor.PortIdDetail
    leafs["header-version"] = lldpNeighbor.HeaderVersion
    leafs["hold-time"] = lldpNeighbor.HoldTime
    leafs["enabled-capabilities"] = lldpNeighbor.EnabledCapabilities
    leafs["platform"] = lldpNeighbor.Platform
    return leafs
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetBundleName() string { return "cisco_ios_xr" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetYangName() string { return "lldp-neighbor" }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) SetParent(parent types.Entity) { lldpNeighbor.parent = parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetParent() types.Entity { return lldpNeighbor.parent }

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetParentYangName() string { return "summary" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetGoName(yname string) string {
    if yname == "peer-mac-address" { return "PeerMacAddress" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "system-name" { return "SystemName" }
    if yname == "system-description" { return "SystemDescription" }
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "enabled-capabilities" { return "EnabledCapabilities" }
    if yname == "auto-negotiation" { return "AutoNegotiation" }
    if yname == "physical-media-capabilities" { return "PhysicalMediaCapabilities" }
    if yname == "media-attachment-unit-type" { return "MediaAttachmentUnitType" }
    if yname == "port-vlan-id" { return "PortVlanId" }
    if yname == "network-addresses" { return "NetworkAddresses" }
    return ""
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-addresses" {
        return &detail.NetworkAddresses
    }
    return nil
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["network-addresses"] = &detail.NetworkAddresses
    return children
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-mac-address"] = detail.PeerMacAddress
    leafs["port-description"] = detail.PortDescription
    leafs["system-name"] = detail.SystemName
    leafs["system-description"] = detail.SystemDescription
    leafs["time-remaining"] = detail.TimeRemaining
    leafs["system-capabilities"] = detail.SystemCapabilities
    leafs["enabled-capabilities"] = detail.EnabledCapabilities
    leafs["auto-negotiation"] = detail.AutoNegotiation
    leafs["physical-media-capabilities"] = detail.PhysicalMediaCapabilities
    leafs["media-attachment-unit-type"] = detail.MediaAttachmentUnitType
    leafs["port-vlan-id"] = detail.PortVlanId
    return leafs
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetYangName() string { return "detail" }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetFilter() yfilter.YFilter { return networkAddresses.YFilter }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) SetFilter(yf yfilter.YFilter) { networkAddresses.YFilter = yf }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetGoName(yname string) string {
    if yname == "lldp-addr-entry" { return "LldpAddrEntry" }
    return ""
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetSegmentPath() string {
    return "network-addresses"
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-addr-entry" {
        for _, c := range networkAddresses.LldpAddrEntry {
            if networkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry{}
        networkAddresses.LldpAddrEntry = append(networkAddresses.LldpAddrEntry, child)
        return &networkAddresses.LldpAddrEntry[len(networkAddresses.LldpAddrEntry)-1]
    }
    return nil
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkAddresses.LldpAddrEntry {
        children[networkAddresses.LldpAddrEntry[i].GetSegmentPath()] = &networkAddresses.LldpAddrEntry[i]
    }
    return children
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetYangName() string { return "network-addresses" }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) SetParent(parent types.Entity) { networkAddresses.parent = parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetParent() types.Entity { return networkAddresses.parent }

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetParentYangName() string { return "detail" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetFilter() yfilter.YFilter { return lldpAddrEntry.YFilter }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetFilter(yf yfilter.YFilter) { lldpAddrEntry.YFilter = yf }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetGoName(yname string) string {
    if yname == "ma-subtype" { return "MaSubtype" }
    if yname == "if-num" { return "IfNum" }
    if yname == "address" { return "Address" }
    return ""
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetSegmentPath() string {
    return "lldp-addr-entry"
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &lldpAddrEntry.Address
    }
    return nil
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &lldpAddrEntry.Address
    return children
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ma-subtype"] = lldpAddrEntry.MaSubtype
    leafs["if-num"] = lldpAddrEntry.IfNum
    return leafs
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetYangName() string { return "lldp-addr-entry" }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) SetParent(parent types.Entity) { lldpAddrEntry.parent = parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParent() types.Entity { return lldpAddrEntry.parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetParentYangName() string { return "network-addresses" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetParentYangName() string { return "lldp-addr-entry" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetFilter() yfilter.YFilter { return mib.YFilter }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) SetFilter(yf yfilter.YFilter) { mib.YFilter = yf }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetGoName(yname string) string {
    if yname == "rem-time-mark" { return "RemTimeMark" }
    if yname == "rem-local-port-num" { return "RemLocalPortNum" }
    if yname == "rem-index" { return "RemIndex" }
    if yname == "chassis-id-sub-type" { return "ChassisIdSubType" }
    if yname == "chassis-id-len" { return "ChassisIdLen" }
    if yname == "port-id-sub-type" { return "PortIdSubType" }
    if yname == "port-id-len" { return "PortIdLen" }
    if yname == "combined-capabilities" { return "CombinedCapabilities" }
    if yname == "unknown-tlv-list" { return "UnknownTlvList" }
    if yname == "org-def-tlv-list" { return "OrgDefTlvList" }
    return ""
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetSegmentPath() string {
    return "mib"
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unknown-tlv-list" {
        return &mib.UnknownTlvList
    }
    if childYangName == "org-def-tlv-list" {
        return &mib.OrgDefTlvList
    }
    return nil
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unknown-tlv-list"] = &mib.UnknownTlvList
    children["org-def-tlv-list"] = &mib.OrgDefTlvList
    return children
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rem-time-mark"] = mib.RemTimeMark
    leafs["rem-local-port-num"] = mib.RemLocalPortNum
    leafs["rem-index"] = mib.RemIndex
    leafs["chassis-id-sub-type"] = mib.ChassisIdSubType
    leafs["chassis-id-len"] = mib.ChassisIdLen
    leafs["port-id-sub-type"] = mib.PortIdSubType
    leafs["port-id-len"] = mib.PortIdLen
    leafs["combined-capabilities"] = mib.CombinedCapabilities
    return leafs
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetBundleName() string { return "cisco_ios_xr" }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetYangName() string { return "mib" }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) SetParent(parent types.Entity) { mib.parent = parent }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetParent() types.Entity { return mib.parent }

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetParentYangName() string { return "lldp-neighbor" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetFilter() yfilter.YFilter { return unknownTlvList.YFilter }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) SetFilter(yf yfilter.YFilter) { unknownTlvList.YFilter = yf }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetGoName(yname string) string {
    if yname == "lldp-unknown-tlv-entry" { return "LldpUnknownTlvEntry" }
    return ""
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetSegmentPath() string {
    return "unknown-tlv-list"
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-unknown-tlv-entry" {
        for _, c := range unknownTlvList.LldpUnknownTlvEntry {
            if unknownTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry{}
        unknownTlvList.LldpUnknownTlvEntry = append(unknownTlvList.LldpUnknownTlvEntry, child)
        return &unknownTlvList.LldpUnknownTlvEntry[len(unknownTlvList.LldpUnknownTlvEntry)-1]
    }
    return nil
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        children[unknownTlvList.LldpUnknownTlvEntry[i].GetSegmentPath()] = &unknownTlvList.LldpUnknownTlvEntry[i]
    }
    return children
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetYangName() string { return "unknown-tlv-list" }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) SetParent(parent types.Entity) { unknownTlvList.parent = parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetParent() types.Entity { return unknownTlvList.parent }

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetFilter() yfilter.YFilter { return lldpUnknownTlvEntry.YFilter }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetFilter(yf yfilter.YFilter) { lldpUnknownTlvEntry.YFilter = yf }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetGoName(yname string) string {
    if yname == "tlv-type" { return "TlvType" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetSegmentPath() string {
    return "lldp-unknown-tlv-entry"
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-type"] = lldpUnknownTlvEntry.TlvType
    leafs["tlv-value"] = lldpUnknownTlvEntry.TlvValue
    return leafs
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetYangName() string { return "lldp-unknown-tlv-entry" }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) SetParent(parent types.Entity) { lldpUnknownTlvEntry.parent = parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParent() types.Entity { return lldpUnknownTlvEntry.parent }

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetParentYangName() string { return "unknown-tlv-list" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetFilter() yfilter.YFilter { return orgDefTlvList.YFilter }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) SetFilter(yf yfilter.YFilter) { orgDefTlvList.YFilter = yf }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetGoName(yname string) string {
    if yname == "lldp-org-def-tlv-entry" { return "LldpOrgDefTlvEntry" }
    return ""
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetSegmentPath() string {
    return "org-def-tlv-list"
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-org-def-tlv-entry" {
        for _, c := range orgDefTlvList.LldpOrgDefTlvEntry {
            if orgDefTlvList.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry{}
        orgDefTlvList.LldpOrgDefTlvEntry = append(orgDefTlvList.LldpOrgDefTlvEntry, child)
        return &orgDefTlvList.LldpOrgDefTlvEntry[len(orgDefTlvList.LldpOrgDefTlvEntry)-1]
    }
    return nil
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        children[orgDefTlvList.LldpOrgDefTlvEntry[i].GetSegmentPath()] = &orgDefTlvList.LldpOrgDefTlvEntry[i]
    }
    return children
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetBundleName() string { return "cisco_ios_xr" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetYangName() string { return "org-def-tlv-list" }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) SetParent(parent types.Entity) { orgDefTlvList.parent = parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetParent() types.Entity { return orgDefTlvList.parent }

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetParentYangName() string { return "mib" }

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetFilter() yfilter.YFilter { return lldpOrgDefTlvEntry.YFilter }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetFilter(yf yfilter.YFilter) { lldpOrgDefTlvEntry.YFilter = yf }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetGoName(yname string) string {
    if yname == "oui" { return "Oui" }
    if yname == "tlv-subtype" { return "TlvSubtype" }
    if yname == "tlv-info-indes" { return "TlvInfoIndes" }
    if yname == "tlv-value" { return "TlvValue" }
    return ""
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetSegmentPath() string {
    return "lldp-org-def-tlv-entry"
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oui"] = lldpOrgDefTlvEntry.Oui
    leafs["tlv-subtype"] = lldpOrgDefTlvEntry.TlvSubtype
    leafs["tlv-info-indes"] = lldpOrgDefTlvEntry.TlvInfoIndes
    leafs["tlv-value"] = lldpOrgDefTlvEntry.TlvValue
    return leafs
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetYangName() string { return "lldp-org-def-tlv-entry" }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) SetParent(parent types.Entity) { lldpOrgDefTlvEntry.parent = parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParent() types.Entity { return lldpOrgDefTlvEntry.parent }

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetParentYangName() string { return "org-def-tlv-list" }

// Lldp_Nodes_Node_Interfaces
// The table of interfaces on which LLDP is
// running on this node
type Lldp_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for an interface on which LLDP is running. The type is
    // slice of Lldp_Nodes_Node_Interfaces_Interface.
    Interface []Lldp_Nodes_Node_Interfaces_Interface
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Lldp_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Lldp_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Lldp_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// Lldp_Nodes_Node_Interfaces_Interface
// Operational data for an interface on which
// LLDP is running
type Lldp_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // TX Enabled. The type is interface{} with range: 0..255.
    TxEnabled interface{}

    // RX Enabled. The type is interface{} with range: 0..255.
    RxEnabled interface{}

    // TX State. The type is string.
    TxState interface{}

    // RX State. The type is string.
    RxState interface{}

    // ifIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Outgoing port identifier. The type is string.
    PortId interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // Local Management Addresses.
    LocalNetworkAddresses Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Lldp_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "tx-enabled" { return "TxEnabled" }
    if yname == "rx-enabled" { return "RxEnabled" }
    if yname == "tx-state" { return "TxState" }
    if yname == "rx-state" { return "RxState" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "port-id" { return "PortId" }
    if yname == "port-id-sub-type" { return "PortIdSubType" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "local-network-addresses" { return "LocalNetworkAddresses" }
    return ""
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-network-addresses" {
        return &self.LocalNetworkAddresses
    }
    return nil
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-network-addresses"] = &self.LocalNetworkAddresses
    return children
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-name-xr"] = self.InterfaceNameXr
    leafs["tx-enabled"] = self.TxEnabled
    leafs["rx-enabled"] = self.RxEnabled
    leafs["tx-state"] = self.TxState
    leafs["rx-state"] = self.RxState
    leafs["if-index"] = self.IfIndex
    leafs["port-id"] = self.PortId
    leafs["port-id-sub-type"] = self.PortIdSubType
    leafs["port-description"] = self.PortDescription
    return leafs
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Lldp_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses
// Local Management Addresses
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetFilter() yfilter.YFilter { return localNetworkAddresses.YFilter }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) SetFilter(yf yfilter.YFilter) { localNetworkAddresses.YFilter = yf }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetGoName(yname string) string {
    if yname == "lldp-addr-entry" { return "LldpAddrEntry" }
    return ""
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetSegmentPath() string {
    return "local-network-addresses"
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lldp-addr-entry" {
        for _, c := range localNetworkAddresses.LldpAddrEntry {
            if localNetworkAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry{}
        localNetworkAddresses.LldpAddrEntry = append(localNetworkAddresses.LldpAddrEntry, child)
        return &localNetworkAddresses.LldpAddrEntry[len(localNetworkAddresses.LldpAddrEntry)-1]
    }
    return nil
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localNetworkAddresses.LldpAddrEntry {
        children[localNetworkAddresses.LldpAddrEntry[i].GetSegmentPath()] = &localNetworkAddresses.LldpAddrEntry[i]
    }
    return children
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetYangName() string { return "local-network-addresses" }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) SetParent(parent types.Entity) { localNetworkAddresses.parent = parent }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetParent() types.Entity { return localNetworkAddresses.parent }

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetParentYangName() string { return "interface" }

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetFilter() yfilter.YFilter { return lldpAddrEntry.YFilter }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) SetFilter(yf yfilter.YFilter) { lldpAddrEntry.YFilter = yf }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetGoName(yname string) string {
    if yname == "ma-subtype" { return "MaSubtype" }
    if yname == "if-num" { return "IfNum" }
    if yname == "address" { return "Address" }
    return ""
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetSegmentPath() string {
    return "lldp-addr-entry"
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &lldpAddrEntry.Address
    }
    return nil
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &lldpAddrEntry.Address
    return children
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ma-subtype"] = lldpAddrEntry.MaSubtype
    leafs["if-num"] = lldpAddrEntry.IfNum
    return leafs
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetBundleName() string { return "cisco_ios_xr" }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetYangName() string { return "lldp-addr-entry" }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) SetParent(parent types.Entity) { lldpAddrEntry.parent = parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetParent() types.Entity { return lldpAddrEntry.parent }

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetParentYangName() string { return "local-network-addresses" }

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetGoName(yname string) string {
    if yname == "address-type" { return "AddressType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetSegmentPath() string {
    return "address"
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-type"] = address.AddressType
    leafs["ipv4-address"] = address.Ipv4Address
    leafs["ipv6-address"] = address.Ipv6Address
    return leafs
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetYangName() string { return "address" }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetParent() types.Entity { return address.parent }

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetParentYangName() string { return "lldp-addr-entry" }

// Lldp_Nodes_Node_Statistics
// The LLDP traffic statistics for this node
type Lldp_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmitted packets. The type is interface{} with range: 0..4294967295.
    TransmittedPackets interface{}

    // Aged out entries. The type is interface{} with range: 0..4294967295.
    AgedOutEntries interface{}

    // Discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Bad packet received and dropped. The type is interface{} with range:
    // 0..4294967295.
    BadPackets interface{}

    // Received packets. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Discarded TLVs. The type is interface{} with range: 0..4294967295.
    DiscardedTlVs interface{}

    // Unrecognized TLVs. The type is interface{} with range: 0..4294967295.
    UnrecognizedTlVs interface{}

    // Out-of-memory conditions. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Transmission errors. The type is interface{} with range: 0..4294967295.
    EncapsulationErrors interface{}

    // Queue overflows. The type is interface{} with range: 0..4294967295.
    QueueOverflowErrors interface{}

    // Table overflows. The type is interface{} with range: 0..4294967295.
    TableOverflowErrors interface{}
}

func (statistics *Lldp_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Lldp_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Lldp_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "transmitted-packets" { return "TransmittedPackets" }
    if yname == "aged-out-entries" { return "AgedOutEntries" }
    if yname == "discarded-packets" { return "DiscardedPackets" }
    if yname == "bad-packets" { return "BadPackets" }
    if yname == "received-packets" { return "ReceivedPackets" }
    if yname == "discarded-tl-vs" { return "DiscardedTlVs" }
    if yname == "unrecognized-tl-vs" { return "UnrecognizedTlVs" }
    if yname == "out-of-memory-errors" { return "OutOfMemoryErrors" }
    if yname == "encapsulation-errors" { return "EncapsulationErrors" }
    if yname == "queue-overflow-errors" { return "QueueOverflowErrors" }
    if yname == "table-overflow-errors" { return "TableOverflowErrors" }
    return ""
}

func (statistics *Lldp_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Lldp_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Lldp_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Lldp_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmitted-packets"] = statistics.TransmittedPackets
    leafs["aged-out-entries"] = statistics.AgedOutEntries
    leafs["discarded-packets"] = statistics.DiscardedPackets
    leafs["bad-packets"] = statistics.BadPackets
    leafs["received-packets"] = statistics.ReceivedPackets
    leafs["discarded-tl-vs"] = statistics.DiscardedTlVs
    leafs["unrecognized-tl-vs"] = statistics.UnrecognizedTlVs
    leafs["out-of-memory-errors"] = statistics.OutOfMemoryErrors
    leafs["encapsulation-errors"] = statistics.EncapsulationErrors
    leafs["queue-overflow-errors"] = statistics.QueueOverflowErrors
    leafs["table-overflow-errors"] = statistics.TableOverflowErrors
    return leafs
}

func (statistics *Lldp_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Lldp_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Lldp_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Lldp_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Lldp_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Lldp_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Lldp_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Lldp_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

