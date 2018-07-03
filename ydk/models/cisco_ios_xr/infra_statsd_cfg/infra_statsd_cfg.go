// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-statsd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   statistics: Global statistics configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_statsd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_statsd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-statsd-cfg statistics}", reflect.TypeOf(Statistics{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-statsd-cfg:statistics", reflect.TypeOf(Statistics{}))
}

// Statistics
// Global statistics configuration
type Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection period for statistics polling.
    Period Statistics_Period
}

func (statistics *Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "Cisco-IOS-XR-infra-statsd-cfg"
    statistics.EntityData.SegmentPath = "Cisco-IOS-XR-infra-statsd-cfg:statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("period", types.YChild{"Period", &statistics.Period})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Statistics_Period
// Collection period for statistics polling
type Statistics_Period struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection polling period for service accounting collectors.
    ServiceAccounting Statistics_Period_ServiceAccounting
}

func (period *Statistics_Period) GetEntityData() *types.CommonEntityData {
    period.EntityData.YFilter = period.YFilter
    period.EntityData.YangName = "period"
    period.EntityData.BundleName = "cisco_ios_xr"
    period.EntityData.ParentYangName = "statistics"
    period.EntityData.SegmentPath = "period"
    period.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    period.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    period.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    period.EntityData.Children = types.NewOrderedMap()
    period.EntityData.Children.Append("service-accounting", types.YChild{"ServiceAccounting", &period.ServiceAccounting})
    period.EntityData.Leafs = types.NewOrderedMap()

    period.EntityData.YListKeys = []string {}

    return &(period.EntityData)
}

// Statistics_Period_ServiceAccounting
// Collection polling period for service
// accounting collectors
type Statistics_Period_ServiceAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection polling period for service accounting collectors. The type is
    // interface{} with range: 30..3600. The default value is 900.
    PollingPeriod interface{}

    // Disable periodic statistics polling for service accounting collectors. The
    // type is interface{}.
    PollingDisable interface{}
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetEntityData() *types.CommonEntityData {
    serviceAccounting.EntityData.YFilter = serviceAccounting.YFilter
    serviceAccounting.EntityData.YangName = "service-accounting"
    serviceAccounting.EntityData.BundleName = "cisco_ios_xr"
    serviceAccounting.EntityData.ParentYangName = "period"
    serviceAccounting.EntityData.SegmentPath = "service-accounting"
    serviceAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccounting.EntityData.Children = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs.Append("polling-period", types.YLeaf{"PollingPeriod", serviceAccounting.PollingPeriod})
    serviceAccounting.EntityData.Leafs.Append("polling-disable", types.YLeaf{"PollingDisable", serviceAccounting.PollingDisable})

    serviceAccounting.EntityData.YListKeys = []string {}

    return &(serviceAccounting.EntityData)
}

