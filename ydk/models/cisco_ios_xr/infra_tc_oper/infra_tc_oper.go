// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-tc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   traffic-collector: Global Traffic Collector configuration
//     commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
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

func (trafficCollector *TrafficCollector) GetFilter() yfilter.YFilter { return trafficCollector.YFilter }

func (trafficCollector *TrafficCollector) SetFilter(yf yfilter.YFilter) { trafficCollector.YFilter = yf }

func (trafficCollector *TrafficCollector) GetGoName(yname string) string {
    if yname == "external-interfaces" { return "ExternalInterfaces" }
    if yname == "summary" { return "Summary" }
    if yname == "vrf-table" { return "VrfTable" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (trafficCollector *TrafficCollector) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-tc-oper:traffic-collector"
}

func (trafficCollector *TrafficCollector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "external-interfaces" {
        return &trafficCollector.ExternalInterfaces
    }
    if childYangName == "summary" {
        return &trafficCollector.Summary
    }
    if childYangName == "vrf-table" {
        return &trafficCollector.VrfTable
    }
    if childYangName == "afs" {
        return &trafficCollector.Afs
    }
    return nil
}

func (trafficCollector *TrafficCollector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["external-interfaces"] = &trafficCollector.ExternalInterfaces
    children["summary"] = &trafficCollector.Summary
    children["vrf-table"] = &trafficCollector.VrfTable
    children["afs"] = &trafficCollector.Afs
    return children
}

func (trafficCollector *TrafficCollector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trafficCollector *TrafficCollector) GetBundleName() string { return "cisco_ios_xr" }

func (trafficCollector *TrafficCollector) GetYangName() string { return "traffic-collector" }

func (trafficCollector *TrafficCollector) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficCollector *TrafficCollector) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficCollector *TrafficCollector) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficCollector *TrafficCollector) SetParent(parent types.Entity) { trafficCollector.parent = parent }

func (trafficCollector *TrafficCollector) GetParent() types.Entity { return trafficCollector.parent }

func (trafficCollector *TrafficCollector) GetParentYangName() string { return "Cisco-IOS-XR-infra-tc-oper" }

