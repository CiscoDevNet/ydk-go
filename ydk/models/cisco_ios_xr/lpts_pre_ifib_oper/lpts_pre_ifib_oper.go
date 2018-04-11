// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-pre-ifib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lpts-pifib: lpts pre-ifib data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_pre_ifib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_pre_ifib_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-pre-ifib-oper lpts-pifib}", reflect.TypeOf(LptsPifib_{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib", reflect.TypeOf(LptsPifib_{}))
}

// LptsPifib represents Lpts pifib
type LptsPifib string

const (
    // ISIS packets
    LptsPifib_isis LptsPifib = "isis"

    // IPv4 fragmented packets
    LptsPifib_ipv4_frag LptsPifib = "ipv4-frag"

    // IPv4 ICMP Echo packets
    LptsPifib_ipv4_echo LptsPifib = "ipv4-echo"

    // All IPv4 packets
    LptsPifib_ipv4_any LptsPifib = "ipv4-any"

    // IPv6 fragmented packets
    LptsPifib_ipv6_frag LptsPifib = "ipv6-frag"

    // IPv6 ICMP Echo packets
    LptsPifib_ipv6_echo LptsPifib = "ipv6-echo"

    // IPv6 ND packets
    LptsPifib_ipv6_nd LptsPifib = "ipv6-nd"

    // All IPv6 packets
    LptsPifib_ipv6_any LptsPifib = "ipv6-any"

    // BFD packets
    LptsPifib_bfd_any LptsPifib = "bfd-any"

    // All packets
    LptsPifib_all LptsPifib = "all"
)

// LptsPifib_
// lpts pre-ifib data
type LptsPifib_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Pre-ifib Nodes.
    Nodes LptsPifib__Nodes
}

func (lptsPifib_ *LptsPifib_) GetEntityData() *types.CommonEntityData {
    lptsPifib_.EntityData.YFilter = lptsPifib_.YFilter
    lptsPifib_.EntityData.YangName = "lpts-pifib"
    lptsPifib_.EntityData.BundleName = "cisco_ios_xr"
    lptsPifib_.EntityData.ParentYangName = "Cisco-IOS-XR-lpts-pre-ifib-oper"
    lptsPifib_.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib"
    lptsPifib_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsPifib_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsPifib_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsPifib_.EntityData.Children = make(map[string]types.YChild)
    lptsPifib_.EntityData.Children["nodes"] = types.YChild{"Nodes", &lptsPifib_.Nodes}
    lptsPifib_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lptsPifib_.EntityData)
}

// LptsPifib__Nodes
// List of Pre-ifib Nodes
type LptsPifib__Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pre-ifib data for particular node. The type is slice of
    // LptsPifib__Nodes_Node.
    Node []LptsPifib__Nodes_Node
}

func (nodes *LptsPifib__Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "lpts-pifib"
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

// LptsPifib__Nodes_Node
// Pre-ifib data for particular node
type LptsPifib__Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Type specific.
    TypeValues LptsPifib__Nodes_Node_TypeValues

    // Dynamic Flows Statistics.
    DynamicFlowsStats LptsPifib__Nodes_Node_DynamicFlowsStats

    // Hardware specific.
    Hardware LptsPifib__Nodes_Node_Hardware
}

func (node *LptsPifib__Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["type-values"] = types.YChild{"TypeValues", &node.TypeValues}
    node.EntityData.Children["dynamic-flows-stats"] = types.YChild{"DynamicFlowsStats", &node.DynamicFlowsStats}
    node.EntityData.Children["Cisco-IOS-XR-platform-pifib-oper:hardware"] = types.YChild{"Hardware", &node.Hardware}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// LptsPifib__Nodes_Node_TypeValues
// Type specific
type LptsPifib__Nodes_Node_TypeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // pifib types. The type is slice of
    // LptsPifib__Nodes_Node_TypeValues_TypeValue.
    TypeValue []LptsPifib__Nodes_Node_TypeValues_TypeValue
}

func (typeValues *LptsPifib__Nodes_Node_TypeValues) GetEntityData() *types.CommonEntityData {
    typeValues.EntityData.YFilter = typeValues.YFilter
    typeValues.EntityData.YangName = "type-values"
    typeValues.EntityData.BundleName = "cisco_ios_xr"
    typeValues.EntityData.ParentYangName = "node"
    typeValues.EntityData.SegmentPath = "type-values"
    typeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeValues.EntityData.Children = make(map[string]types.YChild)
    typeValues.EntityData.Children["type-value"] = types.YChild{"TypeValue", nil}
    for i := range typeValues.TypeValue {
        typeValues.EntityData.Children[types.GetSegmentPath(&typeValues.TypeValue[i])] = types.YChild{"TypeValue", &typeValues.TypeValue[i]}
    }
    typeValues.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(typeValues.EntityData)
}

// LptsPifib__Nodes_Node_TypeValues_TypeValue
// pifib types
type LptsPifib__Nodes_Node_TypeValues_TypeValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Type value. The type is LptsPifib.
    PifibType interface{}

    // Data for single pre-ifib entry. The type is slice of
    // LptsPifib__Nodes_Node_TypeValues_TypeValue_Entry.
    Entry []LptsPifib__Nodes_Node_TypeValues_TypeValue_Entry
}

func (typeValue *LptsPifib__Nodes_Node_TypeValues_TypeValue) GetEntityData() *types.CommonEntityData {
    typeValue.EntityData.YFilter = typeValue.YFilter
    typeValue.EntityData.YangName = "type-value"
    typeValue.EntityData.BundleName = "cisco_ios_xr"
    typeValue.EntityData.ParentYangName = "type-values"
    typeValue.EntityData.SegmentPath = "type-value" + "[pifib-type='" + fmt.Sprintf("%v", typeValue.PifibType) + "']"
    typeValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeValue.EntityData.Children = make(map[string]types.YChild)
    typeValue.EntityData.Children["entry"] = types.YChild{"Entry", nil}
    for i := range typeValue.Entry {
        typeValue.EntityData.Children[types.GetSegmentPath(&typeValue.Entry[i])] = types.YChild{"Entry", &typeValue.Entry[i]}
    }
    typeValue.EntityData.Leafs = make(map[string]types.YLeaf)
    typeValue.EntityData.Leafs["pifib-type"] = types.YLeaf{"PifibType", typeValue.PifibType}
    return &(typeValue.EntityData)
}

