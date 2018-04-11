// This module contains a collection of YANG definitions
// for Flexible NetFlow Operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package flow_monitor_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package flow_monitor_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-flow-monitor-oper flow-monitors}", reflect.TypeOf(FlowMonitors{}))
    ydk.RegisterEntity("Cisco-IOS-XE-flow-monitor-oper:flow-monitors", reflect.TypeOf(FlowMonitors{}))
}

// FlowExporterIpwriteStatsType represents The Netflow export statistics
type FlowExporterIpwriteStatsType string

const (
    // Normal Statistics event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_ok FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-ok"

    // No Forwarding Information Base event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_no_fib FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-no-fib"

    // Adjacency failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_fail_event FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-fail-event"

    // Process switch event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_process FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-process"

    // Enqueue Failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_enqueue_failed FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-enqueue-failed"

    // IPC failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_ipc_failed FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-ipc-failed"

    // Output failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_output_failed FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-output-failed"

    // Maximum Transmission Unit failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_mtu_failed FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-mtu-failed"

    // Encapsulation Fixup failed event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_encapfix_failed FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-encapfix-failed"

    // Cisco Express Forwarding off event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_cef_off FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-cef-off"

    // Other event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_other FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-other"

    // Rate Limit event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_rate_limit FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-rate-limit"

    // No destination event
    FlowExporterIpwriteStatsType_flow_exporter_ipwrite_stats_no_destination FlowExporterIpwriteStatsType = "flow-exporter-ipwrite-stats-no-destination"
)

// FlowMonitorCacheType represents The flow monitor cache type
type FlowMonitorCacheType string

const (
    // Normal Flow monitor cache
    FlowMonitorCacheType_flow_monitor_cache_type_normal FlowMonitorCacheType = "flow-monitor-cache-type-normal"

    // Permanent cache type
    FlowMonitorCacheType_flow_monitor_cache_type_permanent FlowMonitorCacheType = "flow-monitor-cache-type-permanent"

    // Synchronized Flow monitor cache type
    FlowMonitorCacheType_flow_monitor_cache_type_synchronized FlowMonitorCacheType = "flow-monitor-cache-type-synchronized"

    // Immediate Flow monitor cache type
    FlowMonitorCacheType_flow_monitor_cache_type_immediate FlowMonitorCacheType = "flow-monitor-cache-type-immediate"
)

// FlowMonitorCacheState represents Flow monitor cache state
type FlowMonitorCacheState string

const (
    // Flow monitor cache is being deleted
    FlowMonitorCacheState_flow_monitor_cache_state_being_deleted FlowMonitorCacheState = "flow-monitor-cache-state-being-deleted"

    // Flow monitor cache is being allocated
    FlowMonitorCacheState_flow_monitor_cache_state_being_allocated FlowMonitorCacheState = "flow-monitor-cache-state-being-allocated"

    // Flow monitor cache is not allocated
    FlowMonitorCacheState_flow_monitor_cache_state_not_allocated FlowMonitorCacheState = "flow-monitor-cache-state-not-allocated"
)

// FlowMonitors
// All of the flow monitors
type FlowMonitors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Flow monitors. The type is slice of FlowMonitors_FlowMonitor.
    FlowMonitor []FlowMonitors_FlowMonitor

    // List of statistics per exporter. The type is slice of
    // FlowMonitors_FlowExportStatistics.
    FlowExportStatistics []FlowMonitors_FlowExportStatistics

    // List of statistics per flow cache. The type is slice of
    // FlowMonitors_FlowCacheStatistics.
    FlowCacheStatistics []FlowMonitors_FlowCacheStatistics

    // List of statistics per flow monitor. The type is slice of
    // FlowMonitors_FlowMonitorStatistics.
    FlowMonitorStatistics []FlowMonitors_FlowMonitorStatistics
}

