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
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of IP SLA operations with statistics info. The type is slice of
    // IpSlaStats_SlaOperEntry.
    SlaOperEntry []IpSlaStats_SlaOperEntry
}

func (ipSlaStats *IpSlaStats) GetFilter() yfilter.YFilter { return ipSlaStats.YFilter }

func (ipSlaStats *IpSlaStats) SetFilter(yf yfilter.YFilter) { ipSlaStats.YFilter = yf }

func (ipSlaStats *IpSlaStats) GetGoName(yname string) string {
    if yname == "sla-oper-entry" { return "SlaOperEntry" }
    return ""
}

func (ipSlaStats *IpSlaStats) GetSegmentPath() string {
    return "Cisco-IOS-XE-ip-sla-oper:ip-sla-stats"
}

func (ipSlaStats *IpSlaStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sla-oper-entry" {
        for _, c := range ipSlaStats.SlaOperEntry {
            if ipSlaStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IpSlaStats_SlaOperEntry{}
        ipSlaStats.SlaOperEntry = append(ipSlaStats.SlaOperEntry, child)
        return &ipSlaStats.SlaOperEntry[len(ipSlaStats.SlaOperEntry)-1]
    }
    return nil
}

func (ipSlaStats *IpSlaStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipSlaStats.SlaOperEntry {
        children[ipSlaStats.SlaOperEntry[i].GetSegmentPath()] = &ipSlaStats.SlaOperEntry[i]
    }
    return children
}

func (ipSlaStats *IpSlaStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipSlaStats *IpSlaStats) GetBundleName() string { return "cisco_ios_xe" }

func (ipSlaStats *IpSlaStats) GetYangName() string { return "ip-sla-stats" }

func (ipSlaStats *IpSlaStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipSlaStats *IpSlaStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipSlaStats *IpSlaStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipSlaStats *IpSlaStats) SetParent(parent types.Entity) { ipSlaStats.parent = parent }

func (ipSlaStats *IpSlaStats) GetParent() types.Entity { return ipSlaStats.parent }

func (ipSlaStats *IpSlaStats) GetParentYangName() string { return "Cisco-IOS-XE-ip-sla-oper" }

