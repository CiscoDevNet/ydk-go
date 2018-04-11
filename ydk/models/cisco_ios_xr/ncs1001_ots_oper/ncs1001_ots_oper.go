// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1001-ots package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hw-module: ncs1001 hw-module command chain
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ncs1001_ots_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1001_ots_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1001-ots-oper hw-module}", reflect.TypeOf(HwModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1001-ots-oper:hw-module", reflect.TypeOf(HwModule{}))
}

// OtsAmpliTrailData represents Ots ampli trail data
type OtsAmpliTrailData string

const (
    // Displays Booster and Pre trail data
    OtsAmpliTrailData_all OtsAmpliTrailData = "all"

    // Displays Booster trail data
    OtsAmpliTrailData_bst OtsAmpliTrailData = "bst"

    // Displays Pre trail data
    OtsAmpliTrailData_pre OtsAmpliTrailData = "pre"
)

// OtsChannelsTrailData represents Ots channels trail data
type OtsChannelsTrailData string

const (
    // Displays all channels trail data
    OtsChannelsTrailData_all OtsChannelsTrailData = "all"

    // Displays active channels trail data
    OtsChannelsTrailData_active OtsChannelsTrailData = "active"
)

// Chfilter represents Chfilter
type Chfilter string

const (
    // Active
    Chfilter_ch_filter_active Chfilter = "ch-filter-active"

    // All
    Chfilter_ch_filter_all Chfilter = "ch-filter-all"
)

// Trailview represents Trailview
type Trailview string

const (
    // All
    Trailview_trail_view_all Trailview = "trail-view-all"

    // Booster
    Trailview_trail_view_bst Trailview = "trail-view-bst"

    // Pre
    Trailview_trail_view_pre Trailview = "trail-view-pre"
)

// HwModule
// ncs1001 hw-module command chain
type HwModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amplifier trail data information.
    SlotTables HwModule_SlotTables

    // Channels trail data information.
    ChannelsSlots HwModule_ChannelsSlots
}

func (hwModule *HwModule) GetEntityData() *types.CommonEntityData {
    hwModule.EntityData.YFilter = hwModule.YFilter
    hwModule.EntityData.YangName = "hw-module"
    hwModule.EntityData.BundleName = "cisco_ios_xr"
    hwModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1001-ots-oper"
    hwModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1001-ots-oper:hw-module"
    hwModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModule.EntityData.Children = make(map[string]types.YChild)
    hwModule.EntityData.Children["slot-tables"] = types.YChild{"SlotTables", &hwModule.SlotTables}
    hwModule.EntityData.Children["channels-slots"] = types.YChild{"ChannelsSlots", &hwModule.ChannelsSlots}
    hwModule.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwModule.EntityData)
}

// HwModule_SlotTables
// Amplifier trail data information
type HwModule_SlotTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot Id. The type is slice of HwModule_SlotTables_SlotTable.
    SlotTable []HwModule_SlotTables_SlotTable
}

func (slotTables *HwModule_SlotTables) GetEntityData() *types.CommonEntityData {
    slotTables.EntityData.YFilter = slotTables.YFilter
    slotTables.EntityData.YangName = "slot-tables"
    slotTables.EntityData.BundleName = "cisco_ios_xr"
    slotTables.EntityData.ParentYangName = "hw-module"
    slotTables.EntityData.SegmentPath = "slot-tables"
    slotTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slotTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slotTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slotTables.EntityData.Children = make(map[string]types.YChild)
    slotTables.EntityData.Children["slot-table"] = types.YChild{"SlotTable", nil}
    for i := range slotTables.SlotTable {
        slotTables.EntityData.Children[types.GetSegmentPath(&slotTables.SlotTable[i])] = types.YChild{"SlotTable", &slotTables.SlotTable[i]}
    }
    slotTables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slotTables.EntityData)
}

// HwModule_SlotTables_SlotTable
// Slot Id
type HwModule_SlotTables_SlotTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: -2147483648..2147483647.
    SlotId interface{}

    // Trail data information. The type is slice of
    // HwModule_SlotTables_SlotTable_AmplifierTrailData.
    AmplifierTrailData []HwModule_SlotTables_SlotTable_AmplifierTrailData
}

func (slotTable *HwModule_SlotTables_SlotTable) GetEntityData() *types.CommonEntityData {
    slotTable.EntityData.YFilter = slotTable.YFilter
    slotTable.EntityData.YangName = "slot-table"
    slotTable.EntityData.BundleName = "cisco_ios_xr"
    slotTable.EntityData.ParentYangName = "slot-tables"
    slotTable.EntityData.SegmentPath = "slot-table" + "[slot-id='" + fmt.Sprintf("%v", slotTable.SlotId) + "']"
    slotTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slotTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slotTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slotTable.EntityData.Children = make(map[string]types.YChild)
    slotTable.EntityData.Children["amplifier-trail-data"] = types.YChild{"AmplifierTrailData", nil}
    for i := range slotTable.AmplifierTrailData {
        slotTable.EntityData.Children[types.GetSegmentPath(&slotTable.AmplifierTrailData[i])] = types.YChild{"AmplifierTrailData", &slotTable.AmplifierTrailData[i]}
    }
    slotTable.EntityData.Leafs = make(map[string]types.YLeaf)
    slotTable.EntityData.Leafs["slot-id"] = types.YLeaf{"SlotId", slotTable.SlotId}
    return &(slotTable.EntityData)
}

