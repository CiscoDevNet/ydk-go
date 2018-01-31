// This module contains a collection of YANG definitions
// for Cisco IOS-XR freqsync package configuration.
// 
// This module contains definitions
// for the following management objects:
//   frequency-synchronization: frequency synchronization
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg,
//   Cisco-IOS-XR-ifmgr-cfg
//   Cisco-IOS-XR-config-mda-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package freqsync_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package freqsync_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-freqsync-cfg frequency-synchronization}", reflect.TypeOf(FrequencySynchronization{}))
    ydk.RegisterEntity("Cisco-IOS-XR-freqsync-cfg:frequency-synchronization", reflect.TypeOf(FrequencySynchronization{}))
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
    return "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"
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

func (frequencySynchronization *FrequencySynchronization) GetParentYangName() string { return "Cisco-IOS-XR-freqsync-cfg" }

