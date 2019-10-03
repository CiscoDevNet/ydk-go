// This module contains a collection of YANG definitions
// for Cisco IOS-XR iedge4710 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber: Subscriber operational data
//   iedge-license-manager: iedge license manager
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package iedge4710_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iedge4710_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-oper subscriber}", reflect.TypeOf(Subscriber{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-oper:subscriber", reflect.TypeOf(Subscriber{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-oper iedge-license-manager}", reflect.TypeOf(IedgeLicenseManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-oper:iedge-license-manager", reflect.TypeOf(IedgeLicenseManager{}))
}

// SubscriberAuthorStateFilterFlag represents Subscriber author state filter flag
type SubscriberAuthorStateFilterFlag string

const (
    // UnAuthorized
    SubscriberAuthorStateFilterFlag_un_authorized SubscriberAuthorStateFilterFlag = "un-authorized"

    // Authorized
    SubscriberAuthorStateFilterFlag_authorized SubscriberAuthorStateFilterFlag = "authorized"
)

// IedgeOperSessionState represents Subscriber session states
type IedgeOperSessionState string

const (
    // Initialize
    IedgeOperSessionState_initialize IedgeOperSessionState = "initialize"

    // Connecting
    IedgeOperSessionState_connecting IedgeOperSessionState = "connecting"

    // Connected
    IedgeOperSessionState_connected IedgeOperSessionState = "connected"

    // Activated
    IedgeOperSessionState_activated IedgeOperSessionState = "activated"

    // Idle
    IedgeOperSessionState_idle IedgeOperSessionState = "idle"

    // Disconnecting
    IedgeOperSessionState_disconnecting IedgeOperSessionState = "disconnecting"

    // End
    IedgeOperSessionState_end IedgeOperSessionState = "end"
)

// SubscriberStateFilterFlag represents Subscriber state filter flag
type SubscriberStateFilterFlag string

const (
    // Initializing
    SubscriberStateFilterFlag_initializing SubscriberStateFilterFlag = "initializing"

    // Connecting
    SubscriberStateFilterFlag_connecting SubscriberStateFilterFlag = "connecting"

    // Connected
    SubscriberStateFilterFlag_connected SubscriberStateFilterFlag = "connected"

    // Activated
    SubscriberStateFilterFlag_activated SubscriberStateFilterFlag = "activated"

    // Idle
    SubscriberStateFilterFlag_idle SubscriberStateFilterFlag = "idle"

    // Disconnecting
    SubscriberStateFilterFlag_disconnecting SubscriberStateFilterFlag = "disconnecting"

    // End
    SubscriberStateFilterFlag_end SubscriberStateFilterFlag = "end"
)

// IedgePppSub represents PPPoE sub types
type IedgePppSub string

const (
    // PPP termination and aggregation
    IedgePppSub_pta IedgePppSub = "pta"

    // L2TP access controller
    IedgePppSub_lac IedgePppSub = "lac"
)

// IedgeOperServiceStatus represents Subscriber service status
type IedgeOperServiceStatus string

const (
    // Unknown
    IedgeOperServiceStatus_unknown IedgeOperServiceStatus = "unknown"

    // Error
    IedgeOperServiceStatus_error_ IedgeOperServiceStatus = "error"

    // New
    IedgeOperServiceStatus_new_ IedgeOperServiceStatus = "new"

    // Loading
    IedgeOperServiceStatus_loading IedgeOperServiceStatus = "loading"

    // Request Association
    IedgeOperServiceStatus_req_association IedgeOperServiceStatus = "req-association"

    // Associated
    IedgeOperServiceStatus_associated IedgeOperServiceStatus = "associated"

    // Request PD Association
    IedgeOperServiceStatus_req_pd_association IedgeOperServiceStatus = "req-pd-association"

    // Applied
    IedgeOperServiceStatus_applied IedgeOperServiceStatus = "applied"

    // Request Unassociation
    IedgeOperServiceStatus_req_unassociation IedgeOperServiceStatus = "req-unassociation"

    // Request PD Unassociation
    IedgeOperServiceStatus_req_pd_unassociation IedgeOperServiceStatus = "req-pd-unassociation"

    // Unapplied
    IedgeOperServiceStatus_unapplied IedgeOperServiceStatus = "unapplied"

    // Max
    IedgeOperServiceStatus_max IedgeOperServiceStatus = "max"
)

// SubscriberAuthenStateFilterFlag represents Subscriber authen state filter flag
type SubscriberAuthenStateFilterFlag string

const (
    // UnAuthenticated
    SubscriberAuthenStateFilterFlag_un_authenticated SubscriberAuthenStateFilterFlag = "un-authenticated"

    // Authenticated
    SubscriberAuthenStateFilterFlag_authenticated SubscriberAuthenStateFilterFlag = "authenticated"
)

// IedgeOperService represents Service types
type IedgeOperService string

const (
    // Unknown Service
    IedgeOperService_unknown IedgeOperService = "unknown"

    // Profile
    IedgeOperService_profile IedgeOperService = "profile"

    // Template
    IedgeOperService_template IedgeOperService = "template"

    // GRP Template
    IedgeOperService_grp_template IedgeOperService = "grp-template"

    // PPP Template
    IedgeOperService_ppp_template IedgeOperService = "ppp-template"

    // ETH Template
    IedgeOperService_eth_template IedgeOperService = "eth-template"

    // IPSub Template
    IedgeOperService_ip_sub_template IedgeOperService = "ip-sub-template"

    // Multi Template
    IedgeOperService_multi_template IedgeOperService = "multi-template"

    // MAX Template
    IedgeOperService_max_templae IedgeOperService = "max-templae"
)

// IedgeOperSession represents Subscriber session types
type IedgeOperSession string

const (
    // Unknown
    IedgeOperSession_unknown IedgeOperSession = "unknown"

    // PPPoE/PPP client
    IedgeOperSession_pppoe IedgeOperSession = "pppoe"

    // PPP serial client
    IedgeOperSession_ppp IedgeOperSession = "ppp"

    // IP subscriber - packet trigger
    IedgeOperSession_ip_packet_trigger IedgeOperSession = "ip-packet-trigger"

    // IP subscriber - DHCP trigger
    IedgeOperSession_ip_packet_dhcp_trigger IedgeOperSession = "ip-packet-dhcp-trigger"
)

// SubscriberSrgOperFilterFlag represents Subscriber srg oper filter flag
type SubscriberSrgOperFilterFlag string

const (
    // SRG Role None
    SubscriberSrgOperFilterFlag_srg_none SubscriberSrgOperFilterFlag = "srg-none"

    // SRG Role Master
    SubscriberSrgOperFilterFlag_srg_master SubscriberSrgOperFilterFlag = "srg-master"

    // SRG Role Slave
    SubscriberSrgOperFilterFlag_srg_slave SubscriberSrgOperFilterFlag = "srg-slave"

    // SRG Role Master Slave
    SubscriberSrgOperFilterFlag_srg_both SubscriberSrgOperFilterFlag = "srg-both"
)

// SubscriberAddressFamilyFilterFlag represents Subscriber address family filter flag
type SubscriberAddressFamilyFilterFlag string

const (
    // IPv4 only
    SubscriberAddressFamilyFilterFlag_ipv4_only SubscriberAddressFamilyFilterFlag = "ipv4-only"

    // IPv6 only
    SubscriberAddressFamilyFilterFlag_ipv6_only SubscriberAddressFamilyFilterFlag = "ipv6-only"

    // IPv4 all
    SubscriberAddressFamilyFilterFlag_ipv4_all SubscriberAddressFamilyFilterFlag = "ipv4-all"

    // IPv6 all
    SubscriberAddressFamilyFilterFlag_ipv6_all SubscriberAddressFamilyFilterFlag = "ipv6-all"

    // Dual all
    SubscriberAddressFamilyFilterFlag_dual_all SubscriberAddressFamilyFilterFlag = "dual-all"

    // Dual part up
    SubscriberAddressFamilyFilterFlag_dual_part_up SubscriberAddressFamilyFilterFlag = "dual-part-up"

    // Dual up
    SubscriberAddressFamilyFilterFlag_dual_up SubscriberAddressFamilyFilterFlag = "dual-up"

    // LAC
    SubscriberAddressFamilyFilterFlag_lac SubscriberAddressFamilyFilterFlag = "lac"
)

// Subscriber
// Subscriber operational data
type Subscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager operational data.
    Manager Subscriber_Manager

    // Subscriber session operational data.
    Session Subscriber_Session
}

func (subscriber *Subscriber) GetEntityData() *types.CommonEntityData {
    subscriber.EntityData.YFilter = subscriber.YFilter
    subscriber.EntityData.YangName = "subscriber"
    subscriber.EntityData.BundleName = "cisco_ios_xr"
    subscriber.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-oper"
    subscriber.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-oper:subscriber"
    subscriber.EntityData.AbsolutePath = subscriber.EntityData.SegmentPath
    subscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriber.EntityData.Children = types.NewOrderedMap()
    subscriber.EntityData.Children.Append("manager", types.YChild{"Manager", &subscriber.Manager})
    subscriber.EntityData.Children.Append("session", types.YChild{"Session", &subscriber.Session})
    subscriber.EntityData.Leafs = types.NewOrderedMap()

    subscriber.EntityData.YListKeys = []string {}

    return &(subscriber.EntityData)
}

// Subscriber_Manager
// Subscriber manager operational data
type Subscriber_Manager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager list of nodes.
    Nodes Subscriber_Manager_Nodes
}

func (manager *Subscriber_Manager) GetEntityData() *types.CommonEntityData {
    manager.EntityData.YFilter = manager.YFilter
    manager.EntityData.YangName = "manager"
    manager.EntityData.BundleName = "cisco_ios_xr"
    manager.EntityData.ParentYangName = "subscriber"
    manager.EntityData.SegmentPath = "manager"
    manager.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/" + manager.EntityData.SegmentPath
    manager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manager.EntityData.Children = types.NewOrderedMap()
    manager.EntityData.Children.Append("nodes", types.YChild{"Nodes", &manager.Nodes})
    manager.EntityData.Leafs = types.NewOrderedMap()

    manager.EntityData.YListKeys = []string {}

    return &(manager.EntityData)
}

// Subscriber_Manager_Nodes
// Subscriber manager list of nodes
type Subscriber_Manager_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager operational data for a particular node. The type is
    // slice of Subscriber_Manager_Nodes_Node.
    Node []*Subscriber_Manager_Nodes_Node
}

func (nodes *Subscriber_Manager_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "manager"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/" + nodes.EntityData.SegmentPath
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

// Subscriber_Manager_Nodes_Node
// Subscriber manager operational data for a
// particular node
type Subscriber_Manager_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Subscriber manager statistics.
    Statistics Subscriber_Manager_Nodes_Node_Statistics
}

func (node *Subscriber_Manager_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics
// Subscriber manager statistics
type Subscriber_Manager_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA statistics.
    Aaa Subscriber_Manager_Nodes_Node_Statistics_Aaa

    // Aggregate summary of statistics.
    AggregateSummary Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary

    // Disconnect Unique Summary statistics.
    DisconnUnique Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique

    // Geo Redundancy statistics.
    Srg Subscriber_Manager_Nodes_Node_Statistics_Srg
}

func (statistics *Subscriber_Manager_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("aaa", types.YChild{"Aaa", &statistics.Aaa})
    statistics.EntityData.Children.Append("aggregate-summary", types.YChild{"AggregateSummary", &statistics.AggregateSummary})
    statistics.EntityData.Children.Append("disconn-unique", types.YChild{"DisconnUnique", &statistics.DisconnUnique})
    statistics.EntityData.Children.Append("srg", types.YChild{"Srg", &statistics.Srg})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa
// AAA statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aggregate accounting statistics.
    AggregateAccounting Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting

    // Authentication statistics.
    Authentication Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication

    // Aggregate mobility statistics.
    AggregateMobility Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility

    // Aggregate authentication statistics.
    AggregateAuthentication Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication

    // Display all subscriber management statistics.
    AccountingStatsAll Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll

    // Change of authorization (COA) statistics.
    ChangeOfAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization

    // Authorization statistics.
    Authorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization

    // Aggregate authorization statistics.
    AggregateAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization

    // Display all subscriber management total statistics.
    AggregateAccountingStatsAll Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll

    // Accounting statistics.
    Accounting Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting

    // Mobility statistics.
    Mobility Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility

    // Aggregate change of authorization (COA) statistics.
    AggregateChangeOfAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization
}

func (aaa *Subscriber_Manager_Nodes_Node_Statistics_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "statistics"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/" + aaa.EntityData.SegmentPath
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("aggregate-accounting", types.YChild{"AggregateAccounting", &aaa.AggregateAccounting})
    aaa.EntityData.Children.Append("authentication", types.YChild{"Authentication", &aaa.Authentication})
    aaa.EntityData.Children.Append("aggregate-mobility", types.YChild{"AggregateMobility", &aaa.AggregateMobility})
    aaa.EntityData.Children.Append("aggregate-authentication", types.YChild{"AggregateAuthentication", &aaa.AggregateAuthentication})
    aaa.EntityData.Children.Append("accounting-stats-all", types.YChild{"AccountingStatsAll", &aaa.AccountingStatsAll})
    aaa.EntityData.Children.Append("change-of-authorization", types.YChild{"ChangeOfAuthorization", &aaa.ChangeOfAuthorization})
    aaa.EntityData.Children.Append("authorization", types.YChild{"Authorization", &aaa.Authorization})
    aaa.EntityData.Children.Append("aggregate-authorization", types.YChild{"AggregateAuthorization", &aaa.AggregateAuthorization})
    aaa.EntityData.Children.Append("aggregate-accounting-stats-all", types.YChild{"AggregateAccountingStatsAll", &aaa.AggregateAccountingStatsAll})
    aaa.EntityData.Children.Append("accounting", types.YChild{"Accounting", &aaa.Accounting})
    aaa.EntityData.Children.Append("mobility", types.YChild{"Mobility", &aaa.Mobility})
    aaa.EntityData.Children.Append("aggregate-change-of-authorization", types.YChild{"AggregateChangeOfAuthorization", &aaa.AggregateChangeOfAuthorization})
    aaa.EntityData.Leafs = types.NewOrderedMap()

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting
// Aggregate accounting statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight
}

func (aggregateAccounting *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting) GetEntityData() *types.CommonEntityData {
    aggregateAccounting.EntityData.YFilter = aggregateAccounting.YFilter
    aggregateAccounting.EntityData.YangName = "aggregate-accounting"
    aggregateAccounting.EntityData.BundleName = "cisco_ios_xr"
    aggregateAccounting.EntityData.ParentYangName = "aaa"
    aggregateAccounting.EntityData.SegmentPath = "aggregate-accounting"
    aggregateAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateAccounting.EntityData.SegmentPath
    aggregateAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAccounting.EntityData.Children = types.NewOrderedMap()
    aggregateAccounting.EntityData.Children.Append("start", types.YChild{"Start", &aggregateAccounting.Start})
    aggregateAccounting.EntityData.Children.Append("stop", types.YChild{"Stop", &aggregateAccounting.Stop})
    aggregateAccounting.EntityData.Children.Append("interim", types.YChild{"Interim", &aggregateAccounting.Interim})
    aggregateAccounting.EntityData.Children.Append("pass-through", types.YChild{"PassThrough", &aggregateAccounting.PassThrough})
    aggregateAccounting.EntityData.Children.Append("update", types.YChild{"Update", &aggregateAccounting.Update})
    aggregateAccounting.EntityData.Children.Append("interim-inflight", types.YChild{"InterimInflight", &aggregateAccounting.InterimInflight})
    aggregateAccounting.EntityData.Leafs = types.NewOrderedMap()
    aggregateAccounting.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", aggregateAccounting.ActiveSessions})
    aggregateAccounting.EntityData.Leafs.Append("started-sessions", types.YLeaf{"StartedSessions", aggregateAccounting.StartedSessions})
    aggregateAccounting.EntityData.Leafs.Append("stopped-sessions", types.YLeaf{"StoppedSessions", aggregateAccounting.StoppedSessions})
    aggregateAccounting.EntityData.Leafs.Append("policy-plane-errored-requests", types.YLeaf{"PolicyPlaneErroredRequests", aggregateAccounting.PolicyPlaneErroredRequests})
    aggregateAccounting.EntityData.Leafs.Append("policy-plane-unknown-requests", types.YLeaf{"PolicyPlaneUnknownRequests", aggregateAccounting.PolicyPlaneUnknownRequests})

    aggregateAccounting.EntityData.YListKeys = []string {}

    return &(aggregateAccounting.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "aggregate-accounting"
    start.EntityData.SegmentPath = "start"
    start.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + start.EntityData.SegmentPath
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = types.NewOrderedMap()
    start.EntityData.Leafs = types.NewOrderedMap()
    start.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", start.ReceivedRequests})
    start.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", start.ErroredRequests})
    start.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests})
    start.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", start.AaaSentRequests})
    start.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses})
    start.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses})

    start.EntityData.YListKeys = []string {}

    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "aggregate-accounting"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Leafs = types.NewOrderedMap()
    stop.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", stop.ReceivedRequests})
    stop.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", stop.ErroredRequests})
    stop.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests})
    stop.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", stop.AaaSentRequests})
    stop.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses})
    stop.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses})

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "aggregate-accounting"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + interim.EntityData.SegmentPath
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = types.NewOrderedMap()
    interim.EntityData.Leafs = types.NewOrderedMap()
    interim.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", interim.ReceivedRequests})
    interim.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", interim.ErroredRequests})
    interim.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests})
    interim.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", interim.AaaSentRequests})
    interim.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses})
    interim.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses})

    interim.EntityData.YListKeys = []string {}

    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "aggregate-accounting"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + passThrough.EntityData.SegmentPath
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = types.NewOrderedMap()
    passThrough.EntityData.Leafs = types.NewOrderedMap()
    passThrough.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests})
    passThrough.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", passThrough.ErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests})
    passThrough.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses})
    passThrough.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses})

    passThrough.EntityData.YListKeys = []string {}

    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "aggregate-accounting"
    update.EntityData.SegmentPath = "update"
    update.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + update.EntityData.SegmentPath
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = types.NewOrderedMap()
    update.EntityData.Leafs = types.NewOrderedMap()
    update.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", update.ReceivedRequests})
    update.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", update.ErroredRequests})
    update.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests})
    update.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", update.AaaSentRequests})
    update.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses})
    update.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses})

    update.EntityData.YListKeys = []string {}

    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "aggregate-accounting"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting/" + interimInflight.EntityData.SegmentPath
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = types.NewOrderedMap()
    interimInflight.EntityData.Leafs = types.NewOrderedMap()
    interimInflight.EntityData.Leafs.Append("quota-exhausts", types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts})
    interimInflight.EntityData.Leafs.Append("denied-requests", types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests})
    interimInflight.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests})
    interimInflight.EntityData.Leafs.Append("total-quota-of-requests", types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("remaining-quota-of-requests", types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("low-water-mark-quota-of-requests", types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests})

    interimInflight.EntityData.YListKeys = []string {}

    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication
// Authentication statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authentication *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "aaa"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authentication.SentRequests})
    authentication.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authentication.AcceptedRequests})
    authentication.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authentication.SuccessfulRequests})
    authentication.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authentication.RejectedRequests})
    authentication.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authentication.UnreachableRequests})
    authentication.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authentication.ErroredRequests})
    authentication.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authentication.TerminatedRequests})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility
// Aggregate mobility statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (aggregateMobility *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility) GetEntityData() *types.CommonEntityData {
    aggregateMobility.EntityData.YFilter = aggregateMobility.YFilter
    aggregateMobility.EntityData.YangName = "aggregate-mobility"
    aggregateMobility.EntityData.BundleName = "cisco_ios_xr"
    aggregateMobility.EntityData.ParentYangName = "aaa"
    aggregateMobility.EntityData.SegmentPath = "aggregate-mobility"
    aggregateMobility.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateMobility.EntityData.SegmentPath
    aggregateMobility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateMobility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateMobility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateMobility.EntityData.Children = types.NewOrderedMap()
    aggregateMobility.EntityData.Leafs = types.NewOrderedMap()
    aggregateMobility.EntityData.Leafs.Append("send-request-successes", types.YLeaf{"SendRequestSuccesses", aggregateMobility.SendRequestSuccesses})
    aggregateMobility.EntityData.Leafs.Append("send-request-failures", types.YLeaf{"SendRequestFailures", aggregateMobility.SendRequestFailures})
    aggregateMobility.EntityData.Leafs.Append("receive-response-successes", types.YLeaf{"ReceiveResponseSuccesses", aggregateMobility.ReceiveResponseSuccesses})
    aggregateMobility.EntityData.Leafs.Append("receive-response-failures", types.YLeaf{"ReceiveResponseFailures", aggregateMobility.ReceiveResponseFailures})

    aggregateMobility.EntityData.YListKeys = []string {}

    return &(aggregateMobility.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication
// Aggregate authentication statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (aggregateAuthentication *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication) GetEntityData() *types.CommonEntityData {
    aggregateAuthentication.EntityData.YFilter = aggregateAuthentication.YFilter
    aggregateAuthentication.EntityData.YangName = "aggregate-authentication"
    aggregateAuthentication.EntityData.BundleName = "cisco_ios_xr"
    aggregateAuthentication.EntityData.ParentYangName = "aaa"
    aggregateAuthentication.EntityData.SegmentPath = "aggregate-authentication"
    aggregateAuthentication.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateAuthentication.EntityData.SegmentPath
    aggregateAuthentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAuthentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAuthentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAuthentication.EntityData.Children = types.NewOrderedMap()
    aggregateAuthentication.EntityData.Leafs = types.NewOrderedMap()
    aggregateAuthentication.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", aggregateAuthentication.SentRequests})
    aggregateAuthentication.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", aggregateAuthentication.AcceptedRequests})
    aggregateAuthentication.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", aggregateAuthentication.SuccessfulRequests})
    aggregateAuthentication.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", aggregateAuthentication.RejectedRequests})
    aggregateAuthentication.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", aggregateAuthentication.UnreachableRequests})
    aggregateAuthentication.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", aggregateAuthentication.ErroredRequests})
    aggregateAuthentication.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", aggregateAuthentication.TerminatedRequests})

    aggregateAuthentication.EntityData.YListKeys = []string {}

    return &(aggregateAuthentication.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll
// Display all subscriber management
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of stats for accounting.
    AccountingStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics

    // List of stats for authentication.
    AuthenticationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics

    // List of stats for authorization.
    AuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics

    // List of stats for COA.
    ChangeOfAuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics

    // List of stats for Mobility.
    MobilityStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics
}

func (accountingStatsAll *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll) GetEntityData() *types.CommonEntityData {
    accountingStatsAll.EntityData.YFilter = accountingStatsAll.YFilter
    accountingStatsAll.EntityData.YangName = "accounting-stats-all"
    accountingStatsAll.EntityData.BundleName = "cisco_ios_xr"
    accountingStatsAll.EntityData.ParentYangName = "aaa"
    accountingStatsAll.EntityData.SegmentPath = "accounting-stats-all"
    accountingStatsAll.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + accountingStatsAll.EntityData.SegmentPath
    accountingStatsAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatsAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatsAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatsAll.EntityData.Children = types.NewOrderedMap()
    accountingStatsAll.EntityData.Children.Append("accounting-statistics", types.YChild{"AccountingStatistics", &accountingStatsAll.AccountingStatistics})
    accountingStatsAll.EntityData.Children.Append("authentication-statistics", types.YChild{"AuthenticationStatistics", &accountingStatsAll.AuthenticationStatistics})
    accountingStatsAll.EntityData.Children.Append("authorization-statistics", types.YChild{"AuthorizationStatistics", &accountingStatsAll.AuthorizationStatistics})
    accountingStatsAll.EntityData.Children.Append("change-of-authorization-statistics", types.YChild{"ChangeOfAuthorizationStatistics", &accountingStatsAll.ChangeOfAuthorizationStatistics})
    accountingStatsAll.EntityData.Children.Append("mobility-statistics", types.YChild{"MobilityStatistics", &accountingStatsAll.MobilityStatistics})
    accountingStatsAll.EntityData.Leafs = types.NewOrderedMap()

    accountingStatsAll.EntityData.YListKeys = []string {}

    return &(accountingStatsAll.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics
// List of stats for accounting
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight
}

func (accountingStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "accounting-stats-all"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/" + accountingStatistics.EntityData.SegmentPath
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = types.NewOrderedMap()
    accountingStatistics.EntityData.Children.Append("start", types.YChild{"Start", &accountingStatistics.Start})
    accountingStatistics.EntityData.Children.Append("stop", types.YChild{"Stop", &accountingStatistics.Stop})
    accountingStatistics.EntityData.Children.Append("interim", types.YChild{"Interim", &accountingStatistics.Interim})
    accountingStatistics.EntityData.Children.Append("pass-through", types.YChild{"PassThrough", &accountingStatistics.PassThrough})
    accountingStatistics.EntityData.Children.Append("update", types.YChild{"Update", &accountingStatistics.Update})
    accountingStatistics.EntityData.Children.Append("interim-inflight", types.YChild{"InterimInflight", &accountingStatistics.InterimInflight})
    accountingStatistics.EntityData.Leafs = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", accountingStatistics.ActiveSessions})
    accountingStatistics.EntityData.Leafs.Append("started-sessions", types.YLeaf{"StartedSessions", accountingStatistics.StartedSessions})
    accountingStatistics.EntityData.Leafs.Append("stopped-sessions", types.YLeaf{"StoppedSessions", accountingStatistics.StoppedSessions})
    accountingStatistics.EntityData.Leafs.Append("policy-plane-errored-requests", types.YLeaf{"PolicyPlaneErroredRequests", accountingStatistics.PolicyPlaneErroredRequests})
    accountingStatistics.EntityData.Leafs.Append("policy-plane-unknown-requests", types.YLeaf{"PolicyPlaneUnknownRequests", accountingStatistics.PolicyPlaneUnknownRequests})

    accountingStatistics.EntityData.YListKeys = []string {}

    return &(accountingStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting-statistics"
    start.EntityData.SegmentPath = "start"
    start.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + start.EntityData.SegmentPath
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = types.NewOrderedMap()
    start.EntityData.Leafs = types.NewOrderedMap()
    start.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", start.ReceivedRequests})
    start.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", start.ErroredRequests})
    start.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests})
    start.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", start.AaaSentRequests})
    start.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses})
    start.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses})

    start.EntityData.YListKeys = []string {}

    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting-statistics"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Leafs = types.NewOrderedMap()
    stop.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", stop.ReceivedRequests})
    stop.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", stop.ErroredRequests})
    stop.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests})
    stop.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", stop.AaaSentRequests})
    stop.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses})
    stop.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses})

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting-statistics"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + interim.EntityData.SegmentPath
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = types.NewOrderedMap()
    interim.EntityData.Leafs = types.NewOrderedMap()
    interim.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", interim.ReceivedRequests})
    interim.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", interim.ErroredRequests})
    interim.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests})
    interim.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", interim.AaaSentRequests})
    interim.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses})
    interim.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses})

    interim.EntityData.YListKeys = []string {}

    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting-statistics"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + passThrough.EntityData.SegmentPath
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = types.NewOrderedMap()
    passThrough.EntityData.Leafs = types.NewOrderedMap()
    passThrough.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests})
    passThrough.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", passThrough.ErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests})
    passThrough.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses})
    passThrough.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses})

    passThrough.EntityData.YListKeys = []string {}

    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting-statistics"
    update.EntityData.SegmentPath = "update"
    update.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + update.EntityData.SegmentPath
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = types.NewOrderedMap()
    update.EntityData.Leafs = types.NewOrderedMap()
    update.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", update.ReceivedRequests})
    update.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", update.ErroredRequests})
    update.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests})
    update.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", update.AaaSentRequests})
    update.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses})
    update.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses})

    update.EntityData.YListKeys = []string {}

    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting-statistics"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/accounting-statistics/" + interimInflight.EntityData.SegmentPath
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = types.NewOrderedMap()
    interimInflight.EntityData.Leafs = types.NewOrderedMap()
    interimInflight.EntityData.Leafs.Append("quota-exhausts", types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts})
    interimInflight.EntityData.Leafs.Append("denied-requests", types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests})
    interimInflight.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests})
    interimInflight.EntityData.Leafs.Append("total-quota-of-requests", types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("remaining-quota-of-requests", types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("low-water-mark-quota-of-requests", types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests})

    interimInflight.EntityData.YListKeys = []string {}

    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics
// List of stats for authentication
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authenticationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/" + authenticationStatistics.EntityData.SegmentPath
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authenticationStatistics.SentRequests})
    authenticationStatistics.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authenticationStatistics.AcceptedRequests})
    authenticationStatistics.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authenticationStatistics.SuccessfulRequests})
    authenticationStatistics.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authenticationStatistics.RejectedRequests})
    authenticationStatistics.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authenticationStatistics.UnreachableRequests})
    authenticationStatistics.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authenticationStatistics.ErroredRequests})
    authenticationStatistics.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authenticationStatistics.TerminatedRequests})

    authenticationStatistics.EntityData.YListKeys = []string {}

    return &(authenticationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics
// List of stats for authorization
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    authorizationStatistics.EntityData.YFilter = authorizationStatistics.YFilter
    authorizationStatistics.EntityData.YangName = "authorization-statistics"
    authorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authorizationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    authorizationStatistics.EntityData.SegmentPath = "authorization-statistics"
    authorizationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/" + authorizationStatistics.EntityData.SegmentPath
    authorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizationStatistics.EntityData.Children = types.NewOrderedMap()
    authorizationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authorizationStatistics.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authorizationStatistics.SentRequests})
    authorizationStatistics.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authorizationStatistics.AcceptedRequests})
    authorizationStatistics.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authorizationStatistics.SuccessfulRequests})
    authorizationStatistics.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authorizationStatistics.RejectedRequests})
    authorizationStatistics.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authorizationStatistics.UnreachableRequests})
    authorizationStatistics.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authorizationStatistics.ErroredRequests})
    authorizationStatistics.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authorizationStatistics.TerminatedRequests})

    authorizationStatistics.EntityData.YListKeys = []string {}

    return &(authorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics
// List of stats for COA
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
}

func (changeOfAuthorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    changeOfAuthorizationStatistics.EntityData.YFilter = changeOfAuthorizationStatistics.YFilter
    changeOfAuthorizationStatistics.EntityData.YangName = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorizationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    changeOfAuthorizationStatistics.EntityData.SegmentPath = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/" + changeOfAuthorizationStatistics.EntityData.SegmentPath
    changeOfAuthorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorizationStatistics.EntityData.Children = types.NewOrderedMap()
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-logon", types.YChild{"AccountLogon", &changeOfAuthorizationStatistics.AccountLogon})
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-logoff", types.YChild{"AccountLogoff", &changeOfAuthorizationStatistics.AccountLogoff})
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-update", types.YChild{"AccountUpdate", &changeOfAuthorizationStatistics.AccountUpdate})
    changeOfAuthorizationStatistics.EntityData.Children.Append("session-disconnect", types.YChild{"SessionDisconnect", &changeOfAuthorizationStatistics.SessionDisconnect})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-logon", types.YChild{"SingleServiceLogon", &changeOfAuthorizationStatistics.SingleServiceLogon})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-logoff", types.YChild{"SingleServiceLogoff", &changeOfAuthorizationStatistics.SingleServiceLogoff})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-modify", types.YChild{"SingleServiceModify", &changeOfAuthorizationStatistics.SingleServiceModify})
    changeOfAuthorizationStatistics.EntityData.Children.Append("service-multi", types.YChild{"ServiceMulti", &changeOfAuthorizationStatistics.ServiceMulti})
    changeOfAuthorizationStatistics.EntityData.Leafs = types.NewOrderedMap()
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-account-cmd-resps", types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorizationStatistics.UnknownAccountCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-service-cmd-resps", types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorizationStatistics.UnknownServiceCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-cmd-resps", types.YLeaf{"UnknownCmdResps", changeOfAuthorizationStatistics.UnknownCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("attr-list-retrieve-failure-resps", types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorizationStatistics.AttrListRetrieveFailureResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("resp-send-failure", types.YLeaf{"RespSendFailure", changeOfAuthorizationStatistics.RespSendFailure})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("internal-err-resps", types.YLeaf{"InternalErrResps", changeOfAuthorizationStatistics.InternalErrResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("service-profile-push-failure-resps", types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorizationStatistics.ServiceProfilePushFailureResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-cmd-resps", types.YLeaf{"NoCmdResps", changeOfAuthorizationStatistics.NoCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-session-found-resps", types.YLeaf{"NoSessionFoundResps", changeOfAuthorizationStatistics.NoSessionFoundResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-session-peer-resps", types.YLeaf{"NoSessionPeerResps", changeOfAuthorizationStatistics.NoSessionPeerResps})

    changeOfAuthorizationStatistics.EntityData.YListKeys = []string {}

    return &(changeOfAuthorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + accountLogon.EntityData.SegmentPath
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = types.NewOrderedMap()
    accountLogon.EntityData.Leafs = types.NewOrderedMap()
    accountLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests})
    accountLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests})
    accountLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests})

    accountLogon.EntityData.YListKeys = []string {}

    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + accountLogoff.EntityData.SegmentPath
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests})
    accountLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests})
    accountLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests})

    accountLogoff.EntityData.YListKeys = []string {}

    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + accountUpdate.EntityData.SegmentPath
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests})
    accountUpdate.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests})
    accountUpdate.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests})

    accountUpdate.EntityData.YListKeys = []string {}

    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization-statistics"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + sessionDisconnect.EntityData.SegmentPath
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests})
    sessionDisconnect.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests})
    sessionDisconnect.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests})

    sessionDisconnect.EntityData.YListKeys = []string {}

    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + singleServiceLogon.EntityData.SegmentPath
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests})
    singleServiceLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests})
    singleServiceLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests})

    singleServiceLogon.EntityData.YListKeys = []string {}

    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + singleServiceLogoff.EntityData.SegmentPath
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests})

    singleServiceLogoff.EntityData.YListKeys = []string {}

    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + singleServiceModify.EntityData.SegmentPath
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests})
    singleServiceModify.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests})
    singleServiceModify.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests})

    singleServiceModify.EntityData.YListKeys = []string {}

    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization-statistics"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/change-of-authorization-statistics/" + serviceMulti.EntityData.SegmentPath
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests})
    serviceMulti.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests})
    serviceMulti.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests})

    serviceMulti.EntityData.YListKeys = []string {}

    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics
// List of stats for Mobility
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobilityStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics) GetEntityData() *types.CommonEntityData {
    mobilityStatistics.EntityData.YFilter = mobilityStatistics.YFilter
    mobilityStatistics.EntityData.YangName = "mobility-statistics"
    mobilityStatistics.EntityData.BundleName = "cisco_ios_xr"
    mobilityStatistics.EntityData.ParentYangName = "accounting-stats-all"
    mobilityStatistics.EntityData.SegmentPath = "mobility-statistics"
    mobilityStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting-stats-all/" + mobilityStatistics.EntityData.SegmentPath
    mobilityStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobilityStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobilityStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobilityStatistics.EntityData.Children = types.NewOrderedMap()
    mobilityStatistics.EntityData.Leafs = types.NewOrderedMap()
    mobilityStatistics.EntityData.Leafs.Append("send-request-successes", types.YLeaf{"SendRequestSuccesses", mobilityStatistics.SendRequestSuccesses})
    mobilityStatistics.EntityData.Leafs.Append("send-request-failures", types.YLeaf{"SendRequestFailures", mobilityStatistics.SendRequestFailures})
    mobilityStatistics.EntityData.Leafs.Append("receive-response-successes", types.YLeaf{"ReceiveResponseSuccesses", mobilityStatistics.ReceiveResponseSuccesses})
    mobilityStatistics.EntityData.Leafs.Append("receive-response-failures", types.YLeaf{"ReceiveResponseFailures", mobilityStatistics.ReceiveResponseFailures})

    mobilityStatistics.EntityData.YListKeys = []string {}

    return &(mobilityStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization
// Change of authorization (COA) statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti
}

func (changeOfAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    changeOfAuthorization.EntityData.YFilter = changeOfAuthorization.YFilter
    changeOfAuthorization.EntityData.YangName = "change-of-authorization"
    changeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorization.EntityData.ParentYangName = "aaa"
    changeOfAuthorization.EntityData.SegmentPath = "change-of-authorization"
    changeOfAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + changeOfAuthorization.EntityData.SegmentPath
    changeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorization.EntityData.Children = types.NewOrderedMap()
    changeOfAuthorization.EntityData.Children.Append("account-logon", types.YChild{"AccountLogon", &changeOfAuthorization.AccountLogon})
    changeOfAuthorization.EntityData.Children.Append("account-logoff", types.YChild{"AccountLogoff", &changeOfAuthorization.AccountLogoff})
    changeOfAuthorization.EntityData.Children.Append("account-update", types.YChild{"AccountUpdate", &changeOfAuthorization.AccountUpdate})
    changeOfAuthorization.EntityData.Children.Append("session-disconnect", types.YChild{"SessionDisconnect", &changeOfAuthorization.SessionDisconnect})
    changeOfAuthorization.EntityData.Children.Append("single-service-logon", types.YChild{"SingleServiceLogon", &changeOfAuthorization.SingleServiceLogon})
    changeOfAuthorization.EntityData.Children.Append("single-service-logoff", types.YChild{"SingleServiceLogoff", &changeOfAuthorization.SingleServiceLogoff})
    changeOfAuthorization.EntityData.Children.Append("single-service-modify", types.YChild{"SingleServiceModify", &changeOfAuthorization.SingleServiceModify})
    changeOfAuthorization.EntityData.Children.Append("service-multi", types.YChild{"ServiceMulti", &changeOfAuthorization.ServiceMulti})
    changeOfAuthorization.EntityData.Leafs = types.NewOrderedMap()
    changeOfAuthorization.EntityData.Leafs.Append("unknown-account-cmd-resps", types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorization.UnknownAccountCmdResps})
    changeOfAuthorization.EntityData.Leafs.Append("unknown-service-cmd-resps", types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorization.UnknownServiceCmdResps})
    changeOfAuthorization.EntityData.Leafs.Append("unknown-cmd-resps", types.YLeaf{"UnknownCmdResps", changeOfAuthorization.UnknownCmdResps})
    changeOfAuthorization.EntityData.Leafs.Append("attr-list-retrieve-failure-resps", types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorization.AttrListRetrieveFailureResps})
    changeOfAuthorization.EntityData.Leafs.Append("resp-send-failure", types.YLeaf{"RespSendFailure", changeOfAuthorization.RespSendFailure})
    changeOfAuthorization.EntityData.Leafs.Append("internal-err-resps", types.YLeaf{"InternalErrResps", changeOfAuthorization.InternalErrResps})
    changeOfAuthorization.EntityData.Leafs.Append("service-profile-push-failure-resps", types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorization.ServiceProfilePushFailureResps})
    changeOfAuthorization.EntityData.Leafs.Append("no-cmd-resps", types.YLeaf{"NoCmdResps", changeOfAuthorization.NoCmdResps})
    changeOfAuthorization.EntityData.Leafs.Append("no-session-found-resps", types.YLeaf{"NoSessionFoundResps", changeOfAuthorization.NoSessionFoundResps})
    changeOfAuthorization.EntityData.Leafs.Append("no-session-peer-resps", types.YLeaf{"NoSessionPeerResps", changeOfAuthorization.NoSessionPeerResps})

    changeOfAuthorization.EntityData.YListKeys = []string {}

    return &(changeOfAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + accountLogon.EntityData.SegmentPath
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = types.NewOrderedMap()
    accountLogon.EntityData.Leafs = types.NewOrderedMap()
    accountLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests})
    accountLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests})
    accountLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests})

    accountLogon.EntityData.YListKeys = []string {}

    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + accountLogoff.EntityData.SegmentPath
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests})
    accountLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests})
    accountLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests})

    accountLogoff.EntityData.YListKeys = []string {}

    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + accountUpdate.EntityData.SegmentPath
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests})
    accountUpdate.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests})
    accountUpdate.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests})

    accountUpdate.EntityData.YListKeys = []string {}

    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + sessionDisconnect.EntityData.SegmentPath
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests})
    sessionDisconnect.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests})
    sessionDisconnect.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests})

    sessionDisconnect.EntityData.YListKeys = []string {}

    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + singleServiceLogon.EntityData.SegmentPath
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests})
    singleServiceLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests})
    singleServiceLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests})

    singleServiceLogon.EntityData.YListKeys = []string {}

    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + singleServiceLogoff.EntityData.SegmentPath
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests})

    singleServiceLogoff.EntityData.YListKeys = []string {}

    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + singleServiceModify.EntityData.SegmentPath
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests})
    singleServiceModify.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests})
    singleServiceModify.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests})

    singleServiceModify.EntityData.YListKeys = []string {}

    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/change-of-authorization/" + serviceMulti.EntityData.SegmentPath
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests})
    serviceMulti.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests})
    serviceMulti.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests})

    serviceMulti.EntityData.YListKeys = []string {}

    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization
// Authorization statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "aaa"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authorization.SentRequests})
    authorization.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authorization.AcceptedRequests})
    authorization.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authorization.SuccessfulRequests})
    authorization.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authorization.RejectedRequests})
    authorization.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authorization.UnreachableRequests})
    authorization.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authorization.ErroredRequests})
    authorization.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authorization.TerminatedRequests})

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization
// Aggregate authorization statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (aggregateAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization) GetEntityData() *types.CommonEntityData {
    aggregateAuthorization.EntityData.YFilter = aggregateAuthorization.YFilter
    aggregateAuthorization.EntityData.YangName = "aggregate-authorization"
    aggregateAuthorization.EntityData.BundleName = "cisco_ios_xr"
    aggregateAuthorization.EntityData.ParentYangName = "aaa"
    aggregateAuthorization.EntityData.SegmentPath = "aggregate-authorization"
    aggregateAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateAuthorization.EntityData.SegmentPath
    aggregateAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAuthorization.EntityData.Children = types.NewOrderedMap()
    aggregateAuthorization.EntityData.Leafs = types.NewOrderedMap()
    aggregateAuthorization.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", aggregateAuthorization.SentRequests})
    aggregateAuthorization.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", aggregateAuthorization.AcceptedRequests})
    aggregateAuthorization.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", aggregateAuthorization.SuccessfulRequests})
    aggregateAuthorization.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", aggregateAuthorization.RejectedRequests})
    aggregateAuthorization.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", aggregateAuthorization.UnreachableRequests})
    aggregateAuthorization.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", aggregateAuthorization.ErroredRequests})
    aggregateAuthorization.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", aggregateAuthorization.TerminatedRequests})

    aggregateAuthorization.EntityData.YListKeys = []string {}

    return &(aggregateAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll
// Display all subscriber management total
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of stats for accounting.
    AccountingStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics

    // List of stats for authentication.
    AuthenticationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics

    // List of stats for authorization.
    AuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics

    // List of stats for COA.
    ChangeOfAuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics

    // List of stats for Mobility.
    MobilityStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics
}