// HwModule_SlotTables_SlotTable_AmplifierTrailData
// Trail data information
type HwModule_SlotTables_SlotTable_AmplifierTrailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Select trail data. The type is OtsAmpliTrailData.
    AmplifierTrailDataType interface{}

    // ampli trail info. The type is slice of
    // HwModule_SlotTables_SlotTable_AmplifierTrailData_AmpliTrailInfo.
    AmpliTrailInfo []HwModule_SlotTables_SlotTable_AmplifierTrailData_AmpliTrailInfo

    // channel trail info. The type is slice of
    // HwModule_SlotTables_SlotTable_AmplifierTrailData_ChannelTrailInfo.
    ChannelTrailInfo []HwModule_SlotTables_SlotTable_AmplifierTrailData_ChannelTrailInfo
}

func (amplifierTrailData *HwModule_SlotTables_SlotTable_AmplifierTrailData) GetEntityData() *types.CommonEntityData {
    amplifierTrailData.EntityData.YFilter = amplifierTrailData.YFilter
    amplifierTrailData.EntityData.YangName = "amplifier-trail-data"
    amplifierTrailData.EntityData.BundleName = "cisco_ios_xr"
    amplifierTrailData.EntityData.ParentYangName = "slot-table"
    amplifierTrailData.EntityData.SegmentPath = "amplifier-trail-data" + "[amplifier-trail-data-type='" + fmt.Sprintf("%v", amplifierTrailData.AmplifierTrailDataType) + "']"
    amplifierTrailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifierTrailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifierTrailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifierTrailData.EntityData.Children = make(map[string]types.YChild)
    amplifierTrailData.EntityData.Children["ampli-trail-info"] = types.YChild{"AmpliTrailInfo", nil}
    for i := range amplifierTrailData.AmpliTrailInfo {
        amplifierTrailData.EntityData.Children[types.GetSegmentPath(&amplifierTrailData.AmpliTrailInfo[i])] = types.YChild{"AmpliTrailInfo", &amplifierTrailData.AmpliTrailInfo[i]}
    }
    amplifierTrailData.EntityData.Children["channel-trail-info"] = types.YChild{"ChannelTrailInfo", nil}
    for i := range amplifierTrailData.ChannelTrailInfo {
        amplifierTrailData.EntityData.Children[types.GetSegmentPath(&amplifierTrailData.ChannelTrailInfo[i])] = types.YChild{"ChannelTrailInfo", &amplifierTrailData.ChannelTrailInfo[i]}
    }
    amplifierTrailData.EntityData.Leafs = make(map[string]types.YLeaf)
    amplifierTrailData.EntityData.Leafs["amplifier-trail-data-type"] = types.YLeaf{"AmplifierTrailDataType", amplifierTrailData.AmplifierTrailDataType}
    return &(amplifierTrailData.EntityData)
}

// HwModule_SlotTables_SlotTable_AmplifierTrailData_AmpliTrailInfo
// ampli trail info
type HwModule_SlotTables_SlotTable_AmplifierTrailData_AmpliTrailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Eqpt fail. The type is interface{} with range: 0..255.
    EqptFail interface{}

    // View. The type is Trailview.
    View interface{}

    // ComPortNumber. The type is interface{} with range: 0..255.
    ComPortNumber interface{}

    // ComPortName. The type is string with length: 0..32.
    ComPortName interface{}

    // LinePortNumber. The type is interface{} with range: 0..255.
    LinePortNumber interface{}

    // LinePortName. The type is string with length: 0..32.
    LinePortName interface{}

    // BstInRxPower. The type is interface{} with range: -32768..32767.
    BstInRxPower interface{}

    // BstInRxTotalPower. The type is interface{} with range: -32768..32767.
    BstInRxTotalPower interface{}

    // BstInRxPowerThLow. The type is interface{} with range: -32768..32767.
    BstInRxPowerThLow interface{}

    // BstOutTxPower. The type is interface{} with range: -32768..32767.
    BstOutTxPower interface{}

    // BstOutTxTotalPower. The type is interface{} with range: -32768..32767.
    BstOutTxTotalPower interface{}

    // BstOutTxPowerThLow. The type is interface{} with range: -32768..32767.
    BstOutTxPowerThLow interface{}

    // BstWorkingMode. The type is interface{} with range: 0..65535.
    BstWorkingMode interface{}

    // BstSafetyMode. The type is interface{} with range: 0..65535.
    BstSafetyMode interface{}

    // BstGainRange. The type is interface{} with range: 0..65535.
    BstGainRange interface{}

    // BstOsri. The type is interface{} with range: 0..65535.
    BstOsri interface{}

    // BstChannelPower. The type is interface{} with range: -32768..32767.
    BstChannelPower interface{}

    // BstGain. The type is interface{} with range: -32768..32767.
    BstGain interface{}

    // BstTilt. The type is interface{} with range: -32768..32767.
    BstTilt interface{}

    // PreInRxPower. The type is interface{} with range: -32768..32767.
    PreInRxPower interface{}

    // PreInRxTotalPower. The type is interface{} with range: -32768..32767.
    PreInRxTotalPower interface{}

    // PreInRxPowerThLow. The type is interface{} with range: -32768..32767.
    PreInRxPowerThLow interface{}

    // PreOutTxPower. The type is interface{} with range: -32768..32767.
    PreOutTxPower interface{}

    // PreOutTxTotalPower. The type is interface{} with range: -32768..32767.
    PreOutTxTotalPower interface{}

    // PreOutTxPowerThLow. The type is interface{} with range: -32768..32767.
    PreOutTxPowerThLow interface{}

    // PreWorkingMode. The type is interface{} with range: 0..65535.
    PreWorkingMode interface{}

    // PreSafetyMode. The type is interface{} with range: 0..65535.
    PreSafetyMode interface{}

    // PreGainRange. The type is interface{} with range: 0..65535.
    PreGainRange interface{}

    // PreOsri. The type is interface{} with range: 0..65535.
    PreOsri interface{}

    // PreChannelPower. The type is interface{} with range: -32768..32767.
    PreChannelPower interface{}

    // PreGain. The type is interface{} with range: -32768..32767.
    PreGain interface{}

    // PreTilt. The type is interface{} with range: -32768..32767.
    PreTilt interface{}
}

