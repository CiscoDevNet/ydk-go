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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context number table.
    ContextNumbers Context_ContextNumbers

    // Core Context location table.
    ContextLocations Context_ContextLocations

    // context bag.
    All Context_All
}

func (context *Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "Cisco-IOS-XR-infra-dumper-oper"
    context.EntityData.SegmentPath = "Cisco-IOS-XR-infra-dumper-oper:context"
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = make(map[string]types.YChild)
    context.EntityData.Children["context-numbers"] = types.YChild{"ContextNumbers", &context.ContextNumbers}
    context.EntityData.Children["context-locations"] = types.YChild{"ContextLocations", &context.ContextLocations}
    context.EntityData.Children["all"] = types.YChild{"All", &context.All}
    context.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(context.EntityData)
}

// Context_ContextNumbers
// Context number table
type Context_ContextNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context number . The type is slice of Context_ContextNumbers_ContextNumber.
    ContextNumber []Context_ContextNumbers_ContextNumber
}

func (contextNumbers *Context_ContextNumbers) GetEntityData() *types.CommonEntityData {
    contextNumbers.EntityData.YFilter = contextNumbers.YFilter
    contextNumbers.EntityData.YangName = "context-numbers"
    contextNumbers.EntityData.BundleName = "cisco_ios_xr"
    contextNumbers.EntityData.ParentYangName = "context"
    contextNumbers.EntityData.SegmentPath = "context-numbers"
    contextNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextNumbers.EntityData.Children = make(map[string]types.YChild)
    contextNumbers.EntityData.Children["context-number"] = types.YChild{"ContextNumber", nil}
    for i := range contextNumbers.ContextNumber {
        contextNumbers.EntityData.Children[types.GetSegmentPath(&contextNumbers.ContextNumber[i])] = types.YChild{"ContextNumber", &contextNumbers.ContextNumber[i]}
    }
    contextNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contextNumbers.EntityData)
}

// Context_ContextNumbers_ContextNumber
// Context number 
type Context_ContextNumbers_ContextNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context for which crash dump info required. The
    // type is interface{} with range: -2147483648..2147483647.
    ContextNum interface{}

    // Core Context location table.
    Locations Context_ContextNumbers_ContextNumber_Locations

    // context bag.
    All Context_ContextNumbers_ContextNumber_All
}