func (aggregateAccountingStatsAll *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll) GetEntityData() *types.CommonEntityData {
    aggregateAccountingStatsAll.EntityData.YFilter = aggregateAccountingStatsAll.YFilter
    aggregateAccountingStatsAll.EntityData.YangName = "aggregate-accounting-stats-all"
    aggregateAccountingStatsAll.EntityData.BundleName = "cisco_ios_xr"
    aggregateAccountingStatsAll.EntityData.ParentYangName = "aaa"
    aggregateAccountingStatsAll.EntityData.SegmentPath = "aggregate-accounting-stats-all"
    aggregateAccountingStatsAll.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateAccountingStatsAll.EntityData.SegmentPath
    aggregateAccountingStatsAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAccountingStatsAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAccountingStatsAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAccountingStatsAll.EntityData.Children = types.NewOrderedMap()
    aggregateAccountingStatsAll.EntityData.Children.Append("accounting-statistics", types.YChild{"AccountingStatistics", &aggregateAccountingStatsAll.AccountingStatistics})
    aggregateAccountingStatsAll.EntityData.Children.Append("authentication-statistics", types.YChild{"AuthenticationStatistics", &aggregateAccountingStatsAll.AuthenticationStatistics})
    aggregateAccountingStatsAll.EntityData.Children.Append("authorization-statistics", types.YChild{"AuthorizationStatistics", &aggregateAccountingStatsAll.AuthorizationStatistics})
    aggregateAccountingStatsAll.EntityData.Children.Append("change-of-authorization-statistics", types.YChild{"ChangeOfAuthorizationStatistics", &aggregateAccountingStatsAll.ChangeOfAuthorizationStatistics})
    aggregateAccountingStatsAll.EntityData.Children.Append("mobility-statistics", types.YChild{"MobilityStatistics", &aggregateAccountingStatsAll.MobilityStatistics})
    aggregateAccountingStatsAll.EntityData.Leafs = types.NewOrderedMap()

    aggregateAccountingStatsAll.EntityData.YListKeys = []string {}

    return &(aggregateAccountingStatsAll.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics
// List of stats for accounting
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight
}

func (accountingStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/" + accountingStatistics.EntityData.SegmentPath
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = types.NewOrderedMap()
    accountingStatistics.EntityData.Children.Append("start", types.YChild{"Start", &accountingStatistics.Start})
    accountingStatistics.EntityData.Children.Append("stop", types.YChild{"Stop", &accountingStatistics.Stop})
    accountingStatistics.EntityData.Children.Append("interim", types.YChild{"Interim", &accountingStatistics.Interim})
    accountingStatistics.EntityData.Children.Append("pass-through", types.YChild{"PassThrough", &accountingStatistics.PassThrough})
    accountingStatistics.EntityData.Children.Append("update", types.YChild{"Update", &accountingStatistics.Update})
    accountingStatistics.EntityData.Children.Append("interim-inflight", types.YChild{"InterimInflight", &accountingStatistics.InterimInflight})
    accountingStatistics.EntityData.Leafs = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", accountingStatistics.ActiveSessions})
    accountingStatistics.EntityData.Leafs.Append("started-sessions", types.YLeaf{"StartedSessions", accountingStatistics.StartedSessions})
    accountingStatistics.EntityData.Leafs.Append("stopped-sessions", types.YLeaf{"StoppedSessions", accountingStatistics.StoppedSessions})
    accountingStatistics.EntityData.Leafs.Append("policy-plane-errored-requests", types.YLeaf{"PolicyPlaneErroredRequests", accountingStatistics.PolicyPlaneErroredRequests})
    accountingStatistics.EntityData.Leafs.Append("policy-plane-unknown-requests", types.YLeaf{"PolicyPlaneUnknownRequests", accountingStatistics.PolicyPlaneUnknownRequests})

    accountingStatistics.EntityData.YListKeys = []string {}

    return &(accountingStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting-statistics"
    start.EntityData.SegmentPath = "start"
    start.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + start.EntityData.SegmentPath
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = types.NewOrderedMap()
    start.EntityData.Leafs = types.NewOrderedMap()
    start.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", start.ReceivedRequests})
    start.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", start.ErroredRequests})
    start.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests})
    start.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", start.AaaSentRequests})
    start.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses})
    start.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses})

    start.EntityData.YListKeys = []string {}

    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting-statistics"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Leafs = types.NewOrderedMap()
    stop.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", stop.ReceivedRequests})
    stop.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", stop.ErroredRequests})
    stop.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests})
    stop.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", stop.AaaSentRequests})
    stop.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses})
    stop.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses})

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting-statistics"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + interim.EntityData.SegmentPath
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = types.NewOrderedMap()
    interim.EntityData.Leafs = types.NewOrderedMap()
    interim.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", interim.ReceivedRequests})
    interim.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", interim.ErroredRequests})
    interim.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests})
    interim.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", interim.AaaSentRequests})
    interim.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses})
    interim.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses})

    interim.EntityData.YListKeys = []string {}

    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting-statistics"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + passThrough.EntityData.SegmentPath
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = types.NewOrderedMap()
    passThrough.EntityData.Leafs = types.NewOrderedMap()
    passThrough.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests})
    passThrough.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", passThrough.ErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests})
    passThrough.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses})
    passThrough.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses})

    passThrough.EntityData.YListKeys = []string {}

    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting-statistics"
    update.EntityData.SegmentPath = "update"
    update.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + update.EntityData.SegmentPath
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = types.NewOrderedMap()
    update.EntityData.Leafs = types.NewOrderedMap()
    update.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", update.ReceivedRequests})
    update.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", update.ErroredRequests})
    update.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests})
    update.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", update.AaaSentRequests})
    update.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses})
    update.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses})

    update.EntityData.YListKeys = []string {}

    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting-statistics"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/accounting-statistics/" + interimInflight.EntityData.SegmentPath
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = types.NewOrderedMap()
    interimInflight.EntityData.Leafs = types.NewOrderedMap()
    interimInflight.EntityData.Leafs.Append("quota-exhausts", types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts})
    interimInflight.EntityData.Leafs.Append("denied-requests", types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests})
    interimInflight.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests})
    interimInflight.EntityData.Leafs.Append("total-quota-of-requests", types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("remaining-quota-of-requests", types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("low-water-mark-quota-of-requests", types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests})

    interimInflight.EntityData.YListKeys = []string {}

    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics
// List of stats for authentication
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authenticationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/" + authenticationStatistics.EntityData.SegmentPath
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authenticationStatistics.SentRequests})
    authenticationStatistics.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authenticationStatistics.AcceptedRequests})
    authenticationStatistics.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authenticationStatistics.SuccessfulRequests})
    authenticationStatistics.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authenticationStatistics.RejectedRequests})
    authenticationStatistics.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authenticationStatistics.UnreachableRequests})
    authenticationStatistics.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authenticationStatistics.ErroredRequests})
    authenticationStatistics.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authenticationStatistics.TerminatedRequests})

    authenticationStatistics.EntityData.YListKeys = []string {}

    return &(authenticationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics
// List of stats for authorization
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    authorizationStatistics.EntityData.YFilter = authorizationStatistics.YFilter
    authorizationStatistics.EntityData.YangName = "authorization-statistics"
    authorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authorizationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    authorizationStatistics.EntityData.SegmentPath = "authorization-statistics"
    authorizationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/" + authorizationStatistics.EntityData.SegmentPath
    authorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizationStatistics.EntityData.Children = types.NewOrderedMap()
    authorizationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authorizationStatistics.EntityData.Leafs.Append("sent-requests", types.YLeaf{"SentRequests", authorizationStatistics.SentRequests})
    authorizationStatistics.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", authorizationStatistics.AcceptedRequests})
    authorizationStatistics.EntityData.Leafs.Append("successful-requests", types.YLeaf{"SuccessfulRequests", authorizationStatistics.SuccessfulRequests})
    authorizationStatistics.EntityData.Leafs.Append("rejected-requests", types.YLeaf{"RejectedRequests", authorizationStatistics.RejectedRequests})
    authorizationStatistics.EntityData.Leafs.Append("unreachable-requests", types.YLeaf{"UnreachableRequests", authorizationStatistics.UnreachableRequests})
    authorizationStatistics.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", authorizationStatistics.ErroredRequests})
    authorizationStatistics.EntityData.Leafs.Append("terminated-requests", types.YLeaf{"TerminatedRequests", authorizationStatistics.TerminatedRequests})

    authorizationStatistics.EntityData.YListKeys = []string {}

    return &(authorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics
// List of stats for COA
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
}

func (changeOfAuthorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    changeOfAuthorizationStatistics.EntityData.YFilter = changeOfAuthorizationStatistics.YFilter
    changeOfAuthorizationStatistics.EntityData.YangName = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorizationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    changeOfAuthorizationStatistics.EntityData.SegmentPath = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/" + changeOfAuthorizationStatistics.EntityData.SegmentPath
    changeOfAuthorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorizationStatistics.EntityData.Children = types.NewOrderedMap()
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-logon", types.YChild{"AccountLogon", &changeOfAuthorizationStatistics.AccountLogon})
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-logoff", types.YChild{"AccountLogoff", &changeOfAuthorizationStatistics.AccountLogoff})
    changeOfAuthorizationStatistics.EntityData.Children.Append("account-update", types.YChild{"AccountUpdate", &changeOfAuthorizationStatistics.AccountUpdate})
    changeOfAuthorizationStatistics.EntityData.Children.Append("session-disconnect", types.YChild{"SessionDisconnect", &changeOfAuthorizationStatistics.SessionDisconnect})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-logon", types.YChild{"SingleServiceLogon", &changeOfAuthorizationStatistics.SingleServiceLogon})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-logoff", types.YChild{"SingleServiceLogoff", &changeOfAuthorizationStatistics.SingleServiceLogoff})
    changeOfAuthorizationStatistics.EntityData.Children.Append("single-service-modify", types.YChild{"SingleServiceModify", &changeOfAuthorizationStatistics.SingleServiceModify})
    changeOfAuthorizationStatistics.EntityData.Children.Append("service-multi", types.YChild{"ServiceMulti", &changeOfAuthorizationStatistics.ServiceMulti})
    changeOfAuthorizationStatistics.EntityData.Leafs = types.NewOrderedMap()
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-account-cmd-resps", types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorizationStatistics.UnknownAccountCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-service-cmd-resps", types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorizationStatistics.UnknownServiceCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("unknown-cmd-resps", types.YLeaf{"UnknownCmdResps", changeOfAuthorizationStatistics.UnknownCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("attr-list-retrieve-failure-resps", types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorizationStatistics.AttrListRetrieveFailureResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("resp-send-failure", types.YLeaf{"RespSendFailure", changeOfAuthorizationStatistics.RespSendFailure})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("internal-err-resps", types.YLeaf{"InternalErrResps", changeOfAuthorizationStatistics.InternalErrResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("service-profile-push-failure-resps", types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorizationStatistics.ServiceProfilePushFailureResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-cmd-resps", types.YLeaf{"NoCmdResps", changeOfAuthorizationStatistics.NoCmdResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-session-found-resps", types.YLeaf{"NoSessionFoundResps", changeOfAuthorizationStatistics.NoSessionFoundResps})
    changeOfAuthorizationStatistics.EntityData.Leafs.Append("no-session-peer-resps", types.YLeaf{"NoSessionPeerResps", changeOfAuthorizationStatistics.NoSessionPeerResps})

    changeOfAuthorizationStatistics.EntityData.YListKeys = []string {}

    return &(changeOfAuthorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + accountLogon.EntityData.SegmentPath
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = types.NewOrderedMap()
    accountLogon.EntityData.Leafs = types.NewOrderedMap()
    accountLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests})
    accountLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests})
    accountLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests})

    accountLogon.EntityData.YListKeys = []string {}

    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + accountLogoff.EntityData.SegmentPath
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests})
    accountLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests})
    accountLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests})

    accountLogoff.EntityData.YListKeys = []string {}

    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + accountUpdate.EntityData.SegmentPath
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests})
    accountUpdate.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests})
    accountUpdate.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests})

    accountUpdate.EntityData.YListKeys = []string {}

    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization-statistics"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + sessionDisconnect.EntityData.SegmentPath
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests})
    sessionDisconnect.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests})
    sessionDisconnect.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests})

    sessionDisconnect.EntityData.YListKeys = []string {}

    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + singleServiceLogon.EntityData.SegmentPath
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests})
    singleServiceLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests})
    singleServiceLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests})

    singleServiceLogon.EntityData.YListKeys = []string {}

    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + singleServiceLogoff.EntityData.SegmentPath
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests})

    singleServiceLogoff.EntityData.YListKeys = []string {}

    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + singleServiceModify.EntityData.SegmentPath
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests})
    singleServiceModify.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests})
    singleServiceModify.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests})

    singleServiceModify.EntityData.YListKeys = []string {}

    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization-statistics"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/change-of-authorization-statistics/" + serviceMulti.EntityData.SegmentPath
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests})
    serviceMulti.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests})
    serviceMulti.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests})

    serviceMulti.EntityData.YListKeys = []string {}

    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics
// List of stats for Mobility
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobilityStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics) GetEntityData() *types.CommonEntityData {
    mobilityStatistics.EntityData.YFilter = mobilityStatistics.YFilter
    mobilityStatistics.EntityData.YangName = "mobility-statistics"
    mobilityStatistics.EntityData.BundleName = "cisco_ios_xr"
    mobilityStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    mobilityStatistics.EntityData.SegmentPath = "mobility-statistics"
    mobilityStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-accounting-stats-all/" + mobilityStatistics.EntityData.SegmentPath
    mobilityStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobilityStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobilityStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobilityStatistics.EntityData.Children = types.NewOrderedMap()
    mobilityStatistics.EntityData.Leafs = types.NewOrderedMap()
    mobilityStatistics.EntityData.Leafs.Append("send-request-successes", types.YLeaf{"SendRequestSuccesses", mobilityStatistics.SendRequestSuccesses})
    mobilityStatistics.EntityData.Leafs.Append("send-request-failures", types.YLeaf{"SendRequestFailures", mobilityStatistics.SendRequestFailures})
    mobilityStatistics.EntityData.Leafs.Append("receive-response-successes", types.YLeaf{"ReceiveResponseSuccesses", mobilityStatistics.ReceiveResponseSuccesses})
    mobilityStatistics.EntityData.Leafs.Append("receive-response-failures", types.YLeaf{"ReceiveResponseFailures", mobilityStatistics.ReceiveResponseFailures})

    mobilityStatistics.EntityData.YListKeys = []string {}

    return &(mobilityStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting
// Accounting statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight
}

func (accounting *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("start", types.YChild{"Start", &accounting.Start})
    accounting.EntityData.Children.Append("stop", types.YChild{"Stop", &accounting.Stop})
    accounting.EntityData.Children.Append("interim", types.YChild{"Interim", &accounting.Interim})
    accounting.EntityData.Children.Append("pass-through", types.YChild{"PassThrough", &accounting.PassThrough})
    accounting.EntityData.Children.Append("update", types.YChild{"Update", &accounting.Update})
    accounting.EntityData.Children.Append("interim-inflight", types.YChild{"InterimInflight", &accounting.InterimInflight})
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", accounting.ActiveSessions})
    accounting.EntityData.Leafs.Append("started-sessions", types.YLeaf{"StartedSessions", accounting.StartedSessions})
    accounting.EntityData.Leafs.Append("stopped-sessions", types.YLeaf{"StoppedSessions", accounting.StoppedSessions})
    accounting.EntityData.Leafs.Append("policy-plane-errored-requests", types.YLeaf{"PolicyPlaneErroredRequests", accounting.PolicyPlaneErroredRequests})
    accounting.EntityData.Leafs.Append("policy-plane-unknown-requests", types.YLeaf{"PolicyPlaneUnknownRequests", accounting.PolicyPlaneUnknownRequests})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting"
    start.EntityData.SegmentPath = "start"
    start.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + start.EntityData.SegmentPath
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = types.NewOrderedMap()
    start.EntityData.Leafs = types.NewOrderedMap()
    start.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", start.ReceivedRequests})
    start.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", start.ErroredRequests})
    start.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests})
    start.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", start.AaaSentRequests})
    start.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses})
    start.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses})

    start.EntityData.YListKeys = []string {}

    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + stop.EntityData.SegmentPath
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Leafs = types.NewOrderedMap()
    stop.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", stop.ReceivedRequests})
    stop.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", stop.ErroredRequests})
    stop.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests})
    stop.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", stop.AaaSentRequests})
    stop.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses})
    stop.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses})

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + interim.EntityData.SegmentPath
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = types.NewOrderedMap()
    interim.EntityData.Leafs = types.NewOrderedMap()
    interim.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", interim.ReceivedRequests})
    interim.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", interim.ErroredRequests})
    interim.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests})
    interim.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", interim.AaaSentRequests})
    interim.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses})
    interim.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses})

    interim.EntityData.YListKeys = []string {}

    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + passThrough.EntityData.SegmentPath
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = types.NewOrderedMap()
    passThrough.EntityData.Leafs = types.NewOrderedMap()
    passThrough.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests})
    passThrough.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", passThrough.ErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests})
    passThrough.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests})
    passThrough.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses})
    passThrough.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses})

    passThrough.EntityData.YListKeys = []string {}

    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting"
    update.EntityData.SegmentPath = "update"
    update.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + update.EntityData.SegmentPath
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = types.NewOrderedMap()
    update.EntityData.Leafs = types.NewOrderedMap()
    update.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", update.ReceivedRequests})
    update.EntityData.Leafs.Append("errored-requests", types.YLeaf{"ErroredRequests", update.ErroredRequests})
    update.EntityData.Leafs.Append("aaa-errored-requests", types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests})
    update.EntityData.Leafs.Append("aaa-sent-requests", types.YLeaf{"AaaSentRequests", update.AaaSentRequests})
    update.EntityData.Leafs.Append("aaa-succeeded-responses", types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses})
    update.EntityData.Leafs.Append("aaa-failed-responses", types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses})

    update.EntityData.YListKeys = []string {}

    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/accounting/" + interimInflight.EntityData.SegmentPath
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = types.NewOrderedMap()
    interimInflight.EntityData.Leafs = types.NewOrderedMap()
    interimInflight.EntityData.Leafs.Append("quota-exhausts", types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts})
    interimInflight.EntityData.Leafs.Append("denied-requests", types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests})
    interimInflight.EntityData.Leafs.Append("accepted-requests", types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests})
    interimInflight.EntityData.Leafs.Append("total-quota-of-requests", types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("remaining-quota-of-requests", types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests})
    interimInflight.EntityData.Leafs.Append("low-water-mark-quota-of-requests", types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests})

    interimInflight.EntityData.YListKeys = []string {}

    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility
// Mobility statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobility *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility) GetEntityData() *types.CommonEntityData {
    mobility.EntityData.YFilter = mobility.YFilter
    mobility.EntityData.YangName = "mobility"
    mobility.EntityData.BundleName = "cisco_ios_xr"
    mobility.EntityData.ParentYangName = "aaa"
    mobility.EntityData.SegmentPath = "mobility"
    mobility.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + mobility.EntityData.SegmentPath
    mobility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobility.EntityData.Children = types.NewOrderedMap()
    mobility.EntityData.Leafs = types.NewOrderedMap()
    mobility.EntityData.Leafs.Append("send-request-successes", types.YLeaf{"SendRequestSuccesses", mobility.SendRequestSuccesses})
    mobility.EntityData.Leafs.Append("send-request-failures", types.YLeaf{"SendRequestFailures", mobility.SendRequestFailures})
    mobility.EntityData.Leafs.Append("receive-response-successes", types.YLeaf{"ReceiveResponseSuccesses", mobility.ReceiveResponseSuccesses})
    mobility.EntityData.Leafs.Append("receive-response-failures", types.YLeaf{"ReceiveResponseFailures", mobility.ReceiveResponseFailures})

    mobility.EntityData.YListKeys = []string {}

    return &(mobility.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization
// Aggregate change of authorization (COA)
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti
}

func (aggregateChangeOfAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    aggregateChangeOfAuthorization.EntityData.YFilter = aggregateChangeOfAuthorization.YFilter
    aggregateChangeOfAuthorization.EntityData.YangName = "aggregate-change-of-authorization"
    aggregateChangeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    aggregateChangeOfAuthorization.EntityData.ParentYangName = "aaa"
    aggregateChangeOfAuthorization.EntityData.SegmentPath = "aggregate-change-of-authorization"
    aggregateChangeOfAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/" + aggregateChangeOfAuthorization.EntityData.SegmentPath
    aggregateChangeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateChangeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateChangeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateChangeOfAuthorization.EntityData.Children = types.NewOrderedMap()
    aggregateChangeOfAuthorization.EntityData.Children.Append("account-logon", types.YChild{"AccountLogon", &aggregateChangeOfAuthorization.AccountLogon})
    aggregateChangeOfAuthorization.EntityData.Children.Append("account-logoff", types.YChild{"AccountLogoff", &aggregateChangeOfAuthorization.AccountLogoff})
    aggregateChangeOfAuthorization.EntityData.Children.Append("account-update", types.YChild{"AccountUpdate", &aggregateChangeOfAuthorization.AccountUpdate})
    aggregateChangeOfAuthorization.EntityData.Children.Append("session-disconnect", types.YChild{"SessionDisconnect", &aggregateChangeOfAuthorization.SessionDisconnect})
    aggregateChangeOfAuthorization.EntityData.Children.Append("single-service-logon", types.YChild{"SingleServiceLogon", &aggregateChangeOfAuthorization.SingleServiceLogon})
    aggregateChangeOfAuthorization.EntityData.Children.Append("single-service-logoff", types.YChild{"SingleServiceLogoff", &aggregateChangeOfAuthorization.SingleServiceLogoff})
    aggregateChangeOfAuthorization.EntityData.Children.Append("single-service-modify", types.YChild{"SingleServiceModify", &aggregateChangeOfAuthorization.SingleServiceModify})
    aggregateChangeOfAuthorization.EntityData.Children.Append("service-multi", types.YChild{"ServiceMulti", &aggregateChangeOfAuthorization.ServiceMulti})
    aggregateChangeOfAuthorization.EntityData.Leafs = types.NewOrderedMap()
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("unknown-account-cmd-resps", types.YLeaf{"UnknownAccountCmdResps", aggregateChangeOfAuthorization.UnknownAccountCmdResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("unknown-service-cmd-resps", types.YLeaf{"UnknownServiceCmdResps", aggregateChangeOfAuthorization.UnknownServiceCmdResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("unknown-cmd-resps", types.YLeaf{"UnknownCmdResps", aggregateChangeOfAuthorization.UnknownCmdResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("attr-list-retrieve-failure-resps", types.YLeaf{"AttrListRetrieveFailureResps", aggregateChangeOfAuthorization.AttrListRetrieveFailureResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("resp-send-failure", types.YLeaf{"RespSendFailure", aggregateChangeOfAuthorization.RespSendFailure})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("internal-err-resps", types.YLeaf{"InternalErrResps", aggregateChangeOfAuthorization.InternalErrResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("service-profile-push-failure-resps", types.YLeaf{"ServiceProfilePushFailureResps", aggregateChangeOfAuthorization.ServiceProfilePushFailureResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("no-cmd-resps", types.YLeaf{"NoCmdResps", aggregateChangeOfAuthorization.NoCmdResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("no-session-found-resps", types.YLeaf{"NoSessionFoundResps", aggregateChangeOfAuthorization.NoSessionFoundResps})
    aggregateChangeOfAuthorization.EntityData.Leafs.Append("no-session-peer-resps", types.YLeaf{"NoSessionPeerResps", aggregateChangeOfAuthorization.NoSessionPeerResps})

    aggregateChangeOfAuthorization.EntityData.YListKeys = []string {}

    return &(aggregateChangeOfAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + accountLogon.EntityData.SegmentPath
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = types.NewOrderedMap()
    accountLogon.EntityData.Leafs = types.NewOrderedMap()
    accountLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests})
    accountLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests})
    accountLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests})

    accountLogon.EntityData.YListKeys = []string {}

    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + accountLogoff.EntityData.SegmentPath
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs = types.NewOrderedMap()
    accountLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests})
    accountLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests})
    accountLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests})

    accountLogoff.EntityData.YListKeys = []string {}

    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + accountUpdate.EntityData.SegmentPath
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs = types.NewOrderedMap()
    accountUpdate.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests})
    accountUpdate.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests})
    accountUpdate.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests})

    accountUpdate.EntityData.YListKeys = []string {}

    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "aggregate-change-of-authorization"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + sessionDisconnect.EntityData.SegmentPath
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs = types.NewOrderedMap()
    sessionDisconnect.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests})
    sessionDisconnect.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests})
    sessionDisconnect.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests})

    sessionDisconnect.EntityData.YListKeys = []string {}

    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + singleServiceLogon.EntityData.SegmentPath
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogon.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests})
    singleServiceLogon.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests})
    singleServiceLogon.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests})

    singleServiceLogon.EntityData.YListKeys = []string {}

    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + singleServiceLogoff.EntityData.SegmentPath
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs = types.NewOrderedMap()
    singleServiceLogoff.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests})
    singleServiceLogoff.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests})

    singleServiceLogoff.EntityData.YListKeys = []string {}

    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + singleServiceModify.EntityData.SegmentPath
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs = types.NewOrderedMap()
    singleServiceModify.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests})
    singleServiceModify.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests})
    singleServiceModify.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests})

    singleServiceModify.EntityData.YListKeys = []string {}

    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "aggregate-change-of-authorization"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/aaa/aggregate-change-of-authorization/" + serviceMulti.EntityData.SegmentPath
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs = types.NewOrderedMap()
    serviceMulti.EntityData.Leafs.Append("received-requests", types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests})
    serviceMulti.EntityData.Leafs.Append("acknowledged-requests", types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests})
    serviceMulti.EntityData.Leafs.Append("non-acknowledged-requests", types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests})

    serviceMulti.EntityData.YListKeys = []string {}

    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary
// Aggregate summary of statistics
type Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber control policy not applied on interface. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSubscriberControlPolicyOnInterface interface{}

    // No control policy class match during subscriber start. The type is
    // interface{} with range: 0..18446744073709551615.
    NoClassMatchInStartRequest interface{}

    // NAS port attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    NasPortAttributeFormatWarnings interface{}

    // NAS port ID attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    NasPortIdAttributeFormatWarnings interface{}

    // Destination station ID attribute format warnings. The type is interface{}
    // with range: 0..18446744073709551615.
    DestinationStationIdAttributeFormatWarnings interface{}

    // Calling station ID attribute format warnings. The type is interface{} with
    // range: 0..18446744073709551615.
    CallingStationIdAttributeFormatWarnings interface{}

    // Username attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    UsernameAttributeFormatWarnings interface{}

    // User profiles installed. The type is interface{} with range:
    // 0..18446744073709551615.
    InstallUserProfiles interface{}

    // User profile install errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileInstallErrors interface{}

    // User profile removals. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileRemovals interface{}

    // User profile errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileErrors interface{}

    // Session Disconnect Quota Exhausts. The type is interface{} with range:
    // 0..18446744073709551615.
    SessDiscQuotaExhausts interface{}

    // Session Disconnect Request Queued, no quota. The type is interface{} with
    // range: 0..18446744073709551615.
    SessDiscNoQuota interface{}

    // Session Disconnect Request Accepted, quota available. The type is
    // interface{} with range: 0..18446744073709551615.
    SessDiscQuotaAvail interface{}

    // Session Disconnect Requests not Dequeued, reconciliation in progress. The
    // type is interface{} with range: 0..18446744073709551615.
    SessDiscReconIp interface{}

    // Session Disconnect Requests not Dequeued, no quota. The type is interface{}
    // with range: 0..18446744073709551615.
    SessDiscNoneStarted interface{}

    // Session Disconnect Quota. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQuota interface{}

    // Session Disconnect Quota Remaining. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQuotaRemaining interface{}

    // Session Disconnect Requests Queued. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQCount interface{}
}

func (aggregateSummary *Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary) GetEntityData() *types.CommonEntityData {
    aggregateSummary.EntityData.YFilter = aggregateSummary.YFilter
    aggregateSummary.EntityData.YangName = "aggregate-summary"
    aggregateSummary.EntityData.BundleName = "cisco_ios_xr"
    aggregateSummary.EntityData.ParentYangName = "statistics"
    aggregateSummary.EntityData.SegmentPath = "aggregate-summary"
    aggregateSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/" + aggregateSummary.EntityData.SegmentPath
    aggregateSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateSummary.EntityData.Children = types.NewOrderedMap()
    aggregateSummary.EntityData.Leafs = types.NewOrderedMap()
    aggregateSummary.EntityData.Leafs.Append("no-subscriber-control-policy-on-interface", types.YLeaf{"NoSubscriberControlPolicyOnInterface", aggregateSummary.NoSubscriberControlPolicyOnInterface})
    aggregateSummary.EntityData.Leafs.Append("no-class-match-in-start-request", types.YLeaf{"NoClassMatchInStartRequest", aggregateSummary.NoClassMatchInStartRequest})
    aggregateSummary.EntityData.Leafs.Append("nas-port-attribute-format-warnings", types.YLeaf{"NasPortAttributeFormatWarnings", aggregateSummary.NasPortAttributeFormatWarnings})
    aggregateSummary.EntityData.Leafs.Append("nas-port-id-attribute-format-warnings", types.YLeaf{"NasPortIdAttributeFormatWarnings", aggregateSummary.NasPortIdAttributeFormatWarnings})
    aggregateSummary.EntityData.Leafs.Append("destination-station-id-attribute-format-warnings", types.YLeaf{"DestinationStationIdAttributeFormatWarnings", aggregateSummary.DestinationStationIdAttributeFormatWarnings})
    aggregateSummary.EntityData.Leafs.Append("calling-station-id-attribute-format-warnings", types.YLeaf{"CallingStationIdAttributeFormatWarnings", aggregateSummary.CallingStationIdAttributeFormatWarnings})
    aggregateSummary.EntityData.Leafs.Append("username-attribute-format-warnings", types.YLeaf{"UsernameAttributeFormatWarnings", aggregateSummary.UsernameAttributeFormatWarnings})
    aggregateSummary.EntityData.Leafs.Append("install-user-profiles", types.YLeaf{"InstallUserProfiles", aggregateSummary.InstallUserProfiles})
    aggregateSummary.EntityData.Leafs.Append("user-profile-install-errors", types.YLeaf{"UserProfileInstallErrors", aggregateSummary.UserProfileInstallErrors})
    aggregateSummary.EntityData.Leafs.Append("user-profile-removals", types.YLeaf{"UserProfileRemovals", aggregateSummary.UserProfileRemovals})
    aggregateSummary.EntityData.Leafs.Append("user-profile-errors", types.YLeaf{"UserProfileErrors", aggregateSummary.UserProfileErrors})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-quota-exhausts", types.YLeaf{"SessDiscQuotaExhausts", aggregateSummary.SessDiscQuotaExhausts})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-no-quota", types.YLeaf{"SessDiscNoQuota", aggregateSummary.SessDiscNoQuota})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-quota-avail", types.YLeaf{"SessDiscQuotaAvail", aggregateSummary.SessDiscQuotaAvail})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-recon-ip", types.YLeaf{"SessDiscReconIp", aggregateSummary.SessDiscReconIp})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-none-started", types.YLeaf{"SessDiscNoneStarted", aggregateSummary.SessDiscNoneStarted})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-quota", types.YLeaf{"SessDiscQuota", aggregateSummary.SessDiscQuota})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-quota-remaining", types.YLeaf{"SessDiscQuotaRemaining", aggregateSummary.SessDiscQuotaRemaining})
    aggregateSummary.EntityData.Leafs.Append("sess-disc-q-count", types.YLeaf{"SessDiscQCount", aggregateSummary.SessDiscQCount})

    aggregateSummary.EntityData.YListKeys = []string {}

    return &(aggregateSummary.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique
// Disconnect Unique Summary statistics
type Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of disconnect history items. The type is slice of
    // Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique_HistoryData.
    HistoryData []*Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique_HistoryData
}