func (ampliTrailInfo *HwModule_SlotTables_SlotTable_AmplifierTrailData_AmpliTrailInfo) GetEntityData() *types.CommonEntityData {
    ampliTrailInfo.EntityData.YFilter = ampliTrailInfo.YFilter
    ampliTrailInfo.EntityData.YangName = "ampli-trail-info"
    ampliTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    ampliTrailInfo.EntityData.ParentYangName = "amplifier-trail-data"
    ampliTrailInfo.EntityData.SegmentPath = "ampli-trail-info"
    ampliTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampliTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampliTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampliTrailInfo.EntityData.Children = make(map[string]types.YChild)
    ampliTrailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ampliTrailInfo.EntityData.Leafs["eqpt-fail"] = types.YLeaf{"EqptFail", ampliTrailInfo.EqptFail}
    ampliTrailInfo.EntityData.Leafs["view"] = types.YLeaf{"View", ampliTrailInfo.View}
    ampliTrailInfo.EntityData.Leafs["com-port-number"] = types.YLeaf{"ComPortNumber", ampliTrailInfo.ComPortNumber}
    ampliTrailInfo.EntityData.Leafs["com-port-name"] = types.YLeaf{"ComPortName", ampliTrailInfo.ComPortName}
    ampliTrailInfo.EntityData.Leafs["line-port-number"] = types.YLeaf{"LinePortNumber", ampliTrailInfo.LinePortNumber}
    ampliTrailInfo.EntityData.Leafs["line-port-name"] = types.YLeaf{"LinePortName", ampliTrailInfo.LinePortName}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-power"] = types.YLeaf{"BstInRxPower", ampliTrailInfo.BstInRxPower}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-total-power"] = types.YLeaf{"BstInRxTotalPower", ampliTrailInfo.BstInRxTotalPower}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-power-th-low"] = types.YLeaf{"BstInRxPowerThLow", ampliTrailInfo.BstInRxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-power"] = types.YLeaf{"BstOutTxPower", ampliTrailInfo.BstOutTxPower}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-total-power"] = types.YLeaf{"BstOutTxTotalPower", ampliTrailInfo.BstOutTxTotalPower}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-power-th-low"] = types.YLeaf{"BstOutTxPowerThLow", ampliTrailInfo.BstOutTxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["bst-working-mode"] = types.YLeaf{"BstWorkingMode", ampliTrailInfo.BstWorkingMode}
    ampliTrailInfo.EntityData.Leafs["bst-safety-mode"] = types.YLeaf{"BstSafetyMode", ampliTrailInfo.BstSafetyMode}
    ampliTrailInfo.EntityData.Leafs["bst-gain-range"] = types.YLeaf{"BstGainRange", ampliTrailInfo.BstGainRange}
    ampliTrailInfo.EntityData.Leafs["bst-osri"] = types.YLeaf{"BstOsri", ampliTrailInfo.BstOsri}
    ampliTrailInfo.EntityData.Leafs["bst-channel-power"] = types.YLeaf{"BstChannelPower", ampliTrailInfo.BstChannelPower}
    ampliTrailInfo.EntityData.Leafs["bst-gain"] = types.YLeaf{"BstGain", ampliTrailInfo.BstGain}
    ampliTrailInfo.EntityData.Leafs["bst-tilt"] = types.YLeaf{"BstTilt", ampliTrailInfo.BstTilt}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-power"] = types.YLeaf{"PreInRxPower", ampliTrailInfo.PreInRxPower}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-total-power"] = types.YLeaf{"PreInRxTotalPower", ampliTrailInfo.PreInRxTotalPower}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-power-th-low"] = types.YLeaf{"PreInRxPowerThLow", ampliTrailInfo.PreInRxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-power"] = types.YLeaf{"PreOutTxPower", ampliTrailInfo.PreOutTxPower}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-total-power"] = types.YLeaf{"PreOutTxTotalPower", ampliTrailInfo.PreOutTxTotalPower}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-power-th-low"] = types.YLeaf{"PreOutTxPowerThLow", ampliTrailInfo.PreOutTxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["pre-working-mode"] = types.YLeaf{"PreWorkingMode", ampliTrailInfo.PreWorkingMode}
    ampliTrailInfo.EntityData.Leafs["pre-safety-mode"] = types.YLeaf{"PreSafetyMode", ampliTrailInfo.PreSafetyMode}
    ampliTrailInfo.EntityData.Leafs["pre-gain-range"] = types.YLeaf{"PreGainRange", ampliTrailInfo.PreGainRange}
    ampliTrailInfo.EntityData.Leafs["pre-osri"] = types.YLeaf{"PreOsri", ampliTrailInfo.PreOsri}
    ampliTrailInfo.EntityData.Leafs["pre-channel-power"] = types.YLeaf{"PreChannelPower", ampliTrailInfo.PreChannelPower}
    ampliTrailInfo.EntityData.Leafs["pre-gain"] = types.YLeaf{"PreGain", ampliTrailInfo.PreGain}
    ampliTrailInfo.EntityData.Leafs["pre-tilt"] = types.YLeaf{"PreTilt", ampliTrailInfo.PreTilt}
    return &(ampliTrailInfo.EntityData)
}