func (contextNumber *Context_ContextNumbers_ContextNumber) GetEntityData() *types.CommonEntityData {
    contextNumber.EntityData.YFilter = contextNumber.YFilter
    contextNumber.EntityData.YangName = "context-number"
    contextNumber.EntityData.BundleName = "cisco_ios_xr"
    contextNumber.EntityData.ParentYangName = "context-numbers"
    contextNumber.EntityData.SegmentPath = "context-number" + "[context-num='" + fmt.Sprintf("%v", contextNumber.ContextNum) + "']"
    contextNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextNumber.EntityData.Children = make(map[string]types.YChild)
    contextNumber.EntityData.Children["locations"] = types.YChild{"Locations", &contextNumber.Locations}
    contextNumber.EntityData.Children["all"] = types.YChild{"All", &contextNumber.All}
    contextNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    contextNumber.EntityData.Leafs["context-num"] = types.YLeaf{"ContextNum", contextNumber.ContextNum}
    return &(contextNumber.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations
// Core Context location table
type Context_ContextNumbers_ContextNumber_Locations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational Context for a particular location. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location.
    Location []Context_ContextNumbers_ContextNumber_Locations_Location
}

func (locations *Context_ContextNumbers_ContextNumber_Locations) GetEntityData() *types.CommonEntityData {
    locations.EntityData.YFilter = locations.YFilter
    locations.EntityData.YangName = "locations"
    locations.EntityData.BundleName = "cisco_ios_xr"
    locations.EntityData.ParentYangName = "context-number"
    locations.EntityData.SegmentPath = "locations"
    locations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locations.EntityData.Children = make(map[string]types.YChild)
    locations.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range locations.Location {
        locations.EntityData.Children[types.GetSegmentPath(&locations.Location[i])] = types.YChild{"Location", &locations.Location[i]}
    }
    locations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(locations.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location
// Operational Context for a particular location
type Context_ContextNumbers_ContextNumber_Locations_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Context info bag.
    Enter Context_ContextNumbers_ContextNumber_Locations_Location_Enter
}

func (location *Context_ContextNumbers_ContextNumber_Locations_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "locations"
    location.EntityData.SegmentPath = "location" + "[node-name='" + fmt.Sprintf("%v", location.NodeName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["enter"] = types.YChild{"Enter", &location.Enter}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", location.NodeName}
    return &(location.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter
// Context info bag
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo.
    CrashInfo []Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo
}

func (enter *Context_ContextNumbers_ContextNumber_Locations_Location_Enter) GetEntityData() *types.CommonEntityData {
    enter.EntityData.YFilter = enter.YFilter
    enter.EntityData.YangName = "enter"
    enter.EntityData.BundleName = "cisco_ios_xr"
    enter.EntityData.ParentYangName = "location"
    enter.EntityData.SegmentPath = "enter"
    enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enter.EntityData.Children = make(map[string]types.YChild)
    enter.EntityData.Children["crash-info"] = types.YChild{"CrashInfo", nil}
    for i := range enter.CrashInfo {
        enter.EntityData.Children[types.GetSegmentPath(&enter.CrashInfo[i])] = types.YChild{"CrashInfo", &enter.CrashInfo[i]}
    }
    enter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enter.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo
// All crash info
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo struct {
    EntityData types.CommonEntityData
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

func (crashInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo) GetEntityData() *types.CommonEntityData {
    crashInfo.EntityData.YFilter = crashInfo.YFilter
    crashInfo.EntityData.YangName = "crash-info"
    crashInfo.EntityData.BundleName = "cisco_ios_xr"
    crashInfo.EntityData.ParentYangName = "enter"
    crashInfo.EntityData.SegmentPath = "crash-info"
    crashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashInfo.EntityData.Children = make(map[string]types.YChild)
    crashInfo.EntityData.Children["context-info"] = types.YChild{"ContextInfo", nil}
    for i := range crashInfo.ContextInfo {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.ContextInfo[i])] = types.YChild{"ContextInfo", &crashInfo.ContextInfo[i]}
    }
    crashInfo.EntityData.Children["crash-package-information"] = types.YChild{"CrashPackageInformation", nil}
    for i := range crashInfo.CrashPackageInformation {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.CrashPackageInformation[i])] = types.YChild{"CrashPackageInformation", &crashInfo.CrashPackageInformation[i]}
    }
    crashInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    crashInfo.EntityData.Leafs["node"] = types.YLeaf{"Node", crashInfo.Node}
    return &(crashInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo
// Context Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo struct {
    EntityData types.CommonEntityData
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

func (contextInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo) GetEntityData() *types.CommonEntityData {
    contextInfo.EntityData.YFilter = contextInfo.YFilter
    contextInfo.EntityData.YangName = "context-info"
    contextInfo.EntityData.BundleName = "cisco_ios_xr"
    contextInfo.EntityData.ParentYangName = "crash-info"
    contextInfo.EntityData.SegmentPath = "context-info"
    contextInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextInfo.EntityData.Children = make(map[string]types.YChild)
    contextInfo.EntityData.Children["stack-trace"] = types.YChild{"StackTrace", nil}
    for i := range contextInfo.StackTrace {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.StackTrace[i])] = types.YChild{"StackTrace", &contextInfo.StackTrace[i]}
    }
    contextInfo.EntityData.Children["dll-info"] = types.YChild{"DllInfo", nil}
    for i := range contextInfo.DllInfo {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.DllInfo[i])] = types.YChild{"DllInfo", &contextInfo.DllInfo[i]}
    }
    contextInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    contextInfo.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", contextInfo.ProcessName}
    contextInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", contextInfo.Pid}
    contextInfo.EntityData.Leafs["tid"] = types.YLeaf{"Tid", contextInfo.Tid}
    contextInfo.EntityData.Leafs["core-dump-time"] = types.YLeaf{"CoreDumpTime", contextInfo.CoreDumpTime}
    contextInfo.EntityData.Leafs["sig-num"] = types.YLeaf{"SigNum", contextInfo.SigNum}
    contextInfo.EntityData.Leafs["sin-err-str"] = types.YLeaf{"SinErrStr", contextInfo.SinErrStr}
    contextInfo.EntityData.Leafs["sig-send-pid"] = types.YLeaf{"SigSendPid", contextInfo.SigSendPid}
    contextInfo.EntityData.Leafs["sig-code"] = types.YLeaf{"SigCode", contextInfo.SigCode}
    contextInfo.EntityData.Leafs["sin-info"] = types.YLeaf{"SinInfo", contextInfo.SinInfo}
    contextInfo.EntityData.Leafs["core-for-process"] = types.YLeaf{"CoreForProcess", contextInfo.CoreForProcess}
    contextInfo.EntityData.Leafs["registers-info"] = types.YLeaf{"RegistersInfo", contextInfo.RegistersInfo}
    return &(contextInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_StackTrace) GetEntityData() *types.CommonEntityData {
    stackTrace.EntityData.YFilter = stackTrace.YFilter
    stackTrace.EntityData.YangName = "stack-trace"
    stackTrace.EntityData.BundleName = "cisco_ios_xr"
    stackTrace.EntityData.ParentYangName = "context-info"
    stackTrace.EntityData.SegmentPath = "stack-trace"
    stackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stackTrace.EntityData.Children = make(map[string]types.YChild)
    stackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    stackTrace.EntityData.Leafs["stack-trace"] = types.YLeaf{"StackTrace", stackTrace.StackTrace}
    return &(stackTrace.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo struct {
    EntityData types.CommonEntityData
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

func (dllInfo *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_ContextInfo_DllInfo) GetEntityData() *types.CommonEntityData {
    dllInfo.EntityData.YFilter = dllInfo.YFilter
    dllInfo.EntityData.YangName = "dll-info"
    dllInfo.EntityData.BundleName = "cisco_ios_xr"
    dllInfo.EntityData.ParentYangName = "context-info"
    dllInfo.EntityData.SegmentPath = "dll-info"
    dllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dllInfo.EntityData.Children = make(map[string]types.YChild)
    dllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    dllInfo.EntityData.Leafs["path"] = types.YLeaf{"Path", dllInfo.Path}
    dllInfo.EntityData.Leafs["text-addr"] = types.YLeaf{"TextAddr", dllInfo.TextAddr}
    dllInfo.EntityData.Leafs["text-size"] = types.YLeaf{"TextSize", dllInfo.TextSize}
    dllInfo.EntityData.Leafs["data-addr"] = types.YLeaf{"DataAddr", dllInfo.DataAddr}
    dllInfo.EntityData.Leafs["data-size"] = types.YLeaf{"DataSize", dllInfo.DataSize}
    dllInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", dllInfo.Version}
    return &(dllInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_Locations_Location_Enter_CrashInfo_CrashPackageInformation) GetEntityData() *types.CommonEntityData {
    crashPackageInformation.EntityData.YFilter = crashPackageInformation.YFilter
    crashPackageInformation.EntityData.YangName = "crash-package-information"
    crashPackageInformation.EntityData.BundleName = "cisco_ios_xr"
    crashPackageInformation.EntityData.ParentYangName = "crash-info"
    crashPackageInformation.EntityData.SegmentPath = "crash-package-information"
    crashPackageInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashPackageInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashPackageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashPackageInformation.EntityData.Children = make(map[string]types.YChild)
    crashPackageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    crashPackageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", crashPackageInformation.Name}
    crashPackageInformation.EntityData.Leafs["source"] = types.YLeaf{"Source", crashPackageInformation.Source}
    return &(crashPackageInformation.EntityData)
}

// Context_ContextNumbers_ContextNumber_All
// context bag
type Context_ContextNumbers_ContextNumber_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextNumbers_ContextNumber_All_CrashInfo.
    CrashInfo []Context_ContextNumbers_ContextNumber_All_CrashInfo
}

func (all *Context_ContextNumbers_ContextNumber_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "context-number"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["crash-info"] = types.YChild{"CrashInfo", nil}
    for i := range all.CrashInfo {
        all.EntityData.Children[types.GetSegmentPath(&all.CrashInfo[i])] = types.YChild{"CrashInfo", &all.CrashInfo[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Context_ContextNumbers_ContextNumber_All_CrashInfo
// All crash info
type Context_ContextNumbers_ContextNumber_All_CrashInfo struct {
    EntityData types.CommonEntityData
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

func (crashInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo) GetEntityData() *types.CommonEntityData {
    crashInfo.EntityData.YFilter = crashInfo.YFilter
    crashInfo.EntityData.YangName = "crash-info"
    crashInfo.EntityData.BundleName = "cisco_ios_xr"
    crashInfo.EntityData.ParentYangName = "all"
    crashInfo.EntityData.SegmentPath = "crash-info"
    crashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashInfo.EntityData.Children = make(map[string]types.YChild)
    crashInfo.EntityData.Children["context-info"] = types.YChild{"ContextInfo", nil}
    for i := range crashInfo.ContextInfo {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.ContextInfo[i])] = types.YChild{"ContextInfo", &crashInfo.ContextInfo[i]}
    }
    crashInfo.EntityData.Children["crash-package-information"] = types.YChild{"CrashPackageInformation", nil}
    for i := range crashInfo.CrashPackageInformation {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.CrashPackageInformation[i])] = types.YChild{"CrashPackageInformation", &crashInfo.CrashPackageInformation[i]}
    }
    crashInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    crashInfo.EntityData.Leafs["node"] = types.YLeaf{"Node", crashInfo.Node}
    return &(crashInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo
// Context Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo struct {
    EntityData types.CommonEntityData
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

func (contextInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo) GetEntityData() *types.CommonEntityData {
    contextInfo.EntityData.YFilter = contextInfo.YFilter
    contextInfo.EntityData.YangName = "context-info"
    contextInfo.EntityData.BundleName = "cisco_ios_xr"
    contextInfo.EntityData.ParentYangName = "crash-info"
    contextInfo.EntityData.SegmentPath = "context-info"
    contextInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextInfo.EntityData.Children = make(map[string]types.YChild)
    contextInfo.EntityData.Children["stack-trace"] = types.YChild{"StackTrace", nil}
    for i := range contextInfo.StackTrace {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.StackTrace[i])] = types.YChild{"StackTrace", &contextInfo.StackTrace[i]}
    }
    contextInfo.EntityData.Children["dll-info"] = types.YChild{"DllInfo", nil}
    for i := range contextInfo.DllInfo {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.DllInfo[i])] = types.YChild{"DllInfo", &contextInfo.DllInfo[i]}
    }
    contextInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    contextInfo.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", contextInfo.ProcessName}
    contextInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", contextInfo.Pid}
    contextInfo.EntityData.Leafs["tid"] = types.YLeaf{"Tid", contextInfo.Tid}
    contextInfo.EntityData.Leafs["core-dump-time"] = types.YLeaf{"CoreDumpTime", contextInfo.CoreDumpTime}
    contextInfo.EntityData.Leafs["sig-num"] = types.YLeaf{"SigNum", contextInfo.SigNum}
    contextInfo.EntityData.Leafs["sin-err-str"] = types.YLeaf{"SinErrStr", contextInfo.SinErrStr}
    contextInfo.EntityData.Leafs["sig-send-pid"] = types.YLeaf{"SigSendPid", contextInfo.SigSendPid}
    contextInfo.EntityData.Leafs["sig-code"] = types.YLeaf{"SigCode", contextInfo.SigCode}
    contextInfo.EntityData.Leafs["sin-info"] = types.YLeaf{"SinInfo", contextInfo.SinInfo}
    contextInfo.EntityData.Leafs["core-for-process"] = types.YLeaf{"CoreForProcess", contextInfo.CoreForProcess}
    contextInfo.EntityData.Leafs["registers-info"] = types.YLeaf{"RegistersInfo", contextInfo.RegistersInfo}
    return &(contextInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_StackTrace) GetEntityData() *types.CommonEntityData {
    stackTrace.EntityData.YFilter = stackTrace.YFilter
    stackTrace.EntityData.YangName = "stack-trace"
    stackTrace.EntityData.BundleName = "cisco_ios_xr"
    stackTrace.EntityData.ParentYangName = "context-info"
    stackTrace.EntityData.SegmentPath = "stack-trace"
    stackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stackTrace.EntityData.Children = make(map[string]types.YChild)
    stackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    stackTrace.EntityData.Leafs["stack-trace"] = types.YLeaf{"StackTrace", stackTrace.StackTrace}
    return &(stackTrace.EntityData)
}

// Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo struct {
    EntityData types.CommonEntityData
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

func (dllInfo *Context_ContextNumbers_ContextNumber_All_CrashInfo_ContextInfo_DllInfo) GetEntityData() *types.CommonEntityData {
    dllInfo.EntityData.YFilter = dllInfo.YFilter
    dllInfo.EntityData.YangName = "dll-info"
    dllInfo.EntityData.BundleName = "cisco_ios_xr"
    dllInfo.EntityData.ParentYangName = "context-info"
    dllInfo.EntityData.SegmentPath = "dll-info"
    dllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dllInfo.EntityData.Children = make(map[string]types.YChild)
    dllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    dllInfo.EntityData.Leafs["path"] = types.YLeaf{"Path", dllInfo.Path}
    dllInfo.EntityData.Leafs["text-addr"] = types.YLeaf{"TextAddr", dllInfo.TextAddr}
    dllInfo.EntityData.Leafs["text-size"] = types.YLeaf{"TextSize", dllInfo.TextSize}
    dllInfo.EntityData.Leafs["data-addr"] = types.YLeaf{"DataAddr", dllInfo.DataAddr}
    dllInfo.EntityData.Leafs["data-size"] = types.YLeaf{"DataSize", dllInfo.DataSize}
    dllInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", dllInfo.Version}
    return &(dllInfo.EntityData)
}

// Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextNumbers_ContextNumber_All_CrashInfo_CrashPackageInformation) GetEntityData() *types.CommonEntityData {
    crashPackageInformation.EntityData.YFilter = crashPackageInformation.YFilter
    crashPackageInformation.EntityData.YangName = "crash-package-information"
    crashPackageInformation.EntityData.BundleName = "cisco_ios_xr"
    crashPackageInformation.EntityData.ParentYangName = "crash-info"
    crashPackageInformation.EntityData.SegmentPath = "crash-package-information"
    crashPackageInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashPackageInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashPackageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashPackageInformation.EntityData.Children = make(map[string]types.YChild)
    crashPackageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    crashPackageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", crashPackageInformation.Name}
    crashPackageInformation.EntityData.Leafs["source"] = types.YLeaf{"Source", crashPackageInformation.Source}
    return &(crashPackageInformation.EntityData)
}

// Context_ContextLocations
// Core Context location table
type Context_ContextLocations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational Context for a particular location. The type is slice of
    // Context_ContextLocations_ContextLocation.
    ContextLocation []Context_ContextLocations_ContextLocation
}

func (contextLocations *Context_ContextLocations) GetEntityData() *types.CommonEntityData {
    contextLocations.EntityData.YFilter = contextLocations.YFilter
    contextLocations.EntityData.YangName = "context-locations"
    contextLocations.EntityData.BundleName = "cisco_ios_xr"
    contextLocations.EntityData.ParentYangName = "context"
    contextLocations.EntityData.SegmentPath = "context-locations"
    contextLocations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextLocations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextLocations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextLocations.EntityData.Children = make(map[string]types.YChild)
    contextLocations.EntityData.Children["context-location"] = types.YChild{"ContextLocation", nil}
    for i := range contextLocations.ContextLocation {
        contextLocations.EntityData.Children[types.GetSegmentPath(&contextLocations.ContextLocation[i])] = types.YChild{"ContextLocation", &contextLocations.ContextLocation[i]}
    }
    contextLocations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contextLocations.EntityData)
}

// Context_ContextLocations_ContextLocation
// Operational Context for a particular location
type Context_ContextLocations_ContextLocation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Context number Table.
    Numbers Context_ContextLocations_ContextLocation_Numbers

    // context bag.
    All Context_ContextLocations_ContextLocation_All
}

