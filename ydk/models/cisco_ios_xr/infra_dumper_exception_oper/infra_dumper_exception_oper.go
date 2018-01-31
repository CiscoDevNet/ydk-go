// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-dumper-exception package operational data.
// 
// This module contains definitions
// for the following management objects:
//   exception: Show all exception dump configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_dumper_exception_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_dumper_exception_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-dumper-exception-oper exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-dumper-exception-oper:exception", reflect.TypeOf(Exception{}))
}

// Exception
// Show all exception dump configuration
type Exception struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // exception bag.
    Enter Exception_Enter
}

func (exception *Exception) GetFilter() yfilter.YFilter { return exception.YFilter }

func (exception *Exception) SetFilter(yf yfilter.YFilter) { exception.YFilter = yf }

func (exception *Exception) GetGoName(yname string) string {
    if yname == "enter" { return "Enter" }
    return ""
}

func (exception *Exception) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-dumper-exception-oper:exception"
}

func (exception *Exception) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "enter" {
        return &exception.Enter
    }
    return nil
}

func (exception *Exception) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["enter"] = &exception.Enter
    return children
}

func (exception *Exception) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exception *Exception) GetBundleName() string { return "cisco_ios_xr" }

func (exception *Exception) GetYangName() string { return "exception" }

func (exception *Exception) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exception *Exception) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exception *Exception) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exception *Exception) SetParent(parent types.Entity) { exception.parent = parent }

func (exception *Exception) GetParent() types.Entity { return exception.parent }

func (exception *Exception) GetParentYangName() string { return "Cisco-IOS-XR-infra-dumper-exception-oper" }

// Exception_Enter
// exception bag
type Exception_Enter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pak MEM. The type is string.
    PakMem interface{}

    // Sparse . The type is string.
    Sparse interface{}

    // Spr Size. The type is string.
    SprSize interface{}

    // Core Verification . The type is string.
    CoreVerification interface{}

    // Dump Timeout . The type is string.
    DumpTimeOut interface{}

    // Display Configuration.
    DisplayConfig1 Exception_Enter_DisplayConfig1

    // Display Configuration.
    DisplayConfig2 Exception_Enter_DisplayConfig2

    // Display Configuration.
    DisplayConfig3 Exception_Enter_DisplayConfig3

    // Display fallback Configuration.
    DisplayFallBackConfig1 Exception_Enter_DisplayFallBackConfig1

    // Display fallback Configuration.
    DisplayFallBackConfig2 Exception_Enter_DisplayFallBackConfig2

    // Display fallback Configuration.
    DisplayFallBackConfig3 Exception_Enter_DisplayFallBackConfig3

    // Kernel Configuration.
    KernelConfig Exception_Enter_KernelConfig

    // Kernel Route Configuration.
    KernelRouteConfig Exception_Enter_KernelRouteConfig

    // Core Size.
    CoreSize Exception_Enter_CoreSize

    // Memory Threshold .
    MemoryThreshold Exception_Enter_MemoryThreshold

    // Proc Size.
    ProcSize Exception_Enter_ProcSize

    // QSIZE .
    Qsize Exception_Enter_Qsize
}

func (enter *Exception_Enter) GetFilter() yfilter.YFilter { return enter.YFilter }

func (enter *Exception_Enter) SetFilter(yf yfilter.YFilter) { enter.YFilter = yf }

func (enter *Exception_Enter) GetGoName(yname string) string {
    if yname == "pak-mem" { return "PakMem" }
    if yname == "sparse" { return "Sparse" }
    if yname == "spr-size" { return "SprSize" }
    if yname == "core-verification" { return "CoreVerification" }
    if yname == "dump-time-out" { return "DumpTimeOut" }
    if yname == "display-config1" { return "DisplayConfig1" }
    if yname == "display-config2" { return "DisplayConfig2" }
    if yname == "display-config3" { return "DisplayConfig3" }
    if yname == "display-fall-back-config1" { return "DisplayFallBackConfig1" }
    if yname == "display-fall-back-config2" { return "DisplayFallBackConfig2" }
    if yname == "display-fall-back-config3" { return "DisplayFallBackConfig3" }
    if yname == "kernel-config" { return "KernelConfig" }
    if yname == "kernel-route-config" { return "KernelRouteConfig" }
    if yname == "core-size" { return "CoreSize" }
    if yname == "memory-threshold" { return "MemoryThreshold" }
    if yname == "proc-size" { return "ProcSize" }
    if yname == "qsize" { return "Qsize" }
    return ""
}

