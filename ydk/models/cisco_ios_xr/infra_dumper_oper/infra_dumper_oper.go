// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-dumper package operational data.
// 
// This module contains definitions
// for the following management objects:
//   context: Core dump configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_dumper_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_dumper_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-dumper-oper context}", reflect.TypeOf(Context{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-dumper-oper:context", reflect.TypeOf(Context{}))
}

// Context
// Core dump configuration commands
type Context struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context number table.
    ContextNumbers Context_ContextNumbers

    // Core Context location table.
    ContextLocations Context_ContextLocations

    // context bag.
    All Context_All
}

func (context *Context) GetFilter() yfilter.YFilter { return context.YFilter }

func (context *Context) SetFilter(yf yfilter.YFilter) { context.YFilter = yf }

func (context *Context) GetGoName(yname string) string {
    if yname == "context-numbers" { return "ContextNumbers" }
    if yname == "context-locations" { return "ContextLocations" }
    if yname == "all" { return "All" }
    return ""
}

func (context *Context) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-dumper-oper:context"
}

func (context *Context) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-numbers" {
        return &context.ContextNumbers
    }
    if childYangName == "context-locations" {
        return &context.ContextLocations
    }
    if childYangName == "all" {
        return &context.All
    }
    return nil
}

func (context *Context) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["context-numbers"] = &context.ContextNumbers
    children["context-locations"] = &context.ContextLocations
    children["all"] = &context.All
    return children
}

func (context *Context) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (context *Context) GetBundleName() string { return "cisco_ios_xr" }

func (context *Context) GetYangName() string { return "context" }

func (context *Context) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (context *Context) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (context *Context) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (context *Context) SetParent(parent types.Entity) { context.parent = parent }

func (context *Context) GetParent() types.Entity { return context.parent }

func (context *Context) GetParentYangName() string { return "Cisco-IOS-XR-infra-dumper-oper" }

// Context_ContextNumbers
// Context number table
type Context_ContextNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context number . The type is slice of Context_ContextNumbers_ContextNumber.
    ContextNumber []Context_ContextNumbers_ContextNumber
}

func (contextNumbers *Context_ContextNumbers) GetFilter() yfilter.YFilter { return contextNumbers.YFilter }

func (contextNumbers *Context_ContextNumbers) SetFilter(yf yfilter.YFilter) { contextNumbers.YFilter = yf }

func (contextNumbers *Context_ContextNumbers) GetGoName(yname string) string {
    if yname == "context-number" { return "ContextNumber" }
    return ""
}

func (contextNumbers *Context_ContextNumbers) GetSegmentPath() string {
    return "context-numbers"
}

func (contextNumbers *Context_ContextNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-number" {
        for _, c := range contextNumbers.ContextNumber {
            if contextNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber{}
        contextNumbers.ContextNumber = append(contextNumbers.ContextNumber, child)
        return &contextNumbers.ContextNumber[len(contextNumbers.ContextNumber)-1]
    }
    return nil
}

func (contextNumbers *Context_ContextNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextNumbers.ContextNumber {
        children[contextNumbers.ContextNumber[i].GetSegmentPath()] = &contextNumbers.ContextNumber[i]
    }
    return children
}