// IpSlaStats_SlaOperEntry
// The list of IP SLA operations with statistics info
type IpSlaStats_SlaOperEntry struct {
    parent types.Entity
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

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetFilter() yfilter.YFilter { return slaOperEntry.YFilter }

func (slaOperEntry *IpSlaStats_SlaOperEntry) SetFilter(yf yfilter.YFilter) { slaOperEntry.YFilter = yf }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetGoName(yname string) string {
    if yname == "oper-id" { return "OperId" }
    if yname == "oper-type" { return "OperType" }
    if yname == "latest-return-code" { return "LatestReturnCode" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "failure-count" { return "FailureCount" }
    if yname == "latest-oper-start-time" { return "LatestOperStartTime" }
    if yname == "rtt-info" { return "RttInfo" }
    if yname == "measure-stats" { return "MeasureStats" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetSegmentPath() string {
    return "sla-oper-entry" + "[oper-id='" + fmt.Sprintf("%v", slaOperEntry.OperId) + "']"
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rtt-info" {
        return &slaOperEntry.RttInfo
    }
    if childYangName == "measure-stats" {
        return &slaOperEntry.MeasureStats
    }
    if childYangName == "stats" {
        return &slaOperEntry.Stats
    }
    return nil
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rtt-info"] = &slaOperEntry.RttInfo
    children["measure-stats"] = &slaOperEntry.MeasureStats
    children["stats"] = &slaOperEntry.Stats
    return children
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-id"] = slaOperEntry.OperId
    leafs["oper-type"] = slaOperEntry.OperType
    leafs["latest-return-code"] = slaOperEntry.LatestReturnCode
    leafs["success-count"] = slaOperEntry.SuccessCount
    leafs["failure-count"] = slaOperEntry.FailureCount
    leafs["latest-oper-start-time"] = slaOperEntry.LatestOperStartTime
    return leafs
}

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetBundleName() string { return "cisco_ios_xe" }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetYangName() string { return "sla-oper-entry" }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (slaOperEntry *IpSlaStats_SlaOperEntry) SetParent(parent types.Entity) { slaOperEntry.parent = parent }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetParent() types.Entity { return slaOperEntry.parent }

func (slaOperEntry *IpSlaStats_SlaOperEntry) GetParentYangName() string { return "ip-sla-stats" }

// IpSlaStats_SlaOperEntry_RttInfo
// RTT information
type IpSlaStats_SlaOperEntry_RttInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The last Round Trip Time recorded for this SLA.
    LatestRtt IpSlaStats_SlaOperEntry_RttInfo_LatestRtt

    // Time-to-live for the SLA operation.
    TimeToLive IpSlaStats_SlaOperEntry_RttInfo_TimeToLive
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetFilter() yfilter.YFilter { return rttInfo.YFilter }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) SetFilter(yf yfilter.YFilter) { rttInfo.YFilter = yf }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetGoName(yname string) string {
    if yname == "latest-rtt" { return "LatestRtt" }
    if yname == "time-to-live" { return "TimeToLive" }
    return ""
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetSegmentPath() string {
    return "rtt-info"
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "latest-rtt" {
        return &rttInfo.LatestRtt
    }
    if childYangName == "time-to-live" {
        return &rttInfo.TimeToLive
    }
    return nil
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["latest-rtt"] = &rttInfo.LatestRtt
    children["time-to-live"] = &rttInfo.TimeToLive
    return children
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetBundleName() string { return "cisco_ios_xe" }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetYangName() string { return "rtt-info" }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) SetParent(parent types.Entity) { rttInfo.parent = parent }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetParent() types.Entity { return rttInfo.parent }

func (rttInfo *IpSlaStats_SlaOperEntry_RttInfo) GetParentYangName() string { return "sla-oper-entry" }

// IpSlaStats_SlaOperEntry_RttInfo_LatestRtt
// The last Round Trip Time recorded for this SLA
type IpSlaStats_SlaOperEntry_RttInfo_LatestRtt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Round trip time value. The type is interface{} with range:
    // 0..18446744073709551615.
    Rtt interface{}

    // Round trip time is unknown. The type is interface{}.
    Unknown interface{}

    // Round trip time could not be determined. The type is interface{}.
    CouldNotFind interface{}
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetFilter() yfilter.YFilter { return latestRtt.YFilter }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) SetFilter(yf yfilter.YFilter) { latestRtt.YFilter = yf }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetGoName(yname string) string {
    if yname == "rtt" { return "Rtt" }
    if yname == "unknown" { return "Unknown" }
    if yname == "could-not-find" { return "CouldNotFind" }
    return ""
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetSegmentPath() string {
    return "latest-rtt"
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtt"] = latestRtt.Rtt
    leafs["unknown"] = latestRtt.Unknown
    leafs["could-not-find"] = latestRtt.CouldNotFind
    return leafs
}

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetBundleName() string { return "cisco_ios_xe" }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetYangName() string { return "latest-rtt" }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) SetParent(parent types.Entity) { latestRtt.parent = parent }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetParent() types.Entity { return latestRtt.parent }

func (latestRtt *IpSlaStats_SlaOperEntry_RttInfo_LatestRtt) GetParentYangName() string { return "rtt-info" }

// IpSlaStats_SlaOperEntry_RttInfo_TimeToLive
// Time-to-live for the SLA operation
type IpSlaStats_SlaOperEntry_RttInfo_TimeToLive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time to live value. The type is interface{} with range:
    // 0..18446744073709551615.
    Ttl interface{}

    // Time to live unbound. The type is interface{}.
    Forever interface{}
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetFilter() yfilter.YFilter { return timeToLive.YFilter }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) SetFilter(yf yfilter.YFilter) { timeToLive.YFilter = yf }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetGoName(yname string) string {
    if yname == "ttl" { return "Ttl" }
    if yname == "forever" { return "Forever" }
    return ""
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetSegmentPath() string {
    return "time-to-live"
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl"] = timeToLive.Ttl
    leafs["forever"] = timeToLive.Forever
    return leafs
}

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetBundleName() string { return "cisco_ios_xe" }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetYangName() string { return "time-to-live" }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) SetParent(parent types.Entity) { timeToLive.parent = parent }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetParent() types.Entity { return timeToLive.parent }

func (timeToLive *IpSlaStats_SlaOperEntry_RttInfo_TimeToLive) GetParentYangName() string { return "rtt-info" }

// IpSlaStats_SlaOperEntry_MeasureStats
// Measured statistics
type IpSlaStats_SlaOperEntry_MeasureStats struct {
    parent types.Entity
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

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetFilter() yfilter.YFilter { return measureStats.YFilter }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) SetFilter(yf yfilter.YFilter) { measureStats.YFilter = yf }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetGoName(yname string) string {
    if yname == "intv-start-time" { return "IntvStartTime" }
    if yname == "init-count" { return "InitCount" }
    if yname == "complete-count" { return "CompleteCount" }
    if yname == "valid" { return "Valid" }
    return ""
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetSegmentPath() string {
    return "measure-stats"
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intv-start-time"] = measureStats.IntvStartTime
    leafs["init-count"] = measureStats.InitCount
    leafs["complete-count"] = measureStats.CompleteCount
    leafs["valid"] = measureStats.Valid
    return leafs
}

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetBundleName() string { return "cisco_ios_xe" }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetYangName() string { return "measure-stats" }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) SetParent(parent types.Entity) { measureStats.parent = parent }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetParent() types.Entity { return measureStats.parent }

func (measureStats *IpSlaStats_SlaOperEntry_MeasureStats) GetParentYangName() string { return "sla-oper-entry" }

// IpSlaStats_SlaOperEntry_Stats
// Statistics
type IpSlaStats_SlaOperEntry_Stats struct {
    parent types.Entity
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

func (stats *IpSlaStats_SlaOperEntry_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *IpSlaStats_SlaOperEntry_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetGoName(yname string) string {
    if yname == "rtt" { return "Rtt" }
    if yname == "oneway-latency" { return "OnewayLatency" }
    if yname == "jitter" { return "Jitter" }
    if yname == "over-threshold" { return "OverThreshold" }
    if yname == "packet-loss" { return "PacketLoss" }
    if yname == "icmp-packet-loss" { return "IcmpPacketLoss" }
    if yname == "voice-score" { return "VoiceScore" }
    return ""
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rtt" {
        return &stats.Rtt
    }
    if childYangName == "oneway-latency" {
        return &stats.OnewayLatency
    }
    if childYangName == "jitter" {
        return &stats.Jitter
    }
    if childYangName == "over-threshold" {
        return &stats.OverThreshold
    }
    if childYangName == "packet-loss" {
        return &stats.PacketLoss
    }
    if childYangName == "icmp-packet-loss" {
        return &stats.IcmpPacketLoss
    }
    if childYangName == "voice-score" {
        return &stats.VoiceScore
    }
    return nil
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rtt"] = &stats.Rtt
    children["oneway-latency"] = &stats.OnewayLatency
    children["jitter"] = &stats.Jitter
    children["over-threshold"] = &stats.OverThreshold
    children["packet-loss"] = &stats.PacketLoss
    children["icmp-packet-loss"] = &stats.IcmpPacketLoss
    children["voice-score"] = &stats.VoiceScore
    return children
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stats *IpSlaStats_SlaOperEntry_Stats) GetBundleName() string { return "cisco_ios_xe" }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetYangName() string { return "stats" }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stats *IpSlaStats_SlaOperEntry_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetParent() types.Entity { return stats.parent }

func (stats *IpSlaStats_SlaOperEntry_Stats) GetParentYangName() string { return "sla-oper-entry" }

// IpSlaStats_SlaOperEntry_Stats_Rtt
// RTT value
type IpSlaStats_SlaOperEntry_Stats_Rtt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RTT count. The type is interface{} with range: 0..4294967295.
    RttCount interface{}

    // Timing information.
    SlaTimeValues IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetFilter() yfilter.YFilter { return rtt.YFilter }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) SetFilter(yf yfilter.YFilter) { rtt.YFilter = yf }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetGoName(yname string) string {
    if yname == "rtt-count" { return "RttCount" }
    if yname == "sla-time-values" { return "SlaTimeValues" }
    return ""
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetSegmentPath() string {
    return "rtt"
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sla-time-values" {
        return &rtt.SlaTimeValues
    }
    return nil
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sla-time-values"] = &rtt.SlaTimeValues
    return children
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtt-count"] = rtt.RttCount
    return leafs
}

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetBundleName() string { return "cisco_ios_xe" }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetYangName() string { return "rtt" }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) SetParent(parent types.Entity) { rtt.parent = parent }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetParent() types.Entity { return rtt.parent }

