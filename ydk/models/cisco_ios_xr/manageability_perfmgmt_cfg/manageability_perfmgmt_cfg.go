// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-perfmgmt package configuration.
// 
// This module contains definitions
// for the following management objects:
//   perf-mgmt: Performance Management configuration & operations
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package manageability_perfmgmt_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_perfmgmt_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-manageability-perfmgmt-cfg perf-mgmt}", reflect.TypeOf(PerfMgmt{}))
    ydk.RegisterEntity("Cisco-IOS-XR-manageability-perfmgmt-cfg:perf-mgmt", reflect.TypeOf(PerfMgmt{}))
}

// PmThresholdOp represents Pm threshold op
type PmThresholdOp string

const (
    // Equal to
    PmThresholdOp_eq PmThresholdOp = "eq"

    // Not equal to
    PmThresholdOp_ne PmThresholdOp = "ne"

    // Less than
    PmThresholdOp_lt PmThresholdOp = "lt"

    // Less than or equal to
    PmThresholdOp_le PmThresholdOp = "le"

    // Greater than
    PmThresholdOp_gt PmThresholdOp = "gt"

    // Greater than or equal to
    PmThresholdOp_ge PmThresholdOp = "ge"

    // Not in Range
    PmThresholdOp_rg PmThresholdOp = "rg"
)

// PmThresholdRearm represents Pm threshold rearm
type PmThresholdRearm string

const (
    // Rearm Always
    PmThresholdRearm_always PmThresholdRearm = "always"

    // Rearm after window of sampling periods
    PmThresholdRearm_window PmThresholdRearm = "window"

    // Rearm after the first period when condition is
    // not met
    PmThresholdRearm_toggle PmThresholdRearm = "toggle"
)

// PerfMgmt
// Performance Management configuration & operations
type PerfMgmt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resources configuration.
    Resources PerfMgmt_Resources

    // Templates for collection of statistics.
    Statistics PerfMgmt_Statistics

    // Start data collection and/or threshold monitoring.
    Enable PerfMgmt_Enable

    // Configure regular expression group.
    RegExpGroups PerfMgmt_RegExpGroups

    // Container for threshold templates.
    Threshold PerfMgmt_Threshold
}

func (perfMgmt *PerfMgmt) GetFilter() yfilter.YFilter { return perfMgmt.YFilter }

func (perfMgmt *PerfMgmt) SetFilter(yf yfilter.YFilter) { perfMgmt.YFilter = yf }

func (perfMgmt *PerfMgmt) GetGoName(yname string) string {
    if yname == "resources" { return "Resources" }
    if yname == "statistics" { return "Statistics" }
    if yname == "enable" { return "Enable" }
    if yname == "reg-exp-groups" { return "RegExpGroups" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (perfMgmt *PerfMgmt) GetSegmentPath() string {
    return "Cisco-IOS-XR-manageability-perfmgmt-cfg:perf-mgmt"
}

func (perfMgmt *PerfMgmt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "resources" {
        return &perfMgmt.Resources
    }
    if childYangName == "statistics" {
        return &perfMgmt.Statistics
    }
    if childYangName == "enable" {
        return &perfMgmt.Enable
    }
    if childYangName == "reg-exp-groups" {
        return &perfMgmt.RegExpGroups
    }
    if childYangName == "threshold" {
        return &perfMgmt.Threshold
    }
    return nil
}

func (perfMgmt *PerfMgmt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["resources"] = &perfMgmt.Resources
    children["statistics"] = &perfMgmt.Statistics
    children["enable"] = &perfMgmt.Enable
    children["reg-exp-groups"] = &perfMgmt.RegExpGroups
    children["threshold"] = &perfMgmt.Threshold
    return children
}

func (perfMgmt *PerfMgmt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (perfMgmt *PerfMgmt) GetBundleName() string { return "cisco_ios_xr" }

func (perfMgmt *PerfMgmt) GetYangName() string { return "perf-mgmt" }

func (perfMgmt *PerfMgmt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (perfMgmt *PerfMgmt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (perfMgmt *PerfMgmt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (perfMgmt *PerfMgmt) SetParent(parent types.Entity) { perfMgmt.parent = parent }

func (perfMgmt *PerfMgmt) GetParent() types.Entity { return perfMgmt.parent }

func (perfMgmt *PerfMgmt) GetParentYangName() string { return "Cisco-IOS-XR-manageability-perfmgmt-cfg" }

// PerfMgmt_Resources
// Resources configuration
type PerfMgmt_Resources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the TFTP server IP address and directory name.
    TftpResources PerfMgmt_Resources_TftpResources

    // Configure local dump parameters.
    DumpLocal PerfMgmt_Resources_DumpLocal

    // Configure the memory usage limits of performance management.
    MemoryResources PerfMgmt_Resources_MemoryResources
}

func (resources *PerfMgmt_Resources) GetFilter() yfilter.YFilter { return resources.YFilter }

func (resources *PerfMgmt_Resources) SetFilter(yf yfilter.YFilter) { resources.YFilter = yf }

func (resources *PerfMgmt_Resources) GetGoName(yname string) string {
    if yname == "tftp-resources" { return "TftpResources" }
    if yname == "dump-local" { return "DumpLocal" }
    if yname == "memory-resources" { return "MemoryResources" }
    return ""
}

func (resources *PerfMgmt_Resources) GetSegmentPath() string {
    return "resources"
}

func (resources *PerfMgmt_Resources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tftp-resources" {
        return &resources.TftpResources
    }
    if childYangName == "dump-local" {
        return &resources.DumpLocal
    }
    if childYangName == "memory-resources" {
        return &resources.MemoryResources
    }
    return nil
}

func (resources *PerfMgmt_Resources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tftp-resources"] = &resources.TftpResources
    children["dump-local"] = &resources.DumpLocal
    children["memory-resources"] = &resources.MemoryResources
    return children
}

func (resources *PerfMgmt_Resources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (resources *PerfMgmt_Resources) GetBundleName() string { return "cisco_ios_xr" }

func (resources *PerfMgmt_Resources) GetYangName() string { return "resources" }

func (resources *PerfMgmt_Resources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (resources *PerfMgmt_Resources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (resources *PerfMgmt_Resources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (resources *PerfMgmt_Resources) SetParent(parent types.Entity) { resources.parent = parent }

func (resources *PerfMgmt_Resources) GetParent() types.Entity { return resources.parent }

func (resources *PerfMgmt_Resources) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Resources_TftpResources
// Configure the TFTP server IP address and
// directory name
// This type is a presence type.
type PerfMgmt_Resources_TftpResources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of the TFTP server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    ServerAddress interface{}

    // Directory name on TFTP server. The type is string. This attribute is
    // mandatory.
    Directory interface{}

    // VRF name. The type is string with length: 1..32.
    VrfName interface{}
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetFilter() yfilter.YFilter { return tftpResources.YFilter }

func (tftpResources *PerfMgmt_Resources_TftpResources) SetFilter(yf yfilter.YFilter) { tftpResources.YFilter = yf }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "directory" { return "Directory" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetSegmentPath() string {
    return "tftp-resources"
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = tftpResources.ServerAddress
    leafs["directory"] = tftpResources.Directory
    leafs["vrf-name"] = tftpResources.VrfName
    return leafs
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetBundleName() string { return "cisco_ios_xr" }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetYangName() string { return "tftp-resources" }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tftpResources *PerfMgmt_Resources_TftpResources) SetParent(parent types.Entity) { tftpResources.parent = parent }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetParent() types.Entity { return tftpResources.parent }

func (tftpResources *PerfMgmt_Resources_TftpResources) GetParentYangName() string { return "resources" }

// PerfMgmt_Resources_DumpLocal
// Configure local dump parameters
type PerfMgmt_Resources_DumpLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable data dump onto local filesystem. The type is interface{}.
    Enable interface{}
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetFilter() yfilter.YFilter { return dumpLocal.YFilter }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) SetFilter(yf yfilter.YFilter) { dumpLocal.YFilter = yf }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetSegmentPath() string {
    return "dump-local"
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = dumpLocal.Enable
    return leafs
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetBundleName() string { return "cisco_ios_xr" }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetYangName() string { return "dump-local" }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) SetParent(parent types.Entity) { dumpLocal.parent = parent }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetParent() types.Entity { return dumpLocal.parent }

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetParentYangName() string { return "resources" }

// PerfMgmt_Resources_MemoryResources
// Configure the memory usage limits of
// performance management
type PerfMgmt_Resources_MemoryResources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum limit for memory usage (Kbytes) for data buffers. The type is
    // interface{} with range: -2147483648..2147483647. Units are kilobyte.
    MaxLimit interface{}

    // Specify a minimum free memory (Kbytes) to be ensured before allowing a
    // collection request. The type is interface{} with range:
    // -2147483648..2147483647. Units are kilobyte.
    MinReserved interface{}
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetFilter() yfilter.YFilter { return memoryResources.YFilter }

func (memoryResources *PerfMgmt_Resources_MemoryResources) SetFilter(yf yfilter.YFilter) { memoryResources.YFilter = yf }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetGoName(yname string) string {
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "min-reserved" { return "MinReserved" }
    return ""
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetSegmentPath() string {
    return "memory-resources"
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-limit"] = memoryResources.MaxLimit
    leafs["min-reserved"] = memoryResources.MinReserved
    return leafs
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetBundleName() string { return "cisco_ios_xr" }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetYangName() string { return "memory-resources" }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryResources *PerfMgmt_Resources_MemoryResources) SetParent(parent types.Entity) { memoryResources.parent = parent }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetParent() types.Entity { return memoryResources.parent }

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetParentYangName() string { return "resources" }

// PerfMgmt_Statistics
// Templates for collection of statistics
type PerfMgmt_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Generic GenericCounter collection templates.
    GenericCounterInterface PerfMgmt_Statistics_GenericCounterInterface

    // Node Process collection templates.
    ProcessNode PerfMgmt_Statistics_ProcessNode

    // Interface BasicCounter collection templates.
    BasicCounterInterface PerfMgmt_Statistics_BasicCounterInterface

    // OSPF v3 Protocol collection templates.
    Ospfv3Protocol PerfMgmt_Statistics_Ospfv3Protocol

    // Node CPU collection templates.
    CpuNode PerfMgmt_Statistics_CpuNode

    // Interface DataRate collection templates.
    DataRateInterface PerfMgmt_Statistics_DataRateInterface

    // Node Memory collection templates.
    MemoryNode PerfMgmt_Statistics_MemoryNode

    // MPLS LDP collection templates.
    LdpMpls PerfMgmt_Statistics_LdpMpls

    // BGP collection templates.
    Bgp PerfMgmt_Statistics_Bgp

    // OSPF v2 Protocol collection templates.
    Ospfv2Protocol PerfMgmt_Statistics_Ospfv2Protocol
}

func (statistics *PerfMgmt_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *PerfMgmt_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *PerfMgmt_Statistics) GetGoName(yname string) string {
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    if yname == "process-node" { return "ProcessNode" }
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    if yname == "ospfv3-protocol" { return "Ospfv3Protocol" }
    if yname == "cpu-node" { return "CpuNode" }
    if yname == "data-rate-interface" { return "DataRateInterface" }
    if yname == "memory-node" { return "MemoryNode" }
    if yname == "ldp-mpls" { return "LdpMpls" }
    if yname == "bgp" { return "Bgp" }
    if yname == "ospfv2-protocol" { return "Ospfv2Protocol" }
    return ""
}

func (statistics *PerfMgmt_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *PerfMgmt_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface" {
        return &statistics.GenericCounterInterface
    }
    if childYangName == "process-node" {
        return &statistics.ProcessNode
    }
    if childYangName == "basic-counter-interface" {
        return &statistics.BasicCounterInterface
    }
    if childYangName == "ospfv3-protocol" {
        return &statistics.Ospfv3Protocol
    }
    if childYangName == "cpu-node" {
        return &statistics.CpuNode
    }
    if childYangName == "data-rate-interface" {
        return &statistics.DataRateInterface
    }
    if childYangName == "memory-node" {
        return &statistics.MemoryNode
    }
    if childYangName == "ldp-mpls" {
        return &statistics.LdpMpls
    }
    if childYangName == "bgp" {
        return &statistics.Bgp
    }
    if childYangName == "ospfv2-protocol" {
        return &statistics.Ospfv2Protocol
    }
    return nil
}

func (statistics *PerfMgmt_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interface"] = &statistics.GenericCounterInterface
    children["process-node"] = &statistics.ProcessNode
    children["basic-counter-interface"] = &statistics.BasicCounterInterface
    children["ospfv3-protocol"] = &statistics.Ospfv3Protocol
    children["cpu-node"] = &statistics.CpuNode
    children["data-rate-interface"] = &statistics.DataRateInterface
    children["memory-node"] = &statistics.MemoryNode
    children["ldp-mpls"] = &statistics.LdpMpls
    children["bgp"] = &statistics.Bgp
    children["ospfv2-protocol"] = &statistics.Ospfv2Protocol
    return children
}

func (statistics *PerfMgmt_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *PerfMgmt_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *PerfMgmt_Statistics) GetYangName() string { return "statistics" }

func (statistics *PerfMgmt_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *PerfMgmt_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *PerfMgmt_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *PerfMgmt_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *PerfMgmt_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *PerfMgmt_Statistics) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Statistics_GenericCounterInterface
// Interface Generic GenericCounter collection
// templates
type PerfMgmt_Statistics_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_GenericCounterInterface_Templates
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface"
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &genericCounterInterface.Templates
    }
    return nil
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &genericCounterInterface.Templates
    return children
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_GenericCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_GenericCounterInterface_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_GenericCounterInterface_Templates_Template.
    Template []PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_GenericCounterInterface_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetParentYangName() string { return "generic-counter-interface" }

// PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_GenericCounterInterface_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_ProcessNode
// Node Process collection templates
type PerfMgmt_Statistics_ProcessNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_ProcessNode_Templates
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetFilter() yfilter.YFilter { return processNode.YFilter }

func (processNode *PerfMgmt_Statistics_ProcessNode) SetFilter(yf yfilter.YFilter) { processNode.YFilter = yf }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetSegmentPath() string {
    return "process-node"
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &processNode.Templates
    }
    return nil
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &processNode.Templates
    return children
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetBundleName() string { return "cisco_ios_xr" }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetYangName() string { return "process-node" }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNode *PerfMgmt_Statistics_ProcessNode) SetParent(parent types.Entity) { processNode.parent = parent }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetParent() types.Entity { return processNode.parent }

func (processNode *PerfMgmt_Statistics_ProcessNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_ProcessNode_Templates
// Template name
type PerfMgmt_Statistics_ProcessNode_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_ProcessNode_Templates_Template.
    Template []PerfMgmt_Statistics_ProcessNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_ProcessNode_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetParentYangName() string { return "process-node" }

// PerfMgmt_Statistics_ProcessNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_ProcessNode_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_BasicCounterInterface
// Interface BasicCounter collection templates
type PerfMgmt_Statistics_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_BasicCounterInterface_Templates
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface"
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &basicCounterInterface.Templates
    }
    return nil
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &basicCounterInterface.Templates
    return children
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_BasicCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_BasicCounterInterface_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_BasicCounterInterface_Templates_Template.
    Template []PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_BasicCounterInterface_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetParentYangName() string { return "basic-counter-interface" }

// PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_BasicCounterInterface_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_Ospfv3Protocol
// OSPF v3 Protocol collection templates
type PerfMgmt_Statistics_Ospfv3Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Ospfv3Protocol_Templates
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetFilter() yfilter.YFilter { return ospfv3Protocol.YFilter }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) SetFilter(yf yfilter.YFilter) { ospfv3Protocol.YFilter = yf }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetSegmentPath() string {
    return "ospfv3-protocol"
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &ospfv3Protocol.Templates
    }
    return nil
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &ospfv3Protocol.Templates
    return children
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetYangName() string { return "ospfv3-protocol" }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) SetParent(parent types.Entity) { ospfv3Protocol.parent = parent }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetParent() types.Entity { return ospfv3Protocol.parent }

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_Ospfv3Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv3Protocol_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template.
    Template []PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetParentYangName() string { return "ospfv3-protocol" }

// PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_CpuNode
// Node CPU collection templates
type PerfMgmt_Statistics_CpuNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_CpuNode_Templates
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetFilter() yfilter.YFilter { return cpuNode.YFilter }

func (cpuNode *PerfMgmt_Statistics_CpuNode) SetFilter(yf yfilter.YFilter) { cpuNode.YFilter = yf }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetSegmentPath() string {
    return "cpu-node"
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &cpuNode.Templates
    }
    return nil
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &cpuNode.Templates
    return children
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetYangName() string { return "cpu-node" }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNode *PerfMgmt_Statistics_CpuNode) SetParent(parent types.Entity) { cpuNode.parent = parent }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetParent() types.Entity { return cpuNode.parent }

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_CpuNode_Templates
// Template name
type PerfMgmt_Statistics_CpuNode_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_CpuNode_Templates_Template.
    Template []PerfMgmt_Statistics_CpuNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_CpuNode_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Statistics_CpuNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_CpuNode_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_DataRateInterface
// Interface DataRate collection templates
type PerfMgmt_Statistics_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_DataRateInterface_Templates
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface"
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &dataRateInterface.Templates
    }
    return nil
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &dataRateInterface.Templates
    return children
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_DataRateInterface_Templates
// Template name
type PerfMgmt_Statistics_DataRateInterface_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_DataRateInterface_Templates_Template.
    Template []PerfMgmt_Statistics_DataRateInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_DataRateInterface_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetParentYangName() string { return "data-rate-interface" }

// PerfMgmt_Statistics_DataRateInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_DataRateInterface_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_MemoryNode
// Node Memory collection templates
type PerfMgmt_Statistics_MemoryNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_MemoryNode_Templates
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetFilter() yfilter.YFilter { return memoryNode.YFilter }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) SetFilter(yf yfilter.YFilter) { memoryNode.YFilter = yf }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetSegmentPath() string {
    return "memory-node"
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &memoryNode.Templates
    }
    return nil
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &memoryNode.Templates
    return children
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetYangName() string { return "memory-node" }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) SetParent(parent types.Entity) { memoryNode.parent = parent }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetParent() types.Entity { return memoryNode.parent }

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_MemoryNode_Templates
// Template name
type PerfMgmt_Statistics_MemoryNode_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_MemoryNode_Templates_Template.
    Template []PerfMgmt_Statistics_MemoryNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_MemoryNode_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Statistics_MemoryNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_MemoryNode_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_LdpMpls
// MPLS LDP collection templates
type PerfMgmt_Statistics_LdpMpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_LdpMpls_Templates
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetFilter() yfilter.YFilter { return ldpMpls.YFilter }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) SetFilter(yf yfilter.YFilter) { ldpMpls.YFilter = yf }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetSegmentPath() string {
    return "ldp-mpls"
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &ldpMpls.Templates
    }
    return nil
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &ldpMpls.Templates
    return children
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetYangName() string { return "ldp-mpls" }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) SetParent(parent types.Entity) { ldpMpls.parent = parent }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetParent() types.Entity { return ldpMpls.parent }

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_LdpMpls_Templates
// Template name
type PerfMgmt_Statistics_LdpMpls_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_LdpMpls_Templates_Template.
    Template []PerfMgmt_Statistics_LdpMpls_Templates_Template
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_LdpMpls_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetParentYangName() string { return "ldp-mpls" }

// PerfMgmt_Statistics_LdpMpls_Templates_Template
// A template instance
type PerfMgmt_Statistics_LdpMpls_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_Bgp
// BGP collection templates
type PerfMgmt_Statistics_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Bgp_Templates
}

func (bgp *PerfMgmt_Statistics_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Statistics_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Statistics_Bgp) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (bgp *PerfMgmt_Statistics_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Statistics_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &bgp.Templates
    }
    return nil
}

func (bgp *PerfMgmt_Statistics_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &bgp.Templates
    return children
}

func (bgp *PerfMgmt_Statistics_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *PerfMgmt_Statistics_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Statistics_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Statistics_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Statistics_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Statistics_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Statistics_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Statistics_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Statistics_Bgp) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_Bgp_Templates
// Template name
type PerfMgmt_Statistics_Bgp_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Bgp_Templates_Template.
    Template []PerfMgmt_Statistics_Bgp_Templates_Template
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_Bgp_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_Bgp_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_Bgp_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetParentYangName() string { return "bgp" }

// PerfMgmt_Statistics_Bgp_Templates_Template
// A template instance
type PerfMgmt_Statistics_Bgp_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Statistics_Ospfv2Protocol
// OSPF v2 Protocol collection templates
type PerfMgmt_Statistics_Ospfv2Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Ospfv2Protocol_Templates
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetFilter() yfilter.YFilter { return ospfv2Protocol.YFilter }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) SetFilter(yf yfilter.YFilter) { ospfv2Protocol.YFilter = yf }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetGoName(yname string) string {
    if yname == "templates" { return "Templates" }
    return ""
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetSegmentPath() string {
    return "ospfv2-protocol"
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "templates" {
        return &ospfv2Protocol.Templates
    }
    return nil
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["templates"] = &ospfv2Protocol.Templates
    return children
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetYangName() string { return "ospfv2-protocol" }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) SetParent(parent types.Entity) { ospfv2Protocol.parent = parent }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetParent() types.Entity { return ospfv2Protocol.parent }

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetParentYangName() string { return "statistics" }

// PerfMgmt_Statistics_Ospfv2Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv2Protocol_Templates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template.
    Template []PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetFilter() yfilter.YFilter { return templates.YFilter }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) SetFilter(yf yfilter.YFilter) { templates.YFilter = yf }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetGoName(yname string) string {
    if yname == "template" { return "Template" }
    return ""
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetSegmentPath() string {
    return "templates"
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "template" {
        for _, c := range templates.Template {
            if templates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template{}
        templates.Template = append(templates.Template, child)
        return &templates.Template[len(templates.Template)-1]
    }
    return nil
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range templates.Template {
        children[templates.Template[i].GetSegmentPath()] = &templates.Template[i]
    }
    return children
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetBundleName() string { return "cisco_ios_xr" }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetYangName() string { return "templates" }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) SetParent(parent types.Entity) { templates.parent = parent }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetParent() types.Entity { return templates.parent }

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetParentYangName() string { return "ospfv2-protocol" }

// PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable persistent history statistics. The type is interface{}.
    HistoryPersistent interface{}

    // VRF group configured in regular expression to be applied. The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Frequency of each sample in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of samples to be taken. The type is interface{} with range: 1..60.
    SampleSize interface{}
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetFilter() yfilter.YFilter { return template.YFilter }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) SetFilter(yf yfilter.YFilter) { template.YFilter = yf }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "history-persistent" { return "HistoryPersistent" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "sample-size" { return "SampleSize" }
    return ""
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetSegmentPath() string {
    return "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = template.TemplateName
    leafs["reg-exp-group"] = template.RegExpGroup
    leafs["history-persistent"] = template.HistoryPersistent
    leafs["vrf-group"] = template.VrfGroup
    leafs["sample-interval"] = template.SampleInterval
    leafs["sample-size"] = template.SampleSize
    return leafs
}

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetBundleName() string { return "cisco_ios_xr" }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetYangName() string { return "template" }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) SetParent(parent types.Entity) { template.parent = parent }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetParent() types.Entity { return template.parent }

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetParentYangName() string { return "templates" }

// PerfMgmt_Enable
// Start data collection and/or threshold
// monitoring
type PerfMgmt_Enable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start threshold monitoring using a defined template.
    Threshold PerfMgmt_Enable_Threshold

    // Start periodic collection using a defined a template.
    Statistics PerfMgmt_Enable_Statistics

    // Start data collection for a monitored instance.
    MonitorEnable PerfMgmt_Enable_MonitorEnable
}

func (enable *PerfMgmt_Enable) GetFilter() yfilter.YFilter { return enable.YFilter }

func (enable *PerfMgmt_Enable) SetFilter(yf yfilter.YFilter) { enable.YFilter = yf }

func (enable *PerfMgmt_Enable) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "statistics" { return "Statistics" }
    if yname == "monitor-enable" { return "MonitorEnable" }
    return ""
}

func (enable *PerfMgmt_Enable) GetSegmentPath() string {
    return "enable"
}

func (enable *PerfMgmt_Enable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &enable.Threshold
    }
    if childYangName == "statistics" {
        return &enable.Statistics
    }
    if childYangName == "monitor-enable" {
        return &enable.MonitorEnable
    }
    return nil
}

func (enable *PerfMgmt_Enable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &enable.Threshold
    children["statistics"] = &enable.Statistics
    children["monitor-enable"] = &enable.MonitorEnable
    return children
}

func (enable *PerfMgmt_Enable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (enable *PerfMgmt_Enable) GetBundleName() string { return "cisco_ios_xr" }

func (enable *PerfMgmt_Enable) GetYangName() string { return "enable" }

func (enable *PerfMgmt_Enable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enable *PerfMgmt_Enable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enable *PerfMgmt_Enable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enable *PerfMgmt_Enable) SetParent(parent types.Entity) { enable.parent = parent }

func (enable *PerfMgmt_Enable) GetParent() types.Entity { return enable.parent }

func (enable *PerfMgmt_Enable) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Enable_Threshold
// Start threshold monitoring using a defined
// template
type PerfMgmt_Enable_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold monitoring for OSPF v3 Protocol.
    Ospfv3Protocol PerfMgmt_Enable_Threshold_Ospfv3Protocol

    // Threshold monitoring for BGP.
    Bgp PerfMgmt_Enable_Threshold_Bgp

    // Threshold monitoring for Interface  data-rates.
    DataRateInterface PerfMgmt_Enable_Threshold_DataRateInterface

    // Threshold monitoring for OSPF v2 Protocol.
    Ospfv2Protocol PerfMgmt_Enable_Threshold_Ospfv2Protocol

    // Threshold monitoring for memory.
    MemoryNode PerfMgmt_Enable_Threshold_MemoryNode

    // Threshold monitoring for Interface generic-counters.
    GenericCounterInterface PerfMgmt_Enable_Threshold_GenericCounterInterface

    // Threshold monitoring for CPU.
    CpuNode PerfMgmt_Enable_Threshold_CpuNode

    // Threshold monitoring for LDP.
    LdpMpls PerfMgmt_Enable_Threshold_LdpMpls

    // Threshold monitoring for process.
    ProcessNode PerfMgmt_Enable_Threshold_ProcessNode

    // Threshold monitoring for Interface basic-counters.
    BasicCounterInterface PerfMgmt_Enable_Threshold_BasicCounterInterface
}

func (threshold *PerfMgmt_Enable_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *PerfMgmt_Enable_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *PerfMgmt_Enable_Threshold) GetGoName(yname string) string {
    if yname == "ospfv3-protocol" { return "Ospfv3Protocol" }
    if yname == "bgp" { return "Bgp" }
    if yname == "data-rate-interface" { return "DataRateInterface" }
    if yname == "ospfv2-protocol" { return "Ospfv2Protocol" }
    if yname == "memory-node" { return "MemoryNode" }
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    if yname == "cpu-node" { return "CpuNode" }
    if yname == "ldp-mpls" { return "LdpMpls" }
    if yname == "process-node" { return "ProcessNode" }
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    return ""
}

func (threshold *PerfMgmt_Enable_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *PerfMgmt_Enable_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv3-protocol" {
        return &threshold.Ospfv3Protocol
    }
    if childYangName == "bgp" {
        return &threshold.Bgp
    }
    if childYangName == "data-rate-interface" {
        return &threshold.DataRateInterface
    }
    if childYangName == "ospfv2-protocol" {
        return &threshold.Ospfv2Protocol
    }
    if childYangName == "memory-node" {
        return &threshold.MemoryNode
    }
    if childYangName == "generic-counter-interface" {
        return &threshold.GenericCounterInterface
    }
    if childYangName == "cpu-node" {
        return &threshold.CpuNode
    }
    if childYangName == "ldp-mpls" {
        return &threshold.LdpMpls
    }
    if childYangName == "process-node" {
        return &threshold.ProcessNode
    }
    if childYangName == "basic-counter-interface" {
        return &threshold.BasicCounterInterface
    }
    return nil
}

func (threshold *PerfMgmt_Enable_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv3-protocol"] = &threshold.Ospfv3Protocol
    children["bgp"] = &threshold.Bgp
    children["data-rate-interface"] = &threshold.DataRateInterface
    children["ospfv2-protocol"] = &threshold.Ospfv2Protocol
    children["memory-node"] = &threshold.MemoryNode
    children["generic-counter-interface"] = &threshold.GenericCounterInterface
    children["cpu-node"] = &threshold.CpuNode
    children["ldp-mpls"] = &threshold.LdpMpls
    children["process-node"] = &threshold.ProcessNode
    children["basic-counter-interface"] = &threshold.BasicCounterInterface
    return children
}

