// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-session-mon package operational data.
// 
// This module contains definitions
// for the following management objects:
//   session-mon: Sessionmon
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_session_mon_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_session_mon_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-session-mon-oper session-mon}", reflect.TypeOf(SessionMon{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-session-mon-oper:session-mon", reflect.TypeOf(SessionMon{}))
}

// SessionMon
// Sessionmon
type SessionMon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber Sessionmon list of nodes.
    Nodes SessionMon_Nodes
}

func (sessionMon *SessionMon) GetFilter() yfilter.YFilter { return sessionMon.YFilter }

func (sessionMon *SessionMon) SetFilter(yf yfilter.YFilter) { sessionMon.YFilter = yf }

func (sessionMon *SessionMon) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (sessionMon *SessionMon) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon"
}

func (sessionMon *SessionMon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &sessionMon.Nodes
    }
    return nil
}

func (sessionMon *SessionMon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &sessionMon.Nodes
    return children
}

func (sessionMon *SessionMon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionMon *SessionMon) GetBundleName() string { return "cisco_ios_xr" }

func (sessionMon *SessionMon) GetYangName() string { return "session-mon" }

func (sessionMon *SessionMon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionMon *SessionMon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionMon *SessionMon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionMon *SessionMon) SetParent(parent types.Entity) { sessionMon.parent = parent }

func (sessionMon *SessionMon) GetParent() types.Entity { return sessionMon.parent }

func (sessionMon *SessionMon) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-session-mon-oper" }

// SessionMon_Nodes
// Subscriber Sessionmon list of nodes
type SessionMon_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber sessionmon operational data for a particular node. The type is
    // slice of SessionMon_Nodes_Node.
    Node []SessionMon_Nodes_Node
}

func (nodes *SessionMon_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SessionMon_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SessionMon_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SessionMon_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SessionMon_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionMon_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SessionMon_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SessionMon_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SessionMon_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SessionMon_Nodes) GetYangName() string { return "nodes" }

func (nodes *SessionMon_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SessionMon_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SessionMon_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SessionMon_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SessionMon_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SessionMon_Nodes) GetParentYangName() string { return "session-mon" }

// SessionMon_Nodes_Node
// Subscriber sessionmon operational data for a
// particular node
type SessionMon_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nodeid location . The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Session Mon Statistics.
    SessionMonStatistics SessionMon_Nodes_Node_SessionMonStatistics

    // Statistics Table.
    InterfaceAllStatistics SessionMon_Nodes_Node_InterfaceAllStatistics

    // Smart license.
    LicenseStatistics SessionMon_Nodes_Node_LicenseStatistics
}

func (node *SessionMon_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SessionMon_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SessionMon_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "session-mon-statistics" { return "SessionMonStatistics" }
    if yname == "interface-all-statistics" { return "InterfaceAllStatistics" }
    if yname == "license-statistics" { return "LicenseStatistics" }
    return ""
}

func (node *SessionMon_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *SessionMon_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-mon-statistics" {
        return &node.SessionMonStatistics
    }
    if childYangName == "interface-all-statistics" {
        return &node.InterfaceAllStatistics
    }
    if childYangName == "license-statistics" {
        return &node.LicenseStatistics
    }
    return nil
}

func (node *SessionMon_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-mon-statistics"] = &node.SessionMonStatistics
    children["interface-all-statistics"] = &node.InterfaceAllStatistics
    children["license-statistics"] = &node.LicenseStatistics
    return children
}

func (node *SessionMon_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    return leafs
}

func (node *SessionMon_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SessionMon_Nodes_Node) GetYangName() string { return "node" }

func (node *SessionMon_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SessionMon_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SessionMon_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SessionMon_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SessionMon_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SessionMon_Nodes_Node) GetParentYangName() string { return "nodes" }

