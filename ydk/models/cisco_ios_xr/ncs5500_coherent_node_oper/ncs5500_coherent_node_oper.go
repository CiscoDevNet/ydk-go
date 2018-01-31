// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-node package operational data.
// 
// This module contains definitions
// for the following management objects:
//   coherent: Coherent node  operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5500_coherent_node_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5500_coherent_node_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs5500-coherent-node-oper coherent}", reflect.TypeOf(Coherent{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent", reflect.TypeOf(Coherent{}))
}

// Coherent
// Coherent node  operational data
type Coherent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Coherent list of nodes.
    Nodes Coherent_Nodes
}

func (coherent *Coherent) GetFilter() yfilter.YFilter { return coherent.YFilter }

func (coherent *Coherent) SetFilter(yf yfilter.YFilter) { coherent.YFilter = yf }

func (coherent *Coherent) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (coherent *Coherent) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent"
}

func (coherent *Coherent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &coherent.Nodes
    }
    return nil
}

func (coherent *Coherent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &coherent.Nodes
    return children
}

func (coherent *Coherent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (coherent *Coherent) GetBundleName() string { return "cisco_ios_xr" }

func (coherent *Coherent) GetYangName() string { return "coherent" }

func (coherent *Coherent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coherent *Coherent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coherent *Coherent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coherent *Coherent) SetParent(parent types.Entity) { coherent.parent = parent }

func (coherent *Coherent) GetParent() types.Entity { return coherent.parent }

func (coherent *Coherent) GetParentYangName() string { return "Cisco-IOS-XR-ncs5500-coherent-node-oper" }

// Coherent_Nodes
// Coherent list of nodes
type Coherent_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Coherent discovery operational data for a particular node. The type is
    // slice of Coherent_Nodes_Node.
    Node []Coherent_Nodes_Node
}

func (nodes *Coherent_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Coherent_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Coherent_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Coherent_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Coherent_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Coherent_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Coherent_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Coherent_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Coherent_Nodes) GetYangName() string { return "nodes" }

func (nodes *Coherent_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Coherent_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Coherent_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Coherent_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Coherent_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Coherent_Nodes) GetParentYangName() string { return "coherent" }

// Coherent_Nodes_Node
// Coherent discovery operational data for a
// particular node
type Coherent_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Coherent driver performace information.
    CoherentTimeStats Coherent_Nodes_Node_CoherentTimeStats

    // Coherent node data for device _mapping.
    Devicemapping Coherent_Nodes_Node_Devicemapping

    // Coherent node data for driver health.
    Coherenthealth Coherent_Nodes_Node_Coherenthealth

    // PortMode all operational data.
    PortModeAllInfo Coherent_Nodes_Node_PortModeAllInfo
}

func (node *Coherent_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Coherent_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Coherent_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "coherent-time-stats" { return "CoherentTimeStats" }
    if yname == "devicemapping" { return "Devicemapping" }
    if yname == "coherenthealth" { return "Coherenthealth" }
    if yname == "port-mode-all-info" { return "PortModeAllInfo" }
    return ""
}

func (node *Coherent_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Coherent_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "coherent-time-stats" {
        return &node.CoherentTimeStats
    }
    if childYangName == "devicemapping" {
        return &node.Devicemapping
    }
    if childYangName == "coherenthealth" {
        return &node.Coherenthealth
    }
    if childYangName == "port-mode-all-info" {
        return &node.PortModeAllInfo
    }
    return nil
}

func (node *Coherent_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["coherent-time-stats"] = &node.CoherentTimeStats
    children["devicemapping"] = &node.Devicemapping
    children["coherenthealth"] = &node.Coherenthealth
    children["port-mode-all-info"] = &node.PortModeAllInfo
    return children
}

func (node *Coherent_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Coherent_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Coherent_Nodes_Node) GetYangName() string { return "node" }

func (node *Coherent_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Coherent_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Coherent_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Coherent_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Coherent_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Coherent_Nodes_Node) GetParentYangName() string { return "nodes" }

