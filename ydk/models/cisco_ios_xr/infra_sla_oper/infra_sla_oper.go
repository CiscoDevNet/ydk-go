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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of all SLA protocols.
    Protocols Sla_Protocols
}

func (sla *Sla) GetFilter() yfilter.YFilter { return sla.YFilter }

func (sla *Sla) SetFilter(yf yfilter.YFilter) { sla.YFilter = yf }

func (sla *Sla) GetGoName(yname string) string {
    if yname == "protocols" { return "Protocols" }
    return ""
}

func (sla *Sla) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-sla-oper:sla"
}

func (sla *Sla) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocols" {
        return &sla.Protocols
    }
    return nil
}

func (sla *Sla) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocols"] = &sla.Protocols
    return children
}

func (sla *Sla) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sla *Sla) GetBundleName() string { return "cisco_ios_xr" }

func (sla *Sla) GetYangName() string { return "sla" }

func (sla *Sla) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sla *Sla) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sla *Sla) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sla *Sla) SetParent(parent types.Entity) { sla.parent = parent }

func (sla *Sla) GetParent() types.Entity { return sla.parent }

func (sla *Sla) GetParentYangName() string { return "Cisco-IOS-XR-infra-sla-oper" }

// Sla_Protocols
// Table of all SLA protocols
type Sla_Protocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Ethernet SLA protocol.
    Ethernet Sla_Protocols_Ethernet
}

func (protocols *Sla_Protocols) GetFilter() yfilter.YFilter { return protocols.YFilter }

func (protocols *Sla_Protocols) SetFilter(yf yfilter.YFilter) { protocols.YFilter = yf }

func (protocols *Sla_Protocols) GetGoName(yname string) string {
    if yname == "Cisco-IOS-XR-ethernet-cfm-oper:ethernet" { return "Ethernet" }
    return ""
}

func (protocols *Sla_Protocols) GetSegmentPath() string {
    return "protocols"
}

func (protocols *Sla_Protocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-ethernet-cfm-oper:ethernet" {
        return &protocols.Ethernet
    }
    return nil
}

func (protocols *Sla_Protocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-ethernet-cfm-oper:ethernet"] = &protocols.Ethernet
    return children
}

func (protocols *Sla_Protocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocols *Sla_Protocols) GetBundleName() string { return "cisco_ios_xr" }

func (protocols *Sla_Protocols) GetYangName() string { return "protocols" }

func (protocols *Sla_Protocols) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocols *Sla_Protocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocols *Sla_Protocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocols *Sla_Protocols) SetParent(parent types.Entity) { protocols.parent = parent }

func (protocols *Sla_Protocols) GetParent() types.Entity { return protocols.parent }

func (protocols *Sla_Protocols) GetParentYangName() string { return "sla" }

// Sla_Protocols_Ethernet
// The Ethernet SLA protocol
type Sla_Protocols_Ethernet struct {
    parent types.Entity
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

func (ethernet *Sla_Protocols_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *Sla_Protocols_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *Sla_Protocols_Ethernet) GetGoName(yname string) string {
    if yname == "statistics-on-demand-currents" { return "StatisticsOnDemandCurrents" }
    if yname == "operations" { return "Operations" }
    if yname == "statistics-historicals" { return "StatisticsHistoricals" }
    if yname == "statistics-on-demand-historicals" { return "StatisticsOnDemandHistoricals" }
    if yname == "config-errors" { return "ConfigErrors" }
    if yname == "on-demand-operations" { return "OnDemandOperations" }
    if yname == "statistics-currents" { return "StatisticsCurrents" }
    return ""
}

func (ethernet *Sla_Protocols_Ethernet) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-cfm-oper:ethernet"
}

func (ethernet *Sla_Protocols_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-on-demand-currents" {
        return &ethernet.StatisticsOnDemandCurrents
    }
    if childYangName == "operations" {
        return &ethernet.Operations
    }
    if childYangName == "statistics-historicals" {
        return &ethernet.StatisticsHistoricals
    }
    if childYangName == "statistics-on-demand-historicals" {
        return &ethernet.StatisticsOnDemandHistoricals
    }
    if childYangName == "config-errors" {
        return &ethernet.ConfigErrors
    }
    if childYangName == "on-demand-operations" {
        return &ethernet.OnDemandOperations
    }
    if childYangName == "statistics-currents" {
        return &ethernet.StatisticsCurrents
    }
    return nil
}

func (ethernet *Sla_Protocols_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics-on-demand-currents"] = &ethernet.StatisticsOnDemandCurrents
    children["operations"] = &ethernet.Operations
    children["statistics-historicals"] = &ethernet.StatisticsHistoricals
    children["statistics-on-demand-historicals"] = &ethernet.StatisticsOnDemandHistoricals
    children["config-errors"] = &ethernet.ConfigErrors
    children["on-demand-operations"] = &ethernet.OnDemandOperations
    children["statistics-currents"] = &ethernet.StatisticsCurrents
    return children
}

func (ethernet *Sla_Protocols_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernet *Sla_Protocols_Ethernet) GetBundleName() string { return "cisco_ios_xr" }

