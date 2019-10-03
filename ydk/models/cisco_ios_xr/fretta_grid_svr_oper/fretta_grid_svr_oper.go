// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-grid-svr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   grid: GRID operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package fretta_grid_svr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fretta_grid_svr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-fretta-grid-svr-oper grid}", reflect.TypeOf(Grid{}))
    ydk.RegisterEntity("Cisco-IOS-XR-fretta-grid-svr-oper:grid", reflect.TypeOf(Grid{}))
}

// Grid
// GRID operational data
type Grid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Grid_Nodes
}

func (grid *Grid) GetEntityData() *types.CommonEntityData {
    grid.EntityData.YFilter = grid.YFilter
    grid.EntityData.YangName = "grid"
    grid.EntityData.BundleName = "cisco_ios_xr"
    grid.EntityData.ParentYangName = "Cisco-IOS-XR-fretta-grid-svr-oper"
    grid.EntityData.SegmentPath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid"
    grid.EntityData.AbsolutePath = grid.EntityData.SegmentPath
    grid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grid.EntityData.Children = types.NewOrderedMap()
    grid.EntityData.Children.Append("nodes", types.YChild{"Nodes", &grid.Nodes})
    grid.EntityData.Leafs = types.NewOrderedMap()

    grid.EntityData.YListKeys = []string {}

    return &(grid.EntityData)
}

// Grid_Nodes
// Table of nodes
type Grid_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular node. The type is slice of
    // Grid_Nodes_Node.
    Node []*Grid_Nodes_Node
}

func (nodes *Grid_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "grid"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/" + nodes.EntityData.SegmentPath
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

// Grid_Nodes_Node
// Operational data for a particular node
type Grid_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // GRID Client Table.
    ClientXr Grid_Nodes_Node_ClientXr

    // GRID Client Consistency Check.
    Clients Grid_Nodes_Node_Clients
}

func (node *Grid_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("client-xr", types.YChild{"ClientXr", &node.ClientXr})
    node.EntityData.Children.Append("clients", types.YChild{"Clients", &node.Clients})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Grid_Nodes_Node_ClientXr
// GRID Client Table
type Grid_Nodes_Node_ClientXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GRID Client Database. The type is slice of Grid_Nodes_Node_ClientXr_Client.
    Client []*Grid_Nodes_Node_ClientXr_Client
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetEntityData() *types.CommonEntityData {
    clientXr.EntityData.YFilter = clientXr.YFilter
    clientXr.EntityData.YangName = "client-xr"
    clientXr.EntityData.BundleName = "cisco_ios_xr"
    clientXr.EntityData.ParentYangName = "node"
    clientXr.EntityData.SegmentPath = "client-xr"
    clientXr.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/" + clientXr.EntityData.SegmentPath
    clientXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientXr.EntityData.Children = types.NewOrderedMap()
    clientXr.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clientXr.Client {
        clientXr.EntityData.Children.Append(types.GetSegmentPath(clientXr.Client[i]), types.YChild{"Client", clientXr.Client[i]})
    }
    clientXr.EntityData.Leafs = types.NewOrderedMap()

    clientXr.EntityData.YListKeys = []string {}

    return &(clientXr.EntityData)
}

// Grid_Nodes_Node_ClientXr_Client
// GRID Client Database
type Grid_Nodes_Node_ClientXr_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Client name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_ClientXr_Client_ClientData.
    ClientData []*Grid_Nodes_Node_ClientXr_Client_ClientData
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "client-xr"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientName, "client-name")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/client-xr/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("client-data", types.YChild{"ClientData", nil})
    for i := range client.ClientData {
        types.SetYListKey(client.ClientData[i], i)
        client.EntityData.Children.Append(types.GetSegmentPath(client.ClientData[i]), types.YChild{"ClientData", client.ClientData[i]})
    }
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", client.ClientName})

    client.EntityData.YListKeys = []string {"ClientName"}

    return &(client.EntityData)
}

// Grid_Nodes_Node_ClientXr_Client_ClientData
// Client information
type Grid_Nodes_Node_ClientXr_Client_ClientData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetEntityData() *types.CommonEntityData {
    clientData.EntityData.YFilter = clientData.YFilter
    clientData.EntityData.YangName = "client-data"
    clientData.EntityData.BundleName = "cisco_ios_xr"
    clientData.EntityData.ParentYangName = "client"
    clientData.EntityData.SegmentPath = "client-data" + types.AddNoKeyToken(clientData)
    clientData.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/client-xr/client/" + clientData.EntityData.SegmentPath
    clientData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientData.EntityData.Children = types.NewOrderedMap()
    clientData.EntityData.Leafs = types.NewOrderedMap()
    clientData.EntityData.Leafs.Append("res-id", types.YLeaf{"ResId", clientData.ResId})

    clientData.EntityData.YListKeys = []string {}

    return &(clientData.EntityData)
}

// Grid_Nodes_Node_Clients
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GRID Client Consistency Check. The type is slice of
    // Grid_Nodes_Node_Clients_Client.
    Client []*Grid_Nodes_Node_Clients_Client
}

func (clients *Grid_Nodes_Node_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "node"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Grid_Nodes_Node_Clients_Client
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Client name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_Clients_Client_ClientData.
    ClientData []*Grid_Nodes_Node_Clients_Client_ClientData
}

func (client *Grid_Nodes_Node_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientName, "client-name")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("client-data", types.YChild{"ClientData", nil})
    for i := range client.ClientData {
        types.SetYListKey(client.ClientData[i], i)
        client.EntityData.Children.Append(types.GetSegmentPath(client.ClientData[i]), types.YChild{"ClientData", client.ClientData[i]})
    }
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", client.ClientName})

    client.EntityData.YListKeys = []string {"ClientName"}

    return &(client.EntityData)
}

// Grid_Nodes_Node_Clients_Client_ClientData
// Client information
type Grid_Nodes_Node_Clients_Client_ClientData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetEntityData() *types.CommonEntityData {
    clientData.EntityData.YFilter = clientData.YFilter
    clientData.EntityData.YangName = "client-data"
    clientData.EntityData.BundleName = "cisco_ios_xr"
    clientData.EntityData.ParentYangName = "client"
    clientData.EntityData.SegmentPath = "client-data" + types.AddNoKeyToken(clientData)
    clientData.EntityData.AbsolutePath = "Cisco-IOS-XR-fretta-grid-svr-oper:grid/nodes/node/clients/client/" + clientData.EntityData.SegmentPath
    clientData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientData.EntityData.Children = types.NewOrderedMap()
    clientData.EntityData.Leafs = types.NewOrderedMap()
    clientData.EntityData.Leafs.Append("res-id", types.YLeaf{"ResId", clientData.ResId})

    clientData.EntityData.YListKeys = []string {}

    return &(clientData.EntityData)
}

