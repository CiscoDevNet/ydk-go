// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs4k-freqsync package configuration.
// 
// This module contains definitions
// for the following management objects:
//   clock-interface: Configuration for a clock interface
//   frequency-synchronization: frequency synchronization
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg,
//   Cisco-IOS-XR-ifmgr-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs4k_freqsync_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs4k_freqsync_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs4k-freqsync-cfg clock-interface}", reflect.TypeOf(ClockInterface{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs4k-freqsync-cfg:clock-interface", reflect.TypeOf(ClockInterface{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs4k-freqsync-cfg frequency-synchronization}", reflect.TypeOf(FrequencySynchronization{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization", reflect.TypeOf(FrequencySynchronization{}))
}

// FsyncClockSource represents Fsync clock source
type FsyncClockSource string

const (
    // System
    FsyncClockSource_system FsyncClockSource = "system"

    // Independent
    FsyncClockSource_independent FsyncClockSource = "independent"
)

// FsyncSourceSelectionLogging represents Fsync source selection logging
type FsyncSourceSelectionLogging string

const (
    // Log selection changes
    FsyncSourceSelectionLogging_changes FsyncSourceSelectionLogging = "changes"

    // Log selection errors
    FsyncSourceSelectionLogging_errors FsyncSourceSelectionLogging = "errors"
)

// FsyncSystemTimingMode represents Fsync system timing mode
type FsyncSystemTimingMode string

const (
    // Line-interfaces only
    FsyncSystemTimingMode_line_only FsyncSystemTimingMode = "line-only"

    // Clock-interfaces only
    FsyncSystemTimingMode_clock_only FsyncSystemTimingMode = "clock-only"
)

// ClockInterface
// Configuration for a clock interface
type ClockInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface.
    Clocks ClockInterface_Clocks
}

func (clockInterface *ClockInterface) GetFilter() yfilter.YFilter { return clockInterface.YFilter }

func (clockInterface *ClockInterface) SetFilter(yf yfilter.YFilter) { clockInterface.YFilter = yf }

func (clockInterface *ClockInterface) GetGoName(yname string) string {
    if yname == "clocks" { return "Clocks" }
    return ""
}

func (clockInterface *ClockInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs4k-freqsync-cfg:clock-interface"
}

func (clockInterface *ClockInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clocks" {
        return &clockInterface.Clocks
    }
    return nil
}

func (clockInterface *ClockInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clocks"] = &clockInterface.Clocks
    return children
}

func (clockInterface *ClockInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clockInterface *ClockInterface) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterface *ClockInterface) GetYangName() string { return "clock-interface" }

func (clockInterface *ClockInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterface *ClockInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterface *ClockInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterface *ClockInterface) SetParent(parent types.Entity) { clockInterface.parent = parent }

func (clockInterface *ClockInterface) GetParent() types.Entity { return clockInterface.parent }

func (clockInterface *ClockInterface) GetParentYangName() string { return "Cisco-IOS-XR-ncs4k-freqsync-cfg" }

// ClockInterface_Clocks
// Configuration for a clock interface
type ClockInterface_Clocks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface. The type is slice of
    // ClockInterface_Clocks_Clock.
    Clock []ClockInterface_Clocks_Clock
}

func (clocks *ClockInterface_Clocks) GetFilter() yfilter.YFilter { return clocks.YFilter }

func (clocks *ClockInterface_Clocks) SetFilter(yf yfilter.YFilter) { clocks.YFilter = yf }

func (clocks *ClockInterface_Clocks) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    return ""
}

func (clocks *ClockInterface_Clocks) GetSegmentPath() string {
    return "clocks"
}

func (clocks *ClockInterface_Clocks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        for _, c := range clocks.Clock {
            if clocks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ClockInterface_Clocks_Clock{}
        clocks.Clock = append(clocks.Clock, child)
        return &clocks.Clock[len(clocks.Clock)-1]
    }
    return nil
}

func (clocks *ClockInterface_Clocks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clocks.Clock {
        children[clocks.Clock[i].GetSegmentPath()] = &clocks.Clock[i]
    }
    return children
}

func (clocks *ClockInterface_Clocks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clocks *ClockInterface_Clocks) GetBundleName() string { return "cisco_ios_xr" }