func (flowMonitors *FlowMonitors) GetEntityData() *types.CommonEntityData {
    flowMonitors.EntityData.YFilter = flowMonitors.YFilter
    flowMonitors.EntityData.YangName = "flow-monitors"
    flowMonitors.EntityData.BundleName = "cisco_ios_xe"
    flowMonitors.EntityData.ParentYangName = "Cisco-IOS-XE-flow-monitor-oper"
    flowMonitors.EntityData.SegmentPath = "Cisco-IOS-XE-flow-monitor-oper:flow-monitors"
    flowMonitors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowMonitors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowMonitors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowMonitors.EntityData.Children = make(map[string]types.YChild)
    flowMonitors.EntityData.Children["flow-monitor"] = types.YChild{"FlowMonitor", nil}
    for i := range flowMonitors.FlowMonitor {
        flowMonitors.EntityData.Children[types.GetSegmentPath(&flowMonitors.FlowMonitor[i])] = types.YChild{"FlowMonitor", &flowMonitors.FlowMonitor[i]}
    }
    flowMonitors.EntityData.Children["flow-export-statistics"] = types.YChild{"FlowExportStatistics", nil}
    for i := range flowMonitors.FlowExportStatistics {
        flowMonitors.EntityData.Children[types.GetSegmentPath(&flowMonitors.FlowExportStatistics[i])] = types.YChild{"FlowExportStatistics", &flowMonitors.FlowExportStatistics[i]}
    }
    flowMonitors.EntityData.Children["flow-cache-statistics"] = types.YChild{"FlowCacheStatistics", nil}
    for i := range flowMonitors.FlowCacheStatistics {
        flowMonitors.EntityData.Children[types.GetSegmentPath(&flowMonitors.FlowCacheStatistics[i])] = types.YChild{"FlowCacheStatistics", &flowMonitors.FlowCacheStatistics[i]}
    }
    flowMonitors.EntityData.Children["flow-monitor-statistics"] = types.YChild{"FlowMonitorStatistics", nil}
    for i := range flowMonitors.FlowMonitorStatistics {
        flowMonitors.EntityData.Children[types.GetSegmentPath(&flowMonitors.FlowMonitorStatistics[i])] = types.YChild{"FlowMonitorStatistics", &flowMonitors.FlowMonitorStatistics[i]}
    }
    flowMonitors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flowMonitors.EntityData)
}

// FlowMonitors_FlowMonitor
// List of Flow monitors
type FlowMonitors_FlowMonitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the flow monitor. The type is string.
    Name interface{}

    // Time the flow monitor data was collected in seconds. The type is
    // interface{} with range: 0..18446744073709551615.
    TimeCollected interface{}

    // All the flows for this flow monitor.
    Flows FlowMonitors_FlowMonitor_Flows
}

func (flowMonitor *FlowMonitors_FlowMonitor) GetEntityData() *types.CommonEntityData {
    flowMonitor.EntityData.YFilter = flowMonitor.YFilter
    flowMonitor.EntityData.YangName = "flow-monitor"
    flowMonitor.EntityData.BundleName = "cisco_ios_xe"
    flowMonitor.EntityData.ParentYangName = "flow-monitors"
    flowMonitor.EntityData.SegmentPath = "flow-monitor" + "[name='" + fmt.Sprintf("%v", flowMonitor.Name) + "']"
    flowMonitor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowMonitor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowMonitor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowMonitor.EntityData.Children = make(map[string]types.YChild)
    flowMonitor.EntityData.Children["flows"] = types.YChild{"Flows", &flowMonitor.Flows}
    flowMonitor.EntityData.Leafs = make(map[string]types.YLeaf)
    flowMonitor.EntityData.Leafs["name"] = types.YLeaf{"Name", flowMonitor.Name}
    flowMonitor.EntityData.Leafs["time-collected"] = types.YLeaf{"TimeCollected", flowMonitor.TimeCollected}
    return &(flowMonitor.EntityData)
}

// FlowMonitors_FlowMonitor_Flows
// All the flows for this flow monitor
type FlowMonitors_FlowMonitor_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of flows. The type is slice of FlowMonitors_FlowMonitor_Flows_Flow.
    Flow []FlowMonitors_FlowMonitor_Flows_Flow
}

