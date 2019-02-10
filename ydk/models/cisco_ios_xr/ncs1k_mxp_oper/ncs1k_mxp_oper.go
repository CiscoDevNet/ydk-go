// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hw-module: mxp hw-module command chain
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slice information.
    SliceIds HwModule_SliceIds

    // Information for all slices.
    SliceAll HwModule_SliceAll
}

func (hwModule *HwModule) GetEntityData() *types.CommonEntityData {
    hwModule.EntityData.YFilter = hwModule.YFilter
    hwModule.EntityData.YangName = "hw-module"
    hwModule.EntityData.BundleName = "cisco_ios_xr"
    hwModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-mxp-oper"
    hwModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module"
    hwModule.EntityData.AbsolutePath = hwModule.EntityData.SegmentPath
    hwModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModule.EntityData.Children = types.NewOrderedMap()
    hwModule.EntityData.Children.Append("slice-ids", types.YChild{"SliceIds", &hwModule.SliceIds})
    hwModule.EntityData.Children.Append("slice-all", types.YChild{"SliceAll", &hwModule.SliceAll})
    hwModule.EntityData.Leafs = types.NewOrderedMap()

    hwModule.EntityData.YListKeys = []string {}

    return &(hwModule.EntityData)
}

// HwModule_SliceIds
// Slice information
type HwModule_SliceIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per slice num data. The type is slice of HwModule_SliceIds_SliceId.
    SliceId []*HwModule_SliceIds_SliceId
}

func (sliceIds *HwModule_SliceIds) GetEntityData() *types.CommonEntityData {
    sliceIds.EntityData.YFilter = sliceIds.YFilter
    sliceIds.EntityData.YangName = "slice-ids"
    sliceIds.EntityData.BundleName = "cisco_ios_xr"
    sliceIds.EntityData.ParentYangName = "hw-module"
    sliceIds.EntityData.SegmentPath = "slice-ids"
    sliceIds.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/" + sliceIds.EntityData.SegmentPath
    sliceIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceIds.EntityData.Children = types.NewOrderedMap()
    sliceIds.EntityData.Children.Append("slice-id", types.YChild{"SliceId", nil})
    for i := range sliceIds.SliceId {
        sliceIds.EntityData.Children.Append(types.GetSegmentPath(sliceIds.SliceId[i]), types.YChild{"SliceId", sliceIds.SliceId[i]})
    }
    sliceIds.EntityData.Leafs = types.NewOrderedMap()

    sliceIds.EntityData.YListKeys = []string {}

    return &(sliceIds.EntityData)
}

// HwModule_SliceIds_SliceId
// Per slice num data
type HwModule_SliceIds_SliceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Details associated with a particular slice number.
    // The type is interface{} with range: 0..4294967295.
    SliceNum interface{}

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

    // client port. The type is slice of HwModule_SliceIds_SliceId_ClientPort.
    ClientPort []*HwModule_SliceIds_SliceId_ClientPort
}

