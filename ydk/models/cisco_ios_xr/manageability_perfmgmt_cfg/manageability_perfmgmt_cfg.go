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

// PerfMgmt
// Performance Management configuration & operations
type PerfMgmt struct {
    EntityData types.CommonEntityData
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

func (perfMgmt *PerfMgmt) GetEntityData() *types.CommonEntityData {
    perfMgmt.EntityData.YFilter = perfMgmt.YFilter
    perfMgmt.EntityData.YangName = "perf-mgmt"
    perfMgmt.EntityData.BundleName = "cisco_ios_xr"
    perfMgmt.EntityData.ParentYangName = "Cisco-IOS-XR-manageability-perfmgmt-cfg"
    perfMgmt.EntityData.SegmentPath = "Cisco-IOS-XR-manageability-perfmgmt-cfg:perf-mgmt"
    perfMgmt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perfMgmt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perfMgmt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perfMgmt.EntityData.Children = make(map[string]types.YChild)
    perfMgmt.EntityData.Children["resources"] = types.YChild{"Resources", &perfMgmt.Resources}
    perfMgmt.EntityData.Children["statistics"] = types.YChild{"Statistics", &perfMgmt.Statistics}
    perfMgmt.EntityData.Children["enable"] = types.YChild{"Enable", &perfMgmt.Enable}
    perfMgmt.EntityData.Children["reg-exp-groups"] = types.YChild{"RegExpGroups", &perfMgmt.RegExpGroups}
    perfMgmt.EntityData.Children["threshold"] = types.YChild{"Threshold", &perfMgmt.Threshold}
    perfMgmt.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(perfMgmt.EntityData)
}

// PerfMgmt_Resources
// Resources configuration
type PerfMgmt_Resources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the TFTP server IP address and directory name.
    TftpResources PerfMgmt_Resources_TftpResources

    // Configure local dump parameters.
    DumpLocal PerfMgmt_Resources_DumpLocal

    // Configure the memory usage limits of performance management.
    MemoryResources PerfMgmt_Resources_MemoryResources
}

func (resources *PerfMgmt_Resources) GetEntityData() *types.CommonEntityData {
    resources.EntityData.YFilter = resources.YFilter
    resources.EntityData.YangName = "resources"
    resources.EntityData.BundleName = "cisco_ios_xr"
    resources.EntityData.ParentYangName = "perf-mgmt"
    resources.EntityData.SegmentPath = "resources"
    resources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resources.EntityData.Children = make(map[string]types.YChild)
    resources.EntityData.Children["tftp-resources"] = types.YChild{"TftpResources", &resources.TftpResources}
    resources.EntityData.Children["dump-local"] = types.YChild{"DumpLocal", &resources.DumpLocal}
    resources.EntityData.Children["memory-resources"] = types.YChild{"MemoryResources", &resources.MemoryResources}
    resources.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(resources.EntityData)
}

// PerfMgmt_Resources_TftpResources
// Configure the TFTP server IP address and
// directory name
// This type is a presence type.
type PerfMgmt_Resources_TftpResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the TFTP server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    ServerAddress interface{}

    // Directory name on TFTP server. The type is string. This attribute is
    // mandatory.
    Directory interface{}

    // VRF name. The type is string with length: 1..32.
    VrfName interface{}
}

func (tftpResources *PerfMgmt_Resources_TftpResources) GetEntityData() *types.CommonEntityData {
    tftpResources.EntityData.YFilter = tftpResources.YFilter
    tftpResources.EntityData.YangName = "tftp-resources"
    tftpResources.EntityData.BundleName = "cisco_ios_xr"
    tftpResources.EntityData.ParentYangName = "resources"
    tftpResources.EntityData.SegmentPath = "tftp-resources"
    tftpResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpResources.EntityData.Children = make(map[string]types.YChild)
    tftpResources.EntityData.Leafs = make(map[string]types.YLeaf)
    tftpResources.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", tftpResources.ServerAddress}
    tftpResources.EntityData.Leafs["directory"] = types.YLeaf{"Directory", tftpResources.Directory}
    tftpResources.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", tftpResources.VrfName}
    return &(tftpResources.EntityData)
}

// PerfMgmt_Resources_DumpLocal
// Configure local dump parameters
type PerfMgmt_Resources_DumpLocal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable data dump onto local filesystem. The type is interface{}.
    Enable interface{}
}

func (dumpLocal *PerfMgmt_Resources_DumpLocal) GetEntityData() *types.CommonEntityData {
    dumpLocal.EntityData.YFilter = dumpLocal.YFilter
    dumpLocal.EntityData.YangName = "dump-local"
    dumpLocal.EntityData.BundleName = "cisco_ios_xr"
    dumpLocal.EntityData.ParentYangName = "resources"
    dumpLocal.EntityData.SegmentPath = "dump-local"
    dumpLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dumpLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dumpLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dumpLocal.EntityData.Children = make(map[string]types.YChild)
    dumpLocal.EntityData.Leafs = make(map[string]types.YLeaf)
    dumpLocal.EntityData.Leafs["enable"] = types.YLeaf{"Enable", dumpLocal.Enable}
    return &(dumpLocal.EntityData)
}

// PerfMgmt_Resources_MemoryResources
// Configure the memory usage limits of
// performance management
type PerfMgmt_Resources_MemoryResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum limit for memory usage (Kbytes) for data buffers. The type is
    // interface{} with range: -2147483648..2147483647. Units are kilobyte.
    MaxLimit interface{}

    // Specify a minimum free memory (Kbytes) to be ensured before allowing a
    // collection request. The type is interface{} with range:
    // -2147483648..2147483647. Units are kilobyte.
    MinReserved interface{}
}

func (memoryResources *PerfMgmt_Resources_MemoryResources) GetEntityData() *types.CommonEntityData {
    memoryResources.EntityData.YFilter = memoryResources.YFilter
    memoryResources.EntityData.YangName = "memory-resources"
    memoryResources.EntityData.BundleName = "cisco_ios_xr"
    memoryResources.EntityData.ParentYangName = "resources"
    memoryResources.EntityData.SegmentPath = "memory-resources"
    memoryResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryResources.EntityData.Children = make(map[string]types.YChild)
    memoryResources.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryResources.EntityData.Leafs["max-limit"] = types.YLeaf{"MaxLimit", memoryResources.MaxLimit}
    memoryResources.EntityData.Leafs["min-reserved"] = types.YLeaf{"MinReserved", memoryResources.MinReserved}
    return &(memoryResources.EntityData)
}

// PerfMgmt_Statistics
// Templates for collection of statistics
type PerfMgmt_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *PerfMgmt_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "perf-mgmt"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["generic-counter-interface"] = types.YChild{"GenericCounterInterface", &statistics.GenericCounterInterface}
    statistics.EntityData.Children["process-node"] = types.YChild{"ProcessNode", &statistics.ProcessNode}
    statistics.EntityData.Children["basic-counter-interface"] = types.YChild{"BasicCounterInterface", &statistics.BasicCounterInterface}
    statistics.EntityData.Children["ospfv3-protocol"] = types.YChild{"Ospfv3Protocol", &statistics.Ospfv3Protocol}
    statistics.EntityData.Children["cpu-node"] = types.YChild{"CpuNode", &statistics.CpuNode}
    statistics.EntityData.Children["data-rate-interface"] = types.YChild{"DataRateInterface", &statistics.DataRateInterface}
    statistics.EntityData.Children["memory-node"] = types.YChild{"MemoryNode", &statistics.MemoryNode}
    statistics.EntityData.Children["ldp-mpls"] = types.YChild{"LdpMpls", &statistics.LdpMpls}
    statistics.EntityData.Children["bgp"] = types.YChild{"Bgp", &statistics.Bgp}
    statistics.EntityData.Children["ospfv2-protocol"] = types.YChild{"Ospfv2Protocol", &statistics.Ospfv2Protocol}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// PerfMgmt_Statistics_GenericCounterInterface
// Interface Generic GenericCounter collection
// templates
type PerfMgmt_Statistics_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_GenericCounterInterface_Templates
}

func (genericCounterInterface *PerfMgmt_Statistics_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "statistics"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface"
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterface.EntityData.Children["templates"] = types.YChild{"Templates", &genericCounterInterface.Templates}
    genericCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Statistics_GenericCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_GenericCounterInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_GenericCounterInterface_Templates_Template.
    Template []PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_GenericCounterInterface_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "generic-counter-interface"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_GenericCounterInterface_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_ProcessNode
// Node Process collection templates
type PerfMgmt_Statistics_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_ProcessNode_Templates
}

func (processNode *PerfMgmt_Statistics_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "statistics"
    processNode.EntityData.SegmentPath = "process-node"
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = make(map[string]types.YChild)
    processNode.EntityData.Children["templates"] = types.YChild{"Templates", &processNode.Templates}
    processNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNode.EntityData)
}

// PerfMgmt_Statistics_ProcessNode_Templates
// Template name
type PerfMgmt_Statistics_ProcessNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_ProcessNode_Templates_Template.
    Template []PerfMgmt_Statistics_ProcessNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_ProcessNode_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "process-node"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_ProcessNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_ProcessNode_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_BasicCounterInterface
// Interface BasicCounter collection templates
type PerfMgmt_Statistics_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_BasicCounterInterface_Templates
}

func (basicCounterInterface *PerfMgmt_Statistics_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "statistics"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface"
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterface.EntityData.Children["templates"] = types.YChild{"Templates", &basicCounterInterface.Templates}
    basicCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Statistics_BasicCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_BasicCounterInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_BasicCounterInterface_Templates_Template.
    Template []PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_BasicCounterInterface_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "basic-counter-interface"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_BasicCounterInterface_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_Ospfv3Protocol
// OSPF v3 Protocol collection templates
type PerfMgmt_Statistics_Ospfv3Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Ospfv3Protocol_Templates
}

func (ospfv3Protocol *PerfMgmt_Statistics_Ospfv3Protocol) GetEntityData() *types.CommonEntityData {
    ospfv3Protocol.EntityData.YFilter = ospfv3Protocol.YFilter
    ospfv3Protocol.EntityData.YangName = "ospfv3-protocol"
    ospfv3Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv3Protocol.EntityData.ParentYangName = "statistics"
    ospfv3Protocol.EntityData.SegmentPath = "ospfv3-protocol"
    ospfv3Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv3Protocol.EntityData.Children["templates"] = types.YChild{"Templates", &ospfv3Protocol.Templates}
    ospfv3Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Statistics_Ospfv3Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv3Protocol_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template.
    Template []PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
}

func (templates *PerfMgmt_Statistics_Ospfv3Protocol_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "ospfv3-protocol"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_CpuNode
// Node CPU collection templates
type PerfMgmt_Statistics_CpuNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_CpuNode_Templates
}

func (cpuNode *PerfMgmt_Statistics_CpuNode) GetEntityData() *types.CommonEntityData {
    cpuNode.EntityData.YFilter = cpuNode.YFilter
    cpuNode.EntityData.YangName = "cpu-node"
    cpuNode.EntityData.BundleName = "cisco_ios_xr"
    cpuNode.EntityData.ParentYangName = "statistics"
    cpuNode.EntityData.SegmentPath = "cpu-node"
    cpuNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNode.EntityData.Children = make(map[string]types.YChild)
    cpuNode.EntityData.Children["templates"] = types.YChild{"Templates", &cpuNode.Templates}
    cpuNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuNode.EntityData)
}

// PerfMgmt_Statistics_CpuNode_Templates
// Template name
type PerfMgmt_Statistics_CpuNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_CpuNode_Templates_Template.
    Template []PerfMgmt_Statistics_CpuNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_CpuNode_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "cpu-node"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_CpuNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_CpuNode_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_DataRateInterface
// Interface DataRate collection templates
type PerfMgmt_Statistics_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_DataRateInterface_Templates
}

func (dataRateInterface *PerfMgmt_Statistics_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "statistics"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface"
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = make(map[string]types.YChild)
    dataRateInterface.EntityData.Children["templates"] = types.YChild{"Templates", &dataRateInterface.Templates}
    dataRateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Statistics_DataRateInterface_Templates
// Template name
type PerfMgmt_Statistics_DataRateInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_DataRateInterface_Templates_Template.
    Template []PerfMgmt_Statistics_DataRateInterface_Templates_Template
}

func (templates *PerfMgmt_Statistics_DataRateInterface_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "data-rate-interface"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_DataRateInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_DataRateInterface_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_MemoryNode
// Node Memory collection templates
type PerfMgmt_Statistics_MemoryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_MemoryNode_Templates
}

func (memoryNode *PerfMgmt_Statistics_MemoryNode) GetEntityData() *types.CommonEntityData {
    memoryNode.EntityData.YFilter = memoryNode.YFilter
    memoryNode.EntityData.YangName = "memory-node"
    memoryNode.EntityData.BundleName = "cisco_ios_xr"
    memoryNode.EntityData.ParentYangName = "statistics"
    memoryNode.EntityData.SegmentPath = "memory-node"
    memoryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNode.EntityData.Children = make(map[string]types.YChild)
    memoryNode.EntityData.Children["templates"] = types.YChild{"Templates", &memoryNode.Templates}
    memoryNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryNode.EntityData)
}

// PerfMgmt_Statistics_MemoryNode_Templates
// Template name
type PerfMgmt_Statistics_MemoryNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_MemoryNode_Templates_Template.
    Template []PerfMgmt_Statistics_MemoryNode_Templates_Template
}

func (templates *PerfMgmt_Statistics_MemoryNode_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "memory-node"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_MemoryNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_MemoryNode_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_LdpMpls
// MPLS LDP collection templates
type PerfMgmt_Statistics_LdpMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_LdpMpls_Templates
}

func (ldpMpls *PerfMgmt_Statistics_LdpMpls) GetEntityData() *types.CommonEntityData {
    ldpMpls.EntityData.YFilter = ldpMpls.YFilter
    ldpMpls.EntityData.YangName = "ldp-mpls"
    ldpMpls.EntityData.BundleName = "cisco_ios_xr"
    ldpMpls.EntityData.ParentYangName = "statistics"
    ldpMpls.EntityData.SegmentPath = "ldp-mpls"
    ldpMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMpls.EntityData.Children = make(map[string]types.YChild)
    ldpMpls.EntityData.Children["templates"] = types.YChild{"Templates", &ldpMpls.Templates}
    ldpMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldpMpls.EntityData)
}

// PerfMgmt_Statistics_LdpMpls_Templates
// Template name
type PerfMgmt_Statistics_LdpMpls_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_LdpMpls_Templates_Template.
    Template []PerfMgmt_Statistics_LdpMpls_Templates_Template
}

func (templates *PerfMgmt_Statistics_LdpMpls_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "ldp-mpls"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_LdpMpls_Templates_Template
// A template instance
type PerfMgmt_Statistics_LdpMpls_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_Bgp
// BGP collection templates
type PerfMgmt_Statistics_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Bgp_Templates
}

func (bgp *PerfMgmt_Statistics_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "statistics"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Children["templates"] = types.YChild{"Templates", &bgp.Templates}
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgp.EntityData)
}

// PerfMgmt_Statistics_Bgp_Templates
// Template name
type PerfMgmt_Statistics_Bgp_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Bgp_Templates_Template.
    Template []PerfMgmt_Statistics_Bgp_Templates_Template
}

func (templates *PerfMgmt_Statistics_Bgp_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "bgp"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Bgp_Templates_Template
// A template instance
type PerfMgmt_Statistics_Bgp_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Statistics_Ospfv2Protocol
// OSPF v2 Protocol collection templates
type PerfMgmt_Statistics_Ospfv2Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name.
    Templates PerfMgmt_Statistics_Ospfv2Protocol_Templates
}

func (ospfv2Protocol *PerfMgmt_Statistics_Ospfv2Protocol) GetEntityData() *types.CommonEntityData {
    ospfv2Protocol.EntityData.YFilter = ospfv2Protocol.YFilter
    ospfv2Protocol.EntityData.YangName = "ospfv2-protocol"
    ospfv2Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv2Protocol.EntityData.ParentYangName = "statistics"
    ospfv2Protocol.EntityData.SegmentPath = "ospfv2-protocol"
    ospfv2Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv2Protocol.EntityData.Children["templates"] = types.YChild{"Templates", &ospfv2Protocol.Templates}
    ospfv2Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Statistics_Ospfv2Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv2Protocol_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template.
    Template []PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
}

func (templates *PerfMgmt_Statistics_Ospfv2Protocol_Templates) GetEntityData() *types.CommonEntityData {
    templates.EntityData.YFilter = templates.YFilter
    templates.EntityData.YangName = "templates"
    templates.EntityData.BundleName = "cisco_ios_xr"
    templates.EntityData.ParentYangName = "ospfv2-protocol"
    templates.EntityData.SegmentPath = "templates"
    templates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    templates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    templates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    templates.EntityData.Children = make(map[string]types.YChild)
    templates.EntityData.Children["template"] = types.YChild{"Template", nil}
    for i := range templates.Template {
        templates.EntityData.Children[types.GetSegmentPath(&templates.Template[i])] = types.YChild{"Template", &templates.Template[i]}
    }
    templates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + "[template-name='" + fmt.Sprintf("%v", template.TemplateName) + "']"
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = make(map[string]types.YChild)
    template.EntityData.Leafs = make(map[string]types.YLeaf)
    template.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", template.TemplateName}
    template.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", template.RegExpGroup}
    template.EntityData.Leafs["history-persistent"] = types.YLeaf{"HistoryPersistent", template.HistoryPersistent}
    template.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", template.VrfGroup}
    template.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", template.SampleInterval}
    template.EntityData.Leafs["sample-size"] = types.YLeaf{"SampleSize", template.SampleSize}
    return &(template.EntityData)
}

// PerfMgmt_Enable
// Start data collection and/or threshold
// monitoring
type PerfMgmt_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start threshold monitoring using a defined template.
    Threshold PerfMgmt_Enable_Threshold

    // Start periodic collection using a defined a template.
    Statistics PerfMgmt_Enable_Statistics

    // Start data collection for a monitored instance.
    MonitorEnable PerfMgmt_Enable_MonitorEnable
}

func (enable *PerfMgmt_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "perf-mgmt"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Children["threshold"] = types.YChild{"Threshold", &enable.Threshold}
    enable.EntityData.Children["statistics"] = types.YChild{"Statistics", &enable.Statistics}
    enable.EntityData.Children["monitor-enable"] = types.YChild{"MonitorEnable", &enable.MonitorEnable}
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enable.EntityData)
}

// PerfMgmt_Enable_Threshold
// Start threshold monitoring using a defined
// template
type PerfMgmt_Enable_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *PerfMgmt_Enable_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "enable"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["ospfv3-protocol"] = types.YChild{"Ospfv3Protocol", &threshold.Ospfv3Protocol}
    threshold.EntityData.Children["bgp"] = types.YChild{"Bgp", &threshold.Bgp}
    threshold.EntityData.Children["data-rate-interface"] = types.YChild{"DataRateInterface", &threshold.DataRateInterface}
    threshold.EntityData.Children["ospfv2-protocol"] = types.YChild{"Ospfv2Protocol", &threshold.Ospfv2Protocol}
    threshold.EntityData.Children["memory-node"] = types.YChild{"MemoryNode", &threshold.MemoryNode}
    threshold.EntityData.Children["generic-counter-interface"] = types.YChild{"GenericCounterInterface", &threshold.GenericCounterInterface}
    threshold.EntityData.Children["cpu-node"] = types.YChild{"CpuNode", &threshold.CpuNode}
    threshold.EntityData.Children["ldp-mpls"] = types.YChild{"LdpMpls", &threshold.LdpMpls}
    threshold.EntityData.Children["process-node"] = types.YChild{"ProcessNode", &threshold.ProcessNode}
    threshold.EntityData.Children["basic-counter-interface"] = types.YChild{"BasicCounterInterface", &threshold.BasicCounterInterface}
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(threshold.EntityData)
}

// PerfMgmt_Enable_Threshold_Ospfv3Protocol
// Threshold monitoring for OSPF v3 Protocol
type PerfMgmt_Enable_Threshold_Ospfv3Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv3Protocol *PerfMgmt_Enable_Threshold_Ospfv3Protocol) GetEntityData() *types.CommonEntityData {
    ospfv3Protocol.EntityData.YFilter = ospfv3Protocol.YFilter
    ospfv3Protocol.EntityData.YangName = "ospfv3-protocol"
    ospfv3Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv3Protocol.EntityData.ParentYangName = "threshold"
    ospfv3Protocol.EntityData.SegmentPath = "ospfv3-protocol"
    ospfv3Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv3Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Protocol.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv3Protocol.TemplateName}
    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Enable_Threshold_Bgp
// Threshold monitoring for BGP
type PerfMgmt_Enable_Threshold_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (bgp *PerfMgmt_Enable_Threshold_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "threshold"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", bgp.TemplateName}
    return &(bgp.EntityData)
}

// PerfMgmt_Enable_Threshold_DataRateInterface
// Threshold monitoring for Interface  data-rates
type PerfMgmt_Enable_Threshold_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (dataRateInterface *PerfMgmt_Enable_Threshold_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "threshold"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface"
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = make(map[string]types.YChild)
    dataRateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    dataRateInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", dataRateInterface.TemplateName}
    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Enable_Threshold_Ospfv2Protocol
// Threshold monitoring for OSPF v2 Protocol
type PerfMgmt_Enable_Threshold_Ospfv2Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv2Protocol *PerfMgmt_Enable_Threshold_Ospfv2Protocol) GetEntityData() *types.CommonEntityData {
    ospfv2Protocol.EntityData.YFilter = ospfv2Protocol.YFilter
    ospfv2Protocol.EntityData.YangName = "ospfv2-protocol"
    ospfv2Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv2Protocol.EntityData.ParentYangName = "threshold"
    ospfv2Protocol.EntityData.SegmentPath = "ospfv2-protocol"
    ospfv2Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv2Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Protocol.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv2Protocol.TemplateName}
    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode
// Threshold monitoring for memory
type PerfMgmt_Enable_Threshold_MemoryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_MemoryNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_MemoryNode_NodeAll
}

func (memoryNode *PerfMgmt_Enable_Threshold_MemoryNode) GetEntityData() *types.CommonEntityData {
    memoryNode.EntityData.YFilter = memoryNode.YFilter
    memoryNode.EntityData.YangName = "memory-node"
    memoryNode.EntityData.BundleName = "cisco_ios_xr"
    memoryNode.EntityData.ParentYangName = "threshold"
    memoryNode.EntityData.SegmentPath = "memory-node"
    memoryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNode.EntityData.Children = make(map[string]types.YChild)
    memoryNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &memoryNode.Nodes}
    memoryNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &memoryNode.NodeAll}
    memoryNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryNode.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_MemoryNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "memory-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_MemoryNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_MemoryNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "memory-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Threshold_GenericCounterInterface
// Threshold monitoring for Interface
// generic-counters
type PerfMgmt_Enable_Threshold_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (genericCounterInterface *PerfMgmt_Enable_Threshold_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "threshold"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface"
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    genericCounterInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", genericCounterInterface.TemplateName}
    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode
// Threshold monitoring for CPU
type PerfMgmt_Enable_Threshold_CpuNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_CpuNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_CpuNode_NodeAll
}

func (cpuNode *PerfMgmt_Enable_Threshold_CpuNode) GetEntityData() *types.CommonEntityData {
    cpuNode.EntityData.YFilter = cpuNode.YFilter
    cpuNode.EntityData.YangName = "cpu-node"
    cpuNode.EntityData.BundleName = "cisco_ios_xr"
    cpuNode.EntityData.ParentYangName = "threshold"
    cpuNode.EntityData.SegmentPath = "cpu-node"
    cpuNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNode.EntityData.Children = make(map[string]types.YChild)
    cpuNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &cpuNode.Nodes}
    cpuNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &cpuNode.NodeAll}
    cpuNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuNode.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_CpuNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_CpuNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cpu-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_CpuNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_CpuNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "cpu-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Threshold_LdpMpls
// Threshold monitoring for LDP
type PerfMgmt_Enable_Threshold_LdpMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ldpMpls *PerfMgmt_Enable_Threshold_LdpMpls) GetEntityData() *types.CommonEntityData {
    ldpMpls.EntityData.YFilter = ldpMpls.YFilter
    ldpMpls.EntityData.YangName = "ldp-mpls"
    ldpMpls.EntityData.BundleName = "cisco_ios_xr"
    ldpMpls.EntityData.ParentYangName = "threshold"
    ldpMpls.EntityData.SegmentPath = "ldp-mpls"
    ldpMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMpls.EntityData.Children = make(map[string]types.YChild)
    ldpMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    ldpMpls.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ldpMpls.TemplateName}
    return &(ldpMpls.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode
// Threshold monitoring for process
type PerfMgmt_Enable_Threshold_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_Threshold_ProcessNode_Nodes

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Threshold_ProcessNode_NodeAll
}

func (processNode *PerfMgmt_Enable_Threshold_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "threshold"
    processNode.EntityData.SegmentPath = "process-node"
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = make(map[string]types.YChild)
    processNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &processNode.Nodes}
    processNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &processNode.NodeAll}
    processNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNode.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node.
    Node []PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Threshold_ProcessNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "process-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Threshold_ProcessNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Threshold_ProcessNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "process-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Threshold_BasicCounterInterface