// TrafficCollector_ExternalInterfaces
// External Interface
type TrafficCollector_ExternalInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // External Interface . The type is slice of
    // TrafficCollector_ExternalInterfaces_ExternalInterface.
    ExternalInterface []TrafficCollector_ExternalInterfaces_ExternalInterface
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetFilter() yfilter.YFilter { return externalInterfaces.YFilter }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) SetFilter(yf yfilter.YFilter) { externalInterfaces.YFilter = yf }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetGoName(yname string) string {
    if yname == "external-interface" { return "ExternalInterface" }
    return ""
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetSegmentPath() string {
    return "external-interfaces"
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "external-interface" {
        for _, c := range externalInterfaces.ExternalInterface {
            if externalInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_ExternalInterfaces_ExternalInterface{}
        externalInterfaces.ExternalInterface = append(externalInterfaces.ExternalInterface, child)
        return &externalInterfaces.ExternalInterface[len(externalInterfaces.ExternalInterface)-1]
    }
    return nil
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range externalInterfaces.ExternalInterface {
        children[externalInterfaces.ExternalInterface[i].GetSegmentPath()] = &externalInterfaces.ExternalInterface[i]
    }
    return children
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetYangName() string { return "external-interfaces" }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) SetParent(parent types.Entity) { externalInterfaces.parent = parent }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetParent() types.Entity { return externalInterfaces.parent }

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetParentYangName() string { return "traffic-collector" }

// TrafficCollector_ExternalInterfaces_ExternalInterface
// External Interface 
type TrafficCollector_ExternalInterfaces_ExternalInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetFilter() yfilter.YFilter { return externalInterface.YFilter }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) SetFilter(yf yfilter.YFilter) { externalInterface.YFilter = yf }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "vrfid" { return "Vrfid" }
    if yname == "is-interface-enabled" { return "IsInterfaceEnabled" }
    return ""
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetSegmentPath() string {
    return "external-interface" + "[interface-name='" + fmt.Sprintf("%v", externalInterface.InterfaceName) + "']"
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = externalInterface.InterfaceName
    leafs["interface-name-xr"] = externalInterface.InterfaceNameXr
    leafs["interface-handle"] = externalInterface.InterfaceHandle
    leafs["vrfid"] = externalInterface.Vrfid
    leafs["is-interface-enabled"] = externalInterface.IsInterfaceEnabled
    return leafs
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetBundleName() string { return "cisco_ios_xr" }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetYangName() string { return "external-interface" }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) SetParent(parent types.Entity) { externalInterface.parent = parent }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetParent() types.Entity { return externalInterface.parent }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetParentYangName() string { return "external-interfaces" }

// TrafficCollector_Summary
// Traffic Collector summary
type TrafficCollector_Summary struct {
    parent types.Entity
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
    VrfStatistic []TrafficCollector_Summary_VrfStatistic

    // Statistics per message type for STAT collector. The type is slice of
    // TrafficCollector_Summary_CollectionMessageStatistic.
    CollectionMessageStatistic []TrafficCollector_Summary_CollectionMessageStatistic

    // Statistics per message type for Chkpt. The type is slice of
    // TrafficCollector_Summary_CheckpointMessageStatistic.
    CheckpointMessageStatistic []TrafficCollector_Summary_CheckpointMessageStatistic
}

func (summary *TrafficCollector_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *TrafficCollector_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *TrafficCollector_Summary) GetGoName(yname string) string {
    if yname == "collection-interval" { return "CollectionInterval" }
    if yname == "collection-timer-is-running" { return "CollectionTimerIsRunning" }
    if yname == "timeout-interval" { return "TimeoutInterval" }
    if yname == "timeout-timer-is-running" { return "TimeoutTimerIsRunning" }
    if yname == "history-size" { return "HistorySize" }
    if yname == "database-statistics-external-interface" { return "DatabaseStatisticsExternalInterface" }
    if yname == "vrf-statistic" { return "VrfStatistic" }
    if yname == "collection-message-statistic" { return "CollectionMessageStatistic" }
    if yname == "checkpoint-message-statistic" { return "CheckpointMessageStatistic" }
    return ""
}

func (summary *TrafficCollector_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *TrafficCollector_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "database-statistics-external-interface" {
        return &summary.DatabaseStatisticsExternalInterface
    }
    if childYangName == "vrf-statistic" {
        for _, c := range summary.VrfStatistic {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Summary_VrfStatistic{}
        summary.VrfStatistic = append(summary.VrfStatistic, child)
        return &summary.VrfStatistic[len(summary.VrfStatistic)-1]
    }
    if childYangName == "collection-message-statistic" {
        for _, c := range summary.CollectionMessageStatistic {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Summary_CollectionMessageStatistic{}
        summary.CollectionMessageStatistic = append(summary.CollectionMessageStatistic, child)
        return &summary.CollectionMessageStatistic[len(summary.CollectionMessageStatistic)-1]
    }
    if childYangName == "checkpoint-message-statistic" {
        for _, c := range summary.CheckpointMessageStatistic {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Summary_CheckpointMessageStatistic{}
        summary.CheckpointMessageStatistic = append(summary.CheckpointMessageStatistic, child)
        return &summary.CheckpointMessageStatistic[len(summary.CheckpointMessageStatistic)-1]
    }
    return nil
}

func (summary *TrafficCollector_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["database-statistics-external-interface"] = &summary.DatabaseStatisticsExternalInterface
    for i := range summary.VrfStatistic {
        children[summary.VrfStatistic[i].GetSegmentPath()] = &summary.VrfStatistic[i]
    }
    for i := range summary.CollectionMessageStatistic {
        children[summary.CollectionMessageStatistic[i].GetSegmentPath()] = &summary.CollectionMessageStatistic[i]
    }
    for i := range summary.CheckpointMessageStatistic {
        children[summary.CheckpointMessageStatistic[i].GetSegmentPath()] = &summary.CheckpointMessageStatistic[i]
    }
    return children
}

func (summary *TrafficCollector_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["collection-interval"] = summary.CollectionInterval
    leafs["collection-timer-is-running"] = summary.CollectionTimerIsRunning
    leafs["timeout-interval"] = summary.TimeoutInterval
    leafs["timeout-timer-is-running"] = summary.TimeoutTimerIsRunning
    leafs["history-size"] = summary.HistorySize
    return leafs
}

func (summary *TrafficCollector_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *TrafficCollector_Summary) GetYangName() string { return "summary" }

func (summary *TrafficCollector_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *TrafficCollector_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *TrafficCollector_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *TrafficCollector_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *TrafficCollector_Summary) GetParent() types.Entity { return summary.parent }

func (summary *TrafficCollector_Summary) GetParentYangName() string { return "traffic-collector" }

// TrafficCollector_Summary_DatabaseStatisticsExternalInterface
// Database statistics for External Interface
type TrafficCollector_Summary_DatabaseStatisticsExternalInterface struct {
    parent types.Entity
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

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetFilter() yfilter.YFilter { return databaseStatisticsExternalInterface.YFilter }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) SetFilter(yf yfilter.YFilter) { databaseStatisticsExternalInterface.YFilter = yf }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetGoName(yname string) string {
    if yname == "number-of-entries" { return "NumberOfEntries" }
    if yname == "number-of-stale-entries" { return "NumberOfStaleEntries" }
    if yname == "number-of-add-o-perations" { return "NumberOfAddOPerations" }
    if yname == "number-of-delete-o-perations" { return "NumberOfDeleteOPerations" }
    return ""
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetSegmentPath() string {
    return "database-statistics-external-interface"
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-entries"] = databaseStatisticsExternalInterface.NumberOfEntries
    leafs["number-of-stale-entries"] = databaseStatisticsExternalInterface.NumberOfStaleEntries
    leafs["number-of-add-o-perations"] = databaseStatisticsExternalInterface.NumberOfAddOPerations
    leafs["number-of-delete-o-perations"] = databaseStatisticsExternalInterface.NumberOfDeleteOPerations
    return leafs
}

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetBundleName() string { return "cisco_ios_xr" }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetYangName() string { return "database-statistics-external-interface" }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) SetParent(parent types.Entity) { databaseStatisticsExternalInterface.parent = parent }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetParent() types.Entity { return databaseStatisticsExternalInterface.parent }

func (databaseStatisticsExternalInterface *TrafficCollector_Summary_DatabaseStatisticsExternalInterface) GetParentYangName() string { return "summary" }

// TrafficCollector_Summary_VrfStatistic
// VRF table statistics
type TrafficCollector_Summary_VrfStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    VrfName interface{}

    // Database statistics for IPv4 table.
    DatabaseStatisticsIpv4 TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4

    // Database statistics for Tunnel table.
    DatabaseStatisticsTunnel TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetFilter() yfilter.YFilter { return vrfStatistic.YFilter }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) SetFilter(yf yfilter.YFilter) { vrfStatistic.YFilter = yf }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "database-statistics-ipv4" { return "DatabaseStatisticsIpv4" }
    if yname == "database-statistics-tunnel" { return "DatabaseStatisticsTunnel" }
    return ""
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetSegmentPath() string {
    return "vrf-statistic"
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "database-statistics-ipv4" {
        return &vrfStatistic.DatabaseStatisticsIpv4
    }
    if childYangName == "database-statistics-tunnel" {
        return &vrfStatistic.DatabaseStatisticsTunnel
    }
    return nil
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["database-statistics-ipv4"] = &vrfStatistic.DatabaseStatisticsIpv4
    children["database-statistics-tunnel"] = &vrfStatistic.DatabaseStatisticsTunnel
    return children
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrfStatistic.VrfName
    return leafs
}

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetYangName() string { return "vrf-statistic" }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) SetParent(parent types.Entity) { vrfStatistic.parent = parent }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetParent() types.Entity { return vrfStatistic.parent }

func (vrfStatistic *TrafficCollector_Summary_VrfStatistic) GetParentYangName() string { return "summary" }

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4
// Database statistics for IPv4 table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4 struct {
    parent types.Entity
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

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetFilter() yfilter.YFilter { return databaseStatisticsIpv4.YFilter }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) SetFilter(yf yfilter.YFilter) { databaseStatisticsIpv4.YFilter = yf }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetGoName(yname string) string {
    if yname == "number-of-entries" { return "NumberOfEntries" }
    if yname == "number-of-stale-entries" { return "NumberOfStaleEntries" }
    if yname == "number-of-add-o-perations" { return "NumberOfAddOPerations" }
    if yname == "number-of-delete-o-perations" { return "NumberOfDeleteOPerations" }
    return ""
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetSegmentPath() string {
    return "database-statistics-ipv4"
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-entries"] = databaseStatisticsIpv4.NumberOfEntries
    leafs["number-of-stale-entries"] = databaseStatisticsIpv4.NumberOfStaleEntries
    leafs["number-of-add-o-perations"] = databaseStatisticsIpv4.NumberOfAddOPerations
    leafs["number-of-delete-o-perations"] = databaseStatisticsIpv4.NumberOfDeleteOPerations
    return leafs
}

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetYangName() string { return "database-statistics-ipv4" }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) SetParent(parent types.Entity) { databaseStatisticsIpv4.parent = parent }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetParent() types.Entity { return databaseStatisticsIpv4.parent }

func (databaseStatisticsIpv4 *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsIpv4) GetParentYangName() string { return "vrf-statistic" }

// TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel
// Database statistics for Tunnel table
type TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel struct {
    parent types.Entity
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

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetFilter() yfilter.YFilter { return databaseStatisticsTunnel.YFilter }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) SetFilter(yf yfilter.YFilter) { databaseStatisticsTunnel.YFilter = yf }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetGoName(yname string) string {
    if yname == "number-of-entries" { return "NumberOfEntries" }
    if yname == "number-of-stale-entries" { return "NumberOfStaleEntries" }
    if yname == "number-of-add-o-perations" { return "NumberOfAddOPerations" }
    if yname == "number-of-delete-o-perations" { return "NumberOfDeleteOPerations" }
    return ""
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetSegmentPath() string {
    return "database-statistics-tunnel"
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-entries"] = databaseStatisticsTunnel.NumberOfEntries
    leafs["number-of-stale-entries"] = databaseStatisticsTunnel.NumberOfStaleEntries
    leafs["number-of-add-o-perations"] = databaseStatisticsTunnel.NumberOfAddOPerations
    leafs["number-of-delete-o-perations"] = databaseStatisticsTunnel.NumberOfDeleteOPerations
    return leafs
}

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetBundleName() string { return "cisco_ios_xr" }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetYangName() string { return "database-statistics-tunnel" }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) SetParent(parent types.Entity) { databaseStatisticsTunnel.parent = parent }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetParent() types.Entity { return databaseStatisticsTunnel.parent }

func (databaseStatisticsTunnel *TrafficCollector_Summary_VrfStatistic_DatabaseStatisticsTunnel) GetParentYangName() string { return "vrf-statistic" }

// TrafficCollector_Summary_CollectionMessageStatistic
// Statistics per message type for STAT collector
type TrafficCollector_Summary_CollectionMessageStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetFilter() yfilter.YFilter { return collectionMessageStatistic.YFilter }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) SetFilter(yf yfilter.YFilter) { collectionMessageStatistic.YFilter = yf }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetGoName(yname string) string {
    if yname == "packet-sent" { return "PacketSent" }
    if yname == "byte-sent" { return "ByteSent" }
    if yname == "packet-received" { return "PacketReceived" }
    if yname == "byte-received" { return "ByteReceived" }
    if yname == "maximum-roundtrip-latency" { return "MaximumRoundtripLatency" }
    if yname == "maimum-latency-timestamp" { return "MaimumLatencyTimestamp" }
    return ""
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetSegmentPath() string {
    return "collection-message-statistic"
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-sent"] = collectionMessageStatistic.PacketSent
    leafs["byte-sent"] = collectionMessageStatistic.ByteSent
    leafs["packet-received"] = collectionMessageStatistic.PacketReceived
    leafs["byte-received"] = collectionMessageStatistic.ByteReceived
    leafs["maximum-roundtrip-latency"] = collectionMessageStatistic.MaximumRoundtripLatency
    leafs["maimum-latency-timestamp"] = collectionMessageStatistic.MaimumLatencyTimestamp
    return leafs
}

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetYangName() string { return "collection-message-statistic" }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) SetParent(parent types.Entity) { collectionMessageStatistic.parent = parent }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetParent() types.Entity { return collectionMessageStatistic.parent }