func (contextNumbers *Context_ContextNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contextNumbers *Context_ContextNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (contextNumbers *Context_ContextNumbers) GetYangName() string { return "context-numbers" }

func (contextNumbers *Context_ContextNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextNumbers *Context_ContextNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextNumbers *Context_ContextNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextNumbers *Context_ContextNumbers) SetParent(parent types.Entity) { contextNumbers.parent = parent }

func (contextNumbers *Context_ContextNumbers) GetParent() types.Entity { return contextNumbers.parent }

func (contextNumbers *Context_ContextNumbers) GetParentYangName() string { return "context" }

// Context_ContextNumbers_ContextNumber
// Context number 
type Context_ContextNumbers_ContextNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context for which crash dump info required. The
    // type is interface{} with range: -2147483648..2147483647.
    ContextNum interface{}

    // Core Context location table.
    Locations Context_ContextNumbers_ContextNumber_Locations

    // context bag.
    All Context_ContextNumbers_ContextNumber_All
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetFilter() yfilter.YFilter { return contextNumber.YFilter }

func (contextNumber *Context_ContextNumbers_ContextNumber) SetFilter(yf yfilter.YFilter) { contextNumber.YFilter = yf }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetGoName(yname string) string {
    if yname == "context-num" { return "ContextNum" }
    if yname == "locations" { return "Locations" }
    if yname == "all" { return "All" }
    return ""
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetSegmentPath() string {
    return "context-number" + "[context-num='" + fmt.Sprintf("%v", contextNumber.ContextNum) + "']"
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "locations" {
        return &contextNumber.Locations
    }
    if childYangName == "all" {
        return &contextNumber.All
    }
    return nil
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["locations"] = &contextNumber.Locations
    children["all"] = &contextNumber.All
    return children
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-num"] = contextNumber.ContextNum
    return leafs
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetBundleName() string { return "cisco_ios_xr" }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetYangName() string { return "context-number" }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextNumber *Context_ContextNumbers_ContextNumber) SetParent(parent types.Entity) { contextNumber.parent = parent }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetParent() types.Entity { return contextNumber.parent }

func (contextNumber *Context_ContextNumbers_ContextNumber) GetParentYangName() string { return "context-numbers" }

// Context_ContextNumbers_ContextNumber_Locations
// Core Context location table
type Context_ContextNumbers_ContextNumber_Locations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational Context for a particular location. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location.
    Location []Context_ContextNumbers_ContextNumber_Locations_Location
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetFilter() yfilter.YFilter { return locations.YFilter }

func (locations *Context_ContextNumbers_ContextNumber_Locations) SetFilter(yf yfilter.YFilter) { locations.YFilter = yf }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetSegmentPath() string {
    return "locations"
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location" {
        for _, c := range locations.Location {
            if locations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location{}
        locations.Location = append(locations.Location, child)
        return &locations.Location[len(locations.Location)-1]
    }
    return nil
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range locations.Location {
        children[locations.Location[i].GetSegmentPath()] = &locations.Location[i]
    }
    return children
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetBundleName() string { return "cisco_ios_xr" }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetYangName() string { return "locations" }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (locations *Context_ContextNumbers_ContextNumber_Locations) SetParent(parent types.Entity) { locations.parent = parent }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetParent() types.Entity { return locations.parent }

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetParentYangName() string { return "context-number" }

// Context_ContextNumbers_ContextNumber_Locations_Location
// Operational Context for a particular location
type Context_ContextNumbers_ContextNumber_Locations_Location struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Context info bag.
    Enter Context_ContextNumbers_ContextNumber_Locations_Location_Enter
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetFilter() yfilter.YFilter { return location.YFilter }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) SetFilter(yf yfilter.YFilter) { location.YFilter = yf }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "enter" { return "Enter" }
    return ""
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetSegmentPath() string {
    return "location" + "[node-name='" + fmt.Sprintf("%v", location.NodeName) + "']"
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "enter" {
        return &location.Enter
    }
    return nil
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["enter"] = &location.Enter
    return children
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = location.NodeName
    return leafs
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetBundleName() string { return "cisco_ios_xr" }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetYangName() string { return "location" }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) SetParent(parent types.Entity) { location.parent = parent }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetParent() types.Entity { return location.parent }

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetParentYangName() string { return "locations" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter
// Context info bag
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo.
    CrashInfo []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetFilter() yfilter.YFilter { return enter.YFilter }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) SetFilter(yf yfilter.YFilter) { enter.YFilter = yf }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetGoName(yname string) string {
    if yname == "crash-info" { return "CrashInfo" }
    return ""
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetSegmentPath() string {
    return "enter"
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crash-info" {
        for _, c := range enter.CrashInfo {
            if enter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo{}
        enter.CrashInfo = append(enter.CrashInfo, child)
        return &enter.CrashInfo[len(enter.CrashInfo)-1]
    }
    return nil
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range enter.CrashInfo {
        children[enter.CrashInfo[i].GetSegmentPath()] = &enter.CrashInfo[i]
    }
    return children
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetBundleName() string { return "cisco_ios_xr" }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetYangName() string { return "enter" }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) SetParent(parent types.Entity) { enter.parent = parent }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetParent() types.Entity { return enter.parent }

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetParentYangName() string { return "location" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo
// All crash info
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is string.
    Node interface{}

    // Context Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo.
    ContextInfo []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo

    // Crash Package Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation.
    CrashPackageInformation []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetFilter() yfilter.YFilter { return crashInfo.YFilter }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) SetFilter(yf yfilter.YFilter) { crashInfo.YFilter = yf }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "context-info" { return "ContextInfo" }
    if yname == "crash-package-information" { return "CrashPackageInformation" }
    return ""
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetSegmentPath() string {
    return "crash-info"
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-info" {
        for _, c := range crashInfo.ContextInfo {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo{}
        crashInfo.ContextInfo = append(crashInfo.ContextInfo, child)
        return &crashInfo.ContextInfo[len(crashInfo.ContextInfo)-1]
    }
    if childYangName == "crash-package-information" {
        for _, c := range crashInfo.CrashPackageInformation {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation{}
        crashInfo.CrashPackageInformation = append(crashInfo.CrashPackageInformation, child)
        return &crashInfo.CrashPackageInformation[len(crashInfo.CrashPackageInformation)-1]
    }
    return nil
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crashInfo.ContextInfo {
        children[crashInfo.ContextInfo[i].GetSegmentPath()] = &crashInfo.ContextInfo[i]
    }
    for i := range crashInfo.CrashPackageInformation {
        children[crashInfo.CrashPackageInformation[i].GetSegmentPath()] = &crashInfo.CrashPackageInformation[i]
    }
    return children
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = crashInfo.Node
    return leafs
}

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetBundleName() string { return "cisco_ios_xr" }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetYangName() string { return "crash-info" }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) SetParent(parent types.Entity) { crashInfo.parent = parent }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetParent() types.Entity { return crashInfo.parent }

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetParentYangName() string { return "enter" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo
// Context Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Thread ID. The type is interface{} with range: 0..4294967295.
    Tid interface{}

    // Core dump time/Crash time. The type is string.
    CoreDumpTime interface{}

    // Signal number. The type is interface{} with range: 0..4294967295.
    SigNum interface{}

    // Signal error string. The type is string.
    SinErrStr interface{}

    // Sender pid. The type is interface{} with range: 0..4294967295.
    SigSendPid interface{}

    // Signal code. The type is interface{} with range: 0..4294967295.
    SigCode interface{}

    // Signal info. The type is string.
    SinInfo interface{}

    // Core for process at . The type is string.
    CoreForProcess interface{}

    // Registers Info. The type is string.
    RegistersInfo interface{}

    // Stack Trace. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace.
    StackTrace []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace

    // DLL Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo.
    DllInfo []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetFilter() yfilter.YFilter { return contextInfo.YFilter }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) SetFilter(yf yfilter.YFilter) { contextInfo.YFilter = yf }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "pid" { return "Pid" }
    if yname == "tid" { return "Tid" }
    if yname == "core-dump-time" { return "CoreDumpTime" }
    if yname == "sig-num" { return "SigNum" }
    if yname == "sin-err-str" { return "SinErrStr" }
    if yname == "sig-send-pid" { return "SigSendPid" }
    if yname == "sig-code" { return "SigCode" }
    if yname == "sin-info" { return "SinInfo" }
    if yname == "core-for-process" { return "CoreForProcess" }
    if yname == "registers-info" { return "RegistersInfo" }
    if yname == "stack-trace" { return "StackTrace" }
    if yname == "dll-info" { return "DllInfo" }
    return ""
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetSegmentPath() string {
    return "context-info"
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack-trace" {
        for _, c := range contextInfo.StackTrace {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace{}
        contextInfo.StackTrace = append(contextInfo.StackTrace, child)
        return &contextInfo.StackTrace[len(contextInfo.StackTrace)-1]
    }
    if childYangName == "dll-info" {
        for _, c := range contextInfo.DllInfo {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo{}
        contextInfo.DllInfo = append(contextInfo.DllInfo, child)
        return &contextInfo.DllInfo[len(contextInfo.DllInfo)-1]
    }
    return nil
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextInfo.StackTrace {
        children[contextInfo.StackTrace[i].GetSegmentPath()] = &contextInfo.StackTrace[i]
    }
    for i := range contextInfo.DllInfo {
        children[contextInfo.DllInfo[i].GetSegmentPath()] = &contextInfo.DllInfo[i]
    }
    return children
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = contextInfo.ProcessName
    leafs["pid"] = contextInfo.Pid
    leafs["tid"] = contextInfo.Tid
    leafs["core-dump-time"] = contextInfo.CoreDumpTime
    leafs["sig-num"] = contextInfo.SigNum
    leafs["sin-err-str"] = contextInfo.SinErrStr
    leafs["sig-send-pid"] = contextInfo.SigSendPid
    leafs["sig-code"] = contextInfo.SigCode
    leafs["sin-info"] = contextInfo.SinInfo
    leafs["core-for-process"] = contextInfo.CoreForProcess
    leafs["registers-info"] = contextInfo.RegistersInfo
    return leafs
}

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetBundleName() string { return "cisco_ios_xr" }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetYangName() string { return "context-info" }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) SetParent(parent types.Entity) { contextInfo.parent = parent }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetParent() types.Entity { return contextInfo.parent }

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetParentYangName() string { return "crash-info" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetFilter() yfilter.YFilter { return stackTrace.YFilter }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) SetFilter(yf yfilter.YFilter) { stackTrace.YFilter = yf }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetGoName(yname string) string {
    if yname == "stack-trace" { return "StackTrace" }
    return ""
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetSegmentPath() string {
    return "stack-trace"
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stack-trace"] = stackTrace.StackTrace
    return leafs
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetYangName() string { return "stack-trace" }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) SetParent(parent types.Entity) { stackTrace.parent = parent }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetParent() types.Entity { return stackTrace.parent }

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetParentYangName() string { return "context-info" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DLL Path. The type is string.
    Path interface{}

    // text addr. The type is interface{} with range: 0..4294967295.
    TextAddr interface{}

    // text size. The type is interface{} with range: 0..4294967295.
    TextSize interface{}

    // data addr. The type is interface{} with range: 0..4294967295.
    DataAddr interface{}

    // data size. The type is interface{} with range: 0..4294967295.
    DataSize interface{}

    // version. The type is interface{} with range: 0..4294967295.
    Version interface{}
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetFilter() yfilter.YFilter { return dllInfo.YFilter }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) SetFilter(yf yfilter.YFilter) { dllInfo.YFilter = yf }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "text-addr" { return "TextAddr" }
    if yname == "text-size" { return "TextSize" }
    if yname == "data-addr" { return "DataAddr" }
    if yname == "data-size" { return "DataSize" }
    if yname == "version" { return "Version" }
    return ""
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetSegmentPath() string {
    return "dll-info"
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = dllInfo.Path
    leafs["text-addr"] = dllInfo.TextAddr
    leafs["text-size"] = dllInfo.TextSize
    leafs["data-addr"] = dllInfo.DataAddr
    leafs["data-size"] = dllInfo.DataSize
    leafs["version"] = dllInfo.Version
    return leafs
}

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetYangName() string { return "dll-info" }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) SetParent(parent types.Entity) { dllInfo.parent = parent }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetParent() types.Entity { return dllInfo.parent }

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetParentYangName() string { return "context-info" }

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetFilter() yfilter.YFilter { return crashPackageInformation.YFilter }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) SetFilter(yf yfilter.YFilter) { crashPackageInformation.YFilter = yf }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source" { return "Source" }
    return ""
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetSegmentPath() string {
    return "crash-package-information"
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = crashPackageInformation.Name
    leafs["source"] = crashPackageInformation.Source
    return leafs
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetBundleName() string { return "cisco_ios_xr" }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetYangName() string { return "crash-package-information" }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) SetParent(parent types.Entity) { crashPackageInformation.parent = parent }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetParent() types.Entity { return crashPackageInformation.parent }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetParentYangName() string { return "crash-info" }

// Context_ContextNumbers_ContextNumber_All
// context bag
type Context_ContextNumbers_ContextNumber_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo.
    CrashInfo []Context_ContextNumbers_ContextNumber_All_CrashInfo
}

func (all *Context_ContextNumbers_ContextNumber_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *Context_ContextNumbers_ContextNumber_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *Context_ContextNumbers_ContextNumber_All) GetGoName(yname string) string {
    if yname == "crash-info" { return "CrashInfo" }
    return ""
}

func (all *Context_ContextNumbers_ContextNumber_All) GetSegmentPath() string {
    return "all"
}

func (all *Context_ContextNumbers_ContextNumber_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crash-info" {
        for _, c := range all.CrashInfo {
            if all.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_All_CrashInfo{}
        all.CrashInfo = append(all.CrashInfo, child)
        return &all.CrashInfo[len(all.CrashInfo)-1]
    }
    return nil
}

func (all *Context_ContextNumbers_ContextNumber_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range all.CrashInfo {
        children[all.CrashInfo[i].GetSegmentPath()] = &all.CrashInfo[i]
    }
    return children
}

func (all *Context_ContextNumbers_ContextNumber_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (all *Context_ContextNumbers_ContextNumber_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *Context_ContextNumbers_ContextNumber_All) GetYangName() string { return "all" }

func (all *Context_ContextNumbers_ContextNumber_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *Context_ContextNumbers_ContextNumber_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *Context_ContextNumbers_ContextNumber_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *Context_ContextNumbers_ContextNumber_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *Context_ContextNumbers_ContextNumber_All) GetParent() types.Entity { return all.parent }

func (all *Context_ContextNumbers_ContextNumber_All) GetParentYangName() string { return "context-number" }

// Context_ContextNumbers_ContextNumber_All_CrashInfo
// All crash info
type Context_ContextNumbers_ContextNumber_All_CrashInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is string.
    Node interface{}

    // Context Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo.
    ContextInfo []Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo

    // Crash Package Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation.
    CrashPackageInformation []Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetFilter() yfilter.YFilter { return crashInfo.YFilter }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) SetFilter(yf yfilter.YFilter) { crashInfo.YFilter = yf }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "context-info" { return "ContextInfo" }
    if yname == "crash-package-information" { return "CrashPackageInformation" }
    return ""
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetSegmentPath() string {
    return "crash-info"
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-info" {
        for _, c := range crashInfo.ContextInfo {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo{}
        crashInfo.ContextInfo = append(crashInfo.ContextInfo, child)
        return &crashInfo.ContextInfo[len(crashInfo.ContextInfo)-1]
    }
    if childYangName == "crash-package-information" {
        for _, c := range crashInfo.CrashPackageInformation {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation{}
        crashInfo.CrashPackageInformation = append(crashInfo.CrashPackageInformation, child)
        return &crashInfo.CrashPackageInformation[len(crashInfo.CrashPackageInformation)-1]
    }
    return nil
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crashInfo.ContextInfo {
        children[crashInfo.ContextInfo[i].GetSegmentPath()] = &crashInfo.ContextInfo[i]
    }
    for i := range crashInfo.CrashPackageInformation {
        children[crashInfo.CrashPackageInformation[i].GetSegmentPath()] = &crashInfo.CrashPackageInformation[i]
    }
    return children
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = crashInfo.Node
    return leafs
}

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetBundleName() string { return "cisco_ios_xr" }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetYangName() string { return "crash-info" }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) SetParent(parent types.Entity) { crashInfo.parent = parent }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetParent() types.Entity { return crashInfo.parent }

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetParentYangName() string { return "all" }

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo
// Context Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Thread ID. The type is interface{} with range: 0..4294967295.
    Tid interface{}

    // Core dump time/Crash time. The type is string.
    CoreDumpTime interface{}

    // Signal number. The type is interface{} with range: 0..4294967295.
    SigNum interface{}

    // Signal error string. The type is string.
    SinErrStr interface{}

    // Sender pid. The type is interface{} with range: 0..4294967295.
    SigSendPid interface{}

    // Signal code. The type is interface{} with range: 0..4294967295.
    SigCode interface{}

    // Signal info. The type is string.
    SinInfo interface{}

    // Core for process at . The type is string.
    CoreForProcess interface{}

    // Registers Info. The type is string.
    RegistersInfo interface{}

    // Stack Trace. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace.
    StackTrace []Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace

    // DLL Information. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo.
    DllInfo []Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetFilter() yfilter.YFilter { return contextInfo.YFilter }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) SetFilter(yf yfilter.YFilter) { contextInfo.YFilter = yf }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "pid" { return "Pid" }
    if yname == "tid" { return "Tid" }
    if yname == "core-dump-time" { return "CoreDumpTime" }
    if yname == "sig-num" { return "SigNum" }
    if yname == "sin-err-str" { return "SinErrStr" }
    if yname == "sig-send-pid" { return "SigSendPid" }
    if yname == "sig-code" { return "SigCode" }
    if yname == "sin-info" { return "SinInfo" }
    if yname == "core-for-process" { return "CoreForProcess" }
    if yname == "registers-info" { return "RegistersInfo" }
    if yname == "stack-trace" { return "StackTrace" }
    if yname == "dll-info" { return "DllInfo" }
    return ""
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetSegmentPath() string {
    return "context-info"
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack-trace" {
        for _, c := range contextInfo.StackTrace {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace{}
        contextInfo.StackTrace = append(contextInfo.StackTrace, child)
        return &contextInfo.StackTrace[len(contextInfo.StackTrace)-1]
    }
    if childYangName == "dll-info" {
        for _, c := range contextInfo.DllInfo {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo{}
        contextInfo.DllInfo = append(contextInfo.DllInfo, child)
        return &contextInfo.DllInfo[len(contextInfo.DllInfo)-1]
    }
    return nil
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextInfo.StackTrace {
        children[contextInfo.StackTrace[i].GetSegmentPath()] = &contextInfo.StackTrace[i]
    }
    for i := range contextInfo.DllInfo {
        children[contextInfo.DllInfo[i].GetSegmentPath()] = &contextInfo.DllInfo[i]
    }
    return children
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = contextInfo.ProcessName
    leafs["pid"] = contextInfo.Pid
    leafs["tid"] = contextInfo.Tid
    leafs["core-dump-time"] = contextInfo.CoreDumpTime
    leafs["sig-num"] = contextInfo.SigNum
    leafs["sin-err-str"] = contextInfo.SinErrStr
    leafs["sig-send-pid"] = contextInfo.SigSendPid
    leafs["sig-code"] = contextInfo.SigCode
    leafs["sin-info"] = contextInfo.SinInfo
    leafs["core-for-process"] = contextInfo.CoreForProcess
    leafs["registers-info"] = contextInfo.RegistersInfo
    return leafs
}

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetBundleName() string { return "cisco_ios_xr" }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetYangName() string { return "context-info" }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) SetParent(parent types.Entity) { contextInfo.parent = parent }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetParent() types.Entity { return contextInfo.parent }

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetParentYangName() string { return "crash-info" }

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetFilter() yfilter.YFilter { return stackTrace.YFilter }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) SetFilter(yf yfilter.YFilter) { stackTrace.YFilter = yf }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetGoName(yname string) string {
    if yname == "stack-trace" { return "StackTrace" }
    return ""
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetSegmentPath() string {
    return "stack-trace"
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stack-trace"] = stackTrace.StackTrace
    return leafs
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetYangName() string { return "stack-trace" }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) SetParent(parent types.Entity) { stackTrace.parent = parent }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetParent() types.Entity { return stackTrace.parent }

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetParentYangName() string { return "context-info" }

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DLL Path. The type is string.
    Path interface{}

    // text addr. The type is interface{} with range: 0..4294967295.
    TextAddr interface{}

    // text size. The type is interface{} with range: 0..4294967295.
    TextSize interface{}

    // data addr. The type is interface{} with range: 0..4294967295.
    DataAddr interface{}

    // data size. The type is interface{} with range: 0..4294967295.
    DataSize interface{}

    // version. The type is interface{} with range: 0..4294967295.
    Version interface{}
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetFilter() yfilter.YFilter { return dllInfo.YFilter }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) SetFilter(yf yfilter.YFilter) { dllInfo.YFilter = yf }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "text-addr" { return "TextAddr" }
    if yname == "text-size" { return "TextSize" }
    if yname == "data-addr" { return "DataAddr" }
    if yname == "data-size" { return "DataSize" }
    if yname == "version" { return "Version" }
    return ""
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetSegmentPath() string {
    return "dll-info"
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = dllInfo.Path
    leafs["text-addr"] = dllInfo.TextAddr
    leafs["text-size"] = dllInfo.TextSize
    leafs["data-addr"] = dllInfo.DataAddr
    leafs["data-size"] = dllInfo.DataSize
    leafs["version"] = dllInfo.Version
    return leafs
}

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetYangName() string { return "dll-info" }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) SetParent(parent types.Entity) { dllInfo.parent = parent }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetParent() types.Entity { return dllInfo.parent }

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetParentYangName() string { return "context-info" }

// Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetFilter() yfilter.YFilter { return crashPackageInformation.YFilter }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) SetFilter(yf yfilter.YFilter) { crashPackageInformation.YFilter = yf }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source" { return "Source" }
    return ""
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetSegmentPath() string {
    return "crash-package-information"
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = crashPackageInformation.Name
    leafs["source"] = crashPackageInformation.Source
    return leafs
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetBundleName() string { return "cisco_ios_xr" }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetYangName() string { return "crash-package-information" }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) SetParent(parent types.Entity) { crashPackageInformation.parent = parent }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetParent() types.Entity { return crashPackageInformation.parent }

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetParentYangName() string { return "crash-info" }

// Context_ContextLocations
// Core Context location table
type Context_ContextLocations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational Context for a particular location. The type is slice of
    // Context_ContextLocations_ContextLocation.
    ContextLocation []Context_ContextLocations_ContextLocation
}

func (contextLocations *Context_ContextLocations) GetFilter() yfilter.YFilter { return contextLocations.YFilter }

func (contextLocations *Context_ContextLocations) SetFilter(yf yfilter.YFilter) { contextLocations.YFilter = yf }

func (contextLocations *Context_ContextLocations) GetGoName(yname string) string {
    if yname == "context-location" { return "ContextLocation" }
    return ""
}

