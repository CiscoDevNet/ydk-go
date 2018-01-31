// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-protocol-radius package operational data.
// 
// This module contains definitions
// for the following management objects:
//   radius: RADIUS operational data
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-locald-oper
// module with state data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_protocol_radius_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_protocol_radius_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-protocol-radius-oper radius}", reflect.TypeOf(Radius{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-protocol-radius-oper:radius", reflect.TypeOf(Radius{}))
}

// Radius
// RADIUS operational data
type Radius struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains all the nodes.
    Nodes Radius_Nodes
}

func (radius *Radius) GetFilter() yfilter.YFilter { return radius.YFilter }

func (radius *Radius) SetFilter(yf yfilter.YFilter) { radius.YFilter = yf }

func (radius *Radius) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (radius *Radius) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-protocol-radius-oper:radius"
}

func (radius *Radius) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &radius.Nodes
    }
    return nil
}

func (radius *Radius) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &radius.Nodes
    return children
}

func (radius *Radius) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radius *Radius) GetBundleName() string { return "cisco_ios_xr" }

func (radius *Radius) GetYangName() string { return "radius" }

func (radius *Radius) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radius *Radius) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radius *Radius) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radius *Radius) SetParent(parent types.Entity) { radius.parent = parent }

func (radius *Radius) GetParent() types.Entity { return radius.parent }

func (radius *Radius) GetParentYangName() string { return "Cisco-IOS-XR-aaa-protocol-radius-oper" }

// Radius_Nodes
// Contains all the nodes
type Radius_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS operational data for a particular node. The type is slice of
    // Radius_Nodes_Node.
    Node []Radius_Nodes_Node
}

func (nodes *Radius_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Radius_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Radius_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Radius_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Radius_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Radius_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Radius_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Radius_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Radius_Nodes) GetYangName() string { return "nodes" }

func (nodes *Radius_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Radius_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Radius_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Radius_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Radius_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Radius_Nodes) GetParentYangName() string { return "radius" }

// Radius_Nodes_Node
// RADIUS operational data for a particular node
type Radius_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // RADIUS client data.
    Client Radius_Nodes_Node_Client

    // RADIUS dead criteria information.
    DeadCriteria Radius_Nodes_Node_DeadCriteria

    // RADIUS authentication data.
    Authentication Radius_Nodes_Node_Authentication

    // RADIUS accounting data.
    Accounting Radius_Nodes_Node_Accounting

    // RADIUS server group table.
    ServerGroups Radius_Nodes_Node_ServerGroups

    // Dynamic authorization data.
    DynamicAuthorization Radius_Nodes_Node_DynamicAuthorization
}

func (node *Radius_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Radius_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Radius_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "client" { return "Client" }
    if yname == "dead-criteria" { return "DeadCriteria" }
    if yname == "authentication" { return "Authentication" }
    if yname == "accounting" { return "Accounting" }
    if yname == "server-groups" { return "ServerGroups" }
    if yname == "dynamic-authorization" { return "DynamicAuthorization" }
    return ""
}

func (node *Radius_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Radius_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        return &node.Client
    }
    if childYangName == "dead-criteria" {
        return &node.DeadCriteria
    }
    if childYangName == "authentication" {
        return &node.Authentication
    }
    if childYangName == "accounting" {
        return &node.Accounting
    }
    if childYangName == "server-groups" {
        return &node.ServerGroups
    }
    if childYangName == "dynamic-authorization" {
        return &node.DynamicAuthorization
    }
    return nil
}

func (node *Radius_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client"] = &node.Client
    children["dead-criteria"] = &node.DeadCriteria
    children["authentication"] = &node.Authentication
    children["accounting"] = &node.Accounting
    children["server-groups"] = &node.ServerGroups
    children["dynamic-authorization"] = &node.DynamicAuthorization
    return children
}

func (node *Radius_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Radius_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Radius_Nodes_Node) GetYangName() string { return "node" }

func (node *Radius_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Radius_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Radius_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Radius_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Radius_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Radius_Nodes_Node) GetParentYangName() string { return "nodes" }

// Radius_Nodes_Node_Client
// RADIUS client data
type Radius_Nodes_Node_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of RADIUS access responses packets received from unknown addresses.
    // The type is interface{} with range: 0..4294967295.
    UnknownAuthenticationResponses interface{}

    // NAS-Identifier of the RADIUS authentication client. The type is string.
    AuthenticationNasId interface{}

    // Number of RADIUS accounting responses packets received from unknown
    // addresses. The type is interface{} with range: 0..4294967295.
    UnknownAccountingResponses interface{}
}

func (client *Radius_Nodes_Node_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Radius_Nodes_Node_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Radius_Nodes_Node_Client) GetGoName(yname string) string {
    if yname == "unknown-authentication-responses" { return "UnknownAuthenticationResponses" }
    if yname == "authentication-nas-id" { return "AuthenticationNasId" }
    if yname == "unknown-accounting-responses" { return "UnknownAccountingResponses" }
    return ""
}

func (client *Radius_Nodes_Node_Client) GetSegmentPath() string {
    return "client"
}