func (collectionMessageStatistic *TrafficCollector_Summary_CollectionMessageStatistic) GetParentYangName() string { return "summary" }

// TrafficCollector_Summary_CheckpointMessageStatistic
// Statistics per message type for Chkpt
type TrafficCollector_Summary_CheckpointMessageStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetFilter() yfilter.YFilter { return checkpointMessageStatistic.YFilter }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) SetFilter(yf yfilter.YFilter) { checkpointMessageStatistic.YFilter = yf }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetGoName(yname string) string {
    if yname == "packet-sent" { return "PacketSent" }
    if yname == "byte-sent" { return "ByteSent" }
    if yname == "packet-received" { return "PacketReceived" }
    if yname == "byte-received" { return "ByteReceived" }
    if yname == "maximum-roundtrip-latency" { return "MaximumRoundtripLatency" }
    if yname == "maimum-latency-timestamp" { return "MaimumLatencyTimestamp" }
    return ""
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetSegmentPath() string {
    return "checkpoint-message-statistic"
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-sent"] = checkpointMessageStatistic.PacketSent
    leafs["byte-sent"] = checkpointMessageStatistic.ByteSent
    leafs["packet-received"] = checkpointMessageStatistic.PacketReceived
    leafs["byte-received"] = checkpointMessageStatistic.ByteReceived
    leafs["maximum-roundtrip-latency"] = checkpointMessageStatistic.MaximumRoundtripLatency
    leafs["maimum-latency-timestamp"] = checkpointMessageStatistic.MaimumLatencyTimestamp
    return leafs
}

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetYangName() string { return "checkpoint-message-statistic" }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) SetParent(parent types.Entity) { checkpointMessageStatistic.parent = parent }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetParent() types.Entity { return checkpointMessageStatistic.parent }

