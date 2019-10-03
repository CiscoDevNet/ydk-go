// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-node package operational data.
// 
// This module contains definitions
// for the following management objects:
//   coherent: Coherent node  operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    coherent.EntityData.AbsolutePath = coherent.EntityData.SegmentPath
    coherent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherent.EntityData.Children = types.NewOrderedMap()
    coherent.EntityData.Children.Append("nodes", types.YChild{"Nodes", &coherent.Nodes})
    coherent.EntityData.Leafs = types.NewOrderedMap()

    coherent.EntityData.YListKeys = []string {}

    return &(coherent.EntityData)
}

// Coherent_Nodes
// Coherent list of nodes
type Coherent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coherent discovery operational data for a particular node. The type is
    // slice of Coherent_Nodes_Node.
    Node []*Coherent_Nodes_Node
}

func (nodes *Coherent_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "coherent"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/" + nodes.EntityData.SegmentPath
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

// Coherent_Nodes_Node
// Coherent discovery operational data for a
// particular node
type Coherent_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Coherent driver performace information.
    CoherentTimeStats Coherent_Nodes_Node_CoherentTimeStats

    // Coherent node data for driver health.
    Coherenthealth Coherent_Nodes_Node_Coherenthealth
}

func (node *Coherent_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("coherent-time-stats", types.YChild{"CoherentTimeStats", &node.CoherentTimeStats})
    node.EntityData.Children.Append("coherenthealth", types.YChild{"Coherenthealth", &node.Coherenthealth})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

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
    PortStat []*Coherent_Nodes_Node_CoherentTimeStats_PortStat
}

