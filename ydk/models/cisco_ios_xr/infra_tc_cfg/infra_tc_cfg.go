// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-tc package configuration.
// 
// This module contains definitions
// for the following management objects:
//   traffic-collector: Global Traffic Collector configuration
//     commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// HistoryTimeout represents History timeout
type HistoryTimeout string

const (
    // Max timeout
    HistoryTimeout_max HistoryTimeout = "max"
)

// TrafficCollector
// Global Traffic Collector configuration commands
type TrafficCollector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable traffic collector. The type is interface{}.
    EnableTrafficCollector interface{}

    // Configure external interfaces.
    ExternalInterfaces TrafficCollector_ExternalInterfaces

    // Configure statistics related parameters.
    Statistics TrafficCollector_Statistics
}

func (trafficCollector *TrafficCollector) GetEntityData() *types.CommonEntityData {
    trafficCollector.EntityData.YFilter = trafficCollector.YFilter
    trafficCollector.EntityData.YangName = "traffic-collector"
    trafficCollector.EntityData.BundleName = "cisco_ios_xr"
    trafficCollector.EntityData.ParentYangName = "Cisco-IOS-XR-infra-tc-cfg"
    trafficCollector.EntityData.SegmentPath = "Cisco-IOS-XR-infra-tc-cfg:traffic-collector"
    trafficCollector.EntityData.AbsolutePath = trafficCollector.EntityData.SegmentPath
    trafficCollector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficCollector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficCollector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficCollector.EntityData.Children = types.NewOrderedMap()
    trafficCollector.EntityData.Children.Append("external-interfaces", types.YChild{"ExternalInterfaces", &trafficCollector.ExternalInterfaces})
    trafficCollector.EntityData.Children.Append("statistics", types.YChild{"Statistics", &trafficCollector.Statistics})
    trafficCollector.EntityData.Leafs = types.NewOrderedMap()
    trafficCollector.EntityData.Leafs.Append("enable-traffic-collector", types.YLeaf{"EnableTrafficCollector", trafficCollector.EnableTrafficCollector})

    trafficCollector.EntityData.YListKeys = []string {}

    return &(trafficCollector.EntityData)
}

// TrafficCollector_ExternalInterfaces
// Configure external interfaces
type TrafficCollector_ExternalInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure an external internface. The type is slice of
    // TrafficCollector_ExternalInterfaces_ExternalInterface.
    ExternalInterface []*TrafficCollector_ExternalInterfaces_ExternalInterface
}

func (externalInterfaces *TrafficCollector_ExternalInterfaces) GetEntityData() *types.CommonEntityData {
    externalInterfaces.EntityData.YFilter = externalInterfaces.YFilter
    externalInterfaces.EntityData.YangName = "external-interfaces"
    externalInterfaces.EntityData.BundleName = "cisco_ios_xr"
    externalInterfaces.EntityData.ParentYangName = "traffic-collector"
    externalInterfaces.EntityData.SegmentPath = "external-interfaces"
    externalInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-cfg:traffic-collector/" + externalInterfaces.EntityData.SegmentPath
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
// Configure an external internface
type TrafficCollector_ExternalInterfaces_ExternalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of interface. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable traffic collector on this interface. The type is interface{}.
    Enable interface{}
}

func (externalInterface *TrafficCollector_ExternalInterfaces_ExternalInterface) GetEntityData() *types.CommonEntityData {
    externalInterface.EntityData.YFilter = externalInterface.YFilter
    externalInterface.EntityData.YangName = "external-interface"
    externalInterface.EntityData.BundleName = "cisco_ios_xr"
    externalInterface.EntityData.ParentYangName = "external-interfaces"
    externalInterface.EntityData.SegmentPath = "external-interface" + types.AddKeyToken(externalInterface.InterfaceName, "interface-name")
    externalInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-cfg:traffic-collector/external-interfaces/" + externalInterface.EntityData.SegmentPath
    externalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalInterface.EntityData.Children = types.NewOrderedMap()
    externalInterface.EntityData.Leafs = types.NewOrderedMap()
    externalInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", externalInterface.InterfaceName})
    externalInterface.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", externalInterface.Enable})

    externalInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(externalInterface.EntityData)
}

// TrafficCollector_Statistics
// Configure statistics related parameters
type TrafficCollector_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *TrafficCollector_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "traffic-collector"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-tc-cfg:traffic-collector/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("history-size", types.YLeaf{"HistorySize", statistics.HistorySize})
    statistics.EntityData.Leafs.Append("collection-interval", types.YLeaf{"CollectionInterval", statistics.CollectionInterval})
    statistics.EntityData.Leafs.Append("enable-traffic-collector-statistics", types.YLeaf{"EnableTrafficCollectorStatistics", statistics.EnableTrafficCollectorStatistics})
    statistics.EntityData.Leafs.Append("history-timeout", types.YLeaf{"HistoryTimeout", statistics.HistoryTimeout})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

