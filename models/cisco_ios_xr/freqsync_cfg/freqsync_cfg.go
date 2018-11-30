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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    // Log both selection changes and errors
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
    EntityData types.CommonEntityData
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

func (frequencySynchronization *FrequencySynchronization) GetEntityData() *types.CommonEntityData {
    frequencySynchronization.EntityData.YFilter = frequencySynchronization.YFilter
    frequencySynchronization.EntityData.YangName = "frequency-synchronization"
    frequencySynchronization.EntityData.BundleName = "cisco_ios_xr"
    frequencySynchronization.EntityData.ParentYangName = "Cisco-IOS-XR-freqsync-cfg"
    frequencySynchronization.EntityData.SegmentPath = "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"
    frequencySynchronization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySynchronization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySynchronization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySynchronization.EntityData.Children = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", frequencySynchronization.QualityLevelOption})
    frequencySynchronization.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", frequencySynchronization.Enable})
    frequencySynchronization.EntityData.Leafs.Append("source-selection-logging", types.YLeaf{"SourceSelectionLogging", frequencySynchronization.SourceSelectionLogging})
    frequencySynchronization.EntityData.Leafs.Append("clock-interface-source-type", types.YLeaf{"ClockInterfaceSourceType", frequencySynchronization.ClockInterfaceSourceType})
    frequencySynchronization.EntityData.Leafs.Append("system-timing-mode", types.YLeaf{"SystemTimingMode", frequencySynchronization.SystemTimingMode})

    frequencySynchronization.EntityData.YListKeys = []string {}

    return &(frequencySynchronization.EntityData)
}

