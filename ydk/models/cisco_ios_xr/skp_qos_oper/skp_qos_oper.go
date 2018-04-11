// This module contains a collection of YANG definitions
// for Cisco IOS-XR skp-qos package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-qos: QoS Skywarp platform operational data 
//   platform-qos-ea: platform qos ea
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package skp_qos_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package skp_qos_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-skp-qos-oper platform-qos}", reflect.TypeOf(PlatformQos{}))
    ydk.RegisterEntity("Cisco-IOS-XR-skp-qos-oper:platform-qos", reflect.TypeOf(PlatformQos{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-skp-qos-oper platform-qos-ea}", reflect.TypeOf(PlatformQosEa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-skp-qos-oper:platform-qos-ea", reflect.TypeOf(PlatformQosEa{}))
}

// Wred represents Wred
type Wred string

const (
    // wred cos cmd
    Wred_wred_cos_cmd Wred = "wred-cos-cmd"

    // wred dscp cmd
    Wred_wred_dscp_cmd Wred = "wred-dscp-cmd"

    // wred precedence cmd
    Wred_wred_precedence_cmd Wred = "wred-precedence-cmd"

    // wred discard class cmd
    Wred_wred_discard_class_cmd Wred = "wred-discard-class-cmd"

    // wred mpls exp cmd
    Wred_wred_mpls_exp_cmd Wred = "wred-mpls-exp-cmd"

    // red with user min max
    Wred_red_with_user_min_max Wred = "red-with-user-min-max"

    // red with default min max
    Wred_red_with_default_min_max Wred = "red-with-default-min-max"

    // wred dei cmd
    Wred_wred_dei_cmd Wred = "wred-dei-cmd"

    // wred ecn cmd
    Wred_wred_ecn_cmd Wred = "wred-ecn-cmd"

    // wred invalid cmd
    Wred_wred_invalid_cmd Wred = "wred-invalid-cmd"
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

// PlatformQos
// QoS Skywarp platform operational data 
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
    platformQos.EntityData.ParentYangName = "Cisco-IOS-XR-skp-qos-oper"
    platformQos.EntityData.SegmentPath = "Cisco-IOS-XR-skp-qos-oper:platform-qos"
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

    // QoS list of bundle interfaces.
    BundleInterfaces PlatformQos_Nodes_Node_BundleInterfaces

    // QoS system capability.
    Capability PlatformQos_Nodes_Node_Capability

    // QoS list of interfaces.
    Interfaces PlatformQos_Nodes_Node_Interfaces
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
    node.EntityData.Children["bundle-interfaces"] = types.YChild{"BundleInterfaces", &node.BundleInterfaces}
    node.EntityData.Children["capability"] = types.YChild{"Capability", &node.Capability}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
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

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
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
    bundleInterface.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleInterface.MemberInterfaces}
    bundleInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterface.InterfaceName}
    return &(bundleInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetEntityData() *types.CommonEntityData {
    memberInterfaces.EntityData.YFilter = memberInterfaces.YFilter
    memberInterfaces.EntityData.YangName = "member-interfaces"
    memberInterfaces.EntityData.BundleName = "cisco_ios_xr"
    memberInterfaces.EntityData.ParentYangName = "bundle-interface"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS policy direction input.
    BundleInput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput

    // QoS policy direction output.
    BundleOutput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetEntityData() *types.CommonEntityData {
    memberInterface.EntityData.YFilter = memberInterface.YFilter
    memberInterface.EntityData.YangName = "member-interface"
    memberInterface.EntityData.BundleName = "cisco_ios_xr"
    memberInterface.EntityData.ParentYangName = "member-interfaces"
    memberInterface.EntityData.SegmentPath = "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
    memberInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterface.EntityData.Children = make(map[string]types.YChild)
    memberInterface.EntityData.Children["bundle-input"] = types.YChild{"BundleInput", &memberInterface.BundleInput}
    memberInterface.EntityData.Children["bundle-output"] = types.YChild{"BundleOutput", &memberInterface.BundleOutput}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
// QoS policy direction input
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetEntityData() *types.CommonEntityData {
    bundleInput.EntityData.YFilter = bundleInput.YFilter
    bundleInput.EntityData.YangName = "bundle-input"
    bundleInput.EntityData.BundleName = "cisco_ios_xr"
    bundleInput.EntityData.ParentYangName = "member-interface"
    bundleInput.EntityData.SegmentPath = "bundle-input"
    bundleInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInput.EntityData.Children = make(map[string]types.YChild)
    bundleInput.EntityData.Children["header"] = types.YChild{"Header", &bundleInput.Header}
    bundleInput.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &bundleInput.InterfaceParameters}
    bundleInput.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &bundleInput.SkywarpQosPolicyClass}
    bundleInput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInput.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "bundle-input"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "bundle-input"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["interface-config-rate"] = types.YChild{"InterfaceConfigRate", &interfaceParameters.InterfaceConfigRate}
    interfaceParameters.EntityData.Children["interface-program-rate"] = types.YChild{"InterfaceProgramRate", &interfaceParameters.InterfaceProgramRate}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetEntityData() *types.CommonEntityData {
    interfaceConfigRate.EntityData.YFilter = interfaceConfigRate.YFilter
    interfaceConfigRate.EntityData.YangName = "interface-config-rate"
    interfaceConfigRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigRate.EntityData.ParentYangName = "interface-parameters"
    interfaceConfigRate.EntityData.SegmentPath = "interface-config-rate"
    interfaceConfigRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigRate.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfigRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceConfigRate.Value}
    interfaceConfigRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceConfigRate.Unit}
    return &(interfaceConfigRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetEntityData() *types.CommonEntityData {
    interfaceProgramRate.EntityData.YFilter = interfaceProgramRate.YFilter
    interfaceProgramRate.EntityData.YangName = "interface-program-rate"
    interfaceProgramRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceProgramRate.EntityData.ParentYangName = "interface-parameters"
    interfaceProgramRate.EntityData.SegmentPath = "interface-program-rate"
    interfaceProgramRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProgramRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProgramRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProgramRate.EntityData.Children = make(map[string]types.YChild)
    interfaceProgramRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceProgramRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceProgramRate.Value}
    interfaceProgramRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceProgramRate.Unit}
    return &(interfaceProgramRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "bundle-input"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-pclass-st"] = types.YChild{"QosShowPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowPclassSt[i])] = types.YChild{"QosShowPclassSt", &skywarpQosPolicyClass.QosShowPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // QoS Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue

    // QoS EA Shaper parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape

    // QoS WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq

    // QoS Policer parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police

    // QoS Mark parameters.
    Marking PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowPclassSt.EntityData.YFilter = qosShowPclassSt.YFilter
    qosShowPclassSt.EntityData.YangName = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowPclassSt.EntityData.SegmentPath = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowPclassSt.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowPclassSt.Queue}
    qosShowPclassSt.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowPclassSt.Shape}
    qosShowPclassSt.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowPclassSt.Wfq}
    qosShowPclassSt.EntityData.Children["police"] = types.YChild{"Police", &qosShowPclassSt.Police}
    qosShowPclassSt.EntityData.Children["marking"] = types.YChild{"Marking", &qosShowPclassSt.Marking}
    qosShowPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowPclassSt.ClassLevel}
    qosShowPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowPclassSt.ClassName}
    return &(qosShowPclassSt.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-pclass-st"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-pclass-st"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-pclass-st"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["committed-weight"] = types.YChild{"CommittedWeight", &wfq.CommittedWeight}
    wfq.EntityData.Children["programmed-wfq"] = types.YChild{"ProgrammedWfq", &wfq.ProgrammedWfq}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetEntityData() *types.CommonEntityData {
    committedWeight.EntityData.YFilter = committedWeight.YFilter
    committedWeight.EntityData.YangName = "committed-weight"
    committedWeight.EntityData.BundleName = "cisco_ios_xr"
    committedWeight.EntityData.ParentYangName = "wfq"
    committedWeight.EntityData.SegmentPath = "committed-weight"
    committedWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedWeight.EntityData.Children = make(map[string]types.YChild)
    committedWeight.EntityData.Leafs = make(map[string]types.YLeaf)
    committedWeight.EntityData.Leafs["value"] = types.YLeaf{"Value", committedWeight.Value}
    committedWeight.EntityData.Leafs["unit"] = types.YLeaf{"Unit", committedWeight.Unit}
    return &(committedWeight.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetEntityData() *types.CommonEntityData {
    programmedWfq.EntityData.YFilter = programmedWfq.YFilter
    programmedWfq.EntityData.YangName = "programmed-wfq"
    programmedWfq.EntityData.BundleName = "cisco_ios_xr"
    programmedWfq.EntityData.ParentYangName = "wfq"
    programmedWfq.EntityData.SegmentPath = "programmed-wfq"
    programmedWfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedWfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedWfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedWfq.EntityData.Children = make(map[string]types.YChild)
    programmedWfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &programmedWfq.Bandwidth}
    programmedWfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &programmedWfq.SumOfBandwidth}
    programmedWfq.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedWfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", programmedWfq.ExcessRatio}
    return &(programmedWfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "programmed-wfq"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "programmed-wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policer ID. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-pclass-st"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", police.PolicerId}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetEntityData() *types.CommonEntityData {
    marking.EntityData.YFilter = marking.YFilter
    marking.EntityData.YangName = "marking"
    marking.EntityData.BundleName = "cisco_ios_xr"
    marking.EntityData.ParentYangName = "qos-show-pclass-st"
    marking.EntityData.SegmentPath = "marking"
    marking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    marking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    marking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    marking.EntityData.Children = make(map[string]types.YChild)
    marking.EntityData.Children["mark-only"] = types.YChild{"MarkOnly", &marking.MarkOnly}
    marking.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &marking.PoliceConform}
    marking.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &marking.PoliceExceed}
    marking.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(marking.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetEntityData() *types.CommonEntityData {
    markOnly.EntityData.YFilter = markOnly.YFilter
    markOnly.EntityData.YangName = "mark-only"
    markOnly.EntityData.BundleName = "cisco_ios_xr"
    markOnly.EntityData.ParentYangName = "marking"
    markOnly.EntityData.SegmentPath = "mark-only"
    markOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markOnly.EntityData.Children = make(map[string]types.YChild)
    markOnly.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range markOnly.MarkDetail {
        markOnly.EntityData.Children[types.GetSegmentPath(&markOnly.MarkDetail[i])] = types.YChild{"MarkDetail", &markOnly.MarkDetail[i]}
    }
    markOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    markOnly.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", markOnly.ActionType}
    return &(markOnly.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "mark-only"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
// QoS policy direction output
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetEntityData() *types.CommonEntityData {
    bundleOutput.EntityData.YFilter = bundleOutput.YFilter
    bundleOutput.EntityData.YangName = "bundle-output"
    bundleOutput.EntityData.BundleName = "cisco_ios_xr"
    bundleOutput.EntityData.ParentYangName = "member-interface"
    bundleOutput.EntityData.SegmentPath = "bundle-output"
    bundleOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleOutput.EntityData.Children = make(map[string]types.YChild)
    bundleOutput.EntityData.Children["header"] = types.YChild{"Header", &bundleOutput.Header}
    bundleOutput.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &bundleOutput.InterfaceParameters}
    bundleOutput.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &bundleOutput.SkywarpQosPolicyClass}
    bundleOutput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleOutput.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "bundle-output"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "bundle-output"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["interface-config-rate"] = types.YChild{"InterfaceConfigRate", &interfaceParameters.InterfaceConfigRate}
    interfaceParameters.EntityData.Children["interface-program-rate"] = types.YChild{"InterfaceProgramRate", &interfaceParameters.InterfaceProgramRate}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetEntityData() *types.CommonEntityData {
    interfaceConfigRate.EntityData.YFilter = interfaceConfigRate.YFilter
    interfaceConfigRate.EntityData.YangName = "interface-config-rate"
    interfaceConfigRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigRate.EntityData.ParentYangName = "interface-parameters"
    interfaceConfigRate.EntityData.SegmentPath = "interface-config-rate"
    interfaceConfigRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigRate.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfigRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceConfigRate.Value}
    interfaceConfigRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceConfigRate.Unit}
    return &(interfaceConfigRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetEntityData() *types.CommonEntityData {
    interfaceProgramRate.EntityData.YFilter = interfaceProgramRate.YFilter
    interfaceProgramRate.EntityData.YangName = "interface-program-rate"
    interfaceProgramRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceProgramRate.EntityData.ParentYangName = "interface-parameters"
    interfaceProgramRate.EntityData.SegmentPath = "interface-program-rate"
    interfaceProgramRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProgramRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProgramRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProgramRate.EntityData.Children = make(map[string]types.YChild)
    interfaceProgramRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceProgramRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceProgramRate.Value}
    interfaceProgramRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceProgramRate.Unit}
    return &(interfaceProgramRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "bundle-output"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-pclass-st"] = types.YChild{"QosShowPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowPclassSt[i])] = types.YChild{"QosShowPclassSt", &skywarpQosPolicyClass.QosShowPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // QoS Queue parameters.
    Queue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue

    // QoS EA Shaper parameters.
    Shape PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape

    // QoS WFQ parameters.
    Wfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq

    // QoS Policer parameters.
    Police PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police

    // QoS Mark parameters.
    Marking PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowPclassSt.EntityData.YFilter = qosShowPclassSt.YFilter
    qosShowPclassSt.EntityData.YangName = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowPclassSt.EntityData.SegmentPath = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowPclassSt.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowPclassSt.Queue}
    qosShowPclassSt.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowPclassSt.Shape}
    qosShowPclassSt.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowPclassSt.Wfq}
    qosShowPclassSt.EntityData.Children["police"] = types.YChild{"Police", &qosShowPclassSt.Police}
    qosShowPclassSt.EntityData.Children["marking"] = types.YChild{"Marking", &qosShowPclassSt.Marking}
    qosShowPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowPclassSt.ClassLevel}
    qosShowPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowPclassSt.ClassName}
    return &(qosShowPclassSt.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-pclass-st"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-pclass-st"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-pclass-st"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["committed-weight"] = types.YChild{"CommittedWeight", &wfq.CommittedWeight}
    wfq.EntityData.Children["programmed-wfq"] = types.YChild{"ProgrammedWfq", &wfq.ProgrammedWfq}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetEntityData() *types.CommonEntityData {
    committedWeight.EntityData.YFilter = committedWeight.YFilter
    committedWeight.EntityData.YangName = "committed-weight"
    committedWeight.EntityData.BundleName = "cisco_ios_xr"
    committedWeight.EntityData.ParentYangName = "wfq"
    committedWeight.EntityData.SegmentPath = "committed-weight"
    committedWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedWeight.EntityData.Children = make(map[string]types.YChild)
    committedWeight.EntityData.Leafs = make(map[string]types.YLeaf)
    committedWeight.EntityData.Leafs["value"] = types.YLeaf{"Value", committedWeight.Value}
    committedWeight.EntityData.Leafs["unit"] = types.YLeaf{"Unit", committedWeight.Unit}
    return &(committedWeight.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetEntityData() *types.CommonEntityData {
    programmedWfq.EntityData.YFilter = programmedWfq.YFilter
    programmedWfq.EntityData.YangName = "programmed-wfq"
    programmedWfq.EntityData.BundleName = "cisco_ios_xr"
    programmedWfq.EntityData.ParentYangName = "wfq"
    programmedWfq.EntityData.SegmentPath = "programmed-wfq"
    programmedWfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedWfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedWfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedWfq.EntityData.Children = make(map[string]types.YChild)
    programmedWfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &programmedWfq.Bandwidth}
    programmedWfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &programmedWfq.SumOfBandwidth}
    programmedWfq.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedWfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", programmedWfq.ExcessRatio}
    return &(programmedWfq.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "programmed-wfq"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "programmed-wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policer ID. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-pclass-st"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", police.PolicerId}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetEntityData() *types.CommonEntityData {
    marking.EntityData.YFilter = marking.YFilter
    marking.EntityData.YangName = "marking"
    marking.EntityData.BundleName = "cisco_ios_xr"
    marking.EntityData.ParentYangName = "qos-show-pclass-st"
    marking.EntityData.SegmentPath = "marking"
    marking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    marking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    marking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    marking.EntityData.Children = make(map[string]types.YChild)
    marking.EntityData.Children["mark-only"] = types.YChild{"MarkOnly", &marking.MarkOnly}
    marking.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &marking.PoliceConform}
    marking.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &marking.PoliceExceed}
    marking.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(marking.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetEntityData() *types.CommonEntityData {
    markOnly.EntityData.YFilter = markOnly.YFilter
    markOnly.EntityData.YangName = "mark-only"
    markOnly.EntityData.BundleName = "cisco_ios_xr"
    markOnly.EntityData.ParentYangName = "marking"
    markOnly.EntityData.SegmentPath = "mark-only"
    markOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markOnly.EntityData.Children = make(map[string]types.YChild)
    markOnly.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range markOnly.MarkDetail {
        markOnly.EntityData.Children[types.GetSegmentPath(&markOnly.MarkDetail[i])] = types.YChild{"MarkDetail", &markOnly.MarkDetail[i]}
    }
    markOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    markOnly.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", markOnly.ActionType}
    return &(markOnly.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "mark-only"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
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

    // Maximum classes per policy. The type is interface{} with range:
    // 0..4294967295.
    MaxClassesPerPolicy interface{}

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
    capability.EntityData.Leafs["max-classes-per-policy"] = types.YLeaf{"MaxClassesPerPolicy", capability.MaxClassesPerPolicy}
    capability.EntityData.Leafs["max-police-actions-per-class"] = types.YLeaf{"MaxPoliceActionsPerClass", capability.MaxPoliceActionsPerClass}
    capability.EntityData.Leafs["max-marking-actions-per-class"] = types.YLeaf{"MaxMarkingActionsPerClass", capability.MaxMarkingActionsPerClass}
    capability.EntityData.Leafs["max-matches-per-class"] = types.YLeaf{"MaxMatchesPerClass", capability.MaxMatchesPerClass}
    capability.EntityData.Leafs["max-classmap-name-length"] = types.YLeaf{"MaxClassmapNameLength", capability.MaxClassmapNameLength}
    capability.EntityData.Leafs["max-bundle-members"] = types.YLeaf{"MaxBundleMembers", capability.MaxBundleMembers}
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

    // QoS policy direction egress.
    Output PlatformQos_Nodes_Node_Interfaces_Interface_Output

    // QoS policy direction ingress.
    Input PlatformQos_Nodes_Node_Interfaces_Interface_Input
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
    self.EntityData.Children["output"] = types.YChild{"Output", &self.Output}
    self.EntityData.Children["input"] = types.YChild{"Input", &self.Input}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output
// QoS policy direction egress
type PlatformQos_Nodes_Node_Interfaces_Interface_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass
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
    output.EntityData.Children["header"] = types.YChild{"Header", &output.Header}
    output.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &output.InterfaceParameters}
    output.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &output.SkywarpQosPolicyClass}
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(output.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "output"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "output"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["interface-config-rate"] = types.YChild{"InterfaceConfigRate", &interfaceParameters.InterfaceConfigRate}
    interfaceParameters.EntityData.Children["interface-program-rate"] = types.YChild{"InterfaceProgramRate", &interfaceParameters.InterfaceProgramRate}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetEntityData() *types.CommonEntityData {
    interfaceConfigRate.EntityData.YFilter = interfaceConfigRate.YFilter
    interfaceConfigRate.EntityData.YangName = "interface-config-rate"
    interfaceConfigRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigRate.EntityData.ParentYangName = "interface-parameters"
    interfaceConfigRate.EntityData.SegmentPath = "interface-config-rate"
    interfaceConfigRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigRate.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfigRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceConfigRate.Value}
    interfaceConfigRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceConfigRate.Unit}
    return &(interfaceConfigRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetEntityData() *types.CommonEntityData {
    interfaceProgramRate.EntityData.YFilter = interfaceProgramRate.YFilter
    interfaceProgramRate.EntityData.YangName = "interface-program-rate"
    interfaceProgramRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceProgramRate.EntityData.ParentYangName = "interface-parameters"
    interfaceProgramRate.EntityData.SegmentPath = "interface-program-rate"
    interfaceProgramRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProgramRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProgramRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProgramRate.EntityData.Children = make(map[string]types.YChild)
    interfaceProgramRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceProgramRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceProgramRate.Value}
    interfaceProgramRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceProgramRate.Unit}
    return &(interfaceProgramRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "output"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-pclass-st"] = types.YChild{"QosShowPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowPclassSt[i])] = types.YChild{"QosShowPclassSt", &skywarpQosPolicyClass.QosShowPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // QoS Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue

    // QoS EA Shaper parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape

    // QoS WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq

    // QoS Policer parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police

    // QoS Mark parameters.
    Marking PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowPclassSt.EntityData.YFilter = qosShowPclassSt.YFilter
    qosShowPclassSt.EntityData.YangName = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowPclassSt.EntityData.SegmentPath = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowPclassSt.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowPclassSt.Queue}
    qosShowPclassSt.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowPclassSt.Shape}
    qosShowPclassSt.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowPclassSt.Wfq}
    qosShowPclassSt.EntityData.Children["police"] = types.YChild{"Police", &qosShowPclassSt.Police}
    qosShowPclassSt.EntityData.Children["marking"] = types.YChild{"Marking", &qosShowPclassSt.Marking}
    qosShowPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowPclassSt.ClassLevel}
    qosShowPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowPclassSt.ClassName}
    return &(qosShowPclassSt.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-pclass-st"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-pclass-st"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-pclass-st"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["committed-weight"] = types.YChild{"CommittedWeight", &wfq.CommittedWeight}
    wfq.EntityData.Children["programmed-wfq"] = types.YChild{"ProgrammedWfq", &wfq.ProgrammedWfq}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetEntityData() *types.CommonEntityData {
    committedWeight.EntityData.YFilter = committedWeight.YFilter
    committedWeight.EntityData.YangName = "committed-weight"
    committedWeight.EntityData.BundleName = "cisco_ios_xr"
    committedWeight.EntityData.ParentYangName = "wfq"
    committedWeight.EntityData.SegmentPath = "committed-weight"
    committedWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedWeight.EntityData.Children = make(map[string]types.YChild)
    committedWeight.EntityData.Leafs = make(map[string]types.YLeaf)
    committedWeight.EntityData.Leafs["value"] = types.YLeaf{"Value", committedWeight.Value}
    committedWeight.EntityData.Leafs["unit"] = types.YLeaf{"Unit", committedWeight.Unit}
    return &(committedWeight.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetEntityData() *types.CommonEntityData {
    programmedWfq.EntityData.YFilter = programmedWfq.YFilter
    programmedWfq.EntityData.YangName = "programmed-wfq"
    programmedWfq.EntityData.BundleName = "cisco_ios_xr"
    programmedWfq.EntityData.ParentYangName = "wfq"
    programmedWfq.EntityData.SegmentPath = "programmed-wfq"
    programmedWfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedWfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedWfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedWfq.EntityData.Children = make(map[string]types.YChild)
    programmedWfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &programmedWfq.Bandwidth}
    programmedWfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &programmedWfq.SumOfBandwidth}
    programmedWfq.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedWfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", programmedWfq.ExcessRatio}
    return &(programmedWfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "programmed-wfq"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "programmed-wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policer ID. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-pclass-st"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", police.PolicerId}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetEntityData() *types.CommonEntityData {
    marking.EntityData.YFilter = marking.YFilter
    marking.EntityData.YangName = "marking"
    marking.EntityData.BundleName = "cisco_ios_xr"
    marking.EntityData.ParentYangName = "qos-show-pclass-st"
    marking.EntityData.SegmentPath = "marking"
    marking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    marking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    marking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    marking.EntityData.Children = make(map[string]types.YChild)
    marking.EntityData.Children["mark-only"] = types.YChild{"MarkOnly", &marking.MarkOnly}
    marking.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &marking.PoliceConform}
    marking.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &marking.PoliceExceed}
    marking.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(marking.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetEntityData() *types.CommonEntityData {
    markOnly.EntityData.YFilter = markOnly.YFilter
    markOnly.EntityData.YangName = "mark-only"
    markOnly.EntityData.BundleName = "cisco_ios_xr"
    markOnly.EntityData.ParentYangName = "marking"
    markOnly.EntityData.SegmentPath = "mark-only"
    markOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markOnly.EntityData.Children = make(map[string]types.YChild)
    markOnly.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range markOnly.MarkDetail {
        markOnly.EntityData.Children[types.GetSegmentPath(&markOnly.MarkDetail[i])] = types.YChild{"MarkDetail", &markOnly.MarkDetail[i]}
    }
    markOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    markOnly.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", markOnly.ActionType}
    return &(markOnly.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "mark-only"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input
// QoS policy direction ingress
type PlatformQos_Nodes_Node_Interfaces_Interface_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass
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
    input.EntityData.Children["header"] = types.YChild{"Header", &input.Header}
    input.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &input.InterfaceParameters}
    input.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &input.SkywarpQosPolicyClass}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "input"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "input"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Children["interface-config-rate"] = types.YChild{"InterfaceConfigRate", &interfaceParameters.InterfaceConfigRate}
    interfaceParameters.EntityData.Children["interface-program-rate"] = types.YChild{"InterfaceProgramRate", &interfaceParameters.InterfaceProgramRate}
    interfaceParameters.EntityData.Children["port-shaper-rate"] = types.YChild{"PortShaperRate", &interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceParameters.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetEntityData() *types.CommonEntityData {
    interfaceConfigRate.EntityData.YFilter = interfaceConfigRate.YFilter
    interfaceConfigRate.EntityData.YangName = "interface-config-rate"
    interfaceConfigRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceConfigRate.EntityData.ParentYangName = "interface-parameters"
    interfaceConfigRate.EntityData.SegmentPath = "interface-config-rate"
    interfaceConfigRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConfigRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConfigRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConfigRate.EntityData.Children = make(map[string]types.YChild)
    interfaceConfigRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConfigRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceConfigRate.Value}
    interfaceConfigRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceConfigRate.Unit}
    return &(interfaceConfigRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetEntityData() *types.CommonEntityData {
    interfaceProgramRate.EntityData.YFilter = interfaceProgramRate.YFilter
    interfaceProgramRate.EntityData.YangName = "interface-program-rate"
    interfaceProgramRate.EntityData.BundleName = "cisco_ios_xr"
    interfaceProgramRate.EntityData.ParentYangName = "interface-parameters"
    interfaceProgramRate.EntityData.SegmentPath = "interface-program-rate"
    interfaceProgramRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProgramRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProgramRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProgramRate.EntityData.Children = make(map[string]types.YChild)
    interfaceProgramRate.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceProgramRate.EntityData.Leafs["value"] = types.YLeaf{"Value", interfaceProgramRate.Value}
    interfaceProgramRate.EntityData.Leafs["unit"] = types.YLeaf{"Unit", interfaceProgramRate.Unit}
    return &(interfaceProgramRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "input"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-pclass-st"] = types.YChild{"QosShowPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowPclassSt[i])] = types.YChild{"QosShowPclassSt", &skywarpQosPolicyClass.QosShowPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // QoS Queue parameters.
    Queue PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue

    // QoS EA Shaper parameters.
    Shape PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape

    // QoS WFQ parameters.
    Wfq PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq

    // QoS Policer parameters.
    Police PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police

    // QoS Mark parameters.
    Marking PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowPclassSt.EntityData.YFilter = qosShowPclassSt.YFilter
    qosShowPclassSt.EntityData.YangName = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowPclassSt.EntityData.SegmentPath = "qos-show-pclass-st"
    qosShowPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowPclassSt.EntityData.Children["queue"] = types.YChild{"Queue", &qosShowPclassSt.Queue}
    qosShowPclassSt.EntityData.Children["shape"] = types.YChild{"Shape", &qosShowPclassSt.Shape}
    qosShowPclassSt.EntityData.Children["wfq"] = types.YChild{"Wfq", &qosShowPclassSt.Wfq}
    qosShowPclassSt.EntityData.Children["police"] = types.YChild{"Police", &qosShowPclassSt.Police}
    qosShowPclassSt.EntityData.Children["marking"] = types.YChild{"Marking", &qosShowPclassSt.Marking}
    qosShowPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowPclassSt.ClassLevel}
    qosShowPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowPclassSt.ClassName}
    return &(qosShowPclassSt.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "qos-show-pclass-st"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", queue.QueueType}
    return &(queue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "qos-show-pclass-st"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "qos-show-pclass-st"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["committed-weight"] = types.YChild{"CommittedWeight", &wfq.CommittedWeight}
    wfq.EntityData.Children["programmed-wfq"] = types.YChild{"ProgrammedWfq", &wfq.ProgrammedWfq}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-weight"] = types.YLeaf{"ExcessWeight", wfq.ExcessWeight}
    return &(wfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetEntityData() *types.CommonEntityData {
    committedWeight.EntityData.YFilter = committedWeight.YFilter
    committedWeight.EntityData.YangName = "committed-weight"
    committedWeight.EntityData.BundleName = "cisco_ios_xr"
    committedWeight.EntityData.ParentYangName = "wfq"
    committedWeight.EntityData.SegmentPath = "committed-weight"
    committedWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    committedWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    committedWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    committedWeight.EntityData.Children = make(map[string]types.YChild)
    committedWeight.EntityData.Leafs = make(map[string]types.YLeaf)
    committedWeight.EntityData.Leafs["value"] = types.YLeaf{"Value", committedWeight.Value}
    committedWeight.EntityData.Leafs["unit"] = types.YLeaf{"Unit", committedWeight.Unit}
    return &(committedWeight.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetEntityData() *types.CommonEntityData {
    programmedWfq.EntityData.YFilter = programmedWfq.YFilter
    programmedWfq.EntityData.YangName = "programmed-wfq"
    programmedWfq.EntityData.BundleName = "cisco_ios_xr"
    programmedWfq.EntityData.ParentYangName = "wfq"
    programmedWfq.EntityData.SegmentPath = "programmed-wfq"
    programmedWfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    programmedWfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    programmedWfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    programmedWfq.EntityData.Children = make(map[string]types.YChild)
    programmedWfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &programmedWfq.Bandwidth}
    programmedWfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &programmedWfq.SumOfBandwidth}
    programmedWfq.EntityData.Leafs = make(map[string]types.YLeaf)
    programmedWfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", programmedWfq.ExcessRatio}
    return &(programmedWfq.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "programmed-wfq"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "programmed-wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // policer ID. The type is interface{} with range: 0..4294967295.
    PolicerId interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir

    // CBS.
    Cbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "qos-show-pclass-st"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["policer-id"] = types.YLeaf{"PolicerId", police.PolicerId}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetEntityData() *types.CommonEntityData {
    marking.EntityData.YFilter = marking.YFilter
    marking.EntityData.YangName = "marking"
    marking.EntityData.BundleName = "cisco_ios_xr"
    marking.EntityData.ParentYangName = "qos-show-pclass-st"
    marking.EntityData.SegmentPath = "marking"
    marking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    marking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    marking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    marking.EntityData.Children = make(map[string]types.YChild)
    marking.EntityData.Children["mark-only"] = types.YChild{"MarkOnly", &marking.MarkOnly}
    marking.EntityData.Children["police-conform"] = types.YChild{"PoliceConform", &marking.PoliceConform}
    marking.EntityData.Children["police-exceed"] = types.YChild{"PoliceExceed", &marking.PoliceExceed}
    marking.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(marking.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetEntityData() *types.CommonEntityData {
    markOnly.EntityData.YFilter = markOnly.YFilter
    markOnly.EntityData.YangName = "mark-only"
    markOnly.EntityData.BundleName = "cisco_ios_xr"
    markOnly.EntityData.ParentYangName = "marking"
    markOnly.EntityData.SegmentPath = "mark-only"
    markOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    markOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    markOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    markOnly.EntityData.Children = make(map[string]types.YChild)
    markOnly.EntityData.Children["mark-detail"] = types.YChild{"MarkDetail", nil}
    for i := range markOnly.MarkDetail {
        markOnly.EntityData.Children[types.GetSegmentPath(&markOnly.MarkDetail[i])] = types.YChild{"MarkDetail", &markOnly.MarkDetail[i]}
    }
    markOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    markOnly.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", markOnly.ActionType}
    return &(markOnly.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetEntityData() *types.CommonEntityData {
    markDetail.EntityData.YFilter = markDetail.YFilter
    markDetail.EntityData.YangName = "mark-detail"
    markDetail.EntityData.BundleName = "cisco_ios_xr"
    markDetail.EntityData.ParentYangName = "mark-only"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetEntityData() *types.CommonEntityData {
    policeConform.EntityData.YFilter = policeConform.YFilter
    policeConform.EntityData.YangName = "police-conform"
    policeConform.EntityData.BundleName = "cisco_ios_xr"
    policeConform.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetEntityData() *types.CommonEntityData {
    policeExceed.EntityData.YFilter = policeExceed.YFilter
    policeExceed.EntityData.YangName = "police-exceed"
    policeExceed.EntityData.BundleName = "cisco_ios_xr"
    policeExceed.EntityData.ParentYangName = "marking"
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

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa
// platform qos ea
type PlatformQosEa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes with platform specific QoS-EA configuration.
    Nodes PlatformQosEa_Nodes
}

func (platformQosEa *PlatformQosEa) GetEntityData() *types.CommonEntityData {
    platformQosEa.EntityData.YFilter = platformQosEa.YFilter
    platformQosEa.EntityData.YangName = "platform-qos-ea"
    platformQosEa.EntityData.BundleName = "cisco_ios_xr"
    platformQosEa.EntityData.ParentYangName = "Cisco-IOS-XR-skp-qos-oper"
    platformQosEa.EntityData.SegmentPath = "Cisco-IOS-XR-skp-qos-oper:platform-qos-ea"
    platformQosEa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformQosEa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformQosEa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformQosEa.EntityData.Children = make(map[string]types.YChild)
    platformQosEa.EntityData.Children["nodes"] = types.YChild{"Nodes", &platformQosEa.Nodes}
    platformQosEa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformQosEa.EntityData)
}

// PlatformQosEa_Nodes
// List of nodes with platform specific QoS-EA
// configuration
type PlatformQosEa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node with platform specific QoS-EA configuration. The type is slice of
    // PlatformQosEa_Nodes_Node.
    Node []PlatformQosEa_Nodes_Node
}

func (nodes *PlatformQosEa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "platform-qos-ea"
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

// PlatformQosEa_Nodes_Node
// Node with platform specific QoS-EA
// configuration
type PlatformQosEa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // QoS-EA list of bundle interfaces.
    BundleInterfaces PlatformQosEa_Nodes_Node_BundleInterfaces

    // QoS-EA list of interfaces.
    Interfaces PlatformQosEa_Nodes_Node_Interfaces
}

func (node *PlatformQosEa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["bundle-interfaces"] = types.YChild{"BundleInterfaces", &node.BundleInterfaces}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces
// QoS-EA list of bundle interfaces
type PlatformQosEa_Nodes_Node_BundleInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS-EA list of member interfaces.
    MemberInterfaces PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetEntityData() *types.CommonEntityData {
    bundleInterface.EntityData.YFilter = bundleInterface.YFilter
    bundleInterface.EntityData.YangName = "bundle-interface"
    bundleInterface.EntityData.BundleName = "cisco_ios_xr"
    bundleInterface.EntityData.ParentYangName = "bundle-interfaces"
    bundleInterface.EntityData.SegmentPath = "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
    bundleInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterface.EntityData.Children = make(map[string]types.YChild)
    bundleInterface.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleInterface.MemberInterfaces}
    bundleInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterface.InterfaceName}
    return &(bundleInterface.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
// QoS-EA list of member interfaces
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetEntityData() *types.CommonEntityData {
    memberInterfaces.EntityData.YFilter = memberInterfaces.YFilter
    memberInterfaces.EntityData.YangName = "member-interfaces"
    memberInterfaces.EntityData.BundleName = "cisco_ios_xr"
    memberInterfaces.EntityData.ParentYangName = "bundle-interface"
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS-EA policy direction output.
    BundleOutput PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput

    // QoS-EA policy direction input.
    BundleInput PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetEntityData() *types.CommonEntityData {
    memberInterface.EntityData.YFilter = memberInterface.YFilter
    memberInterface.EntityData.YangName = "member-interface"
    memberInterface.EntityData.BundleName = "cisco_ios_xr"
    memberInterface.EntityData.ParentYangName = "member-interfaces"
    memberInterface.EntityData.SegmentPath = "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
    memberInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterface.EntityData.Children = make(map[string]types.YChild)
    memberInterface.EntityData.Children["bundle-output"] = types.YChild{"BundleOutput", &memberInterface.BundleOutput}
    memberInterface.EntityData.Children["bundle-input"] = types.YChild{"BundleInput", &memberInterface.BundleInput}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
// QoS-EA policy direction output
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetEntityData() *types.CommonEntityData {
    bundleOutput.EntityData.YFilter = bundleOutput.YFilter
    bundleOutput.EntityData.YangName = "bundle-output"
    bundleOutput.EntityData.BundleName = "cisco_ios_xr"
    bundleOutput.EntityData.ParentYangName = "member-interface"
    bundleOutput.EntityData.SegmentPath = "bundle-output"
    bundleOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleOutput.EntityData.Children = make(map[string]types.YChild)
    bundleOutput.EntityData.Children["details"] = types.YChild{"Details", &bundleOutput.Details}
    bundleOutput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleOutput.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "bundle-output"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &details.InterfaceParameters}
    details.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &details.SkywarpQosPolicyClass}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Max Hierarchial Depth. The type is interface{} with range: 0..255.
    HierarchicalDepth interface{}

    // Interface Type. The type is string with length: 0..101.
    InterfaceType interface{}

    // Interface Programmed Rate. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRate interface{}

    // Port Shaper Rate. The type is interface{} with range: 0..4294967295.
    PortShaperRate interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UnderLineInterfaceHandle interface{}

    // Bundle Interface ID. The type is interface{} with range: 0..65535.
    BundleId interface{}

    // UIDB Index. The type is interface{} with range: 0..65535.
    UidbIndex interface{}

    // QoS Interface handle. The type is interface{} with range:
    // 0..18446744073709551615.
    QosInterfaceHandle interface{}

    // Local Port. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // Policy Map ID. The type is interface{} with range: 0..65535.
    PolicyMapId interface{}
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "details"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceParameters.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", interfaceParameters.PolicyName}
    interfaceParameters.EntityData.Leafs["hierarchical-depth"] = types.YLeaf{"HierarchicalDepth", interfaceParameters.HierarchicalDepth}
    interfaceParameters.EntityData.Leafs["interface-type"] = types.YLeaf{"InterfaceType", interfaceParameters.InterfaceType}
    interfaceParameters.EntityData.Leafs["interface-rate"] = types.YLeaf{"InterfaceRate", interfaceParameters.InterfaceRate}
    interfaceParameters.EntityData.Leafs["port-shaper-rate"] = types.YLeaf{"PortShaperRate", interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceParameters.InterfaceHandle}
    interfaceParameters.EntityData.Leafs["under-line-interface-handle"] = types.YLeaf{"UnderLineInterfaceHandle", interfaceParameters.UnderLineInterfaceHandle}
    interfaceParameters.EntityData.Leafs["bundle-id"] = types.YLeaf{"BundleId", interfaceParameters.BundleId}
    interfaceParameters.EntityData.Leafs["uidb-index"] = types.YLeaf{"UidbIndex", interfaceParameters.UidbIndex}
    interfaceParameters.EntityData.Leafs["qos-interface-handle"] = types.YLeaf{"QosInterfaceHandle", interfaceParameters.QosInterfaceHandle}
    interfaceParameters.EntityData.Leafs["port"] = types.YLeaf{"Port", interfaceParameters.Port}
    interfaceParameters.EntityData.Leafs["policy-map-id"] = types.YLeaf{"PolicyMapId", interfaceParameters.PolicyMapId}
    return &(interfaceParameters.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "details"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-ea-pclass-st"] = types.YChild{"QosShowEaPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowEaPclassSt[i])] = types.YChild{"QosShowEaPclassSt", &skywarpQosPolicyClass.QosShowEaPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Node Flags. The type is string with length: 0..101.
    NodeFlags interface{}

    // Statistical Flags. The type is string with length: 0..101.
    StatsFlags interface{}

    // QoS EA Class Configuration.
    Config PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config

    // QoS EA Class Result.
    Result PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowEaPclassSt.EntityData.YFilter = qosShowEaPclassSt.YFilter
    qosShowEaPclassSt.EntityData.YangName = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowEaPclassSt.EntityData.SegmentPath = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowEaPclassSt.EntityData.Children["config"] = types.YChild{"Config", &qosShowEaPclassSt.Config}
    qosShowEaPclassSt.EntityData.Children["result"] = types.YChild{"Result", &qosShowEaPclassSt.Result}
    qosShowEaPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaPclassSt.EntityData.Leafs["index"] = types.YLeaf{"Index", qosShowEaPclassSt.Index}
    qosShowEaPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaPclassSt.ClassLevel}
    qosShowEaPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaPclassSt.ClassName}
    qosShowEaPclassSt.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaPclassSt.PolicyName}
    qosShowEaPclassSt.EntityData.Leafs["node-flags"] = types.YLeaf{"NodeFlags", qosShowEaPclassSt.NodeFlags}
    qosShowEaPclassSt.EntityData.Leafs["stats-flags"] = types.YLeaf{"StatsFlags", qosShowEaPclassSt.StatsFlags}
    return &(qosShowEaPclassSt.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Config. The type is string with length: 0..101.
    NodeConfig interface{}

    // QoS EA Policer parameters.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police

    // QoS EA Shaper parameters.
    Shape PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape

    // QoS EA WFQ parameters.
    Wfq PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["police"] = types.YChild{"Police", &config.Police}
    config.EntityData.Children["shape"] = types.YChild{"Shape", &config.Shape}
    config.EntityData.Children["wfq"] = types.YChild{"Wfq", &config.Wfq}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["node-config"] = types.YLeaf{"NodeConfig", config.NodeConfig}
    return &(config.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color Aware. The type is bool.
    ColorAware interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir

    // CBS.
    Cbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "config"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["color-aware"] = types.YLeaf{"ColorAware", police.ColorAware}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "config"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "config"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &wfq.SumOfBandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    return &(wfq.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xr"
    result.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    result.EntityData.Children = make(map[string]types.YChild)
    result.EntityData.Children["queue"] = types.YChild{"Queue", &result.Queue}
    result.EntityData.Children["police"] = types.YChild{"Police", &result.Police}
    result.EntityData.Leafs = make(map[string]types.YLeaf)
    result.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", result.StatsId}
    return &(result.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Commit Tx. The type is interface{} with range: 0..4294967295.
    CommitTx interface{}

    // Excess Tx. The type is interface{} with range: 0..4294967295.
    ExcessTx interface{}

    // Drop. The type is interface{} with range: 0..4294967295.
    Drop interface{}
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "result"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["commit-tx"] = types.YLeaf{"CommitTx", queue.CommitTx}
    queue.EntityData.Leafs["excess-tx"] = types.YLeaf{"ExcessTx", queue.ExcessTx}
    queue.EntityData.Leafs["drop"] = types.YLeaf{"Drop", queue.Drop}
    return &(queue.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Token Bucket ID. The type is interface{} with range: 0..4294967295.
    TokenBucketId interface{}

    // Conform Rate. The type is interface{} with range: 0..4294967295.
    Conform interface{}

    // Exceed Rate. The type is interface{} with range: 0..4294967295.
    Exceed interface{}

    // Violate Rate. The type is interface{} with range: 0..4294967295.
    Violate interface{}
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "result"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["token-bucket-id"] = types.YLeaf{"TokenBucketId", police.TokenBucketId}
    police.EntityData.Leafs["conform"] = types.YLeaf{"Conform", police.Conform}
    police.EntityData.Leafs["exceed"] = types.YLeaf{"Exceed", police.Exceed}
    police.EntityData.Leafs["violate"] = types.YLeaf{"Violate", police.Violate}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
// QoS-EA policy direction input
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetEntityData() *types.CommonEntityData {
    bundleInput.EntityData.YFilter = bundleInput.YFilter
    bundleInput.EntityData.YangName = "bundle-input"
    bundleInput.EntityData.BundleName = "cisco_ios_xr"
    bundleInput.EntityData.ParentYangName = "member-interface"
    bundleInput.EntityData.SegmentPath = "bundle-input"
    bundleInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInput.EntityData.Children = make(map[string]types.YChild)
    bundleInput.EntityData.Children["details"] = types.YChild{"Details", &bundleInput.Details}
    bundleInput.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInput.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "bundle-input"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["header"] = types.YChild{"Header", &details.Header}
    details.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &details.InterfaceParameters}
    details.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &details.SkywarpQosPolicyClass}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Max Hierarchial Depth. The type is interface{} with range: 0..255.
    HierarchicalDepth interface{}

    // Interface Type. The type is string with length: 0..101.
    InterfaceType interface{}

    // Interface Programmed Rate. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRate interface{}

    // Port Shaper Rate. The type is interface{} with range: 0..4294967295.
    PortShaperRate interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UnderLineInterfaceHandle interface{}

    // Bundle Interface ID. The type is interface{} with range: 0..65535.
    BundleId interface{}

    // UIDB Index. The type is interface{} with range: 0..65535.
    UidbIndex interface{}

    // QoS Interface handle. The type is interface{} with range:
    // 0..18446744073709551615.
    QosInterfaceHandle interface{}

    // Local Port. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // Policy Map ID. The type is interface{} with range: 0..65535.
    PolicyMapId interface{}
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "details"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceParameters.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", interfaceParameters.PolicyName}
    interfaceParameters.EntityData.Leafs["hierarchical-depth"] = types.YLeaf{"HierarchicalDepth", interfaceParameters.HierarchicalDepth}
    interfaceParameters.EntityData.Leafs["interface-type"] = types.YLeaf{"InterfaceType", interfaceParameters.InterfaceType}
    interfaceParameters.EntityData.Leafs["interface-rate"] = types.YLeaf{"InterfaceRate", interfaceParameters.InterfaceRate}
    interfaceParameters.EntityData.Leafs["port-shaper-rate"] = types.YLeaf{"PortShaperRate", interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceParameters.InterfaceHandle}
    interfaceParameters.EntityData.Leafs["under-line-interface-handle"] = types.YLeaf{"UnderLineInterfaceHandle", interfaceParameters.UnderLineInterfaceHandle}
    interfaceParameters.EntityData.Leafs["bundle-id"] = types.YLeaf{"BundleId", interfaceParameters.BundleId}
    interfaceParameters.EntityData.Leafs["uidb-index"] = types.YLeaf{"UidbIndex", interfaceParameters.UidbIndex}
    interfaceParameters.EntityData.Leafs["qos-interface-handle"] = types.YLeaf{"QosInterfaceHandle", interfaceParameters.QosInterfaceHandle}
    interfaceParameters.EntityData.Leafs["port"] = types.YLeaf{"Port", interfaceParameters.Port}
    interfaceParameters.EntityData.Leafs["policy-map-id"] = types.YLeaf{"PolicyMapId", interfaceParameters.PolicyMapId}
    return &(interfaceParameters.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "details"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-ea-pclass-st"] = types.YChild{"QosShowEaPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowEaPclassSt[i])] = types.YChild{"QosShowEaPclassSt", &skywarpQosPolicyClass.QosShowEaPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Node Flags. The type is string with length: 0..101.
    NodeFlags interface{}

    // Statistical Flags. The type is string with length: 0..101.
    StatsFlags interface{}

    // QoS EA Class Configuration.
    Config PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config

    // QoS EA Class Result.
    Result PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowEaPclassSt.EntityData.YFilter = qosShowEaPclassSt.YFilter
    qosShowEaPclassSt.EntityData.YangName = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowEaPclassSt.EntityData.SegmentPath = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowEaPclassSt.EntityData.Children["config"] = types.YChild{"Config", &qosShowEaPclassSt.Config}
    qosShowEaPclassSt.EntityData.Children["result"] = types.YChild{"Result", &qosShowEaPclassSt.Result}
    qosShowEaPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaPclassSt.EntityData.Leafs["index"] = types.YLeaf{"Index", qosShowEaPclassSt.Index}
    qosShowEaPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaPclassSt.ClassLevel}
    qosShowEaPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaPclassSt.ClassName}
    qosShowEaPclassSt.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaPclassSt.PolicyName}
    qosShowEaPclassSt.EntityData.Leafs["node-flags"] = types.YLeaf{"NodeFlags", qosShowEaPclassSt.NodeFlags}
    qosShowEaPclassSt.EntityData.Leafs["stats-flags"] = types.YLeaf{"StatsFlags", qosShowEaPclassSt.StatsFlags}
    return &(qosShowEaPclassSt.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Config. The type is string with length: 0..101.
    NodeConfig interface{}

    // QoS EA Policer parameters.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police

    // QoS EA Shaper parameters.
    Shape PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape

    // QoS EA WFQ parameters.
    Wfq PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["police"] = types.YChild{"Police", &config.Police}
    config.EntityData.Children["shape"] = types.YChild{"Shape", &config.Shape}
    config.EntityData.Children["wfq"] = types.YChild{"Wfq", &config.Wfq}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["node-config"] = types.YLeaf{"NodeConfig", config.NodeConfig}
    return &(config.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color Aware. The type is bool.
    ColorAware interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir

    // CBS.
    Cbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "config"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["color-aware"] = types.YLeaf{"ColorAware", police.ColorAware}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "config"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "config"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &wfq.SumOfBandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    return &(wfq.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xr"
    result.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    result.EntityData.Children = make(map[string]types.YChild)
    result.EntityData.Children["queue"] = types.YChild{"Queue", &result.Queue}
    result.EntityData.Children["police"] = types.YChild{"Police", &result.Police}
    result.EntityData.Leafs = make(map[string]types.YLeaf)
    result.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", result.StatsId}
    return &(result.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Commit Tx. The type is interface{} with range: 0..4294967295.
    CommitTx interface{}

    // Excess Tx. The type is interface{} with range: 0..4294967295.
    ExcessTx interface{}

    // Drop. The type is interface{} with range: 0..4294967295.
    Drop interface{}
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "result"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["commit-tx"] = types.YLeaf{"CommitTx", queue.CommitTx}
    queue.EntityData.Leafs["excess-tx"] = types.YLeaf{"ExcessTx", queue.ExcessTx}
    queue.EntityData.Leafs["drop"] = types.YLeaf{"Drop", queue.Drop}
    return &(queue.EntityData)
}

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Token Bucket ID. The type is interface{} with range: 0..4294967295.
    TokenBucketId interface{}

    // Conform Rate. The type is interface{} with range: 0..4294967295.
    Conform interface{}

    // Exceed Rate. The type is interface{} with range: 0..4294967295.
    Exceed interface{}

    // Violate Rate. The type is interface{} with range: 0..4294967295.
    Violate interface{}
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "result"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["token-bucket-id"] = types.YLeaf{"TokenBucketId", police.TokenBucketId}
    police.EntityData.Leafs["conform"] = types.YLeaf{"Conform", police.Conform}
    police.EntityData.Leafs["exceed"] = types.YLeaf{"Exceed", police.Exceed}
    police.EntityData.Leafs["violate"] = types.YLeaf{"Violate", police.Violate}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces
// QoS-EA list of interfaces
type PlatformQosEa_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface_.
    Interface_ []PlatformQosEa_Nodes_Node_Interfaces_Interface
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // QoS-EA policy direction egress.
    Output PlatformQosEa_Nodes_Node_Interfaces_Interface_Output

    // QoS-EA policy direction ingress.
    Input PlatformQosEa_Nodes_Node_Interfaces_Interface_Input
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["output"] = types.YChild{"Output", &self.Output}
    self.EntityData.Children["input"] = types.YChild{"Input", &self.Input}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output
// QoS-EA policy direction egress
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetEntityData() *types.CommonEntityData {
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
    details.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &details.InterfaceParameters}
    details.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &details.SkywarpQosPolicyClass}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Max Hierarchial Depth. The type is interface{} with range: 0..255.
    HierarchicalDepth interface{}

    // Interface Type. The type is string with length: 0..101.
    InterfaceType interface{}

    // Interface Programmed Rate. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRate interface{}

    // Port Shaper Rate. The type is interface{} with range: 0..4294967295.
    PortShaperRate interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UnderLineInterfaceHandle interface{}

    // Bundle Interface ID. The type is interface{} with range: 0..65535.
    BundleId interface{}

    // UIDB Index. The type is interface{} with range: 0..65535.
    UidbIndex interface{}

    // QoS Interface handle. The type is interface{} with range:
    // 0..18446744073709551615.
    QosInterfaceHandle interface{}

    // Local Port. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // Policy Map ID. The type is interface{} with range: 0..65535.
    PolicyMapId interface{}
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "details"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceParameters.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", interfaceParameters.PolicyName}
    interfaceParameters.EntityData.Leafs["hierarchical-depth"] = types.YLeaf{"HierarchicalDepth", interfaceParameters.HierarchicalDepth}
    interfaceParameters.EntityData.Leafs["interface-type"] = types.YLeaf{"InterfaceType", interfaceParameters.InterfaceType}
    interfaceParameters.EntityData.Leafs["interface-rate"] = types.YLeaf{"InterfaceRate", interfaceParameters.InterfaceRate}
    interfaceParameters.EntityData.Leafs["port-shaper-rate"] = types.YLeaf{"PortShaperRate", interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceParameters.InterfaceHandle}
    interfaceParameters.EntityData.Leafs["under-line-interface-handle"] = types.YLeaf{"UnderLineInterfaceHandle", interfaceParameters.UnderLineInterfaceHandle}
    interfaceParameters.EntityData.Leafs["bundle-id"] = types.YLeaf{"BundleId", interfaceParameters.BundleId}
    interfaceParameters.EntityData.Leafs["uidb-index"] = types.YLeaf{"UidbIndex", interfaceParameters.UidbIndex}
    interfaceParameters.EntityData.Leafs["qos-interface-handle"] = types.YLeaf{"QosInterfaceHandle", interfaceParameters.QosInterfaceHandle}
    interfaceParameters.EntityData.Leafs["port"] = types.YLeaf{"Port", interfaceParameters.Port}
    interfaceParameters.EntityData.Leafs["policy-map-id"] = types.YLeaf{"PolicyMapId", interfaceParameters.PolicyMapId}
    return &(interfaceParameters.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "details"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-ea-pclass-st"] = types.YChild{"QosShowEaPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowEaPclassSt[i])] = types.YChild{"QosShowEaPclassSt", &skywarpQosPolicyClass.QosShowEaPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Node Flags. The type is string with length: 0..101.
    NodeFlags interface{}

    // Statistical Flags. The type is string with length: 0..101.
    StatsFlags interface{}

    // QoS EA Class Configuration.
    Config PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config

    // QoS EA Class Result.
    Result PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowEaPclassSt.EntityData.YFilter = qosShowEaPclassSt.YFilter
    qosShowEaPclassSt.EntityData.YangName = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowEaPclassSt.EntityData.SegmentPath = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowEaPclassSt.EntityData.Children["config"] = types.YChild{"Config", &qosShowEaPclassSt.Config}
    qosShowEaPclassSt.EntityData.Children["result"] = types.YChild{"Result", &qosShowEaPclassSt.Result}
    qosShowEaPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaPclassSt.EntityData.Leafs["index"] = types.YLeaf{"Index", qosShowEaPclassSt.Index}
    qosShowEaPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaPclassSt.ClassLevel}
    qosShowEaPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaPclassSt.ClassName}
    qosShowEaPclassSt.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaPclassSt.PolicyName}
    qosShowEaPclassSt.EntityData.Leafs["node-flags"] = types.YLeaf{"NodeFlags", qosShowEaPclassSt.NodeFlags}
    qosShowEaPclassSt.EntityData.Leafs["stats-flags"] = types.YLeaf{"StatsFlags", qosShowEaPclassSt.StatsFlags}
    return &(qosShowEaPclassSt.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Config. The type is string with length: 0..101.
    NodeConfig interface{}

    // QoS EA Policer parameters.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police

    // QoS EA Shaper parameters.
    Shape PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape

    // QoS EA WFQ parameters.
    Wfq PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["police"] = types.YChild{"Police", &config.Police}
    config.EntityData.Children["shape"] = types.YChild{"Shape", &config.Shape}
    config.EntityData.Children["wfq"] = types.YChild{"Wfq", &config.Wfq}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["node-config"] = types.YLeaf{"NodeConfig", config.NodeConfig}
    return &(config.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color Aware. The type is bool.
    ColorAware interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir

    // CBS.
    Cbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "config"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["color-aware"] = types.YLeaf{"ColorAware", police.ColorAware}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "config"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "config"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &wfq.SumOfBandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    return &(wfq.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xr"
    result.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    result.EntityData.Children = make(map[string]types.YChild)
    result.EntityData.Children["queue"] = types.YChild{"Queue", &result.Queue}
    result.EntityData.Children["police"] = types.YChild{"Police", &result.Police}
    result.EntityData.Leafs = make(map[string]types.YLeaf)
    result.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", result.StatsId}
    return &(result.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Commit Tx. The type is interface{} with range: 0..4294967295.
    CommitTx interface{}

    // Excess Tx. The type is interface{} with range: 0..4294967295.
    ExcessTx interface{}

    // Drop. The type is interface{} with range: 0..4294967295.
    Drop interface{}
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "result"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["commit-tx"] = types.YLeaf{"CommitTx", queue.CommitTx}
    queue.EntityData.Leafs["excess-tx"] = types.YLeaf{"ExcessTx", queue.ExcessTx}
    queue.EntityData.Leafs["drop"] = types.YLeaf{"Drop", queue.Drop}
    return &(queue.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Token Bucket ID. The type is interface{} with range: 0..4294967295.
    TokenBucketId interface{}

    // Conform Rate. The type is interface{} with range: 0..4294967295.
    Conform interface{}

    // Exceed Rate. The type is interface{} with range: 0..4294967295.
    Exceed interface{}

    // Violate Rate. The type is interface{} with range: 0..4294967295.
    Violate interface{}
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "result"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["token-bucket-id"] = types.YLeaf{"TokenBucketId", police.TokenBucketId}
    police.EntityData.Leafs["conform"] = types.YLeaf{"Conform", police.Conform}
    police.EntityData.Leafs["exceed"] = types.YLeaf{"Exceed", police.Exceed}
    police.EntityData.Leafs["violate"] = types.YLeaf{"Violate", police.Violate}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input
// QoS-EA policy direction ingress
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetEntityData() *types.CommonEntityData {
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
    details.EntityData.Children["interface-parameters"] = types.YChild{"InterfaceParameters", &details.InterfaceParameters}
    details.EntityData.Children["skywarp-qos-policy-class"] = types.YChild{"SkywarpQosPolicyClass", &details.SkywarpQosPolicyClass}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(details.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with length: 0..101.
    InterfaceName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Direction. The type is string with length: 0..11.
    Direction interface{}

    // Number of classes. The type is interface{} with range: 0..65535.
    Classes interface{}
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetEntityData() *types.CommonEntityData {
    header.EntityData.YFilter = header.YFilter
    header.EntityData.YangName = "header"
    header.EntityData.BundleName = "cisco_ios_xr"
    header.EntityData.ParentYangName = "details"
    header.EntityData.SegmentPath = "header"
    header.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    header.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    header.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    header.EntityData.Children = make(map[string]types.YChild)
    header.EntityData.Leafs = make(map[string]types.YLeaf)
    header.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", header.InterfaceName}
    header.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", header.PolicyName}
    header.EntityData.Leafs["direction"] = types.YLeaf{"Direction", header.Direction}
    header.EntityData.Leafs["classes"] = types.YLeaf{"Classes", header.Classes}
    return &(header.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Max Hierarchial Depth. The type is interface{} with range: 0..255.
    HierarchicalDepth interface{}

    // Interface Type. The type is string with length: 0..101.
    InterfaceType interface{}

    // Interface Programmed Rate. The type is interface{} with range:
    // 0..4294967295.
    InterfaceRate interface{}

    // Port Shaper Rate. The type is interface{} with range: 0..4294967295.
    PortShaperRate interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UnderLineInterfaceHandle interface{}

    // Bundle Interface ID. The type is interface{} with range: 0..65535.
    BundleId interface{}

    // UIDB Index. The type is interface{} with range: 0..65535.
    UidbIndex interface{}

    // QoS Interface handle. The type is interface{} with range:
    // 0..18446744073709551615.
    QosInterfaceHandle interface{}

    // Local Port. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // Policy Map ID. The type is interface{} with range: 0..65535.
    PolicyMapId interface{}
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetEntityData() *types.CommonEntityData {
    interfaceParameters.EntityData.YFilter = interfaceParameters.YFilter
    interfaceParameters.EntityData.YangName = "interface-parameters"
    interfaceParameters.EntityData.BundleName = "cisco_ios_xr"
    interfaceParameters.EntityData.ParentYangName = "details"
    interfaceParameters.EntityData.SegmentPath = "interface-parameters"
    interfaceParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceParameters.EntityData.Children = make(map[string]types.YChild)
    interfaceParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceParameters.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", interfaceParameters.PolicyName}
    interfaceParameters.EntityData.Leafs["hierarchical-depth"] = types.YLeaf{"HierarchicalDepth", interfaceParameters.HierarchicalDepth}
    interfaceParameters.EntityData.Leafs["interface-type"] = types.YLeaf{"InterfaceType", interfaceParameters.InterfaceType}
    interfaceParameters.EntityData.Leafs["interface-rate"] = types.YLeaf{"InterfaceRate", interfaceParameters.InterfaceRate}
    interfaceParameters.EntityData.Leafs["port-shaper-rate"] = types.YLeaf{"PortShaperRate", interfaceParameters.PortShaperRate}
    interfaceParameters.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", interfaceParameters.InterfaceHandle}
    interfaceParameters.EntityData.Leafs["under-line-interface-handle"] = types.YLeaf{"UnderLineInterfaceHandle", interfaceParameters.UnderLineInterfaceHandle}
    interfaceParameters.EntityData.Leafs["bundle-id"] = types.YLeaf{"BundleId", interfaceParameters.BundleId}
    interfaceParameters.EntityData.Leafs["uidb-index"] = types.YLeaf{"UidbIndex", interfaceParameters.UidbIndex}
    interfaceParameters.EntityData.Leafs["qos-interface-handle"] = types.YLeaf{"QosInterfaceHandle", interfaceParameters.QosInterfaceHandle}
    interfaceParameters.EntityData.Leafs["port"] = types.YLeaf{"Port", interfaceParameters.Port}
    interfaceParameters.EntityData.Leafs["policy-map-id"] = types.YLeaf{"PolicyMapId", interfaceParameters.PolicyMapId}
    return &(interfaceParameters.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetEntityData() *types.CommonEntityData {
    skywarpQosPolicyClass.EntityData.YFilter = skywarpQosPolicyClass.YFilter
    skywarpQosPolicyClass.EntityData.YangName = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.BundleName = "cisco_ios_xr"
    skywarpQosPolicyClass.EntityData.ParentYangName = "details"
    skywarpQosPolicyClass.EntityData.SegmentPath = "skywarp-qos-policy-class"
    skywarpQosPolicyClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    skywarpQosPolicyClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    skywarpQosPolicyClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    skywarpQosPolicyClass.EntityData.Children = make(map[string]types.YChild)
    skywarpQosPolicyClass.EntityData.Children["qos-show-ea-pclass-st"] = types.YChild{"QosShowEaPclassSt", nil}
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        skywarpQosPolicyClass.EntityData.Children[types.GetSegmentPath(&skywarpQosPolicyClass.QosShowEaPclassSt[i])] = types.YChild{"QosShowEaPclassSt", &skywarpQosPolicyClass.QosShowEaPclassSt[i]}
    }
    skywarpQosPolicyClass.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(skywarpQosPolicyClass.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Class level. The type is interface{} with range: 0..255.
    ClassLevel interface{}

    // Class name. The type is string with length: 0..65.
    ClassName interface{}

    // Policy name. The type is string with length: 0..65.
    PolicyName interface{}

    // Node Flags. The type is string with length: 0..101.
    NodeFlags interface{}

    // Statistical Flags. The type is string with length: 0..101.
    StatsFlags interface{}

    // QoS EA Class Configuration.
    Config PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config

    // QoS EA Class Result.
    Result PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetEntityData() *types.CommonEntityData {
    qosShowEaPclassSt.EntityData.YFilter = qosShowEaPclassSt.YFilter
    qosShowEaPclassSt.EntityData.YangName = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.BundleName = "cisco_ios_xr"
    qosShowEaPclassSt.EntityData.ParentYangName = "skywarp-qos-policy-class"
    qosShowEaPclassSt.EntityData.SegmentPath = "qos-show-ea-pclass-st"
    qosShowEaPclassSt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qosShowEaPclassSt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qosShowEaPclassSt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qosShowEaPclassSt.EntityData.Children = make(map[string]types.YChild)
    qosShowEaPclassSt.EntityData.Children["config"] = types.YChild{"Config", &qosShowEaPclassSt.Config}
    qosShowEaPclassSt.EntityData.Children["result"] = types.YChild{"Result", &qosShowEaPclassSt.Result}
    qosShowEaPclassSt.EntityData.Leafs = make(map[string]types.YLeaf)
    qosShowEaPclassSt.EntityData.Leafs["index"] = types.YLeaf{"Index", qosShowEaPclassSt.Index}
    qosShowEaPclassSt.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", qosShowEaPclassSt.ClassLevel}
    qosShowEaPclassSt.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", qosShowEaPclassSt.ClassName}
    qosShowEaPclassSt.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", qosShowEaPclassSt.PolicyName}
    qosShowEaPclassSt.EntityData.Leafs["node-flags"] = types.YLeaf{"NodeFlags", qosShowEaPclassSt.NodeFlags}
    qosShowEaPclassSt.EntityData.Leafs["stats-flags"] = types.YLeaf{"StatsFlags", qosShowEaPclassSt.StatsFlags}
    return &(qosShowEaPclassSt.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Config. The type is string with length: 0..101.
    NodeConfig interface{}

    // QoS EA Policer parameters.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police

    // QoS EA Shaper parameters.
    Shape PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape

    // QoS EA WFQ parameters.
    Wfq PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Children["police"] = types.YChild{"Police", &config.Police}
    config.EntityData.Children["shape"] = types.YChild{"Shape", &config.Shape}
    config.EntityData.Children["wfq"] = types.YChild{"Wfq", &config.Wfq}
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["node-config"] = types.YLeaf{"NodeConfig", config.NodeConfig}
    return &(config.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color Aware. The type is bool.
    ColorAware interface{}

    // Policer type. The type is TbAlgorithm.
    PolicerType interface{}

    // CIR.
    Cir PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir

    // CBS.
    Cbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "config"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Children["cir"] = types.YChild{"Cir", &police.Cir}
    police.EntityData.Children["cbs"] = types.YChild{"Cbs", &police.Cbs}
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["color-aware"] = types.YLeaf{"ColorAware", police.ColorAware}
    police.EntityData.Leafs["policer-type"] = types.YLeaf{"PolicerType", police.PolicerType}
    return &(police.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "config"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = make(map[string]types.YChild)
    shape.EntityData.Children["pir"] = types.YChild{"Pir", &shape.Pir}
    shape.EntityData.Children["pbs"] = types.YChild{"Pbs", &shape.Pbs}
    shape.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(shape.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetEntityData() *types.CommonEntityData {
    wfq.EntityData.YFilter = wfq.YFilter
    wfq.EntityData.YangName = "wfq"
    wfq.EntityData.BundleName = "cisco_ios_xr"
    wfq.EntityData.ParentYangName = "config"
    wfq.EntityData.SegmentPath = "wfq"
    wfq.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wfq.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wfq.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wfq.EntityData.Children = make(map[string]types.YChild)
    wfq.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &wfq.Bandwidth}
    wfq.EntityData.Children["sum-of-bandwidth"] = types.YChild{"SumOfBandwidth", &wfq.SumOfBandwidth}
    wfq.EntityData.Leafs = make(map[string]types.YLeaf)
    wfq.EntityData.Leafs["excess-ratio"] = types.YLeaf{"ExcessRatio", wfq.ExcessRatio}
    return &(wfq.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetEntityData() *types.CommonEntityData {
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

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetEntityData() *types.CommonEntityData {
    sumOfBandwidth.EntityData.YFilter = sumOfBandwidth.YFilter
    sumOfBandwidth.EntityData.YangName = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.BundleName = "cisco_ios_xr"
    sumOfBandwidth.EntityData.ParentYangName = "wfq"
    sumOfBandwidth.EntityData.SegmentPath = "sum-of-bandwidth"
    sumOfBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sumOfBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sumOfBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sumOfBandwidth.EntityData.Children = make(map[string]types.YChild)
    sumOfBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    sumOfBandwidth.EntityData.Leafs["value"] = types.YLeaf{"Value", sumOfBandwidth.Value}
    sumOfBandwidth.EntityData.Leafs["unit"] = types.YLeaf{"Unit", sumOfBandwidth.Unit}
    return &(sumOfBandwidth.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetEntityData() *types.CommonEntityData {
    result.EntityData.YFilter = result.YFilter
    result.EntityData.YangName = "result"
    result.EntityData.BundleName = "cisco_ios_xr"
    result.EntityData.ParentYangName = "qos-show-ea-pclass-st"
    result.EntityData.SegmentPath = "result"
    result.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    result.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    result.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    result.EntityData.Children = make(map[string]types.YChild)
    result.EntityData.Children["queue"] = types.YChild{"Queue", &result.Queue}
    result.EntityData.Children["police"] = types.YChild{"Police", &result.Police}
    result.EntityData.Leafs = make(map[string]types.YLeaf)
    result.EntityData.Leafs["stats-id"] = types.YLeaf{"StatsId", result.StatsId}
    return &(result.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Commit Tx. The type is interface{} with range: 0..4294967295.
    CommitTx interface{}

    // Excess Tx. The type is interface{} with range: 0..4294967295.
    ExcessTx interface{}

    // Drop. The type is interface{} with range: 0..4294967295.
    Drop interface{}
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "result"
    queue.EntityData.SegmentPath = "queue"
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = make(map[string]types.YChild)
    queue.EntityData.Leafs = make(map[string]types.YLeaf)
    queue.EntityData.Leafs["queue-id"] = types.YLeaf{"QueueId", queue.QueueId}
    queue.EntityData.Leafs["commit-tx"] = types.YLeaf{"CommitTx", queue.CommitTx}
    queue.EntityData.Leafs["excess-tx"] = types.YLeaf{"ExcessTx", queue.ExcessTx}
    queue.EntityData.Leafs["drop"] = types.YLeaf{"Drop", queue.Drop}
    return &(queue.EntityData)
}

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Token Bucket ID. The type is interface{} with range: 0..4294967295.
    TokenBucketId interface{}

    // Conform Rate. The type is interface{} with range: 0..4294967295.
    Conform interface{}

    // Exceed Rate. The type is interface{} with range: 0..4294967295.
    Exceed interface{}

    // Violate Rate. The type is interface{} with range: 0..4294967295.
    Violate interface{}
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "result"
    police.EntityData.SegmentPath = "police"
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = make(map[string]types.YChild)
    police.EntityData.Leafs = make(map[string]types.YLeaf)
    police.EntityData.Leafs["token-bucket-id"] = types.YLeaf{"TokenBucketId", police.TokenBucketId}
    police.EntityData.Leafs["conform"] = types.YLeaf{"Conform", police.Conform}
    police.EntityData.Leafs["exceed"] = types.YLeaf{"Exceed", police.Exceed}
    police.EntityData.Leafs["violate"] = types.YLeaf{"Violate", police.Violate}
    return &(police.EntityData)
}