func (flows *FlowMonitors_FlowMonitor_Flows) GetEntityData() *types.CommonEntityData {
    flows.EntityData.YFilter = flows.YFilter
    flows.EntityData.YangName = "flows"
    flows.EntityData.BundleName = "cisco_ios_xe"
    flows.EntityData.ParentYangName = "flow-monitor"
    flows.EntityData.SegmentPath = "flows"
    flows.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flows.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flows.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flows.EntityData.Children = make(map[string]types.YChild)
    flows.EntityData.Children["flow"] = types.YChild{"Flow", nil}
    for i := range flows.Flow {
        flows.EntityData.Children[types.GetSegmentPath(&flows.Flow[i])] = types.YChild{"Flow", &flows.Flow[i]}
    }
    flows.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flows.EntityData)
}

// FlowMonitors_FlowMonitor_Flows_Flow
// List of flows
type FlowMonitors_FlowMonitor_Flows_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Source address of the flow. The type is string.
    SourceAddress interface{}

    // This attribute is a key. Destination address of the flow. The type is
    // string.
    DestinationAddress interface{}

    // This attribute is a key. Input interface of the flow. The type is string.
    InterfaceInput interface{}

    // This attribute is a key. Multicast flow. The type is string.
    IsMulticast interface{}

    // This attribute is a key. VRF ID input. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    VrfIdInput interface{}

    // This attribute is a key. Source port number. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    SourcePort interface{}

    // This attribute is a key. Destination port number. The type is interface{}
    // with range: -9223372036854775808..9223372036854775807.
    DestinationPort interface{}

    // This attribute is a key. ip-tos value. The type is string.
    IpTos interface{}

    // This attribute is a key. IP protocol number. The type is interface{} with
    // range: -9223372036854775808..9223372036854775807.
    IpProtocol interface{}

    // Output interface of the flow. The type is string.
    InterfaceOutput interface{}

    // Number of bytes passed through. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Bytes interface{}

    // Number of packets passed through. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Packets interface{}
}

func (flow *FlowMonitors_FlowMonitor_Flows_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xe"
    flow.EntityData.ParentYangName = "flows"
    flow.EntityData.SegmentPath = "flow" + "[source-address='" + fmt.Sprintf("%v", flow.SourceAddress) + "']" + "[destination-address='" + fmt.Sprintf("%v", flow.DestinationAddress) + "']" + "[interface-input='" + fmt.Sprintf("%v", flow.InterfaceInput) + "']" + "[is-multicast='" + fmt.Sprintf("%v", flow.IsMulticast) + "']" + "[vrf-id-input='" + fmt.Sprintf("%v", flow.VrfIdInput) + "']" + "[source-port='" + fmt.Sprintf("%v", flow.SourcePort) + "']" + "[destination-port='" + fmt.Sprintf("%v", flow.DestinationPort) + "']" + "[ip-tos='" + fmt.Sprintf("%v", flow.IpTos) + "']" + "[ip-protocol='" + fmt.Sprintf("%v", flow.IpProtocol) + "']"
    flow.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flow.EntityData.Children = make(map[string]types.YChild)
    flow.EntityData.Leafs = make(map[string]types.YLeaf)
    flow.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", flow.SourceAddress}
    flow.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", flow.DestinationAddress}
    flow.EntityData.Leafs["interface-input"] = types.YLeaf{"InterfaceInput", flow.InterfaceInput}
    flow.EntityData.Leafs["is-multicast"] = types.YLeaf{"IsMulticast", flow.IsMulticast}
    flow.EntityData.Leafs["vrf-id-input"] = types.YLeaf{"VrfIdInput", flow.VrfIdInput}
    flow.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", flow.SourcePort}
    flow.EntityData.Leafs["destination-port"] = types.YLeaf{"DestinationPort", flow.DestinationPort}
    flow.EntityData.Leafs["ip-tos"] = types.YLeaf{"IpTos", flow.IpTos}
    flow.EntityData.Leafs["ip-protocol"] = types.YLeaf{"IpProtocol", flow.IpProtocol}
    flow.EntityData.Leafs["interface-output"] = types.YLeaf{"InterfaceOutput", flow.InterfaceOutput}
    flow.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", flow.Bytes}
    flow.EntityData.Leafs["packets"] = types.YLeaf{"Packets", flow.Packets}
    return &(flow.EntityData)
}