func (enter *Exception_Enter) GetSegmentPath() string {
    return "enter"
}

func (enter *Exception_Enter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "display-config1" {
        return &enter.DisplayConfig1
    }
    if childYangName == "display-config2" {
        return &enter.DisplayConfig2
    }
    if childYangName == "display-config3" {
        return &enter.DisplayConfig3
    }
    if childYangName == "display-fall-back-config1" {
        return &enter.DisplayFallBackConfig1
    }
    if childYangName == "display-fall-back-config2" {
        return &enter.DisplayFallBackConfig2
    }
    if childYangName == "display-fall-back-config3" {
        return &enter.DisplayFallBackConfig3
    }
    if childYangName == "kernel-config" {
        return &enter.KernelConfig
    }
    if childYangName == "kernel-route-config" {
        return &enter.KernelRouteConfig
    }
    if childYangName == "core-size" {
        return &enter.CoreSize
    }
    if childYangName == "memory-threshold" {
        return &enter.MemoryThreshold
    }
    if childYangName == "proc-size" {
        return &enter.ProcSize
    }
    if childYangName == "qsize" {
        return &enter.Qsize
    }
    return nil
}

func (enter *Exception_Enter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["display-config1"] = &enter.DisplayConfig1
    children["display-config2"] = &enter.DisplayConfig2
    children["display-config3"] = &enter.DisplayConfig3
    children["display-fall-back-config1"] = &enter.DisplayFallBackConfig1
    children["display-fall-back-config2"] = &enter.DisplayFallBackConfig2
    children["display-fall-back-config3"] = &enter.DisplayFallBackConfig3
    children["kernel-config"] = &enter.KernelConfig
    children["kernel-route-config"] = &enter.KernelRouteConfig
    children["core-size"] = &enter.CoreSize
    children["memory-threshold"] = &enter.MemoryThreshold
    children["proc-size"] = &enter.ProcSize
    children["qsize"] = &enter.Qsize
    return children
}

func (enter *Exception_Enter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pak-mem"] = enter.PakMem
    leafs["sparse"] = enter.Sparse
    leafs["spr-size"] = enter.SprSize
    leafs["core-verification"] = enter.CoreVerification
    leafs["dump-time-out"] = enter.DumpTimeOut
    return leafs
}

func (enter *Exception_Enter) GetBundleName() string { return "cisco_ios_xr" }

func (enter *Exception_Enter) GetYangName() string { return "enter" }

func (enter *Exception_Enter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enter *Exception_Enter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enter *Exception_Enter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enter *Exception_Enter) SetParent(parent types.Entity) { enter.parent = parent }

func (enter *Exception_Enter) GetParent() types.Entity { return enter.parent }

func (enter *Exception_Enter) GetParentYangName() string { return "exception" }

// Exception_Enter_DisplayConfig1
// Display Configuration
type Exception_Enter_DisplayConfig1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    Choice interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetFilter() yfilter.YFilter { return displayConfig1.YFilter }

