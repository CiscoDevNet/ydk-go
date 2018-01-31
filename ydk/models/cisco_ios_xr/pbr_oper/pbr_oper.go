// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pbr: PBR operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pbr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-oper pbr}", reflect.TypeOf(Pbr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-oper:pbr", reflect.TypeOf(Pbr{}))
}

// PolicyState represents Different Interface states
type PolicyState string

const (
    // active
    PolicyState_active PolicyState = "active"

    // suspended
    PolicyState_suspended PolicyState = "suspended"
)

// Pbr
// PBR operational data
type Pbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific PBR operational data.
    Nodes Pbr_Nodes
}

func (pbr *Pbr) GetFilter() yfilter.YFilter { return pbr.YFilter }

func (pbr *Pbr) SetFilter(yf yfilter.YFilter) { pbr.YFilter = yf }

func (pbr *Pbr) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (pbr *Pbr) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-oper:pbr"
}

func (pbr *Pbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &pbr.Nodes
    }
    return nil
}

func (pbr *Pbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &pbr.Nodes
    return children
}

func (pbr *Pbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pbr *Pbr) GetBundleName() string { return "cisco_ios_xr" }

func (pbr *Pbr) GetYangName() string { return "pbr" }

func (pbr *Pbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbr *Pbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbr *Pbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbr *Pbr) SetParent(parent types.Entity) { pbr.parent = parent }

func (pbr *Pbr) GetParent() types.Entity { return pbr.parent }

func (pbr *Pbr) GetParentYangName() string { return "Cisco-IOS-XR-pbr-oper" }

// Pbr_Nodes
// Node-specific PBR operational data
type Pbr_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBR operational data for a particular node. The type is slice of
    // Pbr_Nodes_Node.
    Node []Pbr_Nodes_Node
}

func (nodes *Pbr_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Pbr_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Pbr_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Pbr_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Pbr_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pbr_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Pbr_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Pbr_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Pbr_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Pbr_Nodes) GetYangName() string { return "nodes" }

func (nodes *Pbr_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Pbr_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Pbr_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Pbr_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Pbr_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Pbr_Nodes) GetParentYangName() string { return "pbr" }

// Pbr_Nodes_Node
// PBR operational data for a particular node
type Pbr_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Operational data for policymaps.
    PolicyMap Pbr_Nodes_Node_PolicyMap
}

func (node *Pbr_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Pbr_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Pbr_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "policy-map" { return "PolicyMap" }
    return ""
}

func (node *Pbr_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Pbr_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-map" {
        return &node.PolicyMap
    }
    return nil
}

func (node *Pbr_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-map"] = &node.PolicyMap
    return children
}

func (node *Pbr_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Pbr_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Pbr_Nodes_Node) GetYangName() string { return "node" }

func (node *Pbr_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Pbr_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Pbr_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Pbr_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Pbr_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Pbr_Nodes_Node) GetParentYangName() string { return "nodes" }

// Pbr_Nodes_Node_PolicyMap
// Operational data for policymaps
type Pbr_Nodes_Node_PolicyMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for all interfaces.
    Interfaces Pbr_Nodes_Node_PolicyMap_Interfaces
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetFilter() yfilter.YFilter { return policyMap.YFilter }

func (policyMap *Pbr_Nodes_Node_PolicyMap) SetFilter(yf yfilter.YFilter) { policyMap.YFilter = yf }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetSegmentPath() string {
    return "policy-map"
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &policyMap.Interfaces
    }
    return nil
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &policyMap.Interfaces
    return children
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetBundleName() string { return "cisco_ios_xr" }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetYangName() string { return "policy-map" }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMap *Pbr_Nodes_Node_PolicyMap) SetParent(parent types.Entity) { policyMap.parent = parent }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetParent() types.Entity { return policyMap.parent }

func (policyMap *Pbr_Nodes_Node_PolicyMap) GetParentYangName() string { return "node" }

// Pbr_Nodes_Node_PolicyMap_Interfaces
// Operational data for all interfaces
type Pbr_Nodes_Node_PolicyMap_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBR action data for a particular interface. The type is slice of
    // Pbr_Nodes_Node_PolicyMap_Interfaces_Interface.
    Interface []Pbr_Nodes_Node_PolicyMap_Interfaces_Interface
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pbr_Nodes_Node_PolicyMap_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pbr_Nodes_Node_PolicyMap_Interfaces) GetParentYangName() string { return "policy-map" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface
// PBR action data for a particular interface
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // PBR direction.
    Direction Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "direction" { return "Direction" }
    return ""
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "direction" {
        return &self.Direction
    }
    return nil
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["direction"] = &self.Direction
    return children
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction
// PBR direction
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBR policy statistics.
    Input Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetFilter() yfilter.YFilter { return direction.YFilter }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) SetFilter(yf yfilter.YFilter) { direction.YFilter = yf }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetSegmentPath() string {
    return "direction"
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &direction.Input
    }
    return nil
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &direction.Input
    return children
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetBundleName() string { return "cisco_ios_xr" }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetYangName() string { return "direction" }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) SetParent(parent types.Entity) { direction.parent = parent }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetParent() types.Entity { return direction.parent }