// HwModule_SlotTables_SlotTable_AmplifierTrailData_ChannelTrailInfo
// channel trail info
type HwModule_SlotTables_SlotTable_AmplifierTrailData_ChannelTrailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Eqpt fail. The type is interface{} with range: 0..255.
    EqptFail interface{}

    // Filter. The type is Chfilter.
    Filter interface{}

    // OchName. The type is string with length: 0..64.
    OchName interface{}

    // Wavelength. The type is interface{} with range: 0..4294967295.
    Wavelength interface{}

    // Frequency. The type is interface{} with range: 0..4294967295.
    Frequency interface{}

    // ComPortNumber. The type is interface{} with range: 0..255.
    ComPortNumber interface{}

    // ComPortName. The type is string with length: 0..32.
    ComPortName interface{}

    // ComRxPowerThLow. The type is interface{} with range: -32768..32767.
    ComRxPowerThLow interface{}

    // LinePortNumber. The type is interface{} with range: 0..255.
    LinePortNumber interface{}

    // LinePortName. The type is string with length: 0..32.
    LinePortName interface{}

    // LineRxPowerThLow. The type is interface{} with range: -32768..32767.
    LineRxPowerThLow interface{}

    // BstInRxPower. The type is interface{} with range: -32768..32767.
    BstInRxPower interface{}

    // BstOutTxPower. The type is interface{} with range: -32768..32767.
    BstOutTxPower interface{}

    // PreInRxPower. The type is interface{} with range: -32768..32767.
    PreInRxPower interface{}

    // PreOutTxPower. The type is interface{} with range: -32768..32767.
    PreOutTxPower interface{}
}

func (channelTrailInfo *HwModule_SlotTables_SlotTable_AmplifierTrailData_ChannelTrailInfo) GetEntityData() *types.CommonEntityData {
    channelTrailInfo.EntityData.YFilter = channelTrailInfo.YFilter
    channelTrailInfo.EntityData.YangName = "channel-trail-info"
    channelTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    channelTrailInfo.EntityData.ParentYangName = "amplifier-trail-data"
    channelTrailInfo.EntityData.SegmentPath = "channel-trail-info"
    channelTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelTrailInfo.EntityData.Children = make(map[string]types.YChild)
    channelTrailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    channelTrailInfo.EntityData.Leafs["eqpt-fail"] = types.YLeaf{"EqptFail", channelTrailInfo.EqptFail}
    channelTrailInfo.EntityData.Leafs["filter"] = types.YLeaf{"Filter", channelTrailInfo.Filter}
    channelTrailInfo.EntityData.Leafs["och-name"] = types.YLeaf{"OchName", channelTrailInfo.OchName}
    channelTrailInfo.EntityData.Leafs["wavelength"] = types.YLeaf{"Wavelength", channelTrailInfo.Wavelength}
    channelTrailInfo.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", channelTrailInfo.Frequency}
    channelTrailInfo.EntityData.Leafs["com-port-number"] = types.YLeaf{"ComPortNumber", channelTrailInfo.ComPortNumber}
    channelTrailInfo.EntityData.Leafs["com-port-name"] = types.YLeaf{"ComPortName", channelTrailInfo.ComPortName}
    channelTrailInfo.EntityData.Leafs["com-rx-power-th-low"] = types.YLeaf{"ComRxPowerThLow", channelTrailInfo.ComRxPowerThLow}
    channelTrailInfo.EntityData.Leafs["line-port-number"] = types.YLeaf{"LinePortNumber", channelTrailInfo.LinePortNumber}
    channelTrailInfo.EntityData.Leafs["line-port-name"] = types.YLeaf{"LinePortName", channelTrailInfo.LinePortName}
    channelTrailInfo.EntityData.Leafs["line-rx-power-th-low"] = types.YLeaf{"LineRxPowerThLow", channelTrailInfo.LineRxPowerThLow}
    channelTrailInfo.EntityData.Leafs["bst-in-rx-power"] = types.YLeaf{"BstInRxPower", channelTrailInfo.BstInRxPower}
    channelTrailInfo.EntityData.Leafs["bst-out-tx-power"] = types.YLeaf{"BstOutTxPower", channelTrailInfo.BstOutTxPower}
    channelTrailInfo.EntityData.Leafs["pre-in-rx-power"] = types.YLeaf{"PreInRxPower", channelTrailInfo.PreInRxPower}
    channelTrailInfo.EntityData.Leafs["pre-out-tx-power"] = types.YLeaf{"PreOutTxPower", channelTrailInfo.PreOutTxPower}
    return &(channelTrailInfo.EntityData)
}