func (client *Radius_Nodes_Node_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *Radius_Nodes_Node_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *Radius_Nodes_Node_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-authentication-responses"] = client.UnknownAuthenticationResponses
    leafs["authentication-nas-id"] = client.AuthenticationNasId
    leafs["unknown-accounting-responses"] = client.UnknownAccountingResponses
    return leafs
}

func (client *Radius_Nodes_Node_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Radius_Nodes_Node_Client) GetYangName() string { return "client" }

func (client *Radius_Nodes_Node_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Radius_Nodes_Node_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Radius_Nodes_Node_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Radius_Nodes_Node_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Radius_Nodes_Node_Client) GetParent() types.Entity { return client.parent }

func (client *Radius_Nodes_Node_Client) GetParentYangName() string { return "node" }

// Radius_Nodes_Node_DeadCriteria
// RADIUS dead criteria information
type Radius_Nodes_Node_DeadCriteria struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS server dead criteria host table.
    Hosts Radius_Nodes_Node_DeadCriteria_Hosts
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetFilter() yfilter.YFilter { return deadCriteria.YFilter }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) SetFilter(yf yfilter.YFilter) { deadCriteria.YFilter = yf }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetGoName(yname string) string {
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetSegmentPath() string {
    return "dead-criteria"
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hosts" {
        return &deadCriteria.Hosts
    }
    return nil
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hosts"] = &deadCriteria.Hosts
    return children
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetBundleName() string { return "cisco_ios_xr" }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetYangName() string { return "dead-criteria" }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) SetParent(parent types.Entity) { deadCriteria.parent = parent }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetParent() types.Entity { return deadCriteria.parent }

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetParentYangName() string { return "node" }

// Radius_Nodes_Node_DeadCriteria_Hosts
// RADIUS server dead criteria host table
type Radius_Nodes_Node_DeadCriteria_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS Server. The type is slice of
    // Radius_Nodes_Node_DeadCriteria_Hosts_Host.
    Host []Radius_Nodes_Node_DeadCriteria_Hosts_Host
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node_DeadCriteria_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetYangName() string { return "hosts" }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetParentYangName() string { return "dead-criteria" }

// Radius_Nodes_Node_DeadCriteria_Hosts_Host
// RADIUS Server
type Radius_Nodes_Node_DeadCriteria_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Authentication Port number (standard port 1645). The type is interface{}
    // with range: 1..65535.
    AuthPortNumber interface{}

    // Accounting Port number (standard port 1646). The type is interface{} with
    // range: 1..65535.
    AcctPortNumber interface{}

    // Time in seconds.
    Time Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time

    // Number of tries.
    Tries Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auth-port-number" { return "AuthPortNumber" }
    if yname == "acct-port-number" { return "AcctPortNumber" }
    if yname == "time" { return "Time" }
    if yname == "tries" { return "Tries" }
    return ""
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetSegmentPath() string {
    return "host"
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time" {
        return &host.Time
    }
    if childYangName == "tries" {
        return &host.Tries
    }
    return nil
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time"] = &host.Time
    children["tries"] = &host.Tries
    return children
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = host.IpAddress
    leafs["auth-port-number"] = host.AuthPortNumber
    leafs["acct-port-number"] = host.AcctPortNumber
    return leafs
}

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetYangName() string { return "host" }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetParentYangName() string { return "hosts" }

// Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time
// Time in seconds
type Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value for time or tries. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // True if computed; false if not. The type is bool.
    IsComputed interface{}
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetFilter() yfilter.YFilter { return time.YFilter }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) SetFilter(yf yfilter.YFilter) { time.YFilter = yf }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "is-computed" { return "IsComputed" }
    return ""
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetSegmentPath() string {
    return "time"
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = time.Value
    leafs["is-computed"] = time.IsComputed
    return leafs
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetBundleName() string { return "cisco_ios_xr" }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetYangName() string { return "time" }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) SetParent(parent types.Entity) { time.parent = parent }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetParent() types.Entity { return time.parent }

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetParentYangName() string { return "host" }

// Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries
// Number of tries
type Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value for time or tries. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // True if computed; false if not. The type is bool.
    IsComputed interface{}
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetFilter() yfilter.YFilter { return tries.YFilter }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) SetFilter(yf yfilter.YFilter) { tries.YFilter = yf }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "is-computed" { return "IsComputed" }
    return ""
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetSegmentPath() string {
    return "tries"
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = tries.Value
    leafs["is-computed"] = tries.IsComputed
    return leafs
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetBundleName() string { return "cisco_ios_xr" }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetYangName() string { return "tries" }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) SetParent(parent types.Entity) { tries.parent = parent }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetParent() types.Entity { return tries.parent }

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetParentYangName() string { return "host" }

// Radius_Nodes_Node_Authentication
// RADIUS authentication data
type Radius_Nodes_Node_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of authentication groups. The type is slice of
    // Radius_Nodes_Node_Authentication_AuthenticationGroup.
    AuthenticationGroup []Radius_Nodes_Node_Authentication_AuthenticationGroup
}