func (sliceId *HwModule_SliceIds_SliceId) GetEntityData() *types.CommonEntityData {
    sliceId.EntityData.YFilter = sliceId.YFilter
    sliceId.EntityData.YangName = "slice-id"
    sliceId.EntityData.BundleName = "cisco_ios_xr"
    sliceId.EntityData.ParentYangName = "slice-ids"
    sliceId.EntityData.SegmentPath = "slice-id" + types.AddKeyToken(sliceId.SliceNum, "slice-num")
    sliceId.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-ids/" + sliceId.EntityData.SegmentPath
    sliceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceId.EntityData.Children = types.NewOrderedMap()
    sliceId.EntityData.Children.Append("client-port", types.YChild{"ClientPort", nil})
    for i := range sliceId.ClientPort {
        types.SetYListKey(sliceId.ClientPort[i], i)
        sliceId.EntityData.Children.Append(types.GetSegmentPath(sliceId.ClientPort[i]), types.YChild{"ClientPort", sliceId.ClientPort[i]})
    }
    sliceId.EntityData.Leafs = types.NewOrderedMap()
    sliceId.EntityData.Leafs.Append("slice-num", types.YLeaf{"SliceNum", sliceId.SliceNum})
    sliceId.EntityData.Leafs.Append("slice-id", types.YLeaf{"SliceId", sliceId.SliceId})
    sliceId.EntityData.Leafs.Append("client-rate", types.YLeaf{"ClientRate", sliceId.ClientRate})
    sliceId.EntityData.Leafs.Append("trunk-rate", types.YLeaf{"TrunkRate", sliceId.TrunkRate})
    sliceId.EntityData.Leafs.Append("hardware-status", types.YLeaf{"HardwareStatus", sliceId.HardwareStatus})
    sliceId.EntityData.Leafs.Append("dp-fpga-fw-type", types.YLeaf{"DpFpgaFwType", sliceId.DpFpgaFwType})
    sliceId.EntityData.Leafs.Append("dp-fpga-fw-ver", types.YLeaf{"DpFpgaFwVer", sliceId.DpFpgaFwVer})
    sliceId.EntityData.Leafs.Append("need-upg", types.YLeaf{"NeedUpg", sliceId.NeedUpg})
    sliceId.EntityData.Leafs.Append("encryption-supported", types.YLeaf{"EncryptionSupported", sliceId.EncryptionSupported})
    sliceId.EntityData.Leafs.Append("lldp-drop-status", types.YLeaf{"LldpDropStatus", sliceId.LldpDropStatus})

    sliceId.EntityData.YListKeys = []string {"SliceNum"}

    return &(sliceId.EntityData)
}

// HwModule_SliceIds_SliceId_ClientPort
// client port
type HwModule_SliceIds_SliceId_ClientPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ClientName. The type is string with length: 0..64.
    ClientName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // trunk port. The type is slice of
    // HwModule_SliceIds_SliceId_ClientPort_TrunkPort.
    TrunkPort []*HwModule_SliceIds_SliceId_ClientPort_TrunkPort
}

func (clientPort *HwModule_SliceIds_SliceId_ClientPort) GetEntityData() *types.CommonEntityData {
    clientPort.EntityData.YFilter = clientPort.YFilter
    clientPort.EntityData.YangName = "client-port"
    clientPort.EntityData.BundleName = "cisco_ios_xr"
    clientPort.EntityData.ParentYangName = "slice-id"
    clientPort.EntityData.SegmentPath = "client-port" + types.AddNoKeyToken(clientPort)
    clientPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-ids/slice-id/" + clientPort.EntityData.SegmentPath
    clientPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientPort.EntityData.Children = types.NewOrderedMap()
    clientPort.EntityData.Children.Append("trunk-port", types.YChild{"TrunkPort", nil})
    for i := range clientPort.TrunkPort {
        types.SetYListKey(clientPort.TrunkPort[i], i)
        clientPort.EntityData.Children.Append(types.GetSegmentPath(clientPort.TrunkPort[i]), types.YChild{"TrunkPort", clientPort.TrunkPort[i]})
    }
    clientPort.EntityData.Leafs = types.NewOrderedMap()
    clientPort.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", clientPort.ClientName})
    clientPort.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", clientPort.IfIndex})

    clientPort.EntityData.YListKeys = []string {}

    return &(clientPort.EntityData)
}

// HwModule_SliceIds_SliceId_ClientPort_TrunkPort
// trunk port
type HwModule_SliceIds_SliceId_ClientPort_TrunkPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // TrunkName. The type is string with length: 0..64.
    TrunkName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Percentage. The type is string with length: 0..8.
    Percentage interface{}
}

