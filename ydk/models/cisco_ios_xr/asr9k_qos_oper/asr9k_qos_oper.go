// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-qos package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-qos: QoS ASR9K platform operational data 
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_qos_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_qos_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-qos-oper platform-qos}", reflect.TypeOf(PlatformQos{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-qos-oper:platform-qos", reflect.TypeOf(PlatformQos{}))
}

// ShapeProfiletypeV2 represents SHAPE profile type
type ShapeProfiletypeV2 string

const (
    // Shape Profile Type Invalid
    ShapeProfiletypeV2_invalid ShapeProfiletypeV2 = "invalid"

    // Shape Profile Type Always
    ShapeProfiletypeV2_always ShapeProfiletypeV2 = "always"

    // Shape Profile Type Never
    ShapeProfiletypeV2_never ShapeProfiletypeV2 = "never"

    // Shape Profile Type Explicit
    ShapeProfiletypeV2_explicit ShapeProfiletypeV2 = "explicit"

    // Shape Profile Type Scale
    ShapeProfiletypeV2_scale ShapeProfiletypeV2 = "scale"

    // Shape Profile Type Grid
    ShapeProfiletypeV2_grid ShapeProfiletypeV2 = "grid"
)

// Action represents Action type
type Action string

const (
    // Police action transmit
    Action_police_transmit Action = "police-transmit"

    // Police action set transmit
    Action_police_set_transmit Action = "police-set-transmit"

    // Police action drop
    Action_police_drop Action = "police-drop"

    // Police action unknown
    Action_police_unknown Action = "police-unknown"
)

// ActionOpcode represents Action opcode
type ActionOpcode string

const (
    // Precedence
    ActionOpcode_precedence ActionOpcode = "precedence"

    // DSCP
    ActionOpcode_dscp ActionOpcode = "dscp"

    // Discard class
    ActionOpcode_discard_class ActionOpcode = "discard-class"

    // QoS group
    ActionOpcode_qos_group ActionOpcode = "qos-group"

    // COS inner
    ActionOpcode_cos_inner ActionOpcode = "cos-inner"

    // COS
    ActionOpcode_cos ActionOpcode = "cos"

    // EXP top
    ActionOpcode_exp_top ActionOpcode = "exp-top"

    // EXP IMP
    ActionOpcode_exp_imp ActionOpcode = "exp-imp"

    // Tunnel precedence
    ActionOpcode_tunnel_precedence ActionOpcode = "tunnel-precedence"

    // Tunnel DSCP
    ActionOpcode_tunnel_dscp ActionOpcode = "tunnel-dscp"

    // ITAG DEI
    ActionOpcode_itag_dei ActionOpcode = "itag-dei"

    // ITAG COS
    ActionOpcode_itag_cos ActionOpcode = "itag-cos"

    // COS imposition
    ActionOpcode_cos_imposition ActionOpcode = "cos-imposition"

    // DEI imposition
    ActionOpcode_dei_imposition ActionOpcode = "dei-imposition"

    // DEI
    ActionOpcode_dei ActionOpcode = "dei"

    // No marking
    ActionOpcode_no_marking ActionOpcode = "no-marking"
)

// Wred represents WRED type
type Wred string

const (
    // WRED based on COS
    Wred_wred_cos Wred = "wred-cos"

    // WRED based on DSCP
    Wred_wred_dscp Wred = "wred-dscp"

    // WRED based on Precedence
    Wred_wred_precedence Wred = "wred-precedence"

    // WRED based on discard class
    Wred_wred_discard_class Wred = "wred-discard-class"

    // WRED based on MPLS EXP
    Wred_wred_mpls_exp Wred = "wred-mpls-exp"

    // RED with user defined min and max
    Wred_red_with_user_min_max Wred = "red-with-user-min-max"

    // RED with default min and max
    Wred_red_with_default_min_max Wred = "red-with-default-min-max"

    // WRED DEI
    Wred_wred_dei Wred = "wred-dei"
)

// TbAlgorithm represents Tokenbucket type
type TbAlgorithm string

const (
    // Inactive, configured but disabled
    TbAlgorithm_inactive TbAlgorithm = "inactive"

    // Single token bucket
    TbAlgorithm_single TbAlgorithm = "single"

    // Single rate three color marker
    TbAlgorithm_single_rate_tcm TbAlgorithm = "single-rate-tcm"

    // Two rate three color marker
    TbAlgorithm_two_rate_tcm TbAlgorithm = "two-rate-tcm"

    // Allows coupling between CIR and PIR tb's
    TbAlgorithm_mef_tcm TbAlgorithm = "mef-tcm"

    // Internal dummy token bucket for coupled-policer
    // child
    TbAlgorithm_dummy TbAlgorithm = "dummy"
)

// Queue represents Queue type
type Queue string

const (
    // Port default queue
    Queue_port_default Queue = "port-default"

    // The default queue in this policy
    Queue_class_default Queue = "class-default"

    // Create the priority 1 queue in the level which
    // will be shared by other p1 classes
    Queue_priority1_queue Queue = "priority1-queue"

    // Create the priority 2 queue in the level, which
    // will be shared by other p2 classes
    Queue_priority2_queue Queue = "priority2-queue"

    // Create the priority 3 queue in the level, which
    // will be shared by other p3 classes
    Queue_priority3_queue Queue = "priority3-queue"

    // Create the priority 4 queue in the level, which
    // will be shared by other p4 classes
    Queue_priority4_queue Queue = "priority4-queue"

    // Create the priority 5 queue in the level, which
    // will be shared by other p5 classes
    Queue_priority5_queue Queue = "priority5-queue"

    // Create the priority 6 queue in the level, which
    // will be shared by other p6 classes
    Queue_priority6_queue Queue = "priority6-queue"

    // Create the priority 7 queue in the level, which
    // will be shared by other p7 classes
    Queue_priority7_queue Queue = "priority7-queue"

    // Current level's priority 1 queue
    Queue_first_p1_class_name Queue = "first-p1-class-name"

    // Current level's priority 2 queue
    Queue_first_p2_class_name Queue = "first-p2-class-name"

    // Current level's priority 3 queue
    Queue_first_p3_class_name Queue = "first-p3-class-name"

    // Current level's priority 4 queue
    Queue_first_p4_class_name Queue = "first-p4-class-name"

    // Current level's priority 5 queue
    Queue_first_p5_class_name Queue = "first-p5-class-name"

    // Current level's priority 6 queue
    Queue_first_p6_class_name Queue = "first-p6-class-name"

    // Current level's priority 7 queue
    Queue_first_p7_class_name Queue = "first-p7-class-name"

    // Port priority 1 queue
    Queue_port_priority1 Queue = "port-priority1"

    // Port priority 2 queue
    Queue_port_priority2 Queue = "port-priority2"

    // Port priority 3 queue
    Queue_port_priority3 Queue = "port-priority3"

    // Port priority 4 queue
    Queue_port_priority4 Queue = "port-priority4"

    // Port priority 5 queue
    Queue_port_priority5 Queue = "port-priority5"

    // Port priority 6 queue
    Queue_port_priority6 Queue = "port-priority6"

    // Port priority 7 queue
    Queue_port_priority7 Queue = "port-priority7"

    // Create a new queue for this class
    Queue_new Queue = "new"

    // Under parent queue
    Queue_parent_class Queue = "parent-class"

    // Priority 1
    Queue_priority1 Queue = "priority1"

    // Priority 2
    Queue_priority2 Queue = "priority2"

    // Priority 3
    Queue_priority3 Queue = "priority3"

    // Priority 4
    Queue_priority4 Queue = "priority4"

    // Priority 5
    Queue_priority5 Queue = "priority5"

    // Priority 6
    Queue_priority6 Queue = "priority6"

    // Priority 7
    Queue_priority7 Queue = "priority7"

    // Priority ignored level
    Queue_priority_ignored_normal Queue = "priority-ignored-normal"

    // Normal priority
    Queue_normal_priority Queue = "normal-priority"

    // Class unknown
    Queue_class_unknown Queue = "class-unknown"

    // Unknown priority
    Queue_unknown_priority Queue = "unknown-priority"
)

// QosUnit represents QoS parameter unit
type QosUnit string

const (
    // Invalid type
    QosUnit_invalid QosUnit = "invalid"

    // Bytes
    QosUnit_bytes QosUnit = "bytes"

    // Kilobytes
    QosUnit_kilobytes QosUnit = "kilobytes"

    // Megabytes
    QosUnit_megabytes QosUnit = "megabytes"

    // Gigabytes
    QosUnit_gigabytes QosUnit = "gigabytes"

    // Bits per second
    QosUnit_bps QosUnit = "bps"

    // Kilo bits per second
    QosUnit_kbps QosUnit = "kbps"

    // Mega bits per second
    QosUnit_mbps QosUnit = "mbps"

    // Giga bits per second
    QosUnit_gbps QosUnit = "gbps"

    // Cells per second
    QosUnit_cells_per_second QosUnit = "cells-per-second"

    // Packets per second
    QosUnit_packets_per_second QosUnit = "packets-per-second"

    // Microsecond
    QosUnit_microsecond QosUnit = "microsecond"

    // Millisecond
    QosUnit_millisecond QosUnit = "millisecond"

    // Number of packets
    QosUnit_packets QosUnit = "packets"

    // Number of cells
    QosUnit_cells QosUnit = "cells"

    // Percentage
    QosUnit_percentage QosUnit = "percentage"

    // Ratio
    QosUnit_ratio QosUnit = "ratio"
)

// Wred1 represents Wred1
type Wred1 string

const (
    // wred cos cmd
    Wred1_wred_cos_cmd Wred1 = "wred-cos-cmd"

    // wred dscp cmd
    Wred1_wred_dscp_cmd Wred1 = "wred-dscp-cmd"

    // wred precedence cmd
    Wred1_wred_precedence_cmd Wred1 = "wred-precedence-cmd"

    // wred discard class cmd
    Wred1_wred_discard_class_cmd Wred1 = "wred-discard-class-cmd"

    // wred mpls exp cmd
    Wred1_wred_mpls_exp_cmd Wred1 = "wred-mpls-exp-cmd"

    // red with user min max
    Wred1_red_with_user_min_max Wred1 = "red-with-user-min-max"

    // red with default min max
    Wred1_red_with_default_min_max Wred1 = "red-with-default-min-max"

    // wred dei cmd
    Wred1_wred_dei_cmd Wred1 = "wred-dei-cmd"

    // wred ecn cmd
    Wred1_wred_ecn_cmd Wred1 = "wred-ecn-cmd"

    // wred invalid cmd
    Wred1_wred_invalid_cmd Wred1 = "wred-invalid-cmd"
)

// PolicyParamUnit represents Policy param unit
type PolicyParamUnit string

const (
    // policy param unit invalid
    PolicyParamUnit_policy_param_unit_invalid PolicyParamUnit = "policy-param-unit-invalid"

    // policy param unit bytes
    PolicyParamUnit_policy_param_unit_bytes PolicyParamUnit = "policy-param-unit-bytes"

    // policy param unit kbytes
    PolicyParamUnit_policy_param_unit_kbytes PolicyParamUnit = "policy-param-unit-kbytes"

    // policy param unit mbytes
    PolicyParamUnit_policy_param_unit_mbytes PolicyParamUnit = "policy-param-unit-mbytes"

    // policy param unit gbytes
    PolicyParamUnit_policy_param_unit_gbytes PolicyParamUnit = "policy-param-unit-gbytes"

    // policy param unit bitsps
    PolicyParamUnit_policy_param_unit_bitsps PolicyParamUnit = "policy-param-unit-bitsps"

    // policy param unit kbitsps
    PolicyParamUnit_policy_param_unit_kbitsps PolicyParamUnit = "policy-param-unit-kbitsps"

    // policy param unit mbitsps
    PolicyParamUnit_policy_param_unit_mbitsps PolicyParamUnit = "policy-param-unit-mbitsps"

    // policy param unit gbitsps
    PolicyParamUnit_policy_param_unit_gbitsps PolicyParamUnit = "policy-param-unit-gbitsps"

    // policy param unit cells ps
    PolicyParamUnit_policy_param_unit_cells_ps PolicyParamUnit = "policy-param-unit-cells-ps"

    // policy param unit packets ps
    PolicyParamUnit_policy_param_unit_packets_ps PolicyParamUnit = "policy-param-unit-packets-ps"

    // policy param unit us
    PolicyParamUnit_policy_param_unit_us PolicyParamUnit = "policy-param-unit-us"

    // policy param unit ms
    PolicyParamUnit_policy_param_unit_ms PolicyParamUnit = "policy-param-unit-ms"

    // policy param unit seconds
    PolicyParamUnit_policy_param_unit_seconds PolicyParamUnit = "policy-param-unit-seconds"

    // policy param unit packets
    PolicyParamUnit_policy_param_unit_packets PolicyParamUnit = "policy-param-unit-packets"

    // policy param unit cells
    PolicyParamUnit_policy_param_unit_cells PolicyParamUnit = "policy-param-unit-cells"

    // policy param unit percent
    PolicyParamUnit_policy_param_unit_percent PolicyParamUnit = "policy-param-unit-percent"

    // policy param unit per thousand
    PolicyParamUnit_policy_param_unit_per_thousand PolicyParamUnit = "policy-param-unit-per-thousand"

    // policy param unit per million
    PolicyParamUnit_policy_param_unit_per_million PolicyParamUnit = "policy-param-unit-per-million"

    // policy param unit hz
    PolicyParamUnit_policy_param_unit_hz PolicyParamUnit = "policy-param-unit-hz"

    // policy param unit khz
    PolicyParamUnit_policy_param_unit_khz PolicyParamUnit = "policy-param-unit-khz"

    // policy param unit mhz
    PolicyParamUnit_policy_param_unit_mhz PolicyParamUnit = "policy-param-unit-mhz"

    // policy param unit ratio
    PolicyParamUnit_policy_param_unit_ratio PolicyParamUnit = "policy-param-unit-ratio"

    // policy param unit max
    PolicyParamUnit_policy_param_unit_max PolicyParamUnit = "policy-param-unit-max"
)

