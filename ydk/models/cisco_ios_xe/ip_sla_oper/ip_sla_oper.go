// This module contains a collection of YANG definitions
// for monitoring of IP SLA statistics of a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_sla_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_sla_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-ip-sla-oper ip-sla-stats}", reflect.TypeOf(IpSlaStats{}))
    ydk.RegisterEntity("Cisco-IOS-XE-ip-sla-oper:ip-sla-stats", reflect.TypeOf(IpSlaStats{}))
}

// TtlType represents IP SLA time-to-live type
type TtlType string

const (
    TtlType_ttl_finite TtlType = "ttl-finite"

    TtlType_ttl_forever TtlType = "ttl-forever"
)

// RttType represents IP SLA RTT type
type RttType string

const (
    RttType_rtt_known RttType = "rtt-known"

    RttType_rtt_unknown RttType = "rtt-unknown"

    RttType_rtt_could_not_find RttType = "rtt-could-not-find"
)

// SlaOperType represents IP SLA operational type
type SlaOperType string

const (
    SlaOperType_oper_type_unknown SlaOperType = "oper-type-unknown"

    SlaOperType_oper_type_udp_echo SlaOperType = "oper-type-udp-echo"

    SlaOperType_oper_type_udp_jitter SlaOperType = "oper-type-udp-jitter"

    SlaOperType_oper_type_icmp_jitter SlaOperType = "oper-type-icmp-jitter"

    SlaOperType_oper_type_ethernet_jitter SlaOperType = "oper-type-ethernet-jitter"

    SlaOperType_oper_type_ethernet_echo SlaOperType = "oper-type-ethernet-echo"

    SlaOperType_oper_type_y1731_delay SlaOperType = "oper-type-y1731-delay"

    SlaOperType_oper_type_y1731_loss SlaOperType = "oper-type-y1731-loss"

    SlaOperType_oper_type_video SlaOperType = "oper-type-video"

    SlaOperType_oper_type_mcast SlaOperType = "oper-type-mcast"

    SlaOperType_oper_type_pong SlaOperType = "oper-type-pong"

    SlaOperType_oper_type_path_jitter SlaOperType = "oper-type-path-jitter"

    SlaOperType_oper_type_icmp_echo SlaOperType = "oper-type-icmp-echo"
)

// SlaReturnCode represents IP SLA return code
type SlaReturnCode string

const (
    SlaReturnCode_ret_code_unknown SlaReturnCode = "ret-code-unknown"

    SlaReturnCode_ret_code_ok SlaReturnCode = "ret-code-ok"

    SlaReturnCode_ret_code_disconnected SlaReturnCode = "ret-code-disconnected"

    SlaReturnCode_ret_code_busy SlaReturnCode = "ret-code-busy"

    SlaReturnCode_ret_code_timeout SlaReturnCode = "ret-code-timeout"

    SlaReturnCode_ret_code_no_connection SlaReturnCode = "ret-code-no-connection"

    SlaReturnCode_ret_code_internal_error SlaReturnCode = "ret-code-internal-error"

    SlaReturnCode_ret_code_operation_failure SlaReturnCode = "ret-code-operation-failure"

    SlaReturnCode_ret_code_could_not_find SlaReturnCode = "ret-code-could-not-find"
)

// AccuracyType represents IP SLA accuracy type
type AccuracyType string

const (
    AccuracyType_accuracy_milliseconds AccuracyType = "accuracy-milliseconds"

    AccuracyType_accuracy_microseconds AccuracyType = "accuracy-microseconds"
)

// IpSlaStats
// Data nodes for All IP SLA Statistics
type IpSlaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of IP SLA operations with statistics info. The type is slice of
    // IpSlaStats_SlaOperEntry.
    SlaOperEntry []*IpSlaStats_SlaOperEntry
}

func (ipSlaStats *IpSlaStats) GetEntityData() *types.CommonEntityData {
    ipSlaStats.EntityData.YFilter = ipSlaStats.YFilter
    ipSlaStats.EntityData.YangName = "ip-sla-stats"
    ipSlaStats.EntityData.BundleName = "cisco_ios_xe"
    ipSlaStats.EntityData.ParentYangName = "Cisco-IOS-XE-ip-sla-oper"
    ipSlaStats.EntityData.SegmentPath = "Cisco-IOS-XE-ip-sla-oper:ip-sla-stats"
    ipSlaStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSlaStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSlaStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSlaStats.EntityData.Children = types.NewOrderedMap()
    ipSlaStats.EntityData.Children.Append("sla-oper-entry", types.YChild{"SlaOperEntry", nil})
    for i := range ipSlaStats.SlaOperEntry {
        ipSlaStats.EntityData.Children.Append(types.GetSegmentPath(ipSlaStats.SlaOperEntry[i]), types.YChild{"SlaOperEntry", ipSlaStats.SlaOperEntry[i]})
    }
    ipSlaStats.EntityData.Leafs = types.NewOrderedMap()

    ipSlaStats.EntityData.YListKeys = []string {}

    return &(ipSlaStats.EntityData)
}