func (checkpointMessageStatistic *TrafficCollector_Summary_CheckpointMessageStatistic) GetParentYangName() string { return "summary" }

// TrafficCollector_VrfTable
// VRF specific operational data
type TrafficCollector_VrfTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DefaultVRF specific operational data.
    DefaultVrf TrafficCollector_VrfTable_DefaultVrf
}

func (vrfTable *TrafficCollector_VrfTable) GetFilter() yfilter.YFilter { return vrfTable.YFilter }

func (vrfTable *TrafficCollector_VrfTable) SetFilter(yf yfilter.YFilter) { vrfTable.YFilter = yf }

func (vrfTable *TrafficCollector_VrfTable) GetGoName(yname string) string {
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (vrfTable *TrafficCollector_VrfTable) GetSegmentPath() string {
    return "vrf-table"
}

func (vrfTable *TrafficCollector_VrfTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-vrf" {
        return &vrfTable.DefaultVrf
    }
    return nil
}

func (vrfTable *TrafficCollector_VrfTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-vrf"] = &vrfTable.DefaultVrf
    return children
}

func (vrfTable *TrafficCollector_VrfTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfTable *TrafficCollector_VrfTable) GetBundleName() string { return "cisco_ios_xr" }

func (vrfTable *TrafficCollector_VrfTable) GetYangName() string { return "vrf-table" }

func (vrfTable *TrafficCollector_VrfTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfTable *TrafficCollector_VrfTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfTable *TrafficCollector_VrfTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfTable *TrafficCollector_VrfTable) SetParent(parent types.Entity) { vrfTable.parent = parent }

func (vrfTable *TrafficCollector_VrfTable) GetParent() types.Entity { return vrfTable.parent }