func (threshold *PerfMgmt_Enable_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (threshold *PerfMgmt_Enable_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *PerfMgmt_Enable_Threshold) GetYangName() string { return "threshold" }

func (threshold *PerfMgmt_Enable_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *PerfMgmt_Enable_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *PerfMgmt_Enable_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *PerfMgmt_Enable_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *PerfMgmt_Enable_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *PerfMgmt_Enable_Threshold) GetParentYangName() string { return "enable" }

// PerfMgmt_Enable_Threshold_Ospfv3Protocol
// Threshold monitoring for OSPF v3 Protocol
type PerfMgmt_Enable_Threshold_Ospfv3Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetFilter() yfilter.YFilter { return ospfv3Protocol.YFilter }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) SetFilter(yf yfilter.YFilter) { ospfv3Protocol.YFilter = yf }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetSegmentPath() string {
    return "ospfv3-protocol"
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv3Protocol.TemplateName
    return leafs
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetYangName() string { return "ospfv3-protocol" }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) SetParent(parent types.Entity) { ospfv3Protocol.parent = parent }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetParent() types.Entity { return ospfv3Protocol.parent }

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_Bgp
// Threshold monitoring for BGP
type PerfMgmt_Enable_Threshold_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = bgp.TemplateName
    return leafs
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_DataRateInterface
// Threshold monitoring for Interface  data-rates
type PerfMgmt_Enable_Threshold_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface"
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = dataRateInterface.TemplateName
    return leafs
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_Ospfv2Protocol
// Threshold monitoring for OSPF v2 Protocol
type PerfMgmt_Enable_Threshold_Ospfv2Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetFilter() yfilter.YFilter { return ospfv2Protocol.YFilter }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) SetFilter(yf yfilter.YFilter) { ospfv2Protocol.YFilter = yf }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetSegmentPath() string {
    return "ospfv2-protocol"
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv2Protocol.TemplateName
    return leafs
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetYangName() string { return "ospfv2-protocol" }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) SetParent(parent types.Entity) { ospfv2Protocol.parent = parent }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetParent() types.Entity { return ospfv2Protocol.parent }

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_MemoryNode
// Threshold monitoring for memory
type PerfMgmt_Enable_Threshold_MemoryNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_MemoryNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_MemoryNode_NodeAll
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetFilter() yfilter.YFilter { return memoryNode.YFilter }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) SetFilter(yf yfilter.YFilter) { memoryNode.YFilter = yf }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "node-all" { return "NodeAll" }
    return ""
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetSegmentPath() string {
    return "memory-node"
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &memoryNode.Nodes
    }
    if childYangName == "node-all" {
        return &memoryNode.NodeAll
    }
    return nil
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &memoryNode.Nodes
    children["node-all"] = &memoryNode.NodeAll
    return children
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetYangName() string { return "memory-node" }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) SetParent(parent types.Entity) { memoryNode.parent = parent }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetParent() types.Entity { return memoryNode.parent }

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Threshold_MemoryNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_MemoryNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Enable_Threshold_GenericCounterInterface
// Threshold monitoring for Interface
// generic-counters
type PerfMgmt_Enable_Threshold_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface"
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = genericCounterInterface.TemplateName
    return leafs
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_CpuNode
// Threshold monitoring for CPU
type PerfMgmt_Enable_Threshold_CpuNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_CpuNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_CpuNode_NodeAll
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetFilter() yfilter.YFilter { return cpuNode.YFilter }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) SetFilter(yf yfilter.YFilter) { cpuNode.YFilter = yf }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "node-all" { return "NodeAll" }
    return ""
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetSegmentPath() string {
    return "cpu-node"
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &cpuNode.Nodes
    }
    if childYangName == "node-all" {
        return &cpuNode.NodeAll
    }
    return nil
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &cpuNode.Nodes
    children["node-all"] = &cpuNode.NodeAll
    return children
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetYangName() string { return "cpu-node" }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) SetParent(parent types.Entity) { cpuNode.parent = parent }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetParent() types.Entity { return cpuNode.parent }

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_CpuNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Threshold_CpuNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_CpuNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Enable_Threshold_LdpMpls
// Threshold monitoring for LDP
type PerfMgmt_Enable_Threshold_LdpMpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetFilter() yfilter.YFilter { return ldpMpls.YFilter }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) SetFilter(yf yfilter.YFilter) { ldpMpls.YFilter = yf }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetSegmentPath() string {
    return "ldp-mpls"
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ldpMpls.TemplateName
    return leafs
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetYangName() string { return "ldp-mpls" }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) SetParent(parent types.Entity) { ldpMpls.parent = parent }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetParent() types.Entity { return ldpMpls.parent }

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_ProcessNode
// Threshold monitoring for process
type PerfMgmt_Enable_Threshold_ProcessNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_ProcessNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_ProcessNode_NodeAll
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetFilter() yfilter.YFilter { return processNode.YFilter }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) SetFilter(yf yfilter.YFilter) { processNode.YFilter = yf }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "node-all" { return "NodeAll" }
    return ""
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetSegmentPath() string {
    return "process-node"
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &processNode.Nodes
    }
    if childYangName == "node-all" {
        return &processNode.NodeAll
    }
    return nil
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &processNode.Nodes
    children["node-all"] = &processNode.NodeAll
    return children
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetBundleName() string { return "cisco_ios_xr" }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetYangName() string { return "process-node" }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) SetParent(parent types.Entity) { processNode.parent = parent }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetParent() types.Entity { return processNode.parent }

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetParentYangName() string { return "process-node" }

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Threshold_ProcessNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_ProcessNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetParentYangName() string { return "process-node" }

// PerfMgmt_Enable_Threshold_BasicCounterInterface
// Threshold monitoring for Interface
// basic-counters
type PerfMgmt_Enable_Threshold_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface"
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = basicCounterInterface.TemplateName
    return leafs
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Enable_Statistics
// Start periodic collection using a defined a
// template
type PerfMgmt_Enable_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics collection for generic-counters.
    GenericCounterInterface PerfMgmt_Enable_Statistics_GenericCounterInterface

    // Data collection for BGP.
    Bgp PerfMgmt_Enable_Statistics_Bgp

    // Data collection for OSPF v2 Protocol.
    Ospfv2Protocol PerfMgmt_Enable_Statistics_Ospfv2Protocol

    // Data collection for OSPF v3 Protocol.
    Ospfv3Protocol PerfMgmt_Enable_Statistics_Ospfv3Protocol

    // Collection for CPU.
    CpuNode PerfMgmt_Enable_Statistics_CpuNode

    // Statistics collection for basic-counters.
    BasicCounterInterface PerfMgmt_Enable_Statistics_BasicCounterInterface

    // Collection for process.
    ProcessNode PerfMgmt_Enable_Statistics_ProcessNode

    // Statistics collection for generic-counters.
    DataRateInterface PerfMgmt_Enable_Statistics_DataRateInterface

    // Collection for memory.
    MemoryNode PerfMgmt_Enable_Statistics_MemoryNode

    // Collection for labels distribution protocol.
    LdpMpls PerfMgmt_Enable_Statistics_LdpMpls
}

func (statistics *PerfMgmt_Enable_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *PerfMgmt_Enable_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *PerfMgmt_Enable_Statistics) GetGoName(yname string) string {
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    if yname == "bgp" { return "Bgp" }
    if yname == "ospfv2-protocol" { return "Ospfv2Protocol" }
    if yname == "ospfv3-protocol" { return "Ospfv3Protocol" }
    if yname == "cpu-node" { return "CpuNode" }
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    if yname == "process-node" { return "ProcessNode" }
    if yname == "data-rate-interface" { return "DataRateInterface" }
    if yname == "memory-node" { return "MemoryNode" }
    if yname == "ldp-mpls" { return "LdpMpls" }
    return ""
}

func (statistics *PerfMgmt_Enable_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *PerfMgmt_Enable_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface" {
        return &statistics.GenericCounterInterface
    }
    if childYangName == "bgp" {
        return &statistics.Bgp
    }
    if childYangName == "ospfv2-protocol" {
        return &statistics.Ospfv2Protocol
    }
    if childYangName == "ospfv3-protocol" {
        return &statistics.Ospfv3Protocol
    }
    if childYangName == "cpu-node" {
        return &statistics.CpuNode
    }
    if childYangName == "basic-counter-interface" {
        return &statistics.BasicCounterInterface
    }
    if childYangName == "process-node" {
        return &statistics.ProcessNode
    }
    if childYangName == "data-rate-interface" {
        return &statistics.DataRateInterface
    }
    if childYangName == "memory-node" {
        return &statistics.MemoryNode
    }
    if childYangName == "ldp-mpls" {
        return &statistics.LdpMpls
    }
    return nil
}

func (statistics *PerfMgmt_Enable_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interface"] = &statistics.GenericCounterInterface
    children["bgp"] = &statistics.Bgp
    children["ospfv2-protocol"] = &statistics.Ospfv2Protocol
    children["ospfv3-protocol"] = &statistics.Ospfv3Protocol
    children["cpu-node"] = &statistics.CpuNode
    children["basic-counter-interface"] = &statistics.BasicCounterInterface
    children["process-node"] = &statistics.ProcessNode
    children["data-rate-interface"] = &statistics.DataRateInterface
    children["memory-node"] = &statistics.MemoryNode
    children["ldp-mpls"] = &statistics.LdpMpls
    return children
}

func (statistics *PerfMgmt_Enable_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *PerfMgmt_Enable_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *PerfMgmt_Enable_Statistics) GetYangName() string { return "statistics" }

func (statistics *PerfMgmt_Enable_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *PerfMgmt_Enable_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *PerfMgmt_Enable_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *PerfMgmt_Enable_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *PerfMgmt_Enable_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *PerfMgmt_Enable_Statistics) GetParentYangName() string { return "enable" }

// PerfMgmt_Enable_Statistics_GenericCounterInterface
// Statistics collection for generic-counters
type PerfMgmt_Enable_Statistics_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface"
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = genericCounterInterface.TemplateName
    return leafs
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_Bgp
// Data collection for BGP
type PerfMgmt_Enable_Statistics_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = bgp.TemplateName
    return leafs
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_Ospfv2Protocol
// Data collection for OSPF v2 Protocol
type PerfMgmt_Enable_Statistics_Ospfv2Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetFilter() yfilter.YFilter { return ospfv2Protocol.YFilter }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) SetFilter(yf yfilter.YFilter) { ospfv2Protocol.YFilter = yf }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetSegmentPath() string {
    return "ospfv2-protocol"
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv2Protocol.TemplateName
    return leafs
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetYangName() string { return "ospfv2-protocol" }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) SetParent(parent types.Entity) { ospfv2Protocol.parent = parent }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetParent() types.Entity { return ospfv2Protocol.parent }

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_Ospfv3Protocol
// Data collection for OSPF v3 Protocol
type PerfMgmt_Enable_Statistics_Ospfv3Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetFilter() yfilter.YFilter { return ospfv3Protocol.YFilter }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) SetFilter(yf yfilter.YFilter) { ospfv3Protocol.YFilter = yf }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetSegmentPath() string {
    return "ospfv3-protocol"
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv3Protocol.TemplateName
    return leafs
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetYangName() string { return "ospfv3-protocol" }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) SetParent(parent types.Entity) { ospfv3Protocol.parent = parent }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetParent() types.Entity { return ospfv3Protocol.parent }

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_CpuNode
// Collection for CPU
type PerfMgmt_Enable_Statistics_CpuNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_CpuNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_CpuNode_Nodes
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetFilter() yfilter.YFilter { return cpuNode.YFilter }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) SetFilter(yf yfilter.YFilter) { cpuNode.YFilter = yf }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetGoName(yname string) string {
    if yname == "node-all" { return "NodeAll" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetSegmentPath() string {
    return "cpu-node"
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-all" {
        return &cpuNode.NodeAll
    }
    if childYangName == "nodes" {
        return &cpuNode.Nodes
    }
    return nil
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-all"] = &cpuNode.NodeAll
    children["nodes"] = &cpuNode.Nodes
    return children
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetYangName() string { return "cpu-node" }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) SetParent(parent types.Entity) { cpuNode.parent = parent }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetParent() types.Entity { return cpuNode.parent }

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_CpuNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_CpuNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Enable_Statistics_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_CpuNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Statistics_BasicCounterInterface
// Statistics collection for basic-counters
type PerfMgmt_Enable_Statistics_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface"
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = basicCounterInterface.TemplateName
    return leafs
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_ProcessNode
// Collection for process
type PerfMgmt_Enable_Statistics_ProcessNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_ProcessNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_ProcessNode_Nodes
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetFilter() yfilter.YFilter { return processNode.YFilter }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) SetFilter(yf yfilter.YFilter) { processNode.YFilter = yf }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetGoName(yname string) string {
    if yname == "node-all" { return "NodeAll" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetSegmentPath() string {
    return "process-node"
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-all" {
        return &processNode.NodeAll
    }
    if childYangName == "nodes" {
        return &processNode.Nodes
    }
    return nil
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-all"] = &processNode.NodeAll
    children["nodes"] = &processNode.Nodes
    return children
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetBundleName() string { return "cisco_ios_xr" }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetYangName() string { return "process-node" }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) SetParent(parent types.Entity) { processNode.parent = parent }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetParent() types.Entity { return processNode.parent }

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_ProcessNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_ProcessNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetParentYangName() string { return "process-node" }

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetParentYangName() string { return "process-node" }

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Statistics_DataRateInterface
// Statistics collection for generic-counters
type PerfMgmt_Enable_Statistics_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface"
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = dataRateInterface.TemplateName
    return leafs
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_MemoryNode
// Collection for memory
type PerfMgmt_Enable_Statistics_MemoryNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_MemoryNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_MemoryNode_Nodes
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetFilter() yfilter.YFilter { return memoryNode.YFilter }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) SetFilter(yf yfilter.YFilter) { memoryNode.YFilter = yf }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetGoName(yname string) string {
    if yname == "node-all" { return "NodeAll" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetSegmentPath() string {
    return "memory-node"
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-all" {
        return &memoryNode.NodeAll
    }
    if childYangName == "nodes" {
        return &memoryNode.Nodes
    }
    return nil
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-all"] = &memoryNode.NodeAll
    children["nodes"] = &memoryNode.Nodes
    return children
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetYangName() string { return "memory-node" }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) SetParent(parent types.Entity) { memoryNode.parent = parent }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetParent() types.Entity { return memoryNode.parent }

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_Statistics_MemoryNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_MemoryNode_NodeAll struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetFilter() yfilter.YFilter { return nodeAll.YFilter }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) SetFilter(yf yfilter.YFilter) { nodeAll.YFilter = yf }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetSegmentPath() string {
    return "node-all"
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = nodeAll.TemplateName
    return leafs
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetBundleName() string { return "cisco_ios_xr" }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetYangName() string { return "node-all" }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) SetParent(parent types.Entity) { nodeAll.parent = parent }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetParent() types.Entity { return nodeAll.parent }

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_Statistics_LdpMpls
// Collection for labels distribution protocol
type PerfMgmt_Enable_Statistics_LdpMpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetFilter() yfilter.YFilter { return ldpMpls.YFilter }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) SetFilter(yf yfilter.YFilter) { ldpMpls.YFilter = yf }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetSegmentPath() string {
    return "ldp-mpls"
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ldpMpls.TemplateName
    return leafs
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetYangName() string { return "ldp-mpls" }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) SetParent(parent types.Entity) { ldpMpls.parent = parent }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetParent() types.Entity { return ldpMpls.parent }

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetParentYangName() string { return "statistics" }

// PerfMgmt_Enable_MonitorEnable
// Start data collection for a monitored instance
type PerfMgmt_Enable_MonitorEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitoring for LDP.
    LdpMpls PerfMgmt_Enable_MonitorEnable_LdpMpls

    // Monitor OSPF v3 Protocol.
    Ospfv3Protocol PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol

    // Monitoring for generic-counters.
    GenericCounters PerfMgmt_Enable_MonitorEnable_GenericCounters

    // Collection for a single process.
    Process PerfMgmt_Enable_MonitorEnable_Process

    // Monitoring for basic-counters.
    BasicCounters PerfMgmt_Enable_MonitorEnable_BasicCounters

    // Collection for memory.
    Memory PerfMgmt_Enable_MonitorEnable_Memory

    // Monitor OSPF v2 Protocol.
    Ospfv2Protocol PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol

    // Collection for CPU.
    Cpu PerfMgmt_Enable_MonitorEnable_Cpu

    // Monitor BGP protocol.
    Bgp PerfMgmt_Enable_MonitorEnable_Bgp

    // Monitoring for data-rates.
    DataRates PerfMgmt_Enable_MonitorEnable_DataRates
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetFilter() yfilter.YFilter { return monitorEnable.YFilter }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) SetFilter(yf yfilter.YFilter) { monitorEnable.YFilter = yf }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetGoName(yname string) string {
    if yname == "ldp-mpls" { return "LdpMpls" }
    if yname == "ospfv3-protocol" { return "Ospfv3Protocol" }
    if yname == "generic-counters" { return "GenericCounters" }
    if yname == "process" { return "Process" }
    if yname == "basic-counters" { return "BasicCounters" }
    if yname == "memory" { return "Memory" }
    if yname == "ospfv2-protocol" { return "Ospfv2Protocol" }
    if yname == "cpu" { return "Cpu" }
    if yname == "bgp" { return "Bgp" }
    if yname == "data-rates" { return "DataRates" }
    return ""
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetSegmentPath() string {
    return "monitor-enable"
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-mpls" {
        return &monitorEnable.LdpMpls
    }
    if childYangName == "ospfv3-protocol" {
        return &monitorEnable.Ospfv3Protocol
    }
    if childYangName == "generic-counters" {
        return &monitorEnable.GenericCounters
    }
    if childYangName == "process" {
        return &monitorEnable.Process
    }
    if childYangName == "basic-counters" {
        return &monitorEnable.BasicCounters
    }
    if childYangName == "memory" {
        return &monitorEnable.Memory
    }
    if childYangName == "ospfv2-protocol" {
        return &monitorEnable.Ospfv2Protocol
    }
    if childYangName == "cpu" {
        return &monitorEnable.Cpu
    }
    if childYangName == "bgp" {
        return &monitorEnable.Bgp
    }
    if childYangName == "data-rates" {
        return &monitorEnable.DataRates
    }
    return nil
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-mpls"] = &monitorEnable.LdpMpls
    children["ospfv3-protocol"] = &monitorEnable.Ospfv3Protocol
    children["generic-counters"] = &monitorEnable.GenericCounters
    children["process"] = &monitorEnable.Process
    children["basic-counters"] = &monitorEnable.BasicCounters
    children["memory"] = &monitorEnable.Memory
    children["ospfv2-protocol"] = &monitorEnable.Ospfv2Protocol
    children["cpu"] = &monitorEnable.Cpu
    children["bgp"] = &monitorEnable.Bgp
    children["data-rates"] = &monitorEnable.DataRates
    return children
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetBundleName() string { return "cisco_ios_xr" }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetYangName() string { return "monitor-enable" }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) SetParent(parent types.Entity) { monitorEnable.parent = parent }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetParent() types.Entity { return monitorEnable.parent }

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetParentYangName() string { return "enable" }

// PerfMgmt_Enable_MonitorEnable_LdpMpls
// Monitoring for LDP
type PerfMgmt_Enable_MonitorEnable_LdpMpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP session specification.
    Sessions PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetFilter() yfilter.YFilter { return ldpMpls.YFilter }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) SetFilter(yf yfilter.YFilter) { ldpMpls.YFilter = yf }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetSegmentPath() string {
    return "ldp-mpls"
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &ldpMpls.Sessions
    }
    return nil
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &ldpMpls.Sessions
    return children
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetYangName() string { return "ldp-mpls" }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) SetParent(parent types.Entity) { ldpMpls.parent = parent }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetParent() types.Entity { return ldpMpls.parent }

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions
// LDP session specification
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of the LDP Session. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session.
    Session []PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetYangName() string { return "sessions" }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetParentYangName() string { return "ldp-mpls" }

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
// IP address of the LDP Session
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the LDP Session. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Session interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session='" + fmt.Sprintf("%v", session.Session) + "']"
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session"] = session.Session
    leafs["template-name"] = session.TemplateName
    return leafs
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetYangName() string { return "session" }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetParentYangName() string { return "sessions" }

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol
// Monitor OSPF v3 Protocol
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor an instance.
    OspfInstances PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetFilter() yfilter.YFilter { return ospfv3Protocol.YFilter }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) SetFilter(yf yfilter.YFilter) { ospfv3Protocol.YFilter = yf }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetGoName(yname string) string {
    if yname == "ospf-instances" { return "OspfInstances" }
    return ""
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetSegmentPath() string {
    return "ospfv3-protocol"
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf-instances" {
        return &ospfv3Protocol.OspfInstances
    }
    return nil
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospf-instances"] = &ospfv3Protocol.OspfInstances
    return children
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetYangName() string { return "ospfv3-protocol" }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) SetParent(parent types.Entity) { ospfv3Protocol.parent = parent }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetParent() types.Entity { return ospfv3Protocol.parent }

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance.
    OspfInstance []PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetFilter() yfilter.YFilter { return ospfInstances.YFilter }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) SetFilter(yf yfilter.YFilter) { ospfInstances.YFilter = yf }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetGoName(yname string) string {
    if yname == "ospf-instance" { return "OspfInstance" }
    return ""
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetSegmentPath() string {
    return "ospf-instances"
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf-instance" {
        for _, c := range ospfInstances.OspfInstance {
            if ospfInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance{}
        ospfInstances.OspfInstance = append(ospfInstances.OspfInstance, child)
        return &ospfInstances.OspfInstance[len(ospfInstances.OspfInstance)-1]
    }
    return nil
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfInstances.OspfInstance {
        children[ospfInstances.OspfInstance[i].GetSegmentPath()] = &ospfInstances.OspfInstance[i]
    }
    return children
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetYangName() string { return "ospf-instances" }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) SetParent(parent types.Entity) { ospfInstances.parent = parent }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetParent() types.Entity { return ospfInstances.parent }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetParentYangName() string { return "ospfv3-protocol" }

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetFilter() yfilter.YFilter { return ospfInstance.YFilter }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) SetFilter(yf yfilter.YFilter) { ospfInstance.YFilter = yf }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetSegmentPath() string {
    return "ospf-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfInstance.InstanceName) + "']"
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfInstance.InstanceName
    leafs["template-name"] = ospfInstance.TemplateName
    return leafs
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetYangName() string { return "ospf-instance" }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) SetParent(parent types.Entity) { ospfInstance.parent = parent }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetParent() types.Entity { return ospfInstance.parent }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetParentYangName() string { return "ospf-instances" }

// PerfMgmt_Enable_MonitorEnable_GenericCounters
// Monitoring for generic-counters
type PerfMgmt_Enable_MonitorEnable_GenericCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetFilter() yfilter.YFilter { return genericCounters.YFilter }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) SetFilter(yf yfilter.YFilter) { genericCounters.YFilter = yf }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetSegmentPath() string {
    return "generic-counters"
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &genericCounters.Interfaces
    }
    return nil
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &genericCounters.Interfaces
    return children
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetYangName() string { return "generic-counters" }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) SetParent(parent types.Entity) { genericCounters.parent = parent }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetParent() types.Entity { return genericCounters.parent }

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface.
    Interface []PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetParentYangName() string { return "generic-counters" }

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["template-name"] = self.TemplateName
    return leafs
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// PerfMgmt_Enable_MonitorEnable_Process
// Collection for a single process
type PerfMgmt_Enable_MonitorEnable_Process struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    ProcessNodes PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetFilter() yfilter.YFilter { return process.YFilter }

func (process *PerfMgmt_Enable_MonitorEnable_Process) SetFilter(yf yfilter.YFilter) { process.YFilter = yf }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetGoName(yname string) string {
    if yname == "process-nodes" { return "ProcessNodes" }
    return ""
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetSegmentPath() string {
    return "process"
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-nodes" {
        return &process.ProcessNodes
    }
    return nil
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process-nodes"] = &process.ProcessNodes
    return children
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetBundleName() string { return "cisco_ios_xr" }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetYangName() string { return "process" }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (process *PerfMgmt_Enable_MonitorEnable_Process) SetParent(parent types.Entity) { process.parent = parent }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetParent() types.Entity { return process.parent }

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode.
    ProcessNode []PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetFilter() yfilter.YFilter { return processNodes.YFilter }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) SetFilter(yf yfilter.YFilter) { processNodes.YFilter = yf }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetGoName(yname string) string {
    if yname == "process-node" { return "ProcessNode" }
    return ""
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetSegmentPath() string {
    return "process-nodes"
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-node" {
        for _, c := range processNodes.ProcessNode {
            if processNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode{}
        processNodes.ProcessNode = append(processNodes.ProcessNode, child)
        return &processNodes.ProcessNode[len(processNodes.ProcessNode)-1]
    }
    return nil
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processNodes.ProcessNode {
        children[processNodes.ProcessNode[i].GetSegmentPath()] = &processNodes.ProcessNode[i]
    }
    return children
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetBundleName() string { return "cisco_ios_xr" }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetYangName() string { return "process-nodes" }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) SetParent(parent types.Entity) { processNodes.parent = parent }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetParent() types.Entity { return processNodes.parent }

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetParentYangName() string { return "process" }

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
// Node instance
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Process ID specification.
    Pids PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetFilter() yfilter.YFilter { return processNode.YFilter }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) SetFilter(yf yfilter.YFilter) { processNode.YFilter = yf }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "pids" { return "Pids" }
    return ""
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetSegmentPath() string {
    return "process-node" + "[node-id='" + fmt.Sprintf("%v", processNode.NodeId) + "']"
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pids" {
        return &processNode.Pids
    }
    return nil
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pids"] = &processNode.Pids
    return children
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = processNode.NodeId
    return leafs
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetBundleName() string { return "cisco_ios_xr" }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetYangName() string { return "process-node" }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) SetParent(parent types.Entity) { processNode.parent = parent }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetParent() types.Entity { return processNode.parent }

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetParentYangName() string { return "process-nodes" }

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
// Process ID specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify an existing template for data collection. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid.
    Pid []PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetFilter() yfilter.YFilter { return pids.YFilter }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) SetFilter(yf yfilter.YFilter) { pids.YFilter = yf }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetGoName(yname string) string {
    if yname == "pid" { return "Pid" }
    return ""
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetSegmentPath() string {
    return "pids"
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pid" {
        for _, c := range pids.Pid {
            if pids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid{}
        pids.Pid = append(pids.Pid, child)
        return &pids.Pid[len(pids.Pid)-1]
    }
    return nil
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pids.Pid {
        children[pids.Pid[i].GetSegmentPath()] = &pids.Pid[i]
    }
    return children
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetBundleName() string { return "cisco_ios_xr" }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetYangName() string { return "pids" }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) SetParent(parent types.Entity) { pids.parent = parent }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetParent() types.Entity { return pids.parent }

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetParentYangName() string { return "process-node" }

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid
// Specify an existing template for data
// collection
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify Process ID. The type is interface{} with
    // range: 0..4294967295.
    Pid interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetFilter() yfilter.YFilter { return pid.YFilter }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) SetFilter(yf yfilter.YFilter) { pid.YFilter = yf }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetGoName(yname string) string {
    if yname == "pid" { return "Pid" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetSegmentPath() string {
    return "pid" + "[pid='" + fmt.Sprintf("%v", pid.Pid) + "']"
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pid"] = pid.Pid
    leafs["template-name"] = pid.TemplateName
    return leafs
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetBundleName() string { return "cisco_ios_xr" }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetYangName() string { return "pid" }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) SetParent(parent types.Entity) { pid.parent = parent }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetParent() types.Entity { return pid.parent }

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetParentYangName() string { return "pids" }

// PerfMgmt_Enable_MonitorEnable_BasicCounters
// Monitoring for basic-counters
type PerfMgmt_Enable_MonitorEnable_BasicCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetFilter() yfilter.YFilter { return basicCounters.YFilter }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) SetFilter(yf yfilter.YFilter) { basicCounters.YFilter = yf }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetSegmentPath() string {
    return "basic-counters"
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &basicCounters.Interfaces
    }
    return nil
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &basicCounters.Interfaces
    return children
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetYangName() string { return "basic-counters" }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) SetParent(parent types.Entity) { basicCounters.parent = parent }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetParent() types.Entity { return basicCounters.parent }

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface.
    Interface []PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetParentYangName() string { return "basic-counters" }

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["template-name"] = self.TemplateName
    return leafs
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// PerfMgmt_Enable_MonitorEnable_Memory
// Collection for memory
type PerfMgmt_Enable_MonitorEnable_Memory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_MonitorEnable_Memory_Nodes
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetFilter() yfilter.YFilter { return memory.YFilter }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) SetFilter(yf yfilter.YFilter) { memory.YFilter = yf }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetSegmentPath() string {
    return "memory"
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &memory.Nodes
    }
    return nil
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &memory.Nodes
    return children
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetBundleName() string { return "cisco_ios_xr" }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetYangName() string { return "memory" }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) SetParent(parent types.Entity) { memory.parent = parent }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetParent() types.Entity { return memory.parent }

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node.
    Node []PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetParentYangName() string { return "memory" }

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol
// Monitor OSPF v2 Protocol
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor an instance.
    OspfInstances PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetFilter() yfilter.YFilter { return ospfv2Protocol.YFilter }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) SetFilter(yf yfilter.YFilter) { ospfv2Protocol.YFilter = yf }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetGoName(yname string) string {
    if yname == "ospf-instances" { return "OspfInstances" }
    return ""
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetSegmentPath() string {
    return "ospfv2-protocol"
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf-instances" {
        return &ospfv2Protocol.OspfInstances
    }
    return nil
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospf-instances"] = &ospfv2Protocol.OspfInstances
    return children
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetYangName() string { return "ospfv2-protocol" }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) SetParent(parent types.Entity) { ospfv2Protocol.parent = parent }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetParent() types.Entity { return ospfv2Protocol.parent }

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance.
    OspfInstance []PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetFilter() yfilter.YFilter { return ospfInstances.YFilter }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) SetFilter(yf yfilter.YFilter) { ospfInstances.YFilter = yf }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetGoName(yname string) string {
    if yname == "ospf-instance" { return "OspfInstance" }
    return ""
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetSegmentPath() string {
    return "ospf-instances"
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospf-instance" {
        for _, c := range ospfInstances.OspfInstance {
            if ospfInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance{}
        ospfInstances.OspfInstance = append(ospfInstances.OspfInstance, child)
        return &ospfInstances.OspfInstance[len(ospfInstances.OspfInstance)-1]
    }
    return nil
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfInstances.OspfInstance {
        children[ospfInstances.OspfInstance[i].GetSegmentPath()] = &ospfInstances.OspfInstance[i]
    }
    return children
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetBundleName() string { return "cisco_ios_xr" }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetYangName() string { return "ospf-instances" }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) SetParent(parent types.Entity) { ospfInstances.parent = parent }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetParent() types.Entity { return ospfInstances.parent }

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetParentYangName() string { return "ospfv2-protocol" }

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetFilter() yfilter.YFilter { return ospfInstance.YFilter }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) SetFilter(yf yfilter.YFilter) { ospfInstance.YFilter = yf }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetSegmentPath() string {
    return "ospf-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfInstance.InstanceName) + "']"
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfInstance.InstanceName
    leafs["template-name"] = ospfInstance.TemplateName
    return leafs
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetBundleName() string { return "cisco_ios_xr" }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetYangName() string { return "ospf-instance" }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) SetParent(parent types.Entity) { ospfInstance.parent = parent }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetParent() types.Entity { return ospfInstance.parent }

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetParentYangName() string { return "ospf-instances" }