func (authentication *Radius_Nodes_Node_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Radius_Nodes_Node_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Radius_Nodes_Node_Authentication) GetGoName(yname string) string {
    if yname == "authentication-group" { return "AuthenticationGroup" }
    return ""
}

func (authentication *Radius_Nodes_Node_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Radius_Nodes_Node_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication-group" {
        for _, c := range authentication.AuthenticationGroup {
            if authentication.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node_Authentication_AuthenticationGroup{}
        authentication.AuthenticationGroup = append(authentication.AuthenticationGroup, child)
        return &authentication.AuthenticationGroup[len(authentication.AuthenticationGroup)-1]
    }
    return nil
}

func (authentication *Radius_Nodes_Node_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authentication.AuthenticationGroup {
        children[authentication.AuthenticationGroup[i].GetSegmentPath()] = &authentication.AuthenticationGroup[i]
    }
    return children
}

func (authentication *Radius_Nodes_Node_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentication *Radius_Nodes_Node_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Radius_Nodes_Node_Authentication) GetYangName() string { return "authentication" }

func (authentication *Radius_Nodes_Node_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Radius_Nodes_Node_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Radius_Nodes_Node_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Radius_Nodes_Node_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Radius_Nodes_Node_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Radius_Nodes_Node_Authentication) GetParentYangName() string { return "node" }

// Radius_Nodes_Node_Authentication_AuthenticationGroup
// List of authentication groups
type Radius_Nodes_Node_Authentication_AuthenticationGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Authentication port number. The type is interface{} with range:
    // 0..4294967295.
    Port interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Authentication data.
    Authentication Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetFilter() yfilter.YFilter { return authenticationGroup.YFilter }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) SetFilter(yf yfilter.YFilter) { authenticationGroup.YFilter = yf }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "port" { return "Port" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "family" { return "Family" }
    if yname == "authentication" { return "Authentication" }
    return ""
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetSegmentPath() string {
    return "authentication-group"
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authentication" {
        return &authenticationGroup.Authentication
    }
    return nil
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authentication"] = &authenticationGroup.Authentication
    return children
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = authenticationGroup.ServerAddress
    leafs["port"] = authenticationGroup.Port
    leafs["ip-address"] = authenticationGroup.IpAddress
    leafs["family"] = authenticationGroup.Family
    return leafs
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetBundleName() string { return "cisco_ios_xr" }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetYangName() string { return "authentication-group" }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) SetParent(parent types.Entity) { authenticationGroup.parent = parent }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetParent() types.Entity { return authenticationGroup.parent }

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetParentYangName() string { return "authentication" }

// Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication
// Authentication data
type Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of access requests. The type is interface{} with range:
    // 0..4294967295.
    AccessRequests interface{}

    // Number of pending access requests. The type is interface{} with range:
    // 0..4294967295.
    PendingAccessRequests interface{}

    // Number of retransmitted access requests. The type is interface{} with
    // range: 0..4294967295.
    AccessRequestRetransmits interface{}

    // Number of access accepts. The type is interface{} with range:
    // 0..4294967295.
    AccessAccepts interface{}

    // Number of access rejects. The type is interface{} with range:
    // 0..4294967295.
    AccessRejects interface{}

    // Number of access challenges. The type is interface{} with range:
    // 0..4294967295.
    AccessChallenges interface{}

    // Number of access packets timed out. The type is interface{} with range:
    // 0..4294967295.
    AccessTimeouts interface{}

    // Number of bad access responses. The type is interface{} with range:
    // 0..4294967295.
    BadAccessResponses interface{}

    // Number of bad access authenticators. The type is interface{} with range:
    // 0..4294967295.
    BadAccessAuthenticators interface{}

    // Number of packets received with unknown type from authentication server.
    // The type is interface{} with range: 0..4294967295.
    UnknownAccessTypes interface{}

    // Number of access responses dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedAccessResponses interface{}

    // Round trip time for authentication in milliseconds. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    Rtt interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AuthenResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthenTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthenTransactionFailure interface{}

    // Number of unexpected authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenUnexpectedResponses interface{}

    // Number of server error authentication responses. The type is interface{}
    // with range: 0..4294967295.
    AuthenServerErrorResponses interface{}

    // Number of incorrect authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenIncorrectResponses interface{}
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetGoName(yname string) string {
    if yname == "access-requests" { return "AccessRequests" }
    if yname == "pending-access-requests" { return "PendingAccessRequests" }
    if yname == "access-request-retransmits" { return "AccessRequestRetransmits" }
    if yname == "access-accepts" { return "AccessAccepts" }
    if yname == "access-rejects" { return "AccessRejects" }
    if yname == "access-challenges" { return "AccessChallenges" }
    if yname == "access-timeouts" { return "AccessTimeouts" }
    if yname == "bad-access-responses" { return "BadAccessResponses" }
    if yname == "bad-access-authenticators" { return "BadAccessAuthenticators" }
    if yname == "unknown-access-types" { return "UnknownAccessTypes" }
    if yname == "dropped-access-responses" { return "DroppedAccessResponses" }
    if yname == "rtt" { return "Rtt" }
    if yname == "authen-response-time" { return "AuthenResponseTime" }
    if yname == "authen-transaction-successess" { return "AuthenTransactionSuccessess" }
    if yname == "authen-transaction-failure" { return "AuthenTransactionFailure" }
    if yname == "authen-unexpected-responses" { return "AuthenUnexpectedResponses" }
    if yname == "authen-server-error-responses" { return "AuthenServerErrorResponses" }
    if yname == "authen-incorrect-responses" { return "AuthenIncorrectResponses" }
    return ""
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-requests"] = authentication.AccessRequests
    leafs["pending-access-requests"] = authentication.PendingAccessRequests
    leafs["access-request-retransmits"] = authentication.AccessRequestRetransmits
    leafs["access-accepts"] = authentication.AccessAccepts
    leafs["access-rejects"] = authentication.AccessRejects
    leafs["access-challenges"] = authentication.AccessChallenges
    leafs["access-timeouts"] = authentication.AccessTimeouts
    leafs["bad-access-responses"] = authentication.BadAccessResponses
    leafs["bad-access-authenticators"] = authentication.BadAccessAuthenticators
    leafs["unknown-access-types"] = authentication.UnknownAccessTypes
    leafs["dropped-access-responses"] = authentication.DroppedAccessResponses
    leafs["rtt"] = authentication.Rtt
    leafs["authen-response-time"] = authentication.AuthenResponseTime
    leafs["authen-transaction-successess"] = authentication.AuthenTransactionSuccessess
    leafs["authen-transaction-failure"] = authentication.AuthenTransactionFailure
    leafs["authen-unexpected-responses"] = authentication.AuthenUnexpectedResponses
    leafs["authen-server-error-responses"] = authentication.AuthenServerErrorResponses
    leafs["authen-incorrect-responses"] = authentication.AuthenIncorrectResponses
    return leafs
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetYangName() string { return "authentication" }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetParentYangName() string { return "authentication-group" }

// Radius_Nodes_Node_Accounting
// RADIUS accounting data
type Radius_Nodes_Node_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of accounting groups. The type is slice of
    // Radius_Nodes_Node_Accounting_AccountingGroup.
    AccountingGroup []Radius_Nodes_Node_Accounting_AccountingGroup
}