// SessionMon_Nodes_Node_SessionMonStatistics
// Session Mon Statistics
type SessionMon_Nodes_Node_SessionMonStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // total. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // pppoe. The type is interface{} with range: 0..4294967295.
    Pppoe interface{}

    // pppoe ds. The type is interface{} with range: 0..4294967295.
    PppoeDs interface{}

    // dhcpv4. The type is interface{} with range: 0..4294967295.
    Dhcpv4 interface{}

    // dhcpv6. The type is interface{} with range: 0..4294967295.
    Dhcpv6 interface{}

    // dhcp ds. The type is interface{} with range: 0..4294967295.
    DhcpDs interface{}

    // ippkt. The type is interface{} with range: 0..4294967295.
    Ippkt interface{}

    // active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // standby sessions. The type is interface{} with range: 0..4294967295.
    StandbySessions interface{}

    // peak active sessions. The type is interface{} with range: 0..4294967295.
    PeakActiveSessions interface{}

    // peak standby sessions. The type is interface{} with range: 0..4294967295.
    PeakStandbySessions interface{}

    // peak start time. The type is interface{} with range: 0..4294967295.
    PeakStartTime interface{}

    // timeout value. The type is interface{} with range: 0..4294967295.
    TimeoutValue interface{}
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetFilter() yfilter.YFilter { return sessionMonStatistics.YFilter }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) SetFilter(yf yfilter.YFilter) { sessionMonStatistics.YFilter = yf }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "pppoe" { return "Pppoe" }
    if yname == "pppoe-ds" { return "PppoeDs" }
    if yname == "dhcpv4" { return "Dhcpv4" }
    if yname == "dhcpv6" { return "Dhcpv6" }
    if yname == "dhcp-ds" { return "DhcpDs" }
    if yname == "ippkt" { return "Ippkt" }
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "standby-sessions" { return "StandbySessions" }
    if yname == "peak-active-sessions" { return "PeakActiveSessions" }
    if yname == "peak-standby-sessions" { return "PeakStandbySessions" }
    if yname == "peak-start-time" { return "PeakStartTime" }
    if yname == "timeout-value" { return "TimeoutValue" }
    return ""
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetSegmentPath() string {
    return "session-mon-statistics"
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = sessionMonStatistics.Total
    leafs["pppoe"] = sessionMonStatistics.Pppoe
    leafs["pppoe-ds"] = sessionMonStatistics.PppoeDs
    leafs["dhcpv4"] = sessionMonStatistics.Dhcpv4
    leafs["dhcpv6"] = sessionMonStatistics.Dhcpv6
    leafs["dhcp-ds"] = sessionMonStatistics.DhcpDs
    leafs["ippkt"] = sessionMonStatistics.Ippkt
    leafs["active-sessions"] = sessionMonStatistics.ActiveSessions
    leafs["standby-sessions"] = sessionMonStatistics.StandbySessions
    leafs["peak-active-sessions"] = sessionMonStatistics.PeakActiveSessions
    leafs["peak-standby-sessions"] = sessionMonStatistics.PeakStandbySessions
    leafs["peak-start-time"] = sessionMonStatistics.PeakStartTime
    leafs["timeout-value"] = sessionMonStatistics.TimeoutValue
    return leafs
}

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetYangName() string { return "session-mon-statistics" }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) SetParent(parent types.Entity) { sessionMonStatistics.parent = parent }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetParent() types.Entity { return sessionMonStatistics.parent }

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetParentYangName() string { return "node" }