// CacState represents CAC/UBRL class states
type CacState string

const (
    // unknown
    CacState_unknown CacState = "unknown"

    // admit
    CacState_admit CacState = "admit"

    // redirect
    CacState_redirect CacState = "redirect"

    // ubrl
    CacState_ubrl CacState = "ubrl"
)

// PolicyState represents Different Interface states
type PolicyState string

const (
    // active
    PolicyState_active PolicyState = "active"

    // suspended
    PolicyState_suspended PolicyState = "suspended"
)

// PlatformQos
// QoS ASR9K platform operational data 
type PlatformQos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes with platform specific QoS configuration.
    Nodes PlatformQos_Nodes
}

func (platformQos *PlatformQos) GetEntityData() *types.CommonEntityData {
    platformQos.EntityData.YFilter = platformQos.YFilter
    platformQos.EntityData.YangName = "platform-qos"
    platformQos.EntityData.BundleName = "cisco_ios_xr"
    platformQos.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-qos-oper"
    platformQos.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-qos-oper:platform-qos"
    platformQos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformQos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformQos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformQos.EntityData.Children = make(map[string]types.YChild)
    platformQos.EntityData.Children["nodes"] = types.YChild{"Nodes", &platformQos.Nodes}
    platformQos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformQos.EntityData)
}

// PlatformQos_Nodes
// List of nodes with platform specific QoS
// configuration
type PlatformQos_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node with platform specific QoS configuration. The type is slice of
    // PlatformQos_Nodes_Node.
    Node []PlatformQos_Nodes_Node
}

func (nodes *PlatformQos_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "platform-qos"
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

// PlatformQos_Nodes_Node
// Node with platform specific QoS configuration
type PlatformQos_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // QoS system capability.
    Capability PlatformQos_Nodes_Node_Capability

    // QoS list of interfaces.
    Interfaces PlatformQos_Nodes_Node_Interfaces

    // QoS list of bundle interfaces.
    BundleInterfaces PlatformQos_Nodes_Node_BundleInterfaces
}

func (node *PlatformQos_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["capability"] = types.YChild{"Capability", &node.Capability}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["bundle-interfaces"] = types.YChild{"BundleInterfaces", &node.BundleInterfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PlatformQos_Nodes_Node_Capability
// QoS system capability
type PlatformQos_Nodes_Node_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum policy maps per system. The type is interface{} with range:
    // 0..4294967295.
    MaxPolicyMaps interface{}

    // Maximum policy hierarchy. The type is interface{} with range:
    // 0..4294967295.
    MaxPolicyHierarchy interface{}

    // Maximum policy name length. The type is interface{} with range:
    // 0..4294967295.
    MaxPolicyNameLength interface{}

    // Maximum classes per child policy. The type is interface{} with range:
    // 0..4294967295.
    MaxClassesPerChildPolicy interface{}

    // Maximum classes per policy. The type is interface{} with range:
    // 0..4294967295.
    MaxClassesPerPolicy interface{}

    // Maximum classes per parent policy. The type is interface{} with range:
    // 0..4294967295.
    MaxClassesPerGrandParentPolicy interface{}

    // Maximum police actions per class. The type is interface{} with range:
    // 0..4294967295.
    MaxPoliceActionsPerClass interface{}

    // Maximum marking action  per class. The type is interface{} with range:
    // 0..4294967295.
    MaxMarkingActionsPerClass interface{}

    // Maximum matches per class. The type is interface{} with range:
    // 0..4294967295.
    MaxMatchesPerClass interface{}

    // Maximum classmap name length. The type is interface{} with range:
    // 0..4294967295.
    MaxClassmapNameLength interface{}

    // Maximum bundle members. The type is interface{} with range: 0..4294967295.
    MaxBundleMembers interface{}

    // Maximum instance name length. The type is interface{} with range:
    // 0..4294967295.
    MaxInstanceNameLength interface{}
}

func (capability *PlatformQos_Nodes_Node_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "cisco_ios_xr"
    capability.EntityData.ParentYangName = "node"
    capability.EntityData.SegmentPath = "capability"
    capability.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capability.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capability.EntityData.Children = make(map[string]types.YChild)
    capability.EntityData.Leafs = make(map[string]types.YLeaf)
    capability.EntityData.Leafs["max-policy-maps"] = types.YLeaf{"MaxPolicyMaps", capability.MaxPolicyMaps}
    capability.EntityData.Leafs["max-policy-hierarchy"] = types.YLeaf{"MaxPolicyHierarchy", capability.MaxPolicyHierarchy}
    capability.EntityData.Leafs["max-policy-name-length"] = types.YLeaf{"MaxPolicyNameLength", capability.MaxPolicyNameLength}
    capability.EntityData.Leafs["max-classes-per-child-policy"] = types.YLeaf{"MaxClassesPerChildPolicy", capability.MaxClassesPerChildPolicy}
    capability.EntityData.Leafs["max-classes-per-policy"] = types.YLeaf{"MaxClassesPerPolicy", capability.MaxClassesPerPolicy}
    capability.EntityData.Leafs["max-classes-per-grand-parent-policy"] = types.YLeaf{"MaxClassesPerGrandParentPolicy", capability.MaxClassesPerGrandParentPolicy}
    capability.EntityData.Leafs["max-police-actions-per-class"] = types.YLeaf{"MaxPoliceActionsPerClass", capability.MaxPoliceActionsPerClass}
    capability.EntityData.Leafs["max-marking-actions-per-class"] = types.YLeaf{"MaxMarkingActionsPerClass", capability.MaxMarkingActionsPerClass}
    capability.EntityData.Leafs["max-matches-per-class"] = types.YLeaf{"MaxMatchesPerClass", capability.MaxMatchesPerClass}
    capability.EntityData.Leafs["max-classmap-name-length"] = types.YLeaf{"MaxClassmapNameLength", capability.MaxClassmapNameLength}
    capability.EntityData.Leafs["max-bundle-members"] = types.YLeaf{"MaxBundleMembers", capability.MaxBundleMembers}
    capability.EntityData.Leafs["max-instance-name-length"] = types.YLeaf{"MaxInstanceNameLength", capability.MaxInstanceNameLength}
    return &(capability.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces
// QoS list of interfaces
type PlatformQos_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_.
    Interface_ []PlatformQos_Nodes_Node_Interfaces_Interface
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
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

// PlatformQos_Nodes_Node_Interfaces_Interface
// QoS interface name
type PlatformQos_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS policy direction ingress.
    Input PlatformQos_Nodes_Node_Interfaces_Interface_Input

    // QoS policy direction egress.
    Output PlatformQos_Nodes_Node_Interfaces_Interface_Output
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["input"] = types.YChild{"Input", &self.Input}
    self.EntityData.Children["output"] = types.YChild{"Output", &self.Output}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input
// QoS policy direction ingress
type PlatformQos_Nodes_Node_Interfaces_Interface_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy direction egress.
    Details PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "interface"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["details"] = types.YChild{"Details", &input.Details}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details
// QoS policy direction egress
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header

    // Trident QoS policy details.
    Policy PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy

    // Typhoon QoS policy details.
    PolicyTyphoon PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon
}

func (details *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "input"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["policy"] = types.YChild{"Policy", &details.Policy}
    details.EntityData.Children["policy-typhoon"] = types.YChild{"PolicyTyphoon", &details.PolicyTyphoon}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header
// QoS policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Interface config and programmed parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters

    // Bandwidth that was programmed.
    ProgrammedBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_ProgrammedBandwidth
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &header.InterfaceParameters}
    header.EntityData.Children["programmed-bandwidth"] = types.YChild{"ProgrammedBandwidth", &header.ProgrammedBandwidth}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters
// Interface config and programmed parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth due to port speed change.
    PortConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortConfigBandwidth

    // Bandwidth obtain from IM.
    AncpConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpConfigBandwidth

    // ANCP bandwidth that was programmed.
    AncpProgrammedBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpProgrammedBandwidth

    // Bandwidth that was programmed.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "header"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["port-config-bandwidth"] = types.YChild{"PortConfigBandwidth", &interfaceParameters.PortConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-config-bandwidth"] = types.YChild{"AncpConfigBandwidth", &interfaceParameters.AncpConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-programmed-bandwidth"] = types.YChild{"AncpProgrammedBandwidth", &interfaceParameters.AncpProgrammedBandwidth}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortConfigBandwidth
// Bandwidth due to port speed change
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portConfigBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortConfigBandwidth) GetEntityData() *types.CommonEntityData {
    portConfigBandwidth.EntityData.YFilter = portConfigBandwidth.YFilter
    portConfigBandwidth.EntityData.YangName = "port-config-bandwidth"
    portConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    portConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    portConfigBandwidth.EntityData.SegmentPath = "port-config-bandwidth"
    portConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    portConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    portConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", portConfigBandwidth.Value}
    portConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portConfigBandwidth.Unit}
    return &(portConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpConfigBandwidth
// Bandwidth obtain from IM
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpConfigBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpConfigBandwidth) GetEntityData() *types.CommonEntityData {
    ancpConfigBandwidth.EntityData.YFilter = ancpConfigBandwidth.YFilter
    ancpConfigBandwidth.EntityData.YangName = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpConfigBandwidth.EntityData.SegmentPath = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpConfigBandwidth.Value}
    ancpConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpConfigBandwidth.Unit}
    return &(ancpConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpProgrammedBandwidth
// ANCP bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpProgrammedBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_AncpProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    ancpProgrammedBandwidth.EntityData.YFilter = ancpProgrammedBandwidth.YFilter
    ancpProgrammedBandwidth.EntityData.YangName = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpProgrammedBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpProgrammedBandwidth.EntityData.SegmentPath = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpProgrammedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpProgrammedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpProgrammedBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpProgrammedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpProgrammedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpProgrammedBandwidth.Value}
    ancpProgrammedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpProgrammedBandwidth.Unit}
    return &(ancpProgrammedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortShaperRate
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
    portShaperRate.EntityData.YFilter = portShaperRate.YFilter
    portShaperRate.EntityData.YangName = "port-shaper-rate"
    portShaperRate.EntityData.BundleName = "cisco_ios_xr"
    portShaperRate.EntityData.ParentYangName = "interface-parameters"
    portShaperRate.EntityData.SegmentPath = "port-shaper-rate"
    portShaperRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portShaperRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portShaperRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portShaperRate.EntityData.Children = make(map[string]types.YChild)
    portShaperRate.EntityData.Leafs = make(map[string]types.YLeaf)
    portShaperRate.EntityData.Leafs["value"] = types.YLeaf{"Value", portShaperRate.Value}
    portShaperRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portShaperRate.Unit}
    return &(portShaperRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_ProgrammedBandwidth
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_ProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (programmedBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Header_ProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    programmedBandwidth.EntityData.YFilter = programmedBandwidth.YFilter
    programmedBandwidth.EntityData.YangName = "programmed-bandwidth"
    programmedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    programmedBandwidth.EntityData.ParentYangName = "header"
    programmedBandwidth.EntityData.SegmentPath = "programmed-bandwidth"
    programmedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedBandwidth.EntityData.Children = make(map[string]types.YChild)
    programmedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", programmedBandwidth.Value}
    programmedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", programmedBandwidth.Unit}
    return &(programmedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy
// Trident QoS policy details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v1. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1.
    QosShowEaStV1 []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1
}

func (policy *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "details"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Children["qos-show-ea-st-v1"] = types.YChild{"QosShowEaStV1", nil}
    for i := range policy.QosShowEaStV1 {
        policy.EntityData.Children[types.GetSegmentPath(&policy.QosShowEaStV1[i])] = types.YChild{"QosShowEaStV1", &policy.QosShowEaStV1[i]}
    }
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policy.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1
// qos show ea st v1
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark
}

func (qosShowEaStV1 *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1) GetEntityData() *types.CommonEntityData {
    qosShowEaStV1.EntityData.YFilter = qosShowEaStV1.YFilter
    qosShowEaStV1.EntityData.YangName = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV1.EntityData.ParentYangName = "policy"
    qosShowEaStV1.EntityData.SegmentPath = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV1.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV1.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV1.Queue}
    qosShowEaStV1.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV1.QueueLimitParameters}
    qosShowEaStV1.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV1.Shape}
    qosShowEaStV1.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV1.Police}
    qosShowEaStV1.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV1.Wfq}
    qosShowEaStV1.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV1.Wred}
    qosShowEaStV1.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV1.Mark}
    qosShowEaStV1.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV1.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV1.ClassLevel}
    qosShowEaStV1.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV1.ClassName}
    qosShowEaStV1.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV1.PolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV1.ParentPolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV1.ParentClassName}
    return &(qosShowEaStV1.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Queue
// Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", queueLimitParameters.ProfileId}
    queueLimitParameters.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", queueLimitParameters.ScalingProfileId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape
// Shape parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cbs

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v1"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir"] = types.YChild{"Cir", &shape.Cir}
    shape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &shape.ConfigBandwidth}
    shape.EntityData.Children["cbs"] = types.YChild{"Cbs", &shape.Cbs}
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", shape.ProfileId}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Shape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police
// Police parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v1"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pir
// PIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Table ID. The type is interface{} with range: 0..255.
    TableId interface{}

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    wred.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", wred.TableId}
    wred.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wred.ProfileId}
    wred.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", wred.ScalingProfileId}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED precedence match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark
// Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v1"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon
// Typhoon QoS policy details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v2. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2.
    QosShowEaStV2 []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2
}

func (policyTyphoon *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon) GetEntityData() *types.CommonEntityData {
    policyTyphoon.EntityData.YFilter = policyTyphoon.YFilter
    policyTyphoon.EntityData.YangName = "policy-typhoon"
    policyTyphoon.EntityData.BundleName = "cisco_ios_xr"
    policyTyphoon.EntityData.ParentYangName = "details"
    policyTyphoon.EntityData.SegmentPath = "policy-typhoon"
    policyTyphoon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyTyphoon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyTyphoon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyTyphoon.EntityData.Children = make(map[string]types.YChild)
    policyTyphoon.EntityData.Children["qos-show-ea-st-v2"] = types.YChild{"QosShowEaStV2", nil}
    for i := range policyTyphoon.QosShowEaStV2 {
        policyTyphoon.EntityData.Children[types.GetSegmentPath(&policyTyphoon.QosShowEaStV2[i])] = types.YChild{"QosShowEaStV2", &policyTyphoon.QosShowEaStV2[i]}
    }
    policyTyphoon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyTyphoon.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2
// qos show ea st v2
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark
}

func (qosShowEaStV2 *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2) GetEntityData() *types.CommonEntityData {
    qosShowEaStV2.EntityData.YFilter = qosShowEaStV2.YFilter
    qosShowEaStV2.EntityData.YangName = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV2.EntityData.ParentYangName = "policy-typhoon"
    qosShowEaStV2.EntityData.SegmentPath = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV2.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV2.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV2.Queue}
    qosShowEaStV2.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV2.QueueLimitParameters}
    qosShowEaStV2.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV2.Shape}
    qosShowEaStV2.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV2.Police}
    qosShowEaStV2.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV2.Wfq}
    qosShowEaStV2.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV2.Wred}
    qosShowEaStV2.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV2.Mark}
    qosShowEaStV2.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV2.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV2.ClassLevel}
    qosShowEaStV2.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV2.ClassName}
    qosShowEaStV2.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV2.PolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV2.ParentPolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV2.ParentClassName}
    return &(qosShowEaStV2.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Queue
// Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", queueLimitParameters.AbsoluteIndex}
    queueLimitParameters.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", queueLimitParameters.TemplateId}
    queueLimitParameters.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", queueLimitParameters.CurveId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape
// Shape parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CIR Shaper type. The type is ShapeProfiletypeV2.
    CirShapeType interface{}

    // PIR Shaper type. The type is ShapeProfiletypeV2.
    PirShapeType interface{}

    // CIR shaper params.
    CirShape PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape

    // PIR shaper params.
    PirShape PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v2"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir-shape"] = types.YChild{"CirShape", &shape.CirShape}
    shape.EntityData.Children["pir-shape"] = types.YChild{"PirShape", &shape.PirShape}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["cir-shape-type"] = types.YLeaf{"CirShapeType", shape.CirShapeType}
    shape.EntityData.Leafs["pir-shape-type"] = types.YLeaf{"PirShapeType", shape.PirShapeType}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape
// CIR shaper params
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
}

func (cirShape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape) GetEntityData() *types.CommonEntityData {
    cirShape.EntityData.YFilter = cirShape.YFilter
    cirShape.EntityData.YangName = "cir-shape"
    cirShape.EntityData.BundleName = "cisco_ios_xr"
    cirShape.EntityData.ParentYangName = "shape"
    cirShape.EntityData.SegmentPath = "cir-shape"
    cirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cirShape.EntityData.Children = make(map[string]types.YChild)
    cirShape.EntityData.Children["cir"] = types.YChild{"Cir", &cirShape.Cir}
    cirShape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &cirShape.ConfigBandwidth}
    cirShape.EntityData.Children["cbs"] = types.YChild{"Cbs", &cirShape.Cbs}
    cirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    cirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", cirShape.ChunkId}
    cirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", cirShape.ProfileId}
    cirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", cirShape.ScaleFactor}
    return &(cirShape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "cir-shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "cir-shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "cir-shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
// PIR shaper params
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
}

func (pirShape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape) GetEntityData() *types.CommonEntityData {
    pirShape.EntityData.YFilter = pirShape.YFilter
    pirShape.EntityData.YangName = "pir-shape"
    pirShape.EntityData.BundleName = "cisco_ios_xr"
    pirShape.EntityData.ParentYangName = "shape"
    pirShape.EntityData.SegmentPath = "pir-shape"
    pirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pirShape.EntityData.Children = make(map[string]types.YChild)
    pirShape.EntityData.Children["pir"] = types.YChild{"Pir", &pirShape.Pir}
    pirShape.EntityData.Children["pbs"] = types.YChild{"Pbs", &pirShape.Pbs}
    pirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    pirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", pirShape.ChunkId}
    pirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pirShape.ProfileId}
    pirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", pirShape.ScaleFactor}
    return &(pirShape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "pir-shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "pir-shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police
// Police parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v2"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir
// PIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // WRED match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // WRED exp match if EXP start value equals to end value Format: <start-value>
    // , else range Format: <start-value> <end-value>. The type is string.
    ExpMatch interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", curve.AbsoluteIndex}
    curve.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", curve.TemplateId}
    curve.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", curve.CurveId}
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    curve.EntityData.Leafs["exp-match"] = types.YLeaf{"ExpMatch", curve.ExpMatch}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark
// Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v2"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output
// QoS policy direction egress
type PlatformQos_Nodes_Node_Interfaces_Interface_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy direction egress.
    Details PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "interface"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Children["details"] = types.YChild{"Details", &output.Details}
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(output.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details
// QoS policy direction egress
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header

    // Trident QoS policy details.
    Policy PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy

    // Typhoon QoS policy details.
    PolicyTyphoon PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon
}

func (details *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "output"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["policy"] = types.YChild{"Policy", &details.Policy}
    details.EntityData.Children["policy-typhoon"] = types.YChild{"PolicyTyphoon", &details.PolicyTyphoon}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header
// QoS policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Interface config and programmed parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters

    // Bandwidth that was programmed.
    ProgrammedBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_ProgrammedBandwidth
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &header.InterfaceParameters}
    header.EntityData.Children["programmed-bandwidth"] = types.YChild{"ProgrammedBandwidth", &header.ProgrammedBandwidth}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters
// Interface config and programmed parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth due to port speed change.
    PortConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortConfigBandwidth

    // Bandwidth obtain from IM.
    AncpConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpConfigBandwidth

    // ANCP bandwidth that was programmed.
    AncpProgrammedBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpProgrammedBandwidth

    // Bandwidth that was programmed.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "header"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["port-config-bandwidth"] = types.YChild{"PortConfigBandwidth", &interfaceParameters.PortConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-config-bandwidth"] = types.YChild{"AncpConfigBandwidth", &interfaceParameters.AncpConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-programmed-bandwidth"] = types.YChild{"AncpProgrammedBandwidth", &interfaceParameters.AncpProgrammedBandwidth}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortConfigBandwidth
// Bandwidth due to port speed change
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portConfigBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortConfigBandwidth) GetEntityData() *types.CommonEntityData {
    portConfigBandwidth.EntityData.YFilter = portConfigBandwidth.YFilter
    portConfigBandwidth.EntityData.YangName = "port-config-bandwidth"
    portConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    portConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    portConfigBandwidth.EntityData.SegmentPath = "port-config-bandwidth"
    portConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    portConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    portConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", portConfigBandwidth.Value}
    portConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portConfigBandwidth.Unit}
    return &(portConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpConfigBandwidth
// Bandwidth obtain from IM
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpConfigBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpConfigBandwidth) GetEntityData() *types.CommonEntityData {
    ancpConfigBandwidth.EntityData.YFilter = ancpConfigBandwidth.YFilter
    ancpConfigBandwidth.EntityData.YangName = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpConfigBandwidth.EntityData.SegmentPath = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpConfigBandwidth.Value}
    ancpConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpConfigBandwidth.Unit}
    return &(ancpConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpProgrammedBandwidth
// ANCP bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpProgrammedBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_AncpProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    ancpProgrammedBandwidth.EntityData.YFilter = ancpProgrammedBandwidth.YFilter
    ancpProgrammedBandwidth.EntityData.YangName = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpProgrammedBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpProgrammedBandwidth.EntityData.SegmentPath = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpProgrammedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpProgrammedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpProgrammedBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpProgrammedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpProgrammedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpProgrammedBandwidth.Value}
    ancpProgrammedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpProgrammedBandwidth.Unit}
    return &(ancpProgrammedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortShaperRate
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
    portShaperRate.EntityData.YFilter = portShaperRate.YFilter
    portShaperRate.EntityData.YangName = "port-shaper-rate"
    portShaperRate.EntityData.BundleName = "cisco_ios_xr"
    portShaperRate.EntityData.ParentYangName = "interface-parameters"
    portShaperRate.EntityData.SegmentPath = "port-shaper-rate"
    portShaperRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portShaperRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portShaperRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portShaperRate.EntityData.Children = make(map[string]types.YChild)
    portShaperRate.EntityData.Leafs = make(map[string]types.YLeaf)
    portShaperRate.EntityData.Leafs["value"] = types.YLeaf{"Value", portShaperRate.Value}
    portShaperRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portShaperRate.Unit}
    return &(portShaperRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_ProgrammedBandwidth
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_ProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (programmedBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Header_ProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    programmedBandwidth.EntityData.YFilter = programmedBandwidth.YFilter
    programmedBandwidth.EntityData.YangName = "programmed-bandwidth"
    programmedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    programmedBandwidth.EntityData.ParentYangName = "header"
    programmedBandwidth.EntityData.SegmentPath = "programmed-bandwidth"
    programmedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedBandwidth.EntityData.Children = make(map[string]types.YChild)
    programmedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", programmedBandwidth.Value}
    programmedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", programmedBandwidth.Unit}
    return &(programmedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy
// Trident QoS policy details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v1. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1.
    QosShowEaStV1 []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1
}

func (policy *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "details"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Children["qos-show-ea-st-v1"] = types.YChild{"QosShowEaStV1", nil}
    for i := range policy.QosShowEaStV1 {
        policy.EntityData.Children[types.GetSegmentPath(&policy.QosShowEaStV1[i])] = types.YChild{"QosShowEaStV1", &policy.QosShowEaStV1[i]}
    }
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policy.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1
// qos show ea st v1
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark
}

func (qosShowEaStV1 *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1) GetEntityData() *types.CommonEntityData {
    qosShowEaStV1.EntityData.YFilter = qosShowEaStV1.YFilter
    qosShowEaStV1.EntityData.YangName = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV1.EntityData.ParentYangName = "policy"
    qosShowEaStV1.EntityData.SegmentPath = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV1.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV1.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV1.Queue}
    qosShowEaStV1.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV1.QueueLimitParameters}
    qosShowEaStV1.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV1.Shape}
    qosShowEaStV1.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV1.Police}
    qosShowEaStV1.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV1.Wfq}
    qosShowEaStV1.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV1.Wred}
    qosShowEaStV1.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV1.Mark}
    qosShowEaStV1.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV1.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV1.ClassLevel}
    qosShowEaStV1.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV1.ClassName}
    qosShowEaStV1.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV1.PolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV1.ParentPolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV1.ParentClassName}
    return &(qosShowEaStV1.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Queue
// Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", queueLimitParameters.ProfileId}
    queueLimitParameters.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", queueLimitParameters.ScalingProfileId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape
// Shape parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cbs

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v1"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir"] = types.YChild{"Cir", &shape.Cir}
    shape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &shape.ConfigBandwidth}
    shape.EntityData.Children["cbs"] = types.YChild{"Cbs", &shape.Cbs}
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", shape.ProfileId}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Shape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police
// Police parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v1"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pir
// PIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Table ID. The type is interface{} with range: 0..255.
    TableId interface{}

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    wred.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", wred.TableId}
    wred.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wred.ProfileId}
    wred.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", wred.ScalingProfileId}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED precedence match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark
// Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v1"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon
// Typhoon QoS policy details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v2. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2.
    QosShowEaStV2 []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2
}

func (policyTyphoon *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon) GetEntityData() *types.CommonEntityData {
    policyTyphoon.EntityData.YFilter = policyTyphoon.YFilter
    policyTyphoon.EntityData.YangName = "policy-typhoon"
    policyTyphoon.EntityData.BundleName = "cisco_ios_xr"
    policyTyphoon.EntityData.ParentYangName = "details"
    policyTyphoon.EntityData.SegmentPath = "policy-typhoon"
    policyTyphoon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyTyphoon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyTyphoon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyTyphoon.EntityData.Children = make(map[string]types.YChild)
    policyTyphoon.EntityData.Children["qos-show-ea-st-v2"] = types.YChild{"QosShowEaStV2", nil}
    for i := range policyTyphoon.QosShowEaStV2 {
        policyTyphoon.EntityData.Children[types.GetSegmentPath(&policyTyphoon.QosShowEaStV2[i])] = types.YChild{"QosShowEaStV2", &policyTyphoon.QosShowEaStV2[i]}
    }
    policyTyphoon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyTyphoon.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2
// qos show ea st v2
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark
}

func (qosShowEaStV2 *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2) GetEntityData() *types.CommonEntityData {
    qosShowEaStV2.EntityData.YFilter = qosShowEaStV2.YFilter
    qosShowEaStV2.EntityData.YangName = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV2.EntityData.ParentYangName = "policy-typhoon"
    qosShowEaStV2.EntityData.SegmentPath = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV2.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV2.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV2.Queue}
    qosShowEaStV2.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV2.QueueLimitParameters}
    qosShowEaStV2.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV2.Shape}
    qosShowEaStV2.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV2.Police}
    qosShowEaStV2.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV2.Wfq}
    qosShowEaStV2.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV2.Wred}
    qosShowEaStV2.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV2.Mark}
    qosShowEaStV2.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV2.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV2.ClassLevel}
    qosShowEaStV2.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV2.ClassName}
    qosShowEaStV2.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV2.PolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV2.ParentPolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV2.ParentClassName}
    return &(qosShowEaStV2.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Queue
// Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", queueLimitParameters.AbsoluteIndex}
    queueLimitParameters.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", queueLimitParameters.TemplateId}
    queueLimitParameters.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", queueLimitParameters.CurveId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape
// Shape parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CIR Shaper type. The type is ShapeProfiletypeV2.
    CirShapeType interface{}

    // PIR Shaper type. The type is ShapeProfiletypeV2.
    PirShapeType interface{}

    // CIR shaper params.
    CirShape PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape

    // PIR shaper params.
    PirShape PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v2"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir-shape"] = types.YChild{"CirShape", &shape.CirShape}
    shape.EntityData.Children["pir-shape"] = types.YChild{"PirShape", &shape.PirShape}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["cir-shape-type"] = types.YLeaf{"CirShapeType", shape.CirShapeType}
    shape.EntityData.Leafs["pir-shape-type"] = types.YLeaf{"PirShapeType", shape.PirShapeType}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape
// CIR shaper params
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
}

func (cirShape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape) GetEntityData() *types.CommonEntityData {
    cirShape.EntityData.YFilter = cirShape.YFilter
    cirShape.EntityData.YangName = "cir-shape"
    cirShape.EntityData.BundleName = "cisco_ios_xr"
    cirShape.EntityData.ParentYangName = "shape"
    cirShape.EntityData.SegmentPath = "cir-shape"
    cirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cirShape.EntityData.Children = make(map[string]types.YChild)
    cirShape.EntityData.Children["cir"] = types.YChild{"Cir", &cirShape.Cir}
    cirShape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &cirShape.ConfigBandwidth}
    cirShape.EntityData.Children["cbs"] = types.YChild{"Cbs", &cirShape.Cbs}
    cirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    cirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", cirShape.ChunkId}
    cirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", cirShape.ProfileId}
    cirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", cirShape.ScaleFactor}
    return &(cirShape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "cir-shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "cir-shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "cir-shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
// PIR shaper params
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
}

func (pirShape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape) GetEntityData() *types.CommonEntityData {
    pirShape.EntityData.YFilter = pirShape.YFilter
    pirShape.EntityData.YangName = "pir-shape"
    pirShape.EntityData.BundleName = "cisco_ios_xr"
    pirShape.EntityData.ParentYangName = "shape"
    pirShape.EntityData.SegmentPath = "pir-shape"
    pirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pirShape.EntityData.Children = make(map[string]types.YChild)
    pirShape.EntityData.Children["pir"] = types.YChild{"Pir", &pirShape.Pir}
    pirShape.EntityData.Children["pbs"] = types.YChild{"Pbs", &pirShape.Pbs}
    pirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    pirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", pirShape.ChunkId}
    pirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pirShape.ProfileId}
    pirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", pirShape.ScaleFactor}
    return &(pirShape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "pir-shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "pir-shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police
// Police parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v2"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir
// PIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // WRED match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // WRED exp match if EXP start value equals to end value Format: <start-value>
    // , else range Format: <start-value> <end-value>. The type is string.
    ExpMatch interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", curve.AbsoluteIndex}
    curve.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", curve.TemplateId}
    curve.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", curve.CurveId}
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    curve.EntityData.Leafs["exp-match"] = types.YLeaf{"ExpMatch", curve.ExpMatch}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark
// Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v2"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces
// QoS list of bundle interfaces
type PlatformQos_Nodes_Node_BundleInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetEntityData() *types.CommonEntityData {
    bundleInterfaces.EntityData.YFilter = bundleInterfaces.YFilter
    bundleInterfaces.EntityData.YangName = "bundle-interfaces"
    bundleInterfaces.EntityData.BundleName = "cisco_ios_xr"
    bundleInterfaces.EntityData.ParentYangName = "node"
    bundleInterfaces.EntityData.SegmentPath = "bundle-interfaces"
    bundleInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterfaces.EntityData.Children = make(map[string]types.YChild)
    bundleInterfaces.EntityData.Children["bundle-interface"] = types.YChild{"BundleInterface", nil}
    for i := range bundleInterfaces.BundleInterface {
        bundleInterfaces.EntityData.Children[types.GetSegmentPath(&bundleInterfaces.BundleInterface[i])] = types.YChild{"BundleInterface", &bundleInterfaces.BundleInterface[i]}
    }
    bundleInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInterfaces.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS policy direction output.
    BundleOutput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput

    // QoS policy direction output.
    BundleInput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetEntityData() *types.CommonEntityData {
    bundleInterface.EntityData.YFilter = bundleInterface.YFilter
    bundleInterface.EntityData.YangName = "bundle-interface"
    bundleInterface.EntityData.BundleName = "cisco_ios_xr"
    bundleInterface.EntityData.ParentYangName = "bundle-interfaces"
    bundleInterface.EntityData.SegmentPath = "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
    bundleInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterface.EntityData.Children = make(map[string]types.YChild)
    bundleInterface.EntityData.Children["bundle-output"] = types.YChild{"BundleOutput", &bundleInterface.BundleOutput}
    bundleInterface.EntityData.Children["bundle-input"] = types.YChild{"BundleInput", &bundleInterface.BundleInput}
    bundleInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterface.InterfaceName}
    return &(bundleInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput
// QoS policy direction output
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput) GetEntityData() *types.CommonEntityData {
    bundleOutput.EntityData.YFilter = bundleOutput.YFilter
    bundleOutput.EntityData.YangName = "bundle-output"
    bundleOutput.EntityData.BundleName = "cisco_ios_xr"
    bundleOutput.EntityData.ParentYangName = "bundle-interface"
    bundleOutput.EntityData.SegmentPath = "bundle-output"
    bundleOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleOutput.EntityData.Children = make(map[string]types.YChild)
    bundleOutput.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleOutput.MemberInterfaces}
    bundleOutput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleOutput.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces) GetEntityData() *types.CommonEntityData {
    memberInterfaces.EntityData.YFilter = memberInterfaces.YFilter
    memberInterfaces.EntityData.YangName = "member-interfaces"
    memberInterfaces.EntityData.BundleName = "cisco_ios_xr"
    memberInterfaces.EntityData.ParentYangName = "bundle-output"
    memberInterfaces.EntityData.SegmentPath = "member-interfaces"
    memberInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterfaces.EntityData.Children = make(map[string]types.YChild)
    memberInterfaces.EntityData.Children["member-interface"] = types.YChild{"MemberInterface", nil}
    for i := range memberInterfaces.MemberInterface {
        memberInterfaces.EntityData.Children[types.GetSegmentPath(&memberInterfaces.MemberInterface[i])] = types.YChild{"MemberInterface", &memberInterfaces.MemberInterface[i]}
    }
    memberInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memberInterfaces.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS policy direction egress.
    Details PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface) GetEntityData() *types.CommonEntityData {
    memberInterface.EntityData.YFilter = memberInterface.YFilter
    memberInterface.EntityData.YangName = "member-interface"
    memberInterface.EntityData.BundleName = "cisco_ios_xr"
    memberInterface.EntityData.ParentYangName = "member-interfaces"
    memberInterface.EntityData.SegmentPath = "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
    memberInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterface.EntityData.Children = make(map[string]types.YChild)
    memberInterface.EntityData.Children["details"] = types.YChild{"Details", &memberInterface.Details}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details
// QoS policy direction egress
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header

    // Trident QoS policy details.
    Policy PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy

    // Typhoon QoS policy details.
    PolicyTyphoon PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon
}

func (details *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "member-interface"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["policy"] = types.YChild{"Policy", &details.Policy}
    details.EntityData.Children["policy-typhoon"] = types.YChild{"PolicyTyphoon", &details.PolicyTyphoon}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header
// QoS policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Interface config and programmed parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters

    // Bandwidth that was programmed.
    ProgrammedBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &header.InterfaceParameters}
    header.EntityData.Children["programmed-bandwidth"] = types.YChild{"ProgrammedBandwidth", &header.ProgrammedBandwidth}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters
// Interface config and programmed parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth due to port speed change.
    PortConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth

    // Bandwidth obtain from IM.
    AncpConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth

    // ANCP bandwidth that was programmed.
    AncpProgrammedBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth

    // Bandwidth that was programmed.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "header"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["port-config-bandwidth"] = types.YChild{"PortConfigBandwidth", &interfaceParameters.PortConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-config-bandwidth"] = types.YChild{"AncpConfigBandwidth", &interfaceParameters.AncpConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-programmed-bandwidth"] = types.YChild{"AncpProgrammedBandwidth", &interfaceParameters.AncpProgrammedBandwidth}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth
// Bandwidth due to port speed change
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portConfigBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth) GetEntityData() *types.CommonEntityData {
    portConfigBandwidth.EntityData.YFilter = portConfigBandwidth.YFilter
    portConfigBandwidth.EntityData.YangName = "port-config-bandwidth"
    portConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    portConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    portConfigBandwidth.EntityData.SegmentPath = "port-config-bandwidth"
    portConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    portConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    portConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", portConfigBandwidth.Value}
    portConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portConfigBandwidth.Unit}
    return &(portConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth
// Bandwidth obtain from IM
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpConfigBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth) GetEntityData() *types.CommonEntityData {
    ancpConfigBandwidth.EntityData.YFilter = ancpConfigBandwidth.YFilter
    ancpConfigBandwidth.EntityData.YangName = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpConfigBandwidth.EntityData.SegmentPath = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpConfigBandwidth.Value}
    ancpConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpConfigBandwidth.Unit}
    return &(ancpConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth
// ANCP bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpProgrammedBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    ancpProgrammedBandwidth.EntityData.YFilter = ancpProgrammedBandwidth.YFilter
    ancpProgrammedBandwidth.EntityData.YangName = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpProgrammedBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpProgrammedBandwidth.EntityData.SegmentPath = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpProgrammedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpProgrammedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpProgrammedBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpProgrammedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpProgrammedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpProgrammedBandwidth.Value}
    ancpProgrammedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpProgrammedBandwidth.Unit}
    return &(ancpProgrammedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
    portShaperRate.EntityData.YFilter = portShaperRate.YFilter
    portShaperRate.EntityData.YangName = "port-shaper-rate"
    portShaperRate.EntityData.BundleName = "cisco_ios_xr"
    portShaperRate.EntityData.ParentYangName = "interface-parameters"
    portShaperRate.EntityData.SegmentPath = "port-shaper-rate"
    portShaperRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portShaperRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portShaperRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portShaperRate.EntityData.Children = make(map[string]types.YChild)
    portShaperRate.EntityData.Leafs = make(map[string]types.YLeaf)
    portShaperRate.EntityData.Leafs["value"] = types.YLeaf{"Value", portShaperRate.Value}
    portShaperRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portShaperRate.Unit}
    return &(portShaperRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (programmedBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    programmedBandwidth.EntityData.YFilter = programmedBandwidth.YFilter
    programmedBandwidth.EntityData.YangName = "programmed-bandwidth"
    programmedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    programmedBandwidth.EntityData.ParentYangName = "header"
    programmedBandwidth.EntityData.SegmentPath = "programmed-bandwidth"
    programmedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedBandwidth.EntityData.Children = make(map[string]types.YChild)
    programmedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", programmedBandwidth.Value}
    programmedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", programmedBandwidth.Unit}
    return &(programmedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy
// Trident QoS policy details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v1. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1.
    QosShowEaStV1 []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1
}

func (policy *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "details"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Children["qos-show-ea-st-v1"] = types.YChild{"QosShowEaStV1", nil}
    for i := range policy.QosShowEaStV1 {
        policy.EntityData.Children[types.GetSegmentPath(&policy.QosShowEaStV1[i])] = types.YChild{"QosShowEaStV1", &policy.QosShowEaStV1[i]}
    }
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policy.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1
// qos show ea st v1
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark
}

func (qosShowEaStV1 *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1) GetEntityData() *types.CommonEntityData {
    qosShowEaStV1.EntityData.YFilter = qosShowEaStV1.YFilter
    qosShowEaStV1.EntityData.YangName = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV1.EntityData.ParentYangName = "policy"
    qosShowEaStV1.EntityData.SegmentPath = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV1.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV1.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV1.Queue}
    qosShowEaStV1.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV1.QueueLimitParameters}
    qosShowEaStV1.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV1.Shape}
    qosShowEaStV1.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV1.Police}
    qosShowEaStV1.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV1.Wfq}
    qosShowEaStV1.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV1.Wred}
    qosShowEaStV1.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV1.Mark}
    qosShowEaStV1.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV1.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV1.ClassLevel}
    qosShowEaStV1.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV1.ClassName}
    qosShowEaStV1.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV1.PolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV1.ParentPolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV1.ParentClassName}
    return &(qosShowEaStV1.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue
// Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", queueLimitParameters.ProfileId}
    queueLimitParameters.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", queueLimitParameters.ScalingProfileId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape
// Shape parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v1"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir"] = types.YChild{"Cir", &shape.Cir}
    shape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &shape.ConfigBandwidth}
    shape.EntityData.Children["cbs"] = types.YChild{"Cbs", &shape.Cbs}
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", shape.ProfileId}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police
// Police parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v1"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir
// PIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Table ID. The type is interface{} with range: 0..255.
    TableId interface{}

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    wred.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", wred.TableId}
    wred.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wred.ProfileId}
    wred.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", wred.ScalingProfileId}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED precedence match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark
// Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v1"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon
// Typhoon QoS policy details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v2. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2.
    QosShowEaStV2 []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2
}

func (policyTyphoon *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon) GetEntityData() *types.CommonEntityData {
    policyTyphoon.EntityData.YFilter = policyTyphoon.YFilter
    policyTyphoon.EntityData.YangName = "policy-typhoon"
    policyTyphoon.EntityData.BundleName = "cisco_ios_xr"
    policyTyphoon.EntityData.ParentYangName = "details"
    policyTyphoon.EntityData.SegmentPath = "policy-typhoon"
    policyTyphoon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyTyphoon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyTyphoon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyTyphoon.EntityData.Children = make(map[string]types.YChild)
    policyTyphoon.EntityData.Children["qos-show-ea-st-v2"] = types.YChild{"QosShowEaStV2", nil}
    for i := range policyTyphoon.QosShowEaStV2 {
        policyTyphoon.EntityData.Children[types.GetSegmentPath(&policyTyphoon.QosShowEaStV2[i])] = types.YChild{"QosShowEaStV2", &policyTyphoon.QosShowEaStV2[i]}
    }
    policyTyphoon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyTyphoon.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2
// qos show ea st v2
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark
}

func (qosShowEaStV2 *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2) GetEntityData() *types.CommonEntityData {
    qosShowEaStV2.EntityData.YFilter = qosShowEaStV2.YFilter
    qosShowEaStV2.EntityData.YangName = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV2.EntityData.ParentYangName = "policy-typhoon"
    qosShowEaStV2.EntityData.SegmentPath = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV2.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV2.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV2.Queue}
    qosShowEaStV2.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV2.QueueLimitParameters}
    qosShowEaStV2.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV2.Shape}
    qosShowEaStV2.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV2.Police}
    qosShowEaStV2.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV2.Wfq}
    qosShowEaStV2.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV2.Wred}
    qosShowEaStV2.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV2.Mark}
    qosShowEaStV2.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV2.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV2.ClassLevel}
    qosShowEaStV2.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV2.ClassName}
    qosShowEaStV2.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV2.PolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV2.ParentPolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV2.ParentClassName}
    return &(qosShowEaStV2.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue
// Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", queueLimitParameters.AbsoluteIndex}
    queueLimitParameters.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", queueLimitParameters.TemplateId}
    queueLimitParameters.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", queueLimitParameters.CurveId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape
// Shape parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CIR Shaper type. The type is ShapeProfiletypeV2.
    CirShapeType interface{}

    // PIR Shaper type. The type is ShapeProfiletypeV2.
    PirShapeType interface{}

    // CIR shaper params.
    CirShape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape

    // PIR shaper params.
    PirShape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v2"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir-shape"] = types.YChild{"CirShape", &shape.CirShape}
    shape.EntityData.Children["pir-shape"] = types.YChild{"PirShape", &shape.PirShape}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["cir-shape-type"] = types.YLeaf{"CirShapeType", shape.CirShapeType}
    shape.EntityData.Leafs["pir-shape-type"] = types.YLeaf{"PirShapeType", shape.PirShapeType}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape
// CIR shaper params
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
}

func (cirShape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape) GetEntityData() *types.CommonEntityData {
    cirShape.EntityData.YFilter = cirShape.YFilter
    cirShape.EntityData.YangName = "cir-shape"
    cirShape.EntityData.BundleName = "cisco_ios_xr"
    cirShape.EntityData.ParentYangName = "shape"
    cirShape.EntityData.SegmentPath = "cir-shape"
    cirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cirShape.EntityData.Children = make(map[string]types.YChild)
    cirShape.EntityData.Children["cir"] = types.YChild{"Cir", &cirShape.Cir}
    cirShape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &cirShape.ConfigBandwidth}
    cirShape.EntityData.Children["cbs"] = types.YChild{"Cbs", &cirShape.Cbs}
    cirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    cirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", cirShape.ChunkId}
    cirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", cirShape.ProfileId}
    cirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", cirShape.ScaleFactor}
    return &(cirShape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "cir-shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "cir-shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "cir-shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
// PIR shaper params
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
}