func (trunkPort *HwModule_SliceIds_SliceId_ClientPort_TrunkPort) GetEntityData() *types.CommonEntityData {
    trunkPort.EntityData.YFilter = trunkPort.YFilter
    trunkPort.EntityData.YangName = "trunk-port"
    trunkPort.EntityData.BundleName = "cisco_ios_xr"
    trunkPort.EntityData.ParentYangName = "client-port"
    trunkPort.EntityData.SegmentPath = "trunk-port" + types.AddNoKeyToken(trunkPort)
    trunkPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-ids/slice-id/client-port/" + trunkPort.EntityData.SegmentPath
    trunkPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunkPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunkPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunkPort.EntityData.Children = types.NewOrderedMap()
    trunkPort.EntityData.Leafs = types.NewOrderedMap()
    trunkPort.EntityData.Leafs.Append("trunk-name", types.YLeaf{"TrunkName", trunkPort.TrunkName})
    trunkPort.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", trunkPort.IfIndex})
    trunkPort.EntityData.Leafs.Append("percentage", types.YLeaf{"Percentage", trunkPort.Percentage})

    trunkPort.EntityData.YListKeys = []string {}

    return &(trunkPort.EntityData)
}

// HwModule_SliceAll
// Information for all slices
type HwModule_SliceAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // slice info. The type is slice of HwModule_SliceAll_SliceInfo.
    SliceInfo []*HwModule_SliceAll_SliceInfo
}

func (sliceAll *HwModule_SliceAll) GetEntityData() *types.CommonEntityData {
    sliceAll.EntityData.YFilter = sliceAll.YFilter
    sliceAll.EntityData.YangName = "slice-all"
    sliceAll.EntityData.BundleName = "cisco_ios_xr"
    sliceAll.EntityData.ParentYangName = "hw-module"
    sliceAll.EntityData.SegmentPath = "slice-all"
    sliceAll.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/" + sliceAll.EntityData.SegmentPath
    sliceAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceAll.EntityData.Children = types.NewOrderedMap()
    sliceAll.EntityData.Children.Append("slice-info", types.YChild{"SliceInfo", nil})
    for i := range sliceAll.SliceInfo {
        types.SetYListKey(sliceAll.SliceInfo[i], i)
        sliceAll.EntityData.Children.Append(types.GetSegmentPath(sliceAll.SliceInfo[i]), types.YChild{"SliceInfo", sliceAll.SliceInfo[i]})
    }
    sliceAll.EntityData.Leafs = types.NewOrderedMap()

    sliceAll.EntityData.YListKeys = []string {}

    return &(sliceAll.EntityData)
}

// HwModule_SliceAll_SliceInfo
// slice info
type HwModule_SliceAll_SliceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    ClientPort []*HwModule_SliceAll_SliceInfo_ClientPort
}

func (sliceInfo *HwModule_SliceAll_SliceInfo) GetEntityData() *types.CommonEntityData {
    sliceInfo.EntityData.YFilter = sliceInfo.YFilter
    sliceInfo.EntityData.YangName = "slice-info"
    sliceInfo.EntityData.BundleName = "cisco_ios_xr"
    sliceInfo.EntityData.ParentYangName = "slice-all"
    sliceInfo.EntityData.SegmentPath = "slice-info" + types.AddNoKeyToken(sliceInfo)
    sliceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-all/" + sliceInfo.EntityData.SegmentPath
    sliceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceInfo.EntityData.Children = types.NewOrderedMap()
    sliceInfo.EntityData.Children.Append("client-port", types.YChild{"ClientPort", nil})
    for i := range sliceInfo.ClientPort {
        types.SetYListKey(sliceInfo.ClientPort[i], i)
        sliceInfo.EntityData.Children.Append(types.GetSegmentPath(sliceInfo.ClientPort[i]), types.YChild{"ClientPort", sliceInfo.ClientPort[i]})
    }
    sliceInfo.EntityData.Leafs = types.NewOrderedMap()
    sliceInfo.EntityData.Leafs.Append("slice-id", types.YLeaf{"SliceId", sliceInfo.SliceId})
    sliceInfo.EntityData.Leafs.Append("client-rate", types.YLeaf{"ClientRate", sliceInfo.ClientRate})
    sliceInfo.EntityData.Leafs.Append("trunk-rate", types.YLeaf{"TrunkRate", sliceInfo.TrunkRate})
    sliceInfo.EntityData.Leafs.Append("hardware-status", types.YLeaf{"HardwareStatus", sliceInfo.HardwareStatus})
    sliceInfo.EntityData.Leafs.Append("dp-fpga-fw-type", types.YLeaf{"DpFpgaFwType", sliceInfo.DpFpgaFwType})
    sliceInfo.EntityData.Leafs.Append("dp-fpga-fw-ver", types.YLeaf{"DpFpgaFwVer", sliceInfo.DpFpgaFwVer})
    sliceInfo.EntityData.Leafs.Append("need-upg", types.YLeaf{"NeedUpg", sliceInfo.NeedUpg})
    sliceInfo.EntityData.Leafs.Append("encryption-supported", types.YLeaf{"EncryptionSupported", sliceInfo.EncryptionSupported})
    sliceInfo.EntityData.Leafs.Append("lldp-drop-status", types.YLeaf{"LldpDropStatus", sliceInfo.LldpDropStatus})

    sliceInfo.EntityData.YListKeys = []string {}

    return &(sliceInfo.EntityData)
}

