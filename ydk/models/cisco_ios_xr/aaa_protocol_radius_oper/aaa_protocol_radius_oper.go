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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    radius.EntityData.AbsolutePath = radius.EntityData.SegmentPath
    radius.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radius.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radius.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radius.EntityData.Children = types.NewOrderedMap()
    radius.EntityData.Children.Append("nodes", types.YChild{"Nodes", &radius.Nodes})
    radius.EntityData.Leafs = types.NewOrderedMap()

    radius.EntityData.YListKeys = []string {}

    return &(radius.EntityData)
}

// Radius_Nodes
// Contains all the nodes
type Radius_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS operational data for a particular node. The type is slice of
    // Radius_Nodes_Node.
    Node []*Radius_Nodes_Node
}

func (nodes *Radius_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "radius"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/" + nodes.EntityData.SegmentPath
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

// Radius_Nodes_Node
// RADIUS operational data for a particular node
type Radius_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Dynamic authorization client data.
    DynamicAuthorizationClients Radius_Nodes_Node_DynamicAuthorizationClients

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
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("client", types.YChild{"Client", &node.Client})
    node.EntityData.Children.Append("dead-criteria", types.YChild{"DeadCriteria", &node.DeadCriteria})
    node.EntityData.Children.Append("authentication", types.YChild{"Authentication", &node.Authentication})
    node.EntityData.Children.Append("accounting", types.YChild{"Accounting", &node.Accounting})
    node.EntityData.Children.Append("dynamic-authorization-clients", types.YChild{"DynamicAuthorizationClients", &node.DynamicAuthorizationClients})
    node.EntityData.Children.Append("server-groups", types.YChild{"ServerGroups", &node.ServerGroups})
    node.EntityData.Children.Append("dynamic-authorization", types.YChild{"DynamicAuthorization", &node.DynamicAuthorization})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

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
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("unknown-authentication-responses", types.YLeaf{"UnknownAuthenticationResponses", client.UnknownAuthenticationResponses})
    client.EntityData.Leafs.Append("authentication-nas-id", types.YLeaf{"AuthenticationNasId", client.AuthenticationNasId})
    client.EntityData.Leafs.Append("unknown-accounting-responses", types.YLeaf{"UnknownAccountingResponses", client.UnknownAccountingResponses})

    client.EntityData.YListKeys = []string {}

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
    deadCriteria.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + deadCriteria.EntityData.SegmentPath
    deadCriteria.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deadCriteria.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deadCriteria.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deadCriteria.EntityData.Children = types.NewOrderedMap()
    deadCriteria.EntityData.Children.Append("hosts", types.YChild{"Hosts", &deadCriteria.Hosts})
    deadCriteria.EntityData.Leafs = types.NewOrderedMap()

    deadCriteria.EntityData.YListKeys = []string {}

    return &(deadCriteria.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts
// RADIUS server dead criteria host table
type Radius_Nodes_Node_DeadCriteria_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS Server. The type is slice of
    // Radius_Nodes_Node_DeadCriteria_Hosts_Host.
    Host []*Radius_Nodes_Node_DeadCriteria_Hosts_Host
}

func (hosts *Radius_Nodes_Node_DeadCriteria_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "dead-criteria"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/dead-criteria/" + hosts.EntityData.SegmentPath
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        types.SetYListKey(hosts.Host[i], i)
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Radius_Nodes_Node_DeadCriteria_Hosts_Host
// RADIUS Server
type Radius_Nodes_Node_DeadCriteria_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    host.EntityData.SegmentPath = "host" + types.AddNoKeyToken(host)
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/dead-criteria/hosts/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Children.Append("time", types.YChild{"Time", &host.Time})
    host.EntityData.Children.Append("tries", types.YChild{"Tries", &host.Tries})
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", host.IpAddress})
    host.EntityData.Leafs.Append("auth-port-number", types.YLeaf{"AuthPortNumber", host.AuthPortNumber})
    host.EntityData.Leafs.Append("acct-port-number", types.YLeaf{"AcctPortNumber", host.AcctPortNumber})

    host.EntityData.YListKeys = []string {}

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
    time.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/dead-criteria/hosts/host/" + time.EntityData.SegmentPath
    time.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    time.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    time.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    time.EntityData.Children = types.NewOrderedMap()
    time.EntityData.Leafs = types.NewOrderedMap()
    time.EntityData.Leafs.Append("value", types.YLeaf{"Value", time.Value})
    time.EntityData.Leafs.Append("is-computed", types.YLeaf{"IsComputed", time.IsComputed})

    time.EntityData.YListKeys = []string {}

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
    tries.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/dead-criteria/hosts/host/" + tries.EntityData.SegmentPath
    tries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tries.EntityData.Children = types.NewOrderedMap()
    tries.EntityData.Leafs = types.NewOrderedMap()
    tries.EntityData.Leafs.Append("value", types.YLeaf{"Value", tries.Value})
    tries.EntityData.Leafs.Append("is-computed", types.YLeaf{"IsComputed", tries.IsComputed})

    tries.EntityData.YListKeys = []string {}

    return &(tries.EntityData)
}

// Radius_Nodes_Node_Authentication
// RADIUS authentication data
type Radius_Nodes_Node_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of authentication groups. The type is slice of
    // Radius_Nodes_Node_Authentication_AuthenticationGroup.
    AuthenticationGroup []*Radius_Nodes_Node_Authentication_AuthenticationGroup
}

