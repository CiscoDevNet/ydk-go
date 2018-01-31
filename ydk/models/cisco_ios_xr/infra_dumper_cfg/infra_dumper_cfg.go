// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-dumper package configuration.
// 
// This module contains definitions
// for the following management objects:
//   exception: Core dump configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_dumper_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_dumper_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-dumper-cfg exception}", reflect.TypeOf(Exception{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-dumper-cfg:exception", reflect.TypeOf(Exception{}))
}

// Context represents Context
type Context string

const (
    // context not configured
    Context_default_ Context = "default"

    // Dump context info only: Overrides other options
    // except for 'no-core'
    Context_context Context = "context"
)

// Skipcpuinfo represents Skipcpuinfo
type Skipcpuinfo string

const (
    // skip-cpu-info not configured
    Skipcpuinfo_default_ Skipcpuinfo = "default"

    // Skip CPU usage snapshot: for time critical
    // processes
    Skipcpuinfo_skip_cpu_info Skipcpuinfo = "skip-cpu-info"
)

// Packetmemory represents Packetmemory
type Packetmemory string

const (
    // packet memory not configured
    Packetmemory_default_ Packetmemory = "default"

    // Dump packet memory of the target process
    Packetmemory_packet_memory Packetmemory = "packet-memory"
)

// Sparse represents Sparse
type Sparse string

const (
    // sparse not configured
    Sparse_default_ Sparse = "default"

    // Dump memory relevant to stack trace only: for
    // time critical processes
    Sparse_sparse Sparse = "sparse"
)

// Mainmemory represents Mainmemory
type Mainmemory string

const (
    // main memory not configured
    Mainmemory_default_ Mainmemory = "default"

    // Dump main memory of the target process
    Mainmemory_main_memory Mainmemory = "main-memory"
)

// Nocore represents Nocore
type Nocore string

const (
    // no core not configured
    Nocore_default_ Nocore = "default"

    // Disable core dump for the target process:
    // Overrides other options
    Nocore_no_core Nocore = "no-core"
)

// Copy represents Copy
type Copy string

const (
    // copy not configured
    Copy_default_ Copy = "default"

    // Dump to local memory: for time critical
    // processes
    Copy_copy Copy = "copy"
)

// Sharedmemory represents Sharedmemory
type Sharedmemory string

const (
    // shared memory not configured
    Sharedmemory_default_ Sharedmemory = "default"

    // Dump shared memory of the target process
    Sharedmemory_shared_memory Sharedmemory = "shared-memory"
)

// Exception
// Core dump configuration commands
type Exception struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify 'true' to enable sparse core dump, 'false' to disable sparse core
    // dump. The type is bool.
    Sparse interface{}

    // Disable core file verification. The type is bool.
    CoreVerification interface{}

    // Only print out stack trace and create no core file beyond this size. The
    // type is interface{} with range: 1..4095.
    CoreSize interface{}

    // Enable kernel debugger. The type is interface{}.
    KernelDebugger interface{}

    // Specify 'true' to dump packet memory for all process, 'false' to disable
    // dump of packet memory. The type is bool.
    PacketMemory interface{}

    // Switch to sparse core dump at this size. The type is interface{} with
    // range: 1..4095.
    SparseSize interface{}

    // Give up core dump if specified free memory can not be secured. The type is
    // interface{} with range: 3..40. Units are percentage.
    MemoryThreshold interface{}

    // Preference of the dump location.
    Choice1 Exception_Choice1

    // Preference of the dump location.
    Choice3 Exception_Choice3

    // Specify per process configuration.
    ProcessNames Exception_ProcessNames

    // Preference of the dump location.
    Choice2 Exception_Choice2
}

func (exception *Exception) GetFilter() yfilter.YFilter { return exception.YFilter }

func (exception *Exception) SetFilter(yf yfilter.YFilter) { exception.YFilter = yf }

func (exception *Exception) GetGoName(yname string) string {
    if yname == "sparse" { return "Sparse" }
    if yname == "core-verification" { return "CoreVerification" }
    if yname == "core-size" { return "CoreSize" }
    if yname == "kernel-debugger" { return "KernelDebugger" }
    if yname == "packet-memory" { return "PacketMemory" }
    if yname == "sparse-size" { return "SparseSize" }
    if yname == "memory-threshold" { return "MemoryThreshold" }
    if yname == "choice1" { return "Choice1" }
    if yname == "choice3" { return "Choice3" }
    if yname == "process-names" { return "ProcessNames" }
    if yname == "choice2" { return "Choice2" }
    return ""
}

