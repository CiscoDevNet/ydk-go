// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-ptp-pd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-ptp: PTP PD operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_ptp_pd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_ptp_pd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-ptp-pd-oper platform-ptp}", reflect.TypeOf(PlatformPtp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-ptp-pd-oper:platform-ptp", reflect.TypeOf(PlatformPtp{}))
}

// PlatformPtp
// PTP PD operational data
type PlatformPtp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PTP PD Servo information.
    PlatformPtpServo PlatformPtp_PlatformPtpServo
}

func (platformPtp *PlatformPtp) GetEntityData() *types.CommonEntityData {
    platformPtp.EntityData.YFilter = platformPtp.YFilter
    platformPtp.EntityData.YangName = "platform-ptp"
    platformPtp.EntityData.BundleName = "cisco_ios_xr"
    platformPtp.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-ptp-pd-oper"
    platformPtp.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-ptp-pd-oper:platform-ptp"
    platformPtp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformPtp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformPtp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformPtp.EntityData.Children = types.NewOrderedMap()
    platformPtp.EntityData.Children.Append("platform-ptp-servo", types.YChild{"PlatformPtpServo", &platformPtp.PlatformPtpServo})
    platformPtp.EntityData.Leafs = types.NewOrderedMap()

    platformPtp.EntityData.YListKeys = []string {}

    return &(platformPtp.EntityData)
}