func (vrfTable *TrafficCollector_VrfTable) GetParentYangName() string { return "traffic-collector" }

// TrafficCollector_VrfTable_DefaultVrf
// DefaultVRF specific operational data
type TrafficCollector_VrfTable_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family specific operational data.
    Afs TrafficCollector_VrfTable_DefaultVrf_Afs
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetGoName(yname string) string {
    if yname == "afs" { return "Afs" }
    return ""
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &defaultVrf.Afs
    }
    return nil
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &defaultVrf.Afs
    return children
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *TrafficCollector_VrfTable_DefaultVrf) GetParentYangName() string { return "vrf-table" }

// TrafficCollector_VrfTable_DefaultVrf_Afs
// Address Family specific operational data
type TrafficCollector_VrfTable_DefaultVrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af.
    Af []TrafficCollector_VrfTable_DefaultVrf_Afs_Af
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetYangName() string { return "afs" }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *TrafficCollector_VrfTable_DefaultVrf_Afs) GetParentYangName() string { return "default-vrf" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af
// Operational data for given Address Family
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &af.Counters
    }
    return nil
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &af.Counters
    return children
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    return leafs
}

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetYangName() string { return "af" }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *TrafficCollector_VrfTable_DefaultVrf_Afs_Af) GetParentYangName() string { return "afs" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters
// Show Counters
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "tunnels" { return "Tunnels" }
    return ""
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &counters.Prefixes
    }
    if childYangName == "tunnels" {
        return &counters.Tunnels
    }
    return nil
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &counters.Prefixes
    children["tunnels"] = &counters.Tunnels
    return children
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetYangName() string { return "counters" }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetParent() types.Entity { return counters.parent }