func (displayConfig1 *Exception_Enter_DisplayConfig1) SetFilter(yf yfilter.YFilter) { displayConfig1.YFilter = yf }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetGoName(yname string) string {
    if yname == "choice" { return "Choice" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetSegmentPath() string {
    return "display-config1"
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice"] = displayConfig1.Choice
    leafs["path"] = displayConfig1.Path
    leafs["compress"] = displayConfig1.Compress
    leafs["file-name"] = displayConfig1.FileName
    leafs["range-low"] = displayConfig1.RangeLow
    leafs["range-high"] = displayConfig1.RangeHigh
    return leafs
}

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetBundleName() string { return "cisco_ios_xr" }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetYangName() string { return "display-config1" }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayConfig1 *Exception_Enter_DisplayConfig1) SetParent(parent types.Entity) { displayConfig1.parent = parent }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetParent() types.Entity { return displayConfig1.parent }

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetParentYangName() string { return "enter" }

// Exception_Enter_DisplayConfig2
// Display Configuration
type Exception_Enter_DisplayConfig2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    Choice interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetFilter() yfilter.YFilter { return displayConfig2.YFilter }

func (displayConfig2 *Exception_Enter_DisplayConfig2) SetFilter(yf yfilter.YFilter) { displayConfig2.YFilter = yf }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetGoName(yname string) string {
    if yname == "choice" { return "Choice" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetSegmentPath() string {
    return "display-config2"
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice"] = displayConfig2.Choice
    leafs["path"] = displayConfig2.Path
    leafs["compress"] = displayConfig2.Compress
    leafs["file-name"] = displayConfig2.FileName
    leafs["range-low"] = displayConfig2.RangeLow
    leafs["range-high"] = displayConfig2.RangeHigh
    return leafs
}

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetBundleName() string { return "cisco_ios_xr" }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetYangName() string { return "display-config2" }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayConfig2 *Exception_Enter_DisplayConfig2) SetParent(parent types.Entity) { displayConfig2.parent = parent }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetParent() types.Entity { return displayConfig2.parent }

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetParentYangName() string { return "enter" }

// Exception_Enter_DisplayConfig3
// Display Configuration
type Exception_Enter_DisplayConfig3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    Choice interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetFilter() yfilter.YFilter { return displayConfig3.YFilter }

func (displayConfig3 *Exception_Enter_DisplayConfig3) SetFilter(yf yfilter.YFilter) { displayConfig3.YFilter = yf }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetGoName(yname string) string {
    if yname == "choice" { return "Choice" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetSegmentPath() string {
    return "display-config3"
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice"] = displayConfig3.Choice
    leafs["path"] = displayConfig3.Path
    leafs["compress"] = displayConfig3.Compress
    leafs["file-name"] = displayConfig3.FileName
    leafs["range-low"] = displayConfig3.RangeLow
    leafs["range-high"] = displayConfig3.RangeHigh
    return leafs
}

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetBundleName() string { return "cisco_ios_xr" }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetYangName() string { return "display-config3" }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayConfig3 *Exception_Enter_DisplayConfig3) SetParent(parent types.Entity) { displayConfig3.parent = parent }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetParent() types.Entity { return displayConfig3.parent }

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetParentYangName() string { return "enter" }

// Exception_Enter_DisplayFallBackConfig1
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    ChoiceFallBack interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Boot Device String . The type is string.
    BootDeviceStr interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetFilter() yfilter.YFilter { return displayFallBackConfig1.YFilter }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) SetFilter(yf yfilter.YFilter) { displayFallBackConfig1.YFilter = yf }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetGoName(yname string) string {
    if yname == "choice-fall-back" { return "ChoiceFallBack" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "boot-device-str" { return "BootDeviceStr" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetSegmentPath() string {
    return "display-fall-back-config1"
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice-fall-back"] = displayFallBackConfig1.ChoiceFallBack
    leafs["path"] = displayFallBackConfig1.Path
    leafs["compress"] = displayFallBackConfig1.Compress
    leafs["file-name"] = displayFallBackConfig1.FileName
    leafs["boot-device-str"] = displayFallBackConfig1.BootDeviceStr
    leafs["range-low"] = displayFallBackConfig1.RangeLow
    leafs["range-high"] = displayFallBackConfig1.RangeHigh
    return leafs
}

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetBundleName() string { return "cisco_ios_xr" }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetYangName() string { return "display-fall-back-config1" }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) SetParent(parent types.Entity) { displayFallBackConfig1.parent = parent }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetParent() types.Entity { return displayFallBackConfig1.parent }

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetParentYangName() string { return "enter" }

// Exception_Enter_DisplayFallBackConfig2
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    ChoiceFallBack interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Boot Device String . The type is string.
    BootDeviceStr interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetFilter() yfilter.YFilter { return displayFallBackConfig2.YFilter }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) SetFilter(yf yfilter.YFilter) { displayFallBackConfig2.YFilter = yf }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetGoName(yname string) string {
    if yname == "choice-fall-back" { return "ChoiceFallBack" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "boot-device-str" { return "BootDeviceStr" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetSegmentPath() string {
    return "display-fall-back-config2"
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice-fall-back"] = displayFallBackConfig2.ChoiceFallBack
    leafs["path"] = displayFallBackConfig2.Path
    leafs["compress"] = displayFallBackConfig2.Compress
    leafs["file-name"] = displayFallBackConfig2.FileName
    leafs["boot-device-str"] = displayFallBackConfig2.BootDeviceStr
    leafs["range-low"] = displayFallBackConfig2.RangeLow
    leafs["range-high"] = displayFallBackConfig2.RangeHigh
    return leafs
}

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetBundleName() string { return "cisco_ios_xr" }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetYangName() string { return "display-fall-back-config2" }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) SetParent(parent types.Entity) { displayFallBackConfig2.parent = parent }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetParent() types.Entity { return displayFallBackConfig2.parent }

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetParentYangName() string { return "enter" }

// Exception_Enter_DisplayFallBackConfig3
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    ChoiceFallBack interface{}

    // Path . The type is string.
    Path interface{}

    // Compress on/off . The type is string.
    Compress interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Boot Device String . The type is string.
    BootDeviceStr interface{}

    // Range Low . The type is interface{} with range: 0..4294967295.
    RangeLow interface{}

    // Range High . The type is interface{} with range: 0..4294967295.
    RangeHigh interface{}
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetFilter() yfilter.YFilter { return displayFallBackConfig3.YFilter }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) SetFilter(yf yfilter.YFilter) { displayFallBackConfig3.YFilter = yf }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetGoName(yname string) string {
    if yname == "choice-fall-back" { return "ChoiceFallBack" }
    if yname == "path" { return "Path" }
    if yname == "compress" { return "Compress" }
    if yname == "file-name" { return "FileName" }
    if yname == "boot-device-str" { return "BootDeviceStr" }
    if yname == "range-low" { return "RangeLow" }
    if yname == "range-high" { return "RangeHigh" }
    return ""
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetSegmentPath() string {
    return "display-fall-back-config3"
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice-fall-back"] = displayFallBackConfig3.ChoiceFallBack
    leafs["path"] = displayFallBackConfig3.Path
    leafs["compress"] = displayFallBackConfig3.Compress
    leafs["file-name"] = displayFallBackConfig3.FileName
    leafs["boot-device-str"] = displayFallBackConfig3.BootDeviceStr
    leafs["range-low"] = displayFallBackConfig3.RangeLow
    leafs["range-high"] = displayFallBackConfig3.RangeHigh
    return leafs
}

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetBundleName() string { return "cisco_ios_xr" }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetYangName() string { return "display-fall-back-config3" }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) SetParent(parent types.Entity) { displayFallBackConfig3.parent = parent }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetParent() types.Entity { return displayFallBackConfig3.parent }

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetParentYangName() string { return "enter" }

// Exception_Enter_KernelConfig
// Kernel Configuration
type Exception_Enter_KernelConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choice . The type is string.
    ChoiceFallBack interface{}

    // Path . The type is string.
    Path interface{}

    // Name of the File . The type is string.
    FileName interface{}

    // Memory . The type is string.
    Memory interface{}
}

func (kernelConfig *Exception_Enter_KernelConfig) GetFilter() yfilter.YFilter { return kernelConfig.YFilter }

func (kernelConfig *Exception_Enter_KernelConfig) SetFilter(yf yfilter.YFilter) { kernelConfig.YFilter = yf }

func (kernelConfig *Exception_Enter_KernelConfig) GetGoName(yname string) string {
    if yname == "choice-fall-back" { return "ChoiceFallBack" }
    if yname == "path" { return "Path" }
    if yname == "file-name" { return "FileName" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (kernelConfig *Exception_Enter_KernelConfig) GetSegmentPath() string {
    return "kernel-config"
}

func (kernelConfig *Exception_Enter_KernelConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (kernelConfig *Exception_Enter_KernelConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (kernelConfig *Exception_Enter_KernelConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["choice-fall-back"] = kernelConfig.ChoiceFallBack
    leafs["path"] = kernelConfig.Path
    leafs["file-name"] = kernelConfig.FileName
    leafs["memory"] = kernelConfig.Memory
    return leafs
}

func (kernelConfig *Exception_Enter_KernelConfig) GetBundleName() string { return "cisco_ios_xr" }

func (kernelConfig *Exception_Enter_KernelConfig) GetYangName() string { return "kernel-config" }

func (kernelConfig *Exception_Enter_KernelConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (kernelConfig *Exception_Enter_KernelConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (kernelConfig *Exception_Enter_KernelConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (kernelConfig *Exception_Enter_KernelConfig) SetParent(parent types.Entity) { kernelConfig.parent = parent }

func (kernelConfig *Exception_Enter_KernelConfig) GetParent() types.Entity { return kernelConfig.parent }

func (kernelConfig *Exception_Enter_KernelConfig) GetParentYangName() string { return "enter" }

// Exception_Enter_KernelRouteConfig
// Kernel Route Configuration
type Exception_Enter_KernelRouteConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot . The type is interface{} with range: 0..4294967295.
    Slot interface{}

    // Port . The type is interface{} with range: 0..4294967295.
    Port interface{}

    // IP Address. The type is string.
    IpAddr interface{}

    // Mask. The type is string.
    Mask interface{}

    // Destination . The type is string.
    Destination interface{}

    // NextHop . The type is string.
    NextHop interface{}
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetFilter() yfilter.YFilter { return kernelRouteConfig.YFilter }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) SetFilter(yf yfilter.YFilter) { kernelRouteConfig.YFilter = yf }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "port" { return "Port" }
    if yname == "ip-addr" { return "IpAddr" }
    if yname == "mask" { return "Mask" }
    if yname == "destination" { return "Destination" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetSegmentPath() string {
    return "kernel-route-config"
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = kernelRouteConfig.Slot
    leafs["port"] = kernelRouteConfig.Port
    leafs["ip-addr"] = kernelRouteConfig.IpAddr
    leafs["mask"] = kernelRouteConfig.Mask
    leafs["destination"] = kernelRouteConfig.Destination
    leafs["next-hop"] = kernelRouteConfig.NextHop
    return leafs
}

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetBundleName() string { return "cisco_ios_xr" }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetYangName() string { return "kernel-route-config" }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) SetParent(parent types.Entity) { kernelRouteConfig.parent = parent }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetParent() types.Entity { return kernelRouteConfig.parent }

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetParentYangName() string { return "enter" }

// Exception_Enter_CoreSize
// Core Size
type Exception_Enter_CoreSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (coreSize *Exception_Enter_CoreSize) GetFilter() yfilter.YFilter { return coreSize.YFilter }

func (coreSize *Exception_Enter_CoreSize) SetFilter(yf yfilter.YFilter) { coreSize.YFilter = yf }

func (coreSize *Exception_Enter_CoreSize) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    return ""
}

func (coreSize *Exception_Enter_CoreSize) GetSegmentPath() string {
    return "core-size"
}

func (coreSize *Exception_Enter_CoreSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coreSize *Exception_Enter_CoreSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coreSize *Exception_Enter_CoreSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = coreSize.String
    return leafs
}

func (coreSize *Exception_Enter_CoreSize) GetBundleName() string { return "cisco_ios_xr" }

func (coreSize *Exception_Enter_CoreSize) GetYangName() string { return "core-size" }

func (coreSize *Exception_Enter_CoreSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coreSize *Exception_Enter_CoreSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coreSize *Exception_Enter_CoreSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coreSize *Exception_Enter_CoreSize) SetParent(parent types.Entity) { coreSize.parent = parent }

func (coreSize *Exception_Enter_CoreSize) GetParent() types.Entity { return coreSize.parent }

func (coreSize *Exception_Enter_CoreSize) GetParentYangName() string { return "enter" }

// Exception_Enter_MemoryThreshold
// Memory Threshold 
type Exception_Enter_MemoryThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetFilter() yfilter.YFilter { return memoryThreshold.YFilter }

func (memoryThreshold *Exception_Enter_MemoryThreshold) SetFilter(yf yfilter.YFilter) { memoryThreshold.YFilter = yf }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    return ""
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetSegmentPath() string {
    return "memory-threshold"
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = memoryThreshold.String
    return leafs
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetYangName() string { return "memory-threshold" }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryThreshold *Exception_Enter_MemoryThreshold) SetParent(parent types.Entity) { memoryThreshold.parent = parent }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetParent() types.Entity { return memoryThreshold.parent }

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetParentYangName() string { return "enter" }

// Exception_Enter_ProcSize
// Proc Size
type Exception_Enter_ProcSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (procSize *Exception_Enter_ProcSize) GetFilter() yfilter.YFilter { return procSize.YFilter }

func (procSize *Exception_Enter_ProcSize) SetFilter(yf yfilter.YFilter) { procSize.YFilter = yf }

func (procSize *Exception_Enter_ProcSize) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    return ""
}

func (procSize *Exception_Enter_ProcSize) GetSegmentPath() string {
    return "proc-size"
}

func (procSize *Exception_Enter_ProcSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (procSize *Exception_Enter_ProcSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (procSize *Exception_Enter_ProcSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = procSize.String
    return leafs
}

func (procSize *Exception_Enter_ProcSize) GetBundleName() string { return "cisco_ios_xr" }

func (procSize *Exception_Enter_ProcSize) GetYangName() string { return "proc-size" }

func (procSize *Exception_Enter_ProcSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (procSize *Exception_Enter_ProcSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (procSize *Exception_Enter_ProcSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (procSize *Exception_Enter_ProcSize) SetParent(parent types.Entity) { procSize.parent = parent }

func (procSize *Exception_Enter_ProcSize) GetParent() types.Entity { return procSize.parent }

func (procSize *Exception_Enter_ProcSize) GetParentYangName() string { return "enter" }

// Exception_Enter_Qsize
// QSIZE 
type Exception_Enter_Qsize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (qsize *Exception_Enter_Qsize) GetFilter() yfilter.YFilter { return qsize.YFilter }

func (qsize *Exception_Enter_Qsize) SetFilter(yf yfilter.YFilter) { qsize.YFilter = yf }

func (qsize *Exception_Enter_Qsize) GetGoName(yname string) string {
    if yname == "string" { return "String" }
    return ""
}

func (qsize *Exception_Enter_Qsize) GetSegmentPath() string {
    return "qsize"
}

func (qsize *Exception_Enter_Qsize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qsize *Exception_Enter_Qsize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qsize *Exception_Enter_Qsize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["string"] = qsize.String
    return leafs
}

func (qsize *Exception_Enter_Qsize) GetBundleName() string { return "cisco_ios_xr" }

func (qsize *Exception_Enter_Qsize) GetYangName() string { return "qsize" }

func (qsize *Exception_Enter_Qsize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qsize *Exception_Enter_Qsize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qsize *Exception_Enter_Qsize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qsize *Exception_Enter_Qsize) SetParent(parent types.Entity) { qsize.parent = parent }

func (qsize *Exception_Enter_Qsize) GetParent() types.Entity { return qsize.parent }

func (qsize *Exception_Enter_Qsize) GetParentYangName() string { return "enter" }

