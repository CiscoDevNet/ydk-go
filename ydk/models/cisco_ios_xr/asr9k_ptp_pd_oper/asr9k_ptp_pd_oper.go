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

    platformPtp.EntityData.Children = make(map[string]types.YChild)
    platformPtp.EntityData.Children["platform-ptp-servo"] = types.YChild{"PlatformPtpServo", &platformPtp.PlatformPtpServo}
    platformPtp.EntityData.Leafs = make(map[string]types.YLeaf)
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

    platformPtpServo.EntityData.Children = make(map[string]types.YChild)
    platformPtpServo.EntityData.Children["last-set-time"] = types.YChild{"LastSetTime", &platformPtpServo.LastSetTime}
    platformPtpServo.EntityData.Children["last-received-t1"] = types.YChild{"LastReceivedT1", &platformPtpServo.LastReceivedT1}
    platformPtpServo.EntityData.Children["last-received-t2"] = types.YChild{"LastReceivedT2", &platformPtpServo.LastReceivedT2}
    platformPtpServo.EntityData.Children["last-received-t3"] = types.YChild{"LastReceivedT3", &platformPtpServo.LastReceivedT3}
    platformPtpServo.EntityData.Children["last-received-t4"] = types.YChild{"LastReceivedT4", &platformPtpServo.LastReceivedT4}
    platformPtpServo.EntityData.Children["pre-received-t1"] = types.YChild{"PreReceivedT1", &platformPtpServo.PreReceivedT1}
    platformPtpServo.EntityData.Children["pre-received-t2"] = types.YChild{"PreReceivedT2", &platformPtpServo.PreReceivedT2}
    platformPtpServo.EntityData.Children["pre-received-t3"] = types.YChild{"PreReceivedT3", &platformPtpServo.PreReceivedT3}
    platformPtpServo.EntityData.Children["pre-received-t4"] = types.YChild{"PreReceivedT4", &platformPtpServo.PreReceivedT4}
    platformPtpServo.EntityData.Leafs = make(map[string]types.YLeaf)
    platformPtpServo.EntityData.Leafs["lock-status"] = types.YLeaf{"LockStatus", platformPtpServo.LockStatus}
    platformPtpServo.EntityData.Leafs["running"] = types.YLeaf{"Running", platformPtpServo.Running}
    platformPtpServo.EntityData.Leafs["device-status"] = types.YLeaf{"DeviceStatus", platformPtpServo.DeviceStatus}
    platformPtpServo.EntityData.Leafs["log-level"] = types.YLeaf{"LogLevel", platformPtpServo.LogLevel}
    platformPtpServo.EntityData.Leafs["phase-accuracy-last"] = types.YLeaf{"PhaseAccuracyLast", platformPtpServo.PhaseAccuracyLast}
    platformPtpServo.EntityData.Leafs["num-sync-timestamp"] = types.YLeaf{"NumSyncTimestamp", platformPtpServo.NumSyncTimestamp}
    platformPtpServo.EntityData.Leafs["num-delay-timestamp"] = types.YLeaf{"NumDelayTimestamp", platformPtpServo.NumDelayTimestamp}
    platformPtpServo.EntityData.Leafs["num-set-time"] = types.YLeaf{"NumSetTime", platformPtpServo.NumSetTime}
    platformPtpServo.EntityData.Leafs["num-step-time"] = types.YLeaf{"NumStepTime", platformPtpServo.NumStepTime}
    platformPtpServo.EntityData.Leafs["num-adjust-freq"] = types.YLeaf{"NumAdjustFreq", platformPtpServo.NumAdjustFreq}
    platformPtpServo.EntityData.Leafs["num-adjust-freq-time"] = types.YLeaf{"NumAdjustFreqTime", platformPtpServo.NumAdjustFreqTime}
    platformPtpServo.EntityData.Leafs["last-adjust-freq"] = types.YLeaf{"LastAdjustFreq", platformPtpServo.LastAdjustFreq}
    platformPtpServo.EntityData.Leafs["last-step-time"] = types.YLeaf{"LastStepTime", platformPtpServo.LastStepTime}
    platformPtpServo.EntityData.Leafs["num-discard-sync-timestamp"] = types.YLeaf{"NumDiscardSyncTimestamp", platformPtpServo.NumDiscardSyncTimestamp}
    platformPtpServo.EntityData.Leafs["num-discard-delay-timestamp"] = types.YLeaf{"NumDiscardDelayTimestamp", platformPtpServo.NumDiscardDelayTimestamp}
    platformPtpServo.EntityData.Leafs["flagof-last-set-time"] = types.YLeaf{"FlagofLastSetTime", platformPtpServo.FlagofLastSetTime}
    platformPtpServo.EntityData.Leafs["offset-from-master"] = types.YLeaf{"OffsetFromMaster", platformPtpServo.OffsetFromMaster}
    platformPtpServo.EntityData.Leafs["mean-path-delay"] = types.YLeaf{"MeanPathDelay", platformPtpServo.MeanPathDelay}
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

    lastSetTime.EntityData.Children = make(map[string]types.YChild)
    lastSetTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lastSetTime.EntityData.Leafs["second"] = types.YLeaf{"Second", lastSetTime.Second}
    lastSetTime.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastSetTime.NanoSecond}
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

    lastReceivedT1.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT1.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT1.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT1.Second}
    lastReceivedT1.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT1.NanoSecond}
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

    lastReceivedT2.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT2.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT2.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT2.Second}
    lastReceivedT2.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT2.NanoSecond}
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

    lastReceivedT3.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT3.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT3.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT3.Second}
    lastReceivedT3.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT3.NanoSecond}
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

    lastReceivedT4.EntityData.Children = make(map[string]types.YChild)
    lastReceivedT4.EntityData.Leafs = make(map[string]types.YLeaf)
    lastReceivedT4.EntityData.Leafs["second"] = types.YLeaf{"Second", lastReceivedT4.Second}
    lastReceivedT4.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", lastReceivedT4.NanoSecond}
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

    preReceivedT1.EntityData.Children = make(map[string]types.YChild)
    preReceivedT1.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT1.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT1.Second}
    preReceivedT1.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT1.NanoSecond}
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

    preReceivedT2.EntityData.Children = make(map[string]types.YChild)
    preReceivedT2.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT2.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT2.Second}
    preReceivedT2.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT2.NanoSecond}
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

    preReceivedT3.EntityData.Children = make(map[string]types.YChild)
    preReceivedT3.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT3.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT3.Second}
    preReceivedT3.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT3.NanoSecond}
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

    preReceivedT4.EntityData.Children = make(map[string]types.YChild)
    preReceivedT4.EntityData.Leafs = make(map[string]types.YLeaf)
    preReceivedT4.EntityData.Leafs["second"] = types.YLeaf{"Second", preReceivedT4.Second}
    preReceivedT4.EntityData.Leafs["nano-second"] = types.YLeaf{"NanoSecond", preReceivedT4.NanoSecond}
    return &(preReceivedT4.EntityData)
}

