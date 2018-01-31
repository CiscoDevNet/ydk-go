// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hw-module: mxp hw-module command chain
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-oper hw-module}", reflect.TypeOf(HwModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-oper:hw-module", reflect.TypeOf(HwModule{}))
}

// ClientDataRate represents Client data rate
type ClientDataRate string

const (
    // Ten Gig
    ClientDataRate_ten_gig ClientDataRate = "ten-gig"

    // Forty Gig
    ClientDataRate_forty_gig ClientDataRate = "forty-gig"

    // Hundread Gig
    ClientDataRate_hundred_gig ClientDataRate = "hundred-gig"

    // Ten and Hundread Gig
    ClientDataRate_ten_and_hundred_gig ClientDataRate = "ten-and-hundred-gig"
)

// TrunkDataRate represents Trunk data rate
type TrunkDataRate string

const (
    // FiftyGig
    TrunkDataRate_fifty_gig TrunkDataRate = "fifty-gig"

    // HundredGig
    TrunkDataRate_hundred_gig TrunkDataRate = "hundred-gig"

    // TwoHundredGig
    TrunkDataRate_two_hundred_gig TrunkDataRate = "two-hundred-gig"

    // TwoHundredFiftyGig
    TrunkDataRate_two_hundred_fifty_gig TrunkDataRate = "two-hundred-fifty-gig"
)

// HwModuleSliceStatus represents Hw module slice status
type HwModuleSliceStatus string

const (
    // Not Provisioned
    HwModuleSliceStatus_not_provisioned HwModuleSliceStatus = "not-provisioned"

    // Provisioning In-Progress
    HwModuleSliceStatus_provisioning_in_progress HwModuleSliceStatus = "provisioning-in-progress"

    // Provisioned
    HwModuleSliceStatus_provisioned HwModuleSliceStatus = "provisioned"

    // Provisioning Failed
    HwModuleSliceStatus_provisioning_failed HwModuleSliceStatus = "provisioning-failed"

    // Provisioning Scheduled
    HwModuleSliceStatus_provisioning_scheduled HwModuleSliceStatus = "provisioning-scheduled"

    // Reprovisioning Aborted
    HwModuleSliceStatus_reprovisioning_aborted HwModuleSliceStatus = "reprovisioning-aborted"
)

// HwModule
// mxp hw-module command chain
type HwModule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slice information.
    SliceIds HwModule_SliceIds

    // Information for all slices.
    SliceAll HwModule_SliceAll
}

func (hwModule *HwModule) GetFilter() yfilter.YFilter { return hwModule.YFilter }

func (hwModule *HwModule) SetFilter(yf yfilter.YFilter) { hwModule.YFilter = yf }

func (hwModule *HwModule) GetGoName(yname string) string {
    if yname == "slice-ids" { return "SliceIds" }
    if yname == "slice-all" { return "SliceAll" }
    return ""
}

func (hwModule *HwModule) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module"
}

func (hwModule *HwModule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-ids" {
        return &hwModule.SliceIds
    }
    if childYangName == "slice-all" {
        return &hwModule.SliceAll
    }
    return nil
}

func (hwModule *HwModule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slice-ids"] = &hwModule.SliceIds
    children["slice-all"] = &hwModule.SliceAll
    return children
}

func (hwModule *HwModule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwModule *HwModule) GetBundleName() string { return "cisco_ios_xr" }

func (hwModule *HwModule) GetYangName() string { return "hw-module" }

func (hwModule *HwModule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwModule *HwModule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwModule *HwModule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwModule *HwModule) SetParent(parent types.Entity) { hwModule.parent = parent }

func (hwModule *HwModule) GetParent() types.Entity { return hwModule.parent }

func (hwModule *HwModule) GetParentYangName() string { return "Cisco-IOS-XR-ncs1k-mxp-oper" }

// HwModule_SliceIds
// Slice information
type HwModule_SliceIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per slice num data. The type is slice of HwModule_SliceIds_SliceId.
    SliceId []HwModule_SliceIds_SliceId
}

func (sliceIds *HwModule_SliceIds) GetFilter() yfilter.YFilter { return sliceIds.YFilter }

func (sliceIds *HwModule_SliceIds) SetFilter(yf yfilter.YFilter) { sliceIds.YFilter = yf }

func (sliceIds *HwModule_SliceIds) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    return ""
}

func (sliceIds *HwModule_SliceIds) GetSegmentPath() string {
    return "slice-ids"
}

