// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp-headless package operational data.
// 
// This module contains definitions
// for the following management objects:
//   headless-func-data: Information related to headless
//     functionality
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_headless_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_headless_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-headless-oper headless-func-data}", reflect.TypeOf(HeadlessFuncData{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data", reflect.TypeOf(HeadlessFuncData{}))
}

// MxpOtnPrbsStatus represents Mxp otn prbs status
type MxpOtnPrbsStatus string

const (
    // Locked
    MxpOtnPrbsStatus_locked MxpOtnPrbsStatus = "locked"

    // Not Locked
    MxpOtnPrbsStatus_not_locked MxpOtnPrbsStatus = "not-locked"

    // Not Applicable
    MxpOtnPrbsStatus_not_applicable MxpOtnPrbsStatus = "not-applicable"

    // Locked With Errors
    MxpOtnPrbsStatus_locked_with_errors MxpOtnPrbsStatus = "locked-with-errors"
)

// HeadlessFuncData
// Information related to headless functionality
type HeadlessFuncData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OTN Statistics collected during last headless operation.
    OtnPortNames HeadlessFuncData_OtnPortNames

    // Ethernet Statistics collected during last headless operation.
    EthernetPortNames HeadlessFuncData_EthernetPortNames
}

func (headlessFuncData *HeadlessFuncData) GetFilter() yfilter.YFilter { return headlessFuncData.YFilter }

func (headlessFuncData *HeadlessFuncData) SetFilter(yf yfilter.YFilter) { headlessFuncData.YFilter = yf }

func (headlessFuncData *HeadlessFuncData) GetGoName(yname string) string {
    if yname == "otn-port-names" { return "OtnPortNames" }
    if yname == "ethernet-port-names" { return "EthernetPortNames" }
    return ""
}

func (headlessFuncData *HeadlessFuncData) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data"
}

func (headlessFuncData *HeadlessFuncData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn-port-names" {
        return &headlessFuncData.OtnPortNames
    }
    if childYangName == "ethernet-port-names" {
        return &headlessFuncData.EthernetPortNames
    }
    return nil
}

func (headlessFuncData *HeadlessFuncData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn-port-names"] = &headlessFuncData.OtnPortNames
    children["ethernet-port-names"] = &headlessFuncData.EthernetPortNames
    return children
}