// LptsPifib__Nodes_Node_TypeValues_TypeValue_Entry
// Data for single pre-ifib entry
type LptsPifib__Nodes_Node_TypeValues_TypeValue_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Single Pre-ifib entry. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Entry interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    Vid interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3Protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Interface Name. The type is string.
    IntfName interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    IntfHandle interface{}

    // Destination IP Address. The type is string.
    DestinationAddr interface{}

    // Source IP Address. The type is string.
    SourceAddr interface{}

    // Destination Key Type. The type is string.
    DestinationType interface{}

    // Destination Port/ICMP Type/IGMP Type. The type is string.
    DestinationValue interface{}

    // Source port. The type is string.
    SourcePort interface{}

    // Is Fragment. The type is interface{} with range: 0..255.
    IsFrag interface{}

    // Is SYN. The type is interface{} with range: 0..255.
    IsSyn interface{}

    // Opcode. The type is string.
    Opcode interface{}

    // Flow type. The type is string.
    FlowType interface{}

    // Listener Tag. The type is string.
    ListenerTag interface{}

    // Local Flag. The type is interface{} with range: 0..255.
    LocalFlag interface{}

    // Is FGID or not. The type is interface{} with range: 0..255.
    IsFgid interface{}

    // Deliver List Short Format. The type is string.
    DeliverListShort interface{}

    // Deliver List Long Format. The type is string.
    DeliverListLong interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    MinTtl interface{}

    // Packets matched to accept. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepts interface{}

    // Packets matched for drop. The type is interface{} with range:
    // 0..18446744073709551615.
    Drops interface{}

    // Is Stale. The type is interface{} with range: 0..255.
    Stale interface{}

    // sub Pre-IFIB type. The type is interface{} with range: 0..255.
    PifibType interface{}

    // Creation or Update Time. The type is string.
    PifibProgramTime interface{}
}

func (entry *LptsPifib__Nodes_Node_TypeValues_TypeValue_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "type-value"
    entry.EntityData.SegmentPath = "entry" + "[entry='" + fmt.Sprintf("%v", entry.Entry) + "']"
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = make(map[string]types.YChild)
    entry.EntityData.Leafs = make(map[string]types.YLeaf)
    entry.EntityData.Leafs["entry"] = types.YLeaf{"Entry", entry.Entry}
    entry.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", entry.VrfName}
    entry.EntityData.Leafs["vid"] = types.YLeaf{"Vid", entry.Vid}
    entry.EntityData.Leafs["l3protocol"] = types.YLeaf{"L3Protocol", entry.L3Protocol}
    entry.EntityData.Leafs["l4protocol"] = types.YLeaf{"L4Protocol", entry.L4Protocol}
    entry.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", entry.IntfName}
    entry.EntityData.Leafs["intf-handle"] = types.YLeaf{"IntfHandle", entry.IntfHandle}
    entry.EntityData.Leafs["destination-addr"] = types.YLeaf{"DestinationAddr", entry.DestinationAddr}
    entry.EntityData.Leafs["source-addr"] = types.YLeaf{"SourceAddr", entry.SourceAddr}
    entry.EntityData.Leafs["destination-type"] = types.YLeaf{"DestinationType", entry.DestinationType}
    entry.EntityData.Leafs["destination-value"] = types.YLeaf{"DestinationValue", entry.DestinationValue}
    entry.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", entry.SourcePort}
    entry.EntityData.Leafs["is-frag"] = types.YLeaf{"IsFrag", entry.IsFrag}
    entry.EntityData.Leafs["is-syn"] = types.YLeaf{"IsSyn", entry.IsSyn}
    entry.EntityData.Leafs["opcode"] = types.YLeaf{"Opcode", entry.Opcode}
    entry.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", entry.FlowType}
    entry.EntityData.Leafs["listener-tag"] = types.YLeaf{"ListenerTag", entry.ListenerTag}
    entry.EntityData.Leafs["local-flag"] = types.YLeaf{"LocalFlag", entry.LocalFlag}
    entry.EntityData.Leafs["is-fgid"] = types.YLeaf{"IsFgid", entry.IsFgid}
    entry.EntityData.Leafs["deliver-list-short"] = types.YLeaf{"DeliverListShort", entry.DeliverListShort}
    entry.EntityData.Leafs["deliver-list-long"] = types.YLeaf{"DeliverListLong", entry.DeliverListLong}
    entry.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", entry.MinTtl}
    entry.EntityData.Leafs["accepts"] = types.YLeaf{"Accepts", entry.Accepts}
    entry.EntityData.Leafs["drops"] = types.YLeaf{"Drops", entry.Drops}
    entry.EntityData.Leafs["stale"] = types.YLeaf{"Stale", entry.Stale}
    entry.EntityData.Leafs["pifib-type"] = types.YLeaf{"PifibType", entry.PifibType}
    entry.EntityData.Leafs["pifib-program-time"] = types.YLeaf{"PifibProgramTime", entry.PifibProgramTime}
    return &(entry.EntityData)
}

