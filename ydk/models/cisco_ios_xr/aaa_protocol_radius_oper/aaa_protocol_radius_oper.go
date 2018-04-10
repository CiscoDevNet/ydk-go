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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains all the nodes.
    Nodes Radius_Nodes
}

func (radius *Radius) GetEntityData() *types.CommonEntityData {
    radius.EntityData.YFilter = radius.YFilter
    radius.EntityData.YangName = "radius"
    radius.EntityData.BundleName = "cisco_ios_xr"
    radius.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-protocol-radius-oper"
    radius.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius"
    radius.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radius.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radius.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radius.EntityData.Children = make(map[string]types.YChild)
    radius.EntityData.Children["nodes"] = types.YChild{"Nodes", &radius.Nodes}
    radius.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(radius.EntityData)
}

// Radius_Nodes
// Contains all the nodes
type Radius_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS operational data for a particular node. The type is slice of
    // Radius_Nodes_Node.
    Node []Radius_Nodes_Node
}

func (nodes *Radius_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "radius"
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

// Radius_Nodes_Node
// RADIUS operational data for a particular node
type Radius_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (node *Radius_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["client"] = types.YChild{"Client", &node.Client}
    node.EntityData.Children["dead-criteria"] = types.YChild{"DeadCriteria", &node.DeadCriteria}
    node.EntityData.Children["authentication"] = types.YChild{"Authentication", &node.Authentication}
    node.EntityData.Children["accounting"] = types.YChild{"Accounting", &node.Accounting}
    node.EntityData.Children["server-groups"] = types.YChild{"ServerGroups", &node.ServerGroups}
    node.EntityData.Children["dynamic-authorization"] = types.YChild{"DynamicAuthorization", &node.DynamicAuthorization}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Radius_Nodes_Node_Client
// RADIUS client data
type Radius_Nodes_Node_Client struct {
    EntityData types.CommonEntityData
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

func (client *Radius_Nodes_Node_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "node"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["unknown-authentication-responses"] = types.YLeaf{"UnknownAuthenticationResponses", client.UnknownAuthenticationResponses}
    client.EntityData.Leafs["authentication-nas-id"] = types.YLeaf{"AuthenticationNasId", client.AuthenticationNasId}
    client.EntityData.Leafs["unknown-accounting-responses"] = types.YLeaf{"UnknownAccountingResponses", client.UnknownAccountingResponses}
    return &(client.EntityData)
}

// Radius_Nodes_Node_DeadCriteria
// RADIUS dead criteria information
type Radius_Nodes_Node_DeadCriteria struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS server dead criteria host table.
    Hosts Radius_Nodes_Node_DeadCriteria_Hosts
}

func (deadCriteria *Radius_Nodes_Node_DeadCriteria) GetEntityData() *types.CommonEntityData {
    deadCriteria.EntityData.YFilter = deadCriteria.YFilter
    deadCriteria.EntityData.YangName = "dead-criteria"
    deadCriteria.EntityData.BundleName = "cisco_ios_xr"
    deadCriteria.EntityData.ParentYangName = "node"
    deadCriteria.EntityData.SegmentPath = "dead-criteria"
    deadCriteria.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deadCriteria.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deadCriteria.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deadCriteria.EntityData.Children = make(map[string]types.YChild)
    deadCriteria.EntityData.Children["hosts"] = types.YChild{"Hosts", &deadCriteria.Hosts}
    deadCriteria.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(deadCriteria.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts
// RADIUS server dead criteria host table
type Radius_Nodes_Node_DeadCriteria_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS Server. The type is slice of
    // Radius_Nodes_Node_DeadCriteria_Hosts_Host.
    Host []Radius_Nodes_Node_DeadCriteria_Hosts_Host
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "dead-criteria"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = make(map[string]types.YChild)
    hosts.EntityData.Children["host"] = types.YChild{"Host", nil}
    for i := range hosts.Host {
        hosts.EntityData.Children[types.GetSegmentPath(&hosts.Host[i])] = types.YChild{"Host", &hosts.Host[i]}
    }
    hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosts.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts_Host
// RADIUS Server
type Radius_Nodes_Node_DeadCriteria_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (host *Radius_Nodes_Node_DeadCriteria_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host"
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = make(map[string]types.YChild)
    host.EntityData.Children["time"] = types.YChild{"Time", &host.Time}
    host.EntityData.Children["tries"] = types.YChild{"Tries", &host.Tries}
    host.EntityData.Leafs = make(map[string]types.YLeaf)
    host.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", host.IpAddress}
    host.EntityData.Leafs["auth-port-number"] = types.YLeaf{"AuthPortNumber", host.AuthPortNumber}
    host.EntityData.Leafs["acct-port-number"] = types.YLeaf{"AcctPortNumber", host.AcctPortNumber}
    return &(host.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time
// Time in seconds
type Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value for time or tries. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // True if computed; false if not. The type is bool.
    IsComputed interface{}
}

func (time *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Time) GetEntityData() *types.CommonEntityData {
    time.EntityData.YFilter = time.YFilter
    time.EntityData.YangName = "time"
    time.EntityData.BundleName = "cisco_ios_xr"
    time.EntityData.ParentYangName = "host"
    time.EntityData.SegmentPath = "time"
    time.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    time.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    time.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    time.EntityData.Children = make(map[string]types.YChild)
    time.EntityData.Leafs = make(map[string]types.YLeaf)
    time.EntityData.Leafs["value"] = types.YLeaf{"Value", time.Value}
    time.EntityData.Leafs["is-computed"] = types.YLeaf{"IsComputed", time.IsComputed}
    return &(time.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries
// Number of tries
type Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value for time or tries. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // True if computed; false if not. The type is bool.
    IsComputed interface{}
}

func (tries *Radius_Nodes_Node_DeadCriteria_Hosts_Host_Tries) GetEntityData() *types.CommonEntityData {
    tries.EntityData.YFilter = tries.YFilter
    tries.EntityData.YangName = "tries"
    tries.EntityData.BundleName = "cisco_ios_xr"
    tries.EntityData.ParentYangName = "host"
    tries.EntityData.SegmentPath = "tries"
    tries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tries.EntityData.Children = make(map[string]types.YChild)
    tries.EntityData.Leafs = make(map[string]types.YLeaf)
    tries.EntityData.Leafs["value"] = types.YLeaf{"Value", tries.Value}
    tries.EntityData.Leafs["is-computed"] = types.YLeaf{"IsComputed", tries.IsComputed}
    return &(tries.EntityData)
}

// Radius_Nodes_Node_Authentication
// RADIUS authentication data
type Radius_Nodes_Node_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of authentication groups. The type is slice of
    // Radius_Nodes_Node_Authentication_AuthenticationGroup.
    AuthenticationGroup []Radius_Nodes_Node_Authentication_AuthenticationGroup
}

func (authentication *Radius_Nodes_Node_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "node"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["authentication-group"] = types.YChild{"AuthenticationGroup", nil}
    for i := range authentication.AuthenticationGroup {
        authentication.EntityData.Children[types.GetSegmentPath(&authentication.AuthenticationGroup[i])] = types.YChild{"AuthenticationGroup", &authentication.AuthenticationGroup[i]}
    }
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authentication.EntityData)
}

// Radius_Nodes_Node_Authentication_AuthenticationGroup
// List of authentication groups
type Radius_Nodes_Node_Authentication_AuthenticationGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerAddress interface{}

    // Authentication port number. The type is interface{} with range:
    // 0..4294967295.
    Port interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Authentication data.
    Authentication Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication_
}

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetEntityData() *types.CommonEntityData {
    authenticationGroup.EntityData.YFilter = authenticationGroup.YFilter
    authenticationGroup.EntityData.YangName = "authentication-group"
    authenticationGroup.EntityData.BundleName = "cisco_ios_xr"
    authenticationGroup.EntityData.ParentYangName = "authentication"
    authenticationGroup.EntityData.SegmentPath = "authentication-group"
    authenticationGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationGroup.EntityData.Children = make(map[string]types.YChild)
    authenticationGroup.EntityData.Children["authentication"] = types.YChild{"Authentication", &authenticationGroup.Authentication}
    authenticationGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    authenticationGroup.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", authenticationGroup.ServerAddress}
    authenticationGroup.EntityData.Leafs["port"] = types.YLeaf{"Port", authenticationGroup.Port}
    authenticationGroup.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", authenticationGroup.IpAddress}
    authenticationGroup.EntityData.Leafs["family"] = types.YLeaf{"Family", authenticationGroup.Family}
    return &(authenticationGroup.EntityData)
}

// Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication_
// Authentication data
type Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication_ struct {
    EntityData types.CommonEntityData
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

func (authentication_ *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication_) GetEntityData() *types.CommonEntityData {
    authentication_.EntityData.YFilter = authentication_.YFilter
    authentication_.EntityData.YangName = "authentication"
    authentication_.EntityData.BundleName = "cisco_ios_xr"
    authentication_.EntityData.ParentYangName = "authentication-group"
    authentication_.EntityData.SegmentPath = "authentication"
    authentication_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication_.EntityData.Children = make(map[string]types.YChild)
    authentication_.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication_.EntityData.Leafs["access-requests"] = types.YLeaf{"AccessRequests", authentication_.AccessRequests}
    authentication_.EntityData.Leafs["pending-access-requests"] = types.YLeaf{"PendingAccessRequests", authentication_.PendingAccessRequests}
    authentication_.EntityData.Leafs["access-request-retransmits"] = types.YLeaf{"AccessRequestRetransmits", authentication_.AccessRequestRetransmits}
    authentication_.EntityData.Leafs["access-accepts"] = types.YLeaf{"AccessAccepts", authentication_.AccessAccepts}
    authentication_.EntityData.Leafs["access-rejects"] = types.YLeaf{"AccessRejects", authentication_.AccessRejects}
    authentication_.EntityData.Leafs["access-challenges"] = types.YLeaf{"AccessChallenges", authentication_.AccessChallenges}
    authentication_.EntityData.Leafs["access-timeouts"] = types.YLeaf{"AccessTimeouts", authentication_.AccessTimeouts}
    authentication_.EntityData.Leafs["bad-access-responses"] = types.YLeaf{"BadAccessResponses", authentication_.BadAccessResponses}
    authentication_.EntityData.Leafs["bad-access-authenticators"] = types.YLeaf{"BadAccessAuthenticators", authentication_.BadAccessAuthenticators}
    authentication_.EntityData.Leafs["unknown-access-types"] = types.YLeaf{"UnknownAccessTypes", authentication_.UnknownAccessTypes}
    authentication_.EntityData.Leafs["dropped-access-responses"] = types.YLeaf{"DroppedAccessResponses", authentication_.DroppedAccessResponses}
    authentication_.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", authentication_.Rtt}
    authentication_.EntityData.Leafs["authen-response-time"] = types.YLeaf{"AuthenResponseTime", authentication_.AuthenResponseTime}
    authentication_.EntityData.Leafs["authen-transaction-successess"] = types.YLeaf{"AuthenTransactionSuccessess", authentication_.AuthenTransactionSuccessess}
    authentication_.EntityData.Leafs["authen-transaction-failure"] = types.YLeaf{"AuthenTransactionFailure", authentication_.AuthenTransactionFailure}
    authentication_.EntityData.Leafs["authen-unexpected-responses"] = types.YLeaf{"AuthenUnexpectedResponses", authentication_.AuthenUnexpectedResponses}
    authentication_.EntityData.Leafs["authen-server-error-responses"] = types.YLeaf{"AuthenServerErrorResponses", authentication_.AuthenServerErrorResponses}
    authentication_.EntityData.Leafs["authen-incorrect-responses"] = types.YLeaf{"AuthenIncorrectResponses", authentication_.AuthenIncorrectResponses}
    return &(authentication_.EntityData)
}

// Radius_Nodes_Node_Accounting
// RADIUS accounting data
type Radius_Nodes_Node_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of accounting groups. The type is slice of
    // Radius_Nodes_Node_Accounting_AccountingGroup.
    AccountingGroup []Radius_Nodes_Node_Accounting_AccountingGroup
}

func (accounting *Radius_Nodes_Node_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "node"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Children["accounting-group"] = types.YChild{"AccountingGroup", nil}
    for i := range accounting.AccountingGroup {
        accounting.EntityData.Children[types.GetSegmentPath(&accounting.AccountingGroup[i])] = types.YChild{"AccountingGroup", &accounting.AccountingGroup[i]}
    }
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accounting.EntityData)
}

// Radius_Nodes_Node_Accounting_AccountingGroup
// List of accounting groups
type Radius_Nodes_Node_Accounting_AccountingGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerAddress interface{}

    // Accounting port number. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Accounting data.
    Accounting Radius_Nodes_Node_Accounting_AccountingGroup_Accounting_
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetEntityData() *types.CommonEntityData {
    accountingGroup.EntityData.YFilter = accountingGroup.YFilter
    accountingGroup.EntityData.YangName = "accounting-group"
    accountingGroup.EntityData.BundleName = "cisco_ios_xr"
    accountingGroup.EntityData.ParentYangName = "accounting"
    accountingGroup.EntityData.SegmentPath = "accounting-group"
    accountingGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingGroup.EntityData.Children = make(map[string]types.YChild)
    accountingGroup.EntityData.Children["accounting"] = types.YChild{"Accounting", &accountingGroup.Accounting}
    accountingGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    accountingGroup.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", accountingGroup.ServerAddress}
    accountingGroup.EntityData.Leafs["port"] = types.YLeaf{"Port", accountingGroup.Port}
    accountingGroup.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", accountingGroup.IpAddress}
    accountingGroup.EntityData.Leafs["family"] = types.YLeaf{"Family", accountingGroup.Family}
    return &(accountingGroup.EntityData)
}

// Radius_Nodes_Node_Accounting_AccountingGroup_Accounting_
// Accounting data
type Radius_Nodes_Node_Accounting_AccountingGroup_Accounting_ struct {
    EntityData types.CommonEntityData
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

func (accounting_ *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting_) GetEntityData() *types.CommonEntityData {
    accounting_.EntityData.YFilter = accounting_.YFilter
    accounting_.EntityData.YangName = "accounting"
    accounting_.EntityData.BundleName = "cisco_ios_xr"
    accounting_.EntityData.ParentYangName = "accounting-group"
    accounting_.EntityData.SegmentPath = "accounting"
    accounting_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting_.EntityData.Children = make(map[string]types.YChild)
    accounting_.EntityData.Leafs = make(map[string]types.YLeaf)
    accounting_.EntityData.Leafs["requests"] = types.YLeaf{"Requests", accounting_.Requests}
    accounting_.EntityData.Leafs["pending-requests"] = types.YLeaf{"PendingRequests", accounting_.PendingRequests}
    accounting_.EntityData.Leafs["retransmits"] = types.YLeaf{"Retransmits", accounting_.Retransmits}
    accounting_.EntityData.Leafs["responses"] = types.YLeaf{"Responses", accounting_.Responses}
    accounting_.EntityData.Leafs["timeouts"] = types.YLeaf{"Timeouts", accounting_.Timeouts}
    accounting_.EntityData.Leafs["bad-responses"] = types.YLeaf{"BadResponses", accounting_.BadResponses}
    accounting_.EntityData.Leafs["bad-authenticators"] = types.YLeaf{"BadAuthenticators", accounting_.BadAuthenticators}
    accounting_.EntityData.Leafs["unknown-packet-types"] = types.YLeaf{"UnknownPacketTypes", accounting_.UnknownPacketTypes}
    accounting_.EntityData.Leafs["dropped-responses"] = types.YLeaf{"DroppedResponses", accounting_.DroppedResponses}
    accounting_.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", accounting_.Rtt}
    accounting_.EntityData.Leafs["acct-unexpected-responses"] = types.YLeaf{"AcctUnexpectedResponses", accounting_.AcctUnexpectedResponses}
    accounting_.EntityData.Leafs["acct-server-error-responses"] = types.YLeaf{"AcctServerErrorResponses", accounting_.AcctServerErrorResponses}
    accounting_.EntityData.Leafs["acct-incorrect-responses"] = types.YLeaf{"AcctIncorrectResponses", accounting_.AcctIncorrectResponses}
    accounting_.EntityData.Leafs["acct-response-time"] = types.YLeaf{"AcctResponseTime", accounting_.AcctResponseTime}
    accounting_.EntityData.Leafs["acct-transaction-successess"] = types.YLeaf{"AcctTransactionSuccessess", accounting_.AcctTransactionSuccessess}
    accounting_.EntityData.Leafs["acct-transaction-failure"] = types.YLeaf{"AcctTransactionFailure", accounting_.AcctTransactionFailure}
    return &(accounting_.EntityData)
}

// Radius_Nodes_Node_ServerGroups
// RADIUS server group table
type Radius_Nodes_Node_ServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS server group data. The type is slice of
    // Radius_Nodes_Node_ServerGroups_ServerGroup.
    ServerGroup []Radius_Nodes_Node_ServerGroups_ServerGroup
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetEntityData() *types.CommonEntityData {
    serverGroups.EntityData.YFilter = serverGroups.YFilter
    serverGroups.EntityData.YangName = "server-groups"
    serverGroups.EntityData.BundleName = "cisco_ios_xr"
    serverGroups.EntityData.ParentYangName = "node"
    serverGroups.EntityData.SegmentPath = "server-groups"
    serverGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroups.EntityData.Children = make(map[string]types.YChild)
    serverGroups.EntityData.Children["server-group"] = types.YChild{"ServerGroup", nil}
    for i := range serverGroups.ServerGroup {
        serverGroups.EntityData.Children[types.GetSegmentPath(&serverGroups.ServerGroup[i])] = types.YChild{"ServerGroup", &serverGroups.ServerGroup[i]}
    }
    serverGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serverGroups.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup
// RADIUS server group data
type Radius_Nodes_Node_ServerGroups_ServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    ServerGroup []Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetEntityData() *types.CommonEntityData {
    serverGroup.EntityData.YFilter = serverGroup.YFilter
    serverGroup.EntityData.YangName = "server-group"
    serverGroup.EntityData.BundleName = "cisco_ios_xr"
    serverGroup.EntityData.ParentYangName = "server-groups"
    serverGroup.EntityData.SegmentPath = "server-group" + "[server-group-name='" + fmt.Sprintf("%v", serverGroup.ServerGroupName) + "']"
    serverGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroup.EntityData.Children = make(map[string]types.YChild)
    serverGroup.EntityData.Children["server-group"] = types.YChild{"ServerGroup", nil}
    for i := range serverGroup.ServerGroup {
        serverGroup.EntityData.Children[types.GetSegmentPath(&serverGroup.ServerGroup[i])] = types.YChild{"ServerGroup", &serverGroup.ServerGroup[i]}
    }
    serverGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    serverGroup.EntityData.Leafs["server-group-name"] = types.YLeaf{"ServerGroupName", serverGroup.ServerGroupName}
    serverGroup.EntityData.Leafs["groups"] = types.YLeaf{"Groups", serverGroup.Groups}
    serverGroup.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", serverGroup.VrfName}
    serverGroup.EntityData.Leafs["dead-time"] = types.YLeaf{"DeadTime", serverGroup.DeadTime}
    serverGroup.EntityData.Leafs["servers"] = types.YLeaf{"Servers", serverGroup.Servers}
    return &(serverGroup.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_
// Server groups
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Server address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    Accounting Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Accounting

    // Authentication data.
    Authentication Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authentication

    // Authorization data.
    Authorization Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authorization
}

func (serverGroup_ *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_) GetEntityData() *types.CommonEntityData {
    serverGroup_.EntityData.YFilter = serverGroup_.YFilter
    serverGroup_.EntityData.YangName = "server-group"
    serverGroup_.EntityData.BundleName = "cisco_ios_xr"
    serverGroup_.EntityData.ParentYangName = "server-group"
    serverGroup_.EntityData.SegmentPath = "server-group"
    serverGroup_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroup_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroup_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroup_.EntityData.Children = make(map[string]types.YChild)
    serverGroup_.EntityData.Children["accounting"] = types.YChild{"Accounting", &serverGroup_.Accounting}
    serverGroup_.EntityData.Children["authentication"] = types.YChild{"Authentication", &serverGroup_.Authentication}
    serverGroup_.EntityData.Children["authorization"] = types.YChild{"Authorization", &serverGroup_.Authorization}
    serverGroup_.EntityData.Leafs = make(map[string]types.YLeaf)
    serverGroup_.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", serverGroup_.ServerAddress}
    serverGroup_.EntityData.Leafs["authentication-port"] = types.YLeaf{"AuthenticationPort", serverGroup_.AuthenticationPort}
    serverGroup_.EntityData.Leafs["accounting-port"] = types.YLeaf{"AccountingPort", serverGroup_.AccountingPort}
    serverGroup_.EntityData.Leafs["is-private"] = types.YLeaf{"IsPrivate", serverGroup_.IsPrivate}
    serverGroup_.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", serverGroup_.IpAddress}
    serverGroup_.EntityData.Leafs["family"] = types.YLeaf{"Family", serverGroup_.Family}
    return &(serverGroup_.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Accounting
// Accounting data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Accounting struct {
    EntityData types.CommonEntityData
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

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "server-group"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    accounting.EntityData.Leafs["requests"] = types.YLeaf{"Requests", accounting.Requests}
    accounting.EntityData.Leafs["pending-requests"] = types.YLeaf{"PendingRequests", accounting.PendingRequests}
    accounting.EntityData.Leafs["retransmits"] = types.YLeaf{"Retransmits", accounting.Retransmits}
    accounting.EntityData.Leafs["responses"] = types.YLeaf{"Responses", accounting.Responses}
    accounting.EntityData.Leafs["timeouts"] = types.YLeaf{"Timeouts", accounting.Timeouts}
    accounting.EntityData.Leafs["bad-responses"] = types.YLeaf{"BadResponses", accounting.BadResponses}
    accounting.EntityData.Leafs["bad-authenticators"] = types.YLeaf{"BadAuthenticators", accounting.BadAuthenticators}
    accounting.EntityData.Leafs["unknown-packet-types"] = types.YLeaf{"UnknownPacketTypes", accounting.UnknownPacketTypes}
    accounting.EntityData.Leafs["dropped-responses"] = types.YLeaf{"DroppedResponses", accounting.DroppedResponses}
    accounting.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", accounting.Rtt}
    accounting.EntityData.Leafs["acct-unexpected-responses"] = types.YLeaf{"AcctUnexpectedResponses", accounting.AcctUnexpectedResponses}
    accounting.EntityData.Leafs["acct-server-error-responses"] = types.YLeaf{"AcctServerErrorResponses", accounting.AcctServerErrorResponses}
    accounting.EntityData.Leafs["acct-incorrect-responses"] = types.YLeaf{"AcctIncorrectResponses", accounting.AcctIncorrectResponses}
    accounting.EntityData.Leafs["acct-response-time"] = types.YLeaf{"AcctResponseTime", accounting.AcctResponseTime}
    accounting.EntityData.Leafs["acct-transaction-successess"] = types.YLeaf{"AcctTransactionSuccessess", accounting.AcctTransactionSuccessess}
    accounting.EntityData.Leafs["acct-transaction-failure"] = types.YLeaf{"AcctTransactionFailure", accounting.AcctTransactionFailure}
    return &(accounting.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authentication
// Authentication data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authentication struct {
    EntityData types.CommonEntityData
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

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "server-group"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["access-requests"] = types.YLeaf{"AccessRequests", authentication.AccessRequests}
    authentication.EntityData.Leafs["pending-access-requests"] = types.YLeaf{"PendingAccessRequests", authentication.PendingAccessRequests}
    authentication.EntityData.Leafs["access-request-retransmits"] = types.YLeaf{"AccessRequestRetransmits", authentication.AccessRequestRetransmits}
    authentication.EntityData.Leafs["access-accepts"] = types.YLeaf{"AccessAccepts", authentication.AccessAccepts}
    authentication.EntityData.Leafs["access-rejects"] = types.YLeaf{"AccessRejects", authentication.AccessRejects}
    authentication.EntityData.Leafs["access-challenges"] = types.YLeaf{"AccessChallenges", authentication.AccessChallenges}
    authentication.EntityData.Leafs["access-timeouts"] = types.YLeaf{"AccessTimeouts", authentication.AccessTimeouts}
    authentication.EntityData.Leafs["bad-access-responses"] = types.YLeaf{"BadAccessResponses", authentication.BadAccessResponses}
    authentication.EntityData.Leafs["bad-access-authenticators"] = types.YLeaf{"BadAccessAuthenticators", authentication.BadAccessAuthenticators}
    authentication.EntityData.Leafs["unknown-access-types"] = types.YLeaf{"UnknownAccessTypes", authentication.UnknownAccessTypes}
    authentication.EntityData.Leafs["dropped-access-responses"] = types.YLeaf{"DroppedAccessResponses", authentication.DroppedAccessResponses}
    authentication.EntityData.Leafs["rtt"] = types.YLeaf{"Rtt", authentication.Rtt}
    authentication.EntityData.Leafs["authen-response-time"] = types.YLeaf{"AuthenResponseTime", authentication.AuthenResponseTime}
    authentication.EntityData.Leafs["authen-transaction-successess"] = types.YLeaf{"AuthenTransactionSuccessess", authentication.AuthenTransactionSuccessess}
    authentication.EntityData.Leafs["authen-transaction-failure"] = types.YLeaf{"AuthenTransactionFailure", authentication.AuthenTransactionFailure}
    authentication.EntityData.Leafs["authen-unexpected-responses"] = types.YLeaf{"AuthenUnexpectedResponses", authentication.AuthenUnexpectedResponses}
    authentication.EntityData.Leafs["authen-server-error-responses"] = types.YLeaf{"AuthenServerErrorResponses", authentication.AuthenServerErrorResponses}
    authentication.EntityData.Leafs["authen-incorrect-responses"] = types.YLeaf{"AuthenIncorrectResponses", authentication.AuthenIncorrectResponses}
    return &(authentication.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authorization
// Authorization data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authorization struct {
    EntityData types.CommonEntityData
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

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup__Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "server-group"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = make(map[string]types.YChild)
    authorization.EntityData.Leafs = make(map[string]types.YLeaf)
    authorization.EntityData.Leafs["author-requests"] = types.YLeaf{"AuthorRequests", authorization.AuthorRequests}
    authorization.EntityData.Leafs["author-request-timeouts"] = types.YLeaf{"AuthorRequestTimeouts", authorization.AuthorRequestTimeouts}
    authorization.EntityData.Leafs["author-unexpected-responses"] = types.YLeaf{"AuthorUnexpectedResponses", authorization.AuthorUnexpectedResponses}
    authorization.EntityData.Leafs["author-server-error-responses"] = types.YLeaf{"AuthorServerErrorResponses", authorization.AuthorServerErrorResponses}
    authorization.EntityData.Leafs["author-incorrect-responses"] = types.YLeaf{"AuthorIncorrectResponses", authorization.AuthorIncorrectResponses}
    authorization.EntityData.Leafs["author-response-time"] = types.YLeaf{"AuthorResponseTime", authorization.AuthorResponseTime}
    authorization.EntityData.Leafs["author-transaction-successess"] = types.YLeaf{"AuthorTransactionSuccessess", authorization.AuthorTransactionSuccessess}
    authorization.EntityData.Leafs["author-transaction-failure"] = types.YLeaf{"AuthorTransactionFailure", authorization.AuthorTransactionFailure}
    return &(authorization.EntityData)
}

// Radius_Nodes_Node_DynamicAuthorization
// Dynamic authorization data
type Radius_Nodes_Node_DynamicAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Invalid disconnected requests. The type is interface{} with range:
    // 0..4294967295.
    DisconnectedInvalidRequests interface{}

    // Invalid change of authorization requests. The type is interface{} with
    // range: 0..4294967295.
    InvalidCoaRequests interface{}
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetEntityData() *types.CommonEntityData {
    dynamicAuthorization.EntityData.YFilter = dynamicAuthorization.YFilter
    dynamicAuthorization.EntityData.YangName = "dynamic-authorization"
    dynamicAuthorization.EntityData.BundleName = "cisco_ios_xr"
    dynamicAuthorization.EntityData.ParentYangName = "node"
    dynamicAuthorization.EntityData.SegmentPath = "dynamic-authorization"
    dynamicAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicAuthorization.EntityData.Children = make(map[string]types.YChild)
    dynamicAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    dynamicAuthorization.EntityData.Leafs["disconnected-invalid-requests"] = types.YLeaf{"DisconnectedInvalidRequests", dynamicAuthorization.DisconnectedInvalidRequests}
    dynamicAuthorization.EntityData.Leafs["invalid-coa-requests"] = types.YLeaf{"InvalidCoaRequests", dynamicAuthorization.InvalidCoaRequests}
    return &(dynamicAuthorization.EntityData)
}

