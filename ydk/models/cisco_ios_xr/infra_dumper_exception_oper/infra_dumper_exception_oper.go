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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // exception bag.
    Enter Exception_Enter
}

func (exception *Exception) GetEntityData() *types.CommonEntityData {
    exception.EntityData.YFilter = exception.YFilter
    exception.EntityData.YangName = "exception"
    exception.EntityData.BundleName = "cisco_ios_xr"
    exception.EntityData.ParentYangName = "Cisco-IOS-XR-infra-dumper-exception-oper"
    exception.EntityData.SegmentPath = "Cisco-IOS-XR-infra-dumper-exception-oper:exception"
    exception.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exception.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exception.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exception.EntityData.Children = types.NewOrderedMap()
    exception.EntityData.Children.Append("enter", types.YChild{"Enter", &exception.Enter})
    exception.EntityData.Leafs = types.NewOrderedMap()

    exception.EntityData.YListKeys = []string {}

    return &(exception.EntityData)
}

// Exception_Enter
// exception bag
type Exception_Enter struct {
    EntityData types.CommonEntityData
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

func (enter *Exception_Enter) GetEntityData() *types.CommonEntityData {
    enter.EntityData.YFilter = enter.YFilter
    enter.EntityData.YangName = "enter"
    enter.EntityData.BundleName = "cisco_ios_xr"
    enter.EntityData.ParentYangName = "exception"
    enter.EntityData.SegmentPath = "enter"
    enter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enter.EntityData.Children = types.NewOrderedMap()
    enter.EntityData.Children.Append("display-config1", types.YChild{"DisplayConfig1", &enter.DisplayConfig1})
    enter.EntityData.Children.Append("display-config2", types.YChild{"DisplayConfig2", &enter.DisplayConfig2})
    enter.EntityData.Children.Append("display-config3", types.YChild{"DisplayConfig3", &enter.DisplayConfig3})
    enter.EntityData.Children.Append("display-fall-back-config1", types.YChild{"DisplayFallBackConfig1", &enter.DisplayFallBackConfig1})
    enter.EntityData.Children.Append("display-fall-back-config2", types.YChild{"DisplayFallBackConfig2", &enter.DisplayFallBackConfig2})
    enter.EntityData.Children.Append("display-fall-back-config3", types.YChild{"DisplayFallBackConfig3", &enter.DisplayFallBackConfig3})
    enter.EntityData.Children.Append("kernel-config", types.YChild{"KernelConfig", &enter.KernelConfig})
    enter.EntityData.Children.Append("kernel-route-config", types.YChild{"KernelRouteConfig", &enter.KernelRouteConfig})
    enter.EntityData.Children.Append("core-size", types.YChild{"CoreSize", &enter.CoreSize})
    enter.EntityData.Children.Append("memory-threshold", types.YChild{"MemoryThreshold", &enter.MemoryThreshold})
    enter.EntityData.Children.Append("proc-size", types.YChild{"ProcSize", &enter.ProcSize})
    enter.EntityData.Children.Append("qsize", types.YChild{"Qsize", &enter.Qsize})
    enter.EntityData.Leafs = types.NewOrderedMap()
    enter.EntityData.Leafs.Append("pak-mem", types.YLeaf{"PakMem", enter.PakMem})
    enter.EntityData.Leafs.Append("sparse", types.YLeaf{"Sparse", enter.Sparse})
    enter.EntityData.Leafs.Append("spr-size", types.YLeaf{"SprSize", enter.SprSize})
    enter.EntityData.Leafs.Append("core-verification", types.YLeaf{"CoreVerification", enter.CoreVerification})
    enter.EntityData.Leafs.Append("dump-time-out", types.YLeaf{"DumpTimeOut", enter.DumpTimeOut})

    enter.EntityData.YListKeys = []string {}

    return &(enter.EntityData)
}

// Exception_Enter_DisplayConfig1
// Display Configuration
type Exception_Enter_DisplayConfig1 struct {
    EntityData types.CommonEntityData
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

func (displayConfig1 *Exception_Enter_DisplayConfig1) GetEntityData() *types.CommonEntityData {
    displayConfig1.EntityData.YFilter = displayConfig1.YFilter
    displayConfig1.EntityData.YangName = "display-config1"
    displayConfig1.EntityData.BundleName = "cisco_ios_xr"
    displayConfig1.EntityData.ParentYangName = "enter"
    displayConfig1.EntityData.SegmentPath = "display-config1"
    displayConfig1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayConfig1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayConfig1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayConfig1.EntityData.Children = types.NewOrderedMap()
    displayConfig1.EntityData.Leafs = types.NewOrderedMap()
    displayConfig1.EntityData.Leafs.Append("choice", types.YLeaf{"Choice", displayConfig1.Choice})
    displayConfig1.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayConfig1.Path})
    displayConfig1.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayConfig1.Compress})
    displayConfig1.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayConfig1.FileName})
    displayConfig1.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayConfig1.RangeLow})
    displayConfig1.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayConfig1.RangeHigh})

    displayConfig1.EntityData.YListKeys = []string {}

    return &(displayConfig1.EntityData)
}

