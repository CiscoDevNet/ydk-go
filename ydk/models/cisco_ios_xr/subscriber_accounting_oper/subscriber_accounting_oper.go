// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-accounting package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-accounting: Subscriber accounting operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_accounting_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_accounting_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-accounting-oper subscriber-accounting}", reflect.TypeOf(SubscriberAccounting{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting", reflect.TypeOf(SubscriberAccounting{}))
}

// SubscriberAccounting
// Subscriber accounting operational data
type SubscriberAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber accounting operational data for a particular location.
    Nodes SubscriberAccounting_Nodes
}

func (subscriberAccounting *SubscriberAccounting) GetEntityData() *types.CommonEntityData {
    subscriberAccounting.EntityData.YFilter = subscriberAccounting.YFilter
    subscriberAccounting.EntityData.YangName = "subscriber-accounting"
    subscriberAccounting.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccounting.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-accounting-oper"
    subscriberAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting"
    subscriberAccounting.EntityData.AbsolutePath = subscriberAccounting.EntityData.SegmentPath
    subscriberAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccounting.EntityData.Children = types.NewOrderedMap()
    subscriberAccounting.EntityData.Children.Append("nodes", types.YChild{"Nodes", &subscriberAccounting.Nodes})
    subscriberAccounting.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccounting.EntityData.YListKeys = []string {}

    return &(subscriberAccounting.EntityData)
}

// SubscriberAccounting_Nodes
// Subscriber accounting operational data for a
// particular location
type SubscriberAccounting_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For example, 0/1/CPU0. The type is slice of
    // SubscriberAccounting_Nodes_Node.
    Node []*SubscriberAccounting_Nodes_Node
}

func (nodes *SubscriberAccounting_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "subscriber-accounting"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/" + nodes.EntityData.SegmentPath
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

// SubscriberAccounting_Nodes_Node
// Location. For example, 0/1/CPU0
type SubscriberAccounting_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node id to filter on. For example, 0/1/CPU0.
    // The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Subscriber accounting session feature data.
    SubscriberAccountingSessionFeatures SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures

    // Display subscriber accounting summary data.
    SubscriberAccountingSummary SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary

    // Subscriber accounting flow feature data.
    SubscriberAccountingFlowFeatures SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures

    // Subscriber accounting SRG feature data.
    SubscriberAccountingSrgFeatures SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures
}

