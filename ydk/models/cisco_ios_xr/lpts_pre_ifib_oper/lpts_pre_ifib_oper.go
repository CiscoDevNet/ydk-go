// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-pre-ifib package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lpts-pifib: lpts pre-ifib data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-pre-ifib-oper lpts-pifib}", reflect.TypeOf(LptsPifib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib", reflect.TypeOf(LptsPifib{}))
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

// LptsPifib
// lpts pre-ifib data
type LptsPifib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Pre-ifib Nodes.
    Nodes LptsPifib_Nodes
}

func (lptsPifib *LptsPifib) GetEntityData() *types.CommonEntityData {
    lptsPifib.EntityData.YFilter = lptsPifib.YFilter
    lptsPifib.EntityData.YangName = "lpts-pifib"
    lptsPifib.EntityData.BundleName = "cisco_ios_xr"
    lptsPifib.EntityData.ParentYangName = "Cisco-IOS-XR-lpts-pre-ifib-oper"
    lptsPifib.EntityData.SegmentPath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib"
    lptsPifib.EntityData.AbsolutePath = lptsPifib.EntityData.SegmentPath
    lptsPifib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsPifib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsPifib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsPifib.EntityData.Children = types.NewOrderedMap()
    lptsPifib.EntityData.Children.Append("nodes", types.YChild{"Nodes", &lptsPifib.Nodes})
    lptsPifib.EntityData.Leafs = types.NewOrderedMap()

    lptsPifib.EntityData.YListKeys = []string {}

    return &(lptsPifib.EntityData)
}

// LptsPifib_Nodes
// List of Pre-ifib Nodes
type LptsPifib_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pre-ifib data for particular node. The type is slice of
    // LptsPifib_Nodes_Node.
    Node []*LptsPifib_Nodes_Node
}

func (nodes *LptsPifib_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "lpts-pifib"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/" + nodes.EntityData.SegmentPath
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

// LptsPifib_Nodes_Node
// Pre-ifib data for particular node
type LptsPifib_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Type specific.
    TypeValues LptsPifib_Nodes_Node_TypeValues

    // Dynamic Flows Statistics.
    DynamicFlowsStats LptsPifib_Nodes_Node_DynamicFlowsStats

    // Hardware specific.
    Hardware LptsPifib_Nodes_Node_Hardware
}

func (node *LptsPifib_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("type-values", types.YChild{"TypeValues", &node.TypeValues})
    node.EntityData.Children.Append("dynamic-flows-stats", types.YChild{"DynamicFlowsStats", &node.DynamicFlowsStats})
    node.EntityData.Children.Append("Cisco-IOS-XR-platform-pifib-oper:hardware", types.YChild{"Hardware", &node.Hardware})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// LptsPifib_Nodes_Node_TypeValues
// Type specific
type LptsPifib_Nodes_Node_TypeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // pifib types. The type is slice of
    // LptsPifib_Nodes_Node_TypeValues_TypeValue.
    TypeValue []*LptsPifib_Nodes_Node_TypeValues_TypeValue
}

func (typeValues *LptsPifib_Nodes_Node_TypeValues) GetEntityData() *types.CommonEntityData {
    typeValues.EntityData.YFilter = typeValues.YFilter
    typeValues.EntityData.YangName = "type-values"
    typeValues.EntityData.BundleName = "cisco_ios_xr"
    typeValues.EntityData.ParentYangName = "node"
    typeValues.EntityData.SegmentPath = "type-values"
    typeValues.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/" + typeValues.EntityData.SegmentPath
    typeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeValues.EntityData.Children = types.NewOrderedMap()
    typeValues.EntityData.Children.Append("type-value", types.YChild{"TypeValue", nil})
    for i := range typeValues.TypeValue {
        typeValues.EntityData.Children.Append(types.GetSegmentPath(typeValues.TypeValue[i]), types.YChild{"TypeValue", typeValues.TypeValue[i]})
    }
    typeValues.EntityData.Leafs = types.NewOrderedMap()

    typeValues.EntityData.YListKeys = []string {}

    return &(typeValues.EntityData)
}

// LptsPifib_Nodes_Node_TypeValues_TypeValue
// pifib types
type LptsPifib_Nodes_Node_TypeValues_TypeValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type value. The type is LptsPifib.
    PifibType interface{}

    // Data for single pre-ifib entry. The type is slice of
    // LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry.
    Entry []*LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry
}

func (typeValue *LptsPifib_Nodes_Node_TypeValues_TypeValue) GetEntityData() *types.CommonEntityData {
    typeValue.EntityData.YFilter = typeValue.YFilter
    typeValue.EntityData.YangName = "type-value"
    typeValue.EntityData.BundleName = "cisco_ios_xr"
    typeValue.EntityData.ParentYangName = "type-values"
    typeValue.EntityData.SegmentPath = "type-value" + types.AddKeyToken(typeValue.PifibType, "pifib-type")
    typeValue.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/type-values/" + typeValue.EntityData.SegmentPath
    typeValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeValue.EntityData.Children = types.NewOrderedMap()
    typeValue.EntityData.Children.Append("entry", types.YChild{"Entry", nil})
    for i := range typeValue.Entry {
        typeValue.EntityData.Children.Append(types.GetSegmentPath(typeValue.Entry[i]), types.YChild{"Entry", typeValue.Entry[i]})
    }
    typeValue.EntityData.Leafs = types.NewOrderedMap()
    typeValue.EntityData.Leafs.Append("pifib-type", types.YLeaf{"PifibType", typeValue.PifibType})

    typeValue.EntityData.YListKeys = []string {"PifibType"}

    return &(typeValue.EntityData)
}

// LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry
// Data for single pre-ifib entry
type LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Single Pre-ifib entry. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Entry interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    Vid interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4protocol interface{}

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

func (entry *LptsPifib_Nodes_Node_TypeValues_TypeValue_Entry) GetEntityData() *types.CommonEntityData {
    entry.EntityData.YFilter = entry.YFilter
    entry.EntityData.YangName = "entry"
    entry.EntityData.BundleName = "cisco_ios_xr"
    entry.EntityData.ParentYangName = "type-value"
    entry.EntityData.SegmentPath = "entry" + types.AddKeyToken(entry.Entry, "entry")
    entry.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/type-values/type-value/" + entry.EntityData.SegmentPath
    entry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entry.EntityData.Children = types.NewOrderedMap()
    entry.EntityData.Leafs = types.NewOrderedMap()
    entry.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", entry.Entry})
    entry.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", entry.VrfName})
    entry.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", entry.Vid})
    entry.EntityData.Leafs.Append("l3protocol", types.YLeaf{"L3protocol", entry.L3protocol})
    entry.EntityData.Leafs.Append("l4protocol", types.YLeaf{"L4protocol", entry.L4protocol})
    entry.EntityData.Leafs.Append("intf-name", types.YLeaf{"IntfName", entry.IntfName})
    entry.EntityData.Leafs.Append("intf-handle", types.YLeaf{"IntfHandle", entry.IntfHandle})
    entry.EntityData.Leafs.Append("destination-addr", types.YLeaf{"DestinationAddr", entry.DestinationAddr})
    entry.EntityData.Leafs.Append("source-addr", types.YLeaf{"SourceAddr", entry.SourceAddr})
    entry.EntityData.Leafs.Append("destination-type", types.YLeaf{"DestinationType", entry.DestinationType})
    entry.EntityData.Leafs.Append("destination-value", types.YLeaf{"DestinationValue", entry.DestinationValue})
    entry.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", entry.SourcePort})
    entry.EntityData.Leafs.Append("is-frag", types.YLeaf{"IsFrag", entry.IsFrag})
    entry.EntityData.Leafs.Append("is-syn", types.YLeaf{"IsSyn", entry.IsSyn})
    entry.EntityData.Leafs.Append("opcode", types.YLeaf{"Opcode", entry.Opcode})
    entry.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", entry.FlowType})
    entry.EntityData.Leafs.Append("listener-tag", types.YLeaf{"ListenerTag", entry.ListenerTag})
    entry.EntityData.Leafs.Append("local-flag", types.YLeaf{"LocalFlag", entry.LocalFlag})
    entry.EntityData.Leafs.Append("is-fgid", types.YLeaf{"IsFgid", entry.IsFgid})
    entry.EntityData.Leafs.Append("deliver-list-short", types.YLeaf{"DeliverListShort", entry.DeliverListShort})
    entry.EntityData.Leafs.Append("deliver-list-long", types.YLeaf{"DeliverListLong", entry.DeliverListLong})
    entry.EntityData.Leafs.Append("min-ttl", types.YLeaf{"MinTtl", entry.MinTtl})
    entry.EntityData.Leafs.Append("accepts", types.YLeaf{"Accepts", entry.Accepts})
    entry.EntityData.Leafs.Append("drops", types.YLeaf{"Drops", entry.Drops})
    entry.EntityData.Leafs.Append("stale", types.YLeaf{"Stale", entry.Stale})
    entry.EntityData.Leafs.Append("pifib-type", types.YLeaf{"PifibType", entry.PifibType})
    entry.EntityData.Leafs.Append("pifib-program-time", types.YLeaf{"PifibProgramTime", entry.PifibProgramTime})

    entry.EntityData.YListKeys = []string {"Entry"}

    return &(entry.EntityData)
}

// LptsPifib_Nodes_Node_DynamicFlowsStats
// Dynamic Flows Statistics
type LptsPifib_Nodes_Node_DynamicFlowsStats struct {
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
    // LptsPifib_Nodes_Node_DynamicFlowsStats_Flow.
    Flow []*LptsPifib_Nodes_Node_DynamicFlowsStats_Flow
}

