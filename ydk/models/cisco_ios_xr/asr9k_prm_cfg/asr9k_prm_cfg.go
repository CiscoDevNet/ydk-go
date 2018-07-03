// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-prm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-qos-mode: QoS mode in hardware module port(s)
//   hardware-module-tcp-mss-adjust: hardware module tcp mss adjust
//   hardware-module-load-balance: hardware module load balance
//   hardware-module-tcam: hardware module tcam
//   hardware-module-efd: hardware module efd
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcp-mss-adjust}", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcp-mss-adjust", reflect.TypeOf(HardwareModuleTcpMssAdjust{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-load-balance}", reflect.TypeOf(HardwareModuleLoadBalance{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-load-balance", reflect.TypeOf(HardwareModuleLoadBalance{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-tcam}", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-tcam", reflect.TypeOf(HardwareModuleTcam{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-prm-cfg hardware-module-efd}", reflect.TypeOf(HardwareModuleEfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-efd", reflect.TypeOf(HardwareModuleEfd{}))
}

// Asr9kEfdMode represents Asr9k efd mode
type Asr9kEfdMode string

const (
    // Only check outer encap
    Asr9kEfdMode_only_outer_encap Asr9kEfdMode = "only-outer-encap"

    // Check outer and inner encap
    Asr9kEfdMode_include_inner_encap Asr9kEfdMode = "include-inner-encap"
)

// Asr9kEfdOperation represents Asr9k efd operation
type Asr9kEfdOperation string

const (
    // Less than
    Asr9kEfdOperation_less_than Asr9kEfdOperation = "less-than"

    // Greater than or equal
    Asr9kEfdOperation_greater_than_or_equal Asr9kEfdOperation = "greater-than-or-equal"
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

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

// HardwareModuleLoadBalance
// hardware module load balance
type HardwareModuleLoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bundle load balance options.
    Bundle HardwareModuleLoadBalance_Bundle
}

func (hardwareModuleLoadBalance *HardwareModuleLoadBalance) GetEntityData() *types.CommonEntityData {
    hardwareModuleLoadBalance.EntityData.YFilter = hardwareModuleLoadBalance.YFilter
    hardwareModuleLoadBalance.EntityData.YangName = "hardware-module-load-balance"
    hardwareModuleLoadBalance.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleLoadBalance.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-prm-cfg"
    hardwareModuleLoadBalance.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-prm-cfg:hardware-module-load-balance"
    hardwareModuleLoadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleLoadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleLoadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleLoadBalance.EntityData.Children = types.NewOrderedMap()
    hardwareModuleLoadBalance.EntityData.Children.Append("bundle", types.YChild{"Bundle", &hardwareModuleLoadBalance.Bundle})
    hardwareModuleLoadBalance.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleLoadBalance.EntityData.YListKeys = []string {}

    return &(hardwareModuleLoadBalance.EntityData)
}

// HardwareModuleLoadBalance_Bundle
// Bundle load balance options
type HardwareModuleLoadBalance_Bundle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Load balance for L2 services.
    L2Service HardwareModuleLoadBalance_Bundle_L2Service
}

func (bundle *HardwareModuleLoadBalance_Bundle) GetEntityData() *types.CommonEntityData {
    bundle.EntityData.YFilter = bundle.YFilter
    bundle.EntityData.YangName = "bundle"
    bundle.EntityData.BundleName = "cisco_ios_xr"
    bundle.EntityData.ParentYangName = "hardware-module-load-balance"
    bundle.EntityData.SegmentPath = "bundle"
    bundle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundle.EntityData.Children = types.NewOrderedMap()
    bundle.EntityData.Children.Append("l2-service", types.YChild{"L2Service", &bundle.L2Service})
    bundle.EntityData.Leafs = types.NewOrderedMap()

    bundle.EntityData.YListKeys = []string {}

    return &(bundle.EntityData)
}

// HardwareModuleLoadBalance_Bundle_L2Service
// Load balance for L2 services
type HardwareModuleLoadBalance_Bundle_L2Service struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Load balance L2 services over bundle with L3 parameters. The type is
    // interface{}.
    L3Parameters interface{}
}

func (l2Service *HardwareModuleLoadBalance_Bundle_L2Service) GetEntityData() *types.CommonEntityData {
    l2Service.EntityData.YFilter = l2Service.YFilter
    l2Service.EntityData.YangName = "l2-service"
    l2Service.EntityData.BundleName = "cisco_ios_xr"
    l2Service.EntityData.ParentYangName = "bundle"
    l2Service.EntityData.SegmentPath = "l2-service"
    l2Service.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Service.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Service.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Service.EntityData.Children = types.NewOrderedMap()
    l2Service.EntityData.Leafs = types.NewOrderedMap()
    l2Service.EntityData.Leafs.Append("l3-parameters", types.YLeaf{"L3Parameters", l2Service.L3Parameters})

    l2Service.EntityData.YListKeys = []string {}

    return &(l2Service.EntityData)
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

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

    // This attribute is a key. Node Name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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

