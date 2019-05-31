// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-prm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-qos-mode: QoS mode in hardware module port(s)
//   hardware-module-processor: hardware module processor
//   hardware-module-tcp-mss-adjust: hardware module tcp mss adjust
//   hardware-module-tcam: hardware module tcam
//   hardware-module-profile: hardware module profile
//   hardware-module-efd: hardware module efd
//   hardware-module-all-qos-mode: hardware module all qos mode
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_prm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_prm_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-qos-mode}", reflect.TypeOf(HardwareModuleQosMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode", reflect.TypeOf(HardwareModuleQosMode{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-processor}", reflect.TypeOf(HardwareModuleProcessor{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-processor", reflect.TypeOf(HardwareModuleProcessor{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcp-mss-adjust}", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcam}", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-profile}", reflect.TypeOf(HardwareModuleProfile{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-profile", reflect.TypeOf(HardwareModuleProfile{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-efd}", reflect.TypeOf(HardwareModuleEfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd", reflect.TypeOf(HardwareModuleEfd{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-all-qos-mode}", reflect.TypeOf(HardwareModuleAllQosMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-all-qos-mode", reflect.TypeOf(HardwareModuleAllQosMode{}))
}

// AdminPrmConfigInternalTcamPartProfile represents Admin prm config internal tcam part profile
type AdminPrmConfigInternalTcamPartProfile string

const (
    // default internal tcam partitions (L2 1k, V4 24k
    // , V6 1.75k entries): Tomahawk
    AdminPrmConfigInternalTcamPartProfile_to_default AdminPrmConfigInternalTcamPartProfile = "to-default"

    // Set internal tcam partitions for service edge
    // (L2 4k, V4 15k, V6 3.25k entries): Tomahawk
    AdminPrmConfigInternalTcamPartProfile_to_profile_se1 AdminPrmConfigInternalTcamPartProfile = "to-profile-se1"
)

// AdminPrmConfigPackageBundle represents Admin prm config package bundle
type AdminPrmConfigPackageBundle string

const (
    // Default Package
    AdminPrmConfigPackageBundle_default_ AdminPrmConfigPackageBundle = "default"

    // Services Package
    AdminPrmConfigPackageBundle_services AdminPrmConfigPackageBundle = "services"
)

// AdminPrmConfigTcamPartProfile represents Admin prm config tcam part profile
type AdminPrmConfigTcamPartProfile string

const (
    // Default tcam partition ods2:ods8 60:40
    AdminPrmConfigTcamPartProfile_default_ AdminPrmConfigTcamPartProfile = "default"

    // Tcam Partition ods2:ods8 30:70
    AdminPrmConfigTcamPartProfile_ods2_30_ods8_70 AdminPrmConfigTcamPartProfile = "ods2-30-ods8-70"

    // Tcam Partition ods2:ods8 40:60
    AdminPrmConfigTcamPartProfile_ods2_40_ods8_60 AdminPrmConfigTcamPartProfile = "ods2-40-ods8-60"

    // Tcam Partition ods2:ods8 50:50
    AdminPrmConfigTcamPartProfile_ods2_50_ods8_50 AdminPrmConfigTcamPartProfile = "ods2-50-ods8-50"

    // Tcam Partition ods2:ods8 70:30
    AdminPrmConfigTcamPartProfile_ods2_70_ods8_30 AdminPrmConfigTcamPartProfile = "ods2-70-ods8-30"
)

// AdminPrmConfigFeatureProfile represents Admin prm config feature profile
type AdminPrmConfigFeatureProfile string

const (
    // Default feature profile
    AdminPrmConfigFeatureProfile_default_ AdminPrmConfigFeatureProfile = "default"

    // L2 feature profile
    AdminPrmConfigFeatureProfile_l2 AdminPrmConfigFeatureProfile = "l2"
)

// PrmProcessorConfig represents Prm processor config
type PrmProcessorConfig string

const (
    // Default cluster setting
    PrmProcessorConfig_mode_default PrmProcessorConfig = "mode-default"

    // Full cluster setting
    PrmProcessorConfig_mode_full PrmProcessorConfig = "mode-full"
)

// Asr9kEfdOperation represents Asr9k efd operation
type Asr9kEfdOperation string

const (
    // Less than
    Asr9kEfdOperation_less_than Asr9kEfdOperation = "less-than"

    // Greater than or equal
    Asr9kEfdOperation_greater_than_or_equal Asr9kEfdOperation = "greater-than-or-equal"
)

// Asr9kEfdMode represents Asr9k efd mode
type Asr9kEfdMode string

const (
    // Only check outer encap
    Asr9kEfdMode_only_outer_encap Asr9kEfdMode = "only-outer-encap"

    // Check outer and inner encap
    Asr9kEfdMode_include_inner_encap Asr9kEfdMode = "include-inner-encap"
)

// AdminPrmConfigScaleProfile represents Admin prm config scale profile
type AdminPrmConfigScaleProfile string

const (
    // Default scale profile
    AdminPrmConfigScaleProfile_default_ AdminPrmConfigScaleProfile = "default"

    // L2 scale profile
    AdminPrmConfigScaleProfile_l2 AdminPrmConfigScaleProfile = "l2"

    // L3 scale profile
    AdminPrmConfigScaleProfile_l3 AdminPrmConfigScaleProfile = "l3"

    // L3XL scale profile
    AdminPrmConfigScaleProfile_l3xl AdminPrmConfigScaleProfile = "l3xl"

    // BNG scale profile
    AdminPrmConfigScaleProfile_bng AdminPrmConfigScaleProfile = "bng"

    // LSR scale profile
    AdminPrmConfigScaleProfile_lsr AdminPrmConfigScaleProfile = "lsr"

    // SAT scale profile
    AdminPrmConfigScaleProfile_sat AdminPrmConfigScaleProfile = "sat"

    // Single-flow perf scale profile
    AdminPrmConfigScaleProfile_sfp AdminPrmConfigScaleProfile = "sfp"
)

// PrmTcamProfile represents Prm tcam profile
type PrmTcamProfile string

const (
    // Profile 0
    PrmTcamProfile_profile0 PrmTcamProfile = "profile0"

    // Profile 1
    PrmTcamProfile_profile1 PrmTcamProfile = "profile1"

    // Profile 2
    PrmTcamProfile_profile2 PrmTcamProfile = "profile2"
)

// HardwareModuleQosMode
// QoS mode in hardware module port(s)
type HardwareModuleQosMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS applicable nodes.
    Nodes HardwareModuleQosMode_Nodes
}

