// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-mda package configuration.
// 
// This module contains definitions
// for the following management objects:
//   active-nodes: Per-node configuration for active nodes
//   preconfigured-nodes: preconfigured nodes
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // The configuration for an active node. The type is slice of
    // ActiveNodes_ActiveNode.
    ActiveNode []ActiveNodes_ActiveNode
}

func (activeNodes *ActiveNodes) GetFilter() yfilter.YFilter { return activeNodes.YFilter }

func (activeNodes *ActiveNodes) SetFilter(yf yfilter.YFilter) { activeNodes.YFilter = yf }

func (activeNodes *ActiveNodes) GetGoName(yname string) string {
    if yname == "active-node" { return "ActiveNode" }
    return ""
}

func (activeNodes *ActiveNodes) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-mda-cfg:active-nodes"
}

func (activeNodes *ActiveNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-node" {
        for _, c := range activeNodes.ActiveNode {
            if activeNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode{}
        activeNodes.ActiveNode = append(activeNodes.ActiveNode, child)
        return &activeNodes.ActiveNode[len(activeNodes.ActiveNode)-1]
    }
    return nil
}

func (activeNodes *ActiveNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range activeNodes.ActiveNode {
        children[activeNodes.ActiveNode[i].GetSegmentPath()] = &activeNodes.ActiveNode[i]
    }
    return children
}

func (activeNodes *ActiveNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (activeNodes *ActiveNodes) GetBundleName() string { return "cisco_ios_xr" }

func (activeNodes *ActiveNodes) GetYangName() string { return "active-nodes" }

func (activeNodes *ActiveNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeNodes *ActiveNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeNodes *ActiveNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeNodes *ActiveNodes) SetParent(parent types.Entity) { activeNodes.parent = parent }

func (activeNodes *ActiveNodes) GetParent() types.Entity { return activeNodes.parent }

func (activeNodes *ActiveNodes) GetParentYangName() string { return "Cisco-IOS-XR-config-mda-cfg" }

// ActiveNodes_ActiveNode
// The configuration for an active node
type ActiveNodes_ActiveNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Configuration for a clock interface.
    ClockInterface ActiveNodes_ActiveNode_ClockInterface

    // Ltrace Memory configuration.
    Ltrace ActiveNodes_ActiveNode_Ltrace

    // lpts node specific configuration commands.
    LptsLocal ActiveNodes_ActiveNode_LptsLocal

    // Per-node SSRP configuration data.
    SsrpGroup ActiveNodes_ActiveNode_SsrpGroup

    // watchdog node threshold.
    CiscoIosXrWatchdCfgWatchdogNodeThreshold ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold

    // Watchdog threshold configuration.
    CiscoIosXrWdCfgWatchdogNodeThreshold ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold
}

func (activeNode *ActiveNodes_ActiveNode) GetFilter() yfilter.YFilter { return activeNode.YFilter }

func (activeNode *ActiveNodes_ActiveNode) SetFilter(yf yfilter.YFilter) { activeNode.YFilter = yf }

func (activeNode *ActiveNodes_ActiveNode) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "Cisco-IOS-XR-freqsync-cfg:clock-interface" { return "ClockInterface" }
    if yname == "Cisco-IOS-XR-infra-ltrace-cfg:ltrace" { return "Ltrace" }
    if yname == "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local" { return "LptsLocal" }
    if yname == "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group" { return "SsrpGroup" }
    if yname == "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold" { return "CiscoIosXrWatchdCfgWatchdogNodeThreshold" }
    if yname == "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold" { return "CiscoIosXrWdCfgWatchdogNodeThreshold" }
    return ""
}

func (activeNode *ActiveNodes_ActiveNode) GetSegmentPath() string {
    return "active-node" + "[node-name='" + fmt.Sprintf("%v", activeNode.NodeName) + "']"
}

func (activeNode *ActiveNodes_ActiveNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-freqsync-cfg:clock-interface" {
        return &activeNode.ClockInterface
    }
    if childYangName == "Cisco-IOS-XR-infra-ltrace-cfg:ltrace" {
        return &activeNode.Ltrace
    }
    if childYangName == "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local" {
        return &activeNode.LptsLocal
    }
    if childYangName == "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group" {
        return &activeNode.SsrpGroup
    }
    if childYangName == "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold" {
        return &activeNode.CiscoIosXrWatchdCfgWatchdogNodeThreshold
    }
    if childYangName == "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold" {
        return &activeNode.CiscoIosXrWdCfgWatchdogNodeThreshold
    }
    return nil
}

func (activeNode *ActiveNodes_ActiveNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-freqsync-cfg:clock-interface"] = &activeNode.ClockInterface
    children["Cisco-IOS-XR-infra-ltrace-cfg:ltrace"] = &activeNode.Ltrace
    children["Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"] = &activeNode.LptsLocal
    children["Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group"] = &activeNode.SsrpGroup
    children["Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"] = &activeNode.CiscoIosXrWatchdCfgWatchdogNodeThreshold
    children["Cisco-IOS-XR-wd-cfg:watchdog-node-threshold"] = &activeNode.CiscoIosXrWdCfgWatchdogNodeThreshold
    return children
}

func (activeNode *ActiveNodes_ActiveNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = activeNode.NodeName
    return leafs
}

func (activeNode *ActiveNodes_ActiveNode) GetBundleName() string { return "cisco_ios_xr" }

func (activeNode *ActiveNodes_ActiveNode) GetYangName() string { return "active-node" }

func (activeNode *ActiveNodes_ActiveNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeNode *ActiveNodes_ActiveNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeNode *ActiveNodes_ActiveNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeNode *ActiveNodes_ActiveNode) SetParent(parent types.Entity) { activeNode.parent = parent }

func (activeNode *ActiveNodes_ActiveNode) GetParent() types.Entity { return activeNode.parent }

func (activeNode *ActiveNodes_ActiveNode) GetParentYangName() string { return "active-nodes" }

// ActiveNodes_ActiveNode_ClockInterface
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface.
    Clocks ActiveNodes_ActiveNode_ClockInterface_Clocks
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetFilter() yfilter.YFilter { return clockInterface.YFilter }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) SetFilter(yf yfilter.YFilter) { clockInterface.YFilter = yf }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetGoName(yname string) string {
    if yname == "clocks" { return "Clocks" }
    return ""
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-freqsync-cfg:clock-interface"
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clocks" {
        return &clockInterface.Clocks
    }
    return nil
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clocks"] = &clockInterface.Clocks
    return children
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetYangName() string { return "clock-interface" }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) SetParent(parent types.Entity) { clockInterface.parent = parent }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetParent() types.Entity { return clockInterface.parent }

func (clockInterface *ActiveNodes_ActiveNode_ClockInterface) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_ClockInterface_Clocks
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface_Clocks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface. The type is slice of
    // ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock.
    Clock []ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetFilter() yfilter.YFilter { return clocks.YFilter }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) SetFilter(yf yfilter.YFilter) { clocks.YFilter = yf }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    return ""
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetSegmentPath() string {
    return "clocks"
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        for _, c := range clocks.Clock {
            if clocks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock{}
        clocks.Clock = append(clocks.Clock, child)
        return &clocks.Clock[len(clocks.Clock)-1]
    }
    return nil
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clocks.Clock {
        children[clocks.Clock[i].GetSegmentPath()] = &clocks.Clock[i]
    }
    return children
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetBundleName() string { return "cisco_ios_xr" }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetYangName() string { return "clocks" }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) SetParent(parent types.Entity) { clocks.parent = parent }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetParent() types.Entity { return clocks.parent }