// PlatformPtp_PlatformPtpServo
// PTP PD Servo information
type PlatformPtp_PlatformPtpServo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lock status of device. The type is interface{} with range: 0..65535.
    LockStatus interface{}

    // running status of apr. The type is bool.
    Running interface{}

    // status of device. The type is string with length: 0..50.
    DeviceStatus interface{}

    // log level of apr. The type is interface{} with range: 0..65535.
    LogLevel interface{}

    // last phase alignment accuracy. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    PhaseAccuracyLast interface{}

    // number of sync timestamp reveiced. The type is interface{} with range:
    // 0..4294967295.
    NumSyncTimestamp interface{}

    // number of delay timestamp reveiced. The type is interface{} with range:
    // 0..4294967295.
    NumDelayTimestamp interface{}

    // number of setTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumSetTime interface{}

    // number of stepTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumStepTime interface{}

    // number of adjustFreq() been called. The type is interface{} with range:
    // 0..4294967295.
    NumAdjustFreq interface{}

    // number of adjustFreqTime() been called. The type is interface{} with range:
    // 0..4294967295.
    NumAdjustFreqTime interface{}

    // last input of adjustFreq. The type is interface{} with range:
    // -2147483648..2147483647.
    LastAdjustFreq interface{}

    // last input of stepTime. The type is interface{} with range:
    // -2147483648..2147483647.
    LastStepTime interface{}

    // number of sync timestamp discarded. The type is interface{} with range:
    // 0..4294967295.
    NumDiscardSyncTimestamp interface{}

    // number of delay timestamp discarded. The type is interface{} with range:
    // 0..4294967295.
    NumDiscardDelayTimestamp interface{}

    // last input flag of setTime. The type is bool.
    FlagofLastSetTime interface{}

    // Time Offset From Master. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    OffsetFromMaster interface{}

    // Mean Path Delay. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    MeanPathDelay interface{}

    // last input of setTime.
    LastSetTime PlatformPtp_PlatformPtpServo_LastSetTime

    // last T1 timestamp reveiced.
    LastReceivedT1 PlatformPtp_PlatformPtpServo_LastReceivedT1

    // last T2 timestamp reveiced.
    LastReceivedT2 PlatformPtp_PlatformPtpServo_LastReceivedT2

    // last T3 timestamp reveiced.
    LastReceivedT3 PlatformPtp_PlatformPtpServo_LastReceivedT3

    // last T4 timestamp reveiced.
    LastReceivedT4 PlatformPtp_PlatformPtpServo_LastReceivedT4

    // pre T1 timestamp reveiced.
    PreReceivedT1 PlatformPtp_PlatformPtpServo_PreReceivedT1

    // pre T2 timestamp reveiced.
    PreReceivedT2 PlatformPtp_PlatformPtpServo_PreReceivedT2

    // pre T3 timestamp reveiced.
    PreReceivedT3 PlatformPtp_PlatformPtpServo_PreReceivedT3

    // pre T4 timestamp reveiced.
    PreReceivedT4 PlatformPtp_PlatformPtpServo_PreReceivedT4
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetEntityData() *types.CommonEntityData {
    platformPtpServo.EntityData.YFilter = platformPtpServo.YFilter
    platformPtpServo.EntityData.YangName = "platform-ptp-servo"
    platformPtpServo.EntityData.BundleName = "cisco_ios_xr"
    platformPtpServo.EntityData.ParentYangName = "platform-ptp"
    platformPtpServo.EntityData.SegmentPath = "platform-ptp-servo"
    platformPtpServo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformPtpServo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformPtpServo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformPtpServo.EntityData.Children = types.NewOrderedMap()
    platformPtpServo.EntityData.Children.Append("last-set-time", types.YChild{"LastSetTime", &platformPtpServo.LastSetTime})
    platformPtpServo.EntityData.Children.Append("last-received-t1", types.YChild{"LastReceivedT1", &platformPtpServo.LastReceivedT1})
    platformPtpServo.EntityData.Children.Append("last-received-t2", types.YChild{"LastReceivedT2", &platformPtpServo.LastReceivedT2})
    platformPtpServo.EntityData.Children.Append("last-received-t3", types.YChild{"LastReceivedT3", &platformPtpServo.LastReceivedT3})
    platformPtpServo.EntityData.Children.Append("last-received-t4", types.YChild{"LastReceivedT4", &platformPtpServo.LastReceivedT4})
    platformPtpServo.EntityData.Children.Append("pre-received-t1", types.YChild{"PreReceivedT1", &platformPtpServo.PreReceivedT1})
    platformPtpServo.EntityData.Children.Append("pre-received-t2", types.YChild{"PreReceivedT2", &platformPtpServo.PreReceivedT2})
    platformPtpServo.EntityData.Children.Append("pre-received-t3", types.YChild{"PreReceivedT3", &platformPtpServo.PreReceivedT3})
    platformPtpServo.EntityData.Children.Append("pre-received-t4", types.YChild{"PreReceivedT4", &platformPtpServo.PreReceivedT4})
    platformPtpServo.EntityData.Leafs = types.NewOrderedMap()
    platformPtpServo.EntityData.Leafs.Append("lock-status", types.YLeaf{"LockStatus", platformPtpServo.LockStatus})
    platformPtpServo.EntityData.Leafs.Append("running", types.YLeaf{"Running", platformPtpServo.Running})
    platformPtpServo.EntityData.Leafs.Append("device-status", types.YLeaf{"DeviceStatus", platformPtpServo.DeviceStatus})
    platformPtpServo.EntityData.Leafs.Append("log-level", types.YLeaf{"LogLevel", platformPtpServo.LogLevel})
    platformPtpServo.EntityData.Leafs.Append("phase-accuracy-last", types.YLeaf{"PhaseAccuracyLast", platformPtpServo.PhaseAccuracyLast})
    platformPtpServo.EntityData.Leafs.Append("num-sync-timestamp", types.YLeaf{"NumSyncTimestamp", platformPtpServo.NumSyncTimestamp})
    platformPtpServo.EntityData.Leafs.Append("num-delay-timestamp", types.YLeaf{"NumDelayTimestamp", platformPtpServo.NumDelayTimestamp})
    platformPtpServo.EntityData.Leafs.Append("num-set-time", types.YLeaf{"NumSetTime", platformPtpServo.NumSetTime})
    platformPtpServo.EntityData.Leafs.Append("num-step-time", types.YLeaf{"NumStepTime", platformPtpServo.NumStepTime})
    platformPtpServo.EntityData.Leafs.Append("num-adjust-freq", types.YLeaf{"NumAdjustFreq", platformPtpServo.NumAdjustFreq})
    platformPtpServo.EntityData.Leafs.Append("num-adjust-freq-time", types.YLeaf{"NumAdjustFreqTime", platformPtpServo.NumAdjustFreqTime})
    platformPtpServo.EntityData.Leafs.Append("last-adjust-freq", types.YLeaf{"LastAdjustFreq", platformPtpServo.LastAdjustFreq})
    platformPtpServo.EntityData.Leafs.Append("last-step-time", types.YLeaf{"LastStepTime", platformPtpServo.LastStepTime})
    platformPtpServo.EntityData.Leafs.Append("num-discard-sync-timestamp", types.YLeaf{"NumDiscardSyncTimestamp", platformPtpServo.NumDiscardSyncTimestamp})
    platformPtpServo.EntityData.Leafs.Append("num-discard-delay-timestamp", types.YLeaf{"NumDiscardDelayTimestamp", platformPtpServo.NumDiscardDelayTimestamp})
    platformPtpServo.EntityData.Leafs.Append("flagof-last-set-time", types.YLeaf{"FlagofLastSetTime", platformPtpServo.FlagofLastSetTime})
    platformPtpServo.EntityData.Leafs.Append("offset-from-master", types.YLeaf{"OffsetFromMaster", platformPtpServo.OffsetFromMaster})
    platformPtpServo.EntityData.Leafs.Append("mean-path-delay", types.YLeaf{"MeanPathDelay", platformPtpServo.MeanPathDelay})

    platformPtpServo.EntityData.YListKeys = []string {}

    return &(platformPtpServo.EntityData)
}