func (sliceIds *HwModule_SliceIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-id" {
        for _, c := range sliceIds.SliceId {
            if sliceIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceIds_SliceId{}
        sliceIds.SliceId = append(sliceIds.SliceId, child)
        return &sliceIds.SliceId[len(sliceIds.SliceId)-1]
    }
    return nil
}

func (sliceIds *HwModule_SliceIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceIds.SliceId {
        children[sliceIds.SliceId[i].GetSegmentPath()] = &sliceIds.SliceId[i]
    }
    return children
}

func (sliceIds *HwModule_SliceIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sliceIds *HwModule_SliceIds) GetBundleName() string { return "cisco_ios_xr" }

func (sliceIds *HwModule_SliceIds) GetYangName() string { return "slice-ids" }

func (sliceIds *HwModule_SliceIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceIds *HwModule_SliceIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceIds *HwModule_SliceIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceIds *HwModule_SliceIds) SetParent(parent types.Entity) { sliceIds.parent = parent }

func (sliceIds *HwModule_SliceIds) GetParent() types.Entity { return sliceIds.parent }

func (sliceIds *HwModule_SliceIds) GetParentYangName() string { return "hw-module" }

// HwModule_SliceIds_SliceId
// Per slice num data
type HwModule_SliceIds_SliceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Details associated with a particular slice number.
    // The type is interface{} with range: -2147483648..2147483647.
    SliceNum interface{}

    // slice info. The type is slice of HwModule_SliceIds_SliceId_SliceInfo.
    SliceInfo []HwModule_SliceIds_SliceId_SliceInfo
}

func (sliceId *HwModule_SliceIds_SliceId) GetFilter() yfilter.YFilter { return sliceId.YFilter }

func (sliceId *HwModule_SliceIds_SliceId) SetFilter(yf yfilter.YFilter) { sliceId.YFilter = yf }

func (sliceId *HwModule_SliceIds_SliceId) GetGoName(yname string) string {
    if yname == "slice-num" { return "SliceNum" }
    if yname == "slice-info" { return "SliceInfo" }
    return ""
}

func (sliceId *HwModule_SliceIds_SliceId) GetSegmentPath() string {
    return "slice-id" + "[slice-num='" + fmt.Sprintf("%v", sliceId.SliceNum) + "']"
}

func (sliceId *HwModule_SliceIds_SliceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-info" {
        for _, c := range sliceId.SliceInfo {
            if sliceId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceIds_SliceId_SliceInfo{}
        sliceId.SliceInfo = append(sliceId.SliceInfo, child)
        return &sliceId.SliceInfo[len(sliceId.SliceInfo)-1]
    }
    return nil
}

func (sliceId *HwModule_SliceIds_SliceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceId.SliceInfo {
        children[sliceId.SliceInfo[i].GetSegmentPath()] = &sliceId.SliceInfo[i]
    }
    return children
}

func (sliceId *HwModule_SliceIds_SliceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-num"] = sliceId.SliceNum
    return leafs
}

func (sliceId *HwModule_SliceIds_SliceId) GetBundleName() string { return "cisco_ios_xr" }

func (sliceId *HwModule_SliceIds_SliceId) GetYangName() string { return "slice-id" }

func (sliceId *HwModule_SliceIds_SliceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceId *HwModule_SliceIds_SliceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceId *HwModule_SliceIds_SliceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceId *HwModule_SliceIds_SliceId) SetParent(parent types.Entity) { sliceId.parent = parent }

func (sliceId *HwModule_SliceIds_SliceId) GetParent() types.Entity { return sliceId.parent }

func (sliceId *HwModule_SliceIds_SliceId) GetParentYangName() string { return "slice-ids" }