// HwModule_ChannelsSlots
// Channels trail data information
type HwModule_ChannelsSlots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot Id. The type is slice of HwModule_ChannelsSlots_ChannelsSlot.
    ChannelsSlot []HwModule_ChannelsSlots_ChannelsSlot
}

func (channelsSlots *HwModule_ChannelsSlots) GetEntityData() *types.CommonEntityData {
    channelsSlots.EntityData.YFilter = channelsSlots.YFilter
    channelsSlots.EntityData.YangName = "channels-slots"
    channelsSlots.EntityData.BundleName = "cisco_ios_xr"
    channelsSlots.EntityData.ParentYangName = "hw-module"
    channelsSlots.EntityData.SegmentPath = "channels-slots"
    channelsSlots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsSlots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsSlots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsSlots.EntityData.Children = make(map[string]types.YChild)
    channelsSlots.EntityData.Children["channels-slot"] = types.YChild{"ChannelsSlot", nil}
    for i := range channelsSlots.ChannelsSlot {
        channelsSlots.EntityData.Children[types.GetSegmentPath(&channelsSlots.ChannelsSlot[i])] = types.YChild{"ChannelsSlot", &channelsSlots.ChannelsSlot[i]}
    }
    channelsSlots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(channelsSlots.EntityData)
}

// HwModule_ChannelsSlots_ChannelsSlot
// Slot Id
type HwModule_ChannelsSlots_ChannelsSlot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: 1..3.
    SlotId interface{}

    // Trail data information. The type is slice of
    // HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData.
    ChannelsTrailData []HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData
}

func (channelsSlot *HwModule_ChannelsSlots_ChannelsSlot) GetEntityData() *types.CommonEntityData {
    channelsSlot.EntityData.YFilter = channelsSlot.YFilter
    channelsSlot.EntityData.YangName = "channels-slot"
    channelsSlot.EntityData.BundleName = "cisco_ios_xr"
    channelsSlot.EntityData.ParentYangName = "channels-slots"
    channelsSlot.EntityData.SegmentPath = "channels-slot" + "[slot-id='" + fmt.Sprintf("%v", channelsSlot.SlotId) + "']"
    channelsSlot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsSlot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsSlot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsSlot.EntityData.Children = make(map[string]types.YChild)
    channelsSlot.EntityData.Children["channels-trail-data"] = types.YChild{"ChannelsTrailData", nil}
    for i := range channelsSlot.ChannelsTrailData {
        channelsSlot.EntityData.Children[types.GetSegmentPath(&channelsSlot.ChannelsTrailData[i])] = types.YChild{"ChannelsTrailData", &channelsSlot.ChannelsTrailData[i]}
    }
    channelsSlot.EntityData.Leafs = make(map[string]types.YLeaf)
    channelsSlot.EntityData.Leafs["slot-id"] = types.YLeaf{"SlotId", channelsSlot.SlotId}
    return &(channelsSlot.EntityData)
}

// HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData
// Trail data information
type HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Select trail data. The type is
    // OtsChannelsTrailData.
    ChannelsTrailDataType interface{}

    // ampli trail info. The type is slice of
    // HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_AmpliTrailInfo.
    AmpliTrailInfo []HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_AmpliTrailInfo

    // channel trail info. The type is slice of
    // HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_ChannelTrailInfo.
    ChannelTrailInfo []HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_ChannelTrailInfo
}