// SessionMon_Nodes_Node_InterfaceAllStatistics
// Statistics Table
type SessionMon_Nodes_Node_InterfaceAllStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics. The type is slice of
    // SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic.
    InterfaceAllStatistic []SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetFilter() yfilter.YFilter { return interfaceAllStatistics.YFilter }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) SetFilter(yf yfilter.YFilter) { interfaceAllStatistics.YFilter = yf }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetGoName(yname string) string {
    if yname == "interface-all-statistic" { return "InterfaceAllStatistic" }
    return ""
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetSegmentPath() string {
    return "interface-all-statistics"
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-all-statistic" {
        for _, c := range interfaceAllStatistics.InterfaceAllStatistic {
            if interfaceAllStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic{}
        interfaceAllStatistics.InterfaceAllStatistic = append(interfaceAllStatistics.InterfaceAllStatistic, child)
        return &interfaceAllStatistics.InterfaceAllStatistic[len(interfaceAllStatistics.InterfaceAllStatistic)-1]
    }
    return nil
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceAllStatistics.InterfaceAllStatistic {
        children[interfaceAllStatistics.InterfaceAllStatistic[i].GetSegmentPath()] = &interfaceAllStatistics.InterfaceAllStatistic[i]
    }
    return children
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetYangName() string { return "interface-all-statistics" }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) SetParent(parent types.Entity) { interfaceAllStatistics.parent = parent }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetParent() types.Entity { return interfaceAllStatistics.parent }

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetParentYangName() string { return "node" }

// SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
// Statistics
type SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // total. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // pppoe. The type is interface{} with range: 0..4294967295.
    Pppoe interface{}

    // pppoe ds. The type is interface{} with range: 0..4294967295.
    PppoeDs interface{}

    // dhcpv4. The type is interface{} with range: 0..4294967295.
    Dhcpv4 interface{}

    // dhcpv6. The type is interface{} with range: 0..4294967295.
    Dhcpv6 interface{}

    // dhcp ds. The type is interface{} with range: 0..4294967295.
    DhcpDs interface{}

    // ippkt. The type is interface{} with range: 0..4294967295.
    Ippkt interface{}

    // active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // standby sessions. The type is interface{} with range: 0..4294967295.
    StandbySessions interface{}

    // peak active sessions. The type is interface{} with range: 0..4294967295.
    PeakActiveSessions interface{}

    // peak standby sessions. The type is interface{} with range: 0..4294967295.
    PeakStandbySessions interface{}

    // peak start time. The type is interface{} with range: 0..4294967295.
    PeakStartTime interface{}

    // timeout value. The type is interface{} with range: 0..4294967295.
    TimeoutValue interface{}
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetFilter() yfilter.YFilter { return interfaceAllStatistic.YFilter }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) SetFilter(yf yfilter.YFilter) { interfaceAllStatistic.YFilter = yf }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "total" { return "Total" }
    if yname == "pppoe" { return "Pppoe" }
    if yname == "pppoe-ds" { return "PppoeDs" }
    if yname == "dhcpv4" { return "Dhcpv4" }
    if yname == "dhcpv6" { return "Dhcpv6" }
    if yname == "dhcp-ds" { return "DhcpDs" }
    if yname == "ippkt" { return "Ippkt" }
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "standby-sessions" { return "StandbySessions" }
    if yname == "peak-active-sessions" { return "PeakActiveSessions" }
    if yname == "peak-standby-sessions" { return "PeakStandbySessions" }
    if yname == "peak-start-time" { return "PeakStartTime" }
    if yname == "timeout-value" { return "TimeoutValue" }
    return ""
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetSegmentPath() string {
    return "interface-all-statistic" + "[interface-name='" + fmt.Sprintf("%v", interfaceAllStatistic.InterfaceName) + "']"
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceAllStatistic.InterfaceName
    leafs["total"] = interfaceAllStatistic.Total
    leafs["pppoe"] = interfaceAllStatistic.Pppoe
    leafs["pppoe-ds"] = interfaceAllStatistic.PppoeDs
    leafs["dhcpv4"] = interfaceAllStatistic.Dhcpv4
    leafs["dhcpv6"] = interfaceAllStatistic.Dhcpv6
    leafs["dhcp-ds"] = interfaceAllStatistic.DhcpDs
    leafs["ippkt"] = interfaceAllStatistic.Ippkt
    leafs["active-sessions"] = interfaceAllStatistic.ActiveSessions
    leafs["standby-sessions"] = interfaceAllStatistic.StandbySessions
    leafs["peak-active-sessions"] = interfaceAllStatistic.PeakActiveSessions
    leafs["peak-standby-sessions"] = interfaceAllStatistic.PeakStandbySessions
    leafs["peak-start-time"] = interfaceAllStatistic.PeakStartTime
    leafs["timeout-value"] = interfaceAllStatistic.TimeoutValue
    return leafs
}

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetYangName() string { return "interface-all-statistic" }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) SetParent(parent types.Entity) { interfaceAllStatistic.parent = parent }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetParent() types.Entity { return interfaceAllStatistic.parent }

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetParentYangName() string { return "interface-all-statistics" }