func (contextLocation *Context_ContextLocations_ContextLocation) GetEntityData() *types.CommonEntityData {
    contextLocation.EntityData.YFilter = contextLocation.YFilter
    contextLocation.EntityData.YangName = "context-location"
    contextLocation.EntityData.BundleName = "cisco_ios_xr"
    contextLocation.EntityData.ParentYangName = "context-locations"
    contextLocation.EntityData.SegmentPath = "context-location" + "[node-name='" + fmt.Sprintf("%v", contextLocation.NodeName) + "']"
    contextLocation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextLocation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextLocation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextLocation.EntityData.Children = make(map[string]types.YChild)
    contextLocation.EntityData.Children["numbers"] = types.YChild{"Numbers", &contextLocation.Numbers}
    contextLocation.EntityData.Children["all"] = types.YChild{"All", &contextLocation.All}
    contextLocation.EntityData.Leafs = make(map[string]types.YLeaf)
    contextLocation.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", contextLocation.NodeName}
    return &(contextLocation.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers
// Context number Table
type Context_ContextLocations_ContextLocation_Numbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context number. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number.
    Number []Context_ContextLocations_ContextLocation_Numbers_Number
}

func (numbers *Context_ContextLocations_ContextLocation_Numbers) GetEntityData() *types.CommonEntityData {
    numbers.EntityData.YFilter = numbers.YFilter
    numbers.EntityData.YangName = "numbers"
    numbers.EntityData.BundleName = "cisco_ios_xr"
    numbers.EntityData.ParentYangName = "context-location"
    numbers.EntityData.SegmentPath = "numbers"
    numbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    numbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    numbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    numbers.EntityData.Children = make(map[string]types.YChild)
    numbers.EntityData.Children["number"] = types.YChild{"Number", nil}
    for i := range numbers.Number {
        numbers.EntityData.Children[types.GetSegmentPath(&numbers.Number[i])] = types.YChild{"Number", &numbers.Number[i]}
    }
    numbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(numbers.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number
// Context number
type Context_ContextLocations_ContextLocation_Numbers_Number struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context for which crash dump info required. The
    // type is interface{} with range: -2147483648..2147483647.
    ContextNum interface{}

    // Context info bag.
    Enter Context_ContextLocations_ContextLocation_Numbers_Number_Enter
}

func (number *Context_ContextLocations_ContextLocation_Numbers_Number) GetEntityData() *types.CommonEntityData {
    number.EntityData.YFilter = number.YFilter
    number.EntityData.YangName = "number"
    number.EntityData.BundleName = "cisco_ios_xr"
    number.EntityData.ParentYangName = "numbers"
    number.EntityData.SegmentPath = "number" + "[context-num='" + fmt.Sprintf("%v", number.ContextNum) + "']"
    number.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    number.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    number.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    number.EntityData.Children = make(map[string]types.YChild)
    number.EntityData.Children["enter"] = types.YChild{"Enter", &number.Enter}
    number.EntityData.Leafs = make(map[string]types.YLeaf)
    number.EntityData.Leafs["context-num"] = types.YLeaf{"ContextNum", number.ContextNum}
    return &(number.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter
// Context info bag
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo.
    CrashInfo []Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo
}

func (enter *Context_ContextLocations_ContextLocation_Numbers_Number_Enter) GetEntityData() *types.CommonEntityData {
    enter.EntityData.YFilter = enter.YFilter
    enter.EntityData.YangName = "enter"
    enter.EntityData.BundleName = "cisco_ios_xr"
    enter.EntityData.ParentYangName = "number"
    enter.EntityData.SegmentPath = "enter"
    enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enter.EntityData.Children = make(map[string]types.YChild)
    enter.EntityData.Children["crash-info"] = types.YChild{"CrashInfo", nil}
    for i := range enter.CrashInfo {
        enter.EntityData.Children[types.GetSegmentPath(&enter.CrashInfo[i])] = types.YChild{"CrashInfo", &enter.CrashInfo[i]}
    }
    enter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enter.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo
// All crash info
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo struct {
    EntityData types.CommonEntityData
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

func (crashInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo) GetEntityData() *types.CommonEntityData {
    crashInfo.EntityData.YFilter = crashInfo.YFilter
    crashInfo.EntityData.YangName = "crash-info"
    crashInfo.EntityData.BundleName = "cisco_ios_xr"
    crashInfo.EntityData.ParentYangName = "enter"
    crashInfo.EntityData.SegmentPath = "crash-info"
    crashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashInfo.EntityData.Children = make(map[string]types.YChild)
    crashInfo.EntityData.Children["context-info"] = types.YChild{"ContextInfo", nil}
    for i := range crashInfo.ContextInfo {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.ContextInfo[i])] = types.YChild{"ContextInfo", &crashInfo.ContextInfo[i]}
    }
    crashInfo.EntityData.Children["crash-package-information"] = types.YChild{"CrashPackageInformation", nil}
    for i := range crashInfo.CrashPackageInformation {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.CrashPackageInformation[i])] = types.YChild{"CrashPackageInformation", &crashInfo.CrashPackageInformation[i]}
    }
    crashInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    crashInfo.EntityData.Leafs["node"] = types.YLeaf{"Node", crashInfo.Node}
    return &(crashInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo
// Context Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo struct {
    EntityData types.CommonEntityData
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

func (contextInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo) GetEntityData() *types.CommonEntityData {
    contextInfo.EntityData.YFilter = contextInfo.YFilter
    contextInfo.EntityData.YangName = "context-info"
    contextInfo.EntityData.BundleName = "cisco_ios_xr"
    contextInfo.EntityData.ParentYangName = "crash-info"
    contextInfo.EntityData.SegmentPath = "context-info"
    contextInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextInfo.EntityData.Children = make(map[string]types.YChild)
    contextInfo.EntityData.Children["stack-trace"] = types.YChild{"StackTrace", nil}
    for i := range contextInfo.StackTrace {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.StackTrace[i])] = types.YChild{"StackTrace", &contextInfo.StackTrace[i]}
    }
    contextInfo.EntityData.Children["dll-info"] = types.YChild{"DllInfo", nil}
    for i := range contextInfo.DllInfo {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.DllInfo[i])] = types.YChild{"DllInfo", &contextInfo.DllInfo[i]}
    }
    contextInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    contextInfo.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", contextInfo.ProcessName}
    contextInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", contextInfo.Pid}
    contextInfo.EntityData.Leafs["tid"] = types.YLeaf{"Tid", contextInfo.Tid}
    contextInfo.EntityData.Leafs["core-dump-time"] = types.YLeaf{"CoreDumpTime", contextInfo.CoreDumpTime}
    contextInfo.EntityData.Leafs["sig-num"] = types.YLeaf{"SigNum", contextInfo.SigNum}
    contextInfo.EntityData.Leafs["sin-err-str"] = types.YLeaf{"SinErrStr", contextInfo.SinErrStr}
    contextInfo.EntityData.Leafs["sig-send-pid"] = types.YLeaf{"SigSendPid", contextInfo.SigSendPid}
    contextInfo.EntityData.Leafs["sig-code"] = types.YLeaf{"SigCode", contextInfo.SigCode}
    contextInfo.EntityData.Leafs["sin-info"] = types.YLeaf{"SinInfo", contextInfo.SinInfo}
    contextInfo.EntityData.Leafs["core-for-process"] = types.YLeaf{"CoreForProcess", contextInfo.CoreForProcess}
    contextInfo.EntityData.Leafs["registers-info"] = types.YLeaf{"RegistersInfo", contextInfo.RegistersInfo}
    return &(contextInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_StackTrace) GetEntityData() *types.CommonEntityData {
    stackTrace.EntityData.YFilter = stackTrace.YFilter
    stackTrace.EntityData.YangName = "stack-trace"
    stackTrace.EntityData.BundleName = "cisco_ios_xr"
    stackTrace.EntityData.ParentYangName = "context-info"
    stackTrace.EntityData.SegmentPath = "stack-trace"
    stackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stackTrace.EntityData.Children = make(map[string]types.YChild)
    stackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    stackTrace.EntityData.Leafs["stack-trace"] = types.YLeaf{"StackTrace", stackTrace.StackTrace}
    return &(stackTrace.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo struct {
    EntityData types.CommonEntityData
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

func (dllInfo *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_ContextInfo_DllInfo) GetEntityData() *types.CommonEntityData {
    dllInfo.EntityData.YFilter = dllInfo.YFilter
    dllInfo.EntityData.YangName = "dll-info"
    dllInfo.EntityData.BundleName = "cisco_ios_xr"
    dllInfo.EntityData.ParentYangName = "context-info"
    dllInfo.EntityData.SegmentPath = "dll-info"
    dllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dllInfo.EntityData.Children = make(map[string]types.YChild)
    dllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    dllInfo.EntityData.Leafs["path"] = types.YLeaf{"Path", dllInfo.Path}
    dllInfo.EntityData.Leafs["text-addr"] = types.YLeaf{"TextAddr", dllInfo.TextAddr}
    dllInfo.EntityData.Leafs["text-size"] = types.YLeaf{"TextSize", dllInfo.TextSize}
    dllInfo.EntityData.Leafs["data-addr"] = types.YLeaf{"DataAddr", dllInfo.DataAddr}
    dllInfo.EntityData.Leafs["data-size"] = types.YLeaf{"DataSize", dllInfo.DataSize}
    dllInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", dllInfo.Version}
    return &(dllInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_Numbers_Number_Enter_CrashInfo_CrashPackageInformation) GetEntityData() *types.CommonEntityData {
    crashPackageInformation.EntityData.YFilter = crashPackageInformation.YFilter
    crashPackageInformation.EntityData.YangName = "crash-package-information"
    crashPackageInformation.EntityData.BundleName = "cisco_ios_xr"
    crashPackageInformation.EntityData.ParentYangName = "crash-info"
    crashPackageInformation.EntityData.SegmentPath = "crash-package-information"
    crashPackageInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashPackageInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashPackageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashPackageInformation.EntityData.Children = make(map[string]types.YChild)
    crashPackageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    crashPackageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", crashPackageInformation.Name}
    crashPackageInformation.EntityData.Leafs["source"] = types.YLeaf{"Source", crashPackageInformation.Source}
    return &(crashPackageInformation.EntityData)
}

// Context_ContextLocations_ContextLocation_All
// context bag
type Context_ContextLocations_ContextLocation_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All crash info. The type is slice of
    // Context_ContextLocations_ContextLocation_All_CrashInfo.
    CrashInfo []Context_ContextLocations_ContextLocation_All_CrashInfo
}

func (all *Context_ContextLocations_ContextLocation_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "context-location"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["crash-info"] = types.YChild{"CrashInfo", nil}
    for i := range all.CrashInfo {
        all.EntityData.Children[types.GetSegmentPath(&all.CrashInfo[i])] = types.YChild{"CrashInfo", &all.CrashInfo[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Context_ContextLocations_ContextLocation_All_CrashInfo
// All crash info
type Context_ContextLocations_ContextLocation_All_CrashInfo struct {
    EntityData types.CommonEntityData
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

func (crashInfo *Context_ContextLocations_ContextLocation_All_CrashInfo) GetEntityData() *types.CommonEntityData {
    crashInfo.EntityData.YFilter = crashInfo.YFilter
    crashInfo.EntityData.YangName = "crash-info"
    crashInfo.EntityData.BundleName = "cisco_ios_xr"
    crashInfo.EntityData.ParentYangName = "all"
    crashInfo.EntityData.SegmentPath = "crash-info"
    crashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashInfo.EntityData.Children = make(map[string]types.YChild)
    crashInfo.EntityData.Children["context-info"] = types.YChild{"ContextInfo", nil}
    for i := range crashInfo.ContextInfo {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.ContextInfo[i])] = types.YChild{"ContextInfo", &crashInfo.ContextInfo[i]}
    }
    crashInfo.EntityData.Children["crash-package-information"] = types.YChild{"CrashPackageInformation", nil}
    for i := range crashInfo.CrashPackageInformation {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.CrashPackageInformation[i])] = types.YChild{"CrashPackageInformation", &crashInfo.CrashPackageInformation[i]}
    }
    crashInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    crashInfo.EntityData.Leafs["node"] = types.YLeaf{"Node", crashInfo.Node}
    return &(crashInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo
// Context Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo struct {
    EntityData types.CommonEntityData
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

func (contextInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo) GetEntityData() *types.CommonEntityData {
    contextInfo.EntityData.YFilter = contextInfo.YFilter
    contextInfo.EntityData.YangName = "context-info"
    contextInfo.EntityData.BundleName = "cisco_ios_xr"
    contextInfo.EntityData.ParentYangName = "crash-info"
    contextInfo.EntityData.SegmentPath = "context-info"
    contextInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextInfo.EntityData.Children = make(map[string]types.YChild)
    contextInfo.EntityData.Children["stack-trace"] = types.YChild{"StackTrace", nil}
    for i := range contextInfo.StackTrace {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.StackTrace[i])] = types.YChild{"StackTrace", &contextInfo.StackTrace[i]}
    }
    contextInfo.EntityData.Children["dll-info"] = types.YChild{"DllInfo", nil}
    for i := range contextInfo.DllInfo {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.DllInfo[i])] = types.YChild{"DllInfo", &contextInfo.DllInfo[i]}
    }
    contextInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    contextInfo.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", contextInfo.ProcessName}
    contextInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", contextInfo.Pid}
    contextInfo.EntityData.Leafs["tid"] = types.YLeaf{"Tid", contextInfo.Tid}
    contextInfo.EntityData.Leafs["core-dump-time"] = types.YLeaf{"CoreDumpTime", contextInfo.CoreDumpTime}
    contextInfo.EntityData.Leafs["sig-num"] = types.YLeaf{"SigNum", contextInfo.SigNum}
    contextInfo.EntityData.Leafs["sin-err-str"] = types.YLeaf{"SinErrStr", contextInfo.SinErrStr}
    contextInfo.EntityData.Leafs["sig-send-pid"] = types.YLeaf{"SigSendPid", contextInfo.SigSendPid}
    contextInfo.EntityData.Leafs["sig-code"] = types.YLeaf{"SigCode", contextInfo.SigCode}
    contextInfo.EntityData.Leafs["sin-info"] = types.YLeaf{"SinInfo", contextInfo.SinInfo}
    contextInfo.EntityData.Leafs["core-for-process"] = types.YLeaf{"CoreForProcess", contextInfo.CoreForProcess}
    contextInfo.EntityData.Leafs["registers-info"] = types.YLeaf{"RegistersInfo", contextInfo.RegistersInfo}
    return &(contextInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_StackTrace) GetEntityData() *types.CommonEntityData {
    stackTrace.EntityData.YFilter = stackTrace.YFilter
    stackTrace.EntityData.YangName = "stack-trace"
    stackTrace.EntityData.BundleName = "cisco_ios_xr"
    stackTrace.EntityData.ParentYangName = "context-info"
    stackTrace.EntityData.SegmentPath = "stack-trace"
    stackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stackTrace.EntityData.Children = make(map[string]types.YChild)
    stackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    stackTrace.EntityData.Leafs["stack-trace"] = types.YLeaf{"StackTrace", stackTrace.StackTrace}
    return &(stackTrace.EntityData)
}

// Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo struct {
    EntityData types.CommonEntityData
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

func (dllInfo *Context_ContextLocations_ContextLocation_All_CrashInfo_ContextInfo_DllInfo) GetEntityData() *types.CommonEntityData {
    dllInfo.EntityData.YFilter = dllInfo.YFilter
    dllInfo.EntityData.YangName = "dll-info"
    dllInfo.EntityData.BundleName = "cisco_ios_xr"
    dllInfo.EntityData.ParentYangName = "context-info"
    dllInfo.EntityData.SegmentPath = "dll-info"
    dllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dllInfo.EntityData.Children = make(map[string]types.YChild)
    dllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    dllInfo.EntityData.Leafs["path"] = types.YLeaf{"Path", dllInfo.Path}
    dllInfo.EntityData.Leafs["text-addr"] = types.YLeaf{"TextAddr", dllInfo.TextAddr}
    dllInfo.EntityData.Leafs["text-size"] = types.YLeaf{"TextSize", dllInfo.TextSize}
    dllInfo.EntityData.Leafs["data-addr"] = types.YLeaf{"DataAddr", dllInfo.DataAddr}
    dllInfo.EntityData.Leafs["data-size"] = types.YLeaf{"DataSize", dllInfo.DataSize}
    dllInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", dllInfo.Version}
    return &(dllInfo.EntityData)
}

// Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_ContextLocations_ContextLocation_All_CrashInfo_CrashPackageInformation) GetEntityData() *types.CommonEntityData {
    crashPackageInformation.EntityData.YFilter = crashPackageInformation.YFilter
    crashPackageInformation.EntityData.YangName = "crash-package-information"
    crashPackageInformation.EntityData.BundleName = "cisco_ios_xr"
    crashPackageInformation.EntityData.ParentYangName = "crash-info"
    crashPackageInformation.EntityData.SegmentPath = "crash-package-information"
    crashPackageInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashPackageInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashPackageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashPackageInformation.EntityData.Children = make(map[string]types.YChild)
    crashPackageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    crashPackageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", crashPackageInformation.Name}
    crashPackageInformation.EntityData.Leafs["source"] = types.YLeaf{"Source", crashPackageInformation.Source}
    return &(crashPackageInformation.EntityData)
}

// Context_All
// context bag
type Context_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All crash info. The type is slice of Context_All_CrashInfo.
    CrashInfo []Context_All_CrashInfo
}

func (all *Context_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "context"
    all.EntityData.SegmentPath = "all"
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = make(map[string]types.YChild)
    all.EntityData.Children["crash-info"] = types.YChild{"CrashInfo", nil}
    for i := range all.CrashInfo {
        all.EntityData.Children[types.GetSegmentPath(&all.CrashInfo[i])] = types.YChild{"CrashInfo", &all.CrashInfo[i]}
    }
    all.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(all.EntityData)
}

// Context_All_CrashInfo
// All crash info
type Context_All_CrashInfo struct {
    EntityData types.CommonEntityData
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

func (crashInfo *Context_All_CrashInfo) GetEntityData() *types.CommonEntityData {
    crashInfo.EntityData.YFilter = crashInfo.YFilter
    crashInfo.EntityData.YangName = "crash-info"
    crashInfo.EntityData.BundleName = "cisco_ios_xr"
    crashInfo.EntityData.ParentYangName = "all"
    crashInfo.EntityData.SegmentPath = "crash-info"
    crashInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashInfo.EntityData.Children = make(map[string]types.YChild)
    crashInfo.EntityData.Children["context-info"] = types.YChild{"ContextInfo", nil}
    for i := range crashInfo.ContextInfo {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.ContextInfo[i])] = types.YChild{"ContextInfo", &crashInfo.ContextInfo[i]}
    }
    crashInfo.EntityData.Children["crash-package-information"] = types.YChild{"CrashPackageInformation", nil}
    for i := range crashInfo.CrashPackageInformation {
        crashInfo.EntityData.Children[types.GetSegmentPath(&crashInfo.CrashPackageInformation[i])] = types.YChild{"CrashPackageInformation", &crashInfo.CrashPackageInformation[i]}
    }
    crashInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    crashInfo.EntityData.Leafs["node"] = types.YLeaf{"Node", crashInfo.Node}
    return &(crashInfo.EntityData)
}