// FlowMonitors_FlowExportStatistics
// List of statistics per exporter
type FlowMonitors_FlowExportStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the flow exporter. The type is string.
    Name interface{}

    // The coontainer for the transport statistics.
    TransportStats FlowMonitors_FlowExportStatistics_TransportStats

    // The container for the export client information. The type is slice of
    // FlowMonitors_FlowExportStatistics_ExportClient.
    ExportClient []FlowMonitors_FlowExportStatistics_ExportClient
}

func (flowExportStatistics *FlowMonitors_FlowExportStatistics) GetEntityData() *types.CommonEntityData {
    flowExportStatistics.EntityData.YFilter = flowExportStatistics.YFilter
    flowExportStatistics.EntityData.YangName = "flow-export-statistics"
    flowExportStatistics.EntityData.BundleName = "cisco_ios_xe"
    flowExportStatistics.EntityData.ParentYangName = "flow-monitors"
    flowExportStatistics.EntityData.SegmentPath = "flow-export-statistics" + "[name='" + fmt.Sprintf("%v", flowExportStatistics.Name) + "']"
    flowExportStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowExportStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowExportStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowExportStatistics.EntityData.Children = make(map[string]types.YChild)
    flowExportStatistics.EntityData.Children["transport-stats"] = types.YChild{"TransportStats", &flowExportStatistics.TransportStats}
    flowExportStatistics.EntityData.Children["export-client"] = types.YChild{"ExportClient", nil}
    for i := range flowExportStatistics.ExportClient {
        flowExportStatistics.EntityData.Children[types.GetSegmentPath(&flowExportStatistics.ExportClient[i])] = types.YChild{"ExportClient", &flowExportStatistics.ExportClient[i]}
    }
    flowExportStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    flowExportStatistics.EntityData.Leafs["name"] = types.YLeaf{"Name", flowExportStatistics.Name}
    return &(flowExportStatistics.EntityData)
}

// FlowMonitors_FlowExportStatistics_TransportStats
// The coontainer for the transport statistics
type FlowMonitors_FlowExportStatistics_TransportStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time when the statistics were last cleared. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    LastCleared interface{}

    // Container of the exporter statistics. The type is slice of
    // FlowMonitors_FlowExportStatistics_TransportStats_FlowExporterStats.
    FlowExporterStats []FlowMonitors_FlowExportStatistics_TransportStats_FlowExporterStats
}

func (transportStats *FlowMonitors_FlowExportStatistics_TransportStats) GetEntityData() *types.CommonEntityData {
    transportStats.EntityData.YFilter = transportStats.YFilter
    transportStats.EntityData.YangName = "transport-stats"
    transportStats.EntityData.BundleName = "cisco_ios_xe"
    transportStats.EntityData.ParentYangName = "flow-export-statistics"
    transportStats.EntityData.SegmentPath = "transport-stats"
    transportStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transportStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transportStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transportStats.EntityData.Children = make(map[string]types.YChild)
    transportStats.EntityData.Children["flow-exporter-stats"] = types.YChild{"FlowExporterStats", nil}
    for i := range transportStats.FlowExporterStats {
        transportStats.EntityData.Children[types.GetSegmentPath(&transportStats.FlowExporterStats[i])] = types.YChild{"FlowExporterStats", &transportStats.FlowExporterStats[i]}
    }
    transportStats.EntityData.Leafs = make(map[string]types.YLeaf)
    transportStats.EntityData.Leafs["last-cleared"] = types.YLeaf{"LastCleared", transportStats.LastCleared}
    return &(transportStats.EntityData)
}

