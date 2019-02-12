// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-tc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   traffic-collector: Global Traffic Collector configuration
//     commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_tc_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_tc_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-tc-oper traffic-collector}", reflect.TypeOf(TrafficCollector{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-tc-oper:traffic-collector", reflect.TypeOf(TrafficCollector{}))
}

// TcOperAfName represents Tc oper af name
type TcOperAfName string

const (
    // IPv4
    TcOperAfName_ipv4 TcOperAfName = "ipv4"

    // IPv6
    TcOperAfName_ipv6 TcOperAfName = "ipv6"
)

// TrafficCollector
// Global Traffic Collector configuration commands
type TrafficCollector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // External Interface.
    ExternalInterfaces TrafficCollector_ExternalInterfaces

    // Traffic Collector summary.
    Summary TrafficCollector_Summary

    // VRF specific operational data.
    VrfTable TrafficCollector_VrfTable

    // Address Family specific operational data.
    Afs TrafficCollector_Afs
}

func (trafficCollector *TrafficCollector) GetEntityData() *types.CommonEntityData {
    trafficCollector.EntityData.YFilter = trafficCollector.YFilter
    trafficCollector.EntityData.YangName = "traffic-collector"
    trafficCollector.EntityData.BundleName = "cisco_ios_xr"
    trafficCollector.EntityData.ParentYangName = "Cisco-IOS-XR-infra-tc-oper"
    trafficCollector.EntityData.SegmentPath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector"
    trafficCollector.EntityData.AbsolutePath = trafficCollector.EntityData.SegmentPath
    trafficCollector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficCollector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficCollector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficCollector.EntityData.Children = types.NewOrderedMap()
    trafficCollector.EntityData.Children.Append("external-interfaces", types.YChild{"ExternalInterfaces", &trafficCollector.ExternalInterfaces})
    trafficCollector.EntityData.Children.Append("summary", types.YChild{"Summary", &trafficCollector.Summary})
    trafficCollector.EntityData.Children.Append("vrf-table", types.YChild{"VrfTable", &trafficCollector.VrfTable})
    trafficCollector.EntityData.Children.Append("afs", types.YChild{"Afs", &trafficCollector.Afs})
    trafficCollector.EntityData.Leafs = types.NewOrderedMap()

    trafficCollector.EntityData.YListKeys = []string {}

    return &(trafficCollector.EntityData)
}

// TrafficCollector_ExternalInterfaces
// External Interface
type TrafficCollector_ExternalInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // External Interface . The type is slice of
    // TrafficCollector_ExternalInterfaces_ExternalInterface.
    ExternalInterface []*TrafficCollector_ExternalInterfaces_ExternalInterface
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetEntityData() *types.CommonEntityData {
    externalInterfaces.EntityData.YFilter = externalInterfaces.YFilter
    externalInterfaces.EntityData.YangName = "external-interfaces"
    externalInterfaces.EntityData.BundleName = "cisco_ios_xr"
    externalInterfaces.EntityData.ParentYangName = "traffic-collector"
    externalInterfaces.EntityData.SegmentPath = "external-interfaces"
    externalInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/" + externalInterfaces.EntityData.SegmentPath
    externalInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalInterfaces.EntityData.Children = types.NewOrderedMap()
    externalInterfaces.EntityData.Children.Append("external-interface", types.YChild{"ExternalInterface", nil})
    for i := range externalInterfaces.ExternalInterface {
        externalInterfaces.EntityData.Children.Append(types.GetSegmentPath(externalInterfaces.ExternalInterface[i]), types.YChild{"ExternalInterface", externalInterfaces.ExternalInterface[i]})
    }
    externalInterfaces.EntityData.Leafs = types.NewOrderedMap()

    externalInterfaces.EntityData.YListKeys = []string {}

    return &(externalInterfaces.EntityData)
}

// TrafficCollector_ExternalInterfaces_ExternalInterface
// External Interface 
type TrafficCollector_ExternalInterfaces_ExternalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Flag to indicate interface enabled or not. The type is bool.
    IsInterfaceEnabled interface{}
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetEntityData() *types.CommonEntityData {
    externalInterface.EntityData.YFilter = externalInterface.YFilter
    externalInterface.EntityData.YangName = "external-interface"
    externalInterface.EntityData.BundleName = "cisco_ios_xr"
    externalInterface.EntityData.ParentYangName = "external-interfaces"
    externalInterface.EntityData.SegmentPath = "external-interface" + types.AddKeyToken(externalInterface.InterfaceName, "interface-name")
    externalInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/external-interfaces/" + externalInterface.EntityData.SegmentPath
    externalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalInterface.EntityData.Children = types.NewOrderedMap()
    externalInterface.EntityData.Leafs = types.NewOrderedMap()
    externalInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", externalInterface.InterfaceName})
    externalInterface.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", externalInterface.InterfaceNameXr})
    externalInterface.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", externalInterface.InterfaceHandle})
    externalInterface.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", externalInterface.Vrfid})
    externalInterface.EntityData.Leafs.Append("is-interface-enabled", types.YLeaf{"IsInterfaceEnabled", externalInterface.IsInterfaceEnabled})

    externalInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(externalInterface.EntityData)
}