func (counters *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters) GetParentYangName() string { return "af" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes) GetParentYangName() string { return "counters" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "ipaddr" { return "Ipaddr" }
    if yname == "mask" { return "Mask" }
    if yname == "label" { return "Label" }
    if yname == "prefix" { return "Prefix" }
    if yname == "label-xr" { return "LabelXr" }
    if yname == "is-active" { return "IsActive" }
    if yname == "base-counter-statistics" { return "BaseCounterStatistics" }
    if yname == "traffic-matrix-counter-statistics" { return "TrafficMatrixCounterStatistics" }
    return ""
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-counter-statistics" {
        return &prefix.BaseCounterStatistics
    }
    if childYangName == "traffic-matrix-counter-statistics" {
        return &prefix.TrafficMatrixCounterStatistics
    }
    return nil
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-counter-statistics"] = &prefix.BaseCounterStatistics
    children["traffic-matrix-counter-statistics"] = &prefix.TrafficMatrixCounterStatistics
    return children
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipaddr"] = prefix.Ipaddr
    leafs["mask"] = prefix.Mask
    leafs["label"] = prefix.Label
    leafs["prefix"] = prefix.Prefix
    leafs["label-xr"] = prefix.LabelXr
    leafs["is-active"] = prefix.IsActive
    return leafs
}

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetFilter() yfilter.YFilter { return baseCounterStatistics.YFilter }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) SetFilter(yf yfilter.YFilter) { baseCounterStatistics.YFilter = yf }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetSegmentPath() string {
    return "base-counter-statistics"
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range baseCounterStatistics.CountHistory {
            if baseCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory{}
        baseCounterStatistics.CountHistory = append(baseCounterStatistics.CountHistory, child)
        return &baseCounterStatistics.CountHistory[len(baseCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseCounterStatistics.CountHistory {
        children[baseCounterStatistics.CountHistory[i].GetSegmentPath()] = &baseCounterStatistics.CountHistory[i]
    }
    return children
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = baseCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = baseCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetYangName() string { return "base-counter-statistics" }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) SetParent(parent types.Entity) { baseCounterStatistics.parent = parent }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetParent() types.Entity { return baseCounterStatistics.parent }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetParentYangName() string { return "prefix" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetParentYangName() string { return "base-counter-statistics" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetFilter() yfilter.YFilter { return trafficMatrixCounterStatistics.YFilter }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) SetFilter(yf yfilter.YFilter) { trafficMatrixCounterStatistics.YFilter = yf }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetSegmentPath() string {
    return "traffic-matrix-counter-statistics"
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range trafficMatrixCounterStatistics.CountHistory {
            if trafficMatrixCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory{}
        trafficMatrixCounterStatistics.CountHistory = append(trafficMatrixCounterStatistics.CountHistory, child)
        return &trafficMatrixCounterStatistics.CountHistory[len(trafficMatrixCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trafficMatrixCounterStatistics.CountHistory {
        children[trafficMatrixCounterStatistics.CountHistory[i].GetSegmentPath()] = &trafficMatrixCounterStatistics.CountHistory[i]
    }
    return children
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetYangName() string { return "traffic-matrix-counter-statistics" }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) SetParent(parent types.Entity) { trafficMatrixCounterStatistics.parent = parent }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetParent() types.Entity { return trafficMatrixCounterStatistics.parent }

func (trafficMatrixCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetParentYangName() string { return "prefix" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetParentYangName() string { return "traffic-matrix-counter-statistics" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels) GetParentYangName() string { return "counters" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "vrfid" { return "Vrfid" }
    if yname == "is-active" { return "IsActive" }
    if yname == "base-counter-statistics" { return "BaseCounterStatistics" }
    return ""
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[interface-name='" + fmt.Sprintf("%v", tunnel.InterfaceName) + "']"
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-counter-statistics" {
        return &tunnel.BaseCounterStatistics
    }
    return nil
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-counter-statistics"] = &tunnel.BaseCounterStatistics
    return children
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = tunnel.InterfaceName
    leafs["interface-name-xr"] = tunnel.InterfaceNameXr
    leafs["interface-handle"] = tunnel.InterfaceHandle
    leafs["vrfid"] = tunnel.Vrfid
    leafs["is-active"] = tunnel.IsActive
    return leafs
}

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetFilter() yfilter.YFilter { return baseCounterStatistics.YFilter }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) SetFilter(yf yfilter.YFilter) { baseCounterStatistics.YFilter = yf }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetSegmentPath() string {
    return "base-counter-statistics"
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range baseCounterStatistics.CountHistory {
            if baseCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory{}
        baseCounterStatistics.CountHistory = append(baseCounterStatistics.CountHistory, child)
        return &baseCounterStatistics.CountHistory[len(baseCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseCounterStatistics.CountHistory {
        children[baseCounterStatistics.CountHistory[i].GetSegmentPath()] = &baseCounterStatistics.CountHistory[i]
    }
    return children
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = baseCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = baseCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetYangName() string { return "base-counter-statistics" }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) SetParent(parent types.Entity) { baseCounterStatistics.parent = parent }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetParent() types.Entity { return baseCounterStatistics.parent }

func (baseCounterStatistics *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetParentYangName() string { return "tunnel" }

// TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_VrfTable_DefaultVrf_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetParentYangName() string { return "base-counter-statistics" }

// TrafficCollector_Afs
// Address Family specific operational data
type TrafficCollector_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for given Address Family. The type is slice of
    // TrafficCollector_Afs_Af.
    Af []TrafficCollector_Afs_Af
}

func (afs *TrafficCollector_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *TrafficCollector_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *TrafficCollector_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *TrafficCollector_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *TrafficCollector_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *TrafficCollector_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *TrafficCollector_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *TrafficCollector_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *TrafficCollector_Afs) GetYangName() string { return "afs" }

func (afs *TrafficCollector_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *TrafficCollector_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *TrafficCollector_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *TrafficCollector_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *TrafficCollector_Afs) GetParent() types.Entity { return afs.parent }

func (afs *TrafficCollector_Afs) GetParentYangName() string { return "traffic-collector" }

// TrafficCollector_Afs_Af
// Operational data for given Address Family
type TrafficCollector_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is TcOperAfName.
    AfName interface{}

    // Show Counters.
    Counters TrafficCollector_Afs_Af_Counters
}

func (af *TrafficCollector_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *TrafficCollector_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *TrafficCollector_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (af *TrafficCollector_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *TrafficCollector_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &af.Counters
    }
    return nil
}

func (af *TrafficCollector_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &af.Counters
    return children
}

func (af *TrafficCollector_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    return leafs
}

func (af *TrafficCollector_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *TrafficCollector_Afs_Af) GetYangName() string { return "af" }

func (af *TrafficCollector_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *TrafficCollector_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *TrafficCollector_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *TrafficCollector_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *TrafficCollector_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *TrafficCollector_Afs_Af) GetParentYangName() string { return "afs" }

// TrafficCollector_Afs_Af_Counters
// Show Counters
type TrafficCollector_Afs_Af_Counters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix Database.
    Prefixes TrafficCollector_Afs_Af_Counters_Prefixes

    // Tunnels.
    Tunnels TrafficCollector_Afs_Af_Counters_Tunnels
}

func (counters *TrafficCollector_Afs_Af_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *TrafficCollector_Afs_Af_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *TrafficCollector_Afs_Af_Counters) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "tunnels" { return "Tunnels" }
    return ""
}

func (counters *TrafficCollector_Afs_Af_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *TrafficCollector_Afs_Af_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &counters.Prefixes
    }
    if childYangName == "tunnels" {
        return &counters.Tunnels
    }
    return nil
}

func (counters *TrafficCollector_Afs_Af_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &counters.Prefixes
    children["tunnels"] = &counters.Tunnels
    return children
}

func (counters *TrafficCollector_Afs_Af_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (counters *TrafficCollector_Afs_Af_Counters) GetBundleName() string { return "cisco_ios_xr" }

func (counters *TrafficCollector_Afs_Af_Counters) GetYangName() string { return "counters" }

func (counters *TrafficCollector_Afs_Af_Counters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (counters *TrafficCollector_Afs_Af_Counters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (counters *TrafficCollector_Afs_Af_Counters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (counters *TrafficCollector_Afs_Af_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *TrafficCollector_Afs_Af_Counters) GetParent() types.Entity { return counters.parent }

func (counters *TrafficCollector_Afs_Af_Counters) GetParentYangName() string { return "af" }

// TrafficCollector_Afs_Af_Counters_Prefixes
// Prefix Database
type TrafficCollector_Afs_Af_Counters_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show Prefix Counter. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix.
    Prefix []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af_Counters_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *TrafficCollector_Afs_Af_Counters_Prefixes) GetParentYangName() string { return "counters" }

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix
// Show Prefix Counter
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ipaddr interface{}

    // Prefix Mask. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Mask interface{}

    // Local Label. The type is interface{} with range: 16..1048575.
    Label interface{}

    // Prefix Address (V4 or V6 Format). The type is string.
    Prefix interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    LabelXr interface{}

    // Prefix is Active and collecting new Statistics. The type is bool.
    IsActive interface{}

    // Base counter statistics.
    BaseCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics

    // Traffic Matrix (TM) counter statistics.
    TrafficMatrixCounterStatistics TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "ipaddr" { return "Ipaddr" }
    if yname == "mask" { return "Mask" }
    if yname == "label" { return "Label" }
    if yname == "prefix" { return "Prefix" }
    if yname == "label-xr" { return "LabelXr" }
    if yname == "is-active" { return "IsActive" }
    if yname == "base-counter-statistics" { return "BaseCounterStatistics" }
    if yname == "traffic-matrix-counter-statistics" { return "TrafficMatrixCounterStatistics" }
    return ""
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-counter-statistics" {
        return &prefix.BaseCounterStatistics
    }
    if childYangName == "traffic-matrix-counter-statistics" {
        return &prefix.TrafficMatrixCounterStatistics
    }
    return nil
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-counter-statistics"] = &prefix.BaseCounterStatistics
    children["traffic-matrix-counter-statistics"] = &prefix.TrafficMatrixCounterStatistics
    return children
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipaddr"] = prefix.Ipaddr
    leafs["mask"] = prefix.Mask
    leafs["label"] = prefix.Label
    leafs["prefix"] = prefix.Prefix
    leafs["label-xr"] = prefix.LabelXr
    leafs["is-active"] = prefix.IsActive
    return leafs
}

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetFilter() yfilter.YFilter { return baseCounterStatistics.YFilter }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) SetFilter(yf yfilter.YFilter) { baseCounterStatistics.YFilter = yf }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetSegmentPath() string {
    return "base-counter-statistics"
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range baseCounterStatistics.CountHistory {
            if baseCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory{}
        baseCounterStatistics.CountHistory = append(baseCounterStatistics.CountHistory, child)
        return &baseCounterStatistics.CountHistory[len(baseCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseCounterStatistics.CountHistory {
        children[baseCounterStatistics.CountHistory[i].GetSegmentPath()] = &baseCounterStatistics.CountHistory[i]
    }
    return children
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = baseCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = baseCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetYangName() string { return "base-counter-statistics" }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) SetParent(parent types.Entity) { baseCounterStatistics.parent = parent }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetParent() types.Entity { return baseCounterStatistics.parent }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics) GetParentYangName() string { return "prefix" }

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_BaseCounterStatistics_CountHistory) GetParentYangName() string { return "base-counter-statistics" }

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics
// Traffic Matrix (TM) counter statistics
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetFilter() yfilter.YFilter { return trafficMatrixCounterStatistics.YFilter }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) SetFilter(yf yfilter.YFilter) { trafficMatrixCounterStatistics.YFilter = yf }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetSegmentPath() string {
    return "traffic-matrix-counter-statistics"
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range trafficMatrixCounterStatistics.CountHistory {
            if trafficMatrixCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory{}
        trafficMatrixCounterStatistics.CountHistory = append(trafficMatrixCounterStatistics.CountHistory, child)
        return &trafficMatrixCounterStatistics.CountHistory[len(trafficMatrixCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trafficMatrixCounterStatistics.CountHistory {
        children[trafficMatrixCounterStatistics.CountHistory[i].GetSegmentPath()] = &trafficMatrixCounterStatistics.CountHistory[i]
    }
    return children
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = trafficMatrixCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = trafficMatrixCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetYangName() string { return "traffic-matrix-counter-statistics" }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) SetParent(parent types.Entity) { trafficMatrixCounterStatistics.parent = parent }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetParent() types.Entity { return trafficMatrixCounterStatistics.parent }

func (trafficMatrixCounterStatistics *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics) GetParentYangName() string { return "prefix" }

// TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Prefixes_Prefix_TrafficMatrixCounterStatistics_CountHistory) GetParentYangName() string { return "traffic-matrix-counter-statistics" }

// TrafficCollector_Afs_Af_Counters_Tunnels
// Tunnels
type TrafficCollector_Afs_Af_Counters_Tunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel.
    Tunnel []TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetFilter() yfilter.YFilter { return tunnels.YFilter }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) SetFilter(yf yfilter.YFilter) { tunnels.YFilter = yf }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetSegmentPath() string {
    return "tunnels"
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        for _, c := range tunnels.Tunnel {
            if tunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel{}
        tunnels.Tunnel = append(tunnels.Tunnel, child)
        return &tunnels.Tunnel[len(tunnels.Tunnel)-1]
    }
    return nil
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnels.Tunnel {
        children[tunnels.Tunnel[i].GetSegmentPath()] = &tunnels.Tunnel[i]
    }
    return children
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetBundleName() string { return "cisco_ios_xr" }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetYangName() string { return "tunnels" }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) SetParent(parent types.Entity) { tunnels.parent = parent }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetParent() types.Entity { return tunnels.parent }

func (tunnels *TrafficCollector_Afs_Af_Counters_Tunnels) GetParentYangName() string { return "counters" }

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel
// Tunnel information
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The Interface Name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-name-xr" { return "InterfaceNameXr" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "vrfid" { return "Vrfid" }
    if yname == "is-active" { return "IsActive" }
    if yname == "base-counter-statistics" { return "BaseCounterStatistics" }
    return ""
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[interface-name='" + fmt.Sprintf("%v", tunnel.InterfaceName) + "']"
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-counter-statistics" {
        return &tunnel.BaseCounterStatistics
    }
    return nil
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-counter-statistics"] = &tunnel.BaseCounterStatistics
    return children
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = tunnel.InterfaceName
    leafs["interface-name-xr"] = tunnel.InterfaceNameXr
    leafs["interface-handle"] = tunnel.InterfaceHandle
    leafs["vrfid"] = tunnel.Vrfid
    leafs["is-active"] = tunnel.IsActive
    return leafs
}

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel) GetParentYangName() string { return "tunnels" }

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics
// Base counter statistics
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Average Rate of Packets/second switched. The type is interface{} with
    // range: 0..18446744073709551615. Units are packet/s.
    TransmitPacketsPerSecondSwitched interface{}

    // Average Rate of Bytes/second switched. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    TransmitBytesPerSecondSwitched interface{}

    // Counter History. The type is slice of
    // TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory.
    CountHistory []TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetFilter() yfilter.YFilter { return baseCounterStatistics.YFilter }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) SetFilter(yf yfilter.YFilter) { baseCounterStatistics.YFilter = yf }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetGoName(yname string) string {
    if yname == "transmit-packets-per-second-switched" { return "TransmitPacketsPerSecondSwitched" }
    if yname == "transmit-bytes-per-second-switched" { return "TransmitBytesPerSecondSwitched" }
    if yname == "count-history" { return "CountHistory" }
    return ""
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetSegmentPath() string {
    return "base-counter-statistics"
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "count-history" {
        for _, c := range baseCounterStatistics.CountHistory {
            if baseCounterStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory{}
        baseCounterStatistics.CountHistory = append(baseCounterStatistics.CountHistory, child)
        return &baseCounterStatistics.CountHistory[len(baseCounterStatistics.CountHistory)-1]
    }
    return nil
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range baseCounterStatistics.CountHistory {
        children[baseCounterStatistics.CountHistory[i].GetSegmentPath()] = &baseCounterStatistics.CountHistory[i]
    }
    return children
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transmit-packets-per-second-switched"] = baseCounterStatistics.TransmitPacketsPerSecondSwitched
    leafs["transmit-bytes-per-second-switched"] = baseCounterStatistics.TransmitBytesPerSecondSwitched
    return leafs
}

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetYangName() string { return "base-counter-statistics" }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) SetParent(parent types.Entity) { baseCounterStatistics.parent = parent }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetParent() types.Entity { return baseCounterStatistics.parent }

func (baseCounterStatistics *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics) GetParentYangName() string { return "tunnel" }

// TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory
// Counter History
type TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetFilter() yfilter.YFilter { return countHistory.YFilter }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) SetFilter(yf yfilter.YFilter) { countHistory.YFilter = yf }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetGoName(yname string) string {
    if yname == "event-start-timestamp" { return "EventStartTimestamp" }
    if yname == "event-end-timestamp" { return "EventEndTimestamp" }
    if yname == "transmit-number-of-packets-switched" { return "TransmitNumberOfPacketsSwitched" }
    if yname == "transmit-number-of-bytes-switched" { return "TransmitNumberOfBytesSwitched" }
    if yname == "is-valid" { return "IsValid" }
    return ""
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetSegmentPath() string {
    return "count-history"
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-start-timestamp"] = countHistory.EventStartTimestamp
    leafs["event-end-timestamp"] = countHistory.EventEndTimestamp
    leafs["transmit-number-of-packets-switched"] = countHistory.TransmitNumberOfPacketsSwitched
    leafs["transmit-number-of-bytes-switched"] = countHistory.TransmitNumberOfBytesSwitched
    leafs["is-valid"] = countHistory.IsValid
    return leafs
}

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetBundleName() string { return "cisco_ios_xr" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetYangName() string { return "count-history" }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) SetParent(parent types.Entity) { countHistory.parent = parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetParent() types.Entity { return countHistory.parent }

func (countHistory *TrafficCollector_Afs_Af_Counters_Tunnels_Tunnel_BaseCounterStatistics_CountHistory) GetParentYangName() string { return "base-counter-statistics" }