func (disconnUnique *Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique) GetEntityData() *types.CommonEntityData {
    disconnUnique.EntityData.YFilter = disconnUnique.YFilter
    disconnUnique.EntityData.YangName = "disconn-unique"
    disconnUnique.EntityData.BundleName = "cisco_ios_xr"
    disconnUnique.EntityData.ParentYangName = "statistics"
    disconnUnique.EntityData.SegmentPath = "disconn-unique"
    disconnUnique.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/" + disconnUnique.EntityData.SegmentPath
    disconnUnique.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disconnUnique.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disconnUnique.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disconnUnique.EntityData.Children = types.NewOrderedMap()
    disconnUnique.EntityData.Children.Append("history-data", types.YChild{"HistoryData", nil})
    for i := range disconnUnique.HistoryData {
        types.SetYListKey(disconnUnique.HistoryData[i], i)
        disconnUnique.EntityData.Children.Append(types.GetSegmentPath(disconnUnique.HistoryData[i]), types.YChild{"HistoryData", disconnUnique.HistoryData[i]})
    }
    disconnUnique.EntityData.Leafs = types.NewOrderedMap()

    disconnUnique.EntityData.YListKeys = []string {}

    return &(disconnUnique.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique_HistoryData
// List of disconnect history items
type Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique_HistoryData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Disconnect reason. The type is string.
    DiscReason interface{}

    // Number of sessions that disconnected this way. The type is interface{} with
    // range: 0..4294967295.
    SessionCount interface{}

    // Latest activity. The type is interface{} with range: 0..4294967295.
    LatestActivity interface{}

    // Interface name. The type is string.
    IfName interface{}
}

func (historyData *Subscriber_Manager_Nodes_Node_Statistics_DisconnUnique_HistoryData) GetEntityData() *types.CommonEntityData {
    historyData.EntityData.YFilter = historyData.YFilter
    historyData.EntityData.YangName = "history-data"
    historyData.EntityData.BundleName = "cisco_ios_xr"
    historyData.EntityData.ParentYangName = "disconn-unique"
    historyData.EntityData.SegmentPath = "history-data" + types.AddNoKeyToken(historyData)
    historyData.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/disconn-unique/" + historyData.EntityData.SegmentPath
    historyData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyData.EntityData.Children = types.NewOrderedMap()
    historyData.EntityData.Leafs = types.NewOrderedMap()
    historyData.EntityData.Leafs.Append("disc-reason", types.YLeaf{"DiscReason", historyData.DiscReason})
    historyData.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", historyData.SessionCount})
    historyData.EntityData.Leafs.Append("latest-activity", types.YLeaf{"LatestActivity", historyData.LatestActivity})
    historyData.EntityData.Leafs.Append("if-name", types.YLeaf{"IfName", historyData.IfName})

    historyData.EntityData.YListKeys = []string {}

    return &(historyData.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Srg
// Geo Redundancy statistics
type Subscriber_Manager_Nodes_Node_Statistics_Srg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Txlist Send Triggered. The type is interface{} with range: 0..4294967295.
    TxlistSendTriggered interface{}

    // Txlist Send Failed. The type is interface{} with range: 0..4294967295.
    TxlistSendFailed interface{}

    // Txlist send failed due to not active. The type is interface{} with range:
    // 0..4294967295.
    TxlistSendFailedNotactive interface{}

    // Txlist Send Success. The type is interface{} with range: 0..4294967295.
    ActualTxlistSent interface{}

    // Element already in Txlist. The type is interface{} with range:
    // 0..4294967295.
    AlreadyinTxlist interface{}

    // Txlist Encode. The type is interface{} with range: 0..4294967295.
    TxlistEncode interface{}

    // Txlist encode Failed. The type is interface{} with range: 0..4294967295.
    TxlistEncodeFail interface{}

    // Txlist Create Update Encode. The type is interface{} with range:
    // 0..4294967295.
    CreateUpdateEncode interface{}

    // Txlist Delete Encode. The type is interface{} with range: 0..4294967295.
    DeleteEncode interface{}

    // Txlist Create/update clean callback. The type is interface{} with range:
    // 0..4294967295.
    CreateUpdCleanCallback interface{}

    // Txlist Delete clean callback. The type is interface{} with range:
    // 0..4294967295.
    DeleteCleanCallback interface{}

    // Slave Recieved Sync. The type is interface{} with range: 0..4294967295.
    SlaveRecvEntry interface{}

    // Decode failed on Slave. The type is interface{} with range: 0..4294967295.
    SlaveDecodeFail interface{}

    // Create Update received on slave. The type is interface{} with range:
    // 0..4294967295.
    SlaveCreateUpdate interface{}

    // Delete received on slave. The type is interface{} with range:
    // 0..4294967295.
    SlaveDelete interface{}

    // SRG context allocated. The type is interface{} with range: 0..4294967295.
    SrgContextMalloc interface{}

    // SRG context freed. The type is interface{} with range: 0..4294967295.
    SrgContextFree interface{}

    // Number of SODs Received. The type is interface{} with range: 0..4294967295.
    SodCount interface{}

    // Number of EODs Received. The type is interface{} with range: 0..4294967295.
    EodCount interface{}

    // Number of Replay Requests Within SOD EOD Window. The type is interface{}
    // with range: 0..4294967295.
    SodEodReplayReqCount interface{}

    // Number of Sessions Marked as Invalid Within SOD EOD Window. The type is
    // interface{} with range: 0..4294967295.
    SodEodDirtyMarkCount interface{}

    // Number of Sessions Invalid Deletes Within SOD EOD Window. The type is
    // interface{} with range: 0..4294967295.
    SodEodDirtyDeleteCount interface{}

    // Number of ACKs sent to Srg. The type is interface{} with range:
    // 0..4294967295.
    AckToSrg interface{}

    // Number of NACKs sent to Srg. The type is interface{} with range:
    // 0..4294967295.
    NackToSrg interface{}

    // Number of NACKs Failed to send to Srg. The type is interface{} with range:
    // 0..4294967295.
    NackToSrgFailCnt interface{}

    // Number of Txlist remove all calls. The type is interface{} with range:
    // 0..4294967295.
    TxlistRemoveAll interface{}

    // Number for Txlist delete for sync msg. The type is interface{} with range:
    // 0..4294967295.
    TxlistDelSync interface{}

    // Number of Txlist delete for sync which are not linked. The type is
    // interface{} with range: 0..4294967295.
    TxlistDelSyncNotlinked interface{}

    // Number of Txlist delete for App msg. The type is interface{} with range:
    // 0..4294967295.
    TxlistDelApp interface{}

    // Number of Txlist delete for App which are not linked. The type is
    // interface{} with range: 0..4294967295.
    TxlistDelAppNotlinked interface{}

    // Number of Txlist Cleanup called on Invalid subscriber srg state. The type
    // is interface{} with range: 0..4294967295.
    TxlistCleanInvalidState interface{}

    // Number of Internal errors upon Master Txlist remove all call. The type is
    // interface{} with range: 0..4294967295.
    TxlistRemoveAllInternalError interface{}

    // Flag indicating SRG Flow control enabled or not. The type is bool.
    IsSrgFlowControlEnabled interface{}

    // Maximum no.of inflight sessions allowed. The type is interface{} with
    // range: 0..4294967295.
    MaxInflightSessoinCount interface{}

    // Threshold Limit to resume the flow control. The type is interface{} with
    // range: 0..4294967295.
    FlowControlResumeThreshold interface{}

    // No.of Sessions inflight at given time. The type is interface{} with range:
    // 0..4294967295.
    InflightSessionCount interface{}

    // No.of inflight sessions added. The type is interface{} with range:
    // 0..4294967295.
    InflightAddCount interface{}

    // Inflight Underrun Counter. The type is interface{} with range:
    // 0..4294967295.
    InflightUnderRunCount interface{}

    // Memory Alloc Failures for Inflight Entry. The type is interface{} with
    // range: 0..4294967295.
    InflightAllocFails interface{}

    // Inflight Entry Insert Failures. The type is interface{} with range:
    // 0..4294967295.
    InflightInsertFailures interface{}

    // Inflight Deletes Count. The type is interface{} with range: 0..4294967295.
    InflightDeletes interface{}

    // Inflight Entries not found during delete. The type is interface{} with
    // range: 0..4294967295.
    InflightNotFound interface{}

    // Inflight Entry Delete Failures. The type is interface{} with range:
    // 0..4294967295.
    InflightDeleteFailures interface{}

    // Total No.of times Pause is Enabled. The type is interface{} with range:
    // 0..4294967295.
    TotalPauseCount interface{}

    // Total No.of times Resume is triggered. The type is interface{} with range:
    // 0..4294967295.
    TotalResumeCount interface{}

    // Total No of times Dont send to Txlist. The type is interface{} with range:
    // 0..4294967295.
    TotalDontSendToTxlist interface{}

    // Total No of times SRG Not Master. The type is interface{} with range:
    // 0..4294967295.
    TotalSrgNotMaster interface{}

    // Total No of times Master EOMS Pending. The type is interface{} with range:
    // 0..4294967295.
    TotalMasterEomsPending interface{}

    // Total No of times Master EOMS Pending Cleared. The type is interface{} with
    // range: 0..4294967295.
    TotalMasterEomsPendingCleared interface{}

    // Amount of time paused during last flow control window. The type is
    // interface{} with range: 0..18446744073709551615.
    LastPausePeriod interface{}

    // Total Amount of time paused during all flow control windows. The type is
    // interface{} with range: 0..18446744073709551615.
    TotalPauseTime interface{}

    // Timestamp of recent Pause Event. The type is interface{} with range:
    // 0..18446744073709551615.
    LastPauseTime interface{}

    // Timestamp of recent Resume Event. The type is interface{} with range:
    // 0..18446744073709551615.
    LastResumeTime interface{}
}

func (srg *Subscriber_Manager_Nodes_Node_Statistics_Srg) GetEntityData() *types.CommonEntityData {
    srg.EntityData.YFilter = srg.YFilter
    srg.EntityData.YangName = "srg"
    srg.EntityData.BundleName = "cisco_ios_xr"
    srg.EntityData.ParentYangName = "statistics"
    srg.EntityData.SegmentPath = "srg"
    srg.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/manager/nodes/node/statistics/" + srg.EntityData.SegmentPath
    srg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srg.EntityData.Children = types.NewOrderedMap()
    srg.EntityData.Leafs = types.NewOrderedMap()
    srg.EntityData.Leafs.Append("txlist-send-triggered", types.YLeaf{"TxlistSendTriggered", srg.TxlistSendTriggered})
    srg.EntityData.Leafs.Append("txlist-send-failed", types.YLeaf{"TxlistSendFailed", srg.TxlistSendFailed})
    srg.EntityData.Leafs.Append("txlist-send-failed-notactive", types.YLeaf{"TxlistSendFailedNotactive", srg.TxlistSendFailedNotactive})
    srg.EntityData.Leafs.Append("actual-txlist-sent", types.YLeaf{"ActualTxlistSent", srg.ActualTxlistSent})
    srg.EntityData.Leafs.Append("alreadyin-txlist", types.YLeaf{"AlreadyinTxlist", srg.AlreadyinTxlist})
    srg.EntityData.Leafs.Append("txlist-encode", types.YLeaf{"TxlistEncode", srg.TxlistEncode})
    srg.EntityData.Leafs.Append("txlist-encode-fail", types.YLeaf{"TxlistEncodeFail", srg.TxlistEncodeFail})
    srg.EntityData.Leafs.Append("create-update-encode", types.YLeaf{"CreateUpdateEncode", srg.CreateUpdateEncode})
    srg.EntityData.Leafs.Append("delete-encode", types.YLeaf{"DeleteEncode", srg.DeleteEncode})
    srg.EntityData.Leafs.Append("create-upd-clean-callback", types.YLeaf{"CreateUpdCleanCallback", srg.CreateUpdCleanCallback})
    srg.EntityData.Leafs.Append("delete-clean-callback", types.YLeaf{"DeleteCleanCallback", srg.DeleteCleanCallback})
    srg.EntityData.Leafs.Append("slave-recv-entry", types.YLeaf{"SlaveRecvEntry", srg.SlaveRecvEntry})
    srg.EntityData.Leafs.Append("slave-decode-fail", types.YLeaf{"SlaveDecodeFail", srg.SlaveDecodeFail})
    srg.EntityData.Leafs.Append("slave-create-update", types.YLeaf{"SlaveCreateUpdate", srg.SlaveCreateUpdate})
    srg.EntityData.Leafs.Append("slave-delete", types.YLeaf{"SlaveDelete", srg.SlaveDelete})
    srg.EntityData.Leafs.Append("srg-context-malloc", types.YLeaf{"SrgContextMalloc", srg.SrgContextMalloc})
    srg.EntityData.Leafs.Append("srg-context-free", types.YLeaf{"SrgContextFree", srg.SrgContextFree})
    srg.EntityData.Leafs.Append("sod-count", types.YLeaf{"SodCount", srg.SodCount})
    srg.EntityData.Leafs.Append("eod-count", types.YLeaf{"EodCount", srg.EodCount})
    srg.EntityData.Leafs.Append("sod-eod-replay-req-count", types.YLeaf{"SodEodReplayReqCount", srg.SodEodReplayReqCount})
    srg.EntityData.Leafs.Append("sod-eod-dirty-mark-count", types.YLeaf{"SodEodDirtyMarkCount", srg.SodEodDirtyMarkCount})
    srg.EntityData.Leafs.Append("sod-eod-dirty-delete-count", types.YLeaf{"SodEodDirtyDeleteCount", srg.SodEodDirtyDeleteCount})
    srg.EntityData.Leafs.Append("ack-to-srg", types.YLeaf{"AckToSrg", srg.AckToSrg})
    srg.EntityData.Leafs.Append("nack-to-srg", types.YLeaf{"NackToSrg", srg.NackToSrg})
    srg.EntityData.Leafs.Append("nack-to-srg-fail-cnt", types.YLeaf{"NackToSrgFailCnt", srg.NackToSrgFailCnt})
    srg.EntityData.Leafs.Append("txlist-remove-all", types.YLeaf{"TxlistRemoveAll", srg.TxlistRemoveAll})
    srg.EntityData.Leafs.Append("txlist-del-sync", types.YLeaf{"TxlistDelSync", srg.TxlistDelSync})
    srg.EntityData.Leafs.Append("txlist-del-sync-notlinked", types.YLeaf{"TxlistDelSyncNotlinked", srg.TxlistDelSyncNotlinked})
    srg.EntityData.Leafs.Append("txlist-del-app", types.YLeaf{"TxlistDelApp", srg.TxlistDelApp})
    srg.EntityData.Leafs.Append("txlist-del-app-notlinked", types.YLeaf{"TxlistDelAppNotlinked", srg.TxlistDelAppNotlinked})
    srg.EntityData.Leafs.Append("txlist-clean-invalid-state", types.YLeaf{"TxlistCleanInvalidState", srg.TxlistCleanInvalidState})
    srg.EntityData.Leafs.Append("txlist-remove-all-internal-error", types.YLeaf{"TxlistRemoveAllInternalError", srg.TxlistRemoveAllInternalError})
    srg.EntityData.Leafs.Append("is-srg-flow-control-enabled", types.YLeaf{"IsSrgFlowControlEnabled", srg.IsSrgFlowControlEnabled})
    srg.EntityData.Leafs.Append("max-inflight-sessoin-count", types.YLeaf{"MaxInflightSessoinCount", srg.MaxInflightSessoinCount})
    srg.EntityData.Leafs.Append("flow-control-resume-threshold", types.YLeaf{"FlowControlResumeThreshold", srg.FlowControlResumeThreshold})
    srg.EntityData.Leafs.Append("inflight-session-count", types.YLeaf{"InflightSessionCount", srg.InflightSessionCount})
    srg.EntityData.Leafs.Append("inflight-add-count", types.YLeaf{"InflightAddCount", srg.InflightAddCount})
    srg.EntityData.Leafs.Append("inflight-under-run-count", types.YLeaf{"InflightUnderRunCount", srg.InflightUnderRunCount})
    srg.EntityData.Leafs.Append("inflight-alloc-fails", types.YLeaf{"InflightAllocFails", srg.InflightAllocFails})
    srg.EntityData.Leafs.Append("inflight-insert-failures", types.YLeaf{"InflightInsertFailures", srg.InflightInsertFailures})
    srg.EntityData.Leafs.Append("inflight-deletes", types.YLeaf{"InflightDeletes", srg.InflightDeletes})
    srg.EntityData.Leafs.Append("inflight-not-found", types.YLeaf{"InflightNotFound", srg.InflightNotFound})
    srg.EntityData.Leafs.Append("inflight-delete-failures", types.YLeaf{"InflightDeleteFailures", srg.InflightDeleteFailures})
    srg.EntityData.Leafs.Append("total-pause-count", types.YLeaf{"TotalPauseCount", srg.TotalPauseCount})
    srg.EntityData.Leafs.Append("total-resume-count", types.YLeaf{"TotalResumeCount", srg.TotalResumeCount})
    srg.EntityData.Leafs.Append("total-dont-send-to-txlist", types.YLeaf{"TotalDontSendToTxlist", srg.TotalDontSendToTxlist})
    srg.EntityData.Leafs.Append("total-srg-not-master", types.YLeaf{"TotalSrgNotMaster", srg.TotalSrgNotMaster})
    srg.EntityData.Leafs.Append("total-master-eoms-pending", types.YLeaf{"TotalMasterEomsPending", srg.TotalMasterEomsPending})
    srg.EntityData.Leafs.Append("total-master-eoms-pending-cleared", types.YLeaf{"TotalMasterEomsPendingCleared", srg.TotalMasterEomsPendingCleared})
    srg.EntityData.Leafs.Append("last-pause-period", types.YLeaf{"LastPausePeriod", srg.LastPausePeriod})
    srg.EntityData.Leafs.Append("total-pause-time", types.YLeaf{"TotalPauseTime", srg.TotalPauseTime})
    srg.EntityData.Leafs.Append("last-pause-time", types.YLeaf{"LastPauseTime", srg.LastPauseTime})
    srg.EntityData.Leafs.Append("last-resume-time", types.YLeaf{"LastResumeTime", srg.LastResumeTime})

    srg.EntityData.YListKeys = []string {}

    return &(srg.EntityData)
}

// Subscriber_Session
// Subscriber session operational data
type Subscriber_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of subscriber session supported nodes.
    Nodes Subscriber_Session_Nodes
}

func (session *Subscriber_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "subscriber"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("nodes", types.YChild{"Nodes", &session.Nodes})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// Subscriber_Session_Nodes
// List of subscriber session supported nodes
type Subscriber_Session_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session operational data for a particular node. The type is
    // slice of Subscriber_Session_Nodes_Node.
    Node []*Subscriber_Session_Nodes_Node
}

func (nodes *Subscriber_Session_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "session"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/" + nodes.EntityData.SegmentPath
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

// Subscriber_Session_Nodes_Node
// Subscriber session operational data for a
// particular node
type Subscriber_Session_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of subscriber session supported srg roles.
    SrgRoles Subscriber_Session_Nodes_Node_SrgRoles

    // Summary information filtered by authorization state.
    AuthorSummaries Subscriber_Session_Nodes_Node_AuthorSummaries

    // Subscriber session summary information.
    Summary Subscriber_Session_Nodes_Node_Summary

    // Summary information filtered by MAC address.
    MacSummaries Subscriber_Session_Nodes_Node_MacSummaries

    // Summary information filtered by interface.
    InterfaceSummaries Subscriber_Session_Nodes_Node_InterfaceSummaries

    // Summary information filtered by authentication state.
    AuthenticationSummaries Subscriber_Session_Nodes_Node_AuthenticationSummaries

    // Summary information filtered by session state.
    StateSummaries Subscriber_Session_Nodes_Node_StateSummaries

    // Summary information filtered by IPv4 address and VRF.
    Ipv4AddressVrfSummaries Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries

    // Summary information filtered by address family.
    AddressFamilySummaries Subscriber_Session_Nodes_Node_AddressFamilySummaries

    // Summary information filtered by username.
    UsernameSummaries Subscriber_Session_Nodes_Node_UsernameSummaries

    // Summary information filtered by access interface.
    AccessInterfaceSummaries Subscriber_Session_Nodes_Node_AccessInterfaceSummaries

    // Summary information filtered by subscriber IPv4 address.
    Ipv4AddressSummaries Subscriber_Session_Nodes_Node_Ipv4AddressSummaries

    // Summary information filtered by VRF.
    VrfSummaries Subscriber_Session_Nodes_Node_VrfSummaries

    // IP subscriber sessions.
    Sessions Subscriber_Session_Nodes_Node_Sessions
}

func (node *Subscriber_Session_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("srg-roles", types.YChild{"SrgRoles", &node.SrgRoles})
    node.EntityData.Children.Append("author-summaries", types.YChild{"AuthorSummaries", &node.AuthorSummaries})
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("mac-summaries", types.YChild{"MacSummaries", &node.MacSummaries})
    node.EntityData.Children.Append("interface-summaries", types.YChild{"InterfaceSummaries", &node.InterfaceSummaries})
    node.EntityData.Children.Append("authentication-summaries", types.YChild{"AuthenticationSummaries", &node.AuthenticationSummaries})
    node.EntityData.Children.Append("state-summaries", types.YChild{"StateSummaries", &node.StateSummaries})
    node.EntityData.Children.Append("ipv4-address-vrf-summaries", types.YChild{"Ipv4AddressVrfSummaries", &node.Ipv4AddressVrfSummaries})
    node.EntityData.Children.Append("address-family-summaries", types.YChild{"AddressFamilySummaries", &node.AddressFamilySummaries})
    node.EntityData.Children.Append("username-summaries", types.YChild{"UsernameSummaries", &node.UsernameSummaries})
    node.EntityData.Children.Append("access-interface-summaries", types.YChild{"AccessInterfaceSummaries", &node.AccessInterfaceSummaries})
    node.EntityData.Children.Append("ipv4-address-summaries", types.YChild{"Ipv4AddressSummaries", &node.Ipv4AddressSummaries})
    node.EntityData.Children.Append("vrf-summaries", types.YChild{"VrfSummaries", &node.VrfSummaries})
    node.EntityData.Children.Append("sessions", types.YChild{"Sessions", &node.Sessions})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles
// List of subscriber session supported srg
// roles
type Subscriber_Session_Nodes_Node_SrgRoles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session operational data based on srg role. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole.
    SrgRole []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole
}

func (srgRoles *Subscriber_Session_Nodes_Node_SrgRoles) GetEntityData() *types.CommonEntityData {
    srgRoles.EntityData.YFilter = srgRoles.YFilter
    srgRoles.EntityData.YangName = "srg-roles"
    srgRoles.EntityData.BundleName = "cisco_ios_xr"
    srgRoles.EntityData.ParentYangName = "node"
    srgRoles.EntityData.SegmentPath = "srg-roles"
    srgRoles.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + srgRoles.EntityData.SegmentPath
    srgRoles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgRoles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgRoles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgRoles.EntityData.Children = types.NewOrderedMap()
    srgRoles.EntityData.Children.Append("srg-role", types.YChild{"SrgRole", nil})
    for i := range srgRoles.SrgRole {
        srgRoles.EntityData.Children.Append(types.GetSegmentPath(srgRoles.SrgRole[i]), types.YChild{"SrgRole", srgRoles.SrgRole[i]})
    }
    srgRoles.EntityData.Leafs = types.NewOrderedMap()

    srgRoles.EntityData.YListKeys = []string {}

    return &(srgRoles.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole
// Subscriber session operational data based on
// srg role
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber srg role. The type is
    // SubscriberSrgOperFilterFlag.
    Srg interface{}

    // Summary information filtered by authorization state.
    AuthorSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries

    // Summary information filtered by username.
    UsernameSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries

    // Summary information filtered by MAC address.
    MacSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries

    // Summary information filtered by interface.
    InterfaceSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries

    // Summary information filtered by session state.
    StateSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries

    // Summary information filtered by authentication state.
    AuthenticationSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries

    // IP subscriber sessions.
    SubscriberSessions Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions

    // Summary information filtered by IPv4 address and VRF.
    Ipv4AddressVrfSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries

    // Subscriber session summary information.
    SrgSummary Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary

    // Summary information filtered by access interface.
    AccessInterfaceSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries

    // Summary information filtered by address family.
    AddressFamilySummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries

    // Summary information filtered by subscriber IPv4 address.
    Ipv4AddressSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries

    // Summary information filtered by VRF.
    VrfSummaries Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries
}

func (srgRole *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole) GetEntityData() *types.CommonEntityData {
    srgRole.EntityData.YFilter = srgRole.YFilter
    srgRole.EntityData.YangName = "srg-role"
    srgRole.EntityData.BundleName = "cisco_ios_xr"
    srgRole.EntityData.ParentYangName = "srg-roles"
    srgRole.EntityData.SegmentPath = "srg-role" + types.AddKeyToken(srgRole.Srg, "srg")
    srgRole.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/" + srgRole.EntityData.SegmentPath
    srgRole.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgRole.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgRole.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgRole.EntityData.Children = types.NewOrderedMap()
    srgRole.EntityData.Children.Append("author-summaries", types.YChild{"AuthorSummaries", &srgRole.AuthorSummaries})
    srgRole.EntityData.Children.Append("username-summaries", types.YChild{"UsernameSummaries", &srgRole.UsernameSummaries})
    srgRole.EntityData.Children.Append("mac-summaries", types.YChild{"MacSummaries", &srgRole.MacSummaries})
    srgRole.EntityData.Children.Append("interface-summaries", types.YChild{"InterfaceSummaries", &srgRole.InterfaceSummaries})
    srgRole.EntityData.Children.Append("state-summaries", types.YChild{"StateSummaries", &srgRole.StateSummaries})
    srgRole.EntityData.Children.Append("authentication-summaries", types.YChild{"AuthenticationSummaries", &srgRole.AuthenticationSummaries})
    srgRole.EntityData.Children.Append("subscriber-sessions", types.YChild{"SubscriberSessions", &srgRole.SubscriberSessions})
    srgRole.EntityData.Children.Append("ipv4-address-vrf-summaries", types.YChild{"Ipv4AddressVrfSummaries", &srgRole.Ipv4AddressVrfSummaries})
    srgRole.EntityData.Children.Append("srg-summary", types.YChild{"SrgSummary", &srgRole.SrgSummary})
    srgRole.EntityData.Children.Append("access-interface-summaries", types.YChild{"AccessInterfaceSummaries", &srgRole.AccessInterfaceSummaries})
    srgRole.EntityData.Children.Append("address-family-summaries", types.YChild{"AddressFamilySummaries", &srgRole.AddressFamilySummaries})
    srgRole.EntityData.Children.Append("ipv4-address-summaries", types.YChild{"Ipv4AddressSummaries", &srgRole.Ipv4AddressSummaries})
    srgRole.EntityData.Children.Append("vrf-summaries", types.YChild{"VrfSummaries", &srgRole.VrfSummaries})
    srgRole.EntityData.Leafs = types.NewOrderedMap()
    srgRole.EntityData.Leafs.Append("srg", types.YLeaf{"Srg", srgRole.Srg})

    srgRole.EntityData.YListKeys = []string {"Srg"}

    return &(srgRole.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries
// Summary information filtered by
// authorization state
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authorization summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary.
    AuthorSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary
}

func (authorSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries) GetEntityData() *types.CommonEntityData {
    authorSummaries.EntityData.YFilter = authorSummaries.YFilter
    authorSummaries.EntityData.YangName = "author-summaries"
    authorSummaries.EntityData.BundleName = "cisco_ios_xr"
    authorSummaries.EntityData.ParentYangName = "srg-role"
    authorSummaries.EntityData.SegmentPath = "author-summaries"
    authorSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + authorSummaries.EntityData.SegmentPath
    authorSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummaries.EntityData.Children = types.NewOrderedMap()
    authorSummaries.EntityData.Children.Append("author-summary", types.YChild{"AuthorSummary", nil})
    for i := range authorSummaries.AuthorSummary {
        authorSummaries.EntityData.Children.Append(types.GetSegmentPath(authorSummaries.AuthorSummary[i]), types.YChild{"AuthorSummary", authorSummaries.AuthorSummary[i]})
    }
    authorSummaries.EntityData.Leafs = types.NewOrderedMap()

    authorSummaries.EntityData.YListKeys = []string {}

    return &(authorSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary
// authorization summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Authorization state. The type is
    // SubscriberAuthorStateFilterFlag.
    AuthorState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr
}

func (authorSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary) GetEntityData() *types.CommonEntityData {
    authorSummary.EntityData.YFilter = authorSummary.YFilter
    authorSummary.EntityData.YangName = "author-summary"
    authorSummary.EntityData.BundleName = "cisco_ios_xr"
    authorSummary.EntityData.ParentYangName = "author-summaries"
    authorSummary.EntityData.SegmentPath = "author-summary" + types.AddKeyToken(authorSummary.AuthorState, "author-state")
    authorSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/" + authorSummary.EntityData.SegmentPath
    authorSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummary.EntityData.Children = types.NewOrderedMap()
    authorSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &authorSummary.StateXr})
    authorSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &authorSummary.AddressFamilyXr})
    authorSummary.EntityData.Leafs = types.NewOrderedMap()
    authorSummary.EntityData.Leafs.Append("author-state", types.YLeaf{"AuthorState", authorSummary.AuthorState})

    authorSummary.EntityData.YListKeys = []string {"AuthorState"}

    return &(authorSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "author-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "author-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/author-summaries/author-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries
// Summary information filtered by username
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary.
    UsernameSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary
}

func (usernameSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries) GetEntityData() *types.CommonEntityData {
    usernameSummaries.EntityData.YFilter = usernameSummaries.YFilter
    usernameSummaries.EntityData.YangName = "username-summaries"
    usernameSummaries.EntityData.BundleName = "cisco_ios_xr"
    usernameSummaries.EntityData.ParentYangName = "srg-role"
    usernameSummaries.EntityData.SegmentPath = "username-summaries"
    usernameSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + usernameSummaries.EntityData.SegmentPath
    usernameSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummaries.EntityData.Children = types.NewOrderedMap()
    usernameSummaries.EntityData.Children.Append("username-summary", types.YChild{"UsernameSummary", nil})
    for i := range usernameSummaries.UsernameSummary {
        usernameSummaries.EntityData.Children.Append(types.GetSegmentPath(usernameSummaries.UsernameSummary[i]), types.YChild{"UsernameSummary", usernameSummaries.UsernameSummary[i]})
    }
    usernameSummaries.EntityData.Leafs = types.NewOrderedMap()

    usernameSummaries.EntityData.YListKeys = []string {}

    return &(usernameSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary
// Username summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber username. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Username interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr
}

func (usernameSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary) GetEntityData() *types.CommonEntityData {
    usernameSummary.EntityData.YFilter = usernameSummary.YFilter
    usernameSummary.EntityData.YangName = "username-summary"
    usernameSummary.EntityData.BundleName = "cisco_ios_xr"
    usernameSummary.EntityData.ParentYangName = "username-summaries"
    usernameSummary.EntityData.SegmentPath = "username-summary" + types.AddKeyToken(usernameSummary.Username, "username")
    usernameSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/" + usernameSummary.EntityData.SegmentPath
    usernameSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummary.EntityData.Children = types.NewOrderedMap()
    usernameSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &usernameSummary.StateXr})
    usernameSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &usernameSummary.AddressFamilyXr})
    usernameSummary.EntityData.Leafs = types.NewOrderedMap()
    usernameSummary.EntityData.Leafs.Append("username", types.YLeaf{"Username", usernameSummary.Username})

    usernameSummary.EntityData.YListKeys = []string {"Username"}

    return &(usernameSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "username-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "username-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/username-summaries/username-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries
// Summary information filtered by MAC address
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary.
    MacSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary
}