// TrafficCollector_Summary
// Traffic Collector summary
type TrafficCollector_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistic collection interval in minutes. The type is interface{} with
    // range: 0..255. Units are minute.
    CollectionInterval interface{}

    // TRUE if collection timer is running. The type is bool.
    CollectionTimerIsRunning interface{}

    // Statistic history timeout interval in hours. The type is interface{} with
    // range: 0..65535. Units are hour.
    TimeoutInterval interface{}

    // TRUE if history timeout timer is running. The type is bool.
    TimeoutTimerIsRunning interface{}

    // Statistics history size. The type is interface{} with range: 0..255.
    HistorySize interface{}

    // Database statistics for External Interface.
    DatabaseStatisticsExternalInterface TrafficCollector_Summary_DatabaseStatisticsExternalInterface

    // VRF table statistics. The type is slice of
    // TrafficCollector_Summary_VrfStatistic.
    VrfStatistic []*TrafficCollector_Summary_VrfStatistic

    // Statistics per message type for STAT collector. The type is slice of
    // TrafficCollector_Summary_CollectionMessageStatistic.
    CollectionMessageStatistic []*TrafficCollector_Summary_CollectionMessageStatistic

    // Statistics per message type for Chkpt. The type is slice of
    // TrafficCollector_Summary_CheckpointMessageStatistic.
    CheckpointMessageStatistic []*TrafficCollector_Summary_CheckpointMessageStatistic
}