// PlatformPtp_PlatformPtpServo_LastSetTime
// last input of setTime
type PlatformPtp_PlatformPtpServo_LastSetTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetEntityData() *types.CommonEntityData {
    lastSetTime.EntityData.YFilter = lastSetTime.YFilter
    lastSetTime.EntityData.YangName = "last-set-time"
    lastSetTime.EntityData.BundleName = "cisco_ios_xr"
    lastSetTime.EntityData.ParentYangName = "platform-ptp-servo"
    lastSetTime.EntityData.SegmentPath = "last-set-time"
    lastSetTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastSetTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastSetTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastSetTime.EntityData.Children = types.NewOrderedMap()
    lastSetTime.EntityData.Leafs = types.NewOrderedMap()
    lastSetTime.EntityData.Leafs.Append("second", types.YLeaf{"Second", lastSetTime.Second})
    lastSetTime.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", lastSetTime.NanoSecond})

    lastSetTime.EntityData.YListKeys = []string {}

    return &(lastSetTime.EntityData)
}

// PlatformPtp_PlatformPtpServo_LastReceivedT1
// last T1 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetEntityData() *types.CommonEntityData {
    lastReceivedT1.EntityData.YFilter = lastReceivedT1.YFilter
    lastReceivedT1.EntityData.YangName = "last-received-t1"
    lastReceivedT1.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT1.EntityData.ParentYangName = "platform-ptp-servo"
    lastReceivedT1.EntityData.SegmentPath = "last-received-t1"
    lastReceivedT1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT1.EntityData.Children = types.NewOrderedMap()
    lastReceivedT1.EntityData.Leafs = types.NewOrderedMap()
    lastReceivedT1.EntityData.Leafs.Append("second", types.YLeaf{"Second", lastReceivedT1.Second})
    lastReceivedT1.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", lastReceivedT1.NanoSecond})

    lastReceivedT1.EntityData.YListKeys = []string {}

    return &(lastReceivedT1.EntityData)
}

// PlatformPtp_PlatformPtpServo_LastReceivedT2
// last T2 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetEntityData() *types.CommonEntityData {
    lastReceivedT2.EntityData.YFilter = lastReceivedT2.YFilter
    lastReceivedT2.EntityData.YangName = "last-received-t2"
    lastReceivedT2.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT2.EntityData.ParentYangName = "platform-ptp-servo"
    lastReceivedT2.EntityData.SegmentPath = "last-received-t2"
    lastReceivedT2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT2.EntityData.Children = types.NewOrderedMap()
    lastReceivedT2.EntityData.Leafs = types.NewOrderedMap()
    lastReceivedT2.EntityData.Leafs.Append("second", types.YLeaf{"Second", lastReceivedT2.Second})
    lastReceivedT2.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", lastReceivedT2.NanoSecond})

    lastReceivedT2.EntityData.YListKeys = []string {}

    return &(lastReceivedT2.EntityData)
}

// PlatformPtp_PlatformPtpServo_LastReceivedT3
// last T3 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetEntityData() *types.CommonEntityData {
    lastReceivedT3.EntityData.YFilter = lastReceivedT3.YFilter
    lastReceivedT3.EntityData.YangName = "last-received-t3"
    lastReceivedT3.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT3.EntityData.ParentYangName = "platform-ptp-servo"
    lastReceivedT3.EntityData.SegmentPath = "last-received-t3"
    lastReceivedT3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT3.EntityData.Children = types.NewOrderedMap()
    lastReceivedT3.EntityData.Leafs = types.NewOrderedMap()
    lastReceivedT3.EntityData.Leafs.Append("second", types.YLeaf{"Second", lastReceivedT3.Second})
    lastReceivedT3.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", lastReceivedT3.NanoSecond})

    lastReceivedT3.EntityData.YListKeys = []string {}

    return &(lastReceivedT3.EntityData)
}

