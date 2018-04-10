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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific NetFlow statistics information.
    Statistics NetFlow_Statistics
}

func (netFlow *NetFlow) GetEntityData() *types.CommonEntityData {
    netFlow.EntityData.YFilter = netFlow.YFilter
    netFlow.EntityData.YangName = "net-flow"
    netFlow.EntityData.BundleName = "cisco_ios_xr"
    netFlow.EntityData.ParentYangName = "Cisco-IOS-XR-dnx-netflow-oper"
    netFlow.EntityData.SegmentPath = "Cisco-IOS-XR-dnx-netflow-oper:net-flow"
    netFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netFlow.EntityData.Children = make(map[string]types.YChild)
    netFlow.EntityData.Children["statistics"] = types.YChild{"Statistics", &netFlow.Statistics}
    netFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netFlow.EntityData)
}

// NetFlow_Statistics
// Node-specific NetFlow statistics information
type NetFlow_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NetFlow statistics information for a particular node. The type is slice of
    // NetFlow_Statistics_Statistic.
    Statistic []NetFlow_Statistics_Statistic
}

func (statistics *NetFlow_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "net-flow"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["statistic"] = types.YChild{"Statistic", nil}
    for i := range statistics.Statistic {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Statistic[i])] = types.YChild{"Statistic", &statistics.Statistic[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// NetFlow_Statistics_Statistic
// NetFlow statistics information for a particular
// node
type NetFlow_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    statistic.EntityData.SegmentPath = "statistic" + "[node='" + fmt.Sprintf("%v", statistic.Node) + "']"
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = make(map[string]types.YChild)
    statistic.EntityData.Children["producer"] = types.YChild{"Producer", &statistic.Producer}
    statistic.EntityData.Children["server"] = types.YChild{"Server", &statistic.Server}
    statistic.EntityData.Leafs = make(map[string]types.YLeaf)
    statistic.EntityData.Leafs["node"] = types.YLeaf{"Node", statistic.Node}
    return &(statistic.EntityData)
}

// NetFlow_Statistics_Statistic_Producer
// NetFlow producer statistics
type NetFlow_Statistics_Statistic_Producer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics information.
    Statistics NetFlow_Statistics_Statistic_Producer_Statistics_
}

func (producer *NetFlow_Statistics_Statistic_Producer) GetEntityData() *types.CommonEntityData {
    producer.EntityData.YFilter = producer.YFilter
    producer.EntityData.YangName = "producer"
    producer.EntityData.BundleName = "cisco_ios_xr"
    producer.EntityData.ParentYangName = "statistic"
    producer.EntityData.SegmentPath = "producer"
    producer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    producer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    producer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    producer.EntityData.Children = make(map[string]types.YChild)
    producer.EntityData.Children["statistics"] = types.YChild{"Statistics", &producer.Statistics}
    producer.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(producer.EntityData)
}

