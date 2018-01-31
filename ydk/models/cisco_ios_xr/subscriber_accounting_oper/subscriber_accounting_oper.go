// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-accounting package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-accounting: Subscriber accounting operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber accounting operational data for a particular location.
    Nodes SubscriberAccounting_Nodes
}

func (subscriberAccounting *SubscriberAccounting) GetFilter() yfilter.YFilter { return subscriberAccounting.YFilter }

func (subscriberAccounting *SubscriberAccounting) SetFilter(yf yfilter.YFilter) { subscriberAccounting.YFilter = yf }

func (subscriberAccounting *SubscriberAccounting) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (subscriberAccounting *SubscriberAccounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-accounting-oper:subscriber-accounting"
}

func (subscriberAccounting *SubscriberAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &subscriberAccounting.Nodes
    }
    return nil
}

func (subscriberAccounting *SubscriberAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &subscriberAccounting.Nodes
    return children
}

func (subscriberAccounting *SubscriberAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberAccounting *SubscriberAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccounting *SubscriberAccounting) GetYangName() string { return "subscriber-accounting" }

func (subscriberAccounting *SubscriberAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccounting *SubscriberAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccounting *SubscriberAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccounting *SubscriberAccounting) SetParent(parent types.Entity) { subscriberAccounting.parent = parent }

func (subscriberAccounting *SubscriberAccounting) GetParent() types.Entity { return subscriberAccounting.parent }

func (subscriberAccounting *SubscriberAccounting) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-accounting-oper" }

// SubscriberAccounting_Nodes
// Subscriber accounting operational data for a
// particular location
type SubscriberAccounting_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location. For example, 0/1/CPU0. The type is slice of
    // SubscriberAccounting_Nodes_Node.
    Node []SubscriberAccounting_Nodes_Node
}

func (nodes *SubscriberAccounting_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SubscriberAccounting_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SubscriberAccounting_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SubscriberAccounting_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SubscriberAccounting_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberAccounting_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SubscriberAccounting_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SubscriberAccounting_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SubscriberAccounting_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SubscriberAccounting_Nodes) GetYangName() string { return "nodes" }

func (nodes *SubscriberAccounting_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SubscriberAccounting_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SubscriberAccounting_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SubscriberAccounting_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SubscriberAccounting_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SubscriberAccounting_Nodes) GetParentYangName() string { return "subscriber-accounting" }

// SubscriberAccounting_Nodes_Node
// Location. For example, 0/1/CPU0
type SubscriberAccounting_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
}

func (node *SubscriberAccounting_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SubscriberAccounting_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SubscriberAccounting_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "subscriber-accounting-session-features" { return "SubscriberAccountingSessionFeatures" }
    if yname == "subscriber-accounting-summary" { return "SubscriberAccountingSummary" }
    if yname == "subscriber-accounting-flow-features" { return "SubscriberAccountingFlowFeatures" }
    return ""
}

func (node *SubscriberAccounting_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *SubscriberAccounting_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscriber-accounting-session-features" {
        return &node.SubscriberAccountingSessionFeatures
    }
    if childYangName == "subscriber-accounting-summary" {
        return &node.SubscriberAccountingSummary
    }
    if childYangName == "subscriber-accounting-flow-features" {
        return &node.SubscriberAccountingFlowFeatures
    }
    return nil
}

func (node *SubscriberAccounting_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["subscriber-accounting-session-features"] = &node.SubscriberAccountingSessionFeatures
    children["subscriber-accounting-summary"] = &node.SubscriberAccountingSummary
    children["subscriber-accounting-flow-features"] = &node.SubscriberAccountingFlowFeatures
    return children
}

func (node *SubscriberAccounting_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *SubscriberAccounting_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SubscriberAccounting_Nodes_Node) GetYangName() string { return "node" }