// PerfMgmt_Enable_MonitorEnable_Cpu
// Collection for CPU
type PerfMgmt_Enable_MonitorEnable_Cpu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_MonitorEnable_Cpu_Nodes
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetFilter() yfilter.YFilter { return cpu.YFilter }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) SetFilter(yf yfilter.YFilter) { cpu.YFilter = yf }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetSegmentPath() string {
    return "cpu"
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &cpu.Nodes
    }
    return nil
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &cpu.Nodes
    return children
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetBundleName() string { return "cisco_ios_xr" }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetYangName() string { return "cpu" }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) SetParent(parent types.Entity) { cpu.parent = parent }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetParent() types.Entity { return cpu.parent }

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node.
    Node []PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetYangName() string { return "nodes" }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetParentYangName() string { return "cpu" }

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-id" { return "NodeId" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-id"] = node.NodeId
    leafs["template-name"] = node.TemplateName
    return leafs
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetYangName() string { return "node" }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetParentYangName() string { return "nodes" }

// PerfMgmt_Enable_MonitorEnable_Bgp
// Monitor BGP protocol
type PerfMgmt_Enable_MonitorEnable_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor BGP protocol for a BGP peer.
    Neighbors PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetGoName(yname string) string {
    if yname == "neighbors" { return "Neighbors" }
    return ""
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbors" {
        return &bgp.Neighbors
    }
    return nil
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["neighbors"] = &bgp.Neighbors
    return children
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors
// Monitor BGP protocol for a BGP peer
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Neighbor being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor.
    Neighbor []PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetBundleName() string { return "cisco_ios_xr" }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetParentYangName() string { return "bgp" }

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
// Neighbor being monitored
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the Neighbor. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[peer-address='" + fmt.Sprintf("%v", neighbor.PeerAddress) + "']"
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = neighbor.PeerAddress
    leafs["template-name"] = neighbor.TemplateName
    return leafs
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// PerfMgmt_Enable_MonitorEnable_DataRates
// Monitoring for data-rates
type PerfMgmt_Enable_MonitorEnable_DataRates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetFilter() yfilter.YFilter { return dataRates.YFilter }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) SetFilter(yf yfilter.YFilter) { dataRates.YFilter = yf }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetSegmentPath() string {
    return "data-rates"
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &dataRates.Interfaces
    }
    return nil
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &dataRates.Interfaces
    return children
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetBundleName() string { return "cisco_ios_xr" }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetYangName() string { return "data-rates" }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) SetParent(parent types.Entity) { dataRates.parent = parent }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetParent() types.Entity { return dataRates.parent }

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetParentYangName() string { return "monitor-enable" }

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface.
    Interface []PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetParentYangName() string { return "data-rates" }

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "template-name" { return "TemplateName" }
    return ""
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["template-name"] = self.TemplateName
    return leafs
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// PerfMgmt_RegExpGroups
// Configure regular expression group
type PerfMgmt_RegExpGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify regular expression group name. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup.
    RegExpGroup []PerfMgmt_RegExpGroups_RegExpGroup
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetFilter() yfilter.YFilter { return regExpGroups.YFilter }

func (regExpGroups *PerfMgmt_RegExpGroups) SetFilter(yf yfilter.YFilter) { regExpGroups.YFilter = yf }

func (regExpGroups *PerfMgmt_RegExpGroups) GetGoName(yname string) string {
    if yname == "reg-exp-group" { return "RegExpGroup" }
    return ""
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetSegmentPath() string {
    return "reg-exp-groups"
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reg-exp-group" {
        for _, c := range regExpGroups.RegExpGroup {
            if regExpGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_RegExpGroups_RegExpGroup{}
        regExpGroups.RegExpGroup = append(regExpGroups.RegExpGroup, child)
        return &regExpGroups.RegExpGroup[len(regExpGroups.RegExpGroup)-1]
    }
    return nil
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range regExpGroups.RegExpGroup {
        children[regExpGroups.RegExpGroup[i].GetSegmentPath()] = &regExpGroups.RegExpGroup[i]
    }
    return children
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetBundleName() string { return "cisco_ios_xr" }

func (regExpGroups *PerfMgmt_RegExpGroups) GetYangName() string { return "reg-exp-groups" }

func (regExpGroups *PerfMgmt_RegExpGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regExpGroups *PerfMgmt_RegExpGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regExpGroups *PerfMgmt_RegExpGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regExpGroups *PerfMgmt_RegExpGroups) SetParent(parent types.Entity) { regExpGroups.parent = parent }

func (regExpGroups *PerfMgmt_RegExpGroups) GetParent() types.Entity { return regExpGroups.parent }

func (regExpGroups *PerfMgmt_RegExpGroups) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_RegExpGroups_RegExpGroup
// Specify regular expression group name
type PerfMgmt_RegExpGroups_RegExpGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Regular expression group name. The type is string
    // with length: 1..32.
    RegExpGroupName interface{}

    // Configure regular expression.
    RegExps PerfMgmt_RegExpGroups_RegExpGroup_RegExps
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetFilter() yfilter.YFilter { return regExpGroup.YFilter }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) SetFilter(yf yfilter.YFilter) { regExpGroup.YFilter = yf }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetGoName(yname string) string {
    if yname == "reg-exp-group-name" { return "RegExpGroupName" }
    if yname == "reg-exps" { return "RegExps" }
    return ""
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetSegmentPath() string {
    return "reg-exp-group" + "[reg-exp-group-name='" + fmt.Sprintf("%v", regExpGroup.RegExpGroupName) + "']"
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reg-exps" {
        return &regExpGroup.RegExps
    }
    return nil
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reg-exps"] = &regExpGroup.RegExps
    return children
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg-exp-group-name"] = regExpGroup.RegExpGroupName
    return leafs
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetBundleName() string { return "cisco_ios_xr" }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetYangName() string { return "reg-exp-group" }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) SetParent(parent types.Entity) { regExpGroup.parent = parent }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetParent() types.Entity { return regExpGroup.parent }

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetParentYangName() string { return "reg-exp-groups" }

// PerfMgmt_RegExpGroups_RegExpGroup_RegExps
// Configure regular expression
type PerfMgmt_RegExpGroups_RegExpGroup_RegExps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify regular expression index number. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp.
    RegExp []PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetFilter() yfilter.YFilter { return regExps.YFilter }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) SetFilter(yf yfilter.YFilter) { regExps.YFilter = yf }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetGoName(yname string) string {
    if yname == "reg-exp" { return "RegExp" }
    return ""
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetSegmentPath() string {
    return "reg-exps"
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reg-exp" {
        for _, c := range regExps.RegExp {
            if regExps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp{}
        regExps.RegExp = append(regExps.RegExp, child)
        return &regExps.RegExp[len(regExps.RegExp)-1]
    }
    return nil
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range regExps.RegExp {
        children[regExps.RegExp[i].GetSegmentPath()] = &regExps.RegExp[i]
    }
    return children
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetBundleName() string { return "cisco_ios_xr" }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetYangName() string { return "reg-exps" }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) SetParent(parent types.Entity) { regExps.parent = parent }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetParent() types.Entity { return regExps.parent }

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetParentYangName() string { return "reg-exp-group" }

// PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp
// Specify regular expression index number
type PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Regular expression index number. The type is
    // interface{} with range: 1..100.
    RegExpIndex interface{}

    // Regular expression string to match. The type is string with length: 1..128.
    // This attribute is mandatory.
    RegExpString interface{}
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetFilter() yfilter.YFilter { return regExp.YFilter }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) SetFilter(yf yfilter.YFilter) { regExp.YFilter = yf }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetGoName(yname string) string {
    if yname == "reg-exp-index" { return "RegExpIndex" }
    if yname == "reg-exp-string" { return "RegExpString" }
    return ""
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetSegmentPath() string {
    return "reg-exp" + "[reg-exp-index='" + fmt.Sprintf("%v", regExp.RegExpIndex) + "']"
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reg-exp-index"] = regExp.RegExpIndex
    leafs["reg-exp-string"] = regExp.RegExpString
    return leafs
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetBundleName() string { return "cisco_ios_xr" }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetYangName() string { return "reg-exp" }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) SetParent(parent types.Entity) { regExp.parent = parent }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetParent() types.Entity { return regExp.parent }

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetParentYangName() string { return "reg-exps" }

// PerfMgmt_Threshold
// Container for threshold templates
type PerfMgmt_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold configuration.
    GenericCounterInterface PerfMgmt_Threshold_GenericCounterInterface

    // MPLS LDP threshold configuration.
    LdpMpls PerfMgmt_Threshold_LdpMpls

    // Interface Basic Counter threshold configuration.
    BasicCounterInterface PerfMgmt_Threshold_BasicCounterInterface

    // BGP threshold configuration.
    Bgp PerfMgmt_Threshold_Bgp

    // OSPF v2 Protocol threshold configuration.
    Ospfv2Protocol PerfMgmt_Threshold_Ospfv2Protocol

    // Node CPU threshold configuration.
    CpuNode PerfMgmt_Threshold_CpuNode

    // Interface Data Rates threshold configuration.
    DataRateInterface PerfMgmt_Threshold_DataRateInterface

    // Node Process threshold configuration.
    ProcessNode PerfMgmt_Threshold_ProcessNode

    // Node Memory threshold configuration.
    MemoryNode PerfMgmt_Threshold_MemoryNode

    // OSPF v2 Protocol threshold configuration.
    Ospfv3Protocol PerfMgmt_Threshold_Ospfv3Protocol
}

func (threshold *PerfMgmt_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *PerfMgmt_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *PerfMgmt_Threshold) GetGoName(yname string) string {
    if yname == "generic-counter-interface" { return "GenericCounterInterface" }
    if yname == "ldp-mpls" { return "LdpMpls" }
    if yname == "basic-counter-interface" { return "BasicCounterInterface" }
    if yname == "bgp" { return "Bgp" }
    if yname == "ospfv2-protocol" { return "Ospfv2Protocol" }
    if yname == "cpu-node" { return "CpuNode" }
    if yname == "data-rate-interface" { return "DataRateInterface" }
    if yname == "process-node" { return "ProcessNode" }
    if yname == "memory-node" { return "MemoryNode" }
    if yname == "ospfv3-protocol" { return "Ospfv3Protocol" }
    return ""
}

func (threshold *PerfMgmt_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *PerfMgmt_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface" {
        return &threshold.GenericCounterInterface
    }
    if childYangName == "ldp-mpls" {
        return &threshold.LdpMpls
    }
    if childYangName == "basic-counter-interface" {
        return &threshold.BasicCounterInterface
    }
    if childYangName == "bgp" {
        return &threshold.Bgp
    }
    if childYangName == "ospfv2-protocol" {
        return &threshold.Ospfv2Protocol
    }
    if childYangName == "cpu-node" {
        return &threshold.CpuNode
    }
    if childYangName == "data-rate-interface" {
        return &threshold.DataRateInterface
    }
    if childYangName == "process-node" {
        return &threshold.ProcessNode
    }
    if childYangName == "memory-node" {
        return &threshold.MemoryNode
    }
    if childYangName == "ospfv3-protocol" {
        return &threshold.Ospfv3Protocol
    }
    return nil
}

func (threshold *PerfMgmt_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interface"] = &threshold.GenericCounterInterface
    children["ldp-mpls"] = &threshold.LdpMpls
    children["basic-counter-interface"] = &threshold.BasicCounterInterface
    children["bgp"] = &threshold.Bgp
    children["ospfv2-protocol"] = &threshold.Ospfv2Protocol
    children["cpu-node"] = &threshold.CpuNode
    children["data-rate-interface"] = &threshold.DataRateInterface
    children["process-node"] = &threshold.ProcessNode
    children["memory-node"] = &threshold.MemoryNode
    children["ospfv3-protocol"] = &threshold.Ospfv3Protocol
    return children
}

func (threshold *PerfMgmt_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (threshold *PerfMgmt_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *PerfMgmt_Threshold) GetYangName() string { return "threshold" }

func (threshold *PerfMgmt_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *PerfMgmt_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *PerfMgmt_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *PerfMgmt_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *PerfMgmt_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *PerfMgmt_Threshold) GetParentYangName() string { return "perf-mgmt" }

// PerfMgmt_Threshold_GenericCounterInterface
// Interface Generic Counter threshold
// configuration
type PerfMgmt_Threshold_GenericCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold templates.
    GenericCounterInterfaceTemplates PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetFilter() yfilter.YFilter { return genericCounterInterface.YFilter }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) SetFilter(yf yfilter.YFilter) { genericCounterInterface.YFilter = yf }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetGoName(yname string) string {
    if yname == "generic-counter-interface-templates" { return "GenericCounterInterfaceTemplates" }
    return ""
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetSegmentPath() string {
    return "generic-counter-interface"
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface-templates" {
        return &genericCounterInterface.GenericCounterInterfaceTemplates
    }
    return nil
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["generic-counter-interface-templates"] = &genericCounterInterface.GenericCounterInterfaceTemplates
    return children
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetYangName() string { return "generic-counter-interface" }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) SetParent(parent types.Entity) { genericCounterInterface.parent = parent }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetParent() types.Entity { return genericCounterInterface.parent }

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates
// Interface Generic Counter threshold templates
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate.
    GenericCounterInterfaceTemplate []PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetFilter() yfilter.YFilter { return genericCounterInterfaceTemplates.YFilter }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) SetFilter(yf yfilter.YFilter) { genericCounterInterfaceTemplates.YFilter = yf }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetGoName(yname string) string {
    if yname == "generic-counter-interface-template" { return "GenericCounterInterfaceTemplate" }
    return ""
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetSegmentPath() string {
    return "generic-counter-interface-templates"
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "generic-counter-interface-template" {
        for _, c := range genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate {
            if genericCounterInterfaceTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate{}
        genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate = append(genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate, child)
        return &genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[len(genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate)-1]
    }
    return nil
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate {
        children[genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i].GetSegmentPath()] = &genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i]
    }
    return children
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetYangName() string { return "generic-counter-interface-templates" }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) SetParent(parent types.Entity) { genericCounterInterfaceTemplates.parent = parent }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetParent() types.Entity { return genericCounterInterfaceTemplates.parent }

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetParentYangName() string { return "generic-counter-interface" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
// Interface Generic Counter threshold template
// instance
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable instance filtering by VRF name regular expression . The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Number of inbound octets/bytes.
    InOctets PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets

    // Number of inbound unicast packets.
    InUcastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts

    // Number of outbound unicast packets.
    OutUcastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts

    // Number of outbound broadcast packets.
    OutBroadcastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts

    // Number of outbound multicast packets.
    OutMulticastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts

    // Number of inbound packets with overrun errors.
    InputOverrun PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun

    // Number of outbound octets/bytes.
    OutOctets PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets

    // Number of outbound packets with underrun errors.
    OutputUnderrun PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun

    // Number of inbound incorrect packets discarded.
    InputTotalErrors PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors

    // Number of outbound correct packets discarded.
    OutputTotalDrops PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops

    // Number of inbound packets discarded with incorrect CRC.
    InputCrc PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc

    // Number of inbound broadcast packets.
    InBroadcastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts

    // Number of inbound multicast packets.
    InMulticastPkts PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts

    // Number of outbound packets.
    OutPackets PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets

    // Number of outbound incorrect packets discarded.
    OutputTotalErrors PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors

    // Number of inbound packets.
    InPackets PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets

    // Number of inbound packets discarded with unknown protocol.
    InputUnknownProto PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto

    // Number of input queue drops.
    InputQueueDrops PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops

    // Number of inbound correct packets discarded.
    InputTotalDrops PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops

    // Number of inbound packets with framing errors.
    InputFrame PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetFilter() yfilter.YFilter { return genericCounterInterfaceTemplate.YFilter }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) SetFilter(yf yfilter.YFilter) { genericCounterInterfaceTemplate.YFilter = yf }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "in-ucast-pkts" { return "InUcastPkts" }
    if yname == "out-ucast-pkts" { return "OutUcastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "input-overrun" { return "InputOverrun" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "output-underrun" { return "OutputUnderrun" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "input-crc" { return "InputCrc" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "input-unknown-proto" { return "InputUnknownProto" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    if yname == "input-frame" { return "InputFrame" }
    return ""
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetSegmentPath() string {
    return "generic-counter-interface-template" + "[template-name='" + fmt.Sprintf("%v", genericCounterInterfaceTemplate.TemplateName) + "']"
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "in-octets" {
        return &genericCounterInterfaceTemplate.InOctets
    }
    if childYangName == "in-ucast-pkts" {
        return &genericCounterInterfaceTemplate.InUcastPkts
    }
    if childYangName == "out-ucast-pkts" {
        return &genericCounterInterfaceTemplate.OutUcastPkts
    }
    if childYangName == "out-broadcast-pkts" {
        return &genericCounterInterfaceTemplate.OutBroadcastPkts
    }
    if childYangName == "out-multicast-pkts" {
        return &genericCounterInterfaceTemplate.OutMulticastPkts
    }
    if childYangName == "input-overrun" {
        return &genericCounterInterfaceTemplate.InputOverrun
    }
    if childYangName == "out-octets" {
        return &genericCounterInterfaceTemplate.OutOctets
    }
    if childYangName == "output-underrun" {
        return &genericCounterInterfaceTemplate.OutputUnderrun
    }
    if childYangName == "input-total-errors" {
        return &genericCounterInterfaceTemplate.InputTotalErrors
    }
    if childYangName == "output-total-drops" {
        return &genericCounterInterfaceTemplate.OutputTotalDrops
    }
    if childYangName == "input-crc" {
        return &genericCounterInterfaceTemplate.InputCrc
    }
    if childYangName == "in-broadcast-pkts" {
        return &genericCounterInterfaceTemplate.InBroadcastPkts
    }
    if childYangName == "in-multicast-pkts" {
        return &genericCounterInterfaceTemplate.InMulticastPkts
    }
    if childYangName == "out-packets" {
        return &genericCounterInterfaceTemplate.OutPackets
    }
    if childYangName == "output-total-errors" {
        return &genericCounterInterfaceTemplate.OutputTotalErrors
    }
    if childYangName == "in-packets" {
        return &genericCounterInterfaceTemplate.InPackets
    }
    if childYangName == "input-unknown-proto" {
        return &genericCounterInterfaceTemplate.InputUnknownProto
    }
    if childYangName == "input-queue-drops" {
        return &genericCounterInterfaceTemplate.InputQueueDrops
    }
    if childYangName == "input-total-drops" {
        return &genericCounterInterfaceTemplate.InputTotalDrops
    }
    if childYangName == "input-frame" {
        return &genericCounterInterfaceTemplate.InputFrame
    }
    return nil
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["in-octets"] = &genericCounterInterfaceTemplate.InOctets
    children["in-ucast-pkts"] = &genericCounterInterfaceTemplate.InUcastPkts
    children["out-ucast-pkts"] = &genericCounterInterfaceTemplate.OutUcastPkts
    children["out-broadcast-pkts"] = &genericCounterInterfaceTemplate.OutBroadcastPkts
    children["out-multicast-pkts"] = &genericCounterInterfaceTemplate.OutMulticastPkts
    children["input-overrun"] = &genericCounterInterfaceTemplate.InputOverrun
    children["out-octets"] = &genericCounterInterfaceTemplate.OutOctets
    children["output-underrun"] = &genericCounterInterfaceTemplate.OutputUnderrun
    children["input-total-errors"] = &genericCounterInterfaceTemplate.InputTotalErrors
    children["output-total-drops"] = &genericCounterInterfaceTemplate.OutputTotalDrops
    children["input-crc"] = &genericCounterInterfaceTemplate.InputCrc
    children["in-broadcast-pkts"] = &genericCounterInterfaceTemplate.InBroadcastPkts
    children["in-multicast-pkts"] = &genericCounterInterfaceTemplate.InMulticastPkts
    children["out-packets"] = &genericCounterInterfaceTemplate.OutPackets
    children["output-total-errors"] = &genericCounterInterfaceTemplate.OutputTotalErrors
    children["in-packets"] = &genericCounterInterfaceTemplate.InPackets
    children["input-unknown-proto"] = &genericCounterInterfaceTemplate.InputUnknownProto
    children["input-queue-drops"] = &genericCounterInterfaceTemplate.InputQueueDrops
    children["input-total-drops"] = &genericCounterInterfaceTemplate.InputTotalDrops
    children["input-frame"] = &genericCounterInterfaceTemplate.InputFrame
    return children
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = genericCounterInterfaceTemplate.TemplateName
    leafs["sample-interval"] = genericCounterInterfaceTemplate.SampleInterval
    leafs["reg-exp-group"] = genericCounterInterfaceTemplate.RegExpGroup
    leafs["vrf-group"] = genericCounterInterfaceTemplate.VrfGroup
    return leafs
}

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetYangName() string { return "generic-counter-interface-template" }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) SetParent(parent types.Entity) { genericCounterInterfaceTemplate.parent = parent }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetParent() types.Entity { return genericCounterInterfaceTemplate.parent }

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetParentYangName() string { return "generic-counter-interface-templates" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetFilter() yfilter.YFilter { return inOctets.YFilter }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) SetFilter(yf yfilter.YFilter) { inOctets.YFilter = yf }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetSegmentPath() string {
    return "in-octets"
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inOctets.Operator
    leafs["value"] = inOctets.Value
    leafs["end-range-value"] = inOctets.EndRangeValue
    leafs["percent"] = inOctets.Percent
    leafs["rearm-type"] = inOctets.RearmType
    leafs["rearm-window"] = inOctets.RearmWindow
    return leafs
}

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetBundleName() string { return "cisco_ios_xr" }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetYangName() string { return "in-octets" }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) SetParent(parent types.Entity) { inOctets.parent = parent }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetParent() types.Entity { return inOctets.parent }

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts
// Number of inbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetFilter() yfilter.YFilter { return inUcastPkts.YFilter }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) SetFilter(yf yfilter.YFilter) { inUcastPkts.YFilter = yf }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetSegmentPath() string {
    return "in-ucast-pkts"
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inUcastPkts.Operator
    leafs["value"] = inUcastPkts.Value
    leafs["end-range-value"] = inUcastPkts.EndRangeValue
    leafs["percent"] = inUcastPkts.Percent
    leafs["rearm-type"] = inUcastPkts.RearmType
    leafs["rearm-window"] = inUcastPkts.RearmWindow
    return leafs
}

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetYangName() string { return "in-ucast-pkts" }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) SetParent(parent types.Entity) { inUcastPkts.parent = parent }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetParent() types.Entity { return inUcastPkts.parent }

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts
// Number of outbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetFilter() yfilter.YFilter { return outUcastPkts.YFilter }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) SetFilter(yf yfilter.YFilter) { outUcastPkts.YFilter = yf }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetSegmentPath() string {
    return "out-ucast-pkts"
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outUcastPkts.Operator
    leafs["value"] = outUcastPkts.Value
    leafs["end-range-value"] = outUcastPkts.EndRangeValue
    leafs["percent"] = outUcastPkts.Percent
    leafs["rearm-type"] = outUcastPkts.RearmType
    leafs["rearm-window"] = outUcastPkts.RearmWindow
    return leafs
}

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetYangName() string { return "out-ucast-pkts" }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) SetParent(parent types.Entity) { outUcastPkts.parent = parent }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetParent() types.Entity { return outUcastPkts.parent }

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts
// Number of outbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetFilter() yfilter.YFilter { return outBroadcastPkts.YFilter }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) SetFilter(yf yfilter.YFilter) { outBroadcastPkts.YFilter = yf }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetSegmentPath() string {
    return "out-broadcast-pkts"
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outBroadcastPkts.Operator
    leafs["value"] = outBroadcastPkts.Value
    leafs["end-range-value"] = outBroadcastPkts.EndRangeValue
    leafs["percent"] = outBroadcastPkts.Percent
    leafs["rearm-type"] = outBroadcastPkts.RearmType
    leafs["rearm-window"] = outBroadcastPkts.RearmWindow
    return leafs
}

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetYangName() string { return "out-broadcast-pkts" }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) SetParent(parent types.Entity) { outBroadcastPkts.parent = parent }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetParent() types.Entity { return outBroadcastPkts.parent }

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts
// Number of outbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetFilter() yfilter.YFilter { return outMulticastPkts.YFilter }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) SetFilter(yf yfilter.YFilter) { outMulticastPkts.YFilter = yf }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetSegmentPath() string {
    return "out-multicast-pkts"
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outMulticastPkts.Operator
    leafs["value"] = outMulticastPkts.Value
    leafs["end-range-value"] = outMulticastPkts.EndRangeValue
    leafs["percent"] = outMulticastPkts.Percent
    leafs["rearm-type"] = outMulticastPkts.RearmType
    leafs["rearm-window"] = outMulticastPkts.RearmWindow
    return leafs
}

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetYangName() string { return "out-multicast-pkts" }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) SetParent(parent types.Entity) { outMulticastPkts.parent = parent }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetParent() types.Entity { return outMulticastPkts.parent }

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun
// Number of inbound packets with overrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetFilter() yfilter.YFilter { return inputOverrun.YFilter }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) SetFilter(yf yfilter.YFilter) { inputOverrun.YFilter = yf }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetSegmentPath() string {
    return "input-overrun"
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputOverrun.Operator
    leafs["value"] = inputOverrun.Value
    leafs["end-range-value"] = inputOverrun.EndRangeValue
    leafs["percent"] = inputOverrun.Percent
    leafs["rearm-type"] = inputOverrun.RearmType
    leafs["rearm-window"] = inputOverrun.RearmWindow
    return leafs
}

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetBundleName() string { return "cisco_ios_xr" }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetYangName() string { return "input-overrun" }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) SetParent(parent types.Entity) { inputOverrun.parent = parent }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetParent() types.Entity { return inputOverrun.parent }

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetFilter() yfilter.YFilter { return outOctets.YFilter }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) SetFilter(yf yfilter.YFilter) { outOctets.YFilter = yf }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetSegmentPath() string {
    return "out-octets"
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outOctets.Operator
    leafs["value"] = outOctets.Value
    leafs["end-range-value"] = outOctets.EndRangeValue
    leafs["percent"] = outOctets.Percent
    leafs["rearm-type"] = outOctets.RearmType
    leafs["rearm-window"] = outOctets.RearmWindow
    return leafs
}

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetBundleName() string { return "cisco_ios_xr" }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetYangName() string { return "out-octets" }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) SetParent(parent types.Entity) { outOctets.parent = parent }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetParent() types.Entity { return outOctets.parent }

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun
// Number of outbound packets with underrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetFilter() yfilter.YFilter { return outputUnderrun.YFilter }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) SetFilter(yf yfilter.YFilter) { outputUnderrun.YFilter = yf }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetSegmentPath() string {
    return "output-underrun"
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputUnderrun.Operator
    leafs["value"] = outputUnderrun.Value
    leafs["end-range-value"] = outputUnderrun.EndRangeValue
    leafs["percent"] = outputUnderrun.Percent
    leafs["rearm-type"] = outputUnderrun.RearmType
    leafs["rearm-window"] = outputUnderrun.RearmWindow
    return leafs
}

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetBundleName() string { return "cisco_ios_xr" }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetYangName() string { return "output-underrun" }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) SetParent(parent types.Entity) { outputUnderrun.parent = parent }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetParent() types.Entity { return outputUnderrun.parent }

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetFilter() yfilter.YFilter { return inputTotalErrors.YFilter }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) SetFilter(yf yfilter.YFilter) { inputTotalErrors.YFilter = yf }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetSegmentPath() string {
    return "input-total-errors"
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputTotalErrors.Operator
    leafs["value"] = inputTotalErrors.Value
    leafs["end-range-value"] = inputTotalErrors.EndRangeValue
    leafs["percent"] = inputTotalErrors.Percent
    leafs["rearm-type"] = inputTotalErrors.RearmType
    leafs["rearm-window"] = inputTotalErrors.RearmWindow
    return leafs
}

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetBundleName() string { return "cisco_ios_xr" }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetYangName() string { return "input-total-errors" }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) SetParent(parent types.Entity) { inputTotalErrors.parent = parent }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetParent() types.Entity { return inputTotalErrors.parent }

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetFilter() yfilter.YFilter { return outputTotalDrops.YFilter }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) SetFilter(yf yfilter.YFilter) { outputTotalDrops.YFilter = yf }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetSegmentPath() string {
    return "output-total-drops"
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputTotalDrops.Operator
    leafs["value"] = outputTotalDrops.Value
    leafs["end-range-value"] = outputTotalDrops.EndRangeValue
    leafs["percent"] = outputTotalDrops.Percent
    leafs["rearm-type"] = outputTotalDrops.RearmType
    leafs["rearm-window"] = outputTotalDrops.RearmWindow
    return leafs
}

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetBundleName() string { return "cisco_ios_xr" }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetYangName() string { return "output-total-drops" }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) SetParent(parent types.Entity) { outputTotalDrops.parent = parent }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetParent() types.Entity { return outputTotalDrops.parent }

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc
// Number of inbound packets discarded with
// incorrect CRC
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetFilter() yfilter.YFilter { return inputCrc.YFilter }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) SetFilter(yf yfilter.YFilter) { inputCrc.YFilter = yf }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetSegmentPath() string {
    return "input-crc"
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputCrc.Operator
    leafs["value"] = inputCrc.Value
    leafs["end-range-value"] = inputCrc.EndRangeValue
    leafs["percent"] = inputCrc.Percent
    leafs["rearm-type"] = inputCrc.RearmType
    leafs["rearm-window"] = inputCrc.RearmWindow
    return leafs
}

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetBundleName() string { return "cisco_ios_xr" }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetYangName() string { return "input-crc" }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) SetParent(parent types.Entity) { inputCrc.parent = parent }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetParent() types.Entity { return inputCrc.parent }

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts
// Number of inbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetFilter() yfilter.YFilter { return inBroadcastPkts.YFilter }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) SetFilter(yf yfilter.YFilter) { inBroadcastPkts.YFilter = yf }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetSegmentPath() string {
    return "in-broadcast-pkts"
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inBroadcastPkts.Operator
    leafs["value"] = inBroadcastPkts.Value
    leafs["end-range-value"] = inBroadcastPkts.EndRangeValue
    leafs["percent"] = inBroadcastPkts.Percent
    leafs["rearm-type"] = inBroadcastPkts.RearmType
    leafs["rearm-window"] = inBroadcastPkts.RearmWindow
    return leafs
}

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetYangName() string { return "in-broadcast-pkts" }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) SetParent(parent types.Entity) { inBroadcastPkts.parent = parent }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetParent() types.Entity { return inBroadcastPkts.parent }

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts
// Number of inbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetFilter() yfilter.YFilter { return inMulticastPkts.YFilter }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) SetFilter(yf yfilter.YFilter) { inMulticastPkts.YFilter = yf }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetSegmentPath() string {
    return "in-multicast-pkts"
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inMulticastPkts.Operator
    leafs["value"] = inMulticastPkts.Value
    leafs["end-range-value"] = inMulticastPkts.EndRangeValue
    leafs["percent"] = inMulticastPkts.Percent
    leafs["rearm-type"] = inMulticastPkts.RearmType
    leafs["rearm-window"] = inMulticastPkts.RearmWindow
    return leafs
}

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetBundleName() string { return "cisco_ios_xr" }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetYangName() string { return "in-multicast-pkts" }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) SetParent(parent types.Entity) { inMulticastPkts.parent = parent }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetParent() types.Entity { return inMulticastPkts.parent }

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetFilter() yfilter.YFilter { return outPackets.YFilter }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) SetFilter(yf yfilter.YFilter) { outPackets.YFilter = yf }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetSegmentPath() string {
    return "out-packets"
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outPackets.Operator
    leafs["value"] = outPackets.Value
    leafs["end-range-value"] = outPackets.EndRangeValue
    leafs["percent"] = outPackets.Percent
    leafs["rearm-type"] = outPackets.RearmType
    leafs["rearm-window"] = outPackets.RearmWindow
    return leafs
}

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetYangName() string { return "out-packets" }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) SetParent(parent types.Entity) { outPackets.parent = parent }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetParent() types.Entity { return outPackets.parent }

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetFilter() yfilter.YFilter { return outputTotalErrors.YFilter }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) SetFilter(yf yfilter.YFilter) { outputTotalErrors.YFilter = yf }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetSegmentPath() string {
    return "output-total-errors"
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputTotalErrors.Operator
    leafs["value"] = outputTotalErrors.Value
    leafs["end-range-value"] = outputTotalErrors.EndRangeValue
    leafs["percent"] = outputTotalErrors.Percent
    leafs["rearm-type"] = outputTotalErrors.RearmType
    leafs["rearm-window"] = outputTotalErrors.RearmWindow
    return leafs
}

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetBundleName() string { return "cisco_ios_xr" }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetYangName() string { return "output-total-errors" }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) SetParent(parent types.Entity) { outputTotalErrors.parent = parent }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetParent() types.Entity { return outputTotalErrors.parent }

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetFilter() yfilter.YFilter { return inPackets.YFilter }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) SetFilter(yf yfilter.YFilter) { inPackets.YFilter = yf }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetSegmentPath() string {
    return "in-packets"
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inPackets.Operator
    leafs["value"] = inPackets.Value
    leafs["end-range-value"] = inPackets.EndRangeValue
    leafs["percent"] = inPackets.Percent
    leafs["rearm-type"] = inPackets.RearmType
    leafs["rearm-window"] = inPackets.RearmWindow
    return leafs
}

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetYangName() string { return "in-packets" }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) SetParent(parent types.Entity) { inPackets.parent = parent }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetParent() types.Entity { return inPackets.parent }

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto
// Number of inbound packets discarded with
// unknown protocol
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetFilter() yfilter.YFilter { return inputUnknownProto.YFilter }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) SetFilter(yf yfilter.YFilter) { inputUnknownProto.YFilter = yf }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetSegmentPath() string {
    return "input-unknown-proto"
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputUnknownProto.Operator
    leafs["value"] = inputUnknownProto.Value
    leafs["end-range-value"] = inputUnknownProto.EndRangeValue
    leafs["percent"] = inputUnknownProto.Percent
    leafs["rearm-type"] = inputUnknownProto.RearmType
    leafs["rearm-window"] = inputUnknownProto.RearmWindow
    return leafs
}

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetBundleName() string { return "cisco_ios_xr" }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetYangName() string { return "input-unknown-proto" }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) SetParent(parent types.Entity) { inputUnknownProto.parent = parent }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetParent() types.Entity { return inputUnknownProto.parent }

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetFilter() yfilter.YFilter { return inputQueueDrops.YFilter }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) SetFilter(yf yfilter.YFilter) { inputQueueDrops.YFilter = yf }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetSegmentPath() string {
    return "input-queue-drops"
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputQueueDrops.Operator
    leafs["value"] = inputQueueDrops.Value
    leafs["end-range-value"] = inputQueueDrops.EndRangeValue
    leafs["percent"] = inputQueueDrops.Percent
    leafs["rearm-type"] = inputQueueDrops.RearmType
    leafs["rearm-window"] = inputQueueDrops.RearmWindow
    return leafs
}

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetBundleName() string { return "cisco_ios_xr" }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetYangName() string { return "input-queue-drops" }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) SetParent(parent types.Entity) { inputQueueDrops.parent = parent }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetParent() types.Entity { return inputQueueDrops.parent }

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetFilter() yfilter.YFilter { return inputTotalDrops.YFilter }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) SetFilter(yf yfilter.YFilter) { inputTotalDrops.YFilter = yf }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetSegmentPath() string {
    return "input-total-drops"
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputTotalDrops.Operator
    leafs["value"] = inputTotalDrops.Value
    leafs["end-range-value"] = inputTotalDrops.EndRangeValue
    leafs["percent"] = inputTotalDrops.Percent
    leafs["rearm-type"] = inputTotalDrops.RearmType
    leafs["rearm-window"] = inputTotalDrops.RearmWindow
    return leafs
}

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetBundleName() string { return "cisco_ios_xr" }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetYangName() string { return "input-total-drops" }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) SetParent(parent types.Entity) { inputTotalDrops.parent = parent }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetParent() types.Entity { return inputTotalDrops.parent }

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame
// Number of inbound packets with framing
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetFilter() yfilter.YFilter { return inputFrame.YFilter }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) SetFilter(yf yfilter.YFilter) { inputFrame.YFilter = yf }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetSegmentPath() string {
    return "input-frame"
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputFrame.Operator
    leafs["value"] = inputFrame.Value
    leafs["end-range-value"] = inputFrame.EndRangeValue
    leafs["percent"] = inputFrame.Percent
    leafs["rearm-type"] = inputFrame.RearmType
    leafs["rearm-window"] = inputFrame.RearmWindow
    return leafs
}

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetBundleName() string { return "cisco_ios_xr" }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetYangName() string { return "input-frame" }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) SetParent(parent types.Entity) { inputFrame.parent = parent }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetParent() types.Entity { return inputFrame.parent }

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetParentYangName() string { return "generic-counter-interface-template" }