func (authentication *Radius_Nodes_Node_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "node"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("authentication-group", types.YChild{"AuthenticationGroup", nil})
    for i := range authentication.AuthenticationGroup {
        types.SetYListKey(authentication.AuthenticationGroup[i], i)
        authentication.EntityData.Children.Append(types.GetSegmentPath(authentication.AuthenticationGroup[i]), types.YChild{"AuthenticationGroup", authentication.AuthenticationGroup[i]})
    }
    authentication.EntityData.Leafs = types.NewOrderedMap()

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Radius_Nodes_Node_Authentication_AuthenticationGroup
// List of authentication groups
type Radius_Nodes_Node_Authentication_AuthenticationGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (authenticationGroup *Radius_Nodes_Node_Authentication_AuthenticationGroup) GetEntityData() *types.CommonEntityData {
    authenticationGroup.EntityData.YFilter = authenticationGroup.YFilter
    authenticationGroup.EntityData.YangName = "authentication-group"
    authenticationGroup.EntityData.BundleName = "cisco_ios_xr"
    authenticationGroup.EntityData.ParentYangName = "authentication"
    authenticationGroup.EntityData.SegmentPath = "authentication-group" + types.AddNoKeyToken(authenticationGroup)
    authenticationGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/authentication/" + authenticationGroup.EntityData.SegmentPath
    authenticationGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationGroup.EntityData.Children = types.NewOrderedMap()
    authenticationGroup.EntityData.Children.Append("authentication", types.YChild{"Authentication", &authenticationGroup.Authentication})
    authenticationGroup.EntityData.Leafs = types.NewOrderedMap()
    authenticationGroup.EntityData.Leafs.Append("port", types.YLeaf{"Port", authenticationGroup.Port})
    authenticationGroup.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", authenticationGroup.IpAddress})
    authenticationGroup.EntityData.Leafs.Append("family", types.YLeaf{"Family", authenticationGroup.Family})

    authenticationGroup.EntityData.YListKeys = []string {}

    return &(authenticationGroup.EntityData)
}

// Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication
// Authentication data
type Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication struct {
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

    // Estimated Throttled Authentication Transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthThrottledTransactions interface{}

    // Maximum Throttled Authentication Transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthMaxTransactions interface{}

    // Automated Test Stats for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthReqs interface{}

    // Automated Test Stats for authentication timeouts. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthTimeouts interface{}

    // Automated Test Stats for authentication response. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthResponse interface{}

    // Automated Test Stats for authentication pending. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthPending interface{}
}

func (authentication *Radius_Nodes_Node_Authentication_AuthenticationGroup_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "authentication-group"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/authentication/authentication-group/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("access-requests", types.YLeaf{"AccessRequests", authentication.AccessRequests})
    authentication.EntityData.Leafs.Append("pending-access-requests", types.YLeaf{"PendingAccessRequests", authentication.PendingAccessRequests})
    authentication.EntityData.Leafs.Append("access-request-retransmits", types.YLeaf{"AccessRequestRetransmits", authentication.AccessRequestRetransmits})
    authentication.EntityData.Leafs.Append("access-accepts", types.YLeaf{"AccessAccepts", authentication.AccessAccepts})
    authentication.EntityData.Leafs.Append("access-rejects", types.YLeaf{"AccessRejects", authentication.AccessRejects})
    authentication.EntityData.Leafs.Append("access-challenges", types.YLeaf{"AccessChallenges", authentication.AccessChallenges})
    authentication.EntityData.Leafs.Append("access-timeouts", types.YLeaf{"AccessTimeouts", authentication.AccessTimeouts})
    authentication.EntityData.Leafs.Append("bad-access-responses", types.YLeaf{"BadAccessResponses", authentication.BadAccessResponses})
    authentication.EntityData.Leafs.Append("bad-access-authenticators", types.YLeaf{"BadAccessAuthenticators", authentication.BadAccessAuthenticators})
    authentication.EntityData.Leafs.Append("unknown-access-types", types.YLeaf{"UnknownAccessTypes", authentication.UnknownAccessTypes})
    authentication.EntityData.Leafs.Append("dropped-access-responses", types.YLeaf{"DroppedAccessResponses", authentication.DroppedAccessResponses})
    authentication.EntityData.Leafs.Append("rtt", types.YLeaf{"Rtt", authentication.Rtt})
    authentication.EntityData.Leafs.Append("authen-transaction-successess", types.YLeaf{"AuthenTransactionSuccessess", authentication.AuthenTransactionSuccessess})
    authentication.EntityData.Leafs.Append("authen-transaction-failure", types.YLeaf{"AuthenTransactionFailure", authentication.AuthenTransactionFailure})
    authentication.EntityData.Leafs.Append("authen-unexpected-responses", types.YLeaf{"AuthenUnexpectedResponses", authentication.AuthenUnexpectedResponses})
    authentication.EntityData.Leafs.Append("authen-server-error-responses", types.YLeaf{"AuthenServerErrorResponses", authentication.AuthenServerErrorResponses})
    authentication.EntityData.Leafs.Append("authen-incorrect-responses", types.YLeaf{"AuthenIncorrectResponses", authentication.AuthenIncorrectResponses})
    authentication.EntityData.Leafs.Append("auth-throttled-transactions", types.YLeaf{"AuthThrottledTransactions", authentication.AuthThrottledTransactions})
    authentication.EntityData.Leafs.Append("auth-max-transactions", types.YLeaf{"AuthMaxTransactions", authentication.AuthMaxTransactions})
    authentication.EntityData.Leafs.Append("total-test-auth-reqs", types.YLeaf{"TotalTestAuthReqs", authentication.TotalTestAuthReqs})
    authentication.EntityData.Leafs.Append("total-test-auth-timeouts", types.YLeaf{"TotalTestAuthTimeouts", authentication.TotalTestAuthTimeouts})
    authentication.EntityData.Leafs.Append("total-test-auth-response", types.YLeaf{"TotalTestAuthResponse", authentication.TotalTestAuthResponse})
    authentication.EntityData.Leafs.Append("total-test-auth-pending", types.YLeaf{"TotalTestAuthPending", authentication.TotalTestAuthPending})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Radius_Nodes_Node_Accounting
// RADIUS accounting data
type Radius_Nodes_Node_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of accounting groups. The type is slice of
    // Radius_Nodes_Node_Accounting_AccountingGroup.
    AccountingGroup []*Radius_Nodes_Node_Accounting_AccountingGroup
}

func (accounting *Radius_Nodes_Node_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "node"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("accounting-group", types.YChild{"AccountingGroup", nil})
    for i := range accounting.AccountingGroup {
        types.SetYListKey(accounting.AccountingGroup[i], i)
        accounting.EntityData.Children.Append(types.GetSegmentPath(accounting.AccountingGroup[i]), types.YChild{"AccountingGroup", accounting.AccountingGroup[i]})
    }
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Radius_Nodes_Node_Accounting_AccountingGroup
// List of accounting groups
type Radius_Nodes_Node_Accounting_AccountingGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Accounting port number. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // IP address buffer. The type is string.
    IpAddress interface{}

    // IP address Family. The type is string.
    Family interface{}

    // Accounting data.
    Accounting Radius_Nodes_Node_Accounting_AccountingGroup_Accounting
}

func (accountingGroup *Radius_Nodes_Node_Accounting_AccountingGroup) GetEntityData() *types.CommonEntityData {
    accountingGroup.EntityData.YFilter = accountingGroup.YFilter
    accountingGroup.EntityData.YangName = "accounting-group"
    accountingGroup.EntityData.BundleName = "cisco_ios_xr"
    accountingGroup.EntityData.ParentYangName = "accounting"
    accountingGroup.EntityData.SegmentPath = "accounting-group" + types.AddNoKeyToken(accountingGroup)
    accountingGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/accounting/" + accountingGroup.EntityData.SegmentPath
    accountingGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingGroup.EntityData.Children = types.NewOrderedMap()
    accountingGroup.EntityData.Children.Append("accounting", types.YChild{"Accounting", &accountingGroup.Accounting})
    accountingGroup.EntityData.Leafs = types.NewOrderedMap()
    accountingGroup.EntityData.Leafs.Append("port", types.YLeaf{"Port", accountingGroup.Port})
    accountingGroup.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", accountingGroup.IpAddress})
    accountingGroup.EntityData.Leafs.Append("family", types.YLeaf{"Family", accountingGroup.Family})

    accountingGroup.EntityData.YListKeys = []string {}

    return &(accountingGroup.EntityData)
}

// Radius_Nodes_Node_Accounting_AccountingGroup_Accounting
// Accounting data
type Radius_Nodes_Node_Accounting_AccountingGroup_Accounting struct {
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

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AcctTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctTransactionFailure interface{}

    // Estimated Throttled Accounting Transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctThrottledTransactions interface{}

    // Maximum Throttled Accounting Transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctMaxThrottleTrans interface{}

    // Automated Test Stats for accounting requests. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctReqs interface{}

    // Automated Test Stats for accounting timeouts. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctTimeouts interface{}

    // Automated Test Stats for accounting response. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctResponse interface{}

    // Automated Test Stats for accounting pending. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctPending interface{}
}

func (accounting *Radius_Nodes_Node_Accounting_AccountingGroup_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "accounting-group"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/accounting/accounting-group/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("requests", types.YLeaf{"Requests", accounting.Requests})
    accounting.EntityData.Leafs.Append("pending-requests", types.YLeaf{"PendingRequests", accounting.PendingRequests})
    accounting.EntityData.Leafs.Append("retransmits", types.YLeaf{"Retransmits", accounting.Retransmits})
    accounting.EntityData.Leafs.Append("responses", types.YLeaf{"Responses", accounting.Responses})
    accounting.EntityData.Leafs.Append("timeouts", types.YLeaf{"Timeouts", accounting.Timeouts})
    accounting.EntityData.Leafs.Append("bad-responses", types.YLeaf{"BadResponses", accounting.BadResponses})
    accounting.EntityData.Leafs.Append("bad-authenticators", types.YLeaf{"BadAuthenticators", accounting.BadAuthenticators})
    accounting.EntityData.Leafs.Append("unknown-packet-types", types.YLeaf{"UnknownPacketTypes", accounting.UnknownPacketTypes})
    accounting.EntityData.Leafs.Append("dropped-responses", types.YLeaf{"DroppedResponses", accounting.DroppedResponses})
    accounting.EntityData.Leafs.Append("rtt", types.YLeaf{"Rtt", accounting.Rtt})
    accounting.EntityData.Leafs.Append("acct-unexpected-responses", types.YLeaf{"AcctUnexpectedResponses", accounting.AcctUnexpectedResponses})
    accounting.EntityData.Leafs.Append("acct-transaction-successess", types.YLeaf{"AcctTransactionSuccessess", accounting.AcctTransactionSuccessess})
    accounting.EntityData.Leafs.Append("acct-transaction-failure", types.YLeaf{"AcctTransactionFailure", accounting.AcctTransactionFailure})
    accounting.EntityData.Leafs.Append("acct-throttled-transactions", types.YLeaf{"AcctThrottledTransactions", accounting.AcctThrottledTransactions})
    accounting.EntityData.Leafs.Append("acct-max-throttle-trans", types.YLeaf{"AcctMaxThrottleTrans", accounting.AcctMaxThrottleTrans})
    accounting.EntityData.Leafs.Append("total-test-acct-reqs", types.YLeaf{"TotalTestAcctReqs", accounting.TotalTestAcctReqs})
    accounting.EntityData.Leafs.Append("total-test-acct-timeouts", types.YLeaf{"TotalTestAcctTimeouts", accounting.TotalTestAcctTimeouts})
    accounting.EntityData.Leafs.Append("total-test-acct-response", types.YLeaf{"TotalTestAcctResponse", accounting.TotalTestAcctResponse})
    accounting.EntityData.Leafs.Append("total-test-acct-pending", types.YLeaf{"TotalTestAcctPending", accounting.TotalTestAcctPending})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Radius_Nodes_Node_DynamicAuthorizationClients
// Dynamic authorization client data
type Radius_Nodes_Node_DynamicAuthorizationClients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of dynamic author clients. The type is slice of
    // Radius_Nodes_Node_DynamicAuthorizationClients_DynamicAuthorClient.
    DynamicAuthorClient []*Radius_Nodes_Node_DynamicAuthorizationClients_DynamicAuthorClient
}

func (dynamicAuthorizationClients *Radius_Nodes_Node_DynamicAuthorizationClients) GetEntityData() *types.CommonEntityData {
    dynamicAuthorizationClients.EntityData.YFilter = dynamicAuthorizationClients.YFilter
    dynamicAuthorizationClients.EntityData.YangName = "dynamic-authorization-clients"
    dynamicAuthorizationClients.EntityData.BundleName = "cisco_ios_xr"
    dynamicAuthorizationClients.EntityData.ParentYangName = "node"
    dynamicAuthorizationClients.EntityData.SegmentPath = "dynamic-authorization-clients"
    dynamicAuthorizationClients.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + dynamicAuthorizationClients.EntityData.SegmentPath
    dynamicAuthorizationClients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicAuthorizationClients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicAuthorizationClients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicAuthorizationClients.EntityData.Children = types.NewOrderedMap()
    dynamicAuthorizationClients.EntityData.Children.Append("dynamic-author-client", types.YChild{"DynamicAuthorClient", nil})
    for i := range dynamicAuthorizationClients.DynamicAuthorClient {
        types.SetYListKey(dynamicAuthorizationClients.DynamicAuthorClient[i], i)
        dynamicAuthorizationClients.EntityData.Children.Append(types.GetSegmentPath(dynamicAuthorizationClients.DynamicAuthorClient[i]), types.YChild{"DynamicAuthorClient", dynamicAuthorizationClients.DynamicAuthorClient[i]})
    }
    dynamicAuthorizationClients.EntityData.Leafs = types.NewOrderedMap()

    dynamicAuthorizationClients.EntityData.YListKeys = []string {}

    return &(dynamicAuthorizationClients.EntityData)
}

// Radius_Nodes_Node_DynamicAuthorizationClients_DynamicAuthorClient
// List of dynamic author clients
type Radius_Nodes_Node_DynamicAuthorizationClients_DynamicAuthorClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Number of RADIUS Disconnect-Requestsreceived from the client. The type is
    // interface{} with range: 0..4294967295.
    DiscReqs interface{}

    // Number of RADIUS Disconnect-ACKs sent to the client. The type is
    // interface{} with range: 0..4294967295.
    DiscAcks interface{}

    // Number of RADIUS Disconnect-NAKs sent to the client. The type is
    // interface{} with range: 0..4294967295.
    DiscNaks interface{}

    // Number of RADIUS Disconnect-Requests received from the client containing an
    // invalid Authenticator. The type is interface{} with range: 0..4294967295.
    DiscBadAuth interface{}

    // Number of RADIUS Disconnect-Requests received from the client that were
    // silently discarded. The type is interface{} with range: 0..4294967295.
    DropDiscReqs interface{}

    // Number of RADIUS CoA-Requests received from the client. The type is
    // interface{} with range: 0..4294967295.
    CoaReqs interface{}

    // Number of RADIUS CoA-ACKs sent to the client. The type is interface{} with
    // range: 0..4294967295.
    CoaAcks interface{}

    // Number of RADIUS CoA-NAKs sent to the client. The type is interface{} with
    // range: 0..4294967295.
    CoaNaks interface{}

    // Number of RADIUS CoA-Requests received from the client containing an
    // invalid Authenticator. The type is interface{} with range: 0..4294967295.
    CoaBadAuth interface{}

    // Number of RADIUS CoA-Requests received from the client that were silently
    // discarded. The type is interface{} with range: 0..4294967295.
    DropCoaReqs interface{}

    // Number of incoming packets of unknown types that were received from the
    // client. The type is interface{} with range: 0..4294967295.
    UnknownTypes interface{}

    // Number of packets dropped due to internal error. The type is interface{}
    // with range: 0..4294967295.
    InternalError interface{}

    // Number of packets dropped due to failure in radius packet decoding. The
    // type is interface{} with range: 0..4294967295.
    PakDecodeFail interface{}

    // Number of requests which encountered vrf parse fail error. The type is
    // interface{} with range: 0..4294967295.
    VrfParseFailErr interface{}

    // Number of requests which encountered unknown vsa error. The type is
    // interface{} with range: 0..4294967295.
    UnknownVsaError interface{}

    // Number of response packets which failed to be send. The type is interface{}
    // with range: 0..4294967295.
    SendMsgFailed interface{}

    // Number of requests sent to command handler. The type is interface{} with
    // range: 0..4294967295.
    RadiusToCh interface{}

    // Number of responses received from command handler. The type is interface{}
    // with range: 0..4294967295.
    ChToRadius interface{}

    // Number of requests which encountered service parse fail error. The type is
    // interface{} with range: 0..4294967295.
    ServiceParseFail interface{}

    // Number of requests which encountered multiple subscribers not allowed
    // error. The type is interface{} with range: 0..4294967295.
    MultiSubsError interface{}

    // Number of requests which has missing service name attribute. The type is
    // interface{} with range: 0..4294967295.
    ServiceNotPresent interface{}

    // Number of requests which failed to be sent to command handler. The type is
    // interface{} with range: 0..4294967295.
    SendToChFail interface{}

    // VRF of RADIUS dynamic authorization client. The type is string.
    VrfName interface{}

    // Address Buffer. The type is string.
    ClientAddress interface{}
}

