// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-mda package configuration.
// 
// This module contains definitions
// for the following management objects:
//   active-nodes: Per-node configuration for active nodes
//   preconfigured-nodes: preconfigured nodes
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package config_mda_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package config_mda_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-mda-cfg active-nodes}", reflect.TypeOf(ActiveNodes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-mda-cfg:active-nodes", reflect.TypeOf(ActiveNodes{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-mda-cfg preconfigured-nodes}", reflect.TypeOf(PreconfiguredNodes{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-mda-cfg:preconfigured-nodes", reflect.TypeOf(PreconfiguredNodes{}))
}

// ActiveNodes
// Per-node configuration for active nodes
type ActiveNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The configuration for an active node. The type is slice of
    // ActiveNodes_ActiveNode.
    ActiveNode []*ActiveNodes_ActiveNode
}

func (activeNodes *ActiveNodes) GetEntityData() *types.CommonEntityData {
    activeNodes.EntityData.YFilter = activeNodes.YFilter
    activeNodes.EntityData.YangName = "active-nodes"
    activeNodes.EntityData.BundleName = "cisco_ios_xr"
    activeNodes.EntityData.ParentYangName = "Cisco-IOS-XR-config-mda-cfg"
    activeNodes.EntityData.SegmentPath = "Cisco-IOS-XR-config-mda-cfg:active-nodes"
    activeNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeNodes.EntityData.Children = types.NewOrderedMap()
    activeNodes.EntityData.Children.Append("active-node", types.YChild{"ActiveNode", nil})
    for i := range activeNodes.ActiveNode {
        activeNodes.EntityData.Children.Append(types.GetSegmentPath(activeNodes.ActiveNode[i]), types.YChild{"ActiveNode", activeNodes.ActiveNode[i]})
    }
    activeNodes.EntityData.Leafs = types.NewOrderedMap()

    activeNodes.EntityData.YListKeys = []string {}

    return &(activeNodes.EntityData)
}

// ActiveNodes_ActiveNode
// The configuration for an active node
type ActiveNodes_ActiveNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // watchdog node threshold.
    WatchdogNodeThreshold ActiveNodes_ActiveNode_WatchdogNodeThreshold

    // Ltrace Memory configuration.
    Ltrace ActiveNodes_ActiveNode_Ltrace

    // Configuration for a clock interface.
    ClockInterface ActiveNodes_ActiveNode_ClockInterface

    // Per-node SSRP configuration data.
    SsrpGroup ActiveNodes_ActiveNode_SsrpGroup

    // lpts node specific configuration commands.
    LptsLocal ActiveNodes_ActiveNode_LptsLocal
}

func (activeNode *ActiveNodes_ActiveNode) GetEntityData() *types.CommonEntityData {
    activeNode.EntityData.YFilter = activeNode.YFilter
    activeNode.EntityData.YangName = "active-node"
    activeNode.EntityData.BundleName = "cisco_ios_xr"
    activeNode.EntityData.ParentYangName = "active-nodes"
    activeNode.EntityData.SegmentPath = "active-node" + types.AddKeyToken(activeNode.NodeName, "node-name")
    activeNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeNode.EntityData.Children = types.NewOrderedMap()
    activeNode.EntityData.Children.Append("Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold", types.YChild{"WatchdogNodeThreshold", &activeNode.WatchdogNodeThreshold})
    activeNode.EntityData.Children.Append("Cisco-IOS-XR-infra-ltrace-cfg:ltrace", types.YChild{"Ltrace", &activeNode.Ltrace})
    activeNode.EntityData.Children.Append("Cisco-IOS-XR-freqsync-cfg:clock-interface", types.YChild{"ClockInterface", &activeNode.ClockInterface})
    activeNode.EntityData.Children.Append("Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group", types.YChild{"SsrpGroup", &activeNode.SsrpGroup})
    activeNode.EntityData.Children.Append("Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local", types.YChild{"LptsLocal", &activeNode.LptsLocal})
    activeNode.EntityData.Leafs = types.NewOrderedMap()
    activeNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", activeNode.NodeName})

    activeNode.EntityData.YListKeys = []string {"NodeName"}

    return &(activeNode.EntityData)
}

// ActiveNodes_ActiveNode_WatchdogNodeThreshold
// watchdog node threshold
type ActiveNodes_ActiveNode_WatchdogNodeThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disk thresholds.
    DiskThreshold ActiveNodes_ActiveNode_WatchdogNodeThreshold_DiskThreshold

    // Memory thresholds.
    MemoryThreshold ActiveNodes_ActiveNode_WatchdogNodeThreshold_MemoryThreshold
}

func (watchdogNodeThreshold *ActiveNodes_ActiveNode_WatchdogNodeThreshold) GetEntityData() *types.CommonEntityData {
    watchdogNodeThreshold.EntityData.YFilter = watchdogNodeThreshold.YFilter
    watchdogNodeThreshold.EntityData.YangName = "watchdog-node-threshold"
    watchdogNodeThreshold.EntityData.BundleName = "cisco_ios_xr"
    watchdogNodeThreshold.EntityData.ParentYangName = "active-node"
    watchdogNodeThreshold.EntityData.SegmentPath = "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"
    watchdogNodeThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdogNodeThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdogNodeThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdogNodeThreshold.EntityData.Children = types.NewOrderedMap()
    watchdogNodeThreshold.EntityData.Children.Append("disk-threshold", types.YChild{"DiskThreshold", &watchdogNodeThreshold.DiskThreshold})
    watchdogNodeThreshold.EntityData.Children.Append("memory-threshold", types.YChild{"MemoryThreshold", &watchdogNodeThreshold.MemoryThreshold})
    watchdogNodeThreshold.EntityData.Leafs = types.NewOrderedMap()

    watchdogNodeThreshold.EntityData.YListKeys = []string {}

    return &(watchdogNodeThreshold.EntityData)
}

// ActiveNodes_ActiveNode_WatchdogNodeThreshold_DiskThreshold
// Disk thresholds
type ActiveNodes_ActiveNode_WatchdogNodeThreshold_DiskThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (diskThreshold *ActiveNodes_ActiveNode_WatchdogNodeThreshold_DiskThreshold) GetEntityData() *types.CommonEntityData {
    diskThreshold.EntityData.YFilter = diskThreshold.YFilter
    diskThreshold.EntityData.YangName = "disk-threshold"
    diskThreshold.EntityData.BundleName = "cisco_ios_xr"
    diskThreshold.EntityData.ParentYangName = "watchdog-node-threshold"
    diskThreshold.EntityData.SegmentPath = "disk-threshold"
    diskThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diskThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diskThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diskThreshold.EntityData.Children = types.NewOrderedMap()
    diskThreshold.EntityData.Leafs = types.NewOrderedMap()
    diskThreshold.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", diskThreshold.Minor})
    diskThreshold.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", diskThreshold.Severe})
    diskThreshold.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", diskThreshold.Critical})

    diskThreshold.EntityData.YListKeys = []string {}

    return &(diskThreshold.EntityData)
}

// ActiveNodes_ActiveNode_WatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type ActiveNodes_ActiveNode_WatchdogNodeThreshold_MemoryThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *ActiveNodes_ActiveNode_WatchdogNodeThreshold_MemoryThreshold) GetEntityData() *types.CommonEntityData {
    memoryThreshold.EntityData.YFilter = memoryThreshold.YFilter
    memoryThreshold.EntityData.YangName = "memory-threshold"
    memoryThreshold.EntityData.BundleName = "cisco_ios_xr"
    memoryThreshold.EntityData.ParentYangName = "watchdog-node-threshold"
    memoryThreshold.EntityData.SegmentPath = "memory-threshold"
    memoryThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryThreshold.EntityData.Children = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", memoryThreshold.Minor})
    memoryThreshold.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", memoryThreshold.Severe})
    memoryThreshold.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", memoryThreshold.Critical})

    memoryThreshold.EntityData.YListKeys = []string {}

    return &(memoryThreshold.EntityData)
}

// ActiveNodes_ActiveNode_Ltrace
// Ltrace Memory configuration
type ActiveNodes_ActiveNode_Ltrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select Ltrace mode and scale-factor.
    AllocationParams ActiveNodes_ActiveNode_Ltrace_AllocationParams
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetEntityData() *types.CommonEntityData {
    ltrace.EntityData.YFilter = ltrace.YFilter
    ltrace.EntityData.YangName = "ltrace"
    ltrace.EntityData.BundleName = "cisco_ios_xr"
    ltrace.EntityData.ParentYangName = "active-node"
    ltrace.EntityData.SegmentPath = "Cisco-IOS-XR-infra-ltrace-cfg:ltrace"
    ltrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ltrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ltrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ltrace.EntityData.Children = types.NewOrderedMap()
    ltrace.EntityData.Children.Append("allocation-params", types.YChild{"AllocationParams", &ltrace.AllocationParams})
    ltrace.EntityData.Leafs = types.NewOrderedMap()

    ltrace.EntityData.YListKeys = []string {}

    return &(ltrace.EntityData)
}

