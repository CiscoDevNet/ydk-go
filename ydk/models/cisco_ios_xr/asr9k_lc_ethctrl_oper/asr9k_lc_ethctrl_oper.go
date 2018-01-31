// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-ethctrl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mlan: Management LAN Operational data space
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lc_ethctrl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lc_ethctrl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lc-ethctrl-oper mlan}", reflect.TypeOf(Mlan{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lc-ethctrl-oper:mlan", reflect.TypeOf(Mlan{}))
}

// Mlan
// Management LAN Operational data space
type Mlan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Mlan_Nodes
}

func (mlan *Mlan) GetFilter() yfilter.YFilter { return mlan.YFilter }

func (mlan *Mlan) SetFilter(yf yfilter.YFilter) { mlan.YFilter = yf }

func (mlan *Mlan) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mlan *Mlan) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lc-ethctrl-oper:mlan"
}

func (mlan *Mlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mlan.Nodes
    }
    return nil
}

func (mlan *Mlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mlan.Nodes
    return children
}

func (mlan *Mlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mlan *Mlan) GetBundleName() string { return "cisco_ios_xr" }

func (mlan *Mlan) GetYangName() string { return "mlan" }

func (mlan *Mlan) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mlan *Mlan) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mlan *Mlan) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mlan *Mlan) SetParent(parent types.Entity) { mlan.parent = parent }

func (mlan *Mlan) GetParent() types.Entity { return mlan.parent }

func (mlan *Mlan) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lc-ethctrl-oper" }

// Mlan_Nodes
// Table of nodes
type Mlan_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of Mlan_Nodes_Node.
    Node []Mlan_Nodes_Node
}

func (nodes *Mlan_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mlan_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mlan_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mlan_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mlan_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mlan_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mlan_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mlan_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mlan_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mlan_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mlan_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mlan_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mlan_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mlan_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mlan_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mlan_Nodes) GetParentYangName() string { return "mlan" }

// Mlan_Nodes_Node
// Number
type Mlan_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Table of port status.
    PortStatusNumbers Mlan_Nodes_Node_PortStatusNumbers

    // Table of switch status.
    SwitchStatusTable Mlan_Nodes_Node_SwitchStatusTable

    // Table of port counters.
    PortCountersNumbers Mlan_Nodes_Node_PortCountersNumbers

    // Table of switch ATU.
    AtuEntryNumbers Mlan_Nodes_Node_AtuEntryNumbers
}

func (node *Mlan_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mlan_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mlan_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port-status-numbers" { return "PortStatusNumbers" }
    if yname == "switch-status-table" { return "SwitchStatusTable" }
    if yname == "port-counters-numbers" { return "PortCountersNumbers" }
    if yname == "atu-entry-numbers" { return "AtuEntryNumbers" }
    return ""
}

func (node *Mlan_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *Mlan_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-status-numbers" {
        return &node.PortStatusNumbers
    }
    if childYangName == "switch-status-table" {
        return &node.SwitchStatusTable
    }
    if childYangName == "port-counters-numbers" {
        return &node.PortCountersNumbers
    }
    if childYangName == "atu-entry-numbers" {
        return &node.AtuEntryNumbers
    }
    return nil
}

func (node *Mlan_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-status-numbers"] = &node.PortStatusNumbers
    children["switch-status-table"] = &node.SwitchStatusTable
    children["port-counters-numbers"] = &node.PortCountersNumbers
    children["atu-entry-numbers"] = &node.AtuEntryNumbers
    return children
}

func (node *Mlan_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *Mlan_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mlan_Nodes_Node) GetYangName() string { return "node" }

func (node *Mlan_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mlan_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mlan_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mlan_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mlan_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mlan_Nodes_Node) GetParentYangName() string { return "nodes" }