// IpSlaStats_SlaOperEntry
// The list of IP SLA operations with statistics info
type IpSlaStats_SlaOperEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Entry unique identifier. The type is interface{}
    // with range: 0..4294967295.
    OperId interface{}

    // Entry type. The type is SlaOperType.
    OperType interface{}

    // Latest return code. The type is SlaReturnCode.
    LatestReturnCode interface{}

    // Success count. The type is interface{} with range: 0..4294967295.
    SuccessCount interface{}

    // Failure count. The type is interface{} with range: 0..4294967295.
    FailureCount interface{}

    // Latest start time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LatestOperStartTime interface{}

    // RTT information.
    RttInfo IpSlaStats_SlaOperEntry_RttInfo

    // Measured statistics.
    MeasureStats IpSlaStats_SlaOperEntry_MeasureStats

    // Statistics.
    Stats IpSlaStats_SlaOperEntry_Stats
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetEntityData() *types.CommonEntityData {
    slaOperEntry.EntityData.YFilter = slaOperEntry.YFilter
    slaOperEntry.EntityData.YangName = "sla-oper-entry"
    slaOperEntry.EntityData.BundleName = "cisco_ios_xe"
    slaOperEntry.EntityData.ParentYangName = "ip-sla-stats"
    slaOperEntry.EntityData.SegmentPath = "sla-oper-entry" + types.AddKeyToken(slaOperEntry.OperId, "oper-id")
    slaOperEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    slaOperEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    slaOperEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    slaOperEntry.EntityData.Children = types.NewOrderedMap()
    slaOperEntry.EntityData.Children.Append("rtt-info", types.YChild{"RttInfo", &slaOperEntry.RttInfo})
    slaOperEntry.EntityData.Children.Append("measure-stats", types.YChild{"MeasureStats", &slaOperEntry.MeasureStats})
    slaOperEntry.EntityData.Children.Append("stats", types.YChild{"Stats", &slaOperEntry.Stats})
    slaOperEntry.EntityData.Leafs = types.NewOrderedMap()
    slaOperEntry.EntityData.Leafs.Append("oper-id", types.YLeaf{"OperId", slaOperEntry.OperId})
    slaOperEntry.EntityData.Leafs.Append("oper-type", types.YLeaf{"OperType", slaOperEntry.OperType})
    slaOperEntry.EntityData.Leafs.Append("latest-return-code", types.YLeaf{"LatestReturnCode", slaOperEntry.LatestReturnCode})
    slaOperEntry.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", slaOperEntry.SuccessCount})
    slaOperEntry.EntityData.Leafs.Append("failure-count", types.YLeaf{"FailureCount", slaOperEntry.FailureCount})
    slaOperEntry.EntityData.Leafs.Append("latest-oper-start-time", types.YLeaf{"LatestOperStartTime", slaOperEntry.LatestOperStartTime})

    slaOperEntry.EntityData.YListKeys = []string {"OperId"}

    return &(slaOperEntry.EntityData)
}

// IpSlaStats_SlaOperEntry_RttInfo
// RTT information
type IpSlaStats_SlaOperEntry_RttInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The last Round Trip Time recorded for this SLA.
    LatestRtt IpSlaStats_SlaOperEntry_RttInfo_LatestRtt

    // Time-to-live for the SLA operation.
    TimeToLive IpSlaStats_SlaOperEntry_RttInfo_TimeToLive
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetEntityData() *types.CommonEntityData {
    rttInfo.EntityData.YFilter = rttInfo.YFilter
    rttInfo.EntityData.YangName = "rtt-info"
    rttInfo.EntityData.BundleName = "cisco_ios_xe"
    rttInfo.EntityData.ParentYangName = "sla-oper-entry"
    rttInfo.EntityData.SegmentPath = "rtt-info"
    rttInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rttInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rttInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rttInfo.EntityData.Children = types.NewOrderedMap()
    rttInfo.EntityData.Children.Append("latest-rtt", types.YChild{"LatestRtt", &rttInfo.LatestRtt})
    rttInfo.EntityData.Children.Append("time-to-live", types.YChild{"TimeToLive", &rttInfo.TimeToLive})
    rttInfo.EntityData.Leafs = types.NewOrderedMap()

    rttInfo.EntityData.YListKeys = []string {}

    return &(rttInfo.EntityData)
}