func (node *SubscriberAccounting_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("subscriber-accounting-session-features", types.YChild{"SubscriberAccountingSessionFeatures", &node.SubscriberAccountingSessionFeatures})
    node.EntityData.Children.Append("subscriber-accounting-summary", types.YChild{"SubscriberAccountingSummary", &node.SubscriberAccountingSummary})
    node.EntityData.Children.Append("subscriber-accounting-flow-features", types.YChild{"SubscriberAccountingFlowFeatures", &node.SubscriberAccountingFlowFeatures})
    node.EntityData.Children.Append("subscriber-accounting-srg-features", types.YChild{"SubscriberAccountingSrgFeatures", &node.SubscriberAccountingSrgFeatures})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

    return &(node.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures
// Subscriber accounting session feature data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Display accounting session features by unique subscriber label. The type is
    // slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature.
    SubscriberAccountingSessionFeature []*SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetEntityData() *types.CommonEntityData {
    subscriberAccountingSessionFeatures.EntityData.YFilter = subscriberAccountingSessionFeatures.YFilter
    subscriberAccountingSessionFeatures.EntityData.YangName = "subscriber-accounting-session-features"
    subscriberAccountingSessionFeatures.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingSessionFeatures.EntityData.ParentYangName = "node"
    subscriberAccountingSessionFeatures.EntityData.SegmentPath = "subscriber-accounting-session-features"
    subscriberAccountingSessionFeatures.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/" + subscriberAccountingSessionFeatures.EntityData.SegmentPath
    subscriberAccountingSessionFeatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingSessionFeatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingSessionFeatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingSessionFeatures.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingSessionFeatures.EntityData.Children.Append("subscriber-accounting-session-feature", types.YChild{"SubscriberAccountingSessionFeature", nil})
    for i := range subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature {
        subscriberAccountingSessionFeatures.EntityData.Children.Append(types.GetSegmentPath(subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature[i]), types.YChild{"SubscriberAccountingSessionFeature", subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature[i]})
    }
    subscriberAccountingSessionFeatures.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccountingSessionFeatures.EntityData.YListKeys = []string {}

    return &(subscriberAccountingSessionFeatures.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature
// Display accounting session features by unique
// subscriber label
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique subscriber label. The type is interface{}
    // with range: 0..4294967295.
    SubLabel interface{}

    // Accounting session feature display data.
    SessionFeatureData SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetEntityData() *types.CommonEntityData {
    subscriberAccountingSessionFeature.EntityData.YFilter = subscriberAccountingSessionFeature.YFilter
    subscriberAccountingSessionFeature.EntityData.YangName = "subscriber-accounting-session-feature"
    subscriberAccountingSessionFeature.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingSessionFeature.EntityData.ParentYangName = "subscriber-accounting-session-features"
    subscriberAccountingSessionFeature.EntityData.SegmentPath = "subscriber-accounting-session-feature" + types.AddKeyToken(subscriberAccountingSessionFeature.SubLabel, "sub-label")
    subscriberAccountingSessionFeature.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-session-features/" + subscriberAccountingSessionFeature.EntityData.SegmentPath
    subscriberAccountingSessionFeature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingSessionFeature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingSessionFeature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingSessionFeature.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingSessionFeature.EntityData.Children.Append("session-feature-data", types.YChild{"SessionFeatureData", &subscriberAccountingSessionFeature.SessionFeatureData})
    subscriberAccountingSessionFeature.EntityData.Leafs = types.NewOrderedMap()
    subscriberAccountingSessionFeature.EntityData.Leafs.Append("sub-label", types.YLeaf{"SubLabel", subscriberAccountingSessionFeature.SubLabel})

    subscriberAccountingSessionFeature.EntityData.YListKeys = []string {"SubLabel"}

    return &(subscriberAccountingSessionFeature.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData
// Accounting session feature display data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique subscriber label used to identify the session. The type is
    // interface{} with range: 0..4294967295.
    UniqueSubscriberLabel interface{}

    // Handle of interface associated with the session. The type is interface{}
    // with range: 0..4294967295.
    InterfaceHandle interface{}

    // True if session is disconnected. The type is bool.
    SessionDisconnected interface{}

    // True if session accounting is enabled. The type is bool.
    SessionAccountingEnabledFlag interface{}

    // Session accounting method list name. The type is string with length:
    // 0..256.
    SessionAccountingMethodList interface{}

    // Session accounting periodic interval. The type is interface{} with range:
    // 0..4294967295.
    SessionAccountingPeriodicInterval interface{}

    // Number of Session Accounting AAA transactions pending. The type is
    // interface{} with range: 0..4294967295.
    SessionAccountingAaaTransPending interface{}

    // Number of Session Accounting AAA request failures. The type is interface{}
    // with range: 0..4294967295.
    SessionAccountingAaaRequestFailed interface{}

    // True if session accounting started. The type is bool.
    SessionAccountingStarted interface{}

    // True if session idle timeout is enabled. The type is bool.
    SessionIdleTimeoutEnabledFlag interface{}

    // Idle timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    IdleTimeoutValue interface{}

    // Idle timeout threshold in minutes per packets. The type is interface{} with
    // range: 0..4294967295. Units are minute.
    IdleTimeoutThreshold interface{}

    // Idle timeout direction. The type is string with length: 0..256.
    IdleTimeoutDirection interface{}

    // True if session is idle. The type is bool.
    SessionIsIdle interface{}

    // Last time session data was received. The type is interface{} with range:
    // 0..4294967295.
    SessionStatsChangedTime interface{}

    // Total time session has been idle. The type is interface{} with range:
    // 0..4294967295.
    SessionTotalIdleTime interface{}

    // Number of Session Idle AAA events. The type is interface{} with range:
    // 0..4294967295.
    SessionToIdleCount interface{}

    // Number of Session Awake AAA events. The type is interface{} with range:
    // 0..4294967295.
    SessionToAwakeCount interface{}

    // Number of Session Idle AAA transactions pending. The type is interface{}
    // with range: 0..4294967295.
    SessionIdleToAaaTransPending interface{}

    // Number of Session Idle AAA request failures. The type is interface{} with
    // range: 0..4294967295.
    SessionIdleToAaaRequestFailed interface{}

    // True if session timeout is enabled. The type is bool.
    SessionTimeoutEnabledFlag interface{}

    // Session timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    SessionTimeoutValue interface{}

    // Number seconds remaining before session times out. The type is interface{}
    // with range: 0..4294967295. Units are second.
    SessionTimeoutTimeRemaining interface{}

    // List of service accounting features. The type is slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature.
    ServiceAccountingFeature []*SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetEntityData() *types.CommonEntityData {
    sessionFeatureData.EntityData.YFilter = sessionFeatureData.YFilter
    sessionFeatureData.EntityData.YangName = "session-feature-data"
    sessionFeatureData.EntityData.BundleName = "cisco_ios_xr"
    sessionFeatureData.EntityData.ParentYangName = "subscriber-accounting-session-feature"
    sessionFeatureData.EntityData.SegmentPath = "session-feature-data"
    sessionFeatureData.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-session-features/subscriber-accounting-session-feature/" + sessionFeatureData.EntityData.SegmentPath
    sessionFeatureData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionFeatureData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionFeatureData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionFeatureData.EntityData.Children = types.NewOrderedMap()
    sessionFeatureData.EntityData.Children.Append("service-accounting-feature", types.YChild{"ServiceAccountingFeature", nil})
    for i := range sessionFeatureData.ServiceAccountingFeature {
        types.SetYListKey(sessionFeatureData.ServiceAccountingFeature[i], i)
        sessionFeatureData.EntityData.Children.Append(types.GetSegmentPath(sessionFeatureData.ServiceAccountingFeature[i]), types.YChild{"ServiceAccountingFeature", sessionFeatureData.ServiceAccountingFeature[i]})
    }
    sessionFeatureData.EntityData.Leafs = types.NewOrderedMap()
    sessionFeatureData.EntityData.Leafs.Append("unique-subscriber-label", types.YLeaf{"UniqueSubscriberLabel", sessionFeatureData.UniqueSubscriberLabel})
    sessionFeatureData.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", sessionFeatureData.InterfaceHandle})
    sessionFeatureData.EntityData.Leafs.Append("session-disconnected", types.YLeaf{"SessionDisconnected", sessionFeatureData.SessionDisconnected})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-enabled-flag", types.YLeaf{"SessionAccountingEnabledFlag", sessionFeatureData.SessionAccountingEnabledFlag})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-method-list", types.YLeaf{"SessionAccountingMethodList", sessionFeatureData.SessionAccountingMethodList})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-periodic-interval", types.YLeaf{"SessionAccountingPeriodicInterval", sessionFeatureData.SessionAccountingPeriodicInterval})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-aaa-trans-pending", types.YLeaf{"SessionAccountingAaaTransPending", sessionFeatureData.SessionAccountingAaaTransPending})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-aaa-request-failed", types.YLeaf{"SessionAccountingAaaRequestFailed", sessionFeatureData.SessionAccountingAaaRequestFailed})
    sessionFeatureData.EntityData.Leafs.Append("session-accounting-started", types.YLeaf{"SessionAccountingStarted", sessionFeatureData.SessionAccountingStarted})
    sessionFeatureData.EntityData.Leafs.Append("session-idle-timeout-enabled-flag", types.YLeaf{"SessionIdleTimeoutEnabledFlag", sessionFeatureData.SessionIdleTimeoutEnabledFlag})
    sessionFeatureData.EntityData.Leafs.Append("idle-timeout-value", types.YLeaf{"IdleTimeoutValue", sessionFeatureData.IdleTimeoutValue})
    sessionFeatureData.EntityData.Leafs.Append("idle-timeout-threshold", types.YLeaf{"IdleTimeoutThreshold", sessionFeatureData.IdleTimeoutThreshold})
    sessionFeatureData.EntityData.Leafs.Append("idle-timeout-direction", types.YLeaf{"IdleTimeoutDirection", sessionFeatureData.IdleTimeoutDirection})
    sessionFeatureData.EntityData.Leafs.Append("session-is-idle", types.YLeaf{"SessionIsIdle", sessionFeatureData.SessionIsIdle})
    sessionFeatureData.EntityData.Leafs.Append("session-stats-changed-time", types.YLeaf{"SessionStatsChangedTime", sessionFeatureData.SessionStatsChangedTime})
    sessionFeatureData.EntityData.Leafs.Append("session-total-idle-time", types.YLeaf{"SessionTotalIdleTime", sessionFeatureData.SessionTotalIdleTime})
    sessionFeatureData.EntityData.Leafs.Append("session-to-idle-count", types.YLeaf{"SessionToIdleCount", sessionFeatureData.SessionToIdleCount})
    sessionFeatureData.EntityData.Leafs.Append("session-to-awake-count", types.YLeaf{"SessionToAwakeCount", sessionFeatureData.SessionToAwakeCount})
    sessionFeatureData.EntityData.Leafs.Append("session-idle-to-aaa-trans-pending", types.YLeaf{"SessionIdleToAaaTransPending", sessionFeatureData.SessionIdleToAaaTransPending})
    sessionFeatureData.EntityData.Leafs.Append("session-idle-to-aaa-request-failed", types.YLeaf{"SessionIdleToAaaRequestFailed", sessionFeatureData.SessionIdleToAaaRequestFailed})
    sessionFeatureData.EntityData.Leafs.Append("session-timeout-enabled-flag", types.YLeaf{"SessionTimeoutEnabledFlag", sessionFeatureData.SessionTimeoutEnabledFlag})
    sessionFeatureData.EntityData.Leafs.Append("session-timeout-value", types.YLeaf{"SessionTimeoutValue", sessionFeatureData.SessionTimeoutValue})
    sessionFeatureData.EntityData.Leafs.Append("session-timeout-time-remaining", types.YLeaf{"SessionTimeoutTimeRemaining", sessionFeatureData.SessionTimeoutTimeRemaining})

    sessionFeatureData.EntityData.YListKeys = []string {}

    return &(sessionFeatureData.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature
// List of service accounting features
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // True if service accounting is enabled. The type is bool.
    ServiceAccountingEnabledFlag interface{}

    // Service accounting service ID. The type is interface{} with range:
    // 0..4294967295.
    ServiceAccountingServiceId interface{}

    // Service accounting method list name. The type is string with length:
    // 0..256.
    ServiceAccountingMethodList interface{}

    // Service accounting periodic interval. The type is interface{} with range:
    // 0..4294967295.
    ServiceAccountingPeriodicInterval interface{}

    // Number of Service Accounting AAA transactions pending for the service. The
    // type is interface{} with range: 0..4294967295.
    SessionAccountingAaaTransPending interface{}

    // Number of Service Accounting AAA request failures for the service. The type
    // is interface{} with range: 0..4294967295.
    SessionAccountingAaaRequestFailed interface{}

    // True if Service accounting started  for the service. The type is bool.
    SessionAccountingStarted interface{}
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetEntityData() *types.CommonEntityData {
    serviceAccountingFeature.EntityData.YFilter = serviceAccountingFeature.YFilter
    serviceAccountingFeature.EntityData.YangName = "service-accounting-feature"
    serviceAccountingFeature.EntityData.BundleName = "cisco_ios_xr"
    serviceAccountingFeature.EntityData.ParentYangName = "session-feature-data"
    serviceAccountingFeature.EntityData.SegmentPath = "service-accounting-feature" + types.AddNoKeyToken(serviceAccountingFeature)
    serviceAccountingFeature.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-session-features/subscriber-accounting-session-feature/session-feature-data/" + serviceAccountingFeature.EntityData.SegmentPath
    serviceAccountingFeature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccountingFeature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccountingFeature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccountingFeature.EntityData.Children = types.NewOrderedMap()
    serviceAccountingFeature.EntityData.Leafs = types.NewOrderedMap()
    serviceAccountingFeature.EntityData.Leafs.Append("service-accounting-enabled-flag", types.YLeaf{"ServiceAccountingEnabledFlag", serviceAccountingFeature.ServiceAccountingEnabledFlag})
    serviceAccountingFeature.EntityData.Leafs.Append("service-accounting-service-id", types.YLeaf{"ServiceAccountingServiceId", serviceAccountingFeature.ServiceAccountingServiceId})
    serviceAccountingFeature.EntityData.Leafs.Append("service-accounting-method-list", types.YLeaf{"ServiceAccountingMethodList", serviceAccountingFeature.ServiceAccountingMethodList})
    serviceAccountingFeature.EntityData.Leafs.Append("service-accounting-periodic-interval", types.YLeaf{"ServiceAccountingPeriodicInterval", serviceAccountingFeature.ServiceAccountingPeriodicInterval})
    serviceAccountingFeature.EntityData.Leafs.Append("session-accounting-aaa-trans-pending", types.YLeaf{"SessionAccountingAaaTransPending", serviceAccountingFeature.SessionAccountingAaaTransPending})
    serviceAccountingFeature.EntityData.Leafs.Append("session-accounting-aaa-request-failed", types.YLeaf{"SessionAccountingAaaRequestFailed", serviceAccountingFeature.SessionAccountingAaaRequestFailed})
    serviceAccountingFeature.EntityData.Leafs.Append("session-accounting-started", types.YLeaf{"SessionAccountingStarted", serviceAccountingFeature.SessionAccountingStarted})

    serviceAccountingFeature.EntityData.YListKeys = []string {}

    return &(serviceAccountingFeature.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary
// Display subscriber accounting summary data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting feature AAA summary counters.
    AaaCounters SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters

    // Accounting feature idle timeout summary counters.
    IdleTimeoutCounters SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters

    // Accounting feature session timeout summary counters.
    SessionTimeoutCounters SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters

    // Accounting feature session context summary counters.
    SessionFlowCounters SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetEntityData() *types.CommonEntityData {
    subscriberAccountingSummary.EntityData.YFilter = subscriberAccountingSummary.YFilter
    subscriberAccountingSummary.EntityData.YangName = "subscriber-accounting-summary"
    subscriberAccountingSummary.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingSummary.EntityData.ParentYangName = "node"
    subscriberAccountingSummary.EntityData.SegmentPath = "subscriber-accounting-summary"
    subscriberAccountingSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/" + subscriberAccountingSummary.EntityData.SegmentPath
    subscriberAccountingSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingSummary.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingSummary.EntityData.Children.Append("aaa-counters", types.YChild{"AaaCounters", &subscriberAccountingSummary.AaaCounters})
    subscriberAccountingSummary.EntityData.Children.Append("idle-timeout-counters", types.YChild{"IdleTimeoutCounters", &subscriberAccountingSummary.IdleTimeoutCounters})
    subscriberAccountingSummary.EntityData.Children.Append("session-timeout-counters", types.YChild{"SessionTimeoutCounters", &subscriberAccountingSummary.SessionTimeoutCounters})
    subscriberAccountingSummary.EntityData.Children.Append("session-flow-counters", types.YChild{"SessionFlowCounters", &subscriberAccountingSummary.SessionFlowCounters})
    subscriberAccountingSummary.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccountingSummary.EntityData.YListKeys = []string {}

    return &(subscriberAccountingSummary.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters
// Accounting feature AAA summary counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Flow Start Requests Sent. The type is interface{} with range:
    // 0..4294967295.
    FlowStart interface{}

    // Number of Flow Disconnect Requests Sent. The type is interface{} with
    // range: 0..4294967295.
    FlowDisconnect interface{}

    // Number of Session Accounting Start Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    SessionAccountingStart interface{}

    // Number of Session Accounting Stop Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    SessionAccountingStop interface{}

    // Number of Session Accounting Update Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    SessionAccountingUpdate interface{}

    // Number of Service Accounting Start Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    ServiceAccountingStart interface{}

    // Number of Service Accounting Stop Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    ServiceAccountingStop interface{}

    // Number of Service Accounting Update Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    ServiceAccountingUpdate interface{}

    // Number of Flow Accounting Start Requests Sent. The type is interface{} with
    // range: 0..4294967295.
    FlowAccountingStart interface{}

    // Number of Flow Accounting Stop Requests Sent. The type is interface{} with
    // range: 0..4294967295.
    FlowAccountingStop interface{}

    // Number of Flow Accounting Update Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    FlowAccountingUpdate interface{}

    // Number of Accounting Callbacks Received. The type is interface{} with
    // range: 0..4294967295.
    AccountingCallback interface{}

    // Number of Session Accounting transactions pending. The type is interface{}
    // with range: 0..4294967295.
    SessionAcctTransPending interface{}

    // Number of Session Accounting requests that failed. The type is interface{}
    // with range: 0..4294967295.
    SessionAcctReqsFailed interface{}

    // Number of Session Accounting sessions out of sync. The type is interface{}
    // with range: 0..4294967295.
    SessionAcctOutOfSync interface{}

    // Number of Session Idle Timeout transactions pending. The type is
    // interface{} with range: 0..4294967295.
    SessionIdleToTransPending interface{}

    // Number of Session Idle Timeout requests that failed. The type is
    // interface{} with range: 0..4294967295.
    SessionIdleToReqsFailed interface{}

    // Number of Session Idle Timeout sessions out of sync. The type is
    // interface{} with range: 0..4294967295.
    SessionIdleToOutOfSync interface{}

    // Number of Service Accounting transactions pending. The type is interface{}
    // with range: 0..4294967295.
    ServiceAcctTransPending interface{}

    // Number of Service Accounting requests that failed. The type is interface{}
    // with range: 0..4294967295.
    ServiceAcctReqsFailed interface{}

    // Number of Service Accounting services out of sync. The type is interface{}
    // with range: 0..4294967295.
    ServiceAcctOutOfSync interface{}

    // Number of Service Idle Timeout transactions pending. The type is
    // interface{} with range: 0..4294967295.
    ServiceIdleToTransPending interface{}

    // Number of Service Idle Timeout requests that failed. The type is
    // interface{} with range: 0..4294967295.
    ServiceIdleToReqsFailed interface{}

    // Number of Service Idle Timeout services out of sync. The type is
    // interface{} with range: 0..4294967295.
    ServiceIdleToOutOfSync interface{}

    // Number of Prepaid Start Requests Sent. The type is interface{} with range:
    // 0..4294967295.
    PrepaidStart interface{}

    // Number of Prepaid Stop Requests Sent. The type is interface{} with range:
    // 0..4294967295.
    PrepaidStop interface{}

    // Number of Prepaid Accounting Start Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    PrepaidAccountingStart interface{}

    // Number of Prepaid Accounting Stop Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    PrepaidAccountingStop interface{}

    // Number of Prepaid Volume Threshold Reached Requests Sent. The type is
    // interface{} with range: 0..4294967295.
    PrepaidVolumeThresholdReached interface{}

    // Number of Prepaid Time Threshold Reached Requests Sent. The type is
    // interface{} with range: 0..4294967295.
    PrepaidTimeThresholdReached interface{}

    // Number of Prepaid Quota Depleted Requests Sent. The type is interface{}
    // with range: 0..4294967295.
    PrepaidQuotaDepleted interface{}

    // Number of Prepaid Authorization Requests Sent. The type is interface{} with
    // range: 0..4294967295.
    PrepaidReauthorization interface{}

    // Number of Idle Timeout Events Sent. The type is interface{} with range:
    // 0..4294967295.
    IdleTimeout interface{}

    // Number of Idle Timeout Callbacks Received. The type is interface{} with
    // range: 0..4294967295.
    IdleTimeoutResponseCallback interface{}

    // Number of Owned Resource Starts. The type is interface{} with range:
    // 0..4294967295.
    OwnedResourceStart interface{}
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetEntityData() *types.CommonEntityData {
    aaaCounters.EntityData.YFilter = aaaCounters.YFilter
    aaaCounters.EntityData.YangName = "aaa-counters"
    aaaCounters.EntityData.BundleName = "cisco_ios_xr"
    aaaCounters.EntityData.ParentYangName = "subscriber-accounting-summary"
    aaaCounters.EntityData.SegmentPath = "aaa-counters"
    aaaCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-summary/" + aaaCounters.EntityData.SegmentPath
    aaaCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaCounters.EntityData.Children = types.NewOrderedMap()
    aaaCounters.EntityData.Leafs = types.NewOrderedMap()
    aaaCounters.EntityData.Leafs.Append("flow-start", types.YLeaf{"FlowStart", aaaCounters.FlowStart})
    aaaCounters.EntityData.Leafs.Append("flow-disconnect", types.YLeaf{"FlowDisconnect", aaaCounters.FlowDisconnect})
    aaaCounters.EntityData.Leafs.Append("session-accounting-start", types.YLeaf{"SessionAccountingStart", aaaCounters.SessionAccountingStart})
    aaaCounters.EntityData.Leafs.Append("session-accounting-stop", types.YLeaf{"SessionAccountingStop", aaaCounters.SessionAccountingStop})
    aaaCounters.EntityData.Leafs.Append("session-accounting-update", types.YLeaf{"SessionAccountingUpdate", aaaCounters.SessionAccountingUpdate})
    aaaCounters.EntityData.Leafs.Append("service-accounting-start", types.YLeaf{"ServiceAccountingStart", aaaCounters.ServiceAccountingStart})
    aaaCounters.EntityData.Leafs.Append("service-accounting-stop", types.YLeaf{"ServiceAccountingStop", aaaCounters.ServiceAccountingStop})
    aaaCounters.EntityData.Leafs.Append("service-accounting-update", types.YLeaf{"ServiceAccountingUpdate", aaaCounters.ServiceAccountingUpdate})
    aaaCounters.EntityData.Leafs.Append("flow-accounting-start", types.YLeaf{"FlowAccountingStart", aaaCounters.FlowAccountingStart})
    aaaCounters.EntityData.Leafs.Append("flow-accounting-stop", types.YLeaf{"FlowAccountingStop", aaaCounters.FlowAccountingStop})
    aaaCounters.EntityData.Leafs.Append("flow-accounting-update", types.YLeaf{"FlowAccountingUpdate", aaaCounters.FlowAccountingUpdate})
    aaaCounters.EntityData.Leafs.Append("accounting-callback", types.YLeaf{"AccountingCallback", aaaCounters.AccountingCallback})
    aaaCounters.EntityData.Leafs.Append("session-acct-trans-pending", types.YLeaf{"SessionAcctTransPending", aaaCounters.SessionAcctTransPending})
    aaaCounters.EntityData.Leafs.Append("session-acct-reqs-failed", types.YLeaf{"SessionAcctReqsFailed", aaaCounters.SessionAcctReqsFailed})
    aaaCounters.EntityData.Leafs.Append("session-acct-out-of-sync", types.YLeaf{"SessionAcctOutOfSync", aaaCounters.SessionAcctOutOfSync})
    aaaCounters.EntityData.Leafs.Append("session-idle-to-trans-pending", types.YLeaf{"SessionIdleToTransPending", aaaCounters.SessionIdleToTransPending})
    aaaCounters.EntityData.Leafs.Append("session-idle-to-reqs-failed", types.YLeaf{"SessionIdleToReqsFailed", aaaCounters.SessionIdleToReqsFailed})
    aaaCounters.EntityData.Leafs.Append("session-idle-to-out-of-sync", types.YLeaf{"SessionIdleToOutOfSync", aaaCounters.SessionIdleToOutOfSync})
    aaaCounters.EntityData.Leafs.Append("service-acct-trans-pending", types.YLeaf{"ServiceAcctTransPending", aaaCounters.ServiceAcctTransPending})
    aaaCounters.EntityData.Leafs.Append("service-acct-reqs-failed", types.YLeaf{"ServiceAcctReqsFailed", aaaCounters.ServiceAcctReqsFailed})
    aaaCounters.EntityData.Leafs.Append("service-acct-out-of-sync", types.YLeaf{"ServiceAcctOutOfSync", aaaCounters.ServiceAcctOutOfSync})
    aaaCounters.EntityData.Leafs.Append("service-idle-to-trans-pending", types.YLeaf{"ServiceIdleToTransPending", aaaCounters.ServiceIdleToTransPending})
    aaaCounters.EntityData.Leafs.Append("service-idle-to-reqs-failed", types.YLeaf{"ServiceIdleToReqsFailed", aaaCounters.ServiceIdleToReqsFailed})
    aaaCounters.EntityData.Leafs.Append("service-idle-to-out-of-sync", types.YLeaf{"ServiceIdleToOutOfSync", aaaCounters.ServiceIdleToOutOfSync})
    aaaCounters.EntityData.Leafs.Append("prepaid-start", types.YLeaf{"PrepaidStart", aaaCounters.PrepaidStart})
    aaaCounters.EntityData.Leafs.Append("prepaid-stop", types.YLeaf{"PrepaidStop", aaaCounters.PrepaidStop})
    aaaCounters.EntityData.Leafs.Append("prepaid-accounting-start", types.YLeaf{"PrepaidAccountingStart", aaaCounters.PrepaidAccountingStart})
    aaaCounters.EntityData.Leafs.Append("prepaid-accounting-stop", types.YLeaf{"PrepaidAccountingStop", aaaCounters.PrepaidAccountingStop})
    aaaCounters.EntityData.Leafs.Append("prepaid-volume-threshold-reached", types.YLeaf{"PrepaidVolumeThresholdReached", aaaCounters.PrepaidVolumeThresholdReached})
    aaaCounters.EntityData.Leafs.Append("prepaid-time-threshold-reached", types.YLeaf{"PrepaidTimeThresholdReached", aaaCounters.PrepaidTimeThresholdReached})
    aaaCounters.EntityData.Leafs.Append("prepaid-quota-depleted", types.YLeaf{"PrepaidQuotaDepleted", aaaCounters.PrepaidQuotaDepleted})
    aaaCounters.EntityData.Leafs.Append("prepaid-reauthorization", types.YLeaf{"PrepaidReauthorization", aaaCounters.PrepaidReauthorization})
    aaaCounters.EntityData.Leafs.Append("idle-timeout", types.YLeaf{"IdleTimeout", aaaCounters.IdleTimeout})
    aaaCounters.EntityData.Leafs.Append("idle-timeout-response-callback", types.YLeaf{"IdleTimeoutResponseCallback", aaaCounters.IdleTimeoutResponseCallback})
    aaaCounters.EntityData.Leafs.Append("owned-resource-start", types.YLeaf{"OwnedResourceStart", aaaCounters.OwnedResourceStart})

    aaaCounters.EntityData.YListKeys = []string {}

    return &(aaaCounters.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters
// Accounting feature idle timeout summary counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Sessions with Idle Timeout Feature. The type is interface{} with
    // range: 0..4294967295.
    ActiveSessionIdleTimers interface{}

    // Number of Idle Sessions. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Number of Sessions Transitions to Idle State. The type is interface{} with
    // range: 0..4294967295.
    TransitionsToIdle interface{}

    // Number of Sessions Transitions to Awake State. The type is interface{} with
    // range: 0..4294967295.
    TransitionsToAwake interface{}

    // Number of Active Flow Idle Timers. The type is interface{} with range:
    // 0..4294967295.
    ActiveFlowIdleTimers interface{}

    // Number of Flow Expired Idle Timers. The type is interface{} with range:
    // 0..4294967295.
    ExpiredFlowIdleTimers interface{}

    // Number of Active Prepaid Idle Timers. The type is interface{} with range:
    // 0..4294967295.
    ActivePrepaidIdleTimers interface{}

    // Number of Expired Prepaid Idle Timers. The type is interface{} with range:
    // 0..4294967295.
    ExpiredPrepaidIdleTimers interface{}
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetEntityData() *types.CommonEntityData {
    idleTimeoutCounters.EntityData.YFilter = idleTimeoutCounters.YFilter
    idleTimeoutCounters.EntityData.YangName = "idle-timeout-counters"
    idleTimeoutCounters.EntityData.BundleName = "cisco_ios_xr"
    idleTimeoutCounters.EntityData.ParentYangName = "subscriber-accounting-summary"
    idleTimeoutCounters.EntityData.SegmentPath = "idle-timeout-counters"
    idleTimeoutCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-summary/" + idleTimeoutCounters.EntityData.SegmentPath
    idleTimeoutCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idleTimeoutCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idleTimeoutCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idleTimeoutCounters.EntityData.Children = types.NewOrderedMap()
    idleTimeoutCounters.EntityData.Leafs = types.NewOrderedMap()
    idleTimeoutCounters.EntityData.Leafs.Append("active-session-idle-timers", types.YLeaf{"ActiveSessionIdleTimers", idleTimeoutCounters.ActiveSessionIdleTimers})
    idleTimeoutCounters.EntityData.Leafs.Append("idle-sessions", types.YLeaf{"IdleSessions", idleTimeoutCounters.IdleSessions})
    idleTimeoutCounters.EntityData.Leafs.Append("transitions-to-idle", types.YLeaf{"TransitionsToIdle", idleTimeoutCounters.TransitionsToIdle})
    idleTimeoutCounters.EntityData.Leafs.Append("transitions-to-awake", types.YLeaf{"TransitionsToAwake", idleTimeoutCounters.TransitionsToAwake})
    idleTimeoutCounters.EntityData.Leafs.Append("active-flow-idle-timers", types.YLeaf{"ActiveFlowIdleTimers", idleTimeoutCounters.ActiveFlowIdleTimers})
    idleTimeoutCounters.EntityData.Leafs.Append("expired-flow-idle-timers", types.YLeaf{"ExpiredFlowIdleTimers", idleTimeoutCounters.ExpiredFlowIdleTimers})
    idleTimeoutCounters.EntityData.Leafs.Append("active-prepaid-idle-timers", types.YLeaf{"ActivePrepaidIdleTimers", idleTimeoutCounters.ActivePrepaidIdleTimers})
    idleTimeoutCounters.EntityData.Leafs.Append("expired-prepaid-idle-timers", types.YLeaf{"ExpiredPrepaidIdleTimers", idleTimeoutCounters.ExpiredPrepaidIdleTimers})

    idleTimeoutCounters.EntityData.YListKeys = []string {}

    return &(idleTimeoutCounters.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters
// Accounting feature session timeout summary
// counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Active Session Timers. The type is interface{} with range:
    // 0..4294967295.
    ActiveSessionTimers interface{}

    // Number of Expired Session Timers. The type is interface{} with range:
    // 0..4294967295.
    ExpiredSessionTimers interface{}
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetEntityData() *types.CommonEntityData {
    sessionTimeoutCounters.EntityData.YFilter = sessionTimeoutCounters.YFilter
    sessionTimeoutCounters.EntityData.YangName = "session-timeout-counters"
    sessionTimeoutCounters.EntityData.BundleName = "cisco_ios_xr"
    sessionTimeoutCounters.EntityData.ParentYangName = "subscriber-accounting-summary"
    sessionTimeoutCounters.EntityData.SegmentPath = "session-timeout-counters"
    sessionTimeoutCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-summary/" + sessionTimeoutCounters.EntityData.SegmentPath
    sessionTimeoutCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionTimeoutCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionTimeoutCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionTimeoutCounters.EntityData.Children = types.NewOrderedMap()
    sessionTimeoutCounters.EntityData.Leafs = types.NewOrderedMap()
    sessionTimeoutCounters.EntityData.Leafs.Append("active-session-timers", types.YLeaf{"ActiveSessionTimers", sessionTimeoutCounters.ActiveSessionTimers})
    sessionTimeoutCounters.EntityData.Leafs.Append("expired-session-timers", types.YLeaf{"ExpiredSessionTimers", sessionTimeoutCounters.ExpiredSessionTimers})

    sessionTimeoutCounters.EntityData.YListKeys = []string {}

    return &(sessionTimeoutCounters.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters
// Accounting feature session context summary
// counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Active Sessions. The type is interface{} with range:
    // 0..4294967295.
    ActiveSessions interface{}

    // Number of Disconnected Sessions. The type is interface{} with range:
    // 0..4294967295.
    DisconnectedSessions interface{}

    // Number of Active Sessions with Accounting. The type is interface{} with
    // range: 0..4294967295.
    ActiveSessionAccountingSessions interface{}

    // Number of Active Flows. The type is interface{} with range: 0..4294967295.
    ActiveFlows interface{}

    // Number of flows for which Quota is received. The type is interface{} with
    // range: 0..4294967295.
    QuotaReceived interface{}
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetEntityData() *types.CommonEntityData {
    sessionFlowCounters.EntityData.YFilter = sessionFlowCounters.YFilter
    sessionFlowCounters.EntityData.YangName = "session-flow-counters"
    sessionFlowCounters.EntityData.BundleName = "cisco_ios_xr"
    sessionFlowCounters.EntityData.ParentYangName = "subscriber-accounting-summary"
    sessionFlowCounters.EntityData.SegmentPath = "session-flow-counters"
    sessionFlowCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-summary/" + sessionFlowCounters.EntityData.SegmentPath
    sessionFlowCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionFlowCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionFlowCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionFlowCounters.EntityData.Children = types.NewOrderedMap()
    sessionFlowCounters.EntityData.Leafs = types.NewOrderedMap()
    sessionFlowCounters.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", sessionFlowCounters.ActiveSessions})
    sessionFlowCounters.EntityData.Leafs.Append("disconnected-sessions", types.YLeaf{"DisconnectedSessions", sessionFlowCounters.DisconnectedSessions})
    sessionFlowCounters.EntityData.Leafs.Append("active-session-accounting-sessions", types.YLeaf{"ActiveSessionAccountingSessions", sessionFlowCounters.ActiveSessionAccountingSessions})
    sessionFlowCounters.EntityData.Leafs.Append("active-flows", types.YLeaf{"ActiveFlows", sessionFlowCounters.ActiveFlows})
    sessionFlowCounters.EntityData.Leafs.Append("quota-received", types.YLeaf{"QuotaReceived", sessionFlowCounters.QuotaReceived})

    sessionFlowCounters.EntityData.YListKeys = []string {}

    return &(sessionFlowCounters.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures
// Subscriber accounting flow feature data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Display accounting flow features by unique subscriber label. The type is
    // slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature.
    SubscriberAccountingFlowFeature []*SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetEntityData() *types.CommonEntityData {
    subscriberAccountingFlowFeatures.EntityData.YFilter = subscriberAccountingFlowFeatures.YFilter
    subscriberAccountingFlowFeatures.EntityData.YangName = "subscriber-accounting-flow-features"
    subscriberAccountingFlowFeatures.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingFlowFeatures.EntityData.ParentYangName = "node"
    subscriberAccountingFlowFeatures.EntityData.SegmentPath = "subscriber-accounting-flow-features"
    subscriberAccountingFlowFeatures.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/" + subscriberAccountingFlowFeatures.EntityData.SegmentPath
    subscriberAccountingFlowFeatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingFlowFeatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingFlowFeatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingFlowFeatures.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingFlowFeatures.EntityData.Children.Append("subscriber-accounting-flow-feature", types.YChild{"SubscriberAccountingFlowFeature", nil})
    for i := range subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature {
        subscriberAccountingFlowFeatures.EntityData.Children.Append(types.GetSegmentPath(subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature[i]), types.YChild{"SubscriberAccountingFlowFeature", subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature[i]})
    }
    subscriberAccountingFlowFeatures.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccountingFlowFeatures.EntityData.YListKeys = []string {}

    return &(subscriberAccountingFlowFeatures.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature
// Display accounting flow features by unique
// subscriber label
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique subscriber class label. The type is
    // interface{} with range: 0..4294967295.
    ClassLabel interface{}

    // Accouting flow feature display data.
    FlowFeatureData SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetEntityData() *types.CommonEntityData {
    subscriberAccountingFlowFeature.EntityData.YFilter = subscriberAccountingFlowFeature.YFilter
    subscriberAccountingFlowFeature.EntityData.YangName = "subscriber-accounting-flow-feature"
    subscriberAccountingFlowFeature.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingFlowFeature.EntityData.ParentYangName = "subscriber-accounting-flow-features"
    subscriberAccountingFlowFeature.EntityData.SegmentPath = "subscriber-accounting-flow-feature" + types.AddKeyToken(subscriberAccountingFlowFeature.ClassLabel, "class-label")
    subscriberAccountingFlowFeature.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-flow-features/" + subscriberAccountingFlowFeature.EntityData.SegmentPath
    subscriberAccountingFlowFeature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingFlowFeature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingFlowFeature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingFlowFeature.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingFlowFeature.EntityData.Children.Append("flow-feature-data", types.YChild{"FlowFeatureData", &subscriberAccountingFlowFeature.FlowFeatureData})
    subscriberAccountingFlowFeature.EntityData.Leafs = types.NewOrderedMap()
    subscriberAccountingFlowFeature.EntityData.Leafs.Append("class-label", types.YLeaf{"ClassLabel", subscriberAccountingFlowFeature.ClassLabel})

    subscriberAccountingFlowFeature.EntityData.YListKeys = []string {"ClassLabel"}

    return &(subscriberAccountingFlowFeature.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData
// Accouting flow feature display data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True if flow accounting is enabled. The type is bool.
    FlowAccountingEnabledFlag interface{}

    // True if flow idle timeout is enabled. The type is bool.
    FlowIdleTimeoutEnabledFlag interface{}

    // True if prepaid is enabled. The type is bool.
    PrepaidEnabledFlag interface{}

    // Flag to specify if absolute timeout for ervice is enabled. The type is
    // bool.
    PrepaidReauthTimerEnabled interface{}

    // Flag to specify if idle timeout for service is enabled. The type is bool.
    PrepaidIdleTimeoutEnabled interface{}

    // Prepaid final unit indication flag. The type is bool.
    PrepaidFinalUnit interface{}

    // Unique class label used to identify the flow. The type is interface{} with
    // range: 0..4294967295.
    UniqueClassLabel interface{}

    // Direction of the flow. 0 = Ingress, 1 = Egress. The type is interface{}
    // with range: 0..4294967295.
    FlowDirection interface{}

    // Flow accounting periodic interval. The type is interface{} with range:
    // 0..4294967295.
    FlowAccountingPeriodicInterval interface{}

    // Flow idle timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    FlowIdleTimeoutValue interface{}

    // Current prepaid time quota in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PrepaidTimeQuota interface{}

    // Current prepaid time threshold in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PrepaidTimeThreshold interface{}

    // Total accumulated prepaid time quota. The type is interface{} with range:
    // 0..4294967295.
    PrepaidTotalTimeQuota interface{}

    // Current prepaid volume threshold in bytes. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    PrepaidVolumeThreshold interface{}

    // The time remaing for QT timer to fire. The type is interface{} with range:
    // 0..4294967295.
    PrepaidRemainingQt interface{}

    // The time remaing for quota absolute timer to fire. The type is interface{}
    // with range: 0..4294967295.
    PrepaidRemainingQat interface{}

    // The time remaing for quota holding timer to fire . The type is interface{}
    // with range: 0..4294967295.
    PrepaidRemainingQit interface{}

    // The time remaining for tariff timer to fire. The type is interface{} with
    // range: 0..4294967295.
    PrepaidRemainingQtt interface{}

    // The time remaining for idle timer wheel to fire. The type is interface{}
    // with range: 0..4294967295.
    PrepaidRemainingWheel interface{}

    // The absolute time at which the traffic switch will occur expressed in UNIX
    // time. The type is interface{} with range: 0..4294967295.
    PrepaidTariffTime interface{}

    // Prepaid idle timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    PrepaidIdleTimeoutValue interface{}

    // The time at which the re-authorization will occur. The type is interface{}
    // with range: 0..4294967295.
    PrepaidReauthTimeoutValue interface{}

    // Prepaid CCFH flag. The type is interface{} with range: 0..4294967295.
    PrepaidCcfh interface{}

    // Prepaid authorization operation result code. The type is interface{} with
    // range: 0..4294967295.
    PrepaidResultCode interface{}

    // Current prepaid input volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeiQuota interface{}

    // Current prepaid output volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeoQuota interface{}

    // Current prepaid total volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumebQuota interface{}

    // Total accumulated input volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidTotalVolumeiQuota interface{}

    // Total accumulated output volume quota in bytes. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    PrepaidTotalVolumeoQuota interface{}

    // Total accumulated total volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidTotalVolumebQuota interface{}

    // Accumulated input volume used quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeUsediQuota interface{}

    // Accumulated output volume used quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeUsedoQuota interface{}

    // Accumulated input volume reference quota in bytes. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeRefiQuota interface{}

    // Accumulated output volume reference quota in bytes. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeRefoQuota interface{}

    // Accumulated bi-directional volume reference quota in bytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeRefbQuota interface{}

    // Total Accumulated input volume reference quota in bytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeAggRefiQuota interface{}

    // Total Accumulated output volume reference quota in bytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeAggRefoQuota interface{}

    // Total Accumulated bi-directional volume reference quota in bytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeAggRefbQuota interface{}

    // Newly arrvied input volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeNewiQuota interface{}

    // Newly arrvied output volume quota in bytes. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeNewoQuota interface{}

    // Newly arrvied bi-directional volume quota in bytes. The type is interface{}
    // with range: 0..18446744073709551615. Units are byte.
    PrepaidVolumeNewbQuota interface{}

    // Total accumulated prepaid pre-tarrif input volume quota in bytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidTariffVolumeiQuota interface{}

    // Total accumulated prepaid pre-tarrif output volume quota in bytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidTariffVolumeoQuota interface{}

    // Total accumulated prepaid pre-tarrif total volume quota in bytes. The type
    // is interface{} with range: 0..18446744073709551615. Units are byte.
    PrepaidTariffVolumebQuota interface{}

    // Flow accounting method list name. The type is string with length: 0..256.
    FlowAccountingMethodListName interface{}

    // Prepaid Config. The type is string with length: 0..256.
    PrepaidCfg interface{}

    // Prepaid time state machine state. The type is string with length: 0..256.
    PrepaidTimeState interface{}

    // Prepaid volume state machine state. The type is string with length: 0..256.
    PrepaidVolumeState interface{}

    // Prepaid charging rule name string. The type is string with length: 0..256.
    PrepaidChargingRule interface{}
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetEntityData() *types.CommonEntityData {
    flowFeatureData.EntityData.YFilter = flowFeatureData.YFilter
    flowFeatureData.EntityData.YangName = "flow-feature-data"
    flowFeatureData.EntityData.BundleName = "cisco_ios_xr"
    flowFeatureData.EntityData.ParentYangName = "subscriber-accounting-flow-feature"
    flowFeatureData.EntityData.SegmentPath = "flow-feature-data"
    flowFeatureData.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-flow-features/subscriber-accounting-flow-feature/" + flowFeatureData.EntityData.SegmentPath
    flowFeatureData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowFeatureData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowFeatureData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowFeatureData.EntityData.Children = types.NewOrderedMap()
    flowFeatureData.EntityData.Leafs = types.NewOrderedMap()
    flowFeatureData.EntityData.Leafs.Append("flow-accounting-enabled-flag", types.YLeaf{"FlowAccountingEnabledFlag", flowFeatureData.FlowAccountingEnabledFlag})
    flowFeatureData.EntityData.Leafs.Append("flow-idle-timeout-enabled-flag", types.YLeaf{"FlowIdleTimeoutEnabledFlag", flowFeatureData.FlowIdleTimeoutEnabledFlag})
    flowFeatureData.EntityData.Leafs.Append("prepaid-enabled-flag", types.YLeaf{"PrepaidEnabledFlag", flowFeatureData.PrepaidEnabledFlag})
    flowFeatureData.EntityData.Leafs.Append("prepaid-reauth-timer-enabled", types.YLeaf{"PrepaidReauthTimerEnabled", flowFeatureData.PrepaidReauthTimerEnabled})
    flowFeatureData.EntityData.Leafs.Append("prepaid-idle-timeout-enabled", types.YLeaf{"PrepaidIdleTimeoutEnabled", flowFeatureData.PrepaidIdleTimeoutEnabled})
    flowFeatureData.EntityData.Leafs.Append("prepaid-final-unit", types.YLeaf{"PrepaidFinalUnit", flowFeatureData.PrepaidFinalUnit})
    flowFeatureData.EntityData.Leafs.Append("unique-class-label", types.YLeaf{"UniqueClassLabel", flowFeatureData.UniqueClassLabel})
    flowFeatureData.EntityData.Leafs.Append("flow-direction", types.YLeaf{"FlowDirection", flowFeatureData.FlowDirection})
    flowFeatureData.EntityData.Leafs.Append("flow-accounting-periodic-interval", types.YLeaf{"FlowAccountingPeriodicInterval", flowFeatureData.FlowAccountingPeriodicInterval})
    flowFeatureData.EntityData.Leafs.Append("flow-idle-timeout-value", types.YLeaf{"FlowIdleTimeoutValue", flowFeatureData.FlowIdleTimeoutValue})
    flowFeatureData.EntityData.Leafs.Append("prepaid-time-quota", types.YLeaf{"PrepaidTimeQuota", flowFeatureData.PrepaidTimeQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-time-threshold", types.YLeaf{"PrepaidTimeThreshold", flowFeatureData.PrepaidTimeThreshold})
    flowFeatureData.EntityData.Leafs.Append("prepaid-total-time-quota", types.YLeaf{"PrepaidTotalTimeQuota", flowFeatureData.PrepaidTotalTimeQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-threshold", types.YLeaf{"PrepaidVolumeThreshold", flowFeatureData.PrepaidVolumeThreshold})
    flowFeatureData.EntityData.Leafs.Append("prepaid-remaining-qt", types.YLeaf{"PrepaidRemainingQt", flowFeatureData.PrepaidRemainingQt})
    flowFeatureData.EntityData.Leafs.Append("prepaid-remaining-qat", types.YLeaf{"PrepaidRemainingQat", flowFeatureData.PrepaidRemainingQat})
    flowFeatureData.EntityData.Leafs.Append("prepaid-remaining-qit", types.YLeaf{"PrepaidRemainingQit", flowFeatureData.PrepaidRemainingQit})
    flowFeatureData.EntityData.Leafs.Append("prepaid-remaining-qtt", types.YLeaf{"PrepaidRemainingQtt", flowFeatureData.PrepaidRemainingQtt})
    flowFeatureData.EntityData.Leafs.Append("prepaid-remaining-wheel", types.YLeaf{"PrepaidRemainingWheel", flowFeatureData.PrepaidRemainingWheel})
    flowFeatureData.EntityData.Leafs.Append("prepaid-tariff-time", types.YLeaf{"PrepaidTariffTime", flowFeatureData.PrepaidTariffTime})
    flowFeatureData.EntityData.Leafs.Append("prepaid-idle-timeout-value", types.YLeaf{"PrepaidIdleTimeoutValue", flowFeatureData.PrepaidIdleTimeoutValue})
    flowFeatureData.EntityData.Leafs.Append("prepaid-reauth-timeout-value", types.YLeaf{"PrepaidReauthTimeoutValue", flowFeatureData.PrepaidReauthTimeoutValue})
    flowFeatureData.EntityData.Leafs.Append("prepaid-ccfh", types.YLeaf{"PrepaidCcfh", flowFeatureData.PrepaidCcfh})
    flowFeatureData.EntityData.Leafs.Append("prepaid-result-code", types.YLeaf{"PrepaidResultCode", flowFeatureData.PrepaidResultCode})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volumei-quota", types.YLeaf{"PrepaidVolumeiQuota", flowFeatureData.PrepaidVolumeiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volumeo-quota", types.YLeaf{"PrepaidVolumeoQuota", flowFeatureData.PrepaidVolumeoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volumeb-quota", types.YLeaf{"PrepaidVolumebQuota", flowFeatureData.PrepaidVolumebQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-total-volumei-quota", types.YLeaf{"PrepaidTotalVolumeiQuota", flowFeatureData.PrepaidTotalVolumeiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-total-volumeo-quota", types.YLeaf{"PrepaidTotalVolumeoQuota", flowFeatureData.PrepaidTotalVolumeoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-total-volumeb-quota", types.YLeaf{"PrepaidTotalVolumebQuota", flowFeatureData.PrepaidTotalVolumebQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-usedi-quota", types.YLeaf{"PrepaidVolumeUsediQuota", flowFeatureData.PrepaidVolumeUsediQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-usedo-quota", types.YLeaf{"PrepaidVolumeUsedoQuota", flowFeatureData.PrepaidVolumeUsedoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-refi-quota", types.YLeaf{"PrepaidVolumeRefiQuota", flowFeatureData.PrepaidVolumeRefiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-refo-quota", types.YLeaf{"PrepaidVolumeRefoQuota", flowFeatureData.PrepaidVolumeRefoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-refb-quota", types.YLeaf{"PrepaidVolumeRefbQuota", flowFeatureData.PrepaidVolumeRefbQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-agg-refi-quota", types.YLeaf{"PrepaidVolumeAggRefiQuota", flowFeatureData.PrepaidVolumeAggRefiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-agg-refo-quota", types.YLeaf{"PrepaidVolumeAggRefoQuota", flowFeatureData.PrepaidVolumeAggRefoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-agg-refb-quota", types.YLeaf{"PrepaidVolumeAggRefbQuota", flowFeatureData.PrepaidVolumeAggRefbQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-newi-quota", types.YLeaf{"PrepaidVolumeNewiQuota", flowFeatureData.PrepaidVolumeNewiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-newo-quota", types.YLeaf{"PrepaidVolumeNewoQuota", flowFeatureData.PrepaidVolumeNewoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-newb-quota", types.YLeaf{"PrepaidVolumeNewbQuota", flowFeatureData.PrepaidVolumeNewbQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-tariff-volumei-quota", types.YLeaf{"PrepaidTariffVolumeiQuota", flowFeatureData.PrepaidTariffVolumeiQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-tariff-volumeo-quota", types.YLeaf{"PrepaidTariffVolumeoQuota", flowFeatureData.PrepaidTariffVolumeoQuota})
    flowFeatureData.EntityData.Leafs.Append("prepaid-tariff-volumeb-quota", types.YLeaf{"PrepaidTariffVolumebQuota", flowFeatureData.PrepaidTariffVolumebQuota})
    flowFeatureData.EntityData.Leafs.Append("flow-accounting-method-list-name", types.YLeaf{"FlowAccountingMethodListName", flowFeatureData.FlowAccountingMethodListName})
    flowFeatureData.EntityData.Leafs.Append("prepaid-cfg", types.YLeaf{"PrepaidCfg", flowFeatureData.PrepaidCfg})
    flowFeatureData.EntityData.Leafs.Append("prepaid-time-state", types.YLeaf{"PrepaidTimeState", flowFeatureData.PrepaidTimeState})
    flowFeatureData.EntityData.Leafs.Append("prepaid-volume-state", types.YLeaf{"PrepaidVolumeState", flowFeatureData.PrepaidVolumeState})
    flowFeatureData.EntityData.Leafs.Append("prepaid-charging-rule", types.YLeaf{"PrepaidChargingRule", flowFeatureData.PrepaidChargingRule})

    flowFeatureData.EntityData.YListKeys = []string {}

    return &(flowFeatureData.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures
// Subscriber accounting SRG feature data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Display accounting SRG features by unique access interface-name. The type
    // is slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature.
    SubscriberAccountingSrgFeature []*SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature
}

func (subscriberAccountingSrgFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures) GetEntityData() *types.CommonEntityData {
    subscriberAccountingSrgFeatures.EntityData.YFilter = subscriberAccountingSrgFeatures.YFilter
    subscriberAccountingSrgFeatures.EntityData.YangName = "subscriber-accounting-srg-features"
    subscriberAccountingSrgFeatures.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingSrgFeatures.EntityData.ParentYangName = "node"
    subscriberAccountingSrgFeatures.EntityData.SegmentPath = "subscriber-accounting-srg-features"
    subscriberAccountingSrgFeatures.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/" + subscriberAccountingSrgFeatures.EntityData.SegmentPath
    subscriberAccountingSrgFeatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingSrgFeatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingSrgFeatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingSrgFeatures.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingSrgFeatures.EntityData.Children.Append("subscriber-accounting-srg-feature", types.YChild{"SubscriberAccountingSrgFeature", nil})
    for i := range subscriberAccountingSrgFeatures.SubscriberAccountingSrgFeature {
        subscriberAccountingSrgFeatures.EntityData.Children.Append(types.GetSegmentPath(subscriberAccountingSrgFeatures.SubscriberAccountingSrgFeature[i]), types.YChild{"SubscriberAccountingSrgFeature", subscriberAccountingSrgFeatures.SubscriberAccountingSrgFeature[i]})
    }
    subscriberAccountingSrgFeatures.EntityData.Leafs = types.NewOrderedMap()

    subscriberAccountingSrgFeatures.EntityData.YListKeys = []string {}

    return &(subscriberAccountingSrgFeatures.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature
// Display accounting SRG features by unique
// access interface-name
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Accounting's SRG Statistics Info Bag.
    SrgStatusCounters SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature_SrgStatusCounters
}

func (subscriberAccountingSrgFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature) GetEntityData() *types.CommonEntityData {
    subscriberAccountingSrgFeature.EntityData.YFilter = subscriberAccountingSrgFeature.YFilter
    subscriberAccountingSrgFeature.EntityData.YangName = "subscriber-accounting-srg-feature"
    subscriberAccountingSrgFeature.EntityData.BundleName = "cisco_ios_xr"
    subscriberAccountingSrgFeature.EntityData.ParentYangName = "subscriber-accounting-srg-features"
    subscriberAccountingSrgFeature.EntityData.SegmentPath = "subscriber-accounting-srg-feature" + types.AddKeyToken(subscriberAccountingSrgFeature.InterfaceName, "interface-name")
    subscriberAccountingSrgFeature.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-srg-features/" + subscriberAccountingSrgFeature.EntityData.SegmentPath
    subscriberAccountingSrgFeature.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAccountingSrgFeature.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAccountingSrgFeature.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAccountingSrgFeature.EntityData.Children = types.NewOrderedMap()
    subscriberAccountingSrgFeature.EntityData.Children.Append("srg-status-counters", types.YChild{"SrgStatusCounters", &subscriberAccountingSrgFeature.SrgStatusCounters})
    subscriberAccountingSrgFeature.EntityData.Leafs = types.NewOrderedMap()
    subscriberAccountingSrgFeature.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", subscriberAccountingSrgFeature.InterfaceName})

    subscriberAccountingSrgFeature.EntityData.YListKeys = []string {"InterfaceName"}

    return &(subscriberAccountingSrgFeature.EntityData)
}

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature_SrgStatusCounters
// Accounting's SRG Statistics Info Bag
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature_SrgStatusCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    InterfaceName interface{}

    // Access Interface handle. The type is interface{} with range: 0..4294967295.
    AccessInterfaceHandle interface{}

    // Accounting Feature Mode. The type is interface{} with range: 0..4294967295.
    AccountingMode interface{}

    // SRG Role for given access interface. The type is interface{} with range:
    // 0..4294967295.
    SrgNodeRole interface{}

    // True/False,if SRG Master Node Sync is still pending. The type is bool.
    SrgMasterEomsPendingFlag interface{}

    // True/False,SRG Node Diameter Notification flag. The type is bool.
    SrgDiameterUpFlag interface{}

    // SRG Role Notification time in DDD MMM DD HH:MM :SS YYYY format eg: Tue Apr
    // 11 21:30:47 2011. The type is string.
    SrgRoleUpdTime interface{}
}

func (srgStatusCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSrgFeatures_SubscriberAccountingSrgFeature_SrgStatusCounters) GetEntityData() *types.CommonEntityData {
    srgStatusCounters.EntityData.YFilter = srgStatusCounters.YFilter
    srgStatusCounters.EntityData.YangName = "srg-status-counters"
    srgStatusCounters.EntityData.BundleName = "cisco_ios_xr"
    srgStatusCounters.EntityData.ParentYangName = "subscriber-accounting-srg-feature"
    srgStatusCounters.EntityData.SegmentPath = "srg-status-counters"
    srgStatusCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting/nodes/node/subscriber-accounting-srg-features/subscriber-accounting-srg-feature/" + srgStatusCounters.EntityData.SegmentPath
    srgStatusCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgStatusCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgStatusCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgStatusCounters.EntityData.Children = types.NewOrderedMap()
    srgStatusCounters.EntityData.Leafs = types.NewOrderedMap()
    srgStatusCounters.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", srgStatusCounters.InterfaceName})
    srgStatusCounters.EntityData.Leafs.Append("access-interface-handle", types.YLeaf{"AccessInterfaceHandle", srgStatusCounters.AccessInterfaceHandle})
    srgStatusCounters.EntityData.Leafs.Append("accounting-mode", types.YLeaf{"AccountingMode", srgStatusCounters.AccountingMode})
    srgStatusCounters.EntityData.Leafs.Append("srg-node-role", types.YLeaf{"SrgNodeRole", srgStatusCounters.SrgNodeRole})
    srgStatusCounters.EntityData.Leafs.Append("srg-master-eoms-pending-flag", types.YLeaf{"SrgMasterEomsPendingFlag", srgStatusCounters.SrgMasterEomsPendingFlag})
    srgStatusCounters.EntityData.Leafs.Append("srg-diameter-up-flag", types.YLeaf{"SrgDiameterUpFlag", srgStatusCounters.SrgDiameterUpFlag})
    srgStatusCounters.EntityData.Leafs.Append("srg-role-upd-time", types.YLeaf{"SrgRoleUpdTime", srgStatusCounters.SrgRoleUpdTime})

    srgStatusCounters.EntityData.YListKeys = []string {}

    return &(srgStatusCounters.EntityData)
}