func (macSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries) GetEntityData() *types.CommonEntityData {
    macSummaries.EntityData.YFilter = macSummaries.YFilter
    macSummaries.EntityData.YangName = "mac-summaries"
    macSummaries.EntityData.BundleName = "cisco_ios_xr"
    macSummaries.EntityData.ParentYangName = "srg-role"
    macSummaries.EntityData.SegmentPath = "mac-summaries"
    macSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + macSummaries.EntityData.SegmentPath
    macSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummaries.EntityData.Children = types.NewOrderedMap()
    macSummaries.EntityData.Children.Append("mac-summary", types.YChild{"MacSummary", nil})
    for i := range macSummaries.MacSummary {
        macSummaries.EntityData.Children.Append(types.GetSegmentPath(macSummaries.MacSummary[i]), types.YChild{"MacSummary", macSummaries.MacSummary[i]})
    }
    macSummaries.EntityData.Leafs = types.NewOrderedMap()

    macSummaries.EntityData.YListKeys = []string {}

    return &(macSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary
// MAC address summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr
}

func (macSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary) GetEntityData() *types.CommonEntityData {
    macSummary.EntityData.YFilter = macSummary.YFilter
    macSummary.EntityData.YangName = "mac-summary"
    macSummary.EntityData.BundleName = "cisco_ios_xr"
    macSummary.EntityData.ParentYangName = "mac-summaries"
    macSummary.EntityData.SegmentPath = "mac-summary" + types.AddKeyToken(macSummary.MacAddress, "mac-address")
    macSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/" + macSummary.EntityData.SegmentPath
    macSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummary.EntityData.Children = types.NewOrderedMap()
    macSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &macSummary.StateXr})
    macSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &macSummary.AddressFamilyXr})
    macSummary.EntityData.Leafs = types.NewOrderedMap()
    macSummary.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macSummary.MacAddress})

    macSummary.EntityData.YListKeys = []string {"MacAddress"}

    return &(macSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "mac-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "mac-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/mac-summaries/mac-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries
// Summary information filtered by interface
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary.
    InterfaceSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary
}

func (interfaceSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries) GetEntityData() *types.CommonEntityData {
    interfaceSummaries.EntityData.YFilter = interfaceSummaries.YFilter
    interfaceSummaries.EntityData.YangName = "interface-summaries"
    interfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummaries.EntityData.ParentYangName = "srg-role"
    interfaceSummaries.EntityData.SegmentPath = "interface-summaries"
    interfaceSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + interfaceSummaries.EntityData.SegmentPath
    interfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummaries.EntityData.Children = types.NewOrderedMap()
    interfaceSummaries.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", nil})
    for i := range interfaceSummaries.InterfaceSummary {
        interfaceSummaries.EntityData.Children.Append(types.GetSegmentPath(interfaceSummaries.InterfaceSummary[i]), types.YChild{"InterfaceSummary", interfaceSummaries.InterfaceSummary[i]})
    }
    interfaceSummaries.EntityData.Leafs = types.NewOrderedMap()

    interfaceSummaries.EntityData.YListKeys = []string {}

    return &(interfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary
// Interface summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
}

func (interfaceSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "interface-summaries"
    interfaceSummary.EntityData.SegmentPath = "interface-summary" + types.AddKeyToken(interfaceSummary.InterfaceName, "interface-name")
    interfaceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/" + interfaceSummary.EntityData.SegmentPath
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &interfaceSummary.StateXr})
    interfaceSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &interfaceSummary.AddressFamilyXr})
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName})

    interfaceSummary.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/interface-summaries/interface-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries
// Summary information filtered by session
// state
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary.
    StateSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary
}

func (stateSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries) GetEntityData() *types.CommonEntityData {
    stateSummaries.EntityData.YFilter = stateSummaries.YFilter
    stateSummaries.EntityData.YangName = "state-summaries"
    stateSummaries.EntityData.BundleName = "cisco_ios_xr"
    stateSummaries.EntityData.ParentYangName = "srg-role"
    stateSummaries.EntityData.SegmentPath = "state-summaries"
    stateSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + stateSummaries.EntityData.SegmentPath
    stateSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummaries.EntityData.Children = types.NewOrderedMap()
    stateSummaries.EntityData.Children.Append("state-summary", types.YChild{"StateSummary", nil})
    for i := range stateSummaries.StateSummary {
        stateSummaries.EntityData.Children.Append(types.GetSegmentPath(stateSummaries.StateSummary[i]), types.YChild{"StateSummary", stateSummaries.StateSummary[i]})
    }
    stateSummaries.EntityData.Leafs = types.NewOrderedMap()

    stateSummaries.EntityData.YListKeys = []string {}

    return &(stateSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber state. The type is
    // SubscriberStateFilterFlag.
    State interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr
}

func (stateSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary) GetEntityData() *types.CommonEntityData {
    stateSummary.EntityData.YFilter = stateSummary.YFilter
    stateSummary.EntityData.YangName = "state-summary"
    stateSummary.EntityData.BundleName = "cisco_ios_xr"
    stateSummary.EntityData.ParentYangName = "state-summaries"
    stateSummary.EntityData.SegmentPath = "state-summary" + types.AddKeyToken(stateSummary.State, "state")
    stateSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/" + stateSummary.EntityData.SegmentPath
    stateSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummary.EntityData.Children = types.NewOrderedMap()
    stateSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &stateSummary.StateXr})
    stateSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &stateSummary.AddressFamilyXr})
    stateSummary.EntityData.Leafs = types.NewOrderedMap()
    stateSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", stateSummary.State})

    stateSummary.EntityData.YListKeys = []string {"State"}

    return &(stateSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "state-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "state-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/state-summaries/state-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries
// Summary information filtered by
// authentication state
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authentication summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary.
    AuthenticationSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary
}

func (authenticationSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries) GetEntityData() *types.CommonEntityData {
    authenticationSummaries.EntityData.YFilter = authenticationSummaries.YFilter
    authenticationSummaries.EntityData.YangName = "authentication-summaries"
    authenticationSummaries.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummaries.EntityData.ParentYangName = "srg-role"
    authenticationSummaries.EntityData.SegmentPath = "authentication-summaries"
    authenticationSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + authenticationSummaries.EntityData.SegmentPath
    authenticationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummaries.EntityData.Children = types.NewOrderedMap()
    authenticationSummaries.EntityData.Children.Append("authentication-summary", types.YChild{"AuthenticationSummary", nil})
    for i := range authenticationSummaries.AuthenticationSummary {
        authenticationSummaries.EntityData.Children.Append(types.GetSegmentPath(authenticationSummaries.AuthenticationSummary[i]), types.YChild{"AuthenticationSummary", authenticationSummaries.AuthenticationSummary[i]})
    }
    authenticationSummaries.EntityData.Leafs = types.NewOrderedMap()

    authenticationSummaries.EntityData.YListKeys = []string {}

    return &(authenticationSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary
// authentication summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Authentication state. The type is
    // SubscriberAuthenStateFilterFlag.
    AuthenticationState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
}

func (authenticationSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary) GetEntityData() *types.CommonEntityData {
    authenticationSummary.EntityData.YFilter = authenticationSummary.YFilter
    authenticationSummary.EntityData.YangName = "authentication-summary"
    authenticationSummary.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummary.EntityData.ParentYangName = "authentication-summaries"
    authenticationSummary.EntityData.SegmentPath = "authentication-summary" + types.AddKeyToken(authenticationSummary.AuthenticationState, "authentication-state")
    authenticationSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/" + authenticationSummary.EntityData.SegmentPath
    authenticationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummary.EntityData.Children = types.NewOrderedMap()
    authenticationSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &authenticationSummary.StateXr})
    authenticationSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &authenticationSummary.AddressFamilyXr})
    authenticationSummary.EntityData.Leafs = types.NewOrderedMap()
    authenticationSummary.EntityData.Leafs.Append("authentication-state", types.YLeaf{"AuthenticationState", authenticationSummary.AuthenticationState})

    authenticationSummary.EntityData.YListKeys = []string {"AuthenticationState"}

    return &(authenticationSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "authentication-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "authentication-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/authentication-summaries/authentication-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions
// IP subscriber sessions
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session information. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession.
    SubscriberSession []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession
}

