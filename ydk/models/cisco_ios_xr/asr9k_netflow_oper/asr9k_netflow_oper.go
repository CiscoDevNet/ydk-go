// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-netflow package operational data.
// 
// This module contains definitions
// for the following management objects:
//   net-flow: NetFlow operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_netflow_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_netflow_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-netflow-oper net-flow}", reflect.TypeOf(NetFlow{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-netflow-oper:net-flow", reflect.TypeOf(NetFlow{}))
}

// NetFlow
// NetFlow operational data
type NetFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific NetFlow statistics information.
    Statistics NetFlow_Statistics
}

func (netFlow *NetFlow) GetEntityData() *types.CommonEntityData {
    netFlow.EntityData.YFilter = netFlow.YFilter
    netFlow.EntityData.YangName = "net-flow"
    netFlow.EntityData.BundleName = "cisco_ios_xr"
    netFlow.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-netflow-oper"
    netFlow.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow"
    netFlow.EntityData.AbsolutePath = netFlow.EntityData.SegmentPath
    netFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netFlow.EntityData.Children = types.NewOrderedMap()
    netFlow.EntityData.Children.Append("statistics", types.YChild{"Statistics", &netFlow.Statistics})
    netFlow.EntityData.Leafs = types.NewOrderedMap()

    netFlow.EntityData.YListKeys = []string {}

    return &(netFlow.EntityData)
}

// NetFlow_Statistics
// Node-specific NetFlow statistics information
type NetFlow_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NetFlow statistics information for a particular node. The type is slice of
    // NetFlow_Statistics_Statistic.
    Statistic []*NetFlow_Statistics_Statistic
}