func (summary *TrafficCollector_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "traffic-collector"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("database-statistics-external-interface", types.YChild{"DatabaseStatisticsExternalInterface", &summary.DatabaseStatisticsExternalInterface})
    summary.EntityData.Children.Append("vrf-statistic", types.YChild{"VrfStatistic", nil})
    for i := range summary.VrfStatistic {
        types.SetYListKey(summary.VrfStatistic[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.VrfStatistic[i]), types.YChild{"VrfStatistic", summary.VrfStatistic[i]})
    }
    summary.EntityData.Children.Append("collection-message-statistic", types.YChild{"CollectionMessageStatistic", nil})
    for i := range summary.CollectionMessageStatistic {
        types.SetYListKey(summary.CollectionMessageStatistic[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.CollectionMessageStatistic[i]), types.YChild{"CollectionMessageStatistic", summary.CollectionMessageStatistic[i]})
    }
    summary.EntityData.Children.Append("checkpoint-message-statistic", types.YChild{"CheckpointMessageStatistic", nil})
    for i := range summary.CheckpointMessageStatistic {
        types.SetYListKey(summary.CheckpointMessageStatistic[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.CheckpointMessageStatistic[i]), types.YChild{"CheckpointMessageStatistic", summary.CheckpointMessageStatistic[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("collection-interval", types.YLeaf{"CollectionInterval", summary.CollectionInterval})
    summary.EntityData.Leafs.Append("collection-timer-is-running", types.YLeaf{"CollectionTimerIsRunning", summary.CollectionTimerIsRunning})
    summary.EntityData.Leafs.Append("timeout-interval", types.YLeaf{"TimeoutInterval", summary.TimeoutInterval})
    summary.EntityData.Leafs.Append("timeout-timer-is-running", types.YLeaf{"TimeoutTimerIsRunning", summary.TimeoutTimerIsRunning})
    summary.EntityData.Leafs.Append("history-size", types.YLeaf{"HistorySize", summary.HistorySize})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// TrafficCollector_Summary_DatabaseStatisticsExternalInterface
// Database statistics for External Interface
type TrafficCollector_Summary_DatabaseStatisticsExternalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetEntityData() *types.CommonEntityData {
    databaseStatisticsExternalInterface.EntityData.YFilter = databaseStatisticsExternalInterface.YFilter
    databaseStatisticsExternalInterface.EntityData.YangName = "database-statistics-external-interface"
    databaseStatisticsExternalInterface.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsExternalInterface.EntityData.ParentYangName = "summary"
    databaseStatisticsExternalInterface.EntityData.SegmentPath = "database-statistics-external-interface"
    databaseStatisticsExternalInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/" + databaseStatisticsExternalInterface.EntityData.SegmentPath
    databaseStatisticsExternalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsExternalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsExternalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsExternalInterface.EntityData.Children = types.NewOrderedMap()
    databaseStatisticsExternalInterface.EntityData.Leafs = types.NewOrderedMap()
    databaseStatisticsExternalInterface.EntityData.Leafs.Append("number-of-entries", types.YLeaf{"NumberOfEntries", databaseStatisticsExternalInterface.NumberOfEntries})
    databaseStatisticsExternalInterface.EntityData.Leafs.Append("number-of-stale-entries", types.YLeaf{"NumberOfStaleEntries", databaseStatisticsExternalInterface.NumberOfStaleEntries})
    databaseStatisticsExternalInterface.EntityData.Leafs.Append("number-of-add-o-perations", types.YLeaf{"NumberOfAddOPerations", databaseStatisticsExternalInterface.NumberOfAddOPerations})
    databaseStatisticsExternalInterface.EntityData.Leafs.Append("number-of-delete-o-perations", types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsExternalInterface.NumberOfDeleteOPerations})

    databaseStatisticsExternalInterface.EntityData.YListKeys = []string {}

    return &(databaseStatisticsExternalInterface.EntityData)
}

// TrafficCollector_Summary_VrfStatistic
// VRF table statistics
type TrafficCollector_Summary_VrfStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // VRF name. The type is string.
    VrfName interface{}

    // Database statistics for IPv4 table.
    DatabaseStatisticsIpv4 TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4

    // Database statistics for Tunnel table.
    DatabaseStatisticsTunnel TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetEntityData() *types.CommonEntityData {
    vrfStatistic.EntityData.YFilter = vrfStatistic.YFilter
    vrfStatistic.EntityData.YangName = "vrf-statistic"
    vrfStatistic.EntityData.BundleName = "cisco_ios_xr"
    vrfStatistic.EntityData.ParentYangName = "summary"
    vrfStatistic.EntityData.SegmentPath = "vrf-statistic" + types.AddNoKeyToken(vrfStatistic)
    vrfStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/" + vrfStatistic.EntityData.SegmentPath
    vrfStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfStatistic.EntityData.Children = types.NewOrderedMap()
    vrfStatistic.EntityData.Children.Append("database-statistics-ipv4", types.YChild{"DatabaseStatisticsIpv4", &vrfStatistic.DatabaseStatisticsIpv4})
    vrfStatistic.EntityData.Children.Append("database-statistics-tunnel", types.YChild{"DatabaseStatisticsTunnel", &vrfStatistic.DatabaseStatisticsTunnel})
    vrfStatistic.EntityData.Leafs = types.NewOrderedMap()
    vrfStatistic.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfStatistic.VrfName})

    vrfStatistic.EntityData.YListKeys = []string {}

    return &(vrfStatistic.EntityData)
}

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4
// Database statistics for IPv4 table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetEntityData() *types.CommonEntityData {
    databaseStatisticsIpv4.EntityData.YFilter = databaseStatisticsIpv4.YFilter
    databaseStatisticsIpv4.EntityData.YangName = "database-statistics-ipv4"
    databaseStatisticsIpv4.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsIpv4.EntityData.ParentYangName = "vrf-statistic"
    databaseStatisticsIpv4.EntityData.SegmentPath = "database-statistics-ipv4"
    databaseStatisticsIpv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/vrf-statistic/" + databaseStatisticsIpv4.EntityData.SegmentPath
    databaseStatisticsIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsIpv4.EntityData.Children = types.NewOrderedMap()
    databaseStatisticsIpv4.EntityData.Leafs = types.NewOrderedMap()
    databaseStatisticsIpv4.EntityData.Leafs.Append("number-of-entries", types.YLeaf{"NumberOfEntries", databaseStatisticsIpv4.NumberOfEntries})
    databaseStatisticsIpv4.EntityData.Leafs.Append("number-of-stale-entries", types.YLeaf{"NumberOfStaleEntries", databaseStatisticsIpv4.NumberOfStaleEntries})
    databaseStatisticsIpv4.EntityData.Leafs.Append("number-of-add-o-perations", types.YLeaf{"NumberOfAddOPerations", databaseStatisticsIpv4.NumberOfAddOPerations})
    databaseStatisticsIpv4.EntityData.Leafs.Append("number-of-delete-o-perations", types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsIpv4.NumberOfDeleteOPerations})

    databaseStatisticsIpv4.EntityData.YListKeys = []string {}

    return &(databaseStatisticsIpv4.EntityData)
}

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
// Database statistics for Tunnel table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of DB entries. The type is interface{} with range: 0..4294967295.
    NumberOfEntries interface{}

    // Number of stale  entries. The type is interface{} with range:
    // 0..4294967295.
    NumberOfStaleEntries interface{}

    // Number of add operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfAddOPerations interface{}

    // Number of delete operations. The type is interface{} with range:
    // 0..18446744073709551615.
    NumberOfDeleteOPerations interface{}
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetEntityData() *types.CommonEntityData {
    databaseStatisticsTunnel.EntityData.YFilter = databaseStatisticsTunnel.YFilter
    databaseStatisticsTunnel.EntityData.YangName = "database-statistics-tunnel"
    databaseStatisticsTunnel.EntityData.BundleName = "cisco_ios_xr"
    databaseStatisticsTunnel.EntityData.ParentYangName = "vrf-statistic"
    databaseStatisticsTunnel.EntityData.SegmentPath = "database-statistics-tunnel"
    databaseStatisticsTunnel.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/vrf-statistic/" + databaseStatisticsTunnel.EntityData.SegmentPath
    databaseStatisticsTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseStatisticsTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseStatisticsTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseStatisticsTunnel.EntityData.Children = types.NewOrderedMap()
    databaseStatisticsTunnel.EntityData.Leafs = types.NewOrderedMap()
    databaseStatisticsTunnel.EntityData.Leafs.Append("number-of-entries", types.YLeaf{"NumberOfEntries", databaseStatisticsTunnel.NumberOfEntries})
    databaseStatisticsTunnel.EntityData.Leafs.Append("number-of-stale-entries", types.YLeaf{"NumberOfStaleEntries", databaseStatisticsTunnel.NumberOfStaleEntries})
    databaseStatisticsTunnel.EntityData.Leafs.Append("number-of-add-o-perations", types.YLeaf{"NumberOfAddOPerations", databaseStatisticsTunnel.NumberOfAddOPerations})
    databaseStatisticsTunnel.EntityData.Leafs.Append("number-of-delete-o-perations", types.YLeaf{"NumberOfDeleteOPerations", databaseStatisticsTunnel.NumberOfDeleteOPerations})

    databaseStatisticsTunnel.EntityData.YListKeys = []string {}

    return &(databaseStatisticsTunnel.EntityData)
}