func (exception *Exception) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-dumper-cfg:exception"
}

func (exception *Exception) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "choice1" {
        return &exception.Choice1
    }
    if childYangName == "choice3" {
        return &exception.Choice3
    }
    if childYangName == "process-names" {
        return &exception.ProcessNames
    }
    if childYangName == "choice2" {
        return &exception.Choice2
    }
    return nil
}

func (exception *Exception) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["choice1"] = &exception.Choice1
    children["choice3"] = &exception.Choice3
    children["process-names"] = &exception.ProcessNames
    children["choice2"] = &exception.Choice2
    return children
}

func (exception *Exception) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sparse"] = exception.Sparse
    leafs["core-verification"] = exception.CoreVerification
    leafs["core-size"] = exception.CoreSize
    leafs["kernel-debugger"] = exception.KernelDebugger
    leafs["packet-memory"] = exception.PacketMemory
    leafs["sparse-size"] = exception.SparseSize
    leafs["memory-threshold"] = exception.MemoryThreshold
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

func (exception *Exception) GetParentYangName() string { return "Cisco-IOS-XR-infra-dumper-cfg" }

// Exception_Choice1
// Preference of the dump location
type Exception_Choice1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice1 *Exception_Choice1) GetFilter() yfilter.YFilter { return choice1.YFilter }

func (choice1 *Exception_Choice1) SetFilter(yf yfilter.YFilter) { choice1.YFilter = yf }

func (choice1 *Exception_Choice1) GetGoName(yname string) string {
    if yname == "compress" { return "Compress" }
    if yname == "lower-limit" { return "LowerLimit" }
    if yname == "higher-limit" { return "HigherLimit" }
    if yname == "file-path" { return "FilePath" }
    if yname == "filename" { return "Filename" }
    return ""
}

func (choice1 *Exception_Choice1) GetSegmentPath() string {
    return "choice1"
}

