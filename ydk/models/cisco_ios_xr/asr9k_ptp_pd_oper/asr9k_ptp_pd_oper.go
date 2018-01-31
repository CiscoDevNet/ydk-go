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
    parent types.Entity
    YFilter yfilter.YFilter

    // PTP PD Servo information.
    PlatformPtpServo PlatformPtp_PlatformPtpServo
}

func (platformPtp *PlatformPtp) GetFilter() yfilter.YFilter { return platformPtp.YFilter }

func (platformPtp *PlatformPtp) SetFilter(yf yfilter.YFilter) { platformPtp.YFilter = yf }

func (platformPtp *PlatformPtp) GetGoName(yname string) string {
    if yname == "platform-ptp-servo" { return "PlatformPtpServo" }
    return ""
}

func (platformPtp *PlatformPtp) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-ptp-pd-oper:platform-ptp"
}

func (platformPtp *PlatformPtp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "platform-ptp-servo" {
        return &platformPtp.PlatformPtpServo
    }
    return nil
}

func (platformPtp *PlatformPtp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["platform-ptp-servo"] = &platformPtp.PlatformPtpServo
    return children
}

func (platformPtp *PlatformPtp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformPtp *PlatformPtp) GetBundleName() string { return "cisco_ios_xr" }

func (platformPtp *PlatformPtp) GetYangName() string { return "platform-ptp" }

func (platformPtp *PlatformPtp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformPtp *PlatformPtp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformPtp *PlatformPtp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformPtp *PlatformPtp) SetParent(parent types.Entity) { platformPtp.parent = parent }

func (platformPtp *PlatformPtp) GetParent() types.Entity { return platformPtp.parent }

func (platformPtp *PlatformPtp) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-ptp-pd-oper" }

// PlatformPtp_PlatformPtpServo
// PTP PD Servo information
type PlatformPtp_PlatformPtpServo struct {
    parent types.Entity
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

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetFilter() yfilter.YFilter { return platformPtpServo.YFilter }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) SetFilter(yf yfilter.YFilter) { platformPtpServo.YFilter = yf }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetGoName(yname string) string {
    if yname == "lock-status" { return "LockStatus" }
    if yname == "running" { return "Running" }
    if yname == "device-status" { return "DeviceStatus" }
    if yname == "log-level" { return "LogLevel" }
    if yname == "phase-accuracy-last" { return "PhaseAccuracyLast" }
    if yname == "num-sync-timestamp" { return "NumSyncTimestamp" }
    if yname == "num-delay-timestamp" { return "NumDelayTimestamp" }
    if yname == "num-set-time" { return "NumSetTime" }
    if yname == "num-step-time" { return "NumStepTime" }
    if yname == "num-adjust-freq" { return "NumAdjustFreq" }
    if yname == "num-adjust-freq-time" { return "NumAdjustFreqTime" }
    if yname == "last-adjust-freq" { return "LastAdjustFreq" }
    if yname == "last-step-time" { return "LastStepTime" }
    if yname == "num-discard-sync-timestamp" { return "NumDiscardSyncTimestamp" }
    if yname == "num-discard-delay-timestamp" { return "NumDiscardDelayTimestamp" }
    if yname == "flagof-last-set-time" { return "FlagofLastSetTime" }
    if yname == "offset-from-master" { return "OffsetFromMaster" }
    if yname == "mean-path-delay" { return "MeanPathDelay" }
    if yname == "last-set-time" { return "LastSetTime" }
    if yname == "last-received-t1" { return "LastReceivedT1" }
    if yname == "last-received-t2" { return "LastReceivedT2" }
    if yname == "last-received-t3" { return "LastReceivedT3" }
    if yname == "last-received-t4" { return "LastReceivedT4" }
    if yname == "pre-received-t1" { return "PreReceivedT1" }
    if yname == "pre-received-t2" { return "PreReceivedT2" }
    if yname == "pre-received-t3" { return "PreReceivedT3" }
    if yname == "pre-received-t4" { return "PreReceivedT4" }
    return ""
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetSegmentPath() string {
    return "platform-ptp-servo"
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-set-time" {
        return &platformPtpServo.LastSetTime
    }
    if childYangName == "last-received-t1" {
        return &platformPtpServo.LastReceivedT1
    }
    if childYangName == "last-received-t2" {
        return &platformPtpServo.LastReceivedT2
    }
    if childYangName == "last-received-t3" {
        return &platformPtpServo.LastReceivedT3
    }
    if childYangName == "last-received-t4" {
        return &platformPtpServo.LastReceivedT4
    }
    if childYangName == "pre-received-t1" {
        return &platformPtpServo.PreReceivedT1
    }
    if childYangName == "pre-received-t2" {
        return &platformPtpServo.PreReceivedT2
    }
    if childYangName == "pre-received-t3" {
        return &platformPtpServo.PreReceivedT3
    }
    if childYangName == "pre-received-t4" {
        return &platformPtpServo.PreReceivedT4
    }
    return nil
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-set-time"] = &platformPtpServo.LastSetTime
    children["last-received-t1"] = &platformPtpServo.LastReceivedT1
    children["last-received-t2"] = &platformPtpServo.LastReceivedT2
    children["last-received-t3"] = &platformPtpServo.LastReceivedT3
    children["last-received-t4"] = &platformPtpServo.LastReceivedT4
    children["pre-received-t1"] = &platformPtpServo.PreReceivedT1
    children["pre-received-t2"] = &platformPtpServo.PreReceivedT2
    children["pre-received-t3"] = &platformPtpServo.PreReceivedT3
    children["pre-received-t4"] = &platformPtpServo.PreReceivedT4
    return children
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lock-status"] = platformPtpServo.LockStatus
    leafs["running"] = platformPtpServo.Running
    leafs["device-status"] = platformPtpServo.DeviceStatus
    leafs["log-level"] = platformPtpServo.LogLevel
    leafs["phase-accuracy-last"] = platformPtpServo.PhaseAccuracyLast
    leafs["num-sync-timestamp"] = platformPtpServo.NumSyncTimestamp
    leafs["num-delay-timestamp"] = platformPtpServo.NumDelayTimestamp
    leafs["num-set-time"] = platformPtpServo.NumSetTime
    leafs["num-step-time"] = platformPtpServo.NumStepTime
    leafs["num-adjust-freq"] = platformPtpServo.NumAdjustFreq
    leafs["num-adjust-freq-time"] = platformPtpServo.NumAdjustFreqTime
    leafs["last-adjust-freq"] = platformPtpServo.LastAdjustFreq
    leafs["last-step-time"] = platformPtpServo.LastStepTime
    leafs["num-discard-sync-timestamp"] = platformPtpServo.NumDiscardSyncTimestamp
    leafs["num-discard-delay-timestamp"] = platformPtpServo.NumDiscardDelayTimestamp
    leafs["flagof-last-set-time"] = platformPtpServo.FlagofLastSetTime
    leafs["offset-from-master"] = platformPtpServo.OffsetFromMaster
    leafs["mean-path-delay"] = platformPtpServo.MeanPathDelay
    return leafs
}

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetBundleName() string { return "cisco_ios_xr" }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetYangName() string { return "platform-ptp-servo" }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) SetParent(parent types.Entity) { platformPtpServo.parent = parent }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetParent() types.Entity { return platformPtpServo.parent }