// PerfMgmt_Threshold_LdpMpls
// MPLS LDP threshold configuration
type PerfMgmt_Threshold_LdpMpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP threshold templates.
    LdpMplsTemplates PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetFilter() yfilter.YFilter { return ldpMpls.YFilter }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) SetFilter(yf yfilter.YFilter) { ldpMpls.YFilter = yf }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetGoName(yname string) string {
    if yname == "ldp-mpls-templates" { return "LdpMplsTemplates" }
    return ""
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetSegmentPath() string {
    return "ldp-mpls"
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-mpls-templates" {
        return &ldpMpls.LdpMplsTemplates
    }
    return nil
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp-mpls-templates"] = &ldpMpls.LdpMplsTemplates
    return children
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetYangName() string { return "ldp-mpls" }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) SetParent(parent types.Entity) { ldpMpls.parent = parent }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetParent() types.Entity { return ldpMpls.parent }

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates
// MPLS LDP threshold templates
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate.
    LdpMplsTemplate []PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetFilter() yfilter.YFilter { return ldpMplsTemplates.YFilter }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) SetFilter(yf yfilter.YFilter) { ldpMplsTemplates.YFilter = yf }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetGoName(yname string) string {
    if yname == "ldp-mpls-template" { return "LdpMplsTemplate" }
    return ""
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetSegmentPath() string {
    return "ldp-mpls-templates"
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp-mpls-template" {
        for _, c := range ldpMplsTemplates.LdpMplsTemplate {
            if ldpMplsTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate{}
        ldpMplsTemplates.LdpMplsTemplate = append(ldpMplsTemplates.LdpMplsTemplate, child)
        return &ldpMplsTemplates.LdpMplsTemplate[len(ldpMplsTemplates.LdpMplsTemplate)-1]
    }
    return nil
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ldpMplsTemplates.LdpMplsTemplate {
        children[ldpMplsTemplates.LdpMplsTemplate[i].GetSegmentPath()] = &ldpMplsTemplates.LdpMplsTemplate[i]
    }
    return children
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetYangName() string { return "ldp-mpls-templates" }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) SetParent(parent types.Entity) { ldpMplsTemplates.parent = parent }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetParent() types.Entity { return ldpMplsTemplates.parent }

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetParentYangName() string { return "ldp-mpls" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
// MPLS LDP threshold template instance
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of Address Withdraw messages received.
    AddressWithdrawMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd

    // Number of Label Withdraw messages received.
    LabelWithdrawMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd

    // Number of Address Withdraw messages sent.
    AddressWithdrawMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent

    // Number of Label Withdraw messages sent.
    LabelWithdrawMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent

    // Number of Notification messages received.
    NotificationMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd

    // Total number of messages received.
    TotalMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd

    // Number of Notification messages sent.
    NotificationMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent

    // Total number of messages sent.
    TotalMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent

    // Number of LAbel Release messages received.
    LabelReleaseMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd

    // Number of Init messages received.
    InitMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd

    // Number of Label Release messages sent.
    LabelReleaseMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent

    // Number of Init messages sent.
    InitMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent

    // Number of Label Mapping messages received.
    LabelMappingMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd

    // Number of Keepalive messages received.
    KeepaliveMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd

    // Number of Label Mapping messages sent.
    LabelMappingMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent

    // Number of Keepalive messages sent.
    KeepaliveMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent

    // Number of Address messages received.
    AddressMsgsRcvd PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd

    // Number of Address messages sent.
    AddressMsgsSent PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetFilter() yfilter.YFilter { return ldpMplsTemplate.YFilter }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) SetFilter(yf yfilter.YFilter) { ldpMplsTemplate.YFilter = yf }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "address-withdraw-msgs-rcvd" { return "AddressWithdrawMsgsRcvd" }
    if yname == "label-withdraw-msgs-rcvd" { return "LabelWithdrawMsgsRcvd" }
    if yname == "address-withdraw-msgs-sent" { return "AddressWithdrawMsgsSent" }
    if yname == "label-withdraw-msgs-sent" { return "LabelWithdrawMsgsSent" }
    if yname == "notification-msgs-rcvd" { return "NotificationMsgsRcvd" }
    if yname == "total-msgs-rcvd" { return "TotalMsgsRcvd" }
    if yname == "notification-msgs-sent" { return "NotificationMsgsSent" }
    if yname == "total-msgs-sent" { return "TotalMsgsSent" }
    if yname == "label-release-msgs-rcvd" { return "LabelReleaseMsgsRcvd" }
    if yname == "init-msgs-rcvd" { return "InitMsgsRcvd" }
    if yname == "label-release-msgs-sent" { return "LabelReleaseMsgsSent" }
    if yname == "init-msgs-sent" { return "InitMsgsSent" }
    if yname == "label-mapping-msgs-rcvd" { return "LabelMappingMsgsRcvd" }
    if yname == "keepalive-msgs-rcvd" { return "KeepaliveMsgsRcvd" }
    if yname == "label-mapping-msgs-sent" { return "LabelMappingMsgsSent" }
    if yname == "keepalive-msgs-sent" { return "KeepaliveMsgsSent" }
    if yname == "address-msgs-rcvd" { return "AddressMsgsRcvd" }
    if yname == "address-msgs-sent" { return "AddressMsgsSent" }
    return ""
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetSegmentPath() string {
    return "ldp-mpls-template" + "[template-name='" + fmt.Sprintf("%v", ldpMplsTemplate.TemplateName) + "']"
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-withdraw-msgs-rcvd" {
        return &ldpMplsTemplate.AddressWithdrawMsgsRcvd
    }
    if childYangName == "label-withdraw-msgs-rcvd" {
        return &ldpMplsTemplate.LabelWithdrawMsgsRcvd
    }
    if childYangName == "address-withdraw-msgs-sent" {
        return &ldpMplsTemplate.AddressWithdrawMsgsSent
    }
    if childYangName == "label-withdraw-msgs-sent" {
        return &ldpMplsTemplate.LabelWithdrawMsgsSent
    }
    if childYangName == "notification-msgs-rcvd" {
        return &ldpMplsTemplate.NotificationMsgsRcvd
    }
    if childYangName == "total-msgs-rcvd" {
        return &ldpMplsTemplate.TotalMsgsRcvd
    }
    if childYangName == "notification-msgs-sent" {
        return &ldpMplsTemplate.NotificationMsgsSent
    }
    if childYangName == "total-msgs-sent" {
        return &ldpMplsTemplate.TotalMsgsSent
    }
    if childYangName == "label-release-msgs-rcvd" {
        return &ldpMplsTemplate.LabelReleaseMsgsRcvd
    }
    if childYangName == "init-msgs-rcvd" {
        return &ldpMplsTemplate.InitMsgsRcvd
    }
    if childYangName == "label-release-msgs-sent" {
        return &ldpMplsTemplate.LabelReleaseMsgsSent
    }
    if childYangName == "init-msgs-sent" {
        return &ldpMplsTemplate.InitMsgsSent
    }
    if childYangName == "label-mapping-msgs-rcvd" {
        return &ldpMplsTemplate.LabelMappingMsgsRcvd
    }
    if childYangName == "keepalive-msgs-rcvd" {
        return &ldpMplsTemplate.KeepaliveMsgsRcvd
    }
    if childYangName == "label-mapping-msgs-sent" {
        return &ldpMplsTemplate.LabelMappingMsgsSent
    }
    if childYangName == "keepalive-msgs-sent" {
        return &ldpMplsTemplate.KeepaliveMsgsSent
    }
    if childYangName == "address-msgs-rcvd" {
        return &ldpMplsTemplate.AddressMsgsRcvd
    }
    if childYangName == "address-msgs-sent" {
        return &ldpMplsTemplate.AddressMsgsSent
    }
    return nil
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-withdraw-msgs-rcvd"] = &ldpMplsTemplate.AddressWithdrawMsgsRcvd
    children["label-withdraw-msgs-rcvd"] = &ldpMplsTemplate.LabelWithdrawMsgsRcvd
    children["address-withdraw-msgs-sent"] = &ldpMplsTemplate.AddressWithdrawMsgsSent
    children["label-withdraw-msgs-sent"] = &ldpMplsTemplate.LabelWithdrawMsgsSent
    children["notification-msgs-rcvd"] = &ldpMplsTemplate.NotificationMsgsRcvd
    children["total-msgs-rcvd"] = &ldpMplsTemplate.TotalMsgsRcvd
    children["notification-msgs-sent"] = &ldpMplsTemplate.NotificationMsgsSent
    children["total-msgs-sent"] = &ldpMplsTemplate.TotalMsgsSent
    children["label-release-msgs-rcvd"] = &ldpMplsTemplate.LabelReleaseMsgsRcvd
    children["init-msgs-rcvd"] = &ldpMplsTemplate.InitMsgsRcvd
    children["label-release-msgs-sent"] = &ldpMplsTemplate.LabelReleaseMsgsSent
    children["init-msgs-sent"] = &ldpMplsTemplate.InitMsgsSent
    children["label-mapping-msgs-rcvd"] = &ldpMplsTemplate.LabelMappingMsgsRcvd
    children["keepalive-msgs-rcvd"] = &ldpMplsTemplate.KeepaliveMsgsRcvd
    children["label-mapping-msgs-sent"] = &ldpMplsTemplate.LabelMappingMsgsSent
    children["keepalive-msgs-sent"] = &ldpMplsTemplate.KeepaliveMsgsSent
    children["address-msgs-rcvd"] = &ldpMplsTemplate.AddressMsgsRcvd
    children["address-msgs-sent"] = &ldpMplsTemplate.AddressMsgsSent
    return children
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ldpMplsTemplate.TemplateName
    leafs["sample-interval"] = ldpMplsTemplate.SampleInterval
    return leafs
}

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetYangName() string { return "ldp-mpls-template" }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) SetParent(parent types.Entity) { ldpMplsTemplate.parent = parent }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetParent() types.Entity { return ldpMplsTemplate.parent }

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetParentYangName() string { return "ldp-mpls-templates" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd
// Number of Address Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetFilter() yfilter.YFilter { return addressWithdrawMsgsRcvd.YFilter }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) SetFilter(yf yfilter.YFilter) { addressWithdrawMsgsRcvd.YFilter = yf }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetSegmentPath() string {
    return "address-withdraw-msgs-rcvd"
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = addressWithdrawMsgsRcvd.Operator
    leafs["value"] = addressWithdrawMsgsRcvd.Value
    leafs["end-range-value"] = addressWithdrawMsgsRcvd.EndRangeValue
    leafs["percent"] = addressWithdrawMsgsRcvd.Percent
    leafs["rearm-type"] = addressWithdrawMsgsRcvd.RearmType
    leafs["rearm-window"] = addressWithdrawMsgsRcvd.RearmWindow
    return leafs
}

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetYangName() string { return "address-withdraw-msgs-rcvd" }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) SetParent(parent types.Entity) { addressWithdrawMsgsRcvd.parent = parent }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetParent() types.Entity { return addressWithdrawMsgsRcvd.parent }

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd
// Number of Label Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetFilter() yfilter.YFilter { return labelWithdrawMsgsRcvd.YFilter }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) SetFilter(yf yfilter.YFilter) { labelWithdrawMsgsRcvd.YFilter = yf }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetSegmentPath() string {
    return "label-withdraw-msgs-rcvd"
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelWithdrawMsgsRcvd.Operator
    leafs["value"] = labelWithdrawMsgsRcvd.Value
    leafs["end-range-value"] = labelWithdrawMsgsRcvd.EndRangeValue
    leafs["percent"] = labelWithdrawMsgsRcvd.Percent
    leafs["rearm-type"] = labelWithdrawMsgsRcvd.RearmType
    leafs["rearm-window"] = labelWithdrawMsgsRcvd.RearmWindow
    return leafs
}

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetYangName() string { return "label-withdraw-msgs-rcvd" }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) SetParent(parent types.Entity) { labelWithdrawMsgsRcvd.parent = parent }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetParent() types.Entity { return labelWithdrawMsgsRcvd.parent }

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent
// Number of Address Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetFilter() yfilter.YFilter { return addressWithdrawMsgsSent.YFilter }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) SetFilter(yf yfilter.YFilter) { addressWithdrawMsgsSent.YFilter = yf }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetSegmentPath() string {
    return "address-withdraw-msgs-sent"
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = addressWithdrawMsgsSent.Operator
    leafs["value"] = addressWithdrawMsgsSent.Value
    leafs["end-range-value"] = addressWithdrawMsgsSent.EndRangeValue
    leafs["percent"] = addressWithdrawMsgsSent.Percent
    leafs["rearm-type"] = addressWithdrawMsgsSent.RearmType
    leafs["rearm-window"] = addressWithdrawMsgsSent.RearmWindow
    return leafs
}

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetYangName() string { return "address-withdraw-msgs-sent" }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) SetParent(parent types.Entity) { addressWithdrawMsgsSent.parent = parent }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetParent() types.Entity { return addressWithdrawMsgsSent.parent }

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent
// Number of Label Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetFilter() yfilter.YFilter { return labelWithdrawMsgsSent.YFilter }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) SetFilter(yf yfilter.YFilter) { labelWithdrawMsgsSent.YFilter = yf }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetSegmentPath() string {
    return "label-withdraw-msgs-sent"
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelWithdrawMsgsSent.Operator
    leafs["value"] = labelWithdrawMsgsSent.Value
    leafs["end-range-value"] = labelWithdrawMsgsSent.EndRangeValue
    leafs["percent"] = labelWithdrawMsgsSent.Percent
    leafs["rearm-type"] = labelWithdrawMsgsSent.RearmType
    leafs["rearm-window"] = labelWithdrawMsgsSent.RearmWindow
    return leafs
}

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetYangName() string { return "label-withdraw-msgs-sent" }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) SetParent(parent types.Entity) { labelWithdrawMsgsSent.parent = parent }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetParent() types.Entity { return labelWithdrawMsgsSent.parent }

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd
// Number of Notification messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetFilter() yfilter.YFilter { return notificationMsgsRcvd.YFilter }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) SetFilter(yf yfilter.YFilter) { notificationMsgsRcvd.YFilter = yf }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetSegmentPath() string {
    return "notification-msgs-rcvd"
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = notificationMsgsRcvd.Operator
    leafs["value"] = notificationMsgsRcvd.Value
    leafs["end-range-value"] = notificationMsgsRcvd.EndRangeValue
    leafs["percent"] = notificationMsgsRcvd.Percent
    leafs["rearm-type"] = notificationMsgsRcvd.RearmType
    leafs["rearm-window"] = notificationMsgsRcvd.RearmWindow
    return leafs
}

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetYangName() string { return "notification-msgs-rcvd" }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) SetParent(parent types.Entity) { notificationMsgsRcvd.parent = parent }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetParent() types.Entity { return notificationMsgsRcvd.parent }

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd
// Total number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..65536.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..65536.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetFilter() yfilter.YFilter { return totalMsgsRcvd.YFilter }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) SetFilter(yf yfilter.YFilter) { totalMsgsRcvd.YFilter = yf }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetSegmentPath() string {
    return "total-msgs-rcvd"
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = totalMsgsRcvd.Operator
    leafs["value"] = totalMsgsRcvd.Value
    leafs["end-range-value"] = totalMsgsRcvd.EndRangeValue
    leafs["percent"] = totalMsgsRcvd.Percent
    leafs["rearm-type"] = totalMsgsRcvd.RearmType
    leafs["rearm-window"] = totalMsgsRcvd.RearmWindow
    return leafs
}

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetYangName() string { return "total-msgs-rcvd" }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) SetParent(parent types.Entity) { totalMsgsRcvd.parent = parent }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetParent() types.Entity { return totalMsgsRcvd.parent }

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent
// Number of Notification messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetFilter() yfilter.YFilter { return notificationMsgsSent.YFilter }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) SetFilter(yf yfilter.YFilter) { notificationMsgsSent.YFilter = yf }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetSegmentPath() string {
    return "notification-msgs-sent"
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = notificationMsgsSent.Operator
    leafs["value"] = notificationMsgsSent.Value
    leafs["end-range-value"] = notificationMsgsSent.EndRangeValue
    leafs["percent"] = notificationMsgsSent.Percent
    leafs["rearm-type"] = notificationMsgsSent.RearmType
    leafs["rearm-window"] = notificationMsgsSent.RearmWindow
    return leafs
}

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetYangName() string { return "notification-msgs-sent" }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) SetParent(parent types.Entity) { notificationMsgsSent.parent = parent }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetParent() types.Entity { return notificationMsgsSent.parent }

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent
// Total number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetFilter() yfilter.YFilter { return totalMsgsSent.YFilter }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) SetFilter(yf yfilter.YFilter) { totalMsgsSent.YFilter = yf }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetSegmentPath() string {
    return "total-msgs-sent"
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = totalMsgsSent.Operator
    leafs["value"] = totalMsgsSent.Value
    leafs["end-range-value"] = totalMsgsSent.EndRangeValue
    leafs["percent"] = totalMsgsSent.Percent
    leafs["rearm-type"] = totalMsgsSent.RearmType
    leafs["rearm-window"] = totalMsgsSent.RearmWindow
    return leafs
}

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetYangName() string { return "total-msgs-sent" }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) SetParent(parent types.Entity) { totalMsgsSent.parent = parent }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetParent() types.Entity { return totalMsgsSent.parent }

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd
// Number of LAbel Release messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetFilter() yfilter.YFilter { return labelReleaseMsgsRcvd.YFilter }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) SetFilter(yf yfilter.YFilter) { labelReleaseMsgsRcvd.YFilter = yf }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetSegmentPath() string {
    return "label-release-msgs-rcvd"
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelReleaseMsgsRcvd.Operator
    leafs["value"] = labelReleaseMsgsRcvd.Value
    leafs["end-range-value"] = labelReleaseMsgsRcvd.EndRangeValue
    leafs["percent"] = labelReleaseMsgsRcvd.Percent
    leafs["rearm-type"] = labelReleaseMsgsRcvd.RearmType
    leafs["rearm-window"] = labelReleaseMsgsRcvd.RearmWindow
    return leafs
}

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetYangName() string { return "label-release-msgs-rcvd" }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) SetParent(parent types.Entity) { labelReleaseMsgsRcvd.parent = parent }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetParent() types.Entity { return labelReleaseMsgsRcvd.parent }

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd
// Number of Init messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetFilter() yfilter.YFilter { return initMsgsRcvd.YFilter }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) SetFilter(yf yfilter.YFilter) { initMsgsRcvd.YFilter = yf }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetSegmentPath() string {
    return "init-msgs-rcvd"
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = initMsgsRcvd.Operator
    leafs["value"] = initMsgsRcvd.Value
    leafs["end-range-value"] = initMsgsRcvd.EndRangeValue
    leafs["percent"] = initMsgsRcvd.Percent
    leafs["rearm-type"] = initMsgsRcvd.RearmType
    leafs["rearm-window"] = initMsgsRcvd.RearmWindow
    return leafs
}

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetYangName() string { return "init-msgs-rcvd" }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) SetParent(parent types.Entity) { initMsgsRcvd.parent = parent }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetParent() types.Entity { return initMsgsRcvd.parent }

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent
// Number of Label Release messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetFilter() yfilter.YFilter { return labelReleaseMsgsSent.YFilter }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) SetFilter(yf yfilter.YFilter) { labelReleaseMsgsSent.YFilter = yf }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetSegmentPath() string {
    return "label-release-msgs-sent"
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelReleaseMsgsSent.Operator
    leafs["value"] = labelReleaseMsgsSent.Value
    leafs["end-range-value"] = labelReleaseMsgsSent.EndRangeValue
    leafs["percent"] = labelReleaseMsgsSent.Percent
    leafs["rearm-type"] = labelReleaseMsgsSent.RearmType
    leafs["rearm-window"] = labelReleaseMsgsSent.RearmWindow
    return leafs
}

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetYangName() string { return "label-release-msgs-sent" }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) SetParent(parent types.Entity) { labelReleaseMsgsSent.parent = parent }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetParent() types.Entity { return labelReleaseMsgsSent.parent }

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent
// Number of Init messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetFilter() yfilter.YFilter { return initMsgsSent.YFilter }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) SetFilter(yf yfilter.YFilter) { initMsgsSent.YFilter = yf }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetSegmentPath() string {
    return "init-msgs-sent"
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = initMsgsSent.Operator
    leafs["value"] = initMsgsSent.Value
    leafs["end-range-value"] = initMsgsSent.EndRangeValue
    leafs["percent"] = initMsgsSent.Percent
    leafs["rearm-type"] = initMsgsSent.RearmType
    leafs["rearm-window"] = initMsgsSent.RearmWindow
    return leafs
}

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetYangName() string { return "init-msgs-sent" }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) SetParent(parent types.Entity) { initMsgsSent.parent = parent }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetParent() types.Entity { return initMsgsSent.parent }

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd
// Number of Label Mapping messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetFilter() yfilter.YFilter { return labelMappingMsgsRcvd.YFilter }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) SetFilter(yf yfilter.YFilter) { labelMappingMsgsRcvd.YFilter = yf }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetSegmentPath() string {
    return "label-mapping-msgs-rcvd"
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelMappingMsgsRcvd.Operator
    leafs["value"] = labelMappingMsgsRcvd.Value
    leafs["end-range-value"] = labelMappingMsgsRcvd.EndRangeValue
    leafs["percent"] = labelMappingMsgsRcvd.Percent
    leafs["rearm-type"] = labelMappingMsgsRcvd.RearmType
    leafs["rearm-window"] = labelMappingMsgsRcvd.RearmWindow
    return leafs
}

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetYangName() string { return "label-mapping-msgs-rcvd" }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) SetParent(parent types.Entity) { labelMappingMsgsRcvd.parent = parent }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetParent() types.Entity { return labelMappingMsgsRcvd.parent }

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd
// Number of Keepalive messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetFilter() yfilter.YFilter { return keepaliveMsgsRcvd.YFilter }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) SetFilter(yf yfilter.YFilter) { keepaliveMsgsRcvd.YFilter = yf }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetSegmentPath() string {
    return "keepalive-msgs-rcvd"
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = keepaliveMsgsRcvd.Operator
    leafs["value"] = keepaliveMsgsRcvd.Value
    leafs["end-range-value"] = keepaliveMsgsRcvd.EndRangeValue
    leafs["percent"] = keepaliveMsgsRcvd.Percent
    leafs["rearm-type"] = keepaliveMsgsRcvd.RearmType
    leafs["rearm-window"] = keepaliveMsgsRcvd.RearmWindow
    return leafs
}

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetYangName() string { return "keepalive-msgs-rcvd" }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) SetParent(parent types.Entity) { keepaliveMsgsRcvd.parent = parent }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetParent() types.Entity { return keepaliveMsgsRcvd.parent }

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent
// Number of Label Mapping messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetFilter() yfilter.YFilter { return labelMappingMsgsSent.YFilter }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) SetFilter(yf yfilter.YFilter) { labelMappingMsgsSent.YFilter = yf }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetSegmentPath() string {
    return "label-mapping-msgs-sent"
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = labelMappingMsgsSent.Operator
    leafs["value"] = labelMappingMsgsSent.Value
    leafs["end-range-value"] = labelMappingMsgsSent.EndRangeValue
    leafs["percent"] = labelMappingMsgsSent.Percent
    leafs["rearm-type"] = labelMappingMsgsSent.RearmType
    leafs["rearm-window"] = labelMappingMsgsSent.RearmWindow
    return leafs
}

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetYangName() string { return "label-mapping-msgs-sent" }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) SetParent(parent types.Entity) { labelMappingMsgsSent.parent = parent }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetParent() types.Entity { return labelMappingMsgsSent.parent }

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent
// Number of Keepalive messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetFilter() yfilter.YFilter { return keepaliveMsgsSent.YFilter }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) SetFilter(yf yfilter.YFilter) { keepaliveMsgsSent.YFilter = yf }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetSegmentPath() string {
    return "keepalive-msgs-sent"
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = keepaliveMsgsSent.Operator
    leafs["value"] = keepaliveMsgsSent.Value
    leafs["end-range-value"] = keepaliveMsgsSent.EndRangeValue
    leafs["percent"] = keepaliveMsgsSent.Percent
    leafs["rearm-type"] = keepaliveMsgsSent.RearmType
    leafs["rearm-window"] = keepaliveMsgsSent.RearmWindow
    return leafs
}

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetYangName() string { return "keepalive-msgs-sent" }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) SetParent(parent types.Entity) { keepaliveMsgsSent.parent = parent }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetParent() types.Entity { return keepaliveMsgsSent.parent }

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd
// Number of Address messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetFilter() yfilter.YFilter { return addressMsgsRcvd.YFilter }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) SetFilter(yf yfilter.YFilter) { addressMsgsRcvd.YFilter = yf }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetSegmentPath() string {
    return "address-msgs-rcvd"
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = addressMsgsRcvd.Operator
    leafs["value"] = addressMsgsRcvd.Value
    leafs["end-range-value"] = addressMsgsRcvd.EndRangeValue
    leafs["percent"] = addressMsgsRcvd.Percent
    leafs["rearm-type"] = addressMsgsRcvd.RearmType
    leafs["rearm-window"] = addressMsgsRcvd.RearmWindow
    return leafs
}

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetBundleName() string { return "cisco_ios_xr" }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetYangName() string { return "address-msgs-rcvd" }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) SetParent(parent types.Entity) { addressMsgsRcvd.parent = parent }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetParent() types.Entity { return addressMsgsRcvd.parent }

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent
// Number of Address messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetFilter() yfilter.YFilter { return addressMsgsSent.YFilter }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) SetFilter(yf yfilter.YFilter) { addressMsgsSent.YFilter = yf }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetSegmentPath() string {
    return "address-msgs-sent"
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = addressMsgsSent.Operator
    leafs["value"] = addressMsgsSent.Value
    leafs["end-range-value"] = addressMsgsSent.EndRangeValue
    leafs["percent"] = addressMsgsSent.Percent
    leafs["rearm-type"] = addressMsgsSent.RearmType
    leafs["rearm-window"] = addressMsgsSent.RearmWindow
    return leafs
}

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetBundleName() string { return "cisco_ios_xr" }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetYangName() string { return "address-msgs-sent" }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) SetParent(parent types.Entity) { addressMsgsSent.parent = parent }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetParent() types.Entity { return addressMsgsSent.parent }

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetParentYangName() string { return "ldp-mpls-template" }