// IpSlaStats_SlaOperEntry_RttInfo_LatestRtt
// The last Round Trip Time recorded for this SLA
type IpSlaStats_SlaOperEntry_RttInfo_LatestRtt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Round trip time value. The type is interface{} with range:
    // 0..18446744073709551615.
    Rtt interface{}

    // Round trip time is unknown. The type is interface{}.
    Unknown interface{}

    // Round trip time could not be determined. The type is interface{}.
    CouldNotFind interface{}
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetEntityData() *types.CommonEntityData {
    latestRtt.EntityData.YFilter = latestRtt.YFilter
    latestRtt.EntityData.YangName = "latest-rtt"
    latestRtt.EntityData.BundleName = "cisco_ios_xe"
    latestRtt.EntityData.ParentYangName = "rtt-info"
    latestRtt.EntityData.SegmentPath = "latest-rtt"
    latestRtt.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    latestRtt.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    latestRtt.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    latestRtt.EntityData.Children = types.NewOrderedMap()
    latestRtt.EntityData.Leafs = types.NewOrderedMap()
    latestRtt.EntityData.Leafs.Append("rtt", types.YLeaf{"Rtt", latestRtt.Rtt})
    latestRtt.EntityData.Leafs.Append("unknown", types.YLeaf{"Unknown", latestRtt.Unknown})
    latestRtt.EntityData.Leafs.Append("could-not-find", types.YLeaf{"CouldNotFind", latestRtt.CouldNotFind})

    latestRtt.EntityData.YListKeys = []string {}

    return &(latestRtt.EntityData)
}

// IpSlaStats_SlaOperEntry_RttInfo_TimeToLive
// Time-to-live for the SLA operation
type IpSlaStats_SlaOperEntry_RttInfo_TimeToLive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range:
    // 0..18446744073709551615.
    Ttl interface{}

    // Time to live unbound. The type is interface{}.
    Forever interface{}
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetEntityData() *types.CommonEntityData {
    timeToLive.EntityData.YFilter = timeToLive.YFilter
    timeToLive.EntityData.YangName = "time-to-live"
    timeToLive.EntityData.BundleName = "cisco_ios_xe"
    timeToLive.EntityData.ParentYangName = "rtt-info"
    timeToLive.EntityData.SegmentPath = "time-to-live"
    timeToLive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    timeToLive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    timeToLive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    timeToLive.EntityData.Children = types.NewOrderedMap()
    timeToLive.EntityData.Leafs = types.NewOrderedMap()
    timeToLive.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", timeToLive.Ttl})
    timeToLive.EntityData.Leafs.Append("forever", types.YLeaf{"Forever", timeToLive.Forever})

    timeToLive.EntityData.YListKeys = []string {}

    return &(timeToLive.EntityData)
}

