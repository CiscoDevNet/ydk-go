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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber Sessionmon list of nodes.
    Nodes SessionMon_Nodes
}

func (sessionMon *SessionMon) GetEntityData() *types.CommonEntityData {
    sessionMon.EntityData.YFilter = sessionMon.YFilter
    sessionMon.EntityData.YangName = "session-mon"
    sessionMon.EntityData.BundleName = "cisco_ios_xr"
    sessionMon.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-session-mon-oper"
    sessionMon.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon"
    sessionMon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMon.EntityData.Children = make(map[string]types.YChild)
    sessionMon.EntityData.Children["nodes"] = types.YChild{"Nodes", &sessionMon.Nodes}
    sessionMon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessionMon.EntityData)
}

// SessionMon_Nodes
// Subscriber Sessionmon list of nodes
type SessionMon_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber sessionmon operational data for a particular node. The type is
    // slice of SessionMon_Nodes_Node.
    Node []SessionMon_Nodes_Node
}

func (nodes *SessionMon_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "session-mon"
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

// SessionMon_Nodes_Node
// Subscriber sessionmon operational data for a
// particular node
type SessionMon_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nodeid location . The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Session Mon Statistics.
    SessionMonStatistics SessionMon_Nodes_Node_SessionMonStatistics

    // Statistics Table.
    InterfaceAllStatistics SessionMon_Nodes_Node_InterfaceAllStatistics

    // Smart license.
    LicenseStatistics SessionMon_Nodes_Node_LicenseStatistics
}