func (accounting *Radius_Nodes_Node_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Radius_Nodes_Node_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Radius_Nodes_Node_Accounting) GetGoName(yname string) string {
    if yname == "accounting-group" { return "AccountingGroup" }
    return ""
}

func (accounting *Radius_Nodes_Node_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Radius_Nodes_Node_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting-group" {
        for _, c := range accounting.AccountingGroup {
            if accounting.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node_Accounting_AccountingGroup{}
        accounting.AccountingGroup = append(accounting.AccountingGroup, child)
        return &accounting.AccountingGroup[len(accounting.AccountingGroup)-1]
    }
    return nil
}

func (accounting *Radius_Nodes_Node_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accounting.AccountingGroup {
        children[accounting.AccountingGroup[i].GetSegmentPath()] = &accounting.AccountingGroup[i]
    }
    return children
}

func (accounting *Radius_Nodes_Node_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accounting *Radius_Nodes_Node_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Radius_Nodes_Node_Accounting) GetYangName() string { return "accounting" }

func (accounting *Radius_Nodes_Node_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Radius_Nodes_Node_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Radius_Nodes_Node_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Radius_Nodes_Node_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Radius_Nodes_Node_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Radius_Nodes_Node_Accounting) GetParentYangName() string { return "node" }

// Radius_Nodes_Node_Accounting_AccountingGroup
// List of accounting groups
type Radius_Nodes_Node_Accounting_AccountingGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Accounting port number. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Accounting data.
    Accounting Radius_Nodes_Node_Accounting_AccountingGroup_Accounting
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetFilter() yfilter.YFilter { return accountingGroup.YFilter }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) SetFilter(yf yfilter.YFilter) { accountingGroup.YFilter = yf }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "port" { return "Port" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "family" { return "Family" }
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetSegmentPath() string {
    return "accounting-group"
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &accountingGroup.Accounting
    }
    return nil
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &accountingGroup.Accounting
    return children
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = accountingGroup.ServerAddress
    leafs["port"] = accountingGroup.Port
    leafs["ip-address"] = accountingGroup.IpAddress
    leafs["family"] = accountingGroup.Family
    return leafs
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetBundleName() string { return "cisco_ios_xr" }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetYangName() string { return "accounting-group" }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) SetParent(parent types.Entity) { accountingGroup.parent = parent }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetParent() types.Entity { return accountingGroup.parent }

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetParentYangName() string { return "accounting" }