func (coherentTimeStats *Coherent_Nodes_Node_CoherentTimeStats) GetEntityData() *types.CommonEntityData {
    coherentTimeStats.EntityData.YFilter = coherentTimeStats.YFilter
    coherentTimeStats.EntityData.YangName = "coherent-time-stats"
    coherentTimeStats.EntityData.BundleName = "cisco_ios_xr"
    coherentTimeStats.EntityData.ParentYangName = "node"
    coherentTimeStats.EntityData.SegmentPath = "coherent-time-stats"
    coherentTimeStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/" + coherentTimeStats.EntityData.SegmentPath
    coherentTimeStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherentTimeStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherentTimeStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherentTimeStats.EntityData.Children = types.NewOrderedMap()
    coherentTimeStats.EntityData.Children.Append("opts-ea-bulk-create", types.YChild{"OptsEaBulkCreate", &coherentTimeStats.OptsEaBulkCreate})
    coherentTimeStats.EntityData.Children.Append("opts-ea-bulk-update", types.YChild{"OptsEaBulkUpdate", &coherentTimeStats.OptsEaBulkUpdate})
    coherentTimeStats.EntityData.Children.Append("dsp-ea-bulk-create", types.YChild{"DspEaBulkCreate", &coherentTimeStats.DspEaBulkCreate})
    coherentTimeStats.EntityData.Children.Append("dsp-ea-bulk-update", types.YChild{"DspEaBulkUpdate", &coherentTimeStats.DspEaBulkUpdate})
    coherentTimeStats.EntityData.Children.Append("port-stat", types.YChild{"PortStat", nil})
    for i := range coherentTimeStats.PortStat {
        types.SetYListKey(coherentTimeStats.PortStat[i], i)
        coherentTimeStats.EntityData.Children.Append(types.GetSegmentPath(coherentTimeStats.PortStat[i]), types.YChild{"PortStat", coherentTimeStats.PortStat[i]})
    }
    coherentTimeStats.EntityData.Leafs = types.NewOrderedMap()
    coherentTimeStats.EntityData.Leafs.Append("driver-init", types.YLeaf{"DriverInit", coherentTimeStats.DriverInit})
    coherentTimeStats.EntityData.Leafs.Append("driver-operational", types.YLeaf{"DriverOperational", coherentTimeStats.DriverOperational})
    coherentTimeStats.EntityData.Leafs.Append("device-created", types.YLeaf{"DeviceCreated", coherentTimeStats.DeviceCreated})
    coherentTimeStats.EntityData.Leafs.Append("optics-controllers-created", types.YLeaf{"OpticsControllersCreated", coherentTimeStats.OpticsControllersCreated})
    coherentTimeStats.EntityData.Leafs.Append("dsp-controllers-created", types.YLeaf{"DspControllersCreated", coherentTimeStats.DspControllersCreated})
    coherentTimeStats.EntityData.Leafs.Append("eth-intf-created", types.YLeaf{"EthIntfCreated", coherentTimeStats.EthIntfCreated})

    coherentTimeStats.EntityData.YListKeys = []string {}

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
    optsEaBulkCreate.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/" + optsEaBulkCreate.EntityData.SegmentPath
    optsEaBulkCreate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optsEaBulkCreate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optsEaBulkCreate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optsEaBulkCreate.EntityData.Children = types.NewOrderedMap()
    optsEaBulkCreate.EntityData.Leafs = types.NewOrderedMap()
    optsEaBulkCreate.EntityData.Leafs.Append("start", types.YLeaf{"Start", optsEaBulkCreate.Start})
    optsEaBulkCreate.EntityData.Leafs.Append("end", types.YLeaf{"End", optsEaBulkCreate.End})
    optsEaBulkCreate.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", optsEaBulkCreate.TimeTaken})
    optsEaBulkCreate.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", optsEaBulkCreate.WorstTime})

    optsEaBulkCreate.EntityData.YListKeys = []string {}

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
    optsEaBulkUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/" + optsEaBulkUpdate.EntityData.SegmentPath
    optsEaBulkUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optsEaBulkUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optsEaBulkUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optsEaBulkUpdate.EntityData.Children = types.NewOrderedMap()
    optsEaBulkUpdate.EntityData.Leafs = types.NewOrderedMap()
    optsEaBulkUpdate.EntityData.Leafs.Append("start", types.YLeaf{"Start", optsEaBulkUpdate.Start})
    optsEaBulkUpdate.EntityData.Leafs.Append("end", types.YLeaf{"End", optsEaBulkUpdate.End})
    optsEaBulkUpdate.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", optsEaBulkUpdate.TimeTaken})
    optsEaBulkUpdate.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", optsEaBulkUpdate.WorstTime})

    optsEaBulkUpdate.EntityData.YListKeys = []string {}

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
    dspEaBulkCreate.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/" + dspEaBulkCreate.EntityData.SegmentPath
    dspEaBulkCreate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dspEaBulkCreate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dspEaBulkCreate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dspEaBulkCreate.EntityData.Children = types.NewOrderedMap()
    dspEaBulkCreate.EntityData.Leafs = types.NewOrderedMap()
    dspEaBulkCreate.EntityData.Leafs.Append("start", types.YLeaf{"Start", dspEaBulkCreate.Start})
    dspEaBulkCreate.EntityData.Leafs.Append("end", types.YLeaf{"End", dspEaBulkCreate.End})
    dspEaBulkCreate.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", dspEaBulkCreate.TimeTaken})
    dspEaBulkCreate.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", dspEaBulkCreate.WorstTime})

    dspEaBulkCreate.EntityData.YListKeys = []string {}

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
    dspEaBulkUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/" + dspEaBulkUpdate.EntityData.SegmentPath
    dspEaBulkUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dspEaBulkUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dspEaBulkUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dspEaBulkUpdate.EntityData.Children = types.NewOrderedMap()
    dspEaBulkUpdate.EntityData.Leafs = types.NewOrderedMap()
    dspEaBulkUpdate.EntityData.Leafs.Append("start", types.YLeaf{"Start", dspEaBulkUpdate.Start})
    dspEaBulkUpdate.EntityData.Leafs.Append("end", types.YLeaf{"End", dspEaBulkUpdate.End})
    dspEaBulkUpdate.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", dspEaBulkUpdate.TimeTaken})
    dspEaBulkUpdate.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", dspEaBulkUpdate.WorstTime})

    dspEaBulkUpdate.EntityData.YListKeys = []string {}

    return &(dspEaBulkUpdate.EntityData)
}