func (clocks *ActiveNodes_ActiveNode_ClockInterface_Clocks) GetParentYangName() string { return "clock-interface" }

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock
// Configuration for a clock interface
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Port interface{}

    // Frequency Synchronization clock configuraiton.
    FrequencySynchronization ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetGoName(yname string) string {
    if yname == "clock-type" { return "ClockType" }
    if yname == "port" { return "Port" }
    if yname == "frequency-synchronization" { return "FrequencySynchronization" }
    return ""
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetSegmentPath() string {
    return "clock" + "[clock-type='" + fmt.Sprintf("%v", clock.ClockType) + "']" + "[port='" + fmt.Sprintf("%v", clock.Port) + "']"
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frequency-synchronization" {
        return &clock.FrequencySynchronization
    }
    return nil
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frequency-synchronization"] = &clock.FrequencySynchronization
    return children
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-type"] = clock.ClockType
    leafs["port"] = clock.Port
    return leafs
}

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetYangName() string { return "clock" }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetParent() types.Entity { return clock.parent }

func (clock *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock) GetParentYangName() string { return "clocks" }

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization
// Frequency Synchronization clock configuraiton
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization struct {
    parent types.Entity
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

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetFilter() yfilter.YFilter { return frequencySynchronization.YFilter }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) SetFilter(yf yfilter.YFilter) { frequencySynchronization.YFilter = yf }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetGoName(yname string) string {
    if yname == "wait-to-restore-time" { return "WaitToRestoreTime" }
    if yname == "priority" { return "Priority" }
    if yname == "selection-input" { return "SelectionInput" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "ssm-disable" { return "SsmDisable" }
    if yname == "output-quality-level" { return "OutputQualityLevel" }
    if yname == "input-quality-level" { return "InputQualityLevel" }
    return ""
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetSegmentPath() string {
    return "frequency-synchronization"
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-quality-level" {
        return &frequencySynchronization.OutputQualityLevel
    }
    if childYangName == "input-quality-level" {
        return &frequencySynchronization.InputQualityLevel
    }
    return nil
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output-quality-level"] = &frequencySynchronization.OutputQualityLevel
    children["input-quality-level"] = &frequencySynchronization.InputQualityLevel
    return children
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wait-to-restore-time"] = frequencySynchronization.WaitToRestoreTime
    leafs["priority"] = frequencySynchronization.Priority
    leafs["selection-input"] = frequencySynchronization.SelectionInput
    leafs["time-of-day-priority"] = frequencySynchronization.TimeOfDayPriority
    leafs["ssm-disable"] = frequencySynchronization.SsmDisable
    return leafs
}

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) SetParent(parent types.Entity) { frequencySynchronization.parent = parent }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetParent() types.Entity { return frequencySynchronization.parent }

func (frequencySynchronization *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetParentYangName() string { return "clock" }

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
// Set the output quality level
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel struct {
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

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetFilter() yfilter.YFilter { return outputQualityLevel.YFilter }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetFilter(yf yfilter.YFilter) { outputQualityLevel.YFilter = yf }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetSegmentPath() string {
    return "output-quality-level"
}

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = outputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = outputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = outputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetYangName() string { return "output-quality-level" }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetParent(parent types.Entity) { outputQualityLevel.parent = parent }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParent() types.Entity { return outputQualityLevel.parent }

func (outputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
// Set the input quality level
type ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel struct {
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

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetFilter() yfilter.YFilter { return inputQualityLevel.YFilter }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetFilter(yf yfilter.YFilter) { inputQualityLevel.YFilter = yf }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetSegmentPath() string {
    return "input-quality-level"
}

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = inputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = inputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = inputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetYangName() string { return "input-quality-level" }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetParent(parent types.Entity) { inputQualityLevel.parent = parent }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParent() types.Entity { return inputQualityLevel.parent }

func (inputQualityLevel *ActiveNodes_ActiveNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// ActiveNodes_ActiveNode_Ltrace
// Ltrace Memory configuration
type ActiveNodes_ActiveNode_Ltrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select Ltrace mode and scale-factor.
    AllocationParams ActiveNodes_ActiveNode_Ltrace_AllocationParams
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetFilter() yfilter.YFilter { return ltrace.YFilter }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) SetFilter(yf yfilter.YFilter) { ltrace.YFilter = yf }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetGoName(yname string) string {
    if yname == "allocation-params" { return "AllocationParams" }
    return ""
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-ltrace-cfg:ltrace"
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "allocation-params" {
        return &ltrace.AllocationParams
    }
    return nil
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["allocation-params"] = &ltrace.AllocationParams
    return children
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetBundleName() string { return "cisco_ios_xr" }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetYangName() string { return "ltrace" }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) SetParent(parent types.Entity) { ltrace.parent = parent }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetParent() types.Entity { return ltrace.parent }

func (ltrace *ActiveNodes_ActiveNode_Ltrace) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_Ltrace_AllocationParams
// Select Ltrace mode and scale-factor
type ActiveNodes_ActiveNode_Ltrace_AllocationParams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select an allocation mode (static:1, dynamic :2). The type is
    // InfraLtraceMode.
    Mode interface{}

    // Select a scaling down factor. The type is InfraLtraceScale.
    ScaleFactor interface{}
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetFilter() yfilter.YFilter { return allocationParams.YFilter }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) SetFilter(yf yfilter.YFilter) { allocationParams.YFilter = yf }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "scale-factor" { return "ScaleFactor" }
    return ""
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetSegmentPath() string {
    return "allocation-params"
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = allocationParams.Mode
    leafs["scale-factor"] = allocationParams.ScaleFactor
    return leafs
}

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetBundleName() string { return "cisco_ios_xr" }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetYangName() string { return "allocation-params" }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) SetParent(parent types.Entity) { allocationParams.parent = parent }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetParent() types.Entity { return allocationParams.parent }

func (allocationParams *ActiveNodes_ActiveNode_Ltrace_AllocationParams) GetParentYangName() string { return "ltrace" }

// ActiveNodes_ActiveNode_LptsLocal
// lpts node specific configuration commands
type ActiveNodes_ActiveNode_LptsLocal struct {
    parent types.Entity
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

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetFilter() yfilter.YFilter { return lptsLocal.YFilter }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) SetFilter(yf yfilter.YFilter) { lptsLocal.YFilter = yf }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetGoName(yname string) string {
    if yname == "ipolicer-local-tables" { return "IpolicerLocalTables" }
    if yname == "dynamic-flows-tables" { return "DynamicFlowsTables" }
    if yname == "ipolicer-local" { return "IpolicerLocal" }
    return ""
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipolicer-local-tables" {
        return &lptsLocal.IpolicerLocalTables
    }
    if childYangName == "dynamic-flows-tables" {
        return &lptsLocal.DynamicFlowsTables
    }
    if childYangName == "ipolicer-local" {
        return &lptsLocal.IpolicerLocal
    }
    return nil
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipolicer-local-tables"] = &lptsLocal.IpolicerLocalTables
    children["dynamic-flows-tables"] = &lptsLocal.DynamicFlowsTables
    children["ipolicer-local"] = &lptsLocal.IpolicerLocal
    return children
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetBundleName() string { return "cisco_ios_xr" }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetYangName() string { return "lpts-local" }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) SetParent(parent types.Entity) { lptsLocal.parent = parent }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetParent() types.Entity { return lptsLocal.parent }

func (lptsLocal *ActiveNodes_ActiveNode_LptsLocal) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pre IFIB (Internal Forwarding Information Base) configuration table. The
    // type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable.
    IpolicerLocalTable []ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetFilter() yfilter.YFilter { return ipolicerLocalTables.YFilter }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) SetFilter(yf yfilter.YFilter) { ipolicerLocalTables.YFilter = yf }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetGoName(yname string) string {
    if yname == "ipolicer-local-table" { return "IpolicerLocalTable" }
    return ""
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetSegmentPath() string {
    return "ipolicer-local-tables"
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipolicer-local-table" {
        for _, c := range ipolicerLocalTables.IpolicerLocalTable {
            if ipolicerLocalTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable{}
        ipolicerLocalTables.IpolicerLocalTable = append(ipolicerLocalTables.IpolicerLocalTable, child)
        return &ipolicerLocalTables.IpolicerLocalTable[len(ipolicerLocalTables.IpolicerLocalTable)-1]
    }
    return nil
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipolicerLocalTables.IpolicerLocalTable {
        children[ipolicerLocalTables.IpolicerLocalTable[i].GetSegmentPath()] = &ipolicerLocalTables.IpolicerLocalTable[i]
    }
    return children
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetYangName() string { return "ipolicer-local-tables" }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) SetParent(parent types.Entity) { ipolicerLocalTables.parent = parent }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetParent() types.Entity { return ipolicerLocalTables.parent }

func (ipolicerLocalTables *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables) GetParentYangName() string { return "lpts-local" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
// Pre IFIB (Internal Forwarding Information
// Base) configuration table
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // NP name.
    Nps ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetFilter() yfilter.YFilter { return ipolicerLocalTable.YFilter }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) SetFilter(yf yfilter.YFilter) { ipolicerLocalTable.YFilter = yf }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetGoName(yname string) string {
    if yname == "id1" { return "Id1" }
    if yname == "nps" { return "Nps" }
    return ""
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetSegmentPath() string {
    return "ipolicer-local-table" + "[id1='" + fmt.Sprintf("%v", ipolicerLocalTable.Id1) + "']"
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nps" {
        return &ipolicerLocalTable.Nps
    }
    return nil
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nps"] = &ipolicerLocalTable.Nps
    return children
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id1"] = ipolicerLocalTable.Id1
    return leafs
}

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetYangName() string { return "ipolicer-local-table" }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) SetParent(parent types.Entity) { ipolicerLocalTable.parent = parent }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetParent() types.Entity { return ipolicerLocalTable.parent }

