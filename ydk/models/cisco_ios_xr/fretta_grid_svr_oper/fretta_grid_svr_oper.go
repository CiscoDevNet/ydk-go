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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Grid_Nodes
}

func (grid *Grid) GetFilter() yfilter.YFilter { return grid.YFilter }

func (grid *Grid) SetFilter(yf yfilter.YFilter) { grid.YFilter = yf }

func (grid *Grid) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (grid *Grid) GetSegmentPath() string {
    return "Cisco-IOS-XR-fretta-grid-svr-oper:grid"
}

func (grid *Grid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &grid.Nodes
    }
    return nil
}

func (grid *Grid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &grid.Nodes
    return children
}

func (grid *Grid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (grid *Grid) GetBundleName() string { return "cisco_ios_xr" }

func (grid *Grid) GetYangName() string { return "grid" }

func (grid *Grid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (grid *Grid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (grid *Grid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (grid *Grid) SetParent(parent types.Entity) { grid.parent = parent }

func (grid *Grid) GetParent() types.Entity { return grid.parent }

func (grid *Grid) GetParentYangName() string { return "Cisco-IOS-XR-fretta-grid-svr-oper" }

// Grid_Nodes
// Table of nodes
type Grid_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a particular node. The type is slice of
    // Grid_Nodes_Node.
    Node []Grid_Nodes_Node
}

func (nodes *Grid_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Grid_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Grid_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Grid_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Grid_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Grid_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Grid_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Grid_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Grid_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Grid_Nodes) GetYangName() string { return "nodes" }

func (nodes *Grid_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Grid_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Grid_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Grid_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Grid_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Grid_Nodes) GetParentYangName() string { return "grid" }

// Grid_Nodes_Node
// Operational data for a particular node
type Grid_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // GRID Client Table.
    ClientXr Grid_Nodes_Node_ClientXr

    // GRID Client Consistency Check.
    Clients Grid_Nodes_Node_Clients
}

func (node *Grid_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Grid_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Grid_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "client-xr" { return "ClientXr" }
    if yname == "clients" { return "Clients" }
    return ""
}

func (node *Grid_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Grid_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-xr" {
        return &node.ClientXr
    }
    if childYangName == "clients" {
        return &node.Clients
    }
    return nil
}

func (node *Grid_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client-xr"] = &node.ClientXr
    children["clients"] = &node.Clients
    return children
}

func (node *Grid_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Grid_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Grid_Nodes_Node) GetYangName() string { return "node" }

func (node *Grid_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Grid_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Grid_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Grid_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Grid_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Grid_Nodes_Node) GetParentYangName() string { return "nodes" }

// Grid_Nodes_Node_ClientXr
// GRID Client Table
type Grid_Nodes_Node_ClientXr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // GRID Client Database. The type is slice of Grid_Nodes_Node_ClientXr_Client.
    Client []Grid_Nodes_Node_ClientXr_Client
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetFilter() yfilter.YFilter { return clientXr.YFilter }

func (clientXr *Grid_Nodes_Node_ClientXr) SetFilter(yf yfilter.YFilter) { clientXr.YFilter = yf }

func (clientXr *Grid_Nodes_Node_ClientXr) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetSegmentPath() string {
    return "client-xr"
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clientXr.Client {
            if clientXr.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Grid_Nodes_Node_ClientXr_Client{}
        clientXr.Client = append(clientXr.Client, child)
        return &clientXr.Client[len(clientXr.Client)-1]
    }
    return nil
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientXr.Client {
        children[clientXr.Client[i].GetSegmentPath()] = &clientXr.Client[i]
    }
    return children
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clientXr *Grid_Nodes_Node_ClientXr) GetBundleName() string { return "cisco_ios_xr" }

func (clientXr *Grid_Nodes_Node_ClientXr) GetYangName() string { return "client-xr" }

func (clientXr *Grid_Nodes_Node_ClientXr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientXr *Grid_Nodes_Node_ClientXr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientXr *Grid_Nodes_Node_ClientXr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientXr *Grid_Nodes_Node_ClientXr) SetParent(parent types.Entity) { clientXr.parent = parent }

func (clientXr *Grid_Nodes_Node_ClientXr) GetParent() types.Entity { return clientXr.parent }

func (clientXr *Grid_Nodes_Node_ClientXr) GetParentYangName() string { return "node" }

// Grid_Nodes_Node_ClientXr_Client
// GRID Client Database
type Grid_Nodes_Node_ClientXr_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_ClientXr_Client_ClientData.
    ClientData []Grid_Nodes_Node_ClientXr_Client_ClientData
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Grid_Nodes_Node_ClientXr_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Grid_Nodes_Node_ClientXr_Client) GetGoName(yname string) string {
    if yname == "client-name" { return "ClientName" }
    if yname == "client-data" { return "ClientData" }
    return ""
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetSegmentPath() string {
    return "client" + "[client-name='" + fmt.Sprintf("%v", client.ClientName) + "']"
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-data" {
        for _, c := range client.ClientData {
            if client.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Grid_Nodes_Node_ClientXr_Client_ClientData{}
        client.ClientData = append(client.ClientData, child)
        return &client.ClientData[len(client.ClientData)-1]
    }
    return nil
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range client.ClientData {
        children[client.ClientData[i].GetSegmentPath()] = &client.ClientData[i]
    }
    return children
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-name"] = client.ClientName
    return leafs
}

func (client *Grid_Nodes_Node_ClientXr_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Grid_Nodes_Node_ClientXr_Client) GetYangName() string { return "client" }

func (client *Grid_Nodes_Node_ClientXr_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Grid_Nodes_Node_ClientXr_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Grid_Nodes_Node_ClientXr_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Grid_Nodes_Node_ClientXr_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Grid_Nodes_Node_ClientXr_Client) GetParent() types.Entity { return client.parent }

func (client *Grid_Nodes_Node_ClientXr_Client) GetParentYangName() string { return "client-xr" }

// Grid_Nodes_Node_ClientXr_Client_ClientData
// Client information
type Grid_Nodes_Node_ClientXr_Client_ClientData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetFilter() yfilter.YFilter { return clientData.YFilter }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) SetFilter(yf yfilter.YFilter) { clientData.YFilter = yf }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetGoName(yname string) string {
    if yname == "res-id" { return "ResId" }
    return ""
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetSegmentPath() string {
    return "client-data"
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["res-id"] = clientData.ResId
    return leafs
}

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetBundleName() string { return "cisco_ios_xr" }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetYangName() string { return "client-data" }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) SetParent(parent types.Entity) { clientData.parent = parent }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetParent() types.Entity { return clientData.parent }