func (rtt *IpSlaStats_SlaOperEntry_Stats_Rtt) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues
// Timing information
type IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues struct {
    parent types.Entity
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

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetFilter() yfilter.YFilter { return slaTimeValues.YFilter }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) SetFilter(yf yfilter.YFilter) { slaTimeValues.YFilter = yf }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    if yname == "accuracy" { return "Accuracy" }
    return ""
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetSegmentPath() string {
    return "sla-time-values"
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = slaTimeValues.Min
    leafs["avg"] = slaTimeValues.Avg
    leafs["max"] = slaTimeValues.Max
    leafs["accuracy"] = slaTimeValues.Accuracy
    return leafs
}

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetBundleName() string { return "cisco_ios_xe" }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetYangName() string { return "sla-time-values" }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) SetParent(parent types.Entity) { slaTimeValues.parent = parent }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetParent() types.Entity { return slaTimeValues.parent }

func (slaTimeValues *IpSlaStats_SlaOperEntry_Stats_Rtt_SlaTimeValues) GetParentYangName() string { return "rtt" }

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency
// Latency information
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sample count. The type is interface{} with range: 0..4294967295.
    SampleCount interface{}

    // Source to Destination for the one-way latency.
    Sd IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd

    // Destination to Source for the one-way latency.
    Ds IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetFilter() yfilter.YFilter { return onewayLatency.YFilter }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) SetFilter(yf yfilter.YFilter) { onewayLatency.YFilter = yf }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetGoName(yname string) string {
    if yname == "sample-count" { return "SampleCount" }
    if yname == "sd" { return "Sd" }
    if yname == "ds" { return "Ds" }
    return ""
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetSegmentPath() string {
    return "oneway-latency"
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sd" {
        return &onewayLatency.Sd
    }
    if childYangName == "ds" {
        return &onewayLatency.Ds
    }
    return nil
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sd"] = &onewayLatency.Sd
    children["ds"] = &onewayLatency.Ds
    return children
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sample-count"] = onewayLatency.SampleCount
    return leafs
}

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetBundleName() string { return "cisco_ios_xe" }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetYangName() string { return "oneway-latency" }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) SetParent(parent types.Entity) { onewayLatency.parent = parent }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetParent() types.Entity { return onewayLatency.parent }

