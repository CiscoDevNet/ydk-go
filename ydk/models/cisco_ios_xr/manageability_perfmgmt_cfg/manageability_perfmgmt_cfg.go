// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-perfmgmt package configuration.
// 
// This module contains definitions
// for the following management objects:
//   perf-mgmt: Performance Management configuration & operations
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    perfMgmt.EntityData.Children = types.NewOrderedMap()
    perfMgmt.EntityData.Children.Append("resources", types.YChild{"Resources", &perfMgmt.Resources})
    perfMgmt.EntityData.Children.Append("statistics", types.YChild{"Statistics", &perfMgmt.Statistics})
    perfMgmt.EntityData.Children.Append("enable", types.YChild{"Enable", &perfMgmt.Enable})
    perfMgmt.EntityData.Children.Append("reg-exp-groups", types.YChild{"RegExpGroups", &perfMgmt.RegExpGroups})
    perfMgmt.EntityData.Children.Append("threshold", types.YChild{"Threshold", &perfMgmt.Threshold})
    perfMgmt.EntityData.Leafs = types.NewOrderedMap()

    perfMgmt.EntityData.YListKeys = []string {}

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

    resources.EntityData.Children = types.NewOrderedMap()
    resources.EntityData.Children.Append("tftp-resources", types.YChild{"TftpResources", &resources.TftpResources})
    resources.EntityData.Children.Append("dump-local", types.YChild{"DumpLocal", &resources.DumpLocal})
    resources.EntityData.Children.Append("memory-resources", types.YChild{"MemoryResources", &resources.MemoryResources})
    resources.EntityData.Leafs = types.NewOrderedMap()

    resources.EntityData.YListKeys = []string {}

    return &(resources.EntityData)
}

// PerfMgmt_Resources_TftpResources
// Configure the TFTP server IP address and
// directory name
// This type is a presence type.
type PerfMgmt_Resources_TftpResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (tftpResources *PerfMgmt_Resources_TftpResources) GetEntityData() *types.CommonEntityData {
    tftpResources.EntityData.YFilter = tftpResources.YFilter
    tftpResources.EntityData.YangName = "tftp-resources"
    tftpResources.EntityData.BundleName = "cisco_ios_xr"
    tftpResources.EntityData.ParentYangName = "resources"
    tftpResources.EntityData.SegmentPath = "tftp-resources"
    tftpResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpResources.EntityData.Children = types.NewOrderedMap()
    tftpResources.EntityData.Leafs = types.NewOrderedMap()
    tftpResources.EntityData.Leafs.Append("server-address", types.YLeaf{"ServerAddress", tftpResources.ServerAddress})
    tftpResources.EntityData.Leafs.Append("directory", types.YLeaf{"Directory", tftpResources.Directory})
    tftpResources.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", tftpResources.VrfName})

    tftpResources.EntityData.YListKeys = []string {}

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

    dumpLocal.EntityData.Children = types.NewOrderedMap()
    dumpLocal.EntityData.Leafs = types.NewOrderedMap()
    dumpLocal.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", dumpLocal.Enable})

    dumpLocal.EntityData.YListKeys = []string {}

    return &(dumpLocal.EntityData)
}

// PerfMgmt_Resources_MemoryResources
// Configure the memory usage limits of
// performance management
type PerfMgmt_Resources_MemoryResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum limit for memory usage (Kbytes) for data buffers. The type is
    // interface{} with range: 0..4294967295. Units are kilobyte.
    MaxLimit interface{}

    // Specify a minimum free memory (Kbytes) to be ensured before allowing a
    // collection request. The type is interface{} with range: 0..4294967295.
    // Units are kilobyte.
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

    memoryResources.EntityData.Children = types.NewOrderedMap()
    memoryResources.EntityData.Leafs = types.NewOrderedMap()
    memoryResources.EntityData.Leafs.Append("max-limit", types.YLeaf{"MaxLimit", memoryResources.MaxLimit})
    memoryResources.EntityData.Leafs.Append("min-reserved", types.YLeaf{"MinReserved", memoryResources.MinReserved})

    memoryResources.EntityData.YListKeys = []string {}

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

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", &statistics.GenericCounterInterface})
    statistics.EntityData.Children.Append("process-node", types.YChild{"ProcessNode", &statistics.ProcessNode})
    statistics.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", &statistics.BasicCounterInterface})
    statistics.EntityData.Children.Append("ospfv3-protocol", types.YChild{"Ospfv3Protocol", &statistics.Ospfv3Protocol})
    statistics.EntityData.Children.Append("cpu-node", types.YChild{"CpuNode", &statistics.CpuNode})
    statistics.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", &statistics.DataRateInterface})
    statistics.EntityData.Children.Append("memory-node", types.YChild{"MemoryNode", &statistics.MemoryNode})
    statistics.EntityData.Children.Append("ldp-mpls", types.YChild{"LdpMpls", &statistics.LdpMpls})
    statistics.EntityData.Children.Append("bgp", types.YChild{"Bgp", &statistics.Bgp})
    statistics.EntityData.Children.Append("ospfv2-protocol", types.YChild{"Ospfv2Protocol", &statistics.Ospfv2Protocol})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

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

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Children.Append("templates", types.YChild{"Templates", &genericCounterInterface.Templates})
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()

    genericCounterInterface.EntityData.YListKeys = []string {}

    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Statistics_GenericCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_GenericCounterInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_GenericCounterInterface_Templates_Template.
    Template []*PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_GenericCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_GenericCounterInterface_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_GenericCounterInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    processNode.EntityData.Children = types.NewOrderedMap()
    processNode.EntityData.Children.Append("templates", types.YChild{"Templates", &processNode.Templates})
    processNode.EntityData.Leafs = types.NewOrderedMap()

    processNode.EntityData.YListKeys = []string {}

    return &(processNode.EntityData)
}

// PerfMgmt_Statistics_ProcessNode_Templates
// Template name
type PerfMgmt_Statistics_ProcessNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_ProcessNode_Templates_Template.
    Template []*PerfMgmt_Statistics_ProcessNode_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_ProcessNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_ProcessNode_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_ProcessNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Children.Append("templates", types.YChild{"Templates", &basicCounterInterface.Templates})
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()

    basicCounterInterface.EntityData.YListKeys = []string {}

    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Statistics_BasicCounterInterface_Templates
// Template name
type PerfMgmt_Statistics_BasicCounterInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_BasicCounterInterface_Templates_Template.
    Template []*PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_BasicCounterInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_BasicCounterInterface_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_BasicCounterInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    ospfv3Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Children.Append("templates", types.YChild{"Templates", &ospfv3Protocol.Templates})
    ospfv3Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Protocol.EntityData.YListKeys = []string {}

    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Statistics_Ospfv3Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv3Protocol_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template.
    Template []*PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_Ospfv3Protocol_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    cpuNode.EntityData.Children = types.NewOrderedMap()
    cpuNode.EntityData.Children.Append("templates", types.YChild{"Templates", &cpuNode.Templates})
    cpuNode.EntityData.Leafs = types.NewOrderedMap()

    cpuNode.EntityData.YListKeys = []string {}

    return &(cpuNode.EntityData)
}

// PerfMgmt_Statistics_CpuNode_Templates
// Template name
type PerfMgmt_Statistics_CpuNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_CpuNode_Templates_Template.
    Template []*PerfMgmt_Statistics_CpuNode_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_CpuNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_CpuNode_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_CpuNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Children.Append("templates", types.YChild{"Templates", &dataRateInterface.Templates})
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()

    dataRateInterface.EntityData.YListKeys = []string {}

    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Statistics_DataRateInterface_Templates
// Template name
type PerfMgmt_Statistics_DataRateInterface_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_DataRateInterface_Templates_Template.
    Template []*PerfMgmt_Statistics_DataRateInterface_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_DataRateInterface_Templates_Template
// A template instance
type PerfMgmt_Statistics_DataRateInterface_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_DataRateInterface_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    memoryNode.EntityData.Children = types.NewOrderedMap()
    memoryNode.EntityData.Children.Append("templates", types.YChild{"Templates", &memoryNode.Templates})
    memoryNode.EntityData.Leafs = types.NewOrderedMap()

    memoryNode.EntityData.YListKeys = []string {}

    return &(memoryNode.EntityData)
}

// PerfMgmt_Statistics_MemoryNode_Templates
// Template name
type PerfMgmt_Statistics_MemoryNode_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_MemoryNode_Templates_Template.
    Template []*PerfMgmt_Statistics_MemoryNode_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_MemoryNode_Templates_Template
// A template instance
type PerfMgmt_Statistics_MemoryNode_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_MemoryNode_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    ldpMpls.EntityData.Children = types.NewOrderedMap()
    ldpMpls.EntityData.Children.Append("templates", types.YChild{"Templates", &ldpMpls.Templates})
    ldpMpls.EntityData.Leafs = types.NewOrderedMap()

    ldpMpls.EntityData.YListKeys = []string {}

    return &(ldpMpls.EntityData)
}

// PerfMgmt_Statistics_LdpMpls_Templates
// Template name
type PerfMgmt_Statistics_LdpMpls_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_LdpMpls_Templates_Template.
    Template []*PerfMgmt_Statistics_LdpMpls_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_LdpMpls_Templates_Template
// A template instance
type PerfMgmt_Statistics_LdpMpls_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_LdpMpls_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("templates", types.YChild{"Templates", &bgp.Templates})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PerfMgmt_Statistics_Bgp_Templates
// Template name
type PerfMgmt_Statistics_Bgp_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Bgp_Templates_Template.
    Template []*PerfMgmt_Statistics_Bgp_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Bgp_Templates_Template
// A template instance
type PerfMgmt_Statistics_Bgp_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_Bgp_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    ospfv2Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Children.Append("templates", types.YChild{"Templates", &ospfv2Protocol.Templates})
    ospfv2Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Protocol.EntityData.YListKeys = []string {}

    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Statistics_Ospfv2Protocol_Templates
// Template name
type PerfMgmt_Statistics_Ospfv2Protocol_Templates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A template instance. The type is slice of
    // PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template.
    Template []*PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
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

    templates.EntityData.Children = types.NewOrderedMap()
    templates.EntityData.Children.Append("template", types.YChild{"Template", nil})
    for i := range templates.Template {
        templates.EntityData.Children.Append(types.GetSegmentPath(templates.Template[i]), types.YChild{"Template", templates.Template[i]})
    }
    templates.EntityData.Leafs = types.NewOrderedMap()

    templates.EntityData.YListKeys = []string {}

    return &(templates.EntityData)
}

// PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template
// A template instance
type PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template struct {
    EntityData types.CommonEntityData
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

func (template *PerfMgmt_Statistics_Ospfv2Protocol_Templates_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "templates"
    template.EntityData.SegmentPath = "template" + types.AddKeyToken(template.TemplateName, "template-name")
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Leafs = types.NewOrderedMap()
    template.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", template.TemplateName})
    template.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", template.RegExpGroup})
    template.EntityData.Leafs.Append("history-persistent", types.YLeaf{"HistoryPersistent", template.HistoryPersistent})
    template.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", template.VrfGroup})
    template.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", template.SampleInterval})
    template.EntityData.Leafs.Append("sample-size", types.YLeaf{"SampleSize", template.SampleSize})

    template.EntityData.YListKeys = []string {"TemplateName"}

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

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("threshold", types.YChild{"Threshold", &enable.Threshold})
    enable.EntityData.Children.Append("statistics", types.YChild{"Statistics", &enable.Statistics})
    enable.EntityData.Children.Append("monitor-enable", types.YChild{"MonitorEnable", &enable.MonitorEnable})
    enable.EntityData.Leafs = types.NewOrderedMap()

    enable.EntityData.YListKeys = []string {}

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

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("ospfv3-protocol", types.YChild{"Ospfv3Protocol", &threshold.Ospfv3Protocol})
    threshold.EntityData.Children.Append("bgp", types.YChild{"Bgp", &threshold.Bgp})
    threshold.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", &threshold.DataRateInterface})
    threshold.EntityData.Children.Append("ospfv2-protocol", types.YChild{"Ospfv2Protocol", &threshold.Ospfv2Protocol})
    threshold.EntityData.Children.Append("memory-node", types.YChild{"MemoryNode", &threshold.MemoryNode})
    threshold.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", &threshold.GenericCounterInterface})
    threshold.EntityData.Children.Append("cpu-node", types.YChild{"CpuNode", &threshold.CpuNode})
    threshold.EntityData.Children.Append("ldp-mpls", types.YChild{"LdpMpls", &threshold.LdpMpls})
    threshold.EntityData.Children.Append("process-node", types.YChild{"ProcessNode", &threshold.ProcessNode})
    threshold.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", &threshold.BasicCounterInterface})
    threshold.EntityData.Leafs = types.NewOrderedMap()

    threshold.EntityData.YListKeys = []string {}

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

    ospfv3Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv3Protocol.TemplateName})

    ospfv3Protocol.EntityData.YListKeys = []string {}

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

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", bgp.TemplateName})

    bgp.EntityData.YListKeys = []string {}

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

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", dataRateInterface.TemplateName})

    dataRateInterface.EntityData.YListKeys = []string {}

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

    ospfv2Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv2Protocol.TemplateName})

    ospfv2Protocol.EntityData.YListKeys = []string {}

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

    memoryNode.EntityData.Children = types.NewOrderedMap()
    memoryNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &memoryNode.Nodes})
    memoryNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &memoryNode.NodeAll})
    memoryNode.EntityData.Leafs = types.NewOrderedMap()

    memoryNode.EntityData.YListKeys = []string {}

    return &(memoryNode.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_MemoryNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

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

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", genericCounterInterface.TemplateName})

    genericCounterInterface.EntityData.YListKeys = []string {}

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

    cpuNode.EntityData.Children = types.NewOrderedMap()
    cpuNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &cpuNode.Nodes})
    cpuNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &cpuNode.NodeAll})
    cpuNode.EntityData.Leafs = types.NewOrderedMap()

    cpuNode.EntityData.YListKeys = []string {}

    return &(cpuNode.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_CpuNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_CpuNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

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

    ldpMpls.EntityData.Children = types.NewOrderedMap()
    ldpMpls.EntityData.Leafs = types.NewOrderedMap()
    ldpMpls.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ldpMpls.TemplateName})

    ldpMpls.EntityData.YListKeys = []string {}

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

    processNode.EntityData.Children = types.NewOrderedMap()
    processNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &processNode.Nodes})
    processNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &processNode.NodeAll})
    processNode.EntityData.Leafs = types.NewOrderedMap()

    processNode.EntityData.YListKeys = []string {}

    return &(processNode.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Threshold_ProcessNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

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

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", basicCounterInterface.TemplateName})

    basicCounterInterface.EntityData.YListKeys = []string {}

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

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", &statistics.GenericCounterInterface})
    statistics.EntityData.Children.Append("bgp", types.YChild{"Bgp", &statistics.Bgp})
    statistics.EntityData.Children.Append("ospfv2-protocol", types.YChild{"Ospfv2Protocol", &statistics.Ospfv2Protocol})
    statistics.EntityData.Children.Append("ospfv3-protocol", types.YChild{"Ospfv3Protocol", &statistics.Ospfv3Protocol})
    statistics.EntityData.Children.Append("cpu-node", types.YChild{"CpuNode", &statistics.CpuNode})
    statistics.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", &statistics.BasicCounterInterface})
    statistics.EntityData.Children.Append("process-node", types.YChild{"ProcessNode", &statistics.ProcessNode})
    statistics.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", &statistics.DataRateInterface})
    statistics.EntityData.Children.Append("memory-node", types.YChild{"MemoryNode", &statistics.MemoryNode})
    statistics.EntityData.Children.Append("ldp-mpls", types.YChild{"LdpMpls", &statistics.LdpMpls})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

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

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    genericCounterInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", genericCounterInterface.TemplateName})

    genericCounterInterface.EntityData.YListKeys = []string {}

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

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", bgp.TemplateName})

    bgp.EntityData.YListKeys = []string {}

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

    ospfv2Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Leafs = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv2Protocol.TemplateName})

    ospfv2Protocol.EntityData.YListKeys = []string {}

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

    ospfv3Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Leafs = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv3Protocol.TemplateName})

    ospfv3Protocol.EntityData.YListKeys = []string {}

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

    cpuNode.EntityData.Children = types.NewOrderedMap()
    cpuNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &cpuNode.NodeAll})
    cpuNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &cpuNode.Nodes})
    cpuNode.EntityData.Leafs = types.NewOrderedMap()

    cpuNode.EntityData.YListKeys = []string {}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_CpuNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_CpuNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()
    basicCounterInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", basicCounterInterface.TemplateName})

    basicCounterInterface.EntityData.YListKeys = []string {}

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

    processNode.EntityData.Children = types.NewOrderedMap()
    processNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &processNode.NodeAll})
    processNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &processNode.Nodes})
    processNode.EntityData.Leafs = types.NewOrderedMap()

    processNode.EntityData.YListKeys = []string {}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_ProcessNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()
    dataRateInterface.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", dataRateInterface.TemplateName})

    dataRateInterface.EntityData.YListKeys = []string {}

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

    memoryNode.EntityData.Children = types.NewOrderedMap()
    memoryNode.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &memoryNode.NodeAll})
    memoryNode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &memoryNode.Nodes})
    memoryNode.EntityData.Leafs = types.NewOrderedMap()

    memoryNode.EntityData.YListKeys = []string {}

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

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", nodeAll.TemplateName})

    nodeAll.EntityData.YListKeys = []string {}

    return &(nodeAll.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes
// Node specification
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node.
    Node []*PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node
// Node instance
type PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_Statistics_MemoryNode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    ldpMpls.EntityData.Children = types.NewOrderedMap()
    ldpMpls.EntityData.Leafs = types.NewOrderedMap()
    ldpMpls.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ldpMpls.TemplateName})

    ldpMpls.EntityData.YListKeys = []string {}

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

    monitorEnable.EntityData.Children = types.NewOrderedMap()
    monitorEnable.EntityData.Children.Append("ldp-mpls", types.YChild{"LdpMpls", &monitorEnable.LdpMpls})
    monitorEnable.EntityData.Children.Append("ospfv3-protocol", types.YChild{"Ospfv3Protocol", &monitorEnable.Ospfv3Protocol})
    monitorEnable.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &monitorEnable.GenericCounters})
    monitorEnable.EntityData.Children.Append("process", types.YChild{"Process", &monitorEnable.Process})
    monitorEnable.EntityData.Children.Append("basic-counters", types.YChild{"BasicCounters", &monitorEnable.BasicCounters})
    monitorEnable.EntityData.Children.Append("memory", types.YChild{"Memory", &monitorEnable.Memory})
    monitorEnable.EntityData.Children.Append("ospfv2-protocol", types.YChild{"Ospfv2Protocol", &monitorEnable.Ospfv2Protocol})
    monitorEnable.EntityData.Children.Append("cpu", types.YChild{"Cpu", &monitorEnable.Cpu})
    monitorEnable.EntityData.Children.Append("bgp", types.YChild{"Bgp", &monitorEnable.Bgp})
    monitorEnable.EntityData.Children.Append("data-rates", types.YChild{"DataRates", &monitorEnable.DataRates})
    monitorEnable.EntityData.Leafs = types.NewOrderedMap()

    monitorEnable.EntityData.YListKeys = []string {}

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

    ldpMpls.EntityData.Children = types.NewOrderedMap()
    ldpMpls.EntityData.Children.Append("sessions", types.YChild{"Sessions", &ldpMpls.Sessions})
    ldpMpls.EntityData.Leafs = types.NewOrderedMap()

    ldpMpls.EntityData.YListKeys = []string {}

    return &(ldpMpls.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions
// LDP session specification
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the LDP Session. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session.
    Session []*PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
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

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session
// IP address of the LDP Session
type PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the LDP Session. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Session interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (session *PerfMgmt_Enable_MonitorEnable_LdpMpls_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.Session, "session")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session", types.YLeaf{"Session", session.Session})
    session.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", session.TemplateName})

    session.EntityData.YListKeys = []string {"Session"}

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

    ospfv3Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Children.Append("ospf-instances", types.YChild{"OspfInstances", &ospfv3Protocol.OspfInstances})
    ospfv3Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Protocol.EntityData.YListKeys = []string {}

    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance.
    OspfInstance []*PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
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

    ospfInstances.EntityData.Children = types.NewOrderedMap()
    ospfInstances.EntityData.Children.Append("ospf-instance", types.YChild{"OspfInstance", nil})
    for i := range ospfInstances.OspfInstance {
        ospfInstances.EntityData.Children.Append(types.GetSegmentPath(ospfInstances.OspfInstance[i]), types.YChild{"OspfInstance", ospfInstances.OspfInstance[i]})
    }
    ospfInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfInstances.EntityData.YListKeys = []string {}

    return &(ospfInstances.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv3Protocol_OspfInstances_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfInstance.EntityData.ParentYangName = "ospf-instances"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + types.AddKeyToken(ospfInstance.InstanceName, "instance-name")
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstance.EntityData.Children = types.NewOrderedMap()
    ospfInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfInstance.InstanceName})
    ospfInstance.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfInstance.TemplateName})

    ospfInstance.EntityData.YListKeys = []string {"InstanceName"}

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

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &genericCounters.Interfaces})
    genericCounters.EntityData.Leafs = types.NewOrderedMap()

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface.
    Interface []*PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_GenericCounters_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", self.TemplateName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

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

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("process-nodes", types.YChild{"ProcessNodes", &process.ProcessNodes})
    process.EntityData.Leafs = types.NewOrderedMap()

    process.EntityData.YListKeys = []string {}

    return &(process.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode.
    ProcessNode []*PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
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

    processNodes.EntityData.Children = types.NewOrderedMap()
    processNodes.EntityData.Children.Append("process-node", types.YChild{"ProcessNode", nil})
    for i := range processNodes.ProcessNode {
        processNodes.EntityData.Children.Append(types.GetSegmentPath(processNodes.ProcessNode[i]), types.YChild{"ProcessNode", processNodes.ProcessNode[i]})
    }
    processNodes.EntityData.Leafs = types.NewOrderedMap()

    processNodes.EntityData.YListKeys = []string {}

    return &(processNodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode
// Node instance
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Process ID specification.
    Pids PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
}

func (processNode *PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode) GetEntityData() *types.CommonEntityData {
    processNode.EntityData.YFilter = processNode.YFilter
    processNode.EntityData.YangName = "process-node"
    processNode.EntityData.BundleName = "cisco_ios_xr"
    processNode.EntityData.ParentYangName = "process-nodes"
    processNode.EntityData.SegmentPath = "process-node" + types.AddKeyToken(processNode.NodeId, "node-id")
    processNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNode.EntityData.Children = types.NewOrderedMap()
    processNode.EntityData.Children.Append("pids", types.YChild{"Pids", &processNode.Pids})
    processNode.EntityData.Leafs = types.NewOrderedMap()
    processNode.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", processNode.NodeId})

    processNode.EntityData.YListKeys = []string {"NodeId"}

    return &(processNode.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids
// Process ID specification
type PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify an existing template for data collection. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid.
    Pid []*PerfMgmt_Enable_MonitorEnable_Process_ProcessNodes_ProcessNode_Pids_Pid
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

    pids.EntityData.Children = types.NewOrderedMap()
    pids.EntityData.Children.Append("pid", types.YChild{"Pid", nil})
    for i := range pids.Pid {
        pids.EntityData.Children.Append(types.GetSegmentPath(pids.Pid[i]), types.YChild{"Pid", pids.Pid[i]})
    }
    pids.EntityData.Leafs = types.NewOrderedMap()

    pids.EntityData.YListKeys = []string {}

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
    pid.EntityData.SegmentPath = "pid" + types.AddKeyToken(pid.Pid, "pid")
    pid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pid.EntityData.Children = types.NewOrderedMap()
    pid.EntityData.Leafs = types.NewOrderedMap()
    pid.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", pid.Pid})
    pid.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", pid.TemplateName})

    pid.EntityData.YListKeys = []string {"Pid"}

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

    basicCounters.EntityData.Children = types.NewOrderedMap()
    basicCounters.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &basicCounters.Interfaces})
    basicCounters.EntityData.Leafs = types.NewOrderedMap()

    basicCounters.EntityData.YListKeys = []string {}

    return &(basicCounters.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface.
    Interface []*PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_BasicCounters_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", self.TemplateName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

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

    memory.EntityData.Children = types.NewOrderedMap()
    memory.EntityData.Children.Append("nodes", types.YChild{"Nodes", &memory.Nodes})
    memory.EntityData.Leafs = types.NewOrderedMap()

    memory.EntityData.YListKeys = []string {}

    return &(memory.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node.
    Node []*PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Memory_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    ospfv2Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Children.Append("ospf-instances", types.YChild{"OspfInstances", &ospfv2Protocol.OspfInstances})
    ospfv2Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Protocol.EntityData.YListKeys = []string {}

    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances
// Monitor an instance
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance.
    OspfInstance []*PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
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

    ospfInstances.EntityData.Children = types.NewOrderedMap()
    ospfInstances.EntityData.Children.Append("ospf-instance", types.YChild{"OspfInstance", nil})
    for i := range ospfInstances.OspfInstance {
        ospfInstances.EntityData.Children.Append(types.GetSegmentPath(ospfInstances.OspfInstance[i]), types.YChild{"OspfInstance", ospfInstances.OspfInstance[i]})
    }
    ospfInstances.EntityData.Leafs = types.NewOrderedMap()

    ospfInstances.EntityData.YListKeys = []string {}

    return &(ospfInstances.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance
// Instance being monitored
type PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF Instance Name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InstanceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (ospfInstance *PerfMgmt_Enable_MonitorEnable_Ospfv2Protocol_OspfInstances_OspfInstance) GetEntityData() *types.CommonEntityData {
    ospfInstance.EntityData.YFilter = ospfInstance.YFilter
    ospfInstance.EntityData.YangName = "ospf-instance"
    ospfInstance.EntityData.BundleName = "cisco_ios_xr"
    ospfInstance.EntityData.ParentYangName = "ospf-instances"
    ospfInstance.EntityData.SegmentPath = "ospf-instance" + types.AddKeyToken(ospfInstance.InstanceName, "instance-name")
    ospfInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInstance.EntityData.Children = types.NewOrderedMap()
    ospfInstance.EntityData.Leafs = types.NewOrderedMap()
    ospfInstance.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", ospfInstance.InstanceName})
    ospfInstance.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfInstance.TemplateName})

    ospfInstance.EntityData.YListKeys = []string {"InstanceName"}

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

    cpu.EntityData.Children = types.NewOrderedMap()
    cpu.EntityData.Children.Append("nodes", types.YChild{"Nodes", &cpu.Nodes})
    cpu.EntityData.Leafs = types.NewOrderedMap()

    cpu.EntityData.YListKeys = []string {}

    return &(cpu.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes
// Node specification
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node instance. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node.
    Node []*PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node
// Node instance
type PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (node *PerfMgmt_Enable_MonitorEnable_Cpu_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeId, "node-id")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", node.NodeId})
    node.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", node.TemplateName})

    node.EntityData.YListKeys = []string {"NodeId"}

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

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &bgp.Neighbors})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors
// Monitor BGP protocol for a BGP peer
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor being monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor.
    Neighbor []*PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
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

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor
// Neighbor being monitored
type PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the Neighbor. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (neighbor *PerfMgmt_Enable_MonitorEnable_Bgp_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.PeerAddress, "peer-address")
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", neighbor.PeerAddress})
    neighbor.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", neighbor.TemplateName})

    neighbor.EntityData.YListKeys = []string {"PeerAddress"}

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

    dataRates.EntityData.Children = types.NewOrderedMap()
    dataRates.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &dataRates.Interfaces})
    dataRates.EntityData.Leafs = types.NewOrderedMap()

    dataRates.EntityData.YListKeys = []string {}

    return &(dataRates.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces
// Monitor an Interface
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being Monitored. The type is slice of
    // PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface.
    Interface []*PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface
// Interface being Monitored
type PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Template name. The type is string.
    TemplateName interface{}
}

func (self *PerfMgmt_Enable_MonitorEnable_DataRates_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", self.TemplateName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// PerfMgmt_RegExpGroups
// Configure regular expression group
type PerfMgmt_RegExpGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify regular expression group name. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup.
    RegExpGroup []*PerfMgmt_RegExpGroups_RegExpGroup
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

    regExpGroups.EntityData.Children = types.NewOrderedMap()
    regExpGroups.EntityData.Children.Append("reg-exp-group", types.YChild{"RegExpGroup", nil})
    for i := range regExpGroups.RegExpGroup {
        regExpGroups.EntityData.Children.Append(types.GetSegmentPath(regExpGroups.RegExpGroup[i]), types.YChild{"RegExpGroup", regExpGroups.RegExpGroup[i]})
    }
    regExpGroups.EntityData.Leafs = types.NewOrderedMap()

    regExpGroups.EntityData.YListKeys = []string {}

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
    regExpGroup.EntityData.SegmentPath = "reg-exp-group" + types.AddKeyToken(regExpGroup.RegExpGroupName, "reg-exp-group-name")
    regExpGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExpGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExpGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExpGroup.EntityData.Children = types.NewOrderedMap()
    regExpGroup.EntityData.Children.Append("reg-exps", types.YChild{"RegExps", &regExpGroup.RegExps})
    regExpGroup.EntityData.Leafs = types.NewOrderedMap()
    regExpGroup.EntityData.Leafs.Append("reg-exp-group-name", types.YLeaf{"RegExpGroupName", regExpGroup.RegExpGroupName})

    regExpGroup.EntityData.YListKeys = []string {"RegExpGroupName"}

    return &(regExpGroup.EntityData)
}