func (hardwareModuleQosMode *HardwareModuleQosMode) GetEntityData() *types.CommonEntityData {
    hardwareModuleQosMode.EntityData.YFilter = hardwareModuleQosMode.YFilter
    hardwareModuleQosMode.EntityData.YangName = "hardware-module-qos-mode"
    hardwareModuleQosMode.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleQosMode.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleQosMode.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode"
    hardwareModuleQosMode.EntityData.AbsolutePath = hardwareModuleQosMode.EntityData.SegmentPath
    hardwareModuleQosMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleQosMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleQosMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleQosMode.EntityData.Children = types.NewOrderedMap()
    hardwareModuleQosMode.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleQosMode.Nodes})
    hardwareModuleQosMode.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleQosMode.EntityData.YListKeys = []string {}

    return &(hardwareModuleQosMode.EntityData)
}

// HardwareModuleQosMode_Nodes
// QoS applicable nodes
type HardwareModuleQosMode_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleQosMode_Nodes_Node.
    Node []*HardwareModuleQosMode_Nodes_Node
}

func (nodes *HardwareModuleQosMode_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-qos-mode"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode/" + nodes.EntityData.SegmentPath
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

// HardwareModuleQosMode_Nodes_Node
// A node
type HardwareModuleQosMode_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Disable child level/flat policy shaping. The type is interface{}.
    ChildShapingDisable interface{}

    // Enable low burst mode for TM entity. The type is interface{}.
    LowburstEnable interface{}
}

func (node *HardwareModuleQosMode_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-qos-mode/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("child-shaping-disable", types.YLeaf{"ChildShapingDisable", node.ChildShapingDisable})
    node.EntityData.Leafs.Append("lowburst-enable", types.YLeaf{"LowburstEnable", node.LowburstEnable})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleProcessor
// hardware module processor
type HardwareModuleProcessor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // applicable nodeTable.
    Nodes HardwareModuleProcessor_Nodes
}

func (hardwareModuleProcessor *HardwareModuleProcessor) GetEntityData() *types.CommonEntityData {
    hardwareModuleProcessor.EntityData.YFilter = hardwareModuleProcessor.YFilter
    hardwareModuleProcessor.EntityData.YangName = "hardware-module-processor"
    hardwareModuleProcessor.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleProcessor.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleProcessor.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-processor"
    hardwareModuleProcessor.EntityData.AbsolutePath = hardwareModuleProcessor.EntityData.SegmentPath
    hardwareModuleProcessor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleProcessor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleProcessor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleProcessor.EntityData.Children = types.NewOrderedMap()
    hardwareModuleProcessor.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleProcessor.Nodes})
    hardwareModuleProcessor.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleProcessor.EntityData.YListKeys = []string {}

    return &(hardwareModuleProcessor.EntityData)
}

// HardwareModuleProcessor_Nodes
// applicable nodeTable
type HardwareModuleProcessor_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // applicable node. The type is slice of HardwareModuleProcessor_Nodes_Node.
    Node []*HardwareModuleProcessor_Nodes_Node
}

func (nodes *HardwareModuleProcessor_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-processor"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-processor/" + nodes.EntityData.SegmentPath
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

// HardwareModuleProcessor_Nodes_Node
// applicable node
type HardwareModuleProcessor_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Processor mode setting. The type is PrmProcessorConfig.
    Mode interface{}
}

func (node *HardwareModuleProcessor_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-processor/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", node.Mode})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleTcpMssAdjust
// hardware module tcp mss adjust
type HardwareModuleTcpMssAdjust struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP MSS Adjust applicable nodes.
    Nodes HardwareModuleTcpMssAdjust_Nodes
}

func (hardwareModuleTcpMssAdjust *HardwareModuleTcpMssAdjust) GetEntityData() *types.CommonEntityData {
    hardwareModuleTcpMssAdjust.EntityData.YFilter = hardwareModuleTcpMssAdjust.YFilter
    hardwareModuleTcpMssAdjust.EntityData.YangName = "hardware-module-tcp-mss-adjust"
    hardwareModuleTcpMssAdjust.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleTcpMssAdjust.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleTcpMssAdjust.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust"
    hardwareModuleTcpMssAdjust.EntityData.AbsolutePath = hardwareModuleTcpMssAdjust.EntityData.SegmentPath
    hardwareModuleTcpMssAdjust.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleTcpMssAdjust.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleTcpMssAdjust.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleTcpMssAdjust.EntityData.Children = types.NewOrderedMap()
    hardwareModuleTcpMssAdjust.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleTcpMssAdjust.Nodes})
    hardwareModuleTcpMssAdjust.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleTcpMssAdjust.EntityData.YListKeys = []string {}

    return &(hardwareModuleTcpMssAdjust.EntityData)
}

