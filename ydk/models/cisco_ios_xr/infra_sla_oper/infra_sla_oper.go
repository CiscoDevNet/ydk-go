// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-sla package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sla: SLA oper commands
//   sla-nodes: sla nodes
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_sla_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_sla_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-sla-oper sla}", reflect.TypeOf(Sla{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-sla-oper:sla", reflect.TypeOf(Sla{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-sla-oper sla-nodes}", reflect.TypeOf(SlaNodes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-sla-oper:sla-nodes", reflect.TypeOf(SlaNodes{}))
}

// Sla
// SLA oper commands
type Sla struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of all SLA protocols.
    Protocols Sla_Protocols
}

func (sla *Sla) GetEntityData() *types.CommonEntityData {
    sla.EntityData.YFilter = sla.YFilter
    sla.EntityData.YangName = "sla"
    sla.EntityData.BundleName = "cisco_ios_xr"
    sla.EntityData.ParentYangName = "Cisco-IOS-XR-infra-sla-oper"
    sla.EntityData.SegmentPath = "Cisco-IOS-XR-infra-sla-oper:sla"
    sla.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sla.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sla.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sla.EntityData.Children = types.NewOrderedMap()
    sla.EntityData.Children.Append("protocols", types.YChild{"Protocols", &sla.Protocols})
    sla.EntityData.Leafs = types.NewOrderedMap()

    sla.EntityData.YListKeys = []string {}

    return &(sla.EntityData)
}

// Sla_Protocols
// Table of all SLA protocols
type Sla_Protocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Ethernet SLA protocol.
    Ethernet Sla_Protocols_Ethernet
}

func (protocols *Sla_Protocols) GetEntityData() *types.CommonEntityData {
    protocols.EntityData.YFilter = protocols.YFilter
    protocols.EntityData.YangName = "protocols"
    protocols.EntityData.BundleName = "cisco_ios_xr"
    protocols.EntityData.ParentYangName = "sla"
    protocols.EntityData.SegmentPath = "protocols"
    protocols.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocols.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocols.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocols.EntityData.Children = types.NewOrderedMap()
    protocols.EntityData.Children.Append("Cisco-IOS-XR-ethernet-cfm-oper:ethernet", types.YChild{"Ethernet", &protocols.Ethernet})
    protocols.EntityData.Leafs = types.NewOrderedMap()

    protocols.EntityData.YListKeys = []string {}

    return &(protocols.EntityData)
}

// Sla_Protocols_Ethernet
// The Ethernet SLA protocol
type Sla_Protocols_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of current statistics for SLA on-demand operations.
    StatisticsOnDemandCurrents Sla_Protocols_Ethernet_StatisticsOnDemandCurrents

    // Table of SLA operations.
    Operations Sla_Protocols_Ethernet_Operations

    // Table of historical statistics for SLA operations.
    StatisticsHistoricals Sla_Protocols_Ethernet_StatisticsHistoricals

    // Table of historical statistics for SLA on-demand operations.
    StatisticsOnDemandHistoricals Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals

    // Table of SLA configuration errors on configured operations.
    ConfigErrors Sla_Protocols_Ethernet_ConfigErrors

    // Table of SLA on-demand operations.
    OnDemandOperations Sla_Protocols_Ethernet_OnDemandOperations

    // Table of current statistics for SLA operations.
    StatisticsCurrents Sla_Protocols_Ethernet_StatisticsCurrents
}