// Mlan_Nodes_Node_PortStatusNumbers
// Table of port status
type Mlan_Nodes_Node_PortStatusNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber.
    PortStatusNumber []Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetFilter() yfilter.YFilter { return portStatusNumbers.YFilter }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) SetFilter(yf yfilter.YFilter) { portStatusNumbers.YFilter = yf }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetGoName(yname string) string {
    if yname == "port-status-number" { return "PortStatusNumber" }
    return ""
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetSegmentPath() string {
    return "port-status-numbers"
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-status-number" {
        for _, c := range portStatusNumbers.PortStatusNumber {
            if portStatusNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber{}
        portStatusNumbers.PortStatusNumber = append(portStatusNumbers.PortStatusNumber, child)
        return &portStatusNumbers.PortStatusNumber[len(portStatusNumbers.PortStatusNumber)-1]
    }
    return nil
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portStatusNumbers.PortStatusNumber {
        children[portStatusNumbers.PortStatusNumber[i].GetSegmentPath()] = &portStatusNumbers.PortStatusNumber[i]
    }
    return children
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetYangName() string { return "port-status-numbers" }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) SetParent(parent types.Entity) { portStatusNumbers.parent = parent }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetParent() types.Entity { return portStatusNumbers.parent }

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetParentYangName() string { return "node" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber
// Number
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. port number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // mlan port status info.
    PortStatus Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetFilter() yfilter.YFilter { return portStatusNumber.YFilter }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) SetFilter(yf yfilter.YFilter) { portStatusNumber.YFilter = yf }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "port-status" { return "PortStatus" }
    return ""
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetSegmentPath() string {
    return "port-status-number" + "[number='" + fmt.Sprintf("%v", portStatusNumber.Number) + "']"
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-status" {
        return &portStatusNumber.PortStatus
    }
    return nil
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-status"] = &portStatusNumber.PortStatus
    return children
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = portStatusNumber.Number
    return leafs
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetBundleName() string { return "cisco_ios_xr" }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetYangName() string { return "port-status-number" }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) SetParent(parent types.Entity) { portStatusNumber.parent = parent }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetParent() types.Entity { return portStatusNumber.parent }

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetParentYangName() string { return "port-status-numbers" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus
// mlan port status info
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port Number. The type is interface{} with range: 0..4294967295.
    PortNum interface{}

    // PHY data valid. The type is interface{} with range: 0..4294967295.
    PhyValid interface{}

    // SERDES data valid. The type is interface{} with range: 0..4294967295.
    SerdesValid interface{}

    // MAC data valid. The type is interface{} with range: 0..4294967295.
    MacValid interface{}

    // Configuration Data.
    Config Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config

    // PHY Registers.
    Phy Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy

    // SERDES Registers.
    Serdes Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes

    // MAC Registers.
    Mac Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetFilter() yfilter.YFilter { return portStatus.YFilter }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) SetFilter(yf yfilter.YFilter) { portStatus.YFilter = yf }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetGoName(yname string) string {
    if yname == "port-num" { return "PortNum" }
    if yname == "phy-valid" { return "PhyValid" }
    if yname == "serdes-valid" { return "SerdesValid" }
    if yname == "mac-valid" { return "MacValid" }
    if yname == "config" { return "Config" }
    if yname == "phy" { return "Phy" }
    if yname == "serdes" { return "Serdes" }
    if yname == "mac" { return "Mac" }
    return ""
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetSegmentPath() string {
    return "port-status"
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &portStatus.Config
    }
    if childYangName == "phy" {
        return &portStatus.Phy
    }
    if childYangName == "serdes" {
        return &portStatus.Serdes
    }
    if childYangName == "mac" {
        return &portStatus.Mac
    }
    return nil
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &portStatus.Config
    children["phy"] = &portStatus.Phy
    children["serdes"] = &portStatus.Serdes
    children["mac"] = &portStatus.Mac
    return children
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-num"] = portStatus.PortNum
    leafs["phy-valid"] = portStatus.PhyValid
    leafs["serdes-valid"] = portStatus.SerdesValid
    leafs["mac-valid"] = portStatus.MacValid
    return leafs
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetBundleName() string { return "cisco_ios_xr" }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetYangName() string { return "port-status" }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) SetParent(parent types.Entity) { portStatus.parent = parent }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetParent() types.Entity { return portStatus.parent }

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetParentYangName() string { return "port-status-number" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config
// Configuration Data
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // duplex. The type is interface{} with range: 0..4294967295.
    Duplex interface{}

    // pauseEn. The type is interface{} with range: 0..65535.
    Pause interface{}

    // myPause. The type is interface{} with range: 0..65535.
    MyPause interface{}

    // loopback. The type is interface{} with range: 0..4294967295.
    Loopback interface{}
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetGoName(yname string) string {
    if yname == "speed" { return "Speed" }
    if yname == "duplex" { return "Duplex" }
    if yname == "pause" { return "Pause" }
    if yname == "my-pause" { return "MyPause" }
    if yname == "loopback" { return "Loopback" }
    return ""
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["speed"] = config.Speed
    leafs["duplex"] = config.Duplex
    leafs["pause"] = config.Pause
    leafs["my-pause"] = config.MyPause
    leafs["loopback"] = config.Loopback
    return leafs
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetYangName() string { return "config" }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetParent() types.Entity { return config.parent }

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetParentYangName() string { return "port-status" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy
// PHY Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reg. The type is slice of interface{} with range: 0..65535.
    Reg []interface{}
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetFilter() yfilter.YFilter { return phy.YFilter }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) SetFilter(yf yfilter.YFilter) { phy.YFilter = yf }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetGoName(yname string) string {
    if yname == "reg" { return "Reg" }
    return ""
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetSegmentPath() string {
    return "phy"
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg"] = phy.Reg
    return leafs
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetBundleName() string { return "cisco_ios_xr" }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetYangName() string { return "phy" }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) SetParent(parent types.Entity) { phy.parent = parent }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetParent() types.Entity { return phy.parent }

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetParentYangName() string { return "port-status" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes
// SERDES Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reg. The type is slice of interface{} with range: 0..65535.
    Reg []interface{}
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetFilter() yfilter.YFilter { return serdes.YFilter }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) SetFilter(yf yfilter.YFilter) { serdes.YFilter = yf }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetGoName(yname string) string {
    if yname == "reg" { return "Reg" }
    return ""
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetSegmentPath() string {
    return "serdes"
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg"] = serdes.Reg
    return leafs
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetBundleName() string { return "cisco_ios_xr" }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetYangName() string { return "serdes" }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) SetParent(parent types.Entity) { serdes.parent = parent }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetParent() types.Entity { return serdes.parent }

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetParentYangName() string { return "port-status" }

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac
// MAC Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reg. The type is slice of interface{} with range: 0..65535.
    Reg []interface{}
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetGoName(yname string) string {
    if yname == "reg" { return "Reg" }
    return ""
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg"] = mac.Reg
    return leafs
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetYangName() string { return "mac" }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetParentYangName() string { return "port-status" }

// Mlan_Nodes_Node_SwitchStatusTable
// Table of switch status
type Mlan_Nodes_Node_SwitchStatusTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // mlan switch status info.
    SwitchStatus Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetFilter() yfilter.YFilter { return switchStatusTable.YFilter }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) SetFilter(yf yfilter.YFilter) { switchStatusTable.YFilter = yf }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetGoName(yname string) string {
    if yname == "switch-status" { return "SwitchStatus" }
    return ""
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetSegmentPath() string {
    return "switch-status-table"
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "switch-status" {
        return &switchStatusTable.SwitchStatus
    }
    return nil
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["switch-status"] = &switchStatusTable.SwitchStatus
    return children
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetBundleName() string { return "cisco_ios_xr" }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetYangName() string { return "switch-status-table" }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) SetParent(parent types.Entity) { switchStatusTable.parent = parent }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetParent() types.Entity { return switchStatusTable.parent }

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetParentYangName() string { return "node" }

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus
// mlan switch status info
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CPU Interface Rate Limit. The type is interface{} with range:
    // -2147483648..2147483647.
    RateLimit interface{}

    // Switch Global Registers.
    SwReg1 Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1

    // Switch Global Registers.
    SwReg2 Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2

    // Switch Status Data.
    SwStatus Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetFilter() yfilter.YFilter { return switchStatus.YFilter }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) SetFilter(yf yfilter.YFilter) { switchStatus.YFilter = yf }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetGoName(yname string) string {
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "sw-reg-1" { return "SwReg1" }
    if yname == "sw-reg-2" { return "SwReg2" }
    if yname == "sw-status" { return "SwStatus" }
    return ""
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetSegmentPath() string {
    return "switch-status"
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sw-reg-1" {
        return &switchStatus.SwReg1
    }
    if childYangName == "sw-reg-2" {
        return &switchStatus.SwReg2
    }
    if childYangName == "sw-status" {
        return &switchStatus.SwStatus
    }
    return nil
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sw-reg-1"] = &switchStatus.SwReg1
    children["sw-reg-2"] = &switchStatus.SwReg2
    children["sw-status"] = &switchStatus.SwStatus
    return children
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate-limit"] = switchStatus.RateLimit
    return leafs
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetBundleName() string { return "cisco_ios_xr" }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetYangName() string { return "switch-status" }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) SetParent(parent types.Entity) { switchStatus.parent = parent }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetParent() types.Entity { return switchStatus.parent }

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetParentYangName() string { return "switch-status-table" }

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1
// Switch Global Registers
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reg. The type is slice of interface{} with range: 0..65535.
    Reg []interface{}
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetFilter() yfilter.YFilter { return swReg1.YFilter }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) SetFilter(yf yfilter.YFilter) { swReg1.YFilter = yf }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetGoName(yname string) string {
    if yname == "reg" { return "Reg" }
    return ""
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetSegmentPath() string {
    return "sw-reg-1"
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg"] = swReg1.Reg
    return leafs
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetBundleName() string { return "cisco_ios_xr" }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetYangName() string { return "sw-reg-1" }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) SetParent(parent types.Entity) { swReg1.parent = parent }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetParent() types.Entity { return swReg1.parent }

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetParentYangName() string { return "switch-status" }

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2
// Switch Global Registers
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reg. The type is slice of interface{} with range: 0..65535.
    Reg []interface{}
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetFilter() yfilter.YFilter { return swReg2.YFilter }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) SetFilter(yf yfilter.YFilter) { swReg2.YFilter = yf }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetGoName(yname string) string {
    if yname == "reg" { return "Reg" }
    return ""
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetSegmentPath() string {
    return "sw-reg-2"
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg"] = swReg2.Reg
    return leafs
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetBundleName() string { return "cisco_ios_xr" }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetYangName() string { return "sw-reg-2" }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) SetParent(parent types.Entity) { swReg2.parent = parent }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetParent() types.Entity { return swReg2.parent }

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetParentYangName() string { return "switch-status" }

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus
// Switch Status Data
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ppu. The type is interface{} with range: 0..4294967295.
    Ppu interface{}

    // mtu. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // mac. The type is string with length: 0..6.
    Mac interface{}

    // cpu port. The type is interface{} with range: 0..65535.
    CpuPort interface{}

    // cpu mac. The type is interface{} with range: 0..65535.
    CpuMac interface{}

    // initialized. The type is interface{} with range: 0..65535.
    Initialized interface{}

    // restarted. The type is interface{} with range: 0..65535.
    Restarted interface{}
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetFilter() yfilter.YFilter { return swStatus.YFilter }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) SetFilter(yf yfilter.YFilter) { swStatus.YFilter = yf }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetGoName(yname string) string {
    if yname == "ppu" { return "Ppu" }
    if yname == "mtu" { return "Mtu" }
    if yname == "mac" { return "Mac" }
    if yname == "cpu-port" { return "CpuPort" }
    if yname == "cpu-mac" { return "CpuMac" }
    if yname == "initialized" { return "Initialized" }
    if yname == "restarted" { return "Restarted" }
    return ""
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetSegmentPath() string {
    return "sw-status"
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ppu"] = swStatus.Ppu
    leafs["mtu"] = swStatus.Mtu
    leafs["mac"] = swStatus.Mac
    leafs["cpu-port"] = swStatus.CpuPort
    leafs["cpu-mac"] = swStatus.CpuMac
    leafs["initialized"] = swStatus.Initialized
    leafs["restarted"] = swStatus.Restarted
    return leafs
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetBundleName() string { return "cisco_ios_xr" }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetYangName() string { return "sw-status" }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) SetParent(parent types.Entity) { swStatus.parent = parent }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetParent() types.Entity { return swStatus.parent }

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetParentYangName() string { return "switch-status" }

// Mlan_Nodes_Node_PortCountersNumbers
// Table of port counters
type Mlan_Nodes_Node_PortCountersNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber.
    PortCountersNumber []Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetFilter() yfilter.YFilter { return portCountersNumbers.YFilter }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) SetFilter(yf yfilter.YFilter) { portCountersNumbers.YFilter = yf }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetGoName(yname string) string {
    if yname == "port-counters-number" { return "PortCountersNumber" }
    return ""
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetSegmentPath() string {
    return "port-counters-numbers"
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-counters-number" {
        for _, c := range portCountersNumbers.PortCountersNumber {
            if portCountersNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber{}
        portCountersNumbers.PortCountersNumber = append(portCountersNumbers.PortCountersNumber, child)
        return &portCountersNumbers.PortCountersNumber[len(portCountersNumbers.PortCountersNumber)-1]
    }
    return nil
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portCountersNumbers.PortCountersNumber {
        children[portCountersNumbers.PortCountersNumber[i].GetSegmentPath()] = &portCountersNumbers.PortCountersNumber[i]
    }
    return children
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetYangName() string { return "port-counters-numbers" }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) SetParent(parent types.Entity) { portCountersNumbers.parent = parent }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetParent() types.Entity { return portCountersNumbers.parent }

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetParentYangName() string { return "node" }

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber
// Number
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. port number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // mlan port counters info.
    PortCounters Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetFilter() yfilter.YFilter { return portCountersNumber.YFilter }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) SetFilter(yf yfilter.YFilter) { portCountersNumber.YFilter = yf }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "port-counters" { return "PortCounters" }
    return ""
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetSegmentPath() string {
    return "port-counters-number" + "[number='" + fmt.Sprintf("%v", portCountersNumber.Number) + "']"
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-counters" {
        return &portCountersNumber.PortCounters
    }
    return nil
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port-counters"] = &portCountersNumber.PortCounters
    return children
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = portCountersNumber.Number
    return leafs
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetBundleName() string { return "cisco_ios_xr" }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetYangName() string { return "port-counters-number" }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) SetParent(parent types.Entity) { portCountersNumber.parent = parent }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetParent() types.Entity { return portCountersNumber.parent }

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetParentYangName() string { return "port-counters-numbers" }

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters
// mlan port counters info
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port Number. The type is interface{} with range: 0..4294967295.
    PortNum interface{}

    // Switch Port Statistics.
    MlanStats Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetFilter() yfilter.YFilter { return portCounters.YFilter }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) SetFilter(yf yfilter.YFilter) { portCounters.YFilter = yf }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetGoName(yname string) string {
    if yname == "port-num" { return "PortNum" }
    if yname == "mlan-stats" { return "MlanStats" }
    return ""
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetSegmentPath() string {
    return "port-counters"
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mlan-stats" {
        return &portCounters.MlanStats
    }
    return nil
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mlan-stats"] = &portCounters.MlanStats
    return children
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-num"] = portCounters.PortNum
    return leafs
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetBundleName() string { return "cisco_ios_xr" }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetYangName() string { return "port-counters" }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) SetParent(parent types.Entity) { portCounters.parent = parent }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetParent() types.Entity { return portCounters.parent }

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetParentYangName() string { return "port-counters-number" }

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats
// Switch Port Statistics
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // inGoodOctets hi. The type is interface{} with range: 0..4294967295.
    InGoodOctetsHi interface{}

    // inGoodOctets. The type is interface{} with range: 0..4294967295.
    InGoodOctets interface{}

    // inBadOctets. The type is interface{} with range: 0..4294967295.
    InBadOctets interface{}

    // inUnicastPkt. The type is interface{} with range: 0..4294967295.
    InUnicastPkt interface{}

    // inBcastPkt. The type is interface{} with range: 0..4294967295.
    InBcastPkt interface{}

    // inMcastPkt. The type is interface{} with range: 0..4294967295.
    InMcastPkt interface{}

    // inPausePkt. The type is interface{} with range: 0..4294967295.
    InPausePkt interface{}

    // inUndersizePkt. The type is interface{} with range: 0..4294967295.
    InUndersizePkt interface{}

    // inFragments. The type is interface{} with range: 0..4294967295.
    InFragments interface{}

    // inOversize. The type is interface{} with range: 0..4294967295.
    InOversize interface{}

    // inJabber. The type is interface{} with range: 0..4294967295.
    InJabber interface{}

    // inRxErr. The type is interface{} with range: 0..4294967295.
    InRxErr interface{}

    // inFcsErr. The type is interface{} with range: 0..4294967295.
    InFcsErr interface{}

    // outOctets hi. The type is interface{} with range: 0..4294967295.
    OutOctetsHi interface{}

    // outOctets. The type is interface{} with range: 0..4294967295.
    OutOctets interface{}

    // outUnicastPkt. The type is interface{} with range: 0..4294967295.
    OutUnicastPkt interface{}

    // outBcastPkt. The type is interface{} with range: 0..4294967295.
    OutBcastPkt interface{}

    // outMcastPkt. The type is interface{} with range: 0..4294967295.
    OutMcastPkt interface{}

    // outPausePkt. The type is interface{} with range: 0..4294967295.
    OutPausePkt interface{}

    // excessive. The type is interface{} with range: 0..4294967295.
    Excessive interface{}

    // collisions. The type is interface{} with range: 0..4294967295.
    Collisions interface{}

    // deferred. The type is interface{} with range: 0..4294967295.
    Deferred interface{}

    // single. The type is interface{} with range: 0..4294967295.
    Single interface{}

    // multiple. The type is interface{} with range: 0..4294967295.
    Multiple interface{}

    // outFcsErr. The type is interface{} with range: 0..4294967295.
    OutFcsErr interface{}

    // late. The type is interface{} with range: 0..4294967295.
    Late interface{}

    // rx tx 64 Octets. The type is interface{} with range: 0..4294967295.
    RxTx64Octets interface{}

    // rx tx 65 127 Octets. The type is interface{} with range: 0..4294967295.
    RxTx65127Octets interface{}

    // rx tx 128 255 Octets. The type is interface{} with range: 0..4294967295.
    RxTx128255Octets interface{}

    // rx tx 256 511 Octets. The type is interface{} with range: 0..4294967295.
    RxTx256511Octets interface{}

    // rx tx 512 1023 Octets. The type is interface{} with range: 0..4294967295.
    RxTx5121023Octets interface{}

    // rx tx 1024 Max Octets. The type is interface{} with range: 0..4294967295.
    RxTx1024MaxOctets interface{}

    // inDiscards. The type is interface{} with range: 0..4294967295.
    InDiscards interface{}

    // inFiltered. The type is interface{} with range: 0..4294967295.
    InFiltered interface{}

    // outFiltered. The type is interface{} with range: 0..4294967295.
    OutFiltered interface{}
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetFilter() yfilter.YFilter { return mlanStats.YFilter }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) SetFilter(yf yfilter.YFilter) { mlanStats.YFilter = yf }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetGoName(yname string) string {
    if yname == "in-good-octets-hi" { return "InGoodOctetsHi" }
    if yname == "in-good-octets" { return "InGoodOctets" }
    if yname == "in-bad-octets" { return "InBadOctets" }
    if yname == "in-unicast-pkt" { return "InUnicastPkt" }
    if yname == "in-bcast-pkt" { return "InBcastPkt" }
    if yname == "in-mcast-pkt" { return "InMcastPkt" }
    if yname == "in-pause-pkt" { return "InPausePkt" }
    if yname == "in-undersize-pkt" { return "InUndersizePkt" }
    if yname == "in-fragments" { return "InFragments" }
    if yname == "in-oversize" { return "InOversize" }
    if yname == "in-jabber" { return "InJabber" }
    if yname == "in-rx-err" { return "InRxErr" }
    if yname == "in-fcs-err" { return "InFcsErr" }
    if yname == "out-octets-hi" { return "OutOctetsHi" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "out-unicast-pkt" { return "OutUnicastPkt" }
    if yname == "out-bcast-pkt" { return "OutBcastPkt" }
    if yname == "out-mcast-pkt" { return "OutMcastPkt" }
    if yname == "out-pause-pkt" { return "OutPausePkt" }
    if yname == "excessive" { return "Excessive" }
    if yname == "collisions" { return "Collisions" }
    if yname == "deferred" { return "Deferred" }
    if yname == "single" { return "Single" }
    if yname == "multiple" { return "Multiple" }
    if yname == "out-fcs-err" { return "OutFcsErr" }
    if yname == "late" { return "Late" }
    if yname == "rx-tx-64-octets" { return "RxTx64Octets" }
    if yname == "rx-tx-65-127-octets" { return "RxTx65127Octets" }
    if yname == "rx-tx-128-255-octets" { return "RxTx128255Octets" }
    if yname == "rx-tx-256-511-octets" { return "RxTx256511Octets" }
    if yname == "rx-tx-512-1023-octets" { return "RxTx5121023Octets" }
    if yname == "rx-tx-1024-max-octets" { return "RxTx1024MaxOctets" }
    if yname == "in-discards" { return "InDiscards" }
    if yname == "in-filtered" { return "InFiltered" }
    if yname == "out-filtered" { return "OutFiltered" }
    return ""
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetSegmentPath() string {
    return "mlan-stats"
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-good-octets-hi"] = mlanStats.InGoodOctetsHi
    leafs["in-good-octets"] = mlanStats.InGoodOctets
    leafs["in-bad-octets"] = mlanStats.InBadOctets
    leafs["in-unicast-pkt"] = mlanStats.InUnicastPkt
    leafs["in-bcast-pkt"] = mlanStats.InBcastPkt
    leafs["in-mcast-pkt"] = mlanStats.InMcastPkt
    leafs["in-pause-pkt"] = mlanStats.InPausePkt
    leafs["in-undersize-pkt"] = mlanStats.InUndersizePkt
    leafs["in-fragments"] = mlanStats.InFragments
    leafs["in-oversize"] = mlanStats.InOversize
    leafs["in-jabber"] = mlanStats.InJabber
    leafs["in-rx-err"] = mlanStats.InRxErr
    leafs["in-fcs-err"] = mlanStats.InFcsErr
    leafs["out-octets-hi"] = mlanStats.OutOctetsHi
    leafs["out-octets"] = mlanStats.OutOctets
    leafs["out-unicast-pkt"] = mlanStats.OutUnicastPkt
    leafs["out-bcast-pkt"] = mlanStats.OutBcastPkt
    leafs["out-mcast-pkt"] = mlanStats.OutMcastPkt
    leafs["out-pause-pkt"] = mlanStats.OutPausePkt
    leafs["excessive"] = mlanStats.Excessive
    leafs["collisions"] = mlanStats.Collisions
    leafs["deferred"] = mlanStats.Deferred
    leafs["single"] = mlanStats.Single
    leafs["multiple"] = mlanStats.Multiple
    leafs["out-fcs-err"] = mlanStats.OutFcsErr
    leafs["late"] = mlanStats.Late
    leafs["rx-tx-64-octets"] = mlanStats.RxTx64Octets
    leafs["rx-tx-65-127-octets"] = mlanStats.RxTx65127Octets
    leafs["rx-tx-128-255-octets"] = mlanStats.RxTx128255Octets
    leafs["rx-tx-256-511-octets"] = mlanStats.RxTx256511Octets
    leafs["rx-tx-512-1023-octets"] = mlanStats.RxTx5121023Octets
    leafs["rx-tx-1024-max-octets"] = mlanStats.RxTx1024MaxOctets
    leafs["in-discards"] = mlanStats.InDiscards
    leafs["in-filtered"] = mlanStats.InFiltered
    leafs["out-filtered"] = mlanStats.OutFiltered
    return leafs
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetBundleName() string { return "cisco_ios_xr" }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetYangName() string { return "mlan-stats" }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) SetParent(parent types.Entity) { mlanStats.parent = parent }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetParent() types.Entity { return mlanStats.parent }

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetParentYangName() string { return "port-counters" }

// Mlan_Nodes_Node_AtuEntryNumbers
// Table of switch ATU
type Mlan_Nodes_Node_AtuEntryNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Entry number. The type is slice of
    // Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber.
    AtuEntryNumber []Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetFilter() yfilter.YFilter { return atuEntryNumbers.YFilter }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) SetFilter(yf yfilter.YFilter) { atuEntryNumbers.YFilter = yf }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetGoName(yname string) string {
    if yname == "atu-entry-number" { return "AtuEntryNumber" }
    return ""
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetSegmentPath() string {
    return "atu-entry-numbers"
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atu-entry-number" {
        for _, c := range atuEntryNumbers.AtuEntryNumber {
            if atuEntryNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber{}
        atuEntryNumbers.AtuEntryNumber = append(atuEntryNumbers.AtuEntryNumber, child)
        return &atuEntryNumbers.AtuEntryNumber[len(atuEntryNumbers.AtuEntryNumber)-1]
    }
    return nil
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atuEntryNumbers.AtuEntryNumber {
        children[atuEntryNumbers.AtuEntryNumber[i].GetSegmentPath()] = &atuEntryNumbers.AtuEntryNumber[i]
    }
    return children
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetYangName() string { return "atu-entry-numbers" }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) SetParent(parent types.Entity) { atuEntryNumbers.parent = parent }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetParent() types.Entity { return atuEntryNumbers.parent }

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetParentYangName() string { return "node" }

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber
// Entry number
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. entry number. The type is interface{} with range:
    // -2147483648..2147483647.
    Entry interface{}

    // mlan switch counters info.
    SwitchCounters Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetFilter() yfilter.YFilter { return atuEntryNumber.YFilter }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) SetFilter(yf yfilter.YFilter) { atuEntryNumber.YFilter = yf }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetGoName(yname string) string {
    if yname == "entry" { return "Entry" }
    if yname == "switch-counters" { return "SwitchCounters" }
    return ""
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetSegmentPath() string {
    return "atu-entry-number" + "[entry='" + fmt.Sprintf("%v", atuEntryNumber.Entry) + "']"
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "switch-counters" {
        return &atuEntryNumber.SwitchCounters
    }
    return nil
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["switch-counters"] = &atuEntryNumber.SwitchCounters
    return children
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry"] = atuEntryNumber.Entry
    return leafs
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetBundleName() string { return "cisco_ios_xr" }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetYangName() string { return "atu-entry-number" }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) SetParent(parent types.Entity) { atuEntryNumber.parent = parent }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetParent() types.Entity { return atuEntryNumber.parent }

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetParentYangName() string { return "atu-entry-numbers" }

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters
// mlan switch counters info
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Index of ATU Entry. The type is interface{} with range: 0..4294967295.
    EntryNum interface{}

    // Switch ATU Data.
    Atu Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetFilter() yfilter.YFilter { return switchCounters.YFilter }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) SetFilter(yf yfilter.YFilter) { switchCounters.YFilter = yf }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetGoName(yname string) string {
    if yname == "entry-num" { return "EntryNum" }
    if yname == "atu" { return "Atu" }
    return ""
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetSegmentPath() string {
    return "switch-counters"
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atu" {
        return &switchCounters.Atu
    }
    return nil
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["atu"] = &switchCounters.Atu
    return children
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry-num"] = switchCounters.EntryNum
    return leafs
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetBundleName() string { return "cisco_ios_xr" }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetYangName() string { return "switch-counters" }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) SetParent(parent types.Entity) { switchCounters.parent = parent }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetParent() types.Entity { return switchCounters.parent }

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetParentYangName() string { return "atu-entry-number" }

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu
// Switch ATU Data
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dbNum. The type is interface{} with range: 0..65535.
    DbNum interface{}

    // priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // trunk. The type is bool.
    Trunk interface{}

    // dpv. The type is interface{} with range: 0..255.
    Dpv interface{}

    // es. The type is interface{} with range: 0..255.
    Es interface{}

    // macaddr. The type is slice of interface{} with range: 0..65535.
    Macaddr []interface{}
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetFilter() yfilter.YFilter { return atu.YFilter }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) SetFilter(yf yfilter.YFilter) { atu.YFilter = yf }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetGoName(yname string) string {
    if yname == "db-num" { return "DbNum" }
    if yname == "priority" { return "Priority" }
    if yname == "trunk" { return "Trunk" }
    if yname == "dpv" { return "Dpv" }
    if yname == "es" { return "Es" }
    if yname == "macaddr" { return "Macaddr" }
    return ""
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetSegmentPath() string {
    return "atu"
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["db-num"] = atu.DbNum
    leafs["priority"] = atu.Priority
    leafs["trunk"] = atu.Trunk
    leafs["dpv"] = atu.Dpv
    leafs["es"] = atu.Es
    leafs["macaddr"] = atu.Macaddr
    return leafs
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetBundleName() string { return "cisco_ios_xr" }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetYangName() string { return "atu" }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) SetParent(parent types.Entity) { atu.parent = parent }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetParent() types.Entity { return atu.parent }

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetParentYangName() string { return "switch-counters" }