func (dynamicFlowsStats *LptsPifib_Nodes_Node_DynamicFlowsStats) GetEntityData() *types.CommonEntityData {
    dynamicFlowsStats.EntityData.YFilter = dynamicFlowsStats.YFilter
    dynamicFlowsStats.EntityData.YangName = "dynamic-flows-stats"
    dynamicFlowsStats.EntityData.BundleName = "cisco_ios_xr"
    dynamicFlowsStats.EntityData.ParentYangName = "node"
    dynamicFlowsStats.EntityData.SegmentPath = "dynamic-flows-stats"
    dynamicFlowsStats.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/" + dynamicFlowsStats.EntityData.SegmentPath
    dynamicFlowsStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicFlowsStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicFlowsStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicFlowsStats.EntityData.Children = types.NewOrderedMap()
    dynamicFlowsStats.EntityData.Children.Append("flow", types.YChild{"Flow", nil})
    for i := range dynamicFlowsStats.Flow {
        types.SetYListKey(dynamicFlowsStats.Flow[i], i)
        dynamicFlowsStats.EntityData.Children.Append(types.GetSegmentPath(dynamicFlowsStats.Flow[i]), types.YChild{"Flow", dynamicFlowsStats.Flow[i]})
    }
    dynamicFlowsStats.EntityData.Leafs = types.NewOrderedMap()
    dynamicFlowsStats.EntityData.Leafs.Append("dynamic-flows-enabled", types.YLeaf{"DynamicFlowsEnabled", dynamicFlowsStats.DynamicFlowsEnabled})
    dynamicFlowsStats.EntityData.Leafs.Append("platform-supported-max", types.YLeaf{"PlatformSupportedMax", dynamicFlowsStats.PlatformSupportedMax})
    dynamicFlowsStats.EntityData.Leafs.Append("platform-configured-max", types.YLeaf{"PlatformConfiguredMax", dynamicFlowsStats.PlatformConfiguredMax})
    dynamicFlowsStats.EntityData.Leafs.Append("platform-total-configured", types.YLeaf{"PlatformTotalConfigured", dynamicFlowsStats.PlatformTotalConfigured})
    dynamicFlowsStats.EntityData.Leafs.Append("total-hw-entries", types.YLeaf{"TotalHwEntries", dynamicFlowsStats.TotalHwEntries})
    dynamicFlowsStats.EntityData.Leafs.Append("total-sw-entries", types.YLeaf{"TotalSwEntries", dynamicFlowsStats.TotalSwEntries})

    dynamicFlowsStats.EntityData.YListKeys = []string {}

    return &(dynamicFlowsStats.EntityData)
}

// LptsPifib_Nodes_Node_DynamicFlowsStats_Flow
// Flow Datalist
type LptsPifib_Nodes_Node_DynamicFlowsStats_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (flow *LptsPifib_Nodes_Node_DynamicFlowsStats_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "dynamic-flows-stats"
    flow.EntityData.SegmentPath = "flow" + types.AddNoKeyToken(flow)
    flow.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/dynamic-flows-stats/" + flow.EntityData.SegmentPath
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-name", types.YLeaf{"FlowName", flow.FlowName})
    flow.EntityData.Leafs.Append("configurable", types.YLeaf{"Configurable", flow.Configurable})
    flow.EntityData.Leafs.Append("configured", types.YLeaf{"Configured", flow.Configured})
    flow.EntityData.Leafs.Append("default-max", types.YLeaf{"DefaultMax", flow.DefaultMax})
    flow.EntityData.Leafs.Append("configured-max", types.YLeaf{"ConfiguredMax", flow.ConfiguredMax})
    flow.EntityData.Leafs.Append("active-max", types.YLeaf{"ActiveMax", flow.ActiveMax})
    flow.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", flow.HardwareCount})
    flow.EntityData.Leafs.Append("software-count", types.YLeaf{"SoftwareCount", flow.SoftwareCount})
    flow.EntityData.Leafs.Append("pending-software-entries", types.YLeaf{"PendingSoftwareEntries", flow.PendingSoftwareEntries})

    flow.EntityData.YListKeys = []string {}

    return &(flow.EntityData)
}

// LptsPifib_Nodes_Node_Hardware
// Hardware specific
type LptsPifib_Nodes_Node_Hardware struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage Table options.
    UsageEntries LptsPifib_Nodes_Node_Hardware_UsageEntries

    // Police details.
    Police LptsPifib_Nodes_Node_Hardware_Police

    // Static Police details.
    StaticPolice LptsPifib_Nodes_Node_Hardware_StaticPolice

    // Bfd details.
    Bfd LptsPifib_Nodes_Node_Hardware_Bfd

    // Hardware Entry type.
    Statistics LptsPifib_Nodes_Node_Hardware_Statistics

    // Hardware Entry options.
    IndexEntries LptsPifib_Nodes_Node_Hardware_IndexEntries
}