func (ethernet *Sla_Protocols_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "cisco_ios_xr"
    ethernet.EntityData.ParentYangName = "protocols"
    ethernet.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-cfm-oper:ethernet"
    ethernet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Children.Append("statistics-on-demand-currents", types.YChild{"StatisticsOnDemandCurrents", &ethernet.StatisticsOnDemandCurrents})
    ethernet.EntityData.Children.Append("operations", types.YChild{"Operations", &ethernet.Operations})
    ethernet.EntityData.Children.Append("statistics-historicals", types.YChild{"StatisticsHistoricals", &ethernet.StatisticsHistoricals})
    ethernet.EntityData.Children.Append("statistics-on-demand-historicals", types.YChild{"StatisticsOnDemandHistoricals", &ethernet.StatisticsOnDemandHistoricals})
    ethernet.EntityData.Children.Append("config-errors", types.YChild{"ConfigErrors", &ethernet.ConfigErrors})
    ethernet.EntityData.Children.Append("on-demand-operations", types.YChild{"OnDemandOperations", &ethernet.OnDemandOperations})
    ethernet.EntityData.Children.Append("statistics-currents", types.YChild{"StatisticsCurrents", &ethernet.StatisticsCurrents})
    ethernet.EntityData.Leafs = types.NewOrderedMap()

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents
// Table of current statistics for SLA on-demand
// operations
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current statistics data for an SLA on-demand operation. The type is slice
    // of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent.
    StatisticsOnDemandCurrent []*Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetEntityData() *types.CommonEntityData {
    statisticsOnDemandCurrents.EntityData.YFilter = statisticsOnDemandCurrents.YFilter
    statisticsOnDemandCurrents.EntityData.YangName = "statistics-on-demand-currents"
    statisticsOnDemandCurrents.EntityData.BundleName = "cisco_ios_xr"
    statisticsOnDemandCurrents.EntityData.ParentYangName = "ethernet"
    statisticsOnDemandCurrents.EntityData.SegmentPath = "statistics-on-demand-currents"
    statisticsOnDemandCurrents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsOnDemandCurrents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsOnDemandCurrents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsOnDemandCurrents.EntityData.Children = types.NewOrderedMap()
    statisticsOnDemandCurrents.EntityData.Children.Append("statistics-on-demand-current", types.YChild{"StatisticsOnDemandCurrent", nil})
    for i := range statisticsOnDemandCurrents.StatisticsOnDemandCurrent {
        statisticsOnDemandCurrents.EntityData.Children.Append(types.GetSegmentPath(statisticsOnDemandCurrents.StatisticsOnDemandCurrent[i]), types.YChild{"StatisticsOnDemandCurrent", statisticsOnDemandCurrents.StatisticsOnDemandCurrent[i]})
    }
    statisticsOnDemandCurrents.EntityData.Leafs = types.NewOrderedMap()

    statisticsOnDemandCurrents.EntityData.YListKeys = []string {}

    return &(statisticsOnDemandCurrents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent
// Current statistics data for an SLA on-demand
// operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation ID. The type is interface{} with range: 1..4294967295.
    OperationId interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule

    // Metrics gathered for the operation. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetEntityData() *types.CommonEntityData {
    statisticsOnDemandCurrent.EntityData.YFilter = statisticsOnDemandCurrent.YFilter
    statisticsOnDemandCurrent.EntityData.YangName = "statistics-on-demand-current"
    statisticsOnDemandCurrent.EntityData.BundleName = "cisco_ios_xr"
    statisticsOnDemandCurrent.EntityData.ParentYangName = "statistics-on-demand-currents"
    statisticsOnDemandCurrent.EntityData.SegmentPath = "statistics-on-demand-current"
    statisticsOnDemandCurrent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsOnDemandCurrent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsOnDemandCurrent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsOnDemandCurrent.EntityData.Children = types.NewOrderedMap()
    statisticsOnDemandCurrent.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &statisticsOnDemandCurrent.SpecificOptions})
    statisticsOnDemandCurrent.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &statisticsOnDemandCurrent.OperationSchedule})
    statisticsOnDemandCurrent.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range statisticsOnDemandCurrent.OperationMetric {
        statisticsOnDemandCurrent.EntityData.Children.Append(types.GetSegmentPath(statisticsOnDemandCurrent.OperationMetric[i]), types.YChild{"OperationMetric", statisticsOnDemandCurrent.OperationMetric[i]})
    }
    statisticsOnDemandCurrent.EntityData.Leafs = types.NewOrderedMap()
    statisticsOnDemandCurrent.EntityData.Leafs.Append("operation-id", types.YLeaf{"OperationId", statisticsOnDemandCurrent.OperationId})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", statisticsOnDemandCurrent.DomainName})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statisticsOnDemandCurrent.InterfaceName})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", statisticsOnDemandCurrent.MepId})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", statisticsOnDemandCurrent.MacAddress})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", statisticsOnDemandCurrent.ProbeType})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", statisticsOnDemandCurrent.DisplayShort})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", statisticsOnDemandCurrent.DisplayLong})
    statisticsOnDemandCurrent.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", statisticsOnDemandCurrent.FlrCalculationInterval})

    statisticsOnDemandCurrent.EntityData.YListKeys = []string {}

    return &(statisticsOnDemandCurrent.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "statistics-on-demand-current"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "statistics-on-demand-current"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket.
    Bucket []*Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "statistics-on-demand-current"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("config", types.YChild{"Config", &operationMetric.Config})
    operationMetric.EntityData.Children.Append("bucket", types.YChild{"Bucket", nil})
    for i := range operationMetric.Bucket {
        operationMetric.EntityData.Children.Append(types.GetSegmentPath(operationMetric.Bucket[i]), types.YChild{"Bucket", operationMetric.Bucket[i]})
    }
    operationMetric.EntityData.Leafs = types.NewOrderedMap()

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "operation-metric"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", config.MetricType})
    config.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", config.BinsCount})
    config.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", config.BinsWidth})
    config.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", config.BucketSize})
    config.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", config.BucketSizeUnit})
    config.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", config.BucketsArchive})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute time that the bucket started being filled at. The type is
    // interface{} with range: 0..4294967295.
    StartAt interface{}

    // Length of time for which the bucket is being filled in seconds. The type is
    // interface{} with range: 0..4294967295. Units are second.
    Duration interface{}

    // Number of packets sent in the probe. The type is interface{} with range:
    // 0..4294967295.
    Sent interface{}

    // Number of lost packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Lost interface{}

    // Number of corrupt packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Corrupt interface{}

    // Number of packets recieved out-of-order in the probe. The type is
    // interface{} with range: 0..4294967295.
    OutOfOrder interface{}

    // Number of duplicate packets received in the probe. The type is interface{}
    // with range: 0..4294967295.
    Duplicates interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Minimum interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Maximum interface{}

    // Absolute time that the minimum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMinimum interface{}

    // Absolute time that the maximum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMaximum interface{}

    // Mean of the results in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Average interface{}

    // Standard deviation of the results in the probe, in microseconds or
    // millionths of a percent. The type is interface{} with range:
    // -2147483648..2147483647.
    StandardDeviation interface{}

    // The count of samples collected in the bucket. The type is interface{} with
    // range: 0..4294967295.
    ResultCount interface{}

    // The number of data packets sent across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataSentCount interface{}

    // The number of data packets lost across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataLostCount interface{}

    // Frame Loss Ratio across the whole bucket, in millionths of a percent. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // percentage.
    OverallFlr interface{}

    // Results suspect due to a probe starting mid-way through a bucket. The type
    // is bool.
    SuspectStartMidBucket interface{}

    // Results suspect due to scheduling latency causing one or more packets to
    // not be sent. The type is bool.
    SuspectScheduleLatency interface{}

    // Results suspect due to failure to send one or more packets. The type is
    // bool.
    SuspectSendFail interface{}

    // Results suspect due to a probe ending prematurely. The type is bool.
    SuspectPrematureEnd interface{}

    // Results suspect as more than 10 seconds time drift detected. The type is
    // bool.
    SuspectClockDrift interface{}

    // Results suspect due to a memory allocation failure. The type is bool.
    SuspectMemoryAllocationFailed interface{}

    // Results suspect as bucket was cleared mid-way through being filled. The
    // type is bool.
    SuspectClearedMidBucket interface{}

    // Results suspect as probe restarted mid-way through the bucket. The type is
    // bool.
    SuspectProbeRestarted interface{}

    // Results suspect as processing of results has been delayed. The type is
    // bool.
    SuspectManagementLatency interface{}

    // Results suspect as the probe has been configured across multiple buckets.
    // The type is bool.
    SuspectMultipleBuckets interface{}

    // Results suspect as misordering has been detected , affecting results. The
    // type is bool.
    SuspectMisordering interface{}

    // Results suspect as FLR calculated based on a low packet count. The type is
    // bool.
    SuspectFlrLowPacketCount interface{}

    // If the probe ended prematurely, the error that caused a probe to end. The
    // type is interface{} with range: 0..4294967295.
    PrematureReason interface{}

    // Description of the error code that caused the probe to end prematurely. For
    // informational purposes only. The type is string.
    PrematureReasonString interface{}

    // The contents of the bucket; bins or samples.
    Contents Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "operation-metric"
    bucket.EntityData.SegmentPath = "bucket"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = types.NewOrderedMap()
    bucket.EntityData.Children.Append("contents", types.YChild{"Contents", &bucket.Contents})
    bucket.EntityData.Leafs = types.NewOrderedMap()
    bucket.EntityData.Leafs.Append("start-at", types.YLeaf{"StartAt", bucket.StartAt})
    bucket.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", bucket.Duration})
    bucket.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", bucket.Sent})
    bucket.EntityData.Leafs.Append("lost", types.YLeaf{"Lost", bucket.Lost})
    bucket.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", bucket.Corrupt})
    bucket.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", bucket.OutOfOrder})
    bucket.EntityData.Leafs.Append("duplicates", types.YLeaf{"Duplicates", bucket.Duplicates})
    bucket.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", bucket.Minimum})
    bucket.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", bucket.Maximum})
    bucket.EntityData.Leafs.Append("time-of-minimum", types.YLeaf{"TimeOfMinimum", bucket.TimeOfMinimum})
    bucket.EntityData.Leafs.Append("time-of-maximum", types.YLeaf{"TimeOfMaximum", bucket.TimeOfMaximum})
    bucket.EntityData.Leafs.Append("average", types.YLeaf{"Average", bucket.Average})
    bucket.EntityData.Leafs.Append("standard-deviation", types.YLeaf{"StandardDeviation", bucket.StandardDeviation})
    bucket.EntityData.Leafs.Append("result-count", types.YLeaf{"ResultCount", bucket.ResultCount})
    bucket.EntityData.Leafs.Append("data-sent-count", types.YLeaf{"DataSentCount", bucket.DataSentCount})
    bucket.EntityData.Leafs.Append("data-lost-count", types.YLeaf{"DataLostCount", bucket.DataLostCount})
    bucket.EntityData.Leafs.Append("overall-flr", types.YLeaf{"OverallFlr", bucket.OverallFlr})
    bucket.EntityData.Leafs.Append("suspect-start-mid-bucket", types.YLeaf{"SuspectStartMidBucket", bucket.SuspectStartMidBucket})
    bucket.EntityData.Leafs.Append("suspect-schedule-latency", types.YLeaf{"SuspectScheduleLatency", bucket.SuspectScheduleLatency})
    bucket.EntityData.Leafs.Append("suspect-send-fail", types.YLeaf{"SuspectSendFail", bucket.SuspectSendFail})
    bucket.EntityData.Leafs.Append("suspect-premature-end", types.YLeaf{"SuspectPrematureEnd", bucket.SuspectPrematureEnd})
    bucket.EntityData.Leafs.Append("suspect-clock-drift", types.YLeaf{"SuspectClockDrift", bucket.SuspectClockDrift})
    bucket.EntityData.Leafs.Append("suspect-memory-allocation-failed", types.YLeaf{"SuspectMemoryAllocationFailed", bucket.SuspectMemoryAllocationFailed})
    bucket.EntityData.Leafs.Append("suspect-cleared-mid-bucket", types.YLeaf{"SuspectClearedMidBucket", bucket.SuspectClearedMidBucket})
    bucket.EntityData.Leafs.Append("suspect-probe-restarted", types.YLeaf{"SuspectProbeRestarted", bucket.SuspectProbeRestarted})
    bucket.EntityData.Leafs.Append("suspect-management-latency", types.YLeaf{"SuspectManagementLatency", bucket.SuspectManagementLatency})
    bucket.EntityData.Leafs.Append("suspect-multiple-buckets", types.YLeaf{"SuspectMultipleBuckets", bucket.SuspectMultipleBuckets})
    bucket.EntityData.Leafs.Append("suspect-misordering", types.YLeaf{"SuspectMisordering", bucket.SuspectMisordering})
    bucket.EntityData.Leafs.Append("suspect-flr-low-packet-count", types.YLeaf{"SuspectFlrLowPacketCount", bucket.SuspectFlrLowPacketCount})
    bucket.EntityData.Leafs.Append("premature-reason", types.YLeaf{"PrematureReason", bucket.PrematureReason})
    bucket.EntityData.Leafs.Append("premature-reason-string", types.YLeaf{"PrematureReasonString", bucket.PrematureReasonString})

    bucket.EntityData.YListKeys = []string {}

    return &(bucket.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetEntityData() *types.CommonEntityData {
    contents.EntityData.YFilter = contents.YFilter
    contents.EntityData.YangName = "contents"
    contents.EntityData.BundleName = "cisco_ios_xr"
    contents.EntityData.ParentYangName = "bucket"
    contents.EntityData.SegmentPath = "contents"
    contents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contents.EntityData.Children = types.NewOrderedMap()
    contents.EntityData.Children.Append("aggregated", types.YChild{"Aggregated", &contents.Aggregated})
    contents.EntityData.Children.Append("unaggregated", types.YChild{"Unaggregated", &contents.Unaggregated})
    contents.EntityData.Leafs = types.NewOrderedMap()
    contents.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", contents.BucketType})

    contents.EntityData.YListKeys = []string {}

    return &(contents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []*Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "contents"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = types.NewOrderedMap()
    aggregated.EntityData.Children.Append("bins", types.YChild{"Bins", nil})
    for i := range aggregated.Bins {
        aggregated.EntityData.Children.Append(types.GetSegmentPath(aggregated.Bins[i]), types.YChild{"Bins", aggregated.Bins[i]})
    }
    aggregated.EntityData.Leafs = types.NewOrderedMap()

    aggregated.EntityData.YListKeys = []string {}

    return &(aggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower bound (inclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    LowerBound interface{}

    // Upper bound (exclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    UpperBound interface{}

    // Lower bound (inclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    LowerBoundTenths interface{}

    // Upper bound (exclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    UpperBoundTenths interface{}

    // The sum of the results in the bin, in microseconds or millionths of a
    // percent. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Sum interface{}

    // The total number of results in the bin. The type is interface{} with range:
    // 0..4294967295.
    Count interface{}
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetEntityData() *types.CommonEntityData {
    bins.EntityData.YFilter = bins.YFilter
    bins.EntityData.YangName = "bins"
    bins.EntityData.BundleName = "cisco_ios_xr"
    bins.EntityData.ParentYangName = "aggregated"
    bins.EntityData.SegmentPath = "bins"
    bins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bins.EntityData.Children = types.NewOrderedMap()
    bins.EntityData.Leafs = types.NewOrderedMap()
    bins.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", bins.LowerBound})
    bins.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", bins.UpperBound})
    bins.EntityData.Leafs.Append("lower-bound-tenths", types.YLeaf{"LowerBoundTenths", bins.LowerBoundTenths})
    bins.EntityData.Leafs.Append("upper-bound-tenths", types.YLeaf{"UpperBoundTenths", bins.UpperBoundTenths})
    bins.EntityData.Leafs.Append("sum", types.YLeaf{"Sum", bins.Sum})
    bins.EntityData.Leafs.Append("count", types.YLeaf{"Count", bins.Count})

    bins.EntityData.YListKeys = []string {}

    return &(bins.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []*Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetEntityData() *types.CommonEntityData {
    unaggregated.EntityData.YFilter = unaggregated.YFilter
    unaggregated.EntityData.YangName = "unaggregated"
    unaggregated.EntityData.BundleName = "cisco_ios_xr"
    unaggregated.EntityData.ParentYangName = "contents"
    unaggregated.EntityData.SegmentPath = "unaggregated"
    unaggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unaggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unaggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unaggregated.EntityData.Children = types.NewOrderedMap()
    unaggregated.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range unaggregated.Sample {
        unaggregated.EntityData.Children.Append(types.GetSegmentPath(unaggregated.Sample[i]), types.YChild{"Sample", unaggregated.Sample[i]})
    }
    unaggregated.EntityData.Leafs = types.NewOrderedMap()

    unaggregated.EntityData.YListKeys = []string {}

    return &(unaggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time (in milliseconds relative to the start time of the bucket) that
    // the sample was sent at. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    SentAt interface{}

    // Whether the sample packet was sucessfully sent. The type is bool.
    Sent interface{}

    // Whether the sample packet timed out. The type is bool.
    TimedOut interface{}

    // Whether the sample packet was corrupt. The type is bool.
    Corrupt interface{}

    // Whether the sample packet was received out-of-order. The type is bool.
    OutOfOrder interface{}

    // Whether a measurement could not be made because no data packets were sent
    // in the sample period. Only applicable for LMM measurements. The type is
    // bool.
    NoDataPackets interface{}

    // The result (in microseconds or millionths of a percent) of the sample, if
    // available. The type is interface{} with range: -2147483648..2147483647.
    Result interface{}

    // For FLR measurements, the number of frames sent, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesSent interface{}

    // For FLR measurements, the number of frames lost, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesLost interface{}
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "unaggregated"
    sample.EntityData.SegmentPath = "sample"
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sent-at", types.YLeaf{"SentAt", sample.SentAt})
    sample.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sample.Sent})
    sample.EntityData.Leafs.Append("timed-out", types.YLeaf{"TimedOut", sample.TimedOut})
    sample.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", sample.Corrupt})
    sample.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", sample.OutOfOrder})
    sample.EntityData.Leafs.Append("no-data-packets", types.YLeaf{"NoDataPackets", sample.NoDataPackets})
    sample.EntityData.Leafs.Append("result", types.YLeaf{"Result", sample.Result})
    sample.EntityData.Leafs.Append("frames-sent", types.YLeaf{"FramesSent", sample.FramesSent})
    sample.EntityData.Leafs.Append("frames-lost", types.YLeaf{"FramesLost", sample.FramesLost})

    sample.EntityData.YListKeys = []string {}

    return &(sample.EntityData)
}

// Sla_Protocols_Ethernet_Operations
// Table of SLA operations
type Sla_Protocols_Ethernet_Operations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SLA operation to get operation data for. The type is slice of
    // Sla_Protocols_Ethernet_Operations_Operation.
    Operation []*Sla_Protocols_Ethernet_Operations_Operation
}

func (operations *Sla_Protocols_Ethernet_Operations) GetEntityData() *types.CommonEntityData {
    operations.EntityData.YFilter = operations.YFilter
    operations.EntityData.YangName = "operations"
    operations.EntityData.BundleName = "cisco_ios_xr"
    operations.EntityData.ParentYangName = "ethernet"
    operations.EntityData.SegmentPath = "operations"
    operations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operations.EntityData.Children = types.NewOrderedMap()
    operations.EntityData.Children.Append("operation", types.YChild{"Operation", nil})
    for i := range operations.Operation {
        operations.EntityData.Children.Append(types.GetSegmentPath(operations.Operation[i]), types.YChild{"Operation", operations.Operation[i]})
    }
    operations.EntityData.Leafs = types.NewOrderedMap()

    operations.EntityData.YListKeys = []string {}

    return &(operations.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation
// SLA operation to get operation data for
type Sla_Protocols_Ethernet_Operations_Operation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Time that the last probe for the operation was run, NULL if never run. The
    // type is interface{} with range: 0..4294967295.
    LastRun interface{}

    // Options that are only valid if the operation has a profile.
    ProfileOptions Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetEntityData() *types.CommonEntityData {
    operation.EntityData.YFilter = operation.YFilter
    operation.EntityData.YangName = "operation"
    operation.EntityData.BundleName = "cisco_ios_xr"
    operation.EntityData.ParentYangName = "operations"
    operation.EntityData.SegmentPath = "operation"
    operation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operation.EntityData.Children = types.NewOrderedMap()
    operation.EntityData.Children.Append("profile-options", types.YChild{"ProfileOptions", &operation.ProfileOptions})
    operation.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &operation.SpecificOptions})
    operation.EntityData.Leafs = types.NewOrderedMap()
    operation.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", operation.ProfileName})
    operation.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", operation.DomainName})
    operation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", operation.InterfaceName})
    operation.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", operation.MepId})
    operation.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", operation.MacAddress})
    operation.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", operation.DisplayShort})
    operation.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", operation.DisplayLong})
    operation.EntityData.Leafs.Append("last-run", types.YLeaf{"LastRun", operation.LastRun})

    operation.EntityData.YListKeys = []string {}

    return &(operation.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions
// Options that are only valid if the operation has
// a profile
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Number of packets sent per burst. The type is interface{} with range:
    // 0..65535.
    PacketsPerBurst interface{}

    // Interval between packets within a burst in milliseconds. The type is
    // interface{} with range: 0..65535. Units are millisecond.
    InterPacketInterval interface{}

    // Number of bursts sent per probe. The type is interface{} with range:
    // 0..4294967295.
    BurstsPerProbe interface{}

    // Interval between bursts within a probe in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    InterBurstInterval interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Configuration of the packet padding.
    PacketPadding Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding

    // Priority at which to send the packet, if configured.
    Priority Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule

    // Array of the metrics that are measured by the operation. The type is slice
    // of
    // Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetEntityData() *types.CommonEntityData {
    profileOptions.EntityData.YFilter = profileOptions.YFilter
    profileOptions.EntityData.YangName = "profile-options"
    profileOptions.EntityData.BundleName = "cisco_ios_xr"
    profileOptions.EntityData.ParentYangName = "operation"
    profileOptions.EntityData.SegmentPath = "profile-options"
    profileOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileOptions.EntityData.Children = types.NewOrderedMap()
    profileOptions.EntityData.Children.Append("packet-padding", types.YChild{"PacketPadding", &profileOptions.PacketPadding})
    profileOptions.EntityData.Children.Append("priority", types.YChild{"Priority", &profileOptions.Priority})
    profileOptions.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &profileOptions.OperationSchedule})
    profileOptions.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range profileOptions.OperationMetric {
        profileOptions.EntityData.Children.Append(types.GetSegmentPath(profileOptions.OperationMetric[i]), types.YChild{"OperationMetric", profileOptions.OperationMetric[i]})
    }
    profileOptions.EntityData.Leafs = types.NewOrderedMap()
    profileOptions.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", profileOptions.ProbeType})
    profileOptions.EntityData.Leafs.Append("packets-per-burst", types.YLeaf{"PacketsPerBurst", profileOptions.PacketsPerBurst})
    profileOptions.EntityData.Leafs.Append("inter-packet-interval", types.YLeaf{"InterPacketInterval", profileOptions.InterPacketInterval})
    profileOptions.EntityData.Leafs.Append("bursts-per-probe", types.YLeaf{"BurstsPerProbe", profileOptions.BurstsPerProbe})
    profileOptions.EntityData.Leafs.Append("inter-burst-interval", types.YLeaf{"InterBurstInterval", profileOptions.InterBurstInterval})
    profileOptions.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", profileOptions.FlrCalculationInterval})

    profileOptions.EntityData.YListKeys = []string {}

    return &(profileOptions.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding
// Configuration of the packet padding
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size that packets are being padded to. The type is interface{} with range:
    // 0..65535.
    PacketPadSize interface{}

    // Test pattern scheme that is used in the packet padding. The type is
    // SlaOperTestPatternScheme.
    TestPatternPadScheme interface{}

    // Hex string that is used in the packet padding. The type is interface{} with
    // range: 0..4294967295.
    TestPatternPadHexString interface{}
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetEntityData() *types.CommonEntityData {
    packetPadding.EntityData.YFilter = packetPadding.YFilter
    packetPadding.EntityData.YangName = "packet-padding"
    packetPadding.EntityData.BundleName = "cisco_ios_xr"
    packetPadding.EntityData.ParentYangName = "profile-options"
    packetPadding.EntityData.SegmentPath = "packet-padding"
    packetPadding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetPadding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetPadding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetPadding.EntityData.Children = types.NewOrderedMap()
    packetPadding.EntityData.Leafs = types.NewOrderedMap()
    packetPadding.EntityData.Leafs.Append("packet-pad-size", types.YLeaf{"PacketPadSize", packetPadding.PacketPadSize})
    packetPadding.EntityData.Leafs.Append("test-pattern-pad-scheme", types.YLeaf{"TestPatternPadScheme", packetPadding.TestPatternPadScheme})
    packetPadding.EntityData.Leafs.Append("test-pattern-pad-hex-string", types.YLeaf{"TestPatternPadHexString", packetPadding.TestPatternPadHexString})

    packetPadding.EntityData.YListKeys = []string {}

    return &(packetPadding.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority
// Priority at which to send the packet, if
// configured
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PriorityType. The type is SlaOperPacketPriority.
    PriorityType interface{}

    // 3-bit COS priority value applied to packets. The type is interface{} with
    // range: 0..255.
    Cos interface{}
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "profile-options"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("priority-type", types.YLeaf{"PriorityType", priority.PriorityType})
    priority.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", priority.Cos})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "profile-options"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric
// Array of the metrics that are measured by the
// operation
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of valid buckets currently in the buckets archive. The type is
    // interface{} with range: 0..4294967295.
    CurrentBucketsArchive interface{}

    // Configuration of the metric.
    MetricConfig Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "profile-options"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("metric-config", types.YChild{"MetricConfig", &operationMetric.MetricConfig})
    operationMetric.EntityData.Leafs = types.NewOrderedMap()
    operationMetric.EntityData.Leafs.Append("current-buckets-archive", types.YLeaf{"CurrentBucketsArchive", operationMetric.CurrentBucketsArchive})

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig
// Configuration of the metric
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetEntityData() *types.CommonEntityData {
    metricConfig.EntityData.YFilter = metricConfig.YFilter
    metricConfig.EntityData.YangName = "metric-config"
    metricConfig.EntityData.BundleName = "cisco_ios_xr"
    metricConfig.EntityData.ParentYangName = "operation-metric"
    metricConfig.EntityData.SegmentPath = "metric-config"
    metricConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricConfig.EntityData.Children = types.NewOrderedMap()
    metricConfig.EntityData.Leafs = types.NewOrderedMap()
    metricConfig.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", metricConfig.MetricType})
    metricConfig.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", metricConfig.BinsCount})
    metricConfig.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", metricConfig.BinsWidth})
    metricConfig.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", metricConfig.BucketSize})
    metricConfig.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", metricConfig.BucketSizeUnit})
    metricConfig.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", metricConfig.BucketsArchive})

    metricConfig.EntityData.YListKeys = []string {}

    return &(metricConfig.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "operation"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals
// Table of historical statistics for SLA
// operations
type Sla_Protocols_Ethernet_StatisticsHistoricals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Historical statistics data for an SLA configured operation. The type is
    // slice of Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical.
    StatisticsHistorical []*Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetEntityData() *types.CommonEntityData {
    statisticsHistoricals.EntityData.YFilter = statisticsHistoricals.YFilter
    statisticsHistoricals.EntityData.YangName = "statistics-historicals"
    statisticsHistoricals.EntityData.BundleName = "cisco_ios_xr"
    statisticsHistoricals.EntityData.ParentYangName = "ethernet"
    statisticsHistoricals.EntityData.SegmentPath = "statistics-historicals"
    statisticsHistoricals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsHistoricals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsHistoricals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsHistoricals.EntityData.Children = types.NewOrderedMap()
    statisticsHistoricals.EntityData.Children.Append("statistics-historical", types.YChild{"StatisticsHistorical", nil})
    for i := range statisticsHistoricals.StatisticsHistorical {
        statisticsHistoricals.EntityData.Children.Append(types.GetSegmentPath(statisticsHistoricals.StatisticsHistorical[i]), types.YChild{"StatisticsHistorical", statisticsHistoricals.StatisticsHistorical[i]})
    }
    statisticsHistoricals.EntityData.Leafs = types.NewOrderedMap()

    statisticsHistoricals.EntityData.YListKeys = []string {}

    return &(statisticsHistoricals.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical
// Historical statistics data for an SLA
// configured operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule

    // Metrics gathered for the operation. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetEntityData() *types.CommonEntityData {
    statisticsHistorical.EntityData.YFilter = statisticsHistorical.YFilter
    statisticsHistorical.EntityData.YangName = "statistics-historical"
    statisticsHistorical.EntityData.BundleName = "cisco_ios_xr"
    statisticsHistorical.EntityData.ParentYangName = "statistics-historicals"
    statisticsHistorical.EntityData.SegmentPath = "statistics-historical"
    statisticsHistorical.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsHistorical.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsHistorical.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsHistorical.EntityData.Children = types.NewOrderedMap()
    statisticsHistorical.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &statisticsHistorical.SpecificOptions})
    statisticsHistorical.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &statisticsHistorical.OperationSchedule})
    statisticsHistorical.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range statisticsHistorical.OperationMetric {
        statisticsHistorical.EntityData.Children.Append(types.GetSegmentPath(statisticsHistorical.OperationMetric[i]), types.YChild{"OperationMetric", statisticsHistorical.OperationMetric[i]})
    }
    statisticsHistorical.EntityData.Leafs = types.NewOrderedMap()
    statisticsHistorical.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", statisticsHistorical.ProfileName})
    statisticsHistorical.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", statisticsHistorical.DomainName})
    statisticsHistorical.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statisticsHistorical.InterfaceName})
    statisticsHistorical.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", statisticsHistorical.MepId})
    statisticsHistorical.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", statisticsHistorical.MacAddress})
    statisticsHistorical.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", statisticsHistorical.ProbeType})
    statisticsHistorical.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", statisticsHistorical.DisplayShort})
    statisticsHistorical.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", statisticsHistorical.DisplayLong})
    statisticsHistorical.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", statisticsHistorical.FlrCalculationInterval})

    statisticsHistorical.EntityData.YListKeys = []string {}

    return &(statisticsHistorical.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "statistics-historical"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "statistics-historical"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket.
    Bucket []*Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "statistics-historical"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("config", types.YChild{"Config", &operationMetric.Config})
    operationMetric.EntityData.Children.Append("bucket", types.YChild{"Bucket", nil})
    for i := range operationMetric.Bucket {
        operationMetric.EntityData.Children.Append(types.GetSegmentPath(operationMetric.Bucket[i]), types.YChild{"Bucket", operationMetric.Bucket[i]})
    }
    operationMetric.EntityData.Leafs = types.NewOrderedMap()

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "operation-metric"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", config.MetricType})
    config.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", config.BinsCount})
    config.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", config.BinsWidth})
    config.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", config.BucketSize})
    config.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", config.BucketSizeUnit})
    config.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", config.BucketsArchive})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute time that the bucket started being filled at. The type is
    // interface{} with range: 0..4294967295.
    StartAt interface{}

    // Length of time for which the bucket is being filled in seconds. The type is
    // interface{} with range: 0..4294967295. Units are second.
    Duration interface{}

    // Number of packets sent in the probe. The type is interface{} with range:
    // 0..4294967295.
    Sent interface{}

    // Number of lost packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Lost interface{}

    // Number of corrupt packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Corrupt interface{}

    // Number of packets recieved out-of-order in the probe. The type is
    // interface{} with range: 0..4294967295.
    OutOfOrder interface{}

    // Number of duplicate packets received in the probe. The type is interface{}
    // with range: 0..4294967295.
    Duplicates interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Minimum interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Maximum interface{}

    // Absolute time that the minimum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMinimum interface{}

    // Absolute time that the maximum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMaximum interface{}

    // Mean of the results in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Average interface{}

    // Standard deviation of the results in the probe, in microseconds or
    // millionths of a percent. The type is interface{} with range:
    // -2147483648..2147483647.
    StandardDeviation interface{}

    // The count of samples collected in the bucket. The type is interface{} with
    // range: 0..4294967295.
    ResultCount interface{}

    // The number of data packets sent across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataSentCount interface{}

    // The number of data packets lost across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataLostCount interface{}

    // Frame Loss Ratio across the whole bucket, in millionths of a percent. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // percentage.
    OverallFlr interface{}

    // Results suspect due to a probe starting mid-way through a bucket. The type
    // is bool.
    SuspectStartMidBucket interface{}

    // Results suspect due to scheduling latency causing one or more packets to
    // not be sent. The type is bool.
    SuspectScheduleLatency interface{}

    // Results suspect due to failure to send one or more packets. The type is
    // bool.
    SuspectSendFail interface{}

    // Results suspect due to a probe ending prematurely. The type is bool.
    SuspectPrematureEnd interface{}

    // Results suspect as more than 10 seconds time drift detected. The type is
    // bool.
    SuspectClockDrift interface{}

    // Results suspect due to a memory allocation failure. The type is bool.
    SuspectMemoryAllocationFailed interface{}

    // Results suspect as bucket was cleared mid-way through being filled. The
    // type is bool.
    SuspectClearedMidBucket interface{}

    // Results suspect as probe restarted mid-way through the bucket. The type is
    // bool.
    SuspectProbeRestarted interface{}

    // Results suspect as processing of results has been delayed. The type is
    // bool.
    SuspectManagementLatency interface{}

    // Results suspect as the probe has been configured across multiple buckets.
    // The type is bool.
    SuspectMultipleBuckets interface{}

    // Results suspect as misordering has been detected , affecting results. The
    // type is bool.
    SuspectMisordering interface{}

    // Results suspect as FLR calculated based on a low packet count. The type is
    // bool.
    SuspectFlrLowPacketCount interface{}

    // If the probe ended prematurely, the error that caused a probe to end. The
    // type is interface{} with range: 0..4294967295.
    PrematureReason interface{}

    // Description of the error code that caused the probe to end prematurely. For
    // informational purposes only. The type is string.
    PrematureReasonString interface{}

    // The contents of the bucket; bins or samples.
    Contents Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "operation-metric"
    bucket.EntityData.SegmentPath = "bucket"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = types.NewOrderedMap()
    bucket.EntityData.Children.Append("contents", types.YChild{"Contents", &bucket.Contents})
    bucket.EntityData.Leafs = types.NewOrderedMap()
    bucket.EntityData.Leafs.Append("start-at", types.YLeaf{"StartAt", bucket.StartAt})
    bucket.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", bucket.Duration})
    bucket.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", bucket.Sent})
    bucket.EntityData.Leafs.Append("lost", types.YLeaf{"Lost", bucket.Lost})
    bucket.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", bucket.Corrupt})
    bucket.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", bucket.OutOfOrder})
    bucket.EntityData.Leafs.Append("duplicates", types.YLeaf{"Duplicates", bucket.Duplicates})
    bucket.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", bucket.Minimum})
    bucket.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", bucket.Maximum})
    bucket.EntityData.Leafs.Append("time-of-minimum", types.YLeaf{"TimeOfMinimum", bucket.TimeOfMinimum})
    bucket.EntityData.Leafs.Append("time-of-maximum", types.YLeaf{"TimeOfMaximum", bucket.TimeOfMaximum})
    bucket.EntityData.Leafs.Append("average", types.YLeaf{"Average", bucket.Average})
    bucket.EntityData.Leafs.Append("standard-deviation", types.YLeaf{"StandardDeviation", bucket.StandardDeviation})
    bucket.EntityData.Leafs.Append("result-count", types.YLeaf{"ResultCount", bucket.ResultCount})
    bucket.EntityData.Leafs.Append("data-sent-count", types.YLeaf{"DataSentCount", bucket.DataSentCount})
    bucket.EntityData.Leafs.Append("data-lost-count", types.YLeaf{"DataLostCount", bucket.DataLostCount})
    bucket.EntityData.Leafs.Append("overall-flr", types.YLeaf{"OverallFlr", bucket.OverallFlr})
    bucket.EntityData.Leafs.Append("suspect-start-mid-bucket", types.YLeaf{"SuspectStartMidBucket", bucket.SuspectStartMidBucket})
    bucket.EntityData.Leafs.Append("suspect-schedule-latency", types.YLeaf{"SuspectScheduleLatency", bucket.SuspectScheduleLatency})
    bucket.EntityData.Leafs.Append("suspect-send-fail", types.YLeaf{"SuspectSendFail", bucket.SuspectSendFail})
    bucket.EntityData.Leafs.Append("suspect-premature-end", types.YLeaf{"SuspectPrematureEnd", bucket.SuspectPrematureEnd})
    bucket.EntityData.Leafs.Append("suspect-clock-drift", types.YLeaf{"SuspectClockDrift", bucket.SuspectClockDrift})
    bucket.EntityData.Leafs.Append("suspect-memory-allocation-failed", types.YLeaf{"SuspectMemoryAllocationFailed", bucket.SuspectMemoryAllocationFailed})
    bucket.EntityData.Leafs.Append("suspect-cleared-mid-bucket", types.YLeaf{"SuspectClearedMidBucket", bucket.SuspectClearedMidBucket})
    bucket.EntityData.Leafs.Append("suspect-probe-restarted", types.YLeaf{"SuspectProbeRestarted", bucket.SuspectProbeRestarted})
    bucket.EntityData.Leafs.Append("suspect-management-latency", types.YLeaf{"SuspectManagementLatency", bucket.SuspectManagementLatency})
    bucket.EntityData.Leafs.Append("suspect-multiple-buckets", types.YLeaf{"SuspectMultipleBuckets", bucket.SuspectMultipleBuckets})
    bucket.EntityData.Leafs.Append("suspect-misordering", types.YLeaf{"SuspectMisordering", bucket.SuspectMisordering})
    bucket.EntityData.Leafs.Append("suspect-flr-low-packet-count", types.YLeaf{"SuspectFlrLowPacketCount", bucket.SuspectFlrLowPacketCount})
    bucket.EntityData.Leafs.Append("premature-reason", types.YLeaf{"PrematureReason", bucket.PrematureReason})
    bucket.EntityData.Leafs.Append("premature-reason-string", types.YLeaf{"PrematureReasonString", bucket.PrematureReasonString})

    bucket.EntityData.YListKeys = []string {}

    return &(bucket.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetEntityData() *types.CommonEntityData {
    contents.EntityData.YFilter = contents.YFilter
    contents.EntityData.YangName = "contents"
    contents.EntityData.BundleName = "cisco_ios_xr"
    contents.EntityData.ParentYangName = "bucket"
    contents.EntityData.SegmentPath = "contents"
    contents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contents.EntityData.Children = types.NewOrderedMap()
    contents.EntityData.Children.Append("aggregated", types.YChild{"Aggregated", &contents.Aggregated})
    contents.EntityData.Children.Append("unaggregated", types.YChild{"Unaggregated", &contents.Unaggregated})
    contents.EntityData.Leafs = types.NewOrderedMap()
    contents.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", contents.BucketType})

    contents.EntityData.YListKeys = []string {}

    return &(contents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []*Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "contents"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = types.NewOrderedMap()
    aggregated.EntityData.Children.Append("bins", types.YChild{"Bins", nil})
    for i := range aggregated.Bins {
        aggregated.EntityData.Children.Append(types.GetSegmentPath(aggregated.Bins[i]), types.YChild{"Bins", aggregated.Bins[i]})
    }
    aggregated.EntityData.Leafs = types.NewOrderedMap()

    aggregated.EntityData.YListKeys = []string {}

    return &(aggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower bound (inclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    LowerBound interface{}

    // Upper bound (exclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    UpperBound interface{}

    // Lower bound (inclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    LowerBoundTenths interface{}

    // Upper bound (exclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    UpperBoundTenths interface{}

    // The sum of the results in the bin, in microseconds or millionths of a
    // percent. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Sum interface{}

    // The total number of results in the bin. The type is interface{} with range:
    // 0..4294967295.
    Count interface{}
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetEntityData() *types.CommonEntityData {
    bins.EntityData.YFilter = bins.YFilter
    bins.EntityData.YangName = "bins"
    bins.EntityData.BundleName = "cisco_ios_xr"
    bins.EntityData.ParentYangName = "aggregated"
    bins.EntityData.SegmentPath = "bins"
    bins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bins.EntityData.Children = types.NewOrderedMap()
    bins.EntityData.Leafs = types.NewOrderedMap()
    bins.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", bins.LowerBound})
    bins.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", bins.UpperBound})
    bins.EntityData.Leafs.Append("lower-bound-tenths", types.YLeaf{"LowerBoundTenths", bins.LowerBoundTenths})
    bins.EntityData.Leafs.Append("upper-bound-tenths", types.YLeaf{"UpperBoundTenths", bins.UpperBoundTenths})
    bins.EntityData.Leafs.Append("sum", types.YLeaf{"Sum", bins.Sum})
    bins.EntityData.Leafs.Append("count", types.YLeaf{"Count", bins.Count})

    bins.EntityData.YListKeys = []string {}

    return &(bins.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []*Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetEntityData() *types.CommonEntityData {
    unaggregated.EntityData.YFilter = unaggregated.YFilter
    unaggregated.EntityData.YangName = "unaggregated"
    unaggregated.EntityData.BundleName = "cisco_ios_xr"
    unaggregated.EntityData.ParentYangName = "contents"
    unaggregated.EntityData.SegmentPath = "unaggregated"
    unaggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unaggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unaggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unaggregated.EntityData.Children = types.NewOrderedMap()
    unaggregated.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range unaggregated.Sample {
        unaggregated.EntityData.Children.Append(types.GetSegmentPath(unaggregated.Sample[i]), types.YChild{"Sample", unaggregated.Sample[i]})
    }
    unaggregated.EntityData.Leafs = types.NewOrderedMap()

    unaggregated.EntityData.YListKeys = []string {}

    return &(unaggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time (in milliseconds relative to the start time of the bucket) that
    // the sample was sent at. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    SentAt interface{}

    // Whether the sample packet was sucessfully sent. The type is bool.
    Sent interface{}

    // Whether the sample packet timed out. The type is bool.
    TimedOut interface{}

    // Whether the sample packet was corrupt. The type is bool.
    Corrupt interface{}

    // Whether the sample packet was received out-of-order. The type is bool.
    OutOfOrder interface{}

    // Whether a measurement could not be made because no data packets were sent
    // in the sample period. Only applicable for LMM measurements. The type is
    // bool.
    NoDataPackets interface{}

    // The result (in microseconds or millionths of a percent) of the sample, if
    // available. The type is interface{} with range: -2147483648..2147483647.
    Result interface{}

    // For FLR measurements, the number of frames sent, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesSent interface{}

    // For FLR measurements, the number of frames lost, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesLost interface{}
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "unaggregated"
    sample.EntityData.SegmentPath = "sample"
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sent-at", types.YLeaf{"SentAt", sample.SentAt})
    sample.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sample.Sent})
    sample.EntityData.Leafs.Append("timed-out", types.YLeaf{"TimedOut", sample.TimedOut})
    sample.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", sample.Corrupt})
    sample.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", sample.OutOfOrder})
    sample.EntityData.Leafs.Append("no-data-packets", types.YLeaf{"NoDataPackets", sample.NoDataPackets})
    sample.EntityData.Leafs.Append("result", types.YLeaf{"Result", sample.Result})
    sample.EntityData.Leafs.Append("frames-sent", types.YLeaf{"FramesSent", sample.FramesSent})
    sample.EntityData.Leafs.Append("frames-lost", types.YLeaf{"FramesLost", sample.FramesLost})

    sample.EntityData.YListKeys = []string {}

    return &(sample.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals
// Table of historical statistics for SLA
// on-demand operations
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Historical statistics data for an SLA on-demand  operation. The type is
    // slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical.
    StatisticsOnDemandHistorical []*Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetEntityData() *types.CommonEntityData {
    statisticsOnDemandHistoricals.EntityData.YFilter = statisticsOnDemandHistoricals.YFilter
    statisticsOnDemandHistoricals.EntityData.YangName = "statistics-on-demand-historicals"
    statisticsOnDemandHistoricals.EntityData.BundleName = "cisco_ios_xr"
    statisticsOnDemandHistoricals.EntityData.ParentYangName = "ethernet"
    statisticsOnDemandHistoricals.EntityData.SegmentPath = "statistics-on-demand-historicals"
    statisticsOnDemandHistoricals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsOnDemandHistoricals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsOnDemandHistoricals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsOnDemandHistoricals.EntityData.Children = types.NewOrderedMap()
    statisticsOnDemandHistoricals.EntityData.Children.Append("statistics-on-demand-historical", types.YChild{"StatisticsOnDemandHistorical", nil})
    for i := range statisticsOnDemandHistoricals.StatisticsOnDemandHistorical {
        statisticsOnDemandHistoricals.EntityData.Children.Append(types.GetSegmentPath(statisticsOnDemandHistoricals.StatisticsOnDemandHistorical[i]), types.YChild{"StatisticsOnDemandHistorical", statisticsOnDemandHistoricals.StatisticsOnDemandHistorical[i]})
    }
    statisticsOnDemandHistoricals.EntityData.Leafs = types.NewOrderedMap()

    statisticsOnDemandHistoricals.EntityData.YListKeys = []string {}

    return &(statisticsOnDemandHistoricals.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical
// Historical statistics data for an SLA
// on-demand  operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation ID. The type is interface{} with range: 1..4294967295.
    OperationId interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule

    // Metrics gathered for the operation. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetEntityData() *types.CommonEntityData {
    statisticsOnDemandHistorical.EntityData.YFilter = statisticsOnDemandHistorical.YFilter
    statisticsOnDemandHistorical.EntityData.YangName = "statistics-on-demand-historical"
    statisticsOnDemandHistorical.EntityData.BundleName = "cisco_ios_xr"
    statisticsOnDemandHistorical.EntityData.ParentYangName = "statistics-on-demand-historicals"
    statisticsOnDemandHistorical.EntityData.SegmentPath = "statistics-on-demand-historical"
    statisticsOnDemandHistorical.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsOnDemandHistorical.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsOnDemandHistorical.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsOnDemandHistorical.EntityData.Children = types.NewOrderedMap()
    statisticsOnDemandHistorical.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &statisticsOnDemandHistorical.SpecificOptions})
    statisticsOnDemandHistorical.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &statisticsOnDemandHistorical.OperationSchedule})
    statisticsOnDemandHistorical.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range statisticsOnDemandHistorical.OperationMetric {
        statisticsOnDemandHistorical.EntityData.Children.Append(types.GetSegmentPath(statisticsOnDemandHistorical.OperationMetric[i]), types.YChild{"OperationMetric", statisticsOnDemandHistorical.OperationMetric[i]})
    }
    statisticsOnDemandHistorical.EntityData.Leafs = types.NewOrderedMap()
    statisticsOnDemandHistorical.EntityData.Leafs.Append("operation-id", types.YLeaf{"OperationId", statisticsOnDemandHistorical.OperationId})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", statisticsOnDemandHistorical.DomainName})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statisticsOnDemandHistorical.InterfaceName})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", statisticsOnDemandHistorical.MepId})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", statisticsOnDemandHistorical.MacAddress})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", statisticsOnDemandHistorical.ProbeType})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", statisticsOnDemandHistorical.DisplayShort})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", statisticsOnDemandHistorical.DisplayLong})
    statisticsOnDemandHistorical.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", statisticsOnDemandHistorical.FlrCalculationInterval})

    statisticsOnDemandHistorical.EntityData.YListKeys = []string {}

    return &(statisticsOnDemandHistorical.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "statistics-on-demand-historical"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "statistics-on-demand-historical"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket.
    Bucket []*Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "statistics-on-demand-historical"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("config", types.YChild{"Config", &operationMetric.Config})
    operationMetric.EntityData.Children.Append("bucket", types.YChild{"Bucket", nil})
    for i := range operationMetric.Bucket {
        operationMetric.EntityData.Children.Append(types.GetSegmentPath(operationMetric.Bucket[i]), types.YChild{"Bucket", operationMetric.Bucket[i]})
    }
    operationMetric.EntityData.Leafs = types.NewOrderedMap()

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "operation-metric"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", config.MetricType})
    config.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", config.BinsCount})
    config.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", config.BinsWidth})
    config.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", config.BucketSize})
    config.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", config.BucketSizeUnit})
    config.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", config.BucketsArchive})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute time that the bucket started being filled at. The type is
    // interface{} with range: 0..4294967295.
    StartAt interface{}

    // Length of time for which the bucket is being filled in seconds. The type is
    // interface{} with range: 0..4294967295. Units are second.
    Duration interface{}

    // Number of packets sent in the probe. The type is interface{} with range:
    // 0..4294967295.
    Sent interface{}

    // Number of lost packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Lost interface{}

    // Number of corrupt packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Corrupt interface{}

    // Number of packets recieved out-of-order in the probe. The type is
    // interface{} with range: 0..4294967295.
    OutOfOrder interface{}

    // Number of duplicate packets received in the probe. The type is interface{}
    // with range: 0..4294967295.
    Duplicates interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Minimum interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Maximum interface{}

    // Absolute time that the minimum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMinimum interface{}

    // Absolute time that the maximum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMaximum interface{}

    // Mean of the results in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Average interface{}

    // Standard deviation of the results in the probe, in microseconds or
    // millionths of a percent. The type is interface{} with range:
    // -2147483648..2147483647.
    StandardDeviation interface{}

    // The count of samples collected in the bucket. The type is interface{} with
    // range: 0..4294967295.
    ResultCount interface{}

    // The number of data packets sent across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataSentCount interface{}

    // The number of data packets lost across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataLostCount interface{}

    // Frame Loss Ratio across the whole bucket, in millionths of a percent. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // percentage.
    OverallFlr interface{}

    // Results suspect due to a probe starting mid-way through a bucket. The type
    // is bool.
    SuspectStartMidBucket interface{}

    // Results suspect due to scheduling latency causing one or more packets to
    // not be sent. The type is bool.
    SuspectScheduleLatency interface{}

    // Results suspect due to failure to send one or more packets. The type is
    // bool.
    SuspectSendFail interface{}

    // Results suspect due to a probe ending prematurely. The type is bool.
    SuspectPrematureEnd interface{}

    // Results suspect as more than 10 seconds time drift detected. The type is
    // bool.
    SuspectClockDrift interface{}

    // Results suspect due to a memory allocation failure. The type is bool.
    SuspectMemoryAllocationFailed interface{}

    // Results suspect as bucket was cleared mid-way through being filled. The
    // type is bool.
    SuspectClearedMidBucket interface{}

    // Results suspect as probe restarted mid-way through the bucket. The type is
    // bool.
    SuspectProbeRestarted interface{}

    // Results suspect as processing of results has been delayed. The type is
    // bool.
    SuspectManagementLatency interface{}

    // Results suspect as the probe has been configured across multiple buckets.
    // The type is bool.
    SuspectMultipleBuckets interface{}

    // Results suspect as misordering has been detected , affecting results. The
    // type is bool.
    SuspectMisordering interface{}

    // Results suspect as FLR calculated based on a low packet count. The type is
    // bool.
    SuspectFlrLowPacketCount interface{}

    // If the probe ended prematurely, the error that caused a probe to end. The
    // type is interface{} with range: 0..4294967295.
    PrematureReason interface{}

    // Description of the error code that caused the probe to end prematurely. For
    // informational purposes only. The type is string.
    PrematureReasonString interface{}

    // The contents of the bucket; bins or samples.
    Contents Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "operation-metric"
    bucket.EntityData.SegmentPath = "bucket"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = types.NewOrderedMap()
    bucket.EntityData.Children.Append("contents", types.YChild{"Contents", &bucket.Contents})
    bucket.EntityData.Leafs = types.NewOrderedMap()
    bucket.EntityData.Leafs.Append("start-at", types.YLeaf{"StartAt", bucket.StartAt})
    bucket.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", bucket.Duration})
    bucket.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", bucket.Sent})
    bucket.EntityData.Leafs.Append("lost", types.YLeaf{"Lost", bucket.Lost})
    bucket.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", bucket.Corrupt})
    bucket.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", bucket.OutOfOrder})
    bucket.EntityData.Leafs.Append("duplicates", types.YLeaf{"Duplicates", bucket.Duplicates})
    bucket.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", bucket.Minimum})
    bucket.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", bucket.Maximum})
    bucket.EntityData.Leafs.Append("time-of-minimum", types.YLeaf{"TimeOfMinimum", bucket.TimeOfMinimum})
    bucket.EntityData.Leafs.Append("time-of-maximum", types.YLeaf{"TimeOfMaximum", bucket.TimeOfMaximum})
    bucket.EntityData.Leafs.Append("average", types.YLeaf{"Average", bucket.Average})
    bucket.EntityData.Leafs.Append("standard-deviation", types.YLeaf{"StandardDeviation", bucket.StandardDeviation})
    bucket.EntityData.Leafs.Append("result-count", types.YLeaf{"ResultCount", bucket.ResultCount})
    bucket.EntityData.Leafs.Append("data-sent-count", types.YLeaf{"DataSentCount", bucket.DataSentCount})
    bucket.EntityData.Leafs.Append("data-lost-count", types.YLeaf{"DataLostCount", bucket.DataLostCount})
    bucket.EntityData.Leafs.Append("overall-flr", types.YLeaf{"OverallFlr", bucket.OverallFlr})
    bucket.EntityData.Leafs.Append("suspect-start-mid-bucket", types.YLeaf{"SuspectStartMidBucket", bucket.SuspectStartMidBucket})
    bucket.EntityData.Leafs.Append("suspect-schedule-latency", types.YLeaf{"SuspectScheduleLatency", bucket.SuspectScheduleLatency})
    bucket.EntityData.Leafs.Append("suspect-send-fail", types.YLeaf{"SuspectSendFail", bucket.SuspectSendFail})
    bucket.EntityData.Leafs.Append("suspect-premature-end", types.YLeaf{"SuspectPrematureEnd", bucket.SuspectPrematureEnd})
    bucket.EntityData.Leafs.Append("suspect-clock-drift", types.YLeaf{"SuspectClockDrift", bucket.SuspectClockDrift})
    bucket.EntityData.Leafs.Append("suspect-memory-allocation-failed", types.YLeaf{"SuspectMemoryAllocationFailed", bucket.SuspectMemoryAllocationFailed})
    bucket.EntityData.Leafs.Append("suspect-cleared-mid-bucket", types.YLeaf{"SuspectClearedMidBucket", bucket.SuspectClearedMidBucket})
    bucket.EntityData.Leafs.Append("suspect-probe-restarted", types.YLeaf{"SuspectProbeRestarted", bucket.SuspectProbeRestarted})
    bucket.EntityData.Leafs.Append("suspect-management-latency", types.YLeaf{"SuspectManagementLatency", bucket.SuspectManagementLatency})
    bucket.EntityData.Leafs.Append("suspect-multiple-buckets", types.YLeaf{"SuspectMultipleBuckets", bucket.SuspectMultipleBuckets})
    bucket.EntityData.Leafs.Append("suspect-misordering", types.YLeaf{"SuspectMisordering", bucket.SuspectMisordering})
    bucket.EntityData.Leafs.Append("suspect-flr-low-packet-count", types.YLeaf{"SuspectFlrLowPacketCount", bucket.SuspectFlrLowPacketCount})
    bucket.EntityData.Leafs.Append("premature-reason", types.YLeaf{"PrematureReason", bucket.PrematureReason})
    bucket.EntityData.Leafs.Append("premature-reason-string", types.YLeaf{"PrematureReasonString", bucket.PrematureReasonString})

    bucket.EntityData.YListKeys = []string {}

    return &(bucket.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetEntityData() *types.CommonEntityData {
    contents.EntityData.YFilter = contents.YFilter
    contents.EntityData.YangName = "contents"
    contents.EntityData.BundleName = "cisco_ios_xr"
    contents.EntityData.ParentYangName = "bucket"
    contents.EntityData.SegmentPath = "contents"
    contents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contents.EntityData.Children = types.NewOrderedMap()
    contents.EntityData.Children.Append("aggregated", types.YChild{"Aggregated", &contents.Aggregated})
    contents.EntityData.Children.Append("unaggregated", types.YChild{"Unaggregated", &contents.Unaggregated})
    contents.EntityData.Leafs = types.NewOrderedMap()
    contents.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", contents.BucketType})

    contents.EntityData.YListKeys = []string {}

    return &(contents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []*Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "contents"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = types.NewOrderedMap()
    aggregated.EntityData.Children.Append("bins", types.YChild{"Bins", nil})
    for i := range aggregated.Bins {
        aggregated.EntityData.Children.Append(types.GetSegmentPath(aggregated.Bins[i]), types.YChild{"Bins", aggregated.Bins[i]})
    }
    aggregated.EntityData.Leafs = types.NewOrderedMap()

    aggregated.EntityData.YListKeys = []string {}

    return &(aggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower bound (inclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    LowerBound interface{}

    // Upper bound (exclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    UpperBound interface{}

    // Lower bound (inclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    LowerBoundTenths interface{}

    // Upper bound (exclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    UpperBoundTenths interface{}

    // The sum of the results in the bin, in microseconds or millionths of a
    // percent. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Sum interface{}

    // The total number of results in the bin. The type is interface{} with range:
    // 0..4294967295.
    Count interface{}
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetEntityData() *types.CommonEntityData {
    bins.EntityData.YFilter = bins.YFilter
    bins.EntityData.YangName = "bins"
    bins.EntityData.BundleName = "cisco_ios_xr"
    bins.EntityData.ParentYangName = "aggregated"
    bins.EntityData.SegmentPath = "bins"
    bins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bins.EntityData.Children = types.NewOrderedMap()
    bins.EntityData.Leafs = types.NewOrderedMap()
    bins.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", bins.LowerBound})
    bins.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", bins.UpperBound})
    bins.EntityData.Leafs.Append("lower-bound-tenths", types.YLeaf{"LowerBoundTenths", bins.LowerBoundTenths})
    bins.EntityData.Leafs.Append("upper-bound-tenths", types.YLeaf{"UpperBoundTenths", bins.UpperBoundTenths})
    bins.EntityData.Leafs.Append("sum", types.YLeaf{"Sum", bins.Sum})
    bins.EntityData.Leafs.Append("count", types.YLeaf{"Count", bins.Count})

    bins.EntityData.YListKeys = []string {}

    return &(bins.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []*Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetEntityData() *types.CommonEntityData {
    unaggregated.EntityData.YFilter = unaggregated.YFilter
    unaggregated.EntityData.YangName = "unaggregated"
    unaggregated.EntityData.BundleName = "cisco_ios_xr"
    unaggregated.EntityData.ParentYangName = "contents"
    unaggregated.EntityData.SegmentPath = "unaggregated"
    unaggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unaggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unaggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unaggregated.EntityData.Children = types.NewOrderedMap()
    unaggregated.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range unaggregated.Sample {
        unaggregated.EntityData.Children.Append(types.GetSegmentPath(unaggregated.Sample[i]), types.YChild{"Sample", unaggregated.Sample[i]})
    }
    unaggregated.EntityData.Leafs = types.NewOrderedMap()

    unaggregated.EntityData.YListKeys = []string {}

    return &(unaggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time (in milliseconds relative to the start time of the bucket) that
    // the sample was sent at. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    SentAt interface{}

    // Whether the sample packet was sucessfully sent. The type is bool.
    Sent interface{}

    // Whether the sample packet timed out. The type is bool.
    TimedOut interface{}

    // Whether the sample packet was corrupt. The type is bool.
    Corrupt interface{}

    // Whether the sample packet was received out-of-order. The type is bool.
    OutOfOrder interface{}

    // Whether a measurement could not be made because no data packets were sent
    // in the sample period. Only applicable for LMM measurements. The type is
    // bool.
    NoDataPackets interface{}

    // The result (in microseconds or millionths of a percent) of the sample, if
    // available. The type is interface{} with range: -2147483648..2147483647.
    Result interface{}

    // For FLR measurements, the number of frames sent, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesSent interface{}

    // For FLR measurements, the number of frames lost, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesLost interface{}
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "unaggregated"
    sample.EntityData.SegmentPath = "sample"
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sent-at", types.YLeaf{"SentAt", sample.SentAt})
    sample.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sample.Sent})
    sample.EntityData.Leafs.Append("timed-out", types.YLeaf{"TimedOut", sample.TimedOut})
    sample.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", sample.Corrupt})
    sample.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", sample.OutOfOrder})
    sample.EntityData.Leafs.Append("no-data-packets", types.YLeaf{"NoDataPackets", sample.NoDataPackets})
    sample.EntityData.Leafs.Append("result", types.YLeaf{"Result", sample.Result})
    sample.EntityData.Leafs.Append("frames-sent", types.YLeaf{"FramesSent", sample.FramesSent})
    sample.EntityData.Leafs.Append("frames-lost", types.YLeaf{"FramesLost", sample.FramesLost})

    sample.EntityData.YListKeys = []string {}

    return &(sample.EntityData)
}

// Sla_Protocols_Ethernet_ConfigErrors
// Table of SLA configuration errors on configured
// operations
type Sla_Protocols_Ethernet_ConfigErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SLA operation to get configuration errors data for. The type is slice of
    // Sla_Protocols_Ethernet_ConfigErrors_ConfigError.
    ConfigError []*Sla_Protocols_Ethernet_ConfigErrors_ConfigError
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetEntityData() *types.CommonEntityData {
    configErrors.EntityData.YFilter = configErrors.YFilter
    configErrors.EntityData.YangName = "config-errors"
    configErrors.EntityData.BundleName = "cisco_ios_xr"
    configErrors.EntityData.ParentYangName = "ethernet"
    configErrors.EntityData.SegmentPath = "config-errors"
    configErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configErrors.EntityData.Children = types.NewOrderedMap()
    configErrors.EntityData.Children.Append("config-error", types.YChild{"ConfigError", nil})
    for i := range configErrors.ConfigError {
        configErrors.EntityData.Children.Append(types.GetSegmentPath(configErrors.ConfigError[i]), types.YChild{"ConfigError", configErrors.ConfigError[i]})
    }
    configErrors.EntityData.Leafs = types.NewOrderedMap()

    configErrors.EntityData.YListKeys = []string {}

    return &(configErrors.EntityData)
}

// Sla_Protocols_Ethernet_ConfigErrors_ConfigError
// SLA operation to get configuration errors data
// for
type Sla_Protocols_Ethernet_ConfigErrors_ConfigError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // The name of the operation profile. The type is string.
    ProfileNameXr interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Is the profile configured to collect RT Delay but the packet type doesn't
    // support it?. The type is bool.
    RtDelayInconsistent interface{}

    // Is the profile configured to collect OW Delay (SD) but the packet type
    // doesn't support it?. The type is bool.
    OwDelaySdInconsistent interface{}

    // Is the profile configured to collect OW Delay (DS) but the packet type
    // doesn't support it?. The type is bool.
    OwDelayDsInconsistent interface{}

    // Is the profile configured to collect RT Jitter but the packet type doesn't
    // support it?. The type is bool.
    RtJitterInconsistent interface{}

    // Is the profile configured to collect OW Jitter (SD) but the packet type
    // doesn't support it?. The type is bool.
    OwJitterSdInconsistent interface{}

    // Is the profile configured to collect OW Delay (DS) but the packet type
    // doesn't support it?. The type is bool.
    OwJitterDsInconsistent interface{}

    // Is the profile configured to collect OW Frame Loss (SD) but the packet type
    // doesn't support it ?. The type is bool.
    OwLossSdInconsistent interface{}

    // Is the profile configured to collect OW Frame Loss (DS) but the packet type
    // doesn't support it ?. The type is bool.
    OwLossDsInconsistent interface{}

    // Is the profile configured to pad packets but the packet type doesn't
    // support it?. The type is bool.
    PacketPadInconsistent interface{}

    // Is the profile configured to pad packets with a pseudo-random string but
    // the packet type doesn't support it?. The type is bool.
    PacketRandPadInconsistent interface{}

    // Is the profile configured to send packets more frequently than the protocol
    // allows?. The type is bool.
    MinPacketIntervalInconsistent interface{}

    // Is the profile configured to use a packet priority scheme that the protocol
    // does not support?. The type is bool.
    PriorityInconsistent interface{}

    // Is the profile configured to use a packet type that isn't supported by any
    // protocols?. The type is bool.
    PacketTypeInconsistent interface{}

    // Is the operation configured to use a profile that is not currently defined
    // for the protocol?. The type is bool.
    ProfileDoesntExist interface{}

    // The profile is configured to use a packet type which doesn't support
    // synthetic loss measurement and the number of packets per FLR calculation
    // has been configured. The type is bool.
    SyntheticLossNotSupported interface{}

    // The profile is configured to use a packet type which does not allow more
    // than 72000 packets per probe and greater than 72000 packets per probe have
    // been configured. The type is bool.
    ProbeTooBig interface{}

    // Displays other issues not indicated from the flags above, for example MIB
    // incompatibility issues. The type is slice of string.
    ErrorString []interface{}
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetEntityData() *types.CommonEntityData {
    configError.EntityData.YFilter = configError.YFilter
    configError.EntityData.YangName = "config-error"
    configError.EntityData.BundleName = "cisco_ios_xr"
    configError.EntityData.ParentYangName = "config-errors"
    configError.EntityData.SegmentPath = "config-error"
    configError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configError.EntityData.Children = types.NewOrderedMap()
    configError.EntityData.Leafs = types.NewOrderedMap()
    configError.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configError.ProfileName})
    configError.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", configError.DomainName})
    configError.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", configError.InterfaceName})
    configError.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", configError.MepId})
    configError.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", configError.MacAddress})
    configError.EntityData.Leafs.Append("profile-name-xr", types.YLeaf{"ProfileNameXr", configError.ProfileNameXr})
    configError.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", configError.DisplayShort})
    configError.EntityData.Leafs.Append("rt-delay-inconsistent", types.YLeaf{"RtDelayInconsistent", configError.RtDelayInconsistent})
    configError.EntityData.Leafs.Append("ow-delay-sd-inconsistent", types.YLeaf{"OwDelaySdInconsistent", configError.OwDelaySdInconsistent})
    configError.EntityData.Leafs.Append("ow-delay-ds-inconsistent", types.YLeaf{"OwDelayDsInconsistent", configError.OwDelayDsInconsistent})
    configError.EntityData.Leafs.Append("rt-jitter-inconsistent", types.YLeaf{"RtJitterInconsistent", configError.RtJitterInconsistent})
    configError.EntityData.Leafs.Append("ow-jitter-sd-inconsistent", types.YLeaf{"OwJitterSdInconsistent", configError.OwJitterSdInconsistent})
    configError.EntityData.Leafs.Append("ow-jitter-ds-inconsistent", types.YLeaf{"OwJitterDsInconsistent", configError.OwJitterDsInconsistent})
    configError.EntityData.Leafs.Append("ow-loss-sd-inconsistent", types.YLeaf{"OwLossSdInconsistent", configError.OwLossSdInconsistent})
    configError.EntityData.Leafs.Append("ow-loss-ds-inconsistent", types.YLeaf{"OwLossDsInconsistent", configError.OwLossDsInconsistent})
    configError.EntityData.Leafs.Append("packet-pad-inconsistent", types.YLeaf{"PacketPadInconsistent", configError.PacketPadInconsistent})
    configError.EntityData.Leafs.Append("packet-rand-pad-inconsistent", types.YLeaf{"PacketRandPadInconsistent", configError.PacketRandPadInconsistent})
    configError.EntityData.Leafs.Append("min-packet-interval-inconsistent", types.YLeaf{"MinPacketIntervalInconsistent", configError.MinPacketIntervalInconsistent})
    configError.EntityData.Leafs.Append("priority-inconsistent", types.YLeaf{"PriorityInconsistent", configError.PriorityInconsistent})
    configError.EntityData.Leafs.Append("packet-type-inconsistent", types.YLeaf{"PacketTypeInconsistent", configError.PacketTypeInconsistent})
    configError.EntityData.Leafs.Append("profile-doesnt-exist", types.YLeaf{"ProfileDoesntExist", configError.ProfileDoesntExist})
    configError.EntityData.Leafs.Append("synthetic-loss-not-supported", types.YLeaf{"SyntheticLossNotSupported", configError.SyntheticLossNotSupported})
    configError.EntityData.Leafs.Append("probe-too-big", types.YLeaf{"ProbeTooBig", configError.ProbeTooBig})
    configError.EntityData.Leafs.Append("error-string", types.YLeaf{"ErrorString", configError.ErrorString})

    configError.EntityData.YListKeys = []string {}

    return &(configError.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations
// Table of SLA on-demand operations
type Sla_Protocols_Ethernet_OnDemandOperations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SLA on-demand operation to get operation data for. The type is slice of
    // Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation.
    OnDemandOperation []*Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetEntityData() *types.CommonEntityData {
    onDemandOperations.EntityData.YFilter = onDemandOperations.YFilter
    onDemandOperations.EntityData.YangName = "on-demand-operations"
    onDemandOperations.EntityData.BundleName = "cisco_ios_xr"
    onDemandOperations.EntityData.ParentYangName = "ethernet"
    onDemandOperations.EntityData.SegmentPath = "on-demand-operations"
    onDemandOperations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandOperations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandOperations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandOperations.EntityData.Children = types.NewOrderedMap()
    onDemandOperations.EntityData.Children.Append("on-demand-operation", types.YChild{"OnDemandOperation", nil})
    for i := range onDemandOperations.OnDemandOperation {
        onDemandOperations.EntityData.Children.Append(types.GetSegmentPath(onDemandOperations.OnDemandOperation[i]), types.YChild{"OnDemandOperation", onDemandOperations.OnDemandOperation[i]})
    }
    onDemandOperations.EntityData.Leafs = types.NewOrderedMap()

    onDemandOperations.EntityData.YListKeys = []string {}

    return &(onDemandOperations.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation
// SLA on-demand operation to get operation data
// for
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operation ID. The type is interface{} with range: 1..4294967295.
    OperationId interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Time that the last probe for the operation was run, NULL if never run. The
    // type is interface{} with range: 0..4294967295.
    LastRun interface{}

    // Options that are only valid if the operation has a profile.
    ProfileOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetEntityData() *types.CommonEntityData {
    onDemandOperation.EntityData.YFilter = onDemandOperation.YFilter
    onDemandOperation.EntityData.YangName = "on-demand-operation"
    onDemandOperation.EntityData.BundleName = "cisco_ios_xr"
    onDemandOperation.EntityData.ParentYangName = "on-demand-operations"
    onDemandOperation.EntityData.SegmentPath = "on-demand-operation"
    onDemandOperation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandOperation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandOperation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandOperation.EntityData.Children = types.NewOrderedMap()
    onDemandOperation.EntityData.Children.Append("profile-options", types.YChild{"ProfileOptions", &onDemandOperation.ProfileOptions})
    onDemandOperation.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &onDemandOperation.SpecificOptions})
    onDemandOperation.EntityData.Leafs = types.NewOrderedMap()
    onDemandOperation.EntityData.Leafs.Append("operation-id", types.YLeaf{"OperationId", onDemandOperation.OperationId})
    onDemandOperation.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", onDemandOperation.DomainName})
    onDemandOperation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", onDemandOperation.InterfaceName})
    onDemandOperation.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", onDemandOperation.MepId})
    onDemandOperation.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", onDemandOperation.MacAddress})
    onDemandOperation.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", onDemandOperation.DisplayShort})
    onDemandOperation.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", onDemandOperation.DisplayLong})
    onDemandOperation.EntityData.Leafs.Append("last-run", types.YLeaf{"LastRun", onDemandOperation.LastRun})

    onDemandOperation.EntityData.YListKeys = []string {}

    return &(onDemandOperation.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions
// Options that are only valid if the operation has
// a profile
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Number of packets sent per burst. The type is interface{} with range:
    // 0..65535.
    PacketsPerBurst interface{}

    // Interval between packets within a burst in milliseconds. The type is
    // interface{} with range: 0..65535. Units are millisecond.
    InterPacketInterval interface{}

    // Number of bursts sent per probe. The type is interface{} with range:
    // 0..4294967295.
    BurstsPerProbe interface{}

    // Interval between bursts within a probe in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    InterBurstInterval interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Configuration of the packet padding.
    PacketPadding Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding

    // Priority at which to send the packet, if configured.
    Priority Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule

    // Array of the metrics that are measured by the operation. The type is slice
    // of
    // Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetEntityData() *types.CommonEntityData {
    profileOptions.EntityData.YFilter = profileOptions.YFilter
    profileOptions.EntityData.YangName = "profile-options"
    profileOptions.EntityData.BundleName = "cisco_ios_xr"
    profileOptions.EntityData.ParentYangName = "on-demand-operation"
    profileOptions.EntityData.SegmentPath = "profile-options"
    profileOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileOptions.EntityData.Children = types.NewOrderedMap()
    profileOptions.EntityData.Children.Append("packet-padding", types.YChild{"PacketPadding", &profileOptions.PacketPadding})
    profileOptions.EntityData.Children.Append("priority", types.YChild{"Priority", &profileOptions.Priority})
    profileOptions.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &profileOptions.OperationSchedule})
    profileOptions.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range profileOptions.OperationMetric {
        profileOptions.EntityData.Children.Append(types.GetSegmentPath(profileOptions.OperationMetric[i]), types.YChild{"OperationMetric", profileOptions.OperationMetric[i]})
    }
    profileOptions.EntityData.Leafs = types.NewOrderedMap()
    profileOptions.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", profileOptions.ProbeType})
    profileOptions.EntityData.Leafs.Append("packets-per-burst", types.YLeaf{"PacketsPerBurst", profileOptions.PacketsPerBurst})
    profileOptions.EntityData.Leafs.Append("inter-packet-interval", types.YLeaf{"InterPacketInterval", profileOptions.InterPacketInterval})
    profileOptions.EntityData.Leafs.Append("bursts-per-probe", types.YLeaf{"BurstsPerProbe", profileOptions.BurstsPerProbe})
    profileOptions.EntityData.Leafs.Append("inter-burst-interval", types.YLeaf{"InterBurstInterval", profileOptions.InterBurstInterval})
    profileOptions.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", profileOptions.FlrCalculationInterval})

    profileOptions.EntityData.YListKeys = []string {}

    return &(profileOptions.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding
// Configuration of the packet padding
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size that packets are being padded to. The type is interface{} with range:
    // 0..65535.
    PacketPadSize interface{}

    // Test pattern scheme that is used in the packet padding. The type is
    // SlaOperTestPatternScheme.
    TestPatternPadScheme interface{}

    // Hex string that is used in the packet padding. The type is interface{} with
    // range: 0..4294967295.
    TestPatternPadHexString interface{}
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetEntityData() *types.CommonEntityData {
    packetPadding.EntityData.YFilter = packetPadding.YFilter
    packetPadding.EntityData.YangName = "packet-padding"
    packetPadding.EntityData.BundleName = "cisco_ios_xr"
    packetPadding.EntityData.ParentYangName = "profile-options"
    packetPadding.EntityData.SegmentPath = "packet-padding"
    packetPadding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetPadding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetPadding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetPadding.EntityData.Children = types.NewOrderedMap()
    packetPadding.EntityData.Leafs = types.NewOrderedMap()
    packetPadding.EntityData.Leafs.Append("packet-pad-size", types.YLeaf{"PacketPadSize", packetPadding.PacketPadSize})
    packetPadding.EntityData.Leafs.Append("test-pattern-pad-scheme", types.YLeaf{"TestPatternPadScheme", packetPadding.TestPatternPadScheme})
    packetPadding.EntityData.Leafs.Append("test-pattern-pad-hex-string", types.YLeaf{"TestPatternPadHexString", packetPadding.TestPatternPadHexString})

    packetPadding.EntityData.YListKeys = []string {}

    return &(packetPadding.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority
// Priority at which to send the packet, if
// configured
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PriorityType. The type is SlaOperPacketPriority.
    PriorityType interface{}

    // 3-bit COS priority value applied to packets. The type is interface{} with
    // range: 0..255.
    Cos interface{}
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "profile-options"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("priority-type", types.YLeaf{"PriorityType", priority.PriorityType})
    priority.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", priority.Cos})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "profile-options"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric
// Array of the metrics that are measured by the
// operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of valid buckets currently in the buckets archive. The type is
    // interface{} with range: 0..4294967295.
    CurrentBucketsArchive interface{}

    // Configuration of the metric.
    MetricConfig Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "profile-options"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("metric-config", types.YChild{"MetricConfig", &operationMetric.MetricConfig})
    operationMetric.EntityData.Leafs = types.NewOrderedMap()
    operationMetric.EntityData.Leafs.Append("current-buckets-archive", types.YLeaf{"CurrentBucketsArchive", operationMetric.CurrentBucketsArchive})

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig
// Configuration of the metric
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetEntityData() *types.CommonEntityData {
    metricConfig.EntityData.YFilter = metricConfig.YFilter
    metricConfig.EntityData.YangName = "metric-config"
    metricConfig.EntityData.BundleName = "cisco_ios_xr"
    metricConfig.EntityData.ParentYangName = "operation-metric"
    metricConfig.EntityData.SegmentPath = "metric-config"
    metricConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricConfig.EntityData.Children = types.NewOrderedMap()
    metricConfig.EntityData.Leafs = types.NewOrderedMap()
    metricConfig.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", metricConfig.MetricType})
    metricConfig.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", metricConfig.BinsCount})
    metricConfig.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", metricConfig.BinsWidth})
    metricConfig.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", metricConfig.BucketSize})
    metricConfig.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", metricConfig.BucketSizeUnit})
    metricConfig.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", metricConfig.BucketsArchive})

    metricConfig.EntityData.YListKeys = []string {}

    return &(metricConfig.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "on-demand-operation"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents
// Table of current statistics for SLA operations
type Sla_Protocols_Ethernet_StatisticsCurrents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current statistics data for an SLA configured operation. The type is slice
    // of Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent.
    StatisticsCurrent []*Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetEntityData() *types.CommonEntityData {
    statisticsCurrents.EntityData.YFilter = statisticsCurrents.YFilter
    statisticsCurrents.EntityData.YangName = "statistics-currents"
    statisticsCurrents.EntityData.BundleName = "cisco_ios_xr"
    statisticsCurrents.EntityData.ParentYangName = "ethernet"
    statisticsCurrents.EntityData.SegmentPath = "statistics-currents"
    statisticsCurrents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsCurrents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsCurrents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsCurrents.EntityData.Children = types.NewOrderedMap()
    statisticsCurrents.EntityData.Children.Append("statistics-current", types.YChild{"StatisticsCurrent", nil})
    for i := range statisticsCurrents.StatisticsCurrent {
        statisticsCurrents.EntityData.Children.Append(types.GetSegmentPath(statisticsCurrents.StatisticsCurrent[i]), types.YChild{"StatisticsCurrent", statisticsCurrents.StatisticsCurrent[i]})
    }
    statisticsCurrents.EntityData.Leafs = types.NewOrderedMap()

    statisticsCurrents.EntityData.YListKeys = []string {}

    return &(statisticsCurrents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent
// Current statistics data for an SLA configured
// operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MEP ID in the range 1 to 8191. Either MEP ID or MAC address must be
    // specified. The type is interface{} with range: 1..8191.
    MepId interface{}

    // Unicast MAC Address in xxxx.xxxx.xxxx format. Either MEP ID or MAC address
    // must be specified. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Type of probe used by the operation. The type is string.
    ProbeType interface{}

    // Short display name used by the operation. The type is string.
    DisplayShort interface{}

    // Long display name used by the operation. The type is string.
    DisplayLong interface{}

    // Interval between FLR calculations for SLM, in milliseconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    FlrCalculationInterval interface{}

    // Options specific to the type of operation.
    SpecificOptions Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions

    // Operation schedule.
    OperationSchedule Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule

    // Metrics gathered for the operation. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric.
    OperationMetric []*Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetEntityData() *types.CommonEntityData {
    statisticsCurrent.EntityData.YFilter = statisticsCurrent.YFilter
    statisticsCurrent.EntityData.YangName = "statistics-current"
    statisticsCurrent.EntityData.BundleName = "cisco_ios_xr"
    statisticsCurrent.EntityData.ParentYangName = "statistics-currents"
    statisticsCurrent.EntityData.SegmentPath = "statistics-current"
    statisticsCurrent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsCurrent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsCurrent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsCurrent.EntityData.Children = types.NewOrderedMap()
    statisticsCurrent.EntityData.Children.Append("specific-options", types.YChild{"SpecificOptions", &statisticsCurrent.SpecificOptions})
    statisticsCurrent.EntityData.Children.Append("operation-schedule", types.YChild{"OperationSchedule", &statisticsCurrent.OperationSchedule})
    statisticsCurrent.EntityData.Children.Append("operation-metric", types.YChild{"OperationMetric", nil})
    for i := range statisticsCurrent.OperationMetric {
        statisticsCurrent.EntityData.Children.Append(types.GetSegmentPath(statisticsCurrent.OperationMetric[i]), types.YChild{"OperationMetric", statisticsCurrent.OperationMetric[i]})
    }
    statisticsCurrent.EntityData.Leafs = types.NewOrderedMap()
    statisticsCurrent.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", statisticsCurrent.ProfileName})
    statisticsCurrent.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", statisticsCurrent.DomainName})
    statisticsCurrent.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", statisticsCurrent.InterfaceName})
    statisticsCurrent.EntityData.Leafs.Append("mep-id", types.YLeaf{"MepId", statisticsCurrent.MepId})
    statisticsCurrent.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", statisticsCurrent.MacAddress})
    statisticsCurrent.EntityData.Leafs.Append("probe-type", types.YLeaf{"ProbeType", statisticsCurrent.ProbeType})
    statisticsCurrent.EntityData.Leafs.Append("display-short", types.YLeaf{"DisplayShort", statisticsCurrent.DisplayShort})
    statisticsCurrent.EntityData.Leafs.Append("display-long", types.YLeaf{"DisplayLong", statisticsCurrent.DisplayLong})
    statisticsCurrent.EntityData.Leafs.Append("flr-calculation-interval", types.YLeaf{"FlrCalculationInterval", statisticsCurrent.FlrCalculationInterval})

    statisticsCurrent.EntityData.YListKeys = []string {}

    return &(statisticsCurrent.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetEntityData() *types.CommonEntityData {
    specificOptions.EntityData.YFilter = specificOptions.YFilter
    specificOptions.EntityData.YangName = "specific-options"
    specificOptions.EntityData.BundleName = "cisco_ios_xr"
    specificOptions.EntityData.ParentYangName = "statistics-current"
    specificOptions.EntityData.SegmentPath = "specific-options"
    specificOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    specificOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    specificOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    specificOptions.EntityData.Children = types.NewOrderedMap()
    specificOptions.EntityData.Children.Append("configured-operation-options", types.YChild{"ConfiguredOperationOptions", &specificOptions.ConfiguredOperationOptions})
    specificOptions.EntityData.Children.Append("ondemand-operation-options", types.YChild{"OndemandOperationOptions", &specificOptions.OndemandOperationOptions})
    specificOptions.EntityData.Leafs = types.NewOrderedMap()
    specificOptions.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", specificOptions.OperType})

    specificOptions.EntityData.YListKeys = []string {}

    return &(specificOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetEntityData() *types.CommonEntityData {
    configuredOperationOptions.EntityData.YFilter = configuredOperationOptions.YFilter
    configuredOperationOptions.EntityData.YangName = "configured-operation-options"
    configuredOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    configuredOperationOptions.EntityData.ParentYangName = "specific-options"
    configuredOperationOptions.EntityData.SegmentPath = "configured-operation-options"
    configuredOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredOperationOptions.EntityData.Children = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    configuredOperationOptions.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", configuredOperationOptions.ProfileName})

    configuredOperationOptions.EntityData.YListKeys = []string {}

    return &(configuredOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetEntityData() *types.CommonEntityData {
    ondemandOperationOptions.EntityData.YFilter = ondemandOperationOptions.YFilter
    ondemandOperationOptions.EntityData.YangName = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.BundleName = "cisco_ios_xr"
    ondemandOperationOptions.EntityData.ParentYangName = "specific-options"
    ondemandOperationOptions.EntityData.SegmentPath = "ondemand-operation-options"
    ondemandOperationOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ondemandOperationOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ondemandOperationOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ondemandOperationOptions.EntityData.Children = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs = types.NewOrderedMap()
    ondemandOperationOptions.EntityData.Leafs.Append("ondemand-operation-id", types.YLeaf{"OndemandOperationId", ondemandOperationOptions.OndemandOperationId})
    ondemandOperationOptions.EntityData.Leafs.Append("probe-count", types.YLeaf{"ProbeCount", ondemandOperationOptions.ProbeCount})

    ondemandOperationOptions.EntityData.YListKeys = []string {}

    return &(ondemandOperationOptions.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start time of the first probe, in seconds since the Unix Epoch. The type is
    // interface{} with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Whether or not the operation start time was explicitly configured. The type
    // is bool.
    StartTimeConfigured interface{}

    // Duration of a probe for the operation in seconds. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ScheduleDuration interface{}

    // Interval between the start times of consecutive probes,  in seconds. The
    // type is interface{} with range: 0..4294967295. Units are second.
    ScheduleInterval interface{}
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetEntityData() *types.CommonEntityData {
    operationSchedule.EntityData.YFilter = operationSchedule.YFilter
    operationSchedule.EntityData.YangName = "operation-schedule"
    operationSchedule.EntityData.BundleName = "cisco_ios_xr"
    operationSchedule.EntityData.ParentYangName = "statistics-current"
    operationSchedule.EntityData.SegmentPath = "operation-schedule"
    operationSchedule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationSchedule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationSchedule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationSchedule.EntityData.Children = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs = types.NewOrderedMap()
    operationSchedule.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", operationSchedule.StartTime})
    operationSchedule.EntityData.Leafs.Append("start-time-configured", types.YLeaf{"StartTimeConfigured", operationSchedule.StartTimeConfigured})
    operationSchedule.EntityData.Leafs.Append("schedule-duration", types.YLeaf{"ScheduleDuration", operationSchedule.ScheduleDuration})
    operationSchedule.EntityData.Leafs.Append("schedule-interval", types.YLeaf{"ScheduleInterval", operationSchedule.ScheduleInterval})

    operationSchedule.EntityData.YListKeys = []string {}

    return &(operationSchedule.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket.
    Bucket []*Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetEntityData() *types.CommonEntityData {
    operationMetric.EntityData.YFilter = operationMetric.YFilter
    operationMetric.EntityData.YangName = "operation-metric"
    operationMetric.EntityData.BundleName = "cisco_ios_xr"
    operationMetric.EntityData.ParentYangName = "statistics-current"
    operationMetric.EntityData.SegmentPath = "operation-metric"
    operationMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationMetric.EntityData.Children = types.NewOrderedMap()
    operationMetric.EntityData.Children.Append("config", types.YChild{"Config", &operationMetric.Config})
    operationMetric.EntityData.Children.Append("bucket", types.YChild{"Bucket", nil})
    for i := range operationMetric.Bucket {
        operationMetric.EntityData.Children.Append(types.GetSegmentPath(operationMetric.Bucket[i]), types.YChild{"Bucket", operationMetric.Bucket[i]})
    }
    operationMetric.EntityData.Leafs = types.NewOrderedMap()

    operationMetric.EntityData.YListKeys = []string {}

    return &(operationMetric.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of metric to which this configuration applies. The type is
    // SlaRecordableMetric.
    MetricType interface{}

    // Total number of bins into which to aggregate. 0 if no aggregation. The type
    // is interface{} with range: 0..65535.
    BinsCount interface{}

    // Width of each bin into which to aggregate. 0 if no aggregation. For SLM,
    // the units of this value are in single units of percent; for LMM they are in
    // tenths of percent; for other measurements they are in milliseconds. The
    // type is interface{} with range: 0..65535.
    BinsWidth interface{}

    // Size of buckets into which measurements are collected. The type is
    // interface{} with range: 0..255.
    BucketSize interface{}

    // Whether bucket size is 'per-probe' or 'probes'. The type is SlaBucketSize.
    BucketSizeUnit interface{}

    // Maximum number of buckets to store in memory. The type is interface{} with
    // range: 0..4294967295.
    BucketsArchive interface{}
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "operation-metric"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", config.MetricType})
    config.EntityData.Leafs.Append("bins-count", types.YLeaf{"BinsCount", config.BinsCount})
    config.EntityData.Leafs.Append("bins-width", types.YLeaf{"BinsWidth", config.BinsWidth})
    config.EntityData.Leafs.Append("bucket-size", types.YLeaf{"BucketSize", config.BucketSize})
    config.EntityData.Leafs.Append("bucket-size-unit", types.YLeaf{"BucketSizeUnit", config.BucketSizeUnit})
    config.EntityData.Leafs.Append("buckets-archive", types.YLeaf{"BucketsArchive", config.BucketsArchive})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute time that the bucket started being filled at. The type is
    // interface{} with range: 0..4294967295.
    StartAt interface{}

    // Length of time for which the bucket is being filled in seconds. The type is
    // interface{} with range: 0..4294967295. Units are second.
    Duration interface{}

    // Number of packets sent in the probe. The type is interface{} with range:
    // 0..4294967295.
    Sent interface{}

    // Number of lost packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Lost interface{}

    // Number of corrupt packets in the probe. The type is interface{} with range:
    // 0..4294967295.
    Corrupt interface{}

    // Number of packets recieved out-of-order in the probe. The type is
    // interface{} with range: 0..4294967295.
    OutOfOrder interface{}

    // Number of duplicate packets received in the probe. The type is interface{}
    // with range: 0..4294967295.
    Duplicates interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Minimum interface{}

    // Overall minimum result in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Maximum interface{}

    // Absolute time that the minimum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMinimum interface{}

    // Absolute time that the maximum value was recorded. The type is interface{}
    // with range: 0..4294967295.
    TimeOfMaximum interface{}

    // Mean of the results in the probe, in microseconds or millionths of a
    // percent. The type is interface{} with range: -2147483648..2147483647.
    Average interface{}

    // Standard deviation of the results in the probe, in microseconds or
    // millionths of a percent. The type is interface{} with range:
    // -2147483648..2147483647.
    StandardDeviation interface{}

    // The count of samples collected in the bucket. The type is interface{} with
    // range: 0..4294967295.
    ResultCount interface{}

    // The number of data packets sent across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataSentCount interface{}

    // The number of data packets lost across the bucket, used in the calculation
    // of overall FLR. The type is interface{} with range: 0..4294967295.
    DataLostCount interface{}

    // Frame Loss Ratio across the whole bucket, in millionths of a percent. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // percentage.
    OverallFlr interface{}

    // Results suspect due to a probe starting mid-way through a bucket. The type
    // is bool.
    SuspectStartMidBucket interface{}

    // Results suspect due to scheduling latency causing one or more packets to
    // not be sent. The type is bool.
    SuspectScheduleLatency interface{}

    // Results suspect due to failure to send one or more packets. The type is
    // bool.
    SuspectSendFail interface{}

    // Results suspect due to a probe ending prematurely. The type is bool.
    SuspectPrematureEnd interface{}

    // Results suspect as more than 10 seconds time drift detected. The type is
    // bool.
    SuspectClockDrift interface{}

    // Results suspect due to a memory allocation failure. The type is bool.
    SuspectMemoryAllocationFailed interface{}

    // Results suspect as bucket was cleared mid-way through being filled. The
    // type is bool.
    SuspectClearedMidBucket interface{}

    // Results suspect as probe restarted mid-way through the bucket. The type is
    // bool.
    SuspectProbeRestarted interface{}

    // Results suspect as processing of results has been delayed. The type is
    // bool.
    SuspectManagementLatency interface{}

    // Results suspect as the probe has been configured across multiple buckets.
    // The type is bool.
    SuspectMultipleBuckets interface{}

    // Results suspect as misordering has been detected , affecting results. The
    // type is bool.
    SuspectMisordering interface{}

    // Results suspect as FLR calculated based on a low packet count. The type is
    // bool.
    SuspectFlrLowPacketCount interface{}

    // If the probe ended prematurely, the error that caused a probe to end. The
    // type is interface{} with range: 0..4294967295.
    PrematureReason interface{}

    // Description of the error code that caused the probe to end prematurely. For
    // informational purposes only. The type is string.
    PrematureReasonString interface{}

    // The contents of the bucket; bins or samples.
    Contents Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetEntityData() *types.CommonEntityData {
    bucket.EntityData.YFilter = bucket.YFilter
    bucket.EntityData.YangName = "bucket"
    bucket.EntityData.BundleName = "cisco_ios_xr"
    bucket.EntityData.ParentYangName = "operation-metric"
    bucket.EntityData.SegmentPath = "bucket"
    bucket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bucket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bucket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bucket.EntityData.Children = types.NewOrderedMap()
    bucket.EntityData.Children.Append("contents", types.YChild{"Contents", &bucket.Contents})
    bucket.EntityData.Leafs = types.NewOrderedMap()
    bucket.EntityData.Leafs.Append("start-at", types.YLeaf{"StartAt", bucket.StartAt})
    bucket.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", bucket.Duration})
    bucket.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", bucket.Sent})
    bucket.EntityData.Leafs.Append("lost", types.YLeaf{"Lost", bucket.Lost})
    bucket.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", bucket.Corrupt})
    bucket.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", bucket.OutOfOrder})
    bucket.EntityData.Leafs.Append("duplicates", types.YLeaf{"Duplicates", bucket.Duplicates})
    bucket.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", bucket.Minimum})
    bucket.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", bucket.Maximum})
    bucket.EntityData.Leafs.Append("time-of-minimum", types.YLeaf{"TimeOfMinimum", bucket.TimeOfMinimum})
    bucket.EntityData.Leafs.Append("time-of-maximum", types.YLeaf{"TimeOfMaximum", bucket.TimeOfMaximum})
    bucket.EntityData.Leafs.Append("average", types.YLeaf{"Average", bucket.Average})
    bucket.EntityData.Leafs.Append("standard-deviation", types.YLeaf{"StandardDeviation", bucket.StandardDeviation})
    bucket.EntityData.Leafs.Append("result-count", types.YLeaf{"ResultCount", bucket.ResultCount})
    bucket.EntityData.Leafs.Append("data-sent-count", types.YLeaf{"DataSentCount", bucket.DataSentCount})
    bucket.EntityData.Leafs.Append("data-lost-count", types.YLeaf{"DataLostCount", bucket.DataLostCount})
    bucket.EntityData.Leafs.Append("overall-flr", types.YLeaf{"OverallFlr", bucket.OverallFlr})
    bucket.EntityData.Leafs.Append("suspect-start-mid-bucket", types.YLeaf{"SuspectStartMidBucket", bucket.SuspectStartMidBucket})
    bucket.EntityData.Leafs.Append("suspect-schedule-latency", types.YLeaf{"SuspectScheduleLatency", bucket.SuspectScheduleLatency})
    bucket.EntityData.Leafs.Append("suspect-send-fail", types.YLeaf{"SuspectSendFail", bucket.SuspectSendFail})
    bucket.EntityData.Leafs.Append("suspect-premature-end", types.YLeaf{"SuspectPrematureEnd", bucket.SuspectPrematureEnd})
    bucket.EntityData.Leafs.Append("suspect-clock-drift", types.YLeaf{"SuspectClockDrift", bucket.SuspectClockDrift})
    bucket.EntityData.Leafs.Append("suspect-memory-allocation-failed", types.YLeaf{"SuspectMemoryAllocationFailed", bucket.SuspectMemoryAllocationFailed})
    bucket.EntityData.Leafs.Append("suspect-cleared-mid-bucket", types.YLeaf{"SuspectClearedMidBucket", bucket.SuspectClearedMidBucket})
    bucket.EntityData.Leafs.Append("suspect-probe-restarted", types.YLeaf{"SuspectProbeRestarted", bucket.SuspectProbeRestarted})
    bucket.EntityData.Leafs.Append("suspect-management-latency", types.YLeaf{"SuspectManagementLatency", bucket.SuspectManagementLatency})
    bucket.EntityData.Leafs.Append("suspect-multiple-buckets", types.YLeaf{"SuspectMultipleBuckets", bucket.SuspectMultipleBuckets})
    bucket.EntityData.Leafs.Append("suspect-misordering", types.YLeaf{"SuspectMisordering", bucket.SuspectMisordering})
    bucket.EntityData.Leafs.Append("suspect-flr-low-packet-count", types.YLeaf{"SuspectFlrLowPacketCount", bucket.SuspectFlrLowPacketCount})
    bucket.EntityData.Leafs.Append("premature-reason", types.YLeaf{"PrematureReason", bucket.PrematureReason})
    bucket.EntityData.Leafs.Append("premature-reason-string", types.YLeaf{"PrematureReasonString", bucket.PrematureReasonString})

    bucket.EntityData.YListKeys = []string {}

    return &(bucket.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetEntityData() *types.CommonEntityData {
    contents.EntityData.YFilter = contents.YFilter
    contents.EntityData.YangName = "contents"
    contents.EntityData.BundleName = "cisco_ios_xr"
    contents.EntityData.ParentYangName = "bucket"
    contents.EntityData.SegmentPath = "contents"
    contents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contents.EntityData.Children = types.NewOrderedMap()
    contents.EntityData.Children.Append("aggregated", types.YChild{"Aggregated", &contents.Aggregated})
    contents.EntityData.Children.Append("unaggregated", types.YChild{"Unaggregated", &contents.Unaggregated})
    contents.EntityData.Leafs = types.NewOrderedMap()
    contents.EntityData.Leafs.Append("bucket-type", types.YLeaf{"BucketType", contents.BucketType})

    contents.EntityData.YListKeys = []string {}

    return &(contents.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []*Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetEntityData() *types.CommonEntityData {
    aggregated.EntityData.YFilter = aggregated.YFilter
    aggregated.EntityData.YangName = "aggregated"
    aggregated.EntityData.BundleName = "cisco_ios_xr"
    aggregated.EntityData.ParentYangName = "contents"
    aggregated.EntityData.SegmentPath = "aggregated"
    aggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregated.EntityData.Children = types.NewOrderedMap()
    aggregated.EntityData.Children.Append("bins", types.YChild{"Bins", nil})
    for i := range aggregated.Bins {
        aggregated.EntityData.Children.Append(types.GetSegmentPath(aggregated.Bins[i]), types.YChild{"Bins", aggregated.Bins[i]})
    }
    aggregated.EntityData.Leafs = types.NewOrderedMap()

    aggregated.EntityData.YListKeys = []string {}

    return &(aggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower bound (inclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    LowerBound interface{}

    // Upper bound (exclusive) of the bin, in milliseconds or single units of
    // percent. This field is not used for LMM measurements. The type is
    // interface{} with range: -2147483648..2147483647.
    UpperBound interface{}

    // Lower bound (inclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    LowerBoundTenths interface{}

    // Upper bound (exclusive) of the bin, in tenths of percent. This field is
    // only used for LMM measurements. The type is interface{} with range:
    // -2147483648..2147483647. Units are percentage.
    UpperBoundTenths interface{}

    // The sum of the results in the bin, in microseconds or millionths of a
    // percent. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Sum interface{}

    // The total number of results in the bin. The type is interface{} with range:
    // 0..4294967295.
    Count interface{}
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetEntityData() *types.CommonEntityData {
    bins.EntityData.YFilter = bins.YFilter
    bins.EntityData.YangName = "bins"
    bins.EntityData.BundleName = "cisco_ios_xr"
    bins.EntityData.ParentYangName = "aggregated"
    bins.EntityData.SegmentPath = "bins"
    bins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bins.EntityData.Children = types.NewOrderedMap()
    bins.EntityData.Leafs = types.NewOrderedMap()
    bins.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", bins.LowerBound})
    bins.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", bins.UpperBound})
    bins.EntityData.Leafs.Append("lower-bound-tenths", types.YLeaf{"LowerBoundTenths", bins.LowerBoundTenths})
    bins.EntityData.Leafs.Append("upper-bound-tenths", types.YLeaf{"UpperBoundTenths", bins.UpperBoundTenths})
    bins.EntityData.Leafs.Append("sum", types.YLeaf{"Sum", bins.Sum})
    bins.EntityData.Leafs.Append("count", types.YLeaf{"Count", bins.Count})

    bins.EntityData.YListKeys = []string {}

    return &(bins.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []*Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetEntityData() *types.CommonEntityData {
    unaggregated.EntityData.YFilter = unaggregated.YFilter
    unaggregated.EntityData.YangName = "unaggregated"
    unaggregated.EntityData.BundleName = "cisco_ios_xr"
    unaggregated.EntityData.ParentYangName = "contents"
    unaggregated.EntityData.SegmentPath = "unaggregated"
    unaggregated.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unaggregated.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unaggregated.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unaggregated.EntityData.Children = types.NewOrderedMap()
    unaggregated.EntityData.Children.Append("sample", types.YChild{"Sample", nil})
    for i := range unaggregated.Sample {
        unaggregated.EntityData.Children.Append(types.GetSegmentPath(unaggregated.Sample[i]), types.YChild{"Sample", unaggregated.Sample[i]})
    }
    unaggregated.EntityData.Leafs = types.NewOrderedMap()

    unaggregated.EntityData.YListKeys = []string {}

    return &(unaggregated.EntityData)
}

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time (in milliseconds relative to the start time of the bucket) that
    // the sample was sent at. The type is interface{} with range: 0..4294967295.
    // Units are millisecond.
    SentAt interface{}

    // Whether the sample packet was sucessfully sent. The type is bool.
    Sent interface{}

    // Whether the sample packet timed out. The type is bool.
    TimedOut interface{}

    // Whether the sample packet was corrupt. The type is bool.
    Corrupt interface{}

    // Whether the sample packet was received out-of-order. The type is bool.
    OutOfOrder interface{}

    // Whether a measurement could not be made because no data packets were sent
    // in the sample period. Only applicable for LMM measurements. The type is
    // bool.
    NoDataPackets interface{}

    // The result (in microseconds or millionths of a percent) of the sample, if
    // available. The type is interface{} with range: -2147483648..2147483647.
    Result interface{}

    // For FLR measurements, the number of frames sent, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesSent interface{}

    // For FLR measurements, the number of frames lost, if available. The type is
    // interface{} with range: 0..4294967295.
    FramesLost interface{}
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetEntityData() *types.CommonEntityData {
    sample.EntityData.YFilter = sample.YFilter
    sample.EntityData.YangName = "sample"
    sample.EntityData.BundleName = "cisco_ios_xr"
    sample.EntityData.ParentYangName = "unaggregated"
    sample.EntityData.SegmentPath = "sample"
    sample.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sample.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sample.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sample.EntityData.Children = types.NewOrderedMap()
    sample.EntityData.Leafs = types.NewOrderedMap()
    sample.EntityData.Leafs.Append("sent-at", types.YLeaf{"SentAt", sample.SentAt})
    sample.EntityData.Leafs.Append("sent", types.YLeaf{"Sent", sample.Sent})
    sample.EntityData.Leafs.Append("timed-out", types.YLeaf{"TimedOut", sample.TimedOut})
    sample.EntityData.Leafs.Append("corrupt", types.YLeaf{"Corrupt", sample.Corrupt})
    sample.EntityData.Leafs.Append("out-of-order", types.YLeaf{"OutOfOrder", sample.OutOfOrder})
    sample.EntityData.Leafs.Append("no-data-packets", types.YLeaf{"NoDataPackets", sample.NoDataPackets})
    sample.EntityData.Leafs.Append("result", types.YLeaf{"Result", sample.Result})
    sample.EntityData.Leafs.Append("frames-sent", types.YLeaf{"FramesSent", sample.FramesSent})
    sample.EntityData.Leafs.Append("frames-lost", types.YLeaf{"FramesLost", sample.FramesLost})

    sample.EntityData.YListKeys = []string {}

    return &(sample.EntityData)
}

// SlaNodes
// sla nodes
type SlaNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (slaNodes *SlaNodes) GetEntityData() *types.CommonEntityData {
    slaNodes.EntityData.YFilter = slaNodes.YFilter
    slaNodes.EntityData.YangName = "sla-nodes"
    slaNodes.EntityData.BundleName = "cisco_ios_xr"
    slaNodes.EntityData.ParentYangName = "Cisco-IOS-XR-infra-sla-oper"
    slaNodes.EntityData.SegmentPath = "Cisco-IOS-XR-infra-sla-oper:sla-nodes"
    slaNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaNodes.EntityData.Children = types.NewOrderedMap()
    slaNodes.EntityData.Leafs = types.NewOrderedMap()

    slaNodes.EntityData.YListKeys = []string {}

    return &(slaNodes.EntityData)
}