// PerfMgmt_Threshold_BasicCounterInterface
// Interface Basic Counter threshold configuration
type PerfMgmt_Threshold_BasicCounterInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Basic Counter threshold templates.
    BasicCounterInterfaceTemplates PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetFilter() yfilter.YFilter { return basicCounterInterface.YFilter }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) SetFilter(yf yfilter.YFilter) { basicCounterInterface.YFilter = yf }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetGoName(yname string) string {
    if yname == "basic-counter-interface-templates" { return "BasicCounterInterfaceTemplates" }
    return ""
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetSegmentPath() string {
    return "basic-counter-interface"
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-counter-interface-templates" {
        return &basicCounterInterface.BasicCounterInterfaceTemplates
    }
    return nil
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["basic-counter-interface-templates"] = &basicCounterInterface.BasicCounterInterfaceTemplates
    return children
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetYangName() string { return "basic-counter-interface" }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) SetParent(parent types.Entity) { basicCounterInterface.parent = parent }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetParent() types.Entity { return basicCounterInterface.parent }

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates
// Interface Basic Counter threshold templates
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Basic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate.
    BasicCounterInterfaceTemplate []PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetFilter() yfilter.YFilter { return basicCounterInterfaceTemplates.YFilter }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) SetFilter(yf yfilter.YFilter) { basicCounterInterfaceTemplates.YFilter = yf }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetGoName(yname string) string {
    if yname == "basic-counter-interface-template" { return "BasicCounterInterfaceTemplate" }
    return ""
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetSegmentPath() string {
    return "basic-counter-interface-templates"
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "basic-counter-interface-template" {
        for _, c := range basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate {
            if basicCounterInterfaceTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate{}
        basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate = append(basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate, child)
        return &basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[len(basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate)-1]
    }
    return nil
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate {
        children[basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i].GetSegmentPath()] = &basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i]
    }
    return children
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetYangName() string { return "basic-counter-interface-templates" }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) SetParent(parent types.Entity) { basicCounterInterfaceTemplates.parent = parent }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetParent() types.Entity { return basicCounterInterfaceTemplates.parent }

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetParentYangName() string { return "basic-counter-interface" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
// Interface Basic Counter threshold template
// instance
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable instance filtering by VRF name regular expression . The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Number of inbound octets/bytes.
    InOctets PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets

    // Number of outbound octets/bytes.
    OutOctets PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets

    // Number of outbound queue drops.
    OutputQueueDrops PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops

    // Number of inbound incorrect packets discarded.
    InputTotalErrors PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors

    // Number of outbound correct packets discarded.
    OutputTotalDrops PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops

    // Number of outbound packets.
    OutPackets PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets

    // Number of outbound incorrect packets discarded.
    OutputTotalErrors PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors

    // Number of inbound packets.
    InPackets PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets

    // Number of input queue drops.
    InputQueueDrops PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops

    // Number of inbound correct packets discarded.
    InputTotalDrops PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetFilter() yfilter.YFilter { return basicCounterInterfaceTemplate.YFilter }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) SetFilter(yf yfilter.YFilter) { basicCounterInterfaceTemplate.YFilter = yf }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "output-queue-drops" { return "OutputQueueDrops" }
    if yname == "input-total-errors" { return "InputTotalErrors" }
    if yname == "output-total-drops" { return "OutputTotalDrops" }
    if yname == "out-packets" { return "OutPackets" }
    if yname == "output-total-errors" { return "OutputTotalErrors" }
    if yname == "in-packets" { return "InPackets" }
    if yname == "input-queue-drops" { return "InputQueueDrops" }
    if yname == "input-total-drops" { return "InputTotalDrops" }
    return ""
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetSegmentPath() string {
    return "basic-counter-interface-template" + "[template-name='" + fmt.Sprintf("%v", basicCounterInterfaceTemplate.TemplateName) + "']"
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "in-octets" {
        return &basicCounterInterfaceTemplate.InOctets
    }
    if childYangName == "out-octets" {
        return &basicCounterInterfaceTemplate.OutOctets
    }
    if childYangName == "output-queue-drops" {
        return &basicCounterInterfaceTemplate.OutputQueueDrops
    }
    if childYangName == "input-total-errors" {
        return &basicCounterInterfaceTemplate.InputTotalErrors
    }
    if childYangName == "output-total-drops" {
        return &basicCounterInterfaceTemplate.OutputTotalDrops
    }
    if childYangName == "out-packets" {
        return &basicCounterInterfaceTemplate.OutPackets
    }
    if childYangName == "output-total-errors" {
        return &basicCounterInterfaceTemplate.OutputTotalErrors
    }
    if childYangName == "in-packets" {
        return &basicCounterInterfaceTemplate.InPackets
    }
    if childYangName == "input-queue-drops" {
        return &basicCounterInterfaceTemplate.InputQueueDrops
    }
    if childYangName == "input-total-drops" {
        return &basicCounterInterfaceTemplate.InputTotalDrops
    }
    return nil
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["in-octets"] = &basicCounterInterfaceTemplate.InOctets
    children["out-octets"] = &basicCounterInterfaceTemplate.OutOctets
    children["output-queue-drops"] = &basicCounterInterfaceTemplate.OutputQueueDrops
    children["input-total-errors"] = &basicCounterInterfaceTemplate.InputTotalErrors
    children["output-total-drops"] = &basicCounterInterfaceTemplate.OutputTotalDrops
    children["out-packets"] = &basicCounterInterfaceTemplate.OutPackets
    children["output-total-errors"] = &basicCounterInterfaceTemplate.OutputTotalErrors
    children["in-packets"] = &basicCounterInterfaceTemplate.InPackets
    children["input-queue-drops"] = &basicCounterInterfaceTemplate.InputQueueDrops
    children["input-total-drops"] = &basicCounterInterfaceTemplate.InputTotalDrops
    return children
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = basicCounterInterfaceTemplate.TemplateName
    leafs["sample-interval"] = basicCounterInterfaceTemplate.SampleInterval
    leafs["reg-exp-group"] = basicCounterInterfaceTemplate.RegExpGroup
    leafs["vrf-group"] = basicCounterInterfaceTemplate.VrfGroup
    return leafs
}

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetYangName() string { return "basic-counter-interface-template" }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) SetParent(parent types.Entity) { basicCounterInterfaceTemplate.parent = parent }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetParent() types.Entity { return basicCounterInterfaceTemplate.parent }

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetParentYangName() string { return "basic-counter-interface-templates" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetFilter() yfilter.YFilter { return inOctets.YFilter }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) SetFilter(yf yfilter.YFilter) { inOctets.YFilter = yf }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetSegmentPath() string {
    return "in-octets"
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inOctets.Operator
    leafs["value"] = inOctets.Value
    leafs["end-range-value"] = inOctets.EndRangeValue
    leafs["percent"] = inOctets.Percent
    leafs["rearm-type"] = inOctets.RearmType
    leafs["rearm-window"] = inOctets.RearmWindow
    return leafs
}

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetBundleName() string { return "cisco_ios_xr" }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetYangName() string { return "in-octets" }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) SetParent(parent types.Entity) { inOctets.parent = parent }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetParent() types.Entity { return inOctets.parent }

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetFilter() yfilter.YFilter { return outOctets.YFilter }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) SetFilter(yf yfilter.YFilter) { outOctets.YFilter = yf }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetSegmentPath() string {
    return "out-octets"
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outOctets.Operator
    leafs["value"] = outOctets.Value
    leafs["end-range-value"] = outOctets.EndRangeValue
    leafs["percent"] = outOctets.Percent
    leafs["rearm-type"] = outOctets.RearmType
    leafs["rearm-window"] = outOctets.RearmWindow
    return leafs
}

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetBundleName() string { return "cisco_ios_xr" }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetYangName() string { return "out-octets" }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) SetParent(parent types.Entity) { outOctets.parent = parent }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetParent() types.Entity { return outOctets.parent }

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops
// Number of outbound queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetFilter() yfilter.YFilter { return outputQueueDrops.YFilter }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) SetFilter(yf yfilter.YFilter) { outputQueueDrops.YFilter = yf }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetSegmentPath() string {
    return "output-queue-drops"
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputQueueDrops.Operator
    leafs["value"] = outputQueueDrops.Value
    leafs["end-range-value"] = outputQueueDrops.EndRangeValue
    leafs["percent"] = outputQueueDrops.Percent
    leafs["rearm-type"] = outputQueueDrops.RearmType
    leafs["rearm-window"] = outputQueueDrops.RearmWindow
    return leafs
}

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetBundleName() string { return "cisco_ios_xr" }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetYangName() string { return "output-queue-drops" }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) SetParent(parent types.Entity) { outputQueueDrops.parent = parent }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetParent() types.Entity { return outputQueueDrops.parent }

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetFilter() yfilter.YFilter { return inputTotalErrors.YFilter }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) SetFilter(yf yfilter.YFilter) { inputTotalErrors.YFilter = yf }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetSegmentPath() string {
    return "input-total-errors"
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputTotalErrors.Operator
    leafs["value"] = inputTotalErrors.Value
    leafs["end-range-value"] = inputTotalErrors.EndRangeValue
    leafs["percent"] = inputTotalErrors.Percent
    leafs["rearm-type"] = inputTotalErrors.RearmType
    leafs["rearm-window"] = inputTotalErrors.RearmWindow
    return leafs
}

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetBundleName() string { return "cisco_ios_xr" }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetYangName() string { return "input-total-errors" }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) SetParent(parent types.Entity) { inputTotalErrors.parent = parent }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetParent() types.Entity { return inputTotalErrors.parent }

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetFilter() yfilter.YFilter { return outputTotalDrops.YFilter }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) SetFilter(yf yfilter.YFilter) { outputTotalDrops.YFilter = yf }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetSegmentPath() string {
    return "output-total-drops"
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputTotalDrops.Operator
    leafs["value"] = outputTotalDrops.Value
    leafs["end-range-value"] = outputTotalDrops.EndRangeValue
    leafs["percent"] = outputTotalDrops.Percent
    leafs["rearm-type"] = outputTotalDrops.RearmType
    leafs["rearm-window"] = outputTotalDrops.RearmWindow
    return leafs
}

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetBundleName() string { return "cisco_ios_xr" }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetYangName() string { return "output-total-drops" }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) SetParent(parent types.Entity) { outputTotalDrops.parent = parent }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetParent() types.Entity { return outputTotalDrops.parent }

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetFilter() yfilter.YFilter { return outPackets.YFilter }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) SetFilter(yf yfilter.YFilter) { outPackets.YFilter = yf }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetSegmentPath() string {
    return "out-packets"
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outPackets.Operator
    leafs["value"] = outPackets.Value
    leafs["end-range-value"] = outPackets.EndRangeValue
    leafs["percent"] = outPackets.Percent
    leafs["rearm-type"] = outPackets.RearmType
    leafs["rearm-window"] = outPackets.RearmWindow
    return leafs
}

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetYangName() string { return "out-packets" }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) SetParent(parent types.Entity) { outPackets.parent = parent }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetParent() types.Entity { return outPackets.parent }

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetFilter() yfilter.YFilter { return outputTotalErrors.YFilter }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) SetFilter(yf yfilter.YFilter) { outputTotalErrors.YFilter = yf }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetSegmentPath() string {
    return "output-total-errors"
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputTotalErrors.Operator
    leafs["value"] = outputTotalErrors.Value
    leafs["end-range-value"] = outputTotalErrors.EndRangeValue
    leafs["percent"] = outputTotalErrors.Percent
    leafs["rearm-type"] = outputTotalErrors.RearmType
    leafs["rearm-window"] = outputTotalErrors.RearmWindow
    return leafs
}

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetBundleName() string { return "cisco_ios_xr" }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetYangName() string { return "output-total-errors" }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) SetParent(parent types.Entity) { outputTotalErrors.parent = parent }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetParent() types.Entity { return outputTotalErrors.parent }

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetFilter() yfilter.YFilter { return inPackets.YFilter }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) SetFilter(yf yfilter.YFilter) { inPackets.YFilter = yf }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetSegmentPath() string {
    return "in-packets"
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inPackets.Operator
    leafs["value"] = inPackets.Value
    leafs["end-range-value"] = inPackets.EndRangeValue
    leafs["percent"] = inPackets.Percent
    leafs["rearm-type"] = inPackets.RearmType
    leafs["rearm-window"] = inPackets.RearmWindow
    return leafs
}

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetYangName() string { return "in-packets" }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) SetParent(parent types.Entity) { inPackets.parent = parent }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetParent() types.Entity { return inPackets.parent }

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetFilter() yfilter.YFilter { return inputQueueDrops.YFilter }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) SetFilter(yf yfilter.YFilter) { inputQueueDrops.YFilter = yf }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetSegmentPath() string {
    return "input-queue-drops"
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputQueueDrops.Operator
    leafs["value"] = inputQueueDrops.Value
    leafs["end-range-value"] = inputQueueDrops.EndRangeValue
    leafs["percent"] = inputQueueDrops.Percent
    leafs["rearm-type"] = inputQueueDrops.RearmType
    leafs["rearm-window"] = inputQueueDrops.RearmWindow
    return leafs
}

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetBundleName() string { return "cisco_ios_xr" }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetYangName() string { return "input-queue-drops" }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) SetParent(parent types.Entity) { inputQueueDrops.parent = parent }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetParent() types.Entity { return inputQueueDrops.parent }

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetFilter() yfilter.YFilter { return inputTotalDrops.YFilter }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) SetFilter(yf yfilter.YFilter) { inputTotalDrops.YFilter = yf }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetSegmentPath() string {
    return "input-total-drops"
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputTotalDrops.Operator
    leafs["value"] = inputTotalDrops.Value
    leafs["end-range-value"] = inputTotalDrops.EndRangeValue
    leafs["percent"] = inputTotalDrops.Percent
    leafs["rearm-type"] = inputTotalDrops.RearmType
    leafs["rearm-window"] = inputTotalDrops.RearmWindow
    return leafs
}

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetBundleName() string { return "cisco_ios_xr" }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetYangName() string { return "input-total-drops" }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) SetParent(parent types.Entity) { inputTotalDrops.parent = parent }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetParent() types.Entity { return inputTotalDrops.parent }

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetParentYangName() string { return "basic-counter-interface-template" }

// PerfMgmt_Threshold_Bgp
// BGP threshold configuration
type PerfMgmt_Threshold_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP threshold templates.
    BgpTemplates PerfMgmt_Threshold_Bgp_BgpTemplates
}

func (bgp *PerfMgmt_Threshold_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PerfMgmt_Threshold_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PerfMgmt_Threshold_Bgp) GetGoName(yname string) string {
    if yname == "bgp-templates" { return "BgpTemplates" }
    return ""
}

func (bgp *PerfMgmt_Threshold_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PerfMgmt_Threshold_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-templates" {
        return &bgp.BgpTemplates
    }
    return nil
}

func (bgp *PerfMgmt_Threshold_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-templates"] = &bgp.BgpTemplates
    return children
}

func (bgp *PerfMgmt_Threshold_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *PerfMgmt_Threshold_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PerfMgmt_Threshold_Bgp) GetYangName() string { return "bgp" }

func (bgp *PerfMgmt_Threshold_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PerfMgmt_Threshold_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PerfMgmt_Threshold_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PerfMgmt_Threshold_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PerfMgmt_Threshold_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PerfMgmt_Threshold_Bgp) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_Bgp_BgpTemplates
// BGP threshold templates
type PerfMgmt_Threshold_Bgp_BgpTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate.
    BgpTemplate []PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetFilter() yfilter.YFilter { return bgpTemplates.YFilter }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) SetFilter(yf yfilter.YFilter) { bgpTemplates.YFilter = yf }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetGoName(yname string) string {
    if yname == "bgp-template" { return "BgpTemplate" }
    return ""
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetSegmentPath() string {
    return "bgp-templates"
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-template" {
        for _, c := range bgpTemplates.BgpTemplate {
            if bgpTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate{}
        bgpTemplates.BgpTemplate = append(bgpTemplates.BgpTemplate, child)
        return &bgpTemplates.BgpTemplate[len(bgpTemplates.BgpTemplate)-1]
    }
    return nil
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bgpTemplates.BgpTemplate {
        children[bgpTemplates.BgpTemplate[i].GetSegmentPath()] = &bgpTemplates.BgpTemplate[i]
    }
    return children
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetYangName() string { return "bgp-templates" }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) SetParent(parent types.Entity) { bgpTemplates.parent = parent }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetParent() types.Entity { return bgpTemplates.parent }

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetParentYangName() string { return "bgp" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
// BGP threshold template instance
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of update messages sent.
    OutputUpdateMessages PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages

    // Number of error notifications received.
    ErrorsReceived PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived

    // Number of times the connection was established.
    ConnEstablished PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished

    // Number of messages sent.
    OutputMessages PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages

    // Number of times the connection was dropped.
    ConnDropped PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped

    // Number of update messages received.
    InputUpdateMessages PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages

    // Number of error notifications sent.
    ErrorsSent PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent

    // Number of messages received.
    InputMessages PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetFilter() yfilter.YFilter { return bgpTemplate.YFilter }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) SetFilter(yf yfilter.YFilter) { bgpTemplate.YFilter = yf }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "output-update-messages" { return "OutputUpdateMessages" }
    if yname == "errors-received" { return "ErrorsReceived" }
    if yname == "conn-established" { return "ConnEstablished" }
    if yname == "output-messages" { return "OutputMessages" }
    if yname == "conn-dropped" { return "ConnDropped" }
    if yname == "input-update-messages" { return "InputUpdateMessages" }
    if yname == "errors-sent" { return "ErrorsSent" }
    if yname == "input-messages" { return "InputMessages" }
    return ""
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetSegmentPath() string {
    return "bgp-template" + "[template-name='" + fmt.Sprintf("%v", bgpTemplate.TemplateName) + "']"
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-update-messages" {
        return &bgpTemplate.OutputUpdateMessages
    }
    if childYangName == "errors-received" {
        return &bgpTemplate.ErrorsReceived
    }
    if childYangName == "conn-established" {
        return &bgpTemplate.ConnEstablished
    }
    if childYangName == "output-messages" {
        return &bgpTemplate.OutputMessages
    }
    if childYangName == "conn-dropped" {
        return &bgpTemplate.ConnDropped
    }
    if childYangName == "input-update-messages" {
        return &bgpTemplate.InputUpdateMessages
    }
    if childYangName == "errors-sent" {
        return &bgpTemplate.ErrorsSent
    }
    if childYangName == "input-messages" {
        return &bgpTemplate.InputMessages
    }
    return nil
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output-update-messages"] = &bgpTemplate.OutputUpdateMessages
    children["errors-received"] = &bgpTemplate.ErrorsReceived
    children["conn-established"] = &bgpTemplate.ConnEstablished
    children["output-messages"] = &bgpTemplate.OutputMessages
    children["conn-dropped"] = &bgpTemplate.ConnDropped
    children["input-update-messages"] = &bgpTemplate.InputUpdateMessages
    children["errors-sent"] = &bgpTemplate.ErrorsSent
    children["input-messages"] = &bgpTemplate.InputMessages
    return children
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = bgpTemplate.TemplateName
    leafs["sample-interval"] = bgpTemplate.SampleInterval
    return leafs
}

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetYangName() string { return "bgp-template" }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) SetParent(parent types.Entity) { bgpTemplate.parent = parent }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetParent() types.Entity { return bgpTemplate.parent }

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetParentYangName() string { return "bgp-templates" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages
// Number of update messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetFilter() yfilter.YFilter { return outputUpdateMessages.YFilter }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) SetFilter(yf yfilter.YFilter) { outputUpdateMessages.YFilter = yf }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetSegmentPath() string {
    return "output-update-messages"
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputUpdateMessages.Operator
    leafs["value"] = outputUpdateMessages.Value
    leafs["end-range-value"] = outputUpdateMessages.EndRangeValue
    leafs["percent"] = outputUpdateMessages.Percent
    leafs["rearm-type"] = outputUpdateMessages.RearmType
    leafs["rearm-window"] = outputUpdateMessages.RearmWindow
    return leafs
}

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetBundleName() string { return "cisco_ios_xr" }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetYangName() string { return "output-update-messages" }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) SetParent(parent types.Entity) { outputUpdateMessages.parent = parent }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetParent() types.Entity { return outputUpdateMessages.parent }

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived
// Number of error notifications received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetFilter() yfilter.YFilter { return errorsReceived.YFilter }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) SetFilter(yf yfilter.YFilter) { errorsReceived.YFilter = yf }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetSegmentPath() string {
    return "errors-received"
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = errorsReceived.Operator
    leafs["value"] = errorsReceived.Value
    leafs["end-range-value"] = errorsReceived.EndRangeValue
    leafs["percent"] = errorsReceived.Percent
    leafs["rearm-type"] = errorsReceived.RearmType
    leafs["rearm-window"] = errorsReceived.RearmWindow
    return leafs
}

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetBundleName() string { return "cisco_ios_xr" }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetYangName() string { return "errors-received" }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) SetParent(parent types.Entity) { errorsReceived.parent = parent }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetParent() types.Entity { return errorsReceived.parent }

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished
// Number of times the connection was
// established
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetFilter() yfilter.YFilter { return connEstablished.YFilter }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) SetFilter(yf yfilter.YFilter) { connEstablished.YFilter = yf }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetSegmentPath() string {
    return "conn-established"
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = connEstablished.Operator
    leafs["value"] = connEstablished.Value
    leafs["end-range-value"] = connEstablished.EndRangeValue
    leafs["percent"] = connEstablished.Percent
    leafs["rearm-type"] = connEstablished.RearmType
    leafs["rearm-window"] = connEstablished.RearmWindow
    return leafs
}

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetBundleName() string { return "cisco_ios_xr" }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetYangName() string { return "conn-established" }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) SetParent(parent types.Entity) { connEstablished.parent = parent }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetParent() types.Entity { return connEstablished.parent }

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages
// Number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetFilter() yfilter.YFilter { return outputMessages.YFilter }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) SetFilter(yf yfilter.YFilter) { outputMessages.YFilter = yf }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetSegmentPath() string {
    return "output-messages"
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputMessages.Operator
    leafs["value"] = outputMessages.Value
    leafs["end-range-value"] = outputMessages.EndRangeValue
    leafs["percent"] = outputMessages.Percent
    leafs["rearm-type"] = outputMessages.RearmType
    leafs["rearm-window"] = outputMessages.RearmWindow
    return leafs
}

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetBundleName() string { return "cisco_ios_xr" }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetYangName() string { return "output-messages" }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) SetParent(parent types.Entity) { outputMessages.parent = parent }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetParent() types.Entity { return outputMessages.parent }

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped
// Number of times the connection was dropped
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetFilter() yfilter.YFilter { return connDropped.YFilter }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) SetFilter(yf yfilter.YFilter) { connDropped.YFilter = yf }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetSegmentPath() string {
    return "conn-dropped"
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = connDropped.Operator
    leafs["value"] = connDropped.Value
    leafs["end-range-value"] = connDropped.EndRangeValue
    leafs["percent"] = connDropped.Percent
    leafs["rearm-type"] = connDropped.RearmType
    leafs["rearm-window"] = connDropped.RearmWindow
    return leafs
}

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetBundleName() string { return "cisco_ios_xr" }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetYangName() string { return "conn-dropped" }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) SetParent(parent types.Entity) { connDropped.parent = parent }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetParent() types.Entity { return connDropped.parent }

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages
// Number of update messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetFilter() yfilter.YFilter { return inputUpdateMessages.YFilter }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) SetFilter(yf yfilter.YFilter) { inputUpdateMessages.YFilter = yf }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetSegmentPath() string {
    return "input-update-messages"
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputUpdateMessages.Operator
    leafs["value"] = inputUpdateMessages.Value
    leafs["end-range-value"] = inputUpdateMessages.EndRangeValue
    leafs["percent"] = inputUpdateMessages.Percent
    leafs["rearm-type"] = inputUpdateMessages.RearmType
    leafs["rearm-window"] = inputUpdateMessages.RearmWindow
    return leafs
}

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetBundleName() string { return "cisco_ios_xr" }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetYangName() string { return "input-update-messages" }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) SetParent(parent types.Entity) { inputUpdateMessages.parent = parent }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetParent() types.Entity { return inputUpdateMessages.parent }

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent
// Number of error notifications sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetFilter() yfilter.YFilter { return errorsSent.YFilter }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) SetFilter(yf yfilter.YFilter) { errorsSent.YFilter = yf }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetSegmentPath() string {
    return "errors-sent"
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = errorsSent.Operator
    leafs["value"] = errorsSent.Value
    leafs["end-range-value"] = errorsSent.EndRangeValue
    leafs["percent"] = errorsSent.Percent
    leafs["rearm-type"] = errorsSent.RearmType
    leafs["rearm-window"] = errorsSent.RearmWindow
    return leafs
}

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetBundleName() string { return "cisco_ios_xr" }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetYangName() string { return "errors-sent" }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) SetParent(parent types.Entity) { errorsSent.parent = parent }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetParent() types.Entity { return errorsSent.parent }

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages
// Number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetFilter() yfilter.YFilter { return inputMessages.YFilter }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) SetFilter(yf yfilter.YFilter) { inputMessages.YFilter = yf }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetSegmentPath() string {
    return "input-messages"
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputMessages.Operator
    leafs["value"] = inputMessages.Value
    leafs["end-range-value"] = inputMessages.EndRangeValue
    leafs["percent"] = inputMessages.Percent
    leafs["rearm-type"] = inputMessages.RearmType
    leafs["rearm-window"] = inputMessages.RearmWindow
    return leafs
}

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetBundleName() string { return "cisco_ios_xr" }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetYangName() string { return "input-messages" }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) SetParent(parent types.Entity) { inputMessages.parent = parent }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetParent() types.Entity { return inputMessages.parent }

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetParentYangName() string { return "bgp-template" }