// Coherent_Nodes_Node_CoherentTimeStats_PortStat
// port stat
type Coherent_Nodes_Node_CoherentTimeStats_PortStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // rsip. The type is interface{} with range: 0..4294967295.
    Rsip interface{}

    // fp port idx. The type is interface{} with range: 0..4294967295.
    FpPortIdx interface{}

    // laser state. The type is bool.
    LaserState interface{}

    // provisioned frequency. The type is interface{} with range: 0..4294967295.
    ProvisionedFrequency interface{}

    // tx power. The type is interface{} with range: -2147483648..2147483647.
    TxPower interface{}

    // cd min. The type is interface{} with range: -2147483648..2147483647.
    CdMin interface{}

    // cd max. The type is interface{} with range: -2147483648..2147483647.
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
    portStat.EntityData.SegmentPath = "port-stat" + types.AddNoKeyToken(portStat)
    portStat.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/" + portStat.EntityData.SegmentPath
    portStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStat.EntityData.Children = types.NewOrderedMap()
    portStat.EntityData.Children.Append("laser-on-stats", types.YChild{"LaserOnStats", &portStat.LaserOnStats})
    portStat.EntityData.Children.Append("laser-off-stats", types.YChild{"LaserOffStats", &portStat.LaserOffStats})
    portStat.EntityData.Children.Append("wl-op-stats", types.YChild{"WlOpStats", &portStat.WlOpStats})
    portStat.EntityData.Children.Append("txpwr-op-stats", types.YChild{"TxpwrOpStats", &portStat.TxpwrOpStats})
    portStat.EntityData.Children.Append("cdmin-op-stats", types.YChild{"CdminOpStats", &portStat.CdminOpStats})
    portStat.EntityData.Children.Append("cdmax-op-stats", types.YChild{"CdmaxOpStats", &portStat.CdmaxOpStats})
    portStat.EntityData.Children.Append("traffictype-op-stats", types.YChild{"TraffictypeOpStats", &portStat.TraffictypeOpStats})
    portStat.EntityData.Leafs = types.NewOrderedMap()
    portStat.EntityData.Leafs.Append("rsip", types.YLeaf{"Rsip", portStat.Rsip})
    portStat.EntityData.Leafs.Append("fp-port-idx", types.YLeaf{"FpPortIdx", portStat.FpPortIdx})
    portStat.EntityData.Leafs.Append("laser-state", types.YLeaf{"LaserState", portStat.LaserState})
    portStat.EntityData.Leafs.Append("provisioned-frequency", types.YLeaf{"ProvisionedFrequency", portStat.ProvisionedFrequency})
    portStat.EntityData.Leafs.Append("tx-power", types.YLeaf{"TxPower", portStat.TxPower})
    portStat.EntityData.Leafs.Append("cd-min", types.YLeaf{"CdMin", portStat.CdMin})
    portStat.EntityData.Leafs.Append("cd-max", types.YLeaf{"CdMax", portStat.CdMax})
    portStat.EntityData.Leafs.Append("traffic-type", types.YLeaf{"TrafficType", portStat.TrafficType})

    portStat.EntityData.YListKeys = []string {}

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
    laserOnStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + laserOnStats.EntityData.SegmentPath
    laserOnStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laserOnStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laserOnStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laserOnStats.EntityData.Children = types.NewOrderedMap()
    laserOnStats.EntityData.Leafs = types.NewOrderedMap()
    laserOnStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", laserOnStats.Start})
    laserOnStats.EntityData.Leafs.Append("end", types.YLeaf{"End", laserOnStats.End})
    laserOnStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", laserOnStats.TimeTaken})
    laserOnStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", laserOnStats.WorstTime})

    laserOnStats.EntityData.YListKeys = []string {}

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
    laserOffStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + laserOffStats.EntityData.SegmentPath
    laserOffStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    laserOffStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    laserOffStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    laserOffStats.EntityData.Children = types.NewOrderedMap()
    laserOffStats.EntityData.Leafs = types.NewOrderedMap()
    laserOffStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", laserOffStats.Start})
    laserOffStats.EntityData.Leafs.Append("end", types.YLeaf{"End", laserOffStats.End})
    laserOffStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", laserOffStats.TimeTaken})
    laserOffStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", laserOffStats.WorstTime})

    laserOffStats.EntityData.YListKeys = []string {}

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
    wlOpStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + wlOpStats.EntityData.SegmentPath
    wlOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wlOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wlOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wlOpStats.EntityData.Children = types.NewOrderedMap()
    wlOpStats.EntityData.Leafs = types.NewOrderedMap()
    wlOpStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", wlOpStats.Start})
    wlOpStats.EntityData.Leafs.Append("end", types.YLeaf{"End", wlOpStats.End})
    wlOpStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", wlOpStats.TimeTaken})
    wlOpStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", wlOpStats.WorstTime})

    wlOpStats.EntityData.YListKeys = []string {}

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
    txpwrOpStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + txpwrOpStats.EntityData.SegmentPath
    txpwrOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txpwrOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txpwrOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txpwrOpStats.EntityData.Children = types.NewOrderedMap()
    txpwrOpStats.EntityData.Leafs = types.NewOrderedMap()
    txpwrOpStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", txpwrOpStats.Start})
    txpwrOpStats.EntityData.Leafs.Append("end", types.YLeaf{"End", txpwrOpStats.End})
    txpwrOpStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", txpwrOpStats.TimeTaken})
    txpwrOpStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", txpwrOpStats.WorstTime})

    txpwrOpStats.EntityData.YListKeys = []string {}

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
    cdminOpStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + cdminOpStats.EntityData.SegmentPath
    cdminOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdminOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdminOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdminOpStats.EntityData.Children = types.NewOrderedMap()
    cdminOpStats.EntityData.Leafs = types.NewOrderedMap()
    cdminOpStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", cdminOpStats.Start})
    cdminOpStats.EntityData.Leafs.Append("end", types.YLeaf{"End", cdminOpStats.End})
    cdminOpStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", cdminOpStats.TimeTaken})
    cdminOpStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", cdminOpStats.WorstTime})

    cdminOpStats.EntityData.YListKeys = []string {}

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
    cdmaxOpStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + cdmaxOpStats.EntityData.SegmentPath
    cdmaxOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cdmaxOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cdmaxOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cdmaxOpStats.EntityData.Children = types.NewOrderedMap()
    cdmaxOpStats.EntityData.Leafs = types.NewOrderedMap()
    cdmaxOpStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", cdmaxOpStats.Start})
    cdmaxOpStats.EntityData.Leafs.Append("end", types.YLeaf{"End", cdmaxOpStats.End})
    cdmaxOpStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", cdmaxOpStats.TimeTaken})
    cdmaxOpStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", cdmaxOpStats.WorstTime})

    cdmaxOpStats.EntityData.YListKeys = []string {}

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
    traffictypeOpStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherent-time-stats/port-stat/" + traffictypeOpStats.EntityData.SegmentPath
    traffictypeOpStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffictypeOpStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffictypeOpStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffictypeOpStats.EntityData.Children = types.NewOrderedMap()
    traffictypeOpStats.EntityData.Leafs = types.NewOrderedMap()
    traffictypeOpStats.EntityData.Leafs.Append("start", types.YLeaf{"Start", traffictypeOpStats.Start})
    traffictypeOpStats.EntityData.Leafs.Append("end", types.YLeaf{"End", traffictypeOpStats.End})
    traffictypeOpStats.EntityData.Leafs.Append("time-taken", types.YLeaf{"TimeTaken", traffictypeOpStats.TimeTaken})
    traffictypeOpStats.EntityData.Leafs.Append("worst-time", types.YLeaf{"WorstTime", traffictypeOpStats.WorstTime})

    traffictypeOpStats.EntityData.YListKeys = []string {}

    return &(traffictypeOpStats.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth
// Coherent node data for driver health
type Coherent_Nodes_Node_Coherenthealth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // is peyto. The type is bool.
    IsPeyto interface{}

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

    // dsp0 version. The type is string with length: 0..255.
    Dsp0Version interface{}

    // dsp1 version. The type is string with length: 0..255.
    Dsp1Version interface{}

    // dsp2 version. The type is string with length: 0..255.
    Dsp2Version interface{}

    // board type. The type is string with length: 0..255.
    BoardType interface{}

    // jlink op. The type is string with length: 0..6144.
    JlinkOp interface{}

    // port data. The type is slice of
    // Coherent_Nodes_Node_Coherenthealth_PortData.
    PortData []*Coherent_Nodes_Node_Coherenthealth_PortData
}