// TrafficCollector_Summary_CollectionMessageStatistic
// Statistics per message type for STAT collector
type TrafficCollector_Summary_CollectionMessageStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketSent interface{}

    // Number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteSent interface{}

    // Number of packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketReceived interface{}

    // Number of bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteReceived interface{}

    // Maximum roundtrip latency in msec. The type is interface{} with range:
    // 0..4294967295.
    MaximumRoundtripLatency interface{}

    // Timestamp of maximum latency. The type is interface{} with range:
    // 0..18446744073709551615.
    MaimumLatencyTimestamp interface{}
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetEntityData() *types.CommonEntityData {
    collectionMessageStatistic.EntityData.YFilter = collectionMessageStatistic.YFilter
    collectionMessageStatistic.EntityData.YangName = "collection-message-statistic"
    collectionMessageStatistic.EntityData.BundleName = "cisco_ios_xr"
    collectionMessageStatistic.EntityData.ParentYangName = "summary"
    collectionMessageStatistic.EntityData.SegmentPath = "collection-message-statistic" + types.AddNoKeyToken(collectionMessageStatistic)
    collectionMessageStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/" + collectionMessageStatistic.EntityData.SegmentPath
    collectionMessageStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    collectionMessageStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    collectionMessageStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    collectionMessageStatistic.EntityData.Children = types.NewOrderedMap()
    collectionMessageStatistic.EntityData.Leafs = types.NewOrderedMap()
    collectionMessageStatistic.EntityData.Leafs.Append("packet-sent", types.YLeaf{"PacketSent", collectionMessageStatistic.PacketSent})
    collectionMessageStatistic.EntityData.Leafs.Append("byte-sent", types.YLeaf{"ByteSent", collectionMessageStatistic.ByteSent})
    collectionMessageStatistic.EntityData.Leafs.Append("packet-received", types.YLeaf{"PacketReceived", collectionMessageStatistic.PacketReceived})
    collectionMessageStatistic.EntityData.Leafs.Append("byte-received", types.YLeaf{"ByteReceived", collectionMessageStatistic.ByteReceived})
    collectionMessageStatistic.EntityData.Leafs.Append("maximum-roundtrip-latency", types.YLeaf{"MaximumRoundtripLatency", collectionMessageStatistic.MaximumRoundtripLatency})
    collectionMessageStatistic.EntityData.Leafs.Append("maimum-latency-timestamp", types.YLeaf{"MaimumLatencyTimestamp", collectionMessageStatistic.MaimumLatencyTimestamp})

    collectionMessageStatistic.EntityData.YListKeys = []string {}

    return &(collectionMessageStatistic.EntityData)
}