func (platformPtpServo *PlatformPtp_PlatformPtpServo) GetParentYangName() string { return "platform-ptp" }

// PlatformPtp_PlatformPtpServo_LastSetTime
// last input of setTime
type PlatformPtp_PlatformPtpServo_LastSetTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetFilter() yfilter.YFilter { return lastSetTime.YFilter }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) SetFilter(yf yfilter.YFilter) { lastSetTime.YFilter = yf }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetSegmentPath() string {
    return "last-set-time"
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastSetTime.Second
    leafs["nano-second"] = lastSetTime.NanoSecond
    return leafs
}

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetBundleName() string { return "cisco_ios_xr" }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetYangName() string { return "last-set-time" }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) SetParent(parent types.Entity) { lastSetTime.parent = parent }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetParent() types.Entity { return lastSetTime.parent }

func (lastSetTime *PlatformPtp_PlatformPtpServo_LastSetTime) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_LastReceivedT1
// last T1 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetFilter() yfilter.YFilter { return lastReceivedT1.YFilter }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) SetFilter(yf yfilter.YFilter) { lastReceivedT1.YFilter = yf }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetSegmentPath() string {
    return "last-received-t1"
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT1.Second
    leafs["nano-second"] = lastReceivedT1.NanoSecond
    return leafs
}

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetYangName() string { return "last-received-t1" }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) SetParent(parent types.Entity) { lastReceivedT1.parent = parent }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetParent() types.Entity { return lastReceivedT1.parent }

func (lastReceivedT1 *PlatformPtp_PlatformPtpServo_LastReceivedT1) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_LastReceivedT2
// last T2 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetFilter() yfilter.YFilter { return lastReceivedT2.YFilter }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) SetFilter(yf yfilter.YFilter) { lastReceivedT2.YFilter = yf }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetSegmentPath() string {
    return "last-received-t2"
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT2.Second
    leafs["nano-second"] = lastReceivedT2.NanoSecond
    return leafs
}

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetYangName() string { return "last-received-t2" }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) SetParent(parent types.Entity) { lastReceivedT2.parent = parent }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetParent() types.Entity { return lastReceivedT2.parent }