// Exception_Enter_DisplayConfig2
// Display Configuration
type Exception_Enter_DisplayConfig2 struct {
    EntityData types.CommonEntityData
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

func (displayConfig2 *Exception_Enter_DisplayConfig2) GetEntityData() *types.CommonEntityData {
    displayConfig2.EntityData.YFilter = displayConfig2.YFilter
    displayConfig2.EntityData.YangName = "display-config2"
    displayConfig2.EntityData.BundleName = "cisco_ios_xr"
    displayConfig2.EntityData.ParentYangName = "enter"
    displayConfig2.EntityData.SegmentPath = "display-config2"
    displayConfig2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayConfig2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayConfig2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayConfig2.EntityData.Children = types.NewOrderedMap()
    displayConfig2.EntityData.Leafs = types.NewOrderedMap()
    displayConfig2.EntityData.Leafs.Append("choice", types.YLeaf{"Choice", displayConfig2.Choice})
    displayConfig2.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayConfig2.Path})
    displayConfig2.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayConfig2.Compress})
    displayConfig2.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayConfig2.FileName})
    displayConfig2.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayConfig2.RangeLow})
    displayConfig2.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayConfig2.RangeHigh})

    displayConfig2.EntityData.YListKeys = []string {}

    return &(displayConfig2.EntityData)
}

// Exception_Enter_DisplayConfig3
// Display Configuration
type Exception_Enter_DisplayConfig3 struct {
    EntityData types.CommonEntityData
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

func (displayConfig3 *Exception_Enter_DisplayConfig3) GetEntityData() *types.CommonEntityData {
    displayConfig3.EntityData.YFilter = displayConfig3.YFilter
    displayConfig3.EntityData.YangName = "display-config3"
    displayConfig3.EntityData.BundleName = "cisco_ios_xr"
    displayConfig3.EntityData.ParentYangName = "enter"
    displayConfig3.EntityData.SegmentPath = "display-config3"
    displayConfig3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayConfig3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayConfig3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayConfig3.EntityData.Children = types.NewOrderedMap()
    displayConfig3.EntityData.Leafs = types.NewOrderedMap()
    displayConfig3.EntityData.Leafs.Append("choice", types.YLeaf{"Choice", displayConfig3.Choice})
    displayConfig3.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayConfig3.Path})
    displayConfig3.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayConfig3.Compress})
    displayConfig3.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayConfig3.FileName})
    displayConfig3.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayConfig3.RangeLow})
    displayConfig3.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayConfig3.RangeHigh})

    displayConfig3.EntityData.YListKeys = []string {}

    return &(displayConfig3.EntityData)
}