// HwModule_SliceIds_SliceId_SliceInfo
// slice info
type HwModule_SliceIds_SliceId_SliceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SliceId. The type is interface{} with range: 0..4294967295.
    SliceId interface{}

    // ClientRate. The type is ClientDataRate.
    ClientRate interface{}

    // TrunkRate. The type is TrunkDataRate.
    TrunkRate interface{}

    // HardwareStatus. The type is HwModuleSliceStatus.
    HardwareStatus interface{}

    // DpFpgaFwType. The type is string with length: 0..10.
    DpFpgaFwType interface{}

    // DpFpgaFwVer. The type is string with length: 0..10.
    DpFpgaFwVer interface{}

    // NeedUpg. The type is interface{} with range: 0..4294967295.
    NeedUpg interface{}

    // EncryptionSupported. The type is bool.
    EncryptionSupported interface{}

    // LldpDropStatus. The type is bool.
    LldpDropStatus interface{}

    // client port. The type is slice of
    // HwModule_SliceIds_SliceId_SliceInfo_ClientPort.
    ClientPort []HwModule_SliceIds_SliceId_SliceInfo_ClientPort
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetFilter() yfilter.YFilter { return sliceInfo.YFilter }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) SetFilter(yf yfilter.YFilter) { sliceInfo.YFilter = yf }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    if yname == "client-rate" { return "ClientRate" }
    if yname == "trunk-rate" { return "TrunkRate" }
    if yname == "hardware-status" { return "HardwareStatus" }
    if yname == "dp-fpga-fw-type" { return "DpFpgaFwType" }
    if yname == "dp-fpga-fw-ver" { return "DpFpgaFwVer" }
    if yname == "need-upg" { return "NeedUpg" }
    if yname == "encryption-supported" { return "EncryptionSupported" }
    if yname == "lldp-drop-status" { return "LldpDropStatus" }
    if yname == "client-port" { return "ClientPort" }
    return ""
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetSegmentPath() string {
    return "slice-info"
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-port" {
        for _, c := range sliceInfo.ClientPort {
            if sliceInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceIds_SliceId_SliceInfo_ClientPort{}
        sliceInfo.ClientPort = append(sliceInfo.ClientPort, child)
        return &sliceInfo.ClientPort[len(sliceInfo.ClientPort)-1]
    }
    return nil
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceInfo.ClientPort {
        children[sliceInfo.ClientPort[i].GetSegmentPath()] = &sliceInfo.ClientPort[i]
    }
    return children
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-id"] = sliceInfo.SliceId
    leafs["client-rate"] = sliceInfo.ClientRate
    leafs["trunk-rate"] = sliceInfo.TrunkRate
    leafs["hardware-status"] = sliceInfo.HardwareStatus
    leafs["dp-fpga-fw-type"] = sliceInfo.DpFpgaFwType
    leafs["dp-fpga-fw-ver"] = sliceInfo.DpFpgaFwVer
    leafs["need-upg"] = sliceInfo.NeedUpg
    leafs["encryption-supported"] = sliceInfo.EncryptionSupported
    leafs["lldp-drop-status"] = sliceInfo.LldpDropStatus
    return leafs
}

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetYangName() string { return "slice-info" }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) SetParent(parent types.Entity) { sliceInfo.parent = parent }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetParent() types.Entity { return sliceInfo.parent }

func (sliceInfo *HwModule_SliceIds_SliceId_SliceInfo) GetParentYangName() string { return "slice-id" }

// HwModule_SliceIds_SliceId_SliceInfo_ClientPort
// client port
type HwModule_SliceIds_SliceId_SliceInfo_ClientPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ClientName. The type is string with length: 0..64.
    ClientName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // trunk port. The type is slice of
    // HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort.
    TrunkPort []HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetFilter() yfilter.YFilter { return clientPort.YFilter }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) SetFilter(yf yfilter.YFilter) { clientPort.YFilter = yf }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetGoName(yname string) string {
    if yname == "client-name" { return "ClientName" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "trunk-port" { return "TrunkPort" }
    return ""
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetSegmentPath() string {
    return "client-port"
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trunk-port" {
        for _, c := range clientPort.TrunkPort {
            if clientPort.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort{}
        clientPort.TrunkPort = append(clientPort.TrunkPort, child)
        return &clientPort.TrunkPort[len(clientPort.TrunkPort)-1]
    }
    return nil
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientPort.TrunkPort {
        children[clientPort.TrunkPort[i].GetSegmentPath()] = &clientPort.TrunkPort[i]
    }
    return children
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-name"] = clientPort.ClientName
    leafs["if-index"] = clientPort.IfIndex
    return leafs
}

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetBundleName() string { return "cisco_ios_xr" }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetYangName() string { return "client-port" }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) SetParent(parent types.Entity) { clientPort.parent = parent }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetParent() types.Entity { return clientPort.parent }

func (clientPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort) GetParentYangName() string { return "slice-info" }

// HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort
// trunk port
type HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TrunkName. The type is string with length: 0..64.
    TrunkName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Percentage. The type is string with length: 0..8.
    Percentage interface{}
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetFilter() yfilter.YFilter { return trunkPort.YFilter }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) SetFilter(yf yfilter.YFilter) { trunkPort.YFilter = yf }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetGoName(yname string) string {
    if yname == "trunk-name" { return "TrunkName" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "percentage" { return "Percentage" }
    return ""
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetSegmentPath() string {
    return "trunk-port"
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trunk-name"] = trunkPort.TrunkName
    leafs["if-index"] = trunkPort.IfIndex
    leafs["percentage"] = trunkPort.Percentage
    return leafs
}

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetBundleName() string { return "cisco_ios_xr" }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetYangName() string { return "trunk-port" }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) SetParent(parent types.Entity) { trunkPort.parent = parent }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetParent() types.Entity { return trunkPort.parent }

func (trunkPort *HwModule_SliceIds_SliceId_SliceInfo_ClientPort_TrunkPort) GetParentYangName() string { return "client-port" }

// HwModule_SliceAll
// Information for all slices
type HwModule_SliceAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // slice info. The type is slice of HwModule_SliceAll_SliceInfo.
    SliceInfo []HwModule_SliceAll_SliceInfo
}