func (ipolicerLocalTable *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetParentYangName() string { return "ipolicer-local-tables" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps
// NP name
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of NP names. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np.
    Np []ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetFilter() yfilter.YFilter { return nps.YFilter }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) SetFilter(yf yfilter.YFilter) { nps.YFilter = yf }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetGoName(yname string) string {
    if yname == "np" { return "Np" }
    return ""
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetSegmentPath() string {
    return "nps"
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        for _, c := range nps.Np {
            if nps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np{}
        nps.Np = append(nps.Np, child)
        return &nps.Np[len(nps.Np)-1]
    }
    return nil
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nps.Np {
        children[nps.Np[i].GetSegmentPath()] = &nps.Np[i]
    }
    return children
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetBundleName() string { return "cisco_ios_xr" }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetYangName() string { return "nps" }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) SetParent(parent types.Entity) { nps.parent = parent }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetParent() types.Entity { return nps.parent }

func (nps *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetParentYangName() string { return "ipolicer-local-table" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np
// Table of NP names
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is interface{} with range:
    // -2147483648..2147483647.
    Id1 interface{}

    // Packets per second. The type is interface{} with range:
    // -2147483648..2147483647. Units are packet/s.
    Rate interface{}
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetGoName(yname string) string {
    if yname == "id1" { return "Id1" }
    if yname == "rate" { return "Rate" }
    return ""
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetSegmentPath() string {
    return "np" + "[id1='" + fmt.Sprintf("%v", np.Id1) + "']"
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id1"] = np.Id1
    leafs["rate"] = np.Rate
    return leafs
}

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetYangName() string { return "np" }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetParent() types.Entity { return np.parent }

func (np *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetParentYangName() string { return "nps" }

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table for Dynamic Flows. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable.
    DynamicFlowsTable []ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetFilter() yfilter.YFilter { return dynamicFlowsTables.YFilter }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) SetFilter(yf yfilter.YFilter) { dynamicFlowsTables.YFilter = yf }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetGoName(yname string) string {
    if yname == "dynamic-flows-table" { return "DynamicFlowsTable" }
    return ""
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetSegmentPath() string {
    return "dynamic-flows-tables"
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dynamic-flows-table" {
        for _, c := range dynamicFlowsTables.DynamicFlowsTable {
            if dynamicFlowsTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable{}
        dynamicFlowsTables.DynamicFlowsTable = append(dynamicFlowsTables.DynamicFlowsTable, child)
        return &dynamicFlowsTables.DynamicFlowsTable[len(dynamicFlowsTables.DynamicFlowsTable)-1]
    }
    return nil
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicFlowsTables.DynamicFlowsTable {
        children[dynamicFlowsTables.DynamicFlowsTable[i].GetSegmentPath()] = &dynamicFlowsTables.DynamicFlowsTable[i]
    }
    return children
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetYangName() string { return "dynamic-flows-tables" }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) SetParent(parent types.Entity) { dynamicFlowsTables.parent = parent }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetParent() types.Entity { return dynamicFlowsTables.parent }

func (dynamicFlowsTables *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables) GetParentYangName() string { return "lpts-local" }

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
// Table for Dynamic Flows
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Dynamic Flows Table Type. The type is
    // LptsDynamicFlowConfig.
    TableType interface{}

    // Selected flow type. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType.
    FlowType []ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetFilter() yfilter.YFilter { return dynamicFlowsTable.YFilter }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) SetFilter(yf yfilter.YFilter) { dynamicFlowsTable.YFilter = yf }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetGoName(yname string) string {
    if yname == "table-type" { return "TableType" }
    if yname == "flow-type" { return "FlowType" }
    return ""
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetSegmentPath() string {
    return "dynamic-flows-table" + "[table-type='" + fmt.Sprintf("%v", dynamicFlowsTable.TableType) + "']"
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-type" {
        for _, c := range dynamicFlowsTable.FlowType {
            if dynamicFlowsTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType{}
        dynamicFlowsTable.FlowType = append(dynamicFlowsTable.FlowType, child)
        return &dynamicFlowsTable.FlowType[len(dynamicFlowsTable.FlowType)-1]
    }
    return nil
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicFlowsTable.FlowType {
        children[dynamicFlowsTable.FlowType[i].GetSegmentPath()] = &dynamicFlowsTable.FlowType[i]
    }
    return children
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["table-type"] = dynamicFlowsTable.TableType
    return leafs
}

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetYangName() string { return "dynamic-flows-table" }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) SetParent(parent types.Entity) { dynamicFlowsTable.parent = parent }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetParent() types.Entity { return dynamicFlowsTable.parent }

func (dynamicFlowsTable *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetParentYangName() string { return "dynamic-flows-tables" }

// ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
// Selected flow type
type ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured Max TCAM value. The type is interface{} with range:
    // -2147483648..2147483647.
    Max interface{}
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetFilter() yfilter.YFilter { return flowType.YFilter }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) SetFilter(yf yfilter.YFilter) { flowType.YFilter = yf }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetGoName(yname string) string {
    if yname == "flow-type" { return "FlowType" }
    if yname == "max" { return "Max" }
    return ""
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetSegmentPath() string {
    return "flow-type" + "[flow-type='" + fmt.Sprintf("%v", flowType.FlowType) + "']"
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-type"] = flowType.FlowType
    leafs["max"] = flowType.Max
    return leafs
}

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetBundleName() string { return "cisco_ios_xr" }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetYangName() string { return "flow-type" }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) SetParent(parent types.Entity) { flowType.parent = parent }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetParent() types.Entity { return flowType.parent }

func (flowType *ActiveNodes_ActiveNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetParentYangName() string { return "dynamic-flows-table" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
// This type is a presence type.
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for Flows.
    Flows ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetFilter() yfilter.YFilter { return ipolicerLocal.YFilter }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) SetFilter(yf yfilter.YFilter) { ipolicerLocal.YFilter = yf }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "flows" { return "Flows" }
    return ""
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetSegmentPath() string {
    return "ipolicer-local"
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flows" {
        return &ipolicerLocal.Flows
    }
    return nil
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flows"] = &ipolicerLocal.Flows
    return children
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ipolicerLocal.Enable
    return leafs
}

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetYangName() string { return "ipolicer-local" }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) SetParent(parent types.Entity) { ipolicerLocal.parent = parent }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetParent() types.Entity { return ipolicerLocal.parent }

func (ipolicerLocal *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal) GetParentYangName() string { return "lpts-local" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows
// Table for Flows
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of
    // ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow.
    Flow []ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetFilter() yfilter.YFilter { return flows.YFilter }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) SetFilter(yf yfilter.YFilter) { flows.YFilter = yf }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetGoName(yname string) string {
    if yname == "flow" { return "Flow" }
    return ""
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetSegmentPath() string {
    return "flows"
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow" {
        for _, c := range flows.Flow {
            if flows.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow{}
        flows.Flow = append(flows.Flow, child)
        return &flows.Flow[len(flows.Flow)-1]
    }
    return nil
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flows.Flow {
        children[flows.Flow[i].GetSegmentPath()] = &flows.Flow[i]
    }
    return children
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetBundleName() string { return "cisco_ios_xr" }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetYangName() string { return "flows" }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) SetParent(parent types.Entity) { flows.parent = parent }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetParent() types.Entity { return flows.parent }

func (flows *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows) GetParentYangName() string { return "ipolicer-local" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow
// selected flow type
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range:
    // -2147483648..2147483647.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetGoName(yname string) string {
    if yname == "flow-type" { return "FlowType" }
    if yname == "rate" { return "Rate" }
    if yname == "precedences" { return "Precedences" }
    return ""
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetSegmentPath() string {
    return "flow" + "[flow-type='" + fmt.Sprintf("%v", flow.FlowType) + "']"
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "precedences" {
        return &flow.Precedences
    }
    return nil
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["precedences"] = &flow.Precedences
    return children
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-type"] = flow.FlowType
    leafs["rate"] = flow.Rate
    return leafs
}

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetYangName() string { return "flow" }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetParent() types.Entity { return flow.parent }

func (flow *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow) GetParentYangName() string { return "flows" }

// ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
// TOS Precedence value(s)
type ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetFilter() yfilter.YFilter { return precedences.YFilter }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) SetFilter(yf yfilter.YFilter) { precedences.YFilter = yf }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetGoName(yname string) string {
    if yname == "precedence" { return "Precedence" }
    return ""
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetSegmentPath() string {
    return "precedences"
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["precedence"] = precedences.Precedence
    return leafs
}

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetBundleName() string { return "cisco_ios_xr" }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetYangName() string { return "precedences" }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) SetParent(parent types.Entity) { precedences.parent = parent }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetParent() types.Entity { return precedences.parent }

func (precedences *ActiveNodes_ActiveNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetParentYangName() string { return "flow" }

// ActiveNodes_ActiveNode_SsrpGroup
// Per-node SSRP configuration data
type ActiveNodes_ActiveNode_SsrpGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of SSRP Group configuration.
    Groups ActiveNodes_ActiveNode_SsrpGroup_Groups
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetFilter() yfilter.YFilter { return ssrpGroup.YFilter }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) SetFilter(yf yfilter.YFilter) { ssrpGroup.YFilter = yf }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    return ""
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp-group"
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &ssrpGroup.Groups
    }
    return nil
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &ssrpGroup.Groups
    return children
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetYangName() string { return "ssrp-group" }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) SetParent(parent types.Entity) { ssrpGroup.parent = parent }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetParent() types.Entity { return ssrpGroup.parent }

func (ssrpGroup *ActiveNodes_ActiveNode_SsrpGroup) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_SsrpGroup_Groups
// Table of SSRP Group configuration
type ActiveNodes_ActiveNode_SsrpGroup_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSRP Group configuration. The type is slice of
    // ActiveNodes_ActiveNode_SsrpGroup_Groups_Group.
    Group []ActiveNodes_ActiveNode_SsrpGroup_Groups_Group
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ActiveNodes_ActiveNode_SsrpGroup_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetYangName() string { return "groups" }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetParent() types.Entity { return groups.parent }

func (groups *ActiveNodes_ActiveNode_SsrpGroup_Groups) GetParentYangName() string { return "ssrp-group" }

// ActiveNodes_ActiveNode_SsrpGroup_Groups_Group
// SSRP Group configuration
type ActiveNodes_ActiveNode_SsrpGroup_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this group. The type is
    // interface{} with range: 1..65535.
    GroupId interface{}

    // This specifies the SSRP profile to use for this group. The type is string.
    Profile interface{}
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "profile" { return "Profile" }
    return ""
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']"
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = group.GroupId
    leafs["profile"] = group.Profile
    return leafs
}

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetYangName() string { return "group" }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *ActiveNodes_ActiveNode_SsrpGroup_Groups_Group) GetParentYangName() string { return "groups" }

// ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold
// watchdog node threshold
type ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory thresholds.
    MemoryThreshold ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetFilter() yfilter.YFilter { return ciscoIOSXRWatchdCfgWatchdogNodeThreshold.YFilter }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) SetFilter(yf yfilter.YFilter) { ciscoIOSXRWatchdCfgWatchdogNodeThreshold.YFilter = yf }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetGoName(yname string) string {
    if yname == "memory-threshold" { return "MemoryThreshold" }
    return ""
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetSegmentPath() string {
    return "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-threshold" {
        return &ciscoIOSXRWatchdCfgWatchdogNodeThreshold.MemoryThreshold
    }
    return nil
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["memory-threshold"] = &ciscoIOSXRWatchdCfgWatchdogNodeThreshold.MemoryThreshold
    return children
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetYangName() string { return "watchdog-node-threshold" }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) SetParent(parent types.Entity) { ciscoIOSXRWatchdCfgWatchdogNodeThreshold.parent = parent }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetParent() types.Entity { return ciscoIOSXRWatchdCfgWatchdogNodeThreshold.parent }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetFilter() yfilter.YFilter { return memoryThreshold.YFilter }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) SetFilter(yf yfilter.YFilter) { memoryThreshold.YFilter = yf }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetSegmentPath() string {
    return "memory-threshold"
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = memoryThreshold.Minor
    leafs["severe"] = memoryThreshold.Severe
    leafs["critical"] = memoryThreshold.Critical
    return leafs
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetYangName() string { return "memory-threshold" }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) SetParent(parent types.Entity) { memoryThreshold.parent = parent }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetParent() types.Entity { return memoryThreshold.parent }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetParentYangName() string { return "watchdog-node-threshold" }

// ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold
// Watchdog threshold configuration
type ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory thresholds.
    MemoryThreshold ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetFilter() yfilter.YFilter { return ciscoIOSXRWdCfgWatchdogNodeThreshold.YFilter }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) SetFilter(yf yfilter.YFilter) { ciscoIOSXRWdCfgWatchdogNodeThreshold.YFilter = yf }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetGoName(yname string) string {
    if yname == "memory-threshold" { return "MemoryThreshold" }
    return ""
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetSegmentPath() string {
    return "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold"
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-threshold" {
        return &ciscoIOSXRWdCfgWatchdogNodeThreshold.MemoryThreshold
    }
    return nil
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["memory-threshold"] = &ciscoIOSXRWdCfgWatchdogNodeThreshold.MemoryThreshold
    return children
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetYangName() string { return "watchdog-node-threshold" }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) SetParent(parent types.Entity) { ciscoIOSXRWdCfgWatchdogNodeThreshold.parent = parent }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetParent() types.Entity { return ciscoIOSXRWdCfgWatchdogNodeThreshold.parent }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetParentYangName() string { return "active-node" }

// ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetFilter() yfilter.YFilter { return memoryThreshold.YFilter }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) SetFilter(yf yfilter.YFilter) { memoryThreshold.YFilter = yf }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetSegmentPath() string {
    return "memory-threshold"
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = memoryThreshold.Minor
    leafs["severe"] = memoryThreshold.Severe
    leafs["critical"] = memoryThreshold.Critical
    return leafs
}

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetYangName() string { return "memory-threshold" }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) SetParent(parent types.Entity) { memoryThreshold.parent = parent }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetParent() types.Entity { return memoryThreshold.parent }