// IpSlaStats_SlaOperEntry_MeasureStats
// Measured statistics
type IpSlaStats_SlaOperEntry_MeasureStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval start time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    IntvStartTime interface{}

    // Initial count. The type is interface{} with range: 0..4294967295.
    InitCount interface{}

    // Complete count. The type is interface{} with range: 0..4294967295.
    CompleteCount interface{}

    // Validity. The type is bool.
    Valid interface{}
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetEntityData() *types.CommonEntityData {
    measureStats.EntityData.YFilter = measureStats.YFilter
    measureStats.EntityData.YangName = "measure-stats"
    measureStats.EntityData.BundleName = "cisco_ios_xe"
    measureStats.EntityData.ParentYangName = "sla-oper-entry"
    measureStats.EntityData.SegmentPath = "measure-stats"
    measureStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    measureStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    measureStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    measureStats.EntityData.Children = types.NewOrderedMap()
    measureStats.EntityData.Leafs = types.NewOrderedMap()
    measureStats.EntityData.Leafs.Append("intv-start-time", types.YLeaf{"IntvStartTime", measureStats.IntvStartTime})
    measureStats.EntityData.Leafs.Append("init-count", types.YLeaf{"InitCount", measureStats.InitCount})
    measureStats.EntityData.Leafs.Append("complete-count", types.YLeaf{"CompleteCount", measureStats.CompleteCount})
    measureStats.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", measureStats.Valid})

    measureStats.EntityData.YListKeys = []string {}

    return &(measureStats.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats
// Statistics
type IpSlaStats_SlaOperEntry_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RTT value.
    Rtt IpSlaStats_SlaOperEntry_Stats_Rtt

    // Latency information.
    OnewayLatency IpSlaStats_SlaOperEntry_Stats_OnewayLatency

    // Jitter information.
    Jitter IpSlaStats_SlaOperEntry_Stats_Jitter

    // Over threshold information.
    OverThreshold IpSlaStats_SlaOperEntry_Stats_OverThreshold

    // Packet loss information.
    PacketLoss IpSlaStats_SlaOperEntry_Stats_PacketLoss

    // ICMP packet loss information.
    IcmpPacketLoss IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss

    // Voice score information.
    VoiceScore IpSlaStats_SlaOperEntry_Stats_VoiceScore
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xe"
    stats.EntityData.ParentYangName = "sla-oper-entry"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Children.Append("rtt", types.YChild{"Rtt", &stats.Rtt})
    stats.EntityData.Children.Append("oneway-latency", types.YChild{"OnewayLatency", &stats.OnewayLatency})
    stats.EntityData.Children.Append("jitter", types.YChild{"Jitter", &stats.Jitter})
    stats.EntityData.Children.Append("over-threshold", types.YChild{"OverThreshold", &stats.OverThreshold})
    stats.EntityData.Children.Append("packet-loss", types.YChild{"PacketLoss", &stats.PacketLoss})
    stats.EntityData.Children.Append("icmp-packet-loss", types.YChild{"IcmpPacketLoss", &stats.IcmpPacketLoss})
    stats.EntityData.Children.Append("voice-score", types.YChild{"VoiceScore", &stats.VoiceScore})
    stats.EntityData.Leafs = types.NewOrderedMap()

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_Rtt
// RTT value
type IpSlaStats_SlaOperEntry_Stats_Rtt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RTT count. The type is interface{} with range: 0..4294967295.
    RttCount interface{}

    // Timing information.
    SlaTimeValues IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetEntityData() *types.CommonEntityData {
    rtt.EntityData.YFilter = rtt.YFilter
    rtt.EntityData.YangName = "rtt"
    rtt.EntityData.BundleName = "cisco_ios_xe"
    rtt.EntityData.ParentYangName = "stats"
    rtt.EntityData.SegmentPath = "rtt"
    rtt.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rtt.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rtt.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rtt.EntityData.Children = types.NewOrderedMap()
    rtt.EntityData.Children.Append("sla-time-values", types.YChild{"SlaTimeValues", &rtt.SlaTimeValues})
    rtt.EntityData.Leafs = types.NewOrderedMap()
    rtt.EntityData.Leafs.Append("rtt-count", types.YLeaf{"RttCount", rtt.RttCount})

    rtt.EntityData.YListKeys = []string {}

    return &(rtt.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues
// Timing information
type IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value reading. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // Average value reading. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // Maximum value reading. The type is interface{} with range: 0..4294967295.
    Max interface{}

    // Reading accuracy. The type is AccuracyType.
    Accuracy interface{}
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetEntityData() *types.CommonEntityData {
    slaTimeValues.EntityData.YFilter = slaTimeValues.YFilter
    slaTimeValues.EntityData.YangName = "sla-time-values"
    slaTimeValues.EntityData.BundleName = "cisco_ios_xe"
    slaTimeValues.EntityData.ParentYangName = "rtt"
    slaTimeValues.EntityData.SegmentPath = "sla-time-values"
    slaTimeValues.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    slaTimeValues.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    slaTimeValues.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    slaTimeValues.EntityData.Children = types.NewOrderedMap()
    slaTimeValues.EntityData.Leafs = types.NewOrderedMap()
    slaTimeValues.EntityData.Leafs.Append("min", types.YLeaf{"Min", slaTimeValues.Min})
    slaTimeValues.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", slaTimeValues.Avg})
    slaTimeValues.EntityData.Leafs.Append("max", types.YLeaf{"Max", slaTimeValues.Max})
    slaTimeValues.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", slaTimeValues.Accuracy})

    slaTimeValues.EntityData.YListKeys = []string {}

    return &(slaTimeValues.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency
// Latency information
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sample count. The type is interface{} with range: 0..4294967295.
    SampleCount interface{}

    // Source to Destination for the one-way latency.
    Sd IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd

    // Destination to Source for the one-way latency.
    Ds IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetEntityData() *types.CommonEntityData {
    onewayLatency.EntityData.YFilter = onewayLatency.YFilter
    onewayLatency.EntityData.YangName = "oneway-latency"
    onewayLatency.EntityData.BundleName = "cisco_ios_xe"
    onewayLatency.EntityData.ParentYangName = "stats"
    onewayLatency.EntityData.SegmentPath = "oneway-latency"
    onewayLatency.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    onewayLatency.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    onewayLatency.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    onewayLatency.EntityData.Children = types.NewOrderedMap()
    onewayLatency.EntityData.Children.Append("sd", types.YChild{"Sd", &onewayLatency.Sd})
    onewayLatency.EntityData.Children.Append("ds", types.YChild{"Ds", &onewayLatency.Ds})
    onewayLatency.EntityData.Leafs = types.NewOrderedMap()
    onewayLatency.EntityData.Leafs.Append("sample-count", types.YLeaf{"SampleCount", onewayLatency.SampleCount})

    onewayLatency.EntityData.YListKeys = []string {}

    return &(onewayLatency.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd
// Source to Destination for the one-way latency
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value reading. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // Average value reading. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // Maximum value reading. The type is interface{} with range: 0..4294967295.
    Max interface{}

    // Reading accuracy. The type is AccuracyType.
    Accuracy interface{}
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetEntityData() *types.CommonEntityData {
    sd.EntityData.YFilter = sd.YFilter
    sd.EntityData.YangName = "sd"
    sd.EntityData.BundleName = "cisco_ios_xe"
    sd.EntityData.ParentYangName = "oneway-latency"
    sd.EntityData.SegmentPath = "sd"
    sd.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sd.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sd.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sd.EntityData.Children = types.NewOrderedMap()
    sd.EntityData.Leafs = types.NewOrderedMap()
    sd.EntityData.Leafs.Append("min", types.YLeaf{"Min", sd.Min})
    sd.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", sd.Avg})
    sd.EntityData.Leafs.Append("max", types.YLeaf{"Max", sd.Max})
    sd.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", sd.Accuracy})

    sd.EntityData.YListKeys = []string {}

    return &(sd.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds
// Destination to Source for the one-way latency
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value reading. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // Average value reading. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // Maximum value reading. The type is interface{} with range: 0..4294967295.
    Max interface{}

    // Reading accuracy. The type is AccuracyType.
    Accuracy interface{}
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetEntityData() *types.CommonEntityData {
    ds.EntityData.YFilter = ds.YFilter
    ds.EntityData.YangName = "ds"
    ds.EntityData.BundleName = "cisco_ios_xe"
    ds.EntityData.ParentYangName = "oneway-latency"
    ds.EntityData.SegmentPath = "ds"
    ds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ds.EntityData.Children = types.NewOrderedMap()
    ds.EntityData.Leafs = types.NewOrderedMap()
    ds.EntityData.Leafs.Append("min", types.YLeaf{"Min", ds.Min})
    ds.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", ds.Avg})
    ds.EntityData.Leafs.Append("max", types.YLeaf{"Max", ds.Max})
    ds.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", ds.Accuracy})

    ds.EntityData.YListKeys = []string {}

    return &(ds.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_Jitter
// Jitter information
type IpSlaStats_SlaOperEntry_Stats_Jitter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sample count. The type is interface{} with range: 0..4294967295.
    SdSampleCount interface{}

    // Sample count. The type is interface{} with range: 0..4294967295.
    DsSampleCount interface{}

    // Source to Destination for the jitter.
    Sd IpSlaStats_SlaOperEntry_Stats_Jitter_Sd

    // Destination to Source for the jitter.
    Ds IpSlaStats_SlaOperEntry_Stats_Jitter_Ds
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetEntityData() *types.CommonEntityData {
    jitter.EntityData.YFilter = jitter.YFilter
    jitter.EntityData.YangName = "jitter"
    jitter.EntityData.BundleName = "cisco_ios_xe"
    jitter.EntityData.ParentYangName = "stats"
    jitter.EntityData.SegmentPath = "jitter"
    jitter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    jitter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    jitter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    jitter.EntityData.Children = types.NewOrderedMap()
    jitter.EntityData.Children.Append("sd", types.YChild{"Sd", &jitter.Sd})
    jitter.EntityData.Children.Append("ds", types.YChild{"Ds", &jitter.Ds})
    jitter.EntityData.Leafs = types.NewOrderedMap()
    jitter.EntityData.Leafs.Append("sd-sample-count", types.YLeaf{"SdSampleCount", jitter.SdSampleCount})
    jitter.EntityData.Leafs.Append("ds-sample-count", types.YLeaf{"DsSampleCount", jitter.DsSampleCount})

    jitter.EntityData.YListKeys = []string {}

    return &(jitter.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_Jitter_Sd
// Source to Destination for the jitter
type IpSlaStats_SlaOperEntry_Stats_Jitter_Sd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value reading. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // Average value reading. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // Maximum value reading. The type is interface{} with range: 0..4294967295.
    Max interface{}

    // Reading accuracy. The type is AccuracyType.
    Accuracy interface{}
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetEntityData() *types.CommonEntityData {
    sd.EntityData.YFilter = sd.YFilter
    sd.EntityData.YangName = "sd"
    sd.EntityData.BundleName = "cisco_ios_xe"
    sd.EntityData.ParentYangName = "jitter"
    sd.EntityData.SegmentPath = "sd"
    sd.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sd.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sd.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sd.EntityData.Children = types.NewOrderedMap()
    sd.EntityData.Leafs = types.NewOrderedMap()
    sd.EntityData.Leafs.Append("min", types.YLeaf{"Min", sd.Min})
    sd.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", sd.Avg})
    sd.EntityData.Leafs.Append("max", types.YLeaf{"Max", sd.Max})
    sd.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", sd.Accuracy})

    sd.EntityData.YListKeys = []string {}

    return &(sd.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_Jitter_Ds
// Destination to Source for the jitter
type IpSlaStats_SlaOperEntry_Stats_Jitter_Ds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum value reading. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // Average value reading. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // Maximum value reading. The type is interface{} with range: 0..4294967295.
    Max interface{}

    // Reading accuracy. The type is AccuracyType.
    Accuracy interface{}
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetEntityData() *types.CommonEntityData {
    ds.EntityData.YFilter = ds.YFilter
    ds.EntityData.YangName = "ds"
    ds.EntityData.BundleName = "cisco_ios_xe"
    ds.EntityData.ParentYangName = "jitter"
    ds.EntityData.SegmentPath = "ds"
    ds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ds.EntityData.Children = types.NewOrderedMap()
    ds.EntityData.Leafs = types.NewOrderedMap()
    ds.EntityData.Leafs.Append("min", types.YLeaf{"Min", ds.Min})
    ds.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", ds.Avg})
    ds.EntityData.Leafs.Append("max", types.YLeaf{"Max", ds.Max})
    ds.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", ds.Accuracy})

    ds.EntityData.YListKeys = []string {}

    return &(ds.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_OverThreshold
// Over threshold information
type IpSlaStats_SlaOperEntry_Stats_OverThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Round Trip Time (RTT) over threshold count (the number of times that the
    // RTT was over the configured threshold). The type is interface{} with range:
    // 0..4294967295.
    RttCount interface{}

    // Round Trip Time over threshold percentage (the percentage that the RTT was
    // over the configured threshold). The type is interface{} with range: 0..255.
    Percent interface{}
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetEntityData() *types.CommonEntityData {
    overThreshold.EntityData.YFilter = overThreshold.YFilter
    overThreshold.EntityData.YangName = "over-threshold"
    overThreshold.EntityData.BundleName = "cisco_ios_xe"
    overThreshold.EntityData.ParentYangName = "stats"
    overThreshold.EntityData.SegmentPath = "over-threshold"
    overThreshold.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    overThreshold.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    overThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    overThreshold.EntityData.Children = types.NewOrderedMap()
    overThreshold.EntityData.Leafs = types.NewOrderedMap()
    overThreshold.EntityData.Leafs.Append("rtt-count", types.YLeaf{"RttCount", overThreshold.RttCount})
    overThreshold.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", overThreshold.Percent})

    overThreshold.EntityData.YListKeys = []string {}

    return &(overThreshold.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_PacketLoss
// Packet loss information
type IpSlaStats_SlaOperEntry_Stats_PacketLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unprocessed packet count. The type is interface{} with range:
    // 0..4294967295.
    UnprocessedPackets interface{}

    // Number of packets lost from Source to Destination. The type is interface{}
    // with range: 0..4294967295.
    SdCount interface{}

    // Number of packets lost from Destination to Source. The type is interface{}
    // with range: 0..4294967295.
    DsCount interface{}

    // Out of sequence packet count. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequence interface{}

    // Dropped packet count. The type is interface{} with range: 0..4294967295.
    Drops interface{}

    // Late arrival packet count. The type is interface{} with range:
    // 0..4294967295.
    LateArrivals interface{}

    // Skipped packet count. The type is interface{} with range: 0..4294967295.
    SkippedPackets interface{}

    // Source to Destination packet loss details.
    SdLoss IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss

    // Destination to Source packet loss details.
    DsLoss IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetEntityData() *types.CommonEntityData {
    packetLoss.EntityData.YFilter = packetLoss.YFilter
    packetLoss.EntityData.YangName = "packet-loss"
    packetLoss.EntityData.BundleName = "cisco_ios_xe"
    packetLoss.EntityData.ParentYangName = "stats"
    packetLoss.EntityData.SegmentPath = "packet-loss"
    packetLoss.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    packetLoss.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    packetLoss.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    packetLoss.EntityData.Children = types.NewOrderedMap()
    packetLoss.EntityData.Children.Append("sd-loss", types.YChild{"SdLoss", &packetLoss.SdLoss})
    packetLoss.EntityData.Children.Append("ds-loss", types.YChild{"DsLoss", &packetLoss.DsLoss})
    packetLoss.EntityData.Leafs = types.NewOrderedMap()
    packetLoss.EntityData.Leafs.Append("unprocessed-packets", types.YLeaf{"UnprocessedPackets", packetLoss.UnprocessedPackets})
    packetLoss.EntityData.Leafs.Append("sd-count", types.YLeaf{"SdCount", packetLoss.SdCount})
    packetLoss.EntityData.Leafs.Append("ds-count", types.YLeaf{"DsCount", packetLoss.DsCount})
    packetLoss.EntityData.Leafs.Append("out-of-sequence", types.YLeaf{"OutOfSequence", packetLoss.OutOfSequence})
    packetLoss.EntityData.Leafs.Append("drops", types.YLeaf{"Drops", packetLoss.Drops})
    packetLoss.EntityData.Leafs.Append("late-arrivals", types.YLeaf{"LateArrivals", packetLoss.LateArrivals})
    packetLoss.EntityData.Leafs.Append("skipped-packets", types.YLeaf{"SkippedPackets", packetLoss.SkippedPackets})

    packetLoss.EntityData.YListKeys = []string {}

    return &(packetLoss.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss
// Source to Destination packet loss details
type IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Loss period count. The type is interface{} with range: 0..4294967295.
    LossPeriodCount interface{}

    // Shortest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMin interface{}

    // Longest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMax interface{}

    // Shortest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMin interface{}

    // Longest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMax interface{}
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetEntityData() *types.CommonEntityData {
    sdLoss.EntityData.YFilter = sdLoss.YFilter
    sdLoss.EntityData.YangName = "sd-loss"
    sdLoss.EntityData.BundleName = "cisco_ios_xe"
    sdLoss.EntityData.ParentYangName = "packet-loss"
    sdLoss.EntityData.SegmentPath = "sd-loss"
    sdLoss.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sdLoss.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sdLoss.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sdLoss.EntityData.Children = types.NewOrderedMap()
    sdLoss.EntityData.Leafs = types.NewOrderedMap()
    sdLoss.EntityData.Leafs.Append("loss-period-count", types.YLeaf{"LossPeriodCount", sdLoss.LossPeriodCount})
    sdLoss.EntityData.Leafs.Append("loss-period-len-min", types.YLeaf{"LossPeriodLenMin", sdLoss.LossPeriodLenMin})
    sdLoss.EntityData.Leafs.Append("loss-period-len-max", types.YLeaf{"LossPeriodLenMax", sdLoss.LossPeriodLenMax})
    sdLoss.EntityData.Leafs.Append("inter-loss-period-len-min", types.YLeaf{"InterLossPeriodLenMin", sdLoss.InterLossPeriodLenMin})
    sdLoss.EntityData.Leafs.Append("inter-loss-period-len-max", types.YLeaf{"InterLossPeriodLenMax", sdLoss.InterLossPeriodLenMax})

    sdLoss.EntityData.YListKeys = []string {}

    return &(sdLoss.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss
// Destination to Source packet loss details
type IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Loss period count. The type is interface{} with range: 0..4294967295.
    LossPeriodCount interface{}

    // Shortest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMin interface{}

    // Longest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMax interface{}

    // Shortest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMin interface{}

    // Longest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMax interface{}
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetEntityData() *types.CommonEntityData {
    dsLoss.EntityData.YFilter = dsLoss.YFilter
    dsLoss.EntityData.YangName = "ds-loss"
    dsLoss.EntityData.BundleName = "cisco_ios_xe"
    dsLoss.EntityData.ParentYangName = "packet-loss"
    dsLoss.EntityData.SegmentPath = "ds-loss"
    dsLoss.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dsLoss.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dsLoss.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dsLoss.EntityData.Children = types.NewOrderedMap()
    dsLoss.EntityData.Leafs = types.NewOrderedMap()
    dsLoss.EntityData.Leafs.Append("loss-period-count", types.YLeaf{"LossPeriodCount", dsLoss.LossPeriodCount})
    dsLoss.EntityData.Leafs.Append("loss-period-len-min", types.YLeaf{"LossPeriodLenMin", dsLoss.LossPeriodLenMin})
    dsLoss.EntityData.Leafs.Append("loss-period-len-max", types.YLeaf{"LossPeriodLenMax", dsLoss.LossPeriodLenMax})
    dsLoss.EntityData.Leafs.Append("inter-loss-period-len-min", types.YLeaf{"InterLossPeriodLenMin", dsLoss.InterLossPeriodLenMin})
    dsLoss.EntityData.Leafs.Append("inter-loss-period-len-max", types.YLeaf{"InterLossPeriodLenMax", dsLoss.InterLossPeriodLenMax})

    dsLoss.EntityData.YListKeys = []string {}

    return &(dsLoss.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss
// ICMP packet loss information
type IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Late arrival packet count. The type is interface{} with range:
    // 0..4294967295.
    LateArrivals interface{}

    // Out of sequence packet count. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequence interface{}

    // Out of sequence packet count. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceSd interface{}

    // Out of sequence packet count. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceDs interface{}

    // Out of sequence packet count. The type is interface{} with range:
    // 0..4294967295.
    OutOfSequenceBoth interface{}

    // Skipped packet count. The type is interface{} with range: 0..4294967295.
    SkippedPackets interface{}

    // Unprocessed packet count. The type is interface{} with range:
    // 0..4294967295.
    UnprocessedPackets interface{}

    // Lost packet count. The type is interface{} with range: 0..4294967295.
    PacketLoss interface{}

    // Loss period count. The type is interface{} with range: 0..4294967295.
    LossPeriodCount interface{}

    // Shortest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMin interface{}

    // Longest loss period length. The type is interface{} with range:
    // 0..4294967295.
    LossPeriodLenMax interface{}

    // Shortest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMin interface{}

    // Longest inter loss period length. The type is interface{} with range:
    // 0..4294967295.
    InterLossPeriodLenMax interface{}
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetEntityData() *types.CommonEntityData {
    icmpPacketLoss.EntityData.YFilter = icmpPacketLoss.YFilter
    icmpPacketLoss.EntityData.YangName = "icmp-packet-loss"
    icmpPacketLoss.EntityData.BundleName = "cisco_ios_xe"
    icmpPacketLoss.EntityData.ParentYangName = "stats"
    icmpPacketLoss.EntityData.SegmentPath = "icmp-packet-loss"
    icmpPacketLoss.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icmpPacketLoss.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icmpPacketLoss.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icmpPacketLoss.EntityData.Children = types.NewOrderedMap()
    icmpPacketLoss.EntityData.Leafs = types.NewOrderedMap()
    icmpPacketLoss.EntityData.Leafs.Append("late-arrivals", types.YLeaf{"LateArrivals", icmpPacketLoss.LateArrivals})
    icmpPacketLoss.EntityData.Leafs.Append("out-of-sequence", types.YLeaf{"OutOfSequence", icmpPacketLoss.OutOfSequence})
    icmpPacketLoss.EntityData.Leafs.Append("out-of-sequence-sd", types.YLeaf{"OutOfSequenceSd", icmpPacketLoss.OutOfSequenceSd})
    icmpPacketLoss.EntityData.Leafs.Append("out-of-sequence-ds", types.YLeaf{"OutOfSequenceDs", icmpPacketLoss.OutOfSequenceDs})
    icmpPacketLoss.EntityData.Leafs.Append("out-of-sequence-both", types.YLeaf{"OutOfSequenceBoth", icmpPacketLoss.OutOfSequenceBoth})
    icmpPacketLoss.EntityData.Leafs.Append("skipped-packets", types.YLeaf{"SkippedPackets", icmpPacketLoss.SkippedPackets})
    icmpPacketLoss.EntityData.Leafs.Append("unprocessed-packets", types.YLeaf{"UnprocessedPackets", icmpPacketLoss.UnprocessedPackets})
    icmpPacketLoss.EntityData.Leafs.Append("packet-loss", types.YLeaf{"PacketLoss", icmpPacketLoss.PacketLoss})
    icmpPacketLoss.EntityData.Leafs.Append("loss-period-count", types.YLeaf{"LossPeriodCount", icmpPacketLoss.LossPeriodCount})
    icmpPacketLoss.EntityData.Leafs.Append("loss-period-len-min", types.YLeaf{"LossPeriodLenMin", icmpPacketLoss.LossPeriodLenMin})
    icmpPacketLoss.EntityData.Leafs.Append("loss-period-len-max", types.YLeaf{"LossPeriodLenMax", icmpPacketLoss.LossPeriodLenMax})
    icmpPacketLoss.EntityData.Leafs.Append("inter-loss-period-len-min", types.YLeaf{"InterLossPeriodLenMin", icmpPacketLoss.InterLossPeriodLenMin})
    icmpPacketLoss.EntityData.Leafs.Append("inter-loss-period-len-max", types.YLeaf{"InterLossPeriodLenMax", icmpPacketLoss.InterLossPeriodLenMax})

    icmpPacketLoss.EntityData.YListKeys = []string {}

    return &(icmpPacketLoss.EntityData)
}

// IpSlaStats_SlaOperEntry_Stats_VoiceScore
// Voice score information
type IpSlaStats_SlaOperEntry_Stats_VoiceScore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Calculated planning impairment factor. The type is interface{} with range:
    // 0..4294967295.
    Icpif interface{}

    // Mean opinion score. The type is interface{} with range: 0..4294967295.
    Mos interface{}
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetEntityData() *types.CommonEntityData {
    voiceScore.EntityData.YFilter = voiceScore.YFilter
    voiceScore.EntityData.YangName = "voice-score"
    voiceScore.EntityData.BundleName = "cisco_ios_xe"
    voiceScore.EntityData.ParentYangName = "stats"
    voiceScore.EntityData.SegmentPath = "voice-score"
    voiceScore.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    voiceScore.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    voiceScore.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    voiceScore.EntityData.Children = types.NewOrderedMap()
    voiceScore.EntityData.Leafs = types.NewOrderedMap()
    voiceScore.EntityData.Leafs.Append("icpif", types.YLeaf{"Icpif", voiceScore.Icpif})
    voiceScore.EntityData.Leafs.Append("mos", types.YLeaf{"Mos", voiceScore.Mos})

    voiceScore.EntityData.YListKeys = []string {}

    return &(voiceScore.EntityData)
}