// FlowMonitors_FlowExportStatistics_TransportStats_FlowExporterStats
// Container of the exporter statistics
type FlowMonitors_FlowExportStatistics_TransportStats_FlowExporterStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of the export statistics. The type is
    // FlowExporterIpwriteStatsType.
    Type_ interface{}

    // The packet counts that have been exported. The type is interface{} with
    // range: 0..18446744073709551615.
    PktCounts interface{}

    // The byte counts that have been exported. The type is interface{} with
    // range: 0..18446744073709551615.
    ByteCounts interface{}
}

func (flowExporterStats *FlowMonitors_FlowExportStatistics_TransportStats_FlowExporterStats) GetEntityData() *types.CommonEntityData {
    flowExporterStats.EntityData.YFilter = flowExporterStats.YFilter
    flowExporterStats.EntityData.YangName = "flow-exporter-stats"
    flowExporterStats.EntityData.BundleName = "cisco_ios_xe"
    flowExporterStats.EntityData.ParentYangName = "transport-stats"
    flowExporterStats.EntityData.SegmentPath = "flow-exporter-stats"
    flowExporterStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowExporterStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowExporterStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowExporterStats.EntityData.Children = make(map[string]types.YChild)
    flowExporterStats.EntityData.Leafs = make(map[string]types.YLeaf)
    flowExporterStats.EntityData.Leafs["type"] = types.YLeaf{"Type_", flowExporterStats.Type_}
    flowExporterStats.EntityData.Leafs["pkt-counts"] = types.YLeaf{"PktCounts", flowExporterStats.PktCounts}
    flowExporterStats.EntityData.Leafs["byte-counts"] = types.YLeaf{"ByteCounts", flowExporterStats.ByteCounts}
    return &(flowExporterStats.EntityData)
}

// FlowMonitors_FlowExportStatistics_ExportClient
// The container for the export client information
type FlowMonitors_FlowExportStatistics_ExportClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the flow export client. The type is string.
    Name interface{}

    // The group that this exporter client belongs to. The type is string.
    Group interface{}

    // The container with the protocol statistics.
    ProtocolStats FlowMonitors_FlowExportStatistics_ExportClient_ProtocolStats
}

func (exportClient *FlowMonitors_FlowExportStatistics_ExportClient) GetEntityData() *types.CommonEntityData {
    exportClient.EntityData.YFilter = exportClient.YFilter
    exportClient.EntityData.YangName = "export-client"
    exportClient.EntityData.BundleName = "cisco_ios_xe"
    exportClient.EntityData.ParentYangName = "flow-export-statistics"
    exportClient.EntityData.SegmentPath = "export-client"
    exportClient.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    exportClient.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    exportClient.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    exportClient.EntityData.Children = make(map[string]types.YChild)
    exportClient.EntityData.Children["protocol-stats"] = types.YChild{"ProtocolStats", &exportClient.ProtocolStats}
    exportClient.EntityData.Leafs = make(map[string]types.YLeaf)
    exportClient.EntityData.Leafs["name"] = types.YLeaf{"Name", exportClient.Name}
    exportClient.EntityData.Leafs["group"] = types.YLeaf{"Group", exportClient.Group}
    return &(exportClient.EntityData)
}

// FlowMonitors_FlowExportStatistics_ExportClient_ProtocolStats
// The container with the protocol statistics
type FlowMonitors_FlowExportStatistics_ExportClient_ProtocolStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of byte added to the exporter. The type is interface{} with range:
    // 0..18446744073709551615.
    BytesAdded interface{}

    // Bytes sent on this exporter. The type is interface{} with range:
    // 0..18446744073709551615.
    BytesSent interface{}

    // Bytes dropped . The type is interface{} with range:
    // 0..18446744073709551615.
    BytesDropped interface{}

    // Number of records added. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsAdded interface{}

    // Number of records sent. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsSent interface{}

    // Number of records dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    RecordsDropped interface{}
}