func (pirShape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape) GetEntityData() *types.CommonEntityData {
    pirShape.EntityData.YFilter = pirShape.YFilter
    pirShape.EntityData.YangName = "pir-shape"
    pirShape.EntityData.BundleName = "cisco_ios_xr"
    pirShape.EntityData.ParentYangName = "shape"
    pirShape.EntityData.SegmentPath = "pir-shape"
    pirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pirShape.EntityData.Children = make(map[string]types.YChild)
    pirShape.EntityData.Children["pir"] = types.YChild{"Pir", &pirShape.Pir}
    pirShape.EntityData.Children["pbs"] = types.YChild{"Pbs", &pirShape.Pbs}
    pirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    pirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", pirShape.ChunkId}
    pirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pirShape.ProfileId}
    pirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", pirShape.ScaleFactor}
    return &(pirShape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "pir-shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "pir-shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police
// Police parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v2"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir
// PIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // WRED match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // WRED exp match if EXP start value equals to end value Format: <start-value>
    // , else range Format: <start-value> <end-value>. The type is string.
    ExpMatch interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", curve.AbsoluteIndex}
    curve.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", curve.TemplateId}
    curve.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", curve.CurveId}
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    curve.EntityData.Leafs["exp-match"] = types.YLeaf{"ExpMatch", curve.ExpMatch}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark
// Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v2"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleOutput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput
// QoS policy direction output
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput) GetEntityData() *types.CommonEntityData {
    bundleInput.EntityData.YFilter = bundleInput.YFilter
    bundleInput.EntityData.YangName = "bundle-input"
    bundleInput.EntityData.BundleName = "cisco_ios_xr"
    bundleInput.EntityData.ParentYangName = "bundle-interface"
    bundleInput.EntityData.SegmentPath = "bundle-input"
    bundleInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInput.EntityData.Children = make(map[string]types.YChild)
    bundleInput.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleInput.MemberInterfaces}
    bundleInput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInput.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces) GetEntityData() *types.CommonEntityData {
    memberInterfaces.EntityData.YFilter = memberInterfaces.YFilter
    memberInterfaces.EntityData.YangName = "member-interfaces"
    memberInterfaces.EntityData.BundleName = "cisco_ios_xr"
    memberInterfaces.EntityData.ParentYangName = "bundle-input"
    memberInterfaces.EntityData.SegmentPath = "member-interfaces"
    memberInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterfaces.EntityData.Children = make(map[string]types.YChild)
    memberInterfaces.EntityData.Children["member-interface"] = types.YChild{"MemberInterface", nil}
    for i := range memberInterfaces.MemberInterface {
        memberInterfaces.EntityData.Children[types.GetSegmentPath(&memberInterfaces.MemberInterface[i])] = types.YChild{"MemberInterface", &memberInterfaces.MemberInterface[i]}
    }
    memberInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(memberInterfaces.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS policy direction egress.
    Details PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface) GetEntityData() *types.CommonEntityData {
    memberInterface.EntityData.YFilter = memberInterface.YFilter
    memberInterface.EntityData.YangName = "member-interface"
    memberInterface.EntityData.BundleName = "cisco_ios_xr"
    memberInterface.EntityData.ParentYangName = "member-interfaces"
    memberInterface.EntityData.SegmentPath = "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
    memberInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterface.EntityData.Children = make(map[string]types.YChild)
    memberInterface.EntityData.Children["details"] = types.YChild{"Details", &memberInterface.Details}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details
// QoS policy direction egress
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header

    // Trident QoS policy details.
    Policy PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy

    // Typhoon QoS policy details.
    PolicyTyphoon PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon
}

func (details *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "member-interface"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["policy"] = types.YChild{"Policy", &details.Policy}
    details.EntityData.Children["policy-typhoon"] = types.YChild{"PolicyTyphoon", &details.PolicyTyphoon}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header
// QoS policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Interface config and programmed parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters

    // Bandwidth that was programmed.
    ProgrammedBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &header.InterfaceParameters}
    header.EntityData.Children["programmed-bandwidth"] = types.YChild{"ProgrammedBandwidth", &header.ProgrammedBandwidth}
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters
// Interface config and programmed parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth due to port speed change.
    PortConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth

    // Bandwidth obtain from IM.
    AncpConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth

    // ANCP bandwidth that was programmed.
    AncpProgrammedBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth

    // Bandwidth that was programmed.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "header"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["port-config-bandwidth"] = types.YChild{"PortConfigBandwidth", &interfaceParameters.PortConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-config-bandwidth"] = types.YChild{"AncpConfigBandwidth", &interfaceParameters.AncpConfigBandwidth}
    interfaceParameters.EntityData.Children["ancp-programmed-bandwidth"] = types.YChild{"AncpProgrammedBandwidth", &interfaceParameters.AncpProgrammedBandwidth}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth
// Bandwidth due to port speed change
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portConfigBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortConfigBandwidth) GetEntityData() *types.CommonEntityData {
    portConfigBandwidth.EntityData.YFilter = portConfigBandwidth.YFilter
    portConfigBandwidth.EntityData.YangName = "port-config-bandwidth"
    portConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    portConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    portConfigBandwidth.EntityData.SegmentPath = "port-config-bandwidth"
    portConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    portConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    portConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", portConfigBandwidth.Value}
    portConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portConfigBandwidth.Unit}
    return &(portConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth
// Bandwidth obtain from IM
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpConfigBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpConfigBandwidth) GetEntityData() *types.CommonEntityData {
    ancpConfigBandwidth.EntityData.YFilter = ancpConfigBandwidth.YFilter
    ancpConfigBandwidth.EntityData.YangName = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpConfigBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpConfigBandwidth.EntityData.SegmentPath = "ancp-config-bandwidth"
    ancpConfigBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpConfigBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpConfigBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpConfigBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpConfigBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpConfigBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpConfigBandwidth.Value}
    ancpConfigBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpConfigBandwidth.Unit}
    return &(ancpConfigBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth
// ANCP bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (ancpProgrammedBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_AncpProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    ancpProgrammedBandwidth.EntityData.YFilter = ancpProgrammedBandwidth.YFilter
    ancpProgrammedBandwidth.EntityData.YangName = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    ancpProgrammedBandwidth.EntityData.ParentYangName = "interface-parameters"
    ancpProgrammedBandwidth.EntityData.SegmentPath = "ancp-programmed-bandwidth"
    ancpProgrammedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ancpProgrammedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ancpProgrammedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ancpProgrammedBandwidth.EntityData.Children = make(map[string]types.YChild)
    ancpProgrammedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    ancpProgrammedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", ancpProgrammedBandwidth.Value}
    ancpProgrammedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", ancpProgrammedBandwidth.Unit}
    return &(ancpProgrammedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
    portShaperRate.EntityData.YFilter = portShaperRate.YFilter
    portShaperRate.EntityData.YangName = "port-shaper-rate"
    portShaperRate.EntityData.BundleName = "cisco_ios_xr"
    portShaperRate.EntityData.ParentYangName = "interface-parameters"
    portShaperRate.EntityData.SegmentPath = "port-shaper-rate"
    portShaperRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portShaperRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portShaperRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portShaperRate.EntityData.Children = make(map[string]types.YChild)
    portShaperRate.EntityData.Leafs = make(map[string]types.YLeaf)
    portShaperRate.EntityData.Leafs["value"] = types.YLeaf{"Value", portShaperRate.Value}
    portShaperRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", portShaperRate.Unit}
    return &(portShaperRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth
// Bandwidth that was programmed
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (programmedBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Header_ProgrammedBandwidth) GetEntityData() *types.CommonEntityData {
    programmedBandwidth.EntityData.YFilter = programmedBandwidth.YFilter
    programmedBandwidth.EntityData.YangName = "programmed-bandwidth"
    programmedBandwidth.EntityData.BundleName = "cisco_ios_xr"
    programmedBandwidth.EntityData.ParentYangName = "header"
    programmedBandwidth.EntityData.SegmentPath = "programmed-bandwidth"
    programmedBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedBandwidth.EntityData.Children = make(map[string]types.YChild)
    programmedBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", programmedBandwidth.Value}
    programmedBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", programmedBandwidth.Unit}
    return &(programmedBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy
// Trident QoS policy details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v1. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1.
    QosShowEaStV1 []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1
}

func (policy *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "details"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Children["qos-show-ea-st-v1"] = types.YChild{"QosShowEaStV1", nil}
    for i := range policy.QosShowEaStV1 {
        policy.EntityData.Children[types.GetSegmentPath(&policy.QosShowEaStV1[i])] = types.YChild{"QosShowEaStV1", &policy.QosShowEaStV1[i]}
    }
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policy.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1
// qos show ea st v1
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark
}

func (qosShowEaStV1 *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1) GetEntityData() *types.CommonEntityData {
    qosShowEaStV1.EntityData.YFilter = qosShowEaStV1.YFilter
    qosShowEaStV1.EntityData.YangName = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV1.EntityData.ParentYangName = "policy"
    qosShowEaStV1.EntityData.SegmentPath = "qos-show-ea-st-v1"
    qosShowEaStV1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV1.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV1.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV1.Queue}
    qosShowEaStV1.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV1.QueueLimitParameters}
    qosShowEaStV1.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV1.Shape}
    qosShowEaStV1.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV1.Police}
    qosShowEaStV1.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV1.Wfq}
    qosShowEaStV1.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV1.Wred}
    qosShowEaStV1.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV1.Mark}
    qosShowEaStV1.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV1.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV1.ClassLevel}
    qosShowEaStV1.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV1.ClassName}
    qosShowEaStV1.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV1.PolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV1.ParentPolicyName}
    qosShowEaStV1.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV1.ParentClassName}
    return &(qosShowEaStV1.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue
// Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v1"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", queueLimitParameters.ProfileId}
    queueLimitParameters.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", queueLimitParameters.ScalingProfileId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape
// Shape parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v1"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir"] = types.YChild{"Cir", &shape.Cir}
    shape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &shape.ConfigBandwidth}
    shape.EntityData.Children["cbs"] = types.YChild{"Cbs", &shape.Cbs}
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", shape.ProfileId}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Shape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police
// Police parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v1"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir
// PIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Table ID. The type is interface{} with range: 0..255.
    TableId interface{}

    // Profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Scaling profile ID. The type is interface{} with range: 0..4294967295.
    ScalingProfileId interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v1"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    wred.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", wred.TableId}
    wred.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wred.ProfileId}
    wred.EntityData.Leafs["scaling-profile-id"] = types.YLeaf{"ScalingProfileId", wred.ScalingProfileId}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED precedence match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark
// Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v1"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_Policy_QosShowEaStV1_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon
// Typhoon QoS policy details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea st v2. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2.
    QosShowEaStV2 []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2
}

func (policyTyphoon *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon) GetEntityData() *types.CommonEntityData {
    policyTyphoon.EntityData.YFilter = policyTyphoon.YFilter
    policyTyphoon.EntityData.YangName = "policy-typhoon"
    policyTyphoon.EntityData.BundleName = "cisco_ios_xr"
    policyTyphoon.EntityData.ParentYangName = "details"
    policyTyphoon.EntityData.SegmentPath = "policy-typhoon"
    policyTyphoon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyTyphoon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyTyphoon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyTyphoon.EntityData.Children = make(map[string]types.YChild)
    policyTyphoon.EntityData.Children["qos-show-ea-st-v2"] = types.YChild{"QosShowEaStV2", nil}
    for i := range policyTyphoon.QosShowEaStV2 {
        policyTyphoon.EntityData.Children[types.GetSegmentPath(&policyTyphoon.QosShowEaStV2[i])] = types.YChild{"QosShowEaStV2", &policyTyphoon.QosShowEaStV2[i]}
    }
    policyTyphoon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policyTyphoon.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2
// qos show ea st v2
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Parent policy name. The type is string with length: 0..65.
    ParentPolicyName interface{}

    // Parent class name. The type is string with length: 0..65.
    ParentClassName interface{}

    // Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue

    // Queue limit parameters.
    QueueLimitParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters

    // Shape parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape

    // Police parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police

    // WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq

    // WRED parameters.
    Wred PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred

    // Mark parameters.
    Mark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark
}

func (qosShowEaStV2 *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2) GetEntityData() *types.CommonEntityData {
    qosShowEaStV2.EntityData.YFilter = qosShowEaStV2.YFilter
    qosShowEaStV2.EntityData.YangName = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaStV2.EntityData.ParentYangName = "policy-typhoon"
    qosShowEaStV2.EntityData.SegmentPath = "qos-show-ea-st-v2"
    qosShowEaStV2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaStV2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaStV2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaStV2.EntityData.Children = make(map[string]types.YChild)
    qosShowEaStV2.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowEaStV2.Queue}
    qosShowEaStV2.EntityData.Children["queue-limit-parameters"] = types.YChild{"QueueLimitParameters", &qosShowEaStV2.QueueLimitParameters}
    qosShowEaStV2.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowEaStV2.Shape}
    qosShowEaStV2.EntityData.Children["police"] = types.YChild{"Police", &qosShowEaStV2.Police}
    qosShowEaStV2.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowEaStV2.Wfq}
    qosShowEaStV2.EntityData.Children["wred"] = types.YChild{"Wred", &qosShowEaStV2.Wred}
    qosShowEaStV2.EntityData.Children["mark"] = types.YChild{"Mark", &qosShowEaStV2.Mark}
    qosShowEaStV2.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaStV2.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaStV2.ClassLevel}
    qosShowEaStV2.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaStV2.ClassName}
    qosShowEaStV2.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaStV2.PolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-policy-name"] = types.YLeaf{"ParentPolicyName", qosShowEaStV2.ParentPolicyName}
    qosShowEaStV2.EntityData.Leafs["parent-class-name"] = types.YLeaf{"ParentClassName", qosShowEaStV2.ParentClassName}
    return &(qosShowEaStV2.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue
// Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue type. The type is Queue.
    QueueType interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    queue.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", queue.ClassName}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters
// Queue limit parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // Queue limit in kbytes.
    QueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit

    // Config queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
}