// Exception_Enter_DisplayFallBackConfig1
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig1 struct {
    EntityData types.CommonEntityData
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

func (displayFallBackConfig1 *Exception_Enter_DisplayFallBackConfig1) GetEntityData() *types.CommonEntityData {
    displayFallBackConfig1.EntityData.YFilter = displayFallBackConfig1.YFilter
    displayFallBackConfig1.EntityData.YangName = "display-fall-back-config1"
    displayFallBackConfig1.EntityData.BundleName = "cisco_ios_xr"
    displayFallBackConfig1.EntityData.ParentYangName = "enter"
    displayFallBackConfig1.EntityData.SegmentPath = "display-fall-back-config1"
    displayFallBackConfig1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayFallBackConfig1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayFallBackConfig1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayFallBackConfig1.EntityData.Children = types.NewOrderedMap()
    displayFallBackConfig1.EntityData.Leafs = types.NewOrderedMap()
    displayFallBackConfig1.EntityData.Leafs.Append("choice-fall-back", types.YLeaf{"ChoiceFallBack", displayFallBackConfig1.ChoiceFallBack})
    displayFallBackConfig1.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayFallBackConfig1.Path})
    displayFallBackConfig1.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayFallBackConfig1.Compress})
    displayFallBackConfig1.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayFallBackConfig1.FileName})
    displayFallBackConfig1.EntityData.Leafs.Append("boot-device-str", types.YLeaf{"BootDeviceStr", displayFallBackConfig1.BootDeviceStr})
    displayFallBackConfig1.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayFallBackConfig1.RangeLow})
    displayFallBackConfig1.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayFallBackConfig1.RangeHigh})

    displayFallBackConfig1.EntityData.YListKeys = []string {}

    return &(displayFallBackConfig1.EntityData)
}

// Exception_Enter_DisplayFallBackConfig2
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig2 struct {
    EntityData types.CommonEntityData
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

func (displayFallBackConfig2 *Exception_Enter_DisplayFallBackConfig2) GetEntityData() *types.CommonEntityData {
    displayFallBackConfig2.EntityData.YFilter = displayFallBackConfig2.YFilter
    displayFallBackConfig2.EntityData.YangName = "display-fall-back-config2"
    displayFallBackConfig2.EntityData.BundleName = "cisco_ios_xr"
    displayFallBackConfig2.EntityData.ParentYangName = "enter"
    displayFallBackConfig2.EntityData.SegmentPath = "display-fall-back-config2"
    displayFallBackConfig2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayFallBackConfig2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayFallBackConfig2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayFallBackConfig2.EntityData.Children = types.NewOrderedMap()
    displayFallBackConfig2.EntityData.Leafs = types.NewOrderedMap()
    displayFallBackConfig2.EntityData.Leafs.Append("choice-fall-back", types.YLeaf{"ChoiceFallBack", displayFallBackConfig2.ChoiceFallBack})
    displayFallBackConfig2.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayFallBackConfig2.Path})
    displayFallBackConfig2.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayFallBackConfig2.Compress})
    displayFallBackConfig2.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayFallBackConfig2.FileName})
    displayFallBackConfig2.EntityData.Leafs.Append("boot-device-str", types.YLeaf{"BootDeviceStr", displayFallBackConfig2.BootDeviceStr})
    displayFallBackConfig2.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayFallBackConfig2.RangeLow})
    displayFallBackConfig2.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayFallBackConfig2.RangeHigh})

    displayFallBackConfig2.EntityData.YListKeys = []string {}

    return &(displayFallBackConfig2.EntityData)
}

