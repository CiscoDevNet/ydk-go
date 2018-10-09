// This module contains a collection of YANG definitions
// for monitoring controller VDSL operational information.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package controller_vdsl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package controller_vdsl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-controller-vdsl-oper vdsl-oper-data}", reflect.TypeOf(VdslOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-controller-vdsl-oper:vdsl-oper-data", reflect.TypeOf(VdslOperData{}))
}

// IdbStates represents Interface state
type IdbStates string

const (
    // Interface down
    IdbStates_down IdbStates = "down"

    // Interface going down
    IdbStates_going_down IdbStates = "going-down"

    // Interface init
    IdbStates_init IdbStates = "init"

    // Interface testing
    IdbStates_testing IdbStates = "testing"

    // Interface up
    IdbStates_up IdbStates = "up"

    // Interface reset
    IdbStates_reset IdbStates = "reset"

    // Interface admindown
    IdbStates_admindown IdbStates = "admindown"

    // Interface deleted
    IdbStates_deleted IdbStates = "deleted"
)

// ModeTc represents Operation Mode
type ModeTc string

const (
    // MODE Ethernet
    ModeTc_ptm ModeTc = "ptm"

    // MODE Atm
    ModeTc_atm ModeTc = "atm"
)

// VdslOperData
// VDSL controller information
type VdslOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VDSL controller information. The type is slice of VdslOperData_VdslInfo.
    VdslInfo []*VdslOperData_VdslInfo
}

func (vdslOperData *VdslOperData) GetEntityData() *types.CommonEntityData {
    vdslOperData.EntityData.YFilter = vdslOperData.YFilter
    vdslOperData.EntityData.YangName = "vdsl-oper-data"
    vdslOperData.EntityData.BundleName = "cisco_ios_xe"
    vdslOperData.EntityData.ParentYangName = "Cisco-IOS-XE-controller-vdsl-oper"
    vdslOperData.EntityData.SegmentPath = "Cisco-IOS-XE-controller-vdsl-oper:vdsl-oper-data"
    vdslOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vdslOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vdslOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vdslOperData.EntityData.Children = types.NewOrderedMap()
    vdslOperData.EntityData.Children.Append("vdsl-info", types.YChild{"VdslInfo", nil})
    for i := range vdslOperData.VdslInfo {
        vdslOperData.EntityData.Children.Append(types.GetSegmentPath(vdslOperData.VdslInfo[i]), types.YChild{"VdslInfo", vdslOperData.VdslInfo[i]})
    }
    vdslOperData.EntityData.Leafs = types.NewOrderedMap()

    vdslOperData.EntityData.YListKeys = []string {}

    return &(vdslOperData.EntityData)
}

// VdslOperData_VdslInfo
// VDSL controller information
type VdslOperData_VdslInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot number. The type is interface{} with range:
    // 0..4294967295.
    SlotNum interface{}

    // This attribute is a key. Sub Slot number. The type is interface{} with
    // range: 0..4294967295.
    SubslotNum interface{}

    // Controller database state. The type is IdbStates.
    State interface{}

    // Operational mode. The type is ModeTc.
    Mode interface{}

    // Modem status. The type is string.
    ModemStatus interface{}

    // Trained Mode. The type is string.
    TrainedMode interface{}

    // Firmware version. The type is string.
    FirmwareVersion interface{}

    // PHY version. The type is string.
    PhyVersion interface{}

    // XTU-R customer-premises equipment statistics.
    CpeStats VdslOperData_VdslInfo_CpeStats

    // XTU-C central office statistics.
    CoStats VdslOperData_VdslInfo_CoStats
}

func (vdslInfo *VdslOperData_VdslInfo) GetEntityData() *types.CommonEntityData {
    vdslInfo.EntityData.YFilter = vdslInfo.YFilter
    vdslInfo.EntityData.YangName = "vdsl-info"
    vdslInfo.EntityData.BundleName = "cisco_ios_xe"
    vdslInfo.EntityData.ParentYangName = "vdsl-oper-data"
    vdslInfo.EntityData.SegmentPath = "vdsl-info" + types.AddKeyToken(vdslInfo.SlotNum, "slot-num") + types.AddKeyToken(vdslInfo.SubslotNum, "subslot-num")
    vdslInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vdslInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vdslInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vdslInfo.EntityData.Children = types.NewOrderedMap()
    vdslInfo.EntityData.Children.Append("cpe-stats", types.YChild{"CpeStats", &vdslInfo.CpeStats})
    vdslInfo.EntityData.Children.Append("co-stats", types.YChild{"CoStats", &vdslInfo.CoStats})
    vdslInfo.EntityData.Leafs = types.NewOrderedMap()
    vdslInfo.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", vdslInfo.SlotNum})
    vdslInfo.EntityData.Leafs.Append("subslot-num", types.YLeaf{"SubslotNum", vdslInfo.SubslotNum})
    vdslInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", vdslInfo.State})
    vdslInfo.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", vdslInfo.Mode})
    vdslInfo.EntityData.Leafs.Append("modem-status", types.YLeaf{"ModemStatus", vdslInfo.ModemStatus})
    vdslInfo.EntityData.Leafs.Append("trained-mode", types.YLeaf{"TrainedMode", vdslInfo.TrainedMode})
    vdslInfo.EntityData.Leafs.Append("firmware-version", types.YLeaf{"FirmwareVersion", vdslInfo.FirmwareVersion})
    vdslInfo.EntityData.Leafs.Append("phy-version", types.YLeaf{"PhyVersion", vdslInfo.PhyVersion})

    vdslInfo.EntityData.YListKeys = []string {"SlotNum", "SubslotNum"}

    return &(vdslInfo.EntityData)
}

