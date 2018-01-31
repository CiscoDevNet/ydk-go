// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-tc package configuration.
// 
// This module contains definitions
// for the following management objects:
//   traffic-collector: Global Traffic Collector configuration
//     commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_tc_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_tc_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-tc-cfg traffic-collector}", reflect.TypeOf(TrafficCollector{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-tc-cfg:traffic-collector", reflect.TypeOf(TrafficCollector{}))
}

// HistoryTimeout represents History timeout
type HistoryTimeout string

const (
    // Max timeout
    HistoryTimeout_max HistoryTimeout = "max"
)

// HistorySize represents History size
type HistorySize string

const (
    // Max history
    HistorySize_max HistorySize = "max"
)

// CollectIonInterval represents Collect ion interval
type CollectIonInterval string

const (
    // Interval1minute
    CollectIonInterval_Y_1_minute CollectIonInterval = "1-minute"

    // Interval2minutes
    CollectIonInterval_Y_2_minutes CollectIonInterval = "2-minutes"

    // Interval3minutes
    CollectIonInterval_Y_3_minutes CollectIonInterval = "3-minutes"

    // Interval4minutes
    CollectIonInterval_Y_4_minutes CollectIonInterval = "4-minutes"

    // Interval5minutes
    CollectIonInterval_Y_5_minutes CollectIonInterval = "5-minutes"

    // Interval6minutes
    CollectIonInterval_Y_6_minutes CollectIonInterval = "6-minutes"

    // Interval10minutes
    CollectIonInterval_Y_10_minutes CollectIonInterval = "10-minutes"

    // Interval12minutes
    CollectIonInterval_Y_12_minutes CollectIonInterval = "12-minutes"

    // Interval15inutes
    CollectIonInterval_Y_15_minutes CollectIonInterval = "15-minutes"

    // Interval20minutes
    CollectIonInterval_Y_20_minutes CollectIonInterval = "20-minutes"

    // Interval30minutes
    CollectIonInterval_Y_30_minutes CollectIonInterval = "30-minutes"

    // Interval60minutes
    CollectIonInterval_Y_60_minutes CollectIonInterval = "60-minutes"
)

// TrafficCollector
// Global Traffic Collector configuration commands
type TrafficCollector struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable traffic collector. The type is interface{}.
    EnableTrafficCollector interface{}

    // Configure external interfaces.
    ExternalInterfaces TrafficCollector_ExternalInterfaces

    // Configure statistics related parameters.
    Statistics TrafficCollector_Statistics
}

func (trafficCollector *TrafficCollector) GetFilter() yfilter.YFilter { return trafficCollector.YFilter }

func (trafficCollector *TrafficCollector) SetFilter(yf yfilter.YFilter) { trafficCollector.YFilter = yf }

func (trafficCollector *TrafficCollector) GetGoName(yname string) string {
    if yname == "enable-traffic-collector" { return "EnableTrafficCollector" }
    if yname == "external-interfaces" { return "ExternalInterfaces" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (trafficCollector *TrafficCollector) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-tc-cfg:traffic-collector"
}

func (trafficCollector *TrafficCollector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "external-interfaces" {
        return &trafficCollector.ExternalInterfaces
    }
    if childYangName == "statistics" {
        return &trafficCollector.Statistics
    }
    return nil
}

func (trafficCollector *TrafficCollector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["external-interfaces"] = &trafficCollector.ExternalInterfaces
    children["statistics"] = &trafficCollector.Statistics
    return children
}

func (trafficCollector *TrafficCollector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-traffic-collector"] = trafficCollector.EnableTrafficCollector
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

func (trafficCollector *TrafficCollector) GetParentYangName() string { return "Cisco-IOS-XR-infra-tc-cfg" }

// TrafficCollector_ExternalInterfaces
// Configure external interfaces
type TrafficCollector_ExternalInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure an external internface. The type is slice of
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
// Configure an external internface
type TrafficCollector_ExternalInterfaces_ExternalInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable traffic collector on this interface. The type is interface{}.
    Enable interface{}
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetFilter() yfilter.YFilter { return externalInterface.YFilter }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) SetFilter(yf yfilter.YFilter) { externalInterface.YFilter = yf }

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
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
    leafs["enable"] = externalInterface.Enable
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

// TrafficCollector_Statistics
// Configure statistics related parameters
type TrafficCollector_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure statistics history size. The type is one of the following types:
    // enumeration HistorySize, or int with range: 1..10.
    HistorySize interface{}

    // Configure statistics collection interval. The type is CollectIonInterval.
    CollectionInterval interface{}

    // Enable traffic collector statistics. The type is interface{}.
    EnableTrafficCollectorStatistics interface{}

    // Configure statistics history timeout interval. The type is one of the
    // following types: enumeration HistoryTimeout, or int with range: 0..720.
    HistoryTimeout interface{}
}

func (statistics *TrafficCollector_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *TrafficCollector_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *TrafficCollector_Statistics) GetGoName(yname string) string {
    if yname == "history-size" { return "HistorySize" }
    if yname == "collection-interval" { return "CollectionInterval" }
    if yname == "enable-traffic-collector-statistics" { return "EnableTrafficCollectorStatistics" }
    if yname == "history-timeout" { return "HistoryTimeout" }
    return ""
}

func (statistics *TrafficCollector_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *TrafficCollector_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *TrafficCollector_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *TrafficCollector_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["history-size"] = statistics.HistorySize
    leafs["collection-interval"] = statistics.CollectionInterval
    leafs["enable-traffic-collector-statistics"] = statistics.EnableTrafficCollectorStatistics
    leafs["history-timeout"] = statistics.HistoryTimeout
    return leafs
}

func (statistics *TrafficCollector_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *TrafficCollector_Statistics) GetYangName() string { return "statistics" }

func (statistics *TrafficCollector_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *TrafficCollector_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *TrafficCollector_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *TrafficCollector_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *TrafficCollector_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *TrafficCollector_Statistics) GetParentYangName() string { return "traffic-collector" }