func (headlessFuncData *HeadlessFuncData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (headlessFuncData *HeadlessFuncData) GetBundleName() string { return "cisco_ios_xr" }

func (headlessFuncData *HeadlessFuncData) GetYangName() string { return "headless-func-data" }

func (headlessFuncData *HeadlessFuncData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (headlessFuncData *HeadlessFuncData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (headlessFuncData *HeadlessFuncData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (headlessFuncData *HeadlessFuncData) SetParent(parent types.Entity) { headlessFuncData.parent = parent }

func (headlessFuncData *HeadlessFuncData) GetParent() types.Entity { return headlessFuncData.parent }

func (headlessFuncData *HeadlessFuncData) GetParentYangName() string { return "Cisco-IOS-XR-ncs1k-mxp-headless-oper" }

// HeadlessFuncData_OtnPortNames
// OTN Statistics collected during last headless
// operation
type HeadlessFuncData_OtnPortNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // port Name. The type is slice of HeadlessFuncData_OtnPortNames_OtnPortName.
    OtnPortName []HeadlessFuncData_OtnPortNames_OtnPortName
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetFilter() yfilter.YFilter { return otnPortNames.YFilter }

func (otnPortNames *HeadlessFuncData_OtnPortNames) SetFilter(yf yfilter.YFilter) { otnPortNames.YFilter = yf }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetGoName(yname string) string {
    if yname == "otn-port-name" { return "OtnPortName" }
    return ""
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetSegmentPath() string {
    return "otn-port-names"
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn-port-name" {
        for _, c := range otnPortNames.OtnPortName {
            if otnPortNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HeadlessFuncData_OtnPortNames_OtnPortName{}
        otnPortNames.OtnPortName = append(otnPortNames.OtnPortName, child)
        return &otnPortNames.OtnPortName[len(otnPortNames.OtnPortName)-1]
    }
    return nil
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range otnPortNames.OtnPortName {
        children[otnPortNames.OtnPortName[i].GetSegmentPath()] = &otnPortNames.OtnPortName[i]
    }
    return children
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetBundleName() string { return "cisco_ios_xr" }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetYangName() string { return "otn-port-names" }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otnPortNames *HeadlessFuncData_OtnPortNames) SetParent(parent types.Entity) { otnPortNames.parent = parent }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetParent() types.Entity { return otnPortNames.parent }

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetParentYangName() string { return "headless-func-data" }

// HeadlessFuncData_OtnPortNames_OtnPortName
// port Name
type HeadlessFuncData_OtnPortNames_OtnPortName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Started Stateful. The type is bool.
    StartedStateful interface{}

    // Headless Start Time. The type is string with length: 0..64.
    HeadlessStartTime interface{}

    // Headless End Time. The type is string with length: 0..64.
    HeadlessEndTime interface{}

    // OTN statistics.
    OtnStatistics HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics

    // PRBS Statistics.
    PrbsStatistics HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetFilter() yfilter.YFilter { return otnPortName.YFilter }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) SetFilter(yf yfilter.YFilter) { otnPortName.YFilter = yf }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "started-stateful" { return "StartedStateful" }
    if yname == "headless-start-time" { return "HeadlessStartTime" }
    if yname == "headless-end-time" { return "HeadlessEndTime" }
    if yname == "otn-statistics" { return "OtnStatistics" }
    if yname == "prbs-statistics" { return "PrbsStatistics" }
    return ""
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetSegmentPath() string {
    return "otn-port-name" + "[name='" + fmt.Sprintf("%v", otnPortName.Name) + "']"
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "otn-statistics" {
        return &otnPortName.OtnStatistics
    }
    if childYangName == "prbs-statistics" {
        return &otnPortName.PrbsStatistics
    }
    return nil
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["otn-statistics"] = &otnPortName.OtnStatistics
    children["prbs-statistics"] = &otnPortName.PrbsStatistics
    return children
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = otnPortName.Name
    leafs["started-stateful"] = otnPortName.StartedStateful
    leafs["headless-start-time"] = otnPortName.HeadlessStartTime
    leafs["headless-end-time"] = otnPortName.HeadlessEndTime
    return leafs
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetBundleName() string { return "cisco_ios_xr" }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetYangName() string { return "otn-port-name" }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) SetParent(parent types.Entity) { otnPortName.parent = parent }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetParent() types.Entity { return otnPortName.parent }

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetParentYangName() string { return "otn-port-names" }

// HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics
// OTN statistics
type HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SmBip. The type is interface{} with range: 0..18446744073709551615.
    SmBip interface{}

    // SmBei. The type is interface{} with range: 0..18446744073709551615.
    SmBei interface{}

    // FecEc. The type is interface{} with range: 0..18446744073709551615.
    FecEc interface{}

    // FecUc. The type is interface{} with range: 0..18446744073709551615.
    FecUc interface{}
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetFilter() yfilter.YFilter { return otnStatistics.YFilter }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) SetFilter(yf yfilter.YFilter) { otnStatistics.YFilter = yf }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetGoName(yname string) string {
    if yname == "sm-bip" { return "SmBip" }
    if yname == "sm-bei" { return "SmBei" }
    if yname == "fec-ec" { return "FecEc" }
    if yname == "fec-uc" { return "FecUc" }
    return ""
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetSegmentPath() string {
    return "otn-statistics"
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sm-bip"] = otnStatistics.SmBip
    leafs["sm-bei"] = otnStatistics.SmBei
    leafs["fec-ec"] = otnStatistics.FecEc
    leafs["fec-uc"] = otnStatistics.FecUc
    return leafs
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetYangName() string { return "otn-statistics" }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) SetParent(parent types.Entity) { otnStatistics.parent = parent }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetParent() types.Entity { return otnStatistics.parent }

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetParentYangName() string { return "otn-port-name" }

// HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics
// PRBS Statistics
type HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // EBC. The type is interface{} with range: 0..18446744073709551615.
    Ebc interface{}

    // SyncStatus. The type is MxpOtnPrbsStatus.
    SyncStatus interface{}
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetFilter() yfilter.YFilter { return prbsStatistics.YFilter }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) SetFilter(yf yfilter.YFilter) { prbsStatistics.YFilter = yf }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetGoName(yname string) string {
    if yname == "ebc" { return "Ebc" }
    if yname == "sync-status" { return "SyncStatus" }
    return ""
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetSegmentPath() string {
    return "prbs-statistics"
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ebc"] = prbsStatistics.Ebc
    leafs["sync-status"] = prbsStatistics.SyncStatus
    return leafs
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetYangName() string { return "prbs-statistics" }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) SetParent(parent types.Entity) { prbsStatistics.parent = parent }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetParent() types.Entity { return prbsStatistics.parent }

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetParentYangName() string { return "otn-port-name" }

// HeadlessFuncData_EthernetPortNames
// Ethernet Statistics collected during last
// headless operation
type HeadlessFuncData_EthernetPortNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port Name. The type is slice of
    // HeadlessFuncData_EthernetPortNames_EthernetPortName.
    EthernetPortName []HeadlessFuncData_EthernetPortNames_EthernetPortName
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetFilter() yfilter.YFilter { return ethernetPortNames.YFilter }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) SetFilter(yf yfilter.YFilter) { ethernetPortNames.YFilter = yf }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetGoName(yname string) string {
    if yname == "ethernet-port-name" { return "EthernetPortName" }
    return ""
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetSegmentPath() string {
    return "ethernet-port-names"
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ethernet-port-name" {
        for _, c := range ethernetPortNames.EthernetPortName {
            if ethernetPortNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HeadlessFuncData_EthernetPortNames_EthernetPortName{}
        ethernetPortNames.EthernetPortName = append(ethernetPortNames.EthernetPortName, child)
        return &ethernetPortNames.EthernetPortName[len(ethernetPortNames.EthernetPortName)-1]
    }
    return nil
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ethernetPortNames.EthernetPortName {
        children[ethernetPortNames.EthernetPortName[i].GetSegmentPath()] = &ethernetPortNames.EthernetPortName[i]
    }
    return children
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetYangName() string { return "ethernet-port-names" }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) SetParent(parent types.Entity) { ethernetPortNames.parent = parent }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetParent() types.Entity { return ethernetPortNames.parent }

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetParentYangName() string { return "headless-func-data" }

// HeadlessFuncData_EthernetPortNames_EthernetPortName
// Port Name
type HeadlessFuncData_EthernetPortNames_EthernetPortName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Started Stateful. The type is bool.
    StartedStateful interface{}

    // Headless Start Time. The type is string with length: 0..64.
    HeadlessStartTime interface{}

    // Headless End Time. The type is string with length: 0..64.
    HeadlessEndTime interface{}

    // Ether Statistics.
    EtherStatistics HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetFilter() yfilter.YFilter { return ethernetPortName.YFilter }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) SetFilter(yf yfilter.YFilter) { ethernetPortName.YFilter = yf }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "started-stateful" { return "StartedStateful" }
    if yname == "headless-start-time" { return "HeadlessStartTime" }
    if yname == "headless-end-time" { return "HeadlessEndTime" }
    if yname == "ether-statistics" { return "EtherStatistics" }
    return ""
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetSegmentPath() string {
    return "ethernet-port-name" + "[name='" + fmt.Sprintf("%v", ethernetPortName.Name) + "']"
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ether-statistics" {
        return &ethernetPortName.EtherStatistics
    }
    return nil
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ether-statistics"] = &ethernetPortName.EtherStatistics
    return children
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ethernetPortName.Name
    leafs["started-stateful"] = ethernetPortName.StartedStateful
    leafs["headless-start-time"] = ethernetPortName.HeadlessStartTime
    leafs["headless-end-time"] = ethernetPortName.HeadlessEndTime
    return leafs
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetYangName() string { return "ethernet-port-name" }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) SetParent(parent types.Entity) { ethernetPortName.parent = parent }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetParent() types.Entity { return ethernetPortName.parent }

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetParentYangName() string { return "ethernet-port-names" }

// HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics
// Ether Statistics
type HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RxPktsOverSized. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsOverSized interface{}

    // RxPktsBadFcs. The type is interface{} with range: 0..18446744073709551615.
    RxPktsBadFcs interface{}

    // RxErrorJabbers. The type is interface{} with range:
    // 0..18446744073709551615.
    RxErrorJabbers interface{}

    // RxPktsMulticast. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsMulticast interface{}

    // RxPktsBroadcast. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsBroadcast interface{}

    // RxPktsUnderSized. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsUnderSized interface{}

    // RxPackets. The type is interface{} with range: 0..18446744073709551615.
    RxPackets interface{}

    // RxTotalBytes. The type is interface{} with range: 0..18446744073709551615.
    RxTotalBytes interface{}

    // RxBytesGood. The type is interface{} with range: 0..18446744073709551615.
    RxBytesGood interface{}

    // RxPktsGood. The type is interface{} with range: 0..18446744073709551615.
    RxPktsGood interface{}

    // TxBytesGood. The type is interface{} with range: 0..18446744073709551615.
    TxBytesGood interface{}

    // TxPktsGood. The type is interface{} with range: 0..18446744073709551615.
    TxPktsGood interface{}

    // RxRecvFragments. The type is interface{} with range:
    // 0..18446744073709551615.
    RxRecvFragments interface{}

    // RxPkts64Bytes. The type is interface{} with range: 0..18446744073709551615.
    RxPkts64Bytes interface{}

    // RxPkts65To127Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts65To127Bytes interface{}

    // RxPkts128to255Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts128To255Bytes interface{}

    // RxPkts256To511Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts256To511Bytes interface{}

    // RxPkts512To1023Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts512To1023Bytes interface{}

    // RxPkts1024To1518Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts1024To1518Bytes interface{}

    // RxPktsUnicast. The type is interface{} with range: 0..18446744073709551615.
    RxPktsUnicast interface{}

    // TxPackets. The type is interface{} with range: 0..18446744073709551615.
    TxPackets interface{}

    // TxTotalBytes. The type is interface{} with range: 0..18446744073709551615.
    TxTotalBytes interface{}

    // TxPktsUnderSized. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPktsUnderSized interface{}

    // TxPktsOverSized. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPktsOverSized interface{}

    // TxFragments. The type is interface{} with range: 0..18446744073709551615.
    TxFragments interface{}

    // TxJabber. The type is interface{} with range: 0..18446744073709551615.
    TxJabber interface{}

    // TxBadFCS. The type is interface{} with range: 0..18446744073709551615.
    TxBadFcs interface{}

    // RxPktDrop. The type is interface{} with range: 0..18446744073709551615.
    RxPktDrop interface{}

    // RxPause. The type is interface{} with range: 0..18446744073709551615.
    RxPause interface{}

    // TxPause. The type is interface{} with range: 0..18446744073709551615.
    TxPause interface{}

    // RxLldpPkt. The type is interface{} with range: 0..18446744073709551615.
    RxLldpPkt interface{}

    // Rx8021QPkt. The type is interface{} with range: 0..18446744073709551615.
    Rx8021QPkt interface{}
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetFilter() yfilter.YFilter { return etherStatistics.YFilter }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) SetFilter(yf yfilter.YFilter) { etherStatistics.YFilter = yf }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetGoName(yname string) string {
    if yname == "rx-pkts-over-sized" { return "RxPktsOverSized" }
    if yname == "rx-pkts-bad-fcs" { return "RxPktsBadFcs" }
    if yname == "rx-error-jabbers" { return "RxErrorJabbers" }
    if yname == "rx-pkts-multicast" { return "RxPktsMulticast" }
    if yname == "rx-pkts-broadcast" { return "RxPktsBroadcast" }
    if yname == "rx-pkts-under-sized" { return "RxPktsUnderSized" }
    if yname == "rx-packets" { return "RxPackets" }
    if yname == "rx-total-bytes" { return "RxTotalBytes" }
    if yname == "rx-bytes-good" { return "RxBytesGood" }
    if yname == "rx-pkts-good" { return "RxPktsGood" }
    if yname == "tx-bytes-good" { return "TxBytesGood" }
    if yname == "tx-pkts-good" { return "TxPktsGood" }
    if yname == "rx-recv-fragments" { return "RxRecvFragments" }
    if yname == "rx-pkts64-bytes" { return "RxPkts64Bytes" }
    if yname == "rx-pkts65-to127-bytes" { return "RxPkts65To127Bytes" }
    if yname == "rx-pkts128to255-bytes" { return "RxPkts128To255Bytes" }
    if yname == "rx-pkts256-to511-bytes" { return "RxPkts256To511Bytes" }
    if yname == "rx-pkts512-to1023-bytes" { return "RxPkts512To1023Bytes" }
    if yname == "rx-pkts1024-to1518-bytes" { return "RxPkts1024To1518Bytes" }
    if yname == "rx-pkts-unicast" { return "RxPktsUnicast" }
    if yname == "tx-packets" { return "TxPackets" }
    if yname == "tx-total-bytes" { return "TxTotalBytes" }
    if yname == "tx-pkts-under-sized" { return "TxPktsUnderSized" }
    if yname == "tx-pkts-over-sized" { return "TxPktsOverSized" }
    if yname == "tx-fragments" { return "TxFragments" }
    if yname == "tx-jabber" { return "TxJabber" }
    if yname == "tx-bad-fcs" { return "TxBadFcs" }
    if yname == "rx-pkt-drop" { return "RxPktDrop" }
    if yname == "rx-pause" { return "RxPause" }
    if yname == "tx-pause" { return "TxPause" }
    if yname == "rx-lldp-pkt" { return "RxLldpPkt" }
    if yname == "rx8021q-pkt" { return "Rx8021QPkt" }
    return ""
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetSegmentPath() string {
    return "ether-statistics"
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-pkts-over-sized"] = etherStatistics.RxPktsOverSized
    leafs["rx-pkts-bad-fcs"] = etherStatistics.RxPktsBadFcs
    leafs["rx-error-jabbers"] = etherStatistics.RxErrorJabbers
    leafs["rx-pkts-multicast"] = etherStatistics.RxPktsMulticast
    leafs["rx-pkts-broadcast"] = etherStatistics.RxPktsBroadcast
    leafs["rx-pkts-under-sized"] = etherStatistics.RxPktsUnderSized
    leafs["rx-packets"] = etherStatistics.RxPackets
    leafs["rx-total-bytes"] = etherStatistics.RxTotalBytes
    leafs["rx-bytes-good"] = etherStatistics.RxBytesGood
    leafs["rx-pkts-good"] = etherStatistics.RxPktsGood
    leafs["tx-bytes-good"] = etherStatistics.TxBytesGood
    leafs["tx-pkts-good"] = etherStatistics.TxPktsGood
    leafs["rx-recv-fragments"] = etherStatistics.RxRecvFragments
    leafs["rx-pkts64-bytes"] = etherStatistics.RxPkts64Bytes
    leafs["rx-pkts65-to127-bytes"] = etherStatistics.RxPkts65To127Bytes
    leafs["rx-pkts128to255-bytes"] = etherStatistics.RxPkts128To255Bytes
    leafs["rx-pkts256-to511-bytes"] = etherStatistics.RxPkts256To511Bytes
    leafs["rx-pkts512-to1023-bytes"] = etherStatistics.RxPkts512To1023Bytes
    leafs["rx-pkts1024-to1518-bytes"] = etherStatistics.RxPkts1024To1518Bytes
    leafs["rx-pkts-unicast"] = etherStatistics.RxPktsUnicast
    leafs["tx-packets"] = etherStatistics.TxPackets
    leafs["tx-total-bytes"] = etherStatistics.TxTotalBytes
    leafs["tx-pkts-under-sized"] = etherStatistics.TxPktsUnderSized
    leafs["tx-pkts-over-sized"] = etherStatistics.TxPktsOverSized
    leafs["tx-fragments"] = etherStatistics.TxFragments
    leafs["tx-jabber"] = etherStatistics.TxJabber
    leafs["tx-bad-fcs"] = etherStatistics.TxBadFcs
    leafs["rx-pkt-drop"] = etherStatistics.RxPktDrop
    leafs["rx-pause"] = etherStatistics.RxPause
    leafs["tx-pause"] = etherStatistics.TxPause
    leafs["rx-lldp-pkt"] = etherStatistics.RxLldpPkt
    leafs["rx8021q-pkt"] = etherStatistics.Rx8021QPkt
    return leafs
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetYangName() string { return "ether-statistics" }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) SetParent(parent types.Entity) { etherStatistics.parent = parent }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetParent() types.Entity { return etherStatistics.parent }

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetParentYangName() string { return "ethernet-port-name" }