// TrafficCollector_Summary_CheckpointMessageStatistic
// Statistics per message type for Chkpt
type TrafficCollector_Summary_CheckpointMessageStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Number of packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketSent interface{}

    // Number of bytes sent. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteSent interface{}

    // Number of packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketReceived interface{}

    // Number of bytes received. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ByteReceived interface{}

    // Maximum roundtrip latency in msec. The type is interface{} with range:
    // 0..4294967295.
    MaximumRoundtripLatency interface{}

    // Timestamp of maximum latency. The type is interface{} with range:
    // 0..18446744073709551615.
    MaimumLatencyTimestamp interface{}
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetEntityData() *types.CommonEntityData {
    checkpointMessageStatistic.EntityData.YFilter = checkpointMessageStatistic.YFilter
    checkpointMessageStatistic.EntityData.YangName = "checkpoint-message-statistic"
    checkpointMessageStatistic.EntityData.BundleName = "cisco_ios_xr"
    checkpointMessageStatistic.EntityData.ParentYangName = "summary"
    checkpointMessageStatistic.EntityData.SegmentPath = "checkpoint-message-statistic" + types.AddNoKeyToken(checkpointMessageStatistic)
    checkpointMessageStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/summary/" + checkpointMessageStatistic.EntityData.SegmentPath
    checkpointMessageStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    checkpointMessageStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    checkpointMessageStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    checkpointMessageStatistic.EntityData.Children = types.NewOrderedMap()
    checkpointMessageStatistic.EntityData.Leafs = types.NewOrderedMap()
    checkpointMessageStatistic.EntityData.Leafs.Append("packet-sent", types.YLeaf{"PacketSent", checkpointMessageStatistic.PacketSent})
    checkpointMessageStatistic.EntityData.Leafs.Append("byte-sent", types.YLeaf{"ByteSent", checkpointMessageStatistic.ByteSent})
    checkpointMessageStatistic.EntityData.Leafs.Append("packet-received", types.YLeaf{"PacketReceived", checkpointMessageStatistic.PacketReceived})
    checkpointMessageStatistic.EntityData.Leafs.Append("byte-received", types.YLeaf{"ByteReceived", checkpointMessageStatistic.ByteReceived})
    checkpointMessageStatistic.EntityData.Leafs.Append("maximum-roundtrip-latency", types.YLeaf{"MaximumRoundtripLatency", checkpointMessageStatistic.MaximumRoundtripLatency})
    checkpointMessageStatistic.EntityData.Leafs.Append("maimum-latency-timestamp", types.YLeaf{"MaimumLatencyTimestamp", checkpointMessageStatistic.MaimumLatencyTimestamp})

    checkpointMessageStatistic.EntityData.YListKeys = []string {}

    return &(checkpointMessageStatistic.EntityData)
}

// TrafficCollector_VrfTable
// VRF specific operational data
type TrafficCollector_VrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DefaultVRF specific operational data.
    DefaultVrf TrafficCollector_VrfTable_DefaultVrf
}