func (clocks *ClockInterface_Clocks) GetYangName() string { return "clocks" }

func (clocks *ClockInterface_Clocks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clocks *ClockInterface_Clocks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clocks *ClockInterface_Clocks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clocks *ClockInterface_Clocks) SetParent(parent types.Entity) { clocks.parent = parent }

func (clocks *ClockInterface_Clocks) GetParent() types.Entity { return clocks.parent }

func (clocks *ClockInterface_Clocks) GetParentYangName() string { return "clock-interface" }

// ClockInterface_Clocks_Clock
// Configuration for a clock interface
type ClockInterface_Clocks_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClockName interface{}

    // Frequency Synchronization clock configuraiton.
    FrequencySynchronization ClockInterface_Clocks_Clock_FrequencySynchronization
}

func (clock *ClockInterface_Clocks_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *ClockInterface_Clocks_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *ClockInterface_Clocks_Clock) GetGoName(yname string) string {
    if yname == "clock-name" { return "ClockName" }
    if yname == "frequency-synchronization" { return "FrequencySynchronization" }
    return ""
}

func (clock *ClockInterface_Clocks_Clock) GetSegmentPath() string {
    return "clock" + "[clock-name='" + fmt.Sprintf("%v", clock.ClockName) + "']"
}

func (clock *ClockInterface_Clocks_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frequency-synchronization" {
        return &clock.FrequencySynchronization
    }
    return nil
}

func (clock *ClockInterface_Clocks_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frequency-synchronization"] = &clock.FrequencySynchronization
    return children
}

func (clock *ClockInterface_Clocks_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-name"] = clock.ClockName
    return leafs
}

func (clock *ClockInterface_Clocks_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *ClockInterface_Clocks_Clock) GetYangName() string { return "clock" }

func (clock *ClockInterface_Clocks_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *ClockInterface_Clocks_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *ClockInterface_Clocks_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *ClockInterface_Clocks_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *ClockInterface_Clocks_Clock) GetParent() types.Entity { return clock.parent }

func (clock *ClockInterface_Clocks_Clock) GetParentYangName() string { return "clocks" }

// ClockInterface_Clocks_Clock_FrequencySynchronization
// Frequency Synchronization clock configuraiton
type ClockInterface_Clocks_Clock_FrequencySynchronization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable SSM on this source. The type is interface{}.
    SsmDisable interface{}

    // Set the wait-to-restore time for this source. The type is interface{} with
    // range: 0..12. The default value is 5.
    WaitToRestoreTime interface{}

    // Set the time-of-day priority of this source. The type is interface{} with
    // range: 1..254. The default value is 100.
    TimeOfDayPriority interface{}

    // Set the priority of this source. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // Assign this source as a selection input. The type is interface{}.
    SelectionInput interface{}

    // Set the input quality level.
    InputQualityLevel ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel

    // Set the output quality level.
    OutputQualityLevel ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetFilter() yfilter.YFilter { return frequencySynchronization.YFilter }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) SetFilter(yf yfilter.YFilter) { frequencySynchronization.YFilter = yf }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetGoName(yname string) string {
    if yname == "ssm-disable" { return "SsmDisable" }
    if yname == "wait-to-restore-time" { return "WaitToRestoreTime" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "priority" { return "Priority" }
    if yname == "selection-input" { return "SelectionInput" }
    if yname == "input-quality-level" { return "InputQualityLevel" }
    if yname == "output-quality-level" { return "OutputQualityLevel" }
    return ""
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetSegmentPath() string {
    return "frequency-synchronization"
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-quality-level" {
        return &frequencySynchronization.InputQualityLevel
    }
    if childYangName == "output-quality-level" {
        return &frequencySynchronization.OutputQualityLevel
    }
    return nil
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-quality-level"] = &frequencySynchronization.InputQualityLevel
    children["output-quality-level"] = &frequencySynchronization.OutputQualityLevel
    return children
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ssm-disable"] = frequencySynchronization.SsmDisable
    leafs["wait-to-restore-time"] = frequencySynchronization.WaitToRestoreTime
    leafs["time-of-day-priority"] = frequencySynchronization.TimeOfDayPriority
    leafs["priority"] = frequencySynchronization.Priority
    leafs["selection-input"] = frequencySynchronization.SelectionInput
    return leafs
}

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) SetParent(parent types.Entity) { frequencySynchronization.parent = parent }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetParent() types.Entity { return frequencySynchronization.parent }

