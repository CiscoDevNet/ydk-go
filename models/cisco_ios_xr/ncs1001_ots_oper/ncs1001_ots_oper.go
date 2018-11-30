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

// OtsChannelsTrailData represents Ots channels trail data
type OtsChannelsTrailData string

const (
    // Displays all channels trail data
    OtsChannelsTrailData_all OtsChannelsTrailData = "all"

    // Displays active channels trail data
    OtsChannelsTrailData_active OtsChannelsTrailData = "active"
)

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

// Chfilter represents Chfilter
type Chfilter string

const (
    // Active
    Chfilter_ch_filter_active Chfilter = "ch-filter-active"

    // All
    Chfilter_ch_filter_all Chfilter = "ch-filter-all"
)

// HwModule
// ncs1001 hw-module command chain
type HwModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amplifier trail data information.
    AmplifierTrails HwModule_AmplifierTrails

    // Channels trail data information.
    ChannelsTrails HwModule_ChannelsTrails
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

    hwModule.EntityData.Children = types.NewOrderedMap()
    hwModule.EntityData.Children.Append("amplifier-trails", types.YChild{"AmplifierTrails", &hwModule.AmplifierTrails})
    hwModule.EntityData.Children.Append("channels-trails", types.YChild{"ChannelsTrails", &hwModule.ChannelsTrails})
    hwModule.EntityData.Leafs = types.NewOrderedMap()

    hwModule.EntityData.YListKeys = []string {}

    return &(hwModule.EntityData)
}

// HwModule_AmplifierTrails
// Amplifier trail data information
type HwModule_AmplifierTrails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot Id. The type is slice of HwModule_AmplifierTrails_AmplifierTrail.
    AmplifierTrail []*HwModule_AmplifierTrails_AmplifierTrail
}

func (amplifierTrails *HwModule_AmplifierTrails) GetEntityData() *types.CommonEntityData {
    amplifierTrails.EntityData.YFilter = amplifierTrails.YFilter
    amplifierTrails.EntityData.YangName = "amplifier-trails"
    amplifierTrails.EntityData.BundleName = "cisco_ios_xr"
    amplifierTrails.EntityData.ParentYangName = "hw-module"
    amplifierTrails.EntityData.SegmentPath = "amplifier-trails"
    amplifierTrails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifierTrails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifierTrails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifierTrails.EntityData.Children = types.NewOrderedMap()
    amplifierTrails.EntityData.Children.Append("amplifier-trail", types.YChild{"AmplifierTrail", nil})
    for i := range amplifierTrails.AmplifierTrail {
        amplifierTrails.EntityData.Children.Append(types.GetSegmentPath(amplifierTrails.AmplifierTrail[i]), types.YChild{"AmplifierTrail", amplifierTrails.AmplifierTrail[i]})
    }
    amplifierTrails.EntityData.Leafs = types.NewOrderedMap()

    amplifierTrails.EntityData.YListKeys = []string {}

    return &(amplifierTrails.EntityData)
}

// HwModule_AmplifierTrails_AmplifierTrail
// Slot Id
type HwModule_AmplifierTrails_AmplifierTrail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: 0..4294967295.
    SlotId interface{}

    // Trail data information. The type is slice of
    // HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData.
    AmplifierTrailData []*HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData
}

func (amplifierTrail *HwModule_AmplifierTrails_AmplifierTrail) GetEntityData() *types.CommonEntityData {
    amplifierTrail.EntityData.YFilter = amplifierTrail.YFilter
    amplifierTrail.EntityData.YangName = "amplifier-trail"
    amplifierTrail.EntityData.BundleName = "cisco_ios_xr"
    amplifierTrail.EntityData.ParentYangName = "amplifier-trails"
    amplifierTrail.EntityData.SegmentPath = "amplifier-trail" + types.AddKeyToken(amplifierTrail.SlotId, "slot-id")
    amplifierTrail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifierTrail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifierTrail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifierTrail.EntityData.Children = types.NewOrderedMap()
    amplifierTrail.EntityData.Children.Append("amplifier-trail-data", types.YChild{"AmplifierTrailData", nil})
    for i := range amplifierTrail.AmplifierTrailData {
        amplifierTrail.EntityData.Children.Append(types.GetSegmentPath(amplifierTrail.AmplifierTrailData[i]), types.YChild{"AmplifierTrailData", amplifierTrail.AmplifierTrailData[i]})
    }
    amplifierTrail.EntityData.Leafs = types.NewOrderedMap()
    amplifierTrail.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", amplifierTrail.SlotId})

    amplifierTrail.EntityData.YListKeys = []string {"SlotId"}

    return &(amplifierTrail.EntityData)
}

// HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData
// Trail data information
type HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Select trail data. The type is OtsAmpliTrailData.
    AmplifierTrailDataType interface{}

    // ampli trail info. The type is slice of
    // HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_AmpliTrailInfo.
    AmpliTrailInfo []*HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_AmpliTrailInfo

    // channel trail info. The type is slice of
    // HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_ChannelTrailInfo.
    ChannelTrailInfo []*HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_ChannelTrailInfo
}

func (amplifierTrailData *HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData) GetEntityData() *types.CommonEntityData {
    amplifierTrailData.EntityData.YFilter = amplifierTrailData.YFilter
    amplifierTrailData.EntityData.YangName = "amplifier-trail-data"
    amplifierTrailData.EntityData.BundleName = "cisco_ios_xr"
    amplifierTrailData.EntityData.ParentYangName = "amplifier-trail"
    amplifierTrailData.EntityData.SegmentPath = "amplifier-trail-data" + types.AddKeyToken(amplifierTrailData.AmplifierTrailDataType, "amplifier-trail-data-type")
    amplifierTrailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    amplifierTrailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    amplifierTrailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    amplifierTrailData.EntityData.Children = types.NewOrderedMap()
    amplifierTrailData.EntityData.Children.Append("ampli-trail-info", types.YChild{"AmpliTrailInfo", nil})
    for i := range amplifierTrailData.AmpliTrailInfo {
        amplifierTrailData.EntityData.Children.Append(types.GetSegmentPath(amplifierTrailData.AmpliTrailInfo[i]), types.YChild{"AmpliTrailInfo", amplifierTrailData.AmpliTrailInfo[i]})
    }
    amplifierTrailData.EntityData.Children.Append("channel-trail-info", types.YChild{"ChannelTrailInfo", nil})
    for i := range amplifierTrailData.ChannelTrailInfo {
        amplifierTrailData.EntityData.Children.Append(types.GetSegmentPath(amplifierTrailData.ChannelTrailInfo[i]), types.YChild{"ChannelTrailInfo", amplifierTrailData.ChannelTrailInfo[i]})
    }
    amplifierTrailData.EntityData.Leafs = types.NewOrderedMap()
    amplifierTrailData.EntityData.Leafs.Append("amplifier-trail-data-type", types.YLeaf{"AmplifierTrailDataType", amplifierTrailData.AmplifierTrailDataType})

    amplifierTrailData.EntityData.YListKeys = []string {"AmplifierTrailDataType"}

    return &(amplifierTrailData.EntityData)
}

// HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_AmpliTrailInfo
// ampli trail info
type HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_AmpliTrailInfo struct {
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

func (ampliTrailInfo *HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_AmpliTrailInfo) GetEntityData() *types.CommonEntityData {
    ampliTrailInfo.EntityData.YFilter = ampliTrailInfo.YFilter
    ampliTrailInfo.EntityData.YangName = "ampli-trail-info"
    ampliTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    ampliTrailInfo.EntityData.ParentYangName = "amplifier-trail-data"
    ampliTrailInfo.EntityData.SegmentPath = "ampli-trail-info"
    ampliTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampliTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampliTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampliTrailInfo.EntityData.Children = types.NewOrderedMap()
    ampliTrailInfo.EntityData.Leafs = types.NewOrderedMap()
    ampliTrailInfo.EntityData.Leafs.Append("eqpt-fail", types.YLeaf{"EqptFail", ampliTrailInfo.EqptFail})
    ampliTrailInfo.EntityData.Leafs.Append("view", types.YLeaf{"View", ampliTrailInfo.View})
    ampliTrailInfo.EntityData.Leafs.Append("com-port-number", types.YLeaf{"ComPortNumber", ampliTrailInfo.ComPortNumber})
    ampliTrailInfo.EntityData.Leafs.Append("com-port-name", types.YLeaf{"ComPortName", ampliTrailInfo.ComPortName})
    ampliTrailInfo.EntityData.Leafs.Append("line-port-number", types.YLeaf{"LinePortNumber", ampliTrailInfo.LinePortNumber})
    ampliTrailInfo.EntityData.Leafs.Append("line-port-name", types.YLeaf{"LinePortName", ampliTrailInfo.LinePortName})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-power", types.YLeaf{"BstInRxPower", ampliTrailInfo.BstInRxPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-total-power", types.YLeaf{"BstInRxTotalPower", ampliTrailInfo.BstInRxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-power-th-low", types.YLeaf{"BstInRxPowerThLow", ampliTrailInfo.BstInRxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-power", types.YLeaf{"BstOutTxPower", ampliTrailInfo.BstOutTxPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-total-power", types.YLeaf{"BstOutTxTotalPower", ampliTrailInfo.BstOutTxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-power-th-low", types.YLeaf{"BstOutTxPowerThLow", ampliTrailInfo.BstOutTxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("bst-working-mode", types.YLeaf{"BstWorkingMode", ampliTrailInfo.BstWorkingMode})
    ampliTrailInfo.EntityData.Leafs.Append("bst-safety-mode", types.YLeaf{"BstSafetyMode", ampliTrailInfo.BstSafetyMode})
    ampliTrailInfo.EntityData.Leafs.Append("bst-gain-range", types.YLeaf{"BstGainRange", ampliTrailInfo.BstGainRange})
    ampliTrailInfo.EntityData.Leafs.Append("bst-osri", types.YLeaf{"BstOsri", ampliTrailInfo.BstOsri})
    ampliTrailInfo.EntityData.Leafs.Append("bst-channel-power", types.YLeaf{"BstChannelPower", ampliTrailInfo.BstChannelPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-gain", types.YLeaf{"BstGain", ampliTrailInfo.BstGain})
    ampliTrailInfo.EntityData.Leafs.Append("bst-tilt", types.YLeaf{"BstTilt", ampliTrailInfo.BstTilt})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-power", types.YLeaf{"PreInRxPower", ampliTrailInfo.PreInRxPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-total-power", types.YLeaf{"PreInRxTotalPower", ampliTrailInfo.PreInRxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-power-th-low", types.YLeaf{"PreInRxPowerThLow", ampliTrailInfo.PreInRxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-power", types.YLeaf{"PreOutTxPower", ampliTrailInfo.PreOutTxPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-total-power", types.YLeaf{"PreOutTxTotalPower", ampliTrailInfo.PreOutTxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-power-th-low", types.YLeaf{"PreOutTxPowerThLow", ampliTrailInfo.PreOutTxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("pre-working-mode", types.YLeaf{"PreWorkingMode", ampliTrailInfo.PreWorkingMode})
    ampliTrailInfo.EntityData.Leafs.Append("pre-safety-mode", types.YLeaf{"PreSafetyMode", ampliTrailInfo.PreSafetyMode})
    ampliTrailInfo.EntityData.Leafs.Append("pre-gain-range", types.YLeaf{"PreGainRange", ampliTrailInfo.PreGainRange})
    ampliTrailInfo.EntityData.Leafs.Append("pre-osri", types.YLeaf{"PreOsri", ampliTrailInfo.PreOsri})
    ampliTrailInfo.EntityData.Leafs.Append("pre-channel-power", types.YLeaf{"PreChannelPower", ampliTrailInfo.PreChannelPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-gain", types.YLeaf{"PreGain", ampliTrailInfo.PreGain})
    ampliTrailInfo.EntityData.Leafs.Append("pre-tilt", types.YLeaf{"PreTilt", ampliTrailInfo.PreTilt})

    ampliTrailInfo.EntityData.YListKeys = []string {}

    return &(ampliTrailInfo.EntityData)
}

// HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_ChannelTrailInfo
// channel trail info
type HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_ChannelTrailInfo struct {
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

func (channelTrailInfo *HwModule_AmplifierTrails_AmplifierTrail_AmplifierTrailData_ChannelTrailInfo) GetEntityData() *types.CommonEntityData {
    channelTrailInfo.EntityData.YFilter = channelTrailInfo.YFilter
    channelTrailInfo.EntityData.YangName = "channel-trail-info"
    channelTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    channelTrailInfo.EntityData.ParentYangName = "amplifier-trail-data"
    channelTrailInfo.EntityData.SegmentPath = "channel-trail-info"
    channelTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelTrailInfo.EntityData.Children = types.NewOrderedMap()
    channelTrailInfo.EntityData.Leafs = types.NewOrderedMap()
    channelTrailInfo.EntityData.Leafs.Append("eqpt-fail", types.YLeaf{"EqptFail", channelTrailInfo.EqptFail})
    channelTrailInfo.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", channelTrailInfo.Filter})
    channelTrailInfo.EntityData.Leafs.Append("och-name", types.YLeaf{"OchName", channelTrailInfo.OchName})
    channelTrailInfo.EntityData.Leafs.Append("wavelength", types.YLeaf{"Wavelength", channelTrailInfo.Wavelength})
    channelTrailInfo.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", channelTrailInfo.Frequency})
    channelTrailInfo.EntityData.Leafs.Append("com-port-number", types.YLeaf{"ComPortNumber", channelTrailInfo.ComPortNumber})
    channelTrailInfo.EntityData.Leafs.Append("com-port-name", types.YLeaf{"ComPortName", channelTrailInfo.ComPortName})
    channelTrailInfo.EntityData.Leafs.Append("com-rx-power-th-low", types.YLeaf{"ComRxPowerThLow", channelTrailInfo.ComRxPowerThLow})
    channelTrailInfo.EntityData.Leafs.Append("line-port-number", types.YLeaf{"LinePortNumber", channelTrailInfo.LinePortNumber})
    channelTrailInfo.EntityData.Leafs.Append("line-port-name", types.YLeaf{"LinePortName", channelTrailInfo.LinePortName})
    channelTrailInfo.EntityData.Leafs.Append("line-rx-power-th-low", types.YLeaf{"LineRxPowerThLow", channelTrailInfo.LineRxPowerThLow})
    channelTrailInfo.EntityData.Leafs.Append("bst-in-rx-power", types.YLeaf{"BstInRxPower", channelTrailInfo.BstInRxPower})
    channelTrailInfo.EntityData.Leafs.Append("bst-out-tx-power", types.YLeaf{"BstOutTxPower", channelTrailInfo.BstOutTxPower})
    channelTrailInfo.EntityData.Leafs.Append("pre-in-rx-power", types.YLeaf{"PreInRxPower", channelTrailInfo.PreInRxPower})
    channelTrailInfo.EntityData.Leafs.Append("pre-out-tx-power", types.YLeaf{"PreOutTxPower", channelTrailInfo.PreOutTxPower})

    channelTrailInfo.EntityData.YListKeys = []string {}

    return &(channelTrailInfo.EntityData)
}

// HwModule_ChannelsTrails
// Channels trail data information
type HwModule_ChannelsTrails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot Id. The type is slice of HwModule_ChannelsTrails_ChannelsTrail.
    ChannelsTrail []*HwModule_ChannelsTrails_ChannelsTrail
}

func (channelsTrails *HwModule_ChannelsTrails) GetEntityData() *types.CommonEntityData {
    channelsTrails.EntityData.YFilter = channelsTrails.YFilter
    channelsTrails.EntityData.YangName = "channels-trails"
    channelsTrails.EntityData.BundleName = "cisco_ios_xr"
    channelsTrails.EntityData.ParentYangName = "hw-module"
    channelsTrails.EntityData.SegmentPath = "channels-trails"
    channelsTrails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsTrails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsTrails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsTrails.EntityData.Children = types.NewOrderedMap()
    channelsTrails.EntityData.Children.Append("channels-trail", types.YChild{"ChannelsTrail", nil})
    for i := range channelsTrails.ChannelsTrail {
        channelsTrails.EntityData.Children.Append(types.GetSegmentPath(channelsTrails.ChannelsTrail[i]), types.YChild{"ChannelsTrail", channelsTrails.ChannelsTrail[i]})
    }
    channelsTrails.EntityData.Leafs = types.NewOrderedMap()

    channelsTrails.EntityData.YListKeys = []string {}

    return &(channelsTrails.EntityData)
}

// HwModule_ChannelsTrails_ChannelsTrail
// Slot Id
type HwModule_ChannelsTrails_ChannelsTrail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: 1..3.
    SlotId interface{}

    // Trail data information. The type is slice of
    // HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData.
    ChannelsTrailData []*HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData
}

func (channelsTrail *HwModule_ChannelsTrails_ChannelsTrail) GetEntityData() *types.CommonEntityData {
    channelsTrail.EntityData.YFilter = channelsTrail.YFilter
    channelsTrail.EntityData.YangName = "channels-trail"
    channelsTrail.EntityData.BundleName = "cisco_ios_xr"
    channelsTrail.EntityData.ParentYangName = "channels-trails"
    channelsTrail.EntityData.SegmentPath = "channels-trail" + types.AddKeyToken(channelsTrail.SlotId, "slot-id")
    channelsTrail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsTrail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsTrail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsTrail.EntityData.Children = types.NewOrderedMap()
    channelsTrail.EntityData.Children.Append("channels-trail-data", types.YChild{"ChannelsTrailData", nil})
    for i := range channelsTrail.ChannelsTrailData {
        channelsTrail.EntityData.Children.Append(types.GetSegmentPath(channelsTrail.ChannelsTrailData[i]), types.YChild{"ChannelsTrailData", channelsTrail.ChannelsTrailData[i]})
    }
    channelsTrail.EntityData.Leafs = types.NewOrderedMap()
    channelsTrail.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", channelsTrail.SlotId})

    channelsTrail.EntityData.YListKeys = []string {"SlotId"}

    return &(channelsTrail.EntityData)
}

// HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData
// Trail data information
type HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Select trail data. The type is
    // OtsChannelsTrailData.
    ChannelsTrailDataType interface{}

    // ampli trail info. The type is slice of
    // HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_AmpliTrailInfo.
    AmpliTrailInfo []*HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_AmpliTrailInfo

    // channel trail info. The type is slice of
    // HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_ChannelTrailInfo.
    ChannelTrailInfo []*HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_ChannelTrailInfo
}

func (channelsTrailData *HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData) GetEntityData() *types.CommonEntityData {
    channelsTrailData.EntityData.YFilter = channelsTrailData.YFilter
    channelsTrailData.EntityData.YangName = "channels-trail-data"
    channelsTrailData.EntityData.BundleName = "cisco_ios_xr"
    channelsTrailData.EntityData.ParentYangName = "channels-trail"
    channelsTrailData.EntityData.SegmentPath = "channels-trail-data" + types.AddKeyToken(channelsTrailData.ChannelsTrailDataType, "channels-trail-data-type")
    channelsTrailData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelsTrailData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelsTrailData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelsTrailData.EntityData.Children = types.NewOrderedMap()
    channelsTrailData.EntityData.Children.Append("ampli-trail-info", types.YChild{"AmpliTrailInfo", nil})
    for i := range channelsTrailData.AmpliTrailInfo {
        channelsTrailData.EntityData.Children.Append(types.GetSegmentPath(channelsTrailData.AmpliTrailInfo[i]), types.YChild{"AmpliTrailInfo", channelsTrailData.AmpliTrailInfo[i]})
    }
    channelsTrailData.EntityData.Children.Append("channel-trail-info", types.YChild{"ChannelTrailInfo", nil})
    for i := range channelsTrailData.ChannelTrailInfo {
        channelsTrailData.EntityData.Children.Append(types.GetSegmentPath(channelsTrailData.ChannelTrailInfo[i]), types.YChild{"ChannelTrailInfo", channelsTrailData.ChannelTrailInfo[i]})
    }
    channelsTrailData.EntityData.Leafs = types.NewOrderedMap()
    channelsTrailData.EntityData.Leafs.Append("channels-trail-data-type", types.YLeaf{"ChannelsTrailDataType", channelsTrailData.ChannelsTrailDataType})

    channelsTrailData.EntityData.YListKeys = []string {"ChannelsTrailDataType"}

    return &(channelsTrailData.EntityData)
}

// HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_AmpliTrailInfo
// ampli trail info
type HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_AmpliTrailInfo struct {
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

func (ampliTrailInfo *HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_AmpliTrailInfo) GetEntityData() *types.CommonEntityData {
    ampliTrailInfo.EntityData.YFilter = ampliTrailInfo.YFilter
    ampliTrailInfo.EntityData.YangName = "ampli-trail-info"
    ampliTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    ampliTrailInfo.EntityData.ParentYangName = "channels-trail-data"
    ampliTrailInfo.EntityData.SegmentPath = "ampli-trail-info"
    ampliTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ampliTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ampliTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ampliTrailInfo.EntityData.Children = types.NewOrderedMap()
    ampliTrailInfo.EntityData.Leafs = types.NewOrderedMap()
    ampliTrailInfo.EntityData.Leafs.Append("eqpt-fail", types.YLeaf{"EqptFail", ampliTrailInfo.EqptFail})
    ampliTrailInfo.EntityData.Leafs.Append("view", types.YLeaf{"View", ampliTrailInfo.View})
    ampliTrailInfo.EntityData.Leafs.Append("com-port-number", types.YLeaf{"ComPortNumber", ampliTrailInfo.ComPortNumber})
    ampliTrailInfo.EntityData.Leafs.Append("com-port-name", types.YLeaf{"ComPortName", ampliTrailInfo.ComPortName})
    ampliTrailInfo.EntityData.Leafs.Append("line-port-number", types.YLeaf{"LinePortNumber", ampliTrailInfo.LinePortNumber})
    ampliTrailInfo.EntityData.Leafs.Append("line-port-name", types.YLeaf{"LinePortName", ampliTrailInfo.LinePortName})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-power", types.YLeaf{"BstInRxPower", ampliTrailInfo.BstInRxPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-total-power", types.YLeaf{"BstInRxTotalPower", ampliTrailInfo.BstInRxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-in-rx-power-th-low", types.YLeaf{"BstInRxPowerThLow", ampliTrailInfo.BstInRxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-power", types.YLeaf{"BstOutTxPower", ampliTrailInfo.BstOutTxPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-total-power", types.YLeaf{"BstOutTxTotalPower", ampliTrailInfo.BstOutTxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-out-tx-power-th-low", types.YLeaf{"BstOutTxPowerThLow", ampliTrailInfo.BstOutTxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("bst-working-mode", types.YLeaf{"BstWorkingMode", ampliTrailInfo.BstWorkingMode})
    ampliTrailInfo.EntityData.Leafs.Append("bst-safety-mode", types.YLeaf{"BstSafetyMode", ampliTrailInfo.BstSafetyMode})
    ampliTrailInfo.EntityData.Leafs.Append("bst-gain-range", types.YLeaf{"BstGainRange", ampliTrailInfo.BstGainRange})
    ampliTrailInfo.EntityData.Leafs.Append("bst-osri", types.YLeaf{"BstOsri", ampliTrailInfo.BstOsri})
    ampliTrailInfo.EntityData.Leafs.Append("bst-channel-power", types.YLeaf{"BstChannelPower", ampliTrailInfo.BstChannelPower})
    ampliTrailInfo.EntityData.Leafs.Append("bst-gain", types.YLeaf{"BstGain", ampliTrailInfo.BstGain})
    ampliTrailInfo.EntityData.Leafs.Append("bst-tilt", types.YLeaf{"BstTilt", ampliTrailInfo.BstTilt})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-power", types.YLeaf{"PreInRxPower", ampliTrailInfo.PreInRxPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-total-power", types.YLeaf{"PreInRxTotalPower", ampliTrailInfo.PreInRxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-in-rx-power-th-low", types.YLeaf{"PreInRxPowerThLow", ampliTrailInfo.PreInRxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-power", types.YLeaf{"PreOutTxPower", ampliTrailInfo.PreOutTxPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-total-power", types.YLeaf{"PreOutTxTotalPower", ampliTrailInfo.PreOutTxTotalPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-out-tx-power-th-low", types.YLeaf{"PreOutTxPowerThLow", ampliTrailInfo.PreOutTxPowerThLow})
    ampliTrailInfo.EntityData.Leafs.Append("pre-working-mode", types.YLeaf{"PreWorkingMode", ampliTrailInfo.PreWorkingMode})
    ampliTrailInfo.EntityData.Leafs.Append("pre-safety-mode", types.YLeaf{"PreSafetyMode", ampliTrailInfo.PreSafetyMode})
    ampliTrailInfo.EntityData.Leafs.Append("pre-gain-range", types.YLeaf{"PreGainRange", ampliTrailInfo.PreGainRange})
    ampliTrailInfo.EntityData.Leafs.Append("pre-osri", types.YLeaf{"PreOsri", ampliTrailInfo.PreOsri})
    ampliTrailInfo.EntityData.Leafs.Append("pre-channel-power", types.YLeaf{"PreChannelPower", ampliTrailInfo.PreChannelPower})
    ampliTrailInfo.EntityData.Leafs.Append("pre-gain", types.YLeaf{"PreGain", ampliTrailInfo.PreGain})
    ampliTrailInfo.EntityData.Leafs.Append("pre-tilt", types.YLeaf{"PreTilt", ampliTrailInfo.PreTilt})

    ampliTrailInfo.EntityData.YListKeys = []string {}

    return &(ampliTrailInfo.EntityData)
}

// HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_ChannelTrailInfo
// channel trail info
type HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_ChannelTrailInfo struct {
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

func (channelTrailInfo *HwModule_ChannelsTrails_ChannelsTrail_ChannelsTrailData_ChannelTrailInfo) GetEntityData() *types.CommonEntityData {
    channelTrailInfo.EntityData.YFilter = channelTrailInfo.YFilter
    channelTrailInfo.EntityData.YangName = "channel-trail-info"
    channelTrailInfo.EntityData.BundleName = "cisco_ios_xr"
    channelTrailInfo.EntityData.ParentYangName = "channels-trail-data"
    channelTrailInfo.EntityData.SegmentPath = "channel-trail-info"
    channelTrailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    channelTrailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    channelTrailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    channelTrailInfo.EntityData.Children = types.NewOrderedMap()
    channelTrailInfo.EntityData.Leafs = types.NewOrderedMap()
    channelTrailInfo.EntityData.Leafs.Append("eqpt-fail", types.YLeaf{"EqptFail", channelTrailInfo.EqptFail})
    channelTrailInfo.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", channelTrailInfo.Filter})
    channelTrailInfo.EntityData.Leafs.Append("och-name", types.YLeaf{"OchName", channelTrailInfo.OchName})
    channelTrailInfo.EntityData.Leafs.Append("wavelength", types.YLeaf{"Wavelength", channelTrailInfo.Wavelength})
    channelTrailInfo.EntityData.Leafs.Append("frequency", types.YLeaf{"Frequency", channelTrailInfo.Frequency})
    channelTrailInfo.EntityData.Leafs.Append("com-port-number", types.YLeaf{"ComPortNumber", channelTrailInfo.ComPortNumber})
    channelTrailInfo.EntityData.Leafs.Append("com-port-name", types.YLeaf{"ComPortName", channelTrailInfo.ComPortName})
    channelTrailInfo.EntityData.Leafs.Append("com-rx-power-th-low", types.YLeaf{"ComRxPowerThLow", channelTrailInfo.ComRxPowerThLow})
    channelTrailInfo.EntityData.Leafs.Append("line-port-number", types.YLeaf{"LinePortNumber", channelTrailInfo.LinePortNumber})
    channelTrailInfo.EntityData.Leafs.Append("line-port-name", types.YLeaf{"LinePortName", channelTrailInfo.LinePortName})
    channelTrailInfo.EntityData.Leafs.Append("line-rx-power-th-low", types.YLeaf{"LineRxPowerThLow", channelTrailInfo.LineRxPowerThLow})
    channelTrailInfo.EntityData.Leafs.Append("bst-in-rx-power", types.YLeaf{"BstInRxPower", channelTrailInfo.BstInRxPower})
    channelTrailInfo.EntityData.Leafs.Append("bst-out-tx-power", types.YLeaf{"BstOutTxPower", channelTrailInfo.BstOutTxPower})
    channelTrailInfo.EntityData.Leafs.Append("pre-in-rx-power", types.YLeaf{"PreInRxPower", channelTrailInfo.PreInRxPower})
    channelTrailInfo.EntityData.Leafs.Append("pre-out-tx-power", types.YLeaf{"PreOutTxPower", channelTrailInfo.PreOutTxPower})

    channelTrailInfo.EntityData.YListKeys = []string {}

    return &(channelTrailInfo.EntityData)
}

