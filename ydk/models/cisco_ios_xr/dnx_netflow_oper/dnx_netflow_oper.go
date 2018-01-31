// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-netflow package operational data.
// 
// This module contains definitions
// for the following management objects:
//   net-flow: NetFlow operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dnx_netflow_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dnx_netflow_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dnx-netflow-oper net-flow}", reflect.TypeOf(NetFlow{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dnx-netflow-oper:net-flow", reflect.TypeOf(NetFlow{}))
}

// NetFlow
// NetFlow operational data
type NetFlow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific NetFlow statistics information.
    Statistics NetFlow_Statistics
}

func (netFlow *NetFlow) GetFilter() yfilter.YFilter { return netFlow.YFilter }

func (netFlow *NetFlow) SetFilter(yf yfilter.YFilter) { netFlow.YFilter = yf }

func (netFlow *NetFlow) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (netFlow *NetFlow) GetSegmentPath() string {
    return "Cisco-IOS-XR-dnx-netflow-oper:net-flow"
}

func (netFlow *NetFlow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &netFlow.Statistics
    }
    return nil
}

func (netFlow *NetFlow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &netFlow.Statistics
    return children
}

func (netFlow *NetFlow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netFlow *NetFlow) GetBundleName() string { return "cisco_ios_xr" }

func (netFlow *NetFlow) GetYangName() string { return "net-flow" }

func (netFlow *NetFlow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netFlow *NetFlow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netFlow *NetFlow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netFlow *NetFlow) SetParent(parent types.Entity) { netFlow.parent = parent }

func (netFlow *NetFlow) GetParent() types.Entity { return netFlow.parent }

func (netFlow *NetFlow) GetParentYangName() string { return "Cisco-IOS-XR-dnx-netflow-oper" }

// NetFlow_Statistics
// Node-specific NetFlow statistics information
type NetFlow_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NetFlow statistics information for a particular node. The type is slice of
    // NetFlow_Statistics_Statistic.
    Statistic []NetFlow_Statistics_Statistic
}

func (statistics *NetFlow_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *NetFlow_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *NetFlow_Statistics) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (statistics *NetFlow_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *NetFlow_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range statistics.Statistic {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_Statistics_Statistic{}
        statistics.Statistic = append(statistics.Statistic, child)
        return &statistics.Statistic[len(statistics.Statistic)-1]
    }
    return nil
}

func (statistics *NetFlow_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Statistic {
        children[statistics.Statistic[i].GetSegmentPath()] = &statistics.Statistic[i]
    }
    return children
}

func (statistics *NetFlow_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *NetFlow_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *NetFlow_Statistics) GetYangName() string { return "statistics" }

func (statistics *NetFlow_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *NetFlow_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *NetFlow_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *NetFlow_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *NetFlow_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *NetFlow_Statistics) GetParentYangName() string { return "net-flow" }

// NetFlow_Statistics_Statistic
// NetFlow statistics information for a particular
// node
type NetFlow_Statistics_Statistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // NetFlow producer statistics.
    Producer NetFlow_Statistics_Statistic_Producer

    // NetFlow server statistics.
    Server NetFlow_Statistics_Statistic_Server
}

func (statistic *NetFlow_Statistics_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *NetFlow_Statistics_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *NetFlow_Statistics_Statistic) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "producer" { return "Producer" }
    if yname == "server" { return "Server" }
    return ""
}

func (statistic *NetFlow_Statistics_Statistic) GetSegmentPath() string {
    return "statistic" + "[node='" + fmt.Sprintf("%v", statistic.Node) + "']"
}

func (statistic *NetFlow_Statistics_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "producer" {
        return &statistic.Producer
    }
    if childYangName == "server" {
        return &statistic.Server
    }
    return nil
}

func (statistic *NetFlow_Statistics_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["producer"] = &statistic.Producer
    children["server"] = &statistic.Server
    return children
}

func (statistic *NetFlow_Statistics_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = statistic.Node
    return leafs
}

func (statistic *NetFlow_Statistics_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *NetFlow_Statistics_Statistic) GetYangName() string { return "statistic" }

func (statistic *NetFlow_Statistics_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *NetFlow_Statistics_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *NetFlow_Statistics_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *NetFlow_Statistics_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *NetFlow_Statistics_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *NetFlow_Statistics_Statistic) GetParentYangName() string { return "statistics" }

// NetFlow_Statistics_Statistic_Producer
// NetFlow producer statistics
type NetFlow_Statistics_Statistic_Producer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics information.
    Statistics NetFlow_Statistics_Statistic_Producer_Statistics
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetFilter() yfilter.YFilter { return producer.YFilter }