// Context_All_CrashInfo_ContextInfo
// Context Information
type Context_All_CrashInfo_ContextInfo struct {
    EntityData types.CommonEntityData
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

func (contextInfo *Context_All_CrashInfo_ContextInfo) GetEntityData() *types.CommonEntityData {
    contextInfo.EntityData.YFilter = contextInfo.YFilter
    contextInfo.EntityData.YangName = "context-info"
    contextInfo.EntityData.BundleName = "cisco_ios_xr"
    contextInfo.EntityData.ParentYangName = "crash-info"
    contextInfo.EntityData.SegmentPath = "context-info"
    contextInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextInfo.EntityData.Children = make(map[string]types.YChild)
    contextInfo.EntityData.Children["stack-trace"] = types.YChild{"StackTrace", nil}
    for i := range contextInfo.StackTrace {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.StackTrace[i])] = types.YChild{"StackTrace", &contextInfo.StackTrace[i]}
    }
    contextInfo.EntityData.Children["dll-info"] = types.YChild{"DllInfo", nil}
    for i := range contextInfo.DllInfo {
        contextInfo.EntityData.Children[types.GetSegmentPath(&contextInfo.DllInfo[i])] = types.YChild{"DllInfo", &contextInfo.DllInfo[i]}
    }
    contextInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    contextInfo.EntityData.Leafs["process-name"] = types.YLeaf{"ProcessName", contextInfo.ProcessName}
    contextInfo.EntityData.Leafs["pid"] = types.YLeaf{"Pid", contextInfo.Pid}
    contextInfo.EntityData.Leafs["tid"] = types.YLeaf{"Tid", contextInfo.Tid}
    contextInfo.EntityData.Leafs["core-dump-time"] = types.YLeaf{"CoreDumpTime", contextInfo.CoreDumpTime}
    contextInfo.EntityData.Leafs["sig-num"] = types.YLeaf{"SigNum", contextInfo.SigNum}
    contextInfo.EntityData.Leafs["sin-err-str"] = types.YLeaf{"SinErrStr", contextInfo.SinErrStr}
    contextInfo.EntityData.Leafs["sig-send-pid"] = types.YLeaf{"SigSendPid", contextInfo.SigSendPid}
    contextInfo.EntityData.Leafs["sig-code"] = types.YLeaf{"SigCode", contextInfo.SigCode}
    contextInfo.EntityData.Leafs["sin-info"] = types.YLeaf{"SinInfo", contextInfo.SinInfo}
    contextInfo.EntityData.Leafs["core-for-process"] = types.YLeaf{"CoreForProcess", contextInfo.CoreForProcess}
    contextInfo.EntityData.Leafs["registers-info"] = types.YLeaf{"RegistersInfo", contextInfo.RegistersInfo}
    return &(contextInfo.EntityData)
}