func (onewayLatency *IpSlaStats_SlaOperEntry_Stats_OnewayLatency) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd
// Source to Destination for the one-way latency
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd struct {
    parent types.Entity
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

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetFilter() yfilter.YFilter { return sd.YFilter }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) SetFilter(yf yfilter.YFilter) { sd.YFilter = yf }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    if yname == "accuracy" { return "Accuracy" }
    return ""
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetSegmentPath() string {
    return "sd"
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = sd.Min
    leafs["avg"] = sd.Avg
    leafs["max"] = sd.Max
    leafs["accuracy"] = sd.Accuracy
    return leafs
}

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetBundleName() string { return "cisco_ios_xe" }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetYangName() string { return "sd" }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) SetParent(parent types.Entity) { sd.parent = parent }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetParent() types.Entity { return sd.parent }

func (sd *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Sd) GetParentYangName() string { return "oneway-latency" }

// IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds
// Destination to Source for the one-way latency
type IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds struct {
    parent types.Entity
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

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetFilter() yfilter.YFilter { return ds.YFilter }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) SetFilter(yf yfilter.YFilter) { ds.YFilter = yf }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    if yname == "accuracy" { return "Accuracy" }
    return ""
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetSegmentPath() string {
    return "ds"
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = ds.Min
    leafs["avg"] = ds.Avg
    leafs["max"] = ds.Max
    leafs["accuracy"] = ds.Accuracy
    return leafs
}

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetBundleName() string { return "cisco_ios_xe" }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetYangName() string { return "ds" }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) SetParent(parent types.Entity) { ds.parent = parent }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetParent() types.Entity { return ds.parent }

func (ds *IpSlaStats_SlaOperEntry_Stats_OnewayLatency_Ds) GetParentYangName() string { return "oneway-latency" }