// Exception_Enter_DisplayFallBackConfig3
// Display fallback Configuration
type Exception_Enter_DisplayFallBackConfig3 struct {
    EntityData types.CommonEntityData
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

func (displayFallBackConfig3 *Exception_Enter_DisplayFallBackConfig3) GetEntityData() *types.CommonEntityData {
    displayFallBackConfig3.EntityData.YFilter = displayFallBackConfig3.YFilter
    displayFallBackConfig3.EntityData.YangName = "display-fall-back-config3"
    displayFallBackConfig3.EntityData.BundleName = "cisco_ios_xr"
    displayFallBackConfig3.EntityData.ParentYangName = "enter"
    displayFallBackConfig3.EntityData.SegmentPath = "display-fall-back-config3"
    displayFallBackConfig3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayFallBackConfig3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayFallBackConfig3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayFallBackConfig3.EntityData.Children = types.NewOrderedMap()
    displayFallBackConfig3.EntityData.Leafs = types.NewOrderedMap()
    displayFallBackConfig3.EntityData.Leafs.Append("choice-fall-back", types.YLeaf{"ChoiceFallBack", displayFallBackConfig3.ChoiceFallBack})
    displayFallBackConfig3.EntityData.Leafs.Append("path", types.YLeaf{"Path", displayFallBackConfig3.Path})
    displayFallBackConfig3.EntityData.Leafs.Append("compress", types.YLeaf{"Compress", displayFallBackConfig3.Compress})
    displayFallBackConfig3.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", displayFallBackConfig3.FileName})
    displayFallBackConfig3.EntityData.Leafs.Append("boot-device-str", types.YLeaf{"BootDeviceStr", displayFallBackConfig3.BootDeviceStr})
    displayFallBackConfig3.EntityData.Leafs.Append("range-low", types.YLeaf{"RangeLow", displayFallBackConfig3.RangeLow})
    displayFallBackConfig3.EntityData.Leafs.Append("range-high", types.YLeaf{"RangeHigh", displayFallBackConfig3.RangeHigh})

    displayFallBackConfig3.EntityData.YListKeys = []string {}

    return &(displayFallBackConfig3.EntityData)
}

// Exception_Enter_KernelConfig
// Kernel Configuration
type Exception_Enter_KernelConfig struct {
    EntityData types.CommonEntityData
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

func (kernelConfig *Exception_Enter_KernelConfig) GetEntityData() *types.CommonEntityData {
    kernelConfig.EntityData.YFilter = kernelConfig.YFilter
    kernelConfig.EntityData.YangName = "kernel-config"
    kernelConfig.EntityData.BundleName = "cisco_ios_xr"
    kernelConfig.EntityData.ParentYangName = "enter"
    kernelConfig.EntityData.SegmentPath = "kernel-config"
    kernelConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    kernelConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    kernelConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    kernelConfig.EntityData.Children = types.NewOrderedMap()
    kernelConfig.EntityData.Leafs = types.NewOrderedMap()
    kernelConfig.EntityData.Leafs.Append("choice-fall-back", types.YLeaf{"ChoiceFallBack", kernelConfig.ChoiceFallBack})
    kernelConfig.EntityData.Leafs.Append("path", types.YLeaf{"Path", kernelConfig.Path})
    kernelConfig.EntityData.Leafs.Append("file-name", types.YLeaf{"FileName", kernelConfig.FileName})
    kernelConfig.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", kernelConfig.Memory})

    kernelConfig.EntityData.YListKeys = []string {}

    return &(kernelConfig.EntityData)
}

// Exception_Enter_KernelRouteConfig
// Kernel Route Configuration
type Exception_Enter_KernelRouteConfig struct {
    EntityData types.CommonEntityData
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

func (kernelRouteConfig *Exception_Enter_KernelRouteConfig) GetEntityData() *types.CommonEntityData {
    kernelRouteConfig.EntityData.YFilter = kernelRouteConfig.YFilter
    kernelRouteConfig.EntityData.YangName = "kernel-route-config"
    kernelRouteConfig.EntityData.BundleName = "cisco_ios_xr"
    kernelRouteConfig.EntityData.ParentYangName = "enter"
    kernelRouteConfig.EntityData.SegmentPath = "kernel-route-config"
    kernelRouteConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    kernelRouteConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    kernelRouteConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    kernelRouteConfig.EntityData.Children = types.NewOrderedMap()
    kernelRouteConfig.EntityData.Leafs = types.NewOrderedMap()
    kernelRouteConfig.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", kernelRouteConfig.Slot})
    kernelRouteConfig.EntityData.Leafs.Append("port", types.YLeaf{"Port", kernelRouteConfig.Port})
    kernelRouteConfig.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", kernelRouteConfig.IpAddr})
    kernelRouteConfig.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", kernelRouteConfig.Mask})
    kernelRouteConfig.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", kernelRouteConfig.Destination})
    kernelRouteConfig.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", kernelRouteConfig.NextHop})

    kernelRouteConfig.EntityData.YListKeys = []string {}

    return &(kernelRouteConfig.EntityData)
}

// Exception_Enter_CoreSize
// Core Size
type Exception_Enter_CoreSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (coreSize *Exception_Enter_CoreSize) GetEntityData() *types.CommonEntityData {
    coreSize.EntityData.YFilter = coreSize.YFilter
    coreSize.EntityData.YangName = "core-size"
    coreSize.EntityData.BundleName = "cisco_ios_xr"
    coreSize.EntityData.ParentYangName = "enter"
    coreSize.EntityData.SegmentPath = "core-size"
    coreSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreSize.EntityData.Children = types.NewOrderedMap()
    coreSize.EntityData.Leafs = types.NewOrderedMap()
    coreSize.EntityData.Leafs.Append("string", types.YLeaf{"String", coreSize.String})

    coreSize.EntityData.YListKeys = []string {}

    return &(coreSize.EntityData)
}

// Exception_Enter_MemoryThreshold
// Memory Threshold 
type Exception_Enter_MemoryThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (memoryThreshold *Exception_Enter_MemoryThreshold) GetEntityData() *types.CommonEntityData {
    memoryThreshold.EntityData.YFilter = memoryThreshold.YFilter
    memoryThreshold.EntityData.YangName = "memory-threshold"
    memoryThreshold.EntityData.BundleName = "cisco_ios_xr"
    memoryThreshold.EntityData.ParentYangName = "enter"
    memoryThreshold.EntityData.SegmentPath = "memory-threshold"
    memoryThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryThreshold.EntityData.Children = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs = types.NewOrderedMap()
    memoryThreshold.EntityData.Leafs.Append("string", types.YLeaf{"String", memoryThreshold.String})

    memoryThreshold.EntityData.YListKeys = []string {}

    return &(memoryThreshold.EntityData)
}

// Exception_Enter_ProcSize
// Proc Size
type Exception_Enter_ProcSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (procSize *Exception_Enter_ProcSize) GetEntityData() *types.CommonEntityData {
    procSize.EntityData.YFilter = procSize.YFilter
    procSize.EntityData.YangName = "proc-size"
    procSize.EntityData.BundleName = "cisco_ios_xr"
    procSize.EntityData.ParentYangName = "enter"
    procSize.EntityData.SegmentPath = "proc-size"
    procSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    procSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    procSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    procSize.EntityData.Children = types.NewOrderedMap()
    procSize.EntityData.Leafs = types.NewOrderedMap()
    procSize.EntityData.Leafs.Append("string", types.YLeaf{"String", procSize.String})

    procSize.EntityData.YListKeys = []string {}

    return &(procSize.EntityData)
}

// Exception_Enter_Qsize
// QSIZE 
type Exception_Enter_Qsize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String  . The type is string.
    String interface{}
}

func (qsize *Exception_Enter_Qsize) GetEntityData() *types.CommonEntityData {
    qsize.EntityData.YFilter = qsize.YFilter
    qsize.EntityData.YangName = "qsize"
    qsize.EntityData.BundleName = "cisco_ios_xr"
    qsize.EntityData.ParentYangName = "enter"
    qsize.EntityData.SegmentPath = "qsize"
    qsize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qsize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qsize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qsize.EntityData.Children = types.NewOrderedMap()
    qsize.EntityData.Leafs = types.NewOrderedMap()
    qsize.EntityData.Leafs.Append("string", types.YLeaf{"String", qsize.String})

    qsize.EntityData.YListKeys = []string {}

    return &(qsize.EntityData)
}