func (node *SubscriberAccounting_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SubscriberAccounting_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SubscriberAccounting_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SubscriberAccounting_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SubscriberAccounting_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SubscriberAccounting_Nodes_Node) GetParentYangName() string { return "nodes" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures
// Subscriber accounting session feature data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Display accounting session features by unique subscriber label. The type is
    // slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature.
    SubscriberAccountingSessionFeature []SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetFilter() yfilter.YFilter { return subscriberAccountingSessionFeatures.YFilter }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) SetFilter(yf yfilter.YFilter) { subscriberAccountingSessionFeatures.YFilter = yf }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetGoName(yname string) string {
    if yname == "subscriber-accounting-session-feature" { return "SubscriberAccountingSessionFeature" }
    return ""
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetSegmentPath() string {
    return "subscriber-accounting-session-features"
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscriber-accounting-session-feature" {
        for _, c := range subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature {
            if subscriberAccountingSessionFeatures.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature{}
        subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature = append(subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature, child)
        return &subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature[len(subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature)-1]
    }
    return nil
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature {
        children[subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature[i].GetSegmentPath()] = &subscriberAccountingSessionFeatures.SubscriberAccountingSessionFeature[i]
    }
    return children
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetYangName() string { return "subscriber-accounting-session-features" }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) SetParent(parent types.Entity) { subscriberAccountingSessionFeatures.parent = parent }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetParent() types.Entity { return subscriberAccountingSessionFeatures.parent }

func (subscriberAccountingSessionFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures) GetParentYangName() string { return "node" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature
// Display accounting session features by unique
// subscriber label
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Unique subscriber label. The type is interface{}
    // with range: -2147483648..2147483647.
    SubLabel interface{}

    // Accounting session feature display data.
    SessionFeatureData SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetFilter() yfilter.YFilter { return subscriberAccountingSessionFeature.YFilter }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) SetFilter(yf yfilter.YFilter) { subscriberAccountingSessionFeature.YFilter = yf }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetGoName(yname string) string {
    if yname == "sub-label" { return "SubLabel" }
    if yname == "session-feature-data" { return "SessionFeatureData" }
    return ""
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetSegmentPath() string {
    return "subscriber-accounting-session-feature" + "[sub-label='" + fmt.Sprintf("%v", subscriberAccountingSessionFeature.SubLabel) + "']"
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-feature-data" {
        return &subscriberAccountingSessionFeature.SessionFeatureData
    }
    return nil
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-feature-data"] = &subscriberAccountingSessionFeature.SessionFeatureData
    return children
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sub-label"] = subscriberAccountingSessionFeature.SubLabel
    return leafs
}

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetYangName() string { return "subscriber-accounting-session-feature" }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) SetParent(parent types.Entity) { subscriberAccountingSessionFeature.parent = parent }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetParent() types.Entity { return subscriberAccountingSessionFeature.parent }