// IpSlaStats_SlaOperEntry_Stats_Jitter
// Jitter information
type IpSlaStats_SlaOperEntry_Stats_Jitter struct {
    parent types.Entity
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

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetFilter() yfilter.YFilter { return jitter.YFilter }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) SetFilter(yf yfilter.YFilter) { jitter.YFilter = yf }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetGoName(yname string) string {
    if yname == "sd-sample-count" { return "SdSampleCount" }
    if yname == "ds-sample-count" { return "DsSampleCount" }
    if yname == "sd" { return "Sd" }
    if yname == "ds" { return "Ds" }
    return ""
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetSegmentPath() string {
    return "jitter"
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sd" {
        return &jitter.Sd
    }
    if childYangName == "ds" {
        return &jitter.Ds
    }
    return nil
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sd"] = &jitter.Sd
    children["ds"] = &jitter.Ds
    return children
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sd-sample-count"] = jitter.SdSampleCount
    leafs["ds-sample-count"] = jitter.DsSampleCount
    return leafs
}

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetBundleName() string { return "cisco_ios_xe" }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetYangName() string { return "jitter" }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) SetParent(parent types.Entity) { jitter.parent = parent }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetParent() types.Entity { return jitter.parent }

func (jitter *IpSlaStats_SlaOperEntry_Stats_Jitter) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_Jitter_Sd
// Source to Destination for the jitter
type IpSlaStats_SlaOperEntry_Stats_Jitter_Sd struct {
    parent types.Entity
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

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetFilter() yfilter.YFilter { return sd.YFilter }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) SetFilter(yf yfilter.YFilter) { sd.YFilter = yf }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    if yname == "accuracy" { return "Accuracy" }
    return ""
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetSegmentPath() string {
    return "sd"
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = sd.Min
    leafs["avg"] = sd.Avg
    leafs["max"] = sd.Max
    leafs["accuracy"] = sd.Accuracy
    return leafs
}

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetBundleName() string { return "cisco_ios_xe" }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetYangName() string { return "sd" }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) SetParent(parent types.Entity) { sd.parent = parent }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetParent() types.Entity { return sd.parent }

func (sd *IpSlaStats_SlaOperEntry_Stats_Jitter_Sd) GetParentYangName() string { return "jitter" }

// IpSlaStats_SlaOperEntry_Stats_Jitter_Ds
// Destination to Source for the jitter
type IpSlaStats_SlaOperEntry_Stats_Jitter_Ds struct {
    parent types.Entity
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

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetFilter() yfilter.YFilter { return ds.YFilter }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) SetFilter(yf yfilter.YFilter) { ds.YFilter = yf }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    if yname == "accuracy" { return "Accuracy" }
    return ""
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetSegmentPath() string {
    return "ds"
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = ds.Min
    leafs["avg"] = ds.Avg
    leafs["max"] = ds.Max
    leafs["accuracy"] = ds.Accuracy
    return leafs
}

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetBundleName() string { return "cisco_ios_xe" }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetYangName() string { return "ds" }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) SetParent(parent types.Entity) { ds.parent = parent }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetParent() types.Entity { return ds.parent }

func (ds *IpSlaStats_SlaOperEntry_Stats_Jitter_Ds) GetParentYangName() string { return "jitter" }

// IpSlaStats_SlaOperEntry_Stats_OverThreshold
// Over threshold information
type IpSlaStats_SlaOperEntry_Stats_OverThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Round Trip Time (RTT) over threshold count (the number of times that the
    // RTT was over the configured threshold). The type is interface{} with range:
    // 0..4294967295.
    RttCount interface{}

    // Round Trip Time over threshold percentage (the percentage that the RTT was
    // over the configured threshold). The type is interface{} with range: 0..255.
    Percent interface{}
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetFilter() yfilter.YFilter { return overThreshold.YFilter }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) SetFilter(yf yfilter.YFilter) { overThreshold.YFilter = yf }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetGoName(yname string) string {
    if yname == "rtt-count" { return "RttCount" }
    if yname == "percent" { return "Percent" }
    return ""
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetSegmentPath() string {
    return "over-threshold"
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rtt-count"] = overThreshold.RttCount
    leafs["percent"] = overThreshold.Percent
    return leafs
}

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetBundleName() string { return "cisco_ios_xe" }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetYangName() string { return "over-threshold" }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) SetParent(parent types.Entity) { overThreshold.parent = parent }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetParent() types.Entity { return overThreshold.parent }