// HwModule_SliceAll_SliceInfo_ClientPort
// client port
type HwModule_SliceAll_SliceInfo_ClientPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // ClientName. The type is string with length: 0..64.
    ClientName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // trunk port. The type is slice of
    // HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort.
    TrunkPort []*HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort
}

func (clientPort *HwModule_SliceAll_SliceInfo_ClientPort) GetEntityData() *types.CommonEntityData {
    clientPort.EntityData.YFilter = clientPort.YFilter
    clientPort.EntityData.YangName = "client-port"
    clientPort.EntityData.BundleName = "cisco_ios_xr"
    clientPort.EntityData.ParentYangName = "slice-info"
    clientPort.EntityData.SegmentPath = "client-port" + types.AddNoKeyToken(clientPort)
    clientPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-all/slice-info/" + clientPort.EntityData.SegmentPath
    clientPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientPort.EntityData.Children = types.NewOrderedMap()
    clientPort.EntityData.Children.Append("trunk-port", types.YChild{"TrunkPort", nil})
    for i := range clientPort.TrunkPort {
        types.SetYListKey(clientPort.TrunkPort[i], i)
        clientPort.EntityData.Children.Append(types.GetSegmentPath(clientPort.TrunkPort[i]), types.YChild{"TrunkPort", clientPort.TrunkPort[i]})
    }
    clientPort.EntityData.Leafs = types.NewOrderedMap()
    clientPort.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", clientPort.ClientName})
    clientPort.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", clientPort.IfIndex})

    clientPort.EntityData.YListKeys = []string {}

    return &(clientPort.EntityData)
}

// HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort
// trunk port
type HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // TrunkName. The type is string with length: 0..64.
    TrunkName interface{}

    // IfIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Percentage. The type is string with length: 0..8.
    Percentage interface{}
}

func (trunkPort *HwModule_SliceAll_SliceInfo_ClientPort_TrunkPort) GetEntityData() *types.CommonEntityData {
    trunkPort.EntityData.YFilter = trunkPort.YFilter
    trunkPort.EntityData.YangName = "trunk-port"
    trunkPort.EntityData.BundleName = "cisco_ios_xr"
    trunkPort.EntityData.ParentYangName = "client-port"
    trunkPort.EntityData.SegmentPath = "trunk-port" + types.AddNoKeyToken(trunkPort)
    trunkPort.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-oper:hw-module/slice-all/slice-info/client-port/" + trunkPort.EntityData.SegmentPath
    trunkPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunkPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunkPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunkPort.EntityData.Children = types.NewOrderedMap()
    trunkPort.EntityData.Leafs = types.NewOrderedMap()
    trunkPort.EntityData.Leafs.Append("trunk-name", types.YLeaf{"TrunkName", trunkPort.TrunkName})
    trunkPort.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", trunkPort.IfIndex})
    trunkPort.EntityData.Leafs.Append("percentage", types.YLeaf{"Percentage", trunkPort.Percentage})

    trunkPort.EntityData.YListKeys = []string {}

    return &(trunkPort.EntityData)
}