// Radius_Nodes_Node_Accounting_AccountingGroup_Accounting
// Accounting data
type Radius_Nodes_Node_Accounting_AccountingGroup_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of accounting requests. The type is interface{} with range:
    // 0..4294967295.
    Requests interface{}

    // Number of pending accounting requests. The type is interface{} with range:
    // 0..4294967295.
    PendingRequests interface{}

    // Number of retransmitted accounting requests. The type is interface{} with
    // range: 0..4294967295.
    Retransmits interface{}

    // Number of accounting responses. The type is interface{} with range:
    // 0..4294967295.
    Responses interface{}

    // Number of accounting packets timed-out. The type is interface{} with range:
    // 0..4294967295.
    Timeouts interface{}

    // Number of bad accounting responses. The type is interface{} with range:
    // 0..4294967295.
    BadResponses interface{}

    // Number of bad accounting authenticators. The type is interface{} with
    // range: 0..4294967295.
    BadAuthenticators interface{}

    // Number of packets received with unknown type from accounting server. The
    // type is interface{} with range: 0..4294967295.
    UnknownPacketTypes interface{}

    // Number of accounting responses dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedResponses interface{}

    // Round trip time for accounting in milliseconds. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    Rtt interface{}

    // Number of unexpected accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctUnexpectedResponses interface{}

    // Number of server error accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctServerErrorResponses interface{}

    // Number of incorrect accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctIncorrectResponses interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AcctResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AcctTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctTransactionFailure interface{}
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetGoName(yname string) string {
    if yname == "requests" { return "Requests" }
    if yname == "pending-requests" { return "PendingRequests" }
    if yname == "retransmits" { return "Retransmits" }
    if yname == "responses" { return "Responses" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "bad-responses" { return "BadResponses" }
    if yname == "bad-authenticators" { return "BadAuthenticators" }
    if yname == "unknown-packet-types" { return "UnknownPacketTypes" }
    if yname == "dropped-responses" { return "DroppedResponses" }
    if yname == "rtt" { return "Rtt" }
    if yname == "acct-unexpected-responses" { return "AcctUnexpectedResponses" }
    if yname == "acct-server-error-responses" { return "AcctServerErrorResponses" }
    if yname == "acct-incorrect-responses" { return "AcctIncorrectResponses" }
    if yname == "acct-response-time" { return "AcctResponseTime" }
    if yname == "acct-transaction-successess" { return "AcctTransactionSuccessess" }
    if yname == "acct-transaction-failure" { return "AcctTransactionFailure" }
    return ""
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["requests"] = accounting.Requests
    leafs["pending-requests"] = accounting.PendingRequests
    leafs["retransmits"] = accounting.Retransmits
    leafs["responses"] = accounting.Responses
    leafs["timeouts"] = accounting.Timeouts
    leafs["bad-responses"] = accounting.BadResponses
    leafs["bad-authenticators"] = accounting.BadAuthenticators
    leafs["unknown-packet-types"] = accounting.UnknownPacketTypes
    leafs["dropped-responses"] = accounting.DroppedResponses
    leafs["rtt"] = accounting.Rtt
    leafs["acct-unexpected-responses"] = accounting.AcctUnexpectedResponses
    leafs["acct-server-error-responses"] = accounting.AcctServerErrorResponses
    leafs["acct-incorrect-responses"] = accounting.AcctIncorrectResponses
    leafs["acct-response-time"] = accounting.AcctResponseTime
    leafs["acct-transaction-successess"] = accounting.AcctTransactionSuccessess
    leafs["acct-transaction-failure"] = accounting.AcctTransactionFailure
    return leafs
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetYangName() string { return "accounting" }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetParentYangName() string { return "accounting-group" }

// Radius_Nodes_Node_ServerGroups
// RADIUS server group table
type Radius_Nodes_Node_ServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS server group data. The type is slice of
    // Radius_Nodes_Node_ServerGroups_ServerGroup.
    ServerGroup []Radius_Nodes_Node_ServerGroups_ServerGroup
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetFilter() yfilter.YFilter { return serverGroups.YFilter }

func (serverGroups *Radius_Nodes_Node_ServerGroups) SetFilter(yf yfilter.YFilter) { serverGroups.YFilter = yf }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetGoName(yname string) string {
    if yname == "server-group" { return "ServerGroup" }
    return ""
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetSegmentPath() string {
    return "server-groups"
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-group" {
        for _, c := range serverGroups.ServerGroup {
            if serverGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node_ServerGroups_ServerGroup{}
        serverGroups.ServerGroup = append(serverGroups.ServerGroup, child)
        return &serverGroups.ServerGroup[len(serverGroups.ServerGroup)-1]
    }
    return nil
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serverGroups.ServerGroup {
        children[serverGroups.ServerGroup[i].GetSegmentPath()] = &serverGroups.ServerGroup[i]
    }
    return children
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetYangName() string { return "server-groups" }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroups *Radius_Nodes_Node_ServerGroups) SetParent(parent types.Entity) { serverGroups.parent = parent }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetParent() types.Entity { return serverGroups.parent }

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetParentYangName() string { return "node" }

// Radius_Nodes_Node_ServerGroups_ServerGroup
// RADIUS server group data
type Radius_Nodes_Node_ServerGroups_ServerGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // Number of groups. The type is interface{} with range: 0..4294967295.
    Groups interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Dead time in minutes. The type is interface{} with range: 0..4294967295.
    // Units are minute.
    DeadTime interface{}

    // Number of servers. The type is interface{} with range: 0..4294967295.
    Servers interface{}

    // Server groups. The type is slice of
    // Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup.
    ServerGroup []Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetFilter() yfilter.YFilter { return serverGroup.YFilter }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) SetFilter(yf yfilter.YFilter) { serverGroup.YFilter = yf }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetGoName(yname string) string {
    if yname == "server-group-name" { return "ServerGroupName" }
    if yname == "groups" { return "Groups" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "dead-time" { return "DeadTime" }
    if yname == "servers" { return "Servers" }
    if yname == "server-group" { return "ServerGroup" }
    return ""
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetSegmentPath() string {
    return "server-group" + "[server-group-name='" + fmt.Sprintf("%v", serverGroup.ServerGroupName) + "']"
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-group" {
        for _, c := range serverGroup.ServerGroup {
            if serverGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup{}
        serverGroup.ServerGroup = append(serverGroup.ServerGroup, child)
        return &serverGroup.ServerGroup[len(serverGroup.ServerGroup)-1]
    }
    return nil
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serverGroup.ServerGroup {
        children[serverGroup.ServerGroup[i].GetSegmentPath()] = &serverGroup.ServerGroup[i]
    }
    return children
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-group-name"] = serverGroup.ServerGroupName
    leafs["groups"] = serverGroup.Groups
    leafs["vrf-name"] = serverGroup.VrfName
    leafs["dead-time"] = serverGroup.DeadTime
    leafs["servers"] = serverGroup.Servers
    return leafs
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetYangName() string { return "server-group" }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) SetParent(parent types.Entity) { serverGroup.parent = parent }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetParent() types.Entity { return serverGroup.parent }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetParentYangName() string { return "server-groups" }

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup
// Server groups
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Server address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Authentication port. The type is interface{} with range: 0..4294967295.
    AuthenticationPort interface{}

    // Accounting port. The type is interface{} with range: 0..4294967295.
    AccountingPort interface{}

    // True if private. The type is bool.
    IsPrivate interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Accounting data.
    Accounting Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting

    // Authentication data.
    Authentication Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication

    // Authorization data.
    Authorization Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetFilter() yfilter.YFilter { return serverGroup.YFilter }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) SetFilter(yf yfilter.YFilter) { serverGroup.YFilter = yf }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "authentication-port" { return "AuthenticationPort" }
    if yname == "accounting-port" { return "AccountingPort" }
    if yname == "is-private" { return "IsPrivate" }
    if yname == "ip-address" { return "IpAddress" }
    if yname == "family" { return "Family" }
    if yname == "accounting" { return "Accounting" }
    if yname == "authentication" { return "Authentication" }
    if yname == "authorization" { return "Authorization" }
    return ""
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetSegmentPath() string {
    return "server-group"
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &serverGroup.Accounting
    }
    if childYangName == "authentication" {
        return &serverGroup.Authentication
    }
    if childYangName == "authorization" {
        return &serverGroup.Authorization
    }
    return nil
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &serverGroup.Accounting
    children["authentication"] = &serverGroup.Authentication
    children["authorization"] = &serverGroup.Authorization
    return children
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = serverGroup.ServerAddress
    leafs["authentication-port"] = serverGroup.AuthenticationPort
    leafs["accounting-port"] = serverGroup.AccountingPort
    leafs["is-private"] = serverGroup.IsPrivate
    leafs["ip-address"] = serverGroup.IpAddress
    leafs["family"] = serverGroup.Family
    return leafs
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetYangName() string { return "server-group" }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) SetParent(parent types.Entity) { serverGroup.parent = parent }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetParent() types.Entity { return serverGroup.parent }

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetParentYangName() string { return "server-group" }

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting
// Accounting data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of accounting requests. The type is interface{} with range:
    // 0..4294967295.
    Requests interface{}

    // Number of pending accounting requests. The type is interface{} with range:
    // 0..4294967295.
    PendingRequests interface{}

    // Number of retransmitted accounting requests. The type is interface{} with
    // range: 0..4294967295.
    Retransmits interface{}

    // Number of accounting responses. The type is interface{} with range:
    // 0..4294967295.
    Responses interface{}

    // Number of accounting packets timed-out. The type is interface{} with range:
    // 0..4294967295.
    Timeouts interface{}

    // Number of bad accounting responses. The type is interface{} with range:
    // 0..4294967295.
    BadResponses interface{}

    // Number of bad accounting authenticators. The type is interface{} with
    // range: 0..4294967295.
    BadAuthenticators interface{}

    // Number of packets received with unknown type from accounting server. The
    // type is interface{} with range: 0..4294967295.
    UnknownPacketTypes interface{}

    // Number of accounting responses dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedResponses interface{}

    // Round trip time for accounting in milliseconds. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    Rtt interface{}

    // Number of unexpected accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctUnexpectedResponses interface{}

    // Number of server error accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctServerErrorResponses interface{}

    // Number of incorrect accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctIncorrectResponses interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AcctResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AcctTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctTransactionFailure interface{}
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetGoName(yname string) string {
    if yname == "requests" { return "Requests" }
    if yname == "pending-requests" { return "PendingRequests" }
    if yname == "retransmits" { return "Retransmits" }
    if yname == "responses" { return "Responses" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "bad-responses" { return "BadResponses" }
    if yname == "bad-authenticators" { return "BadAuthenticators" }
    if yname == "unknown-packet-types" { return "UnknownPacketTypes" }
    if yname == "dropped-responses" { return "DroppedResponses" }
    if yname == "rtt" { return "Rtt" }
    if yname == "acct-unexpected-responses" { return "AcctUnexpectedResponses" }
    if yname == "acct-server-error-responses" { return "AcctServerErrorResponses" }
    if yname == "acct-incorrect-responses" { return "AcctIncorrectResponses" }
    if yname == "acct-response-time" { return "AcctResponseTime" }
    if yname == "acct-transaction-successess" { return "AcctTransactionSuccessess" }
    if yname == "acct-transaction-failure" { return "AcctTransactionFailure" }
    return ""
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["requests"] = accounting.Requests
    leafs["pending-requests"] = accounting.PendingRequests
    leafs["retransmits"] = accounting.Retransmits
    leafs["responses"] = accounting.Responses
    leafs["timeouts"] = accounting.Timeouts
    leafs["bad-responses"] = accounting.BadResponses
    leafs["bad-authenticators"] = accounting.BadAuthenticators
    leafs["unknown-packet-types"] = accounting.UnknownPacketTypes
    leafs["dropped-responses"] = accounting.DroppedResponses
    leafs["rtt"] = accounting.Rtt
    leafs["acct-unexpected-responses"] = accounting.AcctUnexpectedResponses
    leafs["acct-server-error-responses"] = accounting.AcctServerErrorResponses
    leafs["acct-incorrect-responses"] = accounting.AcctIncorrectResponses
    leafs["acct-response-time"] = accounting.AcctResponseTime
    leafs["acct-transaction-successess"] = accounting.AcctTransactionSuccessess
    leafs["acct-transaction-failure"] = accounting.AcctTransactionFailure
    return leafs
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetYangName() string { return "accounting" }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetParentYangName() string { return "server-group" }

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication
// Authentication data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of access requests. The type is interface{} with range:
    // 0..4294967295.
    AccessRequests interface{}

    // Number of pending access requests. The type is interface{} with range:
    // 0..4294967295.
    PendingAccessRequests interface{}

    // Number of retransmitted access requests. The type is interface{} with
    // range: 0..4294967295.
    AccessRequestRetransmits interface{}

    // Number of access accepts. The type is interface{} with range:
    // 0..4294967295.
    AccessAccepts interface{}

    // Number of access rejects. The type is interface{} with range:
    // 0..4294967295.
    AccessRejects interface{}

    // Number of access challenges. The type is interface{} with range:
    // 0..4294967295.
    AccessChallenges interface{}

    // Number of access packets timed out. The type is interface{} with range:
    // 0..4294967295.
    AccessTimeouts interface{}

    // Number of bad access responses. The type is interface{} with range:
    // 0..4294967295.
    BadAccessResponses interface{}

    // Number of bad access authenticators. The type is interface{} with range:
    // 0..4294967295.
    BadAccessAuthenticators interface{}

    // Number of packets received with unknown type from authentication server.
    // The type is interface{} with range: 0..4294967295.
    UnknownAccessTypes interface{}

    // Number of access responses dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedAccessResponses interface{}

    // Round trip time for authentication in milliseconds. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    Rtt interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AuthenResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthenTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthenTransactionFailure interface{}

    // Number of unexpected authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenUnexpectedResponses interface{}

    // Number of server error authentication responses. The type is interface{}
    // with range: 0..4294967295.
    AuthenServerErrorResponses interface{}

    // Number of incorrect authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenIncorrectResponses interface{}
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetGoName(yname string) string {
    if yname == "access-requests" { return "AccessRequests" }
    if yname == "pending-access-requests" { return "PendingAccessRequests" }
    if yname == "access-request-retransmits" { return "AccessRequestRetransmits" }
    if yname == "access-accepts" { return "AccessAccepts" }
    if yname == "access-rejects" { return "AccessRejects" }
    if yname == "access-challenges" { return "AccessChallenges" }
    if yname == "access-timeouts" { return "AccessTimeouts" }
    if yname == "bad-access-responses" { return "BadAccessResponses" }
    if yname == "bad-access-authenticators" { return "BadAccessAuthenticators" }
    if yname == "unknown-access-types" { return "UnknownAccessTypes" }
    if yname == "dropped-access-responses" { return "DroppedAccessResponses" }
    if yname == "rtt" { return "Rtt" }
    if yname == "authen-response-time" { return "AuthenResponseTime" }
    if yname == "authen-transaction-successess" { return "AuthenTransactionSuccessess" }
    if yname == "authen-transaction-failure" { return "AuthenTransactionFailure" }
    if yname == "authen-unexpected-responses" { return "AuthenUnexpectedResponses" }
    if yname == "authen-server-error-responses" { return "AuthenServerErrorResponses" }
    if yname == "authen-incorrect-responses" { return "AuthenIncorrectResponses" }
    return ""
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-requests"] = authentication.AccessRequests
    leafs["pending-access-requests"] = authentication.PendingAccessRequests
    leafs["access-request-retransmits"] = authentication.AccessRequestRetransmits
    leafs["access-accepts"] = authentication.AccessAccepts
    leafs["access-rejects"] = authentication.AccessRejects
    leafs["access-challenges"] = authentication.AccessChallenges
    leafs["access-timeouts"] = authentication.AccessTimeouts
    leafs["bad-access-responses"] = authentication.BadAccessResponses
    leafs["bad-access-authenticators"] = authentication.BadAccessAuthenticators
    leafs["unknown-access-types"] = authentication.UnknownAccessTypes
    leafs["dropped-access-responses"] = authentication.DroppedAccessResponses
    leafs["rtt"] = authentication.Rtt
    leafs["authen-response-time"] = authentication.AuthenResponseTime
    leafs["authen-transaction-successess"] = authentication.AuthenTransactionSuccessess
    leafs["authen-transaction-failure"] = authentication.AuthenTransactionFailure
    leafs["authen-unexpected-responses"] = authentication.AuthenUnexpectedResponses
    leafs["authen-server-error-responses"] = authentication.AuthenServerErrorResponses
    leafs["authen-incorrect-responses"] = authentication.AuthenIncorrectResponses
    return leafs
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetYangName() string { return "authentication" }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetParentYangName() string { return "server-group" }

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization
// Authorization data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthorRequests interface{}

    // Number of access packets timed out. The type is interface{} with range:
    // 0..4294967295.
    AuthorRequestTimeouts interface{}

    // Number of unexpected authorization responses. The type is interface{} with
    // range: 0..4294967295.
    AuthorUnexpectedResponses interface{}

    // Number of server error authorization responses. The type is interface{}
    // with range: 0..4294967295.
    AuthorServerErrorResponses interface{}

    // Number of incorrect authorization responses. The type is interface{} with
    // range: 0..4294967295.
    AuthorIncorrectResponses interface{}

    // Average response time for authorization requests. The type is interface{}
    // with range: 0..4294967295.
    AuthorResponseTime interface{}

    // Number of succeeded authorization transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthorTransactionSuccessess interface{}

    // Number of failed authorization transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthorTransactionFailure interface{}
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetGoName(yname string) string {
    if yname == "author-requests" { return "AuthorRequests" }
    if yname == "author-request-timeouts" { return "AuthorRequestTimeouts" }
    if yname == "author-unexpected-responses" { return "AuthorUnexpectedResponses" }
    if yname == "author-server-error-responses" { return "AuthorServerErrorResponses" }
    if yname == "author-incorrect-responses" { return "AuthorIncorrectResponses" }
    if yname == "author-response-time" { return "AuthorResponseTime" }
    if yname == "author-transaction-successess" { return "AuthorTransactionSuccessess" }
    if yname == "author-transaction-failure" { return "AuthorTransactionFailure" }
    return ""
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetSegmentPath() string {
    return "authorization"
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["author-requests"] = authorization.AuthorRequests
    leafs["author-request-timeouts"] = authorization.AuthorRequestTimeouts
    leafs["author-unexpected-responses"] = authorization.AuthorUnexpectedResponses
    leafs["author-server-error-responses"] = authorization.AuthorServerErrorResponses
    leafs["author-incorrect-responses"] = authorization.AuthorIncorrectResponses
    leafs["author-response-time"] = authorization.AuthorResponseTime
    leafs["author-transaction-successess"] = authorization.AuthorTransactionSuccessess
    leafs["author-transaction-failure"] = authorization.AuthorTransactionFailure
    return leafs
}

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetYangName() string { return "authorization" }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetParentYangName() string { return "server-group" }

// Radius_Nodes_Node_DynamicAuthorization
// Dynamic authorization data
type Radius_Nodes_Node_DynamicAuthorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Invalid disconnected requests. The type is interface{} with range:
    // 0..4294967295.
    DisconnectedInvalidRequests interface{}

    // Invalid change of authorization requests. The type is interface{} with
    // range: 0..4294967295.
    InvalidCoaRequests interface{}
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetFilter() yfilter.YFilter { return dynamicAuthorization.YFilter }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) SetFilter(yf yfilter.YFilter) { dynamicAuthorization.YFilter = yf }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetGoName(yname string) string {
    if yname == "disconnected-invalid-requests" { return "DisconnectedInvalidRequests" }
    if yname == "invalid-coa-requests" { return "InvalidCoaRequests" }
    return ""
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetSegmentPath() string {
    return "dynamic-authorization"
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disconnected-invalid-requests"] = dynamicAuthorization.DisconnectedInvalidRequests
    leafs["invalid-coa-requests"] = dynamicAuthorization.InvalidCoaRequests
    return leafs
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetYangName() string { return "dynamic-authorization" }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) SetParent(parent types.Entity) { dynamicAuthorization.parent = parent }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetParent() types.Entity { return dynamicAuthorization.parent }

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetParentYangName() string { return "node" }