func (queueLimitParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters) GetEntityData() *types.CommonEntityData {
    queueLimitParameters.EntityData.YFilter = queueLimitParameters.YFilter
    queueLimitParameters.EntityData.YangName = "queue-limit-parameters"
    queueLimitParameters.EntityData.BundleName = "cisco_ios_xr"
    queueLimitParameters.EntityData.ParentYangName = "qos-show-ea-st-v2"
    queueLimitParameters.EntityData.SegmentPath = "queue-limit-parameters"
    queueLimitParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimitParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimitParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimitParameters.EntityData.Children = make(map[string]types.YChild)
    queueLimitParameters.EntityData.Children["queue-limit"] = types.YChild{"QueueLimit", &queueLimitParameters.QueueLimit}
    queueLimitParameters.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &queueLimitParameters.ConfigQueueLimit}
    queueLimitParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimitParameters.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", queueLimitParameters.AbsoluteIndex}
    queueLimitParameters.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", queueLimitParameters.TemplateId}
    queueLimitParameters.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", queueLimitParameters.CurveId}
    return &(queueLimitParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit
// Queue limit in kbytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (queueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = make(map[string]types.YChild)
    queueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    queueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", queueLimit.Value}
    queueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", queueLimit.Unit}
    return &(queueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit
// Config queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_QueueLimitParameters_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "queue-limit-parameters"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["value"] = types.YLeaf{"Value", configQueueLimit.Value}
    configQueueLimit.EntityData.Leafs["unit"] = types.YLeaf{"Unit", configQueueLimit.Unit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape
// Shape parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CIR Shaper type. The type is ShapeProfiletypeV2.
    CirShapeType interface{}

    // PIR Shaper type. The type is ShapeProfiletypeV2.
    PirShapeType interface{}

    // CIR shaper params.
    CirShape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape

    // PIR shaper params.
    PirShape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-ea-st-v2"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["cir-shape"] = types.YChild{"CirShape", &shape.CirShape}
    shape.EntityData.Children["pir-shape"] = types.YChild{"PirShape", &shape.PirShape}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    shape.EntityData.Leafs["cir-shape-type"] = types.YLeaf{"CirShapeType", shape.CirShapeType}
    shape.EntityData.Leafs["pir-shape-type"] = types.YLeaf{"PirShapeType", shape.PirShapeType}
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape
// CIR shaper params
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // CIR in kbps.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir

    // Config bandwidth.
    ConfigBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth

    // CBS in bytes.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
}

func (cirShape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape) GetEntityData() *types.CommonEntityData {
    cirShape.EntityData.YFilter = cirShape.YFilter
    cirShape.EntityData.YangName = "cir-shape"
    cirShape.EntityData.BundleName = "cisco_ios_xr"
    cirShape.EntityData.ParentYangName = "shape"
    cirShape.EntityData.SegmentPath = "cir-shape"
    cirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cirShape.EntityData.Children = make(map[string]types.YChild)
    cirShape.EntityData.Children["cir"] = types.YChild{"Cir", &cirShape.Cir}
    cirShape.EntityData.Children["config-bandwidth"] = types.YChild{"ConfigBandwidth", &cirShape.ConfigBandwidth}
    cirShape.EntityData.Children["cbs"] = types.YChild{"Cbs", &cirShape.Cbs}
    cirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    cirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", cirShape.ChunkId}
    cirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", cirShape.ProfileId}
    cirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", cirShape.ScaleFactor}
    return &(cirShape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir
// CIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "cir-shape"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth
// Config bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth rate.
    MinimumRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
}

func (configBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth) GetEntityData() *types.CommonEntityData {
    configBandwidth.EntityData.YFilter = configBandwidth.YFilter
    configBandwidth.EntityData.YangName = "config-bandwidth"
    configBandwidth.EntityData.BundleName = "cisco_ios_xr"
    configBandwidth.EntityData.ParentYangName = "cir-shape"
    configBandwidth.EntityData.SegmentPath = "config-bandwidth"
    configBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configBandwidth.EntityData.Children = make(map[string]types.YChild)
    configBandwidth.EntityData.Children["minimum-rate"] = types.YChild{"MinimumRate", &configBandwidth.MinimumRate}
    configBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate
// Minimum bandwidth rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minimumRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_ConfigBandwidth_MinimumRate) GetEntityData() *types.CommonEntityData {
    minimumRate.EntityData.YFilter = minimumRate.YFilter
    minimumRate.EntityData.YangName = "minimum-rate"
    minimumRate.EntityData.BundleName = "cisco_ios_xr"
    minimumRate.EntityData.ParentYangName = "config-bandwidth"
    minimumRate.EntityData.SegmentPath = "minimum-rate"
    minimumRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minimumRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minimumRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minimumRate.EntityData.Children = make(map[string]types.YChild)
    minimumRate.EntityData.Leafs = make(map[string]types.YLeaf)
    minimumRate.EntityData.Leafs["value"] = types.YLeaf{"Value", minimumRate.Value}
    minimumRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minimumRate.Unit}
    return &(minimumRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs
// CBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_CirShape_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "cir-shape"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape
// PIR shaper params
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Shape profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // Scale Factor. The type is interface{} with range: 0..65535.
    ScaleFactor interface{}

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
}

func (pirShape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape) GetEntityData() *types.CommonEntityData {
    pirShape.EntityData.YFilter = pirShape.YFilter
    pirShape.EntityData.YangName = "pir-shape"
    pirShape.EntityData.BundleName = "cisco_ios_xr"
    pirShape.EntityData.ParentYangName = "shape"
    pirShape.EntityData.SegmentPath = "pir-shape"
    pirShape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pirShape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pirShape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pirShape.EntityData.Children = make(map[string]types.YChild)
    pirShape.EntityData.Children["pir"] = types.YChild{"Pir", &pirShape.Pir}
    pirShape.EntityData.Children["pbs"] = types.YChild{"Pbs", &pirShape.Pbs}
    pirShape.EntityData.Leafs = make(map[string]types.YLeaf)
    pirShape.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", pirShape.ChunkId}
    pirShape.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", pirShape.ProfileId}
    pirShape.EntityData.Leafs["scale-factor"] = types.YLeaf{"ScaleFactor", pirShape.ScaleFactor}
    return &(pirShape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "pir-shape"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Shape_PirShape_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "pir-shape"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police
// Police parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // Police profile ID. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs

    // PIR.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir

    // PBS.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs

    // Police config parameters.
    PoliceConfigParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-ea-st-v2"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Children["pir"] = types.YChild{"Pir", &police.Pir}
    police.EntityData.Children["pbs"] = types.YChild{"Pbs", &police.Pbs}
    police.EntityData.Children["police-config-parameters"] = types.YChild{"PoliceConfigParameters", &police.PoliceConfigParameters}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    police.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", police.ProfileId}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cir) GetEntityData() *types.CommonEntityData {
    cir.EntityData.YFilter = cir.YFilter
    cir.EntityData.YangName = "cir"
    cir.EntityData.BundleName = "cisco_ios_xr"
    cir.EntityData.ParentYangName = "police"
    cir.EntityData.SegmentPath = "cir"
    cir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cir.EntityData.Children = make(map[string]types.YChild)
    cir.EntityData.Leafs = make(map[string]types.YLeaf)
    cir.EntityData.Leafs["value"] = types.YLeaf{"Value", cir.Value}
    cir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cir.Unit}
    return &(cir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Cbs) GetEntityData() *types.CommonEntityData {
    cbs.EntityData.YFilter = cbs.YFilter
    cbs.EntityData.YangName = "cbs"
    cbs.EntityData.BundleName = "cisco_ios_xr"
    cbs.EntityData.ParentYangName = "police"
    cbs.EntityData.SegmentPath = "cbs"
    cbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbs.EntityData.Children = make(map[string]types.YChild)
    cbs.EntityData.Leafs = make(map[string]types.YLeaf)
    cbs.EntityData.Leafs["value"] = types.YLeaf{"Value", cbs.Value}
    cbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", cbs.Unit}
    return &(cbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir
// PIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pir) GetEntityData() *types.CommonEntityData {
    pir.EntityData.YFilter = pir.YFilter
    pir.EntityData.YangName = "pir"
    pir.EntityData.BundleName = "cisco_ios_xr"
    pir.EntityData.ParentYangName = "police"
    pir.EntityData.SegmentPath = "pir"
    pir.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pir.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pir.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pir.EntityData.Children = make(map[string]types.YChild)
    pir.EntityData.Leafs = make(map[string]types.YLeaf)
    pir.EntityData.Leafs["value"] = types.YLeaf{"Value", pir.Value}
    pir.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pir.Unit}
    return &(pir.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs
// PBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_Pbs) GetEntityData() *types.CommonEntityData {
    pbs.EntityData.YFilter = pbs.YFilter
    pbs.EntityData.YangName = "pbs"
    pbs.EntityData.BundleName = "cisco_ios_xr"
    pbs.EntityData.ParentYangName = "police"
    pbs.EntityData.SegmentPath = "pbs"
    pbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbs.EntityData.Children = make(map[string]types.YChild)
    pbs.EntityData.Leafs = make(map[string]types.YLeaf)
    pbs.EntityData.Leafs["value"] = types.YLeaf{"Value", pbs.Value}
    pbs.EntityData.Leafs["unit"] = types.YLeaf{"Unit", pbs.Unit}
    return &(pbs.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters
// Police config parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average rate.
    AverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate

    // Peak rate.
    PeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate

    // Conform burst.
    ConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst

    // Exceed burst.
    ExceedBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
}

func (policeConfigParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters) GetEntityData() *types.CommonEntityData {
    policeConfigParameters.EntityData.YFilter = policeConfigParameters.YFilter
    policeConfigParameters.EntityData.YangName = "police-config-parameters"
    policeConfigParameters.EntityData.BundleName = "cisco_ios_xr"
    policeConfigParameters.EntityData.ParentYangName = "police"
    policeConfigParameters.EntityData.SegmentPath = "police-config-parameters"
    policeConfigParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConfigParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConfigParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConfigParameters.EntityData.Children = make(map[string]types.YChild)
    policeConfigParameters.EntityData.Children["average-rate"] = types.YChild{"AverageRate", &policeConfigParameters.AverageRate}
    policeConfigParameters.EntityData.Children["peak-rate"] = types.YChild{"PeakRate", &policeConfigParameters.PeakRate}
    policeConfigParameters.EntityData.Children["conform-burst"] = types.YChild{"ConformBurst", &policeConfigParameters.ConformBurst}
    policeConfigParameters.EntityData.Children["exceed-burst"] = types.YChild{"ExceedBurst", &policeConfigParameters.ExceedBurst}
    policeConfigParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policeConfigParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate
// Average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (averageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_AverageRate) GetEntityData() *types.CommonEntityData {
    averageRate.EntityData.YFilter = averageRate.YFilter
    averageRate.EntityData.YangName = "average-rate"
    averageRate.EntityData.BundleName = "cisco_ios_xr"
    averageRate.EntityData.ParentYangName = "police-config-parameters"
    averageRate.EntityData.SegmentPath = "average-rate"
    averageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    averageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    averageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    averageRate.EntityData.Children = make(map[string]types.YChild)
    averageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    averageRate.EntityData.Leafs["value"] = types.YLeaf{"Value", averageRate.Value}
    averageRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", averageRate.Unit}
    return &(averageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate
// Peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (peakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police-config-parameters"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = make(map[string]types.YChild)
    peakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    peakRate.EntityData.Leafs["value"] = types.YLeaf{"Value", peakRate.Value}
    peakRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", peakRate.Unit}
    return &(peakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst
// Conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (conformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ConformBurst) GetEntityData() *types.CommonEntityData {
    conformBurst.EntityData.YFilter = conformBurst.YFilter
    conformBurst.EntityData.YangName = "conform-burst"
    conformBurst.EntityData.BundleName = "cisco_ios_xr"
    conformBurst.EntityData.ParentYangName = "police-config-parameters"
    conformBurst.EntityData.SegmentPath = "conform-burst"
    conformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformBurst.EntityData.Children = make(map[string]types.YChild)
    conformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    conformBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", conformBurst.Value}
    conformBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", conformBurst.Unit}
    return &(conformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst
// Exceed burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (exceedBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Police_PoliceConfigParameters_ExceedBurst) GetEntityData() *types.CommonEntityData {
    exceedBurst.EntityData.YFilter = exceedBurst.YFilter
    exceedBurst.EntityData.YangName = "exceed-burst"
    exceedBurst.EntityData.BundleName = "cisco_ios_xr"
    exceedBurst.EntityData.ParentYangName = "police-config-parameters"
    exceedBurst.EntityData.SegmentPath = "exceed-burst"
    exceedBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedBurst.EntityData.Children = make(map[string]types.YChild)
    exceedBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedBurst.EntityData.Leafs["value"] = types.YLeaf{"Value", exceedBurst.Value}
    exceedBurst.EntityData.Leafs["unit"] = types.YLeaf{"Unit", exceedBurst.Unit}
    return &(exceedBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq
// WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WFQ profile. The type is interface{} with range: 0..4294967295.
    ProfileId interface{}

    // Committed weight. The type is interface{} with range: 0..4294967295.
    CommittedWeight interface{}

    // Excess weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Parent Excess ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Chunk ID. The type is interface{} with range: 0..4294967295.
    ChunkId interface{}

    // Level. The type is interface{} with range: 0..255.
    Level interface{}

    // Parent bandwidth.
    ParentBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth

    // CFG Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["parent-bandwidth"] = types.YChild{"ParentBandwidth", &wfq.ParentBandwidth}
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["profile-id"] = types.YLeaf{"ProfileId", wfq.ProfileId}
    wfq.EntityData.Leafs["committed-weight"] = types.YLeaf{"CommittedWeight", wfq.CommittedWeight}
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    wfq.EntityData.Leafs["chunk-id"] = types.YLeaf{"ChunkId", wfq.ChunkId}
    wfq.EntityData.Leafs["level"] = types.YLeaf{"Level", wfq.Level}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth
// Parent bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (parentBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_ParentBandwidth) GetEntityData() *types.CommonEntityData {
    parentBandwidth.EntityData.YFilter = parentBandwidth.YFilter
    parentBandwidth.EntityData.YangName = "parent-bandwidth"
    parentBandwidth.EntityData.BundleName = "cisco_ios_xr"
    parentBandwidth.EntityData.ParentYangName = "wfq"
    parentBandwidth.EntityData.SegmentPath = "parent-bandwidth"
    parentBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentBandwidth.EntityData.Children = make(map[string]types.YChild)
    parentBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    parentBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", parentBandwidth.Value}
    parentBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", parentBandwidth.Unit}
    return &(parentBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth
// CFG Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "wfq"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", bandwidth.Value}
    bandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", bandwidth.Unit}
    return &(bandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // WRED type. The type is Wred.
    Type_ interface{}

    // Number of curves. The type is interface{} with range: 0..65535.
    CurveXr interface{}

    // Curve details. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve.
    Curve []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "qos-show-ea-st-v2"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["curve"] = types.YChild{"Curve", nil}
    for i := range wred.Curve {
        wred.EntityData.Children[types.GetSegmentPath(&wred.Curve[i])] = types.YChild{"Curve", &wred.Curve[i]}
    }
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["type"] = types.YLeaf{"Type_", wred.Type_}
    wred.EntityData.Leafs["curve-xr"] = types.YLeaf{"CurveXr", wred.CurveXr}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve
// Curve details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Absolute Index. The type is interface{} with range: 0..65535.
    AbsoluteIndex interface{}

    // Template ID. The type is interface{} with range: 0..65535.
    TemplateId interface{}

    // Curve ID. The type is interface{} with range: 0..65535.
    CurveId interface{}

    // WRED match if precedence start value equals to end value Format:
    // <start-value> , else range Format: <start-value> <end-value>. The type is
    // string.
    Match interface{}

    // WRED exp match if EXP start value equals to end value Format: <start-value>
    // , else range Format: <start-value> <end-value>. The type is string.
    ExpMatch interface{}

    // Minimum threshold.
    MinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold

    // Minimum threshold WRED context.
    MinThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig

    // Maximum threshold.
    MaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold

    // Maximum threshold WRED context.
    MaxThresholdUserConfig PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
}

func (curve *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve) GetEntityData() *types.CommonEntityData {
    curve.EntityData.YFilter = curve.YFilter
    curve.EntityData.YangName = "curve"
    curve.EntityData.BundleName = "cisco_ios_xr"
    curve.EntityData.ParentYangName = "wred"
    curve.EntityData.SegmentPath = "curve"
    curve.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    curve.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    curve.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    curve.EntityData.Children = make(map[string]types.YChild)
    curve.EntityData.Children["min-threshold"] = types.YChild{"MinThreshold", &curve.MinThreshold}
    curve.EntityData.Children["min-threshold-user-config"] = types.YChild{"MinThresholdUserConfig", &curve.MinThresholdUserConfig}
    curve.EntityData.Children["max-threshold"] = types.YChild{"MaxThreshold", &curve.MaxThreshold}
    curve.EntityData.Children["max-threshold-user-config"] = types.YChild{"MaxThresholdUserConfig", &curve.MaxThresholdUserConfig}
    curve.EntityData.Leafs = make(map[string]types.YLeaf)
    curve.EntityData.Leafs["absolute-index"] = types.YLeaf{"AbsoluteIndex", curve.AbsoluteIndex}
    curve.EntityData.Leafs["template-id"] = types.YLeaf{"TemplateId", curve.TemplateId}
    curve.EntityData.Leafs["curve-id"] = types.YLeaf{"CurveId", curve.CurveId}
    curve.EntityData.Leafs["match"] = types.YLeaf{"Match", curve.Match}
    curve.EntityData.Leafs["exp-match"] = types.YLeaf{"ExpMatch", curve.ExpMatch}
    return &(curve.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold
// Minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThreshold) GetEntityData() *types.CommonEntityData {
    minThreshold.EntityData.YFilter = minThreshold.YFilter
    minThreshold.EntityData.YangName = "min-threshold"
    minThreshold.EntityData.BundleName = "cisco_ios_xr"
    minThreshold.EntityData.ParentYangName = "curve"
    minThreshold.EntityData.SegmentPath = "min-threshold"
    minThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThreshold.EntityData.Children = make(map[string]types.YChild)
    minThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    minThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", minThreshold.Value}
    minThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThreshold.Unit}
    return &(minThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig
// Minimum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (minThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MinThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    minThresholdUserConfig.EntityData.YFilter = minThresholdUserConfig.YFilter
    minThresholdUserConfig.EntityData.YangName = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    minThresholdUserConfig.EntityData.ParentYangName = "curve"
    minThresholdUserConfig.EntityData.SegmentPath = "min-threshold-user-config"
    minThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    minThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    minThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", minThresholdUserConfig.Value}
    minThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", minThresholdUserConfig.Unit}
    return &(minThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold
// Maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThreshold) GetEntityData() *types.CommonEntityData {
    maxThreshold.EntityData.YFilter = maxThreshold.YFilter
    maxThreshold.EntityData.YangName = "max-threshold"
    maxThreshold.EntityData.BundleName = "cisco_ios_xr"
    maxThreshold.EntityData.ParentYangName = "curve"
    maxThreshold.EntityData.SegmentPath = "max-threshold"
    maxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThreshold.EntityData.Children = make(map[string]types.YChild)
    maxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThreshold.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThreshold.Value}
    maxThreshold.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThreshold.Unit}
    return &(maxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig
// Maximum threshold WRED context
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (maxThresholdUserConfig *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Wred_Curve_MaxThresholdUserConfig) GetEntityData() *types.CommonEntityData {
    maxThresholdUserConfig.EntityData.YFilter = maxThresholdUserConfig.YFilter
    maxThresholdUserConfig.EntityData.YangName = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.BundleName = "cisco_ios_xr"
    maxThresholdUserConfig.EntityData.ParentYangName = "curve"
    maxThresholdUserConfig.EntityData.SegmentPath = "max-threshold-user-config"
    maxThresholdUserConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxThresholdUserConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxThresholdUserConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxThresholdUserConfig.EntityData.Children = make(map[string]types.YChild)
    maxThresholdUserConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    maxThresholdUserConfig.EntityData.Leafs["value"] = types.YLeaf{"Value", maxThresholdUserConfig.Value}
    maxThresholdUserConfig.EntityData.Leafs["unit"] = types.YLeaf{"Unit", maxThresholdUserConfig.Unit}
    return &(maxThresholdUserConfig.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark
// Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Child mark only.
    ChildMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark

    // Child police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform

    // Child police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed

    // Child police violate mark.
    PoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate

    // Parent mark only.
    ParentMark PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark

    // Parent police conform mark.
    ParentPoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform

    // Parent police exceed mark.
    ParentPoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed

    // Parent police violate mark.
    ParentPoliceViolate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "qos-show-ea-st-v2"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Children["child-mark"] = types.YChild{"ChildMark", &mark.ChildMark}
    mark.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &mark.PoliceConform}
    mark.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &mark.PoliceExceed}
    mark.EntityData.Children["police-violate"] = types.YChild{"PoliceViolate", &mark.PoliceViolate}
    mark.EntityData.Children["parent-mark"] = types.YChild{"ParentMark", &mark.ParentMark}
    mark.EntityData.Children["parent-police-conform"] = types.YChild{"ParentPoliceConform", &mark.ParentPoliceConform}
    mark.EntityData.Children["parent-police-exceed"] = types.YChild{"ParentPoliceExceed", &mark.ParentPoliceExceed}
    mark.EntityData.Children["parent-police-violate"] = types.YChild{"ParentPoliceViolate", &mark.ParentPoliceViolate}
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark
// Child mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
}

func (childMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark) GetEntityData() *types.CommonEntityData {
    childMark.EntityData.YFilter = childMark.YFilter
    childMark.EntityData.YangName = "child-mark"
    childMark.EntityData.BundleName = "cisco_ios_xr"
    childMark.EntityData.ParentYangName = "mark"
    childMark.EntityData.SegmentPath = "child-mark"
    childMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childMark.EntityData.Children = make(map[string]types.YChild)
    childMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range childMark.MarkDetail {
        childMark.EntityData.Children[types.GetSegmentPath(&childMark.MarkDetail[i])] = types.YChild{"MarkDetail", &childMark.MarkDetail[i]}
    }
    childMark.EntityData.Leafs = make(map[string]types.YLeaf)
    childMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", childMark.ActionType}
    return &(childMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ChildMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "child-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform
// Child police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "mark"
    policeConform.EntityData.SegmentPath = "police-conform"
    policeConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeConform.EntityData.Children = make(map[string]types.YChild)
    policeConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeConform.MarkDetail {
        policeConform.EntityData.Children[types.GetSegmentPath(&policeConform.MarkDetail[i])] = types.YChild{"MarkDetail", &policeConform.MarkDetail[i]}
    }
    policeConform.EntityData.Leafs = make(map[string]types.YLeaf)
    policeConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeConform.ActionType}
    return &(policeConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed
// Child police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "mark"
    policeExceed.EntityData.SegmentPath = "police-exceed"
    policeExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeExceed.EntityData.Children = make(map[string]types.YChild)
    policeExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeExceed.MarkDetail {
        policeExceed.EntityData.Children[types.GetSegmentPath(&policeExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &policeExceed.MarkDetail[i]}
    }
    policeExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    policeExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeExceed.ActionType}
    return &(policeExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate
// Child police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
}

func (policeViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate) GetEntityData() *types.CommonEntityData {
    policeViolate.EntityData.YFilter = policeViolate.YFilter
    policeViolate.EntityData.YangName = "police-violate"
    policeViolate.EntityData.BundleName = "cisco_ios_xr"
    policeViolate.EntityData.ParentYangName = "mark"
    policeViolate.EntityData.SegmentPath = "police-violate"
    policeViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policeViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policeViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policeViolate.EntityData.Children = make(map[string]types.YChild)
    policeViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range policeViolate.MarkDetail {
        policeViolate.EntityData.Children[types.GetSegmentPath(&policeViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &policeViolate.MarkDetail[i]}
    }
    policeViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    policeViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", policeViolate.ActionType}
    return &(policeViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_PoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark
// Parent mark only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
}

func (parentMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark) GetEntityData() *types.CommonEntityData {
    parentMark.EntityData.YFilter = parentMark.YFilter
    parentMark.EntityData.YangName = "parent-mark"
    parentMark.EntityData.BundleName = "cisco_ios_xr"
    parentMark.EntityData.ParentYangName = "mark"
    parentMark.EntityData.SegmentPath = "parent-mark"
    parentMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentMark.EntityData.Children = make(map[string]types.YChild)
    parentMark.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentMark.MarkDetail {
        parentMark.EntityData.Children[types.GetSegmentPath(&parentMark.MarkDetail[i])] = types.YChild{"MarkDetail", &parentMark.MarkDetail[i]}
    }
    parentMark.EntityData.Leafs = make(map[string]types.YLeaf)
    parentMark.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentMark.ActionType}
    return &(parentMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentMark_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-mark"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform
// Parent police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
}

func (parentPoliceConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform) GetEntityData() *types.CommonEntityData {
    parentPoliceConform.EntityData.YFilter = parentPoliceConform.YFilter
    parentPoliceConform.EntityData.YangName = "parent-police-conform"
    parentPoliceConform.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceConform.EntityData.ParentYangName = "mark"
    parentPoliceConform.EntityData.SegmentPath = "parent-police-conform"
    parentPoliceConform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceConform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceConform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceConform.EntityData.Children = make(map[string]types.YChild)
    parentPoliceConform.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceConform.MarkDetail {
        parentPoliceConform.EntityData.Children[types.GetSegmentPath(&parentPoliceConform.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceConform.MarkDetail[i]}
    }
    parentPoliceConform.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceConform.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceConform.ActionType}
    return &(parentPoliceConform.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-conform"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed
// Parent police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
}

func (parentPoliceExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed) GetEntityData() *types.CommonEntityData {
    parentPoliceExceed.EntityData.YFilter = parentPoliceExceed.YFilter
    parentPoliceExceed.EntityData.YangName = "parent-police-exceed"
    parentPoliceExceed.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceExceed.EntityData.ParentYangName = "mark"
    parentPoliceExceed.EntityData.SegmentPath = "parent-police-exceed"
    parentPoliceExceed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceExceed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceExceed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceExceed.EntityData.Children = make(map[string]types.YChild)
    parentPoliceExceed.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceExceed.MarkDetail {
        parentPoliceExceed.EntityData.Children[types.GetSegmentPath(&parentPoliceExceed.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceExceed.MarkDetail[i]}
    }
    parentPoliceExceed.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceExceed.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceExceed.ActionType}
    return &(parentPoliceExceed.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-exceed"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate
// Parent police violate mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
}

func (parentPoliceViolate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate) GetEntityData() *types.CommonEntityData {
    parentPoliceViolate.EntityData.YFilter = parentPoliceViolate.YFilter
    parentPoliceViolate.EntityData.YangName = "parent-police-violate"
    parentPoliceViolate.EntityData.BundleName = "cisco_ios_xr"
    parentPoliceViolate.EntityData.ParentYangName = "mark"
    parentPoliceViolate.EntityData.SegmentPath = "parent-police-violate"
    parentPoliceViolate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentPoliceViolate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentPoliceViolate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentPoliceViolate.EntityData.Children = make(map[string]types.YChild)
    parentPoliceViolate.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range parentPoliceViolate.MarkDetail {
        parentPoliceViolate.EntityData.Children[types.GetSegmentPath(&parentPoliceViolate.MarkDetail[i])] = types.YChild{"MarkDetail", &parentPoliceViolate.MarkDetail[i]}
    }
    parentPoliceViolate.EntityData.Leafs = make(map[string]types.YLeaf)
    parentPoliceViolate.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", parentPoliceViolate.ActionType}
    return &(parentPoliceViolate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_BundleInput_MemberInterfaces_MemberInterface_Details_PolicyTyphoon_QosShowEaStV2_Mark_ParentPoliceViolate_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "parent-police-violate"
    markDetail.EntityData.SegmentPath = "mark-detail"
    markDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markDetail.EntityData.Children = make(map[string]types.YChild)
    markDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    markDetail.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", markDetail.MarkValue}
    markDetail.EntityData.Leafs["action-opcode"] = types.YLeaf{"ActionOpcode", markDetail.ActionOpcode}
    return &(markDetail.EntityData)
}