// Threshold monitoring for Interface
// basic-counters
type PerfMgmt_Enable_Threshold_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (basicCounterInterface *PerfMgmt_Enable_Threshold_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "threshold"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface"
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    basicCounterInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", basicCounterInterface.TemplateName}
    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Enable_Statistics
// Start periodic collection using a defined a
// template
type PerfMgmt_Enable_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *PerfMgmt_Enable_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "enable"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["generic-counter-interface"] = types.YChild{"GenericCounterInterface", &statistics.GenericCounterInterface}
    statistics.EntityData.Children["bgp"] = types.YChild{"Bgp", &statistics.Bgp}
    statistics.EntityData.Children["ospfv2-protocol"] = types.YChild{"Ospfv2Protocol", &statistics.Ospfv2Protocol}
    statistics.EntityData.Children["ospfv3-protocol"] = types.YChild{"Ospfv3Protocol", &statistics.Ospfv3Protocol}
    statistics.EntityData.Children["cpu-node"] = types.YChild{"CpuNode", &statistics.CpuNode}
    statistics.EntityData.Children["basic-counter-interface"] = types.YChild{"BasicCounterInterface", &statistics.BasicCounterInterface}
    statistics.EntityData.Children["process-node"] = types.YChild{"ProcessNode", &statistics.ProcessNode}
    statistics.EntityData.Children["data-rate-interface"] = types.YChild{"DataRateInterface", &statistics.DataRateInterface}
    statistics.EntityData.Children["memory-node"] = types.YChild{"MemoryNode", &statistics.MemoryNode}
    statistics.EntityData.Children["ldp-mpls"] = types.YChild{"LdpMpls", &statistics.LdpMpls}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// PerfMgmt_Enable_Statistics_GenericCounterInterface
// Statistics collection for generic-counters
type PerfMgmt_Enable_Statistics_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (genericCounterInterface *PerfMgmt_Enable_Statistics_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "statistics"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface"
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    genericCounterInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", genericCounterInterface.TemplateName}
    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Enable_Statistics_Bgp
// Data collection for BGP
type PerfMgmt_Enable_Statistics_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (bgp *PerfMgmt_Enable_Statistics_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "statistics"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", bgp.TemplateName}
    return &(bgp.EntityData)
}

// PerfMgmt_Enable_Statistics_Ospfv2Protocol
// Data collection for OSPF v2 Protocol
type PerfMgmt_Enable_Statistics_Ospfv2Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv2Protocol *PerfMgmt_Enable_Statistics_Ospfv2Protocol) GetEntityData() *types.CommonEntityData {
    ospfv2Protocol.EntityData.YFilter = ospfv2Protocol.YFilter
    ospfv2Protocol.EntityData.YangName = "ospfv2-protocol"
    ospfv2Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv2Protocol.EntityData.ParentYangName = "statistics"
    ospfv2Protocol.EntityData.SegmentPath = "ospfv2-protocol"
    ospfv2Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv2Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2Protocol.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv2Protocol.TemplateName}
    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Enable_Statistics_Ospfv3Protocol
// Data collection for OSPF v3 Protocol
type PerfMgmt_Enable_Statistics_Ospfv3Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfv3Protocol *PerfMgmt_Enable_Statistics_Ospfv3Protocol) GetEntityData() *types.CommonEntityData {
    ospfv3Protocol.EntityData.YFilter = ospfv3Protocol.YFilter
    ospfv3Protocol.EntityData.YangName = "ospfv3-protocol"
    ospfv3Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv3Protocol.EntityData.ParentYangName = "statistics"
    ospfv3Protocol.EntityData.SegmentPath = "ospfv3-protocol"
    ospfv3Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv3Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3Protocol.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv3Protocol.TemplateName}
    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode
// Collection for CPU
type PerfMgmt_Enable_Statistics_CpuNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_CpuNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_CpuNode_Nodes
}

func (cpuNode *PerfMgmt_Enable_Statistics_CpuNode) GetEntityData() *types.CommonEntityData {
    cpuNode.EntityData.YFilter = cpuNode.YFilter
    cpuNode.EntityData.YangName = "cpu-node"
    cpuNode.EntityData.BundleName = "cisco_ios_xr"
    cpuNode.EntityData.ParentYangName = "statistics"
    cpuNode.EntityData.SegmentPath = "cpu-node"
    cpuNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNode.EntityData.Children = make(map[string]types.YChild)
    cpuNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &cpuNode.NodeAll}
    cpuNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &cpuNode.Nodes}
    cpuNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuNode.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_CpuNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_CpuNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "cpu-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_CpuNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_CpuNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cpu-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Statistics_BasicCounterInterface
// Statistics collection for basic-counters
type PerfMgmt_Enable_Statistics_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (basicCounterInterface *PerfMgmt_Enable_Statistics_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "statistics"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface"
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    basicCounterInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", basicCounterInterface.TemplateName}
    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode
// Collection for process
type PerfMgmt_Enable_Statistics_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_ProcessNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_ProcessNode_Nodes
}

func (processNode *PerfMgmt_Enable_Statistics_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "statistics"
    processNode.EntityData.SegmentPath = "process-node"
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = make(map[string]types.YChild)
    processNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &processNode.NodeAll}
    processNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &processNode.Nodes}
    processNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNode.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_ProcessNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_ProcessNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "process-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_ProcessNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "process-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Statistics_DataRateInterface
// Statistics collection for generic-counters
type PerfMgmt_Enable_Statistics_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (dataRateInterface *PerfMgmt_Enable_Statistics_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "statistics"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface"
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = make(map[string]types.YChild)
    dataRateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    dataRateInterface.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", dataRateInterface.TemplateName}
    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode
// Collection for memory
type PerfMgmt_Enable_Statistics_MemoryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All the the nodes.
    NodeAll PerfMgmt_Enable_Statistics_MemoryNode_NodeAll

    // Node specification.
    Nodes PerfMgmt_Enable_Statistics_MemoryNode_Nodes
}

func (memoryNode *PerfMgmt_Enable_Statistics_MemoryNode) GetEntityData() *types.CommonEntityData {
    memoryNode.EntityData.YFilter = memoryNode.YFilter
    memoryNode.EntityData.YangName = "memory-node"
    memoryNode.EntityData.BundleName = "cisco_ios_xr"
    memoryNode.EntityData.ParentYangName = "statistics"
    memoryNode.EntityData.SegmentPath = "memory-node"
    memoryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNode.EntityData.Children = make(map[string]types.YChild)
    memoryNode.EntityData.Children["node-all"] = types.YChild{"NodeAll", &memoryNode.NodeAll}
    memoryNode.EntityData.Children["nodes"] = types.YChild{"Nodes", &memoryNode.Nodes}
    memoryNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryNode.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode_NodeAll
// All the the nodes
type PerfMgmt_Enable_Statistics_MemoryNode_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (nodeAll *PerfMgmt_Enable_Statistics_MemoryNode_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "memory-node"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = make(map[string]types.YChild)
    nodeAll.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeAll.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", nodeAll.TemplateName}
    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node.
    Node []PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
}

func (nodes *PerfMgmt_Enable_Statistics_MemoryNode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "memory-node"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_Statistics_LdpMpls
// Collection for labels distribution protocol
type PerfMgmt_Enable_Statistics_LdpMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}
}

func (ldpMpls *PerfMgmt_Enable_Statistics_LdpMpls) GetEntityData() *types.CommonEntityData {
    ldpMpls.EntityData.YFilter = ldpMpls.YFilter
    ldpMpls.EntityData.YangName = "ldp-mpls"
    ldpMpls.EntityData.BundleName = "cisco_ios_xr"
    ldpMpls.EntityData.ParentYangName = "statistics"
    ldpMpls.EntityData.SegmentPath = "ldp-mpls"
    ldpMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMpls.EntityData.Children = make(map[string]types.YChild)
    ldpMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    ldpMpls.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ldpMpls.TemplateName}
    return &(ldpMpls.EntityData)
}