// LptsPifib__Nodes_Node_DynamicFlowsStats
// Dynamic Flows Statistics
type LptsPifib__Nodes_Node_DynamicFlowsStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dynamic Flows Enabled. The type is bool.
    DynamicFlowsEnabled interface{}

    // Platform Max. The type is interface{} with range: 0..4294967295.
    PlatformSupportedMax interface{}

    // Platform Config Limit. The type is interface{} with range: 0..4294967295.
    PlatformConfiguredMax interface{}

    // Platform Total Configured. The type is interface{} with range:
    // 0..4294967295.
    PlatformTotalConfigured interface{}

    // Total HW Entries. The type is interface{} with range: 0..4294967295.
    TotalHwEntries interface{}

    // Total SW Entries. The type is interface{} with range: 0..4294967295.
    TotalSwEntries interface{}

    // Flow Datalist. The type is slice of
    // LptsPifib__Nodes_Node_DynamicFlowsStats_Flow.
    Flow []LptsPifib__Nodes_Node_DynamicFlowsStats_Flow
}

func (dynamicFlowsStats *LptsPifib__Nodes_Node_DynamicFlowsStats) GetEntityData() *types.CommonEntityData {
    dynamicFlowsStats.EntityData.YFilter = dynamicFlowsStats.YFilter
    dynamicFlowsStats.EntityData.YangName = "dynamic-flows-stats"
    dynamicFlowsStats.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsStats.EntityData.ParentYangName = "node"
    dynamicFlowsStats.EntityData.SegmentPath = "dynamic-flows-stats"
    dynamicFlowsStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsStats.EntityData.Children = make(map[string]types.YChild)
    dynamicFlowsStats.EntityData.Children["flow"] = types.YChild{"Flow", nil}
    for i := range dynamicFlowsStats.Flow {
        dynamicFlowsStats.EntityData.Children[types.GetSegmentPath(&dynamicFlowsStats.Flow[i])] = types.YChild{"Flow", &dynamicFlowsStats.Flow[i]}
    }
    dynamicFlowsStats.EntityData.Leafs = make(map[string]types.YLeaf)
    dynamicFlowsStats.EntityData.Leafs["dynamic-flows-enabled"] = types.YLeaf{"DynamicFlowsEnabled", dynamicFlowsStats.DynamicFlowsEnabled}
    dynamicFlowsStats.EntityData.Leafs["platform-supported-max"] = types.YLeaf{"PlatformSupportedMax", dynamicFlowsStats.PlatformSupportedMax}
    dynamicFlowsStats.EntityData.Leafs["platform-configured-max"] = types.YLeaf{"PlatformConfiguredMax", dynamicFlowsStats.PlatformConfiguredMax}
    dynamicFlowsStats.EntityData.Leafs["platform-total-configured"] = types.YLeaf{"PlatformTotalConfigured", dynamicFlowsStats.PlatformTotalConfigured}
    dynamicFlowsStats.EntityData.Leafs["total-hw-entries"] = types.YLeaf{"TotalHwEntries", dynamicFlowsStats.TotalHwEntries}
    dynamicFlowsStats.EntityData.Leafs["total-sw-entries"] = types.YLeaf{"TotalSwEntries", dynamicFlowsStats.TotalSwEntries}
    return &(dynamicFlowsStats.EntityData)
}

// LptsPifib__Nodes_Node_DynamicFlowsStats_Flow
// Flow Datalist
type LptsPifib__Nodes_Node_DynamicFlowsStats_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Name. The type is string.
    FlowName interface{}

    // Is Configurable. The type is bool.
    Configurable interface{}

    // Is Configured. The type is bool.
    Configured interface{}

    // Default Max. The type is interface{} with range: 0..4294967295.
    DefaultMax interface{}

    // Configured Max. The type is string.
    ConfiguredMax interface{}

    // Active Max. The type is interface{} with range: 0..4294967295.
    ActiveMax interface{}

    // Hardware Count. The type is interface{} with range: 0..4294967295.
    HardwareCount interface{}

    // Software Count. The type is interface{} with range: 0..4294967295.
    SoftwareCount interface{}

    // Pending Software Entries. The type is bool.
    PendingSoftwareEntries interface{}
}

func (flow *LptsPifib__Nodes_Node_DynamicFlowsStats_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "dynamic-flows-stats"
    flow.EntityData.SegmentPath = "flow"
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = make(map[string]types.YChild)
    flow.EntityData.Leafs = make(map[string]types.YLeaf)
    flow.EntityData.Leafs["flow-name"] = types.YLeaf{"FlowName", flow.FlowName}
    flow.EntityData.Leafs["configurable"] = types.YLeaf{"Configurable", flow.Configurable}
    flow.EntityData.Leafs["configured"] = types.YLeaf{"Configured", flow.Configured}
    flow.EntityData.Leafs["default-max"] = types.YLeaf{"DefaultMax", flow.DefaultMax}
    flow.EntityData.Leafs["configured-max"] = types.YLeaf{"ConfiguredMax", flow.ConfiguredMax}
    flow.EntityData.Leafs["active-max"] = types.YLeaf{"ActiveMax", flow.ActiveMax}
    flow.EntityData.Leafs["hardware-count"] = types.YLeaf{"HardwareCount", flow.HardwareCount}
    flow.EntityData.Leafs["software-count"] = types.YLeaf{"SoftwareCount", flow.SoftwareCount}
    flow.EntityData.Leafs["pending-software-entries"] = types.YLeaf{"PendingSoftwareEntries", flow.PendingSoftwareEntries}
    return &(flow.EntityData)
}

// LptsPifib__Nodes_Node_Hardware
// Hardware specific
type LptsPifib__Nodes_Node_Hardware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage Table options.
    UsageEntries LptsPifib__Nodes_Node_Hardware_UsageEntries

    // Police details.
    Police LptsPifib__Nodes_Node_Hardware_Police

    // Static Police details.
    StaticPolice LptsPifib__Nodes_Node_Hardware_StaticPolice

    // Bfd details.
    Bfd LptsPifib__Nodes_Node_Hardware_Bfd

    // Hardware Entry type.
    Statistics LptsPifib__Nodes_Node_Hardware_Statistics

    // Hardware Entry options.
    IndexEntries LptsPifib__Nodes_Node_Hardware_IndexEntries
}