func (dynamicAuthorClient *Radius_Nodes_Node_DynamicAuthorizationClients_DynamicAuthorClient) GetEntityData() *types.CommonEntityData {
    dynamicAuthorClient.EntityData.YFilter = dynamicAuthorClient.YFilter
    dynamicAuthorClient.EntityData.YangName = "dynamic-author-client"
    dynamicAuthorClient.EntityData.BundleName = "cisco_ios_xr"
    dynamicAuthorClient.EntityData.ParentYangName = "dynamic-authorization-clients"
    dynamicAuthorClient.EntityData.SegmentPath = "dynamic-author-client" + types.AddNoKeyToken(dynamicAuthorClient)
    dynamicAuthorClient.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/dynamic-authorization-clients/" + dynamicAuthorClient.EntityData.SegmentPath
    dynamicAuthorClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicAuthorClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicAuthorClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicAuthorClient.EntityData.Children = types.NewOrderedMap()
    dynamicAuthorClient.EntityData.Leafs = types.NewOrderedMap()
    dynamicAuthorClient.EntityData.Leafs.Append("disc-reqs", types.YLeaf{"DiscReqs", dynamicAuthorClient.DiscReqs})
    dynamicAuthorClient.EntityData.Leafs.Append("disc-acks", types.YLeaf{"DiscAcks", dynamicAuthorClient.DiscAcks})
    dynamicAuthorClient.EntityData.Leafs.Append("disc-naks", types.YLeaf{"DiscNaks", dynamicAuthorClient.DiscNaks})
    dynamicAuthorClient.EntityData.Leafs.Append("disc-bad-auth", types.YLeaf{"DiscBadAuth", dynamicAuthorClient.DiscBadAuth})
    dynamicAuthorClient.EntityData.Leafs.Append("drop-disc-reqs", types.YLeaf{"DropDiscReqs", dynamicAuthorClient.DropDiscReqs})
    dynamicAuthorClient.EntityData.Leafs.Append("coa-reqs", types.YLeaf{"CoaReqs", dynamicAuthorClient.CoaReqs})
    dynamicAuthorClient.EntityData.Leafs.Append("coa-acks", types.YLeaf{"CoaAcks", dynamicAuthorClient.CoaAcks})
    dynamicAuthorClient.EntityData.Leafs.Append("coa-naks", types.YLeaf{"CoaNaks", dynamicAuthorClient.CoaNaks})
    dynamicAuthorClient.EntityData.Leafs.Append("coa-bad-auth", types.YLeaf{"CoaBadAuth", dynamicAuthorClient.CoaBadAuth})
    dynamicAuthorClient.EntityData.Leafs.Append("drop-coa-reqs", types.YLeaf{"DropCoaReqs", dynamicAuthorClient.DropCoaReqs})
    dynamicAuthorClient.EntityData.Leafs.Append("unknown-types", types.YLeaf{"UnknownTypes", dynamicAuthorClient.UnknownTypes})
    dynamicAuthorClient.EntityData.Leafs.Append("internal-error", types.YLeaf{"InternalError", dynamicAuthorClient.InternalError})
    dynamicAuthorClient.EntityData.Leafs.Append("pak-decode-fail", types.YLeaf{"PakDecodeFail", dynamicAuthorClient.PakDecodeFail})
    dynamicAuthorClient.EntityData.Leafs.Append("vrf-parse-fail-err", types.YLeaf{"VrfParseFailErr", dynamicAuthorClient.VrfParseFailErr})
    dynamicAuthorClient.EntityData.Leafs.Append("unknown-vsa-error", types.YLeaf{"UnknownVsaError", dynamicAuthorClient.UnknownVsaError})
    dynamicAuthorClient.EntityData.Leafs.Append("send-msg-failed", types.YLeaf{"SendMsgFailed", dynamicAuthorClient.SendMsgFailed})
    dynamicAuthorClient.EntityData.Leafs.Append("radius-to-ch", types.YLeaf{"RadiusToCh", dynamicAuthorClient.RadiusToCh})
    dynamicAuthorClient.EntityData.Leafs.Append("ch-to-radius", types.YLeaf{"ChToRadius", dynamicAuthorClient.ChToRadius})
    dynamicAuthorClient.EntityData.Leafs.Append("service-parse-fail", types.YLeaf{"ServiceParseFail", dynamicAuthorClient.ServiceParseFail})
    dynamicAuthorClient.EntityData.Leafs.Append("multi-subs-error", types.YLeaf{"MultiSubsError", dynamicAuthorClient.MultiSubsError})
    dynamicAuthorClient.EntityData.Leafs.Append("service-not-present", types.YLeaf{"ServiceNotPresent", dynamicAuthorClient.ServiceNotPresent})
    dynamicAuthorClient.EntityData.Leafs.Append("send-to-ch-fail", types.YLeaf{"SendToChFail", dynamicAuthorClient.SendToChFail})
    dynamicAuthorClient.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", dynamicAuthorClient.VrfName})
    dynamicAuthorClient.EntityData.Leafs.Append("client-address", types.YLeaf{"ClientAddress", dynamicAuthorClient.ClientAddress})

    dynamicAuthorClient.EntityData.YListKeys = []string {}

    return &(dynamicAuthorClient.EntityData)
}