// PerfMgmt_Enable_MonitorEnable
// Start data collection for a monitored instance
type PerfMgmt_Enable_MonitorEnable struct {
    EntityData types.CommonEntityData
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

func (monitorEnable *PerfMgmt_Enable_MonitorEnable) GetEntityData() *types.CommonEntityData {
    monitorEnable.EntityData.YFilter = monitorEnable.YFilter
    monitorEnable.EntityData.YangName = "monitor-enable"
    monitorEnable.EntityData.BundleName = "cisco_ios_xr"
    monitorEnable.EntityData.ParentYangName = "enable"
    monitorEnable.EntityData.SegmentPath = "monitor-enable"
    monitorEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorEnable.EntityData.Children = make(map[string]types.YChild)
    monitorEnable.EntityData.Children["ldp-mpls"] = types.YChild{"LdpMpls", &monitorEnable.LdpMpls}
    monitorEnable.EntityData.Children["ospfv3-protocol"] = types.YChild{"Ospfv3Protocol", &monitorEnable.Ospfv3Protocol}
    monitorEnable.EntityData.Children["generic-counters"] = types.YChild{"GenericCounters", &monitorEnable.GenericCounters}
    monitorEnable.EntityData.Children["process"] = types.YChild{"Process", &monitorEnable.Process}
    monitorEnable.EntityData.Children["basic-counters"] = types.YChild{"BasicCounters", &monitorEnable.BasicCounters}
    monitorEnable.EntityData.Children["memory"] = types.YChild{"Memory", &monitorEnable.Memory}
    monitorEnable.EntityData.Children["ospfv2-protocol"] = types.YChild{"Ospfv2Protocol", &monitorEnable.Ospfv2Protocol}
    monitorEnable.EntityData.Children["cpu"] = types.YChild{"Cpu", &monitorEnable.Cpu}
    monitorEnable.EntityData.Children["bgp"] = types.YChild{"Bgp", &monitorEnable.Bgp}
    monitorEnable.EntityData.Children["data-rates"] = types.YChild{"DataRates", &monitorEnable.DataRates}
    monitorEnable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(monitorEnable.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_LdpMpls
// Monitoring for LDP
type PerfMgmt_Enable_MonitorEnable_LdpMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP session specification.
    Sessions PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions
}

func (ldpMpls *PerfMgmt_Enable_MonitorEnable_LdpMpls) GetEntityData() *types.CommonEntityData {
    ldpMpls.EntityData.YFilter = ldpMpls.YFilter
    ldpMpls.EntityData.YangName = "ldp-mpls"
    ldpMpls.EntityData.BundleName = "cisco_ios_xr"
    ldpMpls.EntityData.ParentYangName = "monitor-enable"
    ldpMpls.EntityData.SegmentPath = "ldp-mpls"
    ldpMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMpls.EntityData.Children = make(map[string]types.YChild)
    ldpMpls.EntityData.Children["sessions"] = types.YChild{"Sessions", &ldpMpls.Sessions}
    ldpMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldpMpls.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions
// LDP session specification
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the LDP Session. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session.
    Session []PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
}

func (sessions *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "ldp-mpls"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
// IP address of the LDP Session
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the LDP Session. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Session interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[session='" + fmt.Sprintf("%v", session.Session) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["session"] = types.YLeaf{"Session", session.Session}
    session.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", session.TemplateName}
    return &(session.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol
// Monitor OSPF v3 Protocol
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor an instance.
    OspfInstances PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances
}

func (ospfv3Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol) GetEntityData() *types.CommonEntityData {
    ospfv3Protocol.EntityData.YFilter = ospfv3Protocol.YFilter
    ospfv3Protocol.EntityData.YangName = "ospfv3-protocol"
    ospfv3Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv3Protocol.EntityData.ParentYangName = "monitor-enable"
    ospfv3Protocol.EntityData.SegmentPath = "ospfv3-protocol"
    ospfv3Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv3Protocol.EntityData.Children["ospf-instances"] = types.YChild{"OspfInstances", &ospfv3Protocol.OspfInstances}
    ospfv3Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance.
    OspfInstance []PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances) GetEntityData() *types.CommonEntityData {
    ospfInstances.EntityData.YFilter = ospfInstances.YFilter
    ospfInstances.EntityData.YangName = "ospf-instances"
    ospfInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfInstances.EntityData.ParentYangName = "ospfv3-protocol"
    ospfInstances.EntityData.SegmentPath = "ospf-instances"
    ospfInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstances.EntityData.Children = make(map[string]types.YChild)
    ospfInstances.EntityData.Children["ospf-instance"] = types.YChild{"OspfInstance", nil}
    for i := range ospfInstances.OspfInstance {
        ospfInstances.EntityData.Children[types.GetSegmentPath(&ospfInstances.OspfInstance[i])] = types.YChild{"OspfInstance", &ospfInstances.OspfInstance[i]}
    }
    ospfInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfInstances.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfInstance.EntityData.ParentYangName = "ospf-instances"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfInstance.InstanceName) + "']"
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstance.EntityData.Children = make(map[string]types.YChild)
    ospfInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfInstance.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", ospfInstance.InstanceName}
    ospfInstance.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfInstance.TemplateName}
    return &(ospfInstance.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_GenericCounters
// Monitoring for generic-counters
type PerfMgmt_Enable_MonitorEnable_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces
}

func (genericCounters *PerfMgmt_Enable_MonitorEnable_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "monitor-enable"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = make(map[string]types.YChild)
    genericCounters.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &genericCounters.Interfaces}
    genericCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericCounters.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface_.
    Interface_ []PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "generic-counters"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", self.TemplateName}
    return &(self.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process
// Collection for a single process
type PerfMgmt_Enable_MonitorEnable_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    ProcessNodes PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes
}

func (process *PerfMgmt_Enable_MonitorEnable_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "monitor-enable"
    process.EntityData.SegmentPath = "process"
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = make(map[string]types.YChild)
    process.EntityData.Children["process-nodes"] = types.YChild{"ProcessNodes", &process.ProcessNodes}
    process.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(process.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode.
    ProcessNode []PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
}

func (processNodes *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes) GetEntityData() *types.CommonEntityData {
    processNodes.EntityData.YFilter = processNodes.YFilter
    processNodes.EntityData.YangName = "process-nodes"
    processNodes.EntityData.BundleName = "cisco_ios_xr"
    processNodes.EntityData.ParentYangName = "process"
    processNodes.EntityData.SegmentPath = "process-nodes"
    processNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNodes.EntityData.Children = make(map[string]types.YChild)
    processNodes.EntityData.Children["process-node"] = types.YChild{"ProcessNode", nil}
    for i := range processNodes.ProcessNode {
        processNodes.EntityData.Children[types.GetSegmentPath(&processNodes.ProcessNode[i])] = types.YChild{"ProcessNode", &processNodes.ProcessNode[i]}
    }
    processNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
// Node instance
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Process ID specification.
    Pids PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "process-nodes"
    processNode.EntityData.SegmentPath = "process-node" + "[node-id='" + fmt.Sprintf("%v", processNode.NodeId) + "']"
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = make(map[string]types.YChild)
    processNode.EntityData.Children["pids"] = types.YChild{"Pids", &processNode.Pids}
    processNode.EntityData.Leafs = make(map[string]types.YLeaf)
    processNode.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", processNode.NodeId}
    return &(processNode.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
// Process ID specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify an existing template for data collection. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid.
    Pid []PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid
}

func (pids *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids) GetEntityData() *types.CommonEntityData {
    pids.EntityData.YFilter = pids.YFilter
    pids.EntityData.YangName = "pids"
    pids.EntityData.BundleName = "cisco_ios_xr"
    pids.EntityData.ParentYangName = "process-node"
    pids.EntityData.SegmentPath = "pids"
    pids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pids.EntityData.Children = make(map[string]types.YChild)
    pids.EntityData.Children["pid"] = types.YChild{"Pid", nil}
    for i := range pids.Pid {
        pids.EntityData.Children[types.GetSegmentPath(&pids.Pid[i])] = types.YChild{"Pid", &pids.Pid[i]}
    }
    pids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pids.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid
// Specify an existing template for data
// collection
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify Process ID. The type is interface{} with
    // range: 0..4294967295.
    Pid interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (pid *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid) GetEntityData() *types.CommonEntityData {
    pid.EntityData.YFilter = pid.YFilter
    pid.EntityData.YangName = "pid"
    pid.EntityData.BundleName = "cisco_ios_xr"
    pid.EntityData.ParentYangName = "pids"
    pid.EntityData.SegmentPath = "pid" + "[pid='" + fmt.Sprintf("%v", pid.Pid) + "']"
    pid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pid.EntityData.Children = make(map[string]types.YChild)
    pid.EntityData.Leafs = make(map[string]types.YLeaf)
    pid.EntityData.Leafs["pid"] = types.YLeaf{"Pid", pid.Pid}
    pid.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", pid.TemplateName}
    return &(pid.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_BasicCounters
// Monitoring for basic-counters
type PerfMgmt_Enable_MonitorEnable_BasicCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces
}

func (basicCounters *PerfMgmt_Enable_MonitorEnable_BasicCounters) GetEntityData() *types.CommonEntityData {
    basicCounters.EntityData.YFilter = basicCounters.YFilter
    basicCounters.EntityData.YangName = "basic-counters"
    basicCounters.EntityData.BundleName = "cisco_ios_xr"
    basicCounters.EntityData.ParentYangName = "monitor-enable"
    basicCounters.EntityData.SegmentPath = "basic-counters"
    basicCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounters.EntityData.Children = make(map[string]types.YChild)
    basicCounters.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &basicCounters.Interfaces}
    basicCounters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicCounters.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface_.
    Interface_ []PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "basic-counters"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", self.TemplateName}
    return &(self.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Memory
// Collection for memory
type PerfMgmt_Enable_MonitorEnable_Memory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_MonitorEnable_Memory_Nodes
}

func (memory *PerfMgmt_Enable_MonitorEnable_Memory) GetEntityData() *types.CommonEntityData {
    memory.EntityData.YFilter = memory.YFilter
    memory.EntityData.YangName = "memory"
    memory.EntityData.BundleName = "cisco_ios_xr"
    memory.EntityData.ParentYangName = "monitor-enable"
    memory.EntityData.SegmentPath = "memory"
    memory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memory.EntityData.Children = make(map[string]types.YChild)
    memory.EntityData.Children["nodes"] = types.YChild{"Nodes", &memory.Nodes}
    memory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memory.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node.
    Node []PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Memory_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "memory"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol
// Monitor OSPF v2 Protocol
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor an instance.
    OspfInstances PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances
}

func (ospfv2Protocol *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol) GetEntityData() *types.CommonEntityData {
    ospfv2Protocol.EntityData.YFilter = ospfv2Protocol.YFilter
    ospfv2Protocol.EntityData.YangName = "ospfv2-protocol"
    ospfv2Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv2Protocol.EntityData.ParentYangName = "monitor-enable"
    ospfv2Protocol.EntityData.SegmentPath = "ospfv2-protocol"
    ospfv2Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv2Protocol.EntityData.Children["ospf-instances"] = types.YChild{"OspfInstances", &ospfv2Protocol.OspfInstances}
    ospfv2Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance.
    OspfInstance []PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
}

func (ospfInstances *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances) GetEntityData() *types.CommonEntityData {
    ospfInstances.EntityData.YFilter = ospfInstances.YFilter
    ospfInstances.EntityData.YangName = "ospf-instances"
    ospfInstances.EntityData.BundleName = "cisco_ios_xr"
    ospfInstances.EntityData.ParentYangName = "ospfv2-protocol"
    ospfInstances.EntityData.SegmentPath = "ospf-instances"
    ospfInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstances.EntityData.Children = make(map[string]types.YChild)
    ospfInstances.EntityData.Children["ospf-instance"] = types.YChild{"OspfInstance", nil}
    for i := range ospfInstances.OspfInstance {
        ospfInstances.EntityData.Children[types.GetSegmentPath(&ospfInstances.OspfInstance[i])] = types.YChild{"OspfInstance", &ospfInstances.OspfInstance[i]}
    }
    ospfInstances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfInstances.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfInstance.EntityData.ParentYangName = "ospf-instances"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + "[instance-name='" + fmt.Sprintf("%v", ospfInstance.InstanceName) + "']"
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstance.EntityData.Children = make(map[string]types.YChild)
    ospfInstance.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfInstance.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", ospfInstance.InstanceName}
    ospfInstance.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfInstance.TemplateName}
    return &(ospfInstance.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Cpu
// Collection for CPU
type PerfMgmt_Enable_MonitorEnable_Cpu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node specification.
    Nodes PerfMgmt_Enable_MonitorEnable_Cpu_Nodes
}

func (cpu *PerfMgmt_Enable_MonitorEnable_Cpu) GetEntityData() *types.CommonEntityData {
    cpu.EntityData.YFilter = cpu.YFilter
    cpu.EntityData.YangName = "cpu"
    cpu.EntityData.BundleName = "cisco_ios_xr"
    cpu.EntityData.ParentYangName = "monitor-enable"
    cpu.EntityData.SegmentPath = "cpu"
    cpu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpu.EntityData.Children = make(map[string]types.YChild)
    cpu.EntityData.Children["nodes"] = types.YChild{"Nodes", &cpu.Nodes}
    cpu.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpu.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node.
    Node []PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
}

func (nodes *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cpu"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-id='" + fmt.Sprintf("%v", node.NodeId) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-id"] = types.YLeaf{"NodeId", node.NodeId}
    node.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", node.TemplateName}
    return &(node.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Bgp
// Monitor BGP protocol
type PerfMgmt_Enable_MonitorEnable_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor BGP protocol for a BGP peer.
    Neighbors PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors
}

func (bgp *PerfMgmt_Enable_MonitorEnable_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "monitor-enable"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &bgp.Neighbors}
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgp.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors
// Monitor BGP protocol for a BGP peer
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor.
    Neighbor []PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
}

func (neighbors *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "bgp"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
// Neighbor being monitored
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the Neighbor. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[peer-address='" + fmt.Sprintf("%v", neighbor.PeerAddress) + "']"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", neighbor.PeerAddress}
    neighbor.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", neighbor.TemplateName}
    return &(neighbor.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_DataRates
// Monitoring for data-rates
type PerfMgmt_Enable_MonitorEnable_DataRates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor an Interface.
    Interfaces PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces
}

func (dataRates *PerfMgmt_Enable_MonitorEnable_DataRates) GetEntityData() *types.CommonEntityData {
    dataRates.EntityData.YFilter = dataRates.YFilter
    dataRates.EntityData.YangName = "data-rates"
    dataRates.EntityData.BundleName = "cisco_ios_xr"
    dataRates.EntityData.ParentYangName = "monitor-enable"
    dataRates.EntityData.SegmentPath = "data-rates"
    dataRates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRates.EntityData.Children = make(map[string]types.YChild)
    dataRates.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &dataRates.Interfaces}
    dataRates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataRates.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface_.
    Interface_ []PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
}

func (interfaces *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "data-rates"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", self.TemplateName}
    return &(self.EntityData)
}

// PerfMgmt_RegExpGroups
// Configure regular expression group
type PerfMgmt_RegExpGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify regular expression group name. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup.
    RegExpGroup []PerfMgmt_RegExpGroups_RegExpGroup
}

func (regExpGroups *PerfMgmt_RegExpGroups) GetEntityData() *types.CommonEntityData {
    regExpGroups.EntityData.YFilter = regExpGroups.YFilter
    regExpGroups.EntityData.YangName = "reg-exp-groups"
    regExpGroups.EntityData.BundleName = "cisco_ios_xr"
    regExpGroups.EntityData.ParentYangName = "perf-mgmt"
    regExpGroups.EntityData.SegmentPath = "reg-exp-groups"
    regExpGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExpGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExpGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExpGroups.EntityData.Children = make(map[string]types.YChild)
    regExpGroups.EntityData.Children["reg-exp-group"] = types.YChild{"RegExpGroup", nil}
    for i := range regExpGroups.RegExpGroup {
        regExpGroups.EntityData.Children[types.GetSegmentPath(&regExpGroups.RegExpGroup[i])] = types.YChild{"RegExpGroup", &regExpGroups.RegExpGroup[i]}
    }
    regExpGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(regExpGroups.EntityData)
}

// PerfMgmt_RegExpGroups_RegExpGroup
// Specify regular expression group name
type PerfMgmt_RegExpGroups_RegExpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Regular expression group name. The type is string
    // with length: 1..32.
    RegExpGroupName interface{}

    // Configure regular expression.
    RegExps PerfMgmt_RegExpGroups_RegExpGroup_RegExps
}

func (regExpGroup *PerfMgmt_RegExpGroups_RegExpGroup) GetEntityData() *types.CommonEntityData {
    regExpGroup.EntityData.YFilter = regExpGroup.YFilter
    regExpGroup.EntityData.YangName = "reg-exp-group"
    regExpGroup.EntityData.BundleName = "cisco_ios_xr"
    regExpGroup.EntityData.ParentYangName = "reg-exp-groups"
    regExpGroup.EntityData.SegmentPath = "reg-exp-group" + "[reg-exp-group-name='" + fmt.Sprintf("%v", regExpGroup.RegExpGroupName) + "']"
    regExpGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExpGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExpGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExpGroup.EntityData.Children = make(map[string]types.YChild)
    regExpGroup.EntityData.Children["reg-exps"] = types.YChild{"RegExps", &regExpGroup.RegExps}
    regExpGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    regExpGroup.EntityData.Leafs["reg-exp-group-name"] = types.YLeaf{"RegExpGroupName", regExpGroup.RegExpGroupName}
    return &(regExpGroup.EntityData)
}

// PerfMgmt_RegExpGroups_RegExpGroup_RegExps
// Configure regular expression
type PerfMgmt_RegExpGroups_RegExpGroup_RegExps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify regular expression index number. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp.
    RegExp []PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp
}

func (regExps *PerfMgmt_RegExpGroups_RegExpGroup_RegExps) GetEntityData() *types.CommonEntityData {
    regExps.EntityData.YFilter = regExps.YFilter
    regExps.EntityData.YangName = "reg-exps"
    regExps.EntityData.BundleName = "cisco_ios_xr"
    regExps.EntityData.ParentYangName = "reg-exp-group"
    regExps.EntityData.SegmentPath = "reg-exps"
    regExps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExps.EntityData.Children = make(map[string]types.YChild)
    regExps.EntityData.Children["reg-exp"] = types.YChild{"RegExp", nil}
    for i := range regExps.RegExp {
        regExps.EntityData.Children[types.GetSegmentPath(&regExps.RegExp[i])] = types.YChild{"RegExp", &regExps.RegExp[i]}
    }
    regExps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(regExps.EntityData)
}

// PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp
// Specify regular expression index number
type PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Regular expression index number. The type is
    // interface{} with range: 1..100.
    RegExpIndex interface{}

    // Regular expression string to match. The type is string with length: 1..128.
    // This attribute is mandatory.
    RegExpString interface{}
}

func (regExp *PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp) GetEntityData() *types.CommonEntityData {
    regExp.EntityData.YFilter = regExp.YFilter
    regExp.EntityData.YangName = "reg-exp"
    regExp.EntityData.BundleName = "cisco_ios_xr"
    regExp.EntityData.ParentYangName = "reg-exps"
    regExp.EntityData.SegmentPath = "reg-exp" + "[reg-exp-index='" + fmt.Sprintf("%v", regExp.RegExpIndex) + "']"
    regExp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExp.EntityData.Children = make(map[string]types.YChild)
    regExp.EntityData.Leafs = make(map[string]types.YLeaf)
    regExp.EntityData.Leafs["reg-exp-index"] = types.YLeaf{"RegExpIndex", regExp.RegExpIndex}
    regExp.EntityData.Leafs["reg-exp-string"] = types.YLeaf{"RegExpString", regExp.RegExpString}
    return &(regExp.EntityData)
}

// PerfMgmt_Threshold
// Container for threshold templates
type PerfMgmt_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *PerfMgmt_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "perf-mgmt"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["generic-counter-interface"] = types.YChild{"GenericCounterInterface", &threshold.GenericCounterInterface}
    threshold.EntityData.Children["ldp-mpls"] = types.YChild{"LdpMpls", &threshold.LdpMpls}
    threshold.EntityData.Children["basic-counter-interface"] = types.YChild{"BasicCounterInterface", &threshold.BasicCounterInterface}
    threshold.EntityData.Children["bgp"] = types.YChild{"Bgp", &threshold.Bgp}
    threshold.EntityData.Children["ospfv2-protocol"] = types.YChild{"Ospfv2Protocol", &threshold.Ospfv2Protocol}
    threshold.EntityData.Children["cpu-node"] = types.YChild{"CpuNode", &threshold.CpuNode}
    threshold.EntityData.Children["data-rate-interface"] = types.YChild{"DataRateInterface", &threshold.DataRateInterface}
    threshold.EntityData.Children["process-node"] = types.YChild{"ProcessNode", &threshold.ProcessNode}
    threshold.EntityData.Children["memory-node"] = types.YChild{"MemoryNode", &threshold.MemoryNode}
    threshold.EntityData.Children["ospfv3-protocol"] = types.YChild{"Ospfv3Protocol", &threshold.Ospfv3Protocol}
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(threshold.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface
// Interface Generic Counter threshold
// configuration
type PerfMgmt_Threshold_GenericCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold templates.
    GenericCounterInterfaceTemplates PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates
}

func (genericCounterInterface *PerfMgmt_Threshold_GenericCounterInterface) GetEntityData() *types.CommonEntityData {
    genericCounterInterface.EntityData.YFilter = genericCounterInterface.YFilter
    genericCounterInterface.EntityData.YangName = "generic-counter-interface"
    genericCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterface.EntityData.ParentYangName = "threshold"
    genericCounterInterface.EntityData.SegmentPath = "generic-counter-interface"
    genericCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterface.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterface.EntityData.Children["generic-counter-interface-templates"] = types.YChild{"GenericCounterInterfaceTemplates", &genericCounterInterface.GenericCounterInterfaceTemplates}
    genericCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates
// Interface Generic Counter threshold templates
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate.
    GenericCounterInterfaceTemplate []PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
}

func (genericCounterInterfaceTemplates *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates) GetEntityData() *types.CommonEntityData {
    genericCounterInterfaceTemplates.EntityData.YFilter = genericCounterInterfaceTemplates.YFilter
    genericCounterInterfaceTemplates.EntityData.YangName = "generic-counter-interface-templates"
    genericCounterInterfaceTemplates.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterfaceTemplates.EntityData.ParentYangName = "generic-counter-interface"
    genericCounterInterfaceTemplates.EntityData.SegmentPath = "generic-counter-interface-templates"
    genericCounterInterfaceTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterfaceTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterfaceTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterfaceTemplates.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterfaceTemplates.EntityData.Children["generic-counter-interface-template"] = types.YChild{"GenericCounterInterfaceTemplate", nil}
    for i := range genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate {
        genericCounterInterfaceTemplates.EntityData.Children[types.GetSegmentPath(&genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i])] = types.YChild{"GenericCounterInterfaceTemplate", &genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i]}
    }
    genericCounterInterfaceTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(genericCounterInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
// Interface Generic Counter threshold template
// instance
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    genericCounterInterfaceTemplate.EntityData.YFilter = genericCounterInterfaceTemplate.YFilter
    genericCounterInterfaceTemplate.EntityData.YangName = "generic-counter-interface-template"
    genericCounterInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterfaceTemplate.EntityData.ParentYangName = "generic-counter-interface-templates"
    genericCounterInterfaceTemplate.EntityData.SegmentPath = "generic-counter-interface-template" + "[template-name='" + fmt.Sprintf("%v", genericCounterInterfaceTemplate.TemplateName) + "']"
    genericCounterInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterfaceTemplate.EntityData.Children = make(map[string]types.YChild)
    genericCounterInterfaceTemplate.EntityData.Children["in-octets"] = types.YChild{"InOctets", &genericCounterInterfaceTemplate.InOctets}
    genericCounterInterfaceTemplate.EntityData.Children["in-ucast-pkts"] = types.YChild{"InUcastPkts", &genericCounterInterfaceTemplate.InUcastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["out-ucast-pkts"] = types.YChild{"OutUcastPkts", &genericCounterInterfaceTemplate.OutUcastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["out-broadcast-pkts"] = types.YChild{"OutBroadcastPkts", &genericCounterInterfaceTemplate.OutBroadcastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["out-multicast-pkts"] = types.YChild{"OutMulticastPkts", &genericCounterInterfaceTemplate.OutMulticastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["input-overrun"] = types.YChild{"InputOverrun", &genericCounterInterfaceTemplate.InputOverrun}
    genericCounterInterfaceTemplate.EntityData.Children["out-octets"] = types.YChild{"OutOctets", &genericCounterInterfaceTemplate.OutOctets}
    genericCounterInterfaceTemplate.EntityData.Children["output-underrun"] = types.YChild{"OutputUnderrun", &genericCounterInterfaceTemplate.OutputUnderrun}
    genericCounterInterfaceTemplate.EntityData.Children["input-total-errors"] = types.YChild{"InputTotalErrors", &genericCounterInterfaceTemplate.InputTotalErrors}
    genericCounterInterfaceTemplate.EntityData.Children["output-total-drops"] = types.YChild{"OutputTotalDrops", &genericCounterInterfaceTemplate.OutputTotalDrops}
    genericCounterInterfaceTemplate.EntityData.Children["input-crc"] = types.YChild{"InputCrc", &genericCounterInterfaceTemplate.InputCrc}
    genericCounterInterfaceTemplate.EntityData.Children["in-broadcast-pkts"] = types.YChild{"InBroadcastPkts", &genericCounterInterfaceTemplate.InBroadcastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["in-multicast-pkts"] = types.YChild{"InMulticastPkts", &genericCounterInterfaceTemplate.InMulticastPkts}
    genericCounterInterfaceTemplate.EntityData.Children["out-packets"] = types.YChild{"OutPackets", &genericCounterInterfaceTemplate.OutPackets}
    genericCounterInterfaceTemplate.EntityData.Children["output-total-errors"] = types.YChild{"OutputTotalErrors", &genericCounterInterfaceTemplate.OutputTotalErrors}
    genericCounterInterfaceTemplate.EntityData.Children["in-packets"] = types.YChild{"InPackets", &genericCounterInterfaceTemplate.InPackets}
    genericCounterInterfaceTemplate.EntityData.Children["input-unknown-proto"] = types.YChild{"InputUnknownProto", &genericCounterInterfaceTemplate.InputUnknownProto}
    genericCounterInterfaceTemplate.EntityData.Children["input-queue-drops"] = types.YChild{"InputQueueDrops", &genericCounterInterfaceTemplate.InputQueueDrops}
    genericCounterInterfaceTemplate.EntityData.Children["input-total-drops"] = types.YChild{"InputTotalDrops", &genericCounterInterfaceTemplate.InputTotalDrops}
    genericCounterInterfaceTemplate.EntityData.Children["input-frame"] = types.YChild{"InputFrame", &genericCounterInterfaceTemplate.InputFrame}
    genericCounterInterfaceTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    genericCounterInterfaceTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", genericCounterInterfaceTemplate.TemplateName}
    genericCounterInterfaceTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", genericCounterInterfaceTemplate.SampleInterval}
    genericCounterInterfaceTemplate.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", genericCounterInterfaceTemplate.RegExpGroup}
    genericCounterInterfaceTemplate.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", genericCounterInterfaceTemplate.VrfGroup}
    return &(genericCounterInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets struct {
    EntityData types.CommonEntityData
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

func (inOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets) GetEntityData() *types.CommonEntityData {
    inOctets.EntityData.YFilter = inOctets.YFilter
    inOctets.EntityData.YangName = "in-octets"
    inOctets.EntityData.BundleName = "cisco_ios_xr"
    inOctets.EntityData.ParentYangName = "generic-counter-interface-template"
    inOctets.EntityData.SegmentPath = "in-octets"
    inOctets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inOctets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inOctets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inOctets.EntityData.Children = make(map[string]types.YChild)
    inOctets.EntityData.Leafs = make(map[string]types.YLeaf)
    inOctets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inOctets.Operator}
    inOctets.EntityData.Leafs["value"] = types.YLeaf{"Value", inOctets.Value}
    inOctets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inOctets.EndRangeValue}
    inOctets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inOctets.Percent}
    inOctets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inOctets.RearmType}
    inOctets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inOctets.RearmWindow}
    return &(inOctets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts
// Number of inbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts struct {
    EntityData types.CommonEntityData
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

func (inUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts) GetEntityData() *types.CommonEntityData {
    inUcastPkts.EntityData.YFilter = inUcastPkts.YFilter
    inUcastPkts.EntityData.YangName = "in-ucast-pkts"
    inUcastPkts.EntityData.BundleName = "cisco_ios_xr"
    inUcastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    inUcastPkts.EntityData.SegmentPath = "in-ucast-pkts"
    inUcastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inUcastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inUcastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inUcastPkts.EntityData.Children = make(map[string]types.YChild)
    inUcastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    inUcastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inUcastPkts.Operator}
    inUcastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", inUcastPkts.Value}
    inUcastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inUcastPkts.EndRangeValue}
    inUcastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inUcastPkts.Percent}
    inUcastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inUcastPkts.RearmType}
    inUcastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inUcastPkts.RearmWindow}
    return &(inUcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts
// Number of outbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts struct {
    EntityData types.CommonEntityData
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

func (outUcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts) GetEntityData() *types.CommonEntityData {
    outUcastPkts.EntityData.YFilter = outUcastPkts.YFilter
    outUcastPkts.EntityData.YangName = "out-ucast-pkts"
    outUcastPkts.EntityData.BundleName = "cisco_ios_xr"
    outUcastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    outUcastPkts.EntityData.SegmentPath = "out-ucast-pkts"
    outUcastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outUcastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outUcastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outUcastPkts.EntityData.Children = make(map[string]types.YChild)
    outUcastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    outUcastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outUcastPkts.Operator}
    outUcastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", outUcastPkts.Value}
    outUcastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outUcastPkts.EndRangeValue}
    outUcastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outUcastPkts.Percent}
    outUcastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outUcastPkts.RearmType}
    outUcastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outUcastPkts.RearmWindow}
    return &(outUcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts
// Number of outbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts struct {
    EntityData types.CommonEntityData
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

func (outBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts) GetEntityData() *types.CommonEntityData {
    outBroadcastPkts.EntityData.YFilter = outBroadcastPkts.YFilter
    outBroadcastPkts.EntityData.YangName = "out-broadcast-pkts"
    outBroadcastPkts.EntityData.BundleName = "cisco_ios_xr"
    outBroadcastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    outBroadcastPkts.EntityData.SegmentPath = "out-broadcast-pkts"
    outBroadcastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outBroadcastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outBroadcastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outBroadcastPkts.EntityData.Children = make(map[string]types.YChild)
    outBroadcastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    outBroadcastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outBroadcastPkts.Operator}
    outBroadcastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", outBroadcastPkts.Value}
    outBroadcastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outBroadcastPkts.EndRangeValue}
    outBroadcastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outBroadcastPkts.Percent}
    outBroadcastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outBroadcastPkts.RearmType}
    outBroadcastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outBroadcastPkts.RearmWindow}
    return &(outBroadcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts
// Number of outbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts struct {
    EntityData types.CommonEntityData
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

func (outMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts) GetEntityData() *types.CommonEntityData {
    outMulticastPkts.EntityData.YFilter = outMulticastPkts.YFilter
    outMulticastPkts.EntityData.YangName = "out-multicast-pkts"
    outMulticastPkts.EntityData.BundleName = "cisco_ios_xr"
    outMulticastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    outMulticastPkts.EntityData.SegmentPath = "out-multicast-pkts"
    outMulticastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outMulticastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outMulticastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outMulticastPkts.EntityData.Children = make(map[string]types.YChild)
    outMulticastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    outMulticastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outMulticastPkts.Operator}
    outMulticastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", outMulticastPkts.Value}
    outMulticastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outMulticastPkts.EndRangeValue}
    outMulticastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outMulticastPkts.Percent}
    outMulticastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outMulticastPkts.RearmType}
    outMulticastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outMulticastPkts.RearmWindow}
    return &(outMulticastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun
// Number of inbound packets with overrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun struct {
    EntityData types.CommonEntityData
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

func (inputOverrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun) GetEntityData() *types.CommonEntityData {
    inputOverrun.EntityData.YFilter = inputOverrun.YFilter
    inputOverrun.EntityData.YangName = "input-overrun"
    inputOverrun.EntityData.BundleName = "cisco_ios_xr"
    inputOverrun.EntityData.ParentYangName = "generic-counter-interface-template"
    inputOverrun.EntityData.SegmentPath = "input-overrun"
    inputOverrun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputOverrun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputOverrun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputOverrun.EntityData.Children = make(map[string]types.YChild)
    inputOverrun.EntityData.Leafs = make(map[string]types.YLeaf)
    inputOverrun.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputOverrun.Operator}
    inputOverrun.EntityData.Leafs["value"] = types.YLeaf{"Value", inputOverrun.Value}
    inputOverrun.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputOverrun.EndRangeValue}
    inputOverrun.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputOverrun.Percent}
    inputOverrun.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputOverrun.RearmType}
    inputOverrun.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputOverrun.RearmWindow}
    return &(inputOverrun.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets struct {
    EntityData types.CommonEntityData
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

func (outOctets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets) GetEntityData() *types.CommonEntityData {
    outOctets.EntityData.YFilter = outOctets.YFilter
    outOctets.EntityData.YangName = "out-octets"
    outOctets.EntityData.BundleName = "cisco_ios_xr"
    outOctets.EntityData.ParentYangName = "generic-counter-interface-template"
    outOctets.EntityData.SegmentPath = "out-octets"
    outOctets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outOctets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outOctets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outOctets.EntityData.Children = make(map[string]types.YChild)
    outOctets.EntityData.Leafs = make(map[string]types.YLeaf)
    outOctets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outOctets.Operator}
    outOctets.EntityData.Leafs["value"] = types.YLeaf{"Value", outOctets.Value}
    outOctets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outOctets.EndRangeValue}
    outOctets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outOctets.Percent}
    outOctets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outOctets.RearmType}
    outOctets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outOctets.RearmWindow}
    return &(outOctets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun
// Number of outbound packets with underrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun struct {
    EntityData types.CommonEntityData
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

func (outputUnderrun *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun) GetEntityData() *types.CommonEntityData {
    outputUnderrun.EntityData.YFilter = outputUnderrun.YFilter
    outputUnderrun.EntityData.YangName = "output-underrun"
    outputUnderrun.EntityData.BundleName = "cisco_ios_xr"
    outputUnderrun.EntityData.ParentYangName = "generic-counter-interface-template"
    outputUnderrun.EntityData.SegmentPath = "output-underrun"
    outputUnderrun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputUnderrun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputUnderrun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputUnderrun.EntityData.Children = make(map[string]types.YChild)
    outputUnderrun.EntityData.Leafs = make(map[string]types.YLeaf)
    outputUnderrun.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputUnderrun.Operator}
    outputUnderrun.EntityData.Leafs["value"] = types.YLeaf{"Value", outputUnderrun.Value}
    outputUnderrun.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputUnderrun.EndRangeValue}
    outputUnderrun.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputUnderrun.Percent}
    outputUnderrun.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputUnderrun.RearmType}
    outputUnderrun.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputUnderrun.RearmWindow}
    return &(outputUnderrun.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors struct {
    EntityData types.CommonEntityData
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

func (inputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors) GetEntityData() *types.CommonEntityData {
    inputTotalErrors.EntityData.YFilter = inputTotalErrors.YFilter
    inputTotalErrors.EntityData.YangName = "input-total-errors"
    inputTotalErrors.EntityData.BundleName = "cisco_ios_xr"
    inputTotalErrors.EntityData.ParentYangName = "generic-counter-interface-template"
    inputTotalErrors.EntityData.SegmentPath = "input-total-errors"
    inputTotalErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputTotalErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputTotalErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputTotalErrors.EntityData.Children = make(map[string]types.YChild)
    inputTotalErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    inputTotalErrors.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputTotalErrors.Operator}
    inputTotalErrors.EntityData.Leafs["value"] = types.YLeaf{"Value", inputTotalErrors.Value}
    inputTotalErrors.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputTotalErrors.EndRangeValue}
    inputTotalErrors.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputTotalErrors.Percent}
    inputTotalErrors.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputTotalErrors.RearmType}
    inputTotalErrors.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputTotalErrors.RearmWindow}
    return &(inputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops struct {
    EntityData types.CommonEntityData
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

func (outputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops) GetEntityData() *types.CommonEntityData {
    outputTotalDrops.EntityData.YFilter = outputTotalDrops.YFilter
    outputTotalDrops.EntityData.YangName = "output-total-drops"
    outputTotalDrops.EntityData.BundleName = "cisco_ios_xr"
    outputTotalDrops.EntityData.ParentYangName = "generic-counter-interface-template"
    outputTotalDrops.EntityData.SegmentPath = "output-total-drops"
    outputTotalDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputTotalDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputTotalDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputTotalDrops.EntityData.Children = make(map[string]types.YChild)
    outputTotalDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    outputTotalDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputTotalDrops.Operator}
    outputTotalDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", outputTotalDrops.Value}
    outputTotalDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputTotalDrops.EndRangeValue}
    outputTotalDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputTotalDrops.Percent}
    outputTotalDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputTotalDrops.RearmType}
    outputTotalDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputTotalDrops.RearmWindow}
    return &(outputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc
// Number of inbound packets discarded with
// incorrect CRC
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc struct {
    EntityData types.CommonEntityData
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

func (inputCrc *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc) GetEntityData() *types.CommonEntityData {
    inputCrc.EntityData.YFilter = inputCrc.YFilter
    inputCrc.EntityData.YangName = "input-crc"
    inputCrc.EntityData.BundleName = "cisco_ios_xr"
    inputCrc.EntityData.ParentYangName = "generic-counter-interface-template"
    inputCrc.EntityData.SegmentPath = "input-crc"
    inputCrc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputCrc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputCrc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputCrc.EntityData.Children = make(map[string]types.YChild)
    inputCrc.EntityData.Leafs = make(map[string]types.YLeaf)
    inputCrc.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputCrc.Operator}
    inputCrc.EntityData.Leafs["value"] = types.YLeaf{"Value", inputCrc.Value}
    inputCrc.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputCrc.EndRangeValue}
    inputCrc.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputCrc.Percent}
    inputCrc.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputCrc.RearmType}
    inputCrc.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputCrc.RearmWindow}
    return &(inputCrc.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts
// Number of inbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts struct {
    EntityData types.CommonEntityData
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

func (inBroadcastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts) GetEntityData() *types.CommonEntityData {
    inBroadcastPkts.EntityData.YFilter = inBroadcastPkts.YFilter
    inBroadcastPkts.EntityData.YangName = "in-broadcast-pkts"
    inBroadcastPkts.EntityData.BundleName = "cisco_ios_xr"
    inBroadcastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    inBroadcastPkts.EntityData.SegmentPath = "in-broadcast-pkts"
    inBroadcastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inBroadcastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inBroadcastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inBroadcastPkts.EntityData.Children = make(map[string]types.YChild)
    inBroadcastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    inBroadcastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inBroadcastPkts.Operator}
    inBroadcastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", inBroadcastPkts.Value}
    inBroadcastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inBroadcastPkts.EndRangeValue}
    inBroadcastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inBroadcastPkts.Percent}
    inBroadcastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inBroadcastPkts.RearmType}
    inBroadcastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inBroadcastPkts.RearmWindow}
    return &(inBroadcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts
// Number of inbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts struct {
    EntityData types.CommonEntityData
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

func (inMulticastPkts *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts) GetEntityData() *types.CommonEntityData {
    inMulticastPkts.EntityData.YFilter = inMulticastPkts.YFilter
    inMulticastPkts.EntityData.YangName = "in-multicast-pkts"
    inMulticastPkts.EntityData.BundleName = "cisco_ios_xr"
    inMulticastPkts.EntityData.ParentYangName = "generic-counter-interface-template"
    inMulticastPkts.EntityData.SegmentPath = "in-multicast-pkts"
    inMulticastPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inMulticastPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inMulticastPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inMulticastPkts.EntityData.Children = make(map[string]types.YChild)
    inMulticastPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    inMulticastPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inMulticastPkts.Operator}
    inMulticastPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", inMulticastPkts.Value}
    inMulticastPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inMulticastPkts.EndRangeValue}
    inMulticastPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inMulticastPkts.Percent}
    inMulticastPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inMulticastPkts.RearmType}
    inMulticastPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inMulticastPkts.RearmWindow}
    return &(inMulticastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets struct {
    EntityData types.CommonEntityData
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

func (outPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets) GetEntityData() *types.CommonEntityData {
    outPackets.EntityData.YFilter = outPackets.YFilter
    outPackets.EntityData.YangName = "out-packets"
    outPackets.EntityData.BundleName = "cisco_ios_xr"
    outPackets.EntityData.ParentYangName = "generic-counter-interface-template"
    outPackets.EntityData.SegmentPath = "out-packets"
    outPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outPackets.EntityData.Children = make(map[string]types.YChild)
    outPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outPackets.Operator}
    outPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outPackets.Value}
    outPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outPackets.EndRangeValue}
    outPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outPackets.Percent}
    outPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outPackets.RearmType}
    outPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outPackets.RearmWindow}
    return &(outPackets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors struct {
    EntityData types.CommonEntityData
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

func (outputTotalErrors *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors) GetEntityData() *types.CommonEntityData {
    outputTotalErrors.EntityData.YFilter = outputTotalErrors.YFilter
    outputTotalErrors.EntityData.YangName = "output-total-errors"
    outputTotalErrors.EntityData.BundleName = "cisco_ios_xr"
    outputTotalErrors.EntityData.ParentYangName = "generic-counter-interface-template"
    outputTotalErrors.EntityData.SegmentPath = "output-total-errors"
    outputTotalErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputTotalErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputTotalErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputTotalErrors.EntityData.Children = make(map[string]types.YChild)
    outputTotalErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    outputTotalErrors.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputTotalErrors.Operator}
    outputTotalErrors.EntityData.Leafs["value"] = types.YLeaf{"Value", outputTotalErrors.Value}
    outputTotalErrors.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputTotalErrors.EndRangeValue}
    outputTotalErrors.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputTotalErrors.Percent}
    outputTotalErrors.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputTotalErrors.RearmType}
    outputTotalErrors.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputTotalErrors.RearmWindow}
    return &(outputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets struct {
    EntityData types.CommonEntityData
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

func (inPackets *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets) GetEntityData() *types.CommonEntityData {
    inPackets.EntityData.YFilter = inPackets.YFilter
    inPackets.EntityData.YangName = "in-packets"
    inPackets.EntityData.BundleName = "cisco_ios_xr"
    inPackets.EntityData.ParentYangName = "generic-counter-interface-template"
    inPackets.EntityData.SegmentPath = "in-packets"
    inPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inPackets.EntityData.Children = make(map[string]types.YChild)
    inPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inPackets.Operator}
    inPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inPackets.Value}
    inPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inPackets.EndRangeValue}
    inPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inPackets.Percent}
    inPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inPackets.RearmType}
    inPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inPackets.RearmWindow}
    return &(inPackets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto
// Number of inbound packets discarded with
// unknown protocol
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto struct {
    EntityData types.CommonEntityData
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

func (inputUnknownProto *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto) GetEntityData() *types.CommonEntityData {
    inputUnknownProto.EntityData.YFilter = inputUnknownProto.YFilter
    inputUnknownProto.EntityData.YangName = "input-unknown-proto"
    inputUnknownProto.EntityData.BundleName = "cisco_ios_xr"
    inputUnknownProto.EntityData.ParentYangName = "generic-counter-interface-template"
    inputUnknownProto.EntityData.SegmentPath = "input-unknown-proto"
    inputUnknownProto.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputUnknownProto.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputUnknownProto.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputUnknownProto.EntityData.Children = make(map[string]types.YChild)
    inputUnknownProto.EntityData.Leafs = make(map[string]types.YLeaf)
    inputUnknownProto.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputUnknownProto.Operator}
    inputUnknownProto.EntityData.Leafs["value"] = types.YLeaf{"Value", inputUnknownProto.Value}
    inputUnknownProto.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputUnknownProto.EndRangeValue}
    inputUnknownProto.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputUnknownProto.Percent}
    inputUnknownProto.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputUnknownProto.RearmType}
    inputUnknownProto.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputUnknownProto.RearmWindow}
    return &(inputUnknownProto.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops struct {
    EntityData types.CommonEntityData
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

func (inputQueueDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops) GetEntityData() *types.CommonEntityData {
    inputQueueDrops.EntityData.YFilter = inputQueueDrops.YFilter
    inputQueueDrops.EntityData.YangName = "input-queue-drops"
    inputQueueDrops.EntityData.BundleName = "cisco_ios_xr"
    inputQueueDrops.EntityData.ParentYangName = "generic-counter-interface-template"
    inputQueueDrops.EntityData.SegmentPath = "input-queue-drops"
    inputQueueDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputQueueDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputQueueDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputQueueDrops.EntityData.Children = make(map[string]types.YChild)
    inputQueueDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    inputQueueDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputQueueDrops.Operator}
    inputQueueDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", inputQueueDrops.Value}
    inputQueueDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputQueueDrops.EndRangeValue}
    inputQueueDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputQueueDrops.Percent}
    inputQueueDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputQueueDrops.RearmType}
    inputQueueDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputQueueDrops.RearmWindow}
    return &(inputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops struct {
    EntityData types.CommonEntityData
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

func (inputTotalDrops *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops) GetEntityData() *types.CommonEntityData {
    inputTotalDrops.EntityData.YFilter = inputTotalDrops.YFilter
    inputTotalDrops.EntityData.YangName = "input-total-drops"
    inputTotalDrops.EntityData.BundleName = "cisco_ios_xr"
    inputTotalDrops.EntityData.ParentYangName = "generic-counter-interface-template"
    inputTotalDrops.EntityData.SegmentPath = "input-total-drops"
    inputTotalDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputTotalDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputTotalDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputTotalDrops.EntityData.Children = make(map[string]types.YChild)
    inputTotalDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    inputTotalDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputTotalDrops.Operator}
    inputTotalDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", inputTotalDrops.Value}
    inputTotalDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputTotalDrops.EndRangeValue}
    inputTotalDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputTotalDrops.Percent}
    inputTotalDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputTotalDrops.RearmType}
    inputTotalDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputTotalDrops.RearmWindow}
    return &(inputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame
// Number of inbound packets with framing
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame struct {
    EntityData types.CommonEntityData
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

func (inputFrame *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame) GetEntityData() *types.CommonEntityData {
    inputFrame.EntityData.YFilter = inputFrame.YFilter
    inputFrame.EntityData.YangName = "input-frame"
    inputFrame.EntityData.BundleName = "cisco_ios_xr"
    inputFrame.EntityData.ParentYangName = "generic-counter-interface-template"
    inputFrame.EntityData.SegmentPath = "input-frame"
    inputFrame.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputFrame.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputFrame.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputFrame.EntityData.Children = make(map[string]types.YChild)
    inputFrame.EntityData.Leafs = make(map[string]types.YLeaf)
    inputFrame.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputFrame.Operator}
    inputFrame.EntityData.Leafs["value"] = types.YLeaf{"Value", inputFrame.Value}
    inputFrame.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputFrame.EndRangeValue}
    inputFrame.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputFrame.Percent}
    inputFrame.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputFrame.RearmType}
    inputFrame.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputFrame.RearmWindow}
    return &(inputFrame.EntityData)
}

// PerfMgmt_Threshold_LdpMpls
// MPLS LDP threshold configuration
type PerfMgmt_Threshold_LdpMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP threshold templates.
    LdpMplsTemplates PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates
}