func (overThreshold *IpSlaStats_SlaOperEntry_Stats_OverThreshold) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_PacketLoss
// Packet loss information
type IpSlaStats_SlaOperEntry_Stats_PacketLoss struct {
    parent types.Entity
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

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetFilter() yfilter.YFilter { return packetLoss.YFilter }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) SetFilter(yf yfilter.YFilter) { packetLoss.YFilter = yf }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetGoName(yname string) string {
    if yname == "unprocessed-packets" { return "UnprocessedPackets" }
    if yname == "sd-count" { return "SdCount" }
    if yname == "ds-count" { return "DsCount" }
    if yname == "out-of-sequence" { return "OutOfSequence" }
    if yname == "drops" { return "Drops" }
    if yname == "late-arrivals" { return "LateArrivals" }
    if yname == "skipped-packets" { return "SkippedPackets" }
    if yname == "sd-loss" { return "SdLoss" }
    if yname == "ds-loss" { return "DsLoss" }
    return ""
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetSegmentPath() string {
    return "packet-loss"
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sd-loss" {
        return &packetLoss.SdLoss
    }
    if childYangName == "ds-loss" {
        return &packetLoss.DsLoss
    }
    return nil
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sd-loss"] = &packetLoss.SdLoss
    children["ds-loss"] = &packetLoss.DsLoss
    return children
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unprocessed-packets"] = packetLoss.UnprocessedPackets
    leafs["sd-count"] = packetLoss.SdCount
    leafs["ds-count"] = packetLoss.DsCount
    leafs["out-of-sequence"] = packetLoss.OutOfSequence
    leafs["drops"] = packetLoss.Drops
    leafs["late-arrivals"] = packetLoss.LateArrivals
    leafs["skipped-packets"] = packetLoss.SkippedPackets
    return leafs
}

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetBundleName() string { return "cisco_ios_xe" }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetYangName() string { return "packet-loss" }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) SetParent(parent types.Entity) { packetLoss.parent = parent }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetParent() types.Entity { return packetLoss.parent }