// HardwareModuleTcpMssAdjust_Nodes
// TCP MSS Adjust applicable nodes
type HardwareModuleTcpMssAdjust_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleTcpMssAdjust_Nodes_Node.
    Node []*HardwareModuleTcpMssAdjust_Nodes_Node
}

func (nodes *HardwareModuleTcpMssAdjust_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-tcp-mss-adjust"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust/" + nodes.EntityData.SegmentPath
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

// HardwareModuleTcpMssAdjust_Nodes_Node
// A node
type HardwareModuleTcpMssAdjust_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // TCP MSS Adjust NPs.
    Nps HardwareModuleTcpMssAdjust_Nodes_Node_Nps
}

func (node *HardwareModuleTcpMssAdjust_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("nps", types.YChild{"Nps", &node.Nps})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleTcpMssAdjust_Nodes_Node_Nps
// TCP MSS Adjust NPs
type HardwareModuleTcpMssAdjust_Nodes_Node_Nps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NP number. The type is slice of
    // HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np.
    Np []*HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np
}

func (nps *HardwareModuleTcpMssAdjust_Nodes_Node_Nps) GetEntityData() *types.CommonEntityData {
    nps.EntityData.YFilter = nps.YFilter
    nps.EntityData.YangName = "nps"
    nps.EntityData.BundleName = "cisco_ios_xr"
    nps.EntityData.ParentYangName = "node"
    nps.EntityData.SegmentPath = "nps"
    nps.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust/nodes/node/" + nps.EntityData.SegmentPath
    nps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nps.EntityData.Children = types.NewOrderedMap()
    nps.EntityData.Children.Append("np", types.YChild{"Np", nil})
    for i := range nps.Np {
        nps.EntityData.Children.Append(types.GetSegmentPath(nps.Np[i]), types.YChild{"Np", nps.Np[i]})
    }
    nps.EntityData.Leafs = types.NewOrderedMap()

    nps.EntityData.YListKeys = []string {}

    return &(nps.EntityData)
}

// HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np
// NP number
type HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Number between 0-7. The type is interface{} with
    // range: 0..7.
    NpId interface{}

    // TCP MSS Adjust value. The type is interface{} with range: 1280..1535. Units
    // are byte.
    AdjustValue interface{}
}

func (np *HardwareModuleTcpMssAdjust_Nodes_Node_Nps_Np) GetEntityData() *types.CommonEntityData {
    np.EntityData.YFilter = np.YFilter
    np.EntityData.YangName = "np"
    np.EntityData.BundleName = "cisco_ios_xr"
    np.EntityData.ParentYangName = "nps"
    np.EntityData.SegmentPath = "np" + types.AddKeyToken(np.NpId, "np-id")
    np.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust/nodes/node/nps/" + np.EntityData.SegmentPath
    np.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    np.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    np.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    np.EntityData.Children = types.NewOrderedMap()
    np.EntityData.Leafs = types.NewOrderedMap()
    np.EntityData.Leafs.Append("np-id", types.YLeaf{"NpId", np.NpId})
    np.EntityData.Leafs.Append("adjust-value", types.YLeaf{"AdjustValue", np.AdjustValue})

    np.EntityData.YListKeys = []string {"NpId"}

    return &(np.EntityData)
}

// HardwareModuleTcam
// hardware module tcam
type HardwareModuleTcam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global TCAM partition profile for all TCAM applicable nodes. The type is
    // PrmTcamProfile. The default value is profile0.
    GlobalProfile interface{}

    // TCAM applicable nodes.
    Nodes HardwareModuleTcam_Nodes
}

func (hardwareModuleTcam *HardwareModuleTcam) GetEntityData() *types.CommonEntityData {
    hardwareModuleTcam.EntityData.YFilter = hardwareModuleTcam.YFilter
    hardwareModuleTcam.EntityData.YangName = "hardware-module-tcam"
    hardwareModuleTcam.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleTcam.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleTcam.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam"
    hardwareModuleTcam.EntityData.AbsolutePath = hardwareModuleTcam.EntityData.SegmentPath
    hardwareModuleTcam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleTcam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleTcam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleTcam.EntityData.Children = types.NewOrderedMap()
    hardwareModuleTcam.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleTcam.Nodes})
    hardwareModuleTcam.EntityData.Leafs = types.NewOrderedMap()
    hardwareModuleTcam.EntityData.Leafs.Append("global-profile", types.YLeaf{"GlobalProfile", hardwareModuleTcam.GlobalProfile})

    hardwareModuleTcam.EntityData.YListKeys = []string {}

    return &(hardwareModuleTcam.EntityData)
}

// HardwareModuleTcam_Nodes
// TCAM applicable nodes
type HardwareModuleTcam_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A TCAM applicable node. The type is slice of HardwareModuleTcam_Nodes_Node.
    Node []*HardwareModuleTcam_Nodes_Node
}

func (nodes *HardwareModuleTcam_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-tcam"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam/" + nodes.EntityData.SegmentPath
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

// HardwareModuleTcam_Nodes_Node
// A TCAM applicable node
type HardwareModuleTcam_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // A TCAM partition profile. The type is PrmTcamProfile. The default value is
    // profile0.
    Profile interface{}
}

func (node *HardwareModuleTcam_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", node.Profile})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleProfile
// hardware module profile
type HardwareModuleProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Memory resource profile feature. The type is AdminPrmConfigFeatureProfile.
    // The default value is default.
    Feature interface{}

    // Memory resource profile scale active. The type is
    // AdminPrmConfigScaleProfile. The default value is default.
    ScaleActive interface{}

    // Services Package. The type is AdminPrmConfigPackageBundle. The default
    // value is default.
    PackageBundle interface{}

    // Memory resource profile feature active. The type is
    // AdminPrmConfigFeatureProfile. The default value is default.
    FeatureActive interface{}

    // Memory resource profile scale. The type is AdminPrmConfigScaleProfile. The
    // default value is default.
    Scale interface{}

    // TCAM partition sizing applicable nodes.
    Nodes HardwareModuleProfile_Nodes
}