func (hardware *LptsPifib__Nodes_Node_Hardware) GetEntityData() *types.CommonEntityData {
    hardware.EntityData.YFilter = hardware.YFilter
    hardware.EntityData.YangName = "hardware"
    hardware.EntityData.BundleName = "cisco_ios_xr"
    hardware.EntityData.ParentYangName = "node"
    hardware.EntityData.SegmentPath = "Cisco-IOS-XR-platform-pifib-oper:hardware"
    hardware.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardware.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardware.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardware.EntityData.Children = make(map[string]types.YChild)
    hardware.EntityData.Children["usage-entries"] = types.YChild{"UsageEntries", &hardware.UsageEntries}
    hardware.EntityData.Children["police"] = types.YChild{"Police", &hardware.Police}
    hardware.EntityData.Children["static-police"] = types.YChild{"StaticPolice", &hardware.StaticPolice}
    hardware.EntityData.Children["bfd"] = types.YChild{"Bfd", &hardware.Bfd}
    hardware.EntityData.Children["statistics"] = types.YChild{"Statistics", &hardware.Statistics}
    hardware.EntityData.Children["index-entries"] = types.YChild{"IndexEntries", &hardware.IndexEntries}
    hardware.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardware.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_UsageEntries
// Usage Table options
type LptsPifib__Nodes_Node_Hardware_UsageEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage details. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry.
    UsageEntry []LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry
}

func (usageEntries *LptsPifib__Nodes_Node_Hardware_UsageEntries) GetEntityData() *types.CommonEntityData {
    usageEntries.EntityData.YFilter = usageEntries.YFilter
    usageEntries.EntityData.YangName = "usage-entries"
    usageEntries.EntityData.BundleName = "cisco_ios_xr"
    usageEntries.EntityData.ParentYangName = "hardware"
    usageEntries.EntityData.SegmentPath = "usage-entries"
    usageEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageEntries.EntityData.Children = make(map[string]types.YChild)
    usageEntries.EntityData.Children["usage-entry"] = types.YChild{"UsageEntry", nil}
    for i := range usageEntries.UsageEntry {
        usageEntries.EntityData.Children[types.GetSegmentPath(&usageEntries.UsageEntry[i])] = types.YChild{"UsageEntry", &usageEntries.UsageEntry[i]}
    }
    usageEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usageEntries.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry
// Usage details
type LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Region ID. The type is UsageAddressFamily.
    RegionId interface{}

    // Per TCAM type usage info. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo.
    UsageInfo []LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
}