// NetFlow_Statistics_Statistic_Producer_Statistics_
// Statistics information
type NetFlow_Statistics_Statistic_Producer_Statistics_ struct {
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

    // IPFIX315 ingress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipfix315IngressFlows interface{}

    // IPFIX315 egress flows. The type is interface{} with range:
    // 0..18446744073709551615.
    Ipfix315EgressFlows interface{}

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

func (statistics_ *NetFlow_Statistics_Statistic_Producer_Statistics_) GetEntityData() *types.CommonEntityData {
    statistics_.EntityData.YFilter = statistics_.YFilter
    statistics_.EntityData.YangName = "statistics"
    statistics_.EntityData.BundleName = "cisco_ios_xr"
    statistics_.EntityData.ParentYangName = "producer"
    statistics_.EntityData.SegmentPath = "statistics"
    statistics_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics_.EntityData.Children = make(map[string]types.YChild)
    statistics_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics_.EntityData.Leafs["ipv4-ingress-flows"] = types.YLeaf{"Ipv4IngressFlows", statistics_.Ipv4IngressFlows}
    statistics_.EntityData.Leafs["ipv4-egress-flows"] = types.YLeaf{"Ipv4EgressFlows", statistics_.Ipv4EgressFlows}
    statistics_.EntityData.Leafs["ipv6-ingress-flows"] = types.YLeaf{"Ipv6IngressFlows", statistics_.Ipv6IngressFlows}
    statistics_.EntityData.Leafs["ipv6-egress-flows"] = types.YLeaf{"Ipv6EgressFlows", statistics_.Ipv6EgressFlows}
    statistics_.EntityData.Leafs["mpls-ingress-flows"] = types.YLeaf{"MplsIngressFlows", statistics_.MplsIngressFlows}
    statistics_.EntityData.Leafs["mpls-egress-flows"] = types.YLeaf{"MplsEgressFlows", statistics_.MplsEgressFlows}
    statistics_.EntityData.Leafs["ipfix315-ingress-flows"] = types.YLeaf{"Ipfix315IngressFlows", statistics_.Ipfix315IngressFlows}
    statistics_.EntityData.Leafs["ipfix315-egress-flows"] = types.YLeaf{"Ipfix315EgressFlows", statistics_.Ipfix315EgressFlows}
    statistics_.EntityData.Leafs["drops-no-space"] = types.YLeaf{"DropsNoSpace", statistics_.DropsNoSpace}
    statistics_.EntityData.Leafs["drops-others"] = types.YLeaf{"DropsOthers", statistics_.DropsOthers}
    statistics_.EntityData.Leafs["unknown-ingress-flows"] = types.YLeaf{"UnknownIngressFlows", statistics_.UnknownIngressFlows}
    statistics_.EntityData.Leafs["unknown-egress-flows"] = types.YLeaf{"UnknownEgressFlows", statistics_.UnknownEgressFlows}
    statistics_.EntityData.Leafs["waiting-servers"] = types.YLeaf{"WaitingServers", statistics_.WaitingServers}
    statistics_.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", statistics_.LastCleared}
    return &(statistics_.EntityData)
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
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["flow-exporters"] = types.YChild{"FlowExporters", &server.FlowExporters}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(server.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters
// Flow exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exporter information. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter.
    FlowExporter []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
}

func (flowExporters *NetFlow_Statistics_Statistic_Server_FlowExporters) GetEntityData() *types.CommonEntityData {
    flowExporters.EntityData.YFilter = flowExporters.YFilter
    flowExporters.EntityData.YangName = "flow-exporters"
    flowExporters.EntityData.BundleName = "cisco_ios_xr"
    flowExporters.EntityData.ParentYangName = "server"
    flowExporters.EntityData.SegmentPath = "flow-exporters"
    flowExporters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporters.EntityData.Children = make(map[string]types.YChild)
    flowExporters.EntityData.Children["flow-exporter"] = types.YChild{"FlowExporter", nil}
    for i := range flowExporters.FlowExporter {
        flowExporters.EntityData.Children[types.GetSegmentPath(&flowExporters.FlowExporter[i])] = types.YChild{"FlowExporter", &flowExporters.FlowExporter[i]}
    }
    flowExporters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowExporters.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter
// Exporter information
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    flowExporter.EntityData.SegmentPath = "flow-exporter" + "[exporter-name='" + fmt.Sprintf("%v", flowExporter.ExporterName) + "']"
    flowExporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowExporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowExporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowExporter.EntityData.Children = make(map[string]types.YChild)
    flowExporter.EntityData.Children["exporter"] = types.YChild{"Exporter", &flowExporter.Exporter}
    flowExporter.EntityData.Leafs = make(map[string]types.YLeaf)
    flowExporter.EntityData.Leafs["exporter-name"] = types.YLeaf{"ExporterName", flowExporter.ExporterName}
    return &(flowExporter.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter
// Statistics information for the exporter
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of flow exporters. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic.
    Statistic []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_
}

func (exporter *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter) GetEntityData() *types.CommonEntityData {
    exporter.EntityData.YFilter = exporter.YFilter
    exporter.EntityData.YangName = "exporter"
    exporter.EntityData.BundleName = "cisco_ios_xr"
    exporter.EntityData.ParentYangName = "flow-exporter"
    exporter.EntityData.SegmentPath = "exporter"
    exporter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exporter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exporter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exporter.EntityData.Children = make(map[string]types.YChild)
    exporter.EntityData.Children["statistic"] = types.YChild{"Statistic", nil}
    for i := range exporter.Statistic {
        exporter.EntityData.Children[types.GetSegmentPath(&exporter.Statistic[i])] = types.YChild{"Statistic", &exporter.Statistic[i]}
    }
    exporter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exporter.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_
// Array of flow exporters
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exporter name. The type is string.
    Name interface{}

    // Memory usage. The type is interface{} with range: 0..4294967295.
    MemoryUsage interface{}

    // List of flow monitors that use the exporter. The type is slice of string.
    UsedByFlowMonitor []interface{}

    // Statistics of all collectors. The type is slice of
    // NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic__Collector.
    Collector []NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic__Collector
}

func (statistic_ *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic_) GetEntityData() *types.CommonEntityData {
    statistic_.EntityData.YFilter = statistic_.YFilter
    statistic_.EntityData.YangName = "statistic"
    statistic_.EntityData.BundleName = "cisco_ios_xr"
    statistic_.EntityData.ParentYangName = "exporter"
    statistic_.EntityData.SegmentPath = "statistic"
    statistic_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic_.EntityData.Children = make(map[string]types.YChild)
    statistic_.EntityData.Children["collector"] = types.YChild{"Collector", nil}
    for i := range statistic_.Collector {
        statistic_.EntityData.Children[types.GetSegmentPath(&statistic_.Collector[i])] = types.YChild{"Collector", &statistic_.Collector[i]}
    }
    statistic_.EntityData.Leafs = make(map[string]types.YLeaf)
    statistic_.EntityData.Leafs["name"] = types.YLeaf{"Name", statistic_.Name}
    statistic_.EntityData.Leafs["memory-usage"] = types.YLeaf{"MemoryUsage", statistic_.MemoryUsage}
    statistic_.EntityData.Leafs["used-by-flow-monitor"] = types.YLeaf{"UsedByFlowMonitor", statistic_.UsedByFlowMonitor}
    return &(statistic_.EntityData)
}

// NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic__Collector
// Statistics of all collectors
type NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic__Collector struct {
    EntityData types.CommonEntityData
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

func (collector *NetFlow_Statistics_Statistic_Server_FlowExporters_FlowExporter_Exporter_Statistic__Collector) GetEntityData() *types.CommonEntityData {
    collector.EntityData.YFilter = collector.YFilter
    collector.EntityData.YangName = "collector"
    collector.EntityData.BundleName = "cisco_ios_xr"
    collector.EntityData.ParentYangName = "statistic"
    collector.EntityData.SegmentPath = "collector"
    collector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collector.EntityData.Children = make(map[string]types.YChild)
    collector.EntityData.Leafs = make(map[string]types.YLeaf)
    collector.EntityData.Leafs["exporter-state"] = types.YLeaf{"ExporterState", collector.ExporterState}
    collector.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", collector.DestinationAddress}
    collector.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", collector.SourceAddress}
    collector.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", collector.VrfName}
    collector.EntityData.Leafs["destination-port"] = types.YLeaf{"DestinationPort", collector.DestinationPort}
    collector.EntityData.Leafs["souce-port"] = types.YLeaf{"SoucePort", collector.SoucePort}
    collector.EntityData.Leafs["transport-protocol"] = types.YLeaf{"TransportProtocol", collector.TransportProtocol}
    collector.EntityData.Leafs["packets-sent"] = types.YLeaf{"PacketsSent", collector.PacketsSent}
    collector.EntityData.Leafs["flows-sent"] = types.YLeaf{"FlowsSent", collector.FlowsSent}
    collector.EntityData.Leafs["templates-sent"] = types.YLeaf{"TemplatesSent", collector.TemplatesSent}
    collector.EntityData.Leafs["option-templates-sent"] = types.YLeaf{"OptionTemplatesSent", collector.OptionTemplatesSent}
    collector.EntityData.Leafs["option-data-sent"] = types.YLeaf{"OptionDataSent", collector.OptionDataSent}
    collector.EntityData.Leafs["bytes-sent"] = types.YLeaf{"BytesSent", collector.BytesSent}
    collector.EntityData.Leafs["flow-bytes-sent"] = types.YLeaf{"FlowBytesSent", collector.FlowBytesSent}
    collector.EntityData.Leafs["template-bytes-sent"] = types.YLeaf{"TemplateBytesSent", collector.TemplateBytesSent}
    collector.EntityData.Leafs["option-template-bytes-sent"] = types.YLeaf{"OptionTemplateBytesSent", collector.OptionTemplateBytesSent}
    collector.EntityData.Leafs["option-data-bytes-sent"] = types.YLeaf{"OptionDataBytesSent", collector.OptionDataBytesSent}
    collector.EntityData.Leafs["packets-dropped"] = types.YLeaf{"PacketsDropped", collector.PacketsDropped}
    collector.EntityData.Leafs["flows-dropped"] = types.YLeaf{"FlowsDropped", collector.FlowsDropped}
    collector.EntityData.Leafs["templates-dropped"] = types.YLeaf{"TemplatesDropped", collector.TemplatesDropped}
    collector.EntityData.Leafs["option-templates-dropped"] = types.YLeaf{"OptionTemplatesDropped", collector.OptionTemplatesDropped}
    collector.EntityData.Leafs["option-data-dropped"] = types.YLeaf{"OptionDataDropped", collector.OptionDataDropped}
    collector.EntityData.Leafs["bytes-dropped"] = types.YLeaf{"BytesDropped", collector.BytesDropped}
    collector.EntityData.Leafs["flow-bytes-dropped"] = types.YLeaf{"FlowBytesDropped", collector.FlowBytesDropped}
    collector.EntityData.Leafs["template-bytes-dropped"] = types.YLeaf{"TemplateBytesDropped", collector.TemplateBytesDropped}
    collector.EntityData.Leafs["option-template-bytes-dropped"] = types.YLeaf{"OptionTemplateBytesDropped", collector.OptionTemplateBytesDropped}
    collector.EntityData.Leafs["option-data-bytes-dropped"] = types.YLeaf{"OptionDataBytesDropped", collector.OptionDataBytesDropped}
    collector.EntityData.Leafs["last-hour-packest-sent"] = types.YLeaf{"LastHourPackestSent", collector.LastHourPackestSent}
    collector.EntityData.Leafs["last-hour-bytes-sent"] = types.YLeaf{"LastHourBytesSent", collector.LastHourBytesSent}
    collector.EntityData.Leafs["last-hour-flows-sent"] = types.YLeaf{"LastHourFlowsSent", collector.LastHourFlowsSent}
    collector.EntityData.Leafs["last-minute-packets"] = types.YLeaf{"LastMinutePackets", collector.LastMinutePackets}
    collector.EntityData.Leafs["last-minute-bytes-sent"] = types.YLeaf{"LastMinuteBytesSent", collector.LastMinuteBytesSent}
    collector.EntityData.Leafs["last-minute-flows-sent"] = types.YLeaf{"LastMinuteFlowsSent", collector.LastMinuteFlowsSent}
    collector.EntityData.Leafs["last-second-packets-sent"] = types.YLeaf{"LastSecondPacketsSent", collector.LastSecondPacketsSent}
    collector.EntityData.Leafs["last-second-bytes-sent"] = types.YLeaf{"LastSecondBytesSent", collector.LastSecondBytesSent}
    collector.EntityData.Leafs["last-second-flows-sent"] = types.YLeaf{"LastSecondFlowsSent", collector.LastSecondFlowsSent}
    return &(collector.EntityData)
}