func (hardwareModuleProfile *HardwareModuleProfile) GetEntityData() *types.CommonEntityData {
    hardwareModuleProfile.EntityData.YFilter = hardwareModuleProfile.YFilter
    hardwareModuleProfile.EntityData.YangName = "hardware-module-profile"
    hardwareModuleProfile.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleProfile.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleProfile.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-profile"
    hardwareModuleProfile.EntityData.AbsolutePath = hardwareModuleProfile.EntityData.SegmentPath
    hardwareModuleProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleProfile.EntityData.Children = types.NewOrderedMap()
    hardwareModuleProfile.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleProfile.Nodes})
    hardwareModuleProfile.EntityData.Leafs = types.NewOrderedMap()
    hardwareModuleProfile.EntityData.Leafs.Append("feature", types.YLeaf{"Feature", hardwareModuleProfile.Feature})
    hardwareModuleProfile.EntityData.Leafs.Append("scale-active", types.YLeaf{"ScaleActive", hardwareModuleProfile.ScaleActive})
    hardwareModuleProfile.EntityData.Leafs.Append("package-bundle", types.YLeaf{"PackageBundle", hardwareModuleProfile.PackageBundle})
    hardwareModuleProfile.EntityData.Leafs.Append("feature-active", types.YLeaf{"FeatureActive", hardwareModuleProfile.FeatureActive})
    hardwareModuleProfile.EntityData.Leafs.Append("scale", types.YLeaf{"Scale", hardwareModuleProfile.Scale})

    hardwareModuleProfile.EntityData.YListKeys = []string {}

    return &(hardwareModuleProfile.EntityData)
}

// HardwareModuleProfile_Nodes
// TCAM partition sizing applicable nodes
type HardwareModuleProfile_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A TCAM partition sizing applicable node. The type is slice of
    // HardwareModuleProfile_Nodes_Node.
    Node []*HardwareModuleProfile_Nodes_Node
}

func (nodes *HardwareModuleProfile_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-profile"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-profile/" + nodes.EntityData.SegmentPath
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

// HardwareModuleProfile_Nodes_Node
// A TCAM partition sizing applicable node
type HardwareModuleProfile_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node ID. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Tcam partition profile. The type is AdminPrmConfigTcamPartProfile. The
    // default value is default.
    TcamPartition interface{}

    // Internal Tcam partition profile. The type is
    // AdminPrmConfigInternalTcamPartProfile. The default value is to-default.
    InternalTcamPartition interface{}
}

func (node *HardwareModuleProfile_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-profile/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("tcam-partition", types.YLeaf{"TcamPartition", node.TcamPartition})
    node.EntityData.Leafs.Append("internal-tcam-partition", types.YLeaf{"InternalTcamPartition", node.InternalTcamPartition})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleEfd
// hardware module efd
type HardwareModuleEfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All nodes.
    NodeAll HardwareModuleEfd_NodeAll

    // EFD applicable nodes.
    Nodes HardwareModuleEfd_Nodes
}

func (hardwareModuleEfd *HardwareModuleEfd) GetEntityData() *types.CommonEntityData {
    hardwareModuleEfd.EntityData.YFilter = hardwareModuleEfd.YFilter
    hardwareModuleEfd.EntityData.YangName = "hardware-module-efd"
    hardwareModuleEfd.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleEfd.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleEfd.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd"
    hardwareModuleEfd.EntityData.AbsolutePath = hardwareModuleEfd.EntityData.SegmentPath
    hardwareModuleEfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleEfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleEfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleEfd.EntityData.Children = types.NewOrderedMap()
    hardwareModuleEfd.EntityData.Children.Append("node-all", types.YChild{"NodeAll", &hardwareModuleEfd.NodeAll})
    hardwareModuleEfd.EntityData.Children.Append("nodes", types.YChild{"Nodes", &hardwareModuleEfd.Nodes})
    hardwareModuleEfd.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleEfd.EntityData.YListKeys = []string {}

    return &(hardwareModuleEfd.EntityData)
}

// HardwareModuleEfd_NodeAll
// All nodes
type HardwareModuleEfd_NodeAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EFD for this node. The type is interface{}.
    Enable interface{}

    // EFD mode parameter. The type is Asr9kEfdMode.
    Mode interface{}

    // VLAN Priority Mask.
    VlanPriorityMask HardwareModuleEfd_NodeAll_VlanPriorityMask

    // EFD IP parameters.
    IpPrecedence HardwareModuleEfd_NodeAll_IpPrecedence

    // EFD VLAN parameters.
    VlanCos HardwareModuleEfd_NodeAll_VlanCos

    // IP Priority Mask.
    IpPriorityMask HardwareModuleEfd_NodeAll_IpPriorityMask

    // MPLS Priority Mask.
    MplsPriorityMask HardwareModuleEfd_NodeAll_MplsPriorityMask

    // EFD MPLS parameters.
    MplsExp HardwareModuleEfd_NodeAll_MplsExp
}