func (vrfTable *TrafficCollector_VrfTable) GetEntityData() *types.CommonEntityData {
    vrfTable.EntityData.YFilter = vrfTable.YFilter
    vrfTable.EntityData.YangName = "vrf-table"
    vrfTable.EntityData.BundleName = "cisco_ios_xr"
    vrfTable.EntityData.ParentYangName = "traffic-collector"
    vrfTable.EntityData.SegmentPath = "vrf-table"
    vrfTable.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/" + vrfTable.EntityData.SegmentPath
    vrfTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfTable.EntityData.Children = types.NewOrderedMap()
    vrfTable.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &vrfTable.DefaultVrf})
    vrfTable.EntityData.Leafs = types.NewOrderedMap()

    vrfTable.EntityData.YListKeys = []string {}

    return &(vrfTable.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf
// DefaultVRF specific operational data
type TrafficCollector_VrfTable_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family specific operational data.
    Afs TrafficCollector_VrfTable_DefaultVrf_Afs
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "vrf-table"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("afs", types.YChild{"Afs", &defaultVrf.Afs})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs
// Address Family specific operational data
type TrafficCollector_VrfTable_DefaultVrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af.
    Af []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "default-vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af
// Operational data for given Address Family
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("counters", types.YChild{"Counters", &af.Counters})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
// Show Counters
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "af"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("prefixes", types.YChild{"Prefixes", &counters.Prefixes})
    counters.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &counters.Tunnels})
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "counters"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/" + prefixes.EntityData.SegmentPath
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = types.NewOrderedMap()
    prefixes.EntityData.Children.Append("prefix", types.YChild{"Prefix", nil})
    for i := range prefixes.Prefix {
        types.SetYListKey(prefixes.Prefix[i], i)
        prefixes.EntityData.Children.Append(types.GetSegmentPath(prefixes.Prefix[i]), types.YChild{"Prefix", prefixes.Prefix[i]})
    }
    prefixes.EntityData.Leafs = types.NewOrderedMap()

    prefixes.EntityData.YListKeys = []string {}

    return &(prefixes.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IP Address. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // SR Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // LDP Label. The type is interface{} with range: 0..4294967295.
    LdpLabel interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + types.AddNoKeyToken(prefix)
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/prefixes/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("base-counter-statistics", types.YChild{"BaseCounterStatistics", &prefix.BaseCounterStatistics})
    prefix.EntityData.Children.Append("traffic-matrix-counter-statistics", types.YChild{"TrafficMatrixCounterStatistics", &prefix.TrafficMatrixCounterStatistics})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("ipaddr", types.YLeaf{"Ipaddr", prefix.Ipaddr})
    prefix.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", prefix.Mask})
    prefix.EntityData.Leafs.Append("label", types.YLeaf{"Label", prefix.Label})
    prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefix.Prefix})
    prefix.EntityData.Leafs.Append("label-xr", types.YLeaf{"LabelXr", prefix.LabelXr})
    prefix.EntityData.Leafs.Append("ldp-label", types.YLeaf{"LdpLabel", prefix.LdpLabel})
    prefix.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", prefix.IsActive})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "prefix"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/prefixes/prefix/" + baseCounterStatistics.EntityData.SegmentPath
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range baseCounterStatistics.CountHistory {
        types.SetYListKey(baseCounterStatistics.CountHistory[i], i)
        baseCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(baseCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", baseCounterStatistics.CountHistory[i]})
    }
    baseCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched})
    baseCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched})

    baseCounterStatistics.EntityData.YListKeys = []string {}

    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/prefixes/prefix/base-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetEntityData() *types.CommonEntityData {
    trafficMatrixCounterStatistics.EntityData.YFilter = trafficMatrixCounterStatistics.YFilter
    trafficMatrixCounterStatistics.EntityData.YangName = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    trafficMatrixCounterStatistics.EntityData.ParentYangName = "prefix"
    trafficMatrixCounterStatistics.EntityData.SegmentPath = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/prefixes/prefix/" + trafficMatrixCounterStatistics.EntityData.SegmentPath
    trafficMatrixCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMatrixCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMatrixCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMatrixCounterStatistics.EntityData.Children = types.NewOrderedMap()
    trafficMatrixCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range trafficMatrixCounterStatistics.CountHistory {
        types.SetYListKey(trafficMatrixCounterStatistics.CountHistory[i], i)
        trafficMatrixCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(trafficMatrixCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", trafficMatrixCounterStatistics.CountHistory[i]})
    }
    trafficMatrixCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    trafficMatrixCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched})
    trafficMatrixCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched})

    trafficMatrixCounterStatistics.EntityData.YListKeys = []string {}

    return &(trafficMatrixCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "traffic-matrix-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/prefixes/prefix/traffic-matrix-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "counters"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/" + tunnels.EntityData.SegmentPath
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Interface is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.InterfaceName, "interface-name")
    tunnel.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/tunnels/" + tunnel.EntityData.SegmentPath
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("base-counter-statistics", types.YChild{"BaseCounterStatistics", &tunnel.BaseCounterStatistics})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", tunnel.InterfaceName})
    tunnel.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", tunnel.InterfaceNameXr})
    tunnel.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", tunnel.InterfaceHandle})
    tunnel.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", tunnel.Vrfid})
    tunnel.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", tunnel.IsActive})

    tunnel.EntityData.YListKeys = []string {"InterfaceName"}

    return &(tunnel.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "tunnel"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/tunnels/tunnel/" + baseCounterStatistics.EntityData.SegmentPath
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range baseCounterStatistics.CountHistory {
        types.SetYListKey(baseCounterStatistics.CountHistory[i], i)
        baseCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(baseCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", baseCounterStatistics.CountHistory[i]})
    }
    baseCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched})
    baseCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched})

    baseCounterStatistics.EntityData.YListKeys = []string {}

    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/vrf-table/default-vrf/afs/af/counters/tunnels/tunnel/base-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