func (memoryThreshold *ActiveNodes_ActiveNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetParentYangName() string { return "watchdog-node-threshold" }

// PreconfiguredNodes
// preconfigured nodes
type PreconfiguredNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The configuration for a non-active node. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode.
    PreconfiguredNode []PreconfiguredNodes_PreconfiguredNode
}

func (preconfiguredNodes *PreconfiguredNodes) GetFilter() yfilter.YFilter { return preconfiguredNodes.YFilter }

func (preconfiguredNodes *PreconfiguredNodes) SetFilter(yf yfilter.YFilter) { preconfiguredNodes.YFilter = yf }

func (preconfiguredNodes *PreconfiguredNodes) GetGoName(yname string) string {
    if yname == "preconfigured-node" { return "PreconfiguredNode" }
    return ""
}

func (preconfiguredNodes *PreconfiguredNodes) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-mda-cfg:preconfigured-nodes"
}

func (preconfiguredNodes *PreconfiguredNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "preconfigured-node" {
        for _, c := range preconfiguredNodes.PreconfiguredNode {
            if preconfiguredNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode{}
        preconfiguredNodes.PreconfiguredNode = append(preconfiguredNodes.PreconfiguredNode, child)
        return &preconfiguredNodes.PreconfiguredNode[len(preconfiguredNodes.PreconfiguredNode)-1]
    }
    return nil
}

func (preconfiguredNodes *PreconfiguredNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range preconfiguredNodes.PreconfiguredNode {
        children[preconfiguredNodes.PreconfiguredNode[i].GetSegmentPath()] = &preconfiguredNodes.PreconfiguredNode[i]
    }
    return children
}

func (preconfiguredNodes *PreconfiguredNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (preconfiguredNodes *PreconfiguredNodes) GetBundleName() string { return "cisco_ios_xr" }

func (preconfiguredNodes *PreconfiguredNodes) GetYangName() string { return "preconfigured-nodes" }

func (preconfiguredNodes *PreconfiguredNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preconfiguredNodes *PreconfiguredNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preconfiguredNodes *PreconfiguredNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preconfiguredNodes *PreconfiguredNodes) SetParent(parent types.Entity) { preconfiguredNodes.parent = parent }

func (preconfiguredNodes *PreconfiguredNodes) GetParent() types.Entity { return preconfiguredNodes.parent }

func (preconfiguredNodes *PreconfiguredNodes) GetParentYangName() string { return "Cisco-IOS-XR-config-mda-cfg" }

// PreconfiguredNodes_PreconfiguredNode
// The configuration for a non-active node
type PreconfiguredNodes_PreconfiguredNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for this node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Configuration for a clock interface.
    ClockInterface PreconfiguredNodes_PreconfiguredNode_ClockInterface

    // Ltrace Memory configuration.
    Ltrace PreconfiguredNodes_PreconfiguredNode_Ltrace

    // lpts node specific configuration commands.
    LptsLocal PreconfiguredNodes_PreconfiguredNode_LptsLocal

    // watchdog node threshold.
    CiscoIosXrWatchdCfgWatchdogNodeThreshold PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold

    // Watchdog threshold configuration.
    CiscoIosXrWdCfgWatchdogNodeThreshold PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetFilter() yfilter.YFilter { return preconfiguredNode.YFilter }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) SetFilter(yf yfilter.YFilter) { preconfiguredNode.YFilter = yf }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "Cisco-IOS-XR-freqsync-cfg:clock-interface" { return "ClockInterface" }
    if yname == "Cisco-IOS-XR-infra-ltrace-cfg:ltrace" { return "Ltrace" }
    if yname == "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local" { return "LptsLocal" }
    if yname == "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold" { return "CiscoIosXrWatchdCfgWatchdogNodeThreshold" }
    if yname == "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold" { return "CiscoIosXrWdCfgWatchdogNodeThreshold" }
    return ""
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetSegmentPath() string {
    return "preconfigured-node" + "[node-name='" + fmt.Sprintf("%v", preconfiguredNode.NodeName) + "']"
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-freqsync-cfg:clock-interface" {
        return &preconfiguredNode.ClockInterface
    }
    if childYangName == "Cisco-IOS-XR-infra-ltrace-cfg:ltrace" {
        return &preconfiguredNode.Ltrace
    }
    if childYangName == "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local" {
        return &preconfiguredNode.LptsLocal
    }
    if childYangName == "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold" {
        return &preconfiguredNode.CiscoIosXrWatchdCfgWatchdogNodeThreshold
    }
    if childYangName == "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold" {
        return &preconfiguredNode.CiscoIosXrWdCfgWatchdogNodeThreshold
    }
    return nil
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-freqsync-cfg:clock-interface"] = &preconfiguredNode.ClockInterface
    children["Cisco-IOS-XR-infra-ltrace-cfg:ltrace"] = &preconfiguredNode.Ltrace
    children["Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"] = &preconfiguredNode.LptsLocal
    children["Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"] = &preconfiguredNode.CiscoIosXrWatchdCfgWatchdogNodeThreshold
    children["Cisco-IOS-XR-wd-cfg:watchdog-node-threshold"] = &preconfiguredNode.CiscoIosXrWdCfgWatchdogNodeThreshold
    return children
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = preconfiguredNode.NodeName
    return leafs
}

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetBundleName() string { return "cisco_ios_xr" }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetYangName() string { return "preconfigured-node" }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) SetParent(parent types.Entity) { preconfiguredNode.parent = parent }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetParent() types.Entity { return preconfiguredNode.parent }

func (preconfiguredNode *PreconfiguredNodes_PreconfiguredNode) GetParentYangName() string { return "preconfigured-nodes" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface.
    Clocks PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetFilter() yfilter.YFilter { return clockInterface.YFilter }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) SetFilter(yf yfilter.YFilter) { clockInterface.YFilter = yf }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetGoName(yname string) string {
    if yname == "clocks" { return "Clocks" }
    return ""
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetSegmentPath() string {
    return "Cisco-IOS-XR-freqsync-cfg:clock-interface"
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clocks" {
        return &clockInterface.Clocks
    }
    return nil
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clocks"] = &clockInterface.Clocks
    return children
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetYangName() string { return "clock-interface" }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) SetParent(parent types.Entity) { clockInterface.parent = parent }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetParent() types.Entity { return clockInterface.parent }

func (clockInterface *PreconfiguredNodes_PreconfiguredNode_ClockInterface) GetParentYangName() string { return "preconfigured-node" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a clock interface. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock.
    Clock []PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetFilter() yfilter.YFilter { return clocks.YFilter }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) SetFilter(yf yfilter.YFilter) { clocks.YFilter = yf }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    return ""
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetSegmentPath() string {
    return "clocks"
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        for _, c := range clocks.Clock {
            if clocks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock{}
        clocks.Clock = append(clocks.Clock, child)
        return &clocks.Clock[len(clocks.Clock)-1]
    }
    return nil
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clocks.Clock {
        children[clocks.Clock[i].GetSegmentPath()] = &clocks.Clock[i]
    }
    return children
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetBundleName() string { return "cisco_ios_xr" }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetYangName() string { return "clocks" }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) SetParent(parent types.Entity) { clocks.parent = parent }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetParent() types.Entity { return clocks.parent }