func (subscriberAccountingSessionFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature) GetParentYangName() string { return "subscriber-accounting-session-features" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData
// Accounting session feature display data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData struct {
    parent types.Entity
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
    ServiceAccountingFeature []SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetFilter() yfilter.YFilter { return sessionFeatureData.YFilter }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) SetFilter(yf yfilter.YFilter) { sessionFeatureData.YFilter = yf }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetGoName(yname string) string {
    if yname == "unique-subscriber-label" { return "UniqueSubscriberLabel" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "session-disconnected" { return "SessionDisconnected" }
    if yname == "session-accounting-enabled-flag" { return "SessionAccountingEnabledFlag" }
    if yname == "session-accounting-method-list" { return "SessionAccountingMethodList" }
    if yname == "session-accounting-periodic-interval" { return "SessionAccountingPeriodicInterval" }
    if yname == "session-accounting-aaa-trans-pending" { return "SessionAccountingAaaTransPending" }
    if yname == "session-accounting-aaa-request-failed" { return "SessionAccountingAaaRequestFailed" }
    if yname == "session-accounting-started" { return "SessionAccountingStarted" }
    if yname == "session-idle-timeout-enabled-flag" { return "SessionIdleTimeoutEnabledFlag" }
    if yname == "idle-timeout-value" { return "IdleTimeoutValue" }
    if yname == "idle-timeout-threshold" { return "IdleTimeoutThreshold" }
    if yname == "idle-timeout-direction" { return "IdleTimeoutDirection" }
    if yname == "session-is-idle" { return "SessionIsIdle" }
    if yname == "session-stats-changed-time" { return "SessionStatsChangedTime" }
    if yname == "session-total-idle-time" { return "SessionTotalIdleTime" }
    if yname == "session-to-idle-count" { return "SessionToIdleCount" }
    if yname == "session-to-awake-count" { return "SessionToAwakeCount" }
    if yname == "session-idle-to-aaa-trans-pending" { return "SessionIdleToAaaTransPending" }
    if yname == "session-idle-to-aaa-request-failed" { return "SessionIdleToAaaRequestFailed" }
    if yname == "session-timeout-enabled-flag" { return "SessionTimeoutEnabledFlag" }
    if yname == "session-timeout-value" { return "SessionTimeoutValue" }
    if yname == "session-timeout-time-remaining" { return "SessionTimeoutTimeRemaining" }
    if yname == "service-accounting-feature" { return "ServiceAccountingFeature" }
    return ""
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetSegmentPath() string {
    return "session-feature-data"
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-accounting-feature" {
        for _, c := range sessionFeatureData.ServiceAccountingFeature {
            if sessionFeatureData.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature{}
        sessionFeatureData.ServiceAccountingFeature = append(sessionFeatureData.ServiceAccountingFeature, child)
        return &sessionFeatureData.ServiceAccountingFeature[len(sessionFeatureData.ServiceAccountingFeature)-1]
    }
    return nil
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessionFeatureData.ServiceAccountingFeature {
        children[sessionFeatureData.ServiceAccountingFeature[i].GetSegmentPath()] = &sessionFeatureData.ServiceAccountingFeature[i]
    }
    return children
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unique-subscriber-label"] = sessionFeatureData.UniqueSubscriberLabel
    leafs["interface-handle"] = sessionFeatureData.InterfaceHandle
    leafs["session-disconnected"] = sessionFeatureData.SessionDisconnected
    leafs["session-accounting-enabled-flag"] = sessionFeatureData.SessionAccountingEnabledFlag
    leafs["session-accounting-method-list"] = sessionFeatureData.SessionAccountingMethodList
    leafs["session-accounting-periodic-interval"] = sessionFeatureData.SessionAccountingPeriodicInterval
    leafs["session-accounting-aaa-trans-pending"] = sessionFeatureData.SessionAccountingAaaTransPending
    leafs["session-accounting-aaa-request-failed"] = sessionFeatureData.SessionAccountingAaaRequestFailed
    leafs["session-accounting-started"] = sessionFeatureData.SessionAccountingStarted
    leafs["session-idle-timeout-enabled-flag"] = sessionFeatureData.SessionIdleTimeoutEnabledFlag
    leafs["idle-timeout-value"] = sessionFeatureData.IdleTimeoutValue
    leafs["idle-timeout-threshold"] = sessionFeatureData.IdleTimeoutThreshold
    leafs["idle-timeout-direction"] = sessionFeatureData.IdleTimeoutDirection
    leafs["session-is-idle"] = sessionFeatureData.SessionIsIdle
    leafs["session-stats-changed-time"] = sessionFeatureData.SessionStatsChangedTime
    leafs["session-total-idle-time"] = sessionFeatureData.SessionTotalIdleTime
    leafs["session-to-idle-count"] = sessionFeatureData.SessionToIdleCount
    leafs["session-to-awake-count"] = sessionFeatureData.SessionToAwakeCount
    leafs["session-idle-to-aaa-trans-pending"] = sessionFeatureData.SessionIdleToAaaTransPending
    leafs["session-idle-to-aaa-request-failed"] = sessionFeatureData.SessionIdleToAaaRequestFailed
    leafs["session-timeout-enabled-flag"] = sessionFeatureData.SessionTimeoutEnabledFlag
    leafs["session-timeout-value"] = sessionFeatureData.SessionTimeoutValue
    leafs["session-timeout-time-remaining"] = sessionFeatureData.SessionTimeoutTimeRemaining
    return leafs
}

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetBundleName() string { return "cisco_ios_xr" }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetYangName() string { return "session-feature-data" }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) SetParent(parent types.Entity) { sessionFeatureData.parent = parent }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetParent() types.Entity { return sessionFeatureData.parent }

func (sessionFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData) GetParentYangName() string { return "subscriber-accounting-session-feature" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature
// List of service accounting features
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetFilter() yfilter.YFilter { return serviceAccountingFeature.YFilter }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) SetFilter(yf yfilter.YFilter) { serviceAccountingFeature.YFilter = yf }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetGoName(yname string) string {
    if yname == "service-accounting-enabled-flag" { return "ServiceAccountingEnabledFlag" }
    if yname == "service-accounting-service-id" { return "ServiceAccountingServiceId" }
    if yname == "service-accounting-method-list" { return "ServiceAccountingMethodList" }
    if yname == "service-accounting-periodic-interval" { return "ServiceAccountingPeriodicInterval" }
    if yname == "session-accounting-aaa-trans-pending" { return "SessionAccountingAaaTransPending" }
    if yname == "session-accounting-aaa-request-failed" { return "SessionAccountingAaaRequestFailed" }
    if yname == "session-accounting-started" { return "SessionAccountingStarted" }
    return ""
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetSegmentPath() string {
    return "service-accounting-feature"
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-accounting-enabled-flag"] = serviceAccountingFeature.ServiceAccountingEnabledFlag
    leafs["service-accounting-service-id"] = serviceAccountingFeature.ServiceAccountingServiceId
    leafs["service-accounting-method-list"] = serviceAccountingFeature.ServiceAccountingMethodList
    leafs["service-accounting-periodic-interval"] = serviceAccountingFeature.ServiceAccountingPeriodicInterval
    leafs["session-accounting-aaa-trans-pending"] = serviceAccountingFeature.SessionAccountingAaaTransPending
    leafs["session-accounting-aaa-request-failed"] = serviceAccountingFeature.SessionAccountingAaaRequestFailed
    leafs["session-accounting-started"] = serviceAccountingFeature.SessionAccountingStarted
    return leafs
}

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetYangName() string { return "service-accounting-feature" }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) SetParent(parent types.Entity) { serviceAccountingFeature.parent = parent }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetParent() types.Entity { return serviceAccountingFeature.parent }

func (serviceAccountingFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingSessionFeatures_SubscriberAccountingSessionFeature_SessionFeatureData_ServiceAccountingFeature) GetParentYangName() string { return "session-feature-data" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary
// Display subscriber accounting summary data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary struct {
    parent types.Entity
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

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetFilter() yfilter.YFilter { return subscriberAccountingSummary.YFilter }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) SetFilter(yf yfilter.YFilter) { subscriberAccountingSummary.YFilter = yf }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetGoName(yname string) string {
    if yname == "aaa-counters" { return "AaaCounters" }
    if yname == "idle-timeout-counters" { return "IdleTimeoutCounters" }
    if yname == "session-timeout-counters" { return "SessionTimeoutCounters" }
    if yname == "session-flow-counters" { return "SessionFlowCounters" }
    return ""
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetSegmentPath() string {
    return "subscriber-accounting-summary"
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aaa-counters" {
        return &subscriberAccountingSummary.AaaCounters
    }
    if childYangName == "idle-timeout-counters" {
        return &subscriberAccountingSummary.IdleTimeoutCounters
    }
    if childYangName == "session-timeout-counters" {
        return &subscriberAccountingSummary.SessionTimeoutCounters
    }
    if childYangName == "session-flow-counters" {
        return &subscriberAccountingSummary.SessionFlowCounters
    }
    return nil
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aaa-counters"] = &subscriberAccountingSummary.AaaCounters
    children["idle-timeout-counters"] = &subscriberAccountingSummary.IdleTimeoutCounters
    children["session-timeout-counters"] = &subscriberAccountingSummary.SessionTimeoutCounters
    children["session-flow-counters"] = &subscriberAccountingSummary.SessionFlowCounters
    return children
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetYangName() string { return "subscriber-accounting-summary" }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) SetParent(parent types.Entity) { subscriberAccountingSummary.parent = parent }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetParent() types.Entity { return subscriberAccountingSummary.parent }

func (subscriberAccountingSummary *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary) GetParentYangName() string { return "node" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters
// Accounting feature AAA summary counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters struct {
    parent types.Entity
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

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetFilter() yfilter.YFilter { return aaaCounters.YFilter }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) SetFilter(yf yfilter.YFilter) { aaaCounters.YFilter = yf }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetGoName(yname string) string {
    if yname == "flow-start" { return "FlowStart" }
    if yname == "flow-disconnect" { return "FlowDisconnect" }
    if yname == "session-accounting-start" { return "SessionAccountingStart" }
    if yname == "session-accounting-stop" { return "SessionAccountingStop" }
    if yname == "session-accounting-update" { return "SessionAccountingUpdate" }
    if yname == "service-accounting-start" { return "ServiceAccountingStart" }
    if yname == "service-accounting-stop" { return "ServiceAccountingStop" }
    if yname == "service-accounting-update" { return "ServiceAccountingUpdate" }
    if yname == "flow-accounting-start" { return "FlowAccountingStart" }
    if yname == "flow-accounting-stop" { return "FlowAccountingStop" }
    if yname == "flow-accounting-update" { return "FlowAccountingUpdate" }
    if yname == "accounting-callback" { return "AccountingCallback" }
    if yname == "session-acct-trans-pending" { return "SessionAcctTransPending" }
    if yname == "session-acct-reqs-failed" { return "SessionAcctReqsFailed" }
    if yname == "session-acct-out-of-sync" { return "SessionAcctOutOfSync" }
    if yname == "session-idle-to-trans-pending" { return "SessionIdleToTransPending" }
    if yname == "session-idle-to-reqs-failed" { return "SessionIdleToReqsFailed" }
    if yname == "session-idle-to-out-of-sync" { return "SessionIdleToOutOfSync" }
    if yname == "service-acct-trans-pending" { return "ServiceAcctTransPending" }
    if yname == "service-acct-reqs-failed" { return "ServiceAcctReqsFailed" }
    if yname == "service-acct-out-of-sync" { return "ServiceAcctOutOfSync" }
    if yname == "service-idle-to-trans-pending" { return "ServiceIdleToTransPending" }
    if yname == "service-idle-to-reqs-failed" { return "ServiceIdleToReqsFailed" }
    if yname == "service-idle-to-out-of-sync" { return "ServiceIdleToOutOfSync" }
    if yname == "prepaid-start" { return "PrepaidStart" }
    if yname == "prepaid-stop" { return "PrepaidStop" }
    if yname == "prepaid-accounting-start" { return "PrepaidAccountingStart" }
    if yname == "prepaid-accounting-stop" { return "PrepaidAccountingStop" }
    if yname == "prepaid-volume-threshold-reached" { return "PrepaidVolumeThresholdReached" }
    if yname == "prepaid-time-threshold-reached" { return "PrepaidTimeThresholdReached" }
    if yname == "prepaid-quota-depleted" { return "PrepaidQuotaDepleted" }
    if yname == "prepaid-reauthorization" { return "PrepaidReauthorization" }
    if yname == "idle-timeout" { return "IdleTimeout" }
    if yname == "idle-timeout-response-callback" { return "IdleTimeoutResponseCallback" }
    if yname == "owned-resource-start" { return "OwnedResourceStart" }
    return ""
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetSegmentPath() string {
    return "aaa-counters"
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-start"] = aaaCounters.FlowStart
    leafs["flow-disconnect"] = aaaCounters.FlowDisconnect
    leafs["session-accounting-start"] = aaaCounters.SessionAccountingStart
    leafs["session-accounting-stop"] = aaaCounters.SessionAccountingStop
    leafs["session-accounting-update"] = aaaCounters.SessionAccountingUpdate
    leafs["service-accounting-start"] = aaaCounters.ServiceAccountingStart
    leafs["service-accounting-stop"] = aaaCounters.ServiceAccountingStop
    leafs["service-accounting-update"] = aaaCounters.ServiceAccountingUpdate
    leafs["flow-accounting-start"] = aaaCounters.FlowAccountingStart
    leafs["flow-accounting-stop"] = aaaCounters.FlowAccountingStop
    leafs["flow-accounting-update"] = aaaCounters.FlowAccountingUpdate
    leafs["accounting-callback"] = aaaCounters.AccountingCallback
    leafs["session-acct-trans-pending"] = aaaCounters.SessionAcctTransPending
    leafs["session-acct-reqs-failed"] = aaaCounters.SessionAcctReqsFailed
    leafs["session-acct-out-of-sync"] = aaaCounters.SessionAcctOutOfSync
    leafs["session-idle-to-trans-pending"] = aaaCounters.SessionIdleToTransPending
    leafs["session-idle-to-reqs-failed"] = aaaCounters.SessionIdleToReqsFailed
    leafs["session-idle-to-out-of-sync"] = aaaCounters.SessionIdleToOutOfSync
    leafs["service-acct-trans-pending"] = aaaCounters.ServiceAcctTransPending
    leafs["service-acct-reqs-failed"] = aaaCounters.ServiceAcctReqsFailed
    leafs["service-acct-out-of-sync"] = aaaCounters.ServiceAcctOutOfSync
    leafs["service-idle-to-trans-pending"] = aaaCounters.ServiceIdleToTransPending
    leafs["service-idle-to-reqs-failed"] = aaaCounters.ServiceIdleToReqsFailed
    leafs["service-idle-to-out-of-sync"] = aaaCounters.ServiceIdleToOutOfSync
    leafs["prepaid-start"] = aaaCounters.PrepaidStart
    leafs["prepaid-stop"] = aaaCounters.PrepaidStop
    leafs["prepaid-accounting-start"] = aaaCounters.PrepaidAccountingStart
    leafs["prepaid-accounting-stop"] = aaaCounters.PrepaidAccountingStop
    leafs["prepaid-volume-threshold-reached"] = aaaCounters.PrepaidVolumeThresholdReached
    leafs["prepaid-time-threshold-reached"] = aaaCounters.PrepaidTimeThresholdReached
    leafs["prepaid-quota-depleted"] = aaaCounters.PrepaidQuotaDepleted
    leafs["prepaid-reauthorization"] = aaaCounters.PrepaidReauthorization
    leafs["idle-timeout"] = aaaCounters.IdleTimeout
    leafs["idle-timeout-response-callback"] = aaaCounters.IdleTimeoutResponseCallback
    leafs["owned-resource-start"] = aaaCounters.OwnedResourceStart
    return leafs
}

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetBundleName() string { return "cisco_ios_xr" }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetYangName() string { return "aaa-counters" }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) SetParent(parent types.Entity) { aaaCounters.parent = parent }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetParent() types.Entity { return aaaCounters.parent }

func (aaaCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_AaaCounters) GetParentYangName() string { return "subscriber-accounting-summary" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters
// Accounting feature idle timeout summary counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters struct {
    parent types.Entity
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

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetFilter() yfilter.YFilter { return idleTimeoutCounters.YFilter }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) SetFilter(yf yfilter.YFilter) { idleTimeoutCounters.YFilter = yf }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetGoName(yname string) string {
    if yname == "active-session-idle-timers" { return "ActiveSessionIdleTimers" }
    if yname == "idle-sessions" { return "IdleSessions" }
    if yname == "transitions-to-idle" { return "TransitionsToIdle" }
    if yname == "transitions-to-awake" { return "TransitionsToAwake" }
    if yname == "active-flow-idle-timers" { return "ActiveFlowIdleTimers" }
    if yname == "expired-flow-idle-timers" { return "ExpiredFlowIdleTimers" }
    if yname == "active-prepaid-idle-timers" { return "ActivePrepaidIdleTimers" }
    if yname == "expired-prepaid-idle-timers" { return "ExpiredPrepaidIdleTimers" }
    return ""
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetSegmentPath() string {
    return "idle-timeout-counters"
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-session-idle-timers"] = idleTimeoutCounters.ActiveSessionIdleTimers
    leafs["idle-sessions"] = idleTimeoutCounters.IdleSessions
    leafs["transitions-to-idle"] = idleTimeoutCounters.TransitionsToIdle
    leafs["transitions-to-awake"] = idleTimeoutCounters.TransitionsToAwake
    leafs["active-flow-idle-timers"] = idleTimeoutCounters.ActiveFlowIdleTimers
    leafs["expired-flow-idle-timers"] = idleTimeoutCounters.ExpiredFlowIdleTimers
    leafs["active-prepaid-idle-timers"] = idleTimeoutCounters.ActivePrepaidIdleTimers
    leafs["expired-prepaid-idle-timers"] = idleTimeoutCounters.ExpiredPrepaidIdleTimers
    return leafs
}

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetBundleName() string { return "cisco_ios_xr" }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetYangName() string { return "idle-timeout-counters" }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) SetParent(parent types.Entity) { idleTimeoutCounters.parent = parent }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetParent() types.Entity { return idleTimeoutCounters.parent }

func (idleTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_IdleTimeoutCounters) GetParentYangName() string { return "subscriber-accounting-summary" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters
// Accounting feature session timeout summary
// counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of Active Session Timers. The type is interface{} with range:
    // 0..4294967295.
    ActiveSessionTimers interface{}

    // Number of Expired Session Timers. The type is interface{} with range:
    // 0..4294967295.
    ExpiredSessionTimers interface{}
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetFilter() yfilter.YFilter { return sessionTimeoutCounters.YFilter }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) SetFilter(yf yfilter.YFilter) { sessionTimeoutCounters.YFilter = yf }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetGoName(yname string) string {
    if yname == "active-session-timers" { return "ActiveSessionTimers" }
    if yname == "expired-session-timers" { return "ExpiredSessionTimers" }
    return ""
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetSegmentPath() string {
    return "session-timeout-counters"
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-session-timers"] = sessionTimeoutCounters.ActiveSessionTimers
    leafs["expired-session-timers"] = sessionTimeoutCounters.ExpiredSessionTimers
    return leafs
}

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetBundleName() string { return "cisco_ios_xr" }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetYangName() string { return "session-timeout-counters" }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) SetParent(parent types.Entity) { sessionTimeoutCounters.parent = parent }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetParent() types.Entity { return sessionTimeoutCounters.parent }

func (sessionTimeoutCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionTimeoutCounters) GetParentYangName() string { return "subscriber-accounting-summary" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters
// Accounting feature session context summary
// counters
type SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters struct {
    parent types.Entity
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

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetFilter() yfilter.YFilter { return sessionFlowCounters.YFilter }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) SetFilter(yf yfilter.YFilter) { sessionFlowCounters.YFilter = yf }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetGoName(yname string) string {
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "disconnected-sessions" { return "DisconnectedSessions" }
    if yname == "active-session-accounting-sessions" { return "ActiveSessionAccountingSessions" }
    if yname == "active-flows" { return "ActiveFlows" }
    if yname == "quota-received" { return "QuotaReceived" }
    return ""
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetSegmentPath() string {
    return "session-flow-counters"
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-sessions"] = sessionFlowCounters.ActiveSessions
    leafs["disconnected-sessions"] = sessionFlowCounters.DisconnectedSessions
    leafs["active-session-accounting-sessions"] = sessionFlowCounters.ActiveSessionAccountingSessions
    leafs["active-flows"] = sessionFlowCounters.ActiveFlows
    leafs["quota-received"] = sessionFlowCounters.QuotaReceived
    return leafs
}

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetBundleName() string { return "cisco_ios_xr" }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetYangName() string { return "session-flow-counters" }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) SetParent(parent types.Entity) { sessionFlowCounters.parent = parent }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetParent() types.Entity { return sessionFlowCounters.parent }

func (sessionFlowCounters *SubscriberAccounting_Nodes_Node_SubscriberAccountingSummary_SessionFlowCounters) GetParentYangName() string { return "subscriber-accounting-summary" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures
// Subscriber accounting flow feature data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Display accounting flow features by unique subscriber label. The type is
    // slice of
    // SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature.
    SubscriberAccountingFlowFeature []SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetFilter() yfilter.YFilter { return subscriberAccountingFlowFeatures.YFilter }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) SetFilter(yf yfilter.YFilter) { subscriberAccountingFlowFeatures.YFilter = yf }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetGoName(yname string) string {
    if yname == "subscriber-accounting-flow-feature" { return "SubscriberAccountingFlowFeature" }
    return ""
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetSegmentPath() string {
    return "subscriber-accounting-flow-features"
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscriber-accounting-flow-feature" {
        for _, c := range subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature {
            if subscriberAccountingFlowFeatures.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature{}
        subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature = append(subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature, child)
        return &subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature[len(subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature)-1]
    }
    return nil
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature {
        children[subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature[i].GetSegmentPath()] = &subscriberAccountingFlowFeatures.SubscriberAccountingFlowFeature[i]
    }
    return children
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetYangName() string { return "subscriber-accounting-flow-features" }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) SetParent(parent types.Entity) { subscriberAccountingFlowFeatures.parent = parent }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetParent() types.Entity { return subscriberAccountingFlowFeatures.parent }

func (subscriberAccountingFlowFeatures *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures) GetParentYangName() string { return "node" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature
// Display accounting flow features by unique
// subscriber label
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Unique subscriber class label. The type is
    // interface{} with range: -2147483648..2147483647.
    ClassLabel interface{}

    // Accouting flow feature display data.
    FlowFeatureData SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetFilter() yfilter.YFilter { return subscriberAccountingFlowFeature.YFilter }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) SetFilter(yf yfilter.YFilter) { subscriberAccountingFlowFeature.YFilter = yf }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetGoName(yname string) string {
    if yname == "class-label" { return "ClassLabel" }
    if yname == "flow-feature-data" { return "FlowFeatureData" }
    return ""
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetSegmentPath() string {
    return "subscriber-accounting-flow-feature" + "[class-label='" + fmt.Sprintf("%v", subscriberAccountingFlowFeature.ClassLabel) + "']"
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-feature-data" {
        return &subscriberAccountingFlowFeature.FlowFeatureData
    }
    return nil
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-feature-data"] = &subscriberAccountingFlowFeature.FlowFeatureData
    return children
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-label"] = subscriberAccountingFlowFeature.ClassLabel
    return leafs
}

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetYangName() string { return "subscriber-accounting-flow-feature" }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) SetParent(parent types.Entity) { subscriberAccountingFlowFeature.parent = parent }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetParent() types.Entity { return subscriberAccountingFlowFeature.parent }

func (subscriberAccountingFlowFeature *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature) GetParentYangName() string { return "subscriber-accounting-flow-features" }

// SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData
// Accouting flow feature display data
type SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData struct {
    parent types.Entity
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

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetFilter() yfilter.YFilter { return flowFeatureData.YFilter }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) SetFilter(yf yfilter.YFilter) { flowFeatureData.YFilter = yf }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetGoName(yname string) string {
    if yname == "flow-accounting-enabled-flag" { return "FlowAccountingEnabledFlag" }
    if yname == "flow-idle-timeout-enabled-flag" { return "FlowIdleTimeoutEnabledFlag" }
    if yname == "prepaid-enabled-flag" { return "PrepaidEnabledFlag" }
    if yname == "prepaid-reauth-timer-enabled" { return "PrepaidReauthTimerEnabled" }
    if yname == "prepaid-idle-timeout-enabled" { return "PrepaidIdleTimeoutEnabled" }
    if yname == "prepaid-final-unit" { return "PrepaidFinalUnit" }
    if yname == "unique-class-label" { return "UniqueClassLabel" }
    if yname == "flow-direction" { return "FlowDirection" }
    if yname == "flow-accounting-periodic-interval" { return "FlowAccountingPeriodicInterval" }
    if yname == "flow-idle-timeout-value" { return "FlowIdleTimeoutValue" }
    if yname == "prepaid-time-quota" { return "PrepaidTimeQuota" }
    if yname == "prepaid-time-threshold" { return "PrepaidTimeThreshold" }
    if yname == "prepaid-total-time-quota" { return "PrepaidTotalTimeQuota" }
    if yname == "prepaid-volume-threshold" { return "PrepaidVolumeThreshold" }
    if yname == "prepaid-remaining-qt" { return "PrepaidRemainingQt" }
    if yname == "prepaid-remaining-qat" { return "PrepaidRemainingQat" }
    if yname == "prepaid-remaining-qit" { return "PrepaidRemainingQit" }
    if yname == "prepaid-remaining-qtt" { return "PrepaidRemainingQtt" }
    if yname == "prepaid-remaining-wheel" { return "PrepaidRemainingWheel" }
    if yname == "prepaid-tariff-time" { return "PrepaidTariffTime" }
    if yname == "prepaid-idle-timeout-value" { return "PrepaidIdleTimeoutValue" }
    if yname == "prepaid-reauth-timeout-value" { return "PrepaidReauthTimeoutValue" }
    if yname == "prepaid-ccfh" { return "PrepaidCcfh" }
    if yname == "prepaid-result-code" { return "PrepaidResultCode" }
    if yname == "prepaid-volumei-quota" { return "PrepaidVolumeiQuota" }
    if yname == "prepaid-volumeo-quota" { return "PrepaidVolumeoQuota" }
    if yname == "prepaid-volumeb-quota" { return "PrepaidVolumebQuota" }
    if yname == "prepaid-total-volumei-quota" { return "PrepaidTotalVolumeiQuota" }
    if yname == "prepaid-total-volumeo-quota" { return "PrepaidTotalVolumeoQuota" }
    if yname == "prepaid-total-volumeb-quota" { return "PrepaidTotalVolumebQuota" }
    if yname == "prepaid-volume-usedi-quota" { return "PrepaidVolumeUsediQuota" }
    if yname == "prepaid-volume-usedo-quota" { return "PrepaidVolumeUsedoQuota" }
    if yname == "prepaid-volume-refi-quota" { return "PrepaidVolumeRefiQuota" }
    if yname == "prepaid-volume-refo-quota" { return "PrepaidVolumeRefoQuota" }
    if yname == "prepaid-volume-refb-quota" { return "PrepaidVolumeRefbQuota" }
    if yname == "prepaid-volume-newi-quota" { return "PrepaidVolumeNewiQuota" }
    if yname == "prepaid-volume-newo-quota" { return "PrepaidVolumeNewoQuota" }
    if yname == "prepaid-volume-newb-quota" { return "PrepaidVolumeNewbQuota" }
    if yname == "prepaid-tariff-volumei-quota" { return "PrepaidTariffVolumeiQuota" }
    if yname == "prepaid-tariff-volumeo-quota" { return "PrepaidTariffVolumeoQuota" }
    if yname == "prepaid-tariff-volumeb-quota" { return "PrepaidTariffVolumebQuota" }
    if yname == "flow-accounting-method-list-name" { return "FlowAccountingMethodListName" }
    if yname == "prepaid-cfg" { return "PrepaidCfg" }
    if yname == "prepaid-time-state" { return "PrepaidTimeState" }
    if yname == "prepaid-volume-state" { return "PrepaidVolumeState" }
    if yname == "prepaid-charging-rule" { return "PrepaidChargingRule" }
    return ""
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetSegmentPath() string {
    return "flow-feature-data"
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-accounting-enabled-flag"] = flowFeatureData.FlowAccountingEnabledFlag
    leafs["flow-idle-timeout-enabled-flag"] = flowFeatureData.FlowIdleTimeoutEnabledFlag
    leafs["prepaid-enabled-flag"] = flowFeatureData.PrepaidEnabledFlag
    leafs["prepaid-reauth-timer-enabled"] = flowFeatureData.PrepaidReauthTimerEnabled
    leafs["prepaid-idle-timeout-enabled"] = flowFeatureData.PrepaidIdleTimeoutEnabled
    leafs["prepaid-final-unit"] = flowFeatureData.PrepaidFinalUnit
    leafs["unique-class-label"] = flowFeatureData.UniqueClassLabel
    leafs["flow-direction"] = flowFeatureData.FlowDirection
    leafs["flow-accounting-periodic-interval"] = flowFeatureData.FlowAccountingPeriodicInterval
    leafs["flow-idle-timeout-value"] = flowFeatureData.FlowIdleTimeoutValue
    leafs["prepaid-time-quota"] = flowFeatureData.PrepaidTimeQuota
    leafs["prepaid-time-threshold"] = flowFeatureData.PrepaidTimeThreshold
    leafs["prepaid-total-time-quota"] = flowFeatureData.PrepaidTotalTimeQuota
    leafs["prepaid-volume-threshold"] = flowFeatureData.PrepaidVolumeThreshold
    leafs["prepaid-remaining-qt"] = flowFeatureData.PrepaidRemainingQt
    leafs["prepaid-remaining-qat"] = flowFeatureData.PrepaidRemainingQat
    leafs["prepaid-remaining-qit"] = flowFeatureData.PrepaidRemainingQit
    leafs["prepaid-remaining-qtt"] = flowFeatureData.PrepaidRemainingQtt
    leafs["prepaid-remaining-wheel"] = flowFeatureData.PrepaidRemainingWheel
    leafs["prepaid-tariff-time"] = flowFeatureData.PrepaidTariffTime
    leafs["prepaid-idle-timeout-value"] = flowFeatureData.PrepaidIdleTimeoutValue
    leafs["prepaid-reauth-timeout-value"] = flowFeatureData.PrepaidReauthTimeoutValue
    leafs["prepaid-ccfh"] = flowFeatureData.PrepaidCcfh
    leafs["prepaid-result-code"] = flowFeatureData.PrepaidResultCode
    leafs["prepaid-volumei-quota"] = flowFeatureData.PrepaidVolumeiQuota
    leafs["prepaid-volumeo-quota"] = flowFeatureData.PrepaidVolumeoQuota
    leafs["prepaid-volumeb-quota"] = flowFeatureData.PrepaidVolumebQuota
    leafs["prepaid-total-volumei-quota"] = flowFeatureData.PrepaidTotalVolumeiQuota
    leafs["prepaid-total-volumeo-quota"] = flowFeatureData.PrepaidTotalVolumeoQuota
    leafs["prepaid-total-volumeb-quota"] = flowFeatureData.PrepaidTotalVolumebQuota
    leafs["prepaid-volume-usedi-quota"] = flowFeatureData.PrepaidVolumeUsediQuota
    leafs["prepaid-volume-usedo-quota"] = flowFeatureData.PrepaidVolumeUsedoQuota
    leafs["prepaid-volume-refi-quota"] = flowFeatureData.PrepaidVolumeRefiQuota
    leafs["prepaid-volume-refo-quota"] = flowFeatureData.PrepaidVolumeRefoQuota
    leafs["prepaid-volume-refb-quota"] = flowFeatureData.PrepaidVolumeRefbQuota
    leafs["prepaid-volume-newi-quota"] = flowFeatureData.PrepaidVolumeNewiQuota
    leafs["prepaid-volume-newo-quota"] = flowFeatureData.PrepaidVolumeNewoQuota
    leafs["prepaid-volume-newb-quota"] = flowFeatureData.PrepaidVolumeNewbQuota
    leafs["prepaid-tariff-volumei-quota"] = flowFeatureData.PrepaidTariffVolumeiQuota
    leafs["prepaid-tariff-volumeo-quota"] = flowFeatureData.PrepaidTariffVolumeoQuota
    leafs["prepaid-tariff-volumeb-quota"] = flowFeatureData.PrepaidTariffVolumebQuota
    leafs["flow-accounting-method-list-name"] = flowFeatureData.FlowAccountingMethodListName
    leafs["prepaid-cfg"] = flowFeatureData.PrepaidCfg
    leafs["prepaid-time-state"] = flowFeatureData.PrepaidTimeState
    leafs["prepaid-volume-state"] = flowFeatureData.PrepaidVolumeState
    leafs["prepaid-charging-rule"] = flowFeatureData.PrepaidChargingRule
    return leafs
}

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetBundleName() string { return "cisco_ios_xr" }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetYangName() string { return "flow-feature-data" }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) SetParent(parent types.Entity) { flowFeatureData.parent = parent }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetParent() types.Entity { return flowFeatureData.parent }

func (flowFeatureData *SubscriberAccounting_Nodes_Node_SubscriberAccountingFlowFeatures_SubscriberAccountingFlowFeature_FlowFeatureData) GetParentYangName() string { return "subscriber-accounting-flow-feature" }