// Context_All_CrashInfo_ContextInfo_StackTrace
// Stack Trace
type Context_All_CrashInfo_ContextInfo_StackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // stack trace. The type is interface{} with range: 0..4294967295.
    StackTrace interface{}
}

func (stackTrace *Context_All_CrashInfo_ContextInfo_StackTrace) GetEntityData() *types.CommonEntityData {
    stackTrace.EntityData.YFilter = stackTrace.YFilter
    stackTrace.EntityData.YangName = "stack-trace"
    stackTrace.EntityData.BundleName = "cisco_ios_xr"
    stackTrace.EntityData.ParentYangName = "context-info"
    stackTrace.EntityData.SegmentPath = "stack-trace"
    stackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stackTrace.EntityData.Children = make(map[string]types.YChild)
    stackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    stackTrace.EntityData.Leafs["stack-trace"] = types.YLeaf{"StackTrace", stackTrace.StackTrace}
    return &(stackTrace.EntityData)
}

// Context_All_CrashInfo_ContextInfo_DllInfo
// DLL Information
type Context_All_CrashInfo_ContextInfo_DllInfo struct {
    EntityData types.CommonEntityData
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

func (dllInfo *Context_All_CrashInfo_ContextInfo_DllInfo) GetEntityData() *types.CommonEntityData {
    dllInfo.EntityData.YFilter = dllInfo.YFilter
    dllInfo.EntityData.YangName = "dll-info"
    dllInfo.EntityData.BundleName = "cisco_ios_xr"
    dllInfo.EntityData.ParentYangName = "context-info"
    dllInfo.EntityData.SegmentPath = "dll-info"
    dllInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dllInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dllInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dllInfo.EntityData.Children = make(map[string]types.YChild)
    dllInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    dllInfo.EntityData.Leafs["path"] = types.YLeaf{"Path", dllInfo.Path}
    dllInfo.EntityData.Leafs["text-addr"] = types.YLeaf{"TextAddr", dllInfo.TextAddr}
    dllInfo.EntityData.Leafs["text-size"] = types.YLeaf{"TextSize", dllInfo.TextSize}
    dllInfo.EntityData.Leafs["data-addr"] = types.YLeaf{"DataAddr", dllInfo.DataAddr}
    dllInfo.EntityData.Leafs["data-size"] = types.YLeaf{"DataSize", dllInfo.DataSize}
    dllInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", dllInfo.Version}
    return &(dllInfo.EntityData)
}

// Context_All_CrashInfo_CrashPackageInformation
// Crash Package Information
type Context_All_CrashInfo_CrashPackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package Source. The type is string.
    Source interface{}
}

func (crashPackageInformation *Context_All_CrashInfo_CrashPackageInformation) GetEntityData() *types.CommonEntityData {
    crashPackageInformation.EntityData.YFilter = crashPackageInformation.YFilter
    crashPackageInformation.EntityData.YangName = "crash-package-information"
    crashPackageInformation.EntityData.BundleName = "cisco_ios_xr"
    crashPackageInformation.EntityData.ParentYangName = "crash-info"
    crashPackageInformation.EntityData.SegmentPath = "crash-package-information"
    crashPackageInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crashPackageInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crashPackageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crashPackageInformation.EntityData.Children = make(map[string]types.YChild)
    crashPackageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    crashPackageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", crashPackageInformation.Name}
    crashPackageInformation.EntityData.Leafs["source"] = types.YLeaf{"Source", crashPackageInformation.Source}
    return &(crashPackageInformation.EntityData)
}

