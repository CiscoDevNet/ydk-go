// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-session-mon package operational data.
// 
// This module contains definitions
// for the following management objects:
//   session-mon: Sessionmon
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    sessionMon.EntityData.AbsolutePath = sessionMon.EntityData.SegmentPath
    sessionMon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMon.EntityData.Children = types.NewOrderedMap()
    sessionMon.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sessionMon.Nodes})
    sessionMon.EntityData.Leafs = types.NewOrderedMap()

    sessionMon.EntityData.YListKeys = []string {}

    return &(sessionMon.EntityData)
}

// SessionMon_Nodes
// Subscriber Sessionmon list of nodes
type SessionMon_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber sessionmon operational data for a particular node. The type is
    // slice of SessionMon_Nodes_Node.
    Node []*SessionMon_Nodes_Node
}

func (nodes *SessionMon_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "session-mon"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/" + nodes.EntityData.SegmentPath
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

// SessionMon_Nodes_Node
// Subscriber sessionmon operational data for a
// particular node
type SessionMon_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("session-mon-statistics", types.YChild{"SessionMonStatistics", &node.SessionMonStatistics})
    node.EntityData.Children.Append("interface-all-statistics", types.YChild{"InterfaceAllStatistics", &node.InterfaceAllStatistics})
    node.EntityData.Children.Append("license-statistics", types.YChild{"LicenseStatistics", &node.LicenseStatistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})

    node.EntityData.YListKeys = []string {"NodeId"}

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
    sessionMonStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/nodes/node/" + sessionMonStatistics.EntityData.SegmentPath
    sessionMonStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMonStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMonStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMonStatistics.EntityData.Children = types.NewOrderedMap()
    sessionMonStatistics.EntityData.Leafs = types.NewOrderedMap()
    sessionMonStatistics.EntityData.Leafs.Append("total", types.YLeaf{"Total", sessionMonStatistics.Total})
    sessionMonStatistics.EntityData.Leafs.Append("pppoe", types.YLeaf{"Pppoe", sessionMonStatistics.Pppoe})
    sessionMonStatistics.EntityData.Leafs.Append("pppoe-ds", types.YLeaf{"PppoeDs", sessionMonStatistics.PppoeDs})
    sessionMonStatistics.EntityData.Leafs.Append("dhcpv4", types.YLeaf{"Dhcpv4", sessionMonStatistics.Dhcpv4})
    sessionMonStatistics.EntityData.Leafs.Append("dhcpv6", types.YLeaf{"Dhcpv6", sessionMonStatistics.Dhcpv6})
    sessionMonStatistics.EntityData.Leafs.Append("dhcp-ds", types.YLeaf{"DhcpDs", sessionMonStatistics.DhcpDs})
    sessionMonStatistics.EntityData.Leafs.Append("ippkt", types.YLeaf{"Ippkt", sessionMonStatistics.Ippkt})
    sessionMonStatistics.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", sessionMonStatistics.ActiveSessions})
    sessionMonStatistics.EntityData.Leafs.Append("standby-sessions", types.YLeaf{"StandbySessions", sessionMonStatistics.StandbySessions})
    sessionMonStatistics.EntityData.Leafs.Append("peak-active-sessions", types.YLeaf{"PeakActiveSessions", sessionMonStatistics.PeakActiveSessions})
    sessionMonStatistics.EntityData.Leafs.Append("peak-standby-sessions", types.YLeaf{"PeakStandbySessions", sessionMonStatistics.PeakStandbySessions})
    sessionMonStatistics.EntityData.Leafs.Append("peak-start-time", types.YLeaf{"PeakStartTime", sessionMonStatistics.PeakStartTime})
    sessionMonStatistics.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", sessionMonStatistics.TimeoutValue})

    sessionMonStatistics.EntityData.YListKeys = []string {}

    return &(sessionMonStatistics.EntityData)
}

// SessionMon_Nodes_Node_InterfaceAllStatistics
// Statistics Table
type SessionMon_Nodes_Node_InterfaceAllStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics. The type is slice of
    // SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic.
    InterfaceAllStatistic []*SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
}

func (interfaceAllStatistics *SessionMon_Nodes_Node_InterfaceAllStatistics) GetEntityData() *types.CommonEntityData {
    interfaceAllStatistics.EntityData.YFilter = interfaceAllStatistics.YFilter
    interfaceAllStatistics.EntityData.YangName = "interface-all-statistics"
    interfaceAllStatistics.EntityData.BundleName = "cisco_ios_xr"
    interfaceAllStatistics.EntityData.ParentYangName = "node"
    interfaceAllStatistics.EntityData.SegmentPath = "interface-all-statistics"
    interfaceAllStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/nodes/node/" + interfaceAllStatistics.EntityData.SegmentPath
    interfaceAllStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAllStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAllStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAllStatistics.EntityData.Children = types.NewOrderedMap()
    interfaceAllStatistics.EntityData.Children.Append("interface-all-statistic", types.YChild{"InterfaceAllStatistic", nil})
    for i := range interfaceAllStatistics.InterfaceAllStatistic {
        interfaceAllStatistics.EntityData.Children.Append(types.GetSegmentPath(interfaceAllStatistics.InterfaceAllStatistic[i]), types.YChild{"InterfaceAllStatistic", interfaceAllStatistics.InterfaceAllStatistic[i]})
    }
    interfaceAllStatistics.EntityData.Leafs = types.NewOrderedMap()

    interfaceAllStatistics.EntityData.YListKeys = []string {}

    return &(interfaceAllStatistics.EntityData)
}

// SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic
// Statistics
type SessionMon_Nodes_Node_InterfaceAllStatistics_InterfaceAllStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
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
    interfaceAllStatistic.EntityData.SegmentPath = "interface-all-statistic" + types.AddKeyToken(interfaceAllStatistic.InterfaceName, "interface-name")
    interfaceAllStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/nodes/node/interface-all-statistics/" + interfaceAllStatistic.EntityData.SegmentPath
    interfaceAllStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAllStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAllStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAllStatistic.EntityData.Children = types.NewOrderedMap()
    interfaceAllStatistic.EntityData.Leafs = types.NewOrderedMap()
    interfaceAllStatistic.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceAllStatistic.InterfaceName})
    interfaceAllStatistic.EntityData.Leafs.Append("total", types.YLeaf{"Total", interfaceAllStatistic.Total})
    interfaceAllStatistic.EntityData.Leafs.Append("pppoe", types.YLeaf{"Pppoe", interfaceAllStatistic.Pppoe})
    interfaceAllStatistic.EntityData.Leafs.Append("pppoe-ds", types.YLeaf{"PppoeDs", interfaceAllStatistic.PppoeDs})
    interfaceAllStatistic.EntityData.Leafs.Append("dhcpv4", types.YLeaf{"Dhcpv4", interfaceAllStatistic.Dhcpv4})
    interfaceAllStatistic.EntityData.Leafs.Append("dhcpv6", types.YLeaf{"Dhcpv6", interfaceAllStatistic.Dhcpv6})
    interfaceAllStatistic.EntityData.Leafs.Append("dhcp-ds", types.YLeaf{"DhcpDs", interfaceAllStatistic.DhcpDs})
    interfaceAllStatistic.EntityData.Leafs.Append("ippkt", types.YLeaf{"Ippkt", interfaceAllStatistic.Ippkt})
    interfaceAllStatistic.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", interfaceAllStatistic.ActiveSessions})
    interfaceAllStatistic.EntityData.Leafs.Append("standby-sessions", types.YLeaf{"StandbySessions", interfaceAllStatistic.StandbySessions})
    interfaceAllStatistic.EntityData.Leafs.Append("peak-active-sessions", types.YLeaf{"PeakActiveSessions", interfaceAllStatistic.PeakActiveSessions})
    interfaceAllStatistic.EntityData.Leafs.Append("peak-standby-sessions", types.YLeaf{"PeakStandbySessions", interfaceAllStatistic.PeakStandbySessions})
    interfaceAllStatistic.EntityData.Leafs.Append("peak-start-time", types.YLeaf{"PeakStartTime", interfaceAllStatistic.PeakStartTime})
    interfaceAllStatistic.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", interfaceAllStatistic.TimeoutValue})

    interfaceAllStatistic.EntityData.YListKeys = []string {"InterfaceName"}

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
    licenseStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-session-mon-oper:session-mon/nodes/node/" + licenseStatistics.EntityData.SegmentPath
    licenseStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    licenseStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    licenseStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    licenseStatistics.EntityData.Children = types.NewOrderedMap()
    licenseStatistics.EntityData.Leafs = types.NewOrderedMap()
    licenseStatistics.EntityData.Leafs.Append("total", types.YLeaf{"Total", licenseStatistics.Total})
    licenseStatistics.EntityData.Leafs.Append("pppoe", types.YLeaf{"Pppoe", licenseStatistics.Pppoe})
    licenseStatistics.EntityData.Leafs.Append("pppoe-ds", types.YLeaf{"PppoeDs", licenseStatistics.PppoeDs})
    licenseStatistics.EntityData.Leafs.Append("dhcpv4", types.YLeaf{"Dhcpv4", licenseStatistics.Dhcpv4})
    licenseStatistics.EntityData.Leafs.Append("dhcpv6", types.YLeaf{"Dhcpv6", licenseStatistics.Dhcpv6})
    licenseStatistics.EntityData.Leafs.Append("dhcp-ds", types.YLeaf{"DhcpDs", licenseStatistics.DhcpDs})
    licenseStatistics.EntityData.Leafs.Append("ippkt", types.YLeaf{"Ippkt", licenseStatistics.Ippkt})
    licenseStatistics.EntityData.Leafs.Append("active-sessions", types.YLeaf{"ActiveSessions", licenseStatistics.ActiveSessions})
    licenseStatistics.EntityData.Leafs.Append("standby-sessions", types.YLeaf{"StandbySessions", licenseStatistics.StandbySessions})
    licenseStatistics.EntityData.Leafs.Append("peak-active-sessions", types.YLeaf{"PeakActiveSessions", licenseStatistics.PeakActiveSessions})
    licenseStatistics.EntityData.Leafs.Append("peak-standby-sessions", types.YLeaf{"PeakStandbySessions", licenseStatistics.PeakStandbySessions})
    licenseStatistics.EntityData.Leafs.Append("peak-start-time", types.YLeaf{"PeakStartTime", licenseStatistics.PeakStartTime})
    licenseStatistics.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", licenseStatistics.TimeoutValue})

    licenseStatistics.EntityData.YListKeys = []string {}

    return &(licenseStatistics.EntityData)
}