func (protocolStats *FlowMonitors_FlowExportStatistics_ExportClient_ProtocolStats) GetEntityData() *types.CommonEntityData {
    protocolStats.EntityData.YFilter = protocolStats.YFilter
    protocolStats.EntityData.YangName = "protocol-stats"
    protocolStats.EntityData.BundleName = "cisco_ios_xe"
    protocolStats.EntityData.ParentYangName = "export-client"
    protocolStats.EntityData.SegmentPath = "protocol-stats"
    protocolStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protocolStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protocolStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protocolStats.EntityData.Children = make(map[string]types.YChild)
    protocolStats.EntityData.Leafs = make(map[string]types.YLeaf)
    protocolStats.EntityData.Leafs["bytes-added"] = types.YLeaf{"BytesAdded", protocolStats.BytesAdded}
    protocolStats.EntityData.Leafs["bytes-sent"] = types.YLeaf{"BytesSent", protocolStats.BytesSent}
    protocolStats.EntityData.Leafs["bytes-dropped"] = types.YLeaf{"BytesDropped", protocolStats.BytesDropped}
    protocolStats.EntityData.Leafs["records-added"] = types.YLeaf{"RecordsAdded", protocolStats.RecordsAdded}
    protocolStats.EntityData.Leafs["records-sent"] = types.YLeaf{"RecordsSent", protocolStats.RecordsSent}
    protocolStats.EntityData.Leafs["records-dropped"] = types.YLeaf{"RecordsDropped", protocolStats.RecordsDropped}
    return &(protocolStats.EntityData)
}

// FlowMonitors_FlowCacheStatistics
// List of statistics per flow cache
type FlowMonitors_FlowCacheStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the flow cache. The type is string.
    Name interface{}

    // The size of the cache. The type is interface{} with range:
    // 0..18446744073709551615.
    CacheSize interface{}

    // The current number of entries. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentEntries interface{}

    // The high watermark of flows. The type is interface{} with range:
    // 0..18446744073709551615.
    HighWatermark interface{}

    // The number of flows added. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowsAdded interface{}

    // The number of flows that have been aged. The type is interface{} with
    // range: 0..18446744073709551615.
    FlowsAged interface{}

    // The number of flows that have been timed out  whilst still active. The type
    // is interface{} with range: 0..18446744073709551615.
    ActiveFlowsTimedOut interface{}

    // The number of flows that have been timed out for inactivity. The type is
    // interface{} with range: 0..18446744073709551615.
    InactiveFlowsTimedOut interface{}
}

func (flowCacheStatistics *FlowMonitors_FlowCacheStatistics) GetEntityData() *types.CommonEntityData {
    flowCacheStatistics.EntityData.YFilter = flowCacheStatistics.YFilter
    flowCacheStatistics.EntityData.YangName = "flow-cache-statistics"
    flowCacheStatistics.EntityData.BundleName = "cisco_ios_xe"
    flowCacheStatistics.EntityData.ParentYangName = "flow-monitors"
    flowCacheStatistics.EntityData.SegmentPath = "flow-cache-statistics" + "[name='" + fmt.Sprintf("%v", flowCacheStatistics.Name) + "']"
    flowCacheStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowCacheStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowCacheStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowCacheStatistics.EntityData.Children = make(map[string]types.YChild)
    flowCacheStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    flowCacheStatistics.EntityData.Leafs["name"] = types.YLeaf{"Name", flowCacheStatistics.Name}
    flowCacheStatistics.EntityData.Leafs["cache-size"] = types.YLeaf{"CacheSize", flowCacheStatistics.CacheSize}
    flowCacheStatistics.EntityData.Leafs["current-entries"] = types.YLeaf{"CurrentEntries", flowCacheStatistics.CurrentEntries}
    flowCacheStatistics.EntityData.Leafs["high-watermark"] = types.YLeaf{"HighWatermark", flowCacheStatistics.HighWatermark}
    flowCacheStatistics.EntityData.Leafs["flows-added"] = types.YLeaf{"FlowsAdded", flowCacheStatistics.FlowsAdded}
    flowCacheStatistics.EntityData.Leafs["flows-aged"] = types.YLeaf{"FlowsAged", flowCacheStatistics.FlowsAged}
    flowCacheStatistics.EntityData.Leafs["active-flows-timed-out"] = types.YLeaf{"ActiveFlowsTimedOut", flowCacheStatistics.ActiveFlowsTimedOut}
    flowCacheStatistics.EntityData.Leafs["inactive-flows-timed-out"] = types.YLeaf{"InactiveFlowsTimedOut", flowCacheStatistics.InactiveFlowsTimedOut}
    return &(flowCacheStatistics.EntityData)
}