func (sliceAll *HwModule_SliceAll) GetFilter() yfilter.YFilter { return sliceAll.YFilter }

func (sliceAll *HwModule_SliceAll) SetFilter(yf yfilter.YFilter) { sliceAll.YFilter = yf }

func (sliceAll *HwModule_SliceAll) GetGoName(yname string) string {
    if yname == "slice-info" { return "SliceInfo" }
    return ""
}

func (sliceAll *HwModule_SliceAll) GetSegmentPath() string {
    return "slice-all"
}

func (sliceAll *HwModule_SliceAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slice-info" {
        for _, c := range sliceAll.SliceInfo {
            if sliceAll.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceAll_SliceInfo{}
        sliceAll.SliceInfo = append(sliceAll.SliceInfo, child)
        return &sliceAll.SliceInfo[len(sliceAll.SliceInfo)-1]
    }
    return nil
}

func (sliceAll *HwModule_SliceAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceAll.SliceInfo {
        children[sliceAll.SliceInfo[i].GetSegmentPath()] = &sliceAll.SliceInfo[i]
    }
    return children
}

func (sliceAll *HwModule_SliceAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sliceAll *HwModule_SliceAll) GetBundleName() string { return "cisco_ios_xr" }

func (sliceAll *HwModule_SliceAll) GetYangName() string { return "slice-all" }

func (sliceAll *HwModule_SliceAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceAll *HwModule_SliceAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceAll *HwModule_SliceAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceAll *HwModule_SliceAll) SetParent(parent types.Entity) { sliceAll.parent = parent }

func (sliceAll *HwModule_SliceAll) GetParent() types.Entity { return sliceAll.parent }

func (sliceAll *HwModule_SliceAll) GetParentYangName() string { return "hw-module" }

// HwModule_SliceAll_SliceInfo
// slice info
type HwModule_SliceAll_SliceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SliceId. The type is interface{} with range: 0..4294967295.
    SliceId interface{}

    // ClientRate. The type is ClientDataRate.
    ClientRate interface{}

    // TrunkRate. The type is TrunkDataRate.
    TrunkRate interface{}

    // HardwareStatus. The type is HwModuleSliceStatus.
    HardwareStatus interface{}

    // DpFpgaFwType. The type is string with length: 0..10.
    DpFpgaFwType interface{}

    // DpFpgaFwVer. The type is string with length: 0..10.
    DpFpgaFwVer interface{}

    // NeedUpg. The type is interface{} with range: 0..4294967295.
    NeedUpg interface{}

    // EncryptionSupported. The type is bool.
    EncryptionSupported interface{}

    // LldpDropStatus. The type is bool.
    LldpDropStatus interface{}

    // client port. The type is slice of HwModule_SliceAll_SliceInfo_ClientPort.
    ClientPort []HwModule_SliceAll_SliceInfo_ClientPort
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetFilter() yfilter.YFilter { return sliceInfo.YFilter }