func (nodeAll *HardwareModuleEfd_NodeAll) GetEntityData() *types.CommonEntityData {
    nodeAll.EntityData.YFilter = nodeAll.YFilter
    nodeAll.EntityData.YangName = "node-all"
    nodeAll.EntityData.BundleName = "cisco_ios_xr"
    nodeAll.EntityData.ParentYangName = "hardware-module-efd"
    nodeAll.EntityData.SegmentPath = "node-all"
    nodeAll.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/" + nodeAll.EntityData.SegmentPath
    nodeAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeAll.EntityData.Children = types.NewOrderedMap()
    nodeAll.EntityData.Children.Append("vlan-priority-mask", types.YChild{"VlanPriorityMask", &nodeAll.VlanPriorityMask})
    nodeAll.EntityData.Children.Append("ip-precedence", types.YChild{"IpPrecedence", &nodeAll.IpPrecedence})
    nodeAll.EntityData.Children.Append("vlan-cos", types.YChild{"VlanCos", &nodeAll.VlanCos})
    nodeAll.EntityData.Children.Append("ip-priority-mask", types.YChild{"IpPriorityMask", &nodeAll.IpPriorityMask})
    nodeAll.EntityData.Children.Append("mpls-priority-mask", types.YChild{"MplsPriorityMask", &nodeAll.MplsPriorityMask})
    nodeAll.EntityData.Children.Append("mpls-exp", types.YChild{"MplsExp", &nodeAll.MplsExp})
    nodeAll.EntityData.Leafs = types.NewOrderedMap()
    nodeAll.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", nodeAll.Enable})
    nodeAll.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", nodeAll.Mode})

    nodeAll.EntityData.YListKeys = []string {}

    return &(nodeAll.EntityData)
}

// HardwareModuleEfd_NodeAll_VlanPriorityMask
// VLAN Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_VlanPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (vlanPriorityMask *HardwareModuleEfd_NodeAll_VlanPriorityMask) GetEntityData() *types.CommonEntityData {
    vlanPriorityMask.EntityData.YFilter = vlanPriorityMask.YFilter
    vlanPriorityMask.EntityData.YangName = "vlan-priority-mask"
    vlanPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    vlanPriorityMask.EntityData.ParentYangName = "node-all"
    vlanPriorityMask.EntityData.SegmentPath = "vlan-priority-mask"
    vlanPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + vlanPriorityMask.EntityData.SegmentPath
    vlanPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanPriorityMask.EntityData.Children = types.NewOrderedMap()
    vlanPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    vlanPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", vlanPriorityMask.Prec0})
    vlanPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", vlanPriorityMask.Prec1})
    vlanPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", vlanPriorityMask.Prec2})
    vlanPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", vlanPriorityMask.Prec3})
    vlanPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", vlanPriorityMask.Prec4})
    vlanPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", vlanPriorityMask.Prec5})
    vlanPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", vlanPriorityMask.Prec6})
    vlanPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", vlanPriorityMask.Prec7})

    vlanPriorityMask.EntityData.YListKeys = []string {}

    return &(vlanPriorityMask.EntityData)
}

// HardwareModuleEfd_NodeAll_IpPrecedence
// EFD IP parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_IpPrecedence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP TOS precedence threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Precedence interface{}

    // IP operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (ipPrecedence *HardwareModuleEfd_NodeAll_IpPrecedence) GetEntityData() *types.CommonEntityData {
    ipPrecedence.EntityData.YFilter = ipPrecedence.YFilter
    ipPrecedence.EntityData.YangName = "ip-precedence"
    ipPrecedence.EntityData.BundleName = "cisco_ios_xr"
    ipPrecedence.EntityData.ParentYangName = "node-all"
    ipPrecedence.EntityData.SegmentPath = "ip-precedence"
    ipPrecedence.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + ipPrecedence.EntityData.SegmentPath
    ipPrecedence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipPrecedence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipPrecedence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipPrecedence.EntityData.Children = types.NewOrderedMap()
    ipPrecedence.EntityData.Leafs = types.NewOrderedMap()
    ipPrecedence.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", ipPrecedence.Precedence})
    ipPrecedence.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", ipPrecedence.Operation})

    ipPrecedence.EntityData.YListKeys = []string {}

    return &(ipPrecedence.EntityData)
}

// HardwareModuleEfd_NodeAll_VlanCos
// EFD VLAN parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_VlanCos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // VLAN COS threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Cos interface{}

    // VLAN operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (vlanCos *HardwareModuleEfd_NodeAll_VlanCos) GetEntityData() *types.CommonEntityData {
    vlanCos.EntityData.YFilter = vlanCos.YFilter
    vlanCos.EntityData.YangName = "vlan-cos"
    vlanCos.EntityData.BundleName = "cisco_ios_xr"
    vlanCos.EntityData.ParentYangName = "node-all"
    vlanCos.EntityData.SegmentPath = "vlan-cos"
    vlanCos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + vlanCos.EntityData.SegmentPath
    vlanCos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanCos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanCos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanCos.EntityData.Children = types.NewOrderedMap()
    vlanCos.EntityData.Leafs = types.NewOrderedMap()
    vlanCos.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", vlanCos.Cos})
    vlanCos.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", vlanCos.Operation})

    vlanCos.EntityData.YListKeys = []string {}

    return &(vlanCos.EntityData)
}

// HardwareModuleEfd_NodeAll_IpPriorityMask
// IP Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_IpPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (ipPriorityMask *HardwareModuleEfd_NodeAll_IpPriorityMask) GetEntityData() *types.CommonEntityData {
    ipPriorityMask.EntityData.YFilter = ipPriorityMask.YFilter
    ipPriorityMask.EntityData.YangName = "ip-priority-mask"
    ipPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    ipPriorityMask.EntityData.ParentYangName = "node-all"
    ipPriorityMask.EntityData.SegmentPath = "ip-priority-mask"
    ipPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + ipPriorityMask.EntityData.SegmentPath
    ipPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipPriorityMask.EntityData.Children = types.NewOrderedMap()
    ipPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    ipPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", ipPriorityMask.Prec0})
    ipPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", ipPriorityMask.Prec1})
    ipPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", ipPriorityMask.Prec2})
    ipPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", ipPriorityMask.Prec3})
    ipPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", ipPriorityMask.Prec4})
    ipPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", ipPriorityMask.Prec5})
    ipPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", ipPriorityMask.Prec6})
    ipPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", ipPriorityMask.Prec7})

    ipPriorityMask.EntityData.YListKeys = []string {}

    return &(ipPriorityMask.EntityData)
}