func (lastReceivedT2 *PlatformPtp_PlatformPtpServo_LastReceivedT2) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_LastReceivedT3
// last T3 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetFilter() yfilter.YFilter { return lastReceivedT3.YFilter }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) SetFilter(yf yfilter.YFilter) { lastReceivedT3.YFilter = yf }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetSegmentPath() string {
    return "last-received-t3"
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT3.Second
    leafs["nano-second"] = lastReceivedT3.NanoSecond
    return leafs
}

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetYangName() string { return "last-received-t3" }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) SetParent(parent types.Entity) { lastReceivedT3.parent = parent }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetParent() types.Entity { return lastReceivedT3.parent }

func (lastReceivedT3 *PlatformPtp_PlatformPtpServo_LastReceivedT3) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_LastReceivedT4
// last T4 timestamp reveiced
type PlatformPtp_PlatformPtpServo_LastReceivedT4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetFilter() yfilter.YFilter { return lastReceivedT4.YFilter }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) SetFilter(yf yfilter.YFilter) { lastReceivedT4.YFilter = yf }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetSegmentPath() string {
    return "last-received-t4"
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = lastReceivedT4.Second
    leafs["nano-second"] = lastReceivedT4.NanoSecond
    return leafs
}

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetBundleName() string { return "cisco_ios_xr" }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetYangName() string { return "last-received-t4" }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) SetParent(parent types.Entity) { lastReceivedT4.parent = parent }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetParent() types.Entity { return lastReceivedT4.parent }

func (lastReceivedT4 *PlatformPtp_PlatformPtpServo_LastReceivedT4) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_PreReceivedT1
// pre T1 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetFilter() yfilter.YFilter { return preReceivedT1.YFilter }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) SetFilter(yf yfilter.YFilter) { preReceivedT1.YFilter = yf }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetSegmentPath() string {
    return "pre-received-t1"
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT1.Second
    leafs["nano-second"] = preReceivedT1.NanoSecond
    return leafs
}

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetYangName() string { return "pre-received-t1" }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) SetParent(parent types.Entity) { preReceivedT1.parent = parent }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetParent() types.Entity { return preReceivedT1.parent }

func (preReceivedT1 *PlatformPtp_PlatformPtpServo_PreReceivedT1) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_PreReceivedT2
// pre T2 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetFilter() yfilter.YFilter { return preReceivedT2.YFilter }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) SetFilter(yf yfilter.YFilter) { preReceivedT2.YFilter = yf }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetSegmentPath() string {
    return "pre-received-t2"
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT2.Second
    leafs["nano-second"] = preReceivedT2.NanoSecond
    return leafs
}

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetYangName() string { return "pre-received-t2" }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) SetParent(parent types.Entity) { preReceivedT2.parent = parent }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetParent() types.Entity { return preReceivedT2.parent }

func (preReceivedT2 *PlatformPtp_PlatformPtpServo_PreReceivedT2) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_PreReceivedT3
// pre T3 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetFilter() yfilter.YFilter { return preReceivedT3.YFilter }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) SetFilter(yf yfilter.YFilter) { preReceivedT3.YFilter = yf }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetSegmentPath() string {
    return "pre-received-t3"
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT3.Second
    leafs["nano-second"] = preReceivedT3.NanoSecond
    return leafs
}

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetYangName() string { return "pre-received-t3" }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) SetParent(parent types.Entity) { preReceivedT3.parent = parent }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetParent() types.Entity { return preReceivedT3.parent }

func (preReceivedT3 *PlatformPtp_PlatformPtpServo_PreReceivedT3) GetParentYangName() string { return "platform-ptp-servo" }

// PlatformPtp_PlatformPtpServo_PreReceivedT4
// pre T4 timestamp reveiced
type PlatformPtp_PlatformPtpServo_PreReceivedT4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // value of second. The type is interface{} with range: 0..4294967295.
    Second interface{}

    // value of nano second. The type is interface{} with range: 0..4294967295.
    NanoSecond interface{}
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetFilter() yfilter.YFilter { return preReceivedT4.YFilter }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) SetFilter(yf yfilter.YFilter) { preReceivedT4.YFilter = yf }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetGoName(yname string) string {
    if yname == "second" { return "Second" }
    if yname == "nano-second" { return "NanoSecond" }
    return ""
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetSegmentPath() string {
    return "pre-received-t4"
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["second"] = preReceivedT4.Second
    leafs["nano-second"] = preReceivedT4.NanoSecond
    return leafs
}

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetBundleName() string { return "cisco_ios_xr" }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetYangName() string { return "pre-received-t4" }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) SetParent(parent types.Entity) { preReceivedT4.parent = parent }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetParent() types.Entity { return preReceivedT4.parent }

func (preReceivedT4 *PlatformPtp_PlatformPtpServo_PreReceivedT4) GetParentYangName() string { return "platform-ptp-servo" }