// SessionMon_Nodes_Node_LicenseStatistics
// Smart license
type SessionMon_Nodes_Node_LicenseStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // total. The type is interface{} with range: 0..4294967295.
    Total interface{}

    // pppoe. The type is interface{} with range: 0..4294967295.
    Pppoe interface{}

    // pppoe ds. The type is interface{} with range: 0..4294967295.
    PppoeDs interface{}

    // dhcpv4. The type is interface{} with range: 0..4294967295.
    Dhcpv4 interface{}

    // dhcpv6. The type is interface{} with range: 0..4294967295.
    Dhcpv6 interface{}

    // dhcp ds. The type is interface{} with range: 0..4294967295.
    DhcpDs interface{}

    // ippkt. The type is interface{} with range: 0..4294967295.
    Ippkt interface{}

    // active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // standby sessions. The type is interface{} with range: 0..4294967295.
    StandbySessions interface{}

    // peak active sessions. The type is interface{} with range: 0..4294967295.
    PeakActiveSessions interface{}

    // peak standby sessions. The type is interface{} with range: 0..4294967295.
    PeakStandbySessions interface{}

    // peak start time. The type is interface{} with range: 0..4294967295.
    PeakStartTime interface{}

    // timeout value. The type is interface{} with range: 0..4294967295.
    TimeoutValue interface{}
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetFilter() yfilter.YFilter { return licenseStatistics.YFilter }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) SetFilter(yf yfilter.YFilter) { licenseStatistics.YFilter = yf }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "pppoe" { return "Pppoe" }
    if yname == "pppoe-ds" { return "PppoeDs" }
    if yname == "dhcpv4" { return "Dhcpv4" }
    if yname == "dhcpv6" { return "Dhcpv6" }
    if yname == "dhcp-ds" { return "DhcpDs" }
    if yname == "ippkt" { return "Ippkt" }
    if yname == "active-sessions" { return "ActiveSessions" }
    if yname == "standby-sessions" { return "StandbySessions" }
    if yname == "peak-active-sessions" { return "PeakActiveSessions" }
    if yname == "peak-standby-sessions" { return "PeakStandbySessions" }
    if yname == "peak-start-time" { return "PeakStartTime" }
    if yname == "timeout-value" { return "TimeoutValue" }
    return ""
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetSegmentPath() string {
    return "license-statistics"
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = licenseStatistics.Total
    leafs["pppoe"] = licenseStatistics.Pppoe
    leafs["pppoe-ds"] = licenseStatistics.PppoeDs
    leafs["dhcpv4"] = licenseStatistics.Dhcpv4
    leafs["dhcpv6"] = licenseStatistics.Dhcpv6
    leafs["dhcp-ds"] = licenseStatistics.DhcpDs
    leafs["ippkt"] = licenseStatistics.Ippkt
    leafs["active-sessions"] = licenseStatistics.ActiveSessions
    leafs["standby-sessions"] = licenseStatistics.StandbySessions
    leafs["peak-active-sessions"] = licenseStatistics.PeakActiveSessions
    leafs["peak-standby-sessions"] = licenseStatistics.PeakStandbySessions
    leafs["peak-start-time"] = licenseStatistics.PeakStartTime
    leafs["timeout-value"] = licenseStatistics.TimeoutValue
    return leafs
}

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetYangName() string { return "license-statistics" }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) SetParent(parent types.Entity) { licenseStatistics.parent = parent }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetParent() types.Entity { return licenseStatistics.parent }

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetParentYangName() string { return "node" }