// Coherent_Nodes_Node_CoherentTimeStats
// Coherent driver performace information
type Coherent_Nodes_Node_CoherentTimeStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // driver init. The type is string with length: 0..255.
    DriverInit interface{}

    // driver operational. The type is string with length: 0..255.
    DriverOperational interface{}

    // device created. The type is string with length: 0..255.
    DeviceCreated interface{}

    // optics controllers created. The type is string with length: 0..255.
    OpticsControllersCreated interface{}

    // dsp controllers created. The type is string with length: 0..255.
    DspControllersCreated interface{}

    // eth intf created. The type is string with length: 0..255.
    EthIntfCreated interface{}

    // opts ea bulk create.
    OptsEaBulkCreate Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate

    // opts ea bulk update.
    OptsEaBulkUpdate Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate

    // dsp ea bulk create.
    DspEaBulkCreate Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate

    // dsp ea bulk update.
    DspEaBulkUpdate Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate

    // port stat. The type is slice of
    // Coherent_Nodes_Node_CoherentTimeStats_PortStat.
    PortStat []Coherent_Nodes_Node_CoherentTimeStats_PortStat
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetFilter() yfilter.YFilter { return coherentTimeStats.YFilter }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) SetFilter(yf yfilter.YFilter) { coherentTimeStats.YFilter = yf }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetGoName(yname string) string {
    if yname == "driver-init" { return "DriverInit" }
    if yname == "driver-operational" { return "DriverOperational" }
    if yname == "device-created" { return "DeviceCreated" }
    if yname == "optics-controllers-created" { return "OpticsControllersCreated" }
    if yname == "dsp-controllers-created" { return "DspControllersCreated" }
    if yname == "eth-intf-created" { return "EthIntfCreated" }
    if yname == "opts-ea-bulk-create" { return "OptsEaBulkCreate" }
    if yname == "opts-ea-bulk-update" { return "OptsEaBulkUpdate" }
    if yname == "dsp-ea-bulk-create" { return "DspEaBulkCreate" }
    if yname == "dsp-ea-bulk-update" { return "DspEaBulkUpdate" }
    if yname == "port-stat" { return "PortStat" }
    return ""
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetSegmentPath() string {
    return "coherent-time-stats"
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "opts-ea-bulk-create" {
        return &coherentTimeStats.OptsEaBulkCreate
    }
    if childYangName == "opts-ea-bulk-update" {
        return &coherentTimeStats.OptsEaBulkUpdate
    }
    if childYangName == "dsp-ea-bulk-create" {
        return &coherentTimeStats.DspEaBulkCreate
    }
    if childYangName == "dsp-ea-bulk-update" {
        return &coherentTimeStats.DspEaBulkUpdate
    }
    if childYangName == "port-stat" {
        for _, c := range coherentTimeStats.PortStat {
            if coherentTimeStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node_CoherentTimeStats_PortStat{}
        coherentTimeStats.PortStat = append(coherentTimeStats.PortStat, child)
        return &coherentTimeStats.PortStat[len(coherentTimeStats.PortStat)-1]
    }
    return nil
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["opts-ea-bulk-create"] = &coherentTimeStats.OptsEaBulkCreate
    children["opts-ea-bulk-update"] = &coherentTimeStats.OptsEaBulkUpdate
    children["dsp-ea-bulk-create"] = &coherentTimeStats.DspEaBulkCreate
    children["dsp-ea-bulk-update"] = &coherentTimeStats.DspEaBulkUpdate
    for i := range coherentTimeStats.PortStat {
        children[coherentTimeStats.PortStat[i].GetSegmentPath()] = &coherentTimeStats.PortStat[i]
    }
    return children
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["driver-init"] = coherentTimeStats.DriverInit
    leafs["driver-operational"] = coherentTimeStats.DriverOperational
    leafs["device-created"] = coherentTimeStats.DeviceCreated
    leafs["optics-controllers-created"] = coherentTimeStats.OpticsControllersCreated
    leafs["dsp-controllers-created"] = coherentTimeStats.DspControllersCreated
    leafs["eth-intf-created"] = coherentTimeStats.EthIntfCreated
    return leafs
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetBundleName() string { return "cisco_ios_xr" }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetYangName() string { return "coherent-time-stats" }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) SetParent(parent types.Entity) { coherentTimeStats.parent = parent }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetParent() types.Entity { return coherentTimeStats.parent }

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetParentYangName() string { return "node" }

// Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate
// opts ea bulk create
type Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetFilter() yfilter.YFilter { return optsEaBulkCreate.YFilter }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) SetFilter(yf yfilter.YFilter) { optsEaBulkCreate.YFilter = yf }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetSegmentPath() string {
    return "opts-ea-bulk-create"
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = optsEaBulkCreate.Start
    leafs["end"] = optsEaBulkCreate.End
    leafs["time-taken"] = optsEaBulkCreate.TimeTaken
    leafs["worst-time"] = optsEaBulkCreate.WorstTime
    return leafs
}

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetBundleName() string { return "cisco_ios_xr" }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetYangName() string { return "opts-ea-bulk-create" }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) SetParent(parent types.Entity) { optsEaBulkCreate.parent = parent }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetParent() types.Entity { return optsEaBulkCreate.parent }

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetParentYangName() string { return "coherent-time-stats" }

// Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate
// opts ea bulk update
type Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetFilter() yfilter.YFilter { return optsEaBulkUpdate.YFilter }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) SetFilter(yf yfilter.YFilter) { optsEaBulkUpdate.YFilter = yf }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetSegmentPath() string {
    return "opts-ea-bulk-update"
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = optsEaBulkUpdate.Start
    leafs["end"] = optsEaBulkUpdate.End
    leafs["time-taken"] = optsEaBulkUpdate.TimeTaken
    leafs["worst-time"] = optsEaBulkUpdate.WorstTime
    return leafs
}

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetYangName() string { return "opts-ea-bulk-update" }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) SetParent(parent types.Entity) { optsEaBulkUpdate.parent = parent }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetParent() types.Entity { return optsEaBulkUpdate.parent }

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetParentYangName() string { return "coherent-time-stats" }

// Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate
// dsp ea bulk create
type Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetFilter() yfilter.YFilter { return dspEaBulkCreate.YFilter }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) SetFilter(yf yfilter.YFilter) { dspEaBulkCreate.YFilter = yf }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetSegmentPath() string {
    return "dsp-ea-bulk-create"
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = dspEaBulkCreate.Start
    leafs["end"] = dspEaBulkCreate.End
    leafs["time-taken"] = dspEaBulkCreate.TimeTaken
    leafs["worst-time"] = dspEaBulkCreate.WorstTime
    return leafs
}

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetBundleName() string { return "cisco_ios_xr" }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetYangName() string { return "dsp-ea-bulk-create" }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) SetParent(parent types.Entity) { dspEaBulkCreate.parent = parent }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetParent() types.Entity { return dspEaBulkCreate.parent }

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetParentYangName() string { return "coherent-time-stats" }

// Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate
// dsp ea bulk update
type Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetFilter() yfilter.YFilter { return dspEaBulkUpdate.YFilter }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) SetFilter(yf yfilter.YFilter) { dspEaBulkUpdate.YFilter = yf }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetSegmentPath() string {
    return "dsp-ea-bulk-update"
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = dspEaBulkUpdate.Start
    leafs["end"] = dspEaBulkUpdate.End
    leafs["time-taken"] = dspEaBulkUpdate.TimeTaken
    leafs["worst-time"] = dspEaBulkUpdate.WorstTime
    return leafs
}

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetYangName() string { return "dsp-ea-bulk-update" }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) SetParent(parent types.Entity) { dspEaBulkUpdate.parent = parent }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetParent() types.Entity { return dspEaBulkUpdate.parent }

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetParentYangName() string { return "coherent-time-stats" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat
// port stat
type Coherent_Nodes_Node_CoherentTimeStats_PortStat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // laser state. The type is bool.
    LaserState interface{}

    // provisioned frequency. The type is interface{} with range: 0..4294967295.
    ProvisionedFrequency interface{}

    // tx power. The type is interface{} with range: 0..4294967295.
    TxPower interface{}

    // cd min. The type is interface{} with range: 0..4294967295.
    CdMin interface{}

    // cd max. The type is interface{} with range: 0..4294967295.
    CdMax interface{}

    // traffic type. The type is string with length: 0..128.
    TrafficType interface{}

    // laser on stats.
    LaserOnStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats

    // laser off stats.
    LaserOffStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats

    // wl op stats.
    WlOpStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats

    // txpwr op stats.
    TxpwrOpStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats

    // cdmin op stats.
    CdminOpStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats

    // cdmax op stats.
    CdmaxOpStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats

    // traffictype op stats.
    TraffictypeOpStats Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetFilter() yfilter.YFilter { return portStat.YFilter }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) SetFilter(yf yfilter.YFilter) { portStat.YFilter = yf }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetGoName(yname string) string {
    if yname == "laser-state" { return "LaserState" }
    if yname == "provisioned-frequency" { return "ProvisionedFrequency" }
    if yname == "tx-power" { return "TxPower" }
    if yname == "cd-min" { return "CdMin" }
    if yname == "cd-max" { return "CdMax" }
    if yname == "traffic-type" { return "TrafficType" }
    if yname == "laser-on-stats" { return "LaserOnStats" }
    if yname == "laser-off-stats" { return "LaserOffStats" }
    if yname == "wl-op-stats" { return "WlOpStats" }
    if yname == "txpwr-op-stats" { return "TxpwrOpStats" }
    if yname == "cdmin-op-stats" { return "CdminOpStats" }
    if yname == "cdmax-op-stats" { return "CdmaxOpStats" }
    if yname == "traffictype-op-stats" { return "TraffictypeOpStats" }
    return ""
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetSegmentPath() string {
    return "port-stat"
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "laser-on-stats" {
        return &portStat.LaserOnStats
    }
    if childYangName == "laser-off-stats" {
        return &portStat.LaserOffStats
    }
    if childYangName == "wl-op-stats" {
        return &portStat.WlOpStats
    }
    if childYangName == "txpwr-op-stats" {
        return &portStat.TxpwrOpStats
    }
    if childYangName == "cdmin-op-stats" {
        return &portStat.CdminOpStats
    }
    if childYangName == "cdmax-op-stats" {
        return &portStat.CdmaxOpStats
    }
    if childYangName == "traffictype-op-stats" {
        return &portStat.TraffictypeOpStats
    }
    return nil
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["laser-on-stats"] = &portStat.LaserOnStats
    children["laser-off-stats"] = &portStat.LaserOffStats
    children["wl-op-stats"] = &portStat.WlOpStats
    children["txpwr-op-stats"] = &portStat.TxpwrOpStats
    children["cdmin-op-stats"] = &portStat.CdminOpStats
    children["cdmax-op-stats"] = &portStat.CdmaxOpStats
    children["traffictype-op-stats"] = &portStat.TraffictypeOpStats
    return children
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["laser-state"] = portStat.LaserState
    leafs["provisioned-frequency"] = portStat.ProvisionedFrequency
    leafs["tx-power"] = portStat.TxPower
    leafs["cd-min"] = portStat.CdMin
    leafs["cd-max"] = portStat.CdMax
    leafs["traffic-type"] = portStat.TrafficType
    return leafs
}

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetBundleName() string { return "cisco_ios_xr" }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetYangName() string { return "port-stat" }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) SetParent(parent types.Entity) { portStat.parent = parent }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetParent() types.Entity { return portStat.parent }

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetParentYangName() string { return "coherent-time-stats" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats
// laser on stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetFilter() yfilter.YFilter { return laserOnStats.YFilter }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) SetFilter(yf yfilter.YFilter) { laserOnStats.YFilter = yf }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetSegmentPath() string {
    return "laser-on-stats"
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = laserOnStats.Start
    leafs["end"] = laserOnStats.End
    leafs["time-taken"] = laserOnStats.TimeTaken
    leafs["worst-time"] = laserOnStats.WorstTime
    return leafs
}

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetBundleName() string { return "cisco_ios_xr" }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetYangName() string { return "laser-on-stats" }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) SetParent(parent types.Entity) { laserOnStats.parent = parent }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetParent() types.Entity { return laserOnStats.parent }

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats
// laser off stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetFilter() yfilter.YFilter { return laserOffStats.YFilter }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) SetFilter(yf yfilter.YFilter) { laserOffStats.YFilter = yf }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetSegmentPath() string {
    return "laser-off-stats"
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = laserOffStats.Start
    leafs["end"] = laserOffStats.End
    leafs["time-taken"] = laserOffStats.TimeTaken
    leafs["worst-time"] = laserOffStats.WorstTime
    return leafs
}

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetBundleName() string { return "cisco_ios_xr" }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetYangName() string { return "laser-off-stats" }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) SetParent(parent types.Entity) { laserOffStats.parent = parent }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetParent() types.Entity { return laserOffStats.parent }

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats
// wl op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetFilter() yfilter.YFilter { return wlOpStats.YFilter }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) SetFilter(yf yfilter.YFilter) { wlOpStats.YFilter = yf }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetSegmentPath() string {
    return "wl-op-stats"
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = wlOpStats.Start
    leafs["end"] = wlOpStats.End
    leafs["time-taken"] = wlOpStats.TimeTaken
    leafs["worst-time"] = wlOpStats.WorstTime
    return leafs
}

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetBundleName() string { return "cisco_ios_xr" }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetYangName() string { return "wl-op-stats" }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) SetParent(parent types.Entity) { wlOpStats.parent = parent }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetParent() types.Entity { return wlOpStats.parent }

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats
// txpwr op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetFilter() yfilter.YFilter { return txpwrOpStats.YFilter }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) SetFilter(yf yfilter.YFilter) { txpwrOpStats.YFilter = yf }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetSegmentPath() string {
    return "txpwr-op-stats"
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = txpwrOpStats.Start
    leafs["end"] = txpwrOpStats.End
    leafs["time-taken"] = txpwrOpStats.TimeTaken
    leafs["worst-time"] = txpwrOpStats.WorstTime
    return leafs
}

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetBundleName() string { return "cisco_ios_xr" }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetYangName() string { return "txpwr-op-stats" }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) SetParent(parent types.Entity) { txpwrOpStats.parent = parent }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetParent() types.Entity { return txpwrOpStats.parent }

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats
// cdmin op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetFilter() yfilter.YFilter { return cdminOpStats.YFilter }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) SetFilter(yf yfilter.YFilter) { cdminOpStats.YFilter = yf }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetSegmentPath() string {
    return "cdmin-op-stats"
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = cdminOpStats.Start
    leafs["end"] = cdminOpStats.End
    leafs["time-taken"] = cdminOpStats.TimeTaken
    leafs["worst-time"] = cdminOpStats.WorstTime
    return leafs
}

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetBundleName() string { return "cisco_ios_xr" }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetYangName() string { return "cdmin-op-stats" }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) SetParent(parent types.Entity) { cdminOpStats.parent = parent }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetParent() types.Entity { return cdminOpStats.parent }

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats
// cdmax op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetFilter() yfilter.YFilter { return cdmaxOpStats.YFilter }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) SetFilter(yf yfilter.YFilter) { cdmaxOpStats.YFilter = yf }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetSegmentPath() string {
    return "cdmax-op-stats"
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = cdmaxOpStats.Start
    leafs["end"] = cdmaxOpStats.End
    leafs["time-taken"] = cdmaxOpStats.TimeTaken
    leafs["worst-time"] = cdmaxOpStats.WorstTime
    return leafs
}

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetBundleName() string { return "cisco_ios_xr" }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetYangName() string { return "cdmax-op-stats" }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) SetParent(parent types.Entity) { cdmaxOpStats.parent = parent }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetParent() types.Entity { return cdmaxOpStats.parent }

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats
// traffictype op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // start. The type is string with length: 0..255.
    Start interface{}

    // end. The type is string with length: 0..255.
    End interface{}

    // time taken. The type is string with length: 0..255.
    TimeTaken interface{}

    // worst time. The type is string with length: 0..255.
    WorstTime interface{}
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetFilter() yfilter.YFilter { return traffictypeOpStats.YFilter }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) SetFilter(yf yfilter.YFilter) { traffictypeOpStats.YFilter = yf }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "time-taken" { return "TimeTaken" }
    if yname == "worst-time" { return "WorstTime" }
    return ""
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetSegmentPath() string {
    return "traffictype-op-stats"
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = traffictypeOpStats.Start
    leafs["end"] = traffictypeOpStats.End
    leafs["time-taken"] = traffictypeOpStats.TimeTaken
    leafs["worst-time"] = traffictypeOpStats.WorstTime
    return leafs
}

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetBundleName() string { return "cisco_ios_xr" }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetYangName() string { return "traffictype-op-stats" }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) SetParent(parent types.Entity) { traffictypeOpStats.parent = parent }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetParent() types.Entity { return traffictypeOpStats.parent }

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetParentYangName() string { return "port-stat" }