func (ldpMpls *PerfMgmt_Threshold_LdpMpls) GetEntityData() *types.CommonEntityData {
    ldpMpls.EntityData.YFilter = ldpMpls.YFilter
    ldpMpls.EntityData.YangName = "ldp-mpls"
    ldpMpls.EntityData.BundleName = "cisco_ios_xr"
    ldpMpls.EntityData.ParentYangName = "threshold"
    ldpMpls.EntityData.SegmentPath = "ldp-mpls"
    ldpMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMpls.EntityData.Children = make(map[string]types.YChild)
    ldpMpls.EntityData.Children["ldp-mpls-templates"] = types.YChild{"LdpMplsTemplates", &ldpMpls.LdpMplsTemplates}
    ldpMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldpMpls.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates
// MPLS LDP threshold templates
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate.
    LdpMplsTemplate []PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
}

func (ldpMplsTemplates *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates) GetEntityData() *types.CommonEntityData {
    ldpMplsTemplates.EntityData.YFilter = ldpMplsTemplates.YFilter
    ldpMplsTemplates.EntityData.YangName = "ldp-mpls-templates"
    ldpMplsTemplates.EntityData.BundleName = "cisco_ios_xr"
    ldpMplsTemplates.EntityData.ParentYangName = "ldp-mpls"
    ldpMplsTemplates.EntityData.SegmentPath = "ldp-mpls-templates"
    ldpMplsTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMplsTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMplsTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMplsTemplates.EntityData.Children = make(map[string]types.YChild)
    ldpMplsTemplates.EntityData.Children["ldp-mpls-template"] = types.YChild{"LdpMplsTemplate", nil}
    for i := range ldpMplsTemplates.LdpMplsTemplate {
        ldpMplsTemplates.EntityData.Children[types.GetSegmentPath(&ldpMplsTemplates.LdpMplsTemplate[i])] = types.YChild{"LdpMplsTemplate", &ldpMplsTemplates.LdpMplsTemplate[i]}
    }
    ldpMplsTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldpMplsTemplates.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
// MPLS LDP threshold template instance
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetEntityData() *types.CommonEntityData {
    ldpMplsTemplate.EntityData.YFilter = ldpMplsTemplate.YFilter
    ldpMplsTemplate.EntityData.YangName = "ldp-mpls-template"
    ldpMplsTemplate.EntityData.BundleName = "cisco_ios_xr"
    ldpMplsTemplate.EntityData.ParentYangName = "ldp-mpls-templates"
    ldpMplsTemplate.EntityData.SegmentPath = "ldp-mpls-template" + "[template-name='" + fmt.Sprintf("%v", ldpMplsTemplate.TemplateName) + "']"
    ldpMplsTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMplsTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMplsTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMplsTemplate.EntityData.Children = make(map[string]types.YChild)
    ldpMplsTemplate.EntityData.Children["address-withdraw-msgs-rcvd"] = types.YChild{"AddressWithdrawMsgsRcvd", &ldpMplsTemplate.AddressWithdrawMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["label-withdraw-msgs-rcvd"] = types.YChild{"LabelWithdrawMsgsRcvd", &ldpMplsTemplate.LabelWithdrawMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["address-withdraw-msgs-sent"] = types.YChild{"AddressWithdrawMsgsSent", &ldpMplsTemplate.AddressWithdrawMsgsSent}
    ldpMplsTemplate.EntityData.Children["label-withdraw-msgs-sent"] = types.YChild{"LabelWithdrawMsgsSent", &ldpMplsTemplate.LabelWithdrawMsgsSent}
    ldpMplsTemplate.EntityData.Children["notification-msgs-rcvd"] = types.YChild{"NotificationMsgsRcvd", &ldpMplsTemplate.NotificationMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["total-msgs-rcvd"] = types.YChild{"TotalMsgsRcvd", &ldpMplsTemplate.TotalMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["notification-msgs-sent"] = types.YChild{"NotificationMsgsSent", &ldpMplsTemplate.NotificationMsgsSent}
    ldpMplsTemplate.EntityData.Children["total-msgs-sent"] = types.YChild{"TotalMsgsSent", &ldpMplsTemplate.TotalMsgsSent}
    ldpMplsTemplate.EntityData.Children["label-release-msgs-rcvd"] = types.YChild{"LabelReleaseMsgsRcvd", &ldpMplsTemplate.LabelReleaseMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["init-msgs-rcvd"] = types.YChild{"InitMsgsRcvd", &ldpMplsTemplate.InitMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["label-release-msgs-sent"] = types.YChild{"LabelReleaseMsgsSent", &ldpMplsTemplate.LabelReleaseMsgsSent}
    ldpMplsTemplate.EntityData.Children["init-msgs-sent"] = types.YChild{"InitMsgsSent", &ldpMplsTemplate.InitMsgsSent}
    ldpMplsTemplate.EntityData.Children["label-mapping-msgs-rcvd"] = types.YChild{"LabelMappingMsgsRcvd", &ldpMplsTemplate.LabelMappingMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["keepalive-msgs-rcvd"] = types.YChild{"KeepaliveMsgsRcvd", &ldpMplsTemplate.KeepaliveMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["label-mapping-msgs-sent"] = types.YChild{"LabelMappingMsgsSent", &ldpMplsTemplate.LabelMappingMsgsSent}
    ldpMplsTemplate.EntityData.Children["keepalive-msgs-sent"] = types.YChild{"KeepaliveMsgsSent", &ldpMplsTemplate.KeepaliveMsgsSent}
    ldpMplsTemplate.EntityData.Children["address-msgs-rcvd"] = types.YChild{"AddressMsgsRcvd", &ldpMplsTemplate.AddressMsgsRcvd}
    ldpMplsTemplate.EntityData.Children["address-msgs-sent"] = types.YChild{"AddressMsgsSent", &ldpMplsTemplate.AddressMsgsSent}
    ldpMplsTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    ldpMplsTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ldpMplsTemplate.TemplateName}
    ldpMplsTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", ldpMplsTemplate.SampleInterval}
    return &(ldpMplsTemplate.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd
// Number of Address Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetEntityData() *types.CommonEntityData {
    addressWithdrawMsgsRcvd.EntityData.YFilter = addressWithdrawMsgsRcvd.YFilter
    addressWithdrawMsgsRcvd.EntityData.YangName = "address-withdraw-msgs-rcvd"
    addressWithdrawMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    addressWithdrawMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    addressWithdrawMsgsRcvd.EntityData.SegmentPath = "address-withdraw-msgs-rcvd"
    addressWithdrawMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressWithdrawMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressWithdrawMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressWithdrawMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    addressWithdrawMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    addressWithdrawMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", addressWithdrawMsgsRcvd.Operator}
    addressWithdrawMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", addressWithdrawMsgsRcvd.Value}
    addressWithdrawMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", addressWithdrawMsgsRcvd.EndRangeValue}
    addressWithdrawMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", addressWithdrawMsgsRcvd.Percent}
    addressWithdrawMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", addressWithdrawMsgsRcvd.RearmType}
    addressWithdrawMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", addressWithdrawMsgsRcvd.RearmWindow}
    return &(addressWithdrawMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd
// Number of Label Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelWithdrawMsgsRcvd.EntityData.YFilter = labelWithdrawMsgsRcvd.YFilter
    labelWithdrawMsgsRcvd.EntityData.YangName = "label-withdraw-msgs-rcvd"
    labelWithdrawMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelWithdrawMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelWithdrawMsgsRcvd.EntityData.SegmentPath = "label-withdraw-msgs-rcvd"
    labelWithdrawMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelWithdrawMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelWithdrawMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelWithdrawMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    labelWithdrawMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    labelWithdrawMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelWithdrawMsgsRcvd.Operator}
    labelWithdrawMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", labelWithdrawMsgsRcvd.Value}
    labelWithdrawMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelWithdrawMsgsRcvd.EndRangeValue}
    labelWithdrawMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelWithdrawMsgsRcvd.Percent}
    labelWithdrawMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelWithdrawMsgsRcvd.RearmType}
    labelWithdrawMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelWithdrawMsgsRcvd.RearmWindow}
    return &(labelWithdrawMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent
// Number of Address Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent struct {
    EntityData types.CommonEntityData
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

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetEntityData() *types.CommonEntityData {
    addressWithdrawMsgsSent.EntityData.YFilter = addressWithdrawMsgsSent.YFilter
    addressWithdrawMsgsSent.EntityData.YangName = "address-withdraw-msgs-sent"
    addressWithdrawMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    addressWithdrawMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    addressWithdrawMsgsSent.EntityData.SegmentPath = "address-withdraw-msgs-sent"
    addressWithdrawMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressWithdrawMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressWithdrawMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressWithdrawMsgsSent.EntityData.Children = make(map[string]types.YChild)
    addressWithdrawMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    addressWithdrawMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", addressWithdrawMsgsSent.Operator}
    addressWithdrawMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", addressWithdrawMsgsSent.Value}
    addressWithdrawMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", addressWithdrawMsgsSent.EndRangeValue}
    addressWithdrawMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", addressWithdrawMsgsSent.Percent}
    addressWithdrawMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", addressWithdrawMsgsSent.RearmType}
    addressWithdrawMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", addressWithdrawMsgsSent.RearmWindow}
    return &(addressWithdrawMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent
// Number of Label Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent struct {
    EntityData types.CommonEntityData
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

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetEntityData() *types.CommonEntityData {
    labelWithdrawMsgsSent.EntityData.YFilter = labelWithdrawMsgsSent.YFilter
    labelWithdrawMsgsSent.EntityData.YangName = "label-withdraw-msgs-sent"
    labelWithdrawMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelWithdrawMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelWithdrawMsgsSent.EntityData.SegmentPath = "label-withdraw-msgs-sent"
    labelWithdrawMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelWithdrawMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelWithdrawMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelWithdrawMsgsSent.EntityData.Children = make(map[string]types.YChild)
    labelWithdrawMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    labelWithdrawMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelWithdrawMsgsSent.Operator}
    labelWithdrawMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", labelWithdrawMsgsSent.Value}
    labelWithdrawMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelWithdrawMsgsSent.EndRangeValue}
    labelWithdrawMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelWithdrawMsgsSent.Percent}
    labelWithdrawMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelWithdrawMsgsSent.RearmType}
    labelWithdrawMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelWithdrawMsgsSent.RearmWindow}
    return &(labelWithdrawMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd
// Number of Notification messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetEntityData() *types.CommonEntityData {
    notificationMsgsRcvd.EntityData.YFilter = notificationMsgsRcvd.YFilter
    notificationMsgsRcvd.EntityData.YangName = "notification-msgs-rcvd"
    notificationMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    notificationMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    notificationMsgsRcvd.EntityData.SegmentPath = "notification-msgs-rcvd"
    notificationMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    notificationMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    notificationMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", notificationMsgsRcvd.Operator}
    notificationMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", notificationMsgsRcvd.Value}
    notificationMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", notificationMsgsRcvd.EndRangeValue}
    notificationMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", notificationMsgsRcvd.Percent}
    notificationMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", notificationMsgsRcvd.RearmType}
    notificationMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", notificationMsgsRcvd.RearmWindow}
    return &(notificationMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd
// Total number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (totalMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd) GetEntityData() *types.CommonEntityData {
    totalMsgsRcvd.EntityData.YFilter = totalMsgsRcvd.YFilter
    totalMsgsRcvd.EntityData.YangName = "total-msgs-rcvd"
    totalMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    totalMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    totalMsgsRcvd.EntityData.SegmentPath = "total-msgs-rcvd"
    totalMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    totalMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    totalMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    totalMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    totalMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    totalMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", totalMsgsRcvd.Operator}
    totalMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", totalMsgsRcvd.Value}
    totalMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", totalMsgsRcvd.EndRangeValue}
    totalMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", totalMsgsRcvd.Percent}
    totalMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", totalMsgsRcvd.RearmType}
    totalMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", totalMsgsRcvd.RearmWindow}
    return &(totalMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent
// Number of Notification messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent struct {
    EntityData types.CommonEntityData
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

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetEntityData() *types.CommonEntityData {
    notificationMsgsSent.EntityData.YFilter = notificationMsgsSent.YFilter
    notificationMsgsSent.EntityData.YangName = "notification-msgs-sent"
    notificationMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    notificationMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    notificationMsgsSent.EntityData.SegmentPath = "notification-msgs-sent"
    notificationMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationMsgsSent.EntityData.Children = make(map[string]types.YChild)
    notificationMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    notificationMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", notificationMsgsSent.Operator}
    notificationMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", notificationMsgsSent.Value}
    notificationMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", notificationMsgsSent.EndRangeValue}
    notificationMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", notificationMsgsSent.Percent}
    notificationMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", notificationMsgsSent.RearmType}
    notificationMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", notificationMsgsSent.RearmWindow}
    return &(notificationMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent
// Total number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent struct {
    EntityData types.CommonEntityData
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

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetEntityData() *types.CommonEntityData {
    totalMsgsSent.EntityData.YFilter = totalMsgsSent.YFilter
    totalMsgsSent.EntityData.YangName = "total-msgs-sent"
    totalMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    totalMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    totalMsgsSent.EntityData.SegmentPath = "total-msgs-sent"
    totalMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    totalMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    totalMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    totalMsgsSent.EntityData.Children = make(map[string]types.YChild)
    totalMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    totalMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", totalMsgsSent.Operator}
    totalMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", totalMsgsSent.Value}
    totalMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", totalMsgsSent.EndRangeValue}
    totalMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", totalMsgsSent.Percent}
    totalMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", totalMsgsSent.RearmType}
    totalMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", totalMsgsSent.RearmWindow}
    return &(totalMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd
// Number of LAbel Release messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelReleaseMsgsRcvd.EntityData.YFilter = labelReleaseMsgsRcvd.YFilter
    labelReleaseMsgsRcvd.EntityData.YangName = "label-release-msgs-rcvd"
    labelReleaseMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelReleaseMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelReleaseMsgsRcvd.EntityData.SegmentPath = "label-release-msgs-rcvd"
    labelReleaseMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelReleaseMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelReleaseMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelReleaseMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    labelReleaseMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    labelReleaseMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelReleaseMsgsRcvd.Operator}
    labelReleaseMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", labelReleaseMsgsRcvd.Value}
    labelReleaseMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelReleaseMsgsRcvd.EndRangeValue}
    labelReleaseMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelReleaseMsgsRcvd.Percent}
    labelReleaseMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelReleaseMsgsRcvd.RearmType}
    labelReleaseMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelReleaseMsgsRcvd.RearmWindow}
    return &(labelReleaseMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd
// Number of Init messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetEntityData() *types.CommonEntityData {
    initMsgsRcvd.EntityData.YFilter = initMsgsRcvd.YFilter
    initMsgsRcvd.EntityData.YangName = "init-msgs-rcvd"
    initMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    initMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    initMsgsRcvd.EntityData.SegmentPath = "init-msgs-rcvd"
    initMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    initMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    initMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", initMsgsRcvd.Operator}
    initMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", initMsgsRcvd.Value}
    initMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", initMsgsRcvd.EndRangeValue}
    initMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", initMsgsRcvd.Percent}
    initMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", initMsgsRcvd.RearmType}
    initMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", initMsgsRcvd.RearmWindow}
    return &(initMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent
// Number of Label Release messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent struct {
    EntityData types.CommonEntityData
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

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetEntityData() *types.CommonEntityData {
    labelReleaseMsgsSent.EntityData.YFilter = labelReleaseMsgsSent.YFilter
    labelReleaseMsgsSent.EntityData.YangName = "label-release-msgs-sent"
    labelReleaseMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelReleaseMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelReleaseMsgsSent.EntityData.SegmentPath = "label-release-msgs-sent"
    labelReleaseMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelReleaseMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelReleaseMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelReleaseMsgsSent.EntityData.Children = make(map[string]types.YChild)
    labelReleaseMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    labelReleaseMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelReleaseMsgsSent.Operator}
    labelReleaseMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", labelReleaseMsgsSent.Value}
    labelReleaseMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelReleaseMsgsSent.EndRangeValue}
    labelReleaseMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelReleaseMsgsSent.Percent}
    labelReleaseMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelReleaseMsgsSent.RearmType}
    labelReleaseMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelReleaseMsgsSent.RearmWindow}
    return &(labelReleaseMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent
// Number of Init messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent struct {
    EntityData types.CommonEntityData
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

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetEntityData() *types.CommonEntityData {
    initMsgsSent.EntityData.YFilter = initMsgsSent.YFilter
    initMsgsSent.EntityData.YangName = "init-msgs-sent"
    initMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    initMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    initMsgsSent.EntityData.SegmentPath = "init-msgs-sent"
    initMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initMsgsSent.EntityData.Children = make(map[string]types.YChild)
    initMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    initMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", initMsgsSent.Operator}
    initMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", initMsgsSent.Value}
    initMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", initMsgsSent.EndRangeValue}
    initMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", initMsgsSent.Percent}
    initMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", initMsgsSent.RearmType}
    initMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", initMsgsSent.RearmWindow}
    return &(initMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd
// Number of Label Mapping messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelMappingMsgsRcvd.EntityData.YFilter = labelMappingMsgsRcvd.YFilter
    labelMappingMsgsRcvd.EntityData.YangName = "label-mapping-msgs-rcvd"
    labelMappingMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelMappingMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelMappingMsgsRcvd.EntityData.SegmentPath = "label-mapping-msgs-rcvd"
    labelMappingMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMappingMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMappingMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMappingMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    labelMappingMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    labelMappingMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelMappingMsgsRcvd.Operator}
    labelMappingMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", labelMappingMsgsRcvd.Value}
    labelMappingMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelMappingMsgsRcvd.EndRangeValue}
    labelMappingMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelMappingMsgsRcvd.Percent}
    labelMappingMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelMappingMsgsRcvd.RearmType}
    labelMappingMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelMappingMsgsRcvd.RearmWindow}
    return &(labelMappingMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd
// Number of Keepalive messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetEntityData() *types.CommonEntityData {
    keepaliveMsgsRcvd.EntityData.YFilter = keepaliveMsgsRcvd.YFilter
    keepaliveMsgsRcvd.EntityData.YangName = "keepalive-msgs-rcvd"
    keepaliveMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    keepaliveMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    keepaliveMsgsRcvd.EntityData.SegmentPath = "keepalive-msgs-rcvd"
    keepaliveMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepaliveMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepaliveMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepaliveMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    keepaliveMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    keepaliveMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", keepaliveMsgsRcvd.Operator}
    keepaliveMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", keepaliveMsgsRcvd.Value}
    keepaliveMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", keepaliveMsgsRcvd.EndRangeValue}
    keepaliveMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", keepaliveMsgsRcvd.Percent}
    keepaliveMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", keepaliveMsgsRcvd.RearmType}
    keepaliveMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", keepaliveMsgsRcvd.RearmWindow}
    return &(keepaliveMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent
// Number of Label Mapping messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent struct {
    EntityData types.CommonEntityData
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

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetEntityData() *types.CommonEntityData {
    labelMappingMsgsSent.EntityData.YFilter = labelMappingMsgsSent.YFilter
    labelMappingMsgsSent.EntityData.YangName = "label-mapping-msgs-sent"
    labelMappingMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelMappingMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelMappingMsgsSent.EntityData.SegmentPath = "label-mapping-msgs-sent"
    labelMappingMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMappingMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMappingMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMappingMsgsSent.EntityData.Children = make(map[string]types.YChild)
    labelMappingMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    labelMappingMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", labelMappingMsgsSent.Operator}
    labelMappingMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", labelMappingMsgsSent.Value}
    labelMappingMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", labelMappingMsgsSent.EndRangeValue}
    labelMappingMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", labelMappingMsgsSent.Percent}
    labelMappingMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", labelMappingMsgsSent.RearmType}
    labelMappingMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", labelMappingMsgsSent.RearmWindow}
    return &(labelMappingMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent
// Number of Keepalive messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent struct {
    EntityData types.CommonEntityData
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

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetEntityData() *types.CommonEntityData {
    keepaliveMsgsSent.EntityData.YFilter = keepaliveMsgsSent.YFilter
    keepaliveMsgsSent.EntityData.YangName = "keepalive-msgs-sent"
    keepaliveMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    keepaliveMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    keepaliveMsgsSent.EntityData.SegmentPath = "keepalive-msgs-sent"
    keepaliveMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepaliveMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepaliveMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepaliveMsgsSent.EntityData.Children = make(map[string]types.YChild)
    keepaliveMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    keepaliveMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", keepaliveMsgsSent.Operator}
    keepaliveMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", keepaliveMsgsSent.Value}
    keepaliveMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", keepaliveMsgsSent.EndRangeValue}
    keepaliveMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", keepaliveMsgsSent.Percent}
    keepaliveMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", keepaliveMsgsSent.RearmType}
    keepaliveMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", keepaliveMsgsSent.RearmWindow}
    return &(keepaliveMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd
// Number of Address messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd struct {
    EntityData types.CommonEntityData
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

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetEntityData() *types.CommonEntityData {
    addressMsgsRcvd.EntityData.YFilter = addressMsgsRcvd.YFilter
    addressMsgsRcvd.EntityData.YangName = "address-msgs-rcvd"
    addressMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    addressMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    addressMsgsRcvd.EntityData.SegmentPath = "address-msgs-rcvd"
    addressMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressMsgsRcvd.EntityData.Children = make(map[string]types.YChild)
    addressMsgsRcvd.EntityData.Leafs = make(map[string]types.YLeaf)
    addressMsgsRcvd.EntityData.Leafs["operator"] = types.YLeaf{"Operator", addressMsgsRcvd.Operator}
    addressMsgsRcvd.EntityData.Leafs["value"] = types.YLeaf{"Value", addressMsgsRcvd.Value}
    addressMsgsRcvd.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", addressMsgsRcvd.EndRangeValue}
    addressMsgsRcvd.EntityData.Leafs["percent"] = types.YLeaf{"Percent", addressMsgsRcvd.Percent}
    addressMsgsRcvd.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", addressMsgsRcvd.RearmType}
    addressMsgsRcvd.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", addressMsgsRcvd.RearmWindow}
    return &(addressMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent
// Number of Address messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent struct {
    EntityData types.CommonEntityData
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

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetEntityData() *types.CommonEntityData {
    addressMsgsSent.EntityData.YFilter = addressMsgsSent.YFilter
    addressMsgsSent.EntityData.YangName = "address-msgs-sent"
    addressMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    addressMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    addressMsgsSent.EntityData.SegmentPath = "address-msgs-sent"
    addressMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressMsgsSent.EntityData.Children = make(map[string]types.YChild)
    addressMsgsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    addressMsgsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", addressMsgsSent.Operator}
    addressMsgsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", addressMsgsSent.Value}
    addressMsgsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", addressMsgsSent.EndRangeValue}
    addressMsgsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", addressMsgsSent.Percent}
    addressMsgsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", addressMsgsSent.RearmType}
    addressMsgsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", addressMsgsSent.RearmWindow}
    return &(addressMsgsSent.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface
// Interface Basic Counter threshold configuration
type PerfMgmt_Threshold_BasicCounterInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Basic Counter threshold templates.
    BasicCounterInterfaceTemplates PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates
}

func (basicCounterInterface *PerfMgmt_Threshold_BasicCounterInterface) GetEntityData() *types.CommonEntityData {
    basicCounterInterface.EntityData.YFilter = basicCounterInterface.YFilter
    basicCounterInterface.EntityData.YangName = "basic-counter-interface"
    basicCounterInterface.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterface.EntityData.ParentYangName = "threshold"
    basicCounterInterface.EntityData.SegmentPath = "basic-counter-interface"
    basicCounterInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterface.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterface.EntityData.Children["basic-counter-interface-templates"] = types.YChild{"BasicCounterInterfaceTemplates", &basicCounterInterface.BasicCounterInterfaceTemplates}
    basicCounterInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates
// Interface Basic Counter threshold templates
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Basic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate.
    BasicCounterInterfaceTemplate []PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
}

func (basicCounterInterfaceTemplates *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates) GetEntityData() *types.CommonEntityData {
    basicCounterInterfaceTemplates.EntityData.YFilter = basicCounterInterfaceTemplates.YFilter
    basicCounterInterfaceTemplates.EntityData.YangName = "basic-counter-interface-templates"
    basicCounterInterfaceTemplates.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterfaceTemplates.EntityData.ParentYangName = "basic-counter-interface"
    basicCounterInterfaceTemplates.EntityData.SegmentPath = "basic-counter-interface-templates"
    basicCounterInterfaceTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterfaceTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterfaceTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterfaceTemplates.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterfaceTemplates.EntityData.Children["basic-counter-interface-template"] = types.YChild{"BasicCounterInterfaceTemplate", nil}
    for i := range basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate {
        basicCounterInterfaceTemplates.EntityData.Children[types.GetSegmentPath(&basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i])] = types.YChild{"BasicCounterInterfaceTemplate", &basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i]}
    }
    basicCounterInterfaceTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(basicCounterInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
// Interface Basic Counter threshold template
// instance
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    basicCounterInterfaceTemplate.EntityData.YFilter = basicCounterInterfaceTemplate.YFilter
    basicCounterInterfaceTemplate.EntityData.YangName = "basic-counter-interface-template"
    basicCounterInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterfaceTemplate.EntityData.ParentYangName = "basic-counter-interface-templates"
    basicCounterInterfaceTemplate.EntityData.SegmentPath = "basic-counter-interface-template" + "[template-name='" + fmt.Sprintf("%v", basicCounterInterfaceTemplate.TemplateName) + "']"
    basicCounterInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterfaceTemplate.EntityData.Children = make(map[string]types.YChild)
    basicCounterInterfaceTemplate.EntityData.Children["in-octets"] = types.YChild{"InOctets", &basicCounterInterfaceTemplate.InOctets}
    basicCounterInterfaceTemplate.EntityData.Children["out-octets"] = types.YChild{"OutOctets", &basicCounterInterfaceTemplate.OutOctets}
    basicCounterInterfaceTemplate.EntityData.Children["output-queue-drops"] = types.YChild{"OutputQueueDrops", &basicCounterInterfaceTemplate.OutputQueueDrops}
    basicCounterInterfaceTemplate.EntityData.Children["input-total-errors"] = types.YChild{"InputTotalErrors", &basicCounterInterfaceTemplate.InputTotalErrors}
    basicCounterInterfaceTemplate.EntityData.Children["output-total-drops"] = types.YChild{"OutputTotalDrops", &basicCounterInterfaceTemplate.OutputTotalDrops}
    basicCounterInterfaceTemplate.EntityData.Children["out-packets"] = types.YChild{"OutPackets", &basicCounterInterfaceTemplate.OutPackets}
    basicCounterInterfaceTemplate.EntityData.Children["output-total-errors"] = types.YChild{"OutputTotalErrors", &basicCounterInterfaceTemplate.OutputTotalErrors}
    basicCounterInterfaceTemplate.EntityData.Children["in-packets"] = types.YChild{"InPackets", &basicCounterInterfaceTemplate.InPackets}
    basicCounterInterfaceTemplate.EntityData.Children["input-queue-drops"] = types.YChild{"InputQueueDrops", &basicCounterInterfaceTemplate.InputQueueDrops}
    basicCounterInterfaceTemplate.EntityData.Children["input-total-drops"] = types.YChild{"InputTotalDrops", &basicCounterInterfaceTemplate.InputTotalDrops}
    basicCounterInterfaceTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    basicCounterInterfaceTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", basicCounterInterfaceTemplate.TemplateName}
    basicCounterInterfaceTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", basicCounterInterfaceTemplate.SampleInterval}
    basicCounterInterfaceTemplate.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", basicCounterInterfaceTemplate.RegExpGroup}
    basicCounterInterfaceTemplate.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", basicCounterInterfaceTemplate.VrfGroup}
    return &(basicCounterInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets struct {
    EntityData types.CommonEntityData
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

func (inOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets) GetEntityData() *types.CommonEntityData {
    inOctets.EntityData.YFilter = inOctets.YFilter
    inOctets.EntityData.YangName = "in-octets"
    inOctets.EntityData.BundleName = "cisco_ios_xr"
    inOctets.EntityData.ParentYangName = "basic-counter-interface-template"
    inOctets.EntityData.SegmentPath = "in-octets"
    inOctets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inOctets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inOctets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inOctets.EntityData.Children = make(map[string]types.YChild)
    inOctets.EntityData.Leafs = make(map[string]types.YLeaf)
    inOctets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inOctets.Operator}
    inOctets.EntityData.Leafs["value"] = types.YLeaf{"Value", inOctets.Value}
    inOctets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inOctets.EndRangeValue}
    inOctets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inOctets.Percent}
    inOctets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inOctets.RearmType}
    inOctets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inOctets.RearmWindow}
    return &(inOctets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets struct {
    EntityData types.CommonEntityData
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

func (outOctets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets) GetEntityData() *types.CommonEntityData {
    outOctets.EntityData.YFilter = outOctets.YFilter
    outOctets.EntityData.YangName = "out-octets"
    outOctets.EntityData.BundleName = "cisco_ios_xr"
    outOctets.EntityData.ParentYangName = "basic-counter-interface-template"
    outOctets.EntityData.SegmentPath = "out-octets"
    outOctets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outOctets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outOctets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outOctets.EntityData.Children = make(map[string]types.YChild)
    outOctets.EntityData.Leafs = make(map[string]types.YLeaf)
    outOctets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outOctets.Operator}
    outOctets.EntityData.Leafs["value"] = types.YLeaf{"Value", outOctets.Value}
    outOctets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outOctets.EndRangeValue}
    outOctets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outOctets.Percent}
    outOctets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outOctets.RearmType}
    outOctets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outOctets.RearmWindow}
    return &(outOctets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops
// Number of outbound queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops struct {
    EntityData types.CommonEntityData
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

func (outputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops) GetEntityData() *types.CommonEntityData {
    outputQueueDrops.EntityData.YFilter = outputQueueDrops.YFilter
    outputQueueDrops.EntityData.YangName = "output-queue-drops"
    outputQueueDrops.EntityData.BundleName = "cisco_ios_xr"
    outputQueueDrops.EntityData.ParentYangName = "basic-counter-interface-template"
    outputQueueDrops.EntityData.SegmentPath = "output-queue-drops"
    outputQueueDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputQueueDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputQueueDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputQueueDrops.EntityData.Children = make(map[string]types.YChild)
    outputQueueDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    outputQueueDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputQueueDrops.Operator}
    outputQueueDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", outputQueueDrops.Value}
    outputQueueDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputQueueDrops.EndRangeValue}
    outputQueueDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputQueueDrops.Percent}
    outputQueueDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputQueueDrops.RearmType}
    outputQueueDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputQueueDrops.RearmWindow}
    return &(outputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors struct {
    EntityData types.CommonEntityData
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

func (inputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors) GetEntityData() *types.CommonEntityData {
    inputTotalErrors.EntityData.YFilter = inputTotalErrors.YFilter
    inputTotalErrors.EntityData.YangName = "input-total-errors"
    inputTotalErrors.EntityData.BundleName = "cisco_ios_xr"
    inputTotalErrors.EntityData.ParentYangName = "basic-counter-interface-template"
    inputTotalErrors.EntityData.SegmentPath = "input-total-errors"
    inputTotalErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputTotalErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputTotalErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputTotalErrors.EntityData.Children = make(map[string]types.YChild)
    inputTotalErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    inputTotalErrors.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputTotalErrors.Operator}
    inputTotalErrors.EntityData.Leafs["value"] = types.YLeaf{"Value", inputTotalErrors.Value}
    inputTotalErrors.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputTotalErrors.EndRangeValue}
    inputTotalErrors.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputTotalErrors.Percent}
    inputTotalErrors.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputTotalErrors.RearmType}
    inputTotalErrors.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputTotalErrors.RearmWindow}
    return &(inputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops struct {
    EntityData types.CommonEntityData
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

func (outputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops) GetEntityData() *types.CommonEntityData {
    outputTotalDrops.EntityData.YFilter = outputTotalDrops.YFilter
    outputTotalDrops.EntityData.YangName = "output-total-drops"
    outputTotalDrops.EntityData.BundleName = "cisco_ios_xr"
    outputTotalDrops.EntityData.ParentYangName = "basic-counter-interface-template"
    outputTotalDrops.EntityData.SegmentPath = "output-total-drops"
    outputTotalDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputTotalDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputTotalDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputTotalDrops.EntityData.Children = make(map[string]types.YChild)
    outputTotalDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    outputTotalDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputTotalDrops.Operator}
    outputTotalDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", outputTotalDrops.Value}
    outputTotalDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputTotalDrops.EndRangeValue}
    outputTotalDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputTotalDrops.Percent}
    outputTotalDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputTotalDrops.RearmType}
    outputTotalDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputTotalDrops.RearmWindow}
    return &(outputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets struct {
    EntityData types.CommonEntityData
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

func (outPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets) GetEntityData() *types.CommonEntityData {
    outPackets.EntityData.YFilter = outPackets.YFilter
    outPackets.EntityData.YangName = "out-packets"
    outPackets.EntityData.BundleName = "cisco_ios_xr"
    outPackets.EntityData.ParentYangName = "basic-counter-interface-template"
    outPackets.EntityData.SegmentPath = "out-packets"
    outPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outPackets.EntityData.Children = make(map[string]types.YChild)
    outPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outPackets.Operator}
    outPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outPackets.Value}
    outPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outPackets.EndRangeValue}
    outPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outPackets.Percent}
    outPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outPackets.RearmType}
    outPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outPackets.RearmWindow}
    return &(outPackets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors struct {
    EntityData types.CommonEntityData
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

func (outputTotalErrors *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors) GetEntityData() *types.CommonEntityData {
    outputTotalErrors.EntityData.YFilter = outputTotalErrors.YFilter
    outputTotalErrors.EntityData.YangName = "output-total-errors"
    outputTotalErrors.EntityData.BundleName = "cisco_ios_xr"
    outputTotalErrors.EntityData.ParentYangName = "basic-counter-interface-template"
    outputTotalErrors.EntityData.SegmentPath = "output-total-errors"
    outputTotalErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputTotalErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputTotalErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputTotalErrors.EntityData.Children = make(map[string]types.YChild)
    outputTotalErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    outputTotalErrors.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputTotalErrors.Operator}
    outputTotalErrors.EntityData.Leafs["value"] = types.YLeaf{"Value", outputTotalErrors.Value}
    outputTotalErrors.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputTotalErrors.EndRangeValue}
    outputTotalErrors.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputTotalErrors.Percent}
    outputTotalErrors.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputTotalErrors.RearmType}
    outputTotalErrors.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputTotalErrors.RearmWindow}
    return &(outputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets struct {
    EntityData types.CommonEntityData
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

func (inPackets *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets) GetEntityData() *types.CommonEntityData {
    inPackets.EntityData.YFilter = inPackets.YFilter
    inPackets.EntityData.YangName = "in-packets"
    inPackets.EntityData.BundleName = "cisco_ios_xr"
    inPackets.EntityData.ParentYangName = "basic-counter-interface-template"
    inPackets.EntityData.SegmentPath = "in-packets"
    inPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inPackets.EntityData.Children = make(map[string]types.YChild)
    inPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inPackets.Operator}
    inPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inPackets.Value}
    inPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inPackets.EndRangeValue}
    inPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inPackets.Percent}
    inPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inPackets.RearmType}
    inPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inPackets.RearmWindow}
    return &(inPackets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops struct {
    EntityData types.CommonEntityData
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

func (inputQueueDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops) GetEntityData() *types.CommonEntityData {
    inputQueueDrops.EntityData.YFilter = inputQueueDrops.YFilter
    inputQueueDrops.EntityData.YangName = "input-queue-drops"
    inputQueueDrops.EntityData.BundleName = "cisco_ios_xr"
    inputQueueDrops.EntityData.ParentYangName = "basic-counter-interface-template"
    inputQueueDrops.EntityData.SegmentPath = "input-queue-drops"
    inputQueueDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputQueueDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputQueueDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputQueueDrops.EntityData.Children = make(map[string]types.YChild)
    inputQueueDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    inputQueueDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputQueueDrops.Operator}
    inputQueueDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", inputQueueDrops.Value}
    inputQueueDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputQueueDrops.EndRangeValue}
    inputQueueDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputQueueDrops.Percent}
    inputQueueDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputQueueDrops.RearmType}
    inputQueueDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputQueueDrops.RearmWindow}
    return &(inputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops struct {
    EntityData types.CommonEntityData
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

func (inputTotalDrops *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops) GetEntityData() *types.CommonEntityData {
    inputTotalDrops.EntityData.YFilter = inputTotalDrops.YFilter
    inputTotalDrops.EntityData.YangName = "input-total-drops"
    inputTotalDrops.EntityData.BundleName = "cisco_ios_xr"
    inputTotalDrops.EntityData.ParentYangName = "basic-counter-interface-template"
    inputTotalDrops.EntityData.SegmentPath = "input-total-drops"
    inputTotalDrops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputTotalDrops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputTotalDrops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputTotalDrops.EntityData.Children = make(map[string]types.YChild)
    inputTotalDrops.EntityData.Leafs = make(map[string]types.YLeaf)
    inputTotalDrops.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputTotalDrops.Operator}
    inputTotalDrops.EntityData.Leafs["value"] = types.YLeaf{"Value", inputTotalDrops.Value}
    inputTotalDrops.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputTotalDrops.EndRangeValue}
    inputTotalDrops.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputTotalDrops.Percent}
    inputTotalDrops.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputTotalDrops.RearmType}
    inputTotalDrops.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputTotalDrops.RearmWindow}
    return &(inputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_Bgp
// BGP threshold configuration
type PerfMgmt_Threshold_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP threshold templates.
    BgpTemplates PerfMgmt_Threshold_Bgp_BgpTemplates
}