func (node *SessionMon_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["session-mon-statistics"] = types.YChild{"SessionMonStatistics", &node.SessionMonStatistics}
    node.EntityData.Children["interface-all-statistics"] = types.YChild{"InterfaceAllStatistics", &node.InterfaceAllStatistics}
    node.EntityData.Children["license-statistics"] = types.YChild{"LicenseStatistics", &node.LicenseStatistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    return &(node.EntityData)
}

// SessionMon_Nodes_Node_SessionMonStatistics
// Session Mon Statistics
type SessionMon_Nodes_Node_SessionMonStatistics struct {
    EntityData types.CommonEntityData
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

func (sessionMonStatistics *SessionMon_Nodes_Node_SessionMonStatistics) GetEntityData() *types.CommonEntityData {
    sessionMonStatistics.EntityData.YFilter = sessionMonStatistics.YFilter
    sessionMonStatistics.EntityData.YangName = "session-mon-statistics"
    sessionMonStatistics.EntityData.BundleName = "cisco_ios_xr"
    sessionMonStatistics.EntityData.ParentYangName = "node"
    sessionMonStatistics.EntityData.SegmentPath = "session-mon-statistics"
    sessionMonStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMonStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMonStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMonStatistics.EntityData.Children = make(map[string]types.YChild)
    sessionMonStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionMonStatistics.EntityData.Leafs["total"] = types.YLeaf{"Total", sessionMonStatistics.Total}
    sessionMonStatistics.EntityData.Leafs["pppoe"] = types.YLeaf{"Pppoe", sessionMonStatistics.Pppoe}
    sessionMonStatistics.EntityData.Leafs["pppoe-ds"] = types.YLeaf{"PppoeDs", sessionMonStatistics.PppoeDs}
    sessionMonStatistics.EntityData.Leafs["dhcpv4"] = types.YLeaf{"Dhcpv4", sessionMonStatistics.Dhcpv4}
    sessionMonStatistics.EntityData.Leafs["dhcpv6"] = types.YLeaf{"Dhcpv6", sessionMonStatistics.Dhcpv6}
    sessionMonStatistics.EntityData.Leafs["dhcp-ds"] = types.YLeaf{"DhcpDs", sessionMonStatistics.DhcpDs}
    sessionMonStatistics.EntityData.Leafs["ippkt"] = types.YLeaf{"Ippkt", sessionMonStatistics.Ippkt}
    sessionMonStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", sessionMonStatistics.ActiveSessions}
    sessionMonStatistics.EntityData.Leafs["standby-sessions"] = types.YLeaf{"StandbySessions", sessionMonStatistics.StandbySessions}
    sessionMonStatistics.EntityData.Leafs["peak-active-sessions"] = types.YLeaf{"PeakActiveSessions", sessionMonStatistics.PeakActiveSessions}
    sessionMonStatistics.EntityData.Leafs["peak-standby-sessions"] = types.YLeaf{"PeakStandbySessions", sessionMonStatistics.PeakStandbySessions}
    sessionMonStatistics.EntityData.Leafs["peak-start-time"] = types.YLeaf{"PeakStartTime", sessionMonStatistics.PeakStartTime}
    sessionMonStatistics.EntityData.Leafs["timeout-value"] = types.YLeaf{"TimeoutValue", sessionMonStatistics.TimeoutValue}
    return &(sessionMonStatistics.EntityData)
}

// SessionMon_Nodes_Node_InterfaceAllStatistics
// Statistics Table
type SessionMon_Nodes_Node_InterfaceAllStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics. The type is slice of
    // SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic.
    InterfaceAllStatistic []SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetEntityData() *types.CommonEntityData {
    interfaceAllStatistics.EntityData.YFilter = interfaceAllStatistics.YFilter
    interfaceAllStatistics.EntityData.YangName = "interface-all-statistics"
    interfaceAllStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceAllStatistics.EntityData.ParentYangName = "node"
    interfaceAllStatistics.EntityData.SegmentPath = "interface-all-statistics"
    interfaceAllStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAllStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAllStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAllStatistics.EntityData.Children = make(map[string]types.YChild)
    interfaceAllStatistics.EntityData.Children["interface-all-statistic"] = types.YChild{"InterfaceAllStatistic", nil}
    for i := range interfaceAllStatistics.InterfaceAllStatistic {
        interfaceAllStatistics.EntityData.Children[types.GetSegmentPath(&interfaceAllStatistics.InterfaceAllStatistic[i])] = types.YChild{"InterfaceAllStatistic", &interfaceAllStatistics.InterfaceAllStatistic[i]}
    }
    interfaceAllStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAllStatistics.EntityData)
}

// SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
// Statistics
type SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (interfaceAllStatistic *SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic) GetEntityData() *types.CommonEntityData {
    interfaceAllStatistic.EntityData.YFilter = interfaceAllStatistic.YFilter
    interfaceAllStatistic.EntityData.YangName = "interface-all-statistic"
    interfaceAllStatistic.EntityData.BundleName = "cisco_ios_xr"
    interfaceAllStatistic.EntityData.ParentYangName = "interface-all-statistics"
    interfaceAllStatistic.EntityData.SegmentPath = "interface-all-statistic" + "[interface-name='" + fmt.Sprintf("%v", interfaceAllStatistic.InterfaceName) + "']"
    interfaceAllStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAllStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAllStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAllStatistic.EntityData.Children = make(map[string]types.YChild)
    interfaceAllStatistic.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAllStatistic.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceAllStatistic.InterfaceName}
    interfaceAllStatistic.EntityData.Leafs["total"] = types.YLeaf{"Total", interfaceAllStatistic.Total}
    interfaceAllStatistic.EntityData.Leafs["pppoe"] = types.YLeaf{"Pppoe", interfaceAllStatistic.Pppoe}
    interfaceAllStatistic.EntityData.Leafs["pppoe-ds"] = types.YLeaf{"PppoeDs", interfaceAllStatistic.PppoeDs}
    interfaceAllStatistic.EntityData.Leafs["dhcpv4"] = types.YLeaf{"Dhcpv4", interfaceAllStatistic.Dhcpv4}
    interfaceAllStatistic.EntityData.Leafs["dhcpv6"] = types.YLeaf{"Dhcpv6", interfaceAllStatistic.Dhcpv6}
    interfaceAllStatistic.EntityData.Leafs["dhcp-ds"] = types.YLeaf{"DhcpDs", interfaceAllStatistic.DhcpDs}
    interfaceAllStatistic.EntityData.Leafs["ippkt"] = types.YLeaf{"Ippkt", interfaceAllStatistic.Ippkt}
    interfaceAllStatistic.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", interfaceAllStatistic.ActiveSessions}
    interfaceAllStatistic.EntityData.Leafs["standby-sessions"] = types.YLeaf{"StandbySessions", interfaceAllStatistic.StandbySessions}
    interfaceAllStatistic.EntityData.Leafs["peak-active-sessions"] = types.YLeaf{"PeakActiveSessions", interfaceAllStatistic.PeakActiveSessions}
    interfaceAllStatistic.EntityData.Leafs["peak-standby-sessions"] = types.YLeaf{"PeakStandbySessions", interfaceAllStatistic.PeakStandbySessions}
    interfaceAllStatistic.EntityData.Leafs["peak-start-time"] = types.YLeaf{"PeakStartTime", interfaceAllStatistic.PeakStartTime}
    interfaceAllStatistic.EntityData.Leafs["timeout-value"] = types.YLeaf{"TimeoutValue", interfaceAllStatistic.TimeoutValue}
    return &(interfaceAllStatistic.EntityData)
}