func (hardware *LptsPifib_Nodes_Node_Hardware) GetEntityData() *types.CommonEntityData {
    hardware.EntityData.YFilter = hardware.YFilter
    hardware.EntityData.YangName = "hardware"
    hardware.EntityData.BundleName = "cisco_ios_xr"
    hardware.EntityData.ParentYangName = "node"
    hardware.EntityData.SegmentPath = "Cisco-IOS-XR-platform-pifib-oper:hardware"
    hardware.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/" + hardware.EntityData.SegmentPath
    hardware.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardware.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardware.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardware.EntityData.Children = types.NewOrderedMap()
    hardware.EntityData.Children.Append("usage-entries", types.YChild{"UsageEntries", &hardware.UsageEntries})
    hardware.EntityData.Children.Append("police", types.YChild{"Police", &hardware.Police})
    hardware.EntityData.Children.Append("static-police", types.YChild{"StaticPolice", &hardware.StaticPolice})
    hardware.EntityData.Children.Append("bfd", types.YChild{"Bfd", &hardware.Bfd})
    hardware.EntityData.Children.Append("statistics", types.YChild{"Statistics", &hardware.Statistics})
    hardware.EntityData.Children.Append("index-entries", types.YChild{"IndexEntries", &hardware.IndexEntries})
    hardware.EntityData.Leafs = types.NewOrderedMap()

    hardware.EntityData.YListKeys = []string {}

    return &(hardware.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_UsageEntries
// Usage Table options
type LptsPifib_Nodes_Node_Hardware_UsageEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage details. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry.
    UsageEntry []*LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry
}

func (usageEntries *LptsPifib_Nodes_Node_Hardware_UsageEntries) GetEntityData() *types.CommonEntityData {
    usageEntries.EntityData.YFilter = usageEntries.YFilter
    usageEntries.EntityData.YangName = "usage-entries"
    usageEntries.EntityData.BundleName = "cisco_ios_xr"
    usageEntries.EntityData.ParentYangName = "hardware"
    usageEntries.EntityData.SegmentPath = "usage-entries"
    usageEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + usageEntries.EntityData.SegmentPath
    usageEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageEntries.EntityData.Children = types.NewOrderedMap()
    usageEntries.EntityData.Children.Append("usage-entry", types.YChild{"UsageEntry", nil})
    for i := range usageEntries.UsageEntry {
        usageEntries.EntityData.Children.Append(types.GetSegmentPath(usageEntries.UsageEntry[i]), types.YChild{"UsageEntry", usageEntries.UsageEntry[i]})
    }
    usageEntries.EntityData.Leafs = types.NewOrderedMap()

    usageEntries.EntityData.YListKeys = []string {}

    return &(usageEntries.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry
// Usage details
type LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Region ID. The type is UsageAddressFamily.
    RegionId interface{}

    // Per TCAM type usage info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo.
    UsageInfo []*LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
}

func (usageEntry *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry) GetEntityData() *types.CommonEntityData {
    usageEntry.EntityData.YFilter = usageEntry.YFilter
    usageEntry.EntityData.YangName = "usage-entry"
    usageEntry.EntityData.BundleName = "cisco_ios_xr"
    usageEntry.EntityData.ParentYangName = "usage-entries"
    usageEntry.EntityData.SegmentPath = "usage-entry" + types.AddKeyToken(usageEntry.RegionId, "region-id")
    usageEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/usage-entries/" + usageEntry.EntityData.SegmentPath
    usageEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageEntry.EntityData.Children = types.NewOrderedMap()
    usageEntry.EntityData.Children.Append("usage-info", types.YChild{"UsageInfo", nil})
    for i := range usageEntry.UsageInfo {
        types.SetYListKey(usageEntry.UsageInfo[i], i)
        usageEntry.EntityData.Children.Append(types.GetSegmentPath(usageEntry.UsageInfo[i]), types.YChild{"UsageInfo", usageEntry.UsageInfo[i]})
    }
    usageEntry.EntityData.Leafs = types.NewOrderedMap()
    usageEntry.EntityData.Leafs.Append("region-id", types.YLeaf{"RegionId", usageEntry.RegionId})

    usageEntry.EntityData.YListKeys = []string {"RegionId"}

    return &(usageEntry.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo
// Per TCAM type usage info
type LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (usageInfo *LptsPifib_Nodes_Node_Hardware_UsageEntries_UsageEntry_UsageInfo) GetEntityData() *types.CommonEntityData {
    usageInfo.EntityData.YFilter = usageInfo.YFilter
    usageInfo.EntityData.YangName = "usage-info"
    usageInfo.EntityData.BundleName = "cisco_ios_xr"
    usageInfo.EntityData.ParentYangName = "usage-entry"
    usageInfo.EntityData.SegmentPath = "usage-info" + types.AddNoKeyToken(usageInfo)
    usageInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/usage-entries/usage-entry/" + usageInfo.EntityData.SegmentPath
    usageInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usageInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usageInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usageInfo.EntityData.Children = types.NewOrderedMap()
    usageInfo.EntityData.Leafs = types.NewOrderedMap()
    usageInfo.EntityData.Leafs.Append("pipe-id", types.YLeaf{"PipeId", usageInfo.PipeId})
    usageInfo.EntityData.Leafs.Append("region", types.YLeaf{"Region", usageInfo.Region})
    usageInfo.EntityData.Leafs.Append("region-id", types.YLeaf{"RegionId", usageInfo.RegionId})
    usageInfo.EntityData.Leafs.Append("size", types.YLeaf{"Size", usageInfo.Size})
    usageInfo.EntityData.Leafs.Append("used", types.YLeaf{"Used", usageInfo.Used})

    usageInfo.EntityData.YListKeys = []string {}

    return &(usageInfo.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_Police
// Police details
type LptsPifib_Nodes_Node_Hardware_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per flow type police info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo.
    PoliceInfo []*LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo
}

func (police *LptsPifib_Nodes_Node_Hardware_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "hardware"
    police.EntityData.SegmentPath = "police"
    police.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + police.EntityData.SegmentPath
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = types.NewOrderedMap()
    police.EntityData.Children.Append("police-info", types.YChild{"PoliceInfo", nil})
    for i := range police.PoliceInfo {
        types.SetYListKey(police.PoliceInfo[i], i)
        police.EntityData.Children.Append(types.GetSegmentPath(police.PoliceInfo[i]), types.YChild{"PoliceInfo", police.PoliceInfo[i]})
    }
    police.EntityData.Leafs = types.NewOrderedMap()

    police.EntityData.YListKeys = []string {}

    return &(police.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo
// Per flow type police info
type LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AclStr interface{}
}

func (policeInfo *LptsPifib_Nodes_Node_Hardware_Police_PoliceInfo) GetEntityData() *types.CommonEntityData {
    policeInfo.EntityData.YFilter = policeInfo.YFilter
    policeInfo.EntityData.YangName = "police-info"
    policeInfo.EntityData.BundleName = "cisco_ios_xr"
    policeInfo.EntityData.ParentYangName = "police"
    policeInfo.EntityData.SegmentPath = "police-info" + types.AddNoKeyToken(policeInfo)
    policeInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/police/" + policeInfo.EntityData.SegmentPath
    policeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeInfo.EntityData.Children = types.NewOrderedMap()
    policeInfo.EntityData.Leafs = types.NewOrderedMap()
    policeInfo.EntityData.Leafs.Append("avgrate", types.YLeaf{"Avgrate", policeInfo.Avgrate})
    policeInfo.EntityData.Leafs.Append("burst", types.YLeaf{"Burst", policeInfo.Burst})
    policeInfo.EntityData.Leafs.Append("static-avgrate", types.YLeaf{"StaticAvgrate", policeInfo.StaticAvgrate})
    policeInfo.EntityData.Leafs.Append("avgrate-type", types.YLeaf{"AvgrateType", policeInfo.AvgrateType})
    policeInfo.EntityData.Leafs.Append("accepted-stats", types.YLeaf{"AcceptedStats", policeInfo.AcceptedStats})
    policeInfo.EntityData.Leafs.Append("dropped-stats", types.YLeaf{"DroppedStats", policeInfo.DroppedStats})
    policeInfo.EntityData.Leafs.Append("policer", types.YLeaf{"Policer", policeInfo.Policer})
    policeInfo.EntityData.Leafs.Append("iptos-value", types.YLeaf{"IptosValue", policeInfo.IptosValue})
    policeInfo.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", policeInfo.ChangeType})
    policeInfo.EntityData.Leafs.Append("acl-config", types.YLeaf{"AclConfig", policeInfo.AclConfig})
    policeInfo.EntityData.Leafs.Append("acl-str", types.YLeaf{"AclStr", policeInfo.AclStr})

    policeInfo.EntityData.YListKeys = []string {}

    return &(policeInfo.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_StaticPolice
// Static Police details
type LptsPifib_Nodes_Node_Hardware_StaticPolice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per punt reason info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo.
    StaticInfo []*LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo
}

func (staticPolice *LptsPifib_Nodes_Node_Hardware_StaticPolice) GetEntityData() *types.CommonEntityData {
    staticPolice.EntityData.YFilter = staticPolice.YFilter
    staticPolice.EntityData.YangName = "static-police"
    staticPolice.EntityData.BundleName = "cisco_ios_xr"
    staticPolice.EntityData.ParentYangName = "hardware"
    staticPolice.EntityData.SegmentPath = "static-police"
    staticPolice.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + staticPolice.EntityData.SegmentPath
    staticPolice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticPolice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticPolice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticPolice.EntityData.Children = types.NewOrderedMap()
    staticPolice.EntityData.Children.Append("static-info", types.YChild{"StaticInfo", nil})
    for i := range staticPolice.StaticInfo {
        types.SetYListKey(staticPolice.StaticInfo[i], i)
        staticPolice.EntityData.Children.Append(types.GetSegmentPath(staticPolice.StaticInfo[i]), types.YChild{"StaticInfo", staticPolice.StaticInfo[i]})
    }
    staticPolice.EntityData.Leafs = types.NewOrderedMap()

    staticPolice.EntityData.YListKeys = []string {}

    return &(staticPolice.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo
// Per punt reason info
type LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    PuntReasonString interface{}

    // change type. The type is interface{} with range: 0..255.
    ChangeType interface{}
}

func (staticInfo *LptsPifib_Nodes_Node_Hardware_StaticPolice_StaticInfo) GetEntityData() *types.CommonEntityData {
    staticInfo.EntityData.YFilter = staticInfo.YFilter
    staticInfo.EntityData.YangName = "static-info"
    staticInfo.EntityData.BundleName = "cisco_ios_xr"
    staticInfo.EntityData.ParentYangName = "static-police"
    staticInfo.EntityData.SegmentPath = "static-info" + types.AddNoKeyToken(staticInfo)
    staticInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/static-police/" + staticInfo.EntityData.SegmentPath
    staticInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticInfo.EntityData.Children = types.NewOrderedMap()
    staticInfo.EntityData.Leafs = types.NewOrderedMap()
    staticInfo.EntityData.Leafs.Append("punt-reason", types.YLeaf{"PuntReason", staticInfo.PuntReason})
    staticInfo.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", staticInfo.Sid})
    staticInfo.EntityData.Leafs.Append("flow-rate", types.YLeaf{"FlowRate", staticInfo.FlowRate})
    staticInfo.EntityData.Leafs.Append("burst-rate", types.YLeaf{"BurstRate", staticInfo.BurstRate})
    staticInfo.EntityData.Leafs.Append("accepted", types.YLeaf{"Accepted", staticInfo.Accepted})
    staticInfo.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", staticInfo.Dropped})
    staticInfo.EntityData.Leafs.Append("punt-reason-string", types.YLeaf{"PuntReasonString", staticInfo.PuntReasonString})
    staticInfo.EntityData.Leafs.Append("change-type", types.YLeaf{"ChangeType", staticInfo.ChangeType})

    staticInfo.EntityData.YListKeys = []string {}

    return &(staticInfo.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_Bfd
// Bfd details
type LptsPifib_Nodes_Node_Hardware_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per bfd disc entry info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo.
    BfdEntryInfo []*LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo
}

func (bfd *LptsPifib_Nodes_Node_Hardware_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "hardware"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Children.Append("bfd-entry-info", types.YChild{"BfdEntryInfo", nil})
    for i := range bfd.BfdEntryInfo {
        types.SetYListKey(bfd.BfdEntryInfo[i], i)
        bfd.EntityData.Children.Append(types.GetSegmentPath(bfd.BfdEntryInfo[i]), types.YChild{"BfdEntryInfo", bfd.BfdEntryInfo[i]})
    }
    bfd.EntityData.Leafs = types.NewOrderedMap()

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo
// Per bfd disc entry info
type LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (bfdEntryInfo *LptsPifib_Nodes_Node_Hardware_Bfd_BfdEntryInfo) GetEntityData() *types.CommonEntityData {
    bfdEntryInfo.EntityData.YFilter = bfdEntryInfo.YFilter
    bfdEntryInfo.EntityData.YangName = "bfd-entry-info"
    bfdEntryInfo.EntityData.BundleName = "cisco_ios_xr"
    bfdEntryInfo.EntityData.ParentYangName = "bfd"
    bfdEntryInfo.EntityData.SegmentPath = "bfd-entry-info" + types.AddNoKeyToken(bfdEntryInfo)
    bfdEntryInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/bfd/" + bfdEntryInfo.EntityData.SegmentPath
    bfdEntryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdEntryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdEntryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdEntryInfo.EntityData.Children = types.NewOrderedMap()
    bfdEntryInfo.EntityData.Leafs = types.NewOrderedMap()
    bfdEntryInfo.EntityData.Leafs.Append("index", types.YLeaf{"Index", bfdEntryInfo.Index})
    bfdEntryInfo.EntityData.Leafs.Append("is-mcast", types.YLeaf{"IsMcast", bfdEntryInfo.IsMcast})
    bfdEntryInfo.EntityData.Leafs.Append("fgid-or-vqi", types.YLeaf{"FgidOrVqi", bfdEntryInfo.FgidOrVqi})
    bfdEntryInfo.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", bfdEntryInfo.IsValid})
    bfdEntryInfo.EntityData.Leafs.Append("policer-id", types.YLeaf{"PolicerId", bfdEntryInfo.PolicerId})

    bfdEntryInfo.EntityData.YListKeys = []string {}

    return &(bfdEntryInfo.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_Statistics
// Hardware Entry type
type LptsPifib_Nodes_Node_Hardware_Statistics struct {
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

func (statistics *LptsPifib_Nodes_Node_Hardware_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "hardware"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("accepted", types.YLeaf{"Accepted", statistics.Accepted})
    statistics.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", statistics.Dropped})
    statistics.EntityData.Leafs.Append("clear-ts", types.YLeaf{"ClearTs", statistics.ClearTs})
    statistics.EntityData.Leafs.Append("no-stats-mem-err", types.YLeaf{"NoStatsMemErr", statistics.NoStatsMemErr})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_IndexEntries
// Hardware Entry options
type LptsPifib_Nodes_Node_Hardware_IndexEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry options. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry.
    IndexEntry []*LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry
}

func (indexEntries *LptsPifib_Nodes_Node_Hardware_IndexEntries) GetEntityData() *types.CommonEntityData {
    indexEntries.EntityData.YFilter = indexEntries.YFilter
    indexEntries.EntityData.YangName = "index-entries"
    indexEntries.EntityData.BundleName = "cisco_ios_xr"
    indexEntries.EntityData.ParentYangName = "hardware"
    indexEntries.EntityData.SegmentPath = "index-entries"
    indexEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/" + indexEntries.EntityData.SegmentPath
    indexEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexEntries.EntityData.Children = types.NewOrderedMap()
    indexEntries.EntityData.Children.Append("index-entry", types.YChild{"IndexEntry", nil})
    for i := range indexEntries.IndexEntry {
        indexEntries.EntityData.Children.Append(types.GetSegmentPath(indexEntries.IndexEntry[i]), types.YChild{"IndexEntry", indexEntries.IndexEntry[i]})
    }
    indexEntries.EntityData.Leafs = types.NewOrderedMap()

    indexEntries.EntityData.YListKeys = []string {}

    return &(indexEntries.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry
// Entry options
type LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Index. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Layer 3 Protocol. The type is interface{} with range: 0..4294967295.
    L3protocol interface{}

    // Layer 4 Protocol. The type is interface{} with range: 0..4294967295.
    L4protocol interface{}

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
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    AclStr interface{}

    // Stats not available. The type is interface{} with range: 0..255.
    NoStats interface{}

    // Per pipe type hardware info. The type is slice of
    // LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo.
    HwInfo []*LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
}

func (indexEntry *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry) GetEntityData() *types.CommonEntityData {
    indexEntry.EntityData.YFilter = indexEntry.YFilter
    indexEntry.EntityData.YangName = "index-entry"
    indexEntry.EntityData.BundleName = "cisco_ios_xr"
    indexEntry.EntityData.ParentYangName = "index-entries"
    indexEntry.EntityData.SegmentPath = "index-entry" + types.AddKeyToken(indexEntry.Index, "index")
    indexEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/index-entries/" + indexEntry.EntityData.SegmentPath
    indexEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexEntry.EntityData.Children = types.NewOrderedMap()
    indexEntry.EntityData.Children.Append("hw-info", types.YChild{"HwInfo", nil})
    for i := range indexEntry.HwInfo {
        types.SetYListKey(indexEntry.HwInfo[i], i)
        indexEntry.EntityData.Children.Append(types.GetSegmentPath(indexEntry.HwInfo[i]), types.YChild{"HwInfo", indexEntry.HwInfo[i]})
    }
    indexEntry.EntityData.Leafs = types.NewOrderedMap()
    indexEntry.EntityData.Leafs.Append("index", types.YLeaf{"Index", indexEntry.Index})
    indexEntry.EntityData.Leafs.Append("l3protocol", types.YLeaf{"L3protocol", indexEntry.L3protocol})
    indexEntry.EntityData.Leafs.Append("l4protocol", types.YLeaf{"L4protocol", indexEntry.L4protocol})
    indexEntry.EntityData.Leafs.Append("intf-handle", types.YLeaf{"IntfHandle", indexEntry.IntfHandle})
    indexEntry.EntityData.Leafs.Append("intf-name", types.YLeaf{"IntfName", indexEntry.IntfName})
    indexEntry.EntityData.Leafs.Append("uidb-index", types.YLeaf{"UidbIndex", indexEntry.UidbIndex})
    indexEntry.EntityData.Leafs.Append("local-addr", types.YLeaf{"LocalAddr", indexEntry.LocalAddr})
    indexEntry.EntityData.Leafs.Append("local-prefix-len", types.YLeaf{"LocalPrefixLen", indexEntry.LocalPrefixLen})
    indexEntry.EntityData.Leafs.Append("remote-addr", types.YLeaf{"RemoteAddr", indexEntry.RemoteAddr})
    indexEntry.EntityData.Leafs.Append("remote-prefix-len", types.YLeaf{"RemotePrefixLen", indexEntry.RemotePrefixLen})
    indexEntry.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", indexEntry.VrfId})
    indexEntry.EntityData.Leafs.Append("u-value", types.YLeaf{"UValue", indexEntry.UValue})
    indexEntry.EntityData.Leafs.Append("u-len", types.YLeaf{"ULen", indexEntry.ULen})
    indexEntry.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", indexEntry.LocalPort})
    indexEntry.EntityData.Leafs.Append("is-frag", types.YLeaf{"IsFrag", indexEntry.IsFrag})
    indexEntry.EntityData.Leafs.Append("is-syn", types.YLeaf{"IsSyn", indexEntry.IsSyn})
    indexEntry.EntityData.Leafs.Append("action", types.YLeaf{"Action", indexEntry.Action})
    indexEntry.EntityData.Leafs.Append("action-string", types.YLeaf{"ActionString", indexEntry.ActionString})
    indexEntry.EntityData.Leafs.Append("listener-tag", types.YLeaf{"ListenerTag", indexEntry.ListenerTag})
    indexEntry.EntityData.Leafs.Append("is-fgid", types.YLeaf{"IsFgid", indexEntry.IsFgid})
    indexEntry.EntityData.Leafs.Append("is-vrf", types.YLeaf{"IsVrf", indexEntry.IsVrf})
    indexEntry.EntityData.Leafs.Append("is-optimized", types.YLeaf{"IsOptimized", indexEntry.IsOptimized})
    indexEntry.EntityData.Leafs.Append("is-uidb-opt-bit", types.YLeaf{"IsUidbOptBit", indexEntry.IsUidbOptBit})
    indexEntry.EntityData.Leafs.Append("fgid-or-sfp", types.YLeaf{"FgidOrSfp", indexEntry.FgidOrSfp})
    indexEntry.EntityData.Leafs.Append("remote-rack", types.YLeaf{"RemoteRack", indexEntry.RemoteRack})
    indexEntry.EntityData.Leafs.Append("rack-id", types.YLeaf{"RackId", indexEntry.RackId})
    indexEntry.EntityData.Leafs.Append("rslot", types.YLeaf{"Rslot", indexEntry.Rslot})
    indexEntry.EntityData.Leafs.Append("cir", types.YLeaf{"Cir", indexEntry.Cir})
    indexEntry.EntityData.Leafs.Append("flow-type", types.YLeaf{"FlowType", indexEntry.FlowType})
    indexEntry.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", indexEntry.Priority})
    indexEntry.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", indexEntry.Sid})
    indexEntry.EntityData.Leafs.Append("policer-avgrate", types.YLeaf{"PolicerAvgrate", indexEntry.PolicerAvgrate})
    indexEntry.EntityData.Leafs.Append("policer-burst", types.YLeaf{"PolicerBurst", indexEntry.PolicerBurst})
    indexEntry.EntityData.Leafs.Append("lookup-priority", types.YLeaf{"LookupPriority", indexEntry.LookupPriority})
    indexEntry.EntityData.Leafs.Append("storage-priority", types.YLeaf{"StoragePriority", indexEntry.StoragePriority})
    indexEntry.EntityData.Leafs.Append("num-tm-entries", types.YLeaf{"NumTmEntries", indexEntry.NumTmEntries})
    indexEntry.EntityData.Leafs.Append("entry-ptr", types.YLeaf{"EntryPtr", indexEntry.EntryPtr})
    indexEntry.EntityData.Leafs.Append("entry-shadow-ptr", types.YLeaf{"EntryShadowPtr", indexEntry.EntryShadowPtr})
    indexEntry.EntityData.Leafs.Append("list-node-ptr", types.YLeaf{"ListNodePtr", indexEntry.ListNodePtr})
    indexEntry.EntityData.Leafs.Append("state", types.YLeaf{"State", indexEntry.State})
    indexEntry.EntityData.Leafs.Append("retry-fail-cause", types.YLeaf{"RetryFailCause", indexEntry.RetryFailCause})
    indexEntry.EntityData.Leafs.Append("num-retries", types.YLeaf{"NumRetries", indexEntry.NumRetries})
    indexEntry.EntityData.Leafs.Append("min-ttl", types.YLeaf{"MinTtl", indexEntry.MinTtl})
    indexEntry.EntityData.Leafs.Append("u-type", types.YLeaf{"UType", indexEntry.UType})
    indexEntry.EntityData.Leafs.Append("remote-fgid", types.YLeaf{"RemoteFgid", indexEntry.RemoteFgid})
    indexEntry.EntityData.Leafs.Append("acl-str", types.YLeaf{"AclStr", indexEntry.AclStr})
    indexEntry.EntityData.Leafs.Append("no-stats", types.YLeaf{"NoStats", indexEntry.NoStats})

    indexEntry.EntityData.YListKeys = []string {"Index"}

    return &(indexEntry.EntityData)
}

// LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo
// Per pipe type hardware info
type LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hwInfo *LptsPifib_Nodes_Node_Hardware_IndexEntries_IndexEntry_HwInfo) GetEntityData() *types.CommonEntityData {
    hwInfo.EntityData.YFilter = hwInfo.YFilter
    hwInfo.EntityData.YangName = "hw-info"
    hwInfo.EntityData.BundleName = "cisco_ios_xr"
    hwInfo.EntityData.ParentYangName = "index-entry"
    hwInfo.EntityData.SegmentPath = "hw-info" + types.AddNoKeyToken(hwInfo)
    hwInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-lpts-pre-ifib-oper:lpts-pifib/nodes/node/Cisco-IOS-XR-platform-pifib-oper:hardware/index-entries/index-entry/" + hwInfo.EntityData.SegmentPath
    hwInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwInfo.EntityData.Children = types.NewOrderedMap()
    hwInfo.EntityData.Leafs = types.NewOrderedMap()
    hwInfo.EntityData.Leafs.Append("policer", types.YLeaf{"Policer", hwInfo.Policer})
    hwInfo.EntityData.Leafs.Append("stats-ptr", types.YLeaf{"StatsPtr", hwInfo.StatsPtr})
    hwInfo.EntityData.Leafs.Append("accepted", types.YLeaf{"Accepted", hwInfo.Accepted})
    hwInfo.EntityData.Leafs.Append("dropped", types.YLeaf{"Dropped", hwInfo.Dropped})
    hwInfo.EntityData.Leafs.Append("sort-start-offset", types.YLeaf{"SortStartOffset", hwInfo.SortStartOffset})
    hwInfo.EntityData.Leafs.Append("tm-start-offset", types.YLeaf{"TmStartOffset", hwInfo.TmStartOffset})

    hwInfo.EntityData.YListKeys = []string {}

    return &(hwInfo.EntityData)
}