// Coherent_Nodes_Node_Devicemapping
// Coherent node data for device _mapping
type Coherent_Nodes_Node_Devicemapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // idx. The type is interface{} with range: 0..4294967295.
    Idx interface{}

    // dev map. The type is slice of Coherent_Nodes_Node_Devicemapping_DevMap.
    DevMap []Coherent_Nodes_Node_Devicemapping_DevMap
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetFilter() yfilter.YFilter { return devicemapping.YFilter }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) SetFilter(yf yfilter.YFilter) { devicemapping.YFilter = yf }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetGoName(yname string) string {
    if yname == "idx" { return "Idx" }
    if yname == "dev-map" { return "DevMap" }
    return ""
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetSegmentPath() string {
    return "devicemapping"
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dev-map" {
        for _, c := range devicemapping.DevMap {
            if devicemapping.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node_Devicemapping_DevMap{}
        devicemapping.DevMap = append(devicemapping.DevMap, child)
        return &devicemapping.DevMap[len(devicemapping.DevMap)-1]
    }
    return nil
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range devicemapping.DevMap {
        children[devicemapping.DevMap[i].GetSegmentPath()] = &devicemapping.DevMap[i]
    }
    return children
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idx"] = devicemapping.Idx
    return leafs
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetBundleName() string { return "cisco_ios_xr" }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetYangName() string { return "devicemapping" }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) SetParent(parent types.Entity) { devicemapping.parent = parent }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetParent() types.Entity { return devicemapping.parent }

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetParentYangName() string { return "node" }

// Coherent_Nodes_Node_Devicemapping_DevMap
// dev map
type Coherent_Nodes_Node_Devicemapping_DevMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // device address. The type is string with length: 0..128.
    DeviceAddress interface{}

    // ifhandle. The type is string with length: 0..128.
    Ifhandle interface{}

    // intf name. The type is string with length: 0..64.
    IntfName interface{}
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetFilter() yfilter.YFilter { return devMap.YFilter }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) SetFilter(yf yfilter.YFilter) { devMap.YFilter = yf }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetGoName(yname string) string {
    if yname == "device-address" { return "DeviceAddress" }
    if yname == "ifhandle" { return "Ifhandle" }
    if yname == "intf-name" { return "IntfName" }
    return ""
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetSegmentPath() string {
    return "dev-map"
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["device-address"] = devMap.DeviceAddress
    leafs["ifhandle"] = devMap.Ifhandle
    leafs["intf-name"] = devMap.IntfName
    return leafs
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetBundleName() string { return "cisco_ios_xr" }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetYangName() string { return "dev-map" }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) SetParent(parent types.Entity) { devMap.parent = parent }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetParent() types.Entity { return devMap.parent }

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetParentYangName() string { return "devicemapping" }

// Coherent_Nodes_Node_Coherenthealth
// Coherent node data for driver health
type Coherent_Nodes_Node_Coherenthealth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // im state. The type is bool.
    ImState interface{}

    // aipc srvr state. The type is bool.
    AipcSrvrState interface{}

    // sysdb state. The type is bool.
    SysdbState interface{}

    // pm state. The type is bool.
    PmState interface{}

    // optics ea conn. The type is bool.
    OpticsEaConn interface{}

    // dsp ea conn. The type is bool.
    DspEaConn interface{}

    // vether state. The type is bool.
    VetherState interface{}

    // morgoth alive. The type is bool.
    MorgothAlive interface{}

    // prov infra state. The type is bool.
    ProvInfraState interface{}

    // sdk fpga compatible. The type is bool.
    SdkFpgaCompatible interface{}

    // pending provision. The type is bool.
    PendingProvision interface{}

    // pulse sent. The type is bool.
    PulseSent interface{}

    // inside prov loop. The type is bool.
    InsideProvLoop interface{}

    // fpd in progress. The type is bool.
    FpdInProgress interface{}

    // prov run count. The type is interface{} with range: 0..4294967295.
    ProvRunCount interface{}

    // sdk version. The type is string with length: 0..255.
    SdkVersion interface{}

    // morgoth running version. The type is string with length: 0..255.
    MorgothRunningVersion interface{}

    // morgoth downloaded version. The type is string with length: 0..255.
    MorgothDownloadedVersion interface{}

    // morgoth golden version. The type is string with length: 0..255.
    MorgothGoldenVersion interface{}

    // denali0 version. The type is string with length: 0..255.
    Denali0Version interface{}

    // denali1 version. The type is string with length: 0..255.
    Denali1Version interface{}

    // denali2 version. The type is string with length: 0..255.
    Denali2Version interface{}

    // board type. The type is string with length: 0..255.
    BoardType interface{}

    // jlink op. The type is string with length: 0..6144.
    JlinkOp interface{}

    // port data. The type is slice of
    // Coherent_Nodes_Node_Coherenthealth_PortData.
    PortData []Coherent_Nodes_Node_Coherenthealth_PortData
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetFilter() yfilter.YFilter { return coherenthealth.YFilter }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) SetFilter(yf yfilter.YFilter) { coherenthealth.YFilter = yf }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetGoName(yname string) string {
    if yname == "im-state" { return "ImState" }
    if yname == "aipc-srvr-state" { return "AipcSrvrState" }
    if yname == "sysdb-state" { return "SysdbState" }
    if yname == "pm-state" { return "PmState" }
    if yname == "optics-ea-conn" { return "OpticsEaConn" }
    if yname == "dsp-ea-conn" { return "DspEaConn" }
    if yname == "vether-state" { return "VetherState" }
    if yname == "morgoth-alive" { return "MorgothAlive" }
    if yname == "prov-infra-state" { return "ProvInfraState" }
    if yname == "sdk-fpga-compatible" { return "SdkFpgaCompatible" }
    if yname == "pending-provision" { return "PendingProvision" }
    if yname == "pulse-sent" { return "PulseSent" }
    if yname == "inside-prov-loop" { return "InsideProvLoop" }
    if yname == "fpd-in-progress" { return "FpdInProgress" }
    if yname == "prov-run-count" { return "ProvRunCount" }
    if yname == "sdk-version" { return "SdkVersion" }
    if yname == "morgoth-running-version" { return "MorgothRunningVersion" }
    if yname == "morgoth-downloaded-version" { return "MorgothDownloadedVersion" }
    if yname == "morgoth-golden-version" { return "MorgothGoldenVersion" }
    if yname == "denali0-version" { return "Denali0Version" }
    if yname == "denali1-version" { return "Denali1Version" }
    if yname == "denali2-version" { return "Denali2Version" }
    if yname == "board-type" { return "BoardType" }
    if yname == "jlink-op" { return "JlinkOp" }
    if yname == "port-data" { return "PortData" }
    return ""
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetSegmentPath() string {
    return "coherenthealth"
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-data" {
        for _, c := range coherenthealth.PortData {
            if coherenthealth.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node_Coherenthealth_PortData{}
        coherenthealth.PortData = append(coherenthealth.PortData, child)
        return &coherenthealth.PortData[len(coherenthealth.PortData)-1]
    }
    return nil
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range coherenthealth.PortData {
        children[coherenthealth.PortData[i].GetSegmentPath()] = &coherenthealth.PortData[i]
    }
    return children
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["im-state"] = coherenthealth.ImState
    leafs["aipc-srvr-state"] = coherenthealth.AipcSrvrState
    leafs["sysdb-state"] = coherenthealth.SysdbState
    leafs["pm-state"] = coherenthealth.PmState
    leafs["optics-ea-conn"] = coherenthealth.OpticsEaConn
    leafs["dsp-ea-conn"] = coherenthealth.DspEaConn
    leafs["vether-state"] = coherenthealth.VetherState
    leafs["morgoth-alive"] = coherenthealth.MorgothAlive
    leafs["prov-infra-state"] = coherenthealth.ProvInfraState
    leafs["sdk-fpga-compatible"] = coherenthealth.SdkFpgaCompatible
    leafs["pending-provision"] = coherenthealth.PendingProvision
    leafs["pulse-sent"] = coherenthealth.PulseSent
    leafs["inside-prov-loop"] = coherenthealth.InsideProvLoop
    leafs["fpd-in-progress"] = coherenthealth.FpdInProgress
    leafs["prov-run-count"] = coherenthealth.ProvRunCount
    leafs["sdk-version"] = coherenthealth.SdkVersion
    leafs["morgoth-running-version"] = coherenthealth.MorgothRunningVersion
    leafs["morgoth-downloaded-version"] = coherenthealth.MorgothDownloadedVersion
    leafs["morgoth-golden-version"] = coherenthealth.MorgothGoldenVersion
    leafs["denali0-version"] = coherenthealth.Denali0Version
    leafs["denali1-version"] = coherenthealth.Denali1Version
    leafs["denali2-version"] = coherenthealth.Denali2Version
    leafs["board-type"] = coherenthealth.BoardType
    leafs["jlink-op"] = coherenthealth.JlinkOp
    return leafs
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetBundleName() string { return "cisco_ios_xr" }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetYangName() string { return "coherenthealth" }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) SetParent(parent types.Entity) { coherenthealth.parent = parent }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetParent() types.Entity { return coherenthealth.parent }

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetParentYangName() string { return "node" }

// Coherent_Nodes_Node_Coherenthealth_PortData
// port data
type Coherent_Nodes_Node_Coherenthealth_PortData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // optics ctrl created. The type is bool.
    OpticsCtrlCreated interface{}

    // dsp ctrl created. The type is bool.
    DspCtrlCreated interface{}

    // has pluggable. The type is bool.
    HasPluggable interface{}

    // optics admin up. The type is bool.
    OpticsAdminUp interface{}

    // dsp admin up. The type is bool.
    DspAdminUp interface{}

    // laser state. The type is bool.
    LaserState interface{}

    // laser on pending. The type is bool.
    LaserOnPending interface{}

    // provisioning needed. The type is bool.
    ProvisioningNeeded interface{}

    // force reprovision. The type is bool.
    ForceReprovision interface{}

    // fp port idx. The type is interface{} with range: 0..4294967295.
    FpPortIdx interface{}

    // configured frequency. The type is interface{} with range: 0..4294967295.
    ConfiguredFrequency interface{}

    // provisioned frequency. The type is interface{} with range: 0..4294967295.
    ProvisionedFrequency interface{}

    // configured tx power. The type is string with length: 0..128.
    ConfiguredTxPower interface{}

    // provisioned tx power. The type is string with length: 0..128.
    ProvisionedTxPower interface{}

    // configured cd min. The type is string with length: 0..128.
    ConfiguredCdMin interface{}

    // provisioned cd min. The type is string with length: 0..128.
    ProvisionedCdMin interface{}

    // configured cd max. The type is string with length: 0..128.
    ConfiguredCdMax interface{}

    // provisioned cd max. The type is string with length: 0..128.
    ProvisionedCdMax interface{}

    // configured traffic type. The type is string with length: 0..128.
    ConfiguredTrafficType interface{}

    // provisioned traffic type. The type is string with length: 0..128.
    ProvisionedTrafficType interface{}

    // configured loopback mode. The type is string with length: 0..128.
    ConfiguredLoopbackMode interface{}

    // provisioned loopback mode. The type is string with length: 0..128.
    ProvisionedLoopbackMode interface{}

    // expected ctp2 led state. The type is string with length: 0..128.
    ExpectedCtp2LedState interface{}

    // provisioned ctp2 led state. The type is string with length: 0..128.
    ProvisionedCtp2LedState interface{}

    // led op rc. The type is interface{} with range: -2147483648..2147483647.
    LedOpRc interface{}

    // laser op rc. The type is interface{} with range: -2147483648..2147483647.
    LaserOpRc interface{}

    // wlen op rc. The type is interface{} with range: -2147483648..2147483647.
    WlenOpRc interface{}

    // traffic op rc. The type is interface{} with range: -2147483648..2147483647.
    TrafficOpRc interface{}

    // loopback op rc. The type is interface{} with range:
    // -2147483648..2147483647.
    LoopbackOpRc interface{}

    // tx power op rc. The type is interface{} with range:
    // -2147483648..2147483647.
    TxPowerOpRc interface{}

    // cd min op rc. The type is interface{} with range: -2147483648..2147483647.
    CdMinOpRc interface{}

    // cd max op rc. The type is interface{} with range: -2147483648..2147483647.
    CdMaxOpRc interface{}

    // provisioning failed. The type is bool.
    ProvisioningFailed interface{}

    // ctp2 hw alarms. The type is string with length: 0..128.
    Ctp2HwAlarms interface{}

    // denali hw alarms. The type is string with length: 0..256.
    DenaliHwAlarms interface{}

    // is pm port created opt. The type is bool.
    IsPmPortCreatedOpt interface{}

    // rc pm port opt. The type is interface{} with range:
    // -2147483648..2147483647.
    RcPmPortOpt interface{}

    // pm port state opt. The type is interface{} with range:
    // -2147483648..2147483647.
    PmPortStateOpt interface{}

    // rc pm provision opt. The type is interface{} with range:
    // -2147483648..2147483647.
    RcPmProvisionOpt interface{}

    // is alarm port created opt. The type is bool.
    IsAlarmPortCreatedOpt interface{}

    // rc alarm port opt. The type is interface{} with range:
    // -2147483648..2147483647.
    RcAlarmPortOpt interface{}

    // is pm port created dsp. The type is bool.
    IsPmPortCreatedDsp interface{}

    // rc pm port dsp. The type is interface{} with range:
    // -2147483648..2147483647.
    RcPmPortDsp interface{}

    // pm port state dsp. The type is interface{} with range:
    // -2147483648..2147483647.
    PmPortStateDsp interface{}

    // rc pm provision dsp. The type is interface{} with range:
    // -2147483648..2147483647.
    RcPmProvisionDsp interface{}

    // is alarm port created dsp. The type is bool.
    IsAlarmPortCreatedDsp interface{}

    // rc alarm port dsp. The type is interface{} with range:
    // -2147483648..2147483647.
    RcAlarmPortDsp interface{}

    // ctp info.
    CtpInfo Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo

    // interface info.
    InterfaceInfo Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetFilter() yfilter.YFilter { return portData.YFilter }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) SetFilter(yf yfilter.YFilter) { portData.YFilter = yf }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetGoName(yname string) string {
    if yname == "optics-ctrl-created" { return "OpticsCtrlCreated" }
    if yname == "dsp-ctrl-created" { return "DspCtrlCreated" }
    if yname == "has-pluggable" { return "HasPluggable" }
    if yname == "optics-admin-up" { return "OpticsAdminUp" }
    if yname == "dsp-admin-up" { return "DspAdminUp" }
    if yname == "laser-state" { return "LaserState" }
    if yname == "laser-on-pending" { return "LaserOnPending" }
    if yname == "provisioning-needed" { return "ProvisioningNeeded" }
    if yname == "force-reprovision" { return "ForceReprovision" }
    if yname == "fp-port-idx" { return "FpPortIdx" }
    if yname == "configured-frequency" { return "ConfiguredFrequency" }
    if yname == "provisioned-frequency" { return "ProvisionedFrequency" }
    if yname == "configured-tx-power" { return "ConfiguredTxPower" }
    if yname == "provisioned-tx-power" { return "ProvisionedTxPower" }
    if yname == "configured-cd-min" { return "ConfiguredCdMin" }
    if yname == "provisioned-cd-min" { return "ProvisionedCdMin" }
    if yname == "configured-cd-max" { return "ConfiguredCdMax" }
    if yname == "provisioned-cd-max" { return "ProvisionedCdMax" }
    if yname == "configured-traffic-type" { return "ConfiguredTrafficType" }
    if yname == "provisioned-traffic-type" { return "ProvisionedTrafficType" }
    if yname == "configured-loopback-mode" { return "ConfiguredLoopbackMode" }
    if yname == "provisioned-loopback-mode" { return "ProvisionedLoopbackMode" }
    if yname == "expected-ctp2-led-state" { return "ExpectedCtp2LedState" }
    if yname == "provisioned-ctp2-led-state" { return "ProvisionedCtp2LedState" }
    if yname == "led-op-rc" { return "LedOpRc" }
    if yname == "laser-op-rc" { return "LaserOpRc" }
    if yname == "wlen-op-rc" { return "WlenOpRc" }
    if yname == "traffic-op-rc" { return "TrafficOpRc" }
    if yname == "loopback-op-rc" { return "LoopbackOpRc" }
    if yname == "tx-power-op-rc" { return "TxPowerOpRc" }
    if yname == "cd-min-op-rc" { return "CdMinOpRc" }
    if yname == "cd-max-op-rc" { return "CdMaxOpRc" }
    if yname == "provisioning-failed" { return "ProvisioningFailed" }
    if yname == "ctp2-hw-alarms" { return "Ctp2HwAlarms" }
    if yname == "denali-hw-alarms" { return "DenaliHwAlarms" }
    if yname == "is-pm-port-created-opt" { return "IsPmPortCreatedOpt" }
    if yname == "rc-pm-port-opt" { return "RcPmPortOpt" }
    if yname == "pm-port-state-opt" { return "PmPortStateOpt" }
    if yname == "rc-pm-provision-opt" { return "RcPmProvisionOpt" }
    if yname == "is-alarm-port-created-opt" { return "IsAlarmPortCreatedOpt" }
    if yname == "rc-alarm-port-opt" { return "RcAlarmPortOpt" }
    if yname == "is-pm-port-created-dsp" { return "IsPmPortCreatedDsp" }
    if yname == "rc-pm-port-dsp" { return "RcPmPortDsp" }
    if yname == "pm-port-state-dsp" { return "PmPortStateDsp" }
    if yname == "rc-pm-provision-dsp" { return "RcPmProvisionDsp" }
    if yname == "is-alarm-port-created-dsp" { return "IsAlarmPortCreatedDsp" }
    if yname == "rc-alarm-port-dsp" { return "RcAlarmPortDsp" }
    if yname == "ctp-info" { return "CtpInfo" }
    if yname == "interface-info" { return "InterfaceInfo" }
    return ""
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetSegmentPath() string {
    return "port-data"
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ctp-info" {
        return &portData.CtpInfo
    }
    if childYangName == "interface-info" {
        return &portData.InterfaceInfo
    }
    return nil
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ctp-info"] = &portData.CtpInfo
    children["interface-info"] = &portData.InterfaceInfo
    return children
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["optics-ctrl-created"] = portData.OpticsCtrlCreated
    leafs["dsp-ctrl-created"] = portData.DspCtrlCreated
    leafs["has-pluggable"] = portData.HasPluggable
    leafs["optics-admin-up"] = portData.OpticsAdminUp
    leafs["dsp-admin-up"] = portData.DspAdminUp
    leafs["laser-state"] = portData.LaserState
    leafs["laser-on-pending"] = portData.LaserOnPending
    leafs["provisioning-needed"] = portData.ProvisioningNeeded
    leafs["force-reprovision"] = portData.ForceReprovision
    leafs["fp-port-idx"] = portData.FpPortIdx
    leafs["configured-frequency"] = portData.ConfiguredFrequency
    leafs["provisioned-frequency"] = portData.ProvisionedFrequency
    leafs["configured-tx-power"] = portData.ConfiguredTxPower
    leafs["provisioned-tx-power"] = portData.ProvisionedTxPower
    leafs["configured-cd-min"] = portData.ConfiguredCdMin
    leafs["provisioned-cd-min"] = portData.ProvisionedCdMin
    leafs["configured-cd-max"] = portData.ConfiguredCdMax
    leafs["provisioned-cd-max"] = portData.ProvisionedCdMax
    leafs["configured-traffic-type"] = portData.ConfiguredTrafficType
    leafs["provisioned-traffic-type"] = portData.ProvisionedTrafficType
    leafs["configured-loopback-mode"] = portData.ConfiguredLoopbackMode
    leafs["provisioned-loopback-mode"] = portData.ProvisionedLoopbackMode
    leafs["expected-ctp2-led-state"] = portData.ExpectedCtp2LedState
    leafs["provisioned-ctp2-led-state"] = portData.ProvisionedCtp2LedState
    leafs["led-op-rc"] = portData.LedOpRc
    leafs["laser-op-rc"] = portData.LaserOpRc
    leafs["wlen-op-rc"] = portData.WlenOpRc
    leafs["traffic-op-rc"] = portData.TrafficOpRc
    leafs["loopback-op-rc"] = portData.LoopbackOpRc
    leafs["tx-power-op-rc"] = portData.TxPowerOpRc
    leafs["cd-min-op-rc"] = portData.CdMinOpRc
    leafs["cd-max-op-rc"] = portData.CdMaxOpRc
    leafs["provisioning-failed"] = portData.ProvisioningFailed
    leafs["ctp2-hw-alarms"] = portData.Ctp2HwAlarms
    leafs["denali-hw-alarms"] = portData.DenaliHwAlarms
    leafs["is-pm-port-created-opt"] = portData.IsPmPortCreatedOpt
    leafs["rc-pm-port-opt"] = portData.RcPmPortOpt
    leafs["pm-port-state-opt"] = portData.PmPortStateOpt
    leafs["rc-pm-provision-opt"] = portData.RcPmProvisionOpt
    leafs["is-alarm-port-created-opt"] = portData.IsAlarmPortCreatedOpt
    leafs["rc-alarm-port-opt"] = portData.RcAlarmPortOpt
    leafs["is-pm-port-created-dsp"] = portData.IsPmPortCreatedDsp
    leafs["rc-pm-port-dsp"] = portData.RcPmPortDsp
    leafs["pm-port-state-dsp"] = portData.PmPortStateDsp
    leafs["rc-pm-provision-dsp"] = portData.RcPmProvisionDsp
    leafs["is-alarm-port-created-dsp"] = portData.IsAlarmPortCreatedDsp
    leafs["rc-alarm-port-dsp"] = portData.RcAlarmPortDsp
    return leafs
}

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetBundleName() string { return "cisco_ios_xr" }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetYangName() string { return "port-data" }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) SetParent(parent types.Entity) { portData.parent = parent }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetParent() types.Entity { return portData.parent }

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetParentYangName() string { return "coherenthealth" }

// Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo
// ctp info
type Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // deviation. The type is string with length: 0..16.
    Deviation interface{}

    // part number. The type is string with length: 0..16.
    PartNumber interface{}

    // serial number. The type is string with length: 0..16.
    SerialNumber interface{}

    // date code number. The type is string with length: 0..10.
    DateCodeNumber interface{}

    // CLEI code number. The type is string with length: 0..10.
    CleiCodeNumber interface{}

    // vendorname. The type is string with length: 0..16.
    Vendorname interface{}

    // description. The type is string with length: 0..64.
    Description interface{}

    // pid. The type is string with length: 0..16.
    Pid interface{}

    // vid. The type is string with length: 0..16.
    Vid interface{}

    // module hardware version number. The type is interface{} with range:
    // 0..65535.
    ModuleHardwareVersionNumber interface{}

    // module firmware running version number. The type is interface{} with range:
    // 0..65535.
    ModuleFirmwareRunningVersionNumber interface{}

    // module firmware committed version number. The type is interface{} with
    // range: 0..65535.
    ModuleFirmwareCommittedVersionNumber interface{}

    // ctp type. The type is interface{} with range: 0..4294967295.
    CtpType interface{}
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetFilter() yfilter.YFilter { return ctpInfo.YFilter }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) SetFilter(yf yfilter.YFilter) { ctpInfo.YFilter = yf }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetGoName(yname string) string {
    if yname == "deviation" { return "Deviation" }
    if yname == "part-number" { return "PartNumber" }
    if yname == "serial-number" { return "SerialNumber" }
    if yname == "date-code-number" { return "DateCodeNumber" }
    if yname == "clei-code-number" { return "CleiCodeNumber" }
    if yname == "vendorname" { return "Vendorname" }
    if yname == "description" { return "Description" }
    if yname == "pid" { return "Pid" }
    if yname == "vid" { return "Vid" }
    if yname == "module-hardware-version-number" { return "ModuleHardwareVersionNumber" }
    if yname == "module-firmware-running-version-number" { return "ModuleFirmwareRunningVersionNumber" }
    if yname == "module-firmware-committed-version-number" { return "ModuleFirmwareCommittedVersionNumber" }
    if yname == "ctp-type" { return "CtpType" }
    return ""
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetSegmentPath() string {
    return "ctp-info"
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["deviation"] = ctpInfo.Deviation
    leafs["part-number"] = ctpInfo.PartNumber
    leafs["serial-number"] = ctpInfo.SerialNumber
    leafs["date-code-number"] = ctpInfo.DateCodeNumber
    leafs["clei-code-number"] = ctpInfo.CleiCodeNumber
    leafs["vendorname"] = ctpInfo.Vendorname
    leafs["description"] = ctpInfo.Description
    leafs["pid"] = ctpInfo.Pid
    leafs["vid"] = ctpInfo.Vid
    leafs["module-hardware-version-number"] = ctpInfo.ModuleHardwareVersionNumber
    leafs["module-firmware-running-version-number"] = ctpInfo.ModuleFirmwareRunningVersionNumber
    leafs["module-firmware-committed-version-number"] = ctpInfo.ModuleFirmwareCommittedVersionNumber
    leafs["ctp-type"] = ctpInfo.CtpType
    return leafs
}

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetYangName() string { return "ctp-info" }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) SetParent(parent types.Entity) { ctpInfo.parent = parent }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetParent() types.Entity { return ctpInfo.parent }

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetParentYangName() string { return "port-data" }

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo
// interface info
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // eth data. The type is slice of
    // Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData.
    EthData []Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetFilter() yfilter.YFilter { return interfaceInfo.YFilter }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) SetFilter(yf yfilter.YFilter) { interfaceInfo.YFilter = yf }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetGoName(yname string) string {
    if yname == "eth-data" { return "EthData" }
    return ""
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetSegmentPath() string {
    return "interface-info"
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "eth-data" {
        for _, c := range interfaceInfo.EthData {
            if interfaceInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData{}
        interfaceInfo.EthData = append(interfaceInfo.EthData, child)
        return &interfaceInfo.EthData[len(interfaceInfo.EthData)-1]
    }
    return nil
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceInfo.EthData {
        children[interfaceInfo.EthData[i].GetSegmentPath()] = &interfaceInfo.EthData[i]
    }
    return children
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetYangName() string { return "interface-info" }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) SetParent(parent types.Entity) { interfaceInfo.parent = parent }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetParent() types.Entity { return interfaceInfo.parent }

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetParentYangName() string { return "port-data" }

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
// eth data
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ifname. The type is string with length: 0..64.
    Ifname interface{}

    // intf handle. The type is string with length: 0..128.
    IntfHandle interface{}

    // admin state. The type is string with length: 0..128.
    AdminState interface{}

    // admin up. The type is bool.
    AdminUp interface{}

    // is created. The type is bool.
    IsCreated interface{}
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetFilter() yfilter.YFilter { return ethData.YFilter }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) SetFilter(yf yfilter.YFilter) { ethData.YFilter = yf }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetGoName(yname string) string {
    if yname == "ifname" { return "Ifname" }
    if yname == "intf-handle" { return "IntfHandle" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "admin-up" { return "AdminUp" }
    if yname == "is-created" { return "IsCreated" }
    return ""
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetSegmentPath() string {
    return "eth-data"
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifname"] = ethData.Ifname
    leafs["intf-handle"] = ethData.IntfHandle
    leafs["admin-state"] = ethData.AdminState
    leafs["admin-up"] = ethData.AdminUp
    leafs["is-created"] = ethData.IsCreated
    return leafs
}

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetBundleName() string { return "cisco_ios_xr" }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetYangName() string { return "eth-data" }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) SetParent(parent types.Entity) { ethData.parent = parent }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetParent() types.Entity { return ethData.parent }

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetParentYangName() string { return "interface-info" }

// Coherent_Nodes_Node_PortModeAllInfo
// PortMode all operational data
type Coherent_Nodes_Node_PortModeAllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dev map idx. The type is interface{} with range: 0..4294967295.
    Idx interface{}

    // portmode entry. The type is slice of
    // Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry.
    PortmodeEntry []Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetFilter() yfilter.YFilter { return portModeAllInfo.YFilter }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) SetFilter(yf yfilter.YFilter) { portModeAllInfo.YFilter = yf }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetGoName(yname string) string {
    if yname == "idx" { return "Idx" }
    if yname == "portmode-entry" { return "PortmodeEntry" }
    return ""
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetSegmentPath() string {
    return "port-mode-all-info"
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "portmode-entry" {
        for _, c := range portModeAllInfo.PortmodeEntry {
            if portModeAllInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry{}
        portModeAllInfo.PortmodeEntry = append(portModeAllInfo.PortmodeEntry, child)
        return &portModeAllInfo.PortmodeEntry[len(portModeAllInfo.PortmodeEntry)-1]
    }
    return nil
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portModeAllInfo.PortmodeEntry {
        children[portModeAllInfo.PortmodeEntry[i].GetSegmentPath()] = &portModeAllInfo.PortmodeEntry[i]
    }
    return children
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idx"] = portModeAllInfo.Idx
    return leafs
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetYangName() string { return "port-mode-all-info" }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) SetParent(parent types.Entity) { portModeAllInfo.parent = parent }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetParent() types.Entity { return portModeAllInfo.parent }

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetParentYangName() string { return "node" }

// Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry
// portmode entry
type Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry struct {
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

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetFilter() yfilter.YFilter { return portmodeEntry.YFilter }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) SetFilter(yf yfilter.YFilter) { portmodeEntry.YFilter = yf }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetGoName(yname string) string {
    if yname == "intf-name" { return "IntfName" }
    if yname == "speed" { return "Speed" }
    if yname == "fec" { return "Fec" }
    if yname == "diff" { return "Diff" }
    if yname == "modulation" { return "Modulation" }
    return ""
}

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetSegmentPath() string {
    return "portmode-entry"
}

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intf-name"] = portmodeEntry.IntfName
    leafs["speed"] = portmodeEntry.Speed
    leafs["fec"] = portmodeEntry.Fec
    leafs["diff"] = portmodeEntry.Diff
    leafs["modulation"] = portmodeEntry.Modulation
    return leafs
}

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetBundleName() string { return "cisco_ios_xr" }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetYangName() string { return "portmode-entry" }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) SetParent(parent types.Entity) { portmodeEntry.parent = parent }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetParent() types.Entity { return portmodeEntry.parent }

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetParentYangName() string { return "port-mode-all-info" }