func (packetLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss
// Source to Destination packet loss details
type IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss struct {
    parent types.Entity
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

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetFilter() yfilter.YFilter { return sdLoss.YFilter }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) SetFilter(yf yfilter.YFilter) { sdLoss.YFilter = yf }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetGoName(yname string) string {
    if yname == "loss-period-count" { return "LossPeriodCount" }
    if yname == "loss-period-len-min" { return "LossPeriodLenMin" }
    if yname == "loss-period-len-max" { return "LossPeriodLenMax" }
    if yname == "inter-loss-period-len-min" { return "InterLossPeriodLenMin" }
    if yname == "inter-loss-period-len-max" { return "InterLossPeriodLenMax" }
    return ""
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetSegmentPath() string {
    return "sd-loss"
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["loss-period-count"] = sdLoss.LossPeriodCount
    leafs["loss-period-len-min"] = sdLoss.LossPeriodLenMin
    leafs["loss-period-len-max"] = sdLoss.LossPeriodLenMax
    leafs["inter-loss-period-len-min"] = sdLoss.InterLossPeriodLenMin
    leafs["inter-loss-period-len-max"] = sdLoss.InterLossPeriodLenMax
    return leafs
}

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetBundleName() string { return "cisco_ios_xe" }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetYangName() string { return "sd-loss" }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) SetParent(parent types.Entity) { sdLoss.parent = parent }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetParent() types.Entity { return sdLoss.parent }

func (sdLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_SdLoss) GetParentYangName() string { return "packet-loss" }

// IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss
// Destination to Source packet loss details
type IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss struct {
    parent types.Entity
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

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetFilter() yfilter.YFilter { return dsLoss.YFilter }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) SetFilter(yf yfilter.YFilter) { dsLoss.YFilter = yf }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetGoName(yname string) string {
    if yname == "loss-period-count" { return "LossPeriodCount" }
    if yname == "loss-period-len-min" { return "LossPeriodLenMin" }
    if yname == "loss-period-len-max" { return "LossPeriodLenMax" }
    if yname == "inter-loss-period-len-min" { return "InterLossPeriodLenMin" }
    if yname == "inter-loss-period-len-max" { return "InterLossPeriodLenMax" }
    return ""
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetSegmentPath() string {
    return "ds-loss"
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["loss-period-count"] = dsLoss.LossPeriodCount
    leafs["loss-period-len-min"] = dsLoss.LossPeriodLenMin
    leafs["loss-period-len-max"] = dsLoss.LossPeriodLenMax
    leafs["inter-loss-period-len-min"] = dsLoss.InterLossPeriodLenMin
    leafs["inter-loss-period-len-max"] = dsLoss.InterLossPeriodLenMax
    return leafs
}

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetBundleName() string { return "cisco_ios_xe" }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetYangName() string { return "ds-loss" }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) SetParent(parent types.Entity) { dsLoss.parent = parent }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetParent() types.Entity { return dsLoss.parent }

func (dsLoss *IpSlaStats_SlaOperEntry_Stats_PacketLoss_DsLoss) GetParentYangName() string { return "packet-loss" }

// IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss
// ICMP packet loss information
type IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss struct {
    parent types.Entity
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

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetFilter() yfilter.YFilter { return icmpPacketLoss.YFilter }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) SetFilter(yf yfilter.YFilter) { icmpPacketLoss.YFilter = yf }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetGoName(yname string) string {
    if yname == "late-arrivals" { return "LateArrivals" }
    if yname == "out-of-sequence" { return "OutOfSequence" }
    if yname == "out-of-sequence-sd" { return "OutOfSequenceSd" }
    if yname == "out-of-sequence-ds" { return "OutOfSequenceDs" }
    if yname == "out-of-sequence-both" { return "OutOfSequenceBoth" }
    if yname == "skipped-packets" { return "SkippedPackets" }
    if yname == "unprocessed-packets" { return "UnprocessedPackets" }
    if yname == "packet-loss" { return "PacketLoss" }
    if yname == "loss-period-count" { return "LossPeriodCount" }
    if yname == "loss-period-len-min" { return "LossPeriodLenMin" }
    if yname == "loss-period-len-max" { return "LossPeriodLenMax" }
    if yname == "inter-loss-period-len-min" { return "InterLossPeriodLenMin" }
    if yname == "inter-loss-period-len-max" { return "InterLossPeriodLenMax" }
    return ""
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetSegmentPath() string {
    return "icmp-packet-loss"
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["late-arrivals"] = icmpPacketLoss.LateArrivals
    leafs["out-of-sequence"] = icmpPacketLoss.OutOfSequence
    leafs["out-of-sequence-sd"] = icmpPacketLoss.OutOfSequenceSd
    leafs["out-of-sequence-ds"] = icmpPacketLoss.OutOfSequenceDs
    leafs["out-of-sequence-both"] = icmpPacketLoss.OutOfSequenceBoth
    leafs["skipped-packets"] = icmpPacketLoss.SkippedPackets
    leafs["unprocessed-packets"] = icmpPacketLoss.UnprocessedPackets
    leafs["packet-loss"] = icmpPacketLoss.PacketLoss
    leafs["loss-period-count"] = icmpPacketLoss.LossPeriodCount
    leafs["loss-period-len-min"] = icmpPacketLoss.LossPeriodLenMin
    leafs["loss-period-len-max"] = icmpPacketLoss.LossPeriodLenMax
    leafs["inter-loss-period-len-min"] = icmpPacketLoss.InterLossPeriodLenMin
    leafs["inter-loss-period-len-max"] = icmpPacketLoss.InterLossPeriodLenMax
    return leafs
}

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetBundleName() string { return "cisco_ios_xe" }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetYangName() string { return "icmp-packet-loss" }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) SetParent(parent types.Entity) { icmpPacketLoss.parent = parent }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetParent() types.Entity { return icmpPacketLoss.parent }

func (icmpPacketLoss *IpSlaStats_SlaOperEntry_Stats_IcmpPacketLoss) GetParentYangName() string { return "stats" }

// IpSlaStats_SlaOperEntry_Stats_VoiceScore
// Voice score information
type IpSlaStats_SlaOperEntry_Stats_VoiceScore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Calculated planning impairment factor. The type is interface{} with range:
    // 0..4294967295.
    Icpif interface{}

    // Mean opinion score. The type is interface{} with range: 0..4294967295.
    Mos interface{}
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetFilter() yfilter.YFilter { return voiceScore.YFilter }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) SetFilter(yf yfilter.YFilter) { voiceScore.YFilter = yf }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetGoName(yname string) string {
    if yname == "icpif" { return "Icpif" }
    if yname == "mos" { return "Mos" }
    return ""
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetSegmentPath() string {
    return "voice-score"
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icpif"] = voiceScore.Icpif
    leafs["mos"] = voiceScore.Mos
    return leafs
}

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetBundleName() string { return "cisco_ios_xe" }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetYangName() string { return "voice-score" }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) SetParent(parent types.Entity) { voiceScore.parent = parent }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetParent() types.Entity { return voiceScore.parent }

func (voiceScore *IpSlaStats_SlaOperEntry_Stats_VoiceScore) GetParentYangName() string { return "stats" }