func (coherenthealth *Coherent_Nodes_Node_Coherenthealth) GetEntityData() *types.CommonEntityData {
    coherenthealth.EntityData.YFilter = coherenthealth.YFilter
    coherenthealth.EntityData.YangName = "coherenthealth"
    coherenthealth.EntityData.BundleName = "cisco_ios_xr"
    coherenthealth.EntityData.ParentYangName = "node"
    coherenthealth.EntityData.SegmentPath = "coherenthealth"
    coherenthealth.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/" + coherenthealth.EntityData.SegmentPath
    coherenthealth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coherenthealth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coherenthealth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coherenthealth.EntityData.Children = types.NewOrderedMap()
    coherenthealth.EntityData.Children.Append("port-data", types.YChild{"PortData", nil})
    for i := range coherenthealth.PortData {
        types.SetYListKey(coherenthealth.PortData[i], i)
        coherenthealth.EntityData.Children.Append(types.GetSegmentPath(coherenthealth.PortData[i]), types.YChild{"PortData", coherenthealth.PortData[i]})
    }
    coherenthealth.EntityData.Leafs = types.NewOrderedMap()
    coherenthealth.EntityData.Leafs.Append("is-peyto", types.YLeaf{"IsPeyto", coherenthealth.IsPeyto})
    coherenthealth.EntityData.Leafs.Append("im-state", types.YLeaf{"ImState", coherenthealth.ImState})
    coherenthealth.EntityData.Leafs.Append("aipc-srvr-state", types.YLeaf{"AipcSrvrState", coherenthealth.AipcSrvrState})
    coherenthealth.EntityData.Leafs.Append("sysdb-state", types.YLeaf{"SysdbState", coherenthealth.SysdbState})
    coherenthealth.EntityData.Leafs.Append("pm-state", types.YLeaf{"PmState", coherenthealth.PmState})
    coherenthealth.EntityData.Leafs.Append("optics-ea-conn", types.YLeaf{"OpticsEaConn", coherenthealth.OpticsEaConn})
    coherenthealth.EntityData.Leafs.Append("dsp-ea-conn", types.YLeaf{"DspEaConn", coherenthealth.DspEaConn})
    coherenthealth.EntityData.Leafs.Append("vether-state", types.YLeaf{"VetherState", coherenthealth.VetherState})
    coherenthealth.EntityData.Leafs.Append("morgoth-alive", types.YLeaf{"MorgothAlive", coherenthealth.MorgothAlive})
    coherenthealth.EntityData.Leafs.Append("prov-infra-state", types.YLeaf{"ProvInfraState", coherenthealth.ProvInfraState})
    coherenthealth.EntityData.Leafs.Append("sdk-fpga-compatible", types.YLeaf{"SdkFpgaCompatible", coherenthealth.SdkFpgaCompatible})
    coherenthealth.EntityData.Leafs.Append("pending-provision", types.YLeaf{"PendingProvision", coherenthealth.PendingProvision})
    coherenthealth.EntityData.Leafs.Append("pulse-sent", types.YLeaf{"PulseSent", coherenthealth.PulseSent})
    coherenthealth.EntityData.Leafs.Append("inside-prov-loop", types.YLeaf{"InsideProvLoop", coherenthealth.InsideProvLoop})
    coherenthealth.EntityData.Leafs.Append("fpd-in-progress", types.YLeaf{"FpdInProgress", coherenthealth.FpdInProgress})
    coherenthealth.EntityData.Leafs.Append("prov-run-count", types.YLeaf{"ProvRunCount", coherenthealth.ProvRunCount})
    coherenthealth.EntityData.Leafs.Append("sdk-version", types.YLeaf{"SdkVersion", coherenthealth.SdkVersion})
    coherenthealth.EntityData.Leafs.Append("morgoth-running-version", types.YLeaf{"MorgothRunningVersion", coherenthealth.MorgothRunningVersion})
    coherenthealth.EntityData.Leafs.Append("morgoth-downloaded-version", types.YLeaf{"MorgothDownloadedVersion", coherenthealth.MorgothDownloadedVersion})
    coherenthealth.EntityData.Leafs.Append("morgoth-golden-version", types.YLeaf{"MorgothGoldenVersion", coherenthealth.MorgothGoldenVersion})
    coherenthealth.EntityData.Leafs.Append("dsp0-version", types.YLeaf{"Dsp0Version", coherenthealth.Dsp0Version})
    coherenthealth.EntityData.Leafs.Append("dsp1-version", types.YLeaf{"Dsp1Version", coherenthealth.Dsp1Version})
    coherenthealth.EntityData.Leafs.Append("dsp2-version", types.YLeaf{"Dsp2Version", coherenthealth.Dsp2Version})
    coherenthealth.EntityData.Leafs.Append("board-type", types.YLeaf{"BoardType", coherenthealth.BoardType})
    coherenthealth.EntityData.Leafs.Append("jlink-op", types.YLeaf{"JlinkOp", coherenthealth.JlinkOp})

    coherenthealth.EntityData.YListKeys = []string {}

    return &(coherenthealth.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData
// port data
type Coherent_Nodes_Node_Coherenthealth_PortData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // rsip. The type is interface{} with range: 0..4294967295.
    Rsip interface{}

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

    // end of config. The type is bool.
    EndOfConfig interface{}

    // optics laser state. The type is bool.
    OpticsLaserState interface{}

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

    // dsp hw alarms. The type is string with length: 0..256.
    DspHwAlarms interface{}

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
    portData.EntityData.SegmentPath = "port-data" + types.AddNoKeyToken(portData)
    portData.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherenthealth/" + portData.EntityData.SegmentPath
    portData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portData.EntityData.Children = types.NewOrderedMap()
    portData.EntityData.Children.Append("ctp-info", types.YChild{"CtpInfo", &portData.CtpInfo})
    portData.EntityData.Children.Append("interface-info", types.YChild{"InterfaceInfo", &portData.InterfaceInfo})
    portData.EntityData.Leafs = types.NewOrderedMap()
    portData.EntityData.Leafs.Append("rsip", types.YLeaf{"Rsip", portData.Rsip})
    portData.EntityData.Leafs.Append("optics-ctrl-created", types.YLeaf{"OpticsCtrlCreated", portData.OpticsCtrlCreated})
    portData.EntityData.Leafs.Append("dsp-ctrl-created", types.YLeaf{"DspCtrlCreated", portData.DspCtrlCreated})
    portData.EntityData.Leafs.Append("has-pluggable", types.YLeaf{"HasPluggable", portData.HasPluggable})
    portData.EntityData.Leafs.Append("optics-admin-up", types.YLeaf{"OpticsAdminUp", portData.OpticsAdminUp})
    portData.EntityData.Leafs.Append("dsp-admin-up", types.YLeaf{"DspAdminUp", portData.DspAdminUp})
    portData.EntityData.Leafs.Append("end-of-config", types.YLeaf{"EndOfConfig", portData.EndOfConfig})
    portData.EntityData.Leafs.Append("optics-laser-state", types.YLeaf{"OpticsLaserState", portData.OpticsLaserState})
    portData.EntityData.Leafs.Append("laser-state", types.YLeaf{"LaserState", portData.LaserState})
    portData.EntityData.Leafs.Append("laser-on-pending", types.YLeaf{"LaserOnPending", portData.LaserOnPending})
    portData.EntityData.Leafs.Append("provisioning-needed", types.YLeaf{"ProvisioningNeeded", portData.ProvisioningNeeded})
    portData.EntityData.Leafs.Append("force-reprovision", types.YLeaf{"ForceReprovision", portData.ForceReprovision})
    portData.EntityData.Leafs.Append("fp-port-idx", types.YLeaf{"FpPortIdx", portData.FpPortIdx})
    portData.EntityData.Leafs.Append("configured-frequency", types.YLeaf{"ConfiguredFrequency", portData.ConfiguredFrequency})
    portData.EntityData.Leafs.Append("provisioned-frequency", types.YLeaf{"ProvisionedFrequency", portData.ProvisionedFrequency})
    portData.EntityData.Leafs.Append("configured-tx-power", types.YLeaf{"ConfiguredTxPower", portData.ConfiguredTxPower})
    portData.EntityData.Leafs.Append("provisioned-tx-power", types.YLeaf{"ProvisionedTxPower", portData.ProvisionedTxPower})
    portData.EntityData.Leafs.Append("configured-cd-min", types.YLeaf{"ConfiguredCdMin", portData.ConfiguredCdMin})
    portData.EntityData.Leafs.Append("provisioned-cd-min", types.YLeaf{"ProvisionedCdMin", portData.ProvisionedCdMin})
    portData.EntityData.Leafs.Append("configured-cd-max", types.YLeaf{"ConfiguredCdMax", portData.ConfiguredCdMax})
    portData.EntityData.Leafs.Append("provisioned-cd-max", types.YLeaf{"ProvisionedCdMax", portData.ProvisionedCdMax})
    portData.EntityData.Leafs.Append("configured-traffic-type", types.YLeaf{"ConfiguredTrafficType", portData.ConfiguredTrafficType})
    portData.EntityData.Leafs.Append("provisioned-traffic-type", types.YLeaf{"ProvisionedTrafficType", portData.ProvisionedTrafficType})
    portData.EntityData.Leafs.Append("configured-loopback-mode", types.YLeaf{"ConfiguredLoopbackMode", portData.ConfiguredLoopbackMode})
    portData.EntityData.Leafs.Append("provisioned-loopback-mode", types.YLeaf{"ProvisionedLoopbackMode", portData.ProvisionedLoopbackMode})
    portData.EntityData.Leafs.Append("expected-ctp2-led-state", types.YLeaf{"ExpectedCtp2LedState", portData.ExpectedCtp2LedState})
    portData.EntityData.Leafs.Append("provisioned-ctp2-led-state", types.YLeaf{"ProvisionedCtp2LedState", portData.ProvisionedCtp2LedState})
    portData.EntityData.Leafs.Append("led-op-rc", types.YLeaf{"LedOpRc", portData.LedOpRc})
    portData.EntityData.Leafs.Append("laser-op-rc", types.YLeaf{"LaserOpRc", portData.LaserOpRc})
    portData.EntityData.Leafs.Append("wlen-op-rc", types.YLeaf{"WlenOpRc", portData.WlenOpRc})
    portData.EntityData.Leafs.Append("traffic-op-rc", types.YLeaf{"TrafficOpRc", portData.TrafficOpRc})
    portData.EntityData.Leafs.Append("loopback-op-rc", types.YLeaf{"LoopbackOpRc", portData.LoopbackOpRc})
    portData.EntityData.Leafs.Append("tx-power-op-rc", types.YLeaf{"TxPowerOpRc", portData.TxPowerOpRc})
    portData.EntityData.Leafs.Append("cd-min-op-rc", types.YLeaf{"CdMinOpRc", portData.CdMinOpRc})
    portData.EntityData.Leafs.Append("cd-max-op-rc", types.YLeaf{"CdMaxOpRc", portData.CdMaxOpRc})
    portData.EntityData.Leafs.Append("provisioning-failed", types.YLeaf{"ProvisioningFailed", portData.ProvisioningFailed})
    portData.EntityData.Leafs.Append("ctp2-hw-alarms", types.YLeaf{"Ctp2HwAlarms", portData.Ctp2HwAlarms})
    portData.EntityData.Leafs.Append("dsp-hw-alarms", types.YLeaf{"DspHwAlarms", portData.DspHwAlarms})
    portData.EntityData.Leafs.Append("is-pm-port-created-opt", types.YLeaf{"IsPmPortCreatedOpt", portData.IsPmPortCreatedOpt})
    portData.EntityData.Leafs.Append("rc-pm-port-opt", types.YLeaf{"RcPmPortOpt", portData.RcPmPortOpt})
    portData.EntityData.Leafs.Append("pm-port-state-opt", types.YLeaf{"PmPortStateOpt", portData.PmPortStateOpt})
    portData.EntityData.Leafs.Append("rc-pm-provision-opt", types.YLeaf{"RcPmProvisionOpt", portData.RcPmProvisionOpt})
    portData.EntityData.Leafs.Append("is-alarm-port-created-opt", types.YLeaf{"IsAlarmPortCreatedOpt", portData.IsAlarmPortCreatedOpt})
    portData.EntityData.Leafs.Append("rc-alarm-port-opt", types.YLeaf{"RcAlarmPortOpt", portData.RcAlarmPortOpt})
    portData.EntityData.Leafs.Append("is-pm-port-created-dsp", types.YLeaf{"IsPmPortCreatedDsp", portData.IsPmPortCreatedDsp})
    portData.EntityData.Leafs.Append("rc-pm-port-dsp", types.YLeaf{"RcPmPortDsp", portData.RcPmPortDsp})
    portData.EntityData.Leafs.Append("pm-port-state-dsp", types.YLeaf{"PmPortStateDsp", portData.PmPortStateDsp})
    portData.EntityData.Leafs.Append("rc-pm-provision-dsp", types.YLeaf{"RcPmProvisionDsp", portData.RcPmProvisionDsp})
    portData.EntityData.Leafs.Append("is-alarm-port-created-dsp", types.YLeaf{"IsAlarmPortCreatedDsp", portData.IsAlarmPortCreatedDsp})
    portData.EntityData.Leafs.Append("rc-alarm-port-dsp", types.YLeaf{"RcAlarmPortDsp", portData.RcAlarmPortDsp})

    portData.EntityData.YListKeys = []string {}

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

    // module hardware version number. The type is string with length: 0..16.
    ModuleHardwareVersionNumber interface{}

    // module firmware running version number. The type is string with length:
    // 0..16.
    ModuleFirmwareRunningVersionNumber interface{}

    // module firmware committed version number. The type is string with length:
    // 0..16.
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
    ctpInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherenthealth/port-data/" + ctpInfo.EntityData.SegmentPath
    ctpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ctpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ctpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ctpInfo.EntityData.Children = types.NewOrderedMap()
    ctpInfo.EntityData.Leafs = types.NewOrderedMap()
    ctpInfo.EntityData.Leafs.Append("deviation", types.YLeaf{"Deviation", ctpInfo.Deviation})
    ctpInfo.EntityData.Leafs.Append("part-number", types.YLeaf{"PartNumber", ctpInfo.PartNumber})
    ctpInfo.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", ctpInfo.SerialNumber})
    ctpInfo.EntityData.Leafs.Append("date-code-number", types.YLeaf{"DateCodeNumber", ctpInfo.DateCodeNumber})
    ctpInfo.EntityData.Leafs.Append("clei-code-number", types.YLeaf{"CleiCodeNumber", ctpInfo.CleiCodeNumber})
    ctpInfo.EntityData.Leafs.Append("vendorname", types.YLeaf{"Vendorname", ctpInfo.Vendorname})
    ctpInfo.EntityData.Leafs.Append("description", types.YLeaf{"Description", ctpInfo.Description})
    ctpInfo.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", ctpInfo.Pid})
    ctpInfo.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", ctpInfo.Vid})
    ctpInfo.EntityData.Leafs.Append("module-hardware-version-number", types.YLeaf{"ModuleHardwareVersionNumber", ctpInfo.ModuleHardwareVersionNumber})
    ctpInfo.EntityData.Leafs.Append("module-firmware-running-version-number", types.YLeaf{"ModuleFirmwareRunningVersionNumber", ctpInfo.ModuleFirmwareRunningVersionNumber})
    ctpInfo.EntityData.Leafs.Append("module-firmware-committed-version-number", types.YLeaf{"ModuleFirmwareCommittedVersionNumber", ctpInfo.ModuleFirmwareCommittedVersionNumber})
    ctpInfo.EntityData.Leafs.Append("ctp-type", types.YLeaf{"CtpType", ctpInfo.CtpType})

    ctpInfo.EntityData.YListKeys = []string {}

    return &(ctpInfo.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo
// interface info
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // eth data. The type is slice of
    // Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData.
    EthData []*Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
}

func (interfaceInfo *Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo) GetEntityData() *types.CommonEntityData {
    interfaceInfo.EntityData.YFilter = interfaceInfo.YFilter
    interfaceInfo.EntityData.YangName = "interface-info"
    interfaceInfo.EntityData.BundleName = "cisco_ios_xr"
    interfaceInfo.EntityData.ParentYangName = "port-data"
    interfaceInfo.EntityData.SegmentPath = "interface-info"
    interfaceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherenthealth/port-data/" + interfaceInfo.EntityData.SegmentPath
    interfaceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceInfo.EntityData.Children = types.NewOrderedMap()
    interfaceInfo.EntityData.Children.Append("eth-data", types.YChild{"EthData", nil})
    for i := range interfaceInfo.EthData {
        types.SetYListKey(interfaceInfo.EthData[i], i)
        interfaceInfo.EntityData.Children.Append(types.GetSegmentPath(interfaceInfo.EthData[i]), types.YChild{"EthData", interfaceInfo.EthData[i]})
    }
    interfaceInfo.EntityData.Leafs = types.NewOrderedMap()

    interfaceInfo.EntityData.YListKeys = []string {}

    return &(interfaceInfo.EntityData)
}

// Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData
// eth data
type Coherent_Nodes_Node_Coherenthealth_PortData_InterfaceInfo_EthData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    ethData.EntityData.SegmentPath = "eth-data" + types.AddNoKeyToken(ethData)
    ethData.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs5500-coherent-node-oper:coherent/nodes/node/coherenthealth/port-data/interface-info/" + ethData.EntityData.SegmentPath
    ethData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethData.EntityData.Children = types.NewOrderedMap()
    ethData.EntityData.Leafs = types.NewOrderedMap()
    ethData.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", ethData.Ifname})
    ethData.EntityData.Leafs.Append("intf-handle", types.YLeaf{"IntfHandle", ethData.IntfHandle})
    ethData.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", ethData.AdminState})
    ethData.EntityData.Leafs.Append("admin-up", types.YLeaf{"AdminUp", ethData.AdminUp})
    ethData.EntityData.Leafs.Append("is-created", types.YLeaf{"IsCreated", ethData.IsCreated})

    ethData.EntityData.YListKeys = []string {}

    return &(ethData.EntityData)
}