// PerfMgmt_Threshold_Ospfv2Protocol
// OSPF v2 Protocol threshold configuration
type PerfMgmt_Threshold_Ospfv2Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold templates.
    Ospfv2ProtocolTemplates PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetFilter() yfilter.YFilter { return ospfv2Protocol.YFilter }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) SetFilter(yf yfilter.YFilter) { ospfv2Protocol.YFilter = yf }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-templates" { return "Ospfv2ProtocolTemplates" }
    return ""
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetSegmentPath() string {
    return "ospfv2-protocol"
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-templates" {
        return &ospfv2Protocol.Ospfv2ProtocolTemplates
    }
    return nil
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv2-protocol-templates"] = &ospfv2Protocol.Ospfv2ProtocolTemplates
    return children
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetYangName() string { return "ospfv2-protocol" }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) SetParent(parent types.Entity) { ospfv2Protocol.parent = parent }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetParent() types.Entity { return ospfv2Protocol.parent }

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate.
    Ospfv2ProtocolTemplate []PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetFilter() yfilter.YFilter { return ospfv2ProtocolTemplates.YFilter }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolTemplates.YFilter = yf }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetGoName(yname string) string {
    if yname == "ospfv2-protocol-template" { return "Ospfv2ProtocolTemplate" }
    return ""
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetSegmentPath() string {
    return "ospfv2-protocol-templates"
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv2-protocol-template" {
        for _, c := range ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate {
            if ospfv2ProtocolTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate{}
        ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate = append(ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate, child)
        return &ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[len(ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate)-1]
    }
    return nil
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate {
        children[ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i].GetSegmentPath()] = &ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i]
    }
    return children
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetYangName() string { return "ospfv2-protocol-templates" }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) SetParent(parent types.Entity) { ospfv2ProtocolTemplates.parent = parent }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetParent() types.Entity { return ospfv2ProtocolTemplates.parent }

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetParentYangName() string { return "ospfv2-protocol" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of packets received with checksum errors.
    ChecksumErrors PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors

    // Number of LSA received in LSA Acknowledgements.
    InputLsaAcksLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa

    // Number of LSA sent in DBD packets.
    OutputDbDsLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa

    // Number of LSA received in DBD packets.
    InputDbDsLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa

    // Number of LSA Updates received.
    InputLsaUpdates PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates

    // Number of DBD packets sent.
    OutputDbDs PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs

    // Number of LSA sent in LSA Updates.
    OutputLsaUpdatesLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa

    // Number of DBD packets received.
    InputDbDs PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs

    // Number of LSA received in LSA Updates.
    InputLsaUpdatesLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa

    // Total number of packets sent.
    OutputPackets PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets

    // Total number of packets received.
    InputPackets PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets

    // Total number of packets sent.
    OutputHelloPackets PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets

    // Number of Hello packets received.
    InputHelloPackets PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets

    // Number of LS Requests sent.
    OutputLsRequests PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests

    // Number of LSA sent in LSA Acknowledgements.
    OutputLsaAcksLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa

    // Number of LSA Acknowledgements sent.
    OutputLsaAcks PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks

    // Number of LSA Acknowledgements received.
    InputLsaAcks PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks

    // Number of LSA Updates sent.
    OutputLsaUpdates PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates

    // Number of LSA sent in LS Requests.
    OutputLsRequestsLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa

    // Number of LSA received in LS Requests.
    InputLsRequestsLsa PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa

    // Number of LS Requests received.
    InputLsRequests PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetFilter() yfilter.YFilter { return ospfv2ProtocolTemplate.YFilter }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) SetFilter(yf yfilter.YFilter) { ospfv2ProtocolTemplate.YFilter = yf }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "checksum-errors" { return "ChecksumErrors" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    return ""
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetSegmentPath() string {
    return "ospfv2-protocol-template" + "[template-name='" + fmt.Sprintf("%v", ospfv2ProtocolTemplate.TemplateName) + "']"
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "checksum-errors" {
        return &ospfv2ProtocolTemplate.ChecksumErrors
    }
    if childYangName == "input-lsa-acks-lsa" {
        return &ospfv2ProtocolTemplate.InputLsaAcksLsa
    }
    if childYangName == "output-db-ds-lsa" {
        return &ospfv2ProtocolTemplate.OutputDbDsLsa
    }
    if childYangName == "input-db-ds-lsa" {
        return &ospfv2ProtocolTemplate.InputDbDsLsa
    }
    if childYangName == "input-lsa-updates" {
        return &ospfv2ProtocolTemplate.InputLsaUpdates
    }
    if childYangName == "output-db-ds" {
        return &ospfv2ProtocolTemplate.OutputDbDs
    }
    if childYangName == "output-lsa-updates-lsa" {
        return &ospfv2ProtocolTemplate.OutputLsaUpdatesLsa
    }
    if childYangName == "input-db-ds" {
        return &ospfv2ProtocolTemplate.InputDbDs
    }
    if childYangName == "input-lsa-updates-lsa" {
        return &ospfv2ProtocolTemplate.InputLsaUpdatesLsa
    }
    if childYangName == "output-packets" {
        return &ospfv2ProtocolTemplate.OutputPackets
    }
    if childYangName == "input-packets" {
        return &ospfv2ProtocolTemplate.InputPackets
    }
    if childYangName == "output-hello-packets" {
        return &ospfv2ProtocolTemplate.OutputHelloPackets
    }
    if childYangName == "input-hello-packets" {
        return &ospfv2ProtocolTemplate.InputHelloPackets
    }
    if childYangName == "output-ls-requests" {
        return &ospfv2ProtocolTemplate.OutputLsRequests
    }
    if childYangName == "output-lsa-acks-lsa" {
        return &ospfv2ProtocolTemplate.OutputLsaAcksLsa
    }
    if childYangName == "output-lsa-acks" {
        return &ospfv2ProtocolTemplate.OutputLsaAcks
    }
    if childYangName == "input-lsa-acks" {
        return &ospfv2ProtocolTemplate.InputLsaAcks
    }
    if childYangName == "output-lsa-updates" {
        return &ospfv2ProtocolTemplate.OutputLsaUpdates
    }
    if childYangName == "output-ls-requests-lsa" {
        return &ospfv2ProtocolTemplate.OutputLsRequestsLsa
    }
    if childYangName == "input-ls-requests-lsa" {
        return &ospfv2ProtocolTemplate.InputLsRequestsLsa
    }
    if childYangName == "input-ls-requests" {
        return &ospfv2ProtocolTemplate.InputLsRequests
    }
    return nil
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["checksum-errors"] = &ospfv2ProtocolTemplate.ChecksumErrors
    children["input-lsa-acks-lsa"] = &ospfv2ProtocolTemplate.InputLsaAcksLsa
    children["output-db-ds-lsa"] = &ospfv2ProtocolTemplate.OutputDbDsLsa
    children["input-db-ds-lsa"] = &ospfv2ProtocolTemplate.InputDbDsLsa
    children["input-lsa-updates"] = &ospfv2ProtocolTemplate.InputLsaUpdates
    children["output-db-ds"] = &ospfv2ProtocolTemplate.OutputDbDs
    children["output-lsa-updates-lsa"] = &ospfv2ProtocolTemplate.OutputLsaUpdatesLsa
    children["input-db-ds"] = &ospfv2ProtocolTemplate.InputDbDs
    children["input-lsa-updates-lsa"] = &ospfv2ProtocolTemplate.InputLsaUpdatesLsa
    children["output-packets"] = &ospfv2ProtocolTemplate.OutputPackets
    children["input-packets"] = &ospfv2ProtocolTemplate.InputPackets
    children["output-hello-packets"] = &ospfv2ProtocolTemplate.OutputHelloPackets
    children["input-hello-packets"] = &ospfv2ProtocolTemplate.InputHelloPackets
    children["output-ls-requests"] = &ospfv2ProtocolTemplate.OutputLsRequests
    children["output-lsa-acks-lsa"] = &ospfv2ProtocolTemplate.OutputLsaAcksLsa
    children["output-lsa-acks"] = &ospfv2ProtocolTemplate.OutputLsaAcks
    children["input-lsa-acks"] = &ospfv2ProtocolTemplate.InputLsaAcks
    children["output-lsa-updates"] = &ospfv2ProtocolTemplate.OutputLsaUpdates
    children["output-ls-requests-lsa"] = &ospfv2ProtocolTemplate.OutputLsRequestsLsa
    children["input-ls-requests-lsa"] = &ospfv2ProtocolTemplate.InputLsRequestsLsa
    children["input-ls-requests"] = &ospfv2ProtocolTemplate.InputLsRequests
    return children
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv2ProtocolTemplate.TemplateName
    leafs["sample-interval"] = ospfv2ProtocolTemplate.SampleInterval
    return leafs
}

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetYangName() string { return "ospfv2-protocol-template" }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) SetParent(parent types.Entity) { ospfv2ProtocolTemplate.parent = parent }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetParent() types.Entity { return ospfv2ProtocolTemplate.parent }

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetParentYangName() string { return "ospfv2-protocol-templates" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors
// Number of packets received with checksum
// errors
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetFilter() yfilter.YFilter { return checksumErrors.YFilter }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) SetFilter(yf yfilter.YFilter) { checksumErrors.YFilter = yf }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetSegmentPath() string {
    return "checksum-errors"
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = checksumErrors.Operator
    leafs["value"] = checksumErrors.Value
    leafs["end-range-value"] = checksumErrors.EndRangeValue
    leafs["percent"] = checksumErrors.Percent
    leafs["rearm-type"] = checksumErrors.RearmType
    leafs["rearm-window"] = checksumErrors.RearmWindow
    return leafs
}

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetBundleName() string { return "cisco_ios_xr" }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetYangName() string { return "checksum-errors" }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) SetParent(parent types.Entity) { checksumErrors.parent = parent }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetParent() types.Entity { return checksumErrors.parent }

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetFilter() yfilter.YFilter { return inputLsaAcksLsa.YFilter }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) SetFilter(yf yfilter.YFilter) { inputLsaAcksLsa.YFilter = yf }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetSegmentPath() string {
    return "input-lsa-acks-lsa"
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaAcksLsa.Operator
    leafs["value"] = inputLsaAcksLsa.Value
    leafs["end-range-value"] = inputLsaAcksLsa.EndRangeValue
    leafs["percent"] = inputLsaAcksLsa.Percent
    leafs["rearm-type"] = inputLsaAcksLsa.RearmType
    leafs["rearm-window"] = inputLsaAcksLsa.RearmWindow
    return leafs
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetYangName() string { return "input-lsa-acks-lsa" }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) SetParent(parent types.Entity) { inputLsaAcksLsa.parent = parent }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetParent() types.Entity { return inputLsaAcksLsa.parent }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetFilter() yfilter.YFilter { return outputDbDsLsa.YFilter }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) SetFilter(yf yfilter.YFilter) { outputDbDsLsa.YFilter = yf }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetSegmentPath() string {
    return "output-db-ds-lsa"
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputDbDsLsa.Operator
    leafs["value"] = outputDbDsLsa.Value
    leafs["end-range-value"] = outputDbDsLsa.EndRangeValue
    leafs["percent"] = outputDbDsLsa.Percent
    leafs["rearm-type"] = outputDbDsLsa.RearmType
    leafs["rearm-window"] = outputDbDsLsa.RearmWindow
    return leafs
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetYangName() string { return "output-db-ds-lsa" }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) SetParent(parent types.Entity) { outputDbDsLsa.parent = parent }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetParent() types.Entity { return outputDbDsLsa.parent }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetFilter() yfilter.YFilter { return inputDbDsLsa.YFilter }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) SetFilter(yf yfilter.YFilter) { inputDbDsLsa.YFilter = yf }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetSegmentPath() string {
    return "input-db-ds-lsa"
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputDbDsLsa.Operator
    leafs["value"] = inputDbDsLsa.Value
    leafs["end-range-value"] = inputDbDsLsa.EndRangeValue
    leafs["percent"] = inputDbDsLsa.Percent
    leafs["rearm-type"] = inputDbDsLsa.RearmType
    leafs["rearm-window"] = inputDbDsLsa.RearmWindow
    return leafs
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetYangName() string { return "input-db-ds-lsa" }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) SetParent(parent types.Entity) { inputDbDsLsa.parent = parent }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetParent() types.Entity { return inputDbDsLsa.parent }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetFilter() yfilter.YFilter { return inputLsaUpdates.YFilter }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) SetFilter(yf yfilter.YFilter) { inputLsaUpdates.YFilter = yf }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetSegmentPath() string {
    return "input-lsa-updates"
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaUpdates.Operator
    leafs["value"] = inputLsaUpdates.Value
    leafs["end-range-value"] = inputLsaUpdates.EndRangeValue
    leafs["percent"] = inputLsaUpdates.Percent
    leafs["rearm-type"] = inputLsaUpdates.RearmType
    leafs["rearm-window"] = inputLsaUpdates.RearmWindow
    return leafs
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetYangName() string { return "input-lsa-updates" }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) SetParent(parent types.Entity) { inputLsaUpdates.parent = parent }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetParent() types.Entity { return inputLsaUpdates.parent }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetFilter() yfilter.YFilter { return outputDbDs.YFilter }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) SetFilter(yf yfilter.YFilter) { outputDbDs.YFilter = yf }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetSegmentPath() string {
    return "output-db-ds"
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputDbDs.Operator
    leafs["value"] = outputDbDs.Value
    leafs["end-range-value"] = outputDbDs.EndRangeValue
    leafs["percent"] = outputDbDs.Percent
    leafs["rearm-type"] = outputDbDs.RearmType
    leafs["rearm-window"] = outputDbDs.RearmWindow
    return leafs
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetBundleName() string { return "cisco_ios_xr" }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetYangName() string { return "output-db-ds" }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) SetParent(parent types.Entity) { outputDbDs.parent = parent }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetParent() types.Entity { return outputDbDs.parent }

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetFilter() yfilter.YFilter { return outputLsaUpdatesLsa.YFilter }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) SetFilter(yf yfilter.YFilter) { outputLsaUpdatesLsa.YFilter = yf }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetSegmentPath() string {
    return "output-lsa-updates-lsa"
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaUpdatesLsa.Operator
    leafs["value"] = outputLsaUpdatesLsa.Value
    leafs["end-range-value"] = outputLsaUpdatesLsa.EndRangeValue
    leafs["percent"] = outputLsaUpdatesLsa.Percent
    leafs["rearm-type"] = outputLsaUpdatesLsa.RearmType
    leafs["rearm-window"] = outputLsaUpdatesLsa.RearmWindow
    return leafs
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetYangName() string { return "output-lsa-updates-lsa" }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) SetParent(parent types.Entity) { outputLsaUpdatesLsa.parent = parent }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetParent() types.Entity { return outputLsaUpdatesLsa.parent }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetFilter() yfilter.YFilter { return inputDbDs.YFilter }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) SetFilter(yf yfilter.YFilter) { inputDbDs.YFilter = yf }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetSegmentPath() string {
    return "input-db-ds"
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputDbDs.Operator
    leafs["value"] = inputDbDs.Value
    leafs["end-range-value"] = inputDbDs.EndRangeValue
    leafs["percent"] = inputDbDs.Percent
    leafs["rearm-type"] = inputDbDs.RearmType
    leafs["rearm-window"] = inputDbDs.RearmWindow
    return leafs
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetBundleName() string { return "cisco_ios_xr" }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetYangName() string { return "input-db-ds" }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) SetParent(parent types.Entity) { inputDbDs.parent = parent }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetParent() types.Entity { return inputDbDs.parent }

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetFilter() yfilter.YFilter { return inputLsaUpdatesLsa.YFilter }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) SetFilter(yf yfilter.YFilter) { inputLsaUpdatesLsa.YFilter = yf }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetSegmentPath() string {
    return "input-lsa-updates-lsa"
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaUpdatesLsa.Operator
    leafs["value"] = inputLsaUpdatesLsa.Value
    leafs["end-range-value"] = inputLsaUpdatesLsa.EndRangeValue
    leafs["percent"] = inputLsaUpdatesLsa.Percent
    leafs["rearm-type"] = inputLsaUpdatesLsa.RearmType
    leafs["rearm-window"] = inputLsaUpdatesLsa.RearmWindow
    return leafs
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetYangName() string { return "input-lsa-updates-lsa" }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) SetParent(parent types.Entity) { inputLsaUpdatesLsa.parent = parent }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetParent() types.Entity { return inputLsaUpdatesLsa.parent }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetFilter() yfilter.YFilter { return outputPackets.YFilter }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) SetFilter(yf yfilter.YFilter) { outputPackets.YFilter = yf }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetSegmentPath() string {
    return "output-packets"
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputPackets.Operator
    leafs["value"] = outputPackets.Value
    leafs["end-range-value"] = outputPackets.EndRangeValue
    leafs["percent"] = outputPackets.Percent
    leafs["rearm-type"] = outputPackets.RearmType
    leafs["rearm-window"] = outputPackets.RearmWindow
    return leafs
}

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetYangName() string { return "output-packets" }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) SetParent(parent types.Entity) { outputPackets.parent = parent }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetParent() types.Entity { return outputPackets.parent }

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetFilter() yfilter.YFilter { return inputPackets.YFilter }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) SetFilter(yf yfilter.YFilter) { inputPackets.YFilter = yf }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetSegmentPath() string {
    return "input-packets"
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputPackets.Operator
    leafs["value"] = inputPackets.Value
    leafs["end-range-value"] = inputPackets.EndRangeValue
    leafs["percent"] = inputPackets.Percent
    leafs["rearm-type"] = inputPackets.RearmType
    leafs["rearm-window"] = inputPackets.RearmWindow
    return leafs
}

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetYangName() string { return "input-packets" }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) SetParent(parent types.Entity) { inputPackets.parent = parent }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetParent() types.Entity { return inputPackets.parent }

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetFilter() yfilter.YFilter { return outputHelloPackets.YFilter }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) SetFilter(yf yfilter.YFilter) { outputHelloPackets.YFilter = yf }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetSegmentPath() string {
    return "output-hello-packets"
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputHelloPackets.Operator
    leafs["value"] = outputHelloPackets.Value
    leafs["end-range-value"] = outputHelloPackets.EndRangeValue
    leafs["percent"] = outputHelloPackets.Percent
    leafs["rearm-type"] = outputHelloPackets.RearmType
    leafs["rearm-window"] = outputHelloPackets.RearmWindow
    return leafs
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetYangName() string { return "output-hello-packets" }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) SetParent(parent types.Entity) { outputHelloPackets.parent = parent }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetParent() types.Entity { return outputHelloPackets.parent }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetFilter() yfilter.YFilter { return inputHelloPackets.YFilter }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) SetFilter(yf yfilter.YFilter) { inputHelloPackets.YFilter = yf }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetSegmentPath() string {
    return "input-hello-packets"
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputHelloPackets.Operator
    leafs["value"] = inputHelloPackets.Value
    leafs["end-range-value"] = inputHelloPackets.EndRangeValue
    leafs["percent"] = inputHelloPackets.Percent
    leafs["rearm-type"] = inputHelloPackets.RearmType
    leafs["rearm-window"] = inputHelloPackets.RearmWindow
    return leafs
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetYangName() string { return "input-hello-packets" }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) SetParent(parent types.Entity) { inputHelloPackets.parent = parent }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetParent() types.Entity { return inputHelloPackets.parent }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetFilter() yfilter.YFilter { return outputLsRequests.YFilter }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) SetFilter(yf yfilter.YFilter) { outputLsRequests.YFilter = yf }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetSegmentPath() string {
    return "output-ls-requests"
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsRequests.Operator
    leafs["value"] = outputLsRequests.Value
    leafs["end-range-value"] = outputLsRequests.EndRangeValue
    leafs["percent"] = outputLsRequests.Percent
    leafs["rearm-type"] = outputLsRequests.RearmType
    leafs["rearm-window"] = outputLsRequests.RearmWindow
    return leafs
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetYangName() string { return "output-ls-requests" }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) SetParent(parent types.Entity) { outputLsRequests.parent = parent }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetParent() types.Entity { return outputLsRequests.parent }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetFilter() yfilter.YFilter { return outputLsaAcksLsa.YFilter }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) SetFilter(yf yfilter.YFilter) { outputLsaAcksLsa.YFilter = yf }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetSegmentPath() string {
    return "output-lsa-acks-lsa"
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaAcksLsa.Operator
    leafs["value"] = outputLsaAcksLsa.Value
    leafs["end-range-value"] = outputLsaAcksLsa.EndRangeValue
    leafs["percent"] = outputLsaAcksLsa.Percent
    leafs["rearm-type"] = outputLsaAcksLsa.RearmType
    leafs["rearm-window"] = outputLsaAcksLsa.RearmWindow
    return leafs
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetYangName() string { return "output-lsa-acks-lsa" }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) SetParent(parent types.Entity) { outputLsaAcksLsa.parent = parent }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetParent() types.Entity { return outputLsaAcksLsa.parent }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetFilter() yfilter.YFilter { return outputLsaAcks.YFilter }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) SetFilter(yf yfilter.YFilter) { outputLsaAcks.YFilter = yf }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetSegmentPath() string {
    return "output-lsa-acks"
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaAcks.Operator
    leafs["value"] = outputLsaAcks.Value
    leafs["end-range-value"] = outputLsaAcks.EndRangeValue
    leafs["percent"] = outputLsaAcks.Percent
    leafs["rearm-type"] = outputLsaAcks.RearmType
    leafs["rearm-window"] = outputLsaAcks.RearmWindow
    return leafs
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetYangName() string { return "output-lsa-acks" }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) SetParent(parent types.Entity) { outputLsaAcks.parent = parent }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetParent() types.Entity { return outputLsaAcks.parent }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetFilter() yfilter.YFilter { return inputLsaAcks.YFilter }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) SetFilter(yf yfilter.YFilter) { inputLsaAcks.YFilter = yf }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetSegmentPath() string {
    return "input-lsa-acks"
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaAcks.Operator
    leafs["value"] = inputLsaAcks.Value
    leafs["end-range-value"] = inputLsaAcks.EndRangeValue
    leafs["percent"] = inputLsaAcks.Percent
    leafs["rearm-type"] = inputLsaAcks.RearmType
    leafs["rearm-window"] = inputLsaAcks.RearmWindow
    return leafs
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetYangName() string { return "input-lsa-acks" }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) SetParent(parent types.Entity) { inputLsaAcks.parent = parent }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetParent() types.Entity { return inputLsaAcks.parent }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetFilter() yfilter.YFilter { return outputLsaUpdates.YFilter }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) SetFilter(yf yfilter.YFilter) { outputLsaUpdates.YFilter = yf }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetSegmentPath() string {
    return "output-lsa-updates"
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaUpdates.Operator
    leafs["value"] = outputLsaUpdates.Value
    leafs["end-range-value"] = outputLsaUpdates.EndRangeValue
    leafs["percent"] = outputLsaUpdates.Percent
    leafs["rearm-type"] = outputLsaUpdates.RearmType
    leafs["rearm-window"] = outputLsaUpdates.RearmWindow
    return leafs
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetYangName() string { return "output-lsa-updates" }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) SetParent(parent types.Entity) { outputLsaUpdates.parent = parent }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetParent() types.Entity { return outputLsaUpdates.parent }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetFilter() yfilter.YFilter { return outputLsRequestsLsa.YFilter }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) SetFilter(yf yfilter.YFilter) { outputLsRequestsLsa.YFilter = yf }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetSegmentPath() string {
    return "output-ls-requests-lsa"
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsRequestsLsa.Operator
    leafs["value"] = outputLsRequestsLsa.Value
    leafs["end-range-value"] = outputLsRequestsLsa.EndRangeValue
    leafs["percent"] = outputLsRequestsLsa.Percent
    leafs["rearm-type"] = outputLsRequestsLsa.RearmType
    leafs["rearm-window"] = outputLsRequestsLsa.RearmWindow
    return leafs
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetYangName() string { return "output-ls-requests-lsa" }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) SetParent(parent types.Entity) { outputLsRequestsLsa.parent = parent }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetParent() types.Entity { return outputLsRequestsLsa.parent }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetFilter() yfilter.YFilter { return inputLsRequestsLsa.YFilter }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) SetFilter(yf yfilter.YFilter) { inputLsRequestsLsa.YFilter = yf }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetSegmentPath() string {
    return "input-ls-requests-lsa"
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsRequestsLsa.Operator
    leafs["value"] = inputLsRequestsLsa.Value
    leafs["end-range-value"] = inputLsRequestsLsa.EndRangeValue
    leafs["percent"] = inputLsRequestsLsa.Percent
    leafs["rearm-type"] = inputLsRequestsLsa.RearmType
    leafs["rearm-window"] = inputLsRequestsLsa.RearmWindow
    return leafs
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetYangName() string { return "input-ls-requests-lsa" }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) SetParent(parent types.Entity) { inputLsRequestsLsa.parent = parent }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetParent() types.Entity { return inputLsRequestsLsa.parent }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetFilter() yfilter.YFilter { return inputLsRequests.YFilter }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) SetFilter(yf yfilter.YFilter) { inputLsRequests.YFilter = yf }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetSegmentPath() string {
    return "input-ls-requests"
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsRequests.Operator
    leafs["value"] = inputLsRequests.Value
    leafs["end-range-value"] = inputLsRequests.EndRangeValue
    leafs["percent"] = inputLsRequests.Percent
    leafs["rearm-type"] = inputLsRequests.RearmType
    leafs["rearm-window"] = inputLsRequests.RearmWindow
    return leafs
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetYangName() string { return "input-ls-requests" }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) SetParent(parent types.Entity) { inputLsRequests.parent = parent }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetParent() types.Entity { return inputLsRequests.parent }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetParentYangName() string { return "ospfv2-protocol-template" }