func (producer *NetFlow_Statistics_Statistic_Producer) SetFilter(yf yfilter.YFilter) { producer.YFilter = yf }

func (producer *NetFlow_Statistics_Statistic_Producer) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetSegmentPath() string {
    return "producer"
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &producer.Statistics
    }
    return nil
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &producer.Statistics
    return children
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetBundleName() string { return "cisco_ios_xr" }

func (producer *NetFlow_Statistics_Statistic_Producer) GetYangName() string { return "producer" }

func (producer *NetFlow_Statistics_Statistic_Producer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (producer *NetFlow_Statistics_Statistic_Producer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (producer *NetFlow_Statistics_Statistic_Producer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (producer *NetFlow_Statistics_Statistic_Producer) SetParent(parent types.Entity) { producer.parent = parent }

func (producer *NetFlow_Statistics_Statistic_Producer) GetParent() types.Entity { return producer.parent }

func (producer *NetFlow_Statistics_Statistic_Producer) GetParentYangName() string { return "statistic" }

// NetFlow_Statistics_Statistic_Producer_Statistics
// Statistics information
type NetFlow_Statistics_Statistic_Producer_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipv4IngressFlows interface{}

    // IPv4 egress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipv4EgressFlows interface{}

    // IPv6 ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipv6IngressFlows interface{}

    // IPv6 egress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipv6EgressFlows interface{}

    // MPLS ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    MplsIngressFlows interface{}

    // MPLS egress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    MplsEgressFlows interface{}

    // Drops (no space). The type is interface{} with range:
    // 0..18446744073709551615.
    DropsNoSpace interface{}

    // Drops (others). The type is interface{} with range:
    // 0..18446744073709551615.
    DropsOthers interface{}

    // Unknown ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownIngressFlows interface{}

    // Unknown egress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownEgressFlows interface{}

    // Number of waiting servers. The type is interface{} with range:
    // 0..18446744073709551615.
    WaitingServers interface{}

    // Last time Statistics cleared in 'Mon Jan 1 12:00 :00 2xxx' format. The type
    // is string.
    LastCleared interface{}
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-ingress-flows" { return "Ipv4IngressFlows" }
    if yname == "ipv4-egress-flows" { return "Ipv4EgressFlows" }
    if yname == "ipv6-ingress-flows" { return "Ipv6IngressFlows" }
    if yname == "ipv6-egress-flows" { return "Ipv6EgressFlows" }
    if yname == "mpls-ingress-flows" { return "MplsIngressFlows" }
    if yname == "mpls-egress-flows" { return "MplsEgressFlows" }
    if yname == "drops-no-space" { return "DropsNoSpace" }
    if yname == "drops-others" { return "DropsOthers" }
    if yname == "unknown-ingress-flows" { return "UnknownIngressFlows" }
    if yname == "unknown-egress-flows" { return "UnknownEgressFlows" }
    if yname == "waiting-servers" { return "WaitingServers" }
    if yname == "last-cleared" { return "LastCleared" }
    return ""
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-ingress-flows"] = statistics.Ipv4IngressFlows
    leafs["ipv4-egress-flows"] = statistics.Ipv4EgressFlows
    leafs["ipv6-ingress-flows"] = statistics.Ipv6IngressFlows
    leafs["ipv6-egress-flows"] = statistics.Ipv6EgressFlows
    leafs["mpls-ingress-flows"] = statistics.MplsIngressFlows
    leafs["mpls-egress-flows"] = statistics.MplsEgressFlows
    leafs["drops-no-space"] = statistics.DropsNoSpace
    leafs["drops-others"] = statistics.DropsOthers
    leafs["unknown-ingress-flows"] = statistics.UnknownIngressFlows
    leafs["unknown-egress-flows"] = statistics.UnknownEgressFlows
    leafs["waiting-servers"] = statistics.WaitingServers
    leafs["last-cleared"] = statistics.LastCleared
    return leafs
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetYangName() string { return "statistics" }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetParentYangName() string { return "producer" }

// NetFlow_Statistics_Statistic_Server
// NetFlow server statistics
type NetFlow_Statistics_Statistic_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flow exporter information.
    FlowExporters NetFlow_Statistics_Statistic_Server_FlowExporters
}

func (server *NetFlow_Statistics_Statistic_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *NetFlow_Statistics_Statistic_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *NetFlow_Statistics_Statistic_Server) GetGoName(yname string) string {
    if yname == "flow-exporters" { return "FlowExporters" }
    return ""
}

func (server *NetFlow_Statistics_Statistic_Server) GetSegmentPath() string {
    return "server"
}

func (server *NetFlow_Statistics_Statistic_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-exporters" {
        return &server.FlowExporters
    }
    return nil
}

func (server *NetFlow_Statistics_Statistic_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-exporters"] = &server.FlowExporters
    return children
}

func (server *NetFlow_Statistics_Statistic_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (server *NetFlow_Statistics_Statistic_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *NetFlow_Statistics_Statistic_Server) GetYangName() string { return "server" }

func (server *NetFlow_Statistics_Statistic_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *NetFlow_Statistics_Statistic_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *NetFlow_Statistics_Statistic_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *NetFlow_Statistics_Statistic_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *NetFlow_Statistics_Statistic_Server) GetParent() types.Entity { return server.parent }

func (server *NetFlow_Statistics_Statistic_Server) GetParentYangName() string { return "statistic" }

// NetFlow_Statistics_Statistic_Server_FlowExporters
// Flow exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exporter information. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter.
    FlowExporter []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetFilter() yfilter.YFilter { return flowExporters.YFilter }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) SetFilter(yf yfilter.YFilter) { flowExporters.YFilter = yf }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetGoName(yname string) string {
    if yname == "flow-exporter" { return "FlowExporter" }
    return ""
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetSegmentPath() string {
    return "flow-exporters"
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-exporter" {
        for _, c := range flowExporters.FlowExporter {
            if flowExporters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter{}
        flowExporters.FlowExporter = append(flowExporters.FlowExporter, child)
        return &flowExporters.FlowExporter[len(flowExporters.FlowExporter)-1]
    }
    return nil
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flowExporters.FlowExporter {
        children[flowExporters.FlowExporter[i].GetSegmentPath()] = &flowExporters.FlowExporter[i]
    }
    return children
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetBundleName() string { return "cisco_ios_xr" }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetYangName() string { return "flow-exporters" }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) SetParent(parent types.Entity) { flowExporters.parent = parent }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetParent() types.Entity { return flowExporters.parent }

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetParentYangName() string { return "server" }

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
// Exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Exporter name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ExporterName interface{}

    // Statistics information for the exporter.
    Exporter NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetFilter() yfilter.YFilter { return flowExporter.YFilter }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) SetFilter(yf yfilter.YFilter) { flowExporter.YFilter = yf }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetGoName(yname string) string {
    if yname == "exporter-name" { return "ExporterName" }
    if yname == "exporter" { return "Exporter" }
    return ""
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetSegmentPath() string {
    return "flow-exporter" + "[exporter-name='" + fmt.Sprintf("%v", flowExporter.ExporterName) + "']"
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exporter" {
        return &flowExporter.Exporter
    }
    return nil
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["exporter"] = &flowExporter.Exporter
    return children
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exporter-name"] = flowExporter.ExporterName
    return leafs
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetBundleName() string { return "cisco_ios_xr" }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetYangName() string { return "flow-exporter" }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) SetParent(parent types.Entity) { flowExporter.parent = parent }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetParent() types.Entity { return flowExporter.parent }

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetParentYangName() string { return "flow-exporters" }

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter
// Statistics information for the exporter
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of flow exporters. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic.
    Statistic []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetFilter() yfilter.YFilter { return exporter.YFilter }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) SetFilter(yf yfilter.YFilter) { exporter.YFilter = yf }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetSegmentPath() string {
    return "exporter"
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range exporter.Statistic {
            if exporter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic{}
        exporter.Statistic = append(exporter.Statistic, child)
        return &exporter.Statistic[len(exporter.Statistic)-1]
    }
    return nil
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exporter.Statistic {
        children[exporter.Statistic[i].GetSegmentPath()] = &exporter.Statistic[i]
    }
    return children
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetBundleName() string { return "cisco_ios_xr" }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetYangName() string { return "exporter" }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) SetParent(parent types.Entity) { exporter.parent = parent }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetParent() types.Entity { return exporter.parent }

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetParentYangName() string { return "flow-exporter" }

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic
// Array of flow exporters
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exporter name. The type is string.
    Name interface{}

    // Memory usage. The type is interface{} with range: 0..4294967295.
    MemoryUsage interface{}

    // List of flow monitors that use the exporter. The type is slice of string.
    UsedByFlowMonitor []interface{}

    // Statistics of all collectors. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector.
    Collector []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "memory-usage" { return "MemoryUsage" }
    if yname == "used-by-flow-monitor" { return "UsedByFlowMonitor" }
    if yname == "collector" { return "Collector" }
    return ""
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetSegmentPath() string {
    return "statistic"
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "collector" {
        for _, c := range statistic.Collector {
            if statistic.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector{}
        statistic.Collector = append(statistic.Collector, child)
        return &statistic.Collector[len(statistic.Collector)-1]
    }
    return nil
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistic.Collector {
        children[statistic.Collector[i].GetSegmentPath()] = &statistic.Collector[i]
    }
    return children
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = statistic.Name
    leafs["memory-usage"] = statistic.MemoryUsage
    leafs["used-by-flow-monitor"] = statistic.UsedByFlowMonitor
    return leafs
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetYangName() string { return "statistic" }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetParentYangName() string { return "exporter" }

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector
// Statistics of all collectors
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exporter state. The type is string.
    ExporterState interface{}

    // Destination IPv4 address in AAA.BBB.CCC.DDD format. The type is string.
    DestinationAddress interface{}

    // Source IPv4 address in AAA.BBB.CCC.DDD format. The type is string.
    SourceAddress interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Destination port number. The type is interface{} with range: 0..65535.
    DestinationPort interface{}

    // Source port number. The type is interface{} with range: 0..65535.
    SoucePort interface{}

    // Transport protocol. The type is string.
    TransportProtocol interface{}

    // Packets sent. The type is interface{} with range: 0..18446744073709551615.
    PacketsSent interface{}

    // Flows sent. The type is interface{} with range: 0..18446744073709551615.
    FlowsSent interface{}

    // Templates sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TemplatesSent interface{}

    // Option templates sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OptionTemplatesSent interface{}

    // Option data sent. The type is interface{} with range:
    // 0..18446744073709551615.
    OptionDataSent interface{}

    // Bytes sent. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesSent interface{}

    // Flow bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FlowBytesSent interface{}

    // Template bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TemplateBytesSent interface{}

    // Option template bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    OptionTemplateBytesSent interface{}

    // Option data bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    OptionDataBytesSent interface{}

    // Packets dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsDropped interface{}

    // Flows dropped. The type is interface{} with range: 0..18446744073709551615.
    FlowsDropped interface{}

    // Templates dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    TemplatesDropped interface{}

    // Option templates dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    OptionTemplatesDropped interface{}

    // Option data dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    OptionDataDropped interface{}

    // Bytes dropped. The type is interface{} with range: 0..18446744073709551615.
    // Units are byte.
    BytesDropped interface{}

    // Flow bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    FlowBytesDropped interface{}

    // Template bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    TemplateBytesDropped interface{}

    // Option template bytes dropped. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    OptionTemplateBytesDropped interface{}

    // Option data dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    OptionDataBytesDropped interface{}

    // Total packets exported over the last one hour. The type is interface{} with
    // range: 0..18446744073709551615.
    LastHourPackestSent interface{}

    // Total bytes exported over the last one hour. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    LastHourBytesSent interface{}

    // Total flows exported over the of last one hour. The type is interface{}
    // with range: 0..18446744073709551615.
    LastHourFlowsSent interface{}

    // Total packets exported over the last one minute. The type is interface{}
    // with range: 0..18446744073709551615.
    LastMinutePackets interface{}

    // Total bytes exported over the last one minute. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    LastMinuteBytesSent interface{}

    // Total flows exported over the last one minute. The type is interface{} with
    // range: 0..18446744073709551615.
    LastMinuteFlowsSent interface{}

    // Total packets exported over the last one second. The type is interface{}
    // with range: 0..18446744073709551615.
    LastSecondPacketsSent interface{}

    // Total bytes exported over the last one second. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    LastSecondBytesSent interface{}

    // Total flows exported over the last one second. The type is interface{} with
    // range: 0..18446744073709551615.
    LastSecondFlowsSent interface{}
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetFilter() yfilter.YFilter { return collector.YFilter }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) SetFilter(yf yfilter.YFilter) { collector.YFilter = yf }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetGoName(yname string) string {
    if yname == "exporter-state" { return "ExporterState" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "souce-port" { return "SoucePort" }
    if yname == "transport-protocol" { return "TransportProtocol" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "flows-sent" { return "FlowsSent" }
    if yname == "templates-sent" { return "TemplatesSent" }
    if yname == "option-templates-sent" { return "OptionTemplatesSent" }
    if yname == "option-data-sent" { return "OptionDataSent" }
    if yname == "bytes-sent" { return "BytesSent" }
    if yname == "flow-bytes-sent" { return "FlowBytesSent" }
    if yname == "template-bytes-sent" { return "TemplateBytesSent" }
    if yname == "option-template-bytes-sent" { return "OptionTemplateBytesSent" }
    if yname == "option-data-bytes-sent" { return "OptionDataBytesSent" }
    if yname == "packets-dropped" { return "PacketsDropped" }
    if yname == "flows-dropped" { return "FlowsDropped" }
    if yname == "templates-dropped" { return "TemplatesDropped" }
    if yname == "option-templates-dropped" { return "OptionTemplatesDropped" }
    if yname == "option-data-dropped" { return "OptionDataDropped" }
    if yname == "bytes-dropped" { return "BytesDropped" }
    if yname == "flow-bytes-dropped" { return "FlowBytesDropped" }
    if yname == "template-bytes-dropped" { return "TemplateBytesDropped" }
    if yname == "option-template-bytes-dropped" { return "OptionTemplateBytesDropped" }
    if yname == "option-data-bytes-dropped" { return "OptionDataBytesDropped" }
    if yname == "last-hour-packest-sent" { return "LastHourPackestSent" }
    if yname == "last-hour-bytes-sent" { return "LastHourBytesSent" }
    if yname == "last-hour-flows-sent" { return "LastHourFlowsSent" }
    if yname == "last-minute-packets" { return "LastMinutePackets" }
    if yname == "last-minute-bytes-sent" { return "LastMinuteBytesSent" }
    if yname == "last-minute-flows-sent" { return "LastMinuteFlowsSent" }
    if yname == "last-second-packets-sent" { return "LastSecondPacketsSent" }
    if yname == "last-second-bytes-sent" { return "LastSecondBytesSent" }
    if yname == "last-second-flows-sent" { return "LastSecondFlowsSent" }
    return ""
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetSegmentPath() string {
    return "collector"
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exporter-state"] = collector.ExporterState
    leafs["destination-address"] = collector.DestinationAddress
    leafs["source-address"] = collector.SourceAddress
    leafs["vrf-name"] = collector.VrfName
    leafs["destination-port"] = collector.DestinationPort
    leafs["souce-port"] = collector.SoucePort
    leafs["transport-protocol"] = collector.TransportProtocol
    leafs["packets-sent"] = collector.PacketsSent
    leafs["flows-sent"] = collector.FlowsSent
    leafs["templates-sent"] = collector.TemplatesSent
    leafs["option-templates-sent"] = collector.OptionTemplatesSent
    leafs["option-data-sent"] = collector.OptionDataSent
    leafs["bytes-sent"] = collector.BytesSent
    leafs["flow-bytes-sent"] = collector.FlowBytesSent
    leafs["template-bytes-sent"] = collector.TemplateBytesSent
    leafs["option-template-bytes-sent"] = collector.OptionTemplateBytesSent
    leafs["option-data-bytes-sent"] = collector.OptionDataBytesSent
    leafs["packets-dropped"] = collector.PacketsDropped
    leafs["flows-dropped"] = collector.FlowsDropped
    leafs["templates-dropped"] = collector.TemplatesDropped
    leafs["option-templates-dropped"] = collector.OptionTemplatesDropped
    leafs["option-data-dropped"] = collector.OptionDataDropped
    leafs["bytes-dropped"] = collector.BytesDropped
    leafs["flow-bytes-dropped"] = collector.FlowBytesDropped
    leafs["template-bytes-dropped"] = collector.TemplateBytesDropped
    leafs["option-template-bytes-dropped"] = collector.OptionTemplateBytesDropped
    leafs["option-data-bytes-dropped"] = collector.OptionDataBytesDropped
    leafs["last-hour-packest-sent"] = collector.LastHourPackestSent
    leafs["last-hour-bytes-sent"] = collector.LastHourBytesSent
    leafs["last-hour-flows-sent"] = collector.LastHourFlowsSent
    leafs["last-minute-packets"] = collector.LastMinutePackets
    leafs["last-minute-bytes-sent"] = collector.LastMinuteBytesSent
    leafs["last-minute-flows-sent"] = collector.LastMinuteFlowsSent
    leafs["last-second-packets-sent"] = collector.LastSecondPacketsSent
    leafs["last-second-bytes-sent"] = collector.LastSecondBytesSent
    leafs["last-second-flows-sent"] = collector.LastSecondFlowsSent
    return leafs
}

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetBundleName() string { return "cisco_ios_xr" }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetYangName() string { return "collector" }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) SetParent(parent types.Entity) { collector.parent = parent }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetParent() types.Entity { return collector.parent }

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetParentYangName() string { return "statistic" }