// HardwareModuleEfd_NodeAll_MplsPriorityMask
// MPLS Priority Mask
// This type is a presence type.
type HardwareModuleEfd_NodeAll_MplsPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (mplsPriorityMask *HardwareModuleEfd_NodeAll_MplsPriorityMask) GetEntityData() *types.CommonEntityData {
    mplsPriorityMask.EntityData.YFilter = mplsPriorityMask.YFilter
    mplsPriorityMask.EntityData.YangName = "mpls-priority-mask"
    mplsPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    mplsPriorityMask.EntityData.ParentYangName = "node-all"
    mplsPriorityMask.EntityData.SegmentPath = "mpls-priority-mask"
    mplsPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + mplsPriorityMask.EntityData.SegmentPath
    mplsPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsPriorityMask.EntityData.Children = types.NewOrderedMap()
    mplsPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    mplsPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", mplsPriorityMask.Prec0})
    mplsPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", mplsPriorityMask.Prec1})
    mplsPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", mplsPriorityMask.Prec2})
    mplsPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", mplsPriorityMask.Prec3})
    mplsPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", mplsPriorityMask.Prec4})
    mplsPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", mplsPriorityMask.Prec5})
    mplsPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", mplsPriorityMask.Prec6})
    mplsPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", mplsPriorityMask.Prec7})

    mplsPriorityMask.EntityData.YListKeys = []string {}

    return &(mplsPriorityMask.EntityData)
}

// HardwareModuleEfd_NodeAll_MplsExp
// EFD MPLS parameters
// This type is a presence type.
type HardwareModuleEfd_NodeAll_MplsExp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // MPLS EXP threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Exp interface{}

    // MPLS operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (mplsExp *HardwareModuleEfd_NodeAll_MplsExp) GetEntityData() *types.CommonEntityData {
    mplsExp.EntityData.YFilter = mplsExp.YFilter
    mplsExp.EntityData.YangName = "mpls-exp"
    mplsExp.EntityData.BundleName = "cisco_ios_xr"
    mplsExp.EntityData.ParentYangName = "node-all"
    mplsExp.EntityData.SegmentPath = "mpls-exp"
    mplsExp.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/node-all/" + mplsExp.EntityData.SegmentPath
    mplsExp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsExp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsExp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsExp.EntityData.Children = types.NewOrderedMap()
    mplsExp.EntityData.Leafs = types.NewOrderedMap()
    mplsExp.EntityData.Leafs.Append("exp", types.YLeaf{"Exp", mplsExp.Exp})
    mplsExp.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", mplsExp.Operation})

    mplsExp.EntityData.YListKeys = []string {}

    return &(mplsExp.EntityData)
}

// HardwareModuleEfd_Nodes
// EFD applicable nodes
type HardwareModuleEfd_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A node. The type is slice of HardwareModuleEfd_Nodes_Node.
    Node []*HardwareModuleEfd_Nodes_Node
}

func (nodes *HardwareModuleEfd_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "hardware-module-efd"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/" + nodes.EntityData.SegmentPath
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

// HardwareModuleEfd_Nodes_Node
// A node
type HardwareModuleEfd_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node Name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Enable EFD for this node. The type is interface{}.
    Enable interface{}

    // EFD mode parameter. The type is Asr9kEfdMode.
    Mode interface{}

    // VLAN Priority Mask.
    VlanPriorityMask HardwareModuleEfd_Nodes_Node_VlanPriorityMask

    // EFD IP parameters.
    IpPrecedence HardwareModuleEfd_Nodes_Node_IpPrecedence

    // EFD VLAN parameters.
    VlanCos HardwareModuleEfd_Nodes_Node_VlanCos

    // IP Priority Mask.
    IpPriorityMask HardwareModuleEfd_Nodes_Node_IpPriorityMask

    // MPLS Priority Mask.
    MplsPriorityMask HardwareModuleEfd_Nodes_Node_MplsPriorityMask

    // EFD MPLS parameters.
    MplsExp HardwareModuleEfd_Nodes_Node_MplsExp
}