// Radius_Nodes_Node_ServerGroups
// RADIUS server group table
type Radius_Nodes_Node_ServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS server group data. The type is slice of
    // Radius_Nodes_Node_ServerGroups_ServerGroup.
    ServerGroup []*Radius_Nodes_Node_ServerGroups_ServerGroup
}

func (serverGroups *Radius_Nodes_Node_ServerGroups) GetEntityData() *types.CommonEntityData {
    serverGroups.EntityData.YFilter = serverGroups.YFilter
    serverGroups.EntityData.YangName = "server-groups"
    serverGroups.EntityData.BundleName = "cisco_ios_xr"
    serverGroups.EntityData.ParentYangName = "node"
    serverGroups.EntityData.SegmentPath = "server-groups"
    serverGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + serverGroups.EntityData.SegmentPath
    serverGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroups.EntityData.Children = types.NewOrderedMap()
    serverGroups.EntityData.Children.Append("server-group", types.YChild{"ServerGroup", nil})
    for i := range serverGroups.ServerGroup {
        serverGroups.EntityData.Children.Append(types.GetSegmentPath(serverGroups.ServerGroup[i]), types.YChild{"ServerGroup", serverGroups.ServerGroup[i]})
    }
    serverGroups.EntityData.Leafs = types.NewOrderedMap()

    serverGroups.EntityData.YListKeys = []string {}

    return &(serverGroups.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup
// RADIUS server group data
type Radius_Nodes_Node_ServerGroups_ServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    ServerGroup []*Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup) GetEntityData() *types.CommonEntityData {
    serverGroup.EntityData.YFilter = serverGroup.YFilter
    serverGroup.EntityData.YangName = "server-group"
    serverGroup.EntityData.BundleName = "cisco_ios_xr"
    serverGroup.EntityData.ParentYangName = "server-groups"
    serverGroup.EntityData.SegmentPath = "server-group" + types.AddKeyToken(serverGroup.ServerGroupName, "server-group-name")
    serverGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/server-groups/" + serverGroup.EntityData.SegmentPath
    serverGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroup.EntityData.Children = types.NewOrderedMap()
    serverGroup.EntityData.Children.Append("server-group", types.YChild{"ServerGroup", nil})
    for i := range serverGroup.ServerGroup {
        types.SetYListKey(serverGroup.ServerGroup[i], i)
        serverGroup.EntityData.Children.Append(types.GetSegmentPath(serverGroup.ServerGroup[i]), types.YChild{"ServerGroup", serverGroup.ServerGroup[i]})
    }
    serverGroup.EntityData.Leafs = types.NewOrderedMap()
    serverGroup.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", serverGroup.ServerGroupName})
    serverGroup.EntityData.Leafs.Append("groups", types.YLeaf{"Groups", serverGroup.Groups})
    serverGroup.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", serverGroup.VrfName})
    serverGroup.EntityData.Leafs.Append("dead-time", types.YLeaf{"DeadTime", serverGroup.DeadTime})
    serverGroup.EntityData.Leafs.Append("servers", types.YLeaf{"Servers", serverGroup.Servers})

    serverGroup.EntityData.YListKeys = []string {"ServerGroupName"}

    return &(serverGroup.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup
// Server groups
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Redirected Requests. The type is interface{} with range: 0..4294967295.
    RedirectedRequests interface{}

    // Accounting data.
    Accounting Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting

    // Authentication data.
    Authentication Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication

    // Authorization data.
    Authorization Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization
}

func (serverGroup *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup) GetEntityData() *types.CommonEntityData {
    serverGroup.EntityData.YFilter = serverGroup.YFilter
    serverGroup.EntityData.YangName = "server-group"
    serverGroup.EntityData.BundleName = "cisco_ios_xr"
    serverGroup.EntityData.ParentYangName = "server-group"
    serverGroup.EntityData.SegmentPath = "server-group" + types.AddNoKeyToken(serverGroup)
    serverGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/server-groups/server-group/" + serverGroup.EntityData.SegmentPath
    serverGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroup.EntityData.Children = types.NewOrderedMap()
    serverGroup.EntityData.Children.Append("accounting", types.YChild{"Accounting", &serverGroup.Accounting})
    serverGroup.EntityData.Children.Append("authentication", types.YChild{"Authentication", &serverGroup.Authentication})
    serverGroup.EntityData.Children.Append("authorization", types.YChild{"Authorization", &serverGroup.Authorization})
    serverGroup.EntityData.Leafs = types.NewOrderedMap()
    serverGroup.EntityData.Leafs.Append("authentication-port", types.YLeaf{"AuthenticationPort", serverGroup.AuthenticationPort})
    serverGroup.EntityData.Leafs.Append("accounting-port", types.YLeaf{"AccountingPort", serverGroup.AccountingPort})
    serverGroup.EntityData.Leafs.Append("is-private", types.YLeaf{"IsPrivate", serverGroup.IsPrivate})
    serverGroup.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", serverGroup.IpAddress})
    serverGroup.EntityData.Leafs.Append("family", types.YLeaf{"Family", serverGroup.Family})
    serverGroup.EntityData.Leafs.Append("redirected-requests", types.YLeaf{"RedirectedRequests", serverGroup.RedirectedRequests})

    serverGroup.EntityData.YListKeys = []string {}

    return &(serverGroup.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting
// Accounting data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting struct {
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

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AcctTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctTransactionFailure interface{}

    // Estimated Throttled Accounting Transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctThrottledTransactions interface{}

    // Maximum Throttled Accounting Transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctMaxThrottleTrans interface{}

    // Automated Test Stats for accounting requests. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctReqs interface{}

    // Automated Test Stats for accounting timeouts. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctTimeouts interface{}

    // Automated Test Stats for accounting response. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctResponse interface{}

    // Automated Test Stats for accounting pending. The type is interface{} with
    // range: 0..4294967295.
    TotalTestAcctPending interface{}
}

func (accounting *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "server-group"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/server-groups/server-group/server-group/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("requests", types.YLeaf{"Requests", accounting.Requests})
    accounting.EntityData.Leafs.Append("pending-requests", types.YLeaf{"PendingRequests", accounting.PendingRequests})
    accounting.EntityData.Leafs.Append("retransmits", types.YLeaf{"Retransmits", accounting.Retransmits})
    accounting.EntityData.Leafs.Append("responses", types.YLeaf{"Responses", accounting.Responses})
    accounting.EntityData.Leafs.Append("timeouts", types.YLeaf{"Timeouts", accounting.Timeouts})
    accounting.EntityData.Leafs.Append("bad-responses", types.YLeaf{"BadResponses", accounting.BadResponses})
    accounting.EntityData.Leafs.Append("bad-authenticators", types.YLeaf{"BadAuthenticators", accounting.BadAuthenticators})
    accounting.EntityData.Leafs.Append("unknown-packet-types", types.YLeaf{"UnknownPacketTypes", accounting.UnknownPacketTypes})
    accounting.EntityData.Leafs.Append("dropped-responses", types.YLeaf{"DroppedResponses", accounting.DroppedResponses})
    accounting.EntityData.Leafs.Append("rtt", types.YLeaf{"Rtt", accounting.Rtt})
    accounting.EntityData.Leafs.Append("acct-unexpected-responses", types.YLeaf{"AcctUnexpectedResponses", accounting.AcctUnexpectedResponses})
    accounting.EntityData.Leafs.Append("acct-transaction-successess", types.YLeaf{"AcctTransactionSuccessess", accounting.AcctTransactionSuccessess})
    accounting.EntityData.Leafs.Append("acct-transaction-failure", types.YLeaf{"AcctTransactionFailure", accounting.AcctTransactionFailure})
    accounting.EntityData.Leafs.Append("acct-throttled-transactions", types.YLeaf{"AcctThrottledTransactions", accounting.AcctThrottledTransactions})
    accounting.EntityData.Leafs.Append("acct-max-throttle-trans", types.YLeaf{"AcctMaxThrottleTrans", accounting.AcctMaxThrottleTrans})
    accounting.EntityData.Leafs.Append("total-test-acct-reqs", types.YLeaf{"TotalTestAcctReqs", accounting.TotalTestAcctReqs})
    accounting.EntityData.Leafs.Append("total-test-acct-timeouts", types.YLeaf{"TotalTestAcctTimeouts", accounting.TotalTestAcctTimeouts})
    accounting.EntityData.Leafs.Append("total-test-acct-response", types.YLeaf{"TotalTestAcctResponse", accounting.TotalTestAcctResponse})
    accounting.EntityData.Leafs.Append("total-test-acct-pending", types.YLeaf{"TotalTestAcctPending", accounting.TotalTestAcctPending})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication
// Authentication data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication struct {
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

    // Estimated Throttled Authentication Transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthThrottledTransactions interface{}

    // Maximum Throttled Authentication Transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthMaxTransactions interface{}

    // Automated Test Stats for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthReqs interface{}

    // Automated Test Stats for authentication timeouts. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthTimeouts interface{}

    // Automated Test Stats for authentication response. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthResponse interface{}

    // Automated Test Stats for authentication pending. The type is interface{}
    // with range: 0..4294967295.
    TotalTestAuthPending interface{}
}

func (authentication *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "server-group"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/server-groups/server-group/server-group/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("access-requests", types.YLeaf{"AccessRequests", authentication.AccessRequests})
    authentication.EntityData.Leafs.Append("pending-access-requests", types.YLeaf{"PendingAccessRequests", authentication.PendingAccessRequests})
    authentication.EntityData.Leafs.Append("access-request-retransmits", types.YLeaf{"AccessRequestRetransmits", authentication.AccessRequestRetransmits})
    authentication.EntityData.Leafs.Append("access-accepts", types.YLeaf{"AccessAccepts", authentication.AccessAccepts})
    authentication.EntityData.Leafs.Append("access-rejects", types.YLeaf{"AccessRejects", authentication.AccessRejects})
    authentication.EntityData.Leafs.Append("access-challenges", types.YLeaf{"AccessChallenges", authentication.AccessChallenges})
    authentication.EntityData.Leafs.Append("access-timeouts", types.YLeaf{"AccessTimeouts", authentication.AccessTimeouts})
    authentication.EntityData.Leafs.Append("bad-access-responses", types.YLeaf{"BadAccessResponses", authentication.BadAccessResponses})
    authentication.EntityData.Leafs.Append("bad-access-authenticators", types.YLeaf{"BadAccessAuthenticators", authentication.BadAccessAuthenticators})
    authentication.EntityData.Leafs.Append("unknown-access-types", types.YLeaf{"UnknownAccessTypes", authentication.UnknownAccessTypes})
    authentication.EntityData.Leafs.Append("dropped-access-responses", types.YLeaf{"DroppedAccessResponses", authentication.DroppedAccessResponses})
    authentication.EntityData.Leafs.Append("rtt", types.YLeaf{"Rtt", authentication.Rtt})
    authentication.EntityData.Leafs.Append("authen-transaction-successess", types.YLeaf{"AuthenTransactionSuccessess", authentication.AuthenTransactionSuccessess})
    authentication.EntityData.Leafs.Append("authen-transaction-failure", types.YLeaf{"AuthenTransactionFailure", authentication.AuthenTransactionFailure})
    authentication.EntityData.Leafs.Append("authen-unexpected-responses", types.YLeaf{"AuthenUnexpectedResponses", authentication.AuthenUnexpectedResponses})
    authentication.EntityData.Leafs.Append("authen-server-error-responses", types.YLeaf{"AuthenServerErrorResponses", authentication.AuthenServerErrorResponses})
    authentication.EntityData.Leafs.Append("authen-incorrect-responses", types.YLeaf{"AuthenIncorrectResponses", authentication.AuthenIncorrectResponses})
    authentication.EntityData.Leafs.Append("auth-throttled-transactions", types.YLeaf{"AuthThrottledTransactions", authentication.AuthThrottledTransactions})
    authentication.EntityData.Leafs.Append("auth-max-transactions", types.YLeaf{"AuthMaxTransactions", authentication.AuthMaxTransactions})
    authentication.EntityData.Leafs.Append("total-test-auth-reqs", types.YLeaf{"TotalTestAuthReqs", authentication.TotalTestAuthReqs})
    authentication.EntityData.Leafs.Append("total-test-auth-timeouts", types.YLeaf{"TotalTestAuthTimeouts", authentication.TotalTestAuthTimeouts})
    authentication.EntityData.Leafs.Append("total-test-auth-response", types.YLeaf{"TotalTestAuthResponse", authentication.TotalTestAuthResponse})
    authentication.EntityData.Leafs.Append("total-test-auth-pending", types.YLeaf{"TotalTestAuthPending", authentication.TotalTestAuthPending})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization
// Authorization data
type Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization struct {
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

func (authorization *Radius_Nodes_Node_ServerGroups_ServerGroup_ServerGroup_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "server-group"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/server-groups/server-group/server-group/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("author-requests", types.YLeaf{"AuthorRequests", authorization.AuthorRequests})
    authorization.EntityData.Leafs.Append("author-request-timeouts", types.YLeaf{"AuthorRequestTimeouts", authorization.AuthorRequestTimeouts})
    authorization.EntityData.Leafs.Append("author-unexpected-responses", types.YLeaf{"AuthorUnexpectedResponses", authorization.AuthorUnexpectedResponses})
    authorization.EntityData.Leafs.Append("author-server-error-responses", types.YLeaf{"AuthorServerErrorResponses", authorization.AuthorServerErrorResponses})
    authorization.EntityData.Leafs.Append("author-incorrect-responses", types.YLeaf{"AuthorIncorrectResponses", authorization.AuthorIncorrectResponses})
    authorization.EntityData.Leafs.Append("author-response-time", types.YLeaf{"AuthorResponseTime", authorization.AuthorResponseTime})
    authorization.EntityData.Leafs.Append("author-transaction-successess", types.YLeaf{"AuthorTransactionSuccessess", authorization.AuthorTransactionSuccessess})
    authorization.EntityData.Leafs.Append("author-transaction-failure", types.YLeaf{"AuthorTransactionFailure", authorization.AuthorTransactionFailure})

    authorization.EntityData.YListKeys = []string {}

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

    // Radius context not found. The type is interface{} with range:
    // 0..4294967295.
    RadiusContextNotFound interface{}

    // Client context not found. The type is interface{} with range:
    // 0..4294967295.
    ClientContextNotFound interface{}
}

func (dynamicAuthorization *Radius_Nodes_Node_DynamicAuthorization) GetEntityData() *types.CommonEntityData {
    dynamicAuthorization.EntityData.YFilter = dynamicAuthorization.YFilter
    dynamicAuthorization.EntityData.YangName = "dynamic-authorization"
    dynamicAuthorization.EntityData.BundleName = "cisco_ios_xr"
    dynamicAuthorization.EntityData.ParentYangName = "node"
    dynamicAuthorization.EntityData.SegmentPath = "dynamic-authorization"
    dynamicAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius/nodes/node/" + dynamicAuthorization.EntityData.SegmentPath
    dynamicAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicAuthorization.EntityData.Children = types.NewOrderedMap()
    dynamicAuthorization.EntityData.Leafs = types.NewOrderedMap()
    dynamicAuthorization.EntityData.Leafs.Append("disconnected-invalid-requests", types.YLeaf{"DisconnectedInvalidRequests", dynamicAuthorization.DisconnectedInvalidRequests})
    dynamicAuthorization.EntityData.Leafs.Append("invalid-coa-requests", types.YLeaf{"InvalidCoaRequests", dynamicAuthorization.InvalidCoaRequests})
    dynamicAuthorization.EntityData.Leafs.Append("radius-context-not-found", types.YLeaf{"RadiusContextNotFound", dynamicAuthorization.RadiusContextNotFound})
    dynamicAuthorization.EntityData.Leafs.Append("client-context-not-found", types.YLeaf{"ClientContextNotFound", dynamicAuthorization.ClientContextNotFound})

    dynamicAuthorization.EntityData.YListKeys = []string {}

    return &(dynamicAuthorization.EntityData)
}