func (clocks *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks) GetParentYangName() string { return "clock-interface" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock
// Configuration for a clock interface
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Port interface{}

    // Frequency Synchronization clock configuraiton.
    FrequencySynchronization PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetGoName(yname string) string {
    if yname == "clock-type" { return "ClockType" }
    if yname == "port" { return "Port" }
    if yname == "frequency-synchronization" { return "FrequencySynchronization" }
    return ""
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetSegmentPath() string {
    return "clock" + "[clock-type='" + fmt.Sprintf("%v", clock.ClockType) + "']" + "[port='" + fmt.Sprintf("%v", clock.Port) + "']"
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frequency-synchronization" {
        return &clock.FrequencySynchronization
    }
    return nil
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frequency-synchronization"] = &clock.FrequencySynchronization
    return children
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-type"] = clock.ClockType
    leafs["port"] = clock.Port
    return leafs
}

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetYangName() string { return "clock" }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetParent() types.Entity { return clock.parent }

func (clock *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock) GetParentYangName() string { return "clocks" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization
// Frequency Synchronization clock configuraiton
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization struct {
    parent types.Entity
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

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetFilter() yfilter.YFilter { return frequencySynchronization.YFilter }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) SetFilter(yf yfilter.YFilter) { frequencySynchronization.YFilter = yf }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetGoName(yname string) string {
    if yname == "wait-to-restore-time" { return "WaitToRestoreTime" }
    if yname == "priority" { return "Priority" }
    if yname == "selection-input" { return "SelectionInput" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "ssm-disable" { return "SsmDisable" }
    if yname == "output-quality-level" { return "OutputQualityLevel" }
    if yname == "input-quality-level" { return "InputQualityLevel" }
    return ""
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetSegmentPath() string {
    return "frequency-synchronization"
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-quality-level" {
        return &frequencySynchronization.OutputQualityLevel
    }
    if childYangName == "input-quality-level" {
        return &frequencySynchronization.InputQualityLevel
    }
    return nil
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output-quality-level"] = &frequencySynchronization.OutputQualityLevel
    children["input-quality-level"] = &frequencySynchronization.InputQualityLevel
    return children
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wait-to-restore-time"] = frequencySynchronization.WaitToRestoreTime
    leafs["priority"] = frequencySynchronization.Priority
    leafs["selection-input"] = frequencySynchronization.SelectionInput
    leafs["time-of-day-priority"] = frequencySynchronization.TimeOfDayPriority
    leafs["ssm-disable"] = frequencySynchronization.SsmDisable
    return leafs
}

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) SetParent(parent types.Entity) { frequencySynchronization.parent = parent }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetParent() types.Entity { return frequencySynchronization.parent }

func (frequencySynchronization *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization) GetParentYangName() string { return "clock" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel
// Set the output quality level
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel struct {
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

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetFilter() yfilter.YFilter { return outputQualityLevel.YFilter }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetFilter(yf yfilter.YFilter) { outputQualityLevel.YFilter = yf }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetSegmentPath() string {
    return "output-quality-level"
}

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = outputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = outputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = outputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetYangName() string { return "output-quality-level" }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) SetParent(parent types.Entity) { outputQualityLevel.parent = parent }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParent() types.Entity { return outputQualityLevel.parent }

func (outputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_OutputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel
// Set the input quality level
type PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel struct {
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

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetFilter() yfilter.YFilter { return inputQualityLevel.YFilter }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetFilter(yf yfilter.YFilter) { inputQualityLevel.YFilter = yf }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "exact-quality-level-value" { return "ExactQualityLevelValue" }
    if yname == "min-quality-level-value" { return "MinQualityLevelValue" }
    if yname == "max-quality-level-value" { return "MaxQualityLevelValue" }
    return ""
}

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetSegmentPath() string {
    return "input-quality-level"
}

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputQualityLevel.QualityLevelOption
    leafs["exact-quality-level-value"] = inputQualityLevel.ExactQualityLevelValue
    leafs["min-quality-level-value"] = inputQualityLevel.MinQualityLevelValue
    leafs["max-quality-level-value"] = inputQualityLevel.MaxQualityLevelValue
    return leafs
}

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetYangName() string { return "input-quality-level" }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) SetParent(parent types.Entity) { inputQualityLevel.parent = parent }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParent() types.Entity { return inputQualityLevel.parent }

func (inputQualityLevel *PreconfiguredNodes_PreconfiguredNode_ClockInterface_Clocks_Clock_FrequencySynchronization_InputQualityLevel) GetParentYangName() string { return "frequency-synchronization" }

// PreconfiguredNodes_PreconfiguredNode_Ltrace
// Ltrace Memory configuration
type PreconfiguredNodes_PreconfiguredNode_Ltrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select Ltrace mode and scale-factor.
    AllocationParams PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetFilter() yfilter.YFilter { return ltrace.YFilter }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) SetFilter(yf yfilter.YFilter) { ltrace.YFilter = yf }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetGoName(yname string) string {
    if yname == "allocation-params" { return "AllocationParams" }
    return ""
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-ltrace-cfg:ltrace"
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "allocation-params" {
        return &ltrace.AllocationParams
    }
    return nil
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["allocation-params"] = &ltrace.AllocationParams
    return children
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetBundleName() string { return "cisco_ios_xr" }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetYangName() string { return "ltrace" }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) SetParent(parent types.Entity) { ltrace.parent = parent }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetParent() types.Entity { return ltrace.parent }

func (ltrace *PreconfiguredNodes_PreconfiguredNode_Ltrace) GetParentYangName() string { return "preconfigured-node" }

// PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams
// Select Ltrace mode and scale-factor
type PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select an allocation mode (static:1, dynamic :2). The type is
    // InfraLtraceMode.
    Mode interface{}

    // Select a scaling down factor. The type is InfraLtraceScale.
    ScaleFactor interface{}
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetFilter() yfilter.YFilter { return allocationParams.YFilter }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) SetFilter(yf yfilter.YFilter) { allocationParams.YFilter = yf }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "scale-factor" { return "ScaleFactor" }
    return ""
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetSegmentPath() string {
    return "allocation-params"
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = allocationParams.Mode
    leafs["scale-factor"] = allocationParams.ScaleFactor
    return leafs
}

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetBundleName() string { return "cisco_ios_xr" }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetYangName() string { return "allocation-params" }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) SetParent(parent types.Entity) { allocationParams.parent = parent }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetParent() types.Entity { return allocationParams.parent }

func (allocationParams *PreconfiguredNodes_PreconfiguredNode_Ltrace_AllocationParams) GetParentYangName() string { return "ltrace" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal
// lpts node specific configuration commands
type PreconfiguredNodes_PreconfiguredNode_LptsLocal struct {
    parent types.Entity
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

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetFilter() yfilter.YFilter { return lptsLocal.YFilter }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) SetFilter(yf yfilter.YFilter) { lptsLocal.YFilter = yf }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetGoName(yname string) string {
    if yname == "ipolicer-local-tables" { return "IpolicerLocalTables" }
    if yname == "dynamic-flows-tables" { return "DynamicFlowsTables" }
    if yname == "ipolicer-local" { return "IpolicerLocal" }
    return ""
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-pre-ifib-cfg:lpts-local"
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipolicer-local-tables" {
        return &lptsLocal.IpolicerLocalTables
    }
    if childYangName == "dynamic-flows-tables" {
        return &lptsLocal.DynamicFlowsTables
    }
    if childYangName == "ipolicer-local" {
        return &lptsLocal.IpolicerLocal
    }
    return nil
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipolicer-local-tables"] = &lptsLocal.IpolicerLocalTables
    children["dynamic-flows-tables"] = &lptsLocal.DynamicFlowsTables
    children["ipolicer-local"] = &lptsLocal.IpolicerLocal
    return children
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetBundleName() string { return "cisco_ios_xr" }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetYangName() string { return "lpts-local" }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) SetParent(parent types.Entity) { lptsLocal.parent = parent }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetParent() types.Entity { return lptsLocal.parent }