func (subscriberSessions *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions) GetEntityData() *types.CommonEntityData {
    subscriberSessions.EntityData.YFilter = subscriberSessions.YFilter
    subscriberSessions.EntityData.YangName = "subscriber-sessions"
    subscriberSessions.EntityData.BundleName = "cisco_ios_xr"
    subscriberSessions.EntityData.ParentYangName = "srg-role"
    subscriberSessions.EntityData.SegmentPath = "subscriber-sessions"
    subscriberSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + subscriberSessions.EntityData.SegmentPath
    subscriberSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberSessions.EntityData.Children = types.NewOrderedMap()
    subscriberSessions.EntityData.Children.Append("subscriber-session", types.YChild{"SubscriberSession", nil})
    for i := range subscriberSessions.SubscriberSession {
        subscriberSessions.EntityData.Children.Append(types.GetSegmentPath(subscriberSessions.SubscriberSession[i]), types.YChild{"SubscriberSession", subscriberSessions.SubscriberSession[i]})
    }
    subscriberSessions.EntityData.Leafs = types.NewOrderedMap()

    subscriberSessions.EntityData.YListKeys = []string {}

    return &(subscriberSessions.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession
// Subscriber session information
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session ID. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    SessionId interface{}

    // Subscriber session type. The type is IedgeOperSession.
    SessionType interface{}

    // PPPoE sub type. The type is IedgePppSub.
    PppoeSubType interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Circuit ID. The type is string.
    CircuitId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // PPPoE LNS address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LnsAddress interface{}

    // PPPoE LAC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LacAddress interface{}

    // PPPoE LAC tunnel client authentication ID. The type is string.
    TunnelClientAuthenticationId interface{}

    // PPPoE LAC tunnel server authentication ID. The type is string.
    TunnelServerAuthenticationId interface{}

    // Session ip address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SessionIpAddress interface{}

    // Session IPv6 address. The type is string.
    SessionIpv6Address interface{}

    // Session IPv6 prefix. The type is string.
    SessionIpv6Prefix interface{}

    // Session delegated IPv6 prefix. The type is string.
    DelegatedIpv6Prefix interface{}

    // IPv6 Interface ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Ipv6InterfaceId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // NAS port. The type is string.
    NasPort interface{}

    // Username. The type is string.
    Username interface{}

    // Client Username. The type is string.
    Clientname interface{}

    // Formatted Username. The type is string.
    Formattedname interface{}

    // If true, session is authentic. The type is bool.
    IsSessionAuthentic interface{}

    // If true, session is authorized. The type is bool.
    IsSessionAuthor interface{}

    // Session state. The type is IedgeOperSessionState.
    SessionState interface{}

    // Session creation time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SessionCreationTimeEpoch interface{}

    // Time when idle state change occurred in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    IdleStateChangeTime interface{}

    // Total session idle time (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    TotalSessionIdleTime interface{}

    // Access interface name associated with the session. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    AccessInterfaceName interface{}

    // Accounting information.
    Accounting Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting

    // Subscriber control policy applied to this session. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SubPolicyData.
    SubPolicyData []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SubPolicyData

    // List of subscriber services associated to this session. The type is slice
    // of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionServiceInfo.
    SessionServiceInfo []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionServiceInfo

    // Subscriber change of authorization information. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionChangeOfAuthorization.
    SessionChangeOfAuthorization []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionChangeOfAuthorization
}

func (subscriberSession *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession) GetEntityData() *types.CommonEntityData {
    subscriberSession.EntityData.YFilter = subscriberSession.YFilter
    subscriberSession.EntityData.YangName = "subscriber-session"
    subscriberSession.EntityData.BundleName = "cisco_ios_xr"
    subscriberSession.EntityData.ParentYangName = "subscriber-sessions"
    subscriberSession.EntityData.SegmentPath = "subscriber-session" + types.AddKeyToken(subscriberSession.SessionId, "session-id")
    subscriberSession.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/" + subscriberSession.EntityData.SegmentPath
    subscriberSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberSession.EntityData.Children = types.NewOrderedMap()
    subscriberSession.EntityData.Children.Append("accounting", types.YChild{"Accounting", &subscriberSession.Accounting})
    subscriberSession.EntityData.Children.Append("sub-policy-data", types.YChild{"SubPolicyData", nil})
    for i := range subscriberSession.SubPolicyData {
        types.SetYListKey(subscriberSession.SubPolicyData[i], i)
        subscriberSession.EntityData.Children.Append(types.GetSegmentPath(subscriberSession.SubPolicyData[i]), types.YChild{"SubPolicyData", subscriberSession.SubPolicyData[i]})
    }
    subscriberSession.EntityData.Children.Append("session-service-info", types.YChild{"SessionServiceInfo", nil})
    for i := range subscriberSession.SessionServiceInfo {
        types.SetYListKey(subscriberSession.SessionServiceInfo[i], i)
        subscriberSession.EntityData.Children.Append(types.GetSegmentPath(subscriberSession.SessionServiceInfo[i]), types.YChild{"SessionServiceInfo", subscriberSession.SessionServiceInfo[i]})
    }
    subscriberSession.EntityData.Children.Append("session-change-of-authorization", types.YChild{"SessionChangeOfAuthorization", nil})
    for i := range subscriberSession.SessionChangeOfAuthorization {
        types.SetYListKey(subscriberSession.SessionChangeOfAuthorization[i], i)
        subscriberSession.EntityData.Children.Append(types.GetSegmentPath(subscriberSession.SessionChangeOfAuthorization[i]), types.YChild{"SessionChangeOfAuthorization", subscriberSession.SessionChangeOfAuthorization[i]})
    }
    subscriberSession.EntityData.Leafs = types.NewOrderedMap()
    subscriberSession.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", subscriberSession.SessionId})
    subscriberSession.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", subscriberSession.SessionType})
    subscriberSession.EntityData.Leafs.Append("pppoe-sub-type", types.YLeaf{"PppoeSubType", subscriberSession.PppoeSubType})
    subscriberSession.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", subscriberSession.InterfaceName})
    subscriberSession.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", subscriberSession.VrfName})
    subscriberSession.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", subscriberSession.CircuitId})
    subscriberSession.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", subscriberSession.RemoteId})
    subscriberSession.EntityData.Leafs.Append("lns-address", types.YLeaf{"LnsAddress", subscriberSession.LnsAddress})
    subscriberSession.EntityData.Leafs.Append("lac-address", types.YLeaf{"LacAddress", subscriberSession.LacAddress})
    subscriberSession.EntityData.Leafs.Append("tunnel-client-authentication-id", types.YLeaf{"TunnelClientAuthenticationId", subscriberSession.TunnelClientAuthenticationId})
    subscriberSession.EntityData.Leafs.Append("tunnel-server-authentication-id", types.YLeaf{"TunnelServerAuthenticationId", subscriberSession.TunnelServerAuthenticationId})
    subscriberSession.EntityData.Leafs.Append("session-ip-address", types.YLeaf{"SessionIpAddress", subscriberSession.SessionIpAddress})
    subscriberSession.EntityData.Leafs.Append("session-ipv6-address", types.YLeaf{"SessionIpv6Address", subscriberSession.SessionIpv6Address})
    subscriberSession.EntityData.Leafs.Append("session-ipv6-prefix", types.YLeaf{"SessionIpv6Prefix", subscriberSession.SessionIpv6Prefix})
    subscriberSession.EntityData.Leafs.Append("delegated-ipv6-prefix", types.YLeaf{"DelegatedIpv6Prefix", subscriberSession.DelegatedIpv6Prefix})
    subscriberSession.EntityData.Leafs.Append("ipv6-interface-id", types.YLeaf{"Ipv6InterfaceId", subscriberSession.Ipv6InterfaceId})
    subscriberSession.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", subscriberSession.MacAddress})
    subscriberSession.EntityData.Leafs.Append("account-session-id", types.YLeaf{"AccountSessionId", subscriberSession.AccountSessionId})
    subscriberSession.EntityData.Leafs.Append("nas-port", types.YLeaf{"NasPort", subscriberSession.NasPort})
    subscriberSession.EntityData.Leafs.Append("username", types.YLeaf{"Username", subscriberSession.Username})
    subscriberSession.EntityData.Leafs.Append("clientname", types.YLeaf{"Clientname", subscriberSession.Clientname})
    subscriberSession.EntityData.Leafs.Append("formattedname", types.YLeaf{"Formattedname", subscriberSession.Formattedname})
    subscriberSession.EntityData.Leafs.Append("is-session-authentic", types.YLeaf{"IsSessionAuthentic", subscriberSession.IsSessionAuthentic})
    subscriberSession.EntityData.Leafs.Append("is-session-author", types.YLeaf{"IsSessionAuthor", subscriberSession.IsSessionAuthor})
    subscriberSession.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", subscriberSession.SessionState})
    subscriberSession.EntityData.Leafs.Append("session-creation-time-epoch", types.YLeaf{"SessionCreationTimeEpoch", subscriberSession.SessionCreationTimeEpoch})
    subscriberSession.EntityData.Leafs.Append("idle-state-change-time", types.YLeaf{"IdleStateChangeTime", subscriberSession.IdleStateChangeTime})
    subscriberSession.EntityData.Leafs.Append("total-session-idle-time", types.YLeaf{"TotalSessionIdleTime", subscriberSession.TotalSessionIdleTime})
    subscriberSession.EntityData.Leafs.Append("access-interface-name", types.YLeaf{"AccessInterfaceName", subscriberSession.AccessInterfaceName})

    subscriberSession.EntityData.YListKeys = []string {"SessionId"}

    return &(subscriberSession.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting
// Accounting information
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting information. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting_AccountingSession.
    AccountingSession []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting_AccountingSession
}

func (accounting *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "subscriber-session"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/subscriber-session/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("accounting-session", types.YChild{"AccountingSession", nil})
    for i := range accounting.AccountingSession {
        types.SetYListKey(accounting.AccountingSession[i], i)
        accounting.EntityData.Children.Append(types.GetSegmentPath(accounting.AccountingSession[i]), types.YChild{"AccountingSession", accounting.AccountingSession[i]})
    }
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting_AccountingSession
// Accounting information
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting_AccountingSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Accounting State Error Code for Accounting Session. The type is interface{}
    // with range: 0..4294967295.
    AccountingStateRc interface{}

    // Accounting Stop State. The type is interface{} with range: 0..4294967295.
    AccountingStopState interface{}

    // Accounting record context name. The type is string.
    RecordContextName interface{}

    // AAA method list name used to perform accounting. The type is string.
    MethodListName interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // Accounting start time in epoch-seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    AccountingStartTimeEpoch interface{}

    // True if interim accounting is enabled. The type is bool.
    IsInterimAccountingEnabled interface{}

    // Interim accounting interval (in minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    InterimInterval interface{}

    // Time of last successful interim update in epoch-seconds . The type is
    // interface{} with range: 0..18446744073709551615. Units are second.
    LastSuccessfulInterimUpdateTimeEpoch interface{}

    // Time of next interim update attempt (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    NextInterimUpdateAttemptTime interface{}

    // Time of last interim update attempt in epoch-seconds. The type is
    // interface{} with range: 0..18446744073709551615. Units are second.
    LastInterimUpdateAttemptTimeEpoch interface{}

    // Number of interim updates sent. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdates interface{}

    // Number of interim updates accepted. The type is interface{} with range:
    // 0..4294967295.
    AcceptedInterimUpdates interface{}

    // Number of interim updates rejected. The type is interface{} with range:
    // 0..4294967295.
    RejectedInterimUpdates interface{}

    // Number of interim update send failures. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdateFailures interface{}
}

func (accountingSession *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_Accounting_AccountingSession) GetEntityData() *types.CommonEntityData {
    accountingSession.EntityData.YFilter = accountingSession.YFilter
    accountingSession.EntityData.YangName = "accounting-session"
    accountingSession.EntityData.BundleName = "cisco_ios_xr"
    accountingSession.EntityData.ParentYangName = "accounting"
    accountingSession.EntityData.SegmentPath = "accounting-session" + types.AddNoKeyToken(accountingSession)
    accountingSession.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/subscriber-session/accounting/" + accountingSession.EntityData.SegmentPath
    accountingSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingSession.EntityData.Children = types.NewOrderedMap()
    accountingSession.EntityData.Leafs = types.NewOrderedMap()
    accountingSession.EntityData.Leafs.Append("accounting-state-rc", types.YLeaf{"AccountingStateRc", accountingSession.AccountingStateRc})
    accountingSession.EntityData.Leafs.Append("accounting-stop-state", types.YLeaf{"AccountingStopState", accountingSession.AccountingStopState})
    accountingSession.EntityData.Leafs.Append("record-context-name", types.YLeaf{"RecordContextName", accountingSession.RecordContextName})
    accountingSession.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", accountingSession.MethodListName})
    accountingSession.EntityData.Leafs.Append("account-session-id", types.YLeaf{"AccountSessionId", accountingSession.AccountSessionId})
    accountingSession.EntityData.Leafs.Append("accounting-start-time-epoch", types.YLeaf{"AccountingStartTimeEpoch", accountingSession.AccountingStartTimeEpoch})
    accountingSession.EntityData.Leafs.Append("is-interim-accounting-enabled", types.YLeaf{"IsInterimAccountingEnabled", accountingSession.IsInterimAccountingEnabled})
    accountingSession.EntityData.Leafs.Append("interim-interval", types.YLeaf{"InterimInterval", accountingSession.InterimInterval})
    accountingSession.EntityData.Leafs.Append("last-successful-interim-update-time-epoch", types.YLeaf{"LastSuccessfulInterimUpdateTimeEpoch", accountingSession.LastSuccessfulInterimUpdateTimeEpoch})
    accountingSession.EntityData.Leafs.Append("next-interim-update-attempt-time", types.YLeaf{"NextInterimUpdateAttemptTime", accountingSession.NextInterimUpdateAttemptTime})
    accountingSession.EntityData.Leafs.Append("last-interim-update-attempt-time-epoch", types.YLeaf{"LastInterimUpdateAttemptTimeEpoch", accountingSession.LastInterimUpdateAttemptTimeEpoch})
    accountingSession.EntityData.Leafs.Append("sent-interim-updates", types.YLeaf{"SentInterimUpdates", accountingSession.SentInterimUpdates})
    accountingSession.EntityData.Leafs.Append("accepted-interim-updates", types.YLeaf{"AcceptedInterimUpdates", accountingSession.AcceptedInterimUpdates})
    accountingSession.EntityData.Leafs.Append("rejected-interim-updates", types.YLeaf{"RejectedInterimUpdates", accountingSession.RejectedInterimUpdates})
    accountingSession.EntityData.Leafs.Append("sent-interim-update-failures", types.YLeaf{"SentInterimUpdateFailures", accountingSession.SentInterimUpdateFailures})

    accountingSession.EntityData.YListKeys = []string {}

    return &(accountingSession.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SubPolicyData
// Subscriber control policy applied to this
// session
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SubPolicyData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Matching event, condition and executed actions. The type is string.
    PolicyEpoch interface{}
}

func (subPolicyData *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SubPolicyData) GetEntityData() *types.CommonEntityData {
    subPolicyData.EntityData.YFilter = subPolicyData.YFilter
    subPolicyData.EntityData.YangName = "sub-policy-data"
    subPolicyData.EntityData.BundleName = "cisco_ios_xr"
    subPolicyData.EntityData.ParentYangName = "subscriber-session"
    subPolicyData.EntityData.SegmentPath = "sub-policy-data" + types.AddNoKeyToken(subPolicyData)
    subPolicyData.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/subscriber-session/" + subPolicyData.EntityData.SegmentPath
    subPolicyData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subPolicyData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subPolicyData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subPolicyData.EntityData.Children = types.NewOrderedMap()
    subPolicyData.EntityData.Leafs = types.NewOrderedMap()
    subPolicyData.EntityData.Leafs.Append("policy-epoch", types.YLeaf{"PolicyEpoch", subPolicyData.PolicyEpoch})

    subPolicyData.EntityData.YListKeys = []string {}

    return &(subPolicyData.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionServiceInfo
// List of subscriber services associated to this
// session
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionServiceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ServiceName. The type is string.
    ServiceName interface{}

    // ServiceParams. The type is string.
    ServiceParams interface{}

    // ServiceType. The type is IedgeOperService.
    ServiceType interface{}

    // ServiceStatus. The type is IedgeOperServiceStatus.
    ServiceStatus interface{}

    // ServiceIdentifier. The type is interface{} with range: 0..4294967295.
    ServiceId interface{}

    // ServicePrepaid. The type is bool.
    ServicePrepaid interface{}
}

func (sessionServiceInfo *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionServiceInfo) GetEntityData() *types.CommonEntityData {
    sessionServiceInfo.EntityData.YFilter = sessionServiceInfo.YFilter
    sessionServiceInfo.EntityData.YangName = "session-service-info"
    sessionServiceInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionServiceInfo.EntityData.ParentYangName = "subscriber-session"
    sessionServiceInfo.EntityData.SegmentPath = "session-service-info" + types.AddNoKeyToken(sessionServiceInfo)
    sessionServiceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/subscriber-session/" + sessionServiceInfo.EntityData.SegmentPath
    sessionServiceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionServiceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionServiceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionServiceInfo.EntityData.Children = types.NewOrderedMap()
    sessionServiceInfo.EntityData.Leafs = types.NewOrderedMap()
    sessionServiceInfo.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", sessionServiceInfo.ServiceName})
    sessionServiceInfo.EntityData.Leafs.Append("service-params", types.YLeaf{"ServiceParams", sessionServiceInfo.ServiceParams})
    sessionServiceInfo.EntityData.Leafs.Append("service-type", types.YLeaf{"ServiceType", sessionServiceInfo.ServiceType})
    sessionServiceInfo.EntityData.Leafs.Append("service-status", types.YLeaf{"ServiceStatus", sessionServiceInfo.ServiceStatus})
    sessionServiceInfo.EntityData.Leafs.Append("service-id", types.YLeaf{"ServiceId", sessionServiceInfo.ServiceId})
    sessionServiceInfo.EntityData.Leafs.Append("service-prepaid", types.YLeaf{"ServicePrepaid", sessionServiceInfo.ServicePrepaid})

    sessionServiceInfo.EntityData.YListKeys = []string {}

    return &(sessionServiceInfo.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionChangeOfAuthorization
// Subscriber change of authorization information
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Coa Request Acked. The type is bool.
    RequestAcked interface{}

    // Request time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    RequestTimeEpoch interface{}

    // Reply time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ReplyTimeEpoch interface{}
}

func (sessionChangeOfAuthorization *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SubscriberSessions_SubscriberSession_SessionChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    sessionChangeOfAuthorization.EntityData.YFilter = sessionChangeOfAuthorization.YFilter
    sessionChangeOfAuthorization.EntityData.YangName = "session-change-of-authorization"
    sessionChangeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    sessionChangeOfAuthorization.EntityData.ParentYangName = "subscriber-session"
    sessionChangeOfAuthorization.EntityData.SegmentPath = "session-change-of-authorization" + types.AddNoKeyToken(sessionChangeOfAuthorization)
    sessionChangeOfAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/subscriber-sessions/subscriber-session/" + sessionChangeOfAuthorization.EntityData.SegmentPath
    sessionChangeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionChangeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionChangeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionChangeOfAuthorization.EntityData.Children = types.NewOrderedMap()
    sessionChangeOfAuthorization.EntityData.Leafs = types.NewOrderedMap()
    sessionChangeOfAuthorization.EntityData.Leafs.Append("request-acked", types.YLeaf{"RequestAcked", sessionChangeOfAuthorization.RequestAcked})
    sessionChangeOfAuthorization.EntityData.Leafs.Append("request-time-epoch", types.YLeaf{"RequestTimeEpoch", sessionChangeOfAuthorization.RequestTimeEpoch})
    sessionChangeOfAuthorization.EntityData.Leafs.Append("reply-time-epoch", types.YLeaf{"ReplyTimeEpoch", sessionChangeOfAuthorization.ReplyTimeEpoch})

    sessionChangeOfAuthorization.EntityData.YListKeys = []string {}

    return &(sessionChangeOfAuthorization.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries
// Summary information filtered by IPv4
// address and VRF
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address and VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary.
    Ipv4AddressVrfSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
}

func (ipv4AddressVrfSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummaries.EntityData.YFilter = ipv4AddressVrfSummaries.YFilter
    ipv4AddressVrfSummaries.EntityData.YangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummaries.EntityData.ParentYangName = "srg-role"
    ipv4AddressVrfSummaries.EntityData.SegmentPath = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + ipv4AddressVrfSummaries.EntityData.SegmentPath
    ipv4AddressVrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4AddressVrfSummaries.EntityData.Children.Append("ipv4-address-vrf-summary", types.YChild{"Ipv4AddressVrfSummary", nil})
    for i := range ipv4AddressVrfSummaries.Ipv4AddressVrfSummary {
        types.SetYListKey(ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i], i)
        ipv4AddressVrfSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i]), types.YChild{"Ipv4AddressVrfSummary", ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i]})
    }
    ipv4AddressVrfSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4AddressVrfSummaries.EntityData.YListKeys = []string {}

    return &(ipv4AddressVrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
// IPv4 address and VRF summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Subscriber IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
}

func (ipv4AddressVrfSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummary.EntityData.YFilter = ipv4AddressVrfSummary.YFilter
    ipv4AddressVrfSummary.EntityData.YangName = "ipv4-address-vrf-summary"
    ipv4AddressVrfSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummary.EntityData.ParentYangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummary.EntityData.SegmentPath = "ipv4-address-vrf-summary" + types.AddNoKeyToken(ipv4AddressVrfSummary)
    ipv4AddressVrfSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/" + ipv4AddressVrfSummary.EntityData.SegmentPath
    ipv4AddressVrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummary.EntityData.Children = types.NewOrderedMap()
    ipv4AddressVrfSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &ipv4AddressVrfSummary.StateXr})
    ipv4AddressVrfSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &ipv4AddressVrfSummary.AddressFamilyXr})
    ipv4AddressVrfSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressVrfSummary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressVrfSummary.VrfName})
    ipv4AddressVrfSummary.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4AddressVrfSummary.Address})

    ipv4AddressVrfSummary.EntityData.YListKeys = []string {}

    return &(ipv4AddressVrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary
// Subscriber session summary information
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr
}

func (srgSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary) GetEntityData() *types.CommonEntityData {
    srgSummary.EntityData.YFilter = srgSummary.YFilter
    srgSummary.EntityData.YangName = "srg-summary"
    srgSummary.EntityData.BundleName = "cisco_ios_xr"
    srgSummary.EntityData.ParentYangName = "srg-role"
    srgSummary.EntityData.SegmentPath = "srg-summary"
    srgSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + srgSummary.EntityData.SegmentPath
    srgSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgSummary.EntityData.Children = types.NewOrderedMap()
    srgSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &srgSummary.StateXr})
    srgSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &srgSummary.AddressFamilyXr})
    srgSummary.EntityData.Leafs = types.NewOrderedMap()

    srgSummary.EntityData.YListKeys = []string {}

    return &(srgSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "srg-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "srg-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_SrgSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/srg-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries
// Summary information filtered by access
// interface
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary.
    AccessInterfaceSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary
}

func (accessInterfaceSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummaries.EntityData.YFilter = accessInterfaceSummaries.YFilter
    accessInterfaceSummaries.EntityData.YangName = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummaries.EntityData.ParentYangName = "srg-role"
    accessInterfaceSummaries.EntityData.SegmentPath = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + accessInterfaceSummaries.EntityData.SegmentPath
    accessInterfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummaries.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSummaries.EntityData.Children.Append("access-interface-summary", types.YChild{"AccessInterfaceSummary", nil})
    for i := range accessInterfaceSummaries.AccessInterfaceSummary {
        accessInterfaceSummaries.EntityData.Children.Append(types.GetSegmentPath(accessInterfaceSummaries.AccessInterfaceSummary[i]), types.YChild{"AccessInterfaceSummary", accessInterfaceSummaries.AccessInterfaceSummary[i]})
    }
    accessInterfaceSummaries.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaceSummaries.EntityData.YListKeys = []string {}

    return &(accessInterfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary
// Access interface summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
}

func (accessInterfaceSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummary.EntityData.YFilter = accessInterfaceSummary.YFilter
    accessInterfaceSummary.EntityData.YangName = "access-interface-summary"
    accessInterfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummary.EntityData.ParentYangName = "access-interface-summaries"
    accessInterfaceSummary.EntityData.SegmentPath = "access-interface-summary" + types.AddKeyToken(accessInterfaceSummary.InterfaceName, "interface-name")
    accessInterfaceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/" + accessInterfaceSummary.EntityData.SegmentPath
    accessInterfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummary.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &accessInterfaceSummary.StateXr})
    accessInterfaceSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &accessInterfaceSummary.AddressFamilyXr})
    accessInterfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterfaceSummary.InterfaceName})

    accessInterfaceSummary.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "access-interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "access-interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/access-interface-summaries/access-interface-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries
// Summary information filtered by address
// family
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address family summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary.
    AddressFamilySummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary
}

func (addressFamilySummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries) GetEntityData() *types.CommonEntityData {
    addressFamilySummaries.EntityData.YFilter = addressFamilySummaries.YFilter
    addressFamilySummaries.EntityData.YangName = "address-family-summaries"
    addressFamilySummaries.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummaries.EntityData.ParentYangName = "srg-role"
    addressFamilySummaries.EntityData.SegmentPath = "address-family-summaries"
    addressFamilySummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + addressFamilySummaries.EntityData.SegmentPath
    addressFamilySummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummaries.EntityData.Children = types.NewOrderedMap()
    addressFamilySummaries.EntityData.Children.Append("address-family-summary", types.YChild{"AddressFamilySummary", nil})
    for i := range addressFamilySummaries.AddressFamilySummary {
        addressFamilySummaries.EntityData.Children.Append(types.GetSegmentPath(addressFamilySummaries.AddressFamilySummary[i]), types.YChild{"AddressFamilySummary", addressFamilySummaries.AddressFamilySummary[i]})
    }
    addressFamilySummaries.EntityData.Leafs = types.NewOrderedMap()

    addressFamilySummaries.EntityData.YListKeys = []string {}

    return &(addressFamilySummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address family. The type is
    // SubscriberAddressFamilyFilterFlag.
    AddressFamily interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
}

func (addressFamilySummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary) GetEntityData() *types.CommonEntityData {
    addressFamilySummary.EntityData.YFilter = addressFamilySummary.YFilter
    addressFamilySummary.EntityData.YangName = "address-family-summary"
    addressFamilySummary.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummary.EntityData.ParentYangName = "address-family-summaries"
    addressFamilySummary.EntityData.SegmentPath = "address-family-summary" + types.AddKeyToken(addressFamilySummary.AddressFamily, "address-family")
    addressFamilySummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/" + addressFamilySummary.EntityData.SegmentPath
    addressFamilySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummary.EntityData.Children = types.NewOrderedMap()
    addressFamilySummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &addressFamilySummary.StateXr})
    addressFamilySummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &addressFamilySummary.AddressFamilyXr})
    addressFamilySummary.EntityData.Leafs = types.NewOrderedMap()
    addressFamilySummary.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", addressFamilySummary.AddressFamily})

    addressFamilySummary.EntityData.YListKeys = []string {"AddressFamily"}

    return &(addressFamilySummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "address-family-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "address-family-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/address-family-summaries/address-family-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries
// Summary information filtered by subscriber
// IPv4 address
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary.
    Ipv4AddressSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary
}

func (ipv4AddressSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummaries.EntityData.YFilter = ipv4AddressSummaries.YFilter
    ipv4AddressSummaries.EntityData.YangName = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummaries.EntityData.ParentYangName = "srg-role"
    ipv4AddressSummaries.EntityData.SegmentPath = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + ipv4AddressSummaries.EntityData.SegmentPath
    ipv4AddressSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4AddressSummaries.EntityData.Children.Append("ipv4-address-summary", types.YChild{"Ipv4AddressSummary", nil})
    for i := range ipv4AddressSummaries.Ipv4AddressSummary {
        ipv4AddressSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4AddressSummaries.Ipv4AddressSummary[i]), types.YChild{"Ipv4AddressSummary", ipv4AddressSummaries.Ipv4AddressSummary[i]})
    }
    ipv4AddressSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4AddressSummaries.EntityData.YListKeys = []string {}

    return &(ipv4AddressSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary
// IPv4 address summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber IPv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
}

func (ipv4AddressSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummary.EntityData.YFilter = ipv4AddressSummary.YFilter
    ipv4AddressSummary.EntityData.YangName = "ipv4-address-summary"
    ipv4AddressSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummary.EntityData.ParentYangName = "ipv4-address-summaries"
    ipv4AddressSummary.EntityData.SegmentPath = "ipv4-address-summary" + types.AddKeyToken(ipv4AddressSummary.Address, "address")
    ipv4AddressSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/" + ipv4AddressSummary.EntityData.SegmentPath
    ipv4AddressSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummary.EntityData.Children = types.NewOrderedMap()
    ipv4AddressSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &ipv4AddressSummary.StateXr})
    ipv4AddressSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &ipv4AddressSummary.AddressFamilyXr})
    ipv4AddressSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressSummary.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4AddressSummary.Address})

    ipv4AddressSummary.EntityData.YListKeys = []string {"Address"}

    return &(ipv4AddressSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries
// Summary information filtered by VRF
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary.
    VrfSummary []*Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary
}

func (vrfSummaries *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries) GetEntityData() *types.CommonEntityData {
    vrfSummaries.EntityData.YFilter = vrfSummaries.YFilter
    vrfSummaries.EntityData.YangName = "vrf-summaries"
    vrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    vrfSummaries.EntityData.ParentYangName = "srg-role"
    vrfSummaries.EntityData.SegmentPath = "vrf-summaries"
    vrfSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/" + vrfSummaries.EntityData.SegmentPath
    vrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummaries.EntityData.Children = types.NewOrderedMap()
    vrfSummaries.EntityData.Children.Append("vrf-summary", types.YChild{"VrfSummary", nil})
    for i := range vrfSummaries.VrfSummary {
        vrfSummaries.EntityData.Children.Append(types.GetSegmentPath(vrfSummaries.VrfSummary[i]), types.YChild{"VrfSummary", vrfSummaries.VrfSummary[i]})
    }
    vrfSummaries.EntityData.Leafs = types.NewOrderedMap()

    vrfSummaries.EntityData.YListKeys = []string {}

    return &(vrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary
// VRF summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr
}

func (vrfSummary *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "vrf-summaries"
    vrfSummary.EntityData.SegmentPath = "vrf-summary" + types.AddKeyToken(vrfSummary.VrfName, "vrf-name")
    vrfSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/" + vrfSummary.EntityData.SegmentPath
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = types.NewOrderedMap()
    vrfSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &vrfSummary.StateXr})
    vrfSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &vrfSummary.AddressFamilyXr})
    vrfSummary.EntityData.Leafs = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfSummary.VrfName})

    vrfSummary.EntityData.YListKeys = []string {"VrfName"}

    return &(vrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_SrgRoles_SrgRole_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/srg-roles/srg-role/vrf-summaries/vrf-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries
// Summary information filtered by authorization
// state
type Subscriber_Session_Nodes_Node_AuthorSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authorization summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary.
    AuthorSummary []*Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary
}