// PerfMgmt_Threshold_CpuNode
// Node CPU threshold configuration
type PerfMgmt_Threshold_CpuNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node CPU threshold configuration templates.
    CpuNodeTemplates PerfMgmt_Threshold_CpuNode_CpuNodeTemplates
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetFilter() yfilter.YFilter { return cpuNode.YFilter }

func (cpuNode *PerfMgmt_Threshold_CpuNode) SetFilter(yf yfilter.YFilter) { cpuNode.YFilter = yf }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetGoName(yname string) string {
    if yname == "cpu-node-templates" { return "CpuNodeTemplates" }
    return ""
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetSegmentPath() string {
    return "cpu-node"
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-node-templates" {
        return &cpuNode.CpuNodeTemplates
    }
    return nil
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpu-node-templates"] = &cpuNode.CpuNodeTemplates
    return children
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetYangName() string { return "cpu-node" }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNode *PerfMgmt_Threshold_CpuNode) SetParent(parent types.Entity) { cpuNode.parent = parent }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetParent() types.Entity { return cpuNode.parent }

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates
// Node CPU threshold configuration templates
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node CPU threshold configuration template instances. The type is slice of
    // PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate.
    CpuNodeTemplate []PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetFilter() yfilter.YFilter { return cpuNodeTemplates.YFilter }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) SetFilter(yf yfilter.YFilter) { cpuNodeTemplates.YFilter = yf }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetGoName(yname string) string {
    if yname == "cpu-node-template" { return "CpuNodeTemplate" }
    return ""
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetSegmentPath() string {
    return "cpu-node-templates"
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-node-template" {
        for _, c := range cpuNodeTemplates.CpuNodeTemplate {
            if cpuNodeTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate{}
        cpuNodeTemplates.CpuNodeTemplate = append(cpuNodeTemplates.CpuNodeTemplate, child)
        return &cpuNodeTemplates.CpuNodeTemplate[len(cpuNodeTemplates.CpuNodeTemplate)-1]
    }
    return nil
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cpuNodeTemplates.CpuNodeTemplate {
        children[cpuNodeTemplates.CpuNodeTemplate[i].GetSegmentPath()] = &cpuNodeTemplates.CpuNodeTemplate[i]
    }
    return children
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetYangName() string { return "cpu-node-templates" }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) SetParent(parent types.Entity) { cpuNodeTemplates.parent = parent }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetParent() types.Entity { return cpuNodeTemplates.parent }

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetParentYangName() string { return "cpu-node" }

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
// Node CPU threshold configuration template
// instances
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Average %CPU utilization.
    AverageCpuUsed PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed

    // Number of processes.
    NoProcesses PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetFilter() yfilter.YFilter { return cpuNodeTemplate.YFilter }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) SetFilter(yf yfilter.YFilter) { cpuNodeTemplate.YFilter = yf }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    if yname == "no-processes" { return "NoProcesses" }
    return ""
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetSegmentPath() string {
    return "cpu-node-template" + "[template-name='" + fmt.Sprintf("%v", cpuNodeTemplate.TemplateName) + "']"
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "average-cpu-used" {
        return &cpuNodeTemplate.AverageCpuUsed
    }
    if childYangName == "no-processes" {
        return &cpuNodeTemplate.NoProcesses
    }
    return nil
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["average-cpu-used"] = &cpuNodeTemplate.AverageCpuUsed
    children["no-processes"] = &cpuNodeTemplate.NoProcesses
    return children
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = cpuNodeTemplate.TemplateName
    leafs["sample-interval"] = cpuNodeTemplate.SampleInterval
    return leafs
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetYangName() string { return "cpu-node-template" }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) SetParent(parent types.Entity) { cpuNodeTemplate.parent = parent }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetParent() types.Entity { return cpuNodeTemplate.parent }

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetParentYangName() string { return "cpu-node-templates" }

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..100.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..100.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetFilter() yfilter.YFilter { return averageCpuUsed.YFilter }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) SetFilter(yf yfilter.YFilter) { averageCpuUsed.YFilter = yf }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetSegmentPath() string {
    return "average-cpu-used"
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = averageCpuUsed.Operator
    leafs["value"] = averageCpuUsed.Value
    leafs["end-range-value"] = averageCpuUsed.EndRangeValue
    leafs["percent"] = averageCpuUsed.Percent
    leafs["rearm-type"] = averageCpuUsed.RearmType
    leafs["rearm-window"] = averageCpuUsed.RearmWindow
    return leafs
}

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetBundleName() string { return "cisco_ios_xr" }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetYangName() string { return "average-cpu-used" }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) SetParent(parent types.Entity) { averageCpuUsed.parent = parent }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetParent() types.Entity { return averageCpuUsed.parent }

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetParentYangName() string { return "cpu-node-template" }

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses
// Number of processes
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetFilter() yfilter.YFilter { return noProcesses.YFilter }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) SetFilter(yf yfilter.YFilter) { noProcesses.YFilter = yf }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetSegmentPath() string {
    return "no-processes"
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = noProcesses.Operator
    leafs["value"] = noProcesses.Value
    leafs["end-range-value"] = noProcesses.EndRangeValue
    leafs["percent"] = noProcesses.Percent
    leafs["rearm-type"] = noProcesses.RearmType
    leafs["rearm-window"] = noProcesses.RearmWindow
    return leafs
}

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetBundleName() string { return "cisco_ios_xr" }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetYangName() string { return "no-processes" }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) SetParent(parent types.Entity) { noProcesses.parent = parent }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetParent() types.Entity { return noProcesses.parent }

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetParentYangName() string { return "cpu-node-template" }

// PerfMgmt_Threshold_DataRateInterface
// Interface Data Rates threshold configuration
type PerfMgmt_Threshold_DataRateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Data Rates threshold templates.
    DataRateInterfaceTemplates PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetFilter() yfilter.YFilter { return dataRateInterface.YFilter }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) SetFilter(yf yfilter.YFilter) { dataRateInterface.YFilter = yf }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetGoName(yname string) string {
    if yname == "data-rate-interface-templates" { return "DataRateInterfaceTemplates" }
    return ""
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetSegmentPath() string {
    return "data-rate-interface"
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-rate-interface-templates" {
        return &dataRateInterface.DataRateInterfaceTemplates
    }
    return nil
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["data-rate-interface-templates"] = &dataRateInterface.DataRateInterfaceTemplates
    return children
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetYangName() string { return "data-rate-interface" }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) SetParent(parent types.Entity) { dataRateInterface.parent = parent }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetParent() types.Entity { return dataRateInterface.parent }

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates
// Interface Data Rates threshold templates
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Data Rates threshold template instance. The type is slice of
    // PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate.
    DataRateInterfaceTemplate []PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetFilter() yfilter.YFilter { return dataRateInterfaceTemplates.YFilter }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) SetFilter(yf yfilter.YFilter) { dataRateInterfaceTemplates.YFilter = yf }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetGoName(yname string) string {
    if yname == "data-rate-interface-template" { return "DataRateInterfaceTemplate" }
    return ""
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetSegmentPath() string {
    return "data-rate-interface-templates"
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-rate-interface-template" {
        for _, c := range dataRateInterfaceTemplates.DataRateInterfaceTemplate {
            if dataRateInterfaceTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate{}
        dataRateInterfaceTemplates.DataRateInterfaceTemplate = append(dataRateInterfaceTemplates.DataRateInterfaceTemplate, child)
        return &dataRateInterfaceTemplates.DataRateInterfaceTemplate[len(dataRateInterfaceTemplates.DataRateInterfaceTemplate)-1]
    }
    return nil
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dataRateInterfaceTemplates.DataRateInterfaceTemplate {
        children[dataRateInterfaceTemplates.DataRateInterfaceTemplate[i].GetSegmentPath()] = &dataRateInterfaceTemplates.DataRateInterfaceTemplate[i]
    }
    return children
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetYangName() string { return "data-rate-interface-templates" }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) SetParent(parent types.Entity) { dataRateInterfaceTemplates.parent = parent }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetParent() types.Entity { return dataRateInterfaceTemplates.parent }

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetParentYangName() string { return "data-rate-interface" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
// Interface Data Rates threshold template
// instance
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Enable instance filtering by regular expression. The type is string with
    // length: 1..32.
    RegExpGroup interface{}

    // Enable instance filtering by VRF name regular expression . The type is
    // string with length: 1..32.
    VrfGroup interface{}

    // Input data rate in kbps.
    InputDataRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate

    // Bandwidth in kbps.
    Bandwidth PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth

    // Number of Output packets per second.
    OutputPacketRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate

    // Maximum number of input packets per second.
    InputPeakPkts PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts

    // Peak output data rate in kbps.
    OutputPeakRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate

    // Output data rate in kbps.
    OutputDataRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate

    // Number of input packets per second.
    InputPacketRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate

    // Maximum number of output packets per second.
    OutputPeakPkts PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts

    // Peak input data rate in kbps.
    InputPeakRate PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetFilter() yfilter.YFilter { return dataRateInterfaceTemplate.YFilter }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) SetFilter(yf yfilter.YFilter) { dataRateInterfaceTemplate.YFilter = yf }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "reg-exp-group" { return "RegExpGroup" }
    if yname == "vrf-group" { return "VrfGroup" }
    if yname == "input-data-rate" { return "InputDataRate" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "output-packet-rate" { return "OutputPacketRate" }
    if yname == "input-peak-pkts" { return "InputPeakPkts" }
    if yname == "output-peak-rate" { return "OutputPeakRate" }
    if yname == "output-data-rate" { return "OutputDataRate" }
    if yname == "input-packet-rate" { return "InputPacketRate" }
    if yname == "output-peak-pkts" { return "OutputPeakPkts" }
    if yname == "input-peak-rate" { return "InputPeakRate" }
    return ""
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetSegmentPath() string {
    return "data-rate-interface-template" + "[template-name='" + fmt.Sprintf("%v", dataRateInterfaceTemplate.TemplateName) + "']"
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-data-rate" {
        return &dataRateInterfaceTemplate.InputDataRate
    }
    if childYangName == "bandwidth" {
        return &dataRateInterfaceTemplate.Bandwidth
    }
    if childYangName == "output-packet-rate" {
        return &dataRateInterfaceTemplate.OutputPacketRate
    }
    if childYangName == "input-peak-pkts" {
        return &dataRateInterfaceTemplate.InputPeakPkts
    }
    if childYangName == "output-peak-rate" {
        return &dataRateInterfaceTemplate.OutputPeakRate
    }
    if childYangName == "output-data-rate" {
        return &dataRateInterfaceTemplate.OutputDataRate
    }
    if childYangName == "input-packet-rate" {
        return &dataRateInterfaceTemplate.InputPacketRate
    }
    if childYangName == "output-peak-pkts" {
        return &dataRateInterfaceTemplate.OutputPeakPkts
    }
    if childYangName == "input-peak-rate" {
        return &dataRateInterfaceTemplate.InputPeakRate
    }
    return nil
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-data-rate"] = &dataRateInterfaceTemplate.InputDataRate
    children["bandwidth"] = &dataRateInterfaceTemplate.Bandwidth
    children["output-packet-rate"] = &dataRateInterfaceTemplate.OutputPacketRate
    children["input-peak-pkts"] = &dataRateInterfaceTemplate.InputPeakPkts
    children["output-peak-rate"] = &dataRateInterfaceTemplate.OutputPeakRate
    children["output-data-rate"] = &dataRateInterfaceTemplate.OutputDataRate
    children["input-packet-rate"] = &dataRateInterfaceTemplate.InputPacketRate
    children["output-peak-pkts"] = &dataRateInterfaceTemplate.OutputPeakPkts
    children["input-peak-rate"] = &dataRateInterfaceTemplate.InputPeakRate
    return children
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = dataRateInterfaceTemplate.TemplateName
    leafs["sample-interval"] = dataRateInterfaceTemplate.SampleInterval
    leafs["reg-exp-group"] = dataRateInterfaceTemplate.RegExpGroup
    leafs["vrf-group"] = dataRateInterfaceTemplate.VrfGroup
    return leafs
}

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetYangName() string { return "data-rate-interface-template" }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) SetParent(parent types.Entity) { dataRateInterfaceTemplate.parent = parent }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetParent() types.Entity { return dataRateInterfaceTemplate.parent }

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetParentYangName() string { return "data-rate-interface-templates" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate
// Input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetFilter() yfilter.YFilter { return inputDataRate.YFilter }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) SetFilter(yf yfilter.YFilter) { inputDataRate.YFilter = yf }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetSegmentPath() string {
    return "input-data-rate"
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputDataRate.Operator
    leafs["value"] = inputDataRate.Value
    leafs["end-range-value"] = inputDataRate.EndRangeValue
    leafs["percent"] = inputDataRate.Percent
    leafs["rearm-type"] = inputDataRate.RearmType
    leafs["rearm-window"] = inputDataRate.RearmWindow
    return leafs
}

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetBundleName() string { return "cisco_ios_xr" }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetYangName() string { return "input-data-rate" }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) SetParent(parent types.Entity) { inputDataRate.parent = parent }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetParent() types.Entity { return inputDataRate.parent }

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth
// Bandwidth in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = bandwidth.Operator
    leafs["value"] = bandwidth.Value
    leafs["end-range-value"] = bandwidth.EndRangeValue
    leafs["percent"] = bandwidth.Percent
    leafs["rearm-type"] = bandwidth.RearmType
    leafs["rearm-window"] = bandwidth.RearmWindow
    return leafs
}

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate
// Number of Output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetFilter() yfilter.YFilter { return outputPacketRate.YFilter }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) SetFilter(yf yfilter.YFilter) { outputPacketRate.YFilter = yf }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetSegmentPath() string {
    return "output-packet-rate"
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputPacketRate.Operator
    leafs["value"] = outputPacketRate.Value
    leafs["end-range-value"] = outputPacketRate.EndRangeValue
    leafs["percent"] = outputPacketRate.Percent
    leafs["rearm-type"] = outputPacketRate.RearmType
    leafs["rearm-window"] = outputPacketRate.RearmWindow
    return leafs
}

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetBundleName() string { return "cisco_ios_xr" }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetYangName() string { return "output-packet-rate" }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) SetParent(parent types.Entity) { outputPacketRate.parent = parent }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetParent() types.Entity { return outputPacketRate.parent }

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts
// Maximum number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetFilter() yfilter.YFilter { return inputPeakPkts.YFilter }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) SetFilter(yf yfilter.YFilter) { inputPeakPkts.YFilter = yf }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetSegmentPath() string {
    return "input-peak-pkts"
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputPeakPkts.Operator
    leafs["value"] = inputPeakPkts.Value
    leafs["end-range-value"] = inputPeakPkts.EndRangeValue
    leafs["percent"] = inputPeakPkts.Percent
    leafs["rearm-type"] = inputPeakPkts.RearmType
    leafs["rearm-window"] = inputPeakPkts.RearmWindow
    return leafs
}

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetBundleName() string { return "cisco_ios_xr" }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetYangName() string { return "input-peak-pkts" }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) SetParent(parent types.Entity) { inputPeakPkts.parent = parent }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetParent() types.Entity { return inputPeakPkts.parent }

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate
// Peak output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetFilter() yfilter.YFilter { return outputPeakRate.YFilter }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) SetFilter(yf yfilter.YFilter) { outputPeakRate.YFilter = yf }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetSegmentPath() string {
    return "output-peak-rate"
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputPeakRate.Operator
    leafs["value"] = outputPeakRate.Value
    leafs["end-range-value"] = outputPeakRate.EndRangeValue
    leafs["percent"] = outputPeakRate.Percent
    leafs["rearm-type"] = outputPeakRate.RearmType
    leafs["rearm-window"] = outputPeakRate.RearmWindow
    return leafs
}

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetYangName() string { return "output-peak-rate" }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) SetParent(parent types.Entity) { outputPeakRate.parent = parent }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetParent() types.Entity { return outputPeakRate.parent }

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate
// Output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetFilter() yfilter.YFilter { return outputDataRate.YFilter }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) SetFilter(yf yfilter.YFilter) { outputDataRate.YFilter = yf }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetSegmentPath() string {
    return "output-data-rate"
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputDataRate.Operator
    leafs["value"] = outputDataRate.Value
    leafs["end-range-value"] = outputDataRate.EndRangeValue
    leafs["percent"] = outputDataRate.Percent
    leafs["rearm-type"] = outputDataRate.RearmType
    leafs["rearm-window"] = outputDataRate.RearmWindow
    return leafs
}

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetBundleName() string { return "cisco_ios_xr" }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetYangName() string { return "output-data-rate" }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) SetParent(parent types.Entity) { outputDataRate.parent = parent }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetParent() types.Entity { return outputDataRate.parent }

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate
// Number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetFilter() yfilter.YFilter { return inputPacketRate.YFilter }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) SetFilter(yf yfilter.YFilter) { inputPacketRate.YFilter = yf }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetSegmentPath() string {
    return "input-packet-rate"
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputPacketRate.Operator
    leafs["value"] = inputPacketRate.Value
    leafs["end-range-value"] = inputPacketRate.EndRangeValue
    leafs["percent"] = inputPacketRate.Percent
    leafs["rearm-type"] = inputPacketRate.RearmType
    leafs["rearm-window"] = inputPacketRate.RearmWindow
    return leafs
}

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetBundleName() string { return "cisco_ios_xr" }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetYangName() string { return "input-packet-rate" }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) SetParent(parent types.Entity) { inputPacketRate.parent = parent }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetParent() types.Entity { return inputPacketRate.parent }

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts
// Maximum number of output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetFilter() yfilter.YFilter { return outputPeakPkts.YFilter }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) SetFilter(yf yfilter.YFilter) { outputPeakPkts.YFilter = yf }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetSegmentPath() string {
    return "output-peak-pkts"
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputPeakPkts.Operator
    leafs["value"] = outputPeakPkts.Value
    leafs["end-range-value"] = outputPeakPkts.EndRangeValue
    leafs["percent"] = outputPeakPkts.Percent
    leafs["rearm-type"] = outputPeakPkts.RearmType
    leafs["rearm-window"] = outputPeakPkts.RearmWindow
    return leafs
}

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetBundleName() string { return "cisco_ios_xr" }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetYangName() string { return "output-peak-pkts" }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) SetParent(parent types.Entity) { outputPeakPkts.parent = parent }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetParent() types.Entity { return outputPeakPkts.parent }

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate
// Peak input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetFilter() yfilter.YFilter { return inputPeakRate.YFilter }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) SetFilter(yf yfilter.YFilter) { inputPeakRate.YFilter = yf }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetSegmentPath() string {
    return "input-peak-rate"
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputPeakRate.Operator
    leafs["value"] = inputPeakRate.Value
    leafs["end-range-value"] = inputPeakRate.EndRangeValue
    leafs["percent"] = inputPeakRate.Percent
    leafs["rearm-type"] = inputPeakRate.RearmType
    leafs["rearm-window"] = inputPeakRate.RearmWindow
    return leafs
}

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetYangName() string { return "input-peak-rate" }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) SetParent(parent types.Entity) { inputPeakRate.parent = parent }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetParent() types.Entity { return inputPeakRate.parent }

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetParentYangName() string { return "data-rate-interface-template" }

// PerfMgmt_Threshold_ProcessNode
// Node Process threshold configuration
type PerfMgmt_Threshold_ProcessNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory threshold templates.
    ProcessNodeTemplates PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetFilter() yfilter.YFilter { return processNode.YFilter }

func (processNode *PerfMgmt_Threshold_ProcessNode) SetFilter(yf yfilter.YFilter) { processNode.YFilter = yf }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetGoName(yname string) string {
    if yname == "process-node-templates" { return "ProcessNodeTemplates" }
    return ""
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetSegmentPath() string {
    return "process-node"
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-node-templates" {
        return &processNode.ProcessNodeTemplates
    }
    return nil
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["process-node-templates"] = &processNode.ProcessNodeTemplates
    return children
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetBundleName() string { return "cisco_ios_xr" }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetYangName() string { return "process-node" }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNode *PerfMgmt_Threshold_ProcessNode) SetParent(parent types.Entity) { processNode.parent = parent }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetParent() types.Entity { return processNode.parent }

func (processNode *PerfMgmt_Threshold_ProcessNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates
// Node Memory threshold templates
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory threshold template instance. The type is slice of
    // PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate.
    ProcessNodeTemplate []PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetFilter() yfilter.YFilter { return processNodeTemplates.YFilter }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) SetFilter(yf yfilter.YFilter) { processNodeTemplates.YFilter = yf }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetGoName(yname string) string {
    if yname == "process-node-template" { return "ProcessNodeTemplate" }
    return ""
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetSegmentPath() string {
    return "process-node-templates"
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "process-node-template" {
        for _, c := range processNodeTemplates.ProcessNodeTemplate {
            if processNodeTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate{}
        processNodeTemplates.ProcessNodeTemplate = append(processNodeTemplates.ProcessNodeTemplate, child)
        return &processNodeTemplates.ProcessNodeTemplate[len(processNodeTemplates.ProcessNodeTemplate)-1]
    }
    return nil
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range processNodeTemplates.ProcessNodeTemplate {
        children[processNodeTemplates.ProcessNodeTemplate[i].GetSegmentPath()] = &processNodeTemplates.ProcessNodeTemplate[i]
    }
    return children
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetYangName() string { return "process-node-templates" }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) SetParent(parent types.Entity) { processNodeTemplates.parent = parent }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetParent() types.Entity { return processNodeTemplates.parent }

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetParentYangName() string { return "process-node" }

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
// Node Memory threshold template instance
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Average %CPU utilization.
    AverageCpuUsed PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed

    // Max memory (KBytes) used since startup time.
    PeakMemory PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory

    // Number of threads.
    NoThreads PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetFilter() yfilter.YFilter { return processNodeTemplate.YFilter }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) SetFilter(yf yfilter.YFilter) { processNodeTemplate.YFilter = yf }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "average-cpu-used" { return "AverageCpuUsed" }
    if yname == "peak-memory" { return "PeakMemory" }
    if yname == "no-threads" { return "NoThreads" }
    return ""
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetSegmentPath() string {
    return "process-node-template" + "[template-name='" + fmt.Sprintf("%v", processNodeTemplate.TemplateName) + "']"
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "average-cpu-used" {
        return &processNodeTemplate.AverageCpuUsed
    }
    if childYangName == "peak-memory" {
        return &processNodeTemplate.PeakMemory
    }
    if childYangName == "no-threads" {
        return &processNodeTemplate.NoThreads
    }
    return nil
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["average-cpu-used"] = &processNodeTemplate.AverageCpuUsed
    children["peak-memory"] = &processNodeTemplate.PeakMemory
    children["no-threads"] = &processNodeTemplate.NoThreads
    return children
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = processNodeTemplate.TemplateName
    leafs["sample-interval"] = processNodeTemplate.SampleInterval
    return leafs
}

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetYangName() string { return "process-node-template" }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) SetParent(parent types.Entity) { processNodeTemplate.parent = parent }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetParent() types.Entity { return processNodeTemplate.parent }

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetParentYangName() string { return "process-node-templates" }

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..100.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..100.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetFilter() yfilter.YFilter { return averageCpuUsed.YFilter }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) SetFilter(yf yfilter.YFilter) { averageCpuUsed.YFilter = yf }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetSegmentPath() string {
    return "average-cpu-used"
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = averageCpuUsed.Operator
    leafs["value"] = averageCpuUsed.Value
    leafs["end-range-value"] = averageCpuUsed.EndRangeValue
    leafs["percent"] = averageCpuUsed.Percent
    leafs["rearm-type"] = averageCpuUsed.RearmType
    leafs["rearm-window"] = averageCpuUsed.RearmWindow
    return leafs
}

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetBundleName() string { return "cisco_ios_xr" }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetYangName() string { return "average-cpu-used" }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) SetParent(parent types.Entity) { averageCpuUsed.parent = parent }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetParent() types.Entity { return averageCpuUsed.parent }

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetParentYangName() string { return "process-node-template" }

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory
// Max memory (KBytes) used since startup time
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetFilter() yfilter.YFilter { return peakMemory.YFilter }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) SetFilter(yf yfilter.YFilter) { peakMemory.YFilter = yf }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetSegmentPath() string {
    return "peak-memory"
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = peakMemory.Operator
    leafs["value"] = peakMemory.Value
    leafs["end-range-value"] = peakMemory.EndRangeValue
    leafs["percent"] = peakMemory.Percent
    leafs["rearm-type"] = peakMemory.RearmType
    leafs["rearm-window"] = peakMemory.RearmWindow
    return leafs
}

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetBundleName() string { return "cisco_ios_xr" }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetYangName() string { return "peak-memory" }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) SetParent(parent types.Entity) { peakMemory.parent = parent }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetParent() types.Entity { return peakMemory.parent }

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetParentYangName() string { return "process-node-template" }

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads
// Number of threads
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..32767.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..32767.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetFilter() yfilter.YFilter { return noThreads.YFilter }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) SetFilter(yf yfilter.YFilter) { noThreads.YFilter = yf }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetSegmentPath() string {
    return "no-threads"
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = noThreads.Operator
    leafs["value"] = noThreads.Value
    leafs["end-range-value"] = noThreads.EndRangeValue
    leafs["percent"] = noThreads.Percent
    leafs["rearm-type"] = noThreads.RearmType
    leafs["rearm-window"] = noThreads.RearmWindow
    return leafs
}

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetBundleName() string { return "cisco_ios_xr" }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetYangName() string { return "no-threads" }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) SetParent(parent types.Entity) { noThreads.parent = parent }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetParent() types.Entity { return noThreads.parent }

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetParentYangName() string { return "process-node-template" }

// PerfMgmt_Threshold_MemoryNode
// Node Memory threshold configuration
type PerfMgmt_Threshold_MemoryNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory threshold configuration templates.
    MemoryNodeTemplates PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetFilter() yfilter.YFilter { return memoryNode.YFilter }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) SetFilter(yf yfilter.YFilter) { memoryNode.YFilter = yf }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetGoName(yname string) string {
    if yname == "memory-node-templates" { return "MemoryNodeTemplates" }
    return ""
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetSegmentPath() string {
    return "memory-node"
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-node-templates" {
        return &memoryNode.MemoryNodeTemplates
    }
    return nil
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["memory-node-templates"] = &memoryNode.MemoryNodeTemplates
    return children
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetYangName() string { return "memory-node" }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) SetParent(parent types.Entity) { memoryNode.parent = parent }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetParent() types.Entity { return memoryNode.parent }

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates
// Node Memory threshold configuration templates
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Memory threshold configuration template instance. The type is slice of
    // PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate.
    MemoryNodeTemplate []PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetFilter() yfilter.YFilter { return memoryNodeTemplates.YFilter }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) SetFilter(yf yfilter.YFilter) { memoryNodeTemplates.YFilter = yf }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetGoName(yname string) string {
    if yname == "memory-node-template" { return "MemoryNodeTemplate" }
    return ""
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetSegmentPath() string {
    return "memory-node-templates"
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "memory-node-template" {
        for _, c := range memoryNodeTemplates.MemoryNodeTemplate {
            if memoryNodeTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate{}
        memoryNodeTemplates.MemoryNodeTemplate = append(memoryNodeTemplates.MemoryNodeTemplate, child)
        return &memoryNodeTemplates.MemoryNodeTemplate[len(memoryNodeTemplates.MemoryNodeTemplate)-1]
    }
    return nil
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memoryNodeTemplates.MemoryNodeTemplate {
        children[memoryNodeTemplates.MemoryNodeTemplate[i].GetSegmentPath()] = &memoryNodeTemplates.MemoryNodeTemplate[i]
    }
    return children
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetYangName() string { return "memory-node-templates" }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) SetParent(parent types.Entity) { memoryNodeTemplates.parent = parent }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetParent() types.Entity { return memoryNodeTemplates.parent }

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetParentYangName() string { return "memory-node" }

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
// Node Memory threshold configuration template
// instance
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Maximum memory (KBytes) used.
    PeakMemory PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory

    // Current memory (Bytes) in use.
    CurrMemory PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetFilter() yfilter.YFilter { return memoryNodeTemplate.YFilter }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) SetFilter(yf yfilter.YFilter) { memoryNodeTemplate.YFilter = yf }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "peak-memory" { return "PeakMemory" }
    if yname == "curr-memory" { return "CurrMemory" }
    return ""
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetSegmentPath() string {
    return "memory-node-template" + "[template-name='" + fmt.Sprintf("%v", memoryNodeTemplate.TemplateName) + "']"
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peak-memory" {
        return &memoryNodeTemplate.PeakMemory
    }
    if childYangName == "curr-memory" {
        return &memoryNodeTemplate.CurrMemory
    }
    return nil
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peak-memory"] = &memoryNodeTemplate.PeakMemory
    children["curr-memory"] = &memoryNodeTemplate.CurrMemory
    return children
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = memoryNodeTemplate.TemplateName
    leafs["sample-interval"] = memoryNodeTemplate.SampleInterval
    return leafs
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetYangName() string { return "memory-node-template" }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) SetParent(parent types.Entity) { memoryNodeTemplate.parent = parent }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetParent() types.Entity { return memoryNodeTemplate.parent }

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetParentYangName() string { return "memory-node-templates" }

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory
// Maximum memory (KBytes) used
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4194304.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4194304.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetFilter() yfilter.YFilter { return peakMemory.YFilter }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) SetFilter(yf yfilter.YFilter) { peakMemory.YFilter = yf }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetSegmentPath() string {
    return "peak-memory"
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = peakMemory.Operator
    leafs["value"] = peakMemory.Value
    leafs["end-range-value"] = peakMemory.EndRangeValue
    leafs["percent"] = peakMemory.Percent
    leafs["rearm-type"] = peakMemory.RearmType
    leafs["rearm-window"] = peakMemory.RearmWindow
    return leafs
}

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetBundleName() string { return "cisco_ios_xr" }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetYangName() string { return "peak-memory" }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) SetParent(parent types.Entity) { peakMemory.parent = parent }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetParent() types.Entity { return peakMemory.parent }

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetParentYangName() string { return "memory-node-template" }

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory
// Current memory (Bytes) in use
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: -2147483648..2147483647.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: -2147483648..2147483647.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetFilter() yfilter.YFilter { return currMemory.YFilter }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) SetFilter(yf yfilter.YFilter) { currMemory.YFilter = yf }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetSegmentPath() string {
    return "curr-memory"
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = currMemory.Operator
    leafs["value"] = currMemory.Value
    leafs["end-range-value"] = currMemory.EndRangeValue
    leafs["percent"] = currMemory.Percent
    leafs["rearm-type"] = currMemory.RearmType
    leafs["rearm-window"] = currMemory.RearmWindow
    return leafs
}

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetBundleName() string { return "cisco_ios_xr" }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetYangName() string { return "curr-memory" }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) SetParent(parent types.Entity) { currMemory.parent = parent }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetParent() types.Entity { return currMemory.parent }

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetParentYangName() string { return "memory-node-template" }

