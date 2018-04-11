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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coherent list of nodes.
    Nodes Coherent_Nodes
}

func (coherent *Coherent) GetEntityData() *types.CommonEntityData {
    coherent.EntityData.YFilter = coherent.YFilter
    coherent.EntityData.YangName = "coherent"
    coherent.EntityData.BundleName = "cisco_ios_xr"
    coherent.EntityData.ParentYangName = "Cisco-IOS-XR-ncs5500-coherent-node-oper"
    coherent.EntityData.SegmentPath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent"
    coherent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherent.EntityData.Children = make(map[string]types.YChild)
    coherent.EntityData.Children["nodes"] = types.YChild{"Nodes", &coherent.Nodes}
    coherent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(coherent.EntityData)
}

// Coherent_Nodes
// Coherent list of nodes
type Coherent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coherent discovery operational data for a particular node. The type is
    // slice of Coherent_Nodes_Node.
    Node []Coherent_Nodes_Node
}

func (nodes *Coherent_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "coherent"
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

// Coherent_Nodes_Node
// Coherent discovery operational data for a
// particular node
type Coherent_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Coherent_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["coherent-time-stats"] = types.YChild{"CoherentTimeStats", &node.CoherentTimeStats}
    node.EntityData.Children["devicemapping"] = types.YChild{"Devicemapping", &node.Devicemapping}
    node.EntityData.Children["coherenthealth"] = types.YChild{"Coherenthealth", &node.Coherenthealth}
    node.EntityData.Children["port-mode-all-info"] = types.YChild{"PortModeAllInfo", &node.PortModeAllInfo}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats
// Coherent driver performace information
type Coherent_Nodes_Node_CoherentTimeStats struct {
    EntityData types.CommonEntityData
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

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetEntityData() *types.CommonEntityData {
    coherentTimeStats.EntityData.YFilter = coherentTimeStats.YFilter
    coherentTimeStats.EntityData.YangName = "coherent-time-stats"
    coherentTimeStats.EntityData.BundleName = "cisco_ios_xr"
    coherentTimeStats.EntityData.ParentYangName = "node"
    coherentTimeStats.EntityData.SegmentPath = "coherent-time-stats"
    coherentTimeStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherentTimeStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherentTimeStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherentTimeStats.EntityData.Children = make(map[string]types.YChild)
    coherentTimeStats.EntityData.Children["opts-ea-bulk-create"] = types.YChild{"OptsEaBulkCreate", &coherentTimeStats.OptsEaBulkCreate}
    coherentTimeStats.EntityData.Children["opts-ea-bulk-update"] = types.YChild{"OptsEaBulkUpdate", &coherentTimeStats.OptsEaBulkUpdate}
    coherentTimeStats.EntityData.Children["dsp-ea-bulk-create"] = types.YChild{"DspEaBulkCreate", &coherentTimeStats.DspEaBulkCreate}
    coherentTimeStats.EntityData.Children["dsp-ea-bulk-update"] = types.YChild{"DspEaBulkUpdate", &coherentTimeStats.DspEaBulkUpdate}
    coherentTimeStats.EntityData.Children["port-stat"] = types.YChild{"PortStat", nil}
    for i := range coherentTimeStats.PortStat {
        coherentTimeStats.EntityData.Children[types.GetSegmentPath(&coherentTimeStats.PortStat[i])] = types.YChild{"PortStat", &coherentTimeStats.PortStat[i]}
    }
    coherentTimeStats.EntityData.Leafs = make(map[string]types.YLeaf)
    coherentTimeStats.EntityData.Leafs["driver-init"] = types.YLeaf{"DriverInit", coherentTimeStats.DriverInit}
    coherentTimeStats.EntityData.Leafs["driver-operational"] = types.YLeaf{"DriverOperational", coherentTimeStats.DriverOperational}
    coherentTimeStats.EntityData.Leafs["device-created"] = types.YLeaf{"DeviceCreated", coherentTimeStats.DeviceCreated}
    coherentTimeStats.EntityData.Leafs["optics-controllers-created"] = types.YLeaf{"OpticsControllersCreated", coherentTimeStats.OpticsControllersCreated}
    coherentTimeStats.EntityData.Leafs["dsp-controllers-created"] = types.YLeaf{"DspControllersCreated", coherentTimeStats.DspControllersCreated}
    coherentTimeStats.EntityData.Leafs["eth-intf-created"] = types.YLeaf{"EthIntfCreated", coherentTimeStats.EthIntfCreated}
    return &(coherentTimeStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate
// opts ea bulk create
type Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate struct {
    EntityData types.CommonEntityData
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

func (optsEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkCreate) GetEntityData() *types.CommonEntityData {
    optsEaBulkCreate.EntityData.YFilter = optsEaBulkCreate.YFilter
    optsEaBulkCreate.EntityData.YangName = "opts-ea-bulk-create"
    optsEaBulkCreate.EntityData.BundleName = "cisco_ios_xr"
    optsEaBulkCreate.EntityData.ParentYangName = "coherent-time-stats"
    optsEaBulkCreate.EntityData.SegmentPath = "opts-ea-bulk-create"
    optsEaBulkCreate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optsEaBulkCreate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optsEaBulkCreate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optsEaBulkCreate.EntityData.Children = make(map[string]types.YChild)
    optsEaBulkCreate.EntityData.Leafs = make(map[string]types.YLeaf)
    optsEaBulkCreate.EntityData.Leafs["start"] = types.YLeaf{"Start", optsEaBulkCreate.Start}
    optsEaBulkCreate.EntityData.Leafs["end"] = types.YLeaf{"End", optsEaBulkCreate.End}
    optsEaBulkCreate.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", optsEaBulkCreate.TimeTaken}
    optsEaBulkCreate.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", optsEaBulkCreate.WorstTime}
    return &(optsEaBulkCreate.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate
// opts ea bulk update
type Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate struct {
    EntityData types.CommonEntityData
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

func (optsEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_OptsEaBulkUpdate) GetEntityData() *types.CommonEntityData {
    optsEaBulkUpdate.EntityData.YFilter = optsEaBulkUpdate.YFilter
    optsEaBulkUpdate.EntityData.YangName = "opts-ea-bulk-update"
    optsEaBulkUpdate.EntityData.BundleName = "cisco_ios_xr"
    optsEaBulkUpdate.EntityData.ParentYangName = "coherent-time-stats"
    optsEaBulkUpdate.EntityData.SegmentPath = "opts-ea-bulk-update"
    optsEaBulkUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optsEaBulkUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optsEaBulkUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optsEaBulkUpdate.EntityData.Children = make(map[string]types.YChild)
    optsEaBulkUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    optsEaBulkUpdate.EntityData.Leafs["start"] = types.YLeaf{"Start", optsEaBulkUpdate.Start}
    optsEaBulkUpdate.EntityData.Leafs["end"] = types.YLeaf{"End", optsEaBulkUpdate.End}
    optsEaBulkUpdate.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", optsEaBulkUpdate.TimeTaken}
    optsEaBulkUpdate.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", optsEaBulkUpdate.WorstTime}
    return &(optsEaBulkUpdate.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate
// dsp ea bulk create
type Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate struct {
    EntityData types.CommonEntityData
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

func (dspEaBulkCreate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkCreate) GetEntityData() *types.CommonEntityData {
    dspEaBulkCreate.EntityData.YFilter = dspEaBulkCreate.YFilter
    dspEaBulkCreate.EntityData.YangName = "dsp-ea-bulk-create"
    dspEaBulkCreate.EntityData.BundleName = "cisco_ios_xr"
    dspEaBulkCreate.EntityData.ParentYangName = "coherent-time-stats"
    dspEaBulkCreate.EntityData.SegmentPath = "dsp-ea-bulk-create"
    dspEaBulkCreate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dspEaBulkCreate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dspEaBulkCreate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dspEaBulkCreate.EntityData.Children = make(map[string]types.YChild)
    dspEaBulkCreate.EntityData.Leafs = make(map[string]types.YLeaf)
    dspEaBulkCreate.EntityData.Leafs["start"] = types.YLeaf{"Start", dspEaBulkCreate.Start}
    dspEaBulkCreate.EntityData.Leafs["end"] = types.YLeaf{"End", dspEaBulkCreate.End}
    dspEaBulkCreate.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", dspEaBulkCreate.TimeTaken}
    dspEaBulkCreate.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", dspEaBulkCreate.WorstTime}
    return &(dspEaBulkCreate.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate
// dsp ea bulk update
type Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate struct {
    EntityData types.CommonEntityData
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

func (dspEaBulkUpdate *Coherent_Nodes_Node_CoherentTimeStats_DspEaBulkUpdate) GetEntityData() *types.CommonEntityData {
    dspEaBulkUpdate.EntityData.YFilter = dspEaBulkUpdate.YFilter
    dspEaBulkUpdate.EntityData.YangName = "dsp-ea-bulk-update"
    dspEaBulkUpdate.EntityData.BundleName = "cisco_ios_xr"
    dspEaBulkUpdate.EntityData.ParentYangName = "coherent-time-stats"
    dspEaBulkUpdate.EntityData.SegmentPath = "dsp-ea-bulk-update"
    dspEaBulkUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dspEaBulkUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dspEaBulkUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dspEaBulkUpdate.EntityData.Children = make(map[string]types.YChild)
    dspEaBulkUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    dspEaBulkUpdate.EntityData.Leafs["start"] = types.YLeaf{"Start", dspEaBulkUpdate.Start}
    dspEaBulkUpdate.EntityData.Leafs["end"] = types.YLeaf{"End", dspEaBulkUpdate.End}
    dspEaBulkUpdate.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", dspEaBulkUpdate.TimeTaken}
    dspEaBulkUpdate.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", dspEaBulkUpdate.WorstTime}
    return &(dspEaBulkUpdate.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat
// port stat
type Coherent_Nodes_Node_CoherentTimeStats_PortStat struct {
    EntityData types.CommonEntityData
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

func (portStat *Coherent_Nodes_Node_CoherentTimeStats_PortStat) GetEntityData() *types.CommonEntityData {
    portStat.EntityData.YFilter = portStat.YFilter
    portStat.EntityData.YangName = "port-stat"
    portStat.EntityData.BundleName = "cisco_ios_xr"
    portStat.EntityData.ParentYangName = "coherent-time-stats"
    portStat.EntityData.SegmentPath = "port-stat"
    portStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStat.EntityData.Children = make(map[string]types.YChild)
    portStat.EntityData.Children["laser-on-stats"] = types.YChild{"LaserOnStats", &portStat.LaserOnStats}
    portStat.EntityData.Children["laser-off-stats"] = types.YChild{"LaserOffStats", &portStat.LaserOffStats}
    portStat.EntityData.Children["wl-op-stats"] = types.YChild{"WlOpStats", &portStat.WlOpStats}
    portStat.EntityData.Children["txpwr-op-stats"] = types.YChild{"TxpwrOpStats", &portStat.TxpwrOpStats}
    portStat.EntityData.Children["cdmin-op-stats"] = types.YChild{"CdminOpStats", &portStat.CdminOpStats}
    portStat.EntityData.Children["cdmax-op-stats"] = types.YChild{"CdmaxOpStats", &portStat.CdmaxOpStats}
    portStat.EntityData.Children["traffictype-op-stats"] = types.YChild{"TraffictypeOpStats", &portStat.TraffictypeOpStats}
    portStat.EntityData.Leafs = make(map[string]types.YLeaf)
    portStat.EntityData.Leafs["laser-state"] = types.YLeaf{"LaserState", portStat.LaserState}
    portStat.EntityData.Leafs["provisioned-frequency"] = types.YLeaf{"ProvisionedFrequency", portStat.ProvisionedFrequency}
    portStat.EntityData.Leafs["tx-power"] = types.YLeaf{"TxPower", portStat.TxPower}
    portStat.EntityData.Leafs["cd-min"] = types.YLeaf{"CdMin", portStat.CdMin}
    portStat.EntityData.Leafs["cd-max"] = types.YLeaf{"CdMax", portStat.CdMax}
    portStat.EntityData.Leafs["traffic-type"] = types.YLeaf{"TrafficType", portStat.TrafficType}
    return &(portStat.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats
// laser on stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats struct {
    EntityData types.CommonEntityData
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

func (laserOnStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOnStats) GetEntityData() *types.CommonEntityData {
    laserOnStats.EntityData.YFilter = laserOnStats.YFilter
    laserOnStats.EntityData.YangName = "laser-on-stats"
    laserOnStats.EntityData.BundleName = "cisco_ios_xr"
    laserOnStats.EntityData.ParentYangName = "port-stat"
    laserOnStats.EntityData.SegmentPath = "laser-on-stats"
    laserOnStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laserOnStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laserOnStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laserOnStats.EntityData.Children = make(map[string]types.YChild)
    laserOnStats.EntityData.Leafs = make(map[string]types.YLeaf)
    laserOnStats.EntityData.Leafs["start"] = types.YLeaf{"Start", laserOnStats.Start}
    laserOnStats.EntityData.Leafs["end"] = types.YLeaf{"End", laserOnStats.End}
    laserOnStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", laserOnStats.TimeTaken}
    laserOnStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", laserOnStats.WorstTime}
    return &(laserOnStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats
// laser off stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats struct {
    EntityData types.CommonEntityData
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

func (laserOffStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_LaserOffStats) GetEntityData() *types.CommonEntityData {
    laserOffStats.EntityData.YFilter = laserOffStats.YFilter
    laserOffStats.EntityData.YangName = "laser-off-stats"
    laserOffStats.EntityData.BundleName = "cisco_ios_xr"
    laserOffStats.EntityData.ParentYangName = "port-stat"
    laserOffStats.EntityData.SegmentPath = "laser-off-stats"
    laserOffStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laserOffStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laserOffStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laserOffStats.EntityData.Children = make(map[string]types.YChild)
    laserOffStats.EntityData.Leafs = make(map[string]types.YLeaf)
    laserOffStats.EntityData.Leafs["start"] = types.YLeaf{"Start", laserOffStats.Start}
    laserOffStats.EntityData.Leafs["end"] = types.YLeaf{"End", laserOffStats.End}
    laserOffStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", laserOffStats.TimeTaken}
    laserOffStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", laserOffStats.WorstTime}
    return &(laserOffStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats
// wl op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats struct {
    EntityData types.CommonEntityData
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

func (wlOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_WlOpStats) GetEntityData() *types.CommonEntityData {
    wlOpStats.EntityData.YFilter = wlOpStats.YFilter
    wlOpStats.EntityData.YangName = "wl-op-stats"
    wlOpStats.EntityData.BundleName = "cisco_ios_xr"
    wlOpStats.EntityData.ParentYangName = "port-stat"
    wlOpStats.EntityData.SegmentPath = "wl-op-stats"
    wlOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wlOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wlOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wlOpStats.EntityData.Children = make(map[string]types.YChild)
    wlOpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    wlOpStats.EntityData.Leafs["start"] = types.YLeaf{"Start", wlOpStats.Start}
    wlOpStats.EntityData.Leafs["end"] = types.YLeaf{"End", wlOpStats.End}
    wlOpStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", wlOpStats.TimeTaken}
    wlOpStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", wlOpStats.WorstTime}
    return &(wlOpStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats
// txpwr op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats struct {
    EntityData types.CommonEntityData
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

func (txpwrOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TxpwrOpStats) GetEntityData() *types.CommonEntityData {
    txpwrOpStats.EntityData.YFilter = txpwrOpStats.YFilter
    txpwrOpStats.EntityData.YangName = "txpwr-op-stats"
    txpwrOpStats.EntityData.BundleName = "cisco_ios_xr"
    txpwrOpStats.EntityData.ParentYangName = "port-stat"
    txpwrOpStats.EntityData.SegmentPath = "txpwr-op-stats"
    txpwrOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txpwrOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txpwrOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txpwrOpStats.EntityData.Children = make(map[string]types.YChild)
    txpwrOpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txpwrOpStats.EntityData.Leafs["start"] = types.YLeaf{"Start", txpwrOpStats.Start}
    txpwrOpStats.EntityData.Leafs["end"] = types.YLeaf{"End", txpwrOpStats.End}
    txpwrOpStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", txpwrOpStats.TimeTaken}
    txpwrOpStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", txpwrOpStats.WorstTime}
    return &(txpwrOpStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats
// cdmin op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats struct {
    EntityData types.CommonEntityData
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

func (cdminOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdminOpStats) GetEntityData() *types.CommonEntityData {
    cdminOpStats.EntityData.YFilter = cdminOpStats.YFilter
    cdminOpStats.EntityData.YangName = "cdmin-op-stats"
    cdminOpStats.EntityData.BundleName = "cisco_ios_xr"
    cdminOpStats.EntityData.ParentYangName = "port-stat"
    cdminOpStats.EntityData.SegmentPath = "cdmin-op-stats"
    cdminOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdminOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdminOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdminOpStats.EntityData.Children = make(map[string]types.YChild)
    cdminOpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    cdminOpStats.EntityData.Leafs["start"] = types.YLeaf{"Start", cdminOpStats.Start}
    cdminOpStats.EntityData.Leafs["end"] = types.YLeaf{"End", cdminOpStats.End}
    cdminOpStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", cdminOpStats.TimeTaken}
    cdminOpStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", cdminOpStats.WorstTime}
    return &(cdminOpStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats
// cdmax op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats struct {
    EntityData types.CommonEntityData
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

func (cdmaxOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_CdmaxOpStats) GetEntityData() *types.CommonEntityData {
    cdmaxOpStats.EntityData.YFilter = cdmaxOpStats.YFilter
    cdmaxOpStats.EntityData.YangName = "cdmax-op-stats"
    cdmaxOpStats.EntityData.BundleName = "cisco_ios_xr"
    cdmaxOpStats.EntityData.ParentYangName = "port-stat"
    cdmaxOpStats.EntityData.SegmentPath = "cdmax-op-stats"
    cdmaxOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdmaxOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdmaxOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdmaxOpStats.EntityData.Children = make(map[string]types.YChild)
    cdmaxOpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    cdmaxOpStats.EntityData.Leafs["start"] = types.YLeaf{"Start", cdmaxOpStats.Start}
    cdmaxOpStats.EntityData.Leafs["end"] = types.YLeaf{"End", cdmaxOpStats.End}
    cdmaxOpStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", cdmaxOpStats.TimeTaken}
    cdmaxOpStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", cdmaxOpStats.WorstTime}
    return &(cdmaxOpStats.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats
// traffictype op stats
type Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats struct {
    EntityData types.CommonEntityData
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

func (traffictypeOpStats *Coherent_Nodes_Node_CoherentTimeStats_PortStat_TraffictypeOpStats) GetEntityData() *types.CommonEntityData {
    traffictypeOpStats.EntityData.YFilter = traffictypeOpStats.YFilter
    traffictypeOpStats.EntityData.YangName = "traffictype-op-stats"
    traffictypeOpStats.EntityData.BundleName = "cisco_ios_xr"
    traffictypeOpStats.EntityData.ParentYangName = "port-stat"
    traffictypeOpStats.EntityData.SegmentPath = "traffictype-op-stats"
    traffictypeOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffictypeOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffictypeOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffictypeOpStats.EntityData.Children = make(map[string]types.YChild)
    traffictypeOpStats.EntityData.Leafs = make(map[string]types.YLeaf)
    traffictypeOpStats.EntityData.Leafs["start"] = types.YLeaf{"Start", traffictypeOpStats.Start}
    traffictypeOpStats.EntityData.Leafs["end"] = types.YLeaf{"End", traffictypeOpStats.End}
    traffictypeOpStats.EntityData.Leafs["time-taken"] = types.YLeaf{"TimeTaken", traffictypeOpStats.TimeTaken}
    traffictypeOpStats.EntityData.Leafs["worst-time"] = types.YLeaf{"WorstTime", traffictypeOpStats.WorstTime}
    return &(traffictypeOpStats.EntityData)
}

// Coherent_Nodes_Node_Devicemapping
// Coherent node data for device _mapping
type Coherent_Nodes_Node_Devicemapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // idx. The type is interface{} with range: 0..4294967295.
    Idx interface{}

    // dev map. The type is slice of Coherent_Nodes_Node_Devicemapping_DevMap.
    DevMap []Coherent_Nodes_Node_Devicemapping_DevMap
}

func (devicemapping *Coherent_Nodes_Node_Devicemapping) GetEntityData() *types.CommonEntityData {
    devicemapping.EntityData.YFilter = devicemapping.YFilter
    devicemapping.EntityData.YangName = "devicemapping"
    devicemapping.EntityData.BundleName = "cisco_ios_xr"
    devicemapping.EntityData.ParentYangName = "node"
    devicemapping.EntityData.SegmentPath = "devicemapping"
    devicemapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devicemapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devicemapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devicemapping.EntityData.Children = make(map[string]types.YChild)
    devicemapping.EntityData.Children["dev-map"] = types.YChild{"DevMap", nil}
    for i := range devicemapping.DevMap {
        devicemapping.EntityData.Children[types.GetSegmentPath(&devicemapping.DevMap[i])] = types.YChild{"DevMap", &devicemapping.DevMap[i]}
    }
    devicemapping.EntityData.Leafs = make(map[string]types.YLeaf)
    devicemapping.EntityData.Leafs["idx"] = types.YLeaf{"Idx", devicemapping.Idx}
    return &(devicemapping.EntityData)
}

// Coherent_Nodes_Node_Devicemapping_DevMap
// dev map
type Coherent_Nodes_Node_Devicemapping_DevMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // device address. The type is string with length: 0..128.
    DeviceAddress interface{}

    // ifhandle. The type is string with length: 0..128.
    Ifhandle interface{}

    // intf name. The type is string with length: 0..64.
    IntfName interface{}
}

func (devMap *Coherent_Nodes_Node_Devicemapping_DevMap) GetEntityData() *types.CommonEntityData {
    devMap.EntityData.YFilter = devMap.YFilter
    devMap.EntityData.YangName = "dev-map"
    devMap.EntityData.BundleName = "cisco_ios_xr"
    devMap.EntityData.ParentYangName = "devicemapping"
    devMap.EntityData.SegmentPath = "dev-map"
    devMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devMap.EntityData.Children = make(map[string]types.YChild)
    devMap.EntityData.Leafs = make(map[string]types.YLeaf)
    devMap.EntityData.Leafs["device-address"] = types.YLeaf{"DeviceAddress", devMap.DeviceAddress}
    devMap.EntityData.Leafs["ifhandle"] = types.YLeaf{"Ifhandle", devMap.Ifhandle}
    devMap.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", devMap.IntfName}
    return &(devMap.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth
// Coherent node data for driver health
type Coherent_Nodes_Node_Coherenthealth struct {
    EntityData types.CommonEntityData
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

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetEntityData() *types.CommonEntityData {
    coherenthealth.EntityData.YFilter = coherenthealth.YFilter
    coherenthealth.EntityData.YangName = "coherenthealth"
    coherenthealth.EntityData.BundleName = "cisco_ios_xr"
    coherenthealth.EntityData.ParentYangName = "node"
    coherenthealth.EntityData.SegmentPath = "coherenthealth"
    coherenthealth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherenthealth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherenthealth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherenthealth.EntityData.Children = make(map[string]types.YChild)
    coherenthealth.EntityData.Children["port-data"] = types.YChild{"PortData", nil}
    for i := range coherenthealth.PortData {
        coherenthealth.EntityData.Children[types.GetSegmentPath(&coherenthealth.PortData[i])] = types.YChild{"PortData", &coherenthealth.PortData[i]}
    }
    coherenthealth.EntityData.Leafs = make(map[string]types.YLeaf)
    coherenthealth.EntityData.Leafs["im-state"] = types.YLeaf{"ImState", coherenthealth.ImState}
    coherenthealth.EntityData.Leafs["aipc-srvr-state"] = types.YLeaf{"AipcSrvrState", coherenthealth.AipcSrvrState}
    coherenthealth.EntityData.Leafs["sysdb-state"] = types.YLeaf{"SysdbState", coherenthealth.SysdbState}
    coherenthealth.EntityData.Leafs["pm-state"] = types.YLeaf{"PmState", coherenthealth.PmState}
    coherenthealth.EntityData.Leafs["optics-ea-conn"] = types.YLeaf{"OpticsEaConn", coherenthealth.OpticsEaConn}
    coherenthealth.EntityData.Leafs["dsp-ea-conn"] = types.YLeaf{"DspEaConn", coherenthealth.DspEaConn}
    coherenthealth.EntityData.Leafs["vether-state"] = types.YLeaf{"VetherState", coherenthealth.VetherState}
    coherenthealth.EntityData.Leafs["morgoth-alive"] = types.YLeaf{"MorgothAlive", coherenthealth.MorgothAlive}
    coherenthealth.EntityData.Leafs["prov-infra-state"] = types.YLeaf{"ProvInfraState", coherenthealth.ProvInfraState}
    coherenthealth.EntityData.Leafs["sdk-fpga-compatible"] = types.YLeaf{"SdkFpgaCompatible", coherenthealth.SdkFpgaCompatible}
    coherenthealth.EntityData.Leafs["pending-provision"] = types.YLeaf{"PendingProvision", coherenthealth.PendingProvision}
    coherenthealth.EntityData.Leafs["pulse-sent"] = types.YLeaf{"PulseSent", coherenthealth.PulseSent}
    coherenthealth.EntityData.Leafs["inside-prov-loop"] = types.YLeaf{"InsideProvLoop", coherenthealth.InsideProvLoop}
    coherenthealth.EntityData.Leafs["fpd-in-progress"] = types.YLeaf{"FpdInProgress", coherenthealth.FpdInProgress}
    coherenthealth.EntityData.Leafs["prov-run-count"] = types.YLeaf{"ProvRunCount", coherenthealth.ProvRunCount}
    coherenthealth.EntityData.Leafs["sdk-version"] = types.YLeaf{"SdkVersion", coherenthealth.SdkVersion}
    coherenthealth.EntityData.Leafs["morgoth-running-version"] = types.YLeaf{"MorgothRunningVersion", coherenthealth.MorgothRunningVersion}
    coherenthealth.EntityData.Leafs["morgoth-downloaded-version"] = types.YLeaf{"MorgothDownloadedVersion", coherenthealth.MorgothDownloadedVersion}
    coherenthealth.EntityData.Leafs["morgoth-golden-version"] = types.YLeaf{"MorgothGoldenVersion", coherenthealth.MorgothGoldenVersion}
    coherenthealth.EntityData.Leafs["denali0-version"] = types.YLeaf{"Denali0Version", coherenthealth.Denali0Version}
    coherenthealth.EntityData.Leafs["denali1-version"] = types.YLeaf{"Denali1Version", coherenthealth.Denali1Version}
    coherenthealth.EntityData.Leafs["denali2-version"] = types.YLeaf{"Denali2Version", coherenthealth.Denali2Version}
    coherenthealth.EntityData.Leafs["board-type"] = types.YLeaf{"BoardType", coherenthealth.BoardType}
    coherenthealth.EntityData.Leafs["jlink-op"] = types.YLeaf{"JlinkOp", coherenthealth.JlinkOp}
    return &(coherenthealth.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData
// port data
type Coherent_Nodes_Node_Coherenthealth_PortData struct {
    EntityData types.CommonEntityData
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

func (portData *Coherent_Nodes_Node_Coherenthealth_PortData) GetEntityData() *types.CommonEntityData {
    portData.EntityData.YFilter = portData.YFilter
    portData.EntityData.YangName = "port-data"
    portData.EntityData.BundleName = "cisco_ios_xr"
    portData.EntityData.ParentYangName = "coherenthealth"
    portData.EntityData.SegmentPath = "port-data"
    portData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portData.EntityData.Children = make(map[string]types.YChild)
    portData.EntityData.Children["ctp-info"] = types.YChild{"CtpInfo", &portData.CtpInfo}
    portData.EntityData.Children["interface-info"] = types.YChild{"InterfaceInfo", &portData.InterfaceInfo}
    portData.EntityData.Leafs = make(map[string]types.YLeaf)
    portData.EntityData.Leafs["optics-ctrl-created"] = types.YLeaf{"OpticsCtrlCreated", portData.OpticsCtrlCreated}
    portData.EntityData.Leafs["dsp-ctrl-created"] = types.YLeaf{"DspCtrlCreated", portData.DspCtrlCreated}
    portData.EntityData.Leafs["has-pluggable"] = types.YLeaf{"HasPluggable", portData.HasPluggable}
    portData.EntityData.Leafs["optics-admin-up"] = types.YLeaf{"OpticsAdminUp", portData.OpticsAdminUp}
    portData.EntityData.Leafs["dsp-admin-up"] = types.YLeaf{"DspAdminUp", portData.DspAdminUp}
    portData.EntityData.Leafs["laser-state"] = types.YLeaf{"LaserState", portData.LaserState}
    portData.EntityData.Leafs["laser-on-pending"] = types.YLeaf{"LaserOnPending", portData.LaserOnPending}
    portData.EntityData.Leafs["provisioning-needed"] = types.YLeaf{"ProvisioningNeeded", portData.ProvisioningNeeded}
    portData.EntityData.Leafs["force-reprovision"] = types.YLeaf{"ForceReprovision", portData.ForceReprovision}
    portData.EntityData.Leafs["fp-port-idx"] = types.YLeaf{"FpPortIdx", portData.FpPortIdx}
    portData.EntityData.Leafs["configured-frequency"] = types.YLeaf{"ConfiguredFrequency", portData.ConfiguredFrequency}
    portData.EntityData.Leafs["provisioned-frequency"] = types.YLeaf{"ProvisionedFrequency", portData.ProvisionedFrequency}
    portData.EntityData.Leafs["configured-tx-power"] = types.YLeaf{"ConfiguredTxPower", portData.ConfiguredTxPower}
    portData.EntityData.Leafs["provisioned-tx-power"] = types.YLeaf{"ProvisionedTxPower", portData.ProvisionedTxPower}
    portData.EntityData.Leafs["configured-cd-min"] = types.YLeaf{"ConfiguredCdMin", portData.ConfiguredCdMin}
    portData.EntityData.Leafs["provisioned-cd-min"] = types.YLeaf{"ProvisionedCdMin", portData.ProvisionedCdMin}
    portData.EntityData.Leafs["configured-cd-max"] = types.YLeaf{"ConfiguredCdMax", portData.ConfiguredCdMax}
    portData.EntityData.Leafs["provisioned-cd-max"] = types.YLeaf{"ProvisionedCdMax", portData.ProvisionedCdMax}
    portData.EntityData.Leafs["configured-traffic-type"] = types.YLeaf{"ConfiguredTrafficType", portData.ConfiguredTrafficType}
    portData.EntityData.Leafs["provisioned-traffic-type"] = types.YLeaf{"ProvisionedTrafficType", portData.ProvisionedTrafficType}
    portData.EntityData.Leafs["configured-loopback-mode"] = types.YLeaf{"ConfiguredLoopbackMode", portData.ConfiguredLoopbackMode}
    portData.EntityData.Leafs["provisioned-loopback-mode"] = types.YLeaf{"ProvisionedLoopbackMode", portData.ProvisionedLoopbackMode}
    portData.EntityData.Leafs["expected-ctp2-led-state"] = types.YLeaf{"ExpectedCtp2LedState", portData.ExpectedCtp2LedState}
    portData.EntityData.Leafs["provisioned-ctp2-led-state"] = types.YLeaf{"ProvisionedCtp2LedState", portData.ProvisionedCtp2LedState}
    portData.EntityData.Leafs["led-op-rc"] = types.YLeaf{"LedOpRc", portData.LedOpRc}
    portData.EntityData.Leafs["laser-op-rc"] = types.YLeaf{"LaserOpRc", portData.LaserOpRc}
    portData.EntityData.Leafs["wlen-op-rc"] = types.YLeaf{"WlenOpRc", portData.WlenOpRc}
    portData.EntityData.Leafs["traffic-op-rc"] = types.YLeaf{"TrafficOpRc", portData.TrafficOpRc}
    portData.EntityData.Leafs["loopback-op-rc"] = types.YLeaf{"LoopbackOpRc", portData.LoopbackOpRc}
    portData.EntityData.Leafs["tx-power-op-rc"] = types.YLeaf{"TxPowerOpRc", portData.TxPowerOpRc}
    portData.EntityData.Leafs["cd-min-op-rc"] = types.YLeaf{"CdMinOpRc", portData.CdMinOpRc}
    portData.EntityData.Leafs["cd-max-op-rc"] = types.YLeaf{"CdMaxOpRc", portData.CdMaxOpRc}
    portData.EntityData.Leafs["provisioning-failed"] = types.YLeaf{"ProvisioningFailed", portData.ProvisioningFailed}
    portData.EntityData.Leafs["ctp2-hw-alarms"] = types.YLeaf{"Ctp2HwAlarms", portData.Ctp2HwAlarms}
    portData.EntityData.Leafs["denali-hw-alarms"] = types.YLeaf{"DenaliHwAlarms", portData.DenaliHwAlarms}
    portData.EntityData.Leafs["is-pm-port-created-opt"] = types.YLeaf{"IsPmPortCreatedOpt", portData.IsPmPortCreatedOpt}
    portData.EntityData.Leafs["rc-pm-port-opt"] = types.YLeaf{"RcPmPortOpt", portData.RcPmPortOpt}
    portData.EntityData.Leafs["pm-port-state-opt"] = types.YLeaf{"PmPortStateOpt", portData.PmPortStateOpt}
    portData.EntityData.Leafs["rc-pm-provision-opt"] = types.YLeaf{"RcPmProvisionOpt", portData.RcPmProvisionOpt}
    portData.EntityData.Leafs["is-alarm-port-created-opt"] = types.YLeaf{"IsAlarmPortCreatedOpt", portData.IsAlarmPortCreatedOpt}
    portData.EntityData.Leafs["rc-alarm-port-opt"] = types.YLeaf{"RcAlarmPortOpt", portData.RcAlarmPortOpt}
    portData.EntityData.Leafs["is-pm-port-created-dsp"] = types.YLeaf{"IsPmPortCreatedDsp", portData.IsPmPortCreatedDsp}
    portData.EntityData.Leafs["rc-pm-port-dsp"] = types.YLeaf{"RcPmPortDsp", portData.RcPmPortDsp}
    portData.EntityData.Leafs["pm-port-state-dsp"] = types.YLeaf{"PmPortStateDsp", portData.PmPortStateDsp}
    portData.EntityData.Leafs["rc-pm-provision-dsp"] = types.YLeaf{"RcPmProvisionDsp", portData.RcPmProvisionDsp}
    portData.EntityData.Leafs["is-alarm-port-created-dsp"] = types.YLeaf{"IsAlarmPortCreatedDsp", portData.IsAlarmPortCreatedDsp}
    portData.EntityData.Leafs["rc-alarm-port-dsp"] = types.YLeaf{"RcAlarmPortDsp", portData.RcAlarmPortDsp}
    return &(portData.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo
// ctp info
type Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo struct {
    EntityData types.CommonEntityData
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

func (ctpInfo *Coherent_Nodes_Node_Coherenthealth_PortData_CtpInfo) GetEntityData() *types.CommonEntityData {
    ctpInfo.EntityData.YFilter = ctpInfo.YFilter
    ctpInfo.EntityData.YangName = "ctp-info"
    ctpInfo.EntityData.BundleName = "cisco_ios_xr"
    ctpInfo.EntityData.ParentYangName = "port-data"
    ctpInfo.EntityData.SegmentPath = "ctp-info"
    ctpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ctpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ctpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ctpInfo.EntityData.Children = make(map[string]types.YChild)
    ctpInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ctpInfo.EntityData.Leafs["deviation"] = types.YLeaf{"Deviation", ctpInfo.Deviation}
    ctpInfo.EntityData.Leafs["part-number"] = types.YLeaf{"PartNumber", ctpInfo.PartNumber}
    ctpInfo.EntityData.Leafs["serial-number"] = types.YLeaf{"SerialNumber", ctpInfo.SerialNumber}
    ctpInfo.EntityData.Leafs["date-code-number"] = types.YLeaf{"DateCodeNumber", ctpInfo.DateCodeNumber}
    ctpInfo.EntityData.Leafs["clei-code-number"] = types.YLeaf{"CleiCodeNumber", ctpInfo.CleiCodeNumber}
    ctpInfo.EntityData.Leafs["vendorname"] = types.YLeaf{"Vendorname", ctpInfo.Vendorname}
    ctpInfo.EntityData.Leafs["description"] = types.YLeaf{"Description", ctpInfo.Description}
    ctpInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", ctpInfo.Pid}
    ctpInfo.EntityData.Leafs["vid"] = types.YLeaf{"Vid", ctpInfo.Vid}
    ctpInfo.EntityData.Leafs["module-hardware-version-number"] = types.YLeaf{"ModuleHardwareVersionNumber", ctpInfo.ModuleHardwareVersionNumber}
    ctpInfo.EntityData.Leafs["module-firmware-running-version-number"] = types.YLeaf{"ModuleFirmwareRunningVersionNumber", ctpInfo.ModuleFirmwareRunningVersionNumber}
    ctpInfo.EntityData.Leafs["module-firmware-committed-version-number"] = types.YLeaf{"ModuleFirmwareCommittedVersionNumber", ctpInfo.ModuleFirmwareCommittedVersionNumber}
    ctpInfo.EntityData.Leafs["ctp-type"] = types.YLeaf{"CtpType", ctpInfo.CtpType}
    return &(ctpInfo.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo
// interface info
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // eth data. The type is slice of
    // Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData.
    EthData []Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetEntityData() *types.CommonEntityData {
    interfaceInfo.EntityData.YFilter = interfaceInfo.YFilter
    interfaceInfo.EntityData.YangName = "interface-info"
    interfaceInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfo.EntityData.ParentYangName = "port-data"
    interfaceInfo.EntityData.SegmentPath = "interface-info"
    interfaceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfo.EntityData.Children = make(map[string]types.YChild)
    interfaceInfo.EntityData.Children["eth-data"] = types.YChild{"EthData", nil}
    for i := range interfaceInfo.EthData {
        interfaceInfo.EntityData.Children[types.GetSegmentPath(&interfaceInfo.EthData[i])] = types.YChild{"EthData", &interfaceInfo.EthData[i]}
    }
    interfaceInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceInfo.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
// eth data
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData struct {
    EntityData types.CommonEntityData
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

func (ethData *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData) GetEntityData() *types.CommonEntityData {
    ethData.EntityData.YFilter = ethData.YFilter
    ethData.EntityData.YangName = "eth-data"
    ethData.EntityData.BundleName = "cisco_ios_xr"
    ethData.EntityData.ParentYangName = "interface-info"
    ethData.EntityData.SegmentPath = "eth-data"
    ethData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethData.EntityData.Children = make(map[string]types.YChild)
    ethData.EntityData.Leafs = make(map[string]types.YLeaf)
    ethData.EntityData.Leafs["ifname"] = types.YLeaf{"Ifname", ethData.Ifname}
    ethData.EntityData.Leafs["intf-handle"] = types.YLeaf{"IntfHandle", ethData.IntfHandle}
    ethData.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", ethData.AdminState}
    ethData.EntityData.Leafs["admin-up"] = types.YLeaf{"AdminUp", ethData.AdminUp}
    ethData.EntityData.Leafs["is-created"] = types.YLeaf{"IsCreated", ethData.IsCreated}
    return &(ethData.EntityData)
}

// Coherent_Nodes_Node_PortModeAllInfo
// PortMode all operational data
type Coherent_Nodes_Node_PortModeAllInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dev map idx. The type is interface{} with range: 0..4294967295.
    Idx interface{}

    // portmode entry. The type is slice of
    // Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry.
    PortmodeEntry []Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry
}

func (portModeAllInfo *Coherent_Nodes_Node_PortModeAllInfo) GetEntityData() *types.CommonEntityData {
    portModeAllInfo.EntityData.YFilter = portModeAllInfo.YFilter
    portModeAllInfo.EntityData.YangName = "port-mode-all-info"
    portModeAllInfo.EntityData.BundleName = "cisco_ios_xr"
    portModeAllInfo.EntityData.ParentYangName = "node"
    portModeAllInfo.EntityData.SegmentPath = "port-mode-all-info"
    portModeAllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portModeAllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portModeAllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portModeAllInfo.EntityData.Children = make(map[string]types.YChild)
    portModeAllInfo.EntityData.Children["portmode-entry"] = types.YChild{"PortmodeEntry", nil}
    for i := range portModeAllInfo.PortmodeEntry {
        portModeAllInfo.EntityData.Children[types.GetSegmentPath(&portModeAllInfo.PortmodeEntry[i])] = types.YChild{"PortmodeEntry", &portModeAllInfo.PortmodeEntry[i]}
    }
    portModeAllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    portModeAllInfo.EntityData.Leafs["idx"] = types.YLeaf{"Idx", portModeAllInfo.Idx}
    return &(portModeAllInfo.EntityData)
}

// Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry
// portmode entry
type Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry struct {
    EntityData types.CommonEntityData
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

func (portmodeEntry *Coherent_Nodes_Node_PortModeAllInfo_PortmodeEntry) GetEntityData() *types.CommonEntityData {
    portmodeEntry.EntityData.YFilter = portmodeEntry.YFilter
    portmodeEntry.EntityData.YangName = "portmode-entry"
    portmodeEntry.EntityData.BundleName = "cisco_ios_xr"
    portmodeEntry.EntityData.ParentYangName = "port-mode-all-info"
    portmodeEntry.EntityData.SegmentPath = "portmode-entry"
    portmodeEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portmodeEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portmodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portmodeEntry.EntityData.Children = make(map[string]types.YChild)
    portmodeEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    portmodeEntry.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", portmodeEntry.IntfName}
    portmodeEntry.EntityData.Leafs["speed"] = types.YLeaf{"Speed", portmodeEntry.Speed}
    portmodeEntry.EntityData.Leafs["fec"] = types.YLeaf{"Fec", portmodeEntry.Fec}
    portmodeEntry.EntityData.Leafs["diff"] = types.YLeaf{"Diff", portmodeEntry.Diff}
    portmodeEntry.EntityData.Leafs["modulation"] = types.YLeaf{"Modulation", portmodeEntry.Modulation}
    return &(portmodeEntry.EntityData)
}