func (authorSummaries *Subscriber_Session_Nodes_Node_AuthorSummaries) GetEntityData() *types.CommonEntityData {
    authorSummaries.EntityData.YFilter = authorSummaries.YFilter
    authorSummaries.EntityData.YangName = "author-summaries"
    authorSummaries.EntityData.BundleName = "cisco_ios_xr"
    authorSummaries.EntityData.ParentYangName = "node"
    authorSummaries.EntityData.SegmentPath = "author-summaries"
    authorSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + authorSummaries.EntityData.SegmentPath
    authorSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummaries.EntityData.Children = types.NewOrderedMap()
    authorSummaries.EntityData.Children.Append("author-summary", types.YChild{"AuthorSummary", nil})
    for i := range authorSummaries.AuthorSummary {
        authorSummaries.EntityData.Children.Append(types.GetSegmentPath(authorSummaries.AuthorSummary[i]), types.YChild{"AuthorSummary", authorSummaries.AuthorSummary[i]})
    }
    authorSummaries.EntityData.Leafs = types.NewOrderedMap()

    authorSummaries.EntityData.YListKeys = []string {}

    return &(authorSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary
// authorization summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Authorization state. The type is
    // SubscriberAuthorStateFilterFlag.
    AuthorState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr
}

func (authorSummary *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary) GetEntityData() *types.CommonEntityData {
    authorSummary.EntityData.YFilter = authorSummary.YFilter
    authorSummary.EntityData.YangName = "author-summary"
    authorSummary.EntityData.BundleName = "cisco_ios_xr"
    authorSummary.EntityData.ParentYangName = "author-summaries"
    authorSummary.EntityData.SegmentPath = "author-summary" + types.AddKeyToken(authorSummary.AuthorState, "author-state")
    authorSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/" + authorSummary.EntityData.SegmentPath
    authorSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummary.EntityData.Children = types.NewOrderedMap()
    authorSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &authorSummary.StateXr})
    authorSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &authorSummary.AddressFamilyXr})
    authorSummary.EntityData.Leafs = types.NewOrderedMap()
    authorSummary.EntityData.Leafs.Append("author-state", types.YLeaf{"AuthorState", authorSummary.AuthorState})

    authorSummary.EntityData.YListKeys = []string {"AuthorState"}

    return &(authorSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "author-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "author-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/author-summaries/author-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary
// Subscriber session summary information
type Subscriber_Session_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Summary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr
}

func (summary *Subscriber_Session_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &summary.StateXr})
    summary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &summary.AddressFamilyXr})
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Summary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Summary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries
// Summary information filtered by MAC address
type Subscriber_Session_Nodes_Node_MacSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_MacSummaries_MacSummary.
    MacSummary []*Subscriber_Session_Nodes_Node_MacSummaries_MacSummary
}

func (macSummaries *Subscriber_Session_Nodes_Node_MacSummaries) GetEntityData() *types.CommonEntityData {
    macSummaries.EntityData.YFilter = macSummaries.YFilter
    macSummaries.EntityData.YangName = "mac-summaries"
    macSummaries.EntityData.BundleName = "cisco_ios_xr"
    macSummaries.EntityData.ParentYangName = "node"
    macSummaries.EntityData.SegmentPath = "mac-summaries"
    macSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + macSummaries.EntityData.SegmentPath
    macSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummaries.EntityData.Children = types.NewOrderedMap()
    macSummaries.EntityData.Children.Append("mac-summary", types.YChild{"MacSummary", nil})
    for i := range macSummaries.MacSummary {
        macSummaries.EntityData.Children.Append(types.GetSegmentPath(macSummaries.MacSummary[i]), types.YChild{"MacSummary", macSummaries.MacSummary[i]})
    }
    macSummaries.EntityData.Leafs = types.NewOrderedMap()

    macSummaries.EntityData.YListKeys = []string {}

    return &(macSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary
// MAC address summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr
}

func (macSummary *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary) GetEntityData() *types.CommonEntityData {
    macSummary.EntityData.YFilter = macSummary.YFilter
    macSummary.EntityData.YangName = "mac-summary"
    macSummary.EntityData.BundleName = "cisco_ios_xr"
    macSummary.EntityData.ParentYangName = "mac-summaries"
    macSummary.EntityData.SegmentPath = "mac-summary" + types.AddKeyToken(macSummary.MacAddress, "mac-address")
    macSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/" + macSummary.EntityData.SegmentPath
    macSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummary.EntityData.Children = types.NewOrderedMap()
    macSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &macSummary.StateXr})
    macSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &macSummary.AddressFamilyXr})
    macSummary.EntityData.Leafs = types.NewOrderedMap()
    macSummary.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macSummary.MacAddress})

    macSummary.EntityData.YListKeys = []string {"MacAddress"}

    return &(macSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "mac-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "mac-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/mac-summaries/mac-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries
// Summary information filtered by interface
type Subscriber_Session_Nodes_Node_InterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary.
    InterfaceSummary []*Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary
}

func (interfaceSummaries *Subscriber_Session_Nodes_Node_InterfaceSummaries) GetEntityData() *types.CommonEntityData {
    interfaceSummaries.EntityData.YFilter = interfaceSummaries.YFilter
    interfaceSummaries.EntityData.YangName = "interface-summaries"
    interfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummaries.EntityData.ParentYangName = "node"
    interfaceSummaries.EntityData.SegmentPath = "interface-summaries"
    interfaceSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + interfaceSummaries.EntityData.SegmentPath
    interfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummaries.EntityData.Children = types.NewOrderedMap()
    interfaceSummaries.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", nil})
    for i := range interfaceSummaries.InterfaceSummary {
        interfaceSummaries.EntityData.Children.Append(types.GetSegmentPath(interfaceSummaries.InterfaceSummary[i]), types.YChild{"InterfaceSummary", interfaceSummaries.InterfaceSummary[i]})
    }
    interfaceSummaries.EntityData.Leafs = types.NewOrderedMap()

    interfaceSummaries.EntityData.YListKeys = []string {}

    return &(interfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary
// Interface summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
}

func (interfaceSummary *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "interface-summaries"
    interfaceSummary.EntityData.SegmentPath = "interface-summary" + types.AddKeyToken(interfaceSummary.InterfaceName, "interface-name")
    interfaceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/" + interfaceSummary.EntityData.SegmentPath
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &interfaceSummary.StateXr})
    interfaceSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &interfaceSummary.AddressFamilyXr})
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName})

    interfaceSummary.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/interface-summaries/interface-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries
// Summary information filtered by
// authentication state
type Subscriber_Session_Nodes_Node_AuthenticationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authentication summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary.
    AuthenticationSummary []*Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary
}

func (authenticationSummaries *Subscriber_Session_Nodes_Node_AuthenticationSummaries) GetEntityData() *types.CommonEntityData {
    authenticationSummaries.EntityData.YFilter = authenticationSummaries.YFilter
    authenticationSummaries.EntityData.YangName = "authentication-summaries"
    authenticationSummaries.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummaries.EntityData.ParentYangName = "node"
    authenticationSummaries.EntityData.SegmentPath = "authentication-summaries"
    authenticationSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + authenticationSummaries.EntityData.SegmentPath
    authenticationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummaries.EntityData.Children = types.NewOrderedMap()
    authenticationSummaries.EntityData.Children.Append("authentication-summary", types.YChild{"AuthenticationSummary", nil})
    for i := range authenticationSummaries.AuthenticationSummary {
        authenticationSummaries.EntityData.Children.Append(types.GetSegmentPath(authenticationSummaries.AuthenticationSummary[i]), types.YChild{"AuthenticationSummary", authenticationSummaries.AuthenticationSummary[i]})
    }
    authenticationSummaries.EntityData.Leafs = types.NewOrderedMap()

    authenticationSummaries.EntityData.YListKeys = []string {}

    return &(authenticationSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary
// authentication summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Authentication state. The type is
    // SubscriberAuthenStateFilterFlag.
    AuthenticationState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
}

func (authenticationSummary *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary) GetEntityData() *types.CommonEntityData {
    authenticationSummary.EntityData.YFilter = authenticationSummary.YFilter
    authenticationSummary.EntityData.YangName = "authentication-summary"
    authenticationSummary.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummary.EntityData.ParentYangName = "authentication-summaries"
    authenticationSummary.EntityData.SegmentPath = "authentication-summary" + types.AddKeyToken(authenticationSummary.AuthenticationState, "authentication-state")
    authenticationSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/" + authenticationSummary.EntityData.SegmentPath
    authenticationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummary.EntityData.Children = types.NewOrderedMap()
    authenticationSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &authenticationSummary.StateXr})
    authenticationSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &authenticationSummary.AddressFamilyXr})
    authenticationSummary.EntityData.Leafs = types.NewOrderedMap()
    authenticationSummary.EntityData.Leafs.Append("authentication-state", types.YLeaf{"AuthenticationState", authenticationSummary.AuthenticationState})

    authenticationSummary.EntityData.YListKeys = []string {"AuthenticationState"}

    return &(authenticationSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "authentication-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "authentication-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/authentication-summaries/authentication-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries
// Summary information filtered by session state
type Subscriber_Session_Nodes_Node_StateSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary. The type is slice of
    // Subscriber_Session_Nodes_Node_StateSummaries_StateSummary.
    StateSummary []*Subscriber_Session_Nodes_Node_StateSummaries_StateSummary
}

func (stateSummaries *Subscriber_Session_Nodes_Node_StateSummaries) GetEntityData() *types.CommonEntityData {
    stateSummaries.EntityData.YFilter = stateSummaries.YFilter
    stateSummaries.EntityData.YangName = "state-summaries"
    stateSummaries.EntityData.BundleName = "cisco_ios_xr"
    stateSummaries.EntityData.ParentYangName = "node"
    stateSummaries.EntityData.SegmentPath = "state-summaries"
    stateSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + stateSummaries.EntityData.SegmentPath
    stateSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummaries.EntityData.Children = types.NewOrderedMap()
    stateSummaries.EntityData.Children.Append("state-summary", types.YChild{"StateSummary", nil})
    for i := range stateSummaries.StateSummary {
        stateSummaries.EntityData.Children.Append(types.GetSegmentPath(stateSummaries.StateSummary[i]), types.YChild{"StateSummary", stateSummaries.StateSummary[i]})
    }
    stateSummaries.EntityData.Leafs = types.NewOrderedMap()

    stateSummaries.EntityData.YListKeys = []string {}

    return &(stateSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary
// State summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber state. The type is
    // SubscriberStateFilterFlag.
    State interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr
}

func (stateSummary *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary) GetEntityData() *types.CommonEntityData {
    stateSummary.EntityData.YFilter = stateSummary.YFilter
    stateSummary.EntityData.YangName = "state-summary"
    stateSummary.EntityData.BundleName = "cisco_ios_xr"
    stateSummary.EntityData.ParentYangName = "state-summaries"
    stateSummary.EntityData.SegmentPath = "state-summary" + types.AddKeyToken(stateSummary.State, "state")
    stateSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/" + stateSummary.EntityData.SegmentPath
    stateSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummary.EntityData.Children = types.NewOrderedMap()
    stateSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &stateSummary.StateXr})
    stateSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &stateSummary.AddressFamilyXr})
    stateSummary.EntityData.Leafs = types.NewOrderedMap()
    stateSummary.EntityData.Leafs.Append("state", types.YLeaf{"State", stateSummary.State})

    stateSummary.EntityData.YListKeys = []string {"State"}

    return &(stateSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "state-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "state-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/state-summaries/state-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries
// Summary information filtered by IPv4 address
// and VRF
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address and VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary.
    Ipv4AddressVrfSummary []*Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
}

func (ipv4AddressVrfSummaries *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummaries.EntityData.YFilter = ipv4AddressVrfSummaries.YFilter
    ipv4AddressVrfSummaries.EntityData.YangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummaries.EntityData.ParentYangName = "node"
    ipv4AddressVrfSummaries.EntityData.SegmentPath = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + ipv4AddressVrfSummaries.EntityData.SegmentPath
    ipv4AddressVrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4AddressVrfSummaries.EntityData.Children.Append("ipv4-address-vrf-summary", types.YChild{"Ipv4AddressVrfSummary", nil})
    for i := range ipv4AddressVrfSummaries.Ipv4AddressVrfSummary {
        types.SetYListKey(ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i], i)
        ipv4AddressVrfSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i]), types.YChild{"Ipv4AddressVrfSummary", ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i]})
    }
    ipv4AddressVrfSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4AddressVrfSummaries.EntityData.YListKeys = []string {}

    return &(ipv4AddressVrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
// IPv4 address and VRF summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Subscriber IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
}

func (ipv4AddressVrfSummary *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummary.EntityData.YFilter = ipv4AddressVrfSummary.YFilter
    ipv4AddressVrfSummary.EntityData.YangName = "ipv4-address-vrf-summary"
    ipv4AddressVrfSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummary.EntityData.ParentYangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummary.EntityData.SegmentPath = "ipv4-address-vrf-summary" + types.AddNoKeyToken(ipv4AddressVrfSummary)
    ipv4AddressVrfSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/" + ipv4AddressVrfSummary.EntityData.SegmentPath
    ipv4AddressVrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummary.EntityData.Children = types.NewOrderedMap()
    ipv4AddressVrfSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &ipv4AddressVrfSummary.StateXr})
    ipv4AddressVrfSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &ipv4AddressVrfSummary.AddressFamilyXr})
    ipv4AddressVrfSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressVrfSummary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressVrfSummary.VrfName})
    ipv4AddressVrfSummary.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4AddressVrfSummary.Address})

    ipv4AddressVrfSummary.EntityData.YListKeys = []string {}

    return &(ipv4AddressVrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-vrf-summaries/ipv4-address-vrf-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries
// Summary information filtered by address
// family
type Subscriber_Session_Nodes_Node_AddressFamilySummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address family summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary.
    AddressFamilySummary []*Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary
}

func (addressFamilySummaries *Subscriber_Session_Nodes_Node_AddressFamilySummaries) GetEntityData() *types.CommonEntityData {
    addressFamilySummaries.EntityData.YFilter = addressFamilySummaries.YFilter
    addressFamilySummaries.EntityData.YangName = "address-family-summaries"
    addressFamilySummaries.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummaries.EntityData.ParentYangName = "node"
    addressFamilySummaries.EntityData.SegmentPath = "address-family-summaries"
    addressFamilySummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + addressFamilySummaries.EntityData.SegmentPath
    addressFamilySummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummaries.EntityData.Children = types.NewOrderedMap()
    addressFamilySummaries.EntityData.Children.Append("address-family-summary", types.YChild{"AddressFamilySummary", nil})
    for i := range addressFamilySummaries.AddressFamilySummary {
        addressFamilySummaries.EntityData.Children.Append(types.GetSegmentPath(addressFamilySummaries.AddressFamilySummary[i]), types.YChild{"AddressFamilySummary", addressFamilySummaries.AddressFamilySummary[i]})
    }
    addressFamilySummaries.EntityData.Leafs = types.NewOrderedMap()

    addressFamilySummaries.EntityData.YListKeys = []string {}

    return &(addressFamilySummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary
// Address family summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address family. The type is
    // SubscriberAddressFamilyFilterFlag.
    AddressFamily interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
}

func (addressFamilySummary *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary) GetEntityData() *types.CommonEntityData {
    addressFamilySummary.EntityData.YFilter = addressFamilySummary.YFilter
    addressFamilySummary.EntityData.YangName = "address-family-summary"
    addressFamilySummary.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummary.EntityData.ParentYangName = "address-family-summaries"
    addressFamilySummary.EntityData.SegmentPath = "address-family-summary" + types.AddKeyToken(addressFamilySummary.AddressFamily, "address-family")
    addressFamilySummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/" + addressFamilySummary.EntityData.SegmentPath
    addressFamilySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummary.EntityData.Children = types.NewOrderedMap()
    addressFamilySummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &addressFamilySummary.StateXr})
    addressFamilySummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &addressFamilySummary.AddressFamilyXr})
    addressFamilySummary.EntityData.Leafs = types.NewOrderedMap()
    addressFamilySummary.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", addressFamilySummary.AddressFamily})

    addressFamilySummary.EntityData.YListKeys = []string {"AddressFamily"}

    return &(addressFamilySummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "address-family-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "address-family-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/address-family-summaries/address-family-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries
// Summary information filtered by username
type Subscriber_Session_Nodes_Node_UsernameSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username summary. The type is slice of
    // Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary.
    UsernameSummary []*Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary
}

func (usernameSummaries *Subscriber_Session_Nodes_Node_UsernameSummaries) GetEntityData() *types.CommonEntityData {
    usernameSummaries.EntityData.YFilter = usernameSummaries.YFilter
    usernameSummaries.EntityData.YangName = "username-summaries"
    usernameSummaries.EntityData.BundleName = "cisco_ios_xr"
    usernameSummaries.EntityData.ParentYangName = "node"
    usernameSummaries.EntityData.SegmentPath = "username-summaries"
    usernameSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + usernameSummaries.EntityData.SegmentPath
    usernameSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummaries.EntityData.Children = types.NewOrderedMap()
    usernameSummaries.EntityData.Children.Append("username-summary", types.YChild{"UsernameSummary", nil})
    for i := range usernameSummaries.UsernameSummary {
        usernameSummaries.EntityData.Children.Append(types.GetSegmentPath(usernameSummaries.UsernameSummary[i]), types.YChild{"UsernameSummary", usernameSummaries.UsernameSummary[i]})
    }
    usernameSummaries.EntityData.Leafs = types.NewOrderedMap()

    usernameSummaries.EntityData.YListKeys = []string {}

    return &(usernameSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary
// Username summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber username. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Username interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr
}

func (usernameSummary *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary) GetEntityData() *types.CommonEntityData {
    usernameSummary.EntityData.YFilter = usernameSummary.YFilter
    usernameSummary.EntityData.YangName = "username-summary"
    usernameSummary.EntityData.BundleName = "cisco_ios_xr"
    usernameSummary.EntityData.ParentYangName = "username-summaries"
    usernameSummary.EntityData.SegmentPath = "username-summary" + types.AddKeyToken(usernameSummary.Username, "username")
    usernameSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/" + usernameSummary.EntityData.SegmentPath
    usernameSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummary.EntityData.Children = types.NewOrderedMap()
    usernameSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &usernameSummary.StateXr})
    usernameSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &usernameSummary.AddressFamilyXr})
    usernameSummary.EntityData.Leafs = types.NewOrderedMap()
    usernameSummary.EntityData.Leafs.Append("username", types.YLeaf{"Username", usernameSummary.Username})

    usernameSummary.EntityData.YListKeys = []string {"Username"}

    return &(usernameSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "username-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "username-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/username-summaries/username-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries
// Summary information filtered by access
// interface
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary.
    AccessInterfaceSummary []*Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary
}

func (accessInterfaceSummaries *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummaries.EntityData.YFilter = accessInterfaceSummaries.YFilter
    accessInterfaceSummaries.EntityData.YangName = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummaries.EntityData.ParentYangName = "node"
    accessInterfaceSummaries.EntityData.SegmentPath = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + accessInterfaceSummaries.EntityData.SegmentPath
    accessInterfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummaries.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSummaries.EntityData.Children.Append("access-interface-summary", types.YChild{"AccessInterfaceSummary", nil})
    for i := range accessInterfaceSummaries.AccessInterfaceSummary {
        accessInterfaceSummaries.EntityData.Children.Append(types.GetSegmentPath(accessInterfaceSummaries.AccessInterfaceSummary[i]), types.YChild{"AccessInterfaceSummary", accessInterfaceSummaries.AccessInterfaceSummary[i]})
    }
    accessInterfaceSummaries.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaceSummaries.EntityData.YListKeys = []string {}

    return &(accessInterfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary
// Access interface summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
}

func (accessInterfaceSummary *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummary.EntityData.YFilter = accessInterfaceSummary.YFilter
    accessInterfaceSummary.EntityData.YangName = "access-interface-summary"
    accessInterfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummary.EntityData.ParentYangName = "access-interface-summaries"
    accessInterfaceSummary.EntityData.SegmentPath = "access-interface-summary" + types.AddKeyToken(accessInterfaceSummary.InterfaceName, "interface-name")
    accessInterfaceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/" + accessInterfaceSummary.EntityData.SegmentPath
    accessInterfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummary.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &accessInterfaceSummary.StateXr})
    accessInterfaceSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &accessInterfaceSummary.AddressFamilyXr})
    accessInterfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    accessInterfaceSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterfaceSummary.InterfaceName})

    accessInterfaceSummary.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "access-interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "access-interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/access-interface-summaries/access-interface-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries
// Summary information filtered by subscriber
// IPv4 address
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary.
    Ipv4AddressSummary []*Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary
}

func (ipv4AddressSummaries *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummaries.EntityData.YFilter = ipv4AddressSummaries.YFilter
    ipv4AddressSummaries.EntityData.YangName = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummaries.EntityData.ParentYangName = "node"
    ipv4AddressSummaries.EntityData.SegmentPath = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + ipv4AddressSummaries.EntityData.SegmentPath
    ipv4AddressSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4AddressSummaries.EntityData.Children.Append("ipv4-address-summary", types.YChild{"Ipv4AddressSummary", nil})
    for i := range ipv4AddressSummaries.Ipv4AddressSummary {
        ipv4AddressSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4AddressSummaries.Ipv4AddressSummary[i]), types.YChild{"Ipv4AddressSummary", ipv4AddressSummaries.Ipv4AddressSummary[i]})
    }
    ipv4AddressSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4AddressSummaries.EntityData.YListKeys = []string {}

    return &(ipv4AddressSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary
// IPv4 address summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber IPv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
}

func (ipv4AddressSummary *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummary.EntityData.YFilter = ipv4AddressSummary.YFilter
    ipv4AddressSummary.EntityData.YangName = "ipv4-address-summary"
    ipv4AddressSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummary.EntityData.ParentYangName = "ipv4-address-summaries"
    ipv4AddressSummary.EntityData.SegmentPath = "ipv4-address-summary" + types.AddKeyToken(ipv4AddressSummary.Address, "address")
    ipv4AddressSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/" + ipv4AddressSummary.EntityData.SegmentPath
    ipv4AddressSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummary.EntityData.Children = types.NewOrderedMap()
    ipv4AddressSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &ipv4AddressSummary.StateXr})
    ipv4AddressSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &ipv4AddressSummary.AddressFamilyXr})
    ipv4AddressSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressSummary.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4AddressSummary.Address})

    ipv4AddressSummary.EntityData.YListKeys = []string {"Address"}

    return &(ipv4AddressSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/ipv4-address-summaries/ipv4-address-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries
// Summary information filtered by VRF
type Subscriber_Session_Nodes_Node_VrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary.
    VrfSummary []*Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary
}