func (lptsLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal) GetParentYangName() string { return "preconfigured-node" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pre IFIB (Internal Forwarding Information Base) configuration table. The
    // type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable.
    IpolicerLocalTable []PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetFilter() yfilter.YFilter { return ipolicerLocalTables.YFilter }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) SetFilter(yf yfilter.YFilter) { ipolicerLocalTables.YFilter = yf }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetGoName(yname string) string {
    if yname == "ipolicer-local-table" { return "IpolicerLocalTable" }
    return ""
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetSegmentPath() string {
    return "ipolicer-local-tables"
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipolicer-local-table" {
        for _, c := range ipolicerLocalTables.IpolicerLocalTable {
            if ipolicerLocalTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable{}
        ipolicerLocalTables.IpolicerLocalTable = append(ipolicerLocalTables.IpolicerLocalTable, child)
        return &ipolicerLocalTables.IpolicerLocalTable[len(ipolicerLocalTables.IpolicerLocalTable)-1]
    }
    return nil
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipolicerLocalTables.IpolicerLocalTable {
        children[ipolicerLocalTables.IpolicerLocalTable[i].GetSegmentPath()] = &ipolicerLocalTables.IpolicerLocalTable[i]
    }
    return children
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetYangName() string { return "ipolicer-local-tables" }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) SetParent(parent types.Entity) { ipolicerLocalTables.parent = parent }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetParent() types.Entity { return ipolicerLocalTables.parent }

func (ipolicerLocalTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables) GetParentYangName() string { return "lpts-local" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable
// Pre IFIB (Internal Forwarding Information
// Base) configuration table
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Id1 interface{}

    // NP name.
    Nps PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetFilter() yfilter.YFilter { return ipolicerLocalTable.YFilter }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) SetFilter(yf yfilter.YFilter) { ipolicerLocalTable.YFilter = yf }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetGoName(yname string) string {
    if yname == "id1" { return "Id1" }
    if yname == "nps" { return "Nps" }
    return ""
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetSegmentPath() string {
    return "ipolicer-local-table" + "[id1='" + fmt.Sprintf("%v", ipolicerLocalTable.Id1) + "']"
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nps" {
        return &ipolicerLocalTable.Nps
    }
    return nil
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nps"] = &ipolicerLocalTable.Nps
    return children
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id1"] = ipolicerLocalTable.Id1
    return leafs
}

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetYangName() string { return "ipolicer-local-table" }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) SetParent(parent types.Entity) { ipolicerLocalTable.parent = parent }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetParent() types.Entity { return ipolicerLocalTable.parent }

func (ipolicerLocalTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable) GetParentYangName() string { return "ipolicer-local-tables" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps
// NP name
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of NP names. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np.
    Np []PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetFilter() yfilter.YFilter { return nps.YFilter }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) SetFilter(yf yfilter.YFilter) { nps.YFilter = yf }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetGoName(yname string) string {
    if yname == "np" { return "Np" }
    return ""
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetSegmentPath() string {
    return "nps"
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "np" {
        for _, c := range nps.Np {
            if nps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np{}
        nps.Np = append(nps.Np, child)
        return &nps.Np[len(nps.Np)-1]
    }
    return nil
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nps.Np {
        children[nps.Np[i].GetSegmentPath()] = &nps.Np[i]
    }
    return children
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetBundleName() string { return "cisco_ios_xr" }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetYangName() string { return "nps" }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) SetParent(parent types.Entity) { nps.parent = parent }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetParent() types.Entity { return nps.parent }

func (nps *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps) GetParentYangName() string { return "ipolicer-local-table" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np
// Table of NP names
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is interface{} with range:
    // -2147483648..2147483647.
    Id1 interface{}

    // Packets per second. The type is interface{} with range:
    // -2147483648..2147483647. Units are packet/s.
    Rate interface{}
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetFilter() yfilter.YFilter { return np.YFilter }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) SetFilter(yf yfilter.YFilter) { np.YFilter = yf }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetGoName(yname string) string {
    if yname == "id1" { return "Id1" }
    if yname == "rate" { return "Rate" }
    return ""
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetSegmentPath() string {
    return "np" + "[id1='" + fmt.Sprintf("%v", np.Id1) + "']"
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id1"] = np.Id1
    leafs["rate"] = np.Rate
    return leafs
}

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetBundleName() string { return "cisco_ios_xr" }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetYangName() string { return "np" }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) SetParent(parent types.Entity) { np.parent = parent }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetParent() types.Entity { return np.parent }

func (np *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocalTables_IpolicerLocalTable_Nps_Np) GetParentYangName() string { return "nps" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table for Dynamic Flows. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable.
    DynamicFlowsTable []PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetFilter() yfilter.YFilter { return dynamicFlowsTables.YFilter }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) SetFilter(yf yfilter.YFilter) { dynamicFlowsTables.YFilter = yf }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetGoName(yname string) string {
    if yname == "dynamic-flows-table" { return "DynamicFlowsTable" }
    return ""
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetSegmentPath() string {
    return "dynamic-flows-tables"
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dynamic-flows-table" {
        for _, c := range dynamicFlowsTables.DynamicFlowsTable {
            if dynamicFlowsTables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable{}
        dynamicFlowsTables.DynamicFlowsTable = append(dynamicFlowsTables.DynamicFlowsTable, child)
        return &dynamicFlowsTables.DynamicFlowsTable[len(dynamicFlowsTables.DynamicFlowsTable)-1]
    }
    return nil
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicFlowsTables.DynamicFlowsTable {
        children[dynamicFlowsTables.DynamicFlowsTable[i].GetSegmentPath()] = &dynamicFlowsTables.DynamicFlowsTable[i]
    }
    return children
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetYangName() string { return "dynamic-flows-tables" }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) SetParent(parent types.Entity) { dynamicFlowsTables.parent = parent }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetParent() types.Entity { return dynamicFlowsTables.parent }

func (dynamicFlowsTables *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables) GetParentYangName() string { return "lpts-local" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable
// Table for Dynamic Flows
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Dynamic Flows Table Type. The type is
    // LptsDynamicFlowConfig.
    TableType interface{}

    // Selected flow type. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType.
    FlowType []PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetFilter() yfilter.YFilter { return dynamicFlowsTable.YFilter }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) SetFilter(yf yfilter.YFilter) { dynamicFlowsTable.YFilter = yf }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetGoName(yname string) string {
    if yname == "table-type" { return "TableType" }
    if yname == "flow-type" { return "FlowType" }
    return ""
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetSegmentPath() string {
    return "dynamic-flows-table" + "[table-type='" + fmt.Sprintf("%v", dynamicFlowsTable.TableType) + "']"
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-type" {
        for _, c := range dynamicFlowsTable.FlowType {
            if dynamicFlowsTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType{}
        dynamicFlowsTable.FlowType = append(dynamicFlowsTable.FlowType, child)
        return &dynamicFlowsTable.FlowType[len(dynamicFlowsTable.FlowType)-1]
    }
    return nil
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicFlowsTable.FlowType {
        children[dynamicFlowsTable.FlowType[i].GetSegmentPath()] = &dynamicFlowsTable.FlowType[i]
    }
    return children
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["table-type"] = dynamicFlowsTable.TableType
    return leafs
}

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetYangName() string { return "dynamic-flows-table" }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) SetParent(parent types.Entity) { dynamicFlowsTable.parent = parent }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetParent() types.Entity { return dynamicFlowsTable.parent }

func (dynamicFlowsTable *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable) GetParentYangName() string { return "dynamic-flows-tables" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType
// Selected flow type
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured Max TCAM value. The type is interface{} with range:
    // -2147483648..2147483647.
    Max interface{}
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetFilter() yfilter.YFilter { return flowType.YFilter }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) SetFilter(yf yfilter.YFilter) { flowType.YFilter = yf }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetGoName(yname string) string {
    if yname == "flow-type" { return "FlowType" }
    if yname == "max" { return "Max" }
    return ""
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetSegmentPath() string {
    return "flow-type" + "[flow-type='" + fmt.Sprintf("%v", flowType.FlowType) + "']"
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-type"] = flowType.FlowType
    leafs["max"] = flowType.Max
    return leafs
}

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetBundleName() string { return "cisco_ios_xr" }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetYangName() string { return "flow-type" }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) SetParent(parent types.Entity) { flowType.parent = parent }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetParent() types.Entity { return flowType.parent }

func (flowType *PreconfiguredNodes_PreconfiguredNode_LptsLocal_DynamicFlowsTables_DynamicFlowsTable_FlowType) GetParentYangName() string { return "dynamic-flows-table" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal
// Node specific Pre IFIB (Internal Forwarding
// Information Base) Configuration
// This type is a presence type.
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for Flows.
    Flows PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetFilter() yfilter.YFilter { return ipolicerLocal.YFilter }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) SetFilter(yf yfilter.YFilter) { ipolicerLocal.YFilter = yf }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "flows" { return "Flows" }
    return ""
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetSegmentPath() string {
    return "ipolicer-local"
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flows" {
        return &ipolicerLocal.Flows
    }
    return nil
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flows"] = &ipolicerLocal.Flows
    return children
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ipolicerLocal.Enable
    return leafs
}

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetYangName() string { return "ipolicer-local" }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) SetParent(parent types.Entity) { ipolicerLocal.parent = parent }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetParent() types.Entity { return ipolicerLocal.parent }

func (ipolicerLocal *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal) GetParentYangName() string { return "lpts-local" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows
// Table for Flows
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of
    // PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow.
    Flow []PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetFilter() yfilter.YFilter { return flows.YFilter }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) SetFilter(yf yfilter.YFilter) { flows.YFilter = yf }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetGoName(yname string) string {
    if yname == "flow" { return "Flow" }
    return ""
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetSegmentPath() string {
    return "flows"
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow" {
        for _, c := range flows.Flow {
            if flows.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow{}
        flows.Flow = append(flows.Flow, child)
        return &flows.Flow[len(flows.Flow)-1]
    }
    return nil
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flows.Flow {
        children[flows.Flow[i].GetSegmentPath()] = &flows.Flow[i]
    }
    return children
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetBundleName() string { return "cisco_ios_xr" }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetYangName() string { return "flows" }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) SetParent(parent types.Entity) { flows.parent = parent }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetParent() types.Entity { return flows.parent }

func (flows *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows) GetParentYangName() string { return "ipolicer-local" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow
// selected flow type
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range:
    // -2147483648..2147483647.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetGoName(yname string) string {
    if yname == "flow-type" { return "FlowType" }
    if yname == "rate" { return "Rate" }
    if yname == "precedences" { return "Precedences" }
    return ""
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetSegmentPath() string {
    return "flow" + "[flow-type='" + fmt.Sprintf("%v", flow.FlowType) + "']"
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "precedences" {
        return &flow.Precedences
    }
    return nil
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["precedences"] = &flow.Precedences
    return children
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-type"] = flow.FlowType
    leafs["rate"] = flow.Rate
    return leafs
}

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetYangName() string { return "flow" }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetParent() types.Entity { return flow.parent }

func (flow *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow) GetParentYangName() string { return "flows" }

// PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences
// TOS Precedence value(s)
type PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetFilter() yfilter.YFilter { return precedences.YFilter }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) SetFilter(yf yfilter.YFilter) { precedences.YFilter = yf }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetGoName(yname string) string {
    if yname == "precedence" { return "Precedence" }
    return ""
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetSegmentPath() string {
    return "precedences"
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["precedence"] = precedences.Precedence
    return leafs
}

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetBundleName() string { return "cisco_ios_xr" }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetYangName() string { return "precedences" }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) SetParent(parent types.Entity) { precedences.parent = parent }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetParent() types.Entity { return precedences.parent }

func (precedences *PreconfiguredNodes_PreconfiguredNode_LptsLocal_IpolicerLocal_Flows_Flow_Precedences) GetParentYangName() string { return "flow" }

// PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold
// watchdog node threshold
type PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory thresholds.
    MemoryThreshold PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetFilter() yfilter.YFilter { return ciscoIOSXRWatchdCfgWatchdogNodeThreshold.YFilter }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) SetFilter(yf yfilter.YFilter) { ciscoIOSXRWatchdCfgWatchdogNodeThreshold.YFilter = yf }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetGoName(yname string) string {
    if yname == "memory-threshold" { return "MemoryThreshold" }
    return ""
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetSegmentPath() string {
    return "Cisco-IOS-XR-watchd-cfg:watchdog-node-threshold"
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-threshold" {
        return &ciscoIOSXRWatchdCfgWatchdogNodeThreshold.MemoryThreshold
    }
    return nil
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["memory-threshold"] = &ciscoIOSXRWatchdCfgWatchdogNodeThreshold.MemoryThreshold
    return children
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetYangName() string { return "watchdog-node-threshold" }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) SetParent(parent types.Entity) { ciscoIOSXRWatchdCfgWatchdogNodeThreshold.parent = parent }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetParent() types.Entity { return ciscoIOSXRWatchdCfgWatchdogNodeThreshold.parent }

func (ciscoIOSXRWatchdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold) GetParentYangName() string { return "preconfigured-node" }

// PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetFilter() yfilter.YFilter { return memoryThreshold.YFilter }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) SetFilter(yf yfilter.YFilter) { memoryThreshold.YFilter = yf }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetSegmentPath() string {
    return "memory-threshold"
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = memoryThreshold.Minor
    leafs["severe"] = memoryThreshold.Severe
    leafs["critical"] = memoryThreshold.Critical
    return leafs
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetYangName() string { return "memory-threshold" }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) SetParent(parent types.Entity) { memoryThreshold.parent = parent }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetParent() types.Entity { return memoryThreshold.parent }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWatchdCfgWatchdogNodeThreshold_MemoryThreshold) GetParentYangName() string { return "watchdog-node-threshold" }

// PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold
// Watchdog threshold configuration
type PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Memory thresholds.
    MemoryThreshold PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetFilter() yfilter.YFilter { return ciscoIOSXRWdCfgWatchdogNodeThreshold.YFilter }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) SetFilter(yf yfilter.YFilter) { ciscoIOSXRWdCfgWatchdogNodeThreshold.YFilter = yf }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetGoName(yname string) string {
    if yname == "memory-threshold" { return "MemoryThreshold" }
    return ""
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetSegmentPath() string {
    return "Cisco-IOS-XR-wd-cfg:watchdog-node-threshold"
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-threshold" {
        return &ciscoIOSXRWdCfgWatchdogNodeThreshold.MemoryThreshold
    }
    return nil
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["memory-threshold"] = &ciscoIOSXRWdCfgWatchdogNodeThreshold.MemoryThreshold
    return children
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetYangName() string { return "watchdog-node-threshold" }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) SetParent(parent types.Entity) { ciscoIOSXRWdCfgWatchdogNodeThreshold.parent = parent }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetParent() types.Entity { return ciscoIOSXRWdCfgWatchdogNodeThreshold.parent }

func (ciscoIOSXRWdCfgWatchdogNodeThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold) GetParentYangName() string { return "preconfigured-node" }

// PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold
// Memory thresholds
type PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold, Range(5, 40). The type is interface{} with range: 5..40.
    Minor interface{}

    // Threshold, Range(4, minor). The type is interface{} with range: 4..40.
    Severe interface{}

    // Threshold, Range(3, severe). The type is interface{} with range: 3..40.
    Critical interface{}
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetFilter() yfilter.YFilter { return memoryThreshold.YFilter }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) SetFilter(yf yfilter.YFilter) { memoryThreshold.YFilter = yf }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetGoName(yname string) string {
    if yname == "minor" { return "Minor" }
    if yname == "severe" { return "Severe" }
    if yname == "critical" { return "Critical" }
    return ""
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetSegmentPath() string {
    return "memory-threshold"
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minor"] = memoryThreshold.Minor
    leafs["severe"] = memoryThreshold.Severe
    leafs["critical"] = memoryThreshold.Critical
    return leafs
}

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetYangName() string { return "memory-threshold" }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) SetParent(parent types.Entity) { memoryThreshold.parent = parent }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetParent() types.Entity { return memoryThreshold.parent }

func (memoryThreshold *PreconfiguredNodes_PreconfiguredNode_CiscoIOSXRWdCfgWatchdogNodeThreshold_MemoryThreshold) GetParentYangName() string { return "watchdog-node-threshold" }

