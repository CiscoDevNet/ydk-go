// This module contains a collection of YANG definitions
// for Cisco IOS-XR sysdb package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sysdb-connections: Sysdb health on client connections
//   sysdb: sysdb
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysdb_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysdb_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysdb-oper sysdb-connections}", reflect.TypeOf(SysdbConnections{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysdb-oper:sysdb-connections", reflect.TypeOf(SysdbConnections{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-sysdb-oper sysdb}", reflect.TypeOf(Sysdb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysdb-oper:sysdb", reflect.TypeOf(Sysdb{}))
}

// SysdbConnections
// Sysdb health on client connections
type SysdbConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node operational data.
    Nodes SysdbConnections_Nodes
}

func (sysdbConnections *SysdbConnections) GetEntityData() *types.CommonEntityData {
    sysdbConnections.EntityData.YFilter = sysdbConnections.YFilter
    sysdbConnections.EntityData.YangName = "sysdb-connections"
    sysdbConnections.EntityData.BundleName = "cisco_ios_xr"
    sysdbConnections.EntityData.ParentYangName = "Cisco-IOS-XR-sysdb-oper"
    sysdbConnections.EntityData.SegmentPath = "Cisco-IOS-XR-sysdb-oper:sysdb-connections"
    sysdbConnections.EntityData.AbsolutePath = sysdbConnections.EntityData.SegmentPath
    sysdbConnections.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysdbConnections.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysdbConnections.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysdbConnections.EntityData.Children = types.NewOrderedMap()
    sysdbConnections.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sysdbConnections.Nodes})
    sysdbConnections.EntityData.Leafs = types.NewOrderedMap()

    sysdbConnections.EntityData.YListKeys = []string {}

    return &(sysdbConnections.EntityData)
}

// SysdbConnections_Nodes
// Node operational data
type SysdbConnections_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-node Sysdb health on connection. The type is slice of
    // SysdbConnections_Nodes_Node.
    Node []*SysdbConnections_Nodes_Node
}

func (nodes *SysdbConnections_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "sysdb-connections"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysdb-oper:sysdb-connections/" + nodes.EntityData.SegmentPath
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

// SysdbConnections_Nodes_Node
// Per-node Sysdb health on connection
type SysdbConnections_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Per-node Sysdb Client Connections. The type is string.
    Connections interface{}
}

func (node *SysdbConnections_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-sysdb-oper:sysdb-connections/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("connections", types.YLeaf{"Connections", node.Connections})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Sysdb
// sysdb
type Sysdb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sysdb health for configuration space. The type is string.
    ConfigurationSpace interface{}

    // Sysdb health on memory consumption. The type is string.
    Memory interface{}

    // Sysdb health for operational space. The type is string.
    IpcSpace interface{}

    // Sysdb health on cpu consumption. The type is string.
    Cpu interface{}
}

func (sysdb *Sysdb) GetEntityData() *types.CommonEntityData {
    sysdb.EntityData.YFilter = sysdb.YFilter
    sysdb.EntityData.YangName = "sysdb"
    sysdb.EntityData.BundleName = "cisco_ios_xr"
    sysdb.EntityData.ParentYangName = "Cisco-IOS-XR-sysdb-oper"
    sysdb.EntityData.SegmentPath = "Cisco-IOS-XR-sysdb-oper:sysdb"
    sysdb.EntityData.AbsolutePath = sysdb.EntityData.SegmentPath
    sysdb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysdb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysdb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysdb.EntityData.Children = types.NewOrderedMap()
    sysdb.EntityData.Leafs = types.NewOrderedMap()
    sysdb.EntityData.Leafs.Append("configuration-space", types.YLeaf{"ConfigurationSpace", sysdb.ConfigurationSpace})
    sysdb.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", sysdb.Memory})
    sysdb.EntityData.Leafs.Append("ipc-space", types.YLeaf{"IpcSpace", sysdb.IpcSpace})
    sysdb.EntityData.Leafs.Append("cpu", types.YLeaf{"Cpu", sysdb.Cpu})

    sysdb.EntityData.YListKeys = []string {}

    return &(sysdb.EntityData)
}