func (contextLocations *Context_ContextLocations) GetSegmentPath() string {
    return "context-locations"
}

func (contextLocations *Context_ContextLocations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-location" {
        for _, c := range contextLocations.ContextLocation {
            if contextLocations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation{}
        contextLocations.ContextLocation = append(contextLocations.ContextLocation, child)
        return &contextLocations.ContextLocation[len(contextLocations.ContextLocation)-1]
    }
    return nil
}

func (contextLocations *Context_ContextLocations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextLocations.ContextLocation {
        children[contextLocations.ContextLocation[i].GetSegmentPath()] = &contextLocations.ContextLocation[i]
    }
    return children
}

func (contextLocations *Context_ContextLocations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contextLocations *Context_ContextLocations) GetBundleName() string { return "cisco_ios_xr" }

func (contextLocations *Context_ContextLocations) GetYangName() string { return "context-locations" }

func (contextLocations *Context_ContextLocations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextLocations *Context_ContextLocations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextLocations *Context_ContextLocations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextLocations *Context_ContextLocations) SetParent(parent types.Entity) { contextLocations.parent = parent }

func (contextLocations *Context_ContextLocations) GetParent() types.Entity { return contextLocations.parent }

func (contextLocations *Context_ContextLocations) GetParentYangName() string { return "context" }

// Context_ContextLocations_ContextLocation
// Operational Context for a particular location
type Context_ContextLocations_ContextLocation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Context number Table.
    Numbers Context_ContextLocations_ContextLocation_Numbers

    // context bag.
    All Context_ContextLocations_ContextLocation_All
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetFilter() yfilter.YFilter { return contextLocation.YFilter }

func (contextLocation *Context_ContextLocations_ContextLocation) SetFilter(yf yfilter.YFilter) { contextLocation.YFilter = yf }