// ActiveNodes_ActiveNode_Ltrace_AllocationParams
// Select Ltrace mode and scale-factor
type ActiveNodes_ActiveNode_Ltrace_AllocationParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an allocation mode (static:1, dynamic :2). The type is
    // InfraLtraceMode.
    Mode interface{}

    // Select a scaling down factor. The type is InfraLtraceScale.
    ScaleFactor interface{}
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetEntityData() *types.CommonEntityData {
    allocationParams.EntityData.YFilter = allocationParams.YFilter
    allocationParams.EntityData.YangName = "allocation-params"
    allocationParams.EntityData.BundleName = "cisco_ios_xr"
    allocationParams.EntityData.ParentYangName = "ltrace"
    allocationParams.EntityData.SegmentPath = "allocation-params"
    allocationParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocationParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocationParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocationParams.EntityData.Children = types.NewOrderedMap()
    allocationParams.EntityData.Leafs = types.NewOrderedMap()
    allocationParams.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", allocationParams.Mode})
    allocationParams.EntityData.Leafs.Append("scale-factor", types.YLeaf{"ScaleFactor", allocationParams.ScaleFactor})

    allocationParams.EntityData.YListKeys = []string {}

    return &(allocationParams.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a clock interface.
    Clocks ActiveNodes_ActiveNode_ClockInterface_Clocks
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetEntityData() *types.CommonEntityData {
    clockInterface.EntityData.YFilter = clockInterface.YFilter
    clockInterface.EntityData.YangName = "clock-interface"
    clockInterface.EntityData.BundleName = "cisco_ios_xr"
    clockInterface.EntityData.ParentYangName = "active-node"
    clockInterface.EntityData.SegmentPath = "Cisco-IOS-XR-freqsync-cfg:clock-interface"
    clockInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterface.EntityData.Children = types.NewOrderedMap()
    clockInterface.EntityData.Children.Append("clocks", types.YChild{"Clocks", &clockInterface.Clocks})
    clockInterface.EntityData.Leafs = types.NewOrderedMap()

    clockInterface.EntityData.YListKeys = []string {}

    return &(clockInterface.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface_Clocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a clock interface. The type is slice of
    // ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock.
    Clock []*ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetEntityData() *types.CommonEntityData {
    clocks.EntityData.YFilter = clocks.YFilter
    clocks.EntityData.YangName = "clocks"
    clocks.EntityData.BundleName = "cisco_ios_xr"
    clocks.EntityData.ParentYangName = "clock-interface"
    clocks.EntityData.SegmentPath = "clocks"
    clocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clocks.EntityData.Children = types.NewOrderedMap()
    clocks.EntityData.Children.Append("clock", types.YChild{"Clock", nil})
    for i := range clocks.Clock {
        clocks.EntityData.Children.Append(types.GetSegmentPath(clocks.Clock[i]), types.YChild{"Clock", clocks.Clock[i]})
    }
    clocks.EntityData.Leafs = types.NewOrderedMap()

    clocks.EntityData.YListKeys = []string {}

    return &(clocks.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // 0..4294967295.
    Port interface{}

    // Frequency Synchronization clock configuraiton.
    FrequencySynchronization ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization

    // sync-controller value.
    SyncController ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "clocks"
    clock.EntityData.SegmentPath = "clock" + types.AddKeyToken(clock.ClockType, "clock-type") + types.AddKeyToken(clock.Port, "port")
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("frequency-synchronization", types.YChild{"FrequencySynchronization", &clock.FrequencySynchronization})
    clock.EntityData.Children.Append("Cisco-IOS-XR-syncc-controller-cfg:sync-controller", types.YChild{"SyncController", &clock.SyncController})
    clock.EntityData.Leafs = types.NewOrderedMap()
    clock.EntityData.Leafs.Append("clock-type", types.YLeaf{"ClockType", clock.ClockType})
    clock.EntityData.Leafs.Append("port", types.YLeaf{"Port", clock.Port})

    clock.EntityData.YListKeys = []string {"ClockType", "Port"}

    return &(clock.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization
// Frequency Synchronization clock configuraiton
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the wait-to-restore time for this source. The type is interface{} with
    // range: 0..12. The default value is 5.
    WaitToRestoreTime interface{}

    // Set the priority of this source. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // Assign this source as a selection input. The type is interface{}.
    SelectionInput interface{}

    // Set the time-of-day priority of this source. The type is interface{} with
    // range: 1..254. The default value is 100.
    TimeOfDayPriority interface{}

    // Disable SSM on this source. The type is interface{}.
    SsmDisable interface{}

    // Set the output quality level.
    OutputQualityLevel ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel

    // Set the input quality level.
    InputQualityLevel ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetEntityData() *types.CommonEntityData {
    frequencySynchronization.EntityData.YFilter = frequencySynchronization.YFilter
    frequencySynchronization.EntityData.YangName = "frequency-synchronization"
    frequencySynchronization.EntityData.BundleName = "cisco_ios_xr"
    frequencySynchronization.EntityData.ParentYangName = "clock"
    frequencySynchronization.EntityData.SegmentPath = "frequency-synchronization"
    frequencySynchronization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySynchronization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySynchronization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySynchronization.EntityData.Children = types.NewOrderedMap()
    frequencySynchronization.EntityData.Children.Append("output-quality-level", types.YChild{"OutputQualityLevel", &frequencySynchronization.OutputQualityLevel})
    frequencySynchronization.EntityData.Children.Append("input-quality-level", types.YChild{"InputQualityLevel", &frequencySynchronization.InputQualityLevel})
    frequencySynchronization.EntityData.Leafs = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs.Append("wait-to-restore-time", types.YLeaf{"WaitToRestoreTime", frequencySynchronization.WaitToRestoreTime})
    frequencySynchronization.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", frequencySynchronization.Priority})
    frequencySynchronization.EntityData.Leafs.Append("selection-input", types.YLeaf{"SelectionInput", frequencySynchronization.SelectionInput})
    frequencySynchronization.EntityData.Leafs.Append("time-of-day-priority", types.YLeaf{"TimeOfDayPriority", frequencySynchronization.TimeOfDayPriority})
    frequencySynchronization.EntityData.Leafs.Append("ssm-disable", types.YLeaf{"SsmDisable", frequencySynchronization.SsmDisable})

    frequencySynchronization.EntityData.YListKeys = []string {}

    return &(frequencySynchronization.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
// Set the output quality level
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel struct {
    EntityData types.CommonEntityData
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

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetEntityData() *types.CommonEntityData {
    outputQualityLevel.EntityData.YFilter = outputQualityLevel.YFilter
    outputQualityLevel.EntityData.YangName = "output-quality-level"
    outputQualityLevel.EntityData.BundleName = "cisco_ios_xr"
    outputQualityLevel.EntityData.ParentYangName = "frequency-synchronization"
    outputQualityLevel.EntityData.SegmentPath = "output-quality-level"
    outputQualityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputQualityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputQualityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputQualityLevel.EntityData.Children = types.NewOrderedMap()
    outputQualityLevel.EntityData.Leafs = types.NewOrderedMap()
    outputQualityLevel.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", outputQualityLevel.QualityLevelOption})
    outputQualityLevel.EntityData.Leafs.Append("exact-quality-level-value", types.YLeaf{"ExactQualityLevelValue", outputQualityLevel.ExactQualityLevelValue})
    outputQualityLevel.EntityData.Leafs.Append("min-quality-level-value", types.YLeaf{"MinQualityLevelValue", outputQualityLevel.MinQualityLevelValue})
    outputQualityLevel.EntityData.Leafs.Append("max-quality-level-value", types.YLeaf{"MaxQualityLevelValue", outputQualityLevel.MaxQualityLevelValue})

    outputQualityLevel.EntityData.YListKeys = []string {}

    return &(outputQualityLevel.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
// Set the input quality level
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel struct {
    EntityData types.CommonEntityData
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

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetEntityData() *types.CommonEntityData {
    inputQualityLevel.EntityData.YFilter = inputQualityLevel.YFilter
    inputQualityLevel.EntityData.YangName = "input-quality-level"
    inputQualityLevel.EntityData.BundleName = "cisco_ios_xr"
    inputQualityLevel.EntityData.ParentYangName = "frequency-synchronization"
    inputQualityLevel.EntityData.SegmentPath = "input-quality-level"
    inputQualityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputQualityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputQualityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputQualityLevel.EntityData.Children = types.NewOrderedMap()
    inputQualityLevel.EntityData.Leafs = types.NewOrderedMap()
    inputQualityLevel.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", inputQualityLevel.QualityLevelOption})
    inputQualityLevel.EntityData.Leafs.Append("exact-quality-level-value", types.YLeaf{"ExactQualityLevelValue", inputQualityLevel.ExactQualityLevelValue})
    inputQualityLevel.EntityData.Leafs.Append("min-quality-level-value", types.YLeaf{"MinQualityLevelValue", inputQualityLevel.MinQualityLevelValue})
    inputQualityLevel.EntityData.Leafs.Append("max-quality-level-value", types.YLeaf{"MaxQualityLevelValue", inputQualityLevel.MaxQualityLevelValue})

    inputQualityLevel.EntityData.YListKeys = []string {}

    return &(inputQualityLevel.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController
// sync-controller value
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport mode.
    TransportMode ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode
}

func (syncController *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController) GetEntityData() *types.CommonEntityData {
    syncController.EntityData.YFilter = syncController.YFilter
    syncController.EntityData.YangName = "sync-controller"
    syncController.EntityData.BundleName = "cisco_ios_xr"
    syncController.EntityData.ParentYangName = "clock"
    syncController.EntityData.SegmentPath = "Cisco-IOS-XR-syncc-controller-cfg:sync-controller"
    syncController.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncController.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncController.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncController.EntityData.Children = types.NewOrderedMap()
    syncController.EntityData.Children.Append("transport-mode", types.YChild{"TransportMode", &syncController.TransportMode})
    syncController.EntityData.Leafs = types.NewOrderedMap()

    syncController.EntityData.YListKeys = []string {}

    return &(syncController.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode
// Transport mode
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frequency Mode.
    FrequencyMode ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode
}

func (transportMode *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode) GetEntityData() *types.CommonEntityData {
    transportMode.EntityData.YFilter = transportMode.YFilter
    transportMode.EntityData.YangName = "transport-mode"
    transportMode.EntityData.BundleName = "cisco_ios_xr"
    transportMode.EntityData.ParentYangName = "sync-controller"
    transportMode.EntityData.SegmentPath = "transport-mode"
    transportMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportMode.EntityData.Children = types.NewOrderedMap()
    transportMode.EntityData.Children.Append("frequency-mode", types.YChild{"FrequencyMode", &transportMode.FrequencyMode})
    transportMode.EntityData.Leafs = types.NewOrderedMap()

    transportMode.EntityData.YListKeys = []string {}

    return &(transportMode.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode
// Frequency Mode
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable the SyncE Port. The type is interface{}.
    Shutdown interface{}

    // clock-interface sync <value> location <value> port-parameters bits-input 2m
    // -> Option1=0, Option2=2, Option3=0, Option4=0, Option5=0 clock-interface
    // sync <value> location <value> port-parameters bits-input 2m -> Option1=0,
    // Option2=2, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input 64k-input-only -> Option1=0,
    // Option2=3, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa4 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa4 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=0 , Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa5 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa5 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa6 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa6 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa7 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa7 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa8 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa8 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 non-crc-4 ami -> Option1=0,
    // Option2=1, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 non-crc-4 hdb3 -> Option1=0,
    // Option2=1, Option3=0, Option4=1 , Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 d4 ami -> Option1=0,
    // Option2=0, Option3=1, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 d4 b8zs -> Option1=0,
    // Option2=0 , Option3=1, Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 esf ami -> Option1=0,
    // Option2=0, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 esf b8zs -> Option1=0,
    // Option2=0, Option3=0, Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output 2m -> Option1=1 , Option2=2,
    // Option3=0, Option4=0, Option5=0 clock-interface sync <value> location
    // <value> port-parameters bits-output 6m-output-only -> Option1=1 ,
    // Option2=4, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa4 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa4 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa5 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa5 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa6 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa6 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa7 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa7 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa8 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa8 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 non-crc-4 ami -> Option1=1
    // , Option2=1, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 non-crc-4 hdb3 -> Option1=1
    // , Option2=1, Option3=0 , Option4=1, Option5=0clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 0 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 1 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 2 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 3 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 4 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 0 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 1 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 2 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 3 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 4 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 0 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 1 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 2 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 3 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 4 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 0 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 1 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 2 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 3 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 4 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters port-parameters uti -> Option1=2 ,
    // Option2=0, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters port-parameters gps-input -> Option1=3,
    // Option2=0, Option3=0 , Option4=0, Option5=0 .
    PortMode ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode
}

func (frequencyMode *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode) GetEntityData() *types.CommonEntityData {
    frequencyMode.EntityData.YFilter = frequencyMode.YFilter
    frequencyMode.EntityData.YangName = "frequency-mode"
    frequencyMode.EntityData.BundleName = "cisco_ios_xr"
    frequencyMode.EntityData.ParentYangName = "transport-mode"
    frequencyMode.EntityData.SegmentPath = "frequency-mode"
    frequencyMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencyMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencyMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencyMode.EntityData.Children = types.NewOrderedMap()
    frequencyMode.EntityData.Children.Append("port-mode", types.YChild{"PortMode", &frequencyMode.PortMode})
    frequencyMode.EntityData.Leafs = types.NewOrderedMap()
    frequencyMode.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", frequencyMode.Shutdown})

    frequencyMode.EntityData.YListKeys = []string {}

    return &(frequencyMode.EntityData)
}

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode
// clock-interface sync <value> location <value>
// port-parameters bits-input 2m -> Option1=0,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input 2m -> Option1=0,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input 64k-input-only ->
// Option1=0, Option2=3, Option3=0, Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa4 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa4 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=0
// , Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa5 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa5 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa6 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa6 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa7 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa7 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa8 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa8 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=4 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 non-crc-4 ami -> Option1=0, Option2=1,
// Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 non-crc-4 hdb3
// -> Option1=0, Option2=1, Option3=0, Option4=1
// , Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// t1 d4 ami -> Option1=0, Option2=0, Option3=1,
// Option4=0, Option5=0 clock-interface sync
// <value> location <value> port-parameters
// bits-input t1 d4 b8zs -> Option1=0, Option2=0
// , Option3=1, Option4=1, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input t1 esf ami ->
// Option1=0, Option2=0, Option3=0, Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// t1 esf b8zs -> Option1=0, Option2=0,
// Option3=0, Option4=1, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output 2m -> Option1=1 ,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output 6m-output-only ->
// Option1=1 , Option2=4, Option3=0 , Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// e1 crc-4 sa4 ami -> Option1=1 , Option2=1,
// Option3=1 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa4 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=0 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa5 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa5 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=1 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa6 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa6 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=2 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa7 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa7 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=3 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa8 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa8 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=4 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 non-crc-4 ami -> Option1=1 ,
// Option2=1, Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 non-crc-4 hdb3
// -> Option1=1 , Option2=1, Option3=0 ,
// Option4=1, Option5=0clock-interface sync
// <value> location <value> port-parameters
// bits-output t1 d4 ami 0 -> Option1=1 ,
// Option2=0, Option3=1 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 ami 1 ->
// Option1=1 , Option2=0, Option3=1 , Option4=0,
// Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 ami 2 -> Option1=1 , Option2=0,
// Option3=1 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 ami 3 ->
// Option1=1 , Option2=0, Option3=1 , Option4=0,
// Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 ami 4 -> Option1=1 , Option2=0,
// Option3=1 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 0 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 b8zs 1 -> Option1=1 , Option2=0,
// Option3=1 , Option4=1, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 2 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 b8zs 3 -> Option1=1 , Option2=0,
// Option3=1 , Option4=1, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 4 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=4 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 0 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf ami 1 ->
// Option1=1 , Option2=0, Option3=0 , Option4=0,
// Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 2 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf ami 3 ->
// Option1=1 , Option2=0, Option3=0 , Option4=0,
// Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 4 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 0 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf b8zs 1 -> Option1=1 , Option2=0,
// Option3=0 , Option4=1, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 2 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf b8zs 3 -> Option1=1 , Option2=0,
// Option3=0 , Option4=1, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 4 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=4 clock-interface sync <value>
// location <value> port-parameters
// port-parameters uti -> Option1=2 , Option2=0,
// Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters port-parameters gps-input ->
// Option1=3, Option2=0, Option3=0 , Option4=0,
// Option5=0 
// This type is a presence type.
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Option value #1. The type is interface{} with range: 0..3. This attribute
    // is mandatory.
    Option1 interface{}

    // Option value #2. The type is interface{} with range: 0..4. This attribute
    // is mandatory.
    Option2 interface{}

    // Option value #3. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    Option3 interface{}

    // Option value #4. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    Option4 interface{}

    // Option value #5. The type is interface{} with range: 0..4. This attribute
    // is mandatory.
    Option5 interface{}
}

func (portMode *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode) GetEntityData() *types.CommonEntityData {
    portMode.EntityData.YFilter = portMode.YFilter
    portMode.EntityData.YangName = "port-mode"
    portMode.EntityData.BundleName = "cisco_ios_xr"
    portMode.EntityData.ParentYangName = "frequency-mode"
    portMode.EntityData.SegmentPath = "port-mode"
    portMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portMode.EntityData.Children = types.NewOrderedMap()
    portMode.EntityData.Leafs = types.NewOrderedMap()
    portMode.EntityData.Leafs.Append("option1", types.YLeaf{"Option1", portMode.Option1})
    portMode.EntityData.Leafs.Append("option2", types.YLeaf{"Option2", portMode.Option2})
    portMode.EntityData.Leafs.Append("option3", types.YLeaf{"Option3", portMode.Option3})
    portMode.EntityData.Leafs.Append("option4", types.YLeaf{"Option4", portMode.Option4})
    portMode.EntityData.Leafs.Append("option5", types.YLeaf{"Option5", portMode.Option5})

    portMode.EntityData.YListKeys = []string {}

    return &(portMode.EntityData)
}

// ActiveNodes_ActiveNode_SsrpGroup
// Per-node SSRP configuration data
type ActiveNodes_ActiveNode_SsrpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of SSRP Group configuration.
    Groups ActiveNodes_ActiveNode_SsrpGroup_Groups
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetEntityData() *types.CommonEntityData {
    ssrpGroup.EntityData.YFilter = ssrpGroup.YFilter
    ssrpGroup.EntityData.YangName = "ssrp-group"
    ssrpGroup.EntityData.BundleName = "cisco_ios_xr"
    ssrpGroup.EntityData.ParentYangName = "active-node"
    ssrpGroup.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group"
    ssrpGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssrpGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssrpGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssrpGroup.EntityData.Children = types.NewOrderedMap()
    ssrpGroup.EntityData.Children.Append("groups", types.YChild{"Groups", &ssrpGroup.Groups})
    ssrpGroup.EntityData.Leafs = types.NewOrderedMap()

    ssrpGroup.EntityData.YListKeys = []string {}

    return &(ssrpGroup.EntityData)
}

// ActiveNodes_ActiveNode_SsrpGroup_Groups
// Table of SSRP Group configuration
type ActiveNodes_ActiveNode_SsrpGroup_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSRP Group configuration. The type is slice of
    // ActiveNodes_ActiveNode_SsrpGroup_Groups_Group.
    Group []*ActiveNodes_ActiveNode_SsrpGroup_Groups_Group
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "ssrp-group"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// ActiveNodes_ActiveNode_SsrpGroup_Groups_Group
// SSRP Group configuration
type ActiveNodes_ActiveNode_SsrpGroup_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this group. The type is
    // interface{} with range: 1..65535.
    GroupId interface{}

    // This specifies the SSRP profile to use for this group. The type is string.
    Profile interface{}
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", group.Profile})

    group.EntityData.YListKeys = []string {"GroupId"}

    return &(group.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal
// lpts node specific configuration commands
type ActiveNodes_ActiveNode_LptsLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    IpolicerLocalTables ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    DynamicFlowsTables ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    IpolicerLocal ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetEntityData() *types.CommonEntityData {
    lptsLocal.EntityData.YFilter = lptsLocal.YFilter
    lptsLocal.EntityData.YangName = "lpts-local"
    lptsLocal.EntityData.BundleName = "cisco_ios_xr"
    lptsLocal.EntityData.ParentYangName = "active-node"
    lptsLocal.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"
    lptsLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsLocal.EntityData.Children = types.NewOrderedMap()
    lptsLocal.EntityData.Children.Append("ipolicer-local-tables", types.YChild{"IpolicerLocalTables", &lptsLocal.IpolicerLocalTables})
    lptsLocal.EntityData.Children.Append("dynamic-flows-tables", types.YChild{"DynamicFlowsTables", &lptsLocal.DynamicFlowsTables})
    lptsLocal.EntityData.Children.Append("ipolicer-local", types.YChild{"IpolicerLocal", &lptsLocal.IpolicerLocal})
    lptsLocal.EntityData.Leafs = types.NewOrderedMap()

    lptsLocal.EntityData.YListKeys = []string {}

    return &(lptsLocal.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pre IFIB (Internal Forwarding Information Base) configuration table. The
    // type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable.
    IpolicerLocalTable []*ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetEntityData() *types.CommonEntityData {
    ipolicerLocalTables.EntityData.YFilter = ipolicerLocalTables.YFilter
    ipolicerLocalTables.EntityData.YangName = "ipolicer-local-tables"
    ipolicerLocalTables.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocalTables.EntityData.ParentYangName = "lpts-local"
    ipolicerLocalTables.EntityData.SegmentPath = "ipolicer-local-tables"
    ipolicerLocalTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocalTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocalTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocalTables.EntityData.Children = types.NewOrderedMap()
    ipolicerLocalTables.EntityData.Children.Append("ipolicer-local-table", types.YChild{"IpolicerLocalTable", nil})
    for i := range ipolicerLocalTables.IpolicerLocalTable {
        ipolicerLocalTables.EntityData.Children.Append(types.GetSegmentPath(ipolicerLocalTables.IpolicerLocalTable[i]), types.YChild{"IpolicerLocalTable", ipolicerLocalTables.IpolicerLocalTable[i]})
    }
    ipolicerLocalTables.EntityData.Leafs = types.NewOrderedMap()

    ipolicerLocalTables.EntityData.YListKeys = []string {}

    return &(ipolicerLocalTables.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
// Pre IFIB (Internal Forwarding Information
// Base) configuration table
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // NP name.
    NpFlows ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetEntityData() *types.CommonEntityData {
    ipolicerLocalTable.EntityData.YFilter = ipolicerLocalTable.YFilter
    ipolicerLocalTable.EntityData.YangName = "ipolicer-local-table"
    ipolicerLocalTable.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocalTable.EntityData.ParentYangName = "ipolicer-local-tables"
    ipolicerLocalTable.EntityData.SegmentPath = "ipolicer-local-table" + types.AddKeyToken(ipolicerLocalTable.Id1, "id1")
    ipolicerLocalTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocalTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocalTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocalTable.EntityData.Children = types.NewOrderedMap()
    ipolicerLocalTable.EntityData.Children.Append("np-flows", types.YChild{"NpFlows", &ipolicerLocalTable.NpFlows})
    ipolicerLocalTable.EntityData.Leafs = types.NewOrderedMap()
    ipolicerLocalTable.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", ipolicerLocalTable.Id1})

    ipolicerLocalTable.EntityData.YListKeys = []string {"Id1"}

    return &(ipolicerLocalTable.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows
// NP name
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of NP Flow names. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow.
    NpFlow []*ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow
}

func (npFlows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows) GetEntityData() *types.CommonEntityData {
    npFlows.EntityData.YFilter = npFlows.YFilter
    npFlows.EntityData.YangName = "np-flows"
    npFlows.EntityData.BundleName = "cisco_ios_xr"
    npFlows.EntityData.ParentYangName = "ipolicer-local-table"
    npFlows.EntityData.SegmentPath = "np-flows"
    npFlows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npFlows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npFlows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npFlows.EntityData.Children = types.NewOrderedMap()
    npFlows.EntityData.Children.Append("np-flow", types.YChild{"NpFlow", nil})
    for i := range npFlows.NpFlow {
        npFlows.EntityData.Children.Append(types.GetSegmentPath(npFlows.NpFlow[i]), types.YChild{"NpFlow", npFlows.NpFlow[i]})
    }
    npFlows.EntityData.Leafs = types.NewOrderedMap()

    npFlows.EntityData.YListKeys = []string {}

    return &(npFlows.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow
// Table of NP Flow names
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range: 0..4294967295.
    NpRate interface{}
}

func (npFlow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow) GetEntityData() *types.CommonEntityData {
    npFlow.EntityData.YFilter = npFlow.YFilter
    npFlow.EntityData.YangName = "np-flow"
    npFlow.EntityData.BundleName = "cisco_ios_xr"
    npFlow.EntityData.ParentYangName = "np-flows"
    npFlow.EntityData.SegmentPath = "np-flow" + types.AddKeyToken(npFlow.FlowType, "flow-type")
    npFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npFlow.EntityData.Children = types.NewOrderedMap()
    npFlow.EntityData.Leafs = types.NewOrderedMap()
    npFlow.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", npFlow.FlowType})
    npFlow.EntityData.Leafs.Append("np-rate", types.YLeaf{"NpRate", npFlow.NpRate})

    npFlow.EntityData.YListKeys = []string {"FlowType"}

    return &(npFlow.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table for Dynamic Flows. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable.
    DynamicFlowsTable []*ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetEntityData() *types.CommonEntityData {
    dynamicFlowsTables.EntityData.YFilter = dynamicFlowsTables.YFilter
    dynamicFlowsTables.EntityData.YangName = "dynamic-flows-tables"
    dynamicFlowsTables.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsTables.EntityData.ParentYangName = "lpts-local"
    dynamicFlowsTables.EntityData.SegmentPath = "dynamic-flows-tables"
    dynamicFlowsTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsTables.EntityData.Children = types.NewOrderedMap()
    dynamicFlowsTables.EntityData.Children.Append("dynamic-flows-table", types.YChild{"DynamicFlowsTable", nil})
    for i := range dynamicFlowsTables.DynamicFlowsTable {
        dynamicFlowsTables.EntityData.Children.Append(types.GetSegmentPath(dynamicFlowsTables.DynamicFlowsTable[i]), types.YChild{"DynamicFlowsTable", dynamicFlowsTables.DynamicFlowsTable[i]})
    }
    dynamicFlowsTables.EntityData.Leafs = types.NewOrderedMap()

    dynamicFlowsTables.EntityData.YListKeys = []string {}

    return &(dynamicFlowsTables.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
// Table for Dynamic Flows
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Dynamic Flows Table Type. The type is
    // LptsDynamicFlowConfig.
    TableType interface{}

    // Selected flow type. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType.
    FlowType []*ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetEntityData() *types.CommonEntityData {
    dynamicFlowsTable.EntityData.YFilter = dynamicFlowsTable.YFilter
    dynamicFlowsTable.EntityData.YangName = "dynamic-flows-table"
    dynamicFlowsTable.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsTable.EntityData.ParentYangName = "dynamic-flows-tables"
    dynamicFlowsTable.EntityData.SegmentPath = "dynamic-flows-table" + types.AddKeyToken(dynamicFlowsTable.TableType, "table-type")
    dynamicFlowsTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsTable.EntityData.Children = types.NewOrderedMap()
    dynamicFlowsTable.EntityData.Children.Append("flow-type", types.YChild{"FlowType", nil})
    for i := range dynamicFlowsTable.FlowType {
        dynamicFlowsTable.EntityData.Children.Append(types.GetSegmentPath(dynamicFlowsTable.FlowType[i]), types.YChild{"FlowType", dynamicFlowsTable.FlowType[i]})
    }
    dynamicFlowsTable.EntityData.Leafs = types.NewOrderedMap()
    dynamicFlowsTable.EntityData.Leafs.Append("table-type", types.YLeaf{"TableType", dynamicFlowsTable.TableType})

    dynamicFlowsTable.EntityData.YListKeys = []string {"TableType"}

    return &(dynamicFlowsTable.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
// Selected flow type
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured Max TCAM value. The type is interface{} with range:
    // 0..4294967295.
    Max interface{}
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetEntityData() *types.CommonEntityData {
    flowType.EntityData.YFilter = flowType.YFilter
    flowType.EntityData.YangName = "flow-type"
    flowType.EntityData.BundleName = "cisco_ios_xr"
    flowType.EntityData.ParentYangName = "dynamic-flows-table"
    flowType.EntityData.SegmentPath = "flow-type" + types.AddKeyToken(flowType.FlowType, "flow-type")
    flowType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowType.EntityData.Children = types.NewOrderedMap()
    flowType.EntityData.Leafs = types.NewOrderedMap()
    flowType.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", flowType.FlowType})
    flowType.EntityData.Leafs.Append("max", types.YLeaf{"Max", flowType.Max})

    flowType.EntityData.YListKeys = []string {"FlowType"}

    return &(flowType.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
// This type is a presence type.
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for Flows.
    Flows ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetEntityData() *types.CommonEntityData {
    ipolicerLocal.EntityData.YFilter = ipolicerLocal.YFilter
    ipolicerLocal.EntityData.YangName = "ipolicer-local"
    ipolicerLocal.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocal.EntityData.ParentYangName = "lpts-local"
    ipolicerLocal.EntityData.SegmentPath = "ipolicer-local"
    ipolicerLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocal.EntityData.Children = types.NewOrderedMap()
    ipolicerLocal.EntityData.Children.Append("flows", types.YChild{"Flows", &ipolicerLocal.Flows})
    ipolicerLocal.EntityData.Leafs = types.NewOrderedMap()
    ipolicerLocal.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipolicerLocal.Enable})

    ipolicerLocal.EntityData.YListKeys = []string {}

    return &(ipolicerLocal.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows
// Table for Flows
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow.
    Flow []*ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetEntityData() *types.CommonEntityData {
    flows.EntityData.YFilter = flows.YFilter
    flows.EntityData.YangName = "flows"
    flows.EntityData.BundleName = "cisco_ios_xr"
    flows.EntityData.ParentYangName = "ipolicer-local"
    flows.EntityData.SegmentPath = "flows"
    flows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flows.EntityData.Children = types.NewOrderedMap()
    flows.EntityData.Children.Append("flow", types.YChild{"Flow", nil})
    for i := range flows.Flow {
        flows.EntityData.Children.Append(types.GetSegmentPath(flows.Flow[i]), types.YChild{"Flow", flows.Flow[i]})
    }
    flows.EntityData.Leafs = types.NewOrderedMap()

    flows.EntityData.YListKeys = []string {}

    return &(flows.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow
// selected flow type
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "flows"
    flow.EntityData.SegmentPath = "flow" + types.AddKeyToken(flow.FlowType, "flow-type")
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("precedences", types.YChild{"Precedences", &flow.Precedences})
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", flow.FlowType})
    flow.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", flow.Rate})

    flow.EntityData.YListKeys = []string {"FlowType"}

    return &(flow.EntityData)
}

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
// TOS Precedence value(s)
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetEntityData() *types.CommonEntityData {
    precedences.EntityData.YFilter = precedences.YFilter
    precedences.EntityData.YangName = "precedences"
    precedences.EntityData.BundleName = "cisco_ios_xr"
    precedences.EntityData.ParentYangName = "flow"
    precedences.EntityData.SegmentPath = "precedences"
    precedences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    precedences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    precedences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    precedences.EntityData.Children = types.NewOrderedMap()
    precedences.EntityData.Leafs = types.NewOrderedMap()
    precedences.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", precedences.Precedence})

    precedences.EntityData.YListKeys = []string {}

    return &(precedences.EntityData)
}

// PreconfiguredNodes
// preconfigured nodes
type PreconfiguredNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The configuration for a non-active node. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode.
    PreconfiguredNode []*PreconfiguredNodes_PreconfiguredNode
}

func (preconfiguredNodes *PreconfiguredNodes) GetEntityData() *types.CommonEntityData {
    preconfiguredNodes.EntityData.YFilter = preconfiguredNodes.YFilter
    preconfiguredNodes.EntityData.YangName = "preconfigured-nodes"
    preconfiguredNodes.EntityData.BundleName = "cisco_ios_xr"
    preconfiguredNodes.EntityData.ParentYangName = "Cisco-IOS-XR-config-mda-cfg"
    preconfiguredNodes.EntityData.SegmentPath = "Cisco-IOS-XR-config-mda-cfg:preconfigured-nodes"
    preconfiguredNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preconfiguredNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preconfiguredNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preconfiguredNodes.EntityData.Children = types.NewOrderedMap()
    preconfiguredNodes.EntityData.Children.Append("preconfigured-node", types.YChild{"PreconfiguredNode", nil})
    for i := range preconfiguredNodes.PreconfiguredNode {
        preconfiguredNodes.EntityData.Children.Append(types.GetSegmentPath(preconfiguredNodes.PreconfiguredNode[i]), types.YChild{"PreconfiguredNode", preconfiguredNodes.PreconfiguredNode[i]})
    }
    preconfiguredNodes.EntityData.Leafs = types.NewOrderedMap()

    preconfiguredNodes.EntityData.YListKeys = []string {}

    return &(preconfiguredNodes.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode
// The configuration for a non-active node
type PreconfiguredNodes_PreconfiguredNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // watchdog node threshold.
    WatchdogNodeThreshold PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold

    // Ltrace Memory configuration.
    Ltrace PreconfiguredNodes_PreconfiguredNode_Ltrace

    // Configuration for a clock interface.
    ClockInterface PreconfiguredNodes_PreconfiguredNode_ClockInterface

    // lpts node specific configuration commands.
    LptsLocal PreconfiguredNodes_PreconfiguredNode_LptsLocal
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetEntityData() *types.CommonEntityData {
    preconfiguredNode.EntityData.YFilter = preconfiguredNode.YFilter
    preconfiguredNode.EntityData.YangName = "preconfigured-node"
    preconfiguredNode.EntityData.BundleName = "cisco_ios_xr"
    preconfiguredNode.EntityData.ParentYangName = "preconfigured-nodes"
    preconfiguredNode.EntityData.SegmentPath = "preconfigured-node" + types.AddKeyToken(preconfiguredNode.NodeName, "node-name")
    preconfiguredNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preconfiguredNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preconfiguredNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preconfiguredNode.EntityData.Children = types.NewOrderedMap()
    preconfiguredNode.EntityData.Children.Append("Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold", types.YChild{"WatchdogNodeThreshold", &preconfiguredNode.WatchdogNodeThreshold})
    preconfiguredNode.EntityData.Children.Append("Cisco-IOS-XR-infra-ltrace-cfg:ltrace", types.YChild{"Ltrace", &preconfiguredNode.Ltrace})
    preconfiguredNode.EntityData.Children.Append("Cisco-IOS-XR-freqsync-cfg:clock-interface", types.YChild{"ClockInterface", &preconfiguredNode.ClockInterface})
    preconfiguredNode.EntityData.Children.Append("Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local", types.YChild{"LptsLocal", &preconfiguredNode.LptsLocal})
    preconfiguredNode.EntityData.Leafs = types.NewOrderedMap()
    preconfiguredNode.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", preconfiguredNode.NodeName})

    preconfiguredNode.EntityData.YListKeys = []string {"NodeName"}

    return &(preconfiguredNode.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold
// watchdog node threshold
type PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disk thresholds.
    DiskThreshold PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_DiskThreshold

    // Memory thresholds.
    MemoryThreshold PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_MemoryThreshold
}

func (watchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold) GetEntityData() *types.CommonEntityData {
    watchdogNodeThreshold.EntityData.YFilter = watchdogNodeThreshold.YFilter
    watchdogNodeThreshold.EntityData.YangName = "watchdog-node-threshold"
    watchdogNodeThreshold.EntityData.BundleName = "cisco_ios_xr"
    watchdogNodeThreshold.EntityData.ParentYangName = "preconfigured-node"
    watchdogNodeThreshold.EntityData.SegmentPath = "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"
    watchdogNodeThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    watchdogNodeThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    watchdogNodeThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    watchdogNodeThreshold.EntityData.Children = types.NewOrderedMap()
    watchdogNodeThreshold.EntityData.Children.Append("disk-threshold", types.YChild{"DiskThreshold", &watchdogNodeThreshold.DiskThreshold})
    watchdogNodeThreshold.EntityData.Children.Append("memory-threshold", types.YChild{"MemoryThreshold", &watchdogNodeThreshold.MemoryThreshold})
    watchdogNodeThreshold.EntityData.Leafs = types.NewOrderedMap()

    watchdogNodeThreshold.EntityData.YListKeys = []string {}

    return &(watchdogNodeThreshold.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_DiskThreshold
// Disk thresholds
type PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_DiskThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (diskThreshold *PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_DiskThreshold) GetEntityData() *types.CommonEntityData {
    diskThreshold.EntityData.YFilter = diskThreshold.YFilter
    diskThreshold.EntityData.YangName = "disk-threshold"
    diskThreshold.EntityData.BundleName = "cisco_ios_xr"
    diskThreshold.EntityData.ParentYangName = "watchdog-node-threshold"
    diskThreshold.EntityData.SegmentPath = "disk-threshold"
    diskThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diskThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diskThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diskThreshold.EntityData.Children = types.NewOrderedMap()
    diskThreshold.EntityData.Leafs = types.NewOrderedMap()
    diskThreshold.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", diskThreshold.Minor})
    diskThreshold.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", diskThreshold.Severe})
    diskThreshold.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", diskThreshold.Critical})

    diskThreshold.EntityData.YListKeys = []string {}

    return &(diskThreshold.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_MemoryThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_WatchdogNodeThreshold_MemoryThreshold) GetEntityData() *types.CommonEntityData {
    memoryThreshold.EntityData.YFilter = memoryThreshold.YFilter
    memoryThreshold.EntityData.YangName = "memory-threshold"
    memoryThreshold.EntityData.BundleName = "cisco_ios_xr"
    memoryThreshold.EntityData.ParentYangName = "watchdog-node-threshold"
    memoryThreshold.EntityData.SegmentPath = "memory-threshold"
    memoryThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryThreshold.EntityData.Children = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs.Append("minor", types.YLeaf{"Minor", memoryThreshold.Minor})
    memoryThreshold.EntityData.Leafs.Append("severe", types.YLeaf{"Severe", memoryThreshold.Severe})
    memoryThreshold.EntityData.Leafs.Append("critical", types.YLeaf{"Critical", memoryThreshold.Critical})

    memoryThreshold.EntityData.YListKeys = []string {}

    return &(memoryThreshold.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_Ltrace
// Ltrace Memory configuration
type PreconfiguredNodes_PreconfiguredNode_Ltrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select Ltrace mode and scale-factor.
    AllocationParams PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetEntityData() *types.CommonEntityData {
    ltrace.EntityData.YFilter = ltrace.YFilter
    ltrace.EntityData.YangName = "ltrace"
    ltrace.EntityData.BundleName = "cisco_ios_xr"
    ltrace.EntityData.ParentYangName = "preconfigured-node"
    ltrace.EntityData.SegmentPath = "Cisco-IOS-XR-infra-ltrace-cfg:ltrace"
    ltrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ltrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ltrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ltrace.EntityData.Children = types.NewOrderedMap()
    ltrace.EntityData.Children.Append("allocation-params", types.YChild{"AllocationParams", &ltrace.AllocationParams})
    ltrace.EntityData.Leafs = types.NewOrderedMap()

    ltrace.EntityData.YListKeys = []string {}

    return &(ltrace.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams
// Select Ltrace mode and scale-factor
type PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an allocation mode (static:1, dynamic :2). The type is
    // InfraLtraceMode.
    Mode interface{}

    // Select a scaling down factor. The type is InfraLtraceScale.
    ScaleFactor interface{}
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetEntityData() *types.CommonEntityData {
    allocationParams.EntityData.YFilter = allocationParams.YFilter
    allocationParams.EntityData.YangName = "allocation-params"
    allocationParams.EntityData.BundleName = "cisco_ios_xr"
    allocationParams.EntityData.ParentYangName = "ltrace"
    allocationParams.EntityData.SegmentPath = "allocation-params"
    allocationParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allocationParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allocationParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allocationParams.EntityData.Children = types.NewOrderedMap()
    allocationParams.EntityData.Leafs = types.NewOrderedMap()
    allocationParams.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", allocationParams.Mode})
    allocationParams.EntityData.Leafs.Append("scale-factor", types.YLeaf{"ScaleFactor", allocationParams.ScaleFactor})

    allocationParams.EntityData.YListKeys = []string {}

    return &(allocationParams.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a clock interface.
    Clocks PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetEntityData() *types.CommonEntityData {
    clockInterface.EntityData.YFilter = clockInterface.YFilter
    clockInterface.EntityData.YangName = "clock-interface"
    clockInterface.EntityData.BundleName = "cisco_ios_xr"
    clockInterface.EntityData.ParentYangName = "preconfigured-node"
    clockInterface.EntityData.SegmentPath = "Cisco-IOS-XR-freqsync-cfg:clock-interface"
    clockInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterface.EntityData.Children = types.NewOrderedMap()
    clockInterface.EntityData.Children.Append("clocks", types.YChild{"Clocks", &clockInterface.Clocks})
    clockInterface.EntityData.Leafs = types.NewOrderedMap()

    clockInterface.EntityData.YListKeys = []string {}

    return &(clockInterface.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a clock interface. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock.
    Clock []*PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetEntityData() *types.CommonEntityData {
    clocks.EntityData.YFilter = clocks.YFilter
    clocks.EntityData.YangName = "clocks"
    clocks.EntityData.BundleName = "cisco_ios_xr"
    clocks.EntityData.ParentYangName = "clock-interface"
    clocks.EntityData.SegmentPath = "clocks"
    clocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clocks.EntityData.Children = types.NewOrderedMap()
    clocks.EntityData.Children.Append("clock", types.YChild{"Clock", nil})
    for i := range clocks.Clock {
        clocks.EntityData.Children.Append(types.GetSegmentPath(clocks.Clock[i]), types.YChild{"Clock", clocks.Clock[i]})
    }
    clocks.EntityData.Leafs = types.NewOrderedMap()

    clocks.EntityData.YListKeys = []string {}

    return &(clocks.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // 0..4294967295.
    Port interface{}

    // Frequency Synchronization clock configuraiton.
    FrequencySynchronization PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization

    // sync-controller value.
    SyncController PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetEntityData() *types.CommonEntityData {
    clock.EntityData.YFilter = clock.YFilter
    clock.EntityData.YangName = "clock"
    clock.EntityData.BundleName = "cisco_ios_xr"
    clock.EntityData.ParentYangName = "clocks"
    clock.EntityData.SegmentPath = "clock" + types.AddKeyToken(clock.ClockType, "clock-type") + types.AddKeyToken(clock.Port, "port")
    clock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clock.EntityData.Children = types.NewOrderedMap()
    clock.EntityData.Children.Append("frequency-synchronization", types.YChild{"FrequencySynchronization", &clock.FrequencySynchronization})
    clock.EntityData.Children.Append("Cisco-IOS-XR-syncc-controller-cfg:sync-controller", types.YChild{"SyncController", &clock.SyncController})
    clock.EntityData.Leafs = types.NewOrderedMap()
    clock.EntityData.Leafs.Append("clock-type", types.YLeaf{"ClockType", clock.ClockType})
    clock.EntityData.Leafs.Append("port", types.YLeaf{"Port", clock.Port})

    clock.EntityData.YListKeys = []string {"ClockType", "Port"}

    return &(clock.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization
// Frequency Synchronization clock configuraiton
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the wait-to-restore time for this source. The type is interface{} with
    // range: 0..12. The default value is 5.
    WaitToRestoreTime interface{}

    // Set the priority of this source. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // Assign this source as a selection input. The type is interface{}.
    SelectionInput interface{}

    // Set the time-of-day priority of this source. The type is interface{} with
    // range: 1..254. The default value is 100.
    TimeOfDayPriority interface{}

    // Disable SSM on this source. The type is interface{}.
    SsmDisable interface{}

    // Set the output quality level.
    OutputQualityLevel PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel

    // Set the input quality level.
    InputQualityLevel PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetEntityData() *types.CommonEntityData {
    frequencySynchronization.EntityData.YFilter = frequencySynchronization.YFilter
    frequencySynchronization.EntityData.YangName = "frequency-synchronization"
    frequencySynchronization.EntityData.BundleName = "cisco_ios_xr"
    frequencySynchronization.EntityData.ParentYangName = "clock"
    frequencySynchronization.EntityData.SegmentPath = "frequency-synchronization"
    frequencySynchronization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySynchronization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySynchronization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySynchronization.EntityData.Children = types.NewOrderedMap()
    frequencySynchronization.EntityData.Children.Append("output-quality-level", types.YChild{"OutputQualityLevel", &frequencySynchronization.OutputQualityLevel})
    frequencySynchronization.EntityData.Children.Append("input-quality-level", types.YChild{"InputQualityLevel", &frequencySynchronization.InputQualityLevel})
    frequencySynchronization.EntityData.Leafs = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs.Append("wait-to-restore-time", types.YLeaf{"WaitToRestoreTime", frequencySynchronization.WaitToRestoreTime})
    frequencySynchronization.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", frequencySynchronization.Priority})
    frequencySynchronization.EntityData.Leafs.Append("selection-input", types.YLeaf{"SelectionInput", frequencySynchronization.SelectionInput})
    frequencySynchronization.EntityData.Leafs.Append("time-of-day-priority", types.YLeaf{"TimeOfDayPriority", frequencySynchronization.TimeOfDayPriority})
    frequencySynchronization.EntityData.Leafs.Append("ssm-disable", types.YLeaf{"SsmDisable", frequencySynchronization.SsmDisable})

    frequencySynchronization.EntityData.YListKeys = []string {}

    return &(frequencySynchronization.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
// Set the output quality level
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel struct {
    EntityData types.CommonEntityData
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

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetEntityData() *types.CommonEntityData {
    outputQualityLevel.EntityData.YFilter = outputQualityLevel.YFilter
    outputQualityLevel.EntityData.YangName = "output-quality-level"
    outputQualityLevel.EntityData.BundleName = "cisco_ios_xr"
    outputQualityLevel.EntityData.ParentYangName = "frequency-synchronization"
    outputQualityLevel.EntityData.SegmentPath = "output-quality-level"
    outputQualityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputQualityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputQualityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputQualityLevel.EntityData.Children = types.NewOrderedMap()
    outputQualityLevel.EntityData.Leafs = types.NewOrderedMap()
    outputQualityLevel.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", outputQualityLevel.QualityLevelOption})
    outputQualityLevel.EntityData.Leafs.Append("exact-quality-level-value", types.YLeaf{"ExactQualityLevelValue", outputQualityLevel.ExactQualityLevelValue})
    outputQualityLevel.EntityData.Leafs.Append("min-quality-level-value", types.YLeaf{"MinQualityLevelValue", outputQualityLevel.MinQualityLevelValue})
    outputQualityLevel.EntityData.Leafs.Append("max-quality-level-value", types.YLeaf{"MaxQualityLevelValue", outputQualityLevel.MaxQualityLevelValue})

    outputQualityLevel.EntityData.YListKeys = []string {}

    return &(outputQualityLevel.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
// Set the input quality level
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel struct {
    EntityData types.CommonEntityData
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

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetEntityData() *types.CommonEntityData {
    inputQualityLevel.EntityData.YFilter = inputQualityLevel.YFilter
    inputQualityLevel.EntityData.YangName = "input-quality-level"
    inputQualityLevel.EntityData.BundleName = "cisco_ios_xr"
    inputQualityLevel.EntityData.ParentYangName = "frequency-synchronization"
    inputQualityLevel.EntityData.SegmentPath = "input-quality-level"
    inputQualityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputQualityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputQualityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputQualityLevel.EntityData.Children = types.NewOrderedMap()
    inputQualityLevel.EntityData.Leafs = types.NewOrderedMap()
    inputQualityLevel.EntityData.Leafs.Append("quality-level-option", types.YLeaf{"QualityLevelOption", inputQualityLevel.QualityLevelOption})
    inputQualityLevel.EntityData.Leafs.Append("exact-quality-level-value", types.YLeaf{"ExactQualityLevelValue", inputQualityLevel.ExactQualityLevelValue})
    inputQualityLevel.EntityData.Leafs.Append("min-quality-level-value", types.YLeaf{"MinQualityLevelValue", inputQualityLevel.MinQualityLevelValue})
    inputQualityLevel.EntityData.Leafs.Append("max-quality-level-value", types.YLeaf{"MaxQualityLevelValue", inputQualityLevel.MaxQualityLevelValue})

    inputQualityLevel.EntityData.YListKeys = []string {}

    return &(inputQualityLevel.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController
// sync-controller value
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport mode.
    TransportMode PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode
}

func (syncController *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController) GetEntityData() *types.CommonEntityData {
    syncController.EntityData.YFilter = syncController.YFilter
    syncController.EntityData.YangName = "sync-controller"
    syncController.EntityData.BundleName = "cisco_ios_xr"
    syncController.EntityData.ParentYangName = "clock"
    syncController.EntityData.SegmentPath = "Cisco-IOS-XR-syncc-controller-cfg:sync-controller"
    syncController.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syncController.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syncController.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syncController.EntityData.Children = types.NewOrderedMap()
    syncController.EntityData.Children.Append("transport-mode", types.YChild{"TransportMode", &syncController.TransportMode})
    syncController.EntityData.Leafs = types.NewOrderedMap()

    syncController.EntityData.YListKeys = []string {}

    return &(syncController.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode
// Transport mode
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Frequency Mode.
    FrequencyMode PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode
}

func (transportMode *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode) GetEntityData() *types.CommonEntityData {
    transportMode.EntityData.YFilter = transportMode.YFilter
    transportMode.EntityData.YangName = "transport-mode"
    transportMode.EntityData.BundleName = "cisco_ios_xr"
    transportMode.EntityData.ParentYangName = "sync-controller"
    transportMode.EntityData.SegmentPath = "transport-mode"
    transportMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportMode.EntityData.Children = types.NewOrderedMap()
    transportMode.EntityData.Children.Append("frequency-mode", types.YChild{"FrequencyMode", &transportMode.FrequencyMode})
    transportMode.EntityData.Leafs = types.NewOrderedMap()

    transportMode.EntityData.YListKeys = []string {}

    return &(transportMode.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode
// Frequency Mode
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable the SyncE Port. The type is interface{}.
    Shutdown interface{}

    // clock-interface sync <value> location <value> port-parameters bits-input 2m
    // -> Option1=0, Option2=2, Option3=0, Option4=0, Option5=0 clock-interface
    // sync <value> location <value> port-parameters bits-input 2m -> Option1=0,
    // Option2=2, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input 64k-input-only -> Option1=0,
    // Option2=3, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa4 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa4 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=0 , Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa5 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa5 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa6 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa6 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa7 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa7 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa8 ami -> Option1=0,
    // Option2=1, Option3=1, Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 crc-4 sa8 hdb3 -> Option1=0,
    // Option2=1, Option3=1, Option4=1 , Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 non-crc-4 ami -> Option1=0,
    // Option2=1, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input e1 non-crc-4 hdb3 -> Option1=0,
    // Option2=1, Option3=0, Option4=1 , Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 d4 ami -> Option1=0,
    // Option2=0, Option3=1, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 d4 b8zs -> Option1=0,
    // Option2=0 , Option3=1, Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 esf ami -> Option1=0,
    // Option2=0, Option3=0, Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-input t1 esf b8zs -> Option1=0,
    // Option2=0, Option3=0, Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output 2m -> Option1=1 , Option2=2,
    // Option3=0, Option4=0, Option5=0 clock-interface sync <value> location
    // <value> port-parameters bits-output 6m-output-only -> Option1=1 ,
    // Option2=4, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa4 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa4 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa5 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa5 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa6 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa6 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa7 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa7 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa8 ami -> Option1=1
    // , Option2=1, Option3=1 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 crc-4 sa8 hdb3 -> Option1=1
    // , Option2=1, Option3=1 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 non-crc-4 ami -> Option1=1
    // , Option2=1, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output e1 non-crc-4 hdb3 -> Option1=1
    // , Option2=1, Option3=0 , Option4=1, Option5=0clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 0 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 1 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 2 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 3 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 ami 4 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 0 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 1 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 2 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 3 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 d4 b8zs 4 -> Option1=1 ,
    // Option2=0, Option3=1 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 0 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 1 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 2 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 3 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf ami 4 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=0, Option5=4 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 0 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=0 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 1 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=1 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 2 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=2 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 3 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=3 clock-interface sync <value>
    // location <value> port-parameters bits-output t1 esf b8zs 4 -> Option1=1 ,
    // Option2=0, Option3=0 , Option4=1, Option5=4 clock-interface sync <value>
    // location <value> port-parameters port-parameters uti -> Option1=2 ,
    // Option2=0, Option3=0 , Option4=0, Option5=0 clock-interface sync <value>
    // location <value> port-parameters port-parameters gps-input -> Option1=3,
    // Option2=0, Option3=0 , Option4=0, Option5=0 .
    PortMode PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode
}

func (frequencyMode *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode) GetEntityData() *types.CommonEntityData {
    frequencyMode.EntityData.YFilter = frequencyMode.YFilter
    frequencyMode.EntityData.YangName = "frequency-mode"
    frequencyMode.EntityData.BundleName = "cisco_ios_xr"
    frequencyMode.EntityData.ParentYangName = "transport-mode"
    frequencyMode.EntityData.SegmentPath = "frequency-mode"
    frequencyMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencyMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencyMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencyMode.EntityData.Children = types.NewOrderedMap()
    frequencyMode.EntityData.Children.Append("port-mode", types.YChild{"PortMode", &frequencyMode.PortMode})
    frequencyMode.EntityData.Leafs = types.NewOrderedMap()
    frequencyMode.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", frequencyMode.Shutdown})

    frequencyMode.EntityData.YListKeys = []string {}

    return &(frequencyMode.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode
// clock-interface sync <value> location <value>
// port-parameters bits-input 2m -> Option1=0,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input 2m -> Option1=0,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input 64k-input-only ->
// Option1=0, Option2=3, Option3=0, Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa4 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa4 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=0
// , Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa5 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa5 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa6 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa6 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa7 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa7 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 crc-4 sa8 ami -> Option1=0, Option2=1,
// Option3=1, Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 crc-4 sa8 hdb3
// -> Option1=0, Option2=1, Option3=1, Option4=1
// , Option5=4 clock-interface sync <value>
// location <value> port-parameters bits-input
// e1 non-crc-4 ami -> Option1=0, Option2=1,
// Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input e1 non-crc-4 hdb3
// -> Option1=0, Option2=1, Option3=0, Option4=1
// , Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// t1 d4 ami -> Option1=0, Option2=0, Option3=1,
// Option4=0, Option5=0 clock-interface sync
// <value> location <value> port-parameters
// bits-input t1 d4 b8zs -> Option1=0, Option2=0
// , Option3=1, Option4=1, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-input t1 esf ami ->
// Option1=0, Option2=0, Option3=0, Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-input
// t1 esf b8zs -> Option1=0, Option2=0,
// Option3=0, Option4=1, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output 2m -> Option1=1 ,
// Option2=2, Option3=0, Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output 6m-output-only ->
// Option1=1 , Option2=4, Option3=0 , Option4=0,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// e1 crc-4 sa4 ami -> Option1=1 , Option2=1,
// Option3=1 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa4 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=0 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa5 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa5 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=1 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa6 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa6 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=2 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa7 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa7 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=3 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 crc-4 sa8 ami -> Option1=1 ,
// Option2=1, Option3=1 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 crc-4 sa8 hdb3
// -> Option1=1 , Option2=1, Option3=1 ,
// Option4=1, Option5=4 clock-interface sync
// <value> location <value> port-parameters
// bits-output e1 non-crc-4 ami -> Option1=1 ,
// Option2=1, Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output e1 non-crc-4 hdb3
// -> Option1=1 , Option2=1, Option3=0 ,
// Option4=1, Option5=0clock-interface sync
// <value> location <value> port-parameters
// bits-output t1 d4 ami 0 -> Option1=1 ,
// Option2=0, Option3=1 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 ami 1 ->
// Option1=1 , Option2=0, Option3=1 , Option4=0,
// Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 ami 2 -> Option1=1 , Option2=0,
// Option3=1 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 ami 3 ->
// Option1=1 , Option2=0, Option3=1 , Option4=0,
// Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 ami 4 -> Option1=1 , Option2=0,
// Option3=1 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 0 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 b8zs 1 -> Option1=1 , Option2=0,
// Option3=1 , Option4=1, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 2 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 d4 b8zs 3 -> Option1=1 , Option2=0,
// Option3=1 , Option4=1, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 d4 b8zs 4 ->
// Option1=1 , Option2=0, Option3=1 , Option4=1,
// Option5=4 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 0 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf ami 1 ->
// Option1=1 , Option2=0, Option3=0 , Option4=0,
// Option5=1 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 2 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=2
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf ami 3 ->
// Option1=1 , Option2=0, Option3=0 , Option4=0,
// Option5=3 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf ami 4 -> Option1=1 , Option2=0,
// Option3=0 , Option4=0, Option5=4
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 0 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=0 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf b8zs 1 -> Option1=1 , Option2=0,
// Option3=0 , Option4=1, Option5=1
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 2 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=2 clock-interface sync <value>
// location <value> port-parameters bits-output
// t1 esf b8zs 3 -> Option1=1 , Option2=0,
// Option3=0 , Option4=1, Option5=3
// clock-interface sync <value> location <value>
// port-parameters bits-output t1 esf b8zs 4 ->
// Option1=1 , Option2=0, Option3=0 , Option4=1,
// Option5=4 clock-interface sync <value>
// location <value> port-parameters
// port-parameters uti -> Option1=2 , Option2=0,
// Option3=0 , Option4=0, Option5=0
// clock-interface sync <value> location <value>
// port-parameters port-parameters gps-input ->
// Option1=3, Option2=0, Option3=0 , Option4=0,
// Option5=0 
// This type is a presence type.
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Option value #1. The type is interface{} with range: 0..3. This attribute
    // is mandatory.
    Option1 interface{}

    // Option value #2. The type is interface{} with range: 0..4. This attribute
    // is mandatory.
    Option2 interface{}

    // Option value #3. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    Option3 interface{}

    // Option value #4. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    Option4 interface{}

    // Option value #5. The type is interface{} with range: 0..4. This attribute
    // is mandatory.
    Option5 interface{}
}

func (portMode *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_SyncController_TransportMode_FrequencyMode_PortMode) GetEntityData() *types.CommonEntityData {
    portMode.EntityData.YFilter = portMode.YFilter
    portMode.EntityData.YangName = "port-mode"
    portMode.EntityData.BundleName = "cisco_ios_xr"
    portMode.EntityData.ParentYangName = "frequency-mode"
    portMode.EntityData.SegmentPath = "port-mode"
    portMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portMode.EntityData.Children = types.NewOrderedMap()
    portMode.EntityData.Leafs = types.NewOrderedMap()
    portMode.EntityData.Leafs.Append("option1", types.YLeaf{"Option1", portMode.Option1})
    portMode.EntityData.Leafs.Append("option2", types.YLeaf{"Option2", portMode.Option2})
    portMode.EntityData.Leafs.Append("option3", types.YLeaf{"Option3", portMode.Option3})
    portMode.EntityData.Leafs.Append("option4", types.YLeaf{"Option4", portMode.Option4})
    portMode.EntityData.Leafs.Append("option5", types.YLeaf{"Option5", portMode.Option5})

    portMode.EntityData.YListKeys = []string {}

    return &(portMode.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal
// lpts node specific configuration commands
type PreconfiguredNodes_PreconfiguredNode_LptsLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    IpolicerLocalTables PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    DynamicFlowsTables PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables

    // Node specific Pre IFIB (Internal Forwarding Information Base)
    // Configuration.
    IpolicerLocal PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetEntityData() *types.CommonEntityData {
    lptsLocal.EntityData.YFilter = lptsLocal.YFilter
    lptsLocal.EntityData.YangName = "lpts-local"
    lptsLocal.EntityData.BundleName = "cisco_ios_xr"
    lptsLocal.EntityData.ParentYangName = "preconfigured-node"
    lptsLocal.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"
    lptsLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsLocal.EntityData.Children = types.NewOrderedMap()
    lptsLocal.EntityData.Children.Append("ipolicer-local-tables", types.YChild{"IpolicerLocalTables", &lptsLocal.IpolicerLocalTables})
    lptsLocal.EntityData.Children.Append("dynamic-flows-tables", types.YChild{"DynamicFlowsTables", &lptsLocal.DynamicFlowsTables})
    lptsLocal.EntityData.Children.Append("ipolicer-local", types.YChild{"IpolicerLocal", &lptsLocal.IpolicerLocal})
    lptsLocal.EntityData.Leafs = types.NewOrderedMap()

    lptsLocal.EntityData.YListKeys = []string {}

    return &(lptsLocal.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pre IFIB (Internal Forwarding Information Base) configuration table. The
    // type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable.
    IpolicerLocalTable []*PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetEntityData() *types.CommonEntityData {
    ipolicerLocalTables.EntityData.YFilter = ipolicerLocalTables.YFilter
    ipolicerLocalTables.EntityData.YangName = "ipolicer-local-tables"
    ipolicerLocalTables.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocalTables.EntityData.ParentYangName = "lpts-local"
    ipolicerLocalTables.EntityData.SegmentPath = "ipolicer-local-tables"
    ipolicerLocalTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocalTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocalTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocalTables.EntityData.Children = types.NewOrderedMap()
    ipolicerLocalTables.EntityData.Children.Append("ipolicer-local-table", types.YChild{"IpolicerLocalTable", nil})
    for i := range ipolicerLocalTables.IpolicerLocalTable {
        ipolicerLocalTables.EntityData.Children.Append(types.GetSegmentPath(ipolicerLocalTables.IpolicerLocalTable[i]), types.YChild{"IpolicerLocalTable", ipolicerLocalTables.IpolicerLocalTable[i]})
    }
    ipolicerLocalTables.EntityData.Leafs = types.NewOrderedMap()

    ipolicerLocalTables.EntityData.YListKeys = []string {}

    return &(ipolicerLocalTables.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
// Pre IFIB (Internal Forwarding Information
// Base) configuration table
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // NP name.
    NpFlows PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetEntityData() *types.CommonEntityData {
    ipolicerLocalTable.EntityData.YFilter = ipolicerLocalTable.YFilter
    ipolicerLocalTable.EntityData.YangName = "ipolicer-local-table"
    ipolicerLocalTable.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocalTable.EntityData.ParentYangName = "ipolicer-local-tables"
    ipolicerLocalTable.EntityData.SegmentPath = "ipolicer-local-table" + types.AddKeyToken(ipolicerLocalTable.Id1, "id1")
    ipolicerLocalTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocalTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocalTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocalTable.EntityData.Children = types.NewOrderedMap()
    ipolicerLocalTable.EntityData.Children.Append("np-flows", types.YChild{"NpFlows", &ipolicerLocalTable.NpFlows})
    ipolicerLocalTable.EntityData.Leafs = types.NewOrderedMap()
    ipolicerLocalTable.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", ipolicerLocalTable.Id1})

    ipolicerLocalTable.EntityData.YListKeys = []string {"Id1"}

    return &(ipolicerLocalTable.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows
// NP name
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of NP Flow names. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow.
    NpFlow []*PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow
}

func (npFlows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows) GetEntityData() *types.CommonEntityData {
    npFlows.EntityData.YFilter = npFlows.YFilter
    npFlows.EntityData.YangName = "np-flows"
    npFlows.EntityData.BundleName = "cisco_ios_xr"
    npFlows.EntityData.ParentYangName = "ipolicer-local-table"
    npFlows.EntityData.SegmentPath = "np-flows"
    npFlows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npFlows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npFlows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npFlows.EntityData.Children = types.NewOrderedMap()
    npFlows.EntityData.Children.Append("np-flow", types.YChild{"NpFlow", nil})
    for i := range npFlows.NpFlow {
        npFlows.EntityData.Children.Append(types.GetSegmentPath(npFlows.NpFlow[i]), types.YChild{"NpFlow", npFlows.NpFlow[i]})
    }
    npFlows.EntityData.Leafs = types.NewOrderedMap()

    npFlows.EntityData.YListKeys = []string {}

    return &(npFlows.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow
// Table of NP Flow names
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range: 0..4294967295.
    NpRate interface{}
}

func (npFlow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_NpFlows_NpFlow) GetEntityData() *types.CommonEntityData {
    npFlow.EntityData.YFilter = npFlow.YFilter
    npFlow.EntityData.YangName = "np-flow"
    npFlow.EntityData.BundleName = "cisco_ios_xr"
    npFlow.EntityData.ParentYangName = "np-flows"
    npFlow.EntityData.SegmentPath = "np-flow" + types.AddKeyToken(npFlow.FlowType, "flow-type")
    npFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    npFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    npFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    npFlow.EntityData.Children = types.NewOrderedMap()
    npFlow.EntityData.Leafs = types.NewOrderedMap()
    npFlow.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", npFlow.FlowType})
    npFlow.EntityData.Leafs.Append("np-rate", types.YLeaf{"NpRate", npFlow.NpRate})

    npFlow.EntityData.YListKeys = []string {"FlowType"}

    return &(npFlow.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table for Dynamic Flows. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable.
    DynamicFlowsTable []*PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetEntityData() *types.CommonEntityData {
    dynamicFlowsTables.EntityData.YFilter = dynamicFlowsTables.YFilter
    dynamicFlowsTables.EntityData.YangName = "dynamic-flows-tables"
    dynamicFlowsTables.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsTables.EntityData.ParentYangName = "lpts-local"
    dynamicFlowsTables.EntityData.SegmentPath = "dynamic-flows-tables"
    dynamicFlowsTables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsTables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsTables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsTables.EntityData.Children = types.NewOrderedMap()
    dynamicFlowsTables.EntityData.Children.Append("dynamic-flows-table", types.YChild{"DynamicFlowsTable", nil})
    for i := range dynamicFlowsTables.DynamicFlowsTable {
        dynamicFlowsTables.EntityData.Children.Append(types.GetSegmentPath(dynamicFlowsTables.DynamicFlowsTable[i]), types.YChild{"DynamicFlowsTable", dynamicFlowsTables.DynamicFlowsTable[i]})
    }
    dynamicFlowsTables.EntityData.Leafs = types.NewOrderedMap()

    dynamicFlowsTables.EntityData.YListKeys = []string {}

    return &(dynamicFlowsTables.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
// Table for Dynamic Flows
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Dynamic Flows Table Type. The type is
    // LptsDynamicFlowConfig.
    TableType interface{}

    // Selected flow type. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType.
    FlowType []*PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetEntityData() *types.CommonEntityData {
    dynamicFlowsTable.EntityData.YFilter = dynamicFlowsTable.YFilter
    dynamicFlowsTable.EntityData.YangName = "dynamic-flows-table"
    dynamicFlowsTable.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsTable.EntityData.ParentYangName = "dynamic-flows-tables"
    dynamicFlowsTable.EntityData.SegmentPath = "dynamic-flows-table" + types.AddKeyToken(dynamicFlowsTable.TableType, "table-type")
    dynamicFlowsTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsTable.EntityData.Children = types.NewOrderedMap()
    dynamicFlowsTable.EntityData.Children.Append("flow-type", types.YChild{"FlowType", nil})
    for i := range dynamicFlowsTable.FlowType {
        dynamicFlowsTable.EntityData.Children.Append(types.GetSegmentPath(dynamicFlowsTable.FlowType[i]), types.YChild{"FlowType", dynamicFlowsTable.FlowType[i]})
    }
    dynamicFlowsTable.EntityData.Leafs = types.NewOrderedMap()
    dynamicFlowsTable.EntityData.Leafs.Append("table-type", types.YLeaf{"TableType", dynamicFlowsTable.TableType})

    dynamicFlowsTable.EntityData.YListKeys = []string {"TableType"}

    return &(dynamicFlowsTable.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
// Selected flow type
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured Max TCAM value. The type is interface{} with range:
    // 0..4294967295.
    Max interface{}
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetEntityData() *types.CommonEntityData {
    flowType.EntityData.YFilter = flowType.YFilter
    flowType.EntityData.YangName = "flow-type"
    flowType.EntityData.BundleName = "cisco_ios_xr"
    flowType.EntityData.ParentYangName = "dynamic-flows-table"
    flowType.EntityData.SegmentPath = "flow-type" + types.AddKeyToken(flowType.FlowType, "flow-type")
    flowType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowType.EntityData.Children = types.NewOrderedMap()
    flowType.EntityData.Leafs = types.NewOrderedMap()
    flowType.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", flowType.FlowType})
    flowType.EntityData.Leafs.Append("max", types.YLeaf{"Max", flowType.Max})

    flowType.EntityData.YListKeys = []string {"FlowType"}

    return &(flowType.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
// This type is a presence type.
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for Flows.
    Flows PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetEntityData() *types.CommonEntityData {
    ipolicerLocal.EntityData.YFilter = ipolicerLocal.YFilter
    ipolicerLocal.EntityData.YangName = "ipolicer-local"
    ipolicerLocal.EntityData.BundleName = "cisco_ios_xr"
    ipolicerLocal.EntityData.ParentYangName = "lpts-local"
    ipolicerLocal.EntityData.SegmentPath = "ipolicer-local"
    ipolicerLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipolicerLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipolicerLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipolicerLocal.EntityData.Children = types.NewOrderedMap()
    ipolicerLocal.EntityData.Children.Append("flows", types.YChild{"Flows", &ipolicerLocal.Flows})
    ipolicerLocal.EntityData.Leafs = types.NewOrderedMap()
    ipolicerLocal.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipolicerLocal.Enable})

    ipolicerLocal.EntityData.YListKeys = []string {}

    return &(ipolicerLocal.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows
// Table for Flows
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow.
    Flow []*PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetEntityData() *types.CommonEntityData {
    flows.EntityData.YFilter = flows.YFilter
    flows.EntityData.YangName = "flows"
    flows.EntityData.BundleName = "cisco_ios_xr"
    flows.EntityData.ParentYangName = "ipolicer-local"
    flows.EntityData.SegmentPath = "flows"
    flows.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flows.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flows.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flows.EntityData.Children = types.NewOrderedMap()
    flows.EntityData.Children.Append("flow", types.YChild{"Flow", nil})
    for i := range flows.Flow {
        flows.EntityData.Children.Append(types.GetSegmentPath(flows.Flow[i]), types.YChild{"Flow", flows.Flow[i]})
    }
    flows.EntityData.Leafs = types.NewOrderedMap()

    flows.EntityData.YListKeys = []string {}

    return &(flows.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow
// selected flow type
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range: 0..4294967295.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "flows"
    flow.EntityData.SegmentPath = "flow" + types.AddKeyToken(flow.FlowType, "flow-type")
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("precedences", types.YChild{"Precedences", &flow.Precedences})
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", flow.FlowType})
    flow.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", flow.Rate})

    flow.EntityData.YListKeys = []string {"FlowType"}

    return &(flow.EntityData)
}

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
// TOS Precedence value(s)
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetEntityData() *types.CommonEntityData {
    precedences.EntityData.YFilter = precedences.YFilter
    precedences.EntityData.YangName = "precedences"
    precedences.EntityData.BundleName = "cisco_ios_xr"
    precedences.EntityData.ParentYangName = "flow"
    precedences.EntityData.SegmentPath = "precedences"
    precedences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    precedences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    precedences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    precedences.EntityData.Children = types.NewOrderedMap()
    precedences.EntityData.Leafs = types.NewOrderedMap()
    precedences.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", precedences.Precedence})

    precedences.EntityData.YListKeys = []string {}

    return &(precedences.EntityData)
}