// PerfMgmt_Threshold_Ospfv3Protocol
// OSPF v2 Protocol threshold configuration
type PerfMgmt_Threshold_Ospfv3Protocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold templates.
    Ospfv3ProtocolTemplates PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetFilter() yfilter.YFilter { return ospfv3Protocol.YFilter }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) SetFilter(yf yfilter.YFilter) { ospfv3Protocol.YFilter = yf }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetGoName(yname string) string {
    if yname == "ospfv3-protocol-templates" { return "Ospfv3ProtocolTemplates" }
    return ""
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetSegmentPath() string {
    return "ospfv3-protocol"
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv3-protocol-templates" {
        return &ospfv3Protocol.Ospfv3ProtocolTemplates
    }
    return nil
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ospfv3-protocol-templates"] = &ospfv3Protocol.Ospfv3ProtocolTemplates
    return children
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetYangName() string { return "ospfv3-protocol" }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) SetParent(parent types.Entity) { ospfv3Protocol.parent = parent }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetParent() types.Entity { return ospfv3Protocol.parent }

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetParentYangName() string { return "threshold" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate.
    Ospfv3ProtocolTemplate []PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetFilter() yfilter.YFilter { return ospfv3ProtocolTemplates.YFilter }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolTemplates.YFilter = yf }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetGoName(yname string) string {
    if yname == "ospfv3-protocol-template" { return "Ospfv3ProtocolTemplate" }
    return ""
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetSegmentPath() string {
    return "ospfv3-protocol-templates"
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ospfv3-protocol-template" {
        for _, c := range ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate {
            if ospfv3ProtocolTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate{}
        ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate = append(ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate, child)
        return &ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[len(ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate)-1]
    }
    return nil
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate {
        children[ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i].GetSegmentPath()] = &ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i]
    }
    return children
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetYangName() string { return "ospfv3-protocol-templates" }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) SetParent(parent types.Entity) { ospfv3ProtocolTemplates.parent = parent }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetParent() types.Entity { return ospfv3ProtocolTemplates.parent }

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetParentYangName() string { return "ospfv3-protocol" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Number of LSA received in LSA Acknowledgements.
    InputLsaAcksLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa

    // Number of LSA sent in DBD packets.
    OutputDbDsLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa

    // Number of LSA received in DBD packets.
    InputDbDsLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa

    // Number of LSA Updates received.
    InputLsaUpdates PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates

    // Number of DBD packets sent.
    OutputDbDs PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs

    // Number of LSA sent in LSA Updates.
    OutputLsaUpdatesLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa

    // Number of DBD packets received.
    InputDbDs PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs

    // Number of LSA received in LSA Updates.
    InputLsaUpdatesLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa

    // Total number of packets sent.
    OutputPackets PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets

    // Total number of packets received.
    InputPackets PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets

    // Total number of packets sent.
    OutputHelloPackets PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets

    // Number of Hello packets received.
    InputHelloPackets PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets

    // Number of LS Requests sent.
    OutputLsRequests PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests

    // Number of LSA sent in LSA Acknowledgements.
    OutputLsaAcksLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa

    // Number of LSA Acknowledgements sent.
    OutputLsaAcks PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks

    // Number of LSA Acknowledgements received.
    InputLsaAcks PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks

    // Number of LSA Updates sent.
    OutputLsaUpdates PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates

    // Number of LSA sent in LS Requests.
    OutputLsRequestsLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa

    // Number of LSA received in LS Requests.
    InputLsRequestsLsa PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa

    // Number of LS Requests received.
    InputLsRequests PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetFilter() yfilter.YFilter { return ospfv3ProtocolTemplate.YFilter }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) SetFilter(yf yfilter.YFilter) { ospfv3ProtocolTemplate.YFilter = yf }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "sample-interval" { return "SampleInterval" }
    if yname == "input-lsa-acks-lsa" { return "InputLsaAcksLsa" }
    if yname == "output-db-ds-lsa" { return "OutputDbDsLsa" }
    if yname == "input-db-ds-lsa" { return "InputDbDsLsa" }
    if yname == "input-lsa-updates" { return "InputLsaUpdates" }
    if yname == "output-db-ds" { return "OutputDbDs" }
    if yname == "output-lsa-updates-lsa" { return "OutputLsaUpdatesLsa" }
    if yname == "input-db-ds" { return "InputDbDs" }
    if yname == "input-lsa-updates-lsa" { return "InputLsaUpdatesLsa" }
    if yname == "output-packets" { return "OutputPackets" }
    if yname == "input-packets" { return "InputPackets" }
    if yname == "output-hello-packets" { return "OutputHelloPackets" }
    if yname == "input-hello-packets" { return "InputHelloPackets" }
    if yname == "output-ls-requests" { return "OutputLsRequests" }
    if yname == "output-lsa-acks-lsa" { return "OutputLsaAcksLsa" }
    if yname == "output-lsa-acks" { return "OutputLsaAcks" }
    if yname == "input-lsa-acks" { return "InputLsaAcks" }
    if yname == "output-lsa-updates" { return "OutputLsaUpdates" }
    if yname == "output-ls-requests-lsa" { return "OutputLsRequestsLsa" }
    if yname == "input-ls-requests-lsa" { return "InputLsRequestsLsa" }
    if yname == "input-ls-requests" { return "InputLsRequests" }
    return ""
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetSegmentPath() string {
    return "ospfv3-protocol-template" + "[template-name='" + fmt.Sprintf("%v", ospfv3ProtocolTemplate.TemplateName) + "']"
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-lsa-acks-lsa" {
        return &ospfv3ProtocolTemplate.InputLsaAcksLsa
    }
    if childYangName == "output-db-ds-lsa" {
        return &ospfv3ProtocolTemplate.OutputDbDsLsa
    }
    if childYangName == "input-db-ds-lsa" {
        return &ospfv3ProtocolTemplate.InputDbDsLsa
    }
    if childYangName == "input-lsa-updates" {
        return &ospfv3ProtocolTemplate.InputLsaUpdates
    }
    if childYangName == "output-db-ds" {
        return &ospfv3ProtocolTemplate.OutputDbDs
    }
    if childYangName == "output-lsa-updates-lsa" {
        return &ospfv3ProtocolTemplate.OutputLsaUpdatesLsa
    }
    if childYangName == "input-db-ds" {
        return &ospfv3ProtocolTemplate.InputDbDs
    }
    if childYangName == "input-lsa-updates-lsa" {
        return &ospfv3ProtocolTemplate.InputLsaUpdatesLsa
    }
    if childYangName == "output-packets" {
        return &ospfv3ProtocolTemplate.OutputPackets
    }
    if childYangName == "input-packets" {
        return &ospfv3ProtocolTemplate.InputPackets
    }
    if childYangName == "output-hello-packets" {
        return &ospfv3ProtocolTemplate.OutputHelloPackets
    }
    if childYangName == "input-hello-packets" {
        return &ospfv3ProtocolTemplate.InputHelloPackets
    }
    if childYangName == "output-ls-requests" {
        return &ospfv3ProtocolTemplate.OutputLsRequests
    }
    if childYangName == "output-lsa-acks-lsa" {
        return &ospfv3ProtocolTemplate.OutputLsaAcksLsa
    }
    if childYangName == "output-lsa-acks" {
        return &ospfv3ProtocolTemplate.OutputLsaAcks
    }
    if childYangName == "input-lsa-acks" {
        return &ospfv3ProtocolTemplate.InputLsaAcks
    }
    if childYangName == "output-lsa-updates" {
        return &ospfv3ProtocolTemplate.OutputLsaUpdates
    }
    if childYangName == "output-ls-requests-lsa" {
        return &ospfv3ProtocolTemplate.OutputLsRequestsLsa
    }
    if childYangName == "input-ls-requests-lsa" {
        return &ospfv3ProtocolTemplate.InputLsRequestsLsa
    }
    if childYangName == "input-ls-requests" {
        return &ospfv3ProtocolTemplate.InputLsRequests
    }
    return nil
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-lsa-acks-lsa"] = &ospfv3ProtocolTemplate.InputLsaAcksLsa
    children["output-db-ds-lsa"] = &ospfv3ProtocolTemplate.OutputDbDsLsa
    children["input-db-ds-lsa"] = &ospfv3ProtocolTemplate.InputDbDsLsa
    children["input-lsa-updates"] = &ospfv3ProtocolTemplate.InputLsaUpdates
    children["output-db-ds"] = &ospfv3ProtocolTemplate.OutputDbDs
    children["output-lsa-updates-lsa"] = &ospfv3ProtocolTemplate.OutputLsaUpdatesLsa
    children["input-db-ds"] = &ospfv3ProtocolTemplate.InputDbDs
    children["input-lsa-updates-lsa"] = &ospfv3ProtocolTemplate.InputLsaUpdatesLsa
    children["output-packets"] = &ospfv3ProtocolTemplate.OutputPackets
    children["input-packets"] = &ospfv3ProtocolTemplate.InputPackets
    children["output-hello-packets"] = &ospfv3ProtocolTemplate.OutputHelloPackets
    children["input-hello-packets"] = &ospfv3ProtocolTemplate.InputHelloPackets
    children["output-ls-requests"] = &ospfv3ProtocolTemplate.OutputLsRequests
    children["output-lsa-acks-lsa"] = &ospfv3ProtocolTemplate.OutputLsaAcksLsa
    children["output-lsa-acks"] = &ospfv3ProtocolTemplate.OutputLsaAcks
    children["input-lsa-acks"] = &ospfv3ProtocolTemplate.InputLsaAcks
    children["output-lsa-updates"] = &ospfv3ProtocolTemplate.OutputLsaUpdates
    children["output-ls-requests-lsa"] = &ospfv3ProtocolTemplate.OutputLsRequestsLsa
    children["input-ls-requests-lsa"] = &ospfv3ProtocolTemplate.InputLsRequestsLsa
    children["input-ls-requests"] = &ospfv3ProtocolTemplate.InputLsRequests
    return children
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ospfv3ProtocolTemplate.TemplateName
    leafs["sample-interval"] = ospfv3ProtocolTemplate.SampleInterval
    return leafs
}

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetYangName() string { return "ospfv3-protocol-template" }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) SetParent(parent types.Entity) { ospfv3ProtocolTemplate.parent = parent }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetParent() types.Entity { return ospfv3ProtocolTemplate.parent }

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetParentYangName() string { return "ospfv3-protocol-templates" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetFilter() yfilter.YFilter { return inputLsaAcksLsa.YFilter }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) SetFilter(yf yfilter.YFilter) { inputLsaAcksLsa.YFilter = yf }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetSegmentPath() string {
    return "input-lsa-acks-lsa"
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaAcksLsa.Operator
    leafs["value"] = inputLsaAcksLsa.Value
    leafs["end-range-value"] = inputLsaAcksLsa.EndRangeValue
    leafs["percent"] = inputLsaAcksLsa.Percent
    leafs["rearm-type"] = inputLsaAcksLsa.RearmType
    leafs["rearm-window"] = inputLsaAcksLsa.RearmWindow
    return leafs
}

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetYangName() string { return "input-lsa-acks-lsa" }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) SetParent(parent types.Entity) { inputLsaAcksLsa.parent = parent }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetParent() types.Entity { return inputLsaAcksLsa.parent }

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetFilter() yfilter.YFilter { return outputDbDsLsa.YFilter }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) SetFilter(yf yfilter.YFilter) { outputDbDsLsa.YFilter = yf }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetSegmentPath() string {
    return "output-db-ds-lsa"
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputDbDsLsa.Operator
    leafs["value"] = outputDbDsLsa.Value
    leafs["end-range-value"] = outputDbDsLsa.EndRangeValue
    leafs["percent"] = outputDbDsLsa.Percent
    leafs["rearm-type"] = outputDbDsLsa.RearmType
    leafs["rearm-window"] = outputDbDsLsa.RearmWindow
    return leafs
}

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetYangName() string { return "output-db-ds-lsa" }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) SetParent(parent types.Entity) { outputDbDsLsa.parent = parent }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetParent() types.Entity { return outputDbDsLsa.parent }

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetFilter() yfilter.YFilter { return inputDbDsLsa.YFilter }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) SetFilter(yf yfilter.YFilter) { inputDbDsLsa.YFilter = yf }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetSegmentPath() string {
    return "input-db-ds-lsa"
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputDbDsLsa.Operator
    leafs["value"] = inputDbDsLsa.Value
    leafs["end-range-value"] = inputDbDsLsa.EndRangeValue
    leafs["percent"] = inputDbDsLsa.Percent
    leafs["rearm-type"] = inputDbDsLsa.RearmType
    leafs["rearm-window"] = inputDbDsLsa.RearmWindow
    return leafs
}

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetYangName() string { return "input-db-ds-lsa" }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) SetParent(parent types.Entity) { inputDbDsLsa.parent = parent }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetParent() types.Entity { return inputDbDsLsa.parent }

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetFilter() yfilter.YFilter { return inputLsaUpdates.YFilter }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) SetFilter(yf yfilter.YFilter) { inputLsaUpdates.YFilter = yf }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetSegmentPath() string {
    return "input-lsa-updates"
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaUpdates.Operator
    leafs["value"] = inputLsaUpdates.Value
    leafs["end-range-value"] = inputLsaUpdates.EndRangeValue
    leafs["percent"] = inputLsaUpdates.Percent
    leafs["rearm-type"] = inputLsaUpdates.RearmType
    leafs["rearm-window"] = inputLsaUpdates.RearmWindow
    return leafs
}

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetYangName() string { return "input-lsa-updates" }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) SetParent(parent types.Entity) { inputLsaUpdates.parent = parent }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetParent() types.Entity { return inputLsaUpdates.parent }

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetFilter() yfilter.YFilter { return outputDbDs.YFilter }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) SetFilter(yf yfilter.YFilter) { outputDbDs.YFilter = yf }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetSegmentPath() string {
    return "output-db-ds"
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputDbDs.Operator
    leafs["value"] = outputDbDs.Value
    leafs["end-range-value"] = outputDbDs.EndRangeValue
    leafs["percent"] = outputDbDs.Percent
    leafs["rearm-type"] = outputDbDs.RearmType
    leafs["rearm-window"] = outputDbDs.RearmWindow
    return leafs
}

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetBundleName() string { return "cisco_ios_xr" }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetYangName() string { return "output-db-ds" }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) SetParent(parent types.Entity) { outputDbDs.parent = parent }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetParent() types.Entity { return outputDbDs.parent }

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetFilter() yfilter.YFilter { return outputLsaUpdatesLsa.YFilter }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) SetFilter(yf yfilter.YFilter) { outputLsaUpdatesLsa.YFilter = yf }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetSegmentPath() string {
    return "output-lsa-updates-lsa"
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaUpdatesLsa.Operator
    leafs["value"] = outputLsaUpdatesLsa.Value
    leafs["end-range-value"] = outputLsaUpdatesLsa.EndRangeValue
    leafs["percent"] = outputLsaUpdatesLsa.Percent
    leafs["rearm-type"] = outputLsaUpdatesLsa.RearmType
    leafs["rearm-window"] = outputLsaUpdatesLsa.RearmWindow
    return leafs
}

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetYangName() string { return "output-lsa-updates-lsa" }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) SetParent(parent types.Entity) { outputLsaUpdatesLsa.parent = parent }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetParent() types.Entity { return outputLsaUpdatesLsa.parent }

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetFilter() yfilter.YFilter { return inputDbDs.YFilter }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) SetFilter(yf yfilter.YFilter) { inputDbDs.YFilter = yf }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetSegmentPath() string {
    return "input-db-ds"
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputDbDs.Operator
    leafs["value"] = inputDbDs.Value
    leafs["end-range-value"] = inputDbDs.EndRangeValue
    leafs["percent"] = inputDbDs.Percent
    leafs["rearm-type"] = inputDbDs.RearmType
    leafs["rearm-window"] = inputDbDs.RearmWindow
    return leafs
}

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetBundleName() string { return "cisco_ios_xr" }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetYangName() string { return "input-db-ds" }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) SetParent(parent types.Entity) { inputDbDs.parent = parent }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetParent() types.Entity { return inputDbDs.parent }

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetFilter() yfilter.YFilter { return inputLsaUpdatesLsa.YFilter }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) SetFilter(yf yfilter.YFilter) { inputLsaUpdatesLsa.YFilter = yf }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetSegmentPath() string {
    return "input-lsa-updates-lsa"
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaUpdatesLsa.Operator
    leafs["value"] = inputLsaUpdatesLsa.Value
    leafs["end-range-value"] = inputLsaUpdatesLsa.EndRangeValue
    leafs["percent"] = inputLsaUpdatesLsa.Percent
    leafs["rearm-type"] = inputLsaUpdatesLsa.RearmType
    leafs["rearm-window"] = inputLsaUpdatesLsa.RearmWindow
    return leafs
}

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetYangName() string { return "input-lsa-updates-lsa" }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) SetParent(parent types.Entity) { inputLsaUpdatesLsa.parent = parent }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetParent() types.Entity { return inputLsaUpdatesLsa.parent }

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetFilter() yfilter.YFilter { return outputPackets.YFilter }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) SetFilter(yf yfilter.YFilter) { outputPackets.YFilter = yf }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetSegmentPath() string {
    return "output-packets"
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputPackets.Operator
    leafs["value"] = outputPackets.Value
    leafs["end-range-value"] = outputPackets.EndRangeValue
    leafs["percent"] = outputPackets.Percent
    leafs["rearm-type"] = outputPackets.RearmType
    leafs["rearm-window"] = outputPackets.RearmWindow
    return leafs
}

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetYangName() string { return "output-packets" }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) SetParent(parent types.Entity) { outputPackets.parent = parent }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetParent() types.Entity { return outputPackets.parent }

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetFilter() yfilter.YFilter { return inputPackets.YFilter }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) SetFilter(yf yfilter.YFilter) { inputPackets.YFilter = yf }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetSegmentPath() string {
    return "input-packets"
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputPackets.Operator
    leafs["value"] = inputPackets.Value
    leafs["end-range-value"] = inputPackets.EndRangeValue
    leafs["percent"] = inputPackets.Percent
    leafs["rearm-type"] = inputPackets.RearmType
    leafs["rearm-window"] = inputPackets.RearmWindow
    return leafs
}

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetYangName() string { return "input-packets" }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) SetParent(parent types.Entity) { inputPackets.parent = parent }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetParent() types.Entity { return inputPackets.parent }

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetFilter() yfilter.YFilter { return outputHelloPackets.YFilter }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) SetFilter(yf yfilter.YFilter) { outputHelloPackets.YFilter = yf }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetSegmentPath() string {
    return "output-hello-packets"
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputHelloPackets.Operator
    leafs["value"] = outputHelloPackets.Value
    leafs["end-range-value"] = outputHelloPackets.EndRangeValue
    leafs["percent"] = outputHelloPackets.Percent
    leafs["rearm-type"] = outputHelloPackets.RearmType
    leafs["rearm-window"] = outputHelloPackets.RearmWindow
    return leafs
}

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetBundleName() string { return "cisco_ios_xr" }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetYangName() string { return "output-hello-packets" }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) SetParent(parent types.Entity) { outputHelloPackets.parent = parent }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetParent() types.Entity { return outputHelloPackets.parent }

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetFilter() yfilter.YFilter { return inputHelloPackets.YFilter }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) SetFilter(yf yfilter.YFilter) { inputHelloPackets.YFilter = yf }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetSegmentPath() string {
    return "input-hello-packets"
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputHelloPackets.Operator
    leafs["value"] = inputHelloPackets.Value
    leafs["end-range-value"] = inputHelloPackets.EndRangeValue
    leafs["percent"] = inputHelloPackets.Percent
    leafs["rearm-type"] = inputHelloPackets.RearmType
    leafs["rearm-window"] = inputHelloPackets.RearmWindow
    return leafs
}

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetBundleName() string { return "cisco_ios_xr" }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetYangName() string { return "input-hello-packets" }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) SetParent(parent types.Entity) { inputHelloPackets.parent = parent }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetParent() types.Entity { return inputHelloPackets.parent }

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetFilter() yfilter.YFilter { return outputLsRequests.YFilter }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) SetFilter(yf yfilter.YFilter) { outputLsRequests.YFilter = yf }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetSegmentPath() string {
    return "output-ls-requests"
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsRequests.Operator
    leafs["value"] = outputLsRequests.Value
    leafs["end-range-value"] = outputLsRequests.EndRangeValue
    leafs["percent"] = outputLsRequests.Percent
    leafs["rearm-type"] = outputLsRequests.RearmType
    leafs["rearm-window"] = outputLsRequests.RearmWindow
    return leafs
}

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetYangName() string { return "output-ls-requests" }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) SetParent(parent types.Entity) { outputLsRequests.parent = parent }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetParent() types.Entity { return outputLsRequests.parent }

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetFilter() yfilter.YFilter { return outputLsaAcksLsa.YFilter }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) SetFilter(yf yfilter.YFilter) { outputLsaAcksLsa.YFilter = yf }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetSegmentPath() string {
    return "output-lsa-acks-lsa"
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaAcksLsa.Operator
    leafs["value"] = outputLsaAcksLsa.Value
    leafs["end-range-value"] = outputLsaAcksLsa.EndRangeValue
    leafs["percent"] = outputLsaAcksLsa.Percent
    leafs["rearm-type"] = outputLsaAcksLsa.RearmType
    leafs["rearm-window"] = outputLsaAcksLsa.RearmWindow
    return leafs
}

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetYangName() string { return "output-lsa-acks-lsa" }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) SetParent(parent types.Entity) { outputLsaAcksLsa.parent = parent }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetParent() types.Entity { return outputLsaAcksLsa.parent }

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetFilter() yfilter.YFilter { return outputLsaAcks.YFilter }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) SetFilter(yf yfilter.YFilter) { outputLsaAcks.YFilter = yf }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetSegmentPath() string {
    return "output-lsa-acks"
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaAcks.Operator
    leafs["value"] = outputLsaAcks.Value
    leafs["end-range-value"] = outputLsaAcks.EndRangeValue
    leafs["percent"] = outputLsaAcks.Percent
    leafs["rearm-type"] = outputLsaAcks.RearmType
    leafs["rearm-window"] = outputLsaAcks.RearmWindow
    return leafs
}

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetYangName() string { return "output-lsa-acks" }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) SetParent(parent types.Entity) { outputLsaAcks.parent = parent }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetParent() types.Entity { return outputLsaAcks.parent }

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetFilter() yfilter.YFilter { return inputLsaAcks.YFilter }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) SetFilter(yf yfilter.YFilter) { inputLsaAcks.YFilter = yf }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetSegmentPath() string {
    return "input-lsa-acks"
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsaAcks.Operator
    leafs["value"] = inputLsaAcks.Value
    leafs["end-range-value"] = inputLsaAcks.EndRangeValue
    leafs["percent"] = inputLsaAcks.Percent
    leafs["rearm-type"] = inputLsaAcks.RearmType
    leafs["rearm-window"] = inputLsaAcks.RearmWindow
    return leafs
}

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetYangName() string { return "input-lsa-acks" }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) SetParent(parent types.Entity) { inputLsaAcks.parent = parent }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetParent() types.Entity { return inputLsaAcks.parent }

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetFilter() yfilter.YFilter { return outputLsaUpdates.YFilter }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) SetFilter(yf yfilter.YFilter) { outputLsaUpdates.YFilter = yf }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetSegmentPath() string {
    return "output-lsa-updates"
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsaUpdates.Operator
    leafs["value"] = outputLsaUpdates.Value
    leafs["end-range-value"] = outputLsaUpdates.EndRangeValue
    leafs["percent"] = outputLsaUpdates.Percent
    leafs["rearm-type"] = outputLsaUpdates.RearmType
    leafs["rearm-window"] = outputLsaUpdates.RearmWindow
    return leafs
}

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetYangName() string { return "output-lsa-updates" }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) SetParent(parent types.Entity) { outputLsaUpdates.parent = parent }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetParent() types.Entity { return outputLsaUpdates.parent }

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetFilter() yfilter.YFilter { return outputLsRequestsLsa.YFilter }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) SetFilter(yf yfilter.YFilter) { outputLsRequestsLsa.YFilter = yf }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetSegmentPath() string {
    return "output-ls-requests-lsa"
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = outputLsRequestsLsa.Operator
    leafs["value"] = outputLsRequestsLsa.Value
    leafs["end-range-value"] = outputLsRequestsLsa.EndRangeValue
    leafs["percent"] = outputLsRequestsLsa.Percent
    leafs["rearm-type"] = outputLsRequestsLsa.RearmType
    leafs["rearm-window"] = outputLsRequestsLsa.RearmWindow
    return leafs
}

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetYangName() string { return "output-ls-requests-lsa" }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) SetParent(parent types.Entity) { outputLsRequestsLsa.parent = parent }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetParent() types.Entity { return outputLsRequestsLsa.parent }

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetFilter() yfilter.YFilter { return inputLsRequestsLsa.YFilter }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) SetFilter(yf yfilter.YFilter) { inputLsRequestsLsa.YFilter = yf }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetSegmentPath() string {
    return "input-ls-requests-lsa"
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsRequestsLsa.Operator
    leafs["value"] = inputLsRequestsLsa.Value
    leafs["end-range-value"] = inputLsRequestsLsa.EndRangeValue
    leafs["percent"] = inputLsRequestsLsa.Percent
    leafs["rearm-type"] = inputLsRequestsLsa.RearmType
    leafs["rearm-window"] = inputLsRequestsLsa.RearmWindow
    return leafs
}

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetYangName() string { return "input-ls-requests-lsa" }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) SetParent(parent types.Entity) { inputLsRequestsLsa.parent = parent }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetParent() types.Entity { return inputLsRequestsLsa.parent }

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetParentYangName() string { return "ospfv3-protocol-template" }

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operator. The type is PmThresholdOp.
    Operator interface{}

    // Threshold value (or start range value for operator RG). The type is
    // interface{} with range: 0..4294967295.
    Value interface{}

    // Threshold end range value (for operator RG, set to 0 otherwise). The type
    // is interface{} with range: 0..4294967295.
    EndRangeValue interface{}

    // Set to TRUE if Specified threshold values are in percent. The type is bool.
    Percent interface{}

    // Configure the Rearm type. The type is PmThresholdRearm.
    RearmType interface{}

    // Configure the rearm window size (for rearm type Window). The type is
    // interface{} with range: 1..100.
    RearmWindow interface{}
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetFilter() yfilter.YFilter { return inputLsRequests.YFilter }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) SetFilter(yf yfilter.YFilter) { inputLsRequests.YFilter = yf }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    if yname == "value" { return "Value" }
    if yname == "end-range-value" { return "EndRangeValue" }
    if yname == "percent" { return "Percent" }
    if yname == "rearm-type" { return "RearmType" }
    if yname == "rearm-window" { return "RearmWindow" }
    return ""
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetSegmentPath() string {
    return "input-ls-requests"
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator"] = inputLsRequests.Operator
    leafs["value"] = inputLsRequests.Value
    leafs["end-range-value"] = inputLsRequests.EndRangeValue
    leafs["percent"] = inputLsRequests.Percent
    leafs["rearm-type"] = inputLsRequests.RearmType
    leafs["rearm-window"] = inputLsRequests.RearmWindow
    return leafs
}

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetBundleName() string { return "cisco_ios_xr" }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetYangName() string { return "input-ls-requests" }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) SetParent(parent types.Entity) { inputLsRequests.parent = parent }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetParent() types.Entity { return inputLsRequests.parent }

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetParentYangName() string { return "ospfv3-protocol-template" }