func (choice1 *Exception_Choice1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (choice1 *Exception_Choice1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (choice1 *Exception_Choice1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["compress"] = choice1.Compress
    leafs["lower-limit"] = choice1.LowerLimit
    leafs["higher-limit"] = choice1.HigherLimit
    leafs["file-path"] = choice1.FilePath
    leafs["filename"] = choice1.Filename
    return leafs
}

func (choice1 *Exception_Choice1) GetBundleName() string { return "cisco_ios_xr" }

func (choice1 *Exception_Choice1) GetYangName() string { return "choice1" }

func (choice1 *Exception_Choice1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (choice1 *Exception_Choice1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (choice1 *Exception_Choice1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (choice1 *Exception_Choice1) SetParent(parent types.Entity) { choice1.parent = parent }

func (choice1 *Exception_Choice1) GetParent() types.Entity { return choice1.parent }

func (choice1 *Exception_Choice1) GetParentYangName() string { return "exception" }

// Exception_Choice3
// Preference of the dump location
type Exception_Choice3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice3 *Exception_Choice3) GetFilter() yfilter.YFilter { return choice3.YFilter }

func (choice3 *Exception_Choice3) SetFilter(yf yfilter.YFilter) { choice3.YFilter = yf }

func (choice3 *Exception_Choice3) GetGoName(yname string) string {
    if yname == "compress" { return "Compress" }
    if yname == "lower-limit" { return "LowerLimit" }
    if yname == "higher-limit" { return "HigherLimit" }
    if yname == "file-path" { return "FilePath" }
    if yname == "filename" { return "Filename" }
    return ""
}

func (choice3 *Exception_Choice3) GetSegmentPath() string {
    return "choice3"
}

func (choice3 *Exception_Choice3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (choice3 *Exception_Choice3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (choice3 *Exception_Choice3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["compress"] = choice3.Compress
    leafs["lower-limit"] = choice3.LowerLimit
    leafs["higher-limit"] = choice3.HigherLimit
    leafs["file-path"] = choice3.FilePath
    leafs["filename"] = choice3.Filename
    return leafs
}

func (choice3 *Exception_Choice3) GetBundleName() string { return "cisco_ios_xr" }

func (choice3 *Exception_Choice3) GetYangName() string { return "choice3" }

func (choice3 *Exception_Choice3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (choice3 *Exception_Choice3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (choice3 *Exception_Choice3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (choice3 *Exception_Choice3) SetParent(parent types.Entity) { choice3.parent = parent }

func (choice3 *Exception_Choice3) GetParent() types.Entity { return choice3.parent }

func (choice3 *Exception_Choice3) GetParentYangName() string { return "exception" }

// Exception_ProcessNames
// Specify per process configuration
type Exception_ProcessNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify per process configuration. The type is slice of
    // Exception_ProcessNames_ProcessName.
    ProcessName []Exception_ProcessNames_ProcessName
}

func (processNames *Exception_ProcessNames) GetFilter() yfilter.YFilter { return processNames.YFilter }

func (processNames *Exception_ProcessNames) SetFilter(yf yfilter.YFilter) { processNames.YFilter = yf }

func (processNames *Exception_ProcessNames) GetGoName(yname string) string {
    if yname == "process-name" { return "ProcessName" }
    return ""
}

func (processNames *Exception_ProcessNames) GetSegmentPath() string {
    return "process-names"
}

func (processNames *Exception_ProcessNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-name" {
        for _, c := range processNames.ProcessName {
            if processNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Exception_ProcessNames_ProcessName{}
        processNames.ProcessName = append(processNames.ProcessName, child)
        return &processNames.ProcessName[len(processNames.ProcessName)-1]
    }
    return nil
}

func (processNames *Exception_ProcessNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processNames.ProcessName {
        children[processNames.ProcessName[i].GetSegmentPath()] = &processNames.ProcessName[i]
    }
    return children
}

func (processNames *Exception_ProcessNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNames *Exception_ProcessNames) GetBundleName() string { return "cisco_ios_xr" }

func (processNames *Exception_ProcessNames) GetYangName() string { return "process-names" }

func (processNames *Exception_ProcessNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNames *Exception_ProcessNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNames *Exception_ProcessNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNames *Exception_ProcessNames) SetParent(parent types.Entity) { processNames.parent = parent }

func (processNames *Exception_ProcessNames) GetParent() types.Entity { return processNames.parent }

func (processNames *Exception_ProcessNames) GetParentYangName() string { return "exception" }

// Exception_ProcessNames_ProcessName
// Specify per process configuration
type Exception_ProcessNames_ProcessName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify per process configuration. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Processname interface{}

    // Specify per process core option.
    CoreOption Exception_ProcessNames_ProcessName_CoreOption
}

func (processName *Exception_ProcessNames_ProcessName) GetFilter() yfilter.YFilter { return processName.YFilter }

func (processName *Exception_ProcessNames_ProcessName) SetFilter(yf yfilter.YFilter) { processName.YFilter = yf }

func (processName *Exception_ProcessNames_ProcessName) GetGoName(yname string) string {
    if yname == "processname" { return "Processname" }
    if yname == "core-option" { return "CoreOption" }
    return ""
}

func (processName *Exception_ProcessNames_ProcessName) GetSegmentPath() string {
    return "process-name" + "[processname='" + fmt.Sprintf("%v", processName.Processname) + "']"
}

func (processName *Exception_ProcessNames_ProcessName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "core-option" {
        return &processName.CoreOption
    }
    return nil
}

func (processName *Exception_ProcessNames_ProcessName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["core-option"] = &processName.CoreOption
    return children
}

func (processName *Exception_ProcessNames_ProcessName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["processname"] = processName.Processname
    return leafs
}

func (processName *Exception_ProcessNames_ProcessName) GetBundleName() string { return "cisco_ios_xr" }

func (processName *Exception_ProcessNames_ProcessName) GetYangName() string { return "process-name" }

func (processName *Exception_ProcessNames_ProcessName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processName *Exception_ProcessNames_ProcessName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processName *Exception_ProcessNames_ProcessName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processName *Exception_ProcessNames_ProcessName) SetParent(parent types.Entity) { processName.parent = parent }

func (processName *Exception_ProcessNames_ProcessName) GetParent() types.Entity { return processName.parent }

func (processName *Exception_ProcessNames_ProcessName) GetParentYangName() string { return "process-names" }

// Exception_ProcessNames_ProcessName_CoreOption
// Specify per process core option
type Exception_ProcessNames_ProcessName_CoreOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dump main memory of the target process. The type is Mainmemory. The default
    // value is default.
    MainMemoryval interface{}

    // Dump shared memory of the target process. The type is Sharedmemory. The
    // default value is default.
    SharedMemoryval interface{}

    // Dump packet memory of the target process. The type is Packetmemory. The
    // default value is default.
    PacketMemoryval interface{}

    // Dump to local memory: for time critical processes. The type is Copy. The
    // default value is default.
    Copyval interface{}

    // Dump memory relevant to stack trace only: for time critical processes. The
    // type is Sparse. The default value is default.
    Sparseval interface{}

    // Skip CPU usage snapshot: for time critical processes. The type is
    // Skipcpuinfo. The default value is default.
    Skipcpuinfoval interface{}

    // Dump context info only: Overrides other options except for 'no-core'. The
    // type is Context. The default value is default.
    Contextval interface{}

    // Disable core dump for the target process: Overrides other options. The type
    // is Nocore. The default value is default.
    Nocoreval interface{}
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetFilter() yfilter.YFilter { return coreOption.YFilter }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) SetFilter(yf yfilter.YFilter) { coreOption.YFilter = yf }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetGoName(yname string) string {
    if yname == "main-memoryval" { return "MainMemoryval" }
    if yname == "shared-memoryval" { return "SharedMemoryval" }
    if yname == "packet-memoryval" { return "PacketMemoryval" }
    if yname == "copyval" { return "Copyval" }
    if yname == "sparseval" { return "Sparseval" }
    if yname == "skipcpuinfoval" { return "Skipcpuinfoval" }
    if yname == "contextval" { return "Contextval" }
    if yname == "nocoreval" { return "Nocoreval" }
    return ""
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetSegmentPath() string {
    return "core-option"
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["main-memoryval"] = coreOption.MainMemoryval
    leafs["shared-memoryval"] = coreOption.SharedMemoryval
    leafs["packet-memoryval"] = coreOption.PacketMemoryval
    leafs["copyval"] = coreOption.Copyval
    leafs["sparseval"] = coreOption.Sparseval
    leafs["skipcpuinfoval"] = coreOption.Skipcpuinfoval
    leafs["contextval"] = coreOption.Contextval
    leafs["nocoreval"] = coreOption.Nocoreval
    return leafs
}

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetBundleName() string { return "cisco_ios_xr" }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetYangName() string { return "core-option" }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) SetParent(parent types.Entity) { coreOption.parent = parent }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetParent() types.Entity { return coreOption.parent }

func (coreOption *Exception_ProcessNames_ProcessName_CoreOption) GetParentYangName() string { return "process-name" }

// Exception_Choice2
// Preference of the dump location
type Exception_Choice2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify 'true' to compress core files dumped on this path, 'false' to not
    // compress. The type is bool.
    Compress interface{}

    // Lower limit.  This is required if Filename is specified. The type is
    // interface{} with range: 0..4.
    LowerLimit interface{}

    // Higher limit.  This is required if Filename is specified. The type is
    // interface{} with range: 5..64.
    HigherLimit interface{}

    // Protocol and directory. The type is string.
    FilePath interface{}

    // Dump filename. The type is string.
    Filename interface{}
}

func (choice2 *Exception_Choice2) GetFilter() yfilter.YFilter { return choice2.YFilter }

func (choice2 *Exception_Choice2) SetFilter(yf yfilter.YFilter) { choice2.YFilter = yf }

func (choice2 *Exception_Choice2) GetGoName(yname string) string {
    if yname == "compress" { return "Compress" }
    if yname == "lower-limit" { return "LowerLimit" }
    if yname == "higher-limit" { return "HigherLimit" }
    if yname == "file-path" { return "FilePath" }
    if yname == "filename" { return "Filename" }
    return ""
}

func (choice2 *Exception_Choice2) GetSegmentPath() string {
    return "choice2"
}

func (choice2 *Exception_Choice2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (choice2 *Exception_Choice2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (choice2 *Exception_Choice2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["compress"] = choice2.Compress
    leafs["lower-limit"] = choice2.LowerLimit
    leafs["higher-limit"] = choice2.HigherLimit
    leafs["file-path"] = choice2.FilePath
    leafs["filename"] = choice2.Filename
    return leafs
}

func (choice2 *Exception_Choice2) GetBundleName() string { return "cisco_ios_xr" }

func (choice2 *Exception_Choice2) GetYangName() string { return "choice2" }

func (choice2 *Exception_Choice2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (choice2 *Exception_Choice2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (choice2 *Exception_Choice2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (choice2 *Exception_Choice2) SetParent(parent types.Entity) { choice2.parent = parent }

func (choice2 *Exception_Choice2) GetParent() types.Entity { return choice2.parent }

func (choice2 *Exception_Choice2) GetParentYangName() string { return "exception" }