func (channelsTrailData *HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData) GetEntityData() *types.CommonEntityData {
    channelsTrailData.EntityData.YFilter = channelsTrailData.YFilter
    channelsTrailData.EntityData.YangName = "channels-trail-data"
    channelsTrailData.EntityData.BundleName = "cisco_ios_xr"
    channelsTrailData.EntityData.ParentYangName = "channels-slot"
    channelsTrailData.EntityData.SegmentPath = "channels-trail-data" + "[channels-trail-data-type='" + fmt.Sprintf("%v", channelsTrailData.ChannelsTrailDataType) + "']"
    channelsTrailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsTrailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsTrailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsTrailData.EntityData.Children = make(map[string]types.YChild)
    channelsTrailData.EntityData.Children["ampli-trail-info"] = types.YChild{"AmpliTrailInfo", nil}
    for i := range channelsTrailData.AmpliTrailInfo {
        channelsTrailData.EntityData.Children[types.GetSegmentPath(&channelsTrailData.AmpliTrailInfo[i])] = types.YChild{"AmpliTrailInfo", &channelsTrailData.AmpliTrailInfo[i]}
    }
    channelsTrailData.EntityData.Children["channel-trail-info"] = types.YChild{"ChannelTrailInfo", nil}
    for i := range channelsTrailData.ChannelTrailInfo {
        channelsTrailData.EntityData.Children[types.GetSegmentPath(&channelsTrailData.ChannelTrailInfo[i])] = types.YChild{"ChannelTrailInfo", &channelsTrailData.ChannelTrailInfo[i]}
    }
    channelsTrailData.EntityData.Leafs = make(map[string]types.YLeaf)
    channelsTrailData.EntityData.Leafs["channels-trail-data-type"] = types.YLeaf{"ChannelsTrailDataType", channelsTrailData.ChannelsTrailDataType}
    return &(channelsTrailData.EntityData)
}

// HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_AmpliTrailInfo
// ampli trail info
type HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_AmpliTrailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Eqpt fail. The type is interface{} with range: 0..255.
    EqptFail interface{}

    // View. The type is Trailview.
    View interface{}

    // ComPortNumber. The type is interface{} with range: 0..255.
    ComPortNumber interface{}

    // ComPortName. The type is string with length: 0..32.
    ComPortName interface{}

    // LinePortNumber. The type is interface{} with range: 0..255.
    LinePortNumber interface{}

    // LinePortName. The type is string with length: 0..32.
    LinePortName interface{}

    // BstInRxPower. The type is interface{} with range: -32768..32767.
    BstInRxPower interface{}

    // BstInRxTotalPower. The type is interface{} with range: -32768..32767.
    BstInRxTotalPower interface{}

    // BstInRxPowerThLow. The type is interface{} with range: -32768..32767.
    BstInRxPowerThLow interface{}

    // BstOutTxPower. The type is interface{} with range: -32768..32767.
    BstOutTxPower interface{}

    // BstOutTxTotalPower. The type is interface{} with range: -32768..32767.
    BstOutTxTotalPower interface{}

    // BstOutTxPowerThLow. The type is interface{} with range: -32768..32767.
    BstOutTxPowerThLow interface{}

    // BstWorkingMode. The type is interface{} with range: 0..65535.
    BstWorkingMode interface{}

    // BstSafetyMode. The type is interface{} with range: 0..65535.
    BstSafetyMode interface{}

    // BstGainRange. The type is interface{} with range: 0..65535.
    BstGainRange interface{}

    // BstOsri. The type is interface{} with range: 0..65535.
    BstOsri interface{}

    // BstChannelPower. The type is interface{} with range: -32768..32767.
    BstChannelPower interface{}

    // BstGain. The type is interface{} with range: -32768..32767.
    BstGain interface{}

    // BstTilt. The type is interface{} with range: -32768..32767.
    BstTilt interface{}

    // PreInRxPower. The type is interface{} with range: -32768..32767.
    PreInRxPower interface{}

    // PreInRxTotalPower. The type is interface{} with range: -32768..32767.
    PreInRxTotalPower interface{}

    // PreInRxPowerThLow. The type is interface{} with range: -32768..32767.
    PreInRxPowerThLow interface{}

    // PreOutTxPower. The type is interface{} with range: -32768..32767.
    PreOutTxPower interface{}

    // PreOutTxTotalPower. The type is interface{} with range: -32768..32767.
    PreOutTxTotalPower interface{}

    // PreOutTxPowerThLow. The type is interface{} with range: -32768..32767.
    PreOutTxPowerThLow interface{}

    // PreWorkingMode. The type is interface{} with range: 0..65535.
    PreWorkingMode interface{}

    // PreSafetyMode. The type is interface{} with range: 0..65535.
    PreSafetyMode interface{}

    // PreGainRange. The type is interface{} with range: 0..65535.
    PreGainRange interface{}

    // PreOsri. The type is interface{} with range: 0..65535.
    PreOsri interface{}

    // PreChannelPower. The type is interface{} with range: -32768..32767.
    PreChannelPower interface{}

    // PreGain. The type is interface{} with range: -32768..32767.
    PreGain interface{}

    // PreTilt. The type is interface{} with range: -32768..32767.
    PreTilt interface{}
}