func (direction *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction) GetParentYangName() string { return "interface" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input
// PBR policy statistics
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeName. The type is string with length: 0..42.
    NodeName interface{}

    // PolicyName. The type is string with length: 0..65.
    PolicyName interface{}

    // State. The type is PolicyState.
    State interface{}

    // StateDescription. The type is string with length: 0..128.
    StateDescription interface{}

    // Array of classes contained in policy. The type is slice of
    // Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat.
    ClassStat []Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "state" { return "State" }
    if yname == "state-description" { return "StateDescription" }
    if yname == "class-stat" { return "ClassStat" }
    return ""
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetSegmentPath() string {
    return "input"
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-stat" {
        for _, c := range input.ClassStat {
            if input.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat{}
        input.ClassStat = append(input.ClassStat, child)
        return &input.ClassStat[len(input.ClassStat)-1]
    }
    return nil
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range input.ClassStat {
        children[input.ClassStat[i].GetSegmentPath()] = &input.ClassStat[i]
    }
    return children
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = input.NodeName
    leafs["policy-name"] = input.PolicyName
    leafs["state"] = input.State
    leafs["state-description"] = input.StateDescription
    return leafs
}

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetYangName() string { return "input" }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetParent() types.Entity { return input.parent }

func (input *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input) GetParentYangName() string { return "direction" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat
// Array of classes contained in policy
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bitmask to indicate which counter or counters are undetermined. Counters
    // will be marked undetermined when one or more classes share queues with
    // class-default because in such cases the value of counters for each class is
    // invalid. Based on the flag(s) set, the following counters will be marked
    // undetermined. For example, if value of this object returned is 0x00000101,
    // counters TransmitPackets/TransmitBytes/TotalTransmitRate and
    // DropPackets/DropBytes are undetermined .0x00000001 - Transmit
    // (TransmitPackets/TransmitBytes/TotalTransmitRate ), 0x00000002 - Drop
    // (TotalDropPackets/TotalDropBytes/TotalDropRate), 0x00000004 - Httpr
    // (HttprTransmitPackets/HttprTransmitBytes), . The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    CounterValidityBitmask interface{}

    // ClassName. The type is string with length: 0..65.
    ClassName interface{}

    // ClassId. The type is interface{} with range: 0..4294967295.
    ClassId interface{}

    // general stats.
    GeneralStats Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats

    // HTTPR stats.
    HttprStats Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetFilter() yfilter.YFilter { return classStat.YFilter }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) SetFilter(yf yfilter.YFilter) { classStat.YFilter = yf }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetGoName(yname string) string {
    if yname == "counter-validity-bitmask" { return "CounterValidityBitmask" }
    if yname == "class-name" { return "ClassName" }
    if yname == "class-id" { return "ClassId" }
    if yname == "general-stats" { return "GeneralStats" }
    if yname == "httpr-stats" { return "HttprStats" }
    return ""
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetSegmentPath() string {
    return "class-stat"
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "general-stats" {
        return &classStat.GeneralStats
    }
    if childYangName == "httpr-stats" {
        return &classStat.HttprStats
    }
    return nil
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["general-stats"] = &classStat.GeneralStats
    children["httpr-stats"] = &classStat.HttprStats
    return children
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["counter-validity-bitmask"] = classStat.CounterValidityBitmask
    leafs["class-name"] = classStat.ClassName
    leafs["class-id"] = classStat.ClassId
    return leafs
}

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetBundleName() string { return "cisco_ios_xr" }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetYangName() string { return "class-stat" }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) SetParent(parent types.Entity) { classStat.parent = parent }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetParent() types.Entity { return classStat.parent }

func (classStat *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat) GetParentYangName() string { return "input" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats
// general stats
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Transmitted packets (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TransmitPackets interface{}

    // Transmitted bytes (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TransmitBytes interface{}

    // Dropped packets (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalDropPackets interface{}

    // Dropped bytes (packets/bytes). The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TotalDropBytes interface{}

    // Total drop rate (packets/bytes). The type is interface{} with range:
    // 0..4294967295. Units are byte.
    TotalDropRate interface{}

    // Incoming matched data rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    MatchDataRate interface{}

    // Total transmit rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    TotalTransmitRate interface{}

    // Matched pkts before applying policy. The type is interface{} with range:
    // 0..18446744073709551615.
    PrePolicyMatchedPackets interface{}

    // Matched bytes before applying policy. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    PrePolicyMatchedBytes interface{}
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetFilter() yfilter.YFilter { return generalStats.YFilter }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) SetFilter(yf yfilter.YFilter) { generalStats.YFilter = yf }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetGoName(yname string) string {
    if yname == "transmit-packets" { return "TransmitPackets" }
    if yname == "transmit-bytes" { return "TransmitBytes" }
    if yname == "total-drop-packets" { return "TotalDropPackets" }
    if yname == "total-drop-bytes" { return "TotalDropBytes" }
    if yname == "total-drop-rate" { return "TotalDropRate" }
    if yname == "match-data-rate" { return "MatchDataRate" }
    if yname == "total-transmit-rate" { return "TotalTransmitRate" }
    if yname == "pre-policy-matched-packets" { return "PrePolicyMatchedPackets" }
    if yname == "pre-policy-matched-bytes" { return "PrePolicyMatchedBytes" }
    return ""
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetSegmentPath() string {
    return "general-stats"
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets"] = generalStats.TransmitPackets
    leafs["transmit-bytes"] = generalStats.TransmitBytes
    leafs["total-drop-packets"] = generalStats.TotalDropPackets
    leafs["total-drop-bytes"] = generalStats.TotalDropBytes
    leafs["total-drop-rate"] = generalStats.TotalDropRate
    leafs["match-data-rate"] = generalStats.MatchDataRate
    leafs["total-transmit-rate"] = generalStats.TotalTransmitRate
    leafs["pre-policy-matched-packets"] = generalStats.PrePolicyMatchedPackets
    leafs["pre-policy-matched-bytes"] = generalStats.PrePolicyMatchedBytes
    return leafs
}

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetBundleName() string { return "cisco_ios_xr" }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetYangName() string { return "general-stats" }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) SetParent(parent types.Entity) { generalStats.parent = parent }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetParent() types.Entity { return generalStats.parent }

func (generalStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_GeneralStats) GetParentYangName() string { return "class-stat" }

// Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats
// HTTPR stats
type Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TotalNum of pkts HTTP request received. The type is interface{} with range:
    // 0..18446744073709551615.
    RqstRcvdPackets interface{}

    // TotalNum of Bytes HTTP request received. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    RqstRcvdBytes interface{}

    // Dropped  packets. The type is interface{} with range:
    // 0..18446744073709551615.
    DropPackets interface{}

    // Dropped bytes. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    DropBytes interface{}

    // TotalNum of pkts HTTPR response sent. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSentPackets interface{}

    // TotalNum of Bytes HTTPR response sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    RespSentBytes interface{}
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetFilter() yfilter.YFilter { return httprStats.YFilter }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) SetFilter(yf yfilter.YFilter) { httprStats.YFilter = yf }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetGoName(yname string) string {
    if yname == "rqst-rcvd-packets" { return "RqstRcvdPackets" }
    if yname == "rqst-rcvd-bytes" { return "RqstRcvdBytes" }
    if yname == "drop-packets" { return "DropPackets" }
    if yname == "drop-bytes" { return "DropBytes" }
    if yname == "resp-sent-packets" { return "RespSentPackets" }
    if yname == "resp-sent-bytes" { return "RespSentBytes" }
    return ""
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetSegmentPath() string {
    return "httpr-stats"
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rqst-rcvd-packets"] = httprStats.RqstRcvdPackets
    leafs["rqst-rcvd-bytes"] = httprStats.RqstRcvdBytes
    leafs["drop-packets"] = httprStats.DropPackets
    leafs["drop-bytes"] = httprStats.DropBytes
    leafs["resp-sent-packets"] = httprStats.RespSentPackets
    leafs["resp-sent-bytes"] = httprStats.RespSentBytes
    return leafs
}

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetBundleName() string { return "cisco_ios_xr" }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetYangName() string { return "httpr-stats" }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) SetParent(parent types.Entity) { httprStats.parent = parent }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetParent() types.Entity { return httprStats.parent }

func (httprStats *Pbr_Nodes_Node_PolicyMap_Interfaces_Interface_Direction_Input_ClassStat_HttprStats) GetParentYangName() string { return "class-stat" }