func (contextLocation *Context_ContextLocations_ContextLocation) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "numbers" { return "Numbers" }
    if yname == "all" { return "All" }
    return ""
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetSegmentPath() string {
    return "context-location" + "[node-name='" + fmt.Sprintf("%v", contextLocation.NodeName) + "']"
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "numbers" {
        return &contextLocation.Numbers
    }
    if childYangName == "all" {
        return &contextLocation.All
    }
    return nil
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["numbers"] = &contextLocation.Numbers
    children["all"] = &contextLocation.All
    return children
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = contextLocation.NodeName
    return leafs
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetBundleName() string { return "cisco_ios_xr" }

func (contextLocation *Context_ContextLocations_ContextLocation) GetYangName() string { return "context-location" }

func (contextLocation *Context_ContextLocations_ContextLocation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextLocation *Context_ContextLocations_ContextLocation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextLocation *Context_ContextLocations_ContextLocation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextLocation *Context_ContextLocations_ContextLocation) SetParent(parent types.Entity) { contextLocation.parent = parent }

func (contextLocation *Context_ContextLocations_ContextLocation) GetParent() types.Entity { return contextLocation.parent }

func (contextLocation *Context_ContextLocations_ContextLocation) GetParentYangName() string { return "context-locations" }

// Context_ContextLocations_ContextLocation_Numbers
// Context number Table
type Context_ContextLocations_ContextLocation_Numbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context number. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number.
    Number []Context_ContextLocations_ContextLocation_Numbers_Number
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetFilter() yfilter.YFilter { return numbers.YFilter }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) SetFilter(yf yfilter.YFilter) { numbers.YFilter = yf }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    return ""
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetSegmentPath() string {
    return "numbers"
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "number" {
        for _, c := range numbers.Number {
            if numbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number{}
        numbers.Number = append(numbers.Number, child)
        return &numbers.Number[len(numbers.Number)-1]
    }
    return nil
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range numbers.Number {
        children[numbers.Number[i].GetSegmentPath()] = &numbers.Number[i]
    }
    return children
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetBundleName() string { return "cisco_ios_xr" }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetYangName() string { return "numbers" }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) SetParent(parent types.Entity) { numbers.parent = parent }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetParent() types.Entity { return numbers.parent }

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetParentYangName() string { return "context-location" }

// Context_ContextLocations_ContextLocation_Numbers_Number
// Context number
type Context_ContextLocations_ContextLocation_Numbers_Number struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context for which crash dump info required. The
    // type is interface{} with range: -2147483648..2147483647.
    ContextNum interface{}

    // Context info bag.
    Enter Context_ContextLocations_ContextLocation_Numbers_Number_Enter
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetFilter() yfilter.YFilter { return number.YFilter }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) SetFilter(yf yfilter.YFilter) { number.YFilter = yf }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetGoName(yname string) string {
    if yname == "context-num" { return "ContextNum" }
    if yname == "enter" { return "Enter" }
    return ""
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetSegmentPath() string {
    return "number" + "[context-num='" + fmt.Sprintf("%v", number.ContextNum) + "']"
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "enter" {
        return &number.Enter
    }
    return nil
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["enter"] = &number.Enter
    return children
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-num"] = number.ContextNum
    return leafs
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetBundleName() string { return "cisco_ios_xr" }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetYangName() string { return "number" }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) SetParent(parent types.Entity) { number.parent = parent }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetParent() types.Entity { return number.parent }

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetParentYangName() string { return "numbers" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter
// Context info bag
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo.
    CrashInfo []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetFilter() yfilter.YFilter { return enter.YFilter }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) SetFilter(yf yfilter.YFilter) { enter.YFilter = yf }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetGoName(yname string) string {
    if yname == "crash-info" { return "CrashInfo" }
    return ""
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetSegmentPath() string {
    return "enter"
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crash-info" {
        for _, c := range enter.CrashInfo {
            if enter.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo{}
        enter.CrashInfo = append(enter.CrashInfo, child)
        return &enter.CrashInfo[len(enter.CrashInfo)-1]
    }
    return nil
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range enter.CrashInfo {
        children[enter.CrashInfo[i].GetSegmentPath()] = &enter.CrashInfo[i]
    }
    return children
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetBundleName() string { return "cisco_ios_xr" }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetYangName() string { return "enter" }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) SetParent(parent types.Entity) { enter.parent = parent }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetParent() types.Entity { return enter.parent }

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetParentYangName() string { return "number" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo
// All crash info
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is string.
    Node interface{}

    // Context Information. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo.
    ContextInfo []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo

    // Crash Package Information. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation.
    CrashPackageInformation []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetFilter() yfilter.YFilter { return crashInfo.YFilter }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) SetFilter(yf yfilter.YFilter) { crashInfo.YFilter = yf }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "context-info" { return "ContextInfo" }
    if yname == "crash-package-information" { return "CrashPackageInformation" }
    return ""
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetSegmentPath() string {
    return "crash-info"
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-info" {
        for _, c := range crashInfo.ContextInfo {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo{}
        crashInfo.ContextInfo = append(crashInfo.ContextInfo, child)
        return &crashInfo.ContextInfo[len(crashInfo.ContextInfo)-1]
    }
    if childYangName == "crash-package-information" {
        for _, c := range crashInfo.CrashPackageInformation {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation{}
        crashInfo.CrashPackageInformation = append(crashInfo.CrashPackageInformation, child)
        return &crashInfo.CrashPackageInformation[len(crashInfo.CrashPackageInformation)-1]
    }
    return nil
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crashInfo.ContextInfo {
        children[crashInfo.ContextInfo[i].GetSegmentPath()] = &crashInfo.ContextInfo[i]
    }
    for i := range crashInfo.CrashPackageInformation {
        children[crashInfo.CrashPackageInformation[i].GetSegmentPath()] = &crashInfo.CrashPackageInformation[i]
    }
    return children
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = crashInfo.Node
    return leafs
}

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetBundleName() string { return "cisco_ios_xr" }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetYangName() string { return "crash-info" }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) SetParent(parent types.Entity) { crashInfo.parent = parent }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetParent() types.Entity { return crashInfo.parent }

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetParentYangName() string { return "enter" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo
// Context Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Thread ID. The type is interface{} with range: 0..4294967295.
    Tid interface{}

    // Core dump time/Crash time. The type is string.
    CoreDumpTime interface{}

    // Signal number. The type is interface{} with range: 0..4294967295.
    SigNum interface{}

    // Signal error string. The type is string.
    SinErrStr interface{}

    // Sender pid. The type is interface{} with range: 0..4294967295.
    SigSendPid interface{}

    // Signal code. The type is interface{} with range: 0..4294967295.
    SigCode interface{}

    // Signal info. The type is string.
    SinInfo interface{}

    // Core for process at . The type is string.
    CoreForProcess interface{}

    // Registers Info. The type is string.
    RegistersInfo interface{}

    // Stack Trace. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace.
    StackTrace []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace

    // DLL Information. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo.
    DllInfo []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetFilter() yfilter.YFilter { return contextInfo.YFilter }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) SetFilter(yf yfilter.YFilter) { contextInfo.YFilter = yf }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "pid" { return "Pid" }
    if yname == "tid" { return "Tid" }
    if yname == "core-dump-time" { return "CoreDumpTime" }
    if yname == "sig-num" { return "SigNum" }
    if yname == "sin-err-str" { return "SinErrStr" }
    if yname == "sig-send-pid" { return "SigSendPid" }
    if yname == "sig-code" { return "SigCode" }
    if yname == "sin-info" { return "SinInfo" }
    if yname == "core-for-process" { return "CoreForProcess" }
    if yname == "registers-info" { return "RegistersInfo" }
    if yname == "stack-trace" { return "StackTrace" }
    if yname == "dll-info" { return "DllInfo" }
    return ""
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetSegmentPath() string {
    return "context-info"
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack-trace" {
        for _, c := range contextInfo.StackTrace {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace{}
        contextInfo.StackTrace = append(contextInfo.StackTrace, child)
        return &contextInfo.StackTrace[len(contextInfo.StackTrace)-1]
    }
    if childYangName == "dll-info" {
        for _, c := range contextInfo.DllInfo {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo{}
        contextInfo.DllInfo = append(contextInfo.DllInfo, child)
        return &contextInfo.DllInfo[len(contextInfo.DllInfo)-1]
    }
    return nil
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextInfo.StackTrace {
        children[contextInfo.StackTrace[i].GetSegmentPath()] = &contextInfo.StackTrace[i]
    }
    for i := range contextInfo.DllInfo {
        children[contextInfo.DllInfo[i].GetSegmentPath()] = &contextInfo.DllInfo[i]
    }
    return children
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = contextInfo.ProcessName
    leafs["pid"] = contextInfo.Pid
    leafs["tid"] = contextInfo.Tid
    leafs["core-dump-time"] = contextInfo.CoreDumpTime
    leafs["sig-num"] = contextInfo.SigNum
    leafs["sin-err-str"] = contextInfo.SinErrStr
    leafs["sig-send-pid"] = contextInfo.SigSendPid
    leafs["sig-code"] = contextInfo.SigCode
    leafs["sin-info"] = contextInfo.SinInfo
    leafs["core-for-process"] = contextInfo.CoreForProcess
    leafs["registers-info"] = contextInfo.RegistersInfo
    return leafs
}

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetBundleName() string { return "cisco_ios_xr" }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetYangName() string { return "context-info" }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) SetParent(parent types.Entity) { contextInfo.parent = parent }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetParent() types.Entity { return contextInfo.parent }

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetParentYangName() string { return "crash-info" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetFilter() yfilter.YFilter { return stackTrace.YFilter }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) SetFilter(yf yfilter.YFilter) { stackTrace.YFilter = yf }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetGoName(yname string) string {
    if yname == "stack-trace" { return "StackTrace" }
    return ""
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetSegmentPath() string {
    return "stack-trace"
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stack-trace"] = stackTrace.StackTrace
    return leafs
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetYangName() string { return "stack-trace" }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) SetParent(parent types.Entity) { stackTrace.parent = parent }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetParent() types.Entity { return stackTrace.parent }

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetParentYangName() string { return "context-info" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DLL Path. The type is string.
    Path interface{}

    // text addr. The type is interface{} with range: 0..4294967295.
    TextAddr interface{}

    // text size. The type is interface{} with range: 0..4294967295.
    TextSize interface{}

    // data addr. The type is interface{} with range: 0..4294967295.
    DataAddr interface{}

    // data size. The type is interface{} with range: 0..4294967295.
    DataSize interface{}

    // version. The type is interface{} with range: 0..4294967295.
    Version interface{}
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetFilter() yfilter.YFilter { return dllInfo.YFilter }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) SetFilter(yf yfilter.YFilter) { dllInfo.YFilter = yf }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "text-addr" { return "TextAddr" }
    if yname == "text-size" { return "TextSize" }
    if yname == "data-addr" { return "DataAddr" }
    if yname == "data-size" { return "DataSize" }
    if yname == "version" { return "Version" }
    return ""
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetSegmentPath() string {
    return "dll-info"
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = dllInfo.Path
    leafs["text-addr"] = dllInfo.TextAddr
    leafs["text-size"] = dllInfo.TextSize
    leafs["data-addr"] = dllInfo.DataAddr
    leafs["data-size"] = dllInfo.DataSize
    leafs["version"] = dllInfo.Version
    return leafs
}

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetYangName() string { return "dll-info" }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) SetParent(parent types.Entity) { dllInfo.parent = parent }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetParent() types.Entity { return dllInfo.parent }

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetParentYangName() string { return "context-info" }

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetFilter() yfilter.YFilter { return crashPackageInformation.YFilter }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) SetFilter(yf yfilter.YFilter) { crashPackageInformation.YFilter = yf }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source" { return "Source" }
    return ""
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetSegmentPath() string {
    return "crash-package-information"
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = crashPackageInformation.Name
    leafs["source"] = crashPackageInformation.Source
    return leafs
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetBundleName() string { return "cisco_ios_xr" }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetYangName() string { return "crash-package-information" }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) SetParent(parent types.Entity) { crashPackageInformation.parent = parent }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetParent() types.Entity { return crashPackageInformation.parent }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetParentYangName() string { return "crash-info" }

// Context_ContextLocations_ContextLocation_All
// context bag
type Context_ContextLocations_ContextLocation_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo.
    CrashInfo []Context_ContextLocations_ContextLocation_All_CrashInfo
}