// FlowMonitors_FlowMonitorStatistics
// List of statistics per flow monitor
type FlowMonitors_FlowMonitorStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the flow monitor. The type is string.
    MonitorName interface{}

    // The description of the flow monitor. The type is string.
    Description interface{}

    // The name of the record. The type is string.
    RecordName interface{}

    // The active flow exporters. The type is slice of string.
    ActiveFlowExporter []interface{}

    // The inactive flow exporters. The type is slice of string.
    InactiveFlowExporter []interface{}

    // The number of invalid packet counts. The type is interface{} with range:
    // 0..18446744073709551615.
    InvalidPacketCounts interface{}

    // Indicate whether the transaction end ager is enabled. The type is bool.
    TransactionEndAgerEnabled interface{}

    // The protocol distribution is configured. The type is string.
    ProtocolDistConfigured interface{}

    // The size distribution is configured. The type is string.
    SizeDistConfigured interface{}

    // The inactive timer on the normal cache. The type is interface{} with range:
    // 0..4294967295.
    InactiveTimer interface{}

    // The active time on the normal cache. The type is interface{} with range:
    // 0..4294967295.
    ActiveTimer interface{}

    // The update timeout of the permanent type. The type is interface{} with
    // range: 0..4294967295.
    UpdateTimeout interface{}

    // The timeout of the synchronized cache. The type is interface{} with range:
    // 0..4294967295.
    SynchronizedTimeout interface{}

    // The export spread interval. The type is interface{} with range:
    // 0..4294967295.
    ExportSpreadInterval interface{}

    // The timeout for the immediate cache. The type is interface{} with range:
    // 0..4294967295.
    ImmediateTimeout interface{}

    // The grouping of the cache data.
    CacheData FlowMonitors_FlowMonitorStatistics_CacheData
}