func (node *HardwareModuleEfd_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("vlan-priority-mask", types.YChild{"VlanPriorityMask", &node.VlanPriorityMask})
    node.EntityData.Children.Append("ip-precedence", types.YChild{"IpPrecedence", &node.IpPrecedence})
    node.EntityData.Children.Append("vlan-cos", types.YChild{"VlanCos", &node.VlanCos})
    node.EntityData.Children.Append("ip-priority-mask", types.YChild{"IpPriorityMask", &node.IpPriorityMask})
    node.EntityData.Children.Append("mpls-priority-mask", types.YChild{"MplsPriorityMask", &node.MplsPriorityMask})
    node.EntityData.Children.Append("mpls-exp", types.YChild{"MplsExp", &node.MplsExp})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", node.Enable})
    node.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", node.Mode})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// HardwareModuleEfd_Nodes_Node_VlanPriorityMask
// VLAN Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_VlanPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (vlanPriorityMask *HardwareModuleEfd_Nodes_Node_VlanPriorityMask) GetEntityData() *types.CommonEntityData {
    vlanPriorityMask.EntityData.YFilter = vlanPriorityMask.YFilter
    vlanPriorityMask.EntityData.YangName = "vlan-priority-mask"
    vlanPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    vlanPriorityMask.EntityData.ParentYangName = "node"
    vlanPriorityMask.EntityData.SegmentPath = "vlan-priority-mask"
    vlanPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + vlanPriorityMask.EntityData.SegmentPath
    vlanPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanPriorityMask.EntityData.Children = types.NewOrderedMap()
    vlanPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    vlanPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", vlanPriorityMask.Prec0})
    vlanPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", vlanPriorityMask.Prec1})
    vlanPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", vlanPriorityMask.Prec2})
    vlanPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", vlanPriorityMask.Prec3})
    vlanPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", vlanPriorityMask.Prec4})
    vlanPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", vlanPriorityMask.Prec5})
    vlanPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", vlanPriorityMask.Prec6})
    vlanPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", vlanPriorityMask.Prec7})

    vlanPriorityMask.EntityData.YListKeys = []string {}

    return &(vlanPriorityMask.EntityData)
}

// HardwareModuleEfd_Nodes_Node_IpPrecedence
// EFD IP parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_IpPrecedence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP TOS precedence threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Precedence interface{}

    // IP operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (ipPrecedence *HardwareModuleEfd_Nodes_Node_IpPrecedence) GetEntityData() *types.CommonEntityData {
    ipPrecedence.EntityData.YFilter = ipPrecedence.YFilter
    ipPrecedence.EntityData.YangName = "ip-precedence"
    ipPrecedence.EntityData.BundleName = "cisco_ios_xr"
    ipPrecedence.EntityData.ParentYangName = "node"
    ipPrecedence.EntityData.SegmentPath = "ip-precedence"
    ipPrecedence.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + ipPrecedence.EntityData.SegmentPath
    ipPrecedence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipPrecedence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipPrecedence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipPrecedence.EntityData.Children = types.NewOrderedMap()
    ipPrecedence.EntityData.Leafs = types.NewOrderedMap()
    ipPrecedence.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", ipPrecedence.Precedence})
    ipPrecedence.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", ipPrecedence.Operation})

    ipPrecedence.EntityData.YListKeys = []string {}

    return &(ipPrecedence.EntityData)
}

// HardwareModuleEfd_Nodes_Node_VlanCos
// EFD VLAN parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_VlanCos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // VLAN COS threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Cos interface{}

    // VLAN operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (vlanCos *HardwareModuleEfd_Nodes_Node_VlanCos) GetEntityData() *types.CommonEntityData {
    vlanCos.EntityData.YFilter = vlanCos.YFilter
    vlanCos.EntityData.YangName = "vlan-cos"
    vlanCos.EntityData.BundleName = "cisco_ios_xr"
    vlanCos.EntityData.ParentYangName = "node"
    vlanCos.EntityData.SegmentPath = "vlan-cos"
    vlanCos.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + vlanCos.EntityData.SegmentPath
    vlanCos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanCos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanCos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanCos.EntityData.Children = types.NewOrderedMap()
    vlanCos.EntityData.Leafs = types.NewOrderedMap()
    vlanCos.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", vlanCos.Cos})
    vlanCos.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", vlanCos.Operation})

    vlanCos.EntityData.YListKeys = []string {}

    return &(vlanCos.EntityData)
}

// HardwareModuleEfd_Nodes_Node_IpPriorityMask
// IP Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_IpPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (ipPriorityMask *HardwareModuleEfd_Nodes_Node_IpPriorityMask) GetEntityData() *types.CommonEntityData {
    ipPriorityMask.EntityData.YFilter = ipPriorityMask.YFilter
    ipPriorityMask.EntityData.YangName = "ip-priority-mask"
    ipPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    ipPriorityMask.EntityData.ParentYangName = "node"
    ipPriorityMask.EntityData.SegmentPath = "ip-priority-mask"
    ipPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + ipPriorityMask.EntityData.SegmentPath
    ipPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipPriorityMask.EntityData.Children = types.NewOrderedMap()
    ipPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    ipPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", ipPriorityMask.Prec0})
    ipPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", ipPriorityMask.Prec1})
    ipPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", ipPriorityMask.Prec2})
    ipPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", ipPriorityMask.Prec3})
    ipPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", ipPriorityMask.Prec4})
    ipPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", ipPriorityMask.Prec5})
    ipPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", ipPriorityMask.Prec6})
    ipPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", ipPriorityMask.Prec7})

    ipPriorityMask.EntityData.YListKeys = []string {}

    return &(ipPriorityMask.EntityData)
}

// HardwareModuleEfd_Nodes_Node_MplsPriorityMask
// MPLS Priority Mask
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_MplsPriorityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec0 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec1 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec2 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec3 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec4 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec5 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec6 interface{}

    // Prec. The type is interface{} with range: 0..1. The default value is 0.
    Prec7 interface{}
}