func (bgp *PerfMgmt_Threshold_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "threshold"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Children["bgp-templates"] = types.YChild{"BgpTemplates", &bgp.BgpTemplates}
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgp.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates
// BGP threshold templates
type PerfMgmt_Threshold_Bgp_BgpTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate.
    BgpTemplate []PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
}

func (bgpTemplates *PerfMgmt_Threshold_Bgp_BgpTemplates) GetEntityData() *types.CommonEntityData {
    bgpTemplates.EntityData.YFilter = bgpTemplates.YFilter
    bgpTemplates.EntityData.YangName = "bgp-templates"
    bgpTemplates.EntityData.BundleName = "cisco_ios_xr"
    bgpTemplates.EntityData.ParentYangName = "bgp"
    bgpTemplates.EntityData.SegmentPath = "bgp-templates"
    bgpTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpTemplates.EntityData.Children = make(map[string]types.YChild)
    bgpTemplates.EntityData.Children["bgp-template"] = types.YChild{"BgpTemplate", nil}
    for i := range bgpTemplates.BgpTemplate {
        bgpTemplates.EntityData.Children[types.GetSegmentPath(&bgpTemplates.BgpTemplate[i])] = types.YChild{"BgpTemplate", &bgpTemplates.BgpTemplate[i]}
    }
    bgpTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgpTemplates.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
// BGP threshold template instance
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetEntityData() *types.CommonEntityData {
    bgpTemplate.EntityData.YFilter = bgpTemplate.YFilter
    bgpTemplate.EntityData.YangName = "bgp-template"
    bgpTemplate.EntityData.BundleName = "cisco_ios_xr"
    bgpTemplate.EntityData.ParentYangName = "bgp-templates"
    bgpTemplate.EntityData.SegmentPath = "bgp-template" + "[template-name='" + fmt.Sprintf("%v", bgpTemplate.TemplateName) + "']"
    bgpTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpTemplate.EntityData.Children = make(map[string]types.YChild)
    bgpTemplate.EntityData.Children["output-update-messages"] = types.YChild{"OutputUpdateMessages", &bgpTemplate.OutputUpdateMessages}
    bgpTemplate.EntityData.Children["errors-received"] = types.YChild{"ErrorsReceived", &bgpTemplate.ErrorsReceived}
    bgpTemplate.EntityData.Children["conn-established"] = types.YChild{"ConnEstablished", &bgpTemplate.ConnEstablished}
    bgpTemplate.EntityData.Children["output-messages"] = types.YChild{"OutputMessages", &bgpTemplate.OutputMessages}
    bgpTemplate.EntityData.Children["conn-dropped"] = types.YChild{"ConnDropped", &bgpTemplate.ConnDropped}
    bgpTemplate.EntityData.Children["input-update-messages"] = types.YChild{"InputUpdateMessages", &bgpTemplate.InputUpdateMessages}
    bgpTemplate.EntityData.Children["errors-sent"] = types.YChild{"ErrorsSent", &bgpTemplate.ErrorsSent}
    bgpTemplate.EntityData.Children["input-messages"] = types.YChild{"InputMessages", &bgpTemplate.InputMessages}
    bgpTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    bgpTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", bgpTemplate.TemplateName}
    bgpTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", bgpTemplate.SampleInterval}
    return &(bgpTemplate.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages
// Number of update messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages struct {
    EntityData types.CommonEntityData
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

func (outputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages) GetEntityData() *types.CommonEntityData {
    outputUpdateMessages.EntityData.YFilter = outputUpdateMessages.YFilter
    outputUpdateMessages.EntityData.YangName = "output-update-messages"
    outputUpdateMessages.EntityData.BundleName = "cisco_ios_xr"
    outputUpdateMessages.EntityData.ParentYangName = "bgp-template"
    outputUpdateMessages.EntityData.SegmentPath = "output-update-messages"
    outputUpdateMessages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputUpdateMessages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputUpdateMessages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputUpdateMessages.EntityData.Children = make(map[string]types.YChild)
    outputUpdateMessages.EntityData.Leafs = make(map[string]types.YLeaf)
    outputUpdateMessages.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputUpdateMessages.Operator}
    outputUpdateMessages.EntityData.Leafs["value"] = types.YLeaf{"Value", outputUpdateMessages.Value}
    outputUpdateMessages.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputUpdateMessages.EndRangeValue}
    outputUpdateMessages.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputUpdateMessages.Percent}
    outputUpdateMessages.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputUpdateMessages.RearmType}
    outputUpdateMessages.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputUpdateMessages.RearmWindow}
    return &(outputUpdateMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived
// Number of error notifications received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived struct {
    EntityData types.CommonEntityData
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

func (errorsReceived *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived) GetEntityData() *types.CommonEntityData {
    errorsReceived.EntityData.YFilter = errorsReceived.YFilter
    errorsReceived.EntityData.YangName = "errors-received"
    errorsReceived.EntityData.BundleName = "cisco_ios_xr"
    errorsReceived.EntityData.ParentYangName = "bgp-template"
    errorsReceived.EntityData.SegmentPath = "errors-received"
    errorsReceived.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorsReceived.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorsReceived.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorsReceived.EntityData.Children = make(map[string]types.YChild)
    errorsReceived.EntityData.Leafs = make(map[string]types.YLeaf)
    errorsReceived.EntityData.Leafs["operator"] = types.YLeaf{"Operator", errorsReceived.Operator}
    errorsReceived.EntityData.Leafs["value"] = types.YLeaf{"Value", errorsReceived.Value}
    errorsReceived.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", errorsReceived.EndRangeValue}
    errorsReceived.EntityData.Leafs["percent"] = types.YLeaf{"Percent", errorsReceived.Percent}
    errorsReceived.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", errorsReceived.RearmType}
    errorsReceived.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", errorsReceived.RearmWindow}
    return &(errorsReceived.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished
// Number of times the connection was
// established
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished struct {
    EntityData types.CommonEntityData
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

func (connEstablished *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished) GetEntityData() *types.CommonEntityData {
    connEstablished.EntityData.YFilter = connEstablished.YFilter
    connEstablished.EntityData.YangName = "conn-established"
    connEstablished.EntityData.BundleName = "cisco_ios_xr"
    connEstablished.EntityData.ParentYangName = "bgp-template"
    connEstablished.EntityData.SegmentPath = "conn-established"
    connEstablished.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connEstablished.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connEstablished.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connEstablished.EntityData.Children = make(map[string]types.YChild)
    connEstablished.EntityData.Leafs = make(map[string]types.YLeaf)
    connEstablished.EntityData.Leafs["operator"] = types.YLeaf{"Operator", connEstablished.Operator}
    connEstablished.EntityData.Leafs["value"] = types.YLeaf{"Value", connEstablished.Value}
    connEstablished.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", connEstablished.EndRangeValue}
    connEstablished.EntityData.Leafs["percent"] = types.YLeaf{"Percent", connEstablished.Percent}
    connEstablished.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", connEstablished.RearmType}
    connEstablished.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", connEstablished.RearmWindow}
    return &(connEstablished.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages
// Number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages struct {
    EntityData types.CommonEntityData
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

func (outputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages) GetEntityData() *types.CommonEntityData {
    outputMessages.EntityData.YFilter = outputMessages.YFilter
    outputMessages.EntityData.YangName = "output-messages"
    outputMessages.EntityData.BundleName = "cisco_ios_xr"
    outputMessages.EntityData.ParentYangName = "bgp-template"
    outputMessages.EntityData.SegmentPath = "output-messages"
    outputMessages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputMessages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputMessages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputMessages.EntityData.Children = make(map[string]types.YChild)
    outputMessages.EntityData.Leafs = make(map[string]types.YLeaf)
    outputMessages.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputMessages.Operator}
    outputMessages.EntityData.Leafs["value"] = types.YLeaf{"Value", outputMessages.Value}
    outputMessages.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputMessages.EndRangeValue}
    outputMessages.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputMessages.Percent}
    outputMessages.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputMessages.RearmType}
    outputMessages.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputMessages.RearmWindow}
    return &(outputMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped
// Number of times the connection was dropped
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped struct {
    EntityData types.CommonEntityData
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

func (connDropped *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped) GetEntityData() *types.CommonEntityData {
    connDropped.EntityData.YFilter = connDropped.YFilter
    connDropped.EntityData.YangName = "conn-dropped"
    connDropped.EntityData.BundleName = "cisco_ios_xr"
    connDropped.EntityData.ParentYangName = "bgp-template"
    connDropped.EntityData.SegmentPath = "conn-dropped"
    connDropped.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connDropped.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connDropped.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connDropped.EntityData.Children = make(map[string]types.YChild)
    connDropped.EntityData.Leafs = make(map[string]types.YLeaf)
    connDropped.EntityData.Leafs["operator"] = types.YLeaf{"Operator", connDropped.Operator}
    connDropped.EntityData.Leafs["value"] = types.YLeaf{"Value", connDropped.Value}
    connDropped.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", connDropped.EndRangeValue}
    connDropped.EntityData.Leafs["percent"] = types.YLeaf{"Percent", connDropped.Percent}
    connDropped.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", connDropped.RearmType}
    connDropped.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", connDropped.RearmWindow}
    return &(connDropped.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages
// Number of update messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages struct {
    EntityData types.CommonEntityData
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

func (inputUpdateMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages) GetEntityData() *types.CommonEntityData {
    inputUpdateMessages.EntityData.YFilter = inputUpdateMessages.YFilter
    inputUpdateMessages.EntityData.YangName = "input-update-messages"
    inputUpdateMessages.EntityData.BundleName = "cisco_ios_xr"
    inputUpdateMessages.EntityData.ParentYangName = "bgp-template"
    inputUpdateMessages.EntityData.SegmentPath = "input-update-messages"
    inputUpdateMessages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputUpdateMessages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputUpdateMessages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputUpdateMessages.EntityData.Children = make(map[string]types.YChild)
    inputUpdateMessages.EntityData.Leafs = make(map[string]types.YLeaf)
    inputUpdateMessages.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputUpdateMessages.Operator}
    inputUpdateMessages.EntityData.Leafs["value"] = types.YLeaf{"Value", inputUpdateMessages.Value}
    inputUpdateMessages.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputUpdateMessages.EndRangeValue}
    inputUpdateMessages.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputUpdateMessages.Percent}
    inputUpdateMessages.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputUpdateMessages.RearmType}
    inputUpdateMessages.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputUpdateMessages.RearmWindow}
    return &(inputUpdateMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent
// Number of error notifications sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent struct {
    EntityData types.CommonEntityData
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

func (errorsSent *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent) GetEntityData() *types.CommonEntityData {
    errorsSent.EntityData.YFilter = errorsSent.YFilter
    errorsSent.EntityData.YangName = "errors-sent"
    errorsSent.EntityData.BundleName = "cisco_ios_xr"
    errorsSent.EntityData.ParentYangName = "bgp-template"
    errorsSent.EntityData.SegmentPath = "errors-sent"
    errorsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorsSent.EntityData.Children = make(map[string]types.YChild)
    errorsSent.EntityData.Leafs = make(map[string]types.YLeaf)
    errorsSent.EntityData.Leafs["operator"] = types.YLeaf{"Operator", errorsSent.Operator}
    errorsSent.EntityData.Leafs["value"] = types.YLeaf{"Value", errorsSent.Value}
    errorsSent.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", errorsSent.EndRangeValue}
    errorsSent.EntityData.Leafs["percent"] = types.YLeaf{"Percent", errorsSent.Percent}
    errorsSent.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", errorsSent.RearmType}
    errorsSent.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", errorsSent.RearmWindow}
    return &(errorsSent.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages
// Number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages struct {
    EntityData types.CommonEntityData
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

func (inputMessages *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages) GetEntityData() *types.CommonEntityData {
    inputMessages.EntityData.YFilter = inputMessages.YFilter
    inputMessages.EntityData.YangName = "input-messages"
    inputMessages.EntityData.BundleName = "cisco_ios_xr"
    inputMessages.EntityData.ParentYangName = "bgp-template"
    inputMessages.EntityData.SegmentPath = "input-messages"
    inputMessages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputMessages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputMessages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputMessages.EntityData.Children = make(map[string]types.YChild)
    inputMessages.EntityData.Leafs = make(map[string]types.YLeaf)
    inputMessages.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputMessages.Operator}
    inputMessages.EntityData.Leafs["value"] = types.YLeaf{"Value", inputMessages.Value}
    inputMessages.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputMessages.EndRangeValue}
    inputMessages.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputMessages.Percent}
    inputMessages.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputMessages.RearmType}
    inputMessages.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputMessages.RearmWindow}
    return &(inputMessages.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol
// OSPF v2 Protocol threshold configuration
type PerfMgmt_Threshold_Ospfv2Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold templates.
    Ospfv2ProtocolTemplates PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates
}

func (ospfv2Protocol *PerfMgmt_Threshold_Ospfv2Protocol) GetEntityData() *types.CommonEntityData {
    ospfv2Protocol.EntityData.YFilter = ospfv2Protocol.YFilter
    ospfv2Protocol.EntityData.YangName = "ospfv2-protocol"
    ospfv2Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv2Protocol.EntityData.ParentYangName = "threshold"
    ospfv2Protocol.EntityData.SegmentPath = "ospfv2-protocol"
    ospfv2Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv2Protocol.EntityData.Children["ospfv2-protocol-templates"] = types.YChild{"Ospfv2ProtocolTemplates", &ospfv2Protocol.Ospfv2ProtocolTemplates}
    ospfv2Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate.
    Ospfv2ProtocolTemplate []PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
}

func (ospfv2ProtocolTemplates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolTemplates.EntityData.YFilter = ospfv2ProtocolTemplates.YFilter
    ospfv2ProtocolTemplates.EntityData.YangName = "ospfv2-protocol-templates"
    ospfv2ProtocolTemplates.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolTemplates.EntityData.ParentYangName = "ospfv2-protocol"
    ospfv2ProtocolTemplates.EntityData.SegmentPath = "ospfv2-protocol-templates"
    ospfv2ProtocolTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolTemplates.EntityData.Children = make(map[string]types.YChild)
    ospfv2ProtocolTemplates.EntityData.Children["ospfv2-protocol-template"] = types.YChild{"Ospfv2ProtocolTemplate", nil}
    for i := range ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate {
        ospfv2ProtocolTemplates.EntityData.Children[types.GetSegmentPath(&ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i])] = types.YChild{"Ospfv2ProtocolTemplate", &ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i]}
    }
    ospfv2ProtocolTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv2ProtocolTemplates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolTemplate.EntityData.YFilter = ospfv2ProtocolTemplate.YFilter
    ospfv2ProtocolTemplate.EntityData.YangName = "ospfv2-protocol-template"
    ospfv2ProtocolTemplate.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolTemplate.EntityData.ParentYangName = "ospfv2-protocol-templates"
    ospfv2ProtocolTemplate.EntityData.SegmentPath = "ospfv2-protocol-template" + "[template-name='" + fmt.Sprintf("%v", ospfv2ProtocolTemplate.TemplateName) + "']"
    ospfv2ProtocolTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolTemplate.EntityData.Children = make(map[string]types.YChild)
    ospfv2ProtocolTemplate.EntityData.Children["checksum-errors"] = types.YChild{"ChecksumErrors", &ospfv2ProtocolTemplate.ChecksumErrors}
    ospfv2ProtocolTemplate.EntityData.Children["input-lsa-acks-lsa"] = types.YChild{"InputLsaAcksLsa", &ospfv2ProtocolTemplate.InputLsaAcksLsa}
    ospfv2ProtocolTemplate.EntityData.Children["output-db-ds-lsa"] = types.YChild{"OutputDbDsLsa", &ospfv2ProtocolTemplate.OutputDbDsLsa}
    ospfv2ProtocolTemplate.EntityData.Children["input-db-ds-lsa"] = types.YChild{"InputDbDsLsa", &ospfv2ProtocolTemplate.InputDbDsLsa}
    ospfv2ProtocolTemplate.EntityData.Children["input-lsa-updates"] = types.YChild{"InputLsaUpdates", &ospfv2ProtocolTemplate.InputLsaUpdates}
    ospfv2ProtocolTemplate.EntityData.Children["output-db-ds"] = types.YChild{"OutputDbDs", &ospfv2ProtocolTemplate.OutputDbDs}
    ospfv2ProtocolTemplate.EntityData.Children["output-lsa-updates-lsa"] = types.YChild{"OutputLsaUpdatesLsa", &ospfv2ProtocolTemplate.OutputLsaUpdatesLsa}
    ospfv2ProtocolTemplate.EntityData.Children["input-db-ds"] = types.YChild{"InputDbDs", &ospfv2ProtocolTemplate.InputDbDs}
    ospfv2ProtocolTemplate.EntityData.Children["input-lsa-updates-lsa"] = types.YChild{"InputLsaUpdatesLsa", &ospfv2ProtocolTemplate.InputLsaUpdatesLsa}
    ospfv2ProtocolTemplate.EntityData.Children["output-packets"] = types.YChild{"OutputPackets", &ospfv2ProtocolTemplate.OutputPackets}
    ospfv2ProtocolTemplate.EntityData.Children["input-packets"] = types.YChild{"InputPackets", &ospfv2ProtocolTemplate.InputPackets}
    ospfv2ProtocolTemplate.EntityData.Children["output-hello-packets"] = types.YChild{"OutputHelloPackets", &ospfv2ProtocolTemplate.OutputHelloPackets}
    ospfv2ProtocolTemplate.EntityData.Children["input-hello-packets"] = types.YChild{"InputHelloPackets", &ospfv2ProtocolTemplate.InputHelloPackets}
    ospfv2ProtocolTemplate.EntityData.Children["output-ls-requests"] = types.YChild{"OutputLsRequests", &ospfv2ProtocolTemplate.OutputLsRequests}
    ospfv2ProtocolTemplate.EntityData.Children["output-lsa-acks-lsa"] = types.YChild{"OutputLsaAcksLsa", &ospfv2ProtocolTemplate.OutputLsaAcksLsa}
    ospfv2ProtocolTemplate.EntityData.Children["output-lsa-acks"] = types.YChild{"OutputLsaAcks", &ospfv2ProtocolTemplate.OutputLsaAcks}
    ospfv2ProtocolTemplate.EntityData.Children["input-lsa-acks"] = types.YChild{"InputLsaAcks", &ospfv2ProtocolTemplate.InputLsaAcks}
    ospfv2ProtocolTemplate.EntityData.Children["output-lsa-updates"] = types.YChild{"OutputLsaUpdates", &ospfv2ProtocolTemplate.OutputLsaUpdates}
    ospfv2ProtocolTemplate.EntityData.Children["output-ls-requests-lsa"] = types.YChild{"OutputLsRequestsLsa", &ospfv2ProtocolTemplate.OutputLsRequestsLsa}
    ospfv2ProtocolTemplate.EntityData.Children["input-ls-requests-lsa"] = types.YChild{"InputLsRequestsLsa", &ospfv2ProtocolTemplate.InputLsRequestsLsa}
    ospfv2ProtocolTemplate.EntityData.Children["input-ls-requests"] = types.YChild{"InputLsRequests", &ospfv2ProtocolTemplate.InputLsRequests}
    ospfv2ProtocolTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv2ProtocolTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv2ProtocolTemplate.TemplateName}
    ospfv2ProtocolTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", ospfv2ProtocolTemplate.SampleInterval}
    return &(ospfv2ProtocolTemplate.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors
// Number of packets received with checksum
// errors
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors struct {
    EntityData types.CommonEntityData
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

func (checksumErrors *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors) GetEntityData() *types.CommonEntityData {
    checksumErrors.EntityData.YFilter = checksumErrors.YFilter
    checksumErrors.EntityData.YangName = "checksum-errors"
    checksumErrors.EntityData.BundleName = "cisco_ios_xr"
    checksumErrors.EntityData.ParentYangName = "ospfv2-protocol-template"
    checksumErrors.EntityData.SegmentPath = "checksum-errors"
    checksumErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    checksumErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    checksumErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    checksumErrors.EntityData.Children = make(map[string]types.YChild)
    checksumErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    checksumErrors.EntityData.Leafs["operator"] = types.YLeaf{"Operator", checksumErrors.Operator}
    checksumErrors.EntityData.Leafs["value"] = types.YLeaf{"Value", checksumErrors.Value}
    checksumErrors.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", checksumErrors.EndRangeValue}
    checksumErrors.EntityData.Leafs["percent"] = types.YLeaf{"Percent", checksumErrors.Percent}
    checksumErrors.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", checksumErrors.RearmType}
    checksumErrors.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", checksumErrors.RearmWindow}
    return &(checksumErrors.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa) GetEntityData() *types.CommonEntityData {
    inputLsaAcksLsa.EntityData.YFilter = inputLsaAcksLsa.YFilter
    inputLsaAcksLsa.EntityData.YangName = "input-lsa-acks-lsa"
    inputLsaAcksLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsaAcksLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsaAcksLsa.EntityData.SegmentPath = "input-lsa-acks-lsa"
    inputLsaAcksLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaAcksLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaAcksLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaAcksLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsaAcksLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaAcksLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaAcksLsa.Operator}
    inputLsaAcksLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaAcksLsa.Value}
    inputLsaAcksLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaAcksLsa.EndRangeValue}
    inputLsaAcksLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaAcksLsa.Percent}
    inputLsaAcksLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaAcksLsa.RearmType}
    inputLsaAcksLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaAcksLsa.RearmWindow}
    return &(inputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa struct {
    EntityData types.CommonEntityData
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

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa) GetEntityData() *types.CommonEntityData {
    outputDbDsLsa.EntityData.YFilter = outputDbDsLsa.YFilter
    outputDbDsLsa.EntityData.YangName = "output-db-ds-lsa"
    outputDbDsLsa.EntityData.BundleName = "cisco_ios_xr"
    outputDbDsLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputDbDsLsa.EntityData.SegmentPath = "output-db-ds-lsa"
    outputDbDsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputDbDsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputDbDsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputDbDsLsa.EntityData.Children = make(map[string]types.YChild)
    outputDbDsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputDbDsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputDbDsLsa.Operator}
    outputDbDsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputDbDsLsa.Value}
    outputDbDsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputDbDsLsa.EndRangeValue}
    outputDbDsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputDbDsLsa.Percent}
    outputDbDsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputDbDsLsa.RearmType}
    outputDbDsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputDbDsLsa.RearmWindow}
    return &(outputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa struct {
    EntityData types.CommonEntityData
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

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa) GetEntityData() *types.CommonEntityData {
    inputDbDsLsa.EntityData.YFilter = inputDbDsLsa.YFilter
    inputDbDsLsa.EntityData.YangName = "input-db-ds-lsa"
    inputDbDsLsa.EntityData.BundleName = "cisco_ios_xr"
    inputDbDsLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputDbDsLsa.EntityData.SegmentPath = "input-db-ds-lsa"
    inputDbDsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputDbDsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputDbDsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputDbDsLsa.EntityData.Children = make(map[string]types.YChild)
    inputDbDsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputDbDsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputDbDsLsa.Operator}
    inputDbDsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputDbDsLsa.Value}
    inputDbDsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputDbDsLsa.EndRangeValue}
    inputDbDsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputDbDsLsa.Percent}
    inputDbDsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputDbDsLsa.RearmType}
    inputDbDsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputDbDsLsa.RearmWindow}
    return &(inputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates struct {
    EntityData types.CommonEntityData
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

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates) GetEntityData() *types.CommonEntityData {
    inputLsaUpdates.EntityData.YFilter = inputLsaUpdates.YFilter
    inputLsaUpdates.EntityData.YangName = "input-lsa-updates"
    inputLsaUpdates.EntityData.BundleName = "cisco_ios_xr"
    inputLsaUpdates.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsaUpdates.EntityData.SegmentPath = "input-lsa-updates"
    inputLsaUpdates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaUpdates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaUpdates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaUpdates.EntityData.Children = make(map[string]types.YChild)
    inputLsaUpdates.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaUpdates.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaUpdates.Operator}
    inputLsaUpdates.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaUpdates.Value}
    inputLsaUpdates.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaUpdates.EndRangeValue}
    inputLsaUpdates.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaUpdates.Percent}
    inputLsaUpdates.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaUpdates.RearmType}
    inputLsaUpdates.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaUpdates.RearmWindow}
    return &(inputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs struct {
    EntityData types.CommonEntityData
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

func (outputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs) GetEntityData() *types.CommonEntityData {
    outputDbDs.EntityData.YFilter = outputDbDs.YFilter
    outputDbDs.EntityData.YangName = "output-db-ds"
    outputDbDs.EntityData.BundleName = "cisco_ios_xr"
    outputDbDs.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputDbDs.EntityData.SegmentPath = "output-db-ds"
    outputDbDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputDbDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputDbDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputDbDs.EntityData.Children = make(map[string]types.YChild)
    outputDbDs.EntityData.Leafs = make(map[string]types.YLeaf)
    outputDbDs.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputDbDs.Operator}
    outputDbDs.EntityData.Leafs["value"] = types.YLeaf{"Value", outputDbDs.Value}
    outputDbDs.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputDbDs.EndRangeValue}
    outputDbDs.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputDbDs.Percent}
    outputDbDs.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputDbDs.RearmType}
    outputDbDs.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputDbDs.RearmWindow}
    return &(outputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa) GetEntityData() *types.CommonEntityData {
    outputLsaUpdatesLsa.EntityData.YFilter = outputLsaUpdatesLsa.YFilter
    outputLsaUpdatesLsa.EntityData.YangName = "output-lsa-updates-lsa"
    outputLsaUpdatesLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsaUpdatesLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsaUpdatesLsa.EntityData.SegmentPath = "output-lsa-updates-lsa"
    outputLsaUpdatesLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaUpdatesLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaUpdatesLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaUpdatesLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsaUpdatesLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaUpdatesLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaUpdatesLsa.Operator}
    outputLsaUpdatesLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaUpdatesLsa.Value}
    outputLsaUpdatesLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaUpdatesLsa.EndRangeValue}
    outputLsaUpdatesLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaUpdatesLsa.Percent}
    outputLsaUpdatesLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaUpdatesLsa.RearmType}
    outputLsaUpdatesLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaUpdatesLsa.RearmWindow}
    return &(outputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs struct {
    EntityData types.CommonEntityData
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

func (inputDbDs *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs) GetEntityData() *types.CommonEntityData {
    inputDbDs.EntityData.YFilter = inputDbDs.YFilter
    inputDbDs.EntityData.YangName = "input-db-ds"
    inputDbDs.EntityData.BundleName = "cisco_ios_xr"
    inputDbDs.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputDbDs.EntityData.SegmentPath = "input-db-ds"
    inputDbDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputDbDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputDbDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputDbDs.EntityData.Children = make(map[string]types.YChild)
    inputDbDs.EntityData.Leafs = make(map[string]types.YLeaf)
    inputDbDs.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputDbDs.Operator}
    inputDbDs.EntityData.Leafs["value"] = types.YLeaf{"Value", inputDbDs.Value}
    inputDbDs.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputDbDs.EndRangeValue}
    inputDbDs.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputDbDs.Percent}
    inputDbDs.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputDbDs.RearmType}
    inputDbDs.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputDbDs.RearmWindow}
    return &(inputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa) GetEntityData() *types.CommonEntityData {
    inputLsaUpdatesLsa.EntityData.YFilter = inputLsaUpdatesLsa.YFilter
    inputLsaUpdatesLsa.EntityData.YangName = "input-lsa-updates-lsa"
    inputLsaUpdatesLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsaUpdatesLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsaUpdatesLsa.EntityData.SegmentPath = "input-lsa-updates-lsa"
    inputLsaUpdatesLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaUpdatesLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaUpdatesLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaUpdatesLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsaUpdatesLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaUpdatesLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaUpdatesLsa.Operator}
    inputLsaUpdatesLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaUpdatesLsa.Value}
    inputLsaUpdatesLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaUpdatesLsa.EndRangeValue}
    inputLsaUpdatesLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaUpdatesLsa.Percent}
    inputLsaUpdatesLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaUpdatesLsa.RearmType}
    inputLsaUpdatesLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaUpdatesLsa.RearmWindow}
    return &(inputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets struct {
    EntityData types.CommonEntityData
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

func (outputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets) GetEntityData() *types.CommonEntityData {
    outputPackets.EntityData.YFilter = outputPackets.YFilter
    outputPackets.EntityData.YangName = "output-packets"
    outputPackets.EntityData.BundleName = "cisco_ios_xr"
    outputPackets.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputPackets.EntityData.SegmentPath = "output-packets"
    outputPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPackets.EntityData.Children = make(map[string]types.YChild)
    outputPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputPackets.Operator}
    outputPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outputPackets.Value}
    outputPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputPackets.EndRangeValue}
    outputPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputPackets.Percent}
    outputPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputPackets.RearmType}
    outputPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputPackets.RearmWindow}
    return &(outputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets struct {
    EntityData types.CommonEntityData
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

func (inputPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets) GetEntityData() *types.CommonEntityData {
    inputPackets.EntityData.YFilter = inputPackets.YFilter
    inputPackets.EntityData.YangName = "input-packets"
    inputPackets.EntityData.BundleName = "cisco_ios_xr"
    inputPackets.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputPackets.EntityData.SegmentPath = "input-packets"
    inputPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputPackets.EntityData.Children = make(map[string]types.YChild)
    inputPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inputPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputPackets.Operator}
    inputPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inputPackets.Value}
    inputPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputPackets.EndRangeValue}
    inputPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputPackets.Percent}
    inputPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputPackets.RearmType}
    inputPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputPackets.RearmWindow}
    return &(inputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets struct {
    EntityData types.CommonEntityData
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

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets) GetEntityData() *types.CommonEntityData {
    outputHelloPackets.EntityData.YFilter = outputHelloPackets.YFilter
    outputHelloPackets.EntityData.YangName = "output-hello-packets"
    outputHelloPackets.EntityData.BundleName = "cisco_ios_xr"
    outputHelloPackets.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputHelloPackets.EntityData.SegmentPath = "output-hello-packets"
    outputHelloPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputHelloPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputHelloPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputHelloPackets.EntityData.Children = make(map[string]types.YChild)
    outputHelloPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outputHelloPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputHelloPackets.Operator}
    outputHelloPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outputHelloPackets.Value}
    outputHelloPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputHelloPackets.EndRangeValue}
    outputHelloPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputHelloPackets.Percent}
    outputHelloPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputHelloPackets.RearmType}
    outputHelloPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputHelloPackets.RearmWindow}
    return &(outputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets struct {
    EntityData types.CommonEntityData
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

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets) GetEntityData() *types.CommonEntityData {
    inputHelloPackets.EntityData.YFilter = inputHelloPackets.YFilter
    inputHelloPackets.EntityData.YangName = "input-hello-packets"
    inputHelloPackets.EntityData.BundleName = "cisco_ios_xr"
    inputHelloPackets.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputHelloPackets.EntityData.SegmentPath = "input-hello-packets"
    inputHelloPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputHelloPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputHelloPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputHelloPackets.EntityData.Children = make(map[string]types.YChild)
    inputHelloPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inputHelloPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputHelloPackets.Operator}
    inputHelloPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inputHelloPackets.Value}
    inputHelloPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputHelloPackets.EndRangeValue}
    inputHelloPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputHelloPackets.Percent}
    inputHelloPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputHelloPackets.RearmType}
    inputHelloPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputHelloPackets.RearmWindow}
    return &(inputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests struct {
    EntityData types.CommonEntityData
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

func (outputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests) GetEntityData() *types.CommonEntityData {
    outputLsRequests.EntityData.YFilter = outputLsRequests.YFilter
    outputLsRequests.EntityData.YangName = "output-ls-requests"
    outputLsRequests.EntityData.BundleName = "cisco_ios_xr"
    outputLsRequests.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsRequests.EntityData.SegmentPath = "output-ls-requests"
    outputLsRequests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsRequests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsRequests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsRequests.EntityData.Children = make(map[string]types.YChild)
    outputLsRequests.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsRequests.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsRequests.Operator}
    outputLsRequests.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsRequests.Value}
    outputLsRequests.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsRequests.EndRangeValue}
    outputLsRequests.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsRequests.Percent}
    outputLsRequests.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsRequests.RearmType}
    outputLsRequests.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsRequests.RearmWindow}
    return &(outputLsRequests.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa) GetEntityData() *types.CommonEntityData {
    outputLsaAcksLsa.EntityData.YFilter = outputLsaAcksLsa.YFilter
    outputLsaAcksLsa.EntityData.YangName = "output-lsa-acks-lsa"
    outputLsaAcksLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsaAcksLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsaAcksLsa.EntityData.SegmentPath = "output-lsa-acks-lsa"
    outputLsaAcksLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaAcksLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaAcksLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaAcksLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsaAcksLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaAcksLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaAcksLsa.Operator}
    outputLsaAcksLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaAcksLsa.Value}
    outputLsaAcksLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaAcksLsa.EndRangeValue}
    outputLsaAcksLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaAcksLsa.Percent}
    outputLsaAcksLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaAcksLsa.RearmType}
    outputLsaAcksLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaAcksLsa.RearmWindow}
    return &(outputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks struct {
    EntityData types.CommonEntityData
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

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks) GetEntityData() *types.CommonEntityData {
    outputLsaAcks.EntityData.YFilter = outputLsaAcks.YFilter
    outputLsaAcks.EntityData.YangName = "output-lsa-acks"
    outputLsaAcks.EntityData.BundleName = "cisco_ios_xr"
    outputLsaAcks.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsaAcks.EntityData.SegmentPath = "output-lsa-acks"
    outputLsaAcks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaAcks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaAcks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaAcks.EntityData.Children = make(map[string]types.YChild)
    outputLsaAcks.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaAcks.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaAcks.Operator}
    outputLsaAcks.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaAcks.Value}
    outputLsaAcks.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaAcks.EndRangeValue}
    outputLsaAcks.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaAcks.Percent}
    outputLsaAcks.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaAcks.RearmType}
    outputLsaAcks.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaAcks.RearmWindow}
    return &(outputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks struct {
    EntityData types.CommonEntityData
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

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks) GetEntityData() *types.CommonEntityData {
    inputLsaAcks.EntityData.YFilter = inputLsaAcks.YFilter
    inputLsaAcks.EntityData.YangName = "input-lsa-acks"
    inputLsaAcks.EntityData.BundleName = "cisco_ios_xr"
    inputLsaAcks.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsaAcks.EntityData.SegmentPath = "input-lsa-acks"
    inputLsaAcks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaAcks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaAcks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaAcks.EntityData.Children = make(map[string]types.YChild)
    inputLsaAcks.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaAcks.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaAcks.Operator}
    inputLsaAcks.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaAcks.Value}
    inputLsaAcks.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaAcks.EndRangeValue}
    inputLsaAcks.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaAcks.Percent}
    inputLsaAcks.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaAcks.RearmType}
    inputLsaAcks.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaAcks.RearmWindow}
    return &(inputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates struct {
    EntityData types.CommonEntityData
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

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates) GetEntityData() *types.CommonEntityData {
    outputLsaUpdates.EntityData.YFilter = outputLsaUpdates.YFilter
    outputLsaUpdates.EntityData.YangName = "output-lsa-updates"
    outputLsaUpdates.EntityData.BundleName = "cisco_ios_xr"
    outputLsaUpdates.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsaUpdates.EntityData.SegmentPath = "output-lsa-updates"
    outputLsaUpdates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaUpdates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaUpdates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaUpdates.EntityData.Children = make(map[string]types.YChild)
    outputLsaUpdates.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaUpdates.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaUpdates.Operator}
    outputLsaUpdates.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaUpdates.Value}
    outputLsaUpdates.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaUpdates.EndRangeValue}
    outputLsaUpdates.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaUpdates.Percent}
    outputLsaUpdates.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaUpdates.RearmType}
    outputLsaUpdates.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaUpdates.RearmWindow}
    return &(outputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa) GetEntityData() *types.CommonEntityData {
    outputLsRequestsLsa.EntityData.YFilter = outputLsRequestsLsa.YFilter
    outputLsRequestsLsa.EntityData.YangName = "output-ls-requests-lsa"
    outputLsRequestsLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsRequestsLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    outputLsRequestsLsa.EntityData.SegmentPath = "output-ls-requests-lsa"
    outputLsRequestsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsRequestsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsRequestsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsRequestsLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsRequestsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsRequestsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsRequestsLsa.Operator}
    outputLsRequestsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsRequestsLsa.Value}
    outputLsRequestsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsRequestsLsa.EndRangeValue}
    outputLsRequestsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsRequestsLsa.Percent}
    outputLsRequestsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsRequestsLsa.RearmType}
    outputLsRequestsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsRequestsLsa.RearmWindow}
    return &(outputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa) GetEntityData() *types.CommonEntityData {
    inputLsRequestsLsa.EntityData.YFilter = inputLsRequestsLsa.YFilter
    inputLsRequestsLsa.EntityData.YangName = "input-ls-requests-lsa"
    inputLsRequestsLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsRequestsLsa.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsRequestsLsa.EntityData.SegmentPath = "input-ls-requests-lsa"
    inputLsRequestsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsRequestsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsRequestsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsRequestsLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsRequestsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsRequestsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsRequestsLsa.Operator}
    inputLsRequestsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsRequestsLsa.Value}
    inputLsRequestsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsRequestsLsa.EndRangeValue}
    inputLsRequestsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsRequestsLsa.Percent}
    inputLsRequestsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsRequestsLsa.RearmType}
    inputLsRequestsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsRequestsLsa.RearmWindow}
    return &(inputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests struct {
    EntityData types.CommonEntityData
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

func (inputLsRequests *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests) GetEntityData() *types.CommonEntityData {
    inputLsRequests.EntityData.YFilter = inputLsRequests.YFilter
    inputLsRequests.EntityData.YangName = "input-ls-requests"
    inputLsRequests.EntityData.BundleName = "cisco_ios_xr"
    inputLsRequests.EntityData.ParentYangName = "ospfv2-protocol-template"
    inputLsRequests.EntityData.SegmentPath = "input-ls-requests"
    inputLsRequests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsRequests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsRequests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsRequests.EntityData.Children = make(map[string]types.YChild)
    inputLsRequests.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsRequests.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsRequests.Operator}
    inputLsRequests.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsRequests.Value}
    inputLsRequests.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsRequests.EndRangeValue}
    inputLsRequests.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsRequests.Percent}
    inputLsRequests.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsRequests.RearmType}
    inputLsRequests.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsRequests.RearmWindow}
    return &(inputLsRequests.EntityData)
}