func (flowMonitorStatistics *FlowMonitors_FlowMonitorStatistics) GetEntityData() *types.CommonEntityData {
    flowMonitorStatistics.EntityData.YFilter = flowMonitorStatistics.YFilter
    flowMonitorStatistics.EntityData.YangName = "flow-monitor-statistics"
    flowMonitorStatistics.EntityData.BundleName = "cisco_ios_xe"
    flowMonitorStatistics.EntityData.ParentYangName = "flow-monitors"
    flowMonitorStatistics.EntityData.SegmentPath = "flow-monitor-statistics" + "[monitor-name='" + fmt.Sprintf("%v", flowMonitorStatistics.MonitorName) + "']"
    flowMonitorStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowMonitorStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowMonitorStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowMonitorStatistics.EntityData.Children = make(map[string]types.YChild)
    flowMonitorStatistics.EntityData.Children["cache-data"] = types.YChild{"CacheData", &flowMonitorStatistics.CacheData}
    flowMonitorStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    flowMonitorStatistics.EntityData.Leafs["monitor-name"] = types.YLeaf{"MonitorName", flowMonitorStatistics.MonitorName}
    flowMonitorStatistics.EntityData.Leafs["description"] = types.YLeaf{"Description", flowMonitorStatistics.Description}
    flowMonitorStatistics.EntityData.Leafs["record-name"] = types.YLeaf{"RecordName", flowMonitorStatistics.RecordName}
    flowMonitorStatistics.EntityData.Leafs["active-flow-exporter"] = types.YLeaf{"ActiveFlowExporter", flowMonitorStatistics.ActiveFlowExporter}
    flowMonitorStatistics.EntityData.Leafs["inactive-flow-exporter"] = types.YLeaf{"InactiveFlowExporter", flowMonitorStatistics.InactiveFlowExporter}
    flowMonitorStatistics.EntityData.Leafs["invalid-packet-counts"] = types.YLeaf{"InvalidPacketCounts", flowMonitorStatistics.InvalidPacketCounts}
    flowMonitorStatistics.EntityData.Leafs["transaction-end-ager-enabled"] = types.YLeaf{"TransactionEndAgerEnabled", flowMonitorStatistics.TransactionEndAgerEnabled}
    flowMonitorStatistics.EntityData.Leafs["protocol-dist-configured"] = types.YLeaf{"ProtocolDistConfigured", flowMonitorStatistics.ProtocolDistConfigured}
    flowMonitorStatistics.EntityData.Leafs["size-dist-configured"] = types.YLeaf{"SizeDistConfigured", flowMonitorStatistics.SizeDistConfigured}
    flowMonitorStatistics.EntityData.Leafs["inactive-timer"] = types.YLeaf{"InactiveTimer", flowMonitorStatistics.InactiveTimer}
    flowMonitorStatistics.EntityData.Leafs["active-timer"] = types.YLeaf{"ActiveTimer", flowMonitorStatistics.ActiveTimer}
    flowMonitorStatistics.EntityData.Leafs["update-timeout"] = types.YLeaf{"UpdateTimeout", flowMonitorStatistics.UpdateTimeout}
    flowMonitorStatistics.EntityData.Leafs["synchronized-timeout"] = types.YLeaf{"SynchronizedTimeout", flowMonitorStatistics.SynchronizedTimeout}
    flowMonitorStatistics.EntityData.Leafs["export-spread-interval"] = types.YLeaf{"ExportSpreadInterval", flowMonitorStatistics.ExportSpreadInterval}
    flowMonitorStatistics.EntityData.Leafs["immediate-timeout"] = types.YLeaf{"ImmediateTimeout", flowMonitorStatistics.ImmediateTimeout}
    return &(flowMonitorStatistics.EntityData)
}

// FlowMonitors_FlowMonitorStatistics_CacheData
// The grouping of the cache data
type FlowMonitors_FlowMonitorStatistics_CacheData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The state of the flow cache. The type is FlowMonitorCacheState.
    State interface{}

    // The type of the flow cache. The type is string.
    Type_ interface{}

    // The name of the cache. The type is string.
    CacheName interface{}

    // The status of the cache. The type is string.
    Status interface{}

    // The number of entries permissible in the cache. The type is interface{}
    // with range: 0..18446744073709551615.
    NumEntries interface{}

    // The number of bytes in the cache. The type is interface{} with range:
    // 0..18446744073709551615.
    NumBytes interface{}
}

func (cacheData *FlowMonitors_FlowMonitorStatistics_CacheData) GetEntityData() *types.CommonEntityData {
    cacheData.EntityData.YFilter = cacheData.YFilter
    cacheData.EntityData.YangName = "cache-data"
    cacheData.EntityData.BundleName = "cisco_ios_xe"
    cacheData.EntityData.ParentYangName = "flow-monitor-statistics"
    cacheData.EntityData.SegmentPath = "cache-data"
    cacheData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cacheData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cacheData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cacheData.EntityData.Children = make(map[string]types.YChild)
    cacheData.EntityData.Leafs = make(map[string]types.YLeaf)
    cacheData.EntityData.Leafs["state"] = types.YLeaf{"State", cacheData.State}
    cacheData.EntityData.Leafs["type"] = types.YLeaf{"Type_", cacheData.Type_}
    cacheData.EntityData.Leafs["cache-name"] = types.YLeaf{"CacheName", cacheData.CacheName}
    cacheData.EntityData.Leafs["status"] = types.YLeaf{"Status", cacheData.Status}
    cacheData.EntityData.Leafs["num-entries"] = types.YLeaf{"NumEntries", cacheData.NumEntries}
    cacheData.EntityData.Leafs["num-bytes"] = types.YLeaf{"NumBytes", cacheData.NumBytes}
    return &(cacheData.EntityData)
}