func (usageEntry *LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry) GetEntityData() *types.CommonEntityData {
    usageEntry.EntityData.YFilter = usageEntry.YFilter
    usageEntry.EntityData.YangName = "usage-entry"
    usageEntry.EntityData.BundleName = "cisco_ios_xr"
    usageEntry.EntityData.ParentYangName = "usage-entries"
    usageEntry.EntityData.SegmentPath = "usage-entry" + "[region-id='" + fmt.Sprintf("%v", usageEntry.RegionId) + "']"
    usageEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageEntry.EntityData.Children = make(map[string]types.YChild)
    usageEntry.EntityData.Children["usage-info"] = types.YChild{"UsageInfo", nil}
    for i := range usageEntry.UsageInfo {
        usageEntry.EntityData.Children[types.GetSegmentPath(&usageEntry.UsageInfo[i])] = types.YChild{"UsageInfo", &usageEntry.UsageInfo[i]}
    }
    usageEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    usageEntry.EntityData.Leafs["region-id"] = types.YLeaf{"RegionId", usageEntry.RegionId}
    return &(usageEntry.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
// Per TCAM type usage info
type LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pipe ID. The type is interface{} with range: 0..255.
    PipeId interface{}

    // Region Type. The type is interface{} with range: 0..255.
    Region interface{}

    // Region ID. The type is interface{} with range: 0..255.
    RegionId interface{}

    // Maximum Number of Entries in the Region. The type is interface{} with
    // range: 0..4294967295.
    Size interface{}

    // Used Number of Entries in the Region. The type is interface{} with range:
    // 0..4294967295.
    Used interface{}
}

func (usageInfo *LptsPifib__Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetEntityData() *types.CommonEntityData {
    usageInfo.EntityData.YFilter = usageInfo.YFilter
    usageInfo.EntityData.YangName = "usage-info"
    usageInfo.EntityData.BundleName = "cisco_ios_xr"
    usageInfo.EntityData.ParentYangName = "usage-entry"
    usageInfo.EntityData.SegmentPath = "usage-info"
    usageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageInfo.EntityData.Children = make(map[string]types.YChild)
    usageInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    usageInfo.EntityData.Leafs["pipe-id"] = types.YLeaf{"PipeId", usageInfo.PipeId}
    usageInfo.EntityData.Leafs["region"] = types.YLeaf{"Region", usageInfo.Region}
    usageInfo.EntityData.Leafs["region-id"] = types.YLeaf{"RegionId", usageInfo.RegionId}
    usageInfo.EntityData.Leafs["size"] = types.YLeaf{"Size", usageInfo.Size}
    usageInfo.EntityData.Leafs["used"] = types.YLeaf{"Used", usageInfo.Used}
    return &(usageInfo.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_Police
// Police details
type LptsPifib__Nodes_Node_Hardware_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_Police_PoliceInfo.
    PoliceInfo []LptsPifib__Nodes_Node_Hardware_Police_PoliceInfo
}

func (police *LptsPifib__Nodes_Node_Hardware_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "hardware"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["police-info"] = types.YChild{"PoliceInfo", nil}
    for i := range police.PoliceInfo {
        police.EntityData.Children[types.GetSegmentPath(&police.PoliceInfo[i])] = types.YChild{"PoliceInfo", &police.PoliceInfo[i]}
    }
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(police.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_Police_PoliceInfo
// Per flow type police info
type LptsPifib__Nodes_Node_Hardware_Police_PoliceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // avgrate. The type is interface{} with range: 0..4294967295.
    Avgrate interface{}

    // burst. The type is interface{} with range: 0..4294967295.
    Burst interface{}

    // static avgrate. The type is interface{} with range: 0..4294967295.
    StaticAvgrate interface{}

    // avgrate type. The type is interface{} with range: 0..4294967295.
    AvgrateType interface{}

    // accepted stats. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedStats interface{}

    // dropped stats. The type is interface{} with range: 0..18446744073709551615.
    DroppedStats interface{}

    // policer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // iptos value. The type is interface{} with range: 0..255.
    IptosValue interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}

    // acl config. The type is interface{} with range: 0..255.
    AclConfig interface{}

    // acl str. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    AclStr interface{}
}

func (policeInfo *LptsPifib__Nodes_Node_Hardware_Police_PoliceInfo) GetEntityData() *types.CommonEntityData {
    policeInfo.EntityData.YFilter = policeInfo.YFilter
    policeInfo.EntityData.YangName = "police-info"
    policeInfo.EntityData.BundleName = "cisco_ios_xr"
    policeInfo.EntityData.ParentYangName = "police"
    policeInfo.EntityData.SegmentPath = "police-info"
    policeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeInfo.EntityData.Children = make(map[string]types.YChild)
    policeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    policeInfo.EntityData.Leafs["avgrate"] = types.YLeaf{"Avgrate", policeInfo.Avgrate}
    policeInfo.EntityData.Leafs["burst"] = types.YLeaf{"Burst", policeInfo.Burst}
    policeInfo.EntityData.Leafs["static-avgrate"] = types.YLeaf{"StaticAvgrate", policeInfo.StaticAvgrate}
    policeInfo.EntityData.Leafs["avgrate-type"] = types.YLeaf{"AvgrateType", policeInfo.AvgrateType}
    policeInfo.EntityData.Leafs["accepted-stats"] = types.YLeaf{"AcceptedStats", policeInfo.AcceptedStats}
    policeInfo.EntityData.Leafs["dropped-stats"] = types.YLeaf{"DroppedStats", policeInfo.DroppedStats}
    policeInfo.EntityData.Leafs["policer"] = types.YLeaf{"Policer", policeInfo.Policer}
    policeInfo.EntityData.Leafs["iptos-value"] = types.YLeaf{"IptosValue", policeInfo.IptosValue}
    policeInfo.EntityData.Leafs["change-type"] = types.YLeaf{"ChangeType", policeInfo.ChangeType}
    policeInfo.EntityData.Leafs["acl-config"] = types.YLeaf{"AclConfig", policeInfo.AclConfig}
    policeInfo.EntityData.Leafs["acl-str"] = types.YLeaf{"AclStr", policeInfo.AclStr}
    return &(policeInfo.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_StaticPolice
// Static Police details
type LptsPifib__Nodes_Node_Hardware_StaticPolice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per punt reason info. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_StaticPolice_StaticInfo.
    StaticInfo []LptsPifib__Nodes_Node_Hardware_StaticPolice_StaticInfo
}

func (staticPolice *LptsPifib__Nodes_Node_Hardware_StaticPolice) GetEntityData() *types.CommonEntityData {
    staticPolice.EntityData.YFilter = staticPolice.YFilter
    staticPolice.EntityData.YangName = "static-police"
    staticPolice.EntityData.BundleName = "cisco_ios_xr"
    staticPolice.EntityData.ParentYangName = "hardware"
    staticPolice.EntityData.SegmentPath = "static-police"
    staticPolice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticPolice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticPolice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticPolice.EntityData.Children = make(map[string]types.YChild)
    staticPolice.EntityData.Children["static-info"] = types.YChild{"StaticInfo", nil}
    for i := range staticPolice.StaticInfo {
        staticPolice.EntityData.Children[types.GetSegmentPath(&staticPolice.StaticInfo[i])] = types.YChild{"StaticInfo", &staticPolice.StaticInfo[i]}
    }
    staticPolice.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticPolice.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_StaticPolice_StaticInfo
// Per punt reason info
type LptsPifib__Nodes_Node_Hardware_StaticPolice_StaticInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // punt reason. The type is interface{} with range: 0..4294967295.
    PuntReason interface{}

    // sid. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // flow rate. The type is interface{} with range: 0..4294967295.
    FlowRate interface{}

    // burst rate. The type is interface{} with range: 0..4294967295.
    BurstRate interface{}

    // accepted. The type is interface{} with range: 0..18446744073709551615.
    Accepted interface{}

    // dropped. The type is interface{} with range: 0..18446744073709551615.
    Dropped interface{}

    // punt reason string. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    PuntReasonString interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}
}

func (staticInfo *LptsPifib__Nodes_Node_Hardware_StaticPolice_StaticInfo) GetEntityData() *types.CommonEntityData {
    staticInfo.EntityData.YFilter = staticInfo.YFilter
    staticInfo.EntityData.YangName = "static-info"
    staticInfo.EntityData.BundleName = "cisco_ios_xr"
    staticInfo.EntityData.ParentYangName = "static-police"
    staticInfo.EntityData.SegmentPath = "static-info"
    staticInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticInfo.EntityData.Children = make(map[string]types.YChild)
    staticInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    staticInfo.EntityData.Leafs["punt-reason"] = types.YLeaf{"PuntReason", staticInfo.PuntReason}
    staticInfo.EntityData.Leafs["sid"] = types.YLeaf{"Sid", staticInfo.Sid}
    staticInfo.EntityData.Leafs["flow-rate"] = types.YLeaf{"FlowRate", staticInfo.FlowRate}
    staticInfo.EntityData.Leafs["burst-rate"] = types.YLeaf{"BurstRate", staticInfo.BurstRate}
    staticInfo.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", staticInfo.Accepted}
    staticInfo.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", staticInfo.Dropped}
    staticInfo.EntityData.Leafs["punt-reason-string"] = types.YLeaf{"PuntReasonString", staticInfo.PuntReasonString}
    staticInfo.EntityData.Leafs["change-type"] = types.YLeaf{"ChangeType", staticInfo.ChangeType}
    return &(staticInfo.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_Bfd
// Bfd details
type LptsPifib__Nodes_Node_Hardware_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per bfd disc entry info. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_Bfd_BfdEntryInfo.
    BfdEntryInfo []LptsPifib__Nodes_Node_Hardware_Bfd_BfdEntryInfo
}

func (bfd *LptsPifib__Nodes_Node_Hardware_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "hardware"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Children["bfd-entry-info"] = types.YChild{"BfdEntryInfo", nil}
    for i := range bfd.BfdEntryInfo {
        bfd.EntityData.Children[types.GetSegmentPath(&bfd.BfdEntryInfo[i])] = types.YChild{"BfdEntryInfo", &bfd.BfdEntryInfo[i]}
    }
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfd.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_Bfd_BfdEntryInfo
// Per bfd disc entry info
type LptsPifib__Nodes_Node_Hardware_Bfd_BfdEntryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // index. The type is interface{} with range: 0..255.
    Index interface{}

    // is mcast. The type is interface{} with range: 0..255.
    IsMcast interface{}

    // fgid or vqi. The type is interface{} with range: 0..4294967295.
    FgidOrVqi interface{}

    // is valid. The type is interface{} with range: 0..255.
    IsValid interface{}

    // policer id. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}
}

func (bfdEntryInfo *LptsPifib__Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetEntityData() *types.CommonEntityData {
    bfdEntryInfo.EntityData.YFilter = bfdEntryInfo.YFilter
    bfdEntryInfo.EntityData.YangName = "bfd-entry-info"
    bfdEntryInfo.EntityData.BundleName = "cisco_ios_xr"
    bfdEntryInfo.EntityData.ParentYangName = "bfd"
    bfdEntryInfo.EntityData.SegmentPath = "bfd-entry-info"
    bfdEntryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdEntryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdEntryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdEntryInfo.EntityData.Children = make(map[string]types.YChild)
    bfdEntryInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdEntryInfo.EntityData.Leafs["index"] = types.YLeaf{"Index", bfdEntryInfo.Index}
    bfdEntryInfo.EntityData.Leafs["is-mcast"] = types.YLeaf{"IsMcast", bfdEntryInfo.IsMcast}
    bfdEntryInfo.EntityData.Leafs["fgid-or-vqi"] = types.YLeaf{"FgidOrVqi", bfdEntryInfo.FgidOrVqi}
    bfdEntryInfo.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", bfdEntryInfo.IsValid}
    bfdEntryInfo.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", bfdEntryInfo.PolicerId}
    return &(bfdEntryInfo.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_Statistics
// Hardware Entry type
type LptsPifib__Nodes_Node_Hardware_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Deleted-entry accepted packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepted interface{}

    // Deleted-entry dropped packets counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Dropped interface{}

    // Statistics clear timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    ClearTs interface{}

    // No statistics memory error. The type is interface{} with range:
    // 0..18446744073709551615.
    NoStatsMemErr interface{}
}

func (statistics *LptsPifib__Nodes_Node_Hardware_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "hardware"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", statistics.Accepted}
    statistics.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", statistics.Dropped}
    statistics.EntityData.Leafs["clear-ts"] = types.YLeaf{"ClearTs", statistics.ClearTs}
    statistics.EntityData.Leafs["no-stats-mem-err"] = types.YLeaf{"NoStatsMemErr", statistics.NoStatsMemErr}
    return &(statistics.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_IndexEntries
// Hardware Entry options
type LptsPifib__Nodes_Node_Hardware_IndexEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry options. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry.
    IndexEntry []LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry
}

func (indexEntries *LptsPifib__Nodes_Node_Hardware_IndexEntries) GetEntityData() *types.CommonEntityData {
    indexEntries.EntityData.YFilter = indexEntries.YFilter
    indexEntries.EntityData.YangName = "index-entries"
    indexEntries.EntityData.BundleName = "cisco_ios_xr"
    indexEntries.EntityData.ParentYangName = "hardware"
    indexEntries.EntityData.SegmentPath = "index-entries"
    indexEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexEntries.EntityData.Children = make(map[string]types.YChild)
    indexEntries.EntityData.Children["index-entry"] = types.YChild{"IndexEntry", nil}
    for i := range indexEntries.IndexEntry {
        indexEntries.EntityData.Children[types.GetSegmentPath(&indexEntries.IndexEntry[i])] = types.YChild{"IndexEntry", &indexEntries.IndexEntry[i]}
    }
    indexEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(indexEntries.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry
// Entry options
type LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index. The type is interface{} with range:
    // -2147483648..2147483647.
    Index interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3Protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    IntfHandle interface{}

    // Interface Name. The type is string.
    IntfName interface{}

    // Interface uidb index. The type is interface{} with range: 0..4294967295.
    UidbIndex interface{}

    // Local IP Address. The type is string.
    LocalAddr interface{}

    // Local Prefix Length. The type is interface{} with range: 0..4294967295.
    LocalPrefixLen interface{}

    // Remote IP Address. The type is string.
    RemoteAddr interface{}

    // Remote Prefix Length. The type is interface{} with range: 0..4294967295.
    RemotePrefixLen interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Remote Port/ICMP Type/IGMP Type. The type is interface{} with range:
    // 0..4294967295.
    UValue interface{}

    // Union Key Length. The type is interface{} with range: 0..4294967295.
    ULen interface{}

    // Local port. The type is interface{} with range: 0..4294967295.
    LocalPort interface{}

    // Is Fragment. The type is interface{} with range: 0..255.
    IsFrag interface{}

    // Is SYN. The type is interface{} with range: 0..255.
    IsSyn interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Action String. The type is string.
    ActionString interface{}

    // Listener Tag. The type is interface{} with range: 0..255.
    ListenerTag interface{}

    // Is FGID or not. The type is interface{} with range: 0..255.
    IsFgid interface{}

    // Is VRF or not. The type is interface{} with range: 0..255.
    IsVrf interface{}

    // Is optimized or not. The type is interface{} with range: 0..255.
    IsOptimized interface{}

    // Is uidb set for optimized entry or not. The type is interface{} with range:
    // 0..255.
    IsUidbOptBit interface{}

    // fabric group id or swith fabric port. The type is interface{} with range:
    // 0..4294967295.
    FgidOrSfp interface{}

    // Is entry remote or not. The type is interface{} with range: 0..255.
    RemoteRack interface{}

    // Remote racknum if remote. The type is interface{} with range:
    // 0..4294967295.
    RackId interface{}

    // Remote slotnum if remote. The type is interface{} with range:
    // 0..4294967295.
    Rslot interface{}

    // Committed Information Rate. The type is interface{} with range:
    // 0..18446744073709551615.
    Cir interface{}

    // Flow type. The type is interface{} with range: 0..4294967295.
    FlowType interface{}

    // Flow priority or COS. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Stream ID. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // Policer avg. rate limit. The type is interface{} with range: 0..4294967295.
    PolicerAvgrate interface{}

    // Policer burst. The type is interface{} with range: 0..4294967295.
    PolicerBurst interface{}

    // Lookup priority. The type is interface{} with range:
    // -2147483648..2147483647.
    LookupPriority interface{}

    // Storage priority. The type is interface{} with range:
    // -2147483648..2147483647.
    StoragePriority interface{}

    // Number of TCAM entries used. The type is interface{} with range:
    // -2147483648..2147483647.
    NumTmEntries interface{}

    // ptr to ifib_entry_st. The type is interface{} with range: 0..4294967295.
    EntryPtr interface{}

    // ptr to ifib_entry_shadow_st. The type is interface{} with range:
    // 0..4294967295.
    EntryShadowPtr interface{}

    // ptr to dlqueue list node. The type is interface{} with range:
    // 0..4294967295.
    ListNodePtr interface{}

    // state of pifib entry. The type is interface{} with range: 0..255.
    State interface{}

    // failure cause. The type is interface{} with range: 0..255.
    RetryFailCause interface{}

    // retries count. The type is interface{} with range: 0..255.
    NumRetries interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    MinTtl interface{}

    // Union Key Type. The type is interface{} with range: 0..255.
    UType interface{}

    // Remote FGID. The type is interface{} with range: 0..4294967295.
    RemoteFgid interface{}

    // Acl name. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    AclStr interface{}

    // Stats not available. The type is interface{} with range: 0..255.
    NoStats interface{}

    // Per pipe type hardware info. The type is slice of
    // LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo.
    HwInfo []LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
}

func (indexEntry *LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry) GetEntityData() *types.CommonEntityData {
    indexEntry.EntityData.YFilter = indexEntry.YFilter
    indexEntry.EntityData.YangName = "index-entry"
    indexEntry.EntityData.BundleName = "cisco_ios_xr"
    indexEntry.EntityData.ParentYangName = "index-entries"
    indexEntry.EntityData.SegmentPath = "index-entry" + "[index='" + fmt.Sprintf("%v", indexEntry.Index) + "']"
    indexEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexEntry.EntityData.Children = make(map[string]types.YChild)
    indexEntry.EntityData.Children["hw-info"] = types.YChild{"HwInfo", nil}
    for i := range indexEntry.HwInfo {
        indexEntry.EntityData.Children[types.GetSegmentPath(&indexEntry.HwInfo[i])] = types.YChild{"HwInfo", &indexEntry.HwInfo[i]}
    }
    indexEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    indexEntry.EntityData.Leafs["index"] = types.YLeaf{"Index", indexEntry.Index}
    indexEntry.EntityData.Leafs["l3protocol"] = types.YLeaf{"L3Protocol", indexEntry.L3Protocol}
    indexEntry.EntityData.Leafs["l4protocol"] = types.YLeaf{"L4Protocol", indexEntry.L4Protocol}
    indexEntry.EntityData.Leafs["intf-handle"] = types.YLeaf{"IntfHandle", indexEntry.IntfHandle}
    indexEntry.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", indexEntry.IntfName}
    indexEntry.EntityData.Leafs["uidb-index"] = types.YLeaf{"UidbIndex", indexEntry.UidbIndex}
    indexEntry.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", indexEntry.LocalAddr}
    indexEntry.EntityData.Leafs["local-prefix-len"] = types.YLeaf{"LocalPrefixLen", indexEntry.LocalPrefixLen}
    indexEntry.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", indexEntry.RemoteAddr}
    indexEntry.EntityData.Leafs["remote-prefix-len"] = types.YLeaf{"RemotePrefixLen", indexEntry.RemotePrefixLen}
    indexEntry.EntityData.Leafs["vrf-id"] = types.YLeaf{"VrfId", indexEntry.VrfId}
    indexEntry.EntityData.Leafs["u-value"] = types.YLeaf{"UValue", indexEntry.UValue}
    indexEntry.EntityData.Leafs["u-len"] = types.YLeaf{"ULen", indexEntry.ULen}
    indexEntry.EntityData.Leafs["local-port"] = types.YLeaf{"LocalPort", indexEntry.LocalPort}
    indexEntry.EntityData.Leafs["is-frag"] = types.YLeaf{"IsFrag", indexEntry.IsFrag}
    indexEntry.EntityData.Leafs["is-syn"] = types.YLeaf{"IsSyn", indexEntry.IsSyn}
    indexEntry.EntityData.Leafs["action"] = types.YLeaf{"Action", indexEntry.Action}
    indexEntry.EntityData.Leafs["action-string"] = types.YLeaf{"ActionString", indexEntry.ActionString}
    indexEntry.EntityData.Leafs["listener-tag"] = types.YLeaf{"ListenerTag", indexEntry.ListenerTag}
    indexEntry.EntityData.Leafs["is-fgid"] = types.YLeaf{"IsFgid", indexEntry.IsFgid}
    indexEntry.EntityData.Leafs["is-vrf"] = types.YLeaf{"IsVrf", indexEntry.IsVrf}
    indexEntry.EntityData.Leafs["is-optimized"] = types.YLeaf{"IsOptimized", indexEntry.IsOptimized}
    indexEntry.EntityData.Leafs["is-uidb-opt-bit"] = types.YLeaf{"IsUidbOptBit", indexEntry.IsUidbOptBit}
    indexEntry.EntityData.Leafs["fgid-or-sfp"] = types.YLeaf{"FgidOrSfp", indexEntry.FgidOrSfp}
    indexEntry.EntityData.Leafs["remote-rack"] = types.YLeaf{"RemoteRack", indexEntry.RemoteRack}
    indexEntry.EntityData.Leafs["rack-id"] = types.YLeaf{"RackId", indexEntry.RackId}
    indexEntry.EntityData.Leafs["rslot"] = types.YLeaf{"Rslot", indexEntry.Rslot}
    indexEntry.EntityData.Leafs["cir"] = types.YLeaf{"Cir", indexEntry.Cir}
    indexEntry.EntityData.Leafs["flow-type"] = types.YLeaf{"FlowType", indexEntry.FlowType}
    indexEntry.EntityData.Leafs["priority"] = types.YLeaf{"Priority", indexEntry.Priority}
    indexEntry.EntityData.Leafs["sid"] = types.YLeaf{"Sid", indexEntry.Sid}
    indexEntry.EntityData.Leafs["policer-avgrate"] = types.YLeaf{"PolicerAvgrate", indexEntry.PolicerAvgrate}
    indexEntry.EntityData.Leafs["policer-burst"] = types.YLeaf{"PolicerBurst", indexEntry.PolicerBurst}
    indexEntry.EntityData.Leafs["lookup-priority"] = types.YLeaf{"LookupPriority", indexEntry.LookupPriority}
    indexEntry.EntityData.Leafs["storage-priority"] = types.YLeaf{"StoragePriority", indexEntry.StoragePriority}
    indexEntry.EntityData.Leafs["num-tm-entries"] = types.YLeaf{"NumTmEntries", indexEntry.NumTmEntries}
    indexEntry.EntityData.Leafs["entry-ptr"] = types.YLeaf{"EntryPtr", indexEntry.EntryPtr}
    indexEntry.EntityData.Leafs["entry-shadow-ptr"] = types.YLeaf{"EntryShadowPtr", indexEntry.EntryShadowPtr}
    indexEntry.EntityData.Leafs["list-node-ptr"] = types.YLeaf{"ListNodePtr", indexEntry.ListNodePtr}
    indexEntry.EntityData.Leafs["state"] = types.YLeaf{"State", indexEntry.State}
    indexEntry.EntityData.Leafs["retry-fail-cause"] = types.YLeaf{"RetryFailCause", indexEntry.RetryFailCause}
    indexEntry.EntityData.Leafs["num-retries"] = types.YLeaf{"NumRetries", indexEntry.NumRetries}
    indexEntry.EntityData.Leafs["min-ttl"] = types.YLeaf{"MinTtl", indexEntry.MinTtl}
    indexEntry.EntityData.Leafs["u-type"] = types.YLeaf{"UType", indexEntry.UType}
    indexEntry.EntityData.Leafs["remote-fgid"] = types.YLeaf{"RemoteFgid", indexEntry.RemoteFgid}
    indexEntry.EntityData.Leafs["acl-str"] = types.YLeaf{"AclStr", indexEntry.AclStr}
    indexEntry.EntityData.Leafs["no-stats"] = types.YLeaf{"NoStats", indexEntry.NoStats}
    return &(indexEntry.EntityData)
}

// LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
// Per pipe type hardware info
type LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer Pointer. The type is interface{} with range: 0..4294967295.
    Policer interface{}

    // Stats Pointer. The type is interface{} with range: 0..4294967295.
    StatsPtr interface{}

    // Accepted Packets Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Accepted interface{}

    // Dropped Packets Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    Dropped interface{}

    // Relative position in sorting order. The type is interface{} with range:
    // -2147483648..2147483647.
    SortStartOffset interface{}

    // Relative position in TCAM. The type is interface{} with range:
    // -2147483648..2147483647.
    TmStartOffset interface{}
}

func (hwInfo *LptsPifib__Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetEntityData() *types.CommonEntityData {
    hwInfo.EntityData.YFilter = hwInfo.YFilter
    hwInfo.EntityData.YangName = "hw-info"
    hwInfo.EntityData.BundleName = "cisco_ios_xr"
    hwInfo.EntityData.ParentYangName = "index-entry"
    hwInfo.EntityData.SegmentPath = "hw-info"
    hwInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwInfo.EntityData.Children = make(map[string]types.YChild)
    hwInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    hwInfo.EntityData.Leafs["policer"] = types.YLeaf{"Policer", hwInfo.Policer}
    hwInfo.EntityData.Leafs["stats-ptr"] = types.YLeaf{"StatsPtr", hwInfo.StatsPtr}
    hwInfo.EntityData.Leafs["accepted"] = types.YLeaf{"Accepted", hwInfo.Accepted}
    hwInfo.EntityData.Leafs["dropped"] = types.YLeaf{"Dropped", hwInfo.Dropped}
    hwInfo.EntityData.Leafs["sort-start-offset"] = types.YLeaf{"SortStartOffset", hwInfo.SortStartOffset}
    hwInfo.EntityData.Leafs["tm-start-offset"] = types.YLeaf{"TmStartOffset", hwInfo.TmStartOffset}
    return &(hwInfo.EntityData)
}