// PlatformPtp_PlatformPtpServo_LastReceivedT4
// last T4 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetEntityData() *types.CommonEntityData {
    lastReceivedT4.EntityData.YFilter = lastReceivedT4.YFilter
    lastReceivedT4.EntityData.YangName = "last-received-t4"
    lastReceivedT4.EntityData.BundleName = "cisco_ios_xr"
    lastReceivedT4.EntityData.ParentYangName = "platform-ptp-servo"
    lastReceivedT4.EntityData.SegmentPath = "last-received-t4"
    lastReceivedT4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastReceivedT4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastReceivedT4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastReceivedT4.EntityData.Children = types.NewOrderedMap()
    lastReceivedT4.EntityData.Leafs = types.NewOrderedMap()
    lastReceivedT4.EntityData.Leafs.Append("second", types.YLeaf{"Second", lastReceivedT4.Second})
    lastReceivedT4.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", lastReceivedT4.NanoSecond})

    lastReceivedT4.EntityData.YListKeys = []string {}

    return &(lastReceivedT4.EntityData)
}

// PlatformPtp_PlatformPtpServo_PreReceivedT1
// pre T1 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetEntityData() *types.CommonEntityData {
    preReceivedT1.EntityData.YFilter = preReceivedT1.YFilter
    preReceivedT1.EntityData.YangName = "pre-received-t1"
    preReceivedT1.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT1.EntityData.ParentYangName = "platform-ptp-servo"
    preReceivedT1.EntityData.SegmentPath = "pre-received-t1"
    preReceivedT1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT1.EntityData.Children = types.NewOrderedMap()
    preReceivedT1.EntityData.Leafs = types.NewOrderedMap()
    preReceivedT1.EntityData.Leafs.Append("second", types.YLeaf{"Second", preReceivedT1.Second})
    preReceivedT1.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", preReceivedT1.NanoSecond})

    preReceivedT1.EntityData.YListKeys = []string {}

    return &(preReceivedT1.EntityData)
}

// PlatformPtp_PlatformPtpServo_PreReceivedT2
// pre T2 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetEntityData() *types.CommonEntityData {
    preReceivedT2.EntityData.YFilter = preReceivedT2.YFilter
    preReceivedT2.EntityData.YangName = "pre-received-t2"
    preReceivedT2.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT2.EntityData.ParentYangName = "platform-ptp-servo"
    preReceivedT2.EntityData.SegmentPath = "pre-received-t2"
    preReceivedT2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT2.EntityData.Children = types.NewOrderedMap()
    preReceivedT2.EntityData.Leafs = types.NewOrderedMap()
    preReceivedT2.EntityData.Leafs.Append("second", types.YLeaf{"Second", preReceivedT2.Second})
    preReceivedT2.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", preReceivedT2.NanoSecond})

    preReceivedT2.EntityData.YListKeys = []string {}

    return &(preReceivedT2.EntityData)
}

// PlatformPtp_PlatformPtpServo_PreReceivedT3
// pre T3 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetEntityData() *types.CommonEntityData {
    preReceivedT3.EntityData.YFilter = preReceivedT3.YFilter
    preReceivedT3.EntityData.YangName = "pre-received-t3"
    preReceivedT3.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT3.EntityData.ParentYangName = "platform-ptp-servo"
    preReceivedT3.EntityData.SegmentPath = "pre-received-t3"
    preReceivedT3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT3.EntityData.Children = types.NewOrderedMap()
    preReceivedT3.EntityData.Leafs = types.NewOrderedMap()
    preReceivedT3.EntityData.Leafs.Append("second", types.YLeaf{"Second", preReceivedT3.Second})
    preReceivedT3.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", preReceivedT3.NanoSecond})

    preReceivedT3.EntityData.YListKeys = []string {}

    return &(preReceivedT3.EntityData)
}

// PlatformPtp_PlatformPtpServo_PreReceivedT4
// pre T4 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetEntityData() *types.CommonEntityData {
    preReceivedT4.EntityData.YFilter = preReceivedT4.YFilter
    preReceivedT4.EntityData.YangName = "pre-received-t4"
    preReceivedT4.EntityData.BundleName = "cisco_ios_xr"
    preReceivedT4.EntityData.ParentYangName = "platform-ptp-servo"
    preReceivedT4.EntityData.SegmentPath = "pre-received-t4"
    preReceivedT4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preReceivedT4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preReceivedT4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preReceivedT4.EntityData.Children = types.NewOrderedMap()
    preReceivedT4.EntityData.Leafs = types.NewOrderedMap()
    preReceivedT4.EntityData.Leafs.Append("second", types.YLeaf{"Second", preReceivedT4.Second})
    preReceivedT4.EntityData.Leafs.Append("nano-second", types.YLeaf{"NanoSecond", preReceivedT4.NanoSecond})

    preReceivedT4.EntityData.YListKeys = []string {}

    return &(preReceivedT4.EntityData)
}

