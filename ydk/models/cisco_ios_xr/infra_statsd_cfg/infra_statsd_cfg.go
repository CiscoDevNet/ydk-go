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
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection period for statistics polling.
    Period Statistics_Period
}

func (statistics *Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Statistics) GetGoName(yname string) string {
    if yname == "period" { return "Period" }
    return ""
}

func (statistics *Statistics) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-statsd-cfg:statistics"
}

func (statistics *Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "period" {
        return &statistics.Period
    }
    return nil
}

func (statistics *Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["period"] = &statistics.Period
    return children
}

func (statistics *Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Statistics) GetYangName() string { return "statistics" }

func (statistics *Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Statistics) GetParentYangName() string { return "Cisco-IOS-XR-infra-statsd-cfg" }

// Statistics_Period
// Collection period for statistics polling
type Statistics_Period struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection polling period for service accounting collectors.
    ServiceAccounting Statistics_Period_ServiceAccounting
}

func (period *Statistics_Period) GetFilter() yfilter.YFilter { return period.YFilter }

func (period *Statistics_Period) SetFilter(yf yfilter.YFilter) { period.YFilter = yf }

func (period *Statistics_Period) GetGoName(yname string) string {
    if yname == "service-accounting" { return "ServiceAccounting" }
    return ""
}

func (period *Statistics_Period) GetSegmentPath() string {
    return "period"
}

func (period *Statistics_Period) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-accounting" {
        return &period.ServiceAccounting
    }
    return nil
}

func (period *Statistics_Period) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-accounting"] = &period.ServiceAccounting
    return children
}

func (period *Statistics_Period) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (period *Statistics_Period) GetBundleName() string { return "cisco_ios_xr" }

func (period *Statistics_Period) GetYangName() string { return "period" }

func (period *Statistics_Period) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (period *Statistics_Period) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (period *Statistics_Period) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (period *Statistics_Period) SetParent(parent types.Entity) { period.parent = parent }

func (period *Statistics_Period) GetParent() types.Entity { return period.parent }

func (period *Statistics_Period) GetParentYangName() string { return "statistics" }

// Statistics_Period_ServiceAccounting
// Collection polling period for service
// accounting collectors
type Statistics_Period_ServiceAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection polling period for service accounting collectors. The type is
    // interface{} with range: 30..3600. The default value is 900.
    PollingPeriod interface{}

    // Disable periodic statistics polling for service accounting collectors. The
    // type is interface{}.
    PollingDisable interface{}
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetFilter() yfilter.YFilter { return serviceAccounting.YFilter }

func (serviceAccounting *Statistics_Period_ServiceAccounting) SetFilter(yf yfilter.YFilter) { serviceAccounting.YFilter = yf }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetGoName(yname string) string {
    if yname == "polling-period" { return "PollingPeriod" }
    if yname == "polling-disable" { return "PollingDisable" }
    return ""
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetSegmentPath() string {
    return "service-accounting"
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["polling-period"] = serviceAccounting.PollingPeriod
    leafs["polling-disable"] = serviceAccounting.PollingDisable
    return leafs
}

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetYangName() string { return "service-accounting" }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccounting *Statistics_Period_ServiceAccounting) SetParent(parent types.Entity) { serviceAccounting.parent = parent }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetParent() types.Entity { return serviceAccounting.parent }

func (serviceAccounting *Statistics_Period_ServiceAccounting) GetParentYangName() string { return "period" }

