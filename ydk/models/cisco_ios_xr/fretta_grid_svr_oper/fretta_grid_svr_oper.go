// This module contains a collection of YANG definitions
// for Cisco IOS-XR fretta-grid-svr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   grid: GRID operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    grid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    grid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    grid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    grid.EntityData.Children = make(map[string]types.YChild)
    grid.EntityData.Children["nodes"] = types.YChild{"Nodes", &grid.Nodes}
    grid.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(grid.EntityData)
}

// Grid_Nodes
// Table of nodes
type Grid_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular node. The type is slice of
    // Grid_Nodes_Node.
    Node []Grid_Nodes_Node
}

func (nodes *Grid_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "grid"
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

// Grid_Nodes_Node
// Operational data for a particular node
type Grid_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["client-xr"] = types.YChild{"ClientXr", &node.ClientXr}
    node.EntityData.Children["clients"] = types.YChild{"Clients", &node.Clients}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Grid_Nodes_Node_ClientXr
// GRID Client Table
type Grid_Nodes_Node_ClientXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GRID Client Database. The type is slice of Grid_Nodes_Node_ClientXr_Client.
    Client []Grid_Nodes_Node_ClientXr_Client
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetEntityData() *types.CommonEntityData {
    clientXr.EntityData.YFilter = clientXr.YFilter
    clientXr.EntityData.YangName = "client-xr"
    clientXr.EntityData.BundleName = "cisco_ios_xr"
    clientXr.EntityData.ParentYangName = "node"
    clientXr.EntityData.SegmentPath = "client-xr"
    clientXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientXr.EntityData.Children = make(map[string]types.YChild)
    clientXr.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clientXr.Client {
        clientXr.EntityData.Children[types.GetSegmentPath(&clientXr.Client[i])] = types.YChild{"Client", &clientXr.Client[i]}
    }
    clientXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clientXr.EntityData)
}

// Grid_Nodes_Node_ClientXr_Client
// GRID Client Database
type Grid_Nodes_Node_ClientXr_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_ClientXr_Client_ClientData.
    ClientData []Grid_Nodes_Node_ClientXr_Client_ClientData
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "client-xr"
    client.EntityData.SegmentPath = "client" + "[client-name='" + fmt.Sprintf("%v", client.ClientName) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Children["client-data"] = types.YChild{"ClientData", nil}
    for i := range client.ClientData {
        client.EntityData.Children[types.GetSegmentPath(&client.ClientData[i])] = types.YChild{"ClientData", &client.ClientData[i]}
    }
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-name"] = types.YLeaf{"ClientName", client.ClientName}
    return &(client.EntityData)
}

// Grid_Nodes_Node_ClientXr_Client_ClientData
// Client information
type Grid_Nodes_Node_ClientXr_Client_ClientData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetEntityData() *types.CommonEntityData {
    clientData.EntityData.YFilter = clientData.YFilter
    clientData.EntityData.YangName = "client-data"
    clientData.EntityData.BundleName = "cisco_ios_xr"
    clientData.EntityData.ParentYangName = "client"
    clientData.EntityData.SegmentPath = "client-data"
    clientData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientData.EntityData.Children = make(map[string]types.YChild)
    clientData.EntityData.Leafs = make(map[string]types.YLeaf)
    clientData.EntityData.Leafs["res-id"] = types.YLeaf{"ResId", clientData.ResId}
    return &(clientData.EntityData)
}

// Grid_Nodes_Node_Clients
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GRID Client Consistency Check. The type is slice of
    // Grid_Nodes_Node_Clients_Client.
    Client []Grid_Nodes_Node_Clients_Client
}

func (clients *Grid_Nodes_Node_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "node"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clients.EntityData)
}

// Grid_Nodes_Node_Clients_Client
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_Clients_Client_ClientData.
    ClientData []Grid_Nodes_Node_Clients_Client_ClientData
}

func (client *Grid_Nodes_Node_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[client-name='" + fmt.Sprintf("%v", client.ClientName) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Children["client-data"] = types.YChild{"ClientData", nil}
    for i := range client.ClientData {
        client.EntityData.Children[types.GetSegmentPath(&client.ClientData[i])] = types.YChild{"ClientData", &client.ClientData[i]}
    }
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["client-name"] = types.YLeaf{"ClientName", client.ClientName}
    return &(client.EntityData)
}

// Grid_Nodes_Node_Clients_Client_ClientData
// Client information
type Grid_Nodes_Node_Clients_Client_ClientData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetEntityData() *types.CommonEntityData {
    clientData.EntityData.YFilter = clientData.YFilter
    clientData.EntityData.YangName = "client-data"
    clientData.EntityData.BundleName = "cisco_ios_xr"
    clientData.EntityData.ParentYangName = "client"
    clientData.EntityData.SegmentPath = "client-data"
    clientData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientData.EntityData.Children = make(map[string]types.YChild)
    clientData.EntityData.Leafs = make(map[string]types.YLeaf)
    clientData.EntityData.Leafs["res-id"] = types.YLeaf{"ResId", clientData.ResId}
    return &(clientData.EntityData)
}