// SessionMon_Nodes_Node_LicenseStatistics
// Smart license
type SessionMon_Nodes_Node_LicenseStatistics struct {
    EntityData types.CommonEntityData
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

func (licenseStatistics *SessionMon_Nodes_Node_LicenseStatistics) GetEntityData() *types.CommonEntityData {
    licenseStatistics.EntityData.YFilter = licenseStatistics.YFilter
    licenseStatistics.EntityData.YangName = "license-statistics"
    licenseStatistics.EntityData.BundleName = "cisco_ios_xr"
    licenseStatistics.EntityData.ParentYangName = "node"
    licenseStatistics.EntityData.SegmentPath = "license-statistics"
    licenseStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    licenseStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    licenseStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    licenseStatistics.EntityData.Children = make(map[string]types.YChild)
    licenseStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    licenseStatistics.EntityData.Leafs["total"] = types.YLeaf{"Total", licenseStatistics.Total}
    licenseStatistics.EntityData.Leafs["pppoe"] = types.YLeaf{"Pppoe", licenseStatistics.Pppoe}
    licenseStatistics.EntityData.Leafs["pppoe-ds"] = types.YLeaf{"PppoeDs", licenseStatistics.PppoeDs}
    licenseStatistics.EntityData.Leafs["dhcpv4"] = types.YLeaf{"Dhcpv4", licenseStatistics.Dhcpv4}
    licenseStatistics.EntityData.Leafs["dhcpv6"] = types.YLeaf{"Dhcpv6", licenseStatistics.Dhcpv6}
    licenseStatistics.EntityData.Leafs["dhcp-ds"] = types.YLeaf{"DhcpDs", licenseStatistics.DhcpDs}
    licenseStatistics.EntityData.Leafs["ippkt"] = types.YLeaf{"Ippkt", licenseStatistics.Ippkt}
    licenseStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", licenseStatistics.ActiveSessions}
    licenseStatistics.EntityData.Leafs["standby-sessions"] = types.YLeaf{"StandbySessions", licenseStatistics.StandbySessions}
    licenseStatistics.EntityData.Leafs["peak-active-sessions"] = types.YLeaf{"PeakActiveSessions", licenseStatistics.PeakActiveSessions}
    licenseStatistics.EntityData.Leafs["peak-standby-sessions"] = types.YLeaf{"PeakStandbySessions", licenseStatistics.PeakStandbySessions}
    licenseStatistics.EntityData.Leafs["peak-start-time"] = types.YLeaf{"PeakStartTime", licenseStatistics.PeakStartTime}
    licenseStatistics.EntityData.Leafs["timeout-value"] = types.YLeaf{"TimeoutValue", licenseStatistics.TimeoutValue}
    return &(licenseStatistics.EntityData)
}