func (statistics *NetFlow_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "net-flow"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("statistic", types.YChild{"Statistic", nil})
    for i := range statistics.Statistic {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Statistic[i]), types.YChild{"Statistic", statistics.Statistic[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// NetFlow_Statistics_Statistic
// NetFlow statistics information for a particular
// node
type NetFlow_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // NetFlow producer statistics.
    Producer NetFlow_Statistics_Statistic_Producer

    // NetFlow server statistics.
    Server NetFlow_Statistics_Statistic_Server
}

func (statistic *NetFlow_Statistics_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "statistics"
    statistic.EntityData.SegmentPath = "statistic" + types.AddKeyToken(statistic.Node, "node")
    statistic.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/" + statistic.EntityData.SegmentPath
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = types.NewOrderedMap()
    statistic.EntityData.Children.Append("producer", types.YChild{"Producer", &statistic.Producer})
    statistic.EntityData.Children.Append("server", types.YChild{"Server", &statistic.Server})
    statistic.EntityData.Leafs = types.NewOrderedMap()
    statistic.EntityData.Leafs.Append("node", types.YLeaf{"Node", statistic.Node})

    statistic.EntityData.YListKeys = []string {"Node"}

    return &(statistic.EntityData)
}

// NetFlow_Statistics_Statistic_Producer
// NetFlow producer statistics
type NetFlow_Statistics_Statistic_Producer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics information.
    Statistics NetFlow_Statistics_Statistic_Producer_Statistics
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetEntityData() *types.CommonEntityData {
    producer.EntityData.YFilter = producer.YFilter
    producer.EntityData.YangName = "producer"
    producer.EntityData.BundleName = "cisco_ios_xr"
    producer.EntityData.ParentYangName = "statistic"
    producer.EntityData.SegmentPath = "producer"
    producer.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/" + producer.EntityData.SegmentPath
    producer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producer.EntityData.Children = types.NewOrderedMap()
    producer.EntityData.Children.Append("statistics", types.YChild{"Statistics", &producer.Statistics})
    producer.EntityData.Leafs = types.NewOrderedMap()

    producer.EntityData.YListKeys = []string {}

    return &(producer.EntityData)
}

// NetFlow_Statistics_Statistic_Producer_Statistics
// Statistics information
type NetFlow_Statistics_Statistic_Producer_Statistics struct {
    EntityData types.CommonEntityData
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

    // Section ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    SectionIngressFlows interface{}

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

    // Number of Rxed SPP Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SppRxCounts interface{}

    // Number of Rxed Flow Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowPacketCounts interface{}

    // Last time Statistics cleared in 'Mon Jan 1 12:00 :00 2xxx' format. The type
    // is string.
    LastCleared interface{}
}

func (statistics *NetFlow_Statistics_Statistic_Producer_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "producer"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/producer/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("ipv4-ingress-flows", types.YLeaf{"Ipv4IngressFlows", statistics.Ipv4IngressFlows})
    statistics.EntityData.Leafs.Append("ipv4-egress-flows", types.YLeaf{"Ipv4EgressFlows", statistics.Ipv4EgressFlows})
    statistics.EntityData.Leafs.Append("ipv6-ingress-flows", types.YLeaf{"Ipv6IngressFlows", statistics.Ipv6IngressFlows})
    statistics.EntityData.Leafs.Append("ipv6-egress-flows", types.YLeaf{"Ipv6EgressFlows", statistics.Ipv6EgressFlows})
    statistics.EntityData.Leafs.Append("mpls-ingress-flows", types.YLeaf{"MplsIngressFlows", statistics.MplsIngressFlows})
    statistics.EntityData.Leafs.Append("mpls-egress-flows", types.YLeaf{"MplsEgressFlows", statistics.MplsEgressFlows})
    statistics.EntityData.Leafs.Append("section-ingress-flows", types.YLeaf{"SectionIngressFlows", statistics.SectionIngressFlows})
    statistics.EntityData.Leafs.Append("drops-no-space", types.YLeaf{"DropsNoSpace", statistics.DropsNoSpace})
    statistics.EntityData.Leafs.Append("drops-others", types.YLeaf{"DropsOthers", statistics.DropsOthers})
    statistics.EntityData.Leafs.Append("unknown-ingress-flows", types.YLeaf{"UnknownIngressFlows", statistics.UnknownIngressFlows})
    statistics.EntityData.Leafs.Append("unknown-egress-flows", types.YLeaf{"UnknownEgressFlows", statistics.UnknownEgressFlows})
    statistics.EntityData.Leafs.Append("waiting-servers", types.YLeaf{"WaitingServers", statistics.WaitingServers})
    statistics.EntityData.Leafs.Append("spp-rx-counts", types.YLeaf{"SppRxCounts", statistics.SppRxCounts})
    statistics.EntityData.Leafs.Append("flow-packet-counts", types.YLeaf{"FlowPacketCounts", statistics.FlowPacketCounts})
    statistics.EntityData.Leafs.Append("last-cleared", types.YLeaf{"LastCleared", statistics.LastCleared})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// NetFlow_Statistics_Statistic_Server
// NetFlow server statistics
type NetFlow_Statistics_Statistic_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow exporter information.
    FlowExporters NetFlow_Statistics_Statistic_Server_FlowExporters
}

func (server *NetFlow_Statistics_Statistic_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "statistic"
    server.EntityData.SegmentPath = "server"
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("flow-exporters", types.YChild{"FlowExporters", &server.FlowExporters})
    server.EntityData.Leafs = types.NewOrderedMap()

    server.EntityData.YListKeys = []string {}

    return &(server.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters
// Flow exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exporter information. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter.
    FlowExporter []*NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetEntityData() *types.CommonEntityData {
    flowExporters.EntityData.YFilter = flowExporters.YFilter
    flowExporters.EntityData.YangName = "flow-exporters"
    flowExporters.EntityData.BundleName = "cisco_ios_xr"
    flowExporters.EntityData.ParentYangName = "server"
    flowExporters.EntityData.SegmentPath = "flow-exporters"
    flowExporters.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/server/" + flowExporters.EntityData.SegmentPath
    flowExporters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporters.EntityData.Children = types.NewOrderedMap()
    flowExporters.EntityData.Children.Append("flow-exporter", types.YChild{"FlowExporter", nil})
    for i := range flowExporters.FlowExporter {
        flowExporters.EntityData.Children.Append(types.GetSegmentPath(flowExporters.FlowExporter[i]), types.YChild{"FlowExporter", flowExporters.FlowExporter[i]})
    }
    flowExporters.EntityData.Leafs = types.NewOrderedMap()

    flowExporters.EntityData.YListKeys = []string {}

    return &(flowExporters.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
// Exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Exporter name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ExporterName interface{}

    // Statistics information for the exporter.
    Exporter NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter
}

func (flowExporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter) GetEntityData() *types.CommonEntityData {
    flowExporter.EntityData.YFilter = flowExporter.YFilter
    flowExporter.EntityData.YangName = "flow-exporter"
    flowExporter.EntityData.BundleName = "cisco_ios_xr"
    flowExporter.EntityData.ParentYangName = "flow-exporters"
    flowExporter.EntityData.SegmentPath = "flow-exporter" + types.AddKeyToken(flowExporter.ExporterName, "exporter-name")
    flowExporter.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/server/flow-exporters/" + flowExporter.EntityData.SegmentPath
    flowExporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporter.EntityData.Children = types.NewOrderedMap()
    flowExporter.EntityData.Children.Append("exporter", types.YChild{"Exporter", &flowExporter.Exporter})
    flowExporter.EntityData.Leafs = types.NewOrderedMap()
    flowExporter.EntityData.Leafs.Append("exporter-name", types.YLeaf{"ExporterName", flowExporter.ExporterName})

    flowExporter.EntityData.YListKeys = []string {"ExporterName"}

    return &(flowExporter.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter
// Statistics information for the exporter
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of flow exporters. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic.
    Statistic []*NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetEntityData() *types.CommonEntityData {
    exporter.EntityData.YFilter = exporter.YFilter
    exporter.EntityData.YangName = "exporter"
    exporter.EntityData.BundleName = "cisco_ios_xr"
    exporter.EntityData.ParentYangName = "flow-exporter"
    exporter.EntityData.SegmentPath = "exporter"
    exporter.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/server/flow-exporters/flow-exporter/" + exporter.EntityData.SegmentPath
    exporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporter.EntityData.Children = types.NewOrderedMap()
    exporter.EntityData.Children.Append("statistic", types.YChild{"Statistic", nil})
    for i := range exporter.Statistic {
        types.SetYListKey(exporter.Statistic[i], i)
        exporter.EntityData.Children.Append(types.GetSegmentPath(exporter.Statistic[i]), types.YChild{"Statistic", exporter.Statistic[i]})
    }
    exporter.EntityData.Leafs = types.NewOrderedMap()

    exporter.EntityData.YListKeys = []string {}

    return &(exporter.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic
// Array of flow exporters
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Exporter name. The type is string.
    Name interface{}

    // Memory usage. The type is interface{} with range: 0..4294967295.
    MemoryUsage interface{}

    // List of flow monitors that use the exporter. The type is slice of string.
    UsedByFlowMonitor []interface{}

    // Statistics of all collectors. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector.
    Collector []*NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector
}

func (statistic *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "exporter"
    statistic.EntityData.SegmentPath = "statistic" + types.AddNoKeyToken(statistic)
    statistic.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/server/flow-exporters/flow-exporter/exporter/" + statistic.EntityData.SegmentPath
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = types.NewOrderedMap()
    statistic.EntityData.Children.Append("collector", types.YChild{"Collector", nil})
    for i := range statistic.Collector {
        types.SetYListKey(statistic.Collector[i], i)
        statistic.EntityData.Children.Append(types.GetSegmentPath(statistic.Collector[i]), types.YChild{"Collector", statistic.Collector[i]})
    }
    statistic.EntityData.Leafs = types.NewOrderedMap()
    statistic.EntityData.Leafs.Append("name", types.YLeaf{"Name", statistic.Name})
    statistic.EntityData.Leafs.Append("memory-usage", types.YLeaf{"MemoryUsage", statistic.MemoryUsage})
    statistic.EntityData.Leafs.Append("used-by-flow-monitor", types.YLeaf{"UsedByFlowMonitor", statistic.UsedByFlowMonitor})

    statistic.EntityData.YListKeys = []string {}

    return &(statistic.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector
// Statistics of all collectors
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_Collector) GetEntityData() *types.CommonEntityData {
    collector.EntityData.YFilter = collector.YFilter
    collector.EntityData.YangName = "collector"
    collector.EntityData.BundleName = "cisco_ios_xr"
    collector.EntityData.ParentYangName = "statistic"
    collector.EntityData.SegmentPath = "collector" + types.AddNoKeyToken(collector)
    collector.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-netflow-oper:net-flow/statistics/statistic/server/flow-exporters/flow-exporter/exporter/statistic/" + collector.EntityData.SegmentPath
    collector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collector.EntityData.Children = types.NewOrderedMap()
    collector.EntityData.Leafs = types.NewOrderedMap()
    collector.EntityData.Leafs.Append("exporter-state", types.YLeaf{"ExporterState", collector.ExporterState})
    collector.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", collector.DestinationAddress})
    collector.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", collector.SourceAddress})
    collector.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", collector.VrfName})
    collector.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", collector.DestinationPort})
    collector.EntityData.Leafs.Append("souce-port", types.YLeaf{"SoucePort", collector.SoucePort})
    collector.EntityData.Leafs.Append("transport-protocol", types.YLeaf{"TransportProtocol", collector.TransportProtocol})
    collector.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", collector.PacketsSent})
    collector.EntityData.Leafs.Append("flows-sent", types.YLeaf{"FlowsSent", collector.FlowsSent})
    collector.EntityData.Leafs.Append("templates-sent", types.YLeaf{"TemplatesSent", collector.TemplatesSent})
    collector.EntityData.Leafs.Append("option-templates-sent", types.YLeaf{"OptionTemplatesSent", collector.OptionTemplatesSent})
    collector.EntityData.Leafs.Append("option-data-sent", types.YLeaf{"OptionDataSent", collector.OptionDataSent})
    collector.EntityData.Leafs.Append("bytes-sent", types.YLeaf{"BytesSent", collector.BytesSent})
    collector.EntityData.Leafs.Append("flow-bytes-sent", types.YLeaf{"FlowBytesSent", collector.FlowBytesSent})
    collector.EntityData.Leafs.Append("template-bytes-sent", types.YLeaf{"TemplateBytesSent", collector.TemplateBytesSent})
    collector.EntityData.Leafs.Append("option-template-bytes-sent", types.YLeaf{"OptionTemplateBytesSent", collector.OptionTemplateBytesSent})
    collector.EntityData.Leafs.Append("option-data-bytes-sent", types.YLeaf{"OptionDataBytesSent", collector.OptionDataBytesSent})
    collector.EntityData.Leafs.Append("packets-dropped", types.YLeaf{"PacketsDropped", collector.PacketsDropped})
    collector.EntityData.Leafs.Append("flows-dropped", types.YLeaf{"FlowsDropped", collector.FlowsDropped})
    collector.EntityData.Leafs.Append("templates-dropped", types.YLeaf{"TemplatesDropped", collector.TemplatesDropped})
    collector.EntityData.Leafs.Append("option-templates-dropped", types.YLeaf{"OptionTemplatesDropped", collector.OptionTemplatesDropped})
    collector.EntityData.Leafs.Append("option-data-dropped", types.YLeaf{"OptionDataDropped", collector.OptionDataDropped})
    collector.EntityData.Leafs.Append("bytes-dropped", types.YLeaf{"BytesDropped", collector.BytesDropped})
    collector.EntityData.Leafs.Append("flow-bytes-dropped", types.YLeaf{"FlowBytesDropped", collector.FlowBytesDropped})
    collector.EntityData.Leafs.Append("template-bytes-dropped", types.YLeaf{"TemplateBytesDropped", collector.TemplateBytesDropped})
    collector.EntityData.Leafs.Append("option-template-bytes-dropped", types.YLeaf{"OptionTemplateBytesDropped", collector.OptionTemplateBytesDropped})
    collector.EntityData.Leafs.Append("option-data-bytes-dropped", types.YLeaf{"OptionDataBytesDropped", collector.OptionDataBytesDropped})
    collector.EntityData.Leafs.Append("last-hour-packest-sent", types.YLeaf{"LastHourPackestSent", collector.LastHourPackestSent})
    collector.EntityData.Leafs.Append("last-hour-bytes-sent", types.YLeaf{"LastHourBytesSent", collector.LastHourBytesSent})
    collector.EntityData.Leafs.Append("last-hour-flows-sent", types.YLeaf{"LastHourFlowsSent", collector.LastHourFlowsSent})
    collector.EntityData.Leafs.Append("last-minute-packets", types.YLeaf{"LastMinutePackets", collector.LastMinutePackets})
    collector.EntityData.Leafs.Append("last-minute-bytes-sent", types.YLeaf{"LastMinuteBytesSent", collector.LastMinuteBytesSent})
    collector.EntityData.Leafs.Append("last-minute-flows-sent", types.YLeaf{"LastMinuteFlowsSent", collector.LastMinuteFlowsSent})
    collector.EntityData.Leafs.Append("last-second-packets-sent", types.YLeaf{"LastSecondPacketsSent", collector.LastSecondPacketsSent})
    collector.EntityData.Leafs.Append("last-second-bytes-sent", types.YLeaf{"LastSecondBytesSent", collector.LastSecondBytesSent})
    collector.EntityData.Leafs.Append("last-second-flows-sent", types.YLeaf{"LastSecondFlowsSent", collector.LastSecondFlowsSent})

    collector.EntityData.YListKeys = []string {}

    return &(collector.EntityData)
}