func (clientData *Grid_Nodes_Node_ClientXr_Client_ClientData) GetParentYangName() string { return "client" }

// Grid_Nodes_Node_Clients
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // GRID Client Consistency Check. The type is slice of
    // Grid_Nodes_Node_Clients_Client.
    Client []Grid_Nodes_Node_Clients_Client
}

func (clients *Grid_Nodes_Node_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *Grid_Nodes_Node_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *Grid_Nodes_Node_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *Grid_Nodes_Node_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *Grid_Nodes_Node_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Grid_Nodes_Node_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *Grid_Nodes_Node_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *Grid_Nodes_Node_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *Grid_Nodes_Node_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *Grid_Nodes_Node_Clients) GetYangName() string { return "clients" }

func (clients *Grid_Nodes_Node_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *Grid_Nodes_Node_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *Grid_Nodes_Node_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *Grid_Nodes_Node_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *Grid_Nodes_Node_Clients) GetParent() types.Entity { return clients.parent }

func (clients *Grid_Nodes_Node_Clients) GetParentYangName() string { return "node" }

// Grid_Nodes_Node_Clients_Client
// GRID Client Consistency Check
type Grid_Nodes_Node_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Client name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // Client information. The type is slice of
    // Grid_Nodes_Node_Clients_Client_ClientData.
    ClientData []Grid_Nodes_Node_Clients_Client_ClientData
}

func (client *Grid_Nodes_Node_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Grid_Nodes_Node_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Grid_Nodes_Node_Clients_Client) GetGoName(yname string) string {
    if yname == "client-name" { return "ClientName" }
    if yname == "client-data" { return "ClientData" }
    return ""
}

func (client *Grid_Nodes_Node_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-name='" + fmt.Sprintf("%v", client.ClientName) + "']"
}

func (client *Grid_Nodes_Node_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-data" {
        for _, c := range client.ClientData {
            if client.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Grid_Nodes_Node_Clients_Client_ClientData{}
        client.ClientData = append(client.ClientData, child)
        return &client.ClientData[len(client.ClientData)-1]
    }
    return nil
}

func (client *Grid_Nodes_Node_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range client.ClientData {
        children[client.ClientData[i].GetSegmentPath()] = &client.ClientData[i]
    }
    return children
}

func (client *Grid_Nodes_Node_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-name"] = client.ClientName
    return leafs
}

func (client *Grid_Nodes_Node_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Grid_Nodes_Node_Clients_Client) GetYangName() string { return "client" }

func (client *Grid_Nodes_Node_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Grid_Nodes_Node_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Grid_Nodes_Node_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Grid_Nodes_Node_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Grid_Nodes_Node_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *Grid_Nodes_Node_Clients_Client) GetParentYangName() string { return "clients" }

// Grid_Nodes_Node_Clients_Client_ClientData
// Client information
type Grid_Nodes_Node_Clients_Client_ClientData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource ID. The type is interface{} with range: 0..4294967295.
    ResId interface{}
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetFilter() yfilter.YFilter { return clientData.YFilter }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) SetFilter(yf yfilter.YFilter) { clientData.YFilter = yf }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetGoName(yname string) string {
    if yname == "res-id" { return "ResId" }
    return ""
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetSegmentPath() string {
    return "client-data"
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["res-id"] = clientData.ResId
    return leafs
}

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetBundleName() string { return "cisco_ios_xr" }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetYangName() string { return "client-data" }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) SetParent(parent types.Entity) { clientData.parent = parent }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetParent() types.Entity { return clientData.parent }

func (clientData *Grid_Nodes_Node_Clients_Client_ClientData) GetParentYangName() string { return "client" }