func (ethernet *Sla_Protocols_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *Sla_Protocols_Ethernet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernet *Sla_Protocols_Ethernet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernet *Sla_Protocols_Ethernet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernet *Sla_Protocols_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *Sla_Protocols_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *Sla_Protocols_Ethernet) GetParentYangName() string { return "protocols" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents
// Table of current statistics for SLA on-demand
// operations
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current statistics data for an SLA on-demand operation. The type is slice
    // of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent.
    StatisticsOnDemandCurrent []Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetFilter() yfilter.YFilter { return statisticsOnDemandCurrents.YFilter }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) SetFilter(yf yfilter.YFilter) { statisticsOnDemandCurrents.YFilter = yf }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetGoName(yname string) string {
    if yname == "statistics-on-demand-current" { return "StatisticsOnDemandCurrent" }
    return ""
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetSegmentPath() string {
    return "statistics-on-demand-currents"
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-on-demand-current" {
        for _, c := range statisticsOnDemandCurrents.StatisticsOnDemandCurrent {
            if statisticsOnDemandCurrents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent{}
        statisticsOnDemandCurrents.StatisticsOnDemandCurrent = append(statisticsOnDemandCurrents.StatisticsOnDemandCurrent, child)
        return &statisticsOnDemandCurrents.StatisticsOnDemandCurrent[len(statisticsOnDemandCurrents.StatisticsOnDemandCurrent)-1]
    }
    return nil
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticsOnDemandCurrents.StatisticsOnDemandCurrent {
        children[statisticsOnDemandCurrents.StatisticsOnDemandCurrent[i].GetSegmentPath()] = &statisticsOnDemandCurrents.StatisticsOnDemandCurrent[i]
    }
    return children
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetYangName() string { return "statistics-on-demand-currents" }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) SetParent(parent types.Entity) { statisticsOnDemandCurrents.parent = parent }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetParent() types.Entity { return statisticsOnDemandCurrents.parent }

func (statisticsOnDemandCurrents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent
// Current statistics data for an SLA on-demand
// operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetFilter() yfilter.YFilter { return statisticsOnDemandCurrent.YFilter }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) SetFilter(yf yfilter.YFilter) { statisticsOnDemandCurrent.YFilter = yf }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "probe-type" { return "ProbeType" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "specific-options" { return "SpecificOptions" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetSegmentPath() string {
    return "statistics-on-demand-current"
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "specific-options" {
        return &statisticsOnDemandCurrent.SpecificOptions
    }
    if childYangName == "operation-schedule" {
        return &statisticsOnDemandCurrent.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range statisticsOnDemandCurrent.OperationMetric {
            if statisticsOnDemandCurrent.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric{}
        statisticsOnDemandCurrent.OperationMetric = append(statisticsOnDemandCurrent.OperationMetric, child)
        return &statisticsOnDemandCurrent.OperationMetric[len(statisticsOnDemandCurrent.OperationMetric)-1]
    }
    return nil
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["specific-options"] = &statisticsOnDemandCurrent.SpecificOptions
    children["operation-schedule"] = &statisticsOnDemandCurrent.OperationSchedule
    for i := range statisticsOnDemandCurrent.OperationMetric {
        children[statisticsOnDemandCurrent.OperationMetric[i].GetSegmentPath()] = &statisticsOnDemandCurrent.OperationMetric[i]
    }
    return children
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = statisticsOnDemandCurrent.OperationId
    leafs["domain-name"] = statisticsOnDemandCurrent.DomainName
    leafs["interface-name"] = statisticsOnDemandCurrent.InterfaceName
    leafs["mep-id"] = statisticsOnDemandCurrent.MepId
    leafs["mac-address"] = statisticsOnDemandCurrent.MacAddress
    leafs["probe-type"] = statisticsOnDemandCurrent.ProbeType
    leafs["display-short"] = statisticsOnDemandCurrent.DisplayShort
    leafs["display-long"] = statisticsOnDemandCurrent.DisplayLong
    leafs["flr-calculation-interval"] = statisticsOnDemandCurrent.FlrCalculationInterval
    return leafs
}

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetYangName() string { return "statistics-on-demand-current" }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) SetParent(parent types.Entity) { statisticsOnDemandCurrent.parent = parent }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetParent() types.Entity { return statisticsOnDemandCurrent.parent }

func (statisticsOnDemandCurrent *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent) GetParentYangName() string { return "statistics-on-demand-currents" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions) GetParentYangName() string { return "statistics-on-demand-current" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationSchedule) GetParentYangName() string { return "statistics-on-demand-current" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket.
    Bucket []Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "bucket" { return "Bucket" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &operationMetric.Config
    }
    if childYangName == "bucket" {
        for _, c := range operationMetric.Bucket {
            if operationMetric.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket{}
        operationMetric.Bucket = append(operationMetric.Bucket, child)
        return &operationMetric.Bucket[len(operationMetric.Bucket)-1]
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &operationMetric.Config
    for i := range operationMetric.Bucket {
        children[operationMetric.Bucket[i].GetSegmentPath()] = &operationMetric.Bucket[i]
    }
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric) GetParentYangName() string { return "statistics-on-demand-current" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config struct {
    parent types.Entity
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

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetSegmentPath() string {
    return "config"
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = config.MetricType
    leafs["bins-count"] = config.BinsCount
    leafs["bins-width"] = config.BinsWidth
    leafs["bucket-size"] = config.BucketSize
    leafs["bucket-size-unit"] = config.BucketSizeUnit
    leafs["buckets-archive"] = config.BucketsArchive
    return leafs
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetYangName() string { return "config" }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetParent() types.Entity { return config.parent }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Config) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket struct {
    parent types.Entity
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

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetFilter() yfilter.YFilter { return bucket.YFilter }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) SetFilter(yf yfilter.YFilter) { bucket.YFilter = yf }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetGoName(yname string) string {
    if yname == "start-at" { return "StartAt" }
    if yname == "duration" { return "Duration" }
    if yname == "sent" { return "Sent" }
    if yname == "lost" { return "Lost" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "duplicates" { return "Duplicates" }
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    if yname == "time-of-minimum" { return "TimeOfMinimum" }
    if yname == "time-of-maximum" { return "TimeOfMaximum" }
    if yname == "average" { return "Average" }
    if yname == "standard-deviation" { return "StandardDeviation" }
    if yname == "result-count" { return "ResultCount" }
    if yname == "data-sent-count" { return "DataSentCount" }
    if yname == "data-lost-count" { return "DataLostCount" }
    if yname == "overall-flr" { return "OverallFlr" }
    if yname == "suspect-start-mid-bucket" { return "SuspectStartMidBucket" }
    if yname == "suspect-schedule-latency" { return "SuspectScheduleLatency" }
    if yname == "suspect-send-fail" { return "SuspectSendFail" }
    if yname == "suspect-premature-end" { return "SuspectPrematureEnd" }
    if yname == "suspect-clock-drift" { return "SuspectClockDrift" }
    if yname == "suspect-memory-allocation-failed" { return "SuspectMemoryAllocationFailed" }
    if yname == "suspect-cleared-mid-bucket" { return "SuspectClearedMidBucket" }
    if yname == "suspect-probe-restarted" { return "SuspectProbeRestarted" }
    if yname == "suspect-management-latency" { return "SuspectManagementLatency" }
    if yname == "suspect-multiple-buckets" { return "SuspectMultipleBuckets" }
    if yname == "suspect-misordering" { return "SuspectMisordering" }
    if yname == "suspect-flr-low-packet-count" { return "SuspectFlrLowPacketCount" }
    if yname == "premature-reason" { return "PrematureReason" }
    if yname == "premature-reason-string" { return "PrematureReasonString" }
    if yname == "contents" { return "Contents" }
    return ""
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetSegmentPath() string {
    return "bucket"
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contents" {
        return &bucket.Contents
    }
    return nil
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contents"] = &bucket.Contents
    return children
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-at"] = bucket.StartAt
    leafs["duration"] = bucket.Duration
    leafs["sent"] = bucket.Sent
    leafs["lost"] = bucket.Lost
    leafs["corrupt"] = bucket.Corrupt
    leafs["out-of-order"] = bucket.OutOfOrder
    leafs["duplicates"] = bucket.Duplicates
    leafs["minimum"] = bucket.Minimum
    leafs["maximum"] = bucket.Maximum
    leafs["time-of-minimum"] = bucket.TimeOfMinimum
    leafs["time-of-maximum"] = bucket.TimeOfMaximum
    leafs["average"] = bucket.Average
    leafs["standard-deviation"] = bucket.StandardDeviation
    leafs["result-count"] = bucket.ResultCount
    leafs["data-sent-count"] = bucket.DataSentCount
    leafs["data-lost-count"] = bucket.DataLostCount
    leafs["overall-flr"] = bucket.OverallFlr
    leafs["suspect-start-mid-bucket"] = bucket.SuspectStartMidBucket
    leafs["suspect-schedule-latency"] = bucket.SuspectScheduleLatency
    leafs["suspect-send-fail"] = bucket.SuspectSendFail
    leafs["suspect-premature-end"] = bucket.SuspectPrematureEnd
    leafs["suspect-clock-drift"] = bucket.SuspectClockDrift
    leafs["suspect-memory-allocation-failed"] = bucket.SuspectMemoryAllocationFailed
    leafs["suspect-cleared-mid-bucket"] = bucket.SuspectClearedMidBucket
    leafs["suspect-probe-restarted"] = bucket.SuspectProbeRestarted
    leafs["suspect-management-latency"] = bucket.SuspectManagementLatency
    leafs["suspect-multiple-buckets"] = bucket.SuspectMultipleBuckets
    leafs["suspect-misordering"] = bucket.SuspectMisordering
    leafs["suspect-flr-low-packet-count"] = bucket.SuspectFlrLowPacketCount
    leafs["premature-reason"] = bucket.PrematureReason
    leafs["premature-reason-string"] = bucket.PrematureReasonString
    return leafs
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetBundleName() string { return "cisco_ios_xr" }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetYangName() string { return "bucket" }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) SetParent(parent types.Entity) { bucket.parent = parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetParent() types.Entity { return bucket.parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetFilter() yfilter.YFilter { return contents.YFilter }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) SetFilter(yf yfilter.YFilter) { contents.YFilter = yf }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetGoName(yname string) string {
    if yname == "bucket-type" { return "BucketType" }
    if yname == "aggregated" { return "Aggregated" }
    if yname == "unaggregated" { return "Unaggregated" }
    return ""
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetSegmentPath() string {
    return "contents"
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregated" {
        return &contents.Aggregated
    }
    if childYangName == "unaggregated" {
        return &contents.Unaggregated
    }
    return nil
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregated"] = &contents.Aggregated
    children["unaggregated"] = &contents.Unaggregated
    return children
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bucket-type"] = contents.BucketType
    return leafs
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetBundleName() string { return "cisco_ios_xr" }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetYangName() string { return "contents" }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) SetParent(parent types.Entity) { contents.parent = parent }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetParent() types.Entity { return contents.parent }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents) GetParentYangName() string { return "bucket" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetFilter() yfilter.YFilter { return aggregated.YFilter }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) SetFilter(yf yfilter.YFilter) { aggregated.YFilter = yf }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetGoName(yname string) string {
    if yname == "bins" { return "Bins" }
    return ""
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetSegmentPath() string {
    return "aggregated"
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bins" {
        for _, c := range aggregated.Bins {
            if aggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins{}
        aggregated.Bins = append(aggregated.Bins, child)
        return &aggregated.Bins[len(aggregated.Bins)-1]
    }
    return nil
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aggregated.Bins {
        children[aggregated.Bins[i].GetSegmentPath()] = &aggregated.Bins[i]
    }
    return children
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetBundleName() string { return "cisco_ios_xr" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetYangName() string { return "aggregated" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) SetParent(parent types.Entity) { aggregated.parent = parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetParent() types.Entity { return aggregated.parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    parent types.Entity
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

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetFilter() yfilter.YFilter { return bins.YFilter }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) SetFilter(yf yfilter.YFilter) { bins.YFilter = yf }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "lower-bound-tenths" { return "LowerBoundTenths" }
    if yname == "upper-bound-tenths" { return "UpperBoundTenths" }
    if yname == "sum" { return "Sum" }
    if yname == "count" { return "Count" }
    return ""
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetSegmentPath() string {
    return "bins"
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = bins.LowerBound
    leafs["upper-bound"] = bins.UpperBound
    leafs["lower-bound-tenths"] = bins.LowerBoundTenths
    leafs["upper-bound-tenths"] = bins.UpperBoundTenths
    leafs["sum"] = bins.Sum
    leafs["count"] = bins.Count
    return leafs
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleName() string { return "cisco_ios_xr" }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetYangName() string { return "bins" }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) SetParent(parent types.Entity) { bins.parent = parent }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParent() types.Entity { return bins.parent }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParentYangName() string { return "aggregated" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetFilter() yfilter.YFilter { return unaggregated.YFilter }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) SetFilter(yf yfilter.YFilter) { unaggregated.YFilter = yf }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetSegmentPath() string {
    return "unaggregated"
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range unaggregated.Sample {
            if unaggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample{}
        unaggregated.Sample = append(unaggregated.Sample, child)
        return &unaggregated.Sample[len(unaggregated.Sample)-1]
    }
    return nil
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unaggregated.Sample {
        children[unaggregated.Sample[i].GetSegmentPath()] = &unaggregated.Sample[i]
    }
    return children
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetBundleName() string { return "cisco_ios_xr" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetYangName() string { return "unaggregated" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) SetParent(parent types.Entity) { unaggregated.parent = parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetParent() types.Entity { return unaggregated.parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    parent types.Entity
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

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetGoName(yname string) string {
    if yname == "sent-at" { return "SentAt" }
    if yname == "sent" { return "Sent" }
    if yname == "timed-out" { return "TimedOut" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "no-data-packets" { return "NoDataPackets" }
    if yname == "result" { return "Result" }
    if yname == "frames-sent" { return "FramesSent" }
    if yname == "frames-lost" { return "FramesLost" }
    return ""
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetSegmentPath() string {
    return "sample"
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent-at"] = sample.SentAt
    leafs["sent"] = sample.Sent
    leafs["timed-out"] = sample.TimedOut
    leafs["corrupt"] = sample.Corrupt
    leafs["out-of-order"] = sample.OutOfOrder
    leafs["no-data-packets"] = sample.NoDataPackets
    leafs["result"] = sample.Result
    leafs["frames-sent"] = sample.FramesSent
    leafs["frames-lost"] = sample.FramesLost
    return leafs
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetYangName() string { return "sample" }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParent() types.Entity { return sample.parent }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandCurrents_StatisticsOnDemandCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParentYangName() string { return "unaggregated" }

// Sla_Protocols_Ethernet_Operations
// Table of SLA operations
type Sla_Protocols_Ethernet_Operations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SLA operation to get operation data for. The type is slice of
    // Sla_Protocols_Ethernet_Operations_Operation.
    Operation []Sla_Protocols_Ethernet_Operations_Operation
}

func (operations *Sla_Protocols_Ethernet_Operations) GetFilter() yfilter.YFilter { return operations.YFilter }

func (operations *Sla_Protocols_Ethernet_Operations) SetFilter(yf yfilter.YFilter) { operations.YFilter = yf }

func (operations *Sla_Protocols_Ethernet_Operations) GetGoName(yname string) string {
    if yname == "operation" { return "Operation" }
    return ""
}

func (operations *Sla_Protocols_Ethernet_Operations) GetSegmentPath() string {
    return "operations"
}

func (operations *Sla_Protocols_Ethernet_Operations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operation" {
        for _, c := range operations.Operation {
            if operations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_Operations_Operation{}
        operations.Operation = append(operations.Operation, child)
        return &operations.Operation[len(operations.Operation)-1]
    }
    return nil
}

func (operations *Sla_Protocols_Ethernet_Operations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range operations.Operation {
        children[operations.Operation[i].GetSegmentPath()] = &operations.Operation[i]
    }
    return children
}

func (operations *Sla_Protocols_Ethernet_Operations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operations *Sla_Protocols_Ethernet_Operations) GetBundleName() string { return "cisco_ios_xr" }

func (operations *Sla_Protocols_Ethernet_Operations) GetYangName() string { return "operations" }

func (operations *Sla_Protocols_Ethernet_Operations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operations *Sla_Protocols_Ethernet_Operations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operations *Sla_Protocols_Ethernet_Operations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operations *Sla_Protocols_Ethernet_Operations) SetParent(parent types.Entity) { operations.parent = parent }

func (operations *Sla_Protocols_Ethernet_Operations) GetParent() types.Entity { return operations.parent }

func (operations *Sla_Protocols_Ethernet_Operations) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_Operations_Operation
// SLA operation to get operation data for
type Sla_Protocols_Ethernet_Operations_Operation struct {
    parent types.Entity
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

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetFilter() yfilter.YFilter { return operation.YFilter }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) SetFilter(yf yfilter.YFilter) { operation.YFilter = yf }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "last-run" { return "LastRun" }
    if yname == "profile-options" { return "ProfileOptions" }
    if yname == "specific-options" { return "SpecificOptions" }
    return ""
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetSegmentPath() string {
    return "operation"
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile-options" {
        return &operation.ProfileOptions
    }
    if childYangName == "specific-options" {
        return &operation.SpecificOptions
    }
    return nil
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profile-options"] = &operation.ProfileOptions
    children["specific-options"] = &operation.SpecificOptions
    return children
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = operation.ProfileName
    leafs["domain-name"] = operation.DomainName
    leafs["interface-name"] = operation.InterfaceName
    leafs["mep-id"] = operation.MepId
    leafs["mac-address"] = operation.MacAddress
    leafs["display-short"] = operation.DisplayShort
    leafs["display-long"] = operation.DisplayLong
    leafs["last-run"] = operation.LastRun
    return leafs
}

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetBundleName() string { return "cisco_ios_xr" }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetYangName() string { return "operation" }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) SetParent(parent types.Entity) { operation.parent = parent }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetParent() types.Entity { return operation.parent }

func (operation *Sla_Protocols_Ethernet_Operations_Operation) GetParentYangName() string { return "operations" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions
// Options that are only valid if the operation has
// a profile
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetFilter() yfilter.YFilter { return profileOptions.YFilter }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) SetFilter(yf yfilter.YFilter) { profileOptions.YFilter = yf }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetGoName(yname string) string {
    if yname == "probe-type" { return "ProbeType" }
    if yname == "packets-per-burst" { return "PacketsPerBurst" }
    if yname == "inter-packet-interval" { return "InterPacketInterval" }
    if yname == "bursts-per-probe" { return "BurstsPerProbe" }
    if yname == "inter-burst-interval" { return "InterBurstInterval" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "packet-padding" { return "PacketPadding" }
    if yname == "priority" { return "Priority" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetSegmentPath() string {
    return "profile-options"
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-padding" {
        return &profileOptions.PacketPadding
    }
    if childYangName == "priority" {
        return &profileOptions.Priority
    }
    if childYangName == "operation-schedule" {
        return &profileOptions.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range profileOptions.OperationMetric {
            if profileOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric{}
        profileOptions.OperationMetric = append(profileOptions.OperationMetric, child)
        return &profileOptions.OperationMetric[len(profileOptions.OperationMetric)-1]
    }
    return nil
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-padding"] = &profileOptions.PacketPadding
    children["priority"] = &profileOptions.Priority
    children["operation-schedule"] = &profileOptions.OperationSchedule
    for i := range profileOptions.OperationMetric {
        children[profileOptions.OperationMetric[i].GetSegmentPath()] = &profileOptions.OperationMetric[i]
    }
    return children
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probe-type"] = profileOptions.ProbeType
    leafs["packets-per-burst"] = profileOptions.PacketsPerBurst
    leafs["inter-packet-interval"] = profileOptions.InterPacketInterval
    leafs["bursts-per-probe"] = profileOptions.BurstsPerProbe
    leafs["inter-burst-interval"] = profileOptions.InterBurstInterval
    leafs["flr-calculation-interval"] = profileOptions.FlrCalculationInterval
    return leafs
}

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetBundleName() string { return "cisco_ios_xr" }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetYangName() string { return "profile-options" }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) SetParent(parent types.Entity) { profileOptions.parent = parent }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetParent() types.Entity { return profileOptions.parent }

func (profileOptions *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions) GetParentYangName() string { return "operation" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding
// Configuration of the packet padding
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding struct {
    parent types.Entity
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

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetFilter() yfilter.YFilter { return packetPadding.YFilter }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) SetFilter(yf yfilter.YFilter) { packetPadding.YFilter = yf }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetGoName(yname string) string {
    if yname == "packet-pad-size" { return "PacketPadSize" }
    if yname == "test-pattern-pad-scheme" { return "TestPatternPadScheme" }
    if yname == "test-pattern-pad-hex-string" { return "TestPatternPadHexString" }
    return ""
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetSegmentPath() string {
    return "packet-padding"
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-pad-size"] = packetPadding.PacketPadSize
    leafs["test-pattern-pad-scheme"] = packetPadding.TestPatternPadScheme
    leafs["test-pattern-pad-hex-string"] = packetPadding.TestPatternPadHexString
    return leafs
}

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetBundleName() string { return "cisco_ios_xr" }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetYangName() string { return "packet-padding" }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) SetParent(parent types.Entity) { packetPadding.parent = parent }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetParent() types.Entity { return packetPadding.parent }

func (packetPadding *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_PacketPadding) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority
// Priority at which to send the packet, if
// configured
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PriorityType. The type is SlaOperPacketPriority.
    PriorityType interface{}

    // 3-bit COS priority value applied to packets. The type is interface{} with
    // range: 0..255.
    Cos interface{}
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetFilter() yfilter.YFilter { return priority.YFilter }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) SetFilter(yf yfilter.YFilter) { priority.YFilter = yf }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetGoName(yname string) string {
    if yname == "priority-type" { return "PriorityType" }
    if yname == "cos" { return "Cos" }
    return ""
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetSegmentPath() string {
    return "priority"
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority-type"] = priority.PriorityType
    leafs["cos"] = priority.Cos
    return leafs
}

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetBundleName() string { return "cisco_ios_xr" }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetYangName() string { return "priority" }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) SetParent(parent types.Entity) { priority.parent = parent }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetParent() types.Entity { return priority.parent }

func (priority *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_Priority) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationSchedule) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric
// Array of the metrics that are measured by the
// operation
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of valid buckets currently in the buckets archive. The type is
    // interface{} with range: 0..4294967295.
    CurrentBucketsArchive interface{}

    // Configuration of the metric.
    MetricConfig Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetGoName(yname string) string {
    if yname == "current-buckets-archive" { return "CurrentBucketsArchive" }
    if yname == "metric-config" { return "MetricConfig" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric-config" {
        return &operationMetric.MetricConfig
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["metric-config"] = &operationMetric.MetricConfig
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-buckets-archive"] = operationMetric.CurrentBucketsArchive
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig
// Configuration of the metric
type Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig struct {
    parent types.Entity
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

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetFilter() yfilter.YFilter { return metricConfig.YFilter }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) SetFilter(yf yfilter.YFilter) { metricConfig.YFilter = yf }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetSegmentPath() string {
    return "metric-config"
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = metricConfig.MetricType
    leafs["bins-count"] = metricConfig.BinsCount
    leafs["bins-width"] = metricConfig.BinsWidth
    leafs["bucket-size"] = metricConfig.BucketSize
    leafs["bucket-size-unit"] = metricConfig.BucketSizeUnit
    leafs["buckets-archive"] = metricConfig.BucketsArchive
    return leafs
}

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetBundleName() string { return "cisco_ios_xr" }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetYangName() string { return "metric-config" }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) SetParent(parent types.Entity) { metricConfig.parent = parent }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetParent() types.Entity { return metricConfig.parent }

func (metricConfig *Sla_Protocols_Ethernet_Operations_Operation_ProfileOptions_OperationMetric_MetricConfig) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions) GetParentYangName() string { return "operation" }

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_Operations_Operation_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsHistoricals
// Table of historical statistics for SLA
// operations
type Sla_Protocols_Ethernet_StatisticsHistoricals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Historical statistics data for an SLA configured operation. The type is
    // slice of Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical.
    StatisticsHistorical []Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetFilter() yfilter.YFilter { return statisticsHistoricals.YFilter }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) SetFilter(yf yfilter.YFilter) { statisticsHistoricals.YFilter = yf }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetGoName(yname string) string {
    if yname == "statistics-historical" { return "StatisticsHistorical" }
    return ""
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetSegmentPath() string {
    return "statistics-historicals"
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-historical" {
        for _, c := range statisticsHistoricals.StatisticsHistorical {
            if statisticsHistoricals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical{}
        statisticsHistoricals.StatisticsHistorical = append(statisticsHistoricals.StatisticsHistorical, child)
        return &statisticsHistoricals.StatisticsHistorical[len(statisticsHistoricals.StatisticsHistorical)-1]
    }
    return nil
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticsHistoricals.StatisticsHistorical {
        children[statisticsHistoricals.StatisticsHistorical[i].GetSegmentPath()] = &statisticsHistoricals.StatisticsHistorical[i]
    }
    return children
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetYangName() string { return "statistics-historicals" }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) SetParent(parent types.Entity) { statisticsHistoricals.parent = parent }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetParent() types.Entity { return statisticsHistoricals.parent }

func (statisticsHistoricals *Sla_Protocols_Ethernet_StatisticsHistoricals) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical
// Historical statistics data for an SLA
// configured operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetFilter() yfilter.YFilter { return statisticsHistorical.YFilter }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) SetFilter(yf yfilter.YFilter) { statisticsHistorical.YFilter = yf }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "probe-type" { return "ProbeType" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "specific-options" { return "SpecificOptions" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetSegmentPath() string {
    return "statistics-historical"
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "specific-options" {
        return &statisticsHistorical.SpecificOptions
    }
    if childYangName == "operation-schedule" {
        return &statisticsHistorical.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range statisticsHistorical.OperationMetric {
            if statisticsHistorical.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric{}
        statisticsHistorical.OperationMetric = append(statisticsHistorical.OperationMetric, child)
        return &statisticsHistorical.OperationMetric[len(statisticsHistorical.OperationMetric)-1]
    }
    return nil
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["specific-options"] = &statisticsHistorical.SpecificOptions
    children["operation-schedule"] = &statisticsHistorical.OperationSchedule
    for i := range statisticsHistorical.OperationMetric {
        children[statisticsHistorical.OperationMetric[i].GetSegmentPath()] = &statisticsHistorical.OperationMetric[i]
    }
    return children
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = statisticsHistorical.ProfileName
    leafs["domain-name"] = statisticsHistorical.DomainName
    leafs["interface-name"] = statisticsHistorical.InterfaceName
    leafs["mep-id"] = statisticsHistorical.MepId
    leafs["mac-address"] = statisticsHistorical.MacAddress
    leafs["probe-type"] = statisticsHistorical.ProbeType
    leafs["display-short"] = statisticsHistorical.DisplayShort
    leafs["display-long"] = statisticsHistorical.DisplayLong
    leafs["flr-calculation-interval"] = statisticsHistorical.FlrCalculationInterval
    return leafs
}

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetYangName() string { return "statistics-historical" }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) SetParent(parent types.Entity) { statisticsHistorical.parent = parent }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetParent() types.Entity { return statisticsHistorical.parent }

func (statisticsHistorical *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical) GetParentYangName() string { return "statistics-historicals" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions) GetParentYangName() string { return "statistics-historical" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationSchedule) GetParentYangName() string { return "statistics-historical" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket.
    Bucket []Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "bucket" { return "Bucket" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &operationMetric.Config
    }
    if childYangName == "bucket" {
        for _, c := range operationMetric.Bucket {
            if operationMetric.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket{}
        operationMetric.Bucket = append(operationMetric.Bucket, child)
        return &operationMetric.Bucket[len(operationMetric.Bucket)-1]
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &operationMetric.Config
    for i := range operationMetric.Bucket {
        children[operationMetric.Bucket[i].GetSegmentPath()] = &operationMetric.Bucket[i]
    }
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric) GetParentYangName() string { return "statistics-historical" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config struct {
    parent types.Entity
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

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetSegmentPath() string {
    return "config"
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = config.MetricType
    leafs["bins-count"] = config.BinsCount
    leafs["bins-width"] = config.BinsWidth
    leafs["bucket-size"] = config.BucketSize
    leafs["bucket-size-unit"] = config.BucketSizeUnit
    leafs["buckets-archive"] = config.BucketsArchive
    return leafs
}

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetYangName() string { return "config" }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetParent() types.Entity { return config.parent }

func (config *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Config) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket struct {
    parent types.Entity
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

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetFilter() yfilter.YFilter { return bucket.YFilter }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) SetFilter(yf yfilter.YFilter) { bucket.YFilter = yf }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetGoName(yname string) string {
    if yname == "start-at" { return "StartAt" }
    if yname == "duration" { return "Duration" }
    if yname == "sent" { return "Sent" }
    if yname == "lost" { return "Lost" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "duplicates" { return "Duplicates" }
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    if yname == "time-of-minimum" { return "TimeOfMinimum" }
    if yname == "time-of-maximum" { return "TimeOfMaximum" }
    if yname == "average" { return "Average" }
    if yname == "standard-deviation" { return "StandardDeviation" }
    if yname == "result-count" { return "ResultCount" }
    if yname == "data-sent-count" { return "DataSentCount" }
    if yname == "data-lost-count" { return "DataLostCount" }
    if yname == "overall-flr" { return "OverallFlr" }
    if yname == "suspect-start-mid-bucket" { return "SuspectStartMidBucket" }
    if yname == "suspect-schedule-latency" { return "SuspectScheduleLatency" }
    if yname == "suspect-send-fail" { return "SuspectSendFail" }
    if yname == "suspect-premature-end" { return "SuspectPrematureEnd" }
    if yname == "suspect-clock-drift" { return "SuspectClockDrift" }
    if yname == "suspect-memory-allocation-failed" { return "SuspectMemoryAllocationFailed" }
    if yname == "suspect-cleared-mid-bucket" { return "SuspectClearedMidBucket" }
    if yname == "suspect-probe-restarted" { return "SuspectProbeRestarted" }
    if yname == "suspect-management-latency" { return "SuspectManagementLatency" }
    if yname == "suspect-multiple-buckets" { return "SuspectMultipleBuckets" }
    if yname == "suspect-misordering" { return "SuspectMisordering" }
    if yname == "suspect-flr-low-packet-count" { return "SuspectFlrLowPacketCount" }
    if yname == "premature-reason" { return "PrematureReason" }
    if yname == "premature-reason-string" { return "PrematureReasonString" }
    if yname == "contents" { return "Contents" }
    return ""
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetSegmentPath() string {
    return "bucket"
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contents" {
        return &bucket.Contents
    }
    return nil
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contents"] = &bucket.Contents
    return children
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-at"] = bucket.StartAt
    leafs["duration"] = bucket.Duration
    leafs["sent"] = bucket.Sent
    leafs["lost"] = bucket.Lost
    leafs["corrupt"] = bucket.Corrupt
    leafs["out-of-order"] = bucket.OutOfOrder
    leafs["duplicates"] = bucket.Duplicates
    leafs["minimum"] = bucket.Minimum
    leafs["maximum"] = bucket.Maximum
    leafs["time-of-minimum"] = bucket.TimeOfMinimum
    leafs["time-of-maximum"] = bucket.TimeOfMaximum
    leafs["average"] = bucket.Average
    leafs["standard-deviation"] = bucket.StandardDeviation
    leafs["result-count"] = bucket.ResultCount
    leafs["data-sent-count"] = bucket.DataSentCount
    leafs["data-lost-count"] = bucket.DataLostCount
    leafs["overall-flr"] = bucket.OverallFlr
    leafs["suspect-start-mid-bucket"] = bucket.SuspectStartMidBucket
    leafs["suspect-schedule-latency"] = bucket.SuspectScheduleLatency
    leafs["suspect-send-fail"] = bucket.SuspectSendFail
    leafs["suspect-premature-end"] = bucket.SuspectPrematureEnd
    leafs["suspect-clock-drift"] = bucket.SuspectClockDrift
    leafs["suspect-memory-allocation-failed"] = bucket.SuspectMemoryAllocationFailed
    leafs["suspect-cleared-mid-bucket"] = bucket.SuspectClearedMidBucket
    leafs["suspect-probe-restarted"] = bucket.SuspectProbeRestarted
    leafs["suspect-management-latency"] = bucket.SuspectManagementLatency
    leafs["suspect-multiple-buckets"] = bucket.SuspectMultipleBuckets
    leafs["suspect-misordering"] = bucket.SuspectMisordering
    leafs["suspect-flr-low-packet-count"] = bucket.SuspectFlrLowPacketCount
    leafs["premature-reason"] = bucket.PrematureReason
    leafs["premature-reason-string"] = bucket.PrematureReasonString
    return leafs
}

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetBundleName() string { return "cisco_ios_xr" }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetYangName() string { return "bucket" }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) SetParent(parent types.Entity) { bucket.parent = parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetParent() types.Entity { return bucket.parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetFilter() yfilter.YFilter { return contents.YFilter }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) SetFilter(yf yfilter.YFilter) { contents.YFilter = yf }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetGoName(yname string) string {
    if yname == "bucket-type" { return "BucketType" }
    if yname == "aggregated" { return "Aggregated" }
    if yname == "unaggregated" { return "Unaggregated" }
    return ""
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetSegmentPath() string {
    return "contents"
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregated" {
        return &contents.Aggregated
    }
    if childYangName == "unaggregated" {
        return &contents.Unaggregated
    }
    return nil
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregated"] = &contents.Aggregated
    children["unaggregated"] = &contents.Unaggregated
    return children
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bucket-type"] = contents.BucketType
    return leafs
}

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetBundleName() string { return "cisco_ios_xr" }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetYangName() string { return "contents" }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) SetParent(parent types.Entity) { contents.parent = parent }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetParent() types.Entity { return contents.parent }

func (contents *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents) GetParentYangName() string { return "bucket" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetFilter() yfilter.YFilter { return aggregated.YFilter }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) SetFilter(yf yfilter.YFilter) { aggregated.YFilter = yf }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetGoName(yname string) string {
    if yname == "bins" { return "Bins" }
    return ""
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetSegmentPath() string {
    return "aggregated"
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bins" {
        for _, c := range aggregated.Bins {
            if aggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins{}
        aggregated.Bins = append(aggregated.Bins, child)
        return &aggregated.Bins[len(aggregated.Bins)-1]
    }
    return nil
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aggregated.Bins {
        children[aggregated.Bins[i].GetSegmentPath()] = &aggregated.Bins[i]
    }
    return children
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetBundleName() string { return "cisco_ios_xr" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetYangName() string { return "aggregated" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) SetParent(parent types.Entity) { aggregated.parent = parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetParent() types.Entity { return aggregated.parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    parent types.Entity
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

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetFilter() yfilter.YFilter { return bins.YFilter }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) SetFilter(yf yfilter.YFilter) { bins.YFilter = yf }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "lower-bound-tenths" { return "LowerBoundTenths" }
    if yname == "upper-bound-tenths" { return "UpperBoundTenths" }
    if yname == "sum" { return "Sum" }
    if yname == "count" { return "Count" }
    return ""
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetSegmentPath() string {
    return "bins"
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = bins.LowerBound
    leafs["upper-bound"] = bins.UpperBound
    leafs["lower-bound-tenths"] = bins.LowerBoundTenths
    leafs["upper-bound-tenths"] = bins.UpperBoundTenths
    leafs["sum"] = bins.Sum
    leafs["count"] = bins.Count
    return leafs
}

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleName() string { return "cisco_ios_xr" }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetYangName() string { return "bins" }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) SetParent(parent types.Entity) { bins.parent = parent }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParent() types.Entity { return bins.parent }

func (bins *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParentYangName() string { return "aggregated" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetFilter() yfilter.YFilter { return unaggregated.YFilter }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) SetFilter(yf yfilter.YFilter) { unaggregated.YFilter = yf }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetSegmentPath() string {
    return "unaggregated"
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range unaggregated.Sample {
            if unaggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample{}
        unaggregated.Sample = append(unaggregated.Sample, child)
        return &unaggregated.Sample[len(unaggregated.Sample)-1]
    }
    return nil
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unaggregated.Sample {
        children[unaggregated.Sample[i].GetSegmentPath()] = &unaggregated.Sample[i]
    }
    return children
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetBundleName() string { return "cisco_ios_xr" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetYangName() string { return "unaggregated" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) SetParent(parent types.Entity) { unaggregated.parent = parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetParent() types.Entity { return unaggregated.parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    parent types.Entity
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

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetGoName(yname string) string {
    if yname == "sent-at" { return "SentAt" }
    if yname == "sent" { return "Sent" }
    if yname == "timed-out" { return "TimedOut" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "no-data-packets" { return "NoDataPackets" }
    if yname == "result" { return "Result" }
    if yname == "frames-sent" { return "FramesSent" }
    if yname == "frames-lost" { return "FramesLost" }
    return ""
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetSegmentPath() string {
    return "sample"
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent-at"] = sample.SentAt
    leafs["sent"] = sample.Sent
    leafs["timed-out"] = sample.TimedOut
    leafs["corrupt"] = sample.Corrupt
    leafs["out-of-order"] = sample.OutOfOrder
    leafs["no-data-packets"] = sample.NoDataPackets
    leafs["result"] = sample.Result
    leafs["frames-sent"] = sample.FramesSent
    leafs["frames-lost"] = sample.FramesLost
    return leafs
}

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetYangName() string { return "sample" }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParent() types.Entity { return sample.parent }

func (sample *Sla_Protocols_Ethernet_StatisticsHistoricals_StatisticsHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParentYangName() string { return "unaggregated" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals
// Table of historical statistics for SLA
// on-demand operations
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Historical statistics data for an SLA on-demand  operation. The type is
    // slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical.
    StatisticsOnDemandHistorical []Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetFilter() yfilter.YFilter { return statisticsOnDemandHistoricals.YFilter }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) SetFilter(yf yfilter.YFilter) { statisticsOnDemandHistoricals.YFilter = yf }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetGoName(yname string) string {
    if yname == "statistics-on-demand-historical" { return "StatisticsOnDemandHistorical" }
    return ""
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetSegmentPath() string {
    return "statistics-on-demand-historicals"
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-on-demand-historical" {
        for _, c := range statisticsOnDemandHistoricals.StatisticsOnDemandHistorical {
            if statisticsOnDemandHistoricals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical{}
        statisticsOnDemandHistoricals.StatisticsOnDemandHistorical = append(statisticsOnDemandHistoricals.StatisticsOnDemandHistorical, child)
        return &statisticsOnDemandHistoricals.StatisticsOnDemandHistorical[len(statisticsOnDemandHistoricals.StatisticsOnDemandHistorical)-1]
    }
    return nil
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticsOnDemandHistoricals.StatisticsOnDemandHistorical {
        children[statisticsOnDemandHistoricals.StatisticsOnDemandHistorical[i].GetSegmentPath()] = &statisticsOnDemandHistoricals.StatisticsOnDemandHistorical[i]
    }
    return children
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetYangName() string { return "statistics-on-demand-historicals" }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) SetParent(parent types.Entity) { statisticsOnDemandHistoricals.parent = parent }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetParent() types.Entity { return statisticsOnDemandHistoricals.parent }

func (statisticsOnDemandHistoricals *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical
// Historical statistics data for an SLA
// on-demand  operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetFilter() yfilter.YFilter { return statisticsOnDemandHistorical.YFilter }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) SetFilter(yf yfilter.YFilter) { statisticsOnDemandHistorical.YFilter = yf }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "probe-type" { return "ProbeType" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "specific-options" { return "SpecificOptions" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetSegmentPath() string {
    return "statistics-on-demand-historical"
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "specific-options" {
        return &statisticsOnDemandHistorical.SpecificOptions
    }
    if childYangName == "operation-schedule" {
        return &statisticsOnDemandHistorical.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range statisticsOnDemandHistorical.OperationMetric {
            if statisticsOnDemandHistorical.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric{}
        statisticsOnDemandHistorical.OperationMetric = append(statisticsOnDemandHistorical.OperationMetric, child)
        return &statisticsOnDemandHistorical.OperationMetric[len(statisticsOnDemandHistorical.OperationMetric)-1]
    }
    return nil
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["specific-options"] = &statisticsOnDemandHistorical.SpecificOptions
    children["operation-schedule"] = &statisticsOnDemandHistorical.OperationSchedule
    for i := range statisticsOnDemandHistorical.OperationMetric {
        children[statisticsOnDemandHistorical.OperationMetric[i].GetSegmentPath()] = &statisticsOnDemandHistorical.OperationMetric[i]
    }
    return children
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = statisticsOnDemandHistorical.OperationId
    leafs["domain-name"] = statisticsOnDemandHistorical.DomainName
    leafs["interface-name"] = statisticsOnDemandHistorical.InterfaceName
    leafs["mep-id"] = statisticsOnDemandHistorical.MepId
    leafs["mac-address"] = statisticsOnDemandHistorical.MacAddress
    leafs["probe-type"] = statisticsOnDemandHistorical.ProbeType
    leafs["display-short"] = statisticsOnDemandHistorical.DisplayShort
    leafs["display-long"] = statisticsOnDemandHistorical.DisplayLong
    leafs["flr-calculation-interval"] = statisticsOnDemandHistorical.FlrCalculationInterval
    return leafs
}

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetYangName() string { return "statistics-on-demand-historical" }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) SetParent(parent types.Entity) { statisticsOnDemandHistorical.parent = parent }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetParent() types.Entity { return statisticsOnDemandHistorical.parent }

func (statisticsOnDemandHistorical *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical) GetParentYangName() string { return "statistics-on-demand-historicals" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions) GetParentYangName() string { return "statistics-on-demand-historical" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationSchedule) GetParentYangName() string { return "statistics-on-demand-historical" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket.
    Bucket []Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "bucket" { return "Bucket" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &operationMetric.Config
    }
    if childYangName == "bucket" {
        for _, c := range operationMetric.Bucket {
            if operationMetric.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket{}
        operationMetric.Bucket = append(operationMetric.Bucket, child)
        return &operationMetric.Bucket[len(operationMetric.Bucket)-1]
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &operationMetric.Config
    for i := range operationMetric.Bucket {
        children[operationMetric.Bucket[i].GetSegmentPath()] = &operationMetric.Bucket[i]
    }
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric) GetParentYangName() string { return "statistics-on-demand-historical" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config struct {
    parent types.Entity
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

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetSegmentPath() string {
    return "config"
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = config.MetricType
    leafs["bins-count"] = config.BinsCount
    leafs["bins-width"] = config.BinsWidth
    leafs["bucket-size"] = config.BucketSize
    leafs["bucket-size-unit"] = config.BucketSizeUnit
    leafs["buckets-archive"] = config.BucketsArchive
    return leafs
}

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetYangName() string { return "config" }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetParent() types.Entity { return config.parent }

func (config *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Config) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket struct {
    parent types.Entity
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

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetFilter() yfilter.YFilter { return bucket.YFilter }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) SetFilter(yf yfilter.YFilter) { bucket.YFilter = yf }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetGoName(yname string) string {
    if yname == "start-at" { return "StartAt" }
    if yname == "duration" { return "Duration" }
    if yname == "sent" { return "Sent" }
    if yname == "lost" { return "Lost" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "duplicates" { return "Duplicates" }
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    if yname == "time-of-minimum" { return "TimeOfMinimum" }
    if yname == "time-of-maximum" { return "TimeOfMaximum" }
    if yname == "average" { return "Average" }
    if yname == "standard-deviation" { return "StandardDeviation" }
    if yname == "result-count" { return "ResultCount" }
    if yname == "data-sent-count" { return "DataSentCount" }
    if yname == "data-lost-count" { return "DataLostCount" }
    if yname == "overall-flr" { return "OverallFlr" }
    if yname == "suspect-start-mid-bucket" { return "SuspectStartMidBucket" }
    if yname == "suspect-schedule-latency" { return "SuspectScheduleLatency" }
    if yname == "suspect-send-fail" { return "SuspectSendFail" }
    if yname == "suspect-premature-end" { return "SuspectPrematureEnd" }
    if yname == "suspect-clock-drift" { return "SuspectClockDrift" }
    if yname == "suspect-memory-allocation-failed" { return "SuspectMemoryAllocationFailed" }
    if yname == "suspect-cleared-mid-bucket" { return "SuspectClearedMidBucket" }
    if yname == "suspect-probe-restarted" { return "SuspectProbeRestarted" }
    if yname == "suspect-management-latency" { return "SuspectManagementLatency" }
    if yname == "suspect-multiple-buckets" { return "SuspectMultipleBuckets" }
    if yname == "suspect-misordering" { return "SuspectMisordering" }
    if yname == "suspect-flr-low-packet-count" { return "SuspectFlrLowPacketCount" }
    if yname == "premature-reason" { return "PrematureReason" }
    if yname == "premature-reason-string" { return "PrematureReasonString" }
    if yname == "contents" { return "Contents" }
    return ""
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetSegmentPath() string {
    return "bucket"
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contents" {
        return &bucket.Contents
    }
    return nil
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contents"] = &bucket.Contents
    return children
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-at"] = bucket.StartAt
    leafs["duration"] = bucket.Duration
    leafs["sent"] = bucket.Sent
    leafs["lost"] = bucket.Lost
    leafs["corrupt"] = bucket.Corrupt
    leafs["out-of-order"] = bucket.OutOfOrder
    leafs["duplicates"] = bucket.Duplicates
    leafs["minimum"] = bucket.Minimum
    leafs["maximum"] = bucket.Maximum
    leafs["time-of-minimum"] = bucket.TimeOfMinimum
    leafs["time-of-maximum"] = bucket.TimeOfMaximum
    leafs["average"] = bucket.Average
    leafs["standard-deviation"] = bucket.StandardDeviation
    leafs["result-count"] = bucket.ResultCount
    leafs["data-sent-count"] = bucket.DataSentCount
    leafs["data-lost-count"] = bucket.DataLostCount
    leafs["overall-flr"] = bucket.OverallFlr
    leafs["suspect-start-mid-bucket"] = bucket.SuspectStartMidBucket
    leafs["suspect-schedule-latency"] = bucket.SuspectScheduleLatency
    leafs["suspect-send-fail"] = bucket.SuspectSendFail
    leafs["suspect-premature-end"] = bucket.SuspectPrematureEnd
    leafs["suspect-clock-drift"] = bucket.SuspectClockDrift
    leafs["suspect-memory-allocation-failed"] = bucket.SuspectMemoryAllocationFailed
    leafs["suspect-cleared-mid-bucket"] = bucket.SuspectClearedMidBucket
    leafs["suspect-probe-restarted"] = bucket.SuspectProbeRestarted
    leafs["suspect-management-latency"] = bucket.SuspectManagementLatency
    leafs["suspect-multiple-buckets"] = bucket.SuspectMultipleBuckets
    leafs["suspect-misordering"] = bucket.SuspectMisordering
    leafs["suspect-flr-low-packet-count"] = bucket.SuspectFlrLowPacketCount
    leafs["premature-reason"] = bucket.PrematureReason
    leafs["premature-reason-string"] = bucket.PrematureReasonString
    return leafs
}

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetBundleName() string { return "cisco_ios_xr" }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetYangName() string { return "bucket" }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) SetParent(parent types.Entity) { bucket.parent = parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetParent() types.Entity { return bucket.parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetFilter() yfilter.YFilter { return contents.YFilter }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) SetFilter(yf yfilter.YFilter) { contents.YFilter = yf }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetGoName(yname string) string {
    if yname == "bucket-type" { return "BucketType" }
    if yname == "aggregated" { return "Aggregated" }
    if yname == "unaggregated" { return "Unaggregated" }
    return ""
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetSegmentPath() string {
    return "contents"
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregated" {
        return &contents.Aggregated
    }
    if childYangName == "unaggregated" {
        return &contents.Unaggregated
    }
    return nil
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregated"] = &contents.Aggregated
    children["unaggregated"] = &contents.Unaggregated
    return children
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bucket-type"] = contents.BucketType
    return leafs
}

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetBundleName() string { return "cisco_ios_xr" }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetYangName() string { return "contents" }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) SetParent(parent types.Entity) { contents.parent = parent }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetParent() types.Entity { return contents.parent }

func (contents *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents) GetParentYangName() string { return "bucket" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetFilter() yfilter.YFilter { return aggregated.YFilter }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) SetFilter(yf yfilter.YFilter) { aggregated.YFilter = yf }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetGoName(yname string) string {
    if yname == "bins" { return "Bins" }
    return ""
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetSegmentPath() string {
    return "aggregated"
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bins" {
        for _, c := range aggregated.Bins {
            if aggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins{}
        aggregated.Bins = append(aggregated.Bins, child)
        return &aggregated.Bins[len(aggregated.Bins)-1]
    }
    return nil
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aggregated.Bins {
        children[aggregated.Bins[i].GetSegmentPath()] = &aggregated.Bins[i]
    }
    return children
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetBundleName() string { return "cisco_ios_xr" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetYangName() string { return "aggregated" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) SetParent(parent types.Entity) { aggregated.parent = parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetParent() types.Entity { return aggregated.parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    parent types.Entity
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

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetFilter() yfilter.YFilter { return bins.YFilter }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) SetFilter(yf yfilter.YFilter) { bins.YFilter = yf }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "lower-bound-tenths" { return "LowerBoundTenths" }
    if yname == "upper-bound-tenths" { return "UpperBoundTenths" }
    if yname == "sum" { return "Sum" }
    if yname == "count" { return "Count" }
    return ""
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetSegmentPath() string {
    return "bins"
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = bins.LowerBound
    leafs["upper-bound"] = bins.UpperBound
    leafs["lower-bound-tenths"] = bins.LowerBoundTenths
    leafs["upper-bound-tenths"] = bins.UpperBoundTenths
    leafs["sum"] = bins.Sum
    leafs["count"] = bins.Count
    return leafs
}

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleName() string { return "cisco_ios_xr" }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetYangName() string { return "bins" }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) SetParent(parent types.Entity) { bins.parent = parent }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParent() types.Entity { return bins.parent }

func (bins *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParentYangName() string { return "aggregated" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetFilter() yfilter.YFilter { return unaggregated.YFilter }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) SetFilter(yf yfilter.YFilter) { unaggregated.YFilter = yf }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetSegmentPath() string {
    return "unaggregated"
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range unaggregated.Sample {
            if unaggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample{}
        unaggregated.Sample = append(unaggregated.Sample, child)
        return &unaggregated.Sample[len(unaggregated.Sample)-1]
    }
    return nil
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unaggregated.Sample {
        children[unaggregated.Sample[i].GetSegmentPath()] = &unaggregated.Sample[i]
    }
    return children
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetBundleName() string { return "cisco_ios_xr" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetYangName() string { return "unaggregated" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) SetParent(parent types.Entity) { unaggregated.parent = parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetParent() types.Entity { return unaggregated.parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    parent types.Entity
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

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetGoName(yname string) string {
    if yname == "sent-at" { return "SentAt" }
    if yname == "sent" { return "Sent" }
    if yname == "timed-out" { return "TimedOut" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "no-data-packets" { return "NoDataPackets" }
    if yname == "result" { return "Result" }
    if yname == "frames-sent" { return "FramesSent" }
    if yname == "frames-lost" { return "FramesLost" }
    return ""
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetSegmentPath() string {
    return "sample"
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent-at"] = sample.SentAt
    leafs["sent"] = sample.Sent
    leafs["timed-out"] = sample.TimedOut
    leafs["corrupt"] = sample.Corrupt
    leafs["out-of-order"] = sample.OutOfOrder
    leafs["no-data-packets"] = sample.NoDataPackets
    leafs["result"] = sample.Result
    leafs["frames-sent"] = sample.FramesSent
    leafs["frames-lost"] = sample.FramesLost
    return leafs
}

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetYangName() string { return "sample" }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParent() types.Entity { return sample.parent }

func (sample *Sla_Protocols_Ethernet_StatisticsOnDemandHistoricals_StatisticsOnDemandHistorical_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParentYangName() string { return "unaggregated" }

// Sla_Protocols_Ethernet_ConfigErrors
// Table of SLA configuration errors on configured
// operations
type Sla_Protocols_Ethernet_ConfigErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SLA operation to get configuration errors data for. The type is slice of
    // Sla_Protocols_Ethernet_ConfigErrors_ConfigError.
    ConfigError []Sla_Protocols_Ethernet_ConfigErrors_ConfigError
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetFilter() yfilter.YFilter { return configErrors.YFilter }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) SetFilter(yf yfilter.YFilter) { configErrors.YFilter = yf }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetGoName(yname string) string {
    if yname == "config-error" { return "ConfigError" }
    return ""
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetSegmentPath() string {
    return "config-errors"
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-error" {
        for _, c := range configErrors.ConfigError {
            if configErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_ConfigErrors_ConfigError{}
        configErrors.ConfigError = append(configErrors.ConfigError, child)
        return &configErrors.ConfigError[len(configErrors.ConfigError)-1]
    }
    return nil
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configErrors.ConfigError {
        children[configErrors.ConfigError[i].GetSegmentPath()] = &configErrors.ConfigError[i]
    }
    return children
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetYangName() string { return "config-errors" }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) SetParent(parent types.Entity) { configErrors.parent = parent }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetParent() types.Entity { return configErrors.parent }

func (configErrors *Sla_Protocols_Ethernet_ConfigErrors) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_ConfigErrors_ConfigError
// SLA operation to get configuration errors data
// for
type Sla_Protocols_Ethernet_ConfigErrors_ConfigError struct {
    parent types.Entity
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

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetFilter() yfilter.YFilter { return configError.YFilter }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) SetFilter(yf yfilter.YFilter) { configError.YFilter = yf }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "profile-name-xr" { return "ProfileNameXr" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "rt-delay-inconsistent" { return "RtDelayInconsistent" }
    if yname == "ow-delay-sd-inconsistent" { return "OwDelaySdInconsistent" }
    if yname == "ow-delay-ds-inconsistent" { return "OwDelayDsInconsistent" }
    if yname == "rt-jitter-inconsistent" { return "RtJitterInconsistent" }
    if yname == "ow-jitter-sd-inconsistent" { return "OwJitterSdInconsistent" }
    if yname == "ow-jitter-ds-inconsistent" { return "OwJitterDsInconsistent" }
    if yname == "ow-loss-sd-inconsistent" { return "OwLossSdInconsistent" }
    if yname == "ow-loss-ds-inconsistent" { return "OwLossDsInconsistent" }
    if yname == "packet-pad-inconsistent" { return "PacketPadInconsistent" }
    if yname == "packet-rand-pad-inconsistent" { return "PacketRandPadInconsistent" }
    if yname == "min-packet-interval-inconsistent" { return "MinPacketIntervalInconsistent" }
    if yname == "priority-inconsistent" { return "PriorityInconsistent" }
    if yname == "packet-type-inconsistent" { return "PacketTypeInconsistent" }
    if yname == "profile-doesnt-exist" { return "ProfileDoesntExist" }
    if yname == "synthetic-loss-not-supported" { return "SyntheticLossNotSupported" }
    if yname == "probe-too-big" { return "ProbeTooBig" }
    if yname == "error-string" { return "ErrorString" }
    return ""
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetSegmentPath() string {
    return "config-error"
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configError.ProfileName
    leafs["domain-name"] = configError.DomainName
    leafs["interface-name"] = configError.InterfaceName
    leafs["mep-id"] = configError.MepId
    leafs["mac-address"] = configError.MacAddress
    leafs["profile-name-xr"] = configError.ProfileNameXr
    leafs["display-short"] = configError.DisplayShort
    leafs["rt-delay-inconsistent"] = configError.RtDelayInconsistent
    leafs["ow-delay-sd-inconsistent"] = configError.OwDelaySdInconsistent
    leafs["ow-delay-ds-inconsistent"] = configError.OwDelayDsInconsistent
    leafs["rt-jitter-inconsistent"] = configError.RtJitterInconsistent
    leafs["ow-jitter-sd-inconsistent"] = configError.OwJitterSdInconsistent
    leafs["ow-jitter-ds-inconsistent"] = configError.OwJitterDsInconsistent
    leafs["ow-loss-sd-inconsistent"] = configError.OwLossSdInconsistent
    leafs["ow-loss-ds-inconsistent"] = configError.OwLossDsInconsistent
    leafs["packet-pad-inconsistent"] = configError.PacketPadInconsistent
    leafs["packet-rand-pad-inconsistent"] = configError.PacketRandPadInconsistent
    leafs["min-packet-interval-inconsistent"] = configError.MinPacketIntervalInconsistent
    leafs["priority-inconsistent"] = configError.PriorityInconsistent
    leafs["packet-type-inconsistent"] = configError.PacketTypeInconsistent
    leafs["profile-doesnt-exist"] = configError.ProfileDoesntExist
    leafs["synthetic-loss-not-supported"] = configError.SyntheticLossNotSupported
    leafs["probe-too-big"] = configError.ProbeTooBig
    leafs["error-string"] = configError.ErrorString
    return leafs
}

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetBundleName() string { return "cisco_ios_xr" }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetYangName() string { return "config-error" }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) SetParent(parent types.Entity) { configError.parent = parent }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetParent() types.Entity { return configError.parent }

func (configError *Sla_Protocols_Ethernet_ConfigErrors_ConfigError) GetParentYangName() string { return "config-errors" }

// Sla_Protocols_Ethernet_OnDemandOperations
// Table of SLA on-demand operations
type Sla_Protocols_Ethernet_OnDemandOperations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SLA on-demand operation to get operation data for. The type is slice of
    // Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation.
    OnDemandOperation []Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetFilter() yfilter.YFilter { return onDemandOperations.YFilter }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) SetFilter(yf yfilter.YFilter) { onDemandOperations.YFilter = yf }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetGoName(yname string) string {
    if yname == "on-demand-operation" { return "OnDemandOperation" }
    return ""
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetSegmentPath() string {
    return "on-demand-operations"
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "on-demand-operation" {
        for _, c := range onDemandOperations.OnDemandOperation {
            if onDemandOperations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation{}
        onDemandOperations.OnDemandOperation = append(onDemandOperations.OnDemandOperation, child)
        return &onDemandOperations.OnDemandOperation[len(onDemandOperations.OnDemandOperation)-1]
    }
    return nil
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range onDemandOperations.OnDemandOperation {
        children[onDemandOperations.OnDemandOperation[i].GetSegmentPath()] = &onDemandOperations.OnDemandOperation[i]
    }
    return children
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetBundleName() string { return "cisco_ios_xr" }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetYangName() string { return "on-demand-operations" }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) SetParent(parent types.Entity) { onDemandOperations.parent = parent }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetParent() types.Entity { return onDemandOperations.parent }

func (onDemandOperations *Sla_Protocols_Ethernet_OnDemandOperations) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation
// SLA on-demand operation to get operation data
// for
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation struct {
    parent types.Entity
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

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetFilter() yfilter.YFilter { return onDemandOperation.YFilter }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) SetFilter(yf yfilter.YFilter) { onDemandOperation.YFilter = yf }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetGoName(yname string) string {
    if yname == "operation-id" { return "OperationId" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "last-run" { return "LastRun" }
    if yname == "profile-options" { return "ProfileOptions" }
    if yname == "specific-options" { return "SpecificOptions" }
    return ""
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetSegmentPath() string {
    return "on-demand-operation"
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile-options" {
        return &onDemandOperation.ProfileOptions
    }
    if childYangName == "specific-options" {
        return &onDemandOperation.SpecificOptions
    }
    return nil
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profile-options"] = &onDemandOperation.ProfileOptions
    children["specific-options"] = &onDemandOperation.SpecificOptions
    return children
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operation-id"] = onDemandOperation.OperationId
    leafs["domain-name"] = onDemandOperation.DomainName
    leafs["interface-name"] = onDemandOperation.InterfaceName
    leafs["mep-id"] = onDemandOperation.MepId
    leafs["mac-address"] = onDemandOperation.MacAddress
    leafs["display-short"] = onDemandOperation.DisplayShort
    leafs["display-long"] = onDemandOperation.DisplayLong
    leafs["last-run"] = onDemandOperation.LastRun
    return leafs
}

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetBundleName() string { return "cisco_ios_xr" }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetYangName() string { return "on-demand-operation" }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) SetParent(parent types.Entity) { onDemandOperation.parent = parent }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetParent() types.Entity { return onDemandOperation.parent }

func (onDemandOperation *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation) GetParentYangName() string { return "on-demand-operations" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions
// Options that are only valid if the operation has
// a profile
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetFilter() yfilter.YFilter { return profileOptions.YFilter }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) SetFilter(yf yfilter.YFilter) { profileOptions.YFilter = yf }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetGoName(yname string) string {
    if yname == "probe-type" { return "ProbeType" }
    if yname == "packets-per-burst" { return "PacketsPerBurst" }
    if yname == "inter-packet-interval" { return "InterPacketInterval" }
    if yname == "bursts-per-probe" { return "BurstsPerProbe" }
    if yname == "inter-burst-interval" { return "InterBurstInterval" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "packet-padding" { return "PacketPadding" }
    if yname == "priority" { return "Priority" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetSegmentPath() string {
    return "profile-options"
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-padding" {
        return &profileOptions.PacketPadding
    }
    if childYangName == "priority" {
        return &profileOptions.Priority
    }
    if childYangName == "operation-schedule" {
        return &profileOptions.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range profileOptions.OperationMetric {
            if profileOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric{}
        profileOptions.OperationMetric = append(profileOptions.OperationMetric, child)
        return &profileOptions.OperationMetric[len(profileOptions.OperationMetric)-1]
    }
    return nil
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-padding"] = &profileOptions.PacketPadding
    children["priority"] = &profileOptions.Priority
    children["operation-schedule"] = &profileOptions.OperationSchedule
    for i := range profileOptions.OperationMetric {
        children[profileOptions.OperationMetric[i].GetSegmentPath()] = &profileOptions.OperationMetric[i]
    }
    return children
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["probe-type"] = profileOptions.ProbeType
    leafs["packets-per-burst"] = profileOptions.PacketsPerBurst
    leafs["inter-packet-interval"] = profileOptions.InterPacketInterval
    leafs["bursts-per-probe"] = profileOptions.BurstsPerProbe
    leafs["inter-burst-interval"] = profileOptions.InterBurstInterval
    leafs["flr-calculation-interval"] = profileOptions.FlrCalculationInterval
    return leafs
}

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetBundleName() string { return "cisco_ios_xr" }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetYangName() string { return "profile-options" }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) SetParent(parent types.Entity) { profileOptions.parent = parent }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetParent() types.Entity { return profileOptions.parent }

func (profileOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions) GetParentYangName() string { return "on-demand-operation" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding
// Configuration of the packet padding
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding struct {
    parent types.Entity
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

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetFilter() yfilter.YFilter { return packetPadding.YFilter }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) SetFilter(yf yfilter.YFilter) { packetPadding.YFilter = yf }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetGoName(yname string) string {
    if yname == "packet-pad-size" { return "PacketPadSize" }
    if yname == "test-pattern-pad-scheme" { return "TestPatternPadScheme" }
    if yname == "test-pattern-pad-hex-string" { return "TestPatternPadHexString" }
    return ""
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetSegmentPath() string {
    return "packet-padding"
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-pad-size"] = packetPadding.PacketPadSize
    leafs["test-pattern-pad-scheme"] = packetPadding.TestPatternPadScheme
    leafs["test-pattern-pad-hex-string"] = packetPadding.TestPatternPadHexString
    return leafs
}

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetBundleName() string { return "cisco_ios_xr" }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetYangName() string { return "packet-padding" }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) SetParent(parent types.Entity) { packetPadding.parent = parent }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetParent() types.Entity { return packetPadding.parent }

func (packetPadding *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_PacketPadding) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority
// Priority at which to send the packet, if
// configured
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PriorityType. The type is SlaOperPacketPriority.
    PriorityType interface{}

    // 3-bit COS priority value applied to packets. The type is interface{} with
    // range: 0..255.
    Cos interface{}
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetFilter() yfilter.YFilter { return priority.YFilter }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) SetFilter(yf yfilter.YFilter) { priority.YFilter = yf }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetGoName(yname string) string {
    if yname == "priority-type" { return "PriorityType" }
    if yname == "cos" { return "Cos" }
    return ""
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetSegmentPath() string {
    return "priority"
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority-type"] = priority.PriorityType
    leafs["cos"] = priority.Cos
    return leafs
}

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetBundleName() string { return "cisco_ios_xr" }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetYangName() string { return "priority" }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) SetParent(parent types.Entity) { priority.parent = parent }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetParent() types.Entity { return priority.parent }

func (priority *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_Priority) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationSchedule) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric
// Array of the metrics that are measured by the
// operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of valid buckets currently in the buckets archive. The type is
    // interface{} with range: 0..4294967295.
    CurrentBucketsArchive interface{}

    // Configuration of the metric.
    MetricConfig Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetGoName(yname string) string {
    if yname == "current-buckets-archive" { return "CurrentBucketsArchive" }
    if yname == "metric-config" { return "MetricConfig" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric-config" {
        return &operationMetric.MetricConfig
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["metric-config"] = &operationMetric.MetricConfig
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-buckets-archive"] = operationMetric.CurrentBucketsArchive
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric) GetParentYangName() string { return "profile-options" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig
// Configuration of the metric
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig struct {
    parent types.Entity
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

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetFilter() yfilter.YFilter { return metricConfig.YFilter }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) SetFilter(yf yfilter.YFilter) { metricConfig.YFilter = yf }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetSegmentPath() string {
    return "metric-config"
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = metricConfig.MetricType
    leafs["bins-count"] = metricConfig.BinsCount
    leafs["bins-width"] = metricConfig.BinsWidth
    leafs["bucket-size"] = metricConfig.BucketSize
    leafs["bucket-size-unit"] = metricConfig.BucketSizeUnit
    leafs["buckets-archive"] = metricConfig.BucketsArchive
    return leafs
}

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetBundleName() string { return "cisco_ios_xr" }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetYangName() string { return "metric-config" }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) SetParent(parent types.Entity) { metricConfig.parent = parent }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetParent() types.Entity { return metricConfig.parent }

func (metricConfig *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_ProfileOptions_OperationMetric_MetricConfig) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions) GetParentYangName() string { return "on-demand-operation" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_OnDemandOperations_OnDemandOperation_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsCurrents
// Table of current statistics for SLA operations
type Sla_Protocols_Ethernet_StatisticsCurrents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current statistics data for an SLA configured operation. The type is slice
    // of Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent.
    StatisticsCurrent []Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetFilter() yfilter.YFilter { return statisticsCurrents.YFilter }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) SetFilter(yf yfilter.YFilter) { statisticsCurrents.YFilter = yf }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetGoName(yname string) string {
    if yname == "statistics-current" { return "StatisticsCurrent" }
    return ""
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetSegmentPath() string {
    return "statistics-currents"
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-current" {
        for _, c := range statisticsCurrents.StatisticsCurrent {
            if statisticsCurrents.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent{}
        statisticsCurrents.StatisticsCurrent = append(statisticsCurrents.StatisticsCurrent, child)
        return &statisticsCurrents.StatisticsCurrent[len(statisticsCurrents.StatisticsCurrent)-1]
    }
    return nil
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticsCurrents.StatisticsCurrent {
        children[statisticsCurrents.StatisticsCurrent[i].GetSegmentPath()] = &statisticsCurrents.StatisticsCurrent[i]
    }
    return children
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetYangName() string { return "statistics-currents" }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) SetParent(parent types.Entity) { statisticsCurrents.parent = parent }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetParent() types.Entity { return statisticsCurrents.parent }

func (statisticsCurrents *Sla_Protocols_Ethernet_StatisticsCurrents) GetParentYangName() string { return "ethernet" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent
// Current statistics data for an SLA configured
// operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent struct {
    parent types.Entity
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
    OperationMetric []Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetFilter() yfilter.YFilter { return statisticsCurrent.YFilter }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) SetFilter(yf yfilter.YFilter) { statisticsCurrent.YFilter = yf }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mep-id" { return "MepId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "probe-type" { return "ProbeType" }
    if yname == "display-short" { return "DisplayShort" }
    if yname == "display-long" { return "DisplayLong" }
    if yname == "flr-calculation-interval" { return "FlrCalculationInterval" }
    if yname == "specific-options" { return "SpecificOptions" }
    if yname == "operation-schedule" { return "OperationSchedule" }
    if yname == "operation-metric" { return "OperationMetric" }
    return ""
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetSegmentPath() string {
    return "statistics-current"
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "specific-options" {
        return &statisticsCurrent.SpecificOptions
    }
    if childYangName == "operation-schedule" {
        return &statisticsCurrent.OperationSchedule
    }
    if childYangName == "operation-metric" {
        for _, c := range statisticsCurrent.OperationMetric {
            if statisticsCurrent.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric{}
        statisticsCurrent.OperationMetric = append(statisticsCurrent.OperationMetric, child)
        return &statisticsCurrent.OperationMetric[len(statisticsCurrent.OperationMetric)-1]
    }
    return nil
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["specific-options"] = &statisticsCurrent.SpecificOptions
    children["operation-schedule"] = &statisticsCurrent.OperationSchedule
    for i := range statisticsCurrent.OperationMetric {
        children[statisticsCurrent.OperationMetric[i].GetSegmentPath()] = &statisticsCurrent.OperationMetric[i]
    }
    return children
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = statisticsCurrent.ProfileName
    leafs["domain-name"] = statisticsCurrent.DomainName
    leafs["interface-name"] = statisticsCurrent.InterfaceName
    leafs["mep-id"] = statisticsCurrent.MepId
    leafs["mac-address"] = statisticsCurrent.MacAddress
    leafs["probe-type"] = statisticsCurrent.ProbeType
    leafs["display-short"] = statisticsCurrent.DisplayShort
    leafs["display-long"] = statisticsCurrent.DisplayLong
    leafs["flr-calculation-interval"] = statisticsCurrent.FlrCalculationInterval
    return leafs
}

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetYangName() string { return "statistics-current" }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) SetParent(parent types.Entity) { statisticsCurrent.parent = parent }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetParent() types.Entity { return statisticsCurrent.parent }

func (statisticsCurrent *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent) GetParentYangName() string { return "statistics-currents" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions
// Options specific to the type of operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OperType. The type is SlaOperOperation.
    OperType interface{}

    // Parameters for a configured operation.
    ConfiguredOperationOptions Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions

    // Parameters for an ondemand operation.
    OndemandOperationOptions Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetFilter() yfilter.YFilter { return specificOptions.YFilter }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) SetFilter(yf yfilter.YFilter) { specificOptions.YFilter = yf }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetGoName(yname string) string {
    if yname == "oper-type" { return "OperType" }
    if yname == "configured-operation-options" { return "ConfiguredOperationOptions" }
    if yname == "ondemand-operation-options" { return "OndemandOperationOptions" }
    return ""
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetSegmentPath() string {
    return "specific-options"
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "configured-operation-options" {
        return &specificOptions.ConfiguredOperationOptions
    }
    if childYangName == "ondemand-operation-options" {
        return &specificOptions.OndemandOperationOptions
    }
    return nil
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["configured-operation-options"] = &specificOptions.ConfiguredOperationOptions
    children["ondemand-operation-options"] = &specificOptions.OndemandOperationOptions
    return children
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oper-type"] = specificOptions.OperType
    return leafs
}

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetBundleName() string { return "cisco_ios_xr" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetYangName() string { return "specific-options" }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) SetParent(parent types.Entity) { specificOptions.parent = parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetParent() types.Entity { return specificOptions.parent }

func (specificOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions) GetParentYangName() string { return "statistics-current" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions
// Parameters for a configured operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile used by the operation. The type is string.
    ProfileName interface{}
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetFilter() yfilter.YFilter { return configuredOperationOptions.YFilter }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) SetFilter(yf yfilter.YFilter) { configuredOperationOptions.YFilter = yf }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetSegmentPath() string {
    return "configured-operation-options"
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = configuredOperationOptions.ProfileName
    return leafs
}

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetYangName() string { return "configured-operation-options" }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) SetParent(parent types.Entity) { configuredOperationOptions.parent = parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetParent() types.Entity { return configuredOperationOptions.parent }

func (configuredOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_ConfiguredOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions
// Parameters for an ondemand operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ID of the ondemand operation. The type is interface{} with range:
    // 0..4294967295.
    OndemandOperationId interface{}

    // Total number of probes sent during the operation. The type is interface{}
    // with range: 0..255.
    ProbeCount interface{}
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetFilter() yfilter.YFilter { return ondemandOperationOptions.YFilter }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) SetFilter(yf yfilter.YFilter) { ondemandOperationOptions.YFilter = yf }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetGoName(yname string) string {
    if yname == "ondemand-operation-id" { return "OndemandOperationId" }
    if yname == "probe-count" { return "ProbeCount" }
    return ""
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetSegmentPath() string {
    return "ondemand-operation-options"
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ondemand-operation-id"] = ondemandOperationOptions.OndemandOperationId
    leafs["probe-count"] = ondemandOperationOptions.ProbeCount
    return leafs
}

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetBundleName() string { return "cisco_ios_xr" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetYangName() string { return "ondemand-operation-options" }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) SetParent(parent types.Entity) { ondemandOperationOptions.parent = parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetParent() types.Entity { return ondemandOperationOptions.parent }

func (ondemandOperationOptions *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_SpecificOptions_OndemandOperationOptions) GetParentYangName() string { return "specific-options" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule
// Operation schedule
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule struct {
    parent types.Entity
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

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetFilter() yfilter.YFilter { return operationSchedule.YFilter }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) SetFilter(yf yfilter.YFilter) { operationSchedule.YFilter = yf }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetGoName(yname string) string {
    if yname == "start-time" { return "StartTime" }
    if yname == "start-time-configured" { return "StartTimeConfigured" }
    if yname == "schedule-duration" { return "ScheduleDuration" }
    if yname == "schedule-interval" { return "ScheduleInterval" }
    return ""
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetSegmentPath() string {
    return "operation-schedule"
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-time"] = operationSchedule.StartTime
    leafs["start-time-configured"] = operationSchedule.StartTimeConfigured
    leafs["schedule-duration"] = operationSchedule.ScheduleDuration
    leafs["schedule-interval"] = operationSchedule.ScheduleInterval
    return leafs
}

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetBundleName() string { return "cisco_ios_xr" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetYangName() string { return "operation-schedule" }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) SetParent(parent types.Entity) { operationSchedule.parent = parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetParent() types.Entity { return operationSchedule.parent }

func (operationSchedule *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationSchedule) GetParentYangName() string { return "statistics-current" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric
// Metrics gathered for the operation
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of the metric.
    Config Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config

    // Buckets stored for the metric. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket.
    Bucket []Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetFilter() yfilter.YFilter { return operationMetric.YFilter }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) SetFilter(yf yfilter.YFilter) { operationMetric.YFilter = yf }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "bucket" { return "Bucket" }
    return ""
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetSegmentPath() string {
    return "operation-metric"
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &operationMetric.Config
    }
    if childYangName == "bucket" {
        for _, c := range operationMetric.Bucket {
            if operationMetric.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket{}
        operationMetric.Bucket = append(operationMetric.Bucket, child)
        return &operationMetric.Bucket[len(operationMetric.Bucket)-1]
    }
    return nil
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &operationMetric.Config
    for i := range operationMetric.Bucket {
        children[operationMetric.Bucket[i].GetSegmentPath()] = &operationMetric.Bucket[i]
    }
    return children
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetBundleName() string { return "cisco_ios_xr" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetYangName() string { return "operation-metric" }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) SetParent(parent types.Entity) { operationMetric.parent = parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetParent() types.Entity { return operationMetric.parent }

func (operationMetric *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric) GetParentYangName() string { return "statistics-current" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config
// Configuration of the metric
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config struct {
    parent types.Entity
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

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetGoName(yname string) string {
    if yname == "metric-type" { return "MetricType" }
    if yname == "bins-count" { return "BinsCount" }
    if yname == "bins-width" { return "BinsWidth" }
    if yname == "bucket-size" { return "BucketSize" }
    if yname == "bucket-size-unit" { return "BucketSizeUnit" }
    if yname == "buckets-archive" { return "BucketsArchive" }
    return ""
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetSegmentPath() string {
    return "config"
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric-type"] = config.MetricType
    leafs["bins-count"] = config.BinsCount
    leafs["bins-width"] = config.BinsWidth
    leafs["bucket-size"] = config.BucketSize
    leafs["bucket-size-unit"] = config.BucketSizeUnit
    leafs["buckets-archive"] = config.BucketsArchive
    return leafs
}

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetYangName() string { return "config" }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetParent() types.Entity { return config.parent }

func (config *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Config) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket
// Buckets stored for the metric
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket struct {
    parent types.Entity
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

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetFilter() yfilter.YFilter { return bucket.YFilter }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) SetFilter(yf yfilter.YFilter) { bucket.YFilter = yf }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetGoName(yname string) string {
    if yname == "start-at" { return "StartAt" }
    if yname == "duration" { return "Duration" }
    if yname == "sent" { return "Sent" }
    if yname == "lost" { return "Lost" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "duplicates" { return "Duplicates" }
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    if yname == "time-of-minimum" { return "TimeOfMinimum" }
    if yname == "time-of-maximum" { return "TimeOfMaximum" }
    if yname == "average" { return "Average" }
    if yname == "standard-deviation" { return "StandardDeviation" }
    if yname == "result-count" { return "ResultCount" }
    if yname == "data-sent-count" { return "DataSentCount" }
    if yname == "data-lost-count" { return "DataLostCount" }
    if yname == "overall-flr" { return "OverallFlr" }
    if yname == "suspect-start-mid-bucket" { return "SuspectStartMidBucket" }
    if yname == "suspect-schedule-latency" { return "SuspectScheduleLatency" }
    if yname == "suspect-send-fail" { return "SuspectSendFail" }
    if yname == "suspect-premature-end" { return "SuspectPrematureEnd" }
    if yname == "suspect-clock-drift" { return "SuspectClockDrift" }
    if yname == "suspect-memory-allocation-failed" { return "SuspectMemoryAllocationFailed" }
    if yname == "suspect-cleared-mid-bucket" { return "SuspectClearedMidBucket" }
    if yname == "suspect-probe-restarted" { return "SuspectProbeRestarted" }
    if yname == "suspect-management-latency" { return "SuspectManagementLatency" }
    if yname == "suspect-multiple-buckets" { return "SuspectMultipleBuckets" }
    if yname == "suspect-misordering" { return "SuspectMisordering" }
    if yname == "suspect-flr-low-packet-count" { return "SuspectFlrLowPacketCount" }
    if yname == "premature-reason" { return "PrematureReason" }
    if yname == "premature-reason-string" { return "PrematureReasonString" }
    if yname == "contents" { return "Contents" }
    return ""
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetSegmentPath() string {
    return "bucket"
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contents" {
        return &bucket.Contents
    }
    return nil
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["contents"] = &bucket.Contents
    return children
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-at"] = bucket.StartAt
    leafs["duration"] = bucket.Duration
    leafs["sent"] = bucket.Sent
    leafs["lost"] = bucket.Lost
    leafs["corrupt"] = bucket.Corrupt
    leafs["out-of-order"] = bucket.OutOfOrder
    leafs["duplicates"] = bucket.Duplicates
    leafs["minimum"] = bucket.Minimum
    leafs["maximum"] = bucket.Maximum
    leafs["time-of-minimum"] = bucket.TimeOfMinimum
    leafs["time-of-maximum"] = bucket.TimeOfMaximum
    leafs["average"] = bucket.Average
    leafs["standard-deviation"] = bucket.StandardDeviation
    leafs["result-count"] = bucket.ResultCount
    leafs["data-sent-count"] = bucket.DataSentCount
    leafs["data-lost-count"] = bucket.DataLostCount
    leafs["overall-flr"] = bucket.OverallFlr
    leafs["suspect-start-mid-bucket"] = bucket.SuspectStartMidBucket
    leafs["suspect-schedule-latency"] = bucket.SuspectScheduleLatency
    leafs["suspect-send-fail"] = bucket.SuspectSendFail
    leafs["suspect-premature-end"] = bucket.SuspectPrematureEnd
    leafs["suspect-clock-drift"] = bucket.SuspectClockDrift
    leafs["suspect-memory-allocation-failed"] = bucket.SuspectMemoryAllocationFailed
    leafs["suspect-cleared-mid-bucket"] = bucket.SuspectClearedMidBucket
    leafs["suspect-probe-restarted"] = bucket.SuspectProbeRestarted
    leafs["suspect-management-latency"] = bucket.SuspectManagementLatency
    leafs["suspect-multiple-buckets"] = bucket.SuspectMultipleBuckets
    leafs["suspect-misordering"] = bucket.SuspectMisordering
    leafs["suspect-flr-low-packet-count"] = bucket.SuspectFlrLowPacketCount
    leafs["premature-reason"] = bucket.PrematureReason
    leafs["premature-reason-string"] = bucket.PrematureReasonString
    return leafs
}

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetBundleName() string { return "cisco_ios_xr" }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetYangName() string { return "bucket" }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) SetParent(parent types.Entity) { bucket.parent = parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetParent() types.Entity { return bucket.parent }

func (bucket *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket) GetParentYangName() string { return "operation-metric" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents
// The contents of the bucket; bins or samples
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BucketType. The type is SlaOperBucket.
    BucketType interface{}

    // Result bins in an SLA metric bucket.
    Aggregated Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated

    // Result samples in an SLA metric bucket.
    Unaggregated Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetFilter() yfilter.YFilter { return contents.YFilter }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) SetFilter(yf yfilter.YFilter) { contents.YFilter = yf }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetGoName(yname string) string {
    if yname == "bucket-type" { return "BucketType" }
    if yname == "aggregated" { return "Aggregated" }
    if yname == "unaggregated" { return "Unaggregated" }
    return ""
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetSegmentPath() string {
    return "contents"
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggregated" {
        return &contents.Aggregated
    }
    if childYangName == "unaggregated" {
        return &contents.Unaggregated
    }
    return nil
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggregated"] = &contents.Aggregated
    children["unaggregated"] = &contents.Unaggregated
    return children
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bucket-type"] = contents.BucketType
    return leafs
}

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetBundleName() string { return "cisco_ios_xr" }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetYangName() string { return "contents" }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) SetParent(parent types.Entity) { contents.parent = parent }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetParent() types.Entity { return contents.parent }

func (contents *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents) GetParentYangName() string { return "bucket" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated
// Result bins in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The bins of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins.
    Bins []Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetFilter() yfilter.YFilter { return aggregated.YFilter }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) SetFilter(yf yfilter.YFilter) { aggregated.YFilter = yf }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetGoName(yname string) string {
    if yname == "bins" { return "Bins" }
    return ""
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetSegmentPath() string {
    return "aggregated"
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bins" {
        for _, c := range aggregated.Bins {
            if aggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins{}
        aggregated.Bins = append(aggregated.Bins, child)
        return &aggregated.Bins[len(aggregated.Bins)-1]
    }
    return nil
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aggregated.Bins {
        children[aggregated.Bins[i].GetSegmentPath()] = &aggregated.Bins[i]
    }
    return children
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetBundleName() string { return "cisco_ios_xr" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetYangName() string { return "aggregated" }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) SetParent(parent types.Entity) { aggregated.parent = parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetParent() types.Entity { return aggregated.parent }

func (aggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins
// The bins of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins struct {
    parent types.Entity
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

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetFilter() yfilter.YFilter { return bins.YFilter }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) SetFilter(yf yfilter.YFilter) { bins.YFilter = yf }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "lower-bound-tenths" { return "LowerBoundTenths" }
    if yname == "upper-bound-tenths" { return "UpperBoundTenths" }
    if yname == "sum" { return "Sum" }
    if yname == "count" { return "Count" }
    return ""
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetSegmentPath() string {
    return "bins"
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = bins.LowerBound
    leafs["upper-bound"] = bins.UpperBound
    leafs["lower-bound-tenths"] = bins.LowerBoundTenths
    leafs["upper-bound-tenths"] = bins.UpperBoundTenths
    leafs["sum"] = bins.Sum
    leafs["count"] = bins.Count
    return leafs
}

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleName() string { return "cisco_ios_xr" }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetYangName() string { return "bins" }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) SetParent(parent types.Entity) { bins.parent = parent }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParent() types.Entity { return bins.parent }

func (bins *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Aggregated_Bins) GetParentYangName() string { return "aggregated" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated
// Result samples in an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The samples of an SLA metric bucket. The type is slice of
    // Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample.
    Sample []Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetFilter() yfilter.YFilter { return unaggregated.YFilter }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) SetFilter(yf yfilter.YFilter) { unaggregated.YFilter = yf }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetGoName(yname string) string {
    if yname == "sample" { return "Sample" }
    return ""
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetSegmentPath() string {
    return "unaggregated"
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sample" {
        for _, c := range unaggregated.Sample {
            if unaggregated.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample{}
        unaggregated.Sample = append(unaggregated.Sample, child)
        return &unaggregated.Sample[len(unaggregated.Sample)-1]
    }
    return nil
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range unaggregated.Sample {
        children[unaggregated.Sample[i].GetSegmentPath()] = &unaggregated.Sample[i]
    }
    return children
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetBundleName() string { return "cisco_ios_xr" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetYangName() string { return "unaggregated" }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) SetParent(parent types.Entity) { unaggregated.parent = parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetParent() types.Entity { return unaggregated.parent }

func (unaggregated *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated) GetParentYangName() string { return "contents" }

// Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample
// The samples of an SLA metric bucket
type Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample struct {
    parent types.Entity
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

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetFilter() yfilter.YFilter { return sample.YFilter }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetFilter(yf yfilter.YFilter) { sample.YFilter = yf }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetGoName(yname string) string {
    if yname == "sent-at" { return "SentAt" }
    if yname == "sent" { return "Sent" }
    if yname == "timed-out" { return "TimedOut" }
    if yname == "corrupt" { return "Corrupt" }
    if yname == "out-of-order" { return "OutOfOrder" }
    if yname == "no-data-packets" { return "NoDataPackets" }
    if yname == "result" { return "Result" }
    if yname == "frames-sent" { return "FramesSent" }
    if yname == "frames-lost" { return "FramesLost" }
    return ""
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetSegmentPath() string {
    return "sample"
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sent-at"] = sample.SentAt
    leafs["sent"] = sample.Sent
    leafs["timed-out"] = sample.TimedOut
    leafs["corrupt"] = sample.Corrupt
    leafs["out-of-order"] = sample.OutOfOrder
    leafs["no-data-packets"] = sample.NoDataPackets
    leafs["result"] = sample.Result
    leafs["frames-sent"] = sample.FramesSent
    leafs["frames-lost"] = sample.FramesLost
    return leafs
}

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleName() string { return "cisco_ios_xr" }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetYangName() string { return "sample" }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) SetParent(parent types.Entity) { sample.parent = parent }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParent() types.Entity { return sample.parent }

func (sample *Sla_Protocols_Ethernet_StatisticsCurrents_StatisticsCurrent_OperationMetric_Bucket_Contents_Unaggregated_Sample) GetParentYangName() string { return "unaggregated" }

// SlaNodes
// sla nodes
type SlaNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (slaNodes *SlaNodes) GetFilter() yfilter.YFilter { return slaNodes.YFilter }

func (slaNodes *SlaNodes) SetFilter(yf yfilter.YFilter) { slaNodes.YFilter = yf }

func (slaNodes *SlaNodes) GetGoName(yname string) string {
    return ""
}

func (slaNodes *SlaNodes) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-sla-oper:sla-nodes"
}

func (slaNodes *SlaNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (slaNodes *SlaNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (slaNodes *SlaNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaNodes *SlaNodes) GetBundleName() string { return "cisco_ios_xr" }

func (slaNodes *SlaNodes) GetYangName() string { return "sla-nodes" }

func (slaNodes *SlaNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaNodes *SlaNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaNodes *SlaNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaNodes *SlaNodes) SetParent(parent types.Entity) { slaNodes.parent = parent }

func (slaNodes *SlaNodes) GetParent() types.Entity { return slaNodes.parent }

func (slaNodes *SlaNodes) GetParentYangName() string { return "Cisco-IOS-XR-infra-sla-oper" }