func (ampliTrailInfo *HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_AmpliTrailInfo) GetEntityData() *types.CommonEntityData {
    ampliTrailInfo.EntityData.YFilter = ampliTrailInfo.YFilter
    ampliTrailInfo.EntityData.YangName = "ampli-trail-info"
    ampliTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    ampliTrailInfo.EntityData.ParentYangName = "channels-trail-data"
    ampliTrailInfo.EntityData.SegmentPath = "ampli-trail-info"
    ampliTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampliTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampliTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampliTrailInfo.EntityData.Children = make(map[string]types.YChild)
    ampliTrailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ampliTrailInfo.EntityData.Leafs["eqpt-fail"] = types.YLeaf{"EqptFail", ampliTrailInfo.EqptFail}
    ampliTrailInfo.EntityData.Leafs["view"] = types.YLeaf{"View", ampliTrailInfo.View}
    ampliTrailInfo.EntityData.Leafs["com-port-number"] = types.YLeaf{"ComPortNumber", ampliTrailInfo.ComPortNumber}
    ampliTrailInfo.EntityData.Leafs["com-port-name"] = types.YLeaf{"ComPortName", ampliTrailInfo.ComPortName}
    ampliTrailInfo.EntityData.Leafs["line-port-number"] = types.YLeaf{"LinePortNumber", ampliTrailInfo.LinePortNumber}
    ampliTrailInfo.EntityData.Leafs["line-port-name"] = types.YLeaf{"LinePortName", ampliTrailInfo.LinePortName}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-power"] = types.YLeaf{"BstInRxPower", ampliTrailInfo.BstInRxPower}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-total-power"] = types.YLeaf{"BstInRxTotalPower", ampliTrailInfo.BstInRxTotalPower}
    ampliTrailInfo.EntityData.Leafs["bst-in-rx-power-th-low"] = types.YLeaf{"BstInRxPowerThLow", ampliTrailInfo.BstInRxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-power"] = types.YLeaf{"BstOutTxPower", ampliTrailInfo.BstOutTxPower}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-total-power"] = types.YLeaf{"BstOutTxTotalPower", ampliTrailInfo.BstOutTxTotalPower}
    ampliTrailInfo.EntityData.Leafs["bst-out-tx-power-th-low"] = types.YLeaf{"BstOutTxPowerThLow", ampliTrailInfo.BstOutTxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["bst-working-mode"] = types.YLeaf{"BstWorkingMode", ampliTrailInfo.BstWorkingMode}
    ampliTrailInfo.EntityData.Leafs["bst-safety-mode"] = types.YLeaf{"BstSafetyMode", ampliTrailInfo.BstSafetyMode}
    ampliTrailInfo.EntityData.Leafs["bst-gain-range"] = types.YLeaf{"BstGainRange", ampliTrailInfo.BstGainRange}
    ampliTrailInfo.EntityData.Leafs["bst-osri"] = types.YLeaf{"BstOsri", ampliTrailInfo.BstOsri}
    ampliTrailInfo.EntityData.Leafs["bst-channel-power"] = types.YLeaf{"BstChannelPower", ampliTrailInfo.BstChannelPower}
    ampliTrailInfo.EntityData.Leafs["bst-gain"] = types.YLeaf{"BstGain", ampliTrailInfo.BstGain}
    ampliTrailInfo.EntityData.Leafs["bst-tilt"] = types.YLeaf{"BstTilt", ampliTrailInfo.BstTilt}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-power"] = types.YLeaf{"PreInRxPower", ampliTrailInfo.PreInRxPower}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-total-power"] = types.YLeaf{"PreInRxTotalPower", ampliTrailInfo.PreInRxTotalPower}
    ampliTrailInfo.EntityData.Leafs["pre-in-rx-power-th-low"] = types.YLeaf{"PreInRxPowerThLow", ampliTrailInfo.PreInRxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-power"] = types.YLeaf{"PreOutTxPower", ampliTrailInfo.PreOutTxPower}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-total-power"] = types.YLeaf{"PreOutTxTotalPower", ampliTrailInfo.PreOutTxTotalPower}
    ampliTrailInfo.EntityData.Leafs["pre-out-tx-power-th-low"] = types.YLeaf{"PreOutTxPowerThLow", ampliTrailInfo.PreOutTxPowerThLow}
    ampliTrailInfo.EntityData.Leafs["pre-working-mode"] = types.YLeaf{"PreWorkingMode", ampliTrailInfo.PreWorkingMode}
    ampliTrailInfo.EntityData.Leafs["pre-safety-mode"] = types.YLeaf{"PreSafetyMode", ampliTrailInfo.PreSafetyMode}
    ampliTrailInfo.EntityData.Leafs["pre-gain-range"] = types.YLeaf{"PreGainRange", ampliTrailInfo.PreGainRange}
    ampliTrailInfo.EntityData.Leafs["pre-osri"] = types.YLeaf{"PreOsri", ampliTrailInfo.PreOsri}
    ampliTrailInfo.EntityData.Leafs["pre-channel-power"] = types.YLeaf{"PreChannelPower", ampliTrailInfo.PreChannelPower}
    ampliTrailInfo.EntityData.Leafs["pre-gain"] = types.YLeaf{"PreGain", ampliTrailInfo.PreGain}
    ampliTrailInfo.EntityData.Leafs["pre-tilt"] = types.YLeaf{"PreTilt", ampliTrailInfo.PreTilt}
    return &(ampliTrailInfo.EntityData)
}

// HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_ChannelTrailInfo
// channel trail info
type HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_ChannelTrailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Eqpt fail. The type is interface{} with range: 0..255.
    EqptFail interface{}

    // Filter. The type is Chfilter.
    Filter interface{}

    // OchName. The type is string with length: 0..64.
    OchName interface{}

    // Wavelength. The type is interface{} with range: 0..4294967295.
    Wavelength interface{}

    // Frequency. The type is interface{} with range: 0..4294967295.
    Frequency interface{}

    // ComPortNumber. The type is interface{} with range: 0..255.
    ComPortNumber interface{}

    // ComPortName. The type is string with length: 0..32.
    ComPortName interface{}

    // ComRxPowerThLow. The type is interface{} with range: -32768..32767.
    ComRxPowerThLow interface{}

    // LinePortNumber. The type is interface{} with range: 0..255.
    LinePortNumber interface{}

    // LinePortName. The type is string with length: 0..32.
    LinePortName interface{}

    // LineRxPowerThLow. The type is interface{} with range: -32768..32767.
    LineRxPowerThLow interface{}

    // BstInRxPower. The type is interface{} with range: -32768..32767.
    BstInRxPower interface{}

    // BstOutTxPower. The type is interface{} with range: -32768..32767.
    BstOutTxPower interface{}

    // PreInRxPower. The type is interface{} with range: -32768..32767.
    PreInRxPower interface{}

    // PreOutTxPower. The type is interface{} with range: -32768..32767.
    PreOutTxPower interface{}
}

func (channelTrailInfo *HwModule_ChannelsSlots_ChannelsSlot_ChannelsTrailData_ChannelTrailInfo) GetEntityData() *types.CommonEntityData {
    channelTrailInfo.EntityData.YFilter = channelTrailInfo.YFilter
    channelTrailInfo.EntityData.YangName = "channel-trail-info"
    channelTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    channelTrailInfo.EntityData.ParentYangName = "channels-trail-data"
    channelTrailInfo.EntityData.SegmentPath = "channel-trail-info"
    channelTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelTrailInfo.EntityData.Children = make(map[string]types.YChild)
    channelTrailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    channelTrailInfo.EntityData.Leafs["eqpt-fail"] = types.YLeaf{"EqptFail", channelTrailInfo.EqptFail}
    channelTrailInfo.EntityData.Leafs["filter"] = types.YLeaf{"Filter", channelTrailInfo.Filter}
    channelTrailInfo.EntityData.Leafs["och-name"] = types.YLeaf{"OchName", channelTrailInfo.OchName}
    channelTrailInfo.EntityData.Leafs["wavelength"] = types.YLeaf{"Wavelength", channelTrailInfo.Wavelength}
    channelTrailInfo.EntityData.Leafs["frequency"] = types.YLeaf{"Frequency", channelTrailInfo.Frequency}
    channelTrailInfo.EntityData.Leafs["com-port-number"] = types.YLeaf{"ComPortNumber", channelTrailInfo.ComPortNumber}
    channelTrailInfo.EntityData.Leafs["com-port-name"] = types.YLeaf{"ComPortName", channelTrailInfo.ComPortName}
    channelTrailInfo.EntityData.Leafs["com-rx-power-th-low"] = types.YLeaf{"ComRxPowerThLow", channelTrailInfo.ComRxPowerThLow}
    channelTrailInfo.EntityData.Leafs["line-port-number"] = types.YLeaf{"LinePortNumber", channelTrailInfo.LinePortNumber}
    channelTrailInfo.EntityData.Leafs["line-port-name"] = types.YLeaf{"LinePortName", channelTrailInfo.LinePortName}
    channelTrailInfo.EntityData.Leafs["line-rx-power-th-low"] = types.YLeaf{"LineRxPowerThLow", channelTrailInfo.LineRxPowerThLow}
    channelTrailInfo.EntityData.Leafs["bst-in-rx-power"] = types.YLeaf{"BstInRxPower", channelTrailInfo.BstInRxPower}
    channelTrailInfo.EntityData.Leafs["bst-out-tx-power"] = types.YLeaf{"BstOutTxPower", channelTrailInfo.BstOutTxPower}
    channelTrailInfo.EntityData.Leafs["pre-in-rx-power"] = types.YLeaf{"PreInRxPower", channelTrailInfo.PreInRxPower}
    channelTrailInfo.EntityData.Leafs["pre-out-tx-power"] = types.YLeaf{"PreOutTxPower", channelTrailInfo.PreOutTxPower}
    return &(channelTrailInfo.EntityData)
}