func (frequencySynchronization *ClockInterface_Clocks_Clock_FrequencySynchronization) GetParentYangName() string { return "clock" }

// ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
// Set the input quality level
type ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Quality level option. The type is FsyncQlOption.
    QualityLevelOption interface{}

    // Exact quality level value. The type is FsyncQlValue.
    ExactQualityLevelValue interface{}

    // Minimum quality level value. The type is FsyncQlValue.
    MinQualityLevelValue interface{}

    // Maximum quality level value. The type is FsyncQlValue.
    MaxQualityLevelValue interface{}
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetFilter() yfilter.YFilter { return inputQualityLevel.YFilter }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetFilter(yf yfilter.YFilter) { inputQualityLevel.YFilter = yf }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetSegmentPath() string {
    return "input-quality-level"
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = inputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = inputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = inputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetYangName() string { return "input-quality-level" }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetParent(parent types.Entity) { inputQualityLevel.parent = parent }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParent() types.Entity { return inputQualityLevel.parent }

func (inputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
// Set the output quality level
type ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Quality level option. The type is FsyncQlOption.
    QualityLevelOption interface{}

    // Exact quality level value. The type is FsyncQlValue.
    ExactQualityLevelValue interface{}

    // Minimum quality level value. The type is FsyncQlValue.
    MinQualityLevelValue interface{}

    // Maximum quality level value. The type is FsyncQlValue.
    MaxQualityLevelValue interface{}
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetFilter() yfilter.YFilter { return outputQualityLevel.YFilter }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetFilter(yf yfilter.YFilter) { outputQualityLevel.YFilter = yf }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetSegmentPath() string {
    return "output-quality-level"
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = outputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = outputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = outputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetYangName() string { return "output-quality-level" }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetParent(parent types.Entity) { outputQualityLevel.parent = parent }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParent() types.Entity { return outputQualityLevel.parent }

func (outputQualityLevel *ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization
// frequency synchronization
type FrequencySynchronization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Quality level option. The type is FsyncQlOption. The default value is
    // option-1.
    QualityLevelOption interface{}

    // Enable Frequency Synchronization. The type is interface{}.
    Enable interface{}

    // Source selection logging option. The type is FsyncSourceSelectionLogging.
    SourceSelectionLogging interface{}

    // Clock interface source type. The type is FsyncClockSource.
    ClockInterfaceSourceType interface{}

    // System timing mode. The type is FsyncSystemTimingMode.
    SystemTimingMode interface{}
}

func (frequencySynchronization *FrequencySynchronization) GetFilter() yfilter.YFilter { return frequencySynchronization.YFilter }

func (frequencySynchronization *FrequencySynchronization) SetFilter(yf yfilter.YFilter) { frequencySynchronization.YFilter = yf }

func (frequencySynchronization *FrequencySynchronization) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "enable" { return "Enable" }
    if yname == "source-selection-logging" { return "SourceSelectionLogging" }
    if yname == "clock-interface-source-type" { return "ClockInterfaceSourceType" }
    if yname == "system-timing-mode" { return "SystemTimingMode" }
    return ""
}

func (frequencySynchronization *FrequencySynchronization) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization"
}

func (frequencySynchronization *FrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frequencySynchronization *FrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frequencySynchronization *FrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = frequencySynchronization.QualityLevelOption
    leafs["enable"] = frequencySynchronization.Enable
    leafs["source-selection-logging"] = frequencySynchronization.SourceSelectionLogging
    leafs["clock-interface-source-type"] = frequencySynchronization.ClockInterfaceSourceType
    leafs["system-timing-mode"] = frequencySynchronization.SystemTimingMode
    return leafs
}

func (frequencySynchronization *FrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySynchronization *FrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (frequencySynchronization *FrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySynchronization *FrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySynchronization *FrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySynchronization *FrequencySynchronization) SetParent(parent types.Entity) { frequencySynchronization.parent = parent }

func (frequencySynchronization *FrequencySynchronization) GetParent() types.Entity { return frequencySynchronization.parent }

func (frequencySynchronization *FrequencySynchronization) GetParentYangName() string { return "Cisco-IOS-XR-ncs4k-freqsync-cfg" }