func (all *Context_ContextLocations_ContextLocation_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *Context_ContextLocations_ContextLocation_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *Context_ContextLocations_ContextLocation_All) GetGoName(yname string) string {
    if yname == "crash-info" { return "CrashInfo" }
    return ""
}

func (all *Context_ContextLocations_ContextLocation_All) GetSegmentPath() string {
    return "all"
}

func (all *Context_ContextLocations_ContextLocation_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crash-info" {
        for _, c := range all.CrashInfo {
            if all.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_All_CrashInfo{}
        all.CrashInfo = append(all.CrashInfo, child)
        return &all.CrashInfo[len(all.CrashInfo)-1]
    }
    return nil
}

func (all *Context_ContextLocations_ContextLocation_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range all.CrashInfo {
        children[all.CrashInfo[i].GetSegmentPath()] = &all.CrashInfo[i]
    }
    return children
}

func (all *Context_ContextLocations_ContextLocation_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (all *Context_ContextLocations_ContextLocation_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *Context_ContextLocations_ContextLocation_All) GetYangName() string { return "all" }

func (all *Context_ContextLocations_ContextLocation_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *Context_ContextLocations_ContextLocation_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *Context_ContextLocations_ContextLocation_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *Context_ContextLocations_ContextLocation_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *Context_ContextLocations_ContextLocation_All) GetParent() types.Entity { return all.parent }

func (all *Context_ContextLocations_ContextLocation_All) GetParentYangName() string { return "context-location" }

// Context_ContextLocations_ContextLocation_All_CrashInfo
// All crash info
type Context_ContextLocations_ContextLocation_All_CrashInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is string.
    Node interface{}

    // Context Information. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo.
    ContextInfo []Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo

    // Crash Package Information. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation.
    CrashPackageInformation []Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetFilter() yfilter.YFilter { return crashInfo.YFilter }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) SetFilter(yf yfilter.YFilter) { crashInfo.YFilter = yf }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "context-info" { return "ContextInfo" }
    if yname == "crash-package-information" { return "CrashPackageInformation" }
    return ""
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetSegmentPath() string {
    return "crash-info"
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-info" {
        for _, c := range crashInfo.ContextInfo {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo{}
        crashInfo.ContextInfo = append(crashInfo.ContextInfo, child)
        return &crashInfo.ContextInfo[len(crashInfo.ContextInfo)-1]
    }
    if childYangName == "crash-package-information" {
        for _, c := range crashInfo.CrashPackageInformation {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation{}
        crashInfo.CrashPackageInformation = append(crashInfo.CrashPackageInformation, child)
        return &crashInfo.CrashPackageInformation[len(crashInfo.CrashPackageInformation)-1]
    }
    return nil
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crashInfo.ContextInfo {
        children[crashInfo.ContextInfo[i].GetSegmentPath()] = &crashInfo.ContextInfo[i]
    }
    for i := range crashInfo.CrashPackageInformation {
        children[crashInfo.CrashPackageInformation[i].GetSegmentPath()] = &crashInfo.CrashPackageInformation[i]
    }
    return children
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = crashInfo.Node
    return leafs
}

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetBundleName() string { return "cisco_ios_xr" }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetYangName() string { return "crash-info" }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) SetParent(parent types.Entity) { crashInfo.parent = parent }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetParent() types.Entity { return crashInfo.parent }

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetParentYangName() string { return "all" }

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo
// Context Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Thread ID. The type is interface{} with range: 0..4294967295.
    Tid interface{}

    // Core dump time/Crash time. The type is string.
    CoreDumpTime interface{}

    // Signal number. The type is interface{} with range: 0..4294967295.
    SigNum interface{}

    // Signal error string. The type is string.
    SinErrStr interface{}

    // Sender pid. The type is interface{} with range: 0..4294967295.
    SigSendPid interface{}

    // Signal code. The type is interface{} with range: 0..4294967295.
    SigCode interface{}

    // Signal info. The type is string.
    SinInfo interface{}

    // Core for process at . The type is string.
    CoreForProcess interface{}

    // Registers Info. The type is string.
    RegistersInfo interface{}

    // Stack Trace. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace.
    StackTrace []Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace

    // DLL Information. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo.
    DllInfo []Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetFilter() yfilter.YFilter { return contextInfo.YFilter }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) SetFilter(yf yfilter.YFilter) { contextInfo.YFilter = yf }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "pid" { return "Pid" }
    if yname == "tid" { return "Tid" }
    if yname == "core-dump-time" { return "CoreDumpTime" }
    if yname == "sig-num" { return "SigNum" }
    if yname == "sin-err-str" { return "SinErrStr" }
    if yname == "sig-send-pid" { return "SigSendPid" }
    if yname == "sig-code" { return "SigCode" }
    if yname == "sin-info" { return "SinInfo" }
    if yname == "core-for-process" { return "CoreForProcess" }
    if yname == "registers-info" { return "RegistersInfo" }
    if yname == "stack-trace" { return "StackTrace" }
    if yname == "dll-info" { return "DllInfo" }
    return ""
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetSegmentPath() string {
    return "context-info"
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack-trace" {
        for _, c := range contextInfo.StackTrace {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace{}
        contextInfo.StackTrace = append(contextInfo.StackTrace, child)
        return &contextInfo.StackTrace[len(contextInfo.StackTrace)-1]
    }
    if childYangName == "dll-info" {
        for _, c := range contextInfo.DllInfo {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo{}
        contextInfo.DllInfo = append(contextInfo.DllInfo, child)
        return &contextInfo.DllInfo[len(contextInfo.DllInfo)-1]
    }
    return nil
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextInfo.StackTrace {
        children[contextInfo.StackTrace[i].GetSegmentPath()] = &contextInfo.StackTrace[i]
    }
    for i := range contextInfo.DllInfo {
        children[contextInfo.DllInfo[i].GetSegmentPath()] = &contextInfo.DllInfo[i]
    }
    return children
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = contextInfo.ProcessName
    leafs["pid"] = contextInfo.Pid
    leafs["tid"] = contextInfo.Tid
    leafs["core-dump-time"] = contextInfo.CoreDumpTime
    leafs["sig-num"] = contextInfo.SigNum
    leafs["sin-err-str"] = contextInfo.SinErrStr
    leafs["sig-send-pid"] = contextInfo.SigSendPid
    leafs["sig-code"] = contextInfo.SigCode
    leafs["sin-info"] = contextInfo.SinInfo
    leafs["core-for-process"] = contextInfo.CoreForProcess
    leafs["registers-info"] = contextInfo.RegistersInfo
    return leafs
}

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetBundleName() string { return "cisco_ios_xr" }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetYangName() string { return "context-info" }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) SetParent(parent types.Entity) { contextInfo.parent = parent }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetParent() types.Entity { return contextInfo.parent }

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetParentYangName() string { return "crash-info" }

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetFilter() yfilter.YFilter { return stackTrace.YFilter }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) SetFilter(yf yfilter.YFilter) { stackTrace.YFilter = yf }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetGoName(yname string) string {
    if yname == "stack-trace" { return "StackTrace" }
    return ""
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetSegmentPath() string {
    return "stack-trace"
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stack-trace"] = stackTrace.StackTrace
    return leafs
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetYangName() string { return "stack-trace" }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) SetParent(parent types.Entity) { stackTrace.parent = parent }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetParent() types.Entity { return stackTrace.parent }

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetParentYangName() string { return "context-info" }

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DLL Path. The type is string.
    Path interface{}

    // text addr. The type is interface{} with range: 0..4294967295.
    TextAddr interface{}

    // text size. The type is interface{} with range: 0..4294967295.
    TextSize interface{}

    // data addr. The type is interface{} with range: 0..4294967295.
    DataAddr interface{}

    // data size. The type is interface{} with range: 0..4294967295.
    DataSize interface{}

    // version. The type is interface{} with range: 0..4294967295.
    Version interface{}
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetFilter() yfilter.YFilter { return dllInfo.YFilter }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) SetFilter(yf yfilter.YFilter) { dllInfo.YFilter = yf }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "text-addr" { return "TextAddr" }
    if yname == "text-size" { return "TextSize" }
    if yname == "data-addr" { return "DataAddr" }
    if yname == "data-size" { return "DataSize" }
    if yname == "version" { return "Version" }
    return ""
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetSegmentPath() string {
    return "dll-info"
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = dllInfo.Path
    leafs["text-addr"] = dllInfo.TextAddr
    leafs["text-size"] = dllInfo.TextSize
    leafs["data-addr"] = dllInfo.DataAddr
    leafs["data-size"] = dllInfo.DataSize
    leafs["version"] = dllInfo.Version
    return leafs
}

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetYangName() string { return "dll-info" }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) SetParent(parent types.Entity) { dllInfo.parent = parent }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetParent() types.Entity { return dllInfo.parent }

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetParentYangName() string { return "context-info" }

// Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetFilter() yfilter.YFilter { return crashPackageInformation.YFilter }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) SetFilter(yf yfilter.YFilter) { crashPackageInformation.YFilter = yf }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source" { return "Source" }
    return ""
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetSegmentPath() string {
    return "crash-package-information"
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = crashPackageInformation.Name
    leafs["source"] = crashPackageInformation.Source
    return leafs
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetBundleName() string { return "cisco_ios_xr" }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetYangName() string { return "crash-package-information" }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) SetParent(parent types.Entity) { crashPackageInformation.parent = parent }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetParent() types.Entity { return crashPackageInformation.parent }

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetParentYangName() string { return "crash-info" }

// Context_All
// context bag
type Context_All struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All crash info. The type is slice of Context_All_CrashInfo.
    CrashInfo []Context_All_CrashInfo
}

func (all *Context_All) GetFilter() yfilter.YFilter { return all.YFilter }

func (all *Context_All) SetFilter(yf yfilter.YFilter) { all.YFilter = yf }

func (all *Context_All) GetGoName(yname string) string {
    if yname == "crash-info" { return "CrashInfo" }
    return ""
}

func (all *Context_All) GetSegmentPath() string {
    return "all"
}

func (all *Context_All) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "crash-info" {
        for _, c := range all.CrashInfo {
            if all.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_All_CrashInfo{}
        all.CrashInfo = append(all.CrashInfo, child)
        return &all.CrashInfo[len(all.CrashInfo)-1]
    }
    return nil
}

func (all *Context_All) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range all.CrashInfo {
        children[all.CrashInfo[i].GetSegmentPath()] = &all.CrashInfo[i]
    }
    return children
}

func (all *Context_All) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (all *Context_All) GetBundleName() string { return "cisco_ios_xr" }

func (all *Context_All) GetYangName() string { return "all" }

func (all *Context_All) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (all *Context_All) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (all *Context_All) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (all *Context_All) SetParent(parent types.Entity) { all.parent = parent }

func (all *Context_All) GetParent() types.Entity { return all.parent }

func (all *Context_All) GetParentYangName() string { return "context" }

// Context_All_CrashInfo
// All crash info
type Context_All_CrashInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node name. The type is string.
    Node interface{}

    // Context Information. The type is slice of
    // Context_All_CrashInfo_ContextInfo.
    ContextInfo []Context_All_CrashInfo_ContextInfo

    // Crash Package Information. The type is slice of
    // Context_All_CrashInfo_CrashPackageInformation.
    CrashPackageInformation []Context_All_CrashInfo_CrashPackageInformation
}

func (crashInfo *Context_All_CrashInfo) GetFilter() yfilter.YFilter { return crashInfo.YFilter }

func (crashInfo *Context_All_CrashInfo) SetFilter(yf yfilter.YFilter) { crashInfo.YFilter = yf }

func (crashInfo *Context_All_CrashInfo) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "context-info" { return "ContextInfo" }
    if yname == "crash-package-information" { return "CrashPackageInformation" }
    return ""
}

func (crashInfo *Context_All_CrashInfo) GetSegmentPath() string {
    return "crash-info"
}