func (vrfSummaries *Subscriber_Session_Nodes_Node_VrfSummaries) GetEntityData() *types.CommonEntityData {
    vrfSummaries.EntityData.YFilter = vrfSummaries.YFilter
    vrfSummaries.EntityData.YangName = "vrf-summaries"
    vrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    vrfSummaries.EntityData.ParentYangName = "node"
    vrfSummaries.EntityData.SegmentPath = "vrf-summaries"
    vrfSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + vrfSummaries.EntityData.SegmentPath
    vrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummaries.EntityData.Children = types.NewOrderedMap()
    vrfSummaries.EntityData.Children.Append("vrf-summary", types.YChild{"VrfSummary", nil})
    for i := range vrfSummaries.VrfSummary {
        vrfSummaries.EntityData.Children.Append(types.GetSegmentPath(vrfSummaries.VrfSummary[i]), types.YChild{"VrfSummary", vrfSummaries.VrfSummary[i]})
    }
    vrfSummaries.EntityData.Leafs = types.NewOrderedMap()

    vrfSummaries.EntityData.YListKeys = []string {}

    return &(vrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary
// VRF summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr
}

func (vrfSummary *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "vrf-summaries"
    vrfSummary.EntityData.SegmentPath = "vrf-summary" + types.AddKeyToken(vrfSummary.VrfName, "vrf-name")
    vrfSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/" + vrfSummary.EntityData.SegmentPath
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = types.NewOrderedMap()
    vrfSummary.EntityData.Children.Append("state-xr", types.YChild{"StateXr", &vrfSummary.StateXr})
    vrfSummary.EntityData.Children.Append("address-family-xr", types.YChild{"AddressFamilyXr", &vrfSummary.AddressFamilyXr})
    vrfSummary.EntityData.Leafs = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfSummary.VrfName})

    vrfSummary.EntityData.YListKeys = []string {"VrfName"}

    return &(vrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/" + stateXr.EntityData.SegmentPath
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = types.NewOrderedMap()
    stateXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &stateXr.Pppoe})
    stateXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp})
    stateXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket})
    stateXr.EntityData.Leafs = types.NewOrderedMap()

    stateXr.EntityData.YListKeys = []string {}

    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/state-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", pppoe.InitializedSessions})
    pppoe.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions})
    pppoe.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions})
    pppoe.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions})
    pppoe.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", pppoe.IdleSessions})
    pppoe.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions})
    pppoe.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", pppoe.EndSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/state-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/state-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("initialized-sessions", types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connecting-sessions", types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("connected-sessions", types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("activated-sessions", types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("disconnecting-sessions", types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("end-sessions", types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/" + addressFamilyXr.EntityData.SegmentPath
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = types.NewOrderedMap()
    addressFamilyXr.EntityData.Children.Append("pppoe", types.YChild{"Pppoe", &addressFamilyXr.Pppoe})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-dhcp", types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp})
    addressFamilyXr.EntityData.Children.Append("ip-subscriber-packet", types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket})
    addressFamilyXr.EntityData.Leafs = types.NewOrderedMap()

    addressFamilyXr.EntityData.YListKeys = []string {}

    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/address-family-xr/" + pppoe.EntityData.SegmentPath
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = types.NewOrderedMap()
    pppoe.EntityData.Leafs = types.NewOrderedMap()
    pppoe.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", pppoe.InProgressSessions})
    pppoe.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions})
    pppoe.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions})
    pppoe.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions})
    pppoe.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", pppoe.DualUpSessions})
    pppoe.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", pppoe.LacSessions})

    pppoe.EntityData.YListKeys = []string {}

    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/address-family-xr/" + ipSubscriberDhcp.EntityData.SegmentPath
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberDhcp.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions})
    ipSubscriberDhcp.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions})

    ipSubscriberDhcp.EntityData.YListKeys = []string {}

    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/vrf-summaries/vrf-summary/address-family-xr/" + ipSubscriberPacket.EntityData.SegmentPath
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriberPacket.EntityData.Leafs.Append("in-progress-sessions", types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv4-only-sessions", types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("ipv6-only-sessions", types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-part-up-sessions", types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("dual-up-sessions", types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions})
    ipSubscriberPacket.EntityData.Leafs.Append("lac-sessions", types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions})

    ipSubscriberPacket.EntityData.YListKeys = []string {}

    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions
// IP subscriber sessions
type Subscriber_Session_Nodes_Node_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session.
    Session []*Subscriber_Session_Nodes_Node_Sessions_Session
}

func (sessions *Subscriber_Session_Nodes_Node_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "node"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session
// Subscriber session information
type Subscriber_Session_Nodes_Node_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session ID. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    SessionId interface{}

    // Subscriber session type. The type is IedgeOperSession.
    SessionType interface{}

    // PPPoE sub type. The type is IedgePppSub.
    PppoeSubType interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Circuit ID. The type is string.
    CircuitId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // PPPoE LNS address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LnsAddress interface{}

    // PPPoE LAC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LacAddress interface{}

    // PPPoE LAC tunnel client authentication ID. The type is string.
    TunnelClientAuthenticationId interface{}

    // PPPoE LAC tunnel server authentication ID. The type is string.
    TunnelServerAuthenticationId interface{}

    // Session ip address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SessionIpAddress interface{}

    // Session IPv6 address. The type is string.
    SessionIpv6Address interface{}

    // Session IPv6 prefix. The type is string.
    SessionIpv6Prefix interface{}

    // Session delegated IPv6 prefix. The type is string.
    DelegatedIpv6Prefix interface{}

    // IPv6 Interface ID. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Ipv6InterfaceId interface{}

    // MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // NAS port. The type is string.
    NasPort interface{}

    // Username. The type is string.
    Username interface{}

    // Client Username. The type is string.
    Clientname interface{}

    // Formatted Username. The type is string.
    Formattedname interface{}

    // If true, session is authentic. The type is bool.
    IsSessionAuthentic interface{}

    // If true, session is authorized. The type is bool.
    IsSessionAuthor interface{}

    // Session state. The type is IedgeOperSessionState.
    SessionState interface{}

    // Session creation time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    SessionCreationTimeEpoch interface{}

    // Time when idle state change occurred in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    IdleStateChangeTime interface{}

    // Total session idle time (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    TotalSessionIdleTime interface{}

    // Access interface name associated with the session. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    AccessInterfaceName interface{}

    // Accounting information.
    Accounting Subscriber_Session_Nodes_Node_Sessions_Session_Accounting

    // Subscriber control policy applied to this session. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session_SubPolicyData.
    SubPolicyData []*Subscriber_Session_Nodes_Node_Sessions_Session_SubPolicyData

    // List of subscriber services associated to this session. The type is slice
    // of Subscriber_Session_Nodes_Node_Sessions_Session_SessionServiceInfo.
    SessionServiceInfo []*Subscriber_Session_Nodes_Node_Sessions_Session_SessionServiceInfo

    // Subscriber change of authorization information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session_SessionChangeOfAuthorization.
    SessionChangeOfAuthorization []*Subscriber_Session_Nodes_Node_Sessions_Session_SessionChangeOfAuthorization
}

func (session *Subscriber_Session_Nodes_Node_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("accounting", types.YChild{"Accounting", &session.Accounting})
    session.EntityData.Children.Append("sub-policy-data", types.YChild{"SubPolicyData", nil})
    for i := range session.SubPolicyData {
        types.SetYListKey(session.SubPolicyData[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.SubPolicyData[i]), types.YChild{"SubPolicyData", session.SubPolicyData[i]})
    }
    session.EntityData.Children.Append("session-service-info", types.YChild{"SessionServiceInfo", nil})
    for i := range session.SessionServiceInfo {
        types.SetYListKey(session.SessionServiceInfo[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.SessionServiceInfo[i]), types.YChild{"SessionServiceInfo", session.SessionServiceInfo[i]})
    }
    session.EntityData.Children.Append("session-change-of-authorization", types.YChild{"SessionChangeOfAuthorization", nil})
    for i := range session.SessionChangeOfAuthorization {
        types.SetYListKey(session.SessionChangeOfAuthorization[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.SessionChangeOfAuthorization[i]), types.YChild{"SessionChangeOfAuthorization", session.SessionChangeOfAuthorization[i]})
    }
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", session.SessionType})
    session.EntityData.Leafs.Append("pppoe-sub-type", types.YLeaf{"PppoeSubType", session.PppoeSubType})
    session.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", session.InterfaceName})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", session.CircuitId})
    session.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", session.RemoteId})
    session.EntityData.Leafs.Append("lns-address", types.YLeaf{"LnsAddress", session.LnsAddress})
    session.EntityData.Leafs.Append("lac-address", types.YLeaf{"LacAddress", session.LacAddress})
    session.EntityData.Leafs.Append("tunnel-client-authentication-id", types.YLeaf{"TunnelClientAuthenticationId", session.TunnelClientAuthenticationId})
    session.EntityData.Leafs.Append("tunnel-server-authentication-id", types.YLeaf{"TunnelServerAuthenticationId", session.TunnelServerAuthenticationId})
    session.EntityData.Leafs.Append("session-ip-address", types.YLeaf{"SessionIpAddress", session.SessionIpAddress})
    session.EntityData.Leafs.Append("session-ipv6-address", types.YLeaf{"SessionIpv6Address", session.SessionIpv6Address})
    session.EntityData.Leafs.Append("session-ipv6-prefix", types.YLeaf{"SessionIpv6Prefix", session.SessionIpv6Prefix})
    session.EntityData.Leafs.Append("delegated-ipv6-prefix", types.YLeaf{"DelegatedIpv6Prefix", session.DelegatedIpv6Prefix})
    session.EntityData.Leafs.Append("ipv6-interface-id", types.YLeaf{"Ipv6InterfaceId", session.Ipv6InterfaceId})
    session.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", session.MacAddress})
    session.EntityData.Leafs.Append("account-session-id", types.YLeaf{"AccountSessionId", session.AccountSessionId})
    session.EntityData.Leafs.Append("nas-port", types.YLeaf{"NasPort", session.NasPort})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("clientname", types.YLeaf{"Clientname", session.Clientname})
    session.EntityData.Leafs.Append("formattedname", types.YLeaf{"Formattedname", session.Formattedname})
    session.EntityData.Leafs.Append("is-session-authentic", types.YLeaf{"IsSessionAuthentic", session.IsSessionAuthentic})
    session.EntityData.Leafs.Append("is-session-author", types.YLeaf{"IsSessionAuthor", session.IsSessionAuthor})
    session.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", session.SessionState})
    session.EntityData.Leafs.Append("session-creation-time-epoch", types.YLeaf{"SessionCreationTimeEpoch", session.SessionCreationTimeEpoch})
    session.EntityData.Leafs.Append("idle-state-change-time", types.YLeaf{"IdleStateChangeTime", session.IdleStateChangeTime})
    session.EntityData.Leafs.Append("total-session-idle-time", types.YLeaf{"TotalSessionIdleTime", session.TotalSessionIdleTime})
    session.EntityData.Leafs.Append("access-interface-name", types.YLeaf{"AccessInterfaceName", session.AccessInterfaceName})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_Accounting
// Accounting information
type Subscriber_Session_Nodes_Node_Sessions_Session_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session_Accounting_AccountingSession.
    AccountingSession []*Subscriber_Session_Nodes_Node_Sessions_Session_Accounting_AccountingSession
}

func (accounting *Subscriber_Session_Nodes_Node_Sessions_Session_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "session"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/session/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("accounting-session", types.YChild{"AccountingSession", nil})
    for i := range accounting.AccountingSession {
        types.SetYListKey(accounting.AccountingSession[i], i)
        accounting.EntityData.Children.Append(types.GetSegmentPath(accounting.AccountingSession[i]), types.YChild{"AccountingSession", accounting.AccountingSession[i]})
    }
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_Accounting_AccountingSession
// Accounting information
type Subscriber_Session_Nodes_Node_Sessions_Session_Accounting_AccountingSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Accounting State Error Code for Accounting Session. The type is interface{}
    // with range: 0..4294967295.
    AccountingStateRc interface{}

    // Accounting Stop State. The type is interface{} with range: 0..4294967295.
    AccountingStopState interface{}

    // Accounting record context name. The type is string.
    RecordContextName interface{}

    // AAA method list name used to perform accounting. The type is string.
    MethodListName interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // Accounting start time in epoch-seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    AccountingStartTimeEpoch interface{}

    // True if interim accounting is enabled. The type is bool.
    IsInterimAccountingEnabled interface{}

    // Interim accounting interval (in minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    InterimInterval interface{}

    // Time of last successful interim update in epoch-seconds . The type is
    // interface{} with range: 0..18446744073709551615. Units are second.
    LastSuccessfulInterimUpdateTimeEpoch interface{}

    // Time of next interim update attempt (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    NextInterimUpdateAttemptTime interface{}

    // Time of last interim update attempt in epoch-seconds. The type is
    // interface{} with range: 0..18446744073709551615. Units are second.
    LastInterimUpdateAttemptTimeEpoch interface{}

    // Number of interim updates sent. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdates interface{}

    // Number of interim updates accepted. The type is interface{} with range:
    // 0..4294967295.
    AcceptedInterimUpdates interface{}

    // Number of interim updates rejected. The type is interface{} with range:
    // 0..4294967295.
    RejectedInterimUpdates interface{}

    // Number of interim update send failures. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdateFailures interface{}
}

func (accountingSession *Subscriber_Session_Nodes_Node_Sessions_Session_Accounting_AccountingSession) GetEntityData() *types.CommonEntityData {
    accountingSession.EntityData.YFilter = accountingSession.YFilter
    accountingSession.EntityData.YangName = "accounting-session"
    accountingSession.EntityData.BundleName = "cisco_ios_xr"
    accountingSession.EntityData.ParentYangName = "accounting"
    accountingSession.EntityData.SegmentPath = "accounting-session" + types.AddNoKeyToken(accountingSession)
    accountingSession.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/session/accounting/" + accountingSession.EntityData.SegmentPath
    accountingSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingSession.EntityData.Children = types.NewOrderedMap()
    accountingSession.EntityData.Leafs = types.NewOrderedMap()
    accountingSession.EntityData.Leafs.Append("accounting-state-rc", types.YLeaf{"AccountingStateRc", accountingSession.AccountingStateRc})
    accountingSession.EntityData.Leafs.Append("accounting-stop-state", types.YLeaf{"AccountingStopState", accountingSession.AccountingStopState})
    accountingSession.EntityData.Leafs.Append("record-context-name", types.YLeaf{"RecordContextName", accountingSession.RecordContextName})
    accountingSession.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", accountingSession.MethodListName})
    accountingSession.EntityData.Leafs.Append("account-session-id", types.YLeaf{"AccountSessionId", accountingSession.AccountSessionId})
    accountingSession.EntityData.Leafs.Append("accounting-start-time-epoch", types.YLeaf{"AccountingStartTimeEpoch", accountingSession.AccountingStartTimeEpoch})
    accountingSession.EntityData.Leafs.Append("is-interim-accounting-enabled", types.YLeaf{"IsInterimAccountingEnabled", accountingSession.IsInterimAccountingEnabled})
    accountingSession.EntityData.Leafs.Append("interim-interval", types.YLeaf{"InterimInterval", accountingSession.InterimInterval})
    accountingSession.EntityData.Leafs.Append("last-successful-interim-update-time-epoch", types.YLeaf{"LastSuccessfulInterimUpdateTimeEpoch", accountingSession.LastSuccessfulInterimUpdateTimeEpoch})
    accountingSession.EntityData.Leafs.Append("next-interim-update-attempt-time", types.YLeaf{"NextInterimUpdateAttemptTime", accountingSession.NextInterimUpdateAttemptTime})
    accountingSession.EntityData.Leafs.Append("last-interim-update-attempt-time-epoch", types.YLeaf{"LastInterimUpdateAttemptTimeEpoch", accountingSession.LastInterimUpdateAttemptTimeEpoch})
    accountingSession.EntityData.Leafs.Append("sent-interim-updates", types.YLeaf{"SentInterimUpdates", accountingSession.SentInterimUpdates})
    accountingSession.EntityData.Leafs.Append("accepted-interim-updates", types.YLeaf{"AcceptedInterimUpdates", accountingSession.AcceptedInterimUpdates})
    accountingSession.EntityData.Leafs.Append("rejected-interim-updates", types.YLeaf{"RejectedInterimUpdates", accountingSession.RejectedInterimUpdates})
    accountingSession.EntityData.Leafs.Append("sent-interim-update-failures", types.YLeaf{"SentInterimUpdateFailures", accountingSession.SentInterimUpdateFailures})

    accountingSession.EntityData.YListKeys = []string {}

    return &(accountingSession.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_SubPolicyData
// Subscriber control policy applied to this
// session
type Subscriber_Session_Nodes_Node_Sessions_Session_SubPolicyData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Matching event, condition and executed actions. The type is string.
    PolicyEpoch interface{}
}

func (subPolicyData *Subscriber_Session_Nodes_Node_Sessions_Session_SubPolicyData) GetEntityData() *types.CommonEntityData {
    subPolicyData.EntityData.YFilter = subPolicyData.YFilter
    subPolicyData.EntityData.YangName = "sub-policy-data"
    subPolicyData.EntityData.BundleName = "cisco_ios_xr"
    subPolicyData.EntityData.ParentYangName = "session"
    subPolicyData.EntityData.SegmentPath = "sub-policy-data" + types.AddNoKeyToken(subPolicyData)
    subPolicyData.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/session/" + subPolicyData.EntityData.SegmentPath
    subPolicyData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subPolicyData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subPolicyData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subPolicyData.EntityData.Children = types.NewOrderedMap()
    subPolicyData.EntityData.Leafs = types.NewOrderedMap()
    subPolicyData.EntityData.Leafs.Append("policy-epoch", types.YLeaf{"PolicyEpoch", subPolicyData.PolicyEpoch})

    subPolicyData.EntityData.YListKeys = []string {}

    return &(subPolicyData.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_SessionServiceInfo
// List of subscriber services associated to this
// session
type Subscriber_Session_Nodes_Node_Sessions_Session_SessionServiceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ServiceName. The type is string.
    ServiceName interface{}

    // ServiceParams. The type is string.
    ServiceParams interface{}

    // ServiceType. The type is IedgeOperService.
    ServiceType interface{}

    // ServiceStatus. The type is IedgeOperServiceStatus.
    ServiceStatus interface{}

    // ServiceIdentifier. The type is interface{} with range: 0..4294967295.
    ServiceId interface{}

    // ServicePrepaid. The type is bool.
    ServicePrepaid interface{}
}

func (sessionServiceInfo *Subscriber_Session_Nodes_Node_Sessions_Session_SessionServiceInfo) GetEntityData() *types.CommonEntityData {
    sessionServiceInfo.EntityData.YFilter = sessionServiceInfo.YFilter
    sessionServiceInfo.EntityData.YangName = "session-service-info"
    sessionServiceInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionServiceInfo.EntityData.ParentYangName = "session"
    sessionServiceInfo.EntityData.SegmentPath = "session-service-info" + types.AddNoKeyToken(sessionServiceInfo)
    sessionServiceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/session/" + sessionServiceInfo.EntityData.SegmentPath
    sessionServiceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionServiceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionServiceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionServiceInfo.EntityData.Children = types.NewOrderedMap()
    sessionServiceInfo.EntityData.Leafs = types.NewOrderedMap()
    sessionServiceInfo.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", sessionServiceInfo.ServiceName})
    sessionServiceInfo.EntityData.Leafs.Append("service-params", types.YLeaf{"ServiceParams", sessionServiceInfo.ServiceParams})
    sessionServiceInfo.EntityData.Leafs.Append("service-type", types.YLeaf{"ServiceType", sessionServiceInfo.ServiceType})
    sessionServiceInfo.EntityData.Leafs.Append("service-status", types.YLeaf{"ServiceStatus", sessionServiceInfo.ServiceStatus})
    sessionServiceInfo.EntityData.Leafs.Append("service-id", types.YLeaf{"ServiceId", sessionServiceInfo.ServiceId})
    sessionServiceInfo.EntityData.Leafs.Append("service-prepaid", types.YLeaf{"ServicePrepaid", sessionServiceInfo.ServicePrepaid})

    sessionServiceInfo.EntityData.YListKeys = []string {}

    return &(sessionServiceInfo.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_SessionChangeOfAuthorization
// Subscriber change of authorization information
type Subscriber_Session_Nodes_Node_Sessions_Session_SessionChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Coa Request Acked. The type is bool.
    RequestAcked interface{}

    // Request time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    RequestTimeEpoch interface{}

    // Reply time in epoch seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ReplyTimeEpoch interface{}
}

func (sessionChangeOfAuthorization *Subscriber_Session_Nodes_Node_Sessions_Session_SessionChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    sessionChangeOfAuthorization.EntityData.YFilter = sessionChangeOfAuthorization.YFilter
    sessionChangeOfAuthorization.EntityData.YangName = "session-change-of-authorization"
    sessionChangeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    sessionChangeOfAuthorization.EntityData.ParentYangName = "session"
    sessionChangeOfAuthorization.EntityData.SegmentPath = "session-change-of-authorization" + types.AddNoKeyToken(sessionChangeOfAuthorization)
    sessionChangeOfAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:subscriber/session/nodes/node/sessions/session/" + sessionChangeOfAuthorization.EntityData.SegmentPath
    sessionChangeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionChangeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionChangeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionChangeOfAuthorization.EntityData.Children = types.NewOrderedMap()
    sessionChangeOfAuthorization.EntityData.Leafs = types.NewOrderedMap()
    sessionChangeOfAuthorization.EntityData.Leafs.Append("request-acked", types.YLeaf{"RequestAcked", sessionChangeOfAuthorization.RequestAcked})
    sessionChangeOfAuthorization.EntityData.Leafs.Append("request-time-epoch", types.YLeaf{"RequestTimeEpoch", sessionChangeOfAuthorization.RequestTimeEpoch})
    sessionChangeOfAuthorization.EntityData.Leafs.Append("reply-time-epoch", types.YLeaf{"ReplyTimeEpoch", sessionChangeOfAuthorization.ReplyTimeEpoch})

    sessionChangeOfAuthorization.EntityData.YListKeys = []string {}

    return &(sessionChangeOfAuthorization.EntityData)
}

// IedgeLicenseManager
// iedge license manager
type IedgeLicenseManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session License Manager operational data for a location.
    Nodes IedgeLicenseManager_Nodes
}

func (iedgeLicenseManager *IedgeLicenseManager) GetEntityData() *types.CommonEntityData {
    iedgeLicenseManager.EntityData.YFilter = iedgeLicenseManager.YFilter
    iedgeLicenseManager.EntityData.YangName = "iedge-license-manager"
    iedgeLicenseManager.EntityData.BundleName = "cisco_ios_xr"
    iedgeLicenseManager.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-oper"
    iedgeLicenseManager.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-oper:iedge-license-manager"
    iedgeLicenseManager.EntityData.AbsolutePath = iedgeLicenseManager.EntityData.SegmentPath
    iedgeLicenseManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeLicenseManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeLicenseManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeLicenseManager.EntityData.Children = types.NewOrderedMap()
    iedgeLicenseManager.EntityData.Children.Append("nodes", types.YChild{"Nodes", &iedgeLicenseManager.Nodes})
    iedgeLicenseManager.EntityData.Leafs = types.NewOrderedMap()

    iedgeLicenseManager.EntityData.YListKeys = []string {}

    return &(iedgeLicenseManager.EntityData)
}

// IedgeLicenseManager_Nodes
// Session License Manager operational data for a
// location
type IedgeLicenseManager_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For example, 0/1/CPU0. The type is slice of
    // IedgeLicenseManager_Nodes_Node.
    Node []*IedgeLicenseManager_Nodes_Node
}

func (nodes *IedgeLicenseManager_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "iedge-license-manager"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:iedge-license-manager/" + nodes.EntityData.SegmentPath
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

// IedgeLicenseManager_Nodes_Node
// Location. For example, 0/1/CPU0
type IedgeLicenseManager_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node id to filter on. For example, 0/1/CPU0.
    // The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Nodeid interface{}

    // Display Session License Manager summary data.
    IedgeLicenseManagerSummary IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary
}

func (node *IedgeLicenseManager_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Nodeid, "nodeid")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:iedge-license-manager/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("iedge-license-manager-summary", types.YChild{"IedgeLicenseManagerSummary", &node.IedgeLicenseManagerSummary})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("nodeid", types.YLeaf{"Nodeid", node.Nodeid})

    node.EntityData.YListKeys = []string {"Nodeid"}

    return &(node.EntityData)
}

// IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary
// Display Session License Manager summary data
type IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configured session limit. The type is interface{} with range:
    // 0..4294967295.
    SessionLimit interface{}

    // configured session threshold. The type is interface{} with range:
    // 0..4294967295.
    SessionThreshold interface{}

    // number of license. The type is interface{} with range: 0..4294967295.
    SessionLicenseCount interface{}

    // number of sessions. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (iedgeLicenseManagerSummary *IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary) GetEntityData() *types.CommonEntityData {
    iedgeLicenseManagerSummary.EntityData.YFilter = iedgeLicenseManagerSummary.YFilter
    iedgeLicenseManagerSummary.EntityData.YangName = "iedge-license-manager-summary"
    iedgeLicenseManagerSummary.EntityData.BundleName = "cisco_ios_xr"
    iedgeLicenseManagerSummary.EntityData.ParentYangName = "node"
    iedgeLicenseManagerSummary.EntityData.SegmentPath = "iedge-license-manager-summary"
    iedgeLicenseManagerSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-iedge4710-oper:iedge-license-manager/nodes/node/" + iedgeLicenseManagerSummary.EntityData.SegmentPath
    iedgeLicenseManagerSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeLicenseManagerSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeLicenseManagerSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeLicenseManagerSummary.EntityData.Children = types.NewOrderedMap()
    iedgeLicenseManagerSummary.EntityData.Leafs = types.NewOrderedMap()
    iedgeLicenseManagerSummary.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", iedgeLicenseManagerSummary.SessionLimit})
    iedgeLicenseManagerSummary.EntityData.Leafs.Append("session-threshold", types.YLeaf{"SessionThreshold", iedgeLicenseManagerSummary.SessionThreshold})
    iedgeLicenseManagerSummary.EntityData.Leafs.Append("session-license-count", types.YLeaf{"SessionLicenseCount", iedgeLicenseManagerSummary.SessionLicenseCount})
    iedgeLicenseManagerSummary.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", iedgeLicenseManagerSummary.SessionCount})

    iedgeLicenseManagerSummary.EntityData.YListKeys = []string {}

    return &(iedgeLicenseManagerSummary.EntityData)
}