// PerfMgmt_Threshold_CpuNode
// Node CPU threshold configuration
type PerfMgmt_Threshold_CpuNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node CPU threshold configuration templates.
    CpuNodeTemplates PerfMgmt_Threshold_CpuNode_CpuNodeTemplates
}

func (cpuNode *PerfMgmt_Threshold_CpuNode) GetEntityData() *types.CommonEntityData {
    cpuNode.EntityData.YFilter = cpuNode.YFilter
    cpuNode.EntityData.YangName = "cpu-node"
    cpuNode.EntityData.BundleName = "cisco_ios_xr"
    cpuNode.EntityData.ParentYangName = "threshold"
    cpuNode.EntityData.SegmentPath = "cpu-node"
    cpuNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNode.EntityData.Children = make(map[string]types.YChild)
    cpuNode.EntityData.Children["cpu-node-templates"] = types.YChild{"CpuNodeTemplates", &cpuNode.CpuNodeTemplates}
    cpuNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuNode.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates
// Node CPU threshold configuration templates
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node CPU threshold configuration template instances. The type is slice of
    // PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate.
    CpuNodeTemplate []PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
}

func (cpuNodeTemplates *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates) GetEntityData() *types.CommonEntityData {
    cpuNodeTemplates.EntityData.YFilter = cpuNodeTemplates.YFilter
    cpuNodeTemplates.EntityData.YangName = "cpu-node-templates"
    cpuNodeTemplates.EntityData.BundleName = "cisco_ios_xr"
    cpuNodeTemplates.EntityData.ParentYangName = "cpu-node"
    cpuNodeTemplates.EntityData.SegmentPath = "cpu-node-templates"
    cpuNodeTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNodeTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNodeTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNodeTemplates.EntityData.Children = make(map[string]types.YChild)
    cpuNodeTemplates.EntityData.Children["cpu-node-template"] = types.YChild{"CpuNodeTemplate", nil}
    for i := range cpuNodeTemplates.CpuNodeTemplate {
        cpuNodeTemplates.EntityData.Children[types.GetSegmentPath(&cpuNodeTemplates.CpuNodeTemplate[i])] = types.YChild{"CpuNodeTemplate", &cpuNodeTemplates.CpuNodeTemplate[i]}
    }
    cpuNodeTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpuNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
// Node CPU threshold configuration template
// instances
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Average %CPU utilization.
    AverageCpuUsed PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed

    // Number of processes.
    NoProcesses PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses
}

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetEntityData() *types.CommonEntityData {
    cpuNodeTemplate.EntityData.YFilter = cpuNodeTemplate.YFilter
    cpuNodeTemplate.EntityData.YangName = "cpu-node-template"
    cpuNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    cpuNodeTemplate.EntityData.ParentYangName = "cpu-node-templates"
    cpuNodeTemplate.EntityData.SegmentPath = "cpu-node-template" + "[template-name='" + fmt.Sprintf("%v", cpuNodeTemplate.TemplateName) + "']"
    cpuNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNodeTemplate.EntityData.Children = make(map[string]types.YChild)
    cpuNodeTemplate.EntityData.Children["average-cpu-used"] = types.YChild{"AverageCpuUsed", &cpuNodeTemplate.AverageCpuUsed}
    cpuNodeTemplate.EntityData.Children["no-processes"] = types.YChild{"NoProcesses", &cpuNodeTemplate.NoProcesses}
    cpuNodeTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    cpuNodeTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", cpuNodeTemplate.TemplateName}
    cpuNodeTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", cpuNodeTemplate.SampleInterval}
    return &(cpuNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed struct {
    EntityData types.CommonEntityData
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

func (averageCpuUsed *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed) GetEntityData() *types.CommonEntityData {
    averageCpuUsed.EntityData.YFilter = averageCpuUsed.YFilter
    averageCpuUsed.EntityData.YangName = "average-cpu-used"
    averageCpuUsed.EntityData.BundleName = "cisco_ios_xr"
    averageCpuUsed.EntityData.ParentYangName = "cpu-node-template"
    averageCpuUsed.EntityData.SegmentPath = "average-cpu-used"
    averageCpuUsed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageCpuUsed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageCpuUsed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageCpuUsed.EntityData.Children = make(map[string]types.YChild)
    averageCpuUsed.EntityData.Leafs = make(map[string]types.YLeaf)
    averageCpuUsed.EntityData.Leafs["operator"] = types.YLeaf{"Operator", averageCpuUsed.Operator}
    averageCpuUsed.EntityData.Leafs["value"] = types.YLeaf{"Value", averageCpuUsed.Value}
    averageCpuUsed.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", averageCpuUsed.EndRangeValue}
    averageCpuUsed.EntityData.Leafs["percent"] = types.YLeaf{"Percent", averageCpuUsed.Percent}
    averageCpuUsed.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", averageCpuUsed.RearmType}
    averageCpuUsed.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", averageCpuUsed.RearmWindow}
    return &(averageCpuUsed.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses
// Number of processes
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses struct {
    EntityData types.CommonEntityData
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

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetEntityData() *types.CommonEntityData {
    noProcesses.EntityData.YFilter = noProcesses.YFilter
    noProcesses.EntityData.YangName = "no-processes"
    noProcesses.EntityData.BundleName = "cisco_ios_xr"
    noProcesses.EntityData.ParentYangName = "cpu-node-template"
    noProcesses.EntityData.SegmentPath = "no-processes"
    noProcesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    noProcesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    noProcesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    noProcesses.EntityData.Children = make(map[string]types.YChild)
    noProcesses.EntityData.Leafs = make(map[string]types.YLeaf)
    noProcesses.EntityData.Leafs["operator"] = types.YLeaf{"Operator", noProcesses.Operator}
    noProcesses.EntityData.Leafs["value"] = types.YLeaf{"Value", noProcesses.Value}
    noProcesses.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", noProcesses.EndRangeValue}
    noProcesses.EntityData.Leafs["percent"] = types.YLeaf{"Percent", noProcesses.Percent}
    noProcesses.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", noProcesses.RearmType}
    noProcesses.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", noProcesses.RearmWindow}
    return &(noProcesses.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface
// Interface Data Rates threshold configuration
type PerfMgmt_Threshold_DataRateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Data Rates threshold templates.
    DataRateInterfaceTemplates PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates
}

func (dataRateInterface *PerfMgmt_Threshold_DataRateInterface) GetEntityData() *types.CommonEntityData {
    dataRateInterface.EntityData.YFilter = dataRateInterface.YFilter
    dataRateInterface.EntityData.YangName = "data-rate-interface"
    dataRateInterface.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterface.EntityData.ParentYangName = "threshold"
    dataRateInterface.EntityData.SegmentPath = "data-rate-interface"
    dataRateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterface.EntityData.Children = make(map[string]types.YChild)
    dataRateInterface.EntityData.Children["data-rate-interface-templates"] = types.YChild{"DataRateInterfaceTemplates", &dataRateInterface.DataRateInterfaceTemplates}
    dataRateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates
// Interface Data Rates threshold templates
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Data Rates threshold template instance. The type is slice of
    // PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate.
    DataRateInterfaceTemplate []PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
}

func (dataRateInterfaceTemplates *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates) GetEntityData() *types.CommonEntityData {
    dataRateInterfaceTemplates.EntityData.YFilter = dataRateInterfaceTemplates.YFilter
    dataRateInterfaceTemplates.EntityData.YangName = "data-rate-interface-templates"
    dataRateInterfaceTemplates.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterfaceTemplates.EntityData.ParentYangName = "data-rate-interface"
    dataRateInterfaceTemplates.EntityData.SegmentPath = "data-rate-interface-templates"
    dataRateInterfaceTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterfaceTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterfaceTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterfaceTemplates.EntityData.Children = make(map[string]types.YChild)
    dataRateInterfaceTemplates.EntityData.Children["data-rate-interface-template"] = types.YChild{"DataRateInterfaceTemplate", nil}
    for i := range dataRateInterfaceTemplates.DataRateInterfaceTemplate {
        dataRateInterfaceTemplates.EntityData.Children[types.GetSegmentPath(&dataRateInterfaceTemplates.DataRateInterfaceTemplate[i])] = types.YChild{"DataRateInterfaceTemplate", &dataRateInterfaceTemplates.DataRateInterfaceTemplate[i]}
    }
    dataRateInterfaceTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dataRateInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
// Interface Data Rates threshold template
// instance
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    dataRateInterfaceTemplate.EntityData.YFilter = dataRateInterfaceTemplate.YFilter
    dataRateInterfaceTemplate.EntityData.YangName = "data-rate-interface-template"
    dataRateInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterfaceTemplate.EntityData.ParentYangName = "data-rate-interface-templates"
    dataRateInterfaceTemplate.EntityData.SegmentPath = "data-rate-interface-template" + "[template-name='" + fmt.Sprintf("%v", dataRateInterfaceTemplate.TemplateName) + "']"
    dataRateInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterfaceTemplate.EntityData.Children = make(map[string]types.YChild)
    dataRateInterfaceTemplate.EntityData.Children["input-data-rate"] = types.YChild{"InputDataRate", &dataRateInterfaceTemplate.InputDataRate}
    dataRateInterfaceTemplate.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &dataRateInterfaceTemplate.Bandwidth}
    dataRateInterfaceTemplate.EntityData.Children["output-packet-rate"] = types.YChild{"OutputPacketRate", &dataRateInterfaceTemplate.OutputPacketRate}
    dataRateInterfaceTemplate.EntityData.Children["input-peak-pkts"] = types.YChild{"InputPeakPkts", &dataRateInterfaceTemplate.InputPeakPkts}
    dataRateInterfaceTemplate.EntityData.Children["output-peak-rate"] = types.YChild{"OutputPeakRate", &dataRateInterfaceTemplate.OutputPeakRate}
    dataRateInterfaceTemplate.EntityData.Children["output-data-rate"] = types.YChild{"OutputDataRate", &dataRateInterfaceTemplate.OutputDataRate}
    dataRateInterfaceTemplate.EntityData.Children["input-packet-rate"] = types.YChild{"InputPacketRate", &dataRateInterfaceTemplate.InputPacketRate}
    dataRateInterfaceTemplate.EntityData.Children["output-peak-pkts"] = types.YChild{"OutputPeakPkts", &dataRateInterfaceTemplate.OutputPeakPkts}
    dataRateInterfaceTemplate.EntityData.Children["input-peak-rate"] = types.YChild{"InputPeakRate", &dataRateInterfaceTemplate.InputPeakRate}
    dataRateInterfaceTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    dataRateInterfaceTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", dataRateInterfaceTemplate.TemplateName}
    dataRateInterfaceTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", dataRateInterfaceTemplate.SampleInterval}
    dataRateInterfaceTemplate.EntityData.Leafs["reg-exp-group"] = types.YLeaf{"RegExpGroup", dataRateInterfaceTemplate.RegExpGroup}
    dataRateInterfaceTemplate.EntityData.Leafs["vrf-group"] = types.YLeaf{"VrfGroup", dataRateInterfaceTemplate.VrfGroup}
    return &(dataRateInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate
// Input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate struct {
    EntityData types.CommonEntityData
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

func (inputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate) GetEntityData() *types.CommonEntityData {
    inputDataRate.EntityData.YFilter = inputDataRate.YFilter
    inputDataRate.EntityData.YangName = "input-data-rate"
    inputDataRate.EntityData.BundleName = "cisco_ios_xr"
    inputDataRate.EntityData.ParentYangName = "data-rate-interface-template"
    inputDataRate.EntityData.SegmentPath = "input-data-rate"
    inputDataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputDataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputDataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputDataRate.EntityData.Children = make(map[string]types.YChild)
    inputDataRate.EntityData.Leafs = make(map[string]types.YLeaf)
    inputDataRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputDataRate.Operator}
    inputDataRate.EntityData.Leafs["value"] = types.YLeaf{"Value", inputDataRate.Value}
    inputDataRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputDataRate.EndRangeValue}
    inputDataRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputDataRate.Percent}
    inputDataRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputDataRate.RearmType}
    inputDataRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputDataRate.RearmWindow}
    return &(inputDataRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth
// Bandwidth in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth struct {
    EntityData types.CommonEntityData
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

func (bandwidth *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "data-rate-interface-template"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["operator"] = types.YLeaf{"Operator", bandwidth.Operator}
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", bandwidth.EndRangeValue}
    bandwidth.EntityData.Leafs["percent"] = types.YLeaf{"Percent", bandwidth.Percent}
    bandwidth.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", bandwidth.RearmType}
    bandwidth.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", bandwidth.RearmWindow}
    return &(bandwidth.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate
// Number of Output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate struct {
    EntityData types.CommonEntityData
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

func (outputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate) GetEntityData() *types.CommonEntityData {
    outputPacketRate.EntityData.YFilter = outputPacketRate.YFilter
    outputPacketRate.EntityData.YangName = "output-packet-rate"
    outputPacketRate.EntityData.BundleName = "cisco_ios_xr"
    outputPacketRate.EntityData.ParentYangName = "data-rate-interface-template"
    outputPacketRate.EntityData.SegmentPath = "output-packet-rate"
    outputPacketRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPacketRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPacketRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPacketRate.EntityData.Children = make(map[string]types.YChild)
    outputPacketRate.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPacketRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputPacketRate.Operator}
    outputPacketRate.EntityData.Leafs["value"] = types.YLeaf{"Value", outputPacketRate.Value}
    outputPacketRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputPacketRate.EndRangeValue}
    outputPacketRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputPacketRate.Percent}
    outputPacketRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputPacketRate.RearmType}
    outputPacketRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputPacketRate.RearmWindow}
    return &(outputPacketRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts
// Maximum number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts struct {
    EntityData types.CommonEntityData
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

func (inputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts) GetEntityData() *types.CommonEntityData {
    inputPeakPkts.EntityData.YFilter = inputPeakPkts.YFilter
    inputPeakPkts.EntityData.YangName = "input-peak-pkts"
    inputPeakPkts.EntityData.BundleName = "cisco_ios_xr"
    inputPeakPkts.EntityData.ParentYangName = "data-rate-interface-template"
    inputPeakPkts.EntityData.SegmentPath = "input-peak-pkts"
    inputPeakPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputPeakPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputPeakPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputPeakPkts.EntityData.Children = make(map[string]types.YChild)
    inputPeakPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    inputPeakPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputPeakPkts.Operator}
    inputPeakPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", inputPeakPkts.Value}
    inputPeakPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputPeakPkts.EndRangeValue}
    inputPeakPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputPeakPkts.Percent}
    inputPeakPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputPeakPkts.RearmType}
    inputPeakPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputPeakPkts.RearmWindow}
    return &(inputPeakPkts.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate
// Peak output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate struct {
    EntityData types.CommonEntityData
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

func (outputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate) GetEntityData() *types.CommonEntityData {
    outputPeakRate.EntityData.YFilter = outputPeakRate.YFilter
    outputPeakRate.EntityData.YangName = "output-peak-rate"
    outputPeakRate.EntityData.BundleName = "cisco_ios_xr"
    outputPeakRate.EntityData.ParentYangName = "data-rate-interface-template"
    outputPeakRate.EntityData.SegmentPath = "output-peak-rate"
    outputPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPeakRate.EntityData.Children = make(map[string]types.YChild)
    outputPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPeakRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputPeakRate.Operator}
    outputPeakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", outputPeakRate.Value}
    outputPeakRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputPeakRate.EndRangeValue}
    outputPeakRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputPeakRate.Percent}
    outputPeakRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputPeakRate.RearmType}
    outputPeakRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputPeakRate.RearmWindow}
    return &(outputPeakRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate
// Output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate struct {
    EntityData types.CommonEntityData
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

func (outputDataRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate) GetEntityData() *types.CommonEntityData {
    outputDataRate.EntityData.YFilter = outputDataRate.YFilter
    outputDataRate.EntityData.YangName = "output-data-rate"
    outputDataRate.EntityData.BundleName = "cisco_ios_xr"
    outputDataRate.EntityData.ParentYangName = "data-rate-interface-template"
    outputDataRate.EntityData.SegmentPath = "output-data-rate"
    outputDataRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputDataRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputDataRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputDataRate.EntityData.Children = make(map[string]types.YChild)
    outputDataRate.EntityData.Leafs = make(map[string]types.YLeaf)
    outputDataRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputDataRate.Operator}
    outputDataRate.EntityData.Leafs["value"] = types.YLeaf{"Value", outputDataRate.Value}
    outputDataRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputDataRate.EndRangeValue}
    outputDataRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputDataRate.Percent}
    outputDataRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputDataRate.RearmType}
    outputDataRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputDataRate.RearmWindow}
    return &(outputDataRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate
// Number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate struct {
    EntityData types.CommonEntityData
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

func (inputPacketRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate) GetEntityData() *types.CommonEntityData {
    inputPacketRate.EntityData.YFilter = inputPacketRate.YFilter
    inputPacketRate.EntityData.YangName = "input-packet-rate"
    inputPacketRate.EntityData.BundleName = "cisco_ios_xr"
    inputPacketRate.EntityData.ParentYangName = "data-rate-interface-template"
    inputPacketRate.EntityData.SegmentPath = "input-packet-rate"
    inputPacketRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputPacketRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputPacketRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputPacketRate.EntityData.Children = make(map[string]types.YChild)
    inputPacketRate.EntityData.Leafs = make(map[string]types.YLeaf)
    inputPacketRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputPacketRate.Operator}
    inputPacketRate.EntityData.Leafs["value"] = types.YLeaf{"Value", inputPacketRate.Value}
    inputPacketRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputPacketRate.EndRangeValue}
    inputPacketRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputPacketRate.Percent}
    inputPacketRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputPacketRate.RearmType}
    inputPacketRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputPacketRate.RearmWindow}
    return &(inputPacketRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts
// Maximum number of output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts struct {
    EntityData types.CommonEntityData
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

func (outputPeakPkts *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts) GetEntityData() *types.CommonEntityData {
    outputPeakPkts.EntityData.YFilter = outputPeakPkts.YFilter
    outputPeakPkts.EntityData.YangName = "output-peak-pkts"
    outputPeakPkts.EntityData.BundleName = "cisco_ios_xr"
    outputPeakPkts.EntityData.ParentYangName = "data-rate-interface-template"
    outputPeakPkts.EntityData.SegmentPath = "output-peak-pkts"
    outputPeakPkts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPeakPkts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPeakPkts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPeakPkts.EntityData.Children = make(map[string]types.YChild)
    outputPeakPkts.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPeakPkts.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputPeakPkts.Operator}
    outputPeakPkts.EntityData.Leafs["value"] = types.YLeaf{"Value", outputPeakPkts.Value}
    outputPeakPkts.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputPeakPkts.EndRangeValue}
    outputPeakPkts.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputPeakPkts.Percent}
    outputPeakPkts.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputPeakPkts.RearmType}
    outputPeakPkts.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputPeakPkts.RearmWindow}
    return &(outputPeakPkts.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate
// Peak input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate struct {
    EntityData types.CommonEntityData
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

func (inputPeakRate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate) GetEntityData() *types.CommonEntityData {
    inputPeakRate.EntityData.YFilter = inputPeakRate.YFilter
    inputPeakRate.EntityData.YangName = "input-peak-rate"
    inputPeakRate.EntityData.BundleName = "cisco_ios_xr"
    inputPeakRate.EntityData.ParentYangName = "data-rate-interface-template"
    inputPeakRate.EntityData.SegmentPath = "input-peak-rate"
    inputPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputPeakRate.EntityData.Children = make(map[string]types.YChild)
    inputPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    inputPeakRate.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputPeakRate.Operator}
    inputPeakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", inputPeakRate.Value}
    inputPeakRate.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputPeakRate.EndRangeValue}
    inputPeakRate.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputPeakRate.Percent}
    inputPeakRate.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputPeakRate.RearmType}
    inputPeakRate.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputPeakRate.RearmWindow}
    return &(inputPeakRate.EntityData)
}