// TrafficCollector_Afs
// Address Family specific operational data
type TrafficCollector_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_Afs_Af.
    Af []*TrafficCollector_Afs_Af
}

func (afs *TrafficCollector_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "traffic-collector"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/" + afs.EntityData.SegmentPath
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// TrafficCollector_Afs_Af
// Operational data for given Address Family
type TrafficCollector_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_Afs_Af_Counters
}

func (af *TrafficCollector_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("counters", types.YChild{"Counters", &af.Counters})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// TrafficCollector_Afs_Af_Counters
// Show Counters
type TrafficCollector_Afs_Af_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_Afs_Af_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "af"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/" + counters.EntityData.SegmentPath
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("prefixes", types.YChild{"Prefixes", &counters.Prefixes})
    counters.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &counters.Tunnels})
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_Afs_Af_Counters_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []*TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "counters"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/" + prefixes.EntityData.SegmentPath
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = types.NewOrderedMap()
    prefixes.EntityData.Children.Append("prefix", types.YChild{"Prefix", nil})
    for i := range prefixes.Prefix {
        types.SetYListKey(prefixes.Prefix[i], i)
        prefixes.EntityData.Children.Append(types.GetSegmentPath(prefixes.Prefix[i]), types.YChild{"Prefix", prefixes.Prefix[i]})
    }
    prefixes.EntityData.Leafs = types.NewOrderedMap()

    prefixes.EntityData.YListKeys = []string {}

    return &(prefixes.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IP Address. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // SR Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // LDP Label. The type is interface{} with range: 0..4294967295.
    LdpLabel interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + types.AddNoKeyToken(prefix)
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/prefixes/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("base-counter-statistics", types.YChild{"BaseCounterStatistics", &prefix.BaseCounterStatistics})
    prefix.EntityData.Children.Append("traffic-matrix-counter-statistics", types.YChild{"TrafficMatrixCounterStatistics", &prefix.TrafficMatrixCounterStatistics})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("ipaddr", types.YLeaf{"Ipaddr", prefix.Ipaddr})
    prefix.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", prefix.Mask})
    prefix.EntityData.Leafs.Append("label", types.YLeaf{"Label", prefix.Label})
    prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefix.Prefix})
    prefix.EntityData.Leafs.Append("label-xr", types.YLeaf{"LabelXr", prefix.LabelXr})
    prefix.EntityData.Leafs.Append("ldp-label", types.YLeaf{"LdpLabel", prefix.LdpLabel})
    prefix.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", prefix.IsActive})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "prefix"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/prefixes/prefix/" + baseCounterStatistics.EntityData.SegmentPath
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range baseCounterStatistics.CountHistory {
        types.SetYListKey(baseCounterStatistics.CountHistory[i], i)
        baseCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(baseCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", baseCounterStatistics.CountHistory[i]})
    }
    baseCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched})
    baseCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched})

    baseCounterStatistics.EntityData.YListKeys = []string {}

    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/prefixes/prefix/base-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetEntityData() *types.CommonEntityData {
    trafficMatrixCounterStatistics.EntityData.YFilter = trafficMatrixCounterStatistics.YFilter
    trafficMatrixCounterStatistics.EntityData.YangName = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    trafficMatrixCounterStatistics.EntityData.ParentYangName = "prefix"
    trafficMatrixCounterStatistics.EntityData.SegmentPath = "traffic-matrix-counter-statistics"
    trafficMatrixCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/prefixes/prefix/" + trafficMatrixCounterStatistics.EntityData.SegmentPath
    trafficMatrixCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMatrixCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMatrixCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMatrixCounterStatistics.EntityData.Children = types.NewOrderedMap()
    trafficMatrixCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range trafficMatrixCounterStatistics.CountHistory {
        types.SetYListKey(trafficMatrixCounterStatistics.CountHistory[i], i)
        trafficMatrixCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(trafficMatrixCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", trafficMatrixCounterStatistics.CountHistory[i]})
    }
    trafficMatrixCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    trafficMatrixCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched})
    trafficMatrixCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched})

    trafficMatrixCounterStatistics.EntityData.YListKeys = []string {}

    return &(trafficMatrixCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "traffic-matrix-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/prefixes/prefix/traffic-matrix-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_Afs_Af_Counters_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []*TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "counters"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/" + tunnels.EntityData.SegmentPath
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface name in Display format. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface VRF ID. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}

    // Interface is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.InterfaceName, "interface-name")
    tunnel.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/tunnels/" + tunnel.EntityData.SegmentPath
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("base-counter-statistics", types.YChild{"BaseCounterStatistics", &tunnel.BaseCounterStatistics})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", tunnel.InterfaceName})
    tunnel.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", tunnel.InterfaceNameXr})
    tunnel.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", tunnel.InterfaceHandle})
    tunnel.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", tunnel.Vrfid})
    tunnel.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", tunnel.IsActive})

    tunnel.EntityData.YListKeys = []string {"InterfaceName"}

    return &(tunnel.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []*TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetEntityData() *types.CommonEntityData {
    baseCounterStatistics.EntityData.YFilter = baseCounterStatistics.YFilter
    baseCounterStatistics.EntityData.YangName = "base-counter-statistics"
    baseCounterStatistics.EntityData.BundleName = "cisco_ios_xr"
    baseCounterStatistics.EntityData.ParentYangName = "tunnel"
    baseCounterStatistics.EntityData.SegmentPath = "base-counter-statistics"
    baseCounterStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/tunnels/tunnel/" + baseCounterStatistics.EntityData.SegmentPath
    baseCounterStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseCounterStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseCounterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseCounterStatistics.EntityData.Children = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Children.Append("count-history", types.YChild{"CountHistory", nil})
    for i := range baseCounterStatistics.CountHistory {
        types.SetYListKey(baseCounterStatistics.CountHistory[i], i)
        baseCounterStatistics.EntityData.Children.Append(types.GetSegmentPath(baseCounterStatistics.CountHistory[i]), types.YChild{"CountHistory", baseCounterStatistics.CountHistory[i]})
    }
    baseCounterStatistics.EntityData.Leafs = types.NewOrderedMap()
    baseCounterStatistics.EntityData.Leafs.Append("transmit-packets-per-second-switched", types.YLeaf{"TransmitPacketsPerSecondSwitched", baseCounterStatistics.TransmitPacketsPerSecondSwitched})
    baseCounterStatistics.EntityData.Leafs.Append("transmit-bytes-per-second-switched", types.YLeaf{"TransmitBytesPerSecondSwitched", baseCounterStatistics.TransmitBytesPerSecondSwitched})

    baseCounterStatistics.EntityData.YListKeys = []string {}

    return &(baseCounterStatistics.EntityData)
}

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Start timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventStartTimestamp interface{}

    // Event End timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    EventEndTimestamp interface{}

    // Number of packets switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615.
    TransmitNumberOfPacketsSwitched interface{}

    // Number of Bytes switched in this interval. The type is interface{} with
    // range: 0..18446744073709551615. Units are byte.
    TransmitNumberOfBytesSwitched interface{}

    // Flag to indicate if this history entry is valid. The type is bool.
    IsValid interface{}
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetEntityData() *types.CommonEntityData {
    countHistory.EntityData.YFilter = countHistory.YFilter
    countHistory.EntityData.YangName = "count-history"
    countHistory.EntityData.BundleName = "cisco_ios_xr"
    countHistory.EntityData.ParentYangName = "base-counter-statistics"
    countHistory.EntityData.SegmentPath = "count-history" + types.AddNoKeyToken(countHistory)
    countHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-oper:traffic-collector/afs/af/counters/tunnels/tunnel/base-counter-statistics/" + countHistory.EntityData.SegmentPath
    countHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    countHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    countHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    countHistory.EntityData.Children = types.NewOrderedMap()
    countHistory.EntityData.Leafs = types.NewOrderedMap()
    countHistory.EntityData.Leafs.Append("event-start-timestamp", types.YLeaf{"EventStartTimestamp", countHistory.EventStartTimestamp})
    countHistory.EntityData.Leafs.Append("event-end-timestamp", types.YLeaf{"EventEndTimestamp", countHistory.EventEndTimestamp})
    countHistory.EntityData.Leafs.Append("transmit-number-of-packets-switched", types.YLeaf{"TransmitNumberOfPacketsSwitched", countHistory.TransmitNumberOfPacketsSwitched})
    countHistory.EntityData.Leafs.Append("transmit-number-of-bytes-switched", types.YLeaf{"TransmitNumberOfBytesSwitched", countHistory.TransmitNumberOfBytesSwitched})
    countHistory.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", countHistory.IsValid})

    countHistory.EntityData.YListKeys = []string {}

    return &(countHistory.EntityData)
}