func (sliceInfo *HwModule_SliceAll_SliceInfo) SetFilter(yf yfilter.YFilter) { sliceInfo.YFilter = yf }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetGoName(yname string) string {
    if yname == "slice-id" { return "SliceId" }
    if yname == "client-rate" { return "ClientRate" }
    if yname == "trunk-rate" { return "TrunkRate" }
    if yname == "hardware-status" { return "HardwareStatus" }
    if yname == "dp-fpga-fw-type" { return "DpFpgaFwType" }
    if yname == "dp-fpga-fw-ver" { return "DpFpgaFwVer" }
    if yname == "need-upg" { return "NeedUpg" }
    if yname == "encryption-supported" { return "EncryptionSupported" }
    if yname == "lldp-drop-status" { return "LldpDropStatus" }
    if yname == "client-port" { return "ClientPort" }
    return ""
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetSegmentPath() string {
    return "slice-info"
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-port" {
        for _, c := range sliceInfo.ClientPort {
            if sliceInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceAll_SliceInfo_ClientPort{}
        sliceInfo.ClientPort = append(sliceInfo.ClientPort, child)
        return &sliceInfo.ClientPort[len(sliceInfo.ClientPort)-1]
    }
    return nil
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sliceInfo.ClientPort {
        children[sliceInfo.ClientPort[i].GetSegmentPath()] = &sliceInfo.ClientPort[i]
    }
    return children
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slice-id"] = sliceInfo.SliceId
    leafs["client-rate"] = sliceInfo.ClientRate
    leafs["trunk-rate"] = sliceInfo.TrunkRate
    leafs["hardware-status"] = sliceInfo.HardwareStatus
    leafs["dp-fpga-fw-type"] = sliceInfo.DpFpgaFwType
    leafs["dp-fpga-fw-ver"] = sliceInfo.DpFpgaFwVer
    leafs["need-upg"] = sliceInfo.NeedUpg
    leafs["encryption-supported"] = sliceInfo.EncryptionSupported
    leafs["lldp-drop-status"] = sliceInfo.LldpDropStatus
    return leafs
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetYangName() string { return "slice-info" }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sliceInfo *HwModule_SliceAll_SliceInfo) SetParent(parent types.Entity) { sliceInfo.parent = parent }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetParent() types.Entity { return sliceInfo.parent }

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetParentYangName() string { return "slice-all" }

// HwModule_SliceAll_SliceInfo_ClientPort
// client port
type HwModule_SliceAll_SliceInfo_ClientPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ClientName. The type is string with length: 0..64.
    ClientName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // trunk port. The type is slice of
    // HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort.
    TrunkPort []HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetFilter() yfilter.YFilter { return clientPort.YFilter }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) SetFilter(yf yfilter.YFilter) { clientPort.YFilter = yf }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetGoName(yname string) string {
    if yname == "client-name" { return "ClientName" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "trunk-port" { return "TrunkPort" }
    return ""
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetSegmentPath() string {
    return "client-port"
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trunk-port" {
        for _, c := range clientPort.TrunkPort {
            if clientPort.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort{}
        clientPort.TrunkPort = append(clientPort.TrunkPort, child)
        return &clientPort.TrunkPort[len(clientPort.TrunkPort)-1]
    }
    return nil
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clientPort.TrunkPort {
        children[clientPort.TrunkPort[i].GetSegmentPath()] = &clientPort.TrunkPort[i]
    }
    return children
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-name"] = clientPort.ClientName
    leafs["if-index"] = clientPort.IfIndex
    return leafs
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetBundleName() string { return "cisco_ios_xr" }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetYangName() string { return "client-port" }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) SetParent(parent types.Entity) { clientPort.parent = parent }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetParent() types.Entity { return clientPort.parent }

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetParentYangName() string { return "slice-info" }

// HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort
// trunk port
type HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TrunkName. The type is string with length: 0..64.
    TrunkName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Percentage. The type is string with length: 0..8.
    Percentage interface{}
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetFilter() yfilter.YFilter { return trunkPort.YFilter }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) SetFilter(yf yfilter.YFilter) { trunkPort.YFilter = yf }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetGoName(yname string) string {
    if yname == "trunk-name" { return "TrunkName" }
    if yname == "if-index" { return "IfIndex" }
    if yname == "percentage" { return "Percentage" }
    return ""
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetSegmentPath() string {
    return "trunk-port"
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trunk-name"] = trunkPort.TrunkName
    leafs["if-index"] = trunkPort.IfIndex
    leafs["percentage"] = trunkPort.Percentage
    return leafs
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetBundleName() string { return "cisco_ios_xr" }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetYangName() string { return "trunk-port" }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) SetParent(parent types.Entity) { trunkPort.parent = parent }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetParent() types.Entity { return trunkPort.parent }

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetParentYangName() string { return "client-port" }