// PerfMgmt_Threshold_ProcessNode
// Node Process threshold configuration
type PerfMgmt_Threshold_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold templates.
    ProcessNodeTemplates PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates
}

func (processNode *PerfMgmt_Threshold_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "threshold"
    processNode.EntityData.SegmentPath = "process-node"
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = make(map[string]types.YChild)
    processNode.EntityData.Children["process-node-templates"] = types.YChild{"ProcessNodeTemplates", &processNode.ProcessNodeTemplates}
    processNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNode.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates
// Node Memory threshold templates
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold template instance. The type is slice of
    // PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate.
    ProcessNodeTemplate []PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
}

func (processNodeTemplates *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates) GetEntityData() *types.CommonEntityData {
    processNodeTemplates.EntityData.YFilter = processNodeTemplates.YFilter
    processNodeTemplates.EntityData.YangName = "process-node-templates"
    processNodeTemplates.EntityData.BundleName = "cisco_ios_xr"
    processNodeTemplates.EntityData.ParentYangName = "process-node"
    processNodeTemplates.EntityData.SegmentPath = "process-node-templates"
    processNodeTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNodeTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNodeTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNodeTemplates.EntityData.Children = make(map[string]types.YChild)
    processNodeTemplates.EntityData.Children["process-node-template"] = types.YChild{"ProcessNodeTemplate", nil}
    for i := range processNodeTemplates.ProcessNodeTemplate {
        processNodeTemplates.EntityData.Children[types.GetSegmentPath(&processNodeTemplates.ProcessNodeTemplate[i])] = types.YChild{"ProcessNodeTemplate", &processNodeTemplates.ProcessNodeTemplate[i]}
    }
    processNodeTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
// Node Memory threshold template instance
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetEntityData() *types.CommonEntityData {
    processNodeTemplate.EntityData.YFilter = processNodeTemplate.YFilter
    processNodeTemplate.EntityData.YangName = "process-node-template"
    processNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    processNodeTemplate.EntityData.ParentYangName = "process-node-templates"
    processNodeTemplate.EntityData.SegmentPath = "process-node-template" + "[template-name='" + fmt.Sprintf("%v", processNodeTemplate.TemplateName) + "']"
    processNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNodeTemplate.EntityData.Children = make(map[string]types.YChild)
    processNodeTemplate.EntityData.Children["average-cpu-used"] = types.YChild{"AverageCpuUsed", &processNodeTemplate.AverageCpuUsed}
    processNodeTemplate.EntityData.Children["peak-memory"] = types.YChild{"PeakMemory", &processNodeTemplate.PeakMemory}
    processNodeTemplate.EntityData.Children["no-threads"] = types.YChild{"NoThreads", &processNodeTemplate.NoThreads}
    processNodeTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    processNodeTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", processNodeTemplate.TemplateName}
    processNodeTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", processNodeTemplate.SampleInterval}
    return &(processNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed struct {
    EntityData types.CommonEntityData
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

func (averageCpuUsed *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed) GetEntityData() *types.CommonEntityData {
    averageCpuUsed.EntityData.YFilter = averageCpuUsed.YFilter
    averageCpuUsed.EntityData.YangName = "average-cpu-used"
    averageCpuUsed.EntityData.BundleName = "cisco_ios_xr"
    averageCpuUsed.EntityData.ParentYangName = "process-node-template"
    averageCpuUsed.EntityData.SegmentPath = "average-cpu-used"
    averageCpuUsed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageCpuUsed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageCpuUsed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageCpuUsed.EntityData.Children = make(map[string]types.YChild)
    averageCpuUsed.EntityData.Leafs = make(map[string]types.YLeaf)
    averageCpuUsed.EntityData.Leafs["operator"] = types.YLeaf{"Operator", averageCpuUsed.Operator}
    averageCpuUsed.EntityData.Leafs["value"] = types.YLeaf{"Value", averageCpuUsed.Value}
    averageCpuUsed.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", averageCpuUsed.EndRangeValue}
    averageCpuUsed.EntityData.Leafs["percent"] = types.YLeaf{"Percent", averageCpuUsed.Percent}
    averageCpuUsed.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", averageCpuUsed.RearmType}
    averageCpuUsed.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", averageCpuUsed.RearmWindow}
    return &(averageCpuUsed.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory
// Max memory (KBytes) used since startup time
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory struct {
    EntityData types.CommonEntityData
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

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetEntityData() *types.CommonEntityData {
    peakMemory.EntityData.YFilter = peakMemory.YFilter
    peakMemory.EntityData.YangName = "peak-memory"
    peakMemory.EntityData.BundleName = "cisco_ios_xr"
    peakMemory.EntityData.ParentYangName = "process-node-template"
    peakMemory.EntityData.SegmentPath = "peak-memory"
    peakMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakMemory.EntityData.Children = make(map[string]types.YChild)
    peakMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    peakMemory.EntityData.Leafs["operator"] = types.YLeaf{"Operator", peakMemory.Operator}
    peakMemory.EntityData.Leafs["value"] = types.YLeaf{"Value", peakMemory.Value}
    peakMemory.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", peakMemory.EndRangeValue}
    peakMemory.EntityData.Leafs["percent"] = types.YLeaf{"Percent", peakMemory.Percent}
    peakMemory.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", peakMemory.RearmType}
    peakMemory.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", peakMemory.RearmWindow}
    return &(peakMemory.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads
// Number of threads
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads struct {
    EntityData types.CommonEntityData
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

func (noThreads *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads) GetEntityData() *types.CommonEntityData {
    noThreads.EntityData.YFilter = noThreads.YFilter
    noThreads.EntityData.YangName = "no-threads"
    noThreads.EntityData.BundleName = "cisco_ios_xr"
    noThreads.EntityData.ParentYangName = "process-node-template"
    noThreads.EntityData.SegmentPath = "no-threads"
    noThreads.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    noThreads.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    noThreads.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    noThreads.EntityData.Children = make(map[string]types.YChild)
    noThreads.EntityData.Leafs = make(map[string]types.YLeaf)
    noThreads.EntityData.Leafs["operator"] = types.YLeaf{"Operator", noThreads.Operator}
    noThreads.EntityData.Leafs["value"] = types.YLeaf{"Value", noThreads.Value}
    noThreads.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", noThreads.EndRangeValue}
    noThreads.EntityData.Leafs["percent"] = types.YLeaf{"Percent", noThreads.Percent}
    noThreads.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", noThreads.RearmType}
    noThreads.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", noThreads.RearmWindow}
    return &(noThreads.EntityData)
}

// PerfMgmt_Threshold_MemoryNode
// Node Memory threshold configuration
type PerfMgmt_Threshold_MemoryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold configuration templates.
    MemoryNodeTemplates PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates
}

func (memoryNode *PerfMgmt_Threshold_MemoryNode) GetEntityData() *types.CommonEntityData {
    memoryNode.EntityData.YFilter = memoryNode.YFilter
    memoryNode.EntityData.YangName = "memory-node"
    memoryNode.EntityData.BundleName = "cisco_ios_xr"
    memoryNode.EntityData.ParentYangName = "threshold"
    memoryNode.EntityData.SegmentPath = "memory-node"
    memoryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNode.EntityData.Children = make(map[string]types.YChild)
    memoryNode.EntityData.Children["memory-node-templates"] = types.YChild{"MemoryNodeTemplates", &memoryNode.MemoryNodeTemplates}
    memoryNode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryNode.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates
// Node Memory threshold configuration templates
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold configuration template instance. The type is slice of
    // PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate.
    MemoryNodeTemplate []PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
}

func (memoryNodeTemplates *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates) GetEntityData() *types.CommonEntityData {
    memoryNodeTemplates.EntityData.YFilter = memoryNodeTemplates.YFilter
    memoryNodeTemplates.EntityData.YangName = "memory-node-templates"
    memoryNodeTemplates.EntityData.BundleName = "cisco_ios_xr"
    memoryNodeTemplates.EntityData.ParentYangName = "memory-node"
    memoryNodeTemplates.EntityData.SegmentPath = "memory-node-templates"
    memoryNodeTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNodeTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNodeTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNodeTemplates.EntityData.Children = make(map[string]types.YChild)
    memoryNodeTemplates.EntityData.Children["memory-node-template"] = types.YChild{"MemoryNodeTemplate", nil}
    for i := range memoryNodeTemplates.MemoryNodeTemplate {
        memoryNodeTemplates.EntityData.Children[types.GetSegmentPath(&memoryNodeTemplates.MemoryNodeTemplate[i])] = types.YChild{"MemoryNodeTemplate", &memoryNodeTemplates.MemoryNodeTemplate[i]}
    }
    memoryNodeTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memoryNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
// Node Memory threshold configuration template
// instance
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TemplateName interface{}

    // Frequency of sampling in minutes. The type is interface{} with range:
    // 1..60. Units are minute.
    SampleInterval interface{}

    // Maximum memory (KBytes) used.
    PeakMemory PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory

    // Current memory (Bytes) in use.
    CurrMemory PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory
}

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetEntityData() *types.CommonEntityData {
    memoryNodeTemplate.EntityData.YFilter = memoryNodeTemplate.YFilter
    memoryNodeTemplate.EntityData.YangName = "memory-node-template"
    memoryNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    memoryNodeTemplate.EntityData.ParentYangName = "memory-node-templates"
    memoryNodeTemplate.EntityData.SegmentPath = "memory-node-template" + "[template-name='" + fmt.Sprintf("%v", memoryNodeTemplate.TemplateName) + "']"
    memoryNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNodeTemplate.EntityData.Children = make(map[string]types.YChild)
    memoryNodeTemplate.EntityData.Children["peak-memory"] = types.YChild{"PeakMemory", &memoryNodeTemplate.PeakMemory}
    memoryNodeTemplate.EntityData.Children["curr-memory"] = types.YChild{"CurrMemory", &memoryNodeTemplate.CurrMemory}
    memoryNodeTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryNodeTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", memoryNodeTemplate.TemplateName}
    memoryNodeTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", memoryNodeTemplate.SampleInterval}
    return &(memoryNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory
// Maximum memory (KBytes) used
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory struct {
    EntityData types.CommonEntityData
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

func (peakMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory) GetEntityData() *types.CommonEntityData {
    peakMemory.EntityData.YFilter = peakMemory.YFilter
    peakMemory.EntityData.YangName = "peak-memory"
    peakMemory.EntityData.BundleName = "cisco_ios_xr"
    peakMemory.EntityData.ParentYangName = "memory-node-template"
    peakMemory.EntityData.SegmentPath = "peak-memory"
    peakMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakMemory.EntityData.Children = make(map[string]types.YChild)
    peakMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    peakMemory.EntityData.Leafs["operator"] = types.YLeaf{"Operator", peakMemory.Operator}
    peakMemory.EntityData.Leafs["value"] = types.YLeaf{"Value", peakMemory.Value}
    peakMemory.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", peakMemory.EndRangeValue}
    peakMemory.EntityData.Leafs["percent"] = types.YLeaf{"Percent", peakMemory.Percent}
    peakMemory.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", peakMemory.RearmType}
    peakMemory.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", peakMemory.RearmWindow}
    return &(peakMemory.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory
// Current memory (Bytes) in use
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory struct {
    EntityData types.CommonEntityData
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

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetEntityData() *types.CommonEntityData {
    currMemory.EntityData.YFilter = currMemory.YFilter
    currMemory.EntityData.YangName = "curr-memory"
    currMemory.EntityData.BundleName = "cisco_ios_xr"
    currMemory.EntityData.ParentYangName = "memory-node-template"
    currMemory.EntityData.SegmentPath = "curr-memory"
    currMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currMemory.EntityData.Children = make(map[string]types.YChild)
    currMemory.EntityData.Leafs = make(map[string]types.YLeaf)
    currMemory.EntityData.Leafs["operator"] = types.YLeaf{"Operator", currMemory.Operator}
    currMemory.EntityData.Leafs["value"] = types.YLeaf{"Value", currMemory.Value}
    currMemory.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", currMemory.EndRangeValue}
    currMemory.EntityData.Leafs["percent"] = types.YLeaf{"Percent", currMemory.Percent}
    currMemory.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", currMemory.RearmType}
    currMemory.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", currMemory.RearmWindow}
    return &(currMemory.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol
// OSPF v2 Protocol threshold configuration
type PerfMgmt_Threshold_Ospfv3Protocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold templates.
    Ospfv3ProtocolTemplates PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates
}

func (ospfv3Protocol *PerfMgmt_Threshold_Ospfv3Protocol) GetEntityData() *types.CommonEntityData {
    ospfv3Protocol.EntityData.YFilter = ospfv3Protocol.YFilter
    ospfv3Protocol.EntityData.YangName = "ospfv3-protocol"
    ospfv3Protocol.EntityData.BundleName = "cisco_ios_xr"
    ospfv3Protocol.EntityData.ParentYangName = "threshold"
    ospfv3Protocol.EntityData.SegmentPath = "ospfv3-protocol"
    ospfv3Protocol.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3Protocol.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3Protocol.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3Protocol.EntityData.Children = make(map[string]types.YChild)
    ospfv3Protocol.EntityData.Children["ospfv3-protocol-templates"] = types.YChild{"Ospfv3ProtocolTemplates", &ospfv3Protocol.Ospfv3ProtocolTemplates}
    ospfv3Protocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate.
    Ospfv3ProtocolTemplate []PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
}

func (ospfv3ProtocolTemplates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolTemplates.EntityData.YFilter = ospfv3ProtocolTemplates.YFilter
    ospfv3ProtocolTemplates.EntityData.YangName = "ospfv3-protocol-templates"
    ospfv3ProtocolTemplates.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolTemplates.EntityData.ParentYangName = "ospfv3-protocol"
    ospfv3ProtocolTemplates.EntityData.SegmentPath = "ospfv3-protocol-templates"
    ospfv3ProtocolTemplates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolTemplates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolTemplates.EntityData.Children = make(map[string]types.YChild)
    ospfv3ProtocolTemplates.EntityData.Children["ospfv3-protocol-template"] = types.YChild{"Ospfv3ProtocolTemplate", nil}
    for i := range ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate {
        ospfv3ProtocolTemplates.EntityData.Children[types.GetSegmentPath(&ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i])] = types.YChild{"Ospfv3ProtocolTemplate", &ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i]}
    }
    ospfv3ProtocolTemplates.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfv3ProtocolTemplates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Template Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolTemplate.EntityData.YFilter = ospfv3ProtocolTemplate.YFilter
    ospfv3ProtocolTemplate.EntityData.YangName = "ospfv3-protocol-template"
    ospfv3ProtocolTemplate.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolTemplate.EntityData.ParentYangName = "ospfv3-protocol-templates"
    ospfv3ProtocolTemplate.EntityData.SegmentPath = "ospfv3-protocol-template" + "[template-name='" + fmt.Sprintf("%v", ospfv3ProtocolTemplate.TemplateName) + "']"
    ospfv3ProtocolTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolTemplate.EntityData.Children = make(map[string]types.YChild)
    ospfv3ProtocolTemplate.EntityData.Children["input-lsa-acks-lsa"] = types.YChild{"InputLsaAcksLsa", &ospfv3ProtocolTemplate.InputLsaAcksLsa}
    ospfv3ProtocolTemplate.EntityData.Children["output-db-ds-lsa"] = types.YChild{"OutputDbDsLsa", &ospfv3ProtocolTemplate.OutputDbDsLsa}
    ospfv3ProtocolTemplate.EntityData.Children["input-db-ds-lsa"] = types.YChild{"InputDbDsLsa", &ospfv3ProtocolTemplate.InputDbDsLsa}
    ospfv3ProtocolTemplate.EntityData.Children["input-lsa-updates"] = types.YChild{"InputLsaUpdates", &ospfv3ProtocolTemplate.InputLsaUpdates}
    ospfv3ProtocolTemplate.EntityData.Children["output-db-ds"] = types.YChild{"OutputDbDs", &ospfv3ProtocolTemplate.OutputDbDs}
    ospfv3ProtocolTemplate.EntityData.Children["output-lsa-updates-lsa"] = types.YChild{"OutputLsaUpdatesLsa", &ospfv3ProtocolTemplate.OutputLsaUpdatesLsa}
    ospfv3ProtocolTemplate.EntityData.Children["input-db-ds"] = types.YChild{"InputDbDs", &ospfv3ProtocolTemplate.InputDbDs}
    ospfv3ProtocolTemplate.EntityData.Children["input-lsa-updates-lsa"] = types.YChild{"InputLsaUpdatesLsa", &ospfv3ProtocolTemplate.InputLsaUpdatesLsa}
    ospfv3ProtocolTemplate.EntityData.Children["output-packets"] = types.YChild{"OutputPackets", &ospfv3ProtocolTemplate.OutputPackets}
    ospfv3ProtocolTemplate.EntityData.Children["input-packets"] = types.YChild{"InputPackets", &ospfv3ProtocolTemplate.InputPackets}
    ospfv3ProtocolTemplate.EntityData.Children["output-hello-packets"] = types.YChild{"OutputHelloPackets", &ospfv3ProtocolTemplate.OutputHelloPackets}
    ospfv3ProtocolTemplate.EntityData.Children["input-hello-packets"] = types.YChild{"InputHelloPackets", &ospfv3ProtocolTemplate.InputHelloPackets}
    ospfv3ProtocolTemplate.EntityData.Children["output-ls-requests"] = types.YChild{"OutputLsRequests", &ospfv3ProtocolTemplate.OutputLsRequests}
    ospfv3ProtocolTemplate.EntityData.Children["output-lsa-acks-lsa"] = types.YChild{"OutputLsaAcksLsa", &ospfv3ProtocolTemplate.OutputLsaAcksLsa}
    ospfv3ProtocolTemplate.EntityData.Children["output-lsa-acks"] = types.YChild{"OutputLsaAcks", &ospfv3ProtocolTemplate.OutputLsaAcks}
    ospfv3ProtocolTemplate.EntityData.Children["input-lsa-acks"] = types.YChild{"InputLsaAcks", &ospfv3ProtocolTemplate.InputLsaAcks}
    ospfv3ProtocolTemplate.EntityData.Children["output-lsa-updates"] = types.YChild{"OutputLsaUpdates", &ospfv3ProtocolTemplate.OutputLsaUpdates}
    ospfv3ProtocolTemplate.EntityData.Children["output-ls-requests-lsa"] = types.YChild{"OutputLsRequestsLsa", &ospfv3ProtocolTemplate.OutputLsRequestsLsa}
    ospfv3ProtocolTemplate.EntityData.Children["input-ls-requests-lsa"] = types.YChild{"InputLsRequestsLsa", &ospfv3ProtocolTemplate.InputLsRequestsLsa}
    ospfv3ProtocolTemplate.EntityData.Children["input-ls-requests"] = types.YChild{"InputLsRequests", &ospfv3ProtocolTemplate.InputLsRequests}
    ospfv3ProtocolTemplate.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfv3ProtocolTemplate.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", ospfv3ProtocolTemplate.TemplateName}
    ospfv3ProtocolTemplate.EntityData.Leafs["sample-interval"] = types.YLeaf{"SampleInterval", ospfv3ProtocolTemplate.SampleInterval}
    return &(ospfv3ProtocolTemplate.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa) GetEntityData() *types.CommonEntityData {
    inputLsaAcksLsa.EntityData.YFilter = inputLsaAcksLsa.YFilter
    inputLsaAcksLsa.EntityData.YangName = "input-lsa-acks-lsa"
    inputLsaAcksLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsaAcksLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsaAcksLsa.EntityData.SegmentPath = "input-lsa-acks-lsa"
    inputLsaAcksLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaAcksLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaAcksLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaAcksLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsaAcksLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaAcksLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaAcksLsa.Operator}
    inputLsaAcksLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaAcksLsa.Value}
    inputLsaAcksLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaAcksLsa.EndRangeValue}
    inputLsaAcksLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaAcksLsa.Percent}
    inputLsaAcksLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaAcksLsa.RearmType}
    inputLsaAcksLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaAcksLsa.RearmWindow}
    return &(inputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa struct {
    EntityData types.CommonEntityData
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

func (outputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa) GetEntityData() *types.CommonEntityData {
    outputDbDsLsa.EntityData.YFilter = outputDbDsLsa.YFilter
    outputDbDsLsa.EntityData.YangName = "output-db-ds-lsa"
    outputDbDsLsa.EntityData.BundleName = "cisco_ios_xr"
    outputDbDsLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputDbDsLsa.EntityData.SegmentPath = "output-db-ds-lsa"
    outputDbDsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputDbDsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputDbDsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputDbDsLsa.EntityData.Children = make(map[string]types.YChild)
    outputDbDsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputDbDsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputDbDsLsa.Operator}
    outputDbDsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputDbDsLsa.Value}
    outputDbDsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputDbDsLsa.EndRangeValue}
    outputDbDsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputDbDsLsa.Percent}
    outputDbDsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputDbDsLsa.RearmType}
    outputDbDsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputDbDsLsa.RearmWindow}
    return &(outputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa struct {
    EntityData types.CommonEntityData
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

func (inputDbDsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa) GetEntityData() *types.CommonEntityData {
    inputDbDsLsa.EntityData.YFilter = inputDbDsLsa.YFilter
    inputDbDsLsa.EntityData.YangName = "input-db-ds-lsa"
    inputDbDsLsa.EntityData.BundleName = "cisco_ios_xr"
    inputDbDsLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputDbDsLsa.EntityData.SegmentPath = "input-db-ds-lsa"
    inputDbDsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputDbDsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputDbDsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputDbDsLsa.EntityData.Children = make(map[string]types.YChild)
    inputDbDsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputDbDsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputDbDsLsa.Operator}
    inputDbDsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputDbDsLsa.Value}
    inputDbDsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputDbDsLsa.EndRangeValue}
    inputDbDsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputDbDsLsa.Percent}
    inputDbDsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputDbDsLsa.RearmType}
    inputDbDsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputDbDsLsa.RearmWindow}
    return &(inputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates struct {
    EntityData types.CommonEntityData
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

func (inputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates) GetEntityData() *types.CommonEntityData {
    inputLsaUpdates.EntityData.YFilter = inputLsaUpdates.YFilter
    inputLsaUpdates.EntityData.YangName = "input-lsa-updates"
    inputLsaUpdates.EntityData.BundleName = "cisco_ios_xr"
    inputLsaUpdates.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsaUpdates.EntityData.SegmentPath = "input-lsa-updates"
    inputLsaUpdates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaUpdates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaUpdates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaUpdates.EntityData.Children = make(map[string]types.YChild)
    inputLsaUpdates.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaUpdates.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaUpdates.Operator}
    inputLsaUpdates.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaUpdates.Value}
    inputLsaUpdates.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaUpdates.EndRangeValue}
    inputLsaUpdates.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaUpdates.Percent}
    inputLsaUpdates.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaUpdates.RearmType}
    inputLsaUpdates.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaUpdates.RearmWindow}
    return &(inputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs struct {
    EntityData types.CommonEntityData
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

func (outputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs) GetEntityData() *types.CommonEntityData {
    outputDbDs.EntityData.YFilter = outputDbDs.YFilter
    outputDbDs.EntityData.YangName = "output-db-ds"
    outputDbDs.EntityData.BundleName = "cisco_ios_xr"
    outputDbDs.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputDbDs.EntityData.SegmentPath = "output-db-ds"
    outputDbDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputDbDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputDbDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputDbDs.EntityData.Children = make(map[string]types.YChild)
    outputDbDs.EntityData.Leafs = make(map[string]types.YLeaf)
    outputDbDs.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputDbDs.Operator}
    outputDbDs.EntityData.Leafs["value"] = types.YLeaf{"Value", outputDbDs.Value}
    outputDbDs.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputDbDs.EndRangeValue}
    outputDbDs.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputDbDs.Percent}
    outputDbDs.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputDbDs.RearmType}
    outputDbDs.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputDbDs.RearmWindow}
    return &(outputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa) GetEntityData() *types.CommonEntityData {
    outputLsaUpdatesLsa.EntityData.YFilter = outputLsaUpdatesLsa.YFilter
    outputLsaUpdatesLsa.EntityData.YangName = "output-lsa-updates-lsa"
    outputLsaUpdatesLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsaUpdatesLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsaUpdatesLsa.EntityData.SegmentPath = "output-lsa-updates-lsa"
    outputLsaUpdatesLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaUpdatesLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaUpdatesLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaUpdatesLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsaUpdatesLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaUpdatesLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaUpdatesLsa.Operator}
    outputLsaUpdatesLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaUpdatesLsa.Value}
    outputLsaUpdatesLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaUpdatesLsa.EndRangeValue}
    outputLsaUpdatesLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaUpdatesLsa.Percent}
    outputLsaUpdatesLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaUpdatesLsa.RearmType}
    outputLsaUpdatesLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaUpdatesLsa.RearmWindow}
    return &(outputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs struct {
    EntityData types.CommonEntityData
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

func (inputDbDs *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs) GetEntityData() *types.CommonEntityData {
    inputDbDs.EntityData.YFilter = inputDbDs.YFilter
    inputDbDs.EntityData.YangName = "input-db-ds"
    inputDbDs.EntityData.BundleName = "cisco_ios_xr"
    inputDbDs.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputDbDs.EntityData.SegmentPath = "input-db-ds"
    inputDbDs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputDbDs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputDbDs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputDbDs.EntityData.Children = make(map[string]types.YChild)
    inputDbDs.EntityData.Leafs = make(map[string]types.YLeaf)
    inputDbDs.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputDbDs.Operator}
    inputDbDs.EntityData.Leafs["value"] = types.YLeaf{"Value", inputDbDs.Value}
    inputDbDs.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputDbDs.EndRangeValue}
    inputDbDs.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputDbDs.Percent}
    inputDbDs.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputDbDs.RearmType}
    inputDbDs.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputDbDs.RearmWindow}
    return &(inputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsaUpdatesLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa) GetEntityData() *types.CommonEntityData {
    inputLsaUpdatesLsa.EntityData.YFilter = inputLsaUpdatesLsa.YFilter
    inputLsaUpdatesLsa.EntityData.YangName = "input-lsa-updates-lsa"
    inputLsaUpdatesLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsaUpdatesLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsaUpdatesLsa.EntityData.SegmentPath = "input-lsa-updates-lsa"
    inputLsaUpdatesLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaUpdatesLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaUpdatesLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaUpdatesLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsaUpdatesLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaUpdatesLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaUpdatesLsa.Operator}
    inputLsaUpdatesLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaUpdatesLsa.Value}
    inputLsaUpdatesLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaUpdatesLsa.EndRangeValue}
    inputLsaUpdatesLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaUpdatesLsa.Percent}
    inputLsaUpdatesLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaUpdatesLsa.RearmType}
    inputLsaUpdatesLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaUpdatesLsa.RearmWindow}
    return &(inputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets struct {
    EntityData types.CommonEntityData
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

func (outputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets) GetEntityData() *types.CommonEntityData {
    outputPackets.EntityData.YFilter = outputPackets.YFilter
    outputPackets.EntityData.YangName = "output-packets"
    outputPackets.EntityData.BundleName = "cisco_ios_xr"
    outputPackets.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputPackets.EntityData.SegmentPath = "output-packets"
    outputPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPackets.EntityData.Children = make(map[string]types.YChild)
    outputPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputPackets.Operator}
    outputPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outputPackets.Value}
    outputPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputPackets.EndRangeValue}
    outputPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputPackets.Percent}
    outputPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputPackets.RearmType}
    outputPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputPackets.RearmWindow}
    return &(outputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets struct {
    EntityData types.CommonEntityData
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

func (inputPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets) GetEntityData() *types.CommonEntityData {
    inputPackets.EntityData.YFilter = inputPackets.YFilter
    inputPackets.EntityData.YangName = "input-packets"
    inputPackets.EntityData.BundleName = "cisco_ios_xr"
    inputPackets.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputPackets.EntityData.SegmentPath = "input-packets"
    inputPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputPackets.EntityData.Children = make(map[string]types.YChild)
    inputPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inputPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputPackets.Operator}
    inputPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inputPackets.Value}
    inputPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputPackets.EndRangeValue}
    inputPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputPackets.Percent}
    inputPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputPackets.RearmType}
    inputPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputPackets.RearmWindow}
    return &(inputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets struct {
    EntityData types.CommonEntityData
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

func (outputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets) GetEntityData() *types.CommonEntityData {
    outputHelloPackets.EntityData.YFilter = outputHelloPackets.YFilter
    outputHelloPackets.EntityData.YangName = "output-hello-packets"
    outputHelloPackets.EntityData.BundleName = "cisco_ios_xr"
    outputHelloPackets.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputHelloPackets.EntityData.SegmentPath = "output-hello-packets"
    outputHelloPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputHelloPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputHelloPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputHelloPackets.EntityData.Children = make(map[string]types.YChild)
    outputHelloPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    outputHelloPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputHelloPackets.Operator}
    outputHelloPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", outputHelloPackets.Value}
    outputHelloPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputHelloPackets.EndRangeValue}
    outputHelloPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputHelloPackets.Percent}
    outputHelloPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputHelloPackets.RearmType}
    outputHelloPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputHelloPackets.RearmWindow}
    return &(outputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets struct {
    EntityData types.CommonEntityData
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

func (inputHelloPackets *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets) GetEntityData() *types.CommonEntityData {
    inputHelloPackets.EntityData.YFilter = inputHelloPackets.YFilter
    inputHelloPackets.EntityData.YangName = "input-hello-packets"
    inputHelloPackets.EntityData.BundleName = "cisco_ios_xr"
    inputHelloPackets.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputHelloPackets.EntityData.SegmentPath = "input-hello-packets"
    inputHelloPackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputHelloPackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputHelloPackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputHelloPackets.EntityData.Children = make(map[string]types.YChild)
    inputHelloPackets.EntityData.Leafs = make(map[string]types.YLeaf)
    inputHelloPackets.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputHelloPackets.Operator}
    inputHelloPackets.EntityData.Leafs["value"] = types.YLeaf{"Value", inputHelloPackets.Value}
    inputHelloPackets.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputHelloPackets.EndRangeValue}
    inputHelloPackets.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputHelloPackets.Percent}
    inputHelloPackets.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputHelloPackets.RearmType}
    inputHelloPackets.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputHelloPackets.RearmWindow}
    return &(inputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests struct {
    EntityData types.CommonEntityData
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

func (outputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests) GetEntityData() *types.CommonEntityData {
    outputLsRequests.EntityData.YFilter = outputLsRequests.YFilter
    outputLsRequests.EntityData.YangName = "output-ls-requests"
    outputLsRequests.EntityData.BundleName = "cisco_ios_xr"
    outputLsRequests.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsRequests.EntityData.SegmentPath = "output-ls-requests"
    outputLsRequests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsRequests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsRequests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsRequests.EntityData.Children = make(map[string]types.YChild)
    outputLsRequests.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsRequests.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsRequests.Operator}
    outputLsRequests.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsRequests.Value}
    outputLsRequests.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsRequests.EndRangeValue}
    outputLsRequests.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsRequests.Percent}
    outputLsRequests.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsRequests.RearmType}
    outputLsRequests.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsRequests.RearmWindow}
    return &(outputLsRequests.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsaAcksLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa) GetEntityData() *types.CommonEntityData {
    outputLsaAcksLsa.EntityData.YFilter = outputLsaAcksLsa.YFilter
    outputLsaAcksLsa.EntityData.YangName = "output-lsa-acks-lsa"
    outputLsaAcksLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsaAcksLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsaAcksLsa.EntityData.SegmentPath = "output-lsa-acks-lsa"
    outputLsaAcksLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaAcksLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaAcksLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaAcksLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsaAcksLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaAcksLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaAcksLsa.Operator}
    outputLsaAcksLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaAcksLsa.Value}
    outputLsaAcksLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaAcksLsa.EndRangeValue}
    outputLsaAcksLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaAcksLsa.Percent}
    outputLsaAcksLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaAcksLsa.RearmType}
    outputLsaAcksLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaAcksLsa.RearmWindow}
    return &(outputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks struct {
    EntityData types.CommonEntityData
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

func (outputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks) GetEntityData() *types.CommonEntityData {
    outputLsaAcks.EntityData.YFilter = outputLsaAcks.YFilter
    outputLsaAcks.EntityData.YangName = "output-lsa-acks"
    outputLsaAcks.EntityData.BundleName = "cisco_ios_xr"
    outputLsaAcks.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsaAcks.EntityData.SegmentPath = "output-lsa-acks"
    outputLsaAcks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaAcks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaAcks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaAcks.EntityData.Children = make(map[string]types.YChild)
    outputLsaAcks.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaAcks.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaAcks.Operator}
    outputLsaAcks.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaAcks.Value}
    outputLsaAcks.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaAcks.EndRangeValue}
    outputLsaAcks.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaAcks.Percent}
    outputLsaAcks.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaAcks.RearmType}
    outputLsaAcks.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaAcks.RearmWindow}
    return &(outputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks struct {
    EntityData types.CommonEntityData
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

func (inputLsaAcks *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks) GetEntityData() *types.CommonEntityData {
    inputLsaAcks.EntityData.YFilter = inputLsaAcks.YFilter
    inputLsaAcks.EntityData.YangName = "input-lsa-acks"
    inputLsaAcks.EntityData.BundleName = "cisco_ios_xr"
    inputLsaAcks.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsaAcks.EntityData.SegmentPath = "input-lsa-acks"
    inputLsaAcks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsaAcks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsaAcks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsaAcks.EntityData.Children = make(map[string]types.YChild)
    inputLsaAcks.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsaAcks.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsaAcks.Operator}
    inputLsaAcks.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsaAcks.Value}
    inputLsaAcks.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsaAcks.EndRangeValue}
    inputLsaAcks.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsaAcks.Percent}
    inputLsaAcks.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsaAcks.RearmType}
    inputLsaAcks.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsaAcks.RearmWindow}
    return &(inputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates struct {
    EntityData types.CommonEntityData
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

func (outputLsaUpdates *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates) GetEntityData() *types.CommonEntityData {
    outputLsaUpdates.EntityData.YFilter = outputLsaUpdates.YFilter
    outputLsaUpdates.EntityData.YangName = "output-lsa-updates"
    outputLsaUpdates.EntityData.BundleName = "cisco_ios_xr"
    outputLsaUpdates.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsaUpdates.EntityData.SegmentPath = "output-lsa-updates"
    outputLsaUpdates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsaUpdates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsaUpdates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsaUpdates.EntityData.Children = make(map[string]types.YChild)
    outputLsaUpdates.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsaUpdates.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsaUpdates.Operator}
    outputLsaUpdates.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsaUpdates.Value}
    outputLsaUpdates.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsaUpdates.EndRangeValue}
    outputLsaUpdates.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsaUpdates.Percent}
    outputLsaUpdates.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsaUpdates.RearmType}
    outputLsaUpdates.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsaUpdates.RearmWindow}
    return &(outputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa struct {
    EntityData types.CommonEntityData
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

func (outputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa) GetEntityData() *types.CommonEntityData {
    outputLsRequestsLsa.EntityData.YFilter = outputLsRequestsLsa.YFilter
    outputLsRequestsLsa.EntityData.YangName = "output-ls-requests-lsa"
    outputLsRequestsLsa.EntityData.BundleName = "cisco_ios_xr"
    outputLsRequestsLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    outputLsRequestsLsa.EntityData.SegmentPath = "output-ls-requests-lsa"
    outputLsRequestsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputLsRequestsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputLsRequestsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputLsRequestsLsa.EntityData.Children = make(map[string]types.YChild)
    outputLsRequestsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    outputLsRequestsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", outputLsRequestsLsa.Operator}
    outputLsRequestsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", outputLsRequestsLsa.Value}
    outputLsRequestsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", outputLsRequestsLsa.EndRangeValue}
    outputLsRequestsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", outputLsRequestsLsa.Percent}
    outputLsRequestsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", outputLsRequestsLsa.RearmType}
    outputLsRequestsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", outputLsRequestsLsa.RearmWindow}
    return &(outputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa struct {
    EntityData types.CommonEntityData
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

func (inputLsRequestsLsa *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa) GetEntityData() *types.CommonEntityData {
    inputLsRequestsLsa.EntityData.YFilter = inputLsRequestsLsa.YFilter
    inputLsRequestsLsa.EntityData.YangName = "input-ls-requests-lsa"
    inputLsRequestsLsa.EntityData.BundleName = "cisco_ios_xr"
    inputLsRequestsLsa.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsRequestsLsa.EntityData.SegmentPath = "input-ls-requests-lsa"
    inputLsRequestsLsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsRequestsLsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsRequestsLsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsRequestsLsa.EntityData.Children = make(map[string]types.YChild)
    inputLsRequestsLsa.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsRequestsLsa.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsRequestsLsa.Operator}
    inputLsRequestsLsa.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsRequestsLsa.Value}
    inputLsRequestsLsa.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsRequestsLsa.EndRangeValue}
    inputLsRequestsLsa.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsRequestsLsa.Percent}
    inputLsRequestsLsa.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsRequestsLsa.RearmType}
    inputLsRequestsLsa.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsRequestsLsa.RearmWindow}
    return &(inputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests struct {
    EntityData types.CommonEntityData
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

func (inputLsRequests *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests) GetEntityData() *types.CommonEntityData {
    inputLsRequests.EntityData.YFilter = inputLsRequests.YFilter
    inputLsRequests.EntityData.YangName = "input-ls-requests"
    inputLsRequests.EntityData.BundleName = "cisco_ios_xr"
    inputLsRequests.EntityData.ParentYangName = "ospfv3-protocol-template"
    inputLsRequests.EntityData.SegmentPath = "input-ls-requests"
    inputLsRequests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputLsRequests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputLsRequests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputLsRequests.EntityData.Children = make(map[string]types.YChild)
    inputLsRequests.EntityData.Leafs = make(map[string]types.YLeaf)
    inputLsRequests.EntityData.Leafs["operator"] = types.YLeaf{"Operator", inputLsRequests.Operator}
    inputLsRequests.EntityData.Leafs["value"] = types.YLeaf{"Value", inputLsRequests.Value}
    inputLsRequests.EntityData.Leafs["end-range-value"] = types.YLeaf{"EndRangeValue", inputLsRequests.EndRangeValue}
    inputLsRequests.EntityData.Leafs["percent"] = types.YLeaf{"Percent", inputLsRequests.Percent}
    inputLsRequests.EntityData.Leafs["rearm-type"] = types.YLeaf{"RearmType", inputLsRequests.RearmType}
    inputLsRequests.EntityData.Leafs["rearm-window"] = types.YLeaf{"RearmWindow", inputLsRequests.RearmWindow}
    return &(inputLsRequests.EntityData)
}