func (crashInfo *Context_All_CrashInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-info" {
        for _, c := range crashInfo.ContextInfo {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_All_CrashInfo_ContextInfo{}
        crashInfo.ContextInfo = append(crashInfo.ContextInfo, child)
        return &crashInfo.ContextInfo[len(crashInfo.ContextInfo)-1]
    }
    if childYangName == "crash-package-information" {
        for _, c := range crashInfo.CrashPackageInformation {
            if crashInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_All_CrashInfo_CrashPackageInformation{}
        crashInfo.CrashPackageInformation = append(crashInfo.CrashPackageInformation, child)
        return &crashInfo.CrashPackageInformation[len(crashInfo.CrashPackageInformation)-1]
    }
    return nil
}

func (crashInfo *Context_All_CrashInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crashInfo.ContextInfo {
        children[crashInfo.ContextInfo[i].GetSegmentPath()] = &crashInfo.ContextInfo[i]
    }
    for i := range crashInfo.CrashPackageInformation {
        children[crashInfo.CrashPackageInformation[i].GetSegmentPath()] = &crashInfo.CrashPackageInformation[i]
    }
    return children
}

func (crashInfo *Context_All_CrashInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = crashInfo.Node
    return leafs
}

func (crashInfo *Context_All_CrashInfo) GetBundleName() string { return "cisco_ios_xr" }

func (crashInfo *Context_All_CrashInfo) GetYangName() string { return "crash-info" }

func (crashInfo *Context_All_CrashInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashInfo *Context_All_CrashInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashInfo *Context_All_CrashInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashInfo *Context_All_CrashInfo) SetParent(parent types.Entity) { crashInfo.parent = parent }

func (crashInfo *Context_All_CrashInfo) GetParent() types.Entity { return crashInfo.parent }

func (crashInfo *Context_All_CrashInfo) GetParentYangName() string { return "all" }

// Context_All_CrashInfo_ContextInfo
// Context Information
type Context_All_CrashInfo_ContextInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Process name. The type is string.
    ProcessName interface{}

    // Process ID. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Thread ID. The type is interface{} with range: 0..4294967295.
    Tid interface{}

    // Core dump time/Crash time. The type is string.
    CoreDumpTime interface{}

    // Signal number. The type is interface{} with range: 0..4294967295.
    SigNum interface{}

    // Signal error string. The type is string.
    SinErrStr interface{}

    // Sender pid. The type is interface{} with range: 0..4294967295.
    SigSendPid interface{}

    // Signal code. The type is interface{} with range: 0..4294967295.
    SigCode interface{}

    // Signal info. The type is string.
    SinInfo interface{}

    // Core for process at . The type is string.
    CoreForProcess interface{}

    // Registers Info. The type is string.
    RegistersInfo interface{}

    // Stack Trace. The type is slice of
    // Context_All_CrashInfo_ContextInfo_StackTrace.
    StackTrace []Context_All_CrashInfo_ContextInfo_StackTrace

    // DLL Information. The type is slice of
    // Context_All_CrashInfo_ContextInfo_DllInfo.
    DllInfo []Context_All_CrashInfo_ContextInfo_DllInfo
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetFilter() yfilter.YFilter { return contextInfo.YFilter }

func (contextInfo *Context_All_CrashInfo_ContextInfo) SetFilter(yf yfilter.YFilter) { contextInfo.YFilter = yf }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    if yname == "pid" { return "Pid" }
    if yname == "tid" { return "Tid" }
    if yname == "core-dump-time" { return "CoreDumpTime" }
    if yname == "sig-num" { return "SigNum" }
    if yname == "sin-err-str" { return "SinErrStr" }
    if yname == "sig-send-pid" { return "SigSendPid" }
    if yname == "sig-code" { return "SigCode" }
    if yname == "sin-info" { return "SinInfo" }
    if yname == "core-for-process" { return "CoreForProcess" }
    if yname == "registers-info" { return "RegistersInfo" }
    if yname == "stack-trace" { return "StackTrace" }
    if yname == "dll-info" { return "DllInfo" }
    return ""
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetSegmentPath() string {
    return "context-info"
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stack-trace" {
        for _, c := range contextInfo.StackTrace {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_All_CrashInfo_ContextInfo_StackTrace{}
        contextInfo.StackTrace = append(contextInfo.StackTrace, child)
        return &contextInfo.StackTrace[len(contextInfo.StackTrace)-1]
    }
    if childYangName == "dll-info" {
        for _, c := range contextInfo.DllInfo {
            if contextInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Context_All_CrashInfo_ContextInfo_DllInfo{}
        contextInfo.DllInfo = append(contextInfo.DllInfo, child)
        return &contextInfo.DllInfo[len(contextInfo.DllInfo)-1]
    }
    return nil
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextInfo.StackTrace {
        children[contextInfo.StackTrace[i].GetSegmentPath()] = &contextInfo.StackTrace[i]
    }
    for i := range contextInfo.DllInfo {
        children[contextInfo.DllInfo[i].GetSegmentPath()] = &contextInfo.DllInfo[i]
    }
    return children
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["process-name"] = contextInfo.ProcessName
    leafs["pid"] = contextInfo.Pid
    leafs["tid"] = contextInfo.Tid
    leafs["core-dump-time"] = contextInfo.CoreDumpTime
    leafs["sig-num"] = contextInfo.SigNum
    leafs["sin-err-str"] = contextInfo.SinErrStr
    leafs["sig-send-pid"] = contextInfo.SigSendPid
    leafs["sig-code"] = contextInfo.SigCode
    leafs["sin-info"] = contextInfo.SinInfo
    leafs["core-for-process"] = contextInfo.CoreForProcess
    leafs["registers-info"] = contextInfo.RegistersInfo
    return leafs
}

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetBundleName() string { return "cisco_ios_xr" }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetYangName() string { return "context-info" }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextInfo *Context_All_CrashInfo_ContextInfo) SetParent(parent types.Entity) { contextInfo.parent = parent }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetParent() types.Entity { return contextInfo.parent }

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetParentYangName() string { return "crash-info" }

// Context_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_All_CrashInfo_ContextInfo_StackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetFilter() yfilter.YFilter { return stackTrace.YFilter }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) SetFilter(yf yfilter.YFilter) { stackTrace.YFilter = yf }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetGoName(yname string) string {
    if yname == "stack-trace" { return "StackTrace" }
    return ""
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetSegmentPath() string {
    return "stack-trace"
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stack-trace"] = stackTrace.StackTrace
    return leafs
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetYangName() string { return "stack-trace" }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) SetParent(parent types.Entity) { stackTrace.parent = parent }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetParent() types.Entity { return stackTrace.parent }

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetParentYangName() string { return "context-info" }

// Context_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_All_CrashInfo_ContextInfo_DllInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DLL Path. The type is string.
    Path interface{}

    // text addr. The type is interface{} with range: 0..4294967295.
    TextAddr interface{}

    // text size. The type is interface{} with range: 0..4294967295.
    TextSize interface{}

    // data addr. The type is interface{} with range: 0..4294967295.
    DataAddr interface{}

    // data size. The type is interface{} with range: 0..4294967295.
    DataSize interface{}

    // version. The type is interface{} with range: 0..4294967295.
    Version interface{}
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetFilter() yfilter.YFilter { return dllInfo.YFilter }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) SetFilter(yf yfilter.YFilter) { dllInfo.YFilter = yf }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    if yname == "text-addr" { return "TextAddr" }
    if yname == "text-size" { return "TextSize" }
    if yname == "data-addr" { return "DataAddr" }
    if yname == "data-size" { return "DataSize" }
    if yname == "version" { return "Version" }
    return ""
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetSegmentPath() string {
    return "dll-info"
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path"] = dllInfo.Path
    leafs["text-addr"] = dllInfo.TextAddr
    leafs["text-size"] = dllInfo.TextSize
    leafs["data-addr"] = dllInfo.DataAddr
    leafs["data-size"] = dllInfo.DataSize
    leafs["version"] = dllInfo.Version
    return leafs
}

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetBundleName() string { return "cisco_ios_xr" }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetYangName() string { return "dll-info" }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) SetParent(parent types.Entity) { dllInfo.parent = parent }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetParent() types.Entity { return dllInfo.parent }

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetParentYangName() string { return "context-info" }

// Context_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_All_CrashInfo_CrashPackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetFilter() yfilter.YFilter { return crashPackageInformation.YFilter }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) SetFilter(yf yfilter.YFilter) { crashPackageInformation.YFilter = yf }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "source" { return "Source" }
    return ""
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetSegmentPath() string {
    return "crash-package-information"
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = crashPackageInformation.Name
    leafs["source"] = crashPackageInformation.Source
    return leafs
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetBundleName() string { return "cisco_ios_xr" }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetYangName() string { return "crash-package-information" }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) SetParent(parent types.Entity) { crashPackageInformation.parent = parent }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetParent() types.Entity { return crashPackageInformation.parent }

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetParentYangName() string { return "crash-info" }