func (mplsPriorityMask *HardwareModuleEfd_Nodes_Node_MplsPriorityMask) GetEntityData() *types.CommonEntityData {
    mplsPriorityMask.EntityData.YFilter = mplsPriorityMask.YFilter
    mplsPriorityMask.EntityData.YangName = "mpls-priority-mask"
    mplsPriorityMask.EntityData.BundleName = "cisco_ios_xr"
    mplsPriorityMask.EntityData.ParentYangName = "node"
    mplsPriorityMask.EntityData.SegmentPath = "mpls-priority-mask"
    mplsPriorityMask.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + mplsPriorityMask.EntityData.SegmentPath
    mplsPriorityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsPriorityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsPriorityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsPriorityMask.EntityData.Children = types.NewOrderedMap()
    mplsPriorityMask.EntityData.Leafs = types.NewOrderedMap()
    mplsPriorityMask.EntityData.Leafs.Append("prec0", types.YLeaf{"Prec0", mplsPriorityMask.Prec0})
    mplsPriorityMask.EntityData.Leafs.Append("prec1", types.YLeaf{"Prec1", mplsPriorityMask.Prec1})
    mplsPriorityMask.EntityData.Leafs.Append("prec2", types.YLeaf{"Prec2", mplsPriorityMask.Prec2})
    mplsPriorityMask.EntityData.Leafs.Append("prec3", types.YLeaf{"Prec3", mplsPriorityMask.Prec3})
    mplsPriorityMask.EntityData.Leafs.Append("prec4", types.YLeaf{"Prec4", mplsPriorityMask.Prec4})
    mplsPriorityMask.EntityData.Leafs.Append("prec5", types.YLeaf{"Prec5", mplsPriorityMask.Prec5})
    mplsPriorityMask.EntityData.Leafs.Append("prec6", types.YLeaf{"Prec6", mplsPriorityMask.Prec6})
    mplsPriorityMask.EntityData.Leafs.Append("prec7", types.YLeaf{"Prec7", mplsPriorityMask.Prec7})

    mplsPriorityMask.EntityData.YListKeys = []string {}

    return &(mplsPriorityMask.EntityData)
}

// HardwareModuleEfd_Nodes_Node_MplsExp
// EFD MPLS parameters
// This type is a presence type.
type HardwareModuleEfd_Nodes_Node_MplsExp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // MPLS EXP threshold. The type is interface{} with range: 0..7. This
    // attribute is mandatory.
    Exp interface{}

    // MPLS operation. The type is Asr9kEfdOperation. The default value is
    // greater-than-or-equal.
    Operation interface{}
}

func (mplsExp *HardwareModuleEfd_Nodes_Node_MplsExp) GetEntityData() *types.CommonEntityData {
    mplsExp.EntityData.YFilter = mplsExp.YFilter
    mplsExp.EntityData.YangName = "mpls-exp"
    mplsExp.EntityData.BundleName = "cisco_ios_xr"
    mplsExp.EntityData.ParentYangName = "node"
    mplsExp.EntityData.SegmentPath = "mpls-exp"
    mplsExp.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd/nodes/node/" + mplsExp.EntityData.SegmentPath
    mplsExp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsExp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsExp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsExp.EntityData.Children = types.NewOrderedMap()
    mplsExp.EntityData.Leafs = types.NewOrderedMap()
    mplsExp.EntityData.Leafs.Append("exp", types.YLeaf{"Exp", mplsExp.Exp})
    mplsExp.EntityData.Leafs.Append("operation", types.YLeaf{"Operation", mplsExp.Operation})

    mplsExp.EntityData.YListKeys = []string {}

    return &(mplsExp.EntityData)
}

// HardwareModuleAllQosMode
// hardware module all qos mode
type HardwareModuleAllQosMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enable per-priority-buffer-limit. The type is interface{}.
    PerPriorityBufferLimit interface{}

    // Enable ingress queue for MOD-80 4X10G MPA or MOD-400 20X10G MPA. The type
    // is interface{}.
    IngressQueue interface{}

    // Enable aggregate bundle mode . The type is interface{}.
    AggregateBundleMode interface{}

    // Enable bundle aggregate policer mode . The type is interface{}.
    BundleAggregatePolicerMode interface{}

    // Enable L4 wred accounting buffer mode. The type is interface{}.
    WredBufferMode interface{}
}

func (hardwareModuleAllQosMode *HardwareModuleAllQosMode) GetEntityData() *types.CommonEntityData {
    hardwareModuleAllQosMode.EntityData.YFilter = hardwareModuleAllQosMode.YFilter
    hardwareModuleAllQosMode.EntityData.YangName = "hardware-module-all-qos-mode"
    hardwareModuleAllQosMode.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleAllQosMode.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleAllQosMode.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-all-qos-mode"
    hardwareModuleAllQosMode.EntityData.AbsolutePath = hardwareModuleAllQosMode.EntityData.SegmentPath
    hardwareModuleAllQosMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleAllQosMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleAllQosMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleAllQosMode.EntityData.Children = types.NewOrderedMap()
    hardwareModuleAllQosMode.EntityData.Leafs = types.NewOrderedMap()
    hardwareModuleAllQosMode.EntityData.Leafs.Append("per-priority-buffer-limit", types.YLeaf{"PerPriorityBufferLimit", hardwareModuleAllQosMode.PerPriorityBufferLimit})
    hardwareModuleAllQosMode.EntityData.Leafs.Append("ingress-queue", types.YLeaf{"IngressQueue", hardwareModuleAllQosMode.IngressQueue})
    hardwareModuleAllQosMode.EntityData.Leafs.Append("aggregate-bundle-mode", types.YLeaf{"AggregateBundleMode", hardwareModuleAllQosMode.AggregateBundleMode})
    hardwareModuleAllQosMode.EntityData.Leafs.Append("bundle-aggregate-policer-mode", types.YLeaf{"BundleAggregatePolicerMode", hardwareModuleAllQosMode.BundleAggregatePolicerMode})
    hardwareModuleAllQosMode.EntityData.Leafs.Append("wred-buffer-mode", types.YLeaf{"WredBufferMode", hardwareModuleAllQosMode.WredBufferMode})

    hardwareModuleAllQosMode.EntityData.YListKeys = []string {}

    return &(hardwareModuleAllQosMode.EntityData)
}