// VdslOperData_VdslInfo_CpeStats
// XTU-R customer-premises equipment statistics
type VdslOperData_VdslInfo_CpeStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chip vendor ID. The type is string.
    ChipVendor interface{}

    // Line attenuation in decibels. The type is string.
    LineAttenuation interface{}

    // Noise margin in decibels. The type is string.
    NoiseMargin interface{}

    // Attainable rate in kilobits per sec. The type is interface{} with range:
    // 0..4294967295.
    AttainableRate interface{}

    // Actual power in decibel-milliwatts. The type is string.
    ActualPower interface{}

    // Speed in kilobits per sec. The type is interface{} with range:
    // 0..4294967295.
    Speed interface{}
}

func (cpeStats *VdslOperData_VdslInfo_CpeStats) GetEntityData() *types.CommonEntityData {
    cpeStats.EntityData.YFilter = cpeStats.YFilter
    cpeStats.EntityData.YangName = "cpe-stats"
    cpeStats.EntityData.BundleName = "cisco_ios_xe"
    cpeStats.EntityData.ParentYangName = "vdsl-info"
    cpeStats.EntityData.SegmentPath = "cpe-stats"
    cpeStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeStats.EntityData.Children = types.NewOrderedMap()
    cpeStats.EntityData.Leafs = types.NewOrderedMap()
    cpeStats.EntityData.Leafs.Append("chip-vendor", types.YLeaf{"ChipVendor", cpeStats.ChipVendor})
    cpeStats.EntityData.Leafs.Append("line-attenuation", types.YLeaf{"LineAttenuation", cpeStats.LineAttenuation})
    cpeStats.EntityData.Leafs.Append("noise-margin", types.YLeaf{"NoiseMargin", cpeStats.NoiseMargin})
    cpeStats.EntityData.Leafs.Append("attainable-rate", types.YLeaf{"AttainableRate", cpeStats.AttainableRate})
    cpeStats.EntityData.Leafs.Append("actual-power", types.YLeaf{"ActualPower", cpeStats.ActualPower})
    cpeStats.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", cpeStats.Speed})

    cpeStats.EntityData.YListKeys = []string {}

    return &(cpeStats.EntityData)
}

// VdslOperData_VdslInfo_CoStats
// XTU-C central office statistics
type VdslOperData_VdslInfo_CoStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chip vendor ID. The type is string.
    ChipVendor interface{}

    // Line attenuation in decibels. The type is string.
    LineAttenuation interface{}

    // Noise margin in decibels. The type is string.
    NoiseMargin interface{}

    // Attainable rate in kilobits per sec. The type is interface{} with range:
    // 0..4294967295.
    AttainableRate interface{}

    // Actual power in decibel-milliwatts. The type is string.
    ActualPower interface{}

    // Speed in kilobits per sec. The type is interface{} with range:
    // 0..4294967295.
    Speed interface{}
}

func (coStats *VdslOperData_VdslInfo_CoStats) GetEntityData() *types.CommonEntityData {
    coStats.EntityData.YFilter = coStats.YFilter
    coStats.EntityData.YangName = "co-stats"
    coStats.EntityData.BundleName = "cisco_ios_xe"
    coStats.EntityData.ParentYangName = "vdsl-info"
    coStats.EntityData.SegmentPath = "co-stats"
    coStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    coStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    coStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    coStats.EntityData.Children = types.NewOrderedMap()
    coStats.EntityData.Leafs = types.NewOrderedMap()
    coStats.EntityData.Leafs.Append("chip-vendor", types.YLeaf{"ChipVendor", coStats.ChipVendor})
    coStats.EntityData.Leafs.Append("line-attenuation", types.YLeaf{"LineAttenuation", coStats.LineAttenuation})
    coStats.EntityData.Leafs.Append("noise-margin", types.YLeaf{"NoiseMargin", coStats.NoiseMargin})
    coStats.EntityData.Leafs.Append("attainable-rate", types.YLeaf{"AttainableRate", coStats.AttainableRate})
    coStats.EntityData.Leafs.Append("actual-power", types.YLeaf{"ActualPower", coStats.ActualPower})
    coStats.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", coStats.Speed})

    coStats.EntityData.YListKeys = []string {}

    return &(coStats.EntityData)
}