// PerfMgmt_RegExpGroups_RegExpGroup_RegExps
// Configure regular expression
type PerfMgmt_RegExpGroups_RegExpGroup_RegExps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify regular expression index number. The type is slice of
    // PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp.
    RegExp []*PerfMgmt_RegExpGroups_RegExpGroup_RegExps_RegExp
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

    regExps.EntityData.Children = types.NewOrderedMap()
    regExps.EntityData.Children.Append("reg-exp", types.YChild{"RegExp", nil})
    for i := range regExps.RegExp {
        regExps.EntityData.Children.Append(types.GetSegmentPath(regExps.RegExp[i]), types.YChild{"RegExp", regExps.RegExp[i]})
    }
    regExps.EntityData.Leafs = types.NewOrderedMap()

    regExps.EntityData.YListKeys = []string {}

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
    regExp.EntityData.SegmentPath = "reg-exp" + types.AddKeyToken(regExp.RegExpIndex, "reg-exp-index")
    regExp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regExp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regExp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regExp.EntityData.Children = types.NewOrderedMap()
    regExp.EntityData.Leafs = types.NewOrderedMap()
    regExp.EntityData.Leafs.Append("reg-exp-index", types.YLeaf{"RegExpIndex", regExp.RegExpIndex})
    regExp.EntityData.Leafs.Append("reg-exp-string", types.YLeaf{"RegExpString", regExp.RegExpString})

    regExp.EntityData.YListKeys = []string {"RegExpIndex"}

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

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("generic-counter-interface", types.YChild{"GenericCounterInterface", &threshold.GenericCounterInterface})
    threshold.EntityData.Children.Append("ldp-mpls", types.YChild{"LdpMpls", &threshold.LdpMpls})
    threshold.EntityData.Children.Append("basic-counter-interface", types.YChild{"BasicCounterInterface", &threshold.BasicCounterInterface})
    threshold.EntityData.Children.Append("bgp", types.YChild{"Bgp", &threshold.Bgp})
    threshold.EntityData.Children.Append("ospfv2-protocol", types.YChild{"Ospfv2Protocol", &threshold.Ospfv2Protocol})
    threshold.EntityData.Children.Append("cpu-node", types.YChild{"CpuNode", &threshold.CpuNode})
    threshold.EntityData.Children.Append("data-rate-interface", types.YChild{"DataRateInterface", &threshold.DataRateInterface})
    threshold.EntityData.Children.Append("process-node", types.YChild{"ProcessNode", &threshold.ProcessNode})
    threshold.EntityData.Children.Append("memory-node", types.YChild{"MemoryNode", &threshold.MemoryNode})
    threshold.EntityData.Children.Append("ospfv3-protocol", types.YChild{"Ospfv3Protocol", &threshold.Ospfv3Protocol})
    threshold.EntityData.Leafs = types.NewOrderedMap()

    threshold.EntityData.YListKeys = []string {}

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

    genericCounterInterface.EntityData.Children = types.NewOrderedMap()
    genericCounterInterface.EntityData.Children.Append("generic-counter-interface-templates", types.YChild{"GenericCounterInterfaceTemplates", &genericCounterInterface.GenericCounterInterfaceTemplates})
    genericCounterInterface.EntityData.Leafs = types.NewOrderedMap()

    genericCounterInterface.EntityData.YListKeys = []string {}

    return &(genericCounterInterface.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates
// Interface Generic Counter threshold templates
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Generic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate.
    GenericCounterInterfaceTemplate []*PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
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

    genericCounterInterfaceTemplates.EntityData.Children = types.NewOrderedMap()
    genericCounterInterfaceTemplates.EntityData.Children.Append("generic-counter-interface-template", types.YChild{"GenericCounterInterfaceTemplate", nil})
    for i := range genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate {
        genericCounterInterfaceTemplates.EntityData.Children.Append(types.GetSegmentPath(genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i]), types.YChild{"GenericCounterInterfaceTemplate", genericCounterInterfaceTemplates.GenericCounterInterfaceTemplate[i]})
    }
    genericCounterInterfaceTemplates.EntityData.Leafs = types.NewOrderedMap()

    genericCounterInterfaceTemplates.EntityData.YListKeys = []string {}

    return &(genericCounterInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate
// Interface Generic Counter threshold template
// instance
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate struct {
    EntityData types.CommonEntityData
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

func (genericCounterInterfaceTemplate *PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    genericCounterInterfaceTemplate.EntityData.YFilter = genericCounterInterfaceTemplate.YFilter
    genericCounterInterfaceTemplate.EntityData.YangName = "generic-counter-interface-template"
    genericCounterInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    genericCounterInterfaceTemplate.EntityData.ParentYangName = "generic-counter-interface-templates"
    genericCounterInterfaceTemplate.EntityData.SegmentPath = "generic-counter-interface-template" + types.AddKeyToken(genericCounterInterfaceTemplate.TemplateName, "template-name")
    genericCounterInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounterInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounterInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounterInterfaceTemplate.EntityData.Children = types.NewOrderedMap()
    genericCounterInterfaceTemplate.EntityData.Children.Append("in-octets", types.YChild{"InOctets", &genericCounterInterfaceTemplate.InOctets})
    genericCounterInterfaceTemplate.EntityData.Children.Append("in-ucast-pkts", types.YChild{"InUcastPkts", &genericCounterInterfaceTemplate.InUcastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("out-ucast-pkts", types.YChild{"OutUcastPkts", &genericCounterInterfaceTemplate.OutUcastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("out-broadcast-pkts", types.YChild{"OutBroadcastPkts", &genericCounterInterfaceTemplate.OutBroadcastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("out-multicast-pkts", types.YChild{"OutMulticastPkts", &genericCounterInterfaceTemplate.OutMulticastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-overrun", types.YChild{"InputOverrun", &genericCounterInterfaceTemplate.InputOverrun})
    genericCounterInterfaceTemplate.EntityData.Children.Append("out-octets", types.YChild{"OutOctets", &genericCounterInterfaceTemplate.OutOctets})
    genericCounterInterfaceTemplate.EntityData.Children.Append("output-underrun", types.YChild{"OutputUnderrun", &genericCounterInterfaceTemplate.OutputUnderrun})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-total-errors", types.YChild{"InputTotalErrors", &genericCounterInterfaceTemplate.InputTotalErrors})
    genericCounterInterfaceTemplate.EntityData.Children.Append("output-total-drops", types.YChild{"OutputTotalDrops", &genericCounterInterfaceTemplate.OutputTotalDrops})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-crc", types.YChild{"InputCrc", &genericCounterInterfaceTemplate.InputCrc})
    genericCounterInterfaceTemplate.EntityData.Children.Append("in-broadcast-pkts", types.YChild{"InBroadcastPkts", &genericCounterInterfaceTemplate.InBroadcastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("in-multicast-pkts", types.YChild{"InMulticastPkts", &genericCounterInterfaceTemplate.InMulticastPkts})
    genericCounterInterfaceTemplate.EntityData.Children.Append("out-packets", types.YChild{"OutPackets", &genericCounterInterfaceTemplate.OutPackets})
    genericCounterInterfaceTemplate.EntityData.Children.Append("output-total-errors", types.YChild{"OutputTotalErrors", &genericCounterInterfaceTemplate.OutputTotalErrors})
    genericCounterInterfaceTemplate.EntityData.Children.Append("in-packets", types.YChild{"InPackets", &genericCounterInterfaceTemplate.InPackets})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-unknown-proto", types.YChild{"InputUnknownProto", &genericCounterInterfaceTemplate.InputUnknownProto})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-queue-drops", types.YChild{"InputQueueDrops", &genericCounterInterfaceTemplate.InputQueueDrops})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-total-drops", types.YChild{"InputTotalDrops", &genericCounterInterfaceTemplate.InputTotalDrops})
    genericCounterInterfaceTemplate.EntityData.Children.Append("input-frame", types.YChild{"InputFrame", &genericCounterInterfaceTemplate.InputFrame})
    genericCounterInterfaceTemplate.EntityData.Leafs = types.NewOrderedMap()
    genericCounterInterfaceTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", genericCounterInterfaceTemplate.TemplateName})
    genericCounterInterfaceTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", genericCounterInterfaceTemplate.SampleInterval})
    genericCounterInterfaceTemplate.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", genericCounterInterfaceTemplate.RegExpGroup})
    genericCounterInterfaceTemplate.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", genericCounterInterfaceTemplate.VrfGroup})

    genericCounterInterfaceTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(genericCounterInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InOctets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inOctets.EntityData.Children = types.NewOrderedMap()
    inOctets.EntityData.Leafs = types.NewOrderedMap()
    inOctets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inOctets.Operator})
    inOctets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inOctets.Value})
    inOctets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inOctets.EndRangeValue})
    inOctets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inOctets.Percent})
    inOctets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inOctets.RearmType})
    inOctets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inOctets.RearmWindow})

    inOctets.EntityData.YListKeys = []string {}

    return &(inOctets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts
// Number of inbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InUcastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inUcastPkts.EntityData.Children = types.NewOrderedMap()
    inUcastPkts.EntityData.Leafs = types.NewOrderedMap()
    inUcastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inUcastPkts.Operator})
    inUcastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", inUcastPkts.Value})
    inUcastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inUcastPkts.EndRangeValue})
    inUcastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inUcastPkts.Percent})
    inUcastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inUcastPkts.RearmType})
    inUcastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inUcastPkts.RearmWindow})

    inUcastPkts.EntityData.YListKeys = []string {}

    return &(inUcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts
// Number of outbound unicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutUcastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outUcastPkts.EntityData.Children = types.NewOrderedMap()
    outUcastPkts.EntityData.Leafs = types.NewOrderedMap()
    outUcastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outUcastPkts.Operator})
    outUcastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", outUcastPkts.Value})
    outUcastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outUcastPkts.EndRangeValue})
    outUcastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outUcastPkts.Percent})
    outUcastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outUcastPkts.RearmType})
    outUcastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outUcastPkts.RearmWindow})

    outUcastPkts.EntityData.YListKeys = []string {}

    return &(outUcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts
// Number of outbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutBroadcastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outBroadcastPkts.EntityData.Children = types.NewOrderedMap()
    outBroadcastPkts.EntityData.Leafs = types.NewOrderedMap()
    outBroadcastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outBroadcastPkts.Operator})
    outBroadcastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", outBroadcastPkts.Value})
    outBroadcastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outBroadcastPkts.EndRangeValue})
    outBroadcastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outBroadcastPkts.Percent})
    outBroadcastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outBroadcastPkts.RearmType})
    outBroadcastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outBroadcastPkts.RearmWindow})

    outBroadcastPkts.EntityData.YListKeys = []string {}

    return &(outBroadcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts
// Number of outbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutMulticastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outMulticastPkts.EntityData.Children = types.NewOrderedMap()
    outMulticastPkts.EntityData.Leafs = types.NewOrderedMap()
    outMulticastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outMulticastPkts.Operator})
    outMulticastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", outMulticastPkts.Value})
    outMulticastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outMulticastPkts.EndRangeValue})
    outMulticastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outMulticastPkts.Percent})
    outMulticastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outMulticastPkts.RearmType})
    outMulticastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outMulticastPkts.RearmWindow})

    outMulticastPkts.EntityData.YListKeys = []string {}

    return &(outMulticastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun
// Number of inbound packets with overrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputOverrun struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputOverrun.EntityData.Children = types.NewOrderedMap()
    inputOverrun.EntityData.Leafs = types.NewOrderedMap()
    inputOverrun.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputOverrun.Operator})
    inputOverrun.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputOverrun.Value})
    inputOverrun.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputOverrun.EndRangeValue})
    inputOverrun.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputOverrun.Percent})
    inputOverrun.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputOverrun.RearmType})
    inputOverrun.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputOverrun.RearmWindow})

    inputOverrun.EntityData.YListKeys = []string {}

    return &(inputOverrun.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutOctets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outOctets.EntityData.Children = types.NewOrderedMap()
    outOctets.EntityData.Leafs = types.NewOrderedMap()
    outOctets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outOctets.Operator})
    outOctets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outOctets.Value})
    outOctets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outOctets.EndRangeValue})
    outOctets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outOctets.Percent})
    outOctets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outOctets.RearmType})
    outOctets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outOctets.RearmWindow})

    outOctets.EntityData.YListKeys = []string {}

    return &(outOctets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun
// Number of outbound packets with underrun
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputUnderrun struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputUnderrun.EntityData.Children = types.NewOrderedMap()
    outputUnderrun.EntityData.Leafs = types.NewOrderedMap()
    outputUnderrun.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputUnderrun.Operator})
    outputUnderrun.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputUnderrun.Value})
    outputUnderrun.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputUnderrun.EndRangeValue})
    outputUnderrun.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputUnderrun.Percent})
    outputUnderrun.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputUnderrun.RearmType})
    outputUnderrun.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputUnderrun.RearmWindow})

    outputUnderrun.EntityData.YListKeys = []string {}

    return &(outputUnderrun.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputTotalErrors.EntityData.Children = types.NewOrderedMap()
    inputTotalErrors.EntityData.Leafs = types.NewOrderedMap()
    inputTotalErrors.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputTotalErrors.Operator})
    inputTotalErrors.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputTotalErrors.Value})
    inputTotalErrors.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputTotalErrors.EndRangeValue})
    inputTotalErrors.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputTotalErrors.Percent})
    inputTotalErrors.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputTotalErrors.RearmType})
    inputTotalErrors.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputTotalErrors.RearmWindow})

    inputTotalErrors.EntityData.YListKeys = []string {}

    return &(inputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputTotalDrops.EntityData.Children = types.NewOrderedMap()
    outputTotalDrops.EntityData.Leafs = types.NewOrderedMap()
    outputTotalDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputTotalDrops.Operator})
    outputTotalDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputTotalDrops.Value})
    outputTotalDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputTotalDrops.EndRangeValue})
    outputTotalDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputTotalDrops.Percent})
    outputTotalDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputTotalDrops.RearmType})
    outputTotalDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputTotalDrops.RearmWindow})

    outputTotalDrops.EntityData.YListKeys = []string {}

    return &(outputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc
// Number of inbound packets discarded with
// incorrect CRC
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputCrc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputCrc.EntityData.Children = types.NewOrderedMap()
    inputCrc.EntityData.Leafs = types.NewOrderedMap()
    inputCrc.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputCrc.Operator})
    inputCrc.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputCrc.Value})
    inputCrc.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputCrc.EndRangeValue})
    inputCrc.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputCrc.Percent})
    inputCrc.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputCrc.RearmType})
    inputCrc.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputCrc.RearmWindow})

    inputCrc.EntityData.YListKeys = []string {}

    return &(inputCrc.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts
// Number of inbound broadcast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InBroadcastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inBroadcastPkts.EntityData.Children = types.NewOrderedMap()
    inBroadcastPkts.EntityData.Leafs = types.NewOrderedMap()
    inBroadcastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inBroadcastPkts.Operator})
    inBroadcastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", inBroadcastPkts.Value})
    inBroadcastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inBroadcastPkts.EndRangeValue})
    inBroadcastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inBroadcastPkts.Percent})
    inBroadcastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inBroadcastPkts.RearmType})
    inBroadcastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inBroadcastPkts.RearmWindow})

    inBroadcastPkts.EntityData.YListKeys = []string {}

    return &(inBroadcastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts
// Number of inbound multicast packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InMulticastPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inMulticastPkts.EntityData.Children = types.NewOrderedMap()
    inMulticastPkts.EntityData.Leafs = types.NewOrderedMap()
    inMulticastPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inMulticastPkts.Operator})
    inMulticastPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", inMulticastPkts.Value})
    inMulticastPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inMulticastPkts.EndRangeValue})
    inMulticastPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inMulticastPkts.Percent})
    inMulticastPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inMulticastPkts.RearmType})
    inMulticastPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inMulticastPkts.RearmWindow})

    inMulticastPkts.EntityData.YListKeys = []string {}

    return &(inMulticastPkts.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outPackets.EntityData.Children = types.NewOrderedMap()
    outPackets.EntityData.Leafs = types.NewOrderedMap()
    outPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outPackets.Operator})
    outPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outPackets.Value})
    outPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outPackets.EndRangeValue})
    outPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outPackets.Percent})
    outPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outPackets.RearmType})
    outPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outPackets.RearmWindow})

    outPackets.EntityData.YListKeys = []string {}

    return &(outPackets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_OutputTotalErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputTotalErrors.EntityData.Children = types.NewOrderedMap()
    outputTotalErrors.EntityData.Leafs = types.NewOrderedMap()
    outputTotalErrors.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputTotalErrors.Operator})
    outputTotalErrors.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputTotalErrors.Value})
    outputTotalErrors.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputTotalErrors.EndRangeValue})
    outputTotalErrors.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputTotalErrors.Percent})
    outputTotalErrors.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputTotalErrors.RearmType})
    outputTotalErrors.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputTotalErrors.RearmWindow})

    outputTotalErrors.EntityData.YListKeys = []string {}

    return &(outputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inPackets.EntityData.Children = types.NewOrderedMap()
    inPackets.EntityData.Leafs = types.NewOrderedMap()
    inPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inPackets.Operator})
    inPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inPackets.Value})
    inPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inPackets.EndRangeValue})
    inPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inPackets.Percent})
    inPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inPackets.RearmType})
    inPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inPackets.RearmWindow})

    inPackets.EntityData.YListKeys = []string {}

    return &(inPackets.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto
// Number of inbound packets discarded with
// unknown protocol
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputUnknownProto struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputUnknownProto.EntityData.Children = types.NewOrderedMap()
    inputUnknownProto.EntityData.Leafs = types.NewOrderedMap()
    inputUnknownProto.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputUnknownProto.Operator})
    inputUnknownProto.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputUnknownProto.Value})
    inputUnknownProto.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputUnknownProto.EndRangeValue})
    inputUnknownProto.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputUnknownProto.Percent})
    inputUnknownProto.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputUnknownProto.RearmType})
    inputUnknownProto.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputUnknownProto.RearmWindow})

    inputUnknownProto.EntityData.YListKeys = []string {}

    return &(inputUnknownProto.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputQueueDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputQueueDrops.EntityData.Children = types.NewOrderedMap()
    inputQueueDrops.EntityData.Leafs = types.NewOrderedMap()
    inputQueueDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputQueueDrops.Operator})
    inputQueueDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputQueueDrops.Value})
    inputQueueDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputQueueDrops.EndRangeValue})
    inputQueueDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputQueueDrops.Percent})
    inputQueueDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputQueueDrops.RearmType})
    inputQueueDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputQueueDrops.RearmWindow})

    inputQueueDrops.EntityData.YListKeys = []string {}

    return &(inputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputTotalDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputTotalDrops.EntityData.Children = types.NewOrderedMap()
    inputTotalDrops.EntityData.Leafs = types.NewOrderedMap()
    inputTotalDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputTotalDrops.Operator})
    inputTotalDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputTotalDrops.Value})
    inputTotalDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputTotalDrops.EndRangeValue})
    inputTotalDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputTotalDrops.Percent})
    inputTotalDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputTotalDrops.RearmType})
    inputTotalDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputTotalDrops.RearmWindow})

    inputTotalDrops.EntityData.YListKeys = []string {}

    return &(inputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame
// Number of inbound packets with framing
// errors
// This type is a presence type.
type PerfMgmt_Threshold_GenericCounterInterface_GenericCounterInterfaceTemplates_GenericCounterInterfaceTemplate_InputFrame struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputFrame.EntityData.Children = types.NewOrderedMap()
    inputFrame.EntityData.Leafs = types.NewOrderedMap()
    inputFrame.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputFrame.Operator})
    inputFrame.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputFrame.Value})
    inputFrame.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputFrame.EndRangeValue})
    inputFrame.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputFrame.Percent})
    inputFrame.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputFrame.RearmType})
    inputFrame.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputFrame.RearmWindow})

    inputFrame.EntityData.YListKeys = []string {}

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

    ldpMpls.EntityData.Children = types.NewOrderedMap()
    ldpMpls.EntityData.Children.Append("ldp-mpls-templates", types.YChild{"LdpMplsTemplates", &ldpMpls.LdpMplsTemplates})
    ldpMpls.EntityData.Leafs = types.NewOrderedMap()

    ldpMpls.EntityData.YListKeys = []string {}

    return &(ldpMpls.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates
// MPLS LDP threshold templates
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate.
    LdpMplsTemplate []*PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
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

    ldpMplsTemplates.EntityData.Children = types.NewOrderedMap()
    ldpMplsTemplates.EntityData.Children.Append("ldp-mpls-template", types.YChild{"LdpMplsTemplate", nil})
    for i := range ldpMplsTemplates.LdpMplsTemplate {
        ldpMplsTemplates.EntityData.Children.Append(types.GetSegmentPath(ldpMplsTemplates.LdpMplsTemplate[i]), types.YChild{"LdpMplsTemplate", ldpMplsTemplates.LdpMplsTemplate[i]})
    }
    ldpMplsTemplates.EntityData.Leafs = types.NewOrderedMap()

    ldpMplsTemplates.EntityData.YListKeys = []string {}

    return &(ldpMplsTemplates.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate
// MPLS LDP threshold template instance
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate struct {
    EntityData types.CommonEntityData
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

func (ldpMplsTemplate *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate) GetEntityData() *types.CommonEntityData {
    ldpMplsTemplate.EntityData.YFilter = ldpMplsTemplate.YFilter
    ldpMplsTemplate.EntityData.YangName = "ldp-mpls-template"
    ldpMplsTemplate.EntityData.BundleName = "cisco_ios_xr"
    ldpMplsTemplate.EntityData.ParentYangName = "ldp-mpls-templates"
    ldpMplsTemplate.EntityData.SegmentPath = "ldp-mpls-template" + types.AddKeyToken(ldpMplsTemplate.TemplateName, "template-name")
    ldpMplsTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldpMplsTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldpMplsTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldpMplsTemplate.EntityData.Children = types.NewOrderedMap()
    ldpMplsTemplate.EntityData.Children.Append("address-withdraw-msgs-rcvd", types.YChild{"AddressWithdrawMsgsRcvd", &ldpMplsTemplate.AddressWithdrawMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("label-withdraw-msgs-rcvd", types.YChild{"LabelWithdrawMsgsRcvd", &ldpMplsTemplate.LabelWithdrawMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("address-withdraw-msgs-sent", types.YChild{"AddressWithdrawMsgsSent", &ldpMplsTemplate.AddressWithdrawMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("label-withdraw-msgs-sent", types.YChild{"LabelWithdrawMsgsSent", &ldpMplsTemplate.LabelWithdrawMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("notification-msgs-rcvd", types.YChild{"NotificationMsgsRcvd", &ldpMplsTemplate.NotificationMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("total-msgs-rcvd", types.YChild{"TotalMsgsRcvd", &ldpMplsTemplate.TotalMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("notification-msgs-sent", types.YChild{"NotificationMsgsSent", &ldpMplsTemplate.NotificationMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("total-msgs-sent", types.YChild{"TotalMsgsSent", &ldpMplsTemplate.TotalMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("label-release-msgs-rcvd", types.YChild{"LabelReleaseMsgsRcvd", &ldpMplsTemplate.LabelReleaseMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("init-msgs-rcvd", types.YChild{"InitMsgsRcvd", &ldpMplsTemplate.InitMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("label-release-msgs-sent", types.YChild{"LabelReleaseMsgsSent", &ldpMplsTemplate.LabelReleaseMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("init-msgs-sent", types.YChild{"InitMsgsSent", &ldpMplsTemplate.InitMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("label-mapping-msgs-rcvd", types.YChild{"LabelMappingMsgsRcvd", &ldpMplsTemplate.LabelMappingMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("keepalive-msgs-rcvd", types.YChild{"KeepaliveMsgsRcvd", &ldpMplsTemplate.KeepaliveMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("label-mapping-msgs-sent", types.YChild{"LabelMappingMsgsSent", &ldpMplsTemplate.LabelMappingMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("keepalive-msgs-sent", types.YChild{"KeepaliveMsgsSent", &ldpMplsTemplate.KeepaliveMsgsSent})
    ldpMplsTemplate.EntityData.Children.Append("address-msgs-rcvd", types.YChild{"AddressMsgsRcvd", &ldpMplsTemplate.AddressMsgsRcvd})
    ldpMplsTemplate.EntityData.Children.Append("address-msgs-sent", types.YChild{"AddressMsgsSent", &ldpMplsTemplate.AddressMsgsSent})
    ldpMplsTemplate.EntityData.Leafs = types.NewOrderedMap()
    ldpMplsTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ldpMplsTemplate.TemplateName})
    ldpMplsTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", ldpMplsTemplate.SampleInterval})

    ldpMplsTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(ldpMplsTemplate.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd
// Number of Address Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (addressWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsRcvd) GetEntityData() *types.CommonEntityData {
    addressWithdrawMsgsRcvd.EntityData.YFilter = addressWithdrawMsgsRcvd.YFilter
    addressWithdrawMsgsRcvd.EntityData.YangName = "address-withdraw-msgs-rcvd"
    addressWithdrawMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    addressWithdrawMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    addressWithdrawMsgsRcvd.EntityData.SegmentPath = "address-withdraw-msgs-rcvd"
    addressWithdrawMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressWithdrawMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressWithdrawMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressWithdrawMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    addressWithdrawMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", addressWithdrawMsgsRcvd.Operator})
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", addressWithdrawMsgsRcvd.Value})
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", addressWithdrawMsgsRcvd.EndRangeValue})
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", addressWithdrawMsgsRcvd.Percent})
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", addressWithdrawMsgsRcvd.RearmType})
    addressWithdrawMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", addressWithdrawMsgsRcvd.RearmWindow})

    addressWithdrawMsgsRcvd.EntityData.YListKeys = []string {}

    return &(addressWithdrawMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd
// Number of Label Withdraw messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelWithdrawMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelWithdrawMsgsRcvd.EntityData.YFilter = labelWithdrawMsgsRcvd.YFilter
    labelWithdrawMsgsRcvd.EntityData.YangName = "label-withdraw-msgs-rcvd"
    labelWithdrawMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelWithdrawMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelWithdrawMsgsRcvd.EntityData.SegmentPath = "label-withdraw-msgs-rcvd"
    labelWithdrawMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelWithdrawMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelWithdrawMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelWithdrawMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    labelWithdrawMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelWithdrawMsgsRcvd.Operator})
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelWithdrawMsgsRcvd.Value})
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelWithdrawMsgsRcvd.EndRangeValue})
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelWithdrawMsgsRcvd.Percent})
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelWithdrawMsgsRcvd.RearmType})
    labelWithdrawMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelWithdrawMsgsRcvd.RearmWindow})

    labelWithdrawMsgsRcvd.EntityData.YListKeys = []string {}

    return &(labelWithdrawMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent
// Number of Address Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (addressWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressWithdrawMsgsSent) GetEntityData() *types.CommonEntityData {
    addressWithdrawMsgsSent.EntityData.YFilter = addressWithdrawMsgsSent.YFilter
    addressWithdrawMsgsSent.EntityData.YangName = "address-withdraw-msgs-sent"
    addressWithdrawMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    addressWithdrawMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    addressWithdrawMsgsSent.EntityData.SegmentPath = "address-withdraw-msgs-sent"
    addressWithdrawMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressWithdrawMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressWithdrawMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressWithdrawMsgsSent.EntityData.Children = types.NewOrderedMap()
    addressWithdrawMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    addressWithdrawMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", addressWithdrawMsgsSent.Operator})
    addressWithdrawMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", addressWithdrawMsgsSent.Value})
    addressWithdrawMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", addressWithdrawMsgsSent.EndRangeValue})
    addressWithdrawMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", addressWithdrawMsgsSent.Percent})
    addressWithdrawMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", addressWithdrawMsgsSent.RearmType})
    addressWithdrawMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", addressWithdrawMsgsSent.RearmWindow})

    addressWithdrawMsgsSent.EntityData.YListKeys = []string {}

    return &(addressWithdrawMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent
// Number of Label Withdraw messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelWithdrawMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelWithdrawMsgsSent) GetEntityData() *types.CommonEntityData {
    labelWithdrawMsgsSent.EntityData.YFilter = labelWithdrawMsgsSent.YFilter
    labelWithdrawMsgsSent.EntityData.YangName = "label-withdraw-msgs-sent"
    labelWithdrawMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelWithdrawMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelWithdrawMsgsSent.EntityData.SegmentPath = "label-withdraw-msgs-sent"
    labelWithdrawMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelWithdrawMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelWithdrawMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelWithdrawMsgsSent.EntityData.Children = types.NewOrderedMap()
    labelWithdrawMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    labelWithdrawMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelWithdrawMsgsSent.Operator})
    labelWithdrawMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelWithdrawMsgsSent.Value})
    labelWithdrawMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelWithdrawMsgsSent.EndRangeValue})
    labelWithdrawMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelWithdrawMsgsSent.Percent})
    labelWithdrawMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelWithdrawMsgsSent.RearmType})
    labelWithdrawMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelWithdrawMsgsSent.RearmWindow})

    labelWithdrawMsgsSent.EntityData.YListKeys = []string {}

    return &(labelWithdrawMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd
// Number of Notification messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (notificationMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsRcvd) GetEntityData() *types.CommonEntityData {
    notificationMsgsRcvd.EntityData.YFilter = notificationMsgsRcvd.YFilter
    notificationMsgsRcvd.EntityData.YangName = "notification-msgs-rcvd"
    notificationMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    notificationMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    notificationMsgsRcvd.EntityData.SegmentPath = "notification-msgs-rcvd"
    notificationMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    notificationMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    notificationMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", notificationMsgsRcvd.Operator})
    notificationMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", notificationMsgsRcvd.Value})
    notificationMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", notificationMsgsRcvd.EndRangeValue})
    notificationMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", notificationMsgsRcvd.Percent})
    notificationMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", notificationMsgsRcvd.RearmType})
    notificationMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", notificationMsgsRcvd.RearmWindow})

    notificationMsgsRcvd.EntityData.YListKeys = []string {}

    return &(notificationMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd
// Total number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    totalMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    totalMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    totalMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", totalMsgsRcvd.Operator})
    totalMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", totalMsgsRcvd.Value})
    totalMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", totalMsgsRcvd.EndRangeValue})
    totalMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", totalMsgsRcvd.Percent})
    totalMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", totalMsgsRcvd.RearmType})
    totalMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", totalMsgsRcvd.RearmWindow})

    totalMsgsRcvd.EntityData.YListKeys = []string {}

    return &(totalMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent
// Number of Notification messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (notificationMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_NotificationMsgsSent) GetEntityData() *types.CommonEntityData {
    notificationMsgsSent.EntityData.YFilter = notificationMsgsSent.YFilter
    notificationMsgsSent.EntityData.YangName = "notification-msgs-sent"
    notificationMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    notificationMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    notificationMsgsSent.EntityData.SegmentPath = "notification-msgs-sent"
    notificationMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationMsgsSent.EntityData.Children = types.NewOrderedMap()
    notificationMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    notificationMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", notificationMsgsSent.Operator})
    notificationMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", notificationMsgsSent.Value})
    notificationMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", notificationMsgsSent.EndRangeValue})
    notificationMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", notificationMsgsSent.Percent})
    notificationMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", notificationMsgsSent.RearmType})
    notificationMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", notificationMsgsSent.RearmWindow})

    notificationMsgsSent.EntityData.YListKeys = []string {}

    return &(notificationMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent
// Total number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (totalMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_TotalMsgsSent) GetEntityData() *types.CommonEntityData {
    totalMsgsSent.EntityData.YFilter = totalMsgsSent.YFilter
    totalMsgsSent.EntityData.YangName = "total-msgs-sent"
    totalMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    totalMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    totalMsgsSent.EntityData.SegmentPath = "total-msgs-sent"
    totalMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    totalMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    totalMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    totalMsgsSent.EntityData.Children = types.NewOrderedMap()
    totalMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    totalMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", totalMsgsSent.Operator})
    totalMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", totalMsgsSent.Value})
    totalMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", totalMsgsSent.EndRangeValue})
    totalMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", totalMsgsSent.Percent})
    totalMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", totalMsgsSent.RearmType})
    totalMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", totalMsgsSent.RearmWindow})

    totalMsgsSent.EntityData.YListKeys = []string {}

    return &(totalMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd
// Number of LAbel Release messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelReleaseMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelReleaseMsgsRcvd.EntityData.YFilter = labelReleaseMsgsRcvd.YFilter
    labelReleaseMsgsRcvd.EntityData.YangName = "label-release-msgs-rcvd"
    labelReleaseMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelReleaseMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelReleaseMsgsRcvd.EntityData.SegmentPath = "label-release-msgs-rcvd"
    labelReleaseMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelReleaseMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelReleaseMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelReleaseMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    labelReleaseMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelReleaseMsgsRcvd.Operator})
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelReleaseMsgsRcvd.Value})
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelReleaseMsgsRcvd.EndRangeValue})
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelReleaseMsgsRcvd.Percent})
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelReleaseMsgsRcvd.RearmType})
    labelReleaseMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelReleaseMsgsRcvd.RearmWindow})

    labelReleaseMsgsRcvd.EntityData.YListKeys = []string {}

    return &(labelReleaseMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd
// Number of Init messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (initMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsRcvd) GetEntityData() *types.CommonEntityData {
    initMsgsRcvd.EntityData.YFilter = initMsgsRcvd.YFilter
    initMsgsRcvd.EntityData.YangName = "init-msgs-rcvd"
    initMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    initMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    initMsgsRcvd.EntityData.SegmentPath = "init-msgs-rcvd"
    initMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    initMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    initMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", initMsgsRcvd.Operator})
    initMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", initMsgsRcvd.Value})
    initMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", initMsgsRcvd.EndRangeValue})
    initMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", initMsgsRcvd.Percent})
    initMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", initMsgsRcvd.RearmType})
    initMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", initMsgsRcvd.RearmWindow})

    initMsgsRcvd.EntityData.YListKeys = []string {}

    return &(initMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent
// Number of Label Release messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelReleaseMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelReleaseMsgsSent) GetEntityData() *types.CommonEntityData {
    labelReleaseMsgsSent.EntityData.YFilter = labelReleaseMsgsSent.YFilter
    labelReleaseMsgsSent.EntityData.YangName = "label-release-msgs-sent"
    labelReleaseMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelReleaseMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelReleaseMsgsSent.EntityData.SegmentPath = "label-release-msgs-sent"
    labelReleaseMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelReleaseMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelReleaseMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelReleaseMsgsSent.EntityData.Children = types.NewOrderedMap()
    labelReleaseMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    labelReleaseMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelReleaseMsgsSent.Operator})
    labelReleaseMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelReleaseMsgsSent.Value})
    labelReleaseMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelReleaseMsgsSent.EndRangeValue})
    labelReleaseMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelReleaseMsgsSent.Percent})
    labelReleaseMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelReleaseMsgsSent.RearmType})
    labelReleaseMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelReleaseMsgsSent.RearmWindow})

    labelReleaseMsgsSent.EntityData.YListKeys = []string {}

    return &(labelReleaseMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent
// Number of Init messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (initMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_InitMsgsSent) GetEntityData() *types.CommonEntityData {
    initMsgsSent.EntityData.YFilter = initMsgsSent.YFilter
    initMsgsSent.EntityData.YangName = "init-msgs-sent"
    initMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    initMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    initMsgsSent.EntityData.SegmentPath = "init-msgs-sent"
    initMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initMsgsSent.EntityData.Children = types.NewOrderedMap()
    initMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    initMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", initMsgsSent.Operator})
    initMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", initMsgsSent.Value})
    initMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", initMsgsSent.EndRangeValue})
    initMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", initMsgsSent.Percent})
    initMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", initMsgsSent.RearmType})
    initMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", initMsgsSent.RearmWindow})

    initMsgsSent.EntityData.YListKeys = []string {}

    return &(initMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd
// Number of Label Mapping messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelMappingMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsRcvd) GetEntityData() *types.CommonEntityData {
    labelMappingMsgsRcvd.EntityData.YFilter = labelMappingMsgsRcvd.YFilter
    labelMappingMsgsRcvd.EntityData.YangName = "label-mapping-msgs-rcvd"
    labelMappingMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    labelMappingMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    labelMappingMsgsRcvd.EntityData.SegmentPath = "label-mapping-msgs-rcvd"
    labelMappingMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMappingMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMappingMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMappingMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    labelMappingMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    labelMappingMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelMappingMsgsRcvd.Operator})
    labelMappingMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelMappingMsgsRcvd.Value})
    labelMappingMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelMappingMsgsRcvd.EndRangeValue})
    labelMappingMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelMappingMsgsRcvd.Percent})
    labelMappingMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelMappingMsgsRcvd.RearmType})
    labelMappingMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelMappingMsgsRcvd.RearmWindow})

    labelMappingMsgsRcvd.EntityData.YListKeys = []string {}

    return &(labelMappingMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd
// Number of Keepalive messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (keepaliveMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsRcvd) GetEntityData() *types.CommonEntityData {
    keepaliveMsgsRcvd.EntityData.YFilter = keepaliveMsgsRcvd.YFilter
    keepaliveMsgsRcvd.EntityData.YangName = "keepalive-msgs-rcvd"
    keepaliveMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    keepaliveMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    keepaliveMsgsRcvd.EntityData.SegmentPath = "keepalive-msgs-rcvd"
    keepaliveMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepaliveMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepaliveMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepaliveMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    keepaliveMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    keepaliveMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", keepaliveMsgsRcvd.Operator})
    keepaliveMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", keepaliveMsgsRcvd.Value})
    keepaliveMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", keepaliveMsgsRcvd.EndRangeValue})
    keepaliveMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", keepaliveMsgsRcvd.Percent})
    keepaliveMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", keepaliveMsgsRcvd.RearmType})
    keepaliveMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", keepaliveMsgsRcvd.RearmWindow})

    keepaliveMsgsRcvd.EntityData.YListKeys = []string {}

    return &(keepaliveMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent
// Number of Label Mapping messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (labelMappingMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_LabelMappingMsgsSent) GetEntityData() *types.CommonEntityData {
    labelMappingMsgsSent.EntityData.YFilter = labelMappingMsgsSent.YFilter
    labelMappingMsgsSent.EntityData.YangName = "label-mapping-msgs-sent"
    labelMappingMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    labelMappingMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    labelMappingMsgsSent.EntityData.SegmentPath = "label-mapping-msgs-sent"
    labelMappingMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMappingMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMappingMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMappingMsgsSent.EntityData.Children = types.NewOrderedMap()
    labelMappingMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    labelMappingMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", labelMappingMsgsSent.Operator})
    labelMappingMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", labelMappingMsgsSent.Value})
    labelMappingMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", labelMappingMsgsSent.EndRangeValue})
    labelMappingMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", labelMappingMsgsSent.Percent})
    labelMappingMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", labelMappingMsgsSent.RearmType})
    labelMappingMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", labelMappingMsgsSent.RearmWindow})

    labelMappingMsgsSent.EntityData.YListKeys = []string {}

    return &(labelMappingMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent
// Number of Keepalive messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (keepaliveMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_KeepaliveMsgsSent) GetEntityData() *types.CommonEntityData {
    keepaliveMsgsSent.EntityData.YFilter = keepaliveMsgsSent.YFilter
    keepaliveMsgsSent.EntityData.YangName = "keepalive-msgs-sent"
    keepaliveMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    keepaliveMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    keepaliveMsgsSent.EntityData.SegmentPath = "keepalive-msgs-sent"
    keepaliveMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepaliveMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepaliveMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepaliveMsgsSent.EntityData.Children = types.NewOrderedMap()
    keepaliveMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    keepaliveMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", keepaliveMsgsSent.Operator})
    keepaliveMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", keepaliveMsgsSent.Value})
    keepaliveMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", keepaliveMsgsSent.EndRangeValue})
    keepaliveMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", keepaliveMsgsSent.Percent})
    keepaliveMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", keepaliveMsgsSent.RearmType})
    keepaliveMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", keepaliveMsgsSent.RearmWindow})

    keepaliveMsgsSent.EntityData.YListKeys = []string {}

    return &(keepaliveMsgsSent.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd
// Number of Address messages received
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (addressMsgsRcvd *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsRcvd) GetEntityData() *types.CommonEntityData {
    addressMsgsRcvd.EntityData.YFilter = addressMsgsRcvd.YFilter
    addressMsgsRcvd.EntityData.YangName = "address-msgs-rcvd"
    addressMsgsRcvd.EntityData.BundleName = "cisco_ios_xr"
    addressMsgsRcvd.EntityData.ParentYangName = "ldp-mpls-template"
    addressMsgsRcvd.EntityData.SegmentPath = "address-msgs-rcvd"
    addressMsgsRcvd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressMsgsRcvd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressMsgsRcvd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressMsgsRcvd.EntityData.Children = types.NewOrderedMap()
    addressMsgsRcvd.EntityData.Leafs = types.NewOrderedMap()
    addressMsgsRcvd.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", addressMsgsRcvd.Operator})
    addressMsgsRcvd.EntityData.Leafs.Append("value", types.YLeaf{"Value", addressMsgsRcvd.Value})
    addressMsgsRcvd.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", addressMsgsRcvd.EndRangeValue})
    addressMsgsRcvd.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", addressMsgsRcvd.Percent})
    addressMsgsRcvd.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", addressMsgsRcvd.RearmType})
    addressMsgsRcvd.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", addressMsgsRcvd.RearmWindow})

    addressMsgsRcvd.EntityData.YListKeys = []string {}

    return &(addressMsgsRcvd.EntityData)
}

// PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent
// Number of Address messages sent
// This type is a presence type.
type PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (addressMsgsSent *PerfMgmt_Threshold_LdpMpls_LdpMplsTemplates_LdpMplsTemplate_AddressMsgsSent) GetEntityData() *types.CommonEntityData {
    addressMsgsSent.EntityData.YFilter = addressMsgsSent.YFilter
    addressMsgsSent.EntityData.YangName = "address-msgs-sent"
    addressMsgsSent.EntityData.BundleName = "cisco_ios_xr"
    addressMsgsSent.EntityData.ParentYangName = "ldp-mpls-template"
    addressMsgsSent.EntityData.SegmentPath = "address-msgs-sent"
    addressMsgsSent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressMsgsSent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressMsgsSent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressMsgsSent.EntityData.Children = types.NewOrderedMap()
    addressMsgsSent.EntityData.Leafs = types.NewOrderedMap()
    addressMsgsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", addressMsgsSent.Operator})
    addressMsgsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", addressMsgsSent.Value})
    addressMsgsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", addressMsgsSent.EndRangeValue})
    addressMsgsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", addressMsgsSent.Percent})
    addressMsgsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", addressMsgsSent.RearmType})
    addressMsgsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", addressMsgsSent.RearmWindow})

    addressMsgsSent.EntityData.YListKeys = []string {}

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

    basicCounterInterface.EntityData.Children = types.NewOrderedMap()
    basicCounterInterface.EntityData.Children.Append("basic-counter-interface-templates", types.YChild{"BasicCounterInterfaceTemplates", &basicCounterInterface.BasicCounterInterfaceTemplates})
    basicCounterInterface.EntityData.Leafs = types.NewOrderedMap()

    basicCounterInterface.EntityData.YListKeys = []string {}

    return &(basicCounterInterface.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates
// Interface Basic Counter threshold templates
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Basic Counter threshold template instance. The type is slice of
    // PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate.
    BasicCounterInterfaceTemplate []*PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
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

    basicCounterInterfaceTemplates.EntityData.Children = types.NewOrderedMap()
    basicCounterInterfaceTemplates.EntityData.Children.Append("basic-counter-interface-template", types.YChild{"BasicCounterInterfaceTemplate", nil})
    for i := range basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate {
        basicCounterInterfaceTemplates.EntityData.Children.Append(types.GetSegmentPath(basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i]), types.YChild{"BasicCounterInterfaceTemplate", basicCounterInterfaceTemplates.BasicCounterInterfaceTemplate[i]})
    }
    basicCounterInterfaceTemplates.EntityData.Leafs = types.NewOrderedMap()

    basicCounterInterfaceTemplates.EntityData.YListKeys = []string {}

    return &(basicCounterInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate
// Interface Basic Counter threshold template
// instance
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate struct {
    EntityData types.CommonEntityData
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

func (basicCounterInterfaceTemplate *PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    basicCounterInterfaceTemplate.EntityData.YFilter = basicCounterInterfaceTemplate.YFilter
    basicCounterInterfaceTemplate.EntityData.YangName = "basic-counter-interface-template"
    basicCounterInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    basicCounterInterfaceTemplate.EntityData.ParentYangName = "basic-counter-interface-templates"
    basicCounterInterfaceTemplate.EntityData.SegmentPath = "basic-counter-interface-template" + types.AddKeyToken(basicCounterInterfaceTemplate.TemplateName, "template-name")
    basicCounterInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    basicCounterInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    basicCounterInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    basicCounterInterfaceTemplate.EntityData.Children = types.NewOrderedMap()
    basicCounterInterfaceTemplate.EntityData.Children.Append("in-octets", types.YChild{"InOctets", &basicCounterInterfaceTemplate.InOctets})
    basicCounterInterfaceTemplate.EntityData.Children.Append("out-octets", types.YChild{"OutOctets", &basicCounterInterfaceTemplate.OutOctets})
    basicCounterInterfaceTemplate.EntityData.Children.Append("output-queue-drops", types.YChild{"OutputQueueDrops", &basicCounterInterfaceTemplate.OutputQueueDrops})
    basicCounterInterfaceTemplate.EntityData.Children.Append("input-total-errors", types.YChild{"InputTotalErrors", &basicCounterInterfaceTemplate.InputTotalErrors})
    basicCounterInterfaceTemplate.EntityData.Children.Append("output-total-drops", types.YChild{"OutputTotalDrops", &basicCounterInterfaceTemplate.OutputTotalDrops})
    basicCounterInterfaceTemplate.EntityData.Children.Append("out-packets", types.YChild{"OutPackets", &basicCounterInterfaceTemplate.OutPackets})
    basicCounterInterfaceTemplate.EntityData.Children.Append("output-total-errors", types.YChild{"OutputTotalErrors", &basicCounterInterfaceTemplate.OutputTotalErrors})
    basicCounterInterfaceTemplate.EntityData.Children.Append("in-packets", types.YChild{"InPackets", &basicCounterInterfaceTemplate.InPackets})
    basicCounterInterfaceTemplate.EntityData.Children.Append("input-queue-drops", types.YChild{"InputQueueDrops", &basicCounterInterfaceTemplate.InputQueueDrops})
    basicCounterInterfaceTemplate.EntityData.Children.Append("input-total-drops", types.YChild{"InputTotalDrops", &basicCounterInterfaceTemplate.InputTotalDrops})
    basicCounterInterfaceTemplate.EntityData.Leafs = types.NewOrderedMap()
    basicCounterInterfaceTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", basicCounterInterfaceTemplate.TemplateName})
    basicCounterInterfaceTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", basicCounterInterfaceTemplate.SampleInterval})
    basicCounterInterfaceTemplate.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", basicCounterInterfaceTemplate.RegExpGroup})
    basicCounterInterfaceTemplate.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", basicCounterInterfaceTemplate.VrfGroup})

    basicCounterInterfaceTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(basicCounterInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets
// Number of inbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InOctets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inOctets.EntityData.Children = types.NewOrderedMap()
    inOctets.EntityData.Leafs = types.NewOrderedMap()
    inOctets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inOctets.Operator})
    inOctets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inOctets.Value})
    inOctets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inOctets.EndRangeValue})
    inOctets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inOctets.Percent})
    inOctets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inOctets.RearmType})
    inOctets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inOctets.RearmWindow})

    inOctets.EntityData.YListKeys = []string {}

    return &(inOctets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets
// Number of outbound octets/bytes
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutOctets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outOctets.EntityData.Children = types.NewOrderedMap()
    outOctets.EntityData.Leafs = types.NewOrderedMap()
    outOctets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outOctets.Operator})
    outOctets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outOctets.Value})
    outOctets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outOctets.EndRangeValue})
    outOctets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outOctets.Percent})
    outOctets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outOctets.RearmType})
    outOctets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outOctets.RearmWindow})

    outOctets.EntityData.YListKeys = []string {}

    return &(outOctets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops
// Number of outbound queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputQueueDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputQueueDrops.EntityData.Children = types.NewOrderedMap()
    outputQueueDrops.EntityData.Leafs = types.NewOrderedMap()
    outputQueueDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputQueueDrops.Operator})
    outputQueueDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputQueueDrops.Value})
    outputQueueDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputQueueDrops.EndRangeValue})
    outputQueueDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputQueueDrops.Percent})
    outputQueueDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputQueueDrops.RearmType})
    outputQueueDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputQueueDrops.RearmWindow})

    outputQueueDrops.EntityData.YListKeys = []string {}

    return &(outputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors
// Number of inbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputTotalErrors.EntityData.Children = types.NewOrderedMap()
    inputTotalErrors.EntityData.Leafs = types.NewOrderedMap()
    inputTotalErrors.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputTotalErrors.Operator})
    inputTotalErrors.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputTotalErrors.Value})
    inputTotalErrors.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputTotalErrors.EndRangeValue})
    inputTotalErrors.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputTotalErrors.Percent})
    inputTotalErrors.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputTotalErrors.RearmType})
    inputTotalErrors.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputTotalErrors.RearmWindow})

    inputTotalErrors.EntityData.YListKeys = []string {}

    return &(inputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops
// Number of outbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputTotalDrops.EntityData.Children = types.NewOrderedMap()
    outputTotalDrops.EntityData.Leafs = types.NewOrderedMap()
    outputTotalDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputTotalDrops.Operator})
    outputTotalDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputTotalDrops.Value})
    outputTotalDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputTotalDrops.EndRangeValue})
    outputTotalDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputTotalDrops.Percent})
    outputTotalDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputTotalDrops.RearmType})
    outputTotalDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputTotalDrops.RearmWindow})

    outputTotalDrops.EntityData.YListKeys = []string {}

    return &(outputTotalDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets
// Number of outbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outPackets.EntityData.Children = types.NewOrderedMap()
    outPackets.EntityData.Leafs = types.NewOrderedMap()
    outPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outPackets.Operator})
    outPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outPackets.Value})
    outPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outPackets.EndRangeValue})
    outPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outPackets.Percent})
    outPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outPackets.RearmType})
    outPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outPackets.RearmWindow})

    outPackets.EntityData.YListKeys = []string {}

    return &(outPackets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors
// Number of outbound incorrect packets
// discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_OutputTotalErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputTotalErrors.EntityData.Children = types.NewOrderedMap()
    outputTotalErrors.EntityData.Leafs = types.NewOrderedMap()
    outputTotalErrors.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputTotalErrors.Operator})
    outputTotalErrors.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputTotalErrors.Value})
    outputTotalErrors.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputTotalErrors.EndRangeValue})
    outputTotalErrors.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputTotalErrors.Percent})
    outputTotalErrors.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputTotalErrors.RearmType})
    outputTotalErrors.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputTotalErrors.RearmWindow})

    outputTotalErrors.EntityData.YListKeys = []string {}

    return &(outputTotalErrors.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets
// Number of inbound packets
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inPackets.EntityData.Children = types.NewOrderedMap()
    inPackets.EntityData.Leafs = types.NewOrderedMap()
    inPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inPackets.Operator})
    inPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inPackets.Value})
    inPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inPackets.EndRangeValue})
    inPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inPackets.Percent})
    inPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inPackets.RearmType})
    inPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inPackets.RearmWindow})

    inPackets.EntityData.YListKeys = []string {}

    return &(inPackets.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops
// Number of input queue drops
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputQueueDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputQueueDrops.EntityData.Children = types.NewOrderedMap()
    inputQueueDrops.EntityData.Leafs = types.NewOrderedMap()
    inputQueueDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputQueueDrops.Operator})
    inputQueueDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputQueueDrops.Value})
    inputQueueDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputQueueDrops.EndRangeValue})
    inputQueueDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputQueueDrops.Percent})
    inputQueueDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputQueueDrops.RearmType})
    inputQueueDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputQueueDrops.RearmWindow})

    inputQueueDrops.EntityData.YListKeys = []string {}

    return &(inputQueueDrops.EntityData)
}

// PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops
// Number of inbound correct packets discarded
// This type is a presence type.
type PerfMgmt_Threshold_BasicCounterInterface_BasicCounterInterfaceTemplates_BasicCounterInterfaceTemplate_InputTotalDrops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputTotalDrops.EntityData.Children = types.NewOrderedMap()
    inputTotalDrops.EntityData.Leafs = types.NewOrderedMap()
    inputTotalDrops.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputTotalDrops.Operator})
    inputTotalDrops.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputTotalDrops.Value})
    inputTotalDrops.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputTotalDrops.EndRangeValue})
    inputTotalDrops.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputTotalDrops.Percent})
    inputTotalDrops.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputTotalDrops.RearmType})
    inputTotalDrops.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputTotalDrops.RearmWindow})

    inputTotalDrops.EntityData.YListKeys = []string {}

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

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("bgp-templates", types.YChild{"BgpTemplates", &bgp.BgpTemplates})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates
// BGP threshold templates
type PerfMgmt_Threshold_Bgp_BgpTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate.
    BgpTemplate []*PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
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

    bgpTemplates.EntityData.Children = types.NewOrderedMap()
    bgpTemplates.EntityData.Children.Append("bgp-template", types.YChild{"BgpTemplate", nil})
    for i := range bgpTemplates.BgpTemplate {
        bgpTemplates.EntityData.Children.Append(types.GetSegmentPath(bgpTemplates.BgpTemplate[i]), types.YChild{"BgpTemplate", bgpTemplates.BgpTemplate[i]})
    }
    bgpTemplates.EntityData.Leafs = types.NewOrderedMap()

    bgpTemplates.EntityData.YListKeys = []string {}

    return &(bgpTemplates.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate
// BGP threshold template instance
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate struct {
    EntityData types.CommonEntityData
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

func (bgpTemplate *PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate) GetEntityData() *types.CommonEntityData {
    bgpTemplate.EntityData.YFilter = bgpTemplate.YFilter
    bgpTemplate.EntityData.YangName = "bgp-template"
    bgpTemplate.EntityData.BundleName = "cisco_ios_xr"
    bgpTemplate.EntityData.ParentYangName = "bgp-templates"
    bgpTemplate.EntityData.SegmentPath = "bgp-template" + types.AddKeyToken(bgpTemplate.TemplateName, "template-name")
    bgpTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgpTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgpTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgpTemplate.EntityData.Children = types.NewOrderedMap()
    bgpTemplate.EntityData.Children.Append("output-update-messages", types.YChild{"OutputUpdateMessages", &bgpTemplate.OutputUpdateMessages})
    bgpTemplate.EntityData.Children.Append("errors-received", types.YChild{"ErrorsReceived", &bgpTemplate.ErrorsReceived})
    bgpTemplate.EntityData.Children.Append("conn-established", types.YChild{"ConnEstablished", &bgpTemplate.ConnEstablished})
    bgpTemplate.EntityData.Children.Append("output-messages", types.YChild{"OutputMessages", &bgpTemplate.OutputMessages})
    bgpTemplate.EntityData.Children.Append("conn-dropped", types.YChild{"ConnDropped", &bgpTemplate.ConnDropped})
    bgpTemplate.EntityData.Children.Append("input-update-messages", types.YChild{"InputUpdateMessages", &bgpTemplate.InputUpdateMessages})
    bgpTemplate.EntityData.Children.Append("errors-sent", types.YChild{"ErrorsSent", &bgpTemplate.ErrorsSent})
    bgpTemplate.EntityData.Children.Append("input-messages", types.YChild{"InputMessages", &bgpTemplate.InputMessages})
    bgpTemplate.EntityData.Leafs = types.NewOrderedMap()
    bgpTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", bgpTemplate.TemplateName})
    bgpTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", bgpTemplate.SampleInterval})

    bgpTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(bgpTemplate.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages
// Number of update messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputUpdateMessages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputUpdateMessages.EntityData.Children = types.NewOrderedMap()
    outputUpdateMessages.EntityData.Leafs = types.NewOrderedMap()
    outputUpdateMessages.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputUpdateMessages.Operator})
    outputUpdateMessages.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputUpdateMessages.Value})
    outputUpdateMessages.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputUpdateMessages.EndRangeValue})
    outputUpdateMessages.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputUpdateMessages.Percent})
    outputUpdateMessages.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputUpdateMessages.RearmType})
    outputUpdateMessages.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputUpdateMessages.RearmWindow})

    outputUpdateMessages.EntityData.YListKeys = []string {}

    return &(outputUpdateMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived
// Number of error notifications received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsReceived struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    errorsReceived.EntityData.Children = types.NewOrderedMap()
    errorsReceived.EntityData.Leafs = types.NewOrderedMap()
    errorsReceived.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", errorsReceived.Operator})
    errorsReceived.EntityData.Leafs.Append("value", types.YLeaf{"Value", errorsReceived.Value})
    errorsReceived.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", errorsReceived.EndRangeValue})
    errorsReceived.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", errorsReceived.Percent})
    errorsReceived.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", errorsReceived.RearmType})
    errorsReceived.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", errorsReceived.RearmWindow})

    errorsReceived.EntityData.YListKeys = []string {}

    return &(errorsReceived.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished
// Number of times the connection was
// established
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnEstablished struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    connEstablished.EntityData.Children = types.NewOrderedMap()
    connEstablished.EntityData.Leafs = types.NewOrderedMap()
    connEstablished.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", connEstablished.Operator})
    connEstablished.EntityData.Leafs.Append("value", types.YLeaf{"Value", connEstablished.Value})
    connEstablished.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", connEstablished.EndRangeValue})
    connEstablished.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", connEstablished.Percent})
    connEstablished.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", connEstablished.RearmType})
    connEstablished.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", connEstablished.RearmWindow})

    connEstablished.EntityData.YListKeys = []string {}

    return &(connEstablished.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages
// Number of messages sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_OutputMessages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputMessages.EntityData.Children = types.NewOrderedMap()
    outputMessages.EntityData.Leafs = types.NewOrderedMap()
    outputMessages.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputMessages.Operator})
    outputMessages.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputMessages.Value})
    outputMessages.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputMessages.EndRangeValue})
    outputMessages.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputMessages.Percent})
    outputMessages.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputMessages.RearmType})
    outputMessages.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputMessages.RearmWindow})

    outputMessages.EntityData.YListKeys = []string {}

    return &(outputMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped
// Number of times the connection was dropped
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ConnDropped struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    connDropped.EntityData.Children = types.NewOrderedMap()
    connDropped.EntityData.Leafs = types.NewOrderedMap()
    connDropped.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", connDropped.Operator})
    connDropped.EntityData.Leafs.Append("value", types.YLeaf{"Value", connDropped.Value})
    connDropped.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", connDropped.EndRangeValue})
    connDropped.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", connDropped.Percent})
    connDropped.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", connDropped.RearmType})
    connDropped.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", connDropped.RearmWindow})

    connDropped.EntityData.YListKeys = []string {}

    return &(connDropped.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages
// Number of update messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputUpdateMessages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputUpdateMessages.EntityData.Children = types.NewOrderedMap()
    inputUpdateMessages.EntityData.Leafs = types.NewOrderedMap()
    inputUpdateMessages.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputUpdateMessages.Operator})
    inputUpdateMessages.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputUpdateMessages.Value})
    inputUpdateMessages.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputUpdateMessages.EndRangeValue})
    inputUpdateMessages.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputUpdateMessages.Percent})
    inputUpdateMessages.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputUpdateMessages.RearmType})
    inputUpdateMessages.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputUpdateMessages.RearmWindow})

    inputUpdateMessages.EntityData.YListKeys = []string {}

    return &(inputUpdateMessages.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent
// Number of error notifications sent
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_ErrorsSent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    errorsSent.EntityData.Children = types.NewOrderedMap()
    errorsSent.EntityData.Leafs = types.NewOrderedMap()
    errorsSent.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", errorsSent.Operator})
    errorsSent.EntityData.Leafs.Append("value", types.YLeaf{"Value", errorsSent.Value})
    errorsSent.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", errorsSent.EndRangeValue})
    errorsSent.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", errorsSent.Percent})
    errorsSent.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", errorsSent.RearmType})
    errorsSent.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", errorsSent.RearmWindow})

    errorsSent.EntityData.YListKeys = []string {}

    return &(errorsSent.EntityData)
}

// PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages
// Number of messages received
// This type is a presence type.
type PerfMgmt_Threshold_Bgp_BgpTemplates_BgpTemplate_InputMessages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputMessages.EntityData.Children = types.NewOrderedMap()
    inputMessages.EntityData.Leafs = types.NewOrderedMap()
    inputMessages.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputMessages.Operator})
    inputMessages.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputMessages.Value})
    inputMessages.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputMessages.EndRangeValue})
    inputMessages.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputMessages.Percent})
    inputMessages.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputMessages.RearmType})
    inputMessages.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputMessages.RearmWindow})

    inputMessages.EntityData.YListKeys = []string {}

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

    ospfv2Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv2Protocol.EntityData.Children.Append("ospfv2-protocol-templates", types.YChild{"Ospfv2ProtocolTemplates", &ospfv2Protocol.Ospfv2ProtocolTemplates})
    ospfv2Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv2Protocol.EntityData.YListKeys = []string {}

    return &(ospfv2Protocol.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate.
    Ospfv2ProtocolTemplate []*PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
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

    ospfv2ProtocolTemplates.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolTemplates.EntityData.Children.Append("ospfv2-protocol-template", types.YChild{"Ospfv2ProtocolTemplate", nil})
    for i := range ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate {
        ospfv2ProtocolTemplates.EntityData.Children.Append(types.GetSegmentPath(ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i]), types.YChild{"Ospfv2ProtocolTemplate", ospfv2ProtocolTemplates.Ospfv2ProtocolTemplate[i]})
    }
    ospfv2ProtocolTemplates.EntityData.Leafs = types.NewOrderedMap()

    ospfv2ProtocolTemplates.EntityData.YListKeys = []string {}

    return &(ospfv2ProtocolTemplates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate struct {
    EntityData types.CommonEntityData
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

func (ospfv2ProtocolTemplate *PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate) GetEntityData() *types.CommonEntityData {
    ospfv2ProtocolTemplate.EntityData.YFilter = ospfv2ProtocolTemplate.YFilter
    ospfv2ProtocolTemplate.EntityData.YangName = "ospfv2-protocol-template"
    ospfv2ProtocolTemplate.EntityData.BundleName = "cisco_ios_xr"
    ospfv2ProtocolTemplate.EntityData.ParentYangName = "ospfv2-protocol-templates"
    ospfv2ProtocolTemplate.EntityData.SegmentPath = "ospfv2-protocol-template" + types.AddKeyToken(ospfv2ProtocolTemplate.TemplateName, "template-name")
    ospfv2ProtocolTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv2ProtocolTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv2ProtocolTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv2ProtocolTemplate.EntityData.Children = types.NewOrderedMap()
    ospfv2ProtocolTemplate.EntityData.Children.Append("checksum-errors", types.YChild{"ChecksumErrors", &ospfv2ProtocolTemplate.ChecksumErrors})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-lsa-acks-lsa", types.YChild{"InputLsaAcksLsa", &ospfv2ProtocolTemplate.InputLsaAcksLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-db-ds-lsa", types.YChild{"OutputDbDsLsa", &ospfv2ProtocolTemplate.OutputDbDsLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-db-ds-lsa", types.YChild{"InputDbDsLsa", &ospfv2ProtocolTemplate.InputDbDsLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-lsa-updates", types.YChild{"InputLsaUpdates", &ospfv2ProtocolTemplate.InputLsaUpdates})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-db-ds", types.YChild{"OutputDbDs", &ospfv2ProtocolTemplate.OutputDbDs})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-lsa-updates-lsa", types.YChild{"OutputLsaUpdatesLsa", &ospfv2ProtocolTemplate.OutputLsaUpdatesLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-db-ds", types.YChild{"InputDbDs", &ospfv2ProtocolTemplate.InputDbDs})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-lsa-updates-lsa", types.YChild{"InputLsaUpdatesLsa", &ospfv2ProtocolTemplate.InputLsaUpdatesLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-packets", types.YChild{"OutputPackets", &ospfv2ProtocolTemplate.OutputPackets})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-packets", types.YChild{"InputPackets", &ospfv2ProtocolTemplate.InputPackets})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-hello-packets", types.YChild{"OutputHelloPackets", &ospfv2ProtocolTemplate.OutputHelloPackets})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-hello-packets", types.YChild{"InputHelloPackets", &ospfv2ProtocolTemplate.InputHelloPackets})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-ls-requests", types.YChild{"OutputLsRequests", &ospfv2ProtocolTemplate.OutputLsRequests})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-lsa-acks-lsa", types.YChild{"OutputLsaAcksLsa", &ospfv2ProtocolTemplate.OutputLsaAcksLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-lsa-acks", types.YChild{"OutputLsaAcks", &ospfv2ProtocolTemplate.OutputLsaAcks})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-lsa-acks", types.YChild{"InputLsaAcks", &ospfv2ProtocolTemplate.InputLsaAcks})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-lsa-updates", types.YChild{"OutputLsaUpdates", &ospfv2ProtocolTemplate.OutputLsaUpdates})
    ospfv2ProtocolTemplate.EntityData.Children.Append("output-ls-requests-lsa", types.YChild{"OutputLsRequestsLsa", &ospfv2ProtocolTemplate.OutputLsRequestsLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-ls-requests-lsa", types.YChild{"InputLsRequestsLsa", &ospfv2ProtocolTemplate.InputLsRequestsLsa})
    ospfv2ProtocolTemplate.EntityData.Children.Append("input-ls-requests", types.YChild{"InputLsRequests", &ospfv2ProtocolTemplate.InputLsRequests})
    ospfv2ProtocolTemplate.EntityData.Leafs = types.NewOrderedMap()
    ospfv2ProtocolTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv2ProtocolTemplate.TemplateName})
    ospfv2ProtocolTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", ospfv2ProtocolTemplate.SampleInterval})

    ospfv2ProtocolTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(ospfv2ProtocolTemplate.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors
// Number of packets received with checksum
// errors
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_ChecksumErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    checksumErrors.EntityData.Children = types.NewOrderedMap()
    checksumErrors.EntityData.Leafs = types.NewOrderedMap()
    checksumErrors.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", checksumErrors.Operator})
    checksumErrors.EntityData.Leafs.Append("value", types.YLeaf{"Value", checksumErrors.Value})
    checksumErrors.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", checksumErrors.EndRangeValue})
    checksumErrors.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", checksumErrors.Percent})
    checksumErrors.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", checksumErrors.RearmType})
    checksumErrors.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", checksumErrors.RearmWindow})

    checksumErrors.EntityData.YListKeys = []string {}

    return &(checksumErrors.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcksLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaAcksLsa.EntityData.Children = types.NewOrderedMap()
    inputLsaAcksLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsaAcksLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaAcksLsa.Operator})
    inputLsaAcksLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaAcksLsa.Value})
    inputLsaAcksLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaAcksLsa.EndRangeValue})
    inputLsaAcksLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaAcksLsa.Percent})
    inputLsaAcksLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaAcksLsa.RearmType})
    inputLsaAcksLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaAcksLsa.RearmWindow})

    inputLsaAcksLsa.EntityData.YListKeys = []string {}

    return &(inputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputDbDsLsa.EntityData.Children = types.NewOrderedMap()
    outputDbDsLsa.EntityData.Leafs = types.NewOrderedMap()
    outputDbDsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputDbDsLsa.Operator})
    outputDbDsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputDbDsLsa.Value})
    outputDbDsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputDbDsLsa.EndRangeValue})
    outputDbDsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputDbDsLsa.Percent})
    outputDbDsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputDbDsLsa.RearmType})
    outputDbDsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputDbDsLsa.RearmWindow})

    outputDbDsLsa.EntityData.YListKeys = []string {}

    return &(outputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputDbDsLsa.EntityData.Children = types.NewOrderedMap()
    inputDbDsLsa.EntityData.Leafs = types.NewOrderedMap()
    inputDbDsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputDbDsLsa.Operator})
    inputDbDsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputDbDsLsa.Value})
    inputDbDsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputDbDsLsa.EndRangeValue})
    inputDbDsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputDbDsLsa.Percent})
    inputDbDsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputDbDsLsa.RearmType})
    inputDbDsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputDbDsLsa.RearmWindow})

    inputDbDsLsa.EntityData.YListKeys = []string {}

    return &(inputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaUpdates.EntityData.Children = types.NewOrderedMap()
    inputLsaUpdates.EntityData.Leafs = types.NewOrderedMap()
    inputLsaUpdates.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaUpdates.Operator})
    inputLsaUpdates.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaUpdates.Value})
    inputLsaUpdates.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaUpdates.EndRangeValue})
    inputLsaUpdates.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaUpdates.Percent})
    inputLsaUpdates.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaUpdates.RearmType})
    inputLsaUpdates.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaUpdates.RearmWindow})

    inputLsaUpdates.EntityData.YListKeys = []string {}

    return &(inputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputDbDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputDbDs.EntityData.Children = types.NewOrderedMap()
    outputDbDs.EntityData.Leafs = types.NewOrderedMap()
    outputDbDs.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputDbDs.Operator})
    outputDbDs.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputDbDs.Value})
    outputDbDs.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputDbDs.EndRangeValue})
    outputDbDs.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputDbDs.Percent})
    outputDbDs.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputDbDs.RearmType})
    outputDbDs.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputDbDs.RearmWindow})

    outputDbDs.EntityData.YListKeys = []string {}

    return &(outputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaUpdatesLsa.EntityData.Children = types.NewOrderedMap()
    outputLsaUpdatesLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsaUpdatesLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaUpdatesLsa.Operator})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaUpdatesLsa.Value})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaUpdatesLsa.EndRangeValue})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaUpdatesLsa.Percent})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaUpdatesLsa.RearmType})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaUpdatesLsa.RearmWindow})

    outputLsaUpdatesLsa.EntityData.YListKeys = []string {}

    return &(outputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputDbDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputDbDs.EntityData.Children = types.NewOrderedMap()
    inputDbDs.EntityData.Leafs = types.NewOrderedMap()
    inputDbDs.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputDbDs.Operator})
    inputDbDs.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputDbDs.Value})
    inputDbDs.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputDbDs.EndRangeValue})
    inputDbDs.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputDbDs.Percent})
    inputDbDs.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputDbDs.RearmType})
    inputDbDs.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputDbDs.RearmWindow})

    inputDbDs.EntityData.YListKeys = []string {}

    return &(inputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaUpdatesLsa.EntityData.Children = types.NewOrderedMap()
    inputLsaUpdatesLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsaUpdatesLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaUpdatesLsa.Operator})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaUpdatesLsa.Value})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaUpdatesLsa.EndRangeValue})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaUpdatesLsa.Percent})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaUpdatesLsa.RearmType})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaUpdatesLsa.RearmWindow})

    inputLsaUpdatesLsa.EntityData.YListKeys = []string {}

    return &(inputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputPackets.EntityData.Children = types.NewOrderedMap()
    outputPackets.EntityData.Leafs = types.NewOrderedMap()
    outputPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputPackets.Operator})
    outputPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputPackets.Value})
    outputPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputPackets.EndRangeValue})
    outputPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputPackets.Percent})
    outputPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputPackets.RearmType})
    outputPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputPackets.RearmWindow})

    outputPackets.EntityData.YListKeys = []string {}

    return &(outputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputPackets.EntityData.Children = types.NewOrderedMap()
    inputPackets.EntityData.Leafs = types.NewOrderedMap()
    inputPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputPackets.Operator})
    inputPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputPackets.Value})
    inputPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputPackets.EndRangeValue})
    inputPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputPackets.Percent})
    inputPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputPackets.RearmType})
    inputPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputPackets.RearmWindow})

    inputPackets.EntityData.YListKeys = []string {}

    return &(inputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputHelloPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputHelloPackets.EntityData.Children = types.NewOrderedMap()
    outputHelloPackets.EntityData.Leafs = types.NewOrderedMap()
    outputHelloPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputHelloPackets.Operator})
    outputHelloPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputHelloPackets.Value})
    outputHelloPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputHelloPackets.EndRangeValue})
    outputHelloPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputHelloPackets.Percent})
    outputHelloPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputHelloPackets.RearmType})
    outputHelloPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputHelloPackets.RearmWindow})

    outputHelloPackets.EntityData.YListKeys = []string {}

    return &(outputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputHelloPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputHelloPackets.EntityData.Children = types.NewOrderedMap()
    inputHelloPackets.EntityData.Leafs = types.NewOrderedMap()
    inputHelloPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputHelloPackets.Operator})
    inputHelloPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputHelloPackets.Value})
    inputHelloPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputHelloPackets.EndRangeValue})
    inputHelloPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputHelloPackets.Percent})
    inputHelloPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputHelloPackets.RearmType})
    inputHelloPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputHelloPackets.RearmWindow})

    inputHelloPackets.EntityData.YListKeys = []string {}

    return &(inputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsRequests.EntityData.Children = types.NewOrderedMap()
    outputLsRequests.EntityData.Leafs = types.NewOrderedMap()
    outputLsRequests.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsRequests.Operator})
    outputLsRequests.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsRequests.Value})
    outputLsRequests.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsRequests.EndRangeValue})
    outputLsRequests.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsRequests.Percent})
    outputLsRequests.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsRequests.RearmType})
    outputLsRequests.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsRequests.RearmWindow})

    outputLsRequests.EntityData.YListKeys = []string {}

    return &(outputLsRequests.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcksLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaAcksLsa.EntityData.Children = types.NewOrderedMap()
    outputLsaAcksLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsaAcksLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaAcksLsa.Operator})
    outputLsaAcksLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaAcksLsa.Value})
    outputLsaAcksLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaAcksLsa.EndRangeValue})
    outputLsaAcksLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaAcksLsa.Percent})
    outputLsaAcksLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaAcksLsa.RearmType})
    outputLsaAcksLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaAcksLsa.RearmWindow})

    outputLsaAcksLsa.EntityData.YListKeys = []string {}

    return &(outputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaAcks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaAcks.EntityData.Children = types.NewOrderedMap()
    outputLsaAcks.EntityData.Leafs = types.NewOrderedMap()
    outputLsaAcks.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaAcks.Operator})
    outputLsaAcks.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaAcks.Value})
    outputLsaAcks.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaAcks.EndRangeValue})
    outputLsaAcks.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaAcks.Percent})
    outputLsaAcks.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaAcks.RearmType})
    outputLsaAcks.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaAcks.RearmWindow})

    outputLsaAcks.EntityData.YListKeys = []string {}

    return &(outputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsaAcks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaAcks.EntityData.Children = types.NewOrderedMap()
    inputLsaAcks.EntityData.Leafs = types.NewOrderedMap()
    inputLsaAcks.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaAcks.Operator})
    inputLsaAcks.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaAcks.Value})
    inputLsaAcks.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaAcks.EndRangeValue})
    inputLsaAcks.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaAcks.Percent})
    inputLsaAcks.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaAcks.RearmType})
    inputLsaAcks.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaAcks.RearmWindow})

    inputLsaAcks.EntityData.YListKeys = []string {}

    return &(inputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsaUpdates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaUpdates.EntityData.Children = types.NewOrderedMap()
    outputLsaUpdates.EntityData.Leafs = types.NewOrderedMap()
    outputLsaUpdates.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaUpdates.Operator})
    outputLsaUpdates.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaUpdates.Value})
    outputLsaUpdates.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaUpdates.EndRangeValue})
    outputLsaUpdates.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaUpdates.Percent})
    outputLsaUpdates.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaUpdates.RearmType})
    outputLsaUpdates.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaUpdates.RearmWindow})

    outputLsaUpdates.EntityData.YListKeys = []string {}

    return &(outputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_OutputLsRequestsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsRequestsLsa.EntityData.Children = types.NewOrderedMap()
    outputLsRequestsLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsRequestsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsRequestsLsa.Operator})
    outputLsRequestsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsRequestsLsa.Value})
    outputLsRequestsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsRequestsLsa.EndRangeValue})
    outputLsRequestsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsRequestsLsa.Percent})
    outputLsRequestsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsRequestsLsa.RearmType})
    outputLsRequestsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsRequestsLsa.RearmWindow})

    outputLsRequestsLsa.EntityData.YListKeys = []string {}

    return &(outputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequestsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsRequestsLsa.EntityData.Children = types.NewOrderedMap()
    inputLsRequestsLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsRequestsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsRequestsLsa.Operator})
    inputLsRequestsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsRequestsLsa.Value})
    inputLsRequestsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsRequestsLsa.EndRangeValue})
    inputLsRequestsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsRequestsLsa.Percent})
    inputLsRequestsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsRequestsLsa.RearmType})
    inputLsRequestsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsRequestsLsa.RearmWindow})

    inputLsRequestsLsa.EntityData.YListKeys = []string {}

    return &(inputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv2Protocol_Ospfv2ProtocolTemplates_Ospfv2ProtocolTemplate_InputLsRequests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsRequests.EntityData.Children = types.NewOrderedMap()
    inputLsRequests.EntityData.Leafs = types.NewOrderedMap()
    inputLsRequests.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsRequests.Operator})
    inputLsRequests.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsRequests.Value})
    inputLsRequests.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsRequests.EndRangeValue})
    inputLsRequests.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsRequests.Percent})
    inputLsRequests.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsRequests.RearmType})
    inputLsRequests.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsRequests.RearmWindow})

    inputLsRequests.EntityData.YListKeys = []string {}

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

    cpuNode.EntityData.Children = types.NewOrderedMap()
    cpuNode.EntityData.Children.Append("cpu-node-templates", types.YChild{"CpuNodeTemplates", &cpuNode.CpuNodeTemplates})
    cpuNode.EntityData.Leafs = types.NewOrderedMap()

    cpuNode.EntityData.YListKeys = []string {}

    return &(cpuNode.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates
// Node CPU threshold configuration templates
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node CPU threshold configuration template instances. The type is slice of
    // PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate.
    CpuNodeTemplate []*PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
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

    cpuNodeTemplates.EntityData.Children = types.NewOrderedMap()
    cpuNodeTemplates.EntityData.Children.Append("cpu-node-template", types.YChild{"CpuNodeTemplate", nil})
    for i := range cpuNodeTemplates.CpuNodeTemplate {
        cpuNodeTemplates.EntityData.Children.Append(types.GetSegmentPath(cpuNodeTemplates.CpuNodeTemplate[i]), types.YChild{"CpuNodeTemplate", cpuNodeTemplates.CpuNodeTemplate[i]})
    }
    cpuNodeTemplates.EntityData.Leafs = types.NewOrderedMap()

    cpuNodeTemplates.EntityData.YListKeys = []string {}

    return &(cpuNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate
// Node CPU threshold configuration template
// instances
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate struct {
    EntityData types.CommonEntityData
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

func (cpuNodeTemplate *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate) GetEntityData() *types.CommonEntityData {
    cpuNodeTemplate.EntityData.YFilter = cpuNodeTemplate.YFilter
    cpuNodeTemplate.EntityData.YangName = "cpu-node-template"
    cpuNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    cpuNodeTemplate.EntityData.ParentYangName = "cpu-node-templates"
    cpuNodeTemplate.EntityData.SegmentPath = "cpu-node-template" + types.AddKeyToken(cpuNodeTemplate.TemplateName, "template-name")
    cpuNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cpuNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cpuNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cpuNodeTemplate.EntityData.Children = types.NewOrderedMap()
    cpuNodeTemplate.EntityData.Children.Append("average-cpu-used", types.YChild{"AverageCpuUsed", &cpuNodeTemplate.AverageCpuUsed})
    cpuNodeTemplate.EntityData.Children.Append("no-processes", types.YChild{"NoProcesses", &cpuNodeTemplate.NoProcesses})
    cpuNodeTemplate.EntityData.Leafs = types.NewOrderedMap()
    cpuNodeTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", cpuNodeTemplate.TemplateName})
    cpuNodeTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", cpuNodeTemplate.SampleInterval})

    cpuNodeTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(cpuNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_AverageCpuUsed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    averageCpuUsed.EntityData.Children = types.NewOrderedMap()
    averageCpuUsed.EntityData.Leafs = types.NewOrderedMap()
    averageCpuUsed.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", averageCpuUsed.Operator})
    averageCpuUsed.EntityData.Leafs.Append("value", types.YLeaf{"Value", averageCpuUsed.Value})
    averageCpuUsed.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", averageCpuUsed.EndRangeValue})
    averageCpuUsed.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", averageCpuUsed.Percent})
    averageCpuUsed.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", averageCpuUsed.RearmType})
    averageCpuUsed.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", averageCpuUsed.RearmWindow})

    averageCpuUsed.EntityData.YListKeys = []string {}

    return &(averageCpuUsed.EntityData)
}

// PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses
// Number of processes
// This type is a presence type.
type PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (noProcesses *PerfMgmt_Threshold_CpuNode_CpuNodeTemplates_CpuNodeTemplate_NoProcesses) GetEntityData() *types.CommonEntityData {
    noProcesses.EntityData.YFilter = noProcesses.YFilter
    noProcesses.EntityData.YangName = "no-processes"
    noProcesses.EntityData.BundleName = "cisco_ios_xr"
    noProcesses.EntityData.ParentYangName = "cpu-node-template"
    noProcesses.EntityData.SegmentPath = "no-processes"
    noProcesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    noProcesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    noProcesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    noProcesses.EntityData.Children = types.NewOrderedMap()
    noProcesses.EntityData.Leafs = types.NewOrderedMap()
    noProcesses.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", noProcesses.Operator})
    noProcesses.EntityData.Leafs.Append("value", types.YLeaf{"Value", noProcesses.Value})
    noProcesses.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", noProcesses.EndRangeValue})
    noProcesses.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", noProcesses.Percent})
    noProcesses.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", noProcesses.RearmType})
    noProcesses.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", noProcesses.RearmWindow})

    noProcesses.EntityData.YListKeys = []string {}

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

    dataRateInterface.EntityData.Children = types.NewOrderedMap()
    dataRateInterface.EntityData.Children.Append("data-rate-interface-templates", types.YChild{"DataRateInterfaceTemplates", &dataRateInterface.DataRateInterfaceTemplates})
    dataRateInterface.EntityData.Leafs = types.NewOrderedMap()

    dataRateInterface.EntityData.YListKeys = []string {}

    return &(dataRateInterface.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates
// Interface Data Rates threshold templates
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Data Rates threshold template instance. The type is slice of
    // PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate.
    DataRateInterfaceTemplate []*PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
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

    dataRateInterfaceTemplates.EntityData.Children = types.NewOrderedMap()
    dataRateInterfaceTemplates.EntityData.Children.Append("data-rate-interface-template", types.YChild{"DataRateInterfaceTemplate", nil})
    for i := range dataRateInterfaceTemplates.DataRateInterfaceTemplate {
        dataRateInterfaceTemplates.EntityData.Children.Append(types.GetSegmentPath(dataRateInterfaceTemplates.DataRateInterfaceTemplate[i]), types.YChild{"DataRateInterfaceTemplate", dataRateInterfaceTemplates.DataRateInterfaceTemplate[i]})
    }
    dataRateInterfaceTemplates.EntityData.Leafs = types.NewOrderedMap()

    dataRateInterfaceTemplates.EntityData.YListKeys = []string {}

    return &(dataRateInterfaceTemplates.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate
// Interface Data Rates threshold template
// instance
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate struct {
    EntityData types.CommonEntityData
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

func (dataRateInterfaceTemplate *PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate) GetEntityData() *types.CommonEntityData {
    dataRateInterfaceTemplate.EntityData.YFilter = dataRateInterfaceTemplate.YFilter
    dataRateInterfaceTemplate.EntityData.YangName = "data-rate-interface-template"
    dataRateInterfaceTemplate.EntityData.BundleName = "cisco_ios_xr"
    dataRateInterfaceTemplate.EntityData.ParentYangName = "data-rate-interface-templates"
    dataRateInterfaceTemplate.EntityData.SegmentPath = "data-rate-interface-template" + types.AddKeyToken(dataRateInterfaceTemplate.TemplateName, "template-name")
    dataRateInterfaceTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataRateInterfaceTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataRateInterfaceTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataRateInterfaceTemplate.EntityData.Children = types.NewOrderedMap()
    dataRateInterfaceTemplate.EntityData.Children.Append("input-data-rate", types.YChild{"InputDataRate", &dataRateInterfaceTemplate.InputDataRate})
    dataRateInterfaceTemplate.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &dataRateInterfaceTemplate.Bandwidth})
    dataRateInterfaceTemplate.EntityData.Children.Append("output-packet-rate", types.YChild{"OutputPacketRate", &dataRateInterfaceTemplate.OutputPacketRate})
    dataRateInterfaceTemplate.EntityData.Children.Append("input-peak-pkts", types.YChild{"InputPeakPkts", &dataRateInterfaceTemplate.InputPeakPkts})
    dataRateInterfaceTemplate.EntityData.Children.Append("output-peak-rate", types.YChild{"OutputPeakRate", &dataRateInterfaceTemplate.OutputPeakRate})
    dataRateInterfaceTemplate.EntityData.Children.Append("output-data-rate", types.YChild{"OutputDataRate", &dataRateInterfaceTemplate.OutputDataRate})
    dataRateInterfaceTemplate.EntityData.Children.Append("input-packet-rate", types.YChild{"InputPacketRate", &dataRateInterfaceTemplate.InputPacketRate})
    dataRateInterfaceTemplate.EntityData.Children.Append("output-peak-pkts", types.YChild{"OutputPeakPkts", &dataRateInterfaceTemplate.OutputPeakPkts})
    dataRateInterfaceTemplate.EntityData.Children.Append("input-peak-rate", types.YChild{"InputPeakRate", &dataRateInterfaceTemplate.InputPeakRate})
    dataRateInterfaceTemplate.EntityData.Leafs = types.NewOrderedMap()
    dataRateInterfaceTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", dataRateInterfaceTemplate.TemplateName})
    dataRateInterfaceTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", dataRateInterfaceTemplate.SampleInterval})
    dataRateInterfaceTemplate.EntityData.Leafs.Append("reg-exp-group", types.YLeaf{"RegExpGroup", dataRateInterfaceTemplate.RegExpGroup})
    dataRateInterfaceTemplate.EntityData.Leafs.Append("vrf-group", types.YLeaf{"VrfGroup", dataRateInterfaceTemplate.VrfGroup})

    dataRateInterfaceTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(dataRateInterfaceTemplate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate
// Input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputDataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputDataRate.EntityData.Children = types.NewOrderedMap()
    inputDataRate.EntityData.Leafs = types.NewOrderedMap()
    inputDataRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputDataRate.Operator})
    inputDataRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputDataRate.Value})
    inputDataRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputDataRate.EndRangeValue})
    inputDataRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputDataRate.Percent})
    inputDataRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputDataRate.RearmType})
    inputDataRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputDataRate.RearmWindow})

    inputDataRate.EntityData.YListKeys = []string {}

    return &(inputDataRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth
// Bandwidth in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Leafs = types.NewOrderedMap()
    bandwidth.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", bandwidth.Operator})
    bandwidth.EntityData.Leafs.Append("value", types.YLeaf{"Value", bandwidth.Value})
    bandwidth.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", bandwidth.EndRangeValue})
    bandwidth.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", bandwidth.Percent})
    bandwidth.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", bandwidth.RearmType})
    bandwidth.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", bandwidth.RearmWindow})

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate
// Number of Output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPacketRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputPacketRate.EntityData.Children = types.NewOrderedMap()
    outputPacketRate.EntityData.Leafs = types.NewOrderedMap()
    outputPacketRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputPacketRate.Operator})
    outputPacketRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputPacketRate.Value})
    outputPacketRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputPacketRate.EndRangeValue})
    outputPacketRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputPacketRate.Percent})
    outputPacketRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputPacketRate.RearmType})
    outputPacketRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputPacketRate.RearmWindow})

    outputPacketRate.EntityData.YListKeys = []string {}

    return &(outputPacketRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts
// Maximum number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputPeakPkts.EntityData.Children = types.NewOrderedMap()
    inputPeakPkts.EntityData.Leafs = types.NewOrderedMap()
    inputPeakPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputPeakPkts.Operator})
    inputPeakPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputPeakPkts.Value})
    inputPeakPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputPeakPkts.EndRangeValue})
    inputPeakPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputPeakPkts.Percent})
    inputPeakPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputPeakPkts.RearmType})
    inputPeakPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputPeakPkts.RearmWindow})

    inputPeakPkts.EntityData.YListKeys = []string {}

    return &(inputPeakPkts.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate
// Peak output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputPeakRate.EntityData.Children = types.NewOrderedMap()
    outputPeakRate.EntityData.Leafs = types.NewOrderedMap()
    outputPeakRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputPeakRate.Operator})
    outputPeakRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputPeakRate.Value})
    outputPeakRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputPeakRate.EndRangeValue})
    outputPeakRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputPeakRate.Percent})
    outputPeakRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputPeakRate.RearmType})
    outputPeakRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputPeakRate.RearmWindow})

    outputPeakRate.EntityData.YListKeys = []string {}

    return &(outputPeakRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate
// Output data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputDataRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputDataRate.EntityData.Children = types.NewOrderedMap()
    outputDataRate.EntityData.Leafs = types.NewOrderedMap()
    outputDataRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputDataRate.Operator})
    outputDataRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputDataRate.Value})
    outputDataRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputDataRate.EndRangeValue})
    outputDataRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputDataRate.Percent})
    outputDataRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputDataRate.RearmType})
    outputDataRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputDataRate.RearmWindow})

    outputDataRate.EntityData.YListKeys = []string {}

    return &(outputDataRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate
// Number of input packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPacketRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputPacketRate.EntityData.Children = types.NewOrderedMap()
    inputPacketRate.EntityData.Leafs = types.NewOrderedMap()
    inputPacketRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputPacketRate.Operator})
    inputPacketRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputPacketRate.Value})
    inputPacketRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputPacketRate.EndRangeValue})
    inputPacketRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputPacketRate.Percent})
    inputPacketRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputPacketRate.RearmType})
    inputPacketRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputPacketRate.RearmWindow})

    inputPacketRate.EntityData.YListKeys = []string {}

    return &(inputPacketRate.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts
// Maximum number of output packets per second
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_OutputPeakPkts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputPeakPkts.EntityData.Children = types.NewOrderedMap()
    outputPeakPkts.EntityData.Leafs = types.NewOrderedMap()
    outputPeakPkts.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputPeakPkts.Operator})
    outputPeakPkts.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputPeakPkts.Value})
    outputPeakPkts.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputPeakPkts.EndRangeValue})
    outputPeakPkts.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputPeakPkts.Percent})
    outputPeakPkts.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputPeakPkts.RearmType})
    outputPeakPkts.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputPeakPkts.RearmWindow})

    outputPeakPkts.EntityData.YListKeys = []string {}

    return &(outputPeakPkts.EntityData)
}

// PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate
// Peak input data rate in kbps
// This type is a presence type.
type PerfMgmt_Threshold_DataRateInterface_DataRateInterfaceTemplates_DataRateInterfaceTemplate_InputPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputPeakRate.EntityData.Children = types.NewOrderedMap()
    inputPeakRate.EntityData.Leafs = types.NewOrderedMap()
    inputPeakRate.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputPeakRate.Operator})
    inputPeakRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputPeakRate.Value})
    inputPeakRate.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputPeakRate.EndRangeValue})
    inputPeakRate.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputPeakRate.Percent})
    inputPeakRate.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputPeakRate.RearmType})
    inputPeakRate.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputPeakRate.RearmWindow})

    inputPeakRate.EntityData.YListKeys = []string {}

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

    processNode.EntityData.Children = types.NewOrderedMap()
    processNode.EntityData.Children.Append("process-node-templates", types.YChild{"ProcessNodeTemplates", &processNode.ProcessNodeTemplates})
    processNode.EntityData.Leafs = types.NewOrderedMap()

    processNode.EntityData.YListKeys = []string {}

    return &(processNode.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates
// Node Memory threshold templates
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold template instance. The type is slice of
    // PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate.
    ProcessNodeTemplate []*PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
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

    processNodeTemplates.EntityData.Children = types.NewOrderedMap()
    processNodeTemplates.EntityData.Children.Append("process-node-template", types.YChild{"ProcessNodeTemplate", nil})
    for i := range processNodeTemplates.ProcessNodeTemplate {
        processNodeTemplates.EntityData.Children.Append(types.GetSegmentPath(processNodeTemplates.ProcessNodeTemplate[i]), types.YChild{"ProcessNodeTemplate", processNodeTemplates.ProcessNodeTemplate[i]})
    }
    processNodeTemplates.EntityData.Leafs = types.NewOrderedMap()

    processNodeTemplates.EntityData.YListKeys = []string {}

    return &(processNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate
// Node Memory threshold template instance
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate struct {
    EntityData types.CommonEntityData
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

func (processNodeTemplate *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate) GetEntityData() *types.CommonEntityData {
    processNodeTemplate.EntityData.YFilter = processNodeTemplate.YFilter
    processNodeTemplate.EntityData.YangName = "process-node-template"
    processNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    processNodeTemplate.EntityData.ParentYangName = "process-node-templates"
    processNodeTemplate.EntityData.SegmentPath = "process-node-template" + types.AddKeyToken(processNodeTemplate.TemplateName, "template-name")
    processNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processNodeTemplate.EntityData.Children = types.NewOrderedMap()
    processNodeTemplate.EntityData.Children.Append("average-cpu-used", types.YChild{"AverageCpuUsed", &processNodeTemplate.AverageCpuUsed})
    processNodeTemplate.EntityData.Children.Append("peak-memory", types.YChild{"PeakMemory", &processNodeTemplate.PeakMemory})
    processNodeTemplate.EntityData.Children.Append("no-threads", types.YChild{"NoThreads", &processNodeTemplate.NoThreads})
    processNodeTemplate.EntityData.Leafs = types.NewOrderedMap()
    processNodeTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", processNodeTemplate.TemplateName})
    processNodeTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", processNodeTemplate.SampleInterval})

    processNodeTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(processNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed
// Average %CPU utilization
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_AverageCpuUsed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    averageCpuUsed.EntityData.Children = types.NewOrderedMap()
    averageCpuUsed.EntityData.Leafs = types.NewOrderedMap()
    averageCpuUsed.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", averageCpuUsed.Operator})
    averageCpuUsed.EntityData.Leafs.Append("value", types.YLeaf{"Value", averageCpuUsed.Value})
    averageCpuUsed.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", averageCpuUsed.EndRangeValue})
    averageCpuUsed.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", averageCpuUsed.Percent})
    averageCpuUsed.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", averageCpuUsed.RearmType})
    averageCpuUsed.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", averageCpuUsed.RearmWindow})

    averageCpuUsed.EntityData.YListKeys = []string {}

    return &(averageCpuUsed.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory
// Max memory (KBytes) used since startup time
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (peakMemory *PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_PeakMemory) GetEntityData() *types.CommonEntityData {
    peakMemory.EntityData.YFilter = peakMemory.YFilter
    peakMemory.EntityData.YangName = "peak-memory"
    peakMemory.EntityData.BundleName = "cisco_ios_xr"
    peakMemory.EntityData.ParentYangName = "process-node-template"
    peakMemory.EntityData.SegmentPath = "peak-memory"
    peakMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakMemory.EntityData.Children = types.NewOrderedMap()
    peakMemory.EntityData.Leafs = types.NewOrderedMap()
    peakMemory.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", peakMemory.Operator})
    peakMemory.EntityData.Leafs.Append("value", types.YLeaf{"Value", peakMemory.Value})
    peakMemory.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", peakMemory.EndRangeValue})
    peakMemory.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", peakMemory.Percent})
    peakMemory.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", peakMemory.RearmType})
    peakMemory.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", peakMemory.RearmWindow})

    peakMemory.EntityData.YListKeys = []string {}

    return &(peakMemory.EntityData)
}

// PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads
// Number of threads
// This type is a presence type.
type PerfMgmt_Threshold_ProcessNode_ProcessNodeTemplates_ProcessNodeTemplate_NoThreads struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    noThreads.EntityData.Children = types.NewOrderedMap()
    noThreads.EntityData.Leafs = types.NewOrderedMap()
    noThreads.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", noThreads.Operator})
    noThreads.EntityData.Leafs.Append("value", types.YLeaf{"Value", noThreads.Value})
    noThreads.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", noThreads.EndRangeValue})
    noThreads.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", noThreads.Percent})
    noThreads.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", noThreads.RearmType})
    noThreads.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", noThreads.RearmWindow})

    noThreads.EntityData.YListKeys = []string {}

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

    memoryNode.EntityData.Children = types.NewOrderedMap()
    memoryNode.EntityData.Children.Append("memory-node-templates", types.YChild{"MemoryNodeTemplates", &memoryNode.MemoryNodeTemplates})
    memoryNode.EntityData.Leafs = types.NewOrderedMap()

    memoryNode.EntityData.YListKeys = []string {}

    return &(memoryNode.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates
// Node Memory threshold configuration templates
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Memory threshold configuration template instance. The type is slice of
    // PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate.
    MemoryNodeTemplate []*PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
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

    memoryNodeTemplates.EntityData.Children = types.NewOrderedMap()
    memoryNodeTemplates.EntityData.Children.Append("memory-node-template", types.YChild{"MemoryNodeTemplate", nil})
    for i := range memoryNodeTemplates.MemoryNodeTemplate {
        memoryNodeTemplates.EntityData.Children.Append(types.GetSegmentPath(memoryNodeTemplates.MemoryNodeTemplate[i]), types.YChild{"MemoryNodeTemplate", memoryNodeTemplates.MemoryNodeTemplate[i]})
    }
    memoryNodeTemplates.EntityData.Leafs = types.NewOrderedMap()

    memoryNodeTemplates.EntityData.YListKeys = []string {}

    return &(memoryNodeTemplates.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate
// Node Memory threshold configuration template
// instance
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate struct {
    EntityData types.CommonEntityData
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

func (memoryNodeTemplate *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate) GetEntityData() *types.CommonEntityData {
    memoryNodeTemplate.EntityData.YFilter = memoryNodeTemplate.YFilter
    memoryNodeTemplate.EntityData.YangName = "memory-node-template"
    memoryNodeTemplate.EntityData.BundleName = "cisco_ios_xr"
    memoryNodeTemplate.EntityData.ParentYangName = "memory-node-templates"
    memoryNodeTemplate.EntityData.SegmentPath = "memory-node-template" + types.AddKeyToken(memoryNodeTemplate.TemplateName, "template-name")
    memoryNodeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memoryNodeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memoryNodeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memoryNodeTemplate.EntityData.Children = types.NewOrderedMap()
    memoryNodeTemplate.EntityData.Children.Append("peak-memory", types.YChild{"PeakMemory", &memoryNodeTemplate.PeakMemory})
    memoryNodeTemplate.EntityData.Children.Append("curr-memory", types.YChild{"CurrMemory", &memoryNodeTemplate.CurrMemory})
    memoryNodeTemplate.EntityData.Leafs = types.NewOrderedMap()
    memoryNodeTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", memoryNodeTemplate.TemplateName})
    memoryNodeTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", memoryNodeTemplate.SampleInterval})

    memoryNodeTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(memoryNodeTemplate.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory
// Maximum memory (KBytes) used
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_PeakMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    peakMemory.EntityData.Children = types.NewOrderedMap()
    peakMemory.EntityData.Leafs = types.NewOrderedMap()
    peakMemory.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", peakMemory.Operator})
    peakMemory.EntityData.Leafs.Append("value", types.YLeaf{"Value", peakMemory.Value})
    peakMemory.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", peakMemory.EndRangeValue})
    peakMemory.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", peakMemory.Percent})
    peakMemory.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", peakMemory.RearmType})
    peakMemory.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", peakMemory.RearmWindow})

    peakMemory.EntityData.YListKeys = []string {}

    return &(peakMemory.EntityData)
}

// PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory
// Current memory (Bytes) in use
// This type is a presence type.
type PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (currMemory *PerfMgmt_Threshold_MemoryNode_MemoryNodeTemplates_MemoryNodeTemplate_CurrMemory) GetEntityData() *types.CommonEntityData {
    currMemory.EntityData.YFilter = currMemory.YFilter
    currMemory.EntityData.YangName = "curr-memory"
    currMemory.EntityData.BundleName = "cisco_ios_xr"
    currMemory.EntityData.ParentYangName = "memory-node-template"
    currMemory.EntityData.SegmentPath = "curr-memory"
    currMemory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currMemory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currMemory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currMemory.EntityData.Children = types.NewOrderedMap()
    currMemory.EntityData.Leafs = types.NewOrderedMap()
    currMemory.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", currMemory.Operator})
    currMemory.EntityData.Leafs.Append("value", types.YLeaf{"Value", currMemory.Value})
    currMemory.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", currMemory.EndRangeValue})
    currMemory.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", currMemory.Percent})
    currMemory.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", currMemory.RearmType})
    currMemory.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", currMemory.RearmWindow})

    currMemory.EntityData.YListKeys = []string {}

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

    ospfv3Protocol.EntityData.Children = types.NewOrderedMap()
    ospfv3Protocol.EntityData.Children.Append("ospfv3-protocol-templates", types.YChild{"Ospfv3ProtocolTemplates", &ospfv3Protocol.Ospfv3ProtocolTemplates})
    ospfv3Protocol.EntityData.Leafs = types.NewOrderedMap()

    ospfv3Protocol.EntityData.YListKeys = []string {}

    return &(ospfv3Protocol.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates
// OSPF v2 Protocol threshold templates
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF v2 Protocol threshold template instance. The type is slice of
    // PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate.
    Ospfv3ProtocolTemplate []*PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
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

    ospfv3ProtocolTemplates.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolTemplates.EntityData.Children.Append("ospfv3-protocol-template", types.YChild{"Ospfv3ProtocolTemplate", nil})
    for i := range ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate {
        ospfv3ProtocolTemplates.EntityData.Children.Append(types.GetSegmentPath(ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i]), types.YChild{"Ospfv3ProtocolTemplate", ospfv3ProtocolTemplates.Ospfv3ProtocolTemplate[i]})
    }
    ospfv3ProtocolTemplates.EntityData.Leafs = types.NewOrderedMap()

    ospfv3ProtocolTemplates.EntityData.YListKeys = []string {}

    return &(ospfv3ProtocolTemplates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate
// OSPF v2 Protocol threshold template instance
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate struct {
    EntityData types.CommonEntityData
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

func (ospfv3ProtocolTemplate *PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate) GetEntityData() *types.CommonEntityData {
    ospfv3ProtocolTemplate.EntityData.YFilter = ospfv3ProtocolTemplate.YFilter
    ospfv3ProtocolTemplate.EntityData.YangName = "ospfv3-protocol-template"
    ospfv3ProtocolTemplate.EntityData.BundleName = "cisco_ios_xr"
    ospfv3ProtocolTemplate.EntityData.ParentYangName = "ospfv3-protocol-templates"
    ospfv3ProtocolTemplate.EntityData.SegmentPath = "ospfv3-protocol-template" + types.AddKeyToken(ospfv3ProtocolTemplate.TemplateName, "template-name")
    ospfv3ProtocolTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3ProtocolTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3ProtocolTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3ProtocolTemplate.EntityData.Children = types.NewOrderedMap()
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-lsa-acks-lsa", types.YChild{"InputLsaAcksLsa", &ospfv3ProtocolTemplate.InputLsaAcksLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-db-ds-lsa", types.YChild{"OutputDbDsLsa", &ospfv3ProtocolTemplate.OutputDbDsLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-db-ds-lsa", types.YChild{"InputDbDsLsa", &ospfv3ProtocolTemplate.InputDbDsLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-lsa-updates", types.YChild{"InputLsaUpdates", &ospfv3ProtocolTemplate.InputLsaUpdates})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-db-ds", types.YChild{"OutputDbDs", &ospfv3ProtocolTemplate.OutputDbDs})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-lsa-updates-lsa", types.YChild{"OutputLsaUpdatesLsa", &ospfv3ProtocolTemplate.OutputLsaUpdatesLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-db-ds", types.YChild{"InputDbDs", &ospfv3ProtocolTemplate.InputDbDs})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-lsa-updates-lsa", types.YChild{"InputLsaUpdatesLsa", &ospfv3ProtocolTemplate.InputLsaUpdatesLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-packets", types.YChild{"OutputPackets", &ospfv3ProtocolTemplate.OutputPackets})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-packets", types.YChild{"InputPackets", &ospfv3ProtocolTemplate.InputPackets})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-hello-packets", types.YChild{"OutputHelloPackets", &ospfv3ProtocolTemplate.OutputHelloPackets})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-hello-packets", types.YChild{"InputHelloPackets", &ospfv3ProtocolTemplate.InputHelloPackets})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-ls-requests", types.YChild{"OutputLsRequests", &ospfv3ProtocolTemplate.OutputLsRequests})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-lsa-acks-lsa", types.YChild{"OutputLsaAcksLsa", &ospfv3ProtocolTemplate.OutputLsaAcksLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-lsa-acks", types.YChild{"OutputLsaAcks", &ospfv3ProtocolTemplate.OutputLsaAcks})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-lsa-acks", types.YChild{"InputLsaAcks", &ospfv3ProtocolTemplate.InputLsaAcks})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-lsa-updates", types.YChild{"OutputLsaUpdates", &ospfv3ProtocolTemplate.OutputLsaUpdates})
    ospfv3ProtocolTemplate.EntityData.Children.Append("output-ls-requests-lsa", types.YChild{"OutputLsRequestsLsa", &ospfv3ProtocolTemplate.OutputLsRequestsLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-ls-requests-lsa", types.YChild{"InputLsRequestsLsa", &ospfv3ProtocolTemplate.InputLsRequestsLsa})
    ospfv3ProtocolTemplate.EntityData.Children.Append("input-ls-requests", types.YChild{"InputLsRequests", &ospfv3ProtocolTemplate.InputLsRequests})
    ospfv3ProtocolTemplate.EntityData.Leafs = types.NewOrderedMap()
    ospfv3ProtocolTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ospfv3ProtocolTemplate.TemplateName})
    ospfv3ProtocolTemplate.EntityData.Leafs.Append("sample-interval", types.YLeaf{"SampleInterval", ospfv3ProtocolTemplate.SampleInterval})

    ospfv3ProtocolTemplate.EntityData.YListKeys = []string {"TemplateName"}

    return &(ospfv3ProtocolTemplate.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa
// Number of LSA received in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcksLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaAcksLsa.EntityData.Children = types.NewOrderedMap()
    inputLsaAcksLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsaAcksLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaAcksLsa.Operator})
    inputLsaAcksLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaAcksLsa.Value})
    inputLsaAcksLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaAcksLsa.EndRangeValue})
    inputLsaAcksLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaAcksLsa.Percent})
    inputLsaAcksLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaAcksLsa.RearmType})
    inputLsaAcksLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaAcksLsa.RearmWindow})

    inputLsaAcksLsa.EntityData.YListKeys = []string {}

    return &(inputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa
// Number of LSA sent in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputDbDsLsa.EntityData.Children = types.NewOrderedMap()
    outputDbDsLsa.EntityData.Leafs = types.NewOrderedMap()
    outputDbDsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputDbDsLsa.Operator})
    outputDbDsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputDbDsLsa.Value})
    outputDbDsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputDbDsLsa.EndRangeValue})
    outputDbDsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputDbDsLsa.Percent})
    outputDbDsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputDbDsLsa.RearmType})
    outputDbDsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputDbDsLsa.RearmWindow})

    outputDbDsLsa.EntityData.YListKeys = []string {}

    return &(outputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa
// Number of LSA received in DBD packets
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputDbDsLsa.EntityData.Children = types.NewOrderedMap()
    inputDbDsLsa.EntityData.Leafs = types.NewOrderedMap()
    inputDbDsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputDbDsLsa.Operator})
    inputDbDsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputDbDsLsa.Value})
    inputDbDsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputDbDsLsa.EndRangeValue})
    inputDbDsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputDbDsLsa.Percent})
    inputDbDsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputDbDsLsa.RearmType})
    inputDbDsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputDbDsLsa.RearmWindow})

    inputDbDsLsa.EntityData.YListKeys = []string {}

    return &(inputDbDsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates
// Number of LSA Updates received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaUpdates.EntityData.Children = types.NewOrderedMap()
    inputLsaUpdates.EntityData.Leafs = types.NewOrderedMap()
    inputLsaUpdates.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaUpdates.Operator})
    inputLsaUpdates.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaUpdates.Value})
    inputLsaUpdates.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaUpdates.EndRangeValue})
    inputLsaUpdates.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaUpdates.Percent})
    inputLsaUpdates.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaUpdates.RearmType})
    inputLsaUpdates.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaUpdates.RearmWindow})

    inputLsaUpdates.EntityData.YListKeys = []string {}

    return &(inputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs
// Number of DBD packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputDbDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputDbDs.EntityData.Children = types.NewOrderedMap()
    outputDbDs.EntityData.Leafs = types.NewOrderedMap()
    outputDbDs.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputDbDs.Operator})
    outputDbDs.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputDbDs.Value})
    outputDbDs.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputDbDs.EndRangeValue})
    outputDbDs.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputDbDs.Percent})
    outputDbDs.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputDbDs.RearmType})
    outputDbDs.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputDbDs.RearmWindow})

    outputDbDs.EntityData.YListKeys = []string {}

    return &(outputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa
// Number of LSA sent in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaUpdatesLsa.EntityData.Children = types.NewOrderedMap()
    outputLsaUpdatesLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsaUpdatesLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaUpdatesLsa.Operator})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaUpdatesLsa.Value})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaUpdatesLsa.EndRangeValue})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaUpdatesLsa.Percent})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaUpdatesLsa.RearmType})
    outputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaUpdatesLsa.RearmWindow})

    outputLsaUpdatesLsa.EntityData.YListKeys = []string {}

    return &(outputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs
// Number of DBD packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputDbDs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputDbDs.EntityData.Children = types.NewOrderedMap()
    inputDbDs.EntityData.Leafs = types.NewOrderedMap()
    inputDbDs.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputDbDs.Operator})
    inputDbDs.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputDbDs.Value})
    inputDbDs.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputDbDs.EndRangeValue})
    inputDbDs.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputDbDs.Percent})
    inputDbDs.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputDbDs.RearmType})
    inputDbDs.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputDbDs.RearmWindow})

    inputDbDs.EntityData.YListKeys = []string {}

    return &(inputDbDs.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa
// Number of LSA received in LSA Updates
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaUpdatesLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaUpdatesLsa.EntityData.Children = types.NewOrderedMap()
    inputLsaUpdatesLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsaUpdatesLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaUpdatesLsa.Operator})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaUpdatesLsa.Value})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaUpdatesLsa.EndRangeValue})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaUpdatesLsa.Percent})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaUpdatesLsa.RearmType})
    inputLsaUpdatesLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaUpdatesLsa.RearmWindow})

    inputLsaUpdatesLsa.EntityData.YListKeys = []string {}

    return &(inputLsaUpdatesLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputPackets.EntityData.Children = types.NewOrderedMap()
    outputPackets.EntityData.Leafs = types.NewOrderedMap()
    outputPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputPackets.Operator})
    outputPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputPackets.Value})
    outputPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputPackets.EndRangeValue})
    outputPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputPackets.Percent})
    outputPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputPackets.RearmType})
    outputPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputPackets.RearmWindow})

    outputPackets.EntityData.YListKeys = []string {}

    return &(outputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets
// Total number of packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputPackets.EntityData.Children = types.NewOrderedMap()
    inputPackets.EntityData.Leafs = types.NewOrderedMap()
    inputPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputPackets.Operator})
    inputPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputPackets.Value})
    inputPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputPackets.EndRangeValue})
    inputPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputPackets.Percent})
    inputPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputPackets.RearmType})
    inputPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputPackets.RearmWindow})

    inputPackets.EntityData.YListKeys = []string {}

    return &(inputPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets
// Total number of packets sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputHelloPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputHelloPackets.EntityData.Children = types.NewOrderedMap()
    outputHelloPackets.EntityData.Leafs = types.NewOrderedMap()
    outputHelloPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputHelloPackets.Operator})
    outputHelloPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputHelloPackets.Value})
    outputHelloPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputHelloPackets.EndRangeValue})
    outputHelloPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputHelloPackets.Percent})
    outputHelloPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputHelloPackets.RearmType})
    outputHelloPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputHelloPackets.RearmWindow})

    outputHelloPackets.EntityData.YListKeys = []string {}

    return &(outputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets
// Number of Hello packets received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputHelloPackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputHelloPackets.EntityData.Children = types.NewOrderedMap()
    inputHelloPackets.EntityData.Leafs = types.NewOrderedMap()
    inputHelloPackets.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputHelloPackets.Operator})
    inputHelloPackets.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputHelloPackets.Value})
    inputHelloPackets.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputHelloPackets.EndRangeValue})
    inputHelloPackets.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputHelloPackets.Percent})
    inputHelloPackets.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputHelloPackets.RearmType})
    inputHelloPackets.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputHelloPackets.RearmWindow})

    inputHelloPackets.EntityData.YListKeys = []string {}

    return &(inputHelloPackets.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests
// Number of LS Requests sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsRequests.EntityData.Children = types.NewOrderedMap()
    outputLsRequests.EntityData.Leafs = types.NewOrderedMap()
    outputLsRequests.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsRequests.Operator})
    outputLsRequests.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsRequests.Value})
    outputLsRequests.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsRequests.EndRangeValue})
    outputLsRequests.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsRequests.Percent})
    outputLsRequests.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsRequests.RearmType})
    outputLsRequests.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsRequests.RearmWindow})

    outputLsRequests.EntityData.YListKeys = []string {}

    return &(outputLsRequests.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa
// Number of LSA sent in LSA Acknowledgements
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcksLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaAcksLsa.EntityData.Children = types.NewOrderedMap()
    outputLsaAcksLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsaAcksLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaAcksLsa.Operator})
    outputLsaAcksLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaAcksLsa.Value})
    outputLsaAcksLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaAcksLsa.EndRangeValue})
    outputLsaAcksLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaAcksLsa.Percent})
    outputLsaAcksLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaAcksLsa.RearmType})
    outputLsaAcksLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaAcksLsa.RearmWindow})

    outputLsaAcksLsa.EntityData.YListKeys = []string {}

    return &(outputLsaAcksLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks
// Number of LSA Acknowledgements sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaAcks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaAcks.EntityData.Children = types.NewOrderedMap()
    outputLsaAcks.EntityData.Leafs = types.NewOrderedMap()
    outputLsaAcks.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaAcks.Operator})
    outputLsaAcks.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaAcks.Value})
    outputLsaAcks.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaAcks.EndRangeValue})
    outputLsaAcks.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaAcks.Percent})
    outputLsaAcks.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaAcks.RearmType})
    outputLsaAcks.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaAcks.RearmWindow})

    outputLsaAcks.EntityData.YListKeys = []string {}

    return &(outputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks
// Number of LSA Acknowledgements received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsaAcks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsaAcks.EntityData.Children = types.NewOrderedMap()
    inputLsaAcks.EntityData.Leafs = types.NewOrderedMap()
    inputLsaAcks.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsaAcks.Operator})
    inputLsaAcks.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsaAcks.Value})
    inputLsaAcks.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsaAcks.EndRangeValue})
    inputLsaAcks.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsaAcks.Percent})
    inputLsaAcks.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsaAcks.RearmType})
    inputLsaAcks.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsaAcks.RearmWindow})

    inputLsaAcks.EntityData.YListKeys = []string {}

    return &(inputLsaAcks.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates
// Number of LSA Updates sent
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsaUpdates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsaUpdates.EntityData.Children = types.NewOrderedMap()
    outputLsaUpdates.EntityData.Leafs = types.NewOrderedMap()
    outputLsaUpdates.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsaUpdates.Operator})
    outputLsaUpdates.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsaUpdates.Value})
    outputLsaUpdates.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsaUpdates.EndRangeValue})
    outputLsaUpdates.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsaUpdates.Percent})
    outputLsaUpdates.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsaUpdates.RearmType})
    outputLsaUpdates.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsaUpdates.RearmWindow})

    outputLsaUpdates.EntityData.YListKeys = []string {}

    return &(outputLsaUpdates.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa
// Number of LSA sent in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_OutputLsRequestsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    outputLsRequestsLsa.EntityData.Children = types.NewOrderedMap()
    outputLsRequestsLsa.EntityData.Leafs = types.NewOrderedMap()
    outputLsRequestsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", outputLsRequestsLsa.Operator})
    outputLsRequestsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", outputLsRequestsLsa.Value})
    outputLsRequestsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", outputLsRequestsLsa.EndRangeValue})
    outputLsRequestsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", outputLsRequestsLsa.Percent})
    outputLsRequestsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", outputLsRequestsLsa.RearmType})
    outputLsRequestsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", outputLsRequestsLsa.RearmWindow})

    outputLsRequestsLsa.EntityData.YListKeys = []string {}

    return &(outputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa
// Number of LSA received in LS Requests
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequestsLsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsRequestsLsa.EntityData.Children = types.NewOrderedMap()
    inputLsRequestsLsa.EntityData.Leafs = types.NewOrderedMap()
    inputLsRequestsLsa.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsRequestsLsa.Operator})
    inputLsRequestsLsa.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsRequestsLsa.Value})
    inputLsRequestsLsa.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsRequestsLsa.EndRangeValue})
    inputLsRequestsLsa.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsRequestsLsa.Percent})
    inputLsRequestsLsa.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsRequestsLsa.RearmType})
    inputLsRequestsLsa.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsRequestsLsa.RearmWindow})

    inputLsRequestsLsa.EntityData.YListKeys = []string {}

    return &(inputLsRequestsLsa.EntityData)
}

// PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests
// Number of LS Requests received
// This type is a presence type.
type PerfMgmt_Threshold_Ospfv3Protocol_Ospfv3ProtocolTemplates_Ospfv3ProtocolTemplate_InputLsRequests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

    inputLsRequests.EntityData.Children = types.NewOrderedMap()
    inputLsRequests.EntityData.Leafs = types.NewOrderedMap()
    inputLsRequests.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", inputLsRequests.Operator})
    inputLsRequests.EntityData.Leafs.Append("value", types.YLeaf{"Value", inputLsRequests.Value})
    inputLsRequests.EntityData.Leafs.Append("end-range-value", types.YLeaf{"EndRangeValue", inputLsRequests.EndRangeValue})
    inputLsRequests.EntityData.Leafs.Append("percent", types.YLeaf{"Percent", inputLsRequests.Percent})
    inputLsRequests.EntityData.Leafs.Append("rearm-type", types.YLeaf{"RearmType", inputLsRequests.RearmType})
    inputLsRequests.EntityData.Leafs.Append("rearm-window", types.YLeaf{"RearmWindow", inputLsRequests.RearmWindow})

    inputLsRequests.EntityData.YListKeys = []string {}

    return &(inputLsRequests.EntityData)
}

